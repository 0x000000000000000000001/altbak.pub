#!/usr/bin/env python3
"""Rust-only typed-AST diagnostics: counting-allocator runs and CPU samples.

alloc:
    Run the `--cfg diag_alloc` binary with DIAG_ALLOC=1 and save the raw
    per-phase JSON plus a summary of allocation and free totals. Keep these
    numbers separate from the timed campaign: the counters perturb timings.

sample:
    Run the timed binary with the requested phase(s) and collect a macOS
    `sample` CPU profile while it works. DIAG_SAMPLES extends the run so the
    sampled interval covers steady-state passes. The raw sample output and a
    top-of-stack summary are written to the output directory.

Both subcommands require an explicit corpus file in the driver JSON format and
never touch the twelve-module fixture or its oracle.
"""
import argparse
import json
import os
import re
import select
import statistics
import subprocess
import sys
import time
from pathlib import Path

ROOT = Path(__file__).resolve().parents[3]
SAMPLE = Path('/usr/bin/sample')


def run_label(value):
    return re.sub(r'[^A-Za-z0-9_.-]+', '_', value)


def driver_env(corpus, phases=None, warmups=None, samples=None, alloc=False):
    env = {k: v for k, v in os.environ.items() if not k.startswith('DIAG_')}
    env['DIAG_CORPUS'] = str(corpus)
    if phases:
        env['DIAG_PHASES'] = phases
    if warmups is not None:
        env['DIAG_WARMUPS'] = str(warmups)
    if samples is not None:
        env['DIAG_SAMPLES'] = str(samples)
    if alloc:
        env['DIAG_ALLOC'] = '1'
    return env


def parse_report(stdout):
    for line in reversed(stdout.strip().splitlines()):
        line = line.strip()
        if line.startswith('{') and line.endswith('}'):
            return json.loads(line)
    raise ValueError('driver produced no JSON report')


def summarize_alloc(report):
    summary = {'backend': report['backend'], 'allocator': report.get('allocator'),
               'alloc_diag': report.get('alloc_diag'), 'modules': report['modules'],
               'timed_cases': report['timed_cases'], 'phases': {}}
    for phase, data in report['phases'].items():
        samples = [s for s in data['samples'] if 'allocated_bytes' in s]
        if not samples:
            continue
        summary['phases'][phase] = {
            'samples': len(samples),
            'allocated_bytes': [s['allocated_bytes'] for s in samples],
            'allocations': [s['allocations'] for s in samples],
            'deallocated_bytes': [s['deallocated_bytes'] for s in samples],
            'deallocations': [s['deallocations'] for s in samples],
            'drop_allocated_bytes': [s.get('drop_allocated_bytes') for s in samples],
            'drop_deallocated_bytes': [s.get('drop_deallocated_bytes') for s in samples],
            'median_allocated_bytes': statistics.median(s['allocated_bytes'] for s in samples),
            'median_allocations': statistics.median(s['allocations'] for s in samples),
            'median_drop_deallocated_bytes': statistics.median(
                s.get('drop_deallocated_bytes', 0) for s in samples),
            'min_time_us': data['time_us'],
        }
    return summary


def command_alloc(args):
    out = args.out.resolve()
    out.mkdir(parents=True, exist_ok=True)
    phases = args.phases
    env = driver_env(args.corpus, phases=phases, warmups=args.warmups, samples=args.samples, alloc=True)
    raw = subprocess.run([str(args.binary)], env=env, cwd=args.cwd, capture_output=True, text=True)
    if raw.returncode:
        raise RuntimeError(f'alloc run failed ({raw.returncode}): {raw.stderr[-4000:]}')
    report = parse_report(raw.stdout)
    label = run_label(phases or 'parse,decode,combined')
    (out / f'alloc-{label}.json').write_text(raw.stdout)
    summary = summarize_alloc(report)
    (out / f'alloc-{label}-summary.json').write_text(json.dumps(summary, indent=2) + '\n')
    print(json.dumps({phase: {key: value for key, value in data.items()
                              if key.startswith('median') or key == 'min_time_us'}
                      for phase, data in summary['phases'].items()}, indent=2))
    return 0


THREAD_RE = re.compile(r'^\s*(\d+)\s+(Thread_\d+)\s*(.*)$')
FRAME_RE = re.compile(r'^\s*([+!:|\s]*?)\s*(\d+)\s+(.*?)\s+\(in\s+([^)]+)\)(.*)$')
OFFSET_RE = re.compile(r'\s+\+\s+\d+\s+\[[0-9a-fA-Fx]+\]\s*$')
DROP_MARKERS = ('drop_in_place', 'drop_slow', 'mi_free', 'mi_free_aligned')


def parse_call_graph(text):
    """Per-thread frames with depth and exclusive sample attribution.

    macOS `sample` prints inclusive counts per frame; exclusive samples are
    count minus the sum of immediate children. Only the `Call graph:` section
    is parsed (the trailing `Sort by top of stack` section reports cumulative
    counts that are not sample counts).
    """
    threads = []
    current = None
    in_graph = False
    for line in text.splitlines():
        if line.startswith('Call graph:'):
            in_graph = True
            continue
        if not in_graph:
            continue
        if line.startswith('Sort by top') or line.startswith('Binary Images'):
            break
        header = THREAD_RE.match(line)
        if header:
            current = {'id': header.group(2), 'name': header.group(3).strip(), 'frames': []}
            threads.append(current)
            continue
        frame = FRAME_RE.match(line)
        if frame is None or current is None:
            continue
        prefix, count, symbol, binary = frame.group(1), int(frame.group(2)), frame.group(3), frame.group(4)
        symbol = OFFSET_RE.sub('', symbol).strip()
        # macOS sample indents the inclusive count by two characters per depth
        # after the fixed four-space prefix and one marker column.
        current['frames'].append({'depth': max(0, (frame.start(3) - 6) // 2), 'count': count,
                                  'symbol': symbol, 'binary': binary})
    return threads


def exclusive_frames(thread):
    """Unused by the current summary; kept only as documentation of why raw
    inclusive counts cannot be differenced naively (collapsed branches can
    repeat the same symbol at several depths)."""
    pending = []
    done = []
    for frame in thread['frames']:
        while pending and pending[-1]['depth'] >= frame['depth']:
            child = pending.pop()
            child['exclusive'] = child['count'] - child['child_sum']
            done.append(child)
        frame['child_sum'] = 0
        if pending:
            pending[-1]['child_sum'] += frame['count']
        pending.append(frame)
    while pending:
        child = pending.pop()
        child['exclusive'] = child['count'] - child['child_sum']
        done.append(child)
    return done


SORT_RE = re.compile(r'^\s+(.*?)\s+\(in\s+([^)]+)\)\s+(\d+)\s*$')
IDLE_SYMBOLS = {'__psynch_cvwait', '__ulock_wait', '__workq_kernreturn', 'start_wqthread',
                'mach_msg2_trap', 'semaphore_wait_trap', 'guarded_pwrite', 'select',
                '__select', 'pthread_mutex_lock'}
FINGERPRINT_MARKERS = ('fingerprint', 'canonical_hash', 'stringify', 'sha256')


def self_time_entries(text):
    """`Sort by top of stack, same collapsed` entries: leaf/self samples.

    Counts are summed across all sampled threads; idle worker threads parked in
    kernel waits are separated by the caller, and only entries with at least
    five samples are printed by `sample`."""
    section = text.split('Sort by top of stack', 1)
    if len(section) == 1:
        return []
    entries = []
    for line in section[1].splitlines():
        match = SORT_RE.match(line)
        if match:
            entries.append({'symbol': match.group(1).strip(), 'binary': match.group(2).strip(),
                            'samples': int(match.group(3))})
    return entries


def sample_summary(path):
    text = path.read_text(errors='replace')
    threads = parse_call_graph(text)
    graph = {}
    working = next((thread for thread in threads if 'purust-main' in thread['name']), None)
    if working is not None and working['frames']:
        frames = working['frames']
        graph['thread_samples'] = frames[0]['count']
        graph['profile_batch_inclusive'] = max((f['count'] for f in frames
                                                if 'profile_batch' in f['symbol']), default=0)
        graph['driver_closure_inclusive'] = max((f['count'] for f in frames
                                                 if 'Test_JsonTypedAst_drive' in f['symbol']), default=0)
        graph['fingerprint_symbols_in_thread'] = sorted({f['symbol'] for f in frames
                                                         if any(marker in f['symbol'].lower()
                                                                for marker in FINGERPRINT_MARKERS)})
    entries = self_time_entries(text)
    work = [entry for entry in entries if entry['binary'] == 'purust_output']
    system = [entry for entry in entries if entry['binary'] != 'purust_output']
    idle = [entry for entry in system if entry['symbol'] in IDLE_SYMBOLS]
    on_behalf = [entry for entry in system if entry['symbol'] not in IDLE_SYMBOLS]
    destruction = [entry for entry in work
                   if any(marker in entry['symbol'] for marker in DROP_MARKERS)]
    compute = [entry for entry in work if entry not in destruction]
    return {
        'threads': [{'id': thread['id'], 'name': thread['name'],
                     'samples': thread['frames'][0]['count'] if thread['frames'] else 0}
                    for thread in threads],
        'graph': graph,
        'working_thread_self_samples': sum(entry['samples'] for entry in work),
        'compute': sorted(compute, key=lambda entry: -entry['samples'])[:25],
        'destruction': sorted(destruction, key=lambda entry: -entry['samples'])[:15],
        'system_on_behalf': sorted(on_behalf, key=lambda entry: -entry['samples'])[:10],
        'idle_threads': sorted(idle, key=lambda entry: -entry['samples'])[:5],
        'fingerprint_symbols': sorted({entry['symbol'] for entry in entries
                                       if any(marker in entry['symbol'].lower()
                                              for marker in FINGERPRINT_MARKERS)}),
        'self_time_note': 'leaf samples from `sample` (>=5), idle worker threads separated',
    }


def command_sample(args):
    out = args.out.resolve()
    out.mkdir(parents=True, exist_ok=True)
    label = run_label(args.phase)
    json_path = out / f'sample-{label}.json'
    log_path = out / f'sample-{label}.driver.log'
    sample_path = out / f'sample-{label}.txt'
    env = driver_env(args.corpus)
    env['DIAG_PROFILE_PHASE'] = args.phase
    env['DIAG_PROFILE_PASSES'] = str(args.passes)
    read_fd, write_fd = os.pipe()
    process = None
    try:
        with json_path.open('w') as stdout, log_path.open('w') as stderr:
            env['DIAG_READY_FD'] = str(write_fd)
            process = subprocess.Popen([str(args.binary)], env=env, cwd=args.cwd,
                                       stdout=stdout, stderr=stderr, pass_fds=(write_fd,))
        os.close(write_fd)
        write_fd = -1
        ready, _, _ = select.select([read_fd], [], [], args.ready_timeout)
        if not ready:
            raise RuntimeError(f'driver did not signal readiness in {args.ready_timeout}s; see {log_path}')
        os.read(read_fd, 1)
        started = time.time()
        result = subprocess.run([str(SAMPLE), str(process.pid), str(args.seconds),
                                 '-file', str(sample_path), '-mayDie'],
                                capture_output=True, text=True)
        print(f'sample finished in {time.time()-started:.1f}s: {result.stderr.strip()[-200:]}', flush=True)
        process.terminate()
        try:
            process.wait(timeout=30)
        except subprocess.TimeoutExpired:
            process.kill()
            process.wait()
    finally:
        if write_fd != -1:
            os.close(write_fd)
        os.close(read_fd)
        if process is not None and process.poll() is None:
            process.terminate()
            process.wait()
    driver = None
    if json_path.stat().st_size:
        try:
            driver = json.loads(json_path.read_text())
        except ValueError:
            driver = None
    summary = {'phase': args.phase, 'seconds': args.seconds, 'passes': args.passes,
               'profile_mode': True, 'killed_after_sample': True,
               'fingerprinting_excluded': 'validation before readiness and after the loop; kernel has none',
               'driver_json': driver}
    if sample_path.exists():
        summary['sample'] = sample_summary(sample_path)
    (out / f'sample-{label}-summary.json').write_text(json.dumps(summary, indent=2) + '\n')
    sample = summary.get('sample', {})
    print(json.dumps({'phase': args.phase, 'passes': args.passes,
                      'working_thread_self_samples': sample.get('working_thread_self_samples'),
                      'compute_top': [{'samples': item['samples'], 'symbol': item['symbol']}
                                      for item in sample.get('compute', [])[:10]],
                      'destruction_top': [{'samples': item['samples'], 'symbol': item['symbol']}
                                          for item in sample.get('destruction', [])[:5]],
                      'fingerprint_symbols': sample.get('fingerprint_symbols', [])}, indent=2))
    return 0


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    sub = parser.add_subparsers(dest='command', required=True)

    alloc = sub.add_parser('alloc', help='counting-allocator run (separate build)')
    alloc.add_argument('--binary', type=Path, required=True)
    alloc.add_argument('--corpus', type=Path, required=True)
    alloc.add_argument('--out', type=Path, required=True)
    alloc.add_argument('--cwd', type=Path)
    alloc.add_argument('--phases', default='parse,decode,combined')
    alloc.add_argument('--warmups', type=int, default=0)
    alloc.add_argument('--samples', type=int, default=3)
    alloc.set_defaults(func=command_alloc)

    sample = sub.add_parser('sample', help='macOS CPU sample of one phase')
    sample.add_argument('--binary', type=Path, required=True)
    sample.add_argument('--corpus', type=Path, required=True)
    sample.add_argument('--out', type=Path, required=True)
    sample.add_argument('--cwd', type=Path)
    sample.add_argument('--phase', default='combined')
    sample.add_argument('--seconds', type=int, default=20)
    sample.add_argument('--passes', type=int, default=1000000,
                        help='profiling kernel passes before natural exit; the helper kills first')
    sample.add_argument('--ready-timeout', type=float, default=300.0)
    sample.set_defaults(func=command_sample)

    args = parser.parse_args()
    if getattr(args, 'cwd', None) is None:
        args.cwd = ROOT
    return args.func(args)


if __name__ == '__main__':
    sys.exit(main())

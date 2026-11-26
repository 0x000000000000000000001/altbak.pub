#!/usr/bin/env python3
"""Build, validate and archive the missing C library references (three processes)."""
import datetime
import hashlib
import json
import os
from pathlib import Path
import platform
import statistics
import subprocess
import sys

ROOT = Path(__file__).resolve().parents[2]
WORK = ROOT / 'var/benchmark/c-library-20261126'
ARCHIVE = ROOT / 'docs/benchmark-results/2026-11-26-c-library.json'
SOURCE = ROOT / 'tmp/c-library'


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def run(command, log):
    result = subprocess.run(command, cwd=ROOT, text=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
    log.write_text(result.stdout)
    if result.returncode:
        raise RuntimeError(f'{command}: exit {result.returncode}; see {log}')
    return result.stdout


def main():
    if ARCHIVE.exists():
        raise SystemExit('Archive already exists; use a new campaign path instead of overwriting measurements')
    for directory in ('build', 'validation', 'runs', 'assembly'):
        (WORK / directory).mkdir(parents=True, exist_ok=True)
    baseline = (ROOT / 'README.md').read_bytes()
    (WORK / 'README.before.md').write_bytes(baseline)
    # Guard original workloads, harnesses, historical C source and archived baselines.
    protected = [ROOT / 'tmp/bench_native.c', ROOT / 'bin/benchmark/library-cases.json',
                 ROOT / 'docs/benchmark-results/2026-11-26-go-native-callbacks.json']
    protected += [path for path in (ROOT / 'src').rglob('*') if path.is_file()]
    guards = {str(p.relative_to(ROOT)): sha(p) for p in protected}
    names = ('benchmark.c', 'arrays.c', 'strings.c', 'records.c')
    sources = [str((SOURCE / name).relative_to(ROOT)) for name in names]
    flags = ['clang', '-std=c11', '-O3', '-Wall', '-Wextra', '-Werror']
    binary = WORK / 'build/benchmark-c-library'
    build = flags + sources + ['-o', str(binary)]
    run(build, WORK / 'build/compiler.log')
    run([str(binary), '--verify'], WORK / 'validation/official-oracles.log')
    sanitized = WORK / 'build/benchmark-c-library-sanitized'
    run(flags + ['-fsanitize=address,undefined'] + sources + ['-o', str(sanitized)], WORK / 'build/sanitizers.log')
    run([str(sanitized), '--verify'], WORK / 'validation/official-oracles-sanitized.log')
    for kernel, test in [('arrays.c', 'verify-arrays.c'), ('records.c', 'test-records.c')]:
        target = WORK / 'build' / ('verify-' + kernel[:-2])
        run(flags + ['-fsanitize=address,undefined', str(SOURCE / kernel), str(SOURCE / test), '-o', str(target)],
            WORK / 'build' / (kernel + '.log'))
        run([str(target)], WORK / 'validation' / (kernel + '.log'))
    run(['node', str(SOURCE / 'test-strings.mjs')], WORK / 'validation/strings.log')
    for name in names:
        run(flags + ['-S', str(SOURCE / name), '-o', str(WORK / 'assembly' / (name + '.s'))],
            WORK / 'build' / (name + '-assembly.log'))
    before = {str(p.relative_to(ROOT)): sha(p) for p in SOURCE.iterdir() if p.is_file()}
    binary_sha = sha(binary)
    definitions = json.loads((ROOT / 'bin/benchmark/library-cases.json').read_text())
    expected = {case['label']: case['expected'] for case in definitions}
    started = datetime.datetime.now(datetime.timezone.utc).isoformat()
    load_start = os.getloadavg()
    runs = []
    for process in range(1, 4):
        print(f'Measuring C library process {process}/3', flush=True)
        raw = run([str(binary)], WORK / 'runs' / f'{process}.jsonl')
        rows = [json.loads(line) for line in raw.splitlines()]
        assert [row['case'] for row in rows] == list(expected)
        for row in rows:
            assert row['value'] == expected[row['case']]
            assert len(row['batches_ns']) == 10 and min(row['batches_ns']) > 0
            exact = min(row['batches_ns']) / 1000 / row['iterations']
            assert abs(exact - row['minimum_us']) < 1e-10
            row['minimum_us'] = exact  # Keep the integer-clock calculation, not printed rounding.
        runs.append(rows)
        assert sha(binary) == binary_sha
    medians = {name: statistics.median(next(row['minimum_us'] for row in run if row['case'] == name)
                                      for run in runs) for name in expected}
    assert guards == {name: sha(ROOT / name) for name in guards}
    assert before == {str(p.relative_to(ROOT)): sha(p) for p in SOURCE.iterdir() if p.is_file()}
    assert (ROOT / 'README.md').read_bytes() == baseline
    files = list(SOURCE.iterdir()) + list((WORK / 'validation').glob('*')) + list((WORK / 'assembly').glob('*'))
    result = {
        'title': '23 additional optimized imperative C library references',
        'started_at_utc': started, 'completed_at_utc': datetime.datetime.now(datetime.timezone.utc).isoformat(),
        'host': {'system': platform.platform(), 'machine': platform.machine(), 'load_start': load_start,
                 'load_end': os.getloadavg(), 'isolated': False},
        'toolchain': subprocess.check_output(['clang', '--version'], text=True),
        'build_command': build, 'binary_sha256': binary_sha,
        'protocol': {'processes': 3, 'global_warmups': 3, 'local_warmups': 3, 'batches': 10,
                     'batch_target_ms': 10, 'maximum_iterations': 16777216,
                     'aggregate': 'median of three process minima', 'lto': False,
                     'preparation': 'outside timed action, matching PureScript prepare',
                     'input_opacity': 'noinline address + memory compiler barrier every action; separate translation units',
                     'consumption': 'seq_cst atomic store per result; validate final value after each batch'},
        'readme_before_sha256': hashlib.sha256(baseline).hexdigest(),
        'protected_sources_sha256': guards, 'library_cases': definitions,
        'processes': runs, 'medians_us': medians,
        'proofs': {str(p.relative_to(ROOT)): {'sha256': sha(p), 'text': p.read_text()}
                   for p in files if p.is_file()},
        'historical_c_core_remeasured': False, 'go_or_js_remeasured': False,
    }
    ARCHIVE.write_text(json.dumps(result, indent=2, ensure_ascii=False, allow_nan=False) + '\n')
    print(json.dumps(medians, indent=2), flush=True)
    print(f'Archived {ARCHIVE.relative_to(ROOT)}', flush=True)


if __name__ == '__main__':
    main()

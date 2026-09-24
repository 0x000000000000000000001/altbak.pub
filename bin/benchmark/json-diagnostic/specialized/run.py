#!/usr/bin/env python3
"""Build and run the specialised-decoder A/B audit.

The audit compares two decoders that consume the same parsed Json and must
produce the same final values (checked against the frozen oracle):

  generated    the Argonaut decoder emitted by gopurs (selected workspace)
  specialized  the hand-written schema-specialised decoder in decode.go

Both run in the same process, in alternating passes, so parser state, allocator
state and background load are shared. This script copies an official diagnostic
workspace, grafts the audit runner, builds `benchmark-audit`, and launches
alternating processes.

Usage:
  python3 run.py setup --workspace NEW --source VAR/json-dec-cache7-20260923
  python3 run.py campaign --workspace NEW --output VAR/json-dec-specialized-20260924/run1
"""
import argparse
import datetime
import hashlib
import json
import os
import platform
import shutil
import statistics
import subprocess
import sys
import time
from pathlib import Path

SPECIALIZED = Path(__file__).resolve().parent
ROOT = SPECIALIZED.parents[3]
DEFAULT_SOURCE = ROOT / 'var/benchmark/json-dec-cache7-20260923'
FIXTURES = ROOT / 'test/fixtures/json-decoding'


def environment(extra=None):
    env = {k: v for k, v in os.environ.items()
           if not k.startswith(('GOPURS_', 'NEUTRAL_', 'DIAG_'))
           and k not in ['PPROF', 'GODEBUG', 'GOMEMLIMIT', 'GOFLAGS', 'GOEXPERIMENT', 'NODE_OPTIONS']}
    env.update(GOMAXPROCS='1', GOGC='100', GOWORK='off')
    if extra:
        env.update(extra)
    return env


def call(command, cwd, log, env):
    print(' '.join(str(part) for part in command), flush=True)
    with log.open('w') as out:
        result = subprocess.run([str(part) for part in command], cwd=cwd, env=env,
                                stdout=out, stderr=subprocess.STDOUT)
    if result.returncode:
        print(log.read_text()[-10000:], file=sys.stderr)
        raise RuntimeError(f'Failed: {log}')


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def generated_sources(work):
    output = work / 'output'
    paths = list(output.rglob('*.go')) + [output / 'go.mod']
    if (output / 'go.sum').exists():
        paths.append(output / 'go.sum')
    return {str(path.relative_to(work)): sha(path) for path in sorted(paths)}


def verify_build(work):
    provenance = json.loads((work / 'audit-build.json').read_text())
    if provenance['generated_sources'] != generated_sources(work):
        raise SystemExit(f'Generated audit sources changed: {work}')
    if provenance['binary_sha256'] != sha(work / 'output/benchmark-audit'):
        raise SystemExit(f'Audit binary changed: {work}')
    if provenance['source_manifest_sha256'] != sha(work / 'manifest.json'):
        raise SystemExit(f'Source build manifest changed: {work}')
    return provenance


def setup(args):
    work = args.workspace.resolve()
    source = args.source.resolve()
    if work.exists():
        raise SystemExit(f'workspace already exists: {work}')
    if not (source / 'output' / 'purescript').is_dir():
        raise SystemExit(f'source workspace is not a built diagnostic: {source}')
    shutil.copytree(source, work, ignore=shutil.ignore_patterns('logs'))
    (work / 'logs').mkdir(exist_ok=True)
    shutil.copyfile(SPECIALIZED / 'decode.go', work / 'output/purescript/zz_specialized.go')
    shutil.copyfile(SPECIALIZED / 'audit_main.go', work / 'output/purescript/zz_audit_main.go')
    shutil.copyfile(SPECIALIZED / 'audit_main_test.go', work / 'output/purescript/zz_audit_main_test.go')
    shutil.copyfile(SPECIALIZED / 'main_audit.go', work / 'output/main/main.go')
    build_env = environment({'GOMAXPROCS': '14'})
    call(['go', 'test', '-count=1', '-run', '^TestAuditValidatesEveryTimedOutput$', './purescript'],
         work / 'output', work / 'logs/go-audit-test.log', build_env)
    call(['go', 'build', '-pgo=off', '-o', 'benchmark-audit', './main'],
         work / 'output', work / 'logs/go-audit.log', build_env)
    binary = work / 'output' / 'benchmark-audit'
    provenance = {
        'source_workspace': str(source),
        'source_manifest_sha256': sha(work / 'manifest.json'),
        'source_build': json.loads((work / 'manifest.json').read_text()),
        'generated_sources': generated_sources(work),
        'binary_sha256': sha(binary),
        'audit_sources': {path.name: sha(path) for path in sorted(SPECIALIZED.glob('*'))
                          if path.suffix in {'.py', '.go'}},
        'go': subprocess.check_output(['go', 'version'], env=build_env, text=True).strip(),
        'command': ['go', 'build', '-pgo=off', '-o', 'benchmark-audit', './main'],
    }
    (work / 'audit-build.json').write_text(json.dumps(provenance, indent=2) + '\n')
    print(f'audit binary: {binary}', flush=True)


def campaign(args):
    work = args.workspace.resolve()
    workspaces = {'current': work}
    if args.compare_workspace:
        workspaces['baseline'] = args.compare_workspace.resolve()
    builds = {name: verify_build(path) for name, path in workspaces.items()}
    out = args.output.resolve()
    out.mkdir(parents=True, exist_ok=False)
    (out / 'corpus.json').write_bytes((FIXTURES / 'corpus.json').read_bytes())
    oracle_path = FIXTURES / 'expected.json'
    oracle = json.loads(oracle_path.read_text())
    env = environment({'DIAG_CORPUS': str(out / 'corpus.json'),
                       'DIAG_ORACLE': str(oracle_path),
                       'DIAG_WARMUPS': str(args.warmups),
                       'DIAG_SAMPLES': str(args.samples),
                       'DIAG_PHASES': args.phases})
    if args.gogc:
        env['GOGC'] = args.gogc
    phases = args.phases.split(',')
    if not phases or len(set(phases)) != len(phases) or any(p not in ['parse', 'decode', 'combined'] for p in phases):
        raise SystemExit('phases must be distinct parse/decode/combined names')
    runs = {name: [] for name in workspaces}
    schedule = []
    for index in range(args.processes):
        offset = index % len(phases)
        phase_order = phases[offset:] + phases[:offset]
        if (index // len(phases)) % 2:
            phase_order.reverse()
        process_env = dict(env, DIAG_PHASES=','.join(phase_order), DIAG_START_MODE=str(index % 2))
        order = list(workspaces)
        if index % 2:
            order.reverse()
        for name in order:
            directory = workspaces[name]
            log = out / f'{index:02d}-{name}-go.log'
            observation = {'pair': index, 'workspace': name, 'phase_order': phase_order,
                           'start_mode': 'specialized' if index % 2 else 'generated',
                           'utc': datetime.datetime.now(datetime.timezone.utc).isoformat(),
                           'load_before': os.getloadavg()}
            start = time.monotonic()
            call([directory / 'output/benchmark-audit'], directory, log, process_env)
            observation.update(wall_seconds=time.monotonic() - start, load_after=os.getloadavg())
            report = json.loads(log.read_text())
            for mode, fingerprints in report['fingerprints'].items():
                if fingerprints != oracle['fingerprints']:
                    raise SystemExit(f'run {index}: {mode} fingerprints differ from the oracle')
            if report['json_fingerprints'] != oracle['json_fingerprints']:
                raise SystemExit(f'run {index}: json fingerprints differ from the oracle')
            if report['retained_results_per_pass'] != oracle['timed_cases'] or report['validated_timed_results'] != oracle['timed_cases'] * args.samples * 2 * len(phases):
                raise SystemExit('Timed-output validation protocol mismatch')
            runs[name].append(report)
            schedule.append(observation)
            summary = {phase: {mode: round(data['time_us'] / 1000, 3)
                               for mode, data in modes.items()}
                       for phase, modes in report['phases'].items()}
            print(index, name, summary, flush=True)
    medians, allocations = {}, {}
    for name, reports in runs.items():
        medians[name], allocations[name] = {}, {}
        for phase in phases:
            medians[name][phase], allocations[name][phase] = {}, {}
            for mode in reports[0]['phases'][phase]:
                medians[name][phase][mode] = statistics.median(
                    run['phases'][phase][mode]['time_us'] for run in reports)
                allocations[name][phase][mode] = min(
                    sample['allocated_bytes']
                    for run in reports for sample in run['phases'][phase][mode]['samples'])
    for directory in workspaces.values():
        verify_build(directory)
    results = {
        'audit': 'specialized-decoder',
        'protocol': {
            'processes_per_workspace': args.processes,
            'warmups_per_mode': args.warmups,
            'samples_per_mode': args.samples,
            'alternating_passes_in_process': True,
            'alternating_start_mode': True,
            'counterbalanced_workspace_order': bool(args.compare_workspace),
            'rotating_and_reversing_phase_order': True,
            'retained_outputs_validated_after_each_pass': True,
            'retention_buffer_allocation_timed': False,
            'fingerprinting_timed': False,
            'phases': args.phases,
            'cell': 'median of process minima',
            'GOMAXPROCS': 1,
            'GOGC': args.gogc or '100',
            'monotonic_timing': True,
        },
        'medians_us': medians,
        'min_allocated_bytes_per_corpus': allocations,
        'oracle': {'names': oracle['names'], 'timed_cases': oracle['timed_cases']},
        'corpus_sha256': sha(out / 'corpus.json'),
        'oracle_sha256': sha(oracle_path),
        'runner_sha256': sha(Path(__file__).resolve()),
        'machine': {'platform': platform.platform(), 'cpu_count': os.cpu_count()},
        'environment': {key: env[key] for key in sorted(env) if key.startswith(('GO', 'DIAG_'))},
        'builds': builds,
        'schedule': schedule,
        'runs': runs,
    }
    (out / 'results.json').write_text(json.dumps(results, indent=2) + '\n')
    print(json.dumps(medians, indent=2), flush=True)
    print(json.dumps(allocations, indent=2), flush=True)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('action', choices=['setup', 'campaign'])
    parser.add_argument('--workspace', type=Path, required=True)
    parser.add_argument('--source', type=Path, default=DEFAULT_SOURCE)
    parser.add_argument('--compare-workspace', type=Path, help='paired campaign against a second preserved audit build')
    parser.add_argument('--output', type=Path)
    parser.add_argument('--processes', type=int, default=6)
    parser.add_argument('--warmups', type=int, default=2)
    parser.add_argument('--samples', type=int, default=5)
    parser.add_argument('--phases', default='parse,decode,combined')
    parser.add_argument('--gogc', help='override GOGC for the campaign')
    args = parser.parse_args()
    if args.processes < 1 or args.samples < 1 or args.warmups < 0:
        parser.error('processes and samples must be positive; warmups must be nonnegative')
    if args.action == 'setup':
        setup(args)
    else:
        if not args.output:
            parser.error('campaign requires --output NEW_DIRECTORY')
        campaign(args)


if __name__ == '__main__':
    main()

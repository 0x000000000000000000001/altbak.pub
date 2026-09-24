#!/usr/bin/env python3
"""Build and run the specialised-decoder A/B audit.

The audit compares two decoders that consume the same parsed Json and must
produce the same final values (checked against the frozen oracle):

  generated    the Argonaut decoder emitted by gopurs (cache7 workspace)
  specialized  the hand-written schema-specialised decoder in decode.go

Both run in the same process, in alternating passes, so parser state, allocator
state and background load are shared. This script copies the official cache7
workspace, grafts the audit runner, builds `benchmark-audit`, and launches
alternating processes.

Usage:
  python3 run.py setup --workspace NEW --source VAR/json-dec-cache7-20260923
  python3 run.py campaign --workspace NEW --output VAR/json-dec-specialized-20260924/run1
"""
import argparse
import json
import os
import shutil
import statistics
import subprocess
import sys
from pathlib import Path

SPECIALIZED = Path(__file__).resolve().parent
ROOT = SPECIALIZED.parents[3]
DEFAULT_SOURCE = ROOT / 'var/benchmark/json-dec-cache7-20260923'
FIXTURES = ROOT / 'test/fixtures/json-decoding'


def environment(extra=None):
    env = {k: v for k, v in os.environ.items()
           if not k.startswith(('GOPURS_', 'NEUTRAL_'))
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
    shutil.copyfile(SPECIALIZED / 'main_audit.go', work / 'output/main/main.go')
    build_env = environment({'GOMAXPROCS': '14'})
    call(['go', 'build', '-pgo=off', '-o', 'benchmark-audit', './main'],
         work / 'output', work / 'logs/go-audit.log', build_env)
    binary = work / 'output' / 'benchmark-audit'
    print(f'audit binary: {binary}', flush=True)


def campaign(args):
    work = args.workspace.resolve()
    binary = work / 'output' / 'benchmark-audit'
    if not binary.exists():
        raise SystemExit(f'audit binary not found: {binary}')
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
    runs = []
    for index in range(args.processes):
        log = out / f'{index:02d}-go.log'
        call([str(binary)], work, log, env)
        report = json.loads(log.read_text())
        for mode, fingerprints in report['fingerprints'].items():
            if fingerprints != oracle['fingerprints']:
                raise SystemExit(f'run {index}: {mode} fingerprints differ from the oracle')
        if report['json_fingerprints'] != oracle['json_fingerprints']:
            raise SystemExit(f'run {index}: json fingerprints differ from the oracle')
        runs.append(report)
        summary = {phase: {mode: round(data['time_us'] / 1000, 3)
                           for mode, data in modes.items()}
                   for phase, modes in report['phases'].items()}
        print(index, summary, flush=True)
    medians = {}
    allocations = {}
    for phase in runs[0]['phases']:
        medians[phase] = {}
        allocations[phase] = {}
        for mode in runs[0]['phases'][phase]:
            medians[phase][mode] = statistics.median(
                run['phases'][phase][mode]['time_us'] for run in runs)
            allocations[phase][mode] = min(
                sample['allocated_bytes']
                for run in runs for sample in run['phases'][phase][mode]['samples'])
    results = {
        'audit': 'specialized-decoder',
        'protocol': {
            'processes': args.processes,
            'warmups_per_mode': args.warmups,
            'samples_per_mode': args.samples,
            'alternating_passes_in_process': True,
            'phases': args.phases,
            'cell': 'median of process minima',
            'GOMAXPROCS': 1,
            'GOGC': args.gogc or '100',
            'monotonic_timing': True,
        },
        'medians_us': medians,
        'min_allocated_bytes_per_corpus': allocations,
        'oracle': {'names': oracle['names'], 'timed_cases': oracle['timed_cases']},
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
    parser.add_argument('--output', type=Path)
    parser.add_argument('--processes', type=int, default=9)
    parser.add_argument('--warmups', type=int, default=2)
    parser.add_argument('--samples', type=int, default=5)
    parser.add_argument('--phases', default='decode,combined')
    parser.add_argument('--gogc', help='override GOGC for the campaign')
    args = parser.parse_args()
    if args.action == 'setup':
        setup(args)
    else:
        if not args.output:
            parser.error('campaign requires --output NEW_DIRECTORY')
        campaign(args)


if __name__ == '__main__':
    main()

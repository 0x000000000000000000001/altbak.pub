#!/usr/bin/env python3
"""Pair preserved official Go diagnostics, retaining each driver's output checks.

Source/compiler changes between builds are intentional. Validate each pinned
binary against its own manifest, hash its emitted Go, and reject any artifact
change during the campaign. Both drivers retain and validate every timed result;
their setup fingerprints are additionally checked against the frozen oracle here.
"""
import argparse
import datetime
import importlib.util
import json
import os
from pathlib import Path
import statistics
import subprocess
import time

RUNNER = Path(__file__).resolve().parents[1] / 'json-diagnostic.py'
spec = importlib.util.spec_from_file_location('diagnostic', RUNNER)
diagnostic = importlib.util.module_from_spec(spec)
spec.loader.exec_module(diagnostic)


def provenance(work, suite):
    manifest = json.loads((work / 'manifest.json').read_text())
    diagnostic.check_suite(manifest, suite)
    binary = diagnostic.sha(work / 'benchmark')
    if binary != manifest['binary_sha256']:
        raise ValueError(f'Binary differs from manifest: {work}')
    sources = {str(p.relative_to(work)): diagnostic.sha(p)
               for p in sorted((work / 'output').rglob('*.go'))}
    return {'workspace': str(work), 'manifest': manifest,
            'manifest_sha256': diagnostic.sha(work / 'manifest.json'),
            'binary_sha256': binary, 'generated_go': sources}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--suite', required=True, choices=diagnostic.SUITES)
    parser.add_argument('--baseline', type=Path, required=True)
    parser.add_argument('--current', type=Path, required=True)
    parser.add_argument('--output', type=Path, required=True)
    parser.add_argument('--processes', type=int, default=6)
    args = parser.parse_args()
    if args.processes < 1:
        parser.error('processes must be positive')
    workspaces = {'baseline': args.baseline.resolve(), 'current': args.current.resolve()}
    builds = {name: provenance(work, args.suite) for name, work in workspaces.items()}
    out = args.output.resolve()
    out.mkdir(parents=True, exist_ok=False)
    corpus, info = diagnostic.corpus_data(args.suite)
    (out / 'corpus.json').write_bytes(corpus)
    oracle = json.loads((diagnostic.fixtures(args.suite) / 'expected.json').read_text())
    env = diagnostic.environment()
    env['DIAG_CORPUS'] = str(out / 'corpus.json')
    phases = ['parse', 'decode', 'combined']
    runs = {name: [] for name in workspaces}
    schedule = []
    for index in range(args.processes):
        phase_order = phases
        # JsonTypedAst's preserved driver has a fixed phase order. Do not claim
        # phase counterbalancing for a binary that ignores DIAG_PHASES.
        if args.suite == 'JsonDecoding':
            offset = index % 3
            phase_order = phases[offset:] + phases[:offset]
            if (index // 3) % 2:
                phase_order.reverse()
            env['DIAG_PHASES'] = ','.join(phase_order)
        order = list(workspaces)
        if index % 2:
            order.reverse()
        for name in order:
            work = workspaces[name]
            log = out / f'{index:02d}-{name}.json'
            observation = {'pair': index, 'workspace': name, 'phases': phase_order,
                           'utc': datetime.datetime.now(datetime.timezone.utc).isoformat(),
                           'load_before': os.getloadavg()}
            start = time.monotonic()
            with log.open('w') as stream:
                subprocess.run([work / 'benchmark'], cwd=work, env=env, stdout=stream, check=True)
            observation.update(wall_seconds=time.monotonic() - start, load_after=os.getloadavg())
            report = json.loads(log.read_text())
            diagnostic.validate_result(report, oracle, args.suite)
            for phase in phases:
                if len(report['phases'][phase]['samples']) != 5:
                    raise ValueError('Unexpected driver sample count')
            runs[name].append(report)
            schedule.append(observation)
            print(index, name, {phase: round(report['phases'][phase]['time_us']/1000, 3)
                               for phase in phases}, flush=True)
    medians = {name: {phase: statistics.median(r['phases'][phase]['time_us'] for r in reports)
                      for phase in phases} for name, reports in runs.items()}
    allocations = {name: {phase: min(s['allocated_bytes'] for r in reports
                                    for s in r['phases'][phase]['samples'])
                         for phase in phases} for name, reports in runs.items()}
    for name, work in workspaces.items():
        if provenance(work, args.suite) != builds[name]:
            raise ValueError(f'Build changed during measurements: {name}')
    result = {'suite': args.suite, 'protocol': {
        'processes_per_workspace': args.processes, 'warmups_per_phase': 2, 'samples_per_phase': 5,
        'GOMAXPROCS': 1, 'GOGC': 100, 'cell': 'median of process minima',
        'counterbalanced_workspace_order': True,
        'rotating_and_reversing_phase_order': args.suite == 'JsonDecoding',
        'retained_outputs_validated_after_each_pass': True,
        'fingerprinting_timed': False, 'retention_buffer_allocation_timed': False},
        'builds': builds, 'runs': runs, 'schedule': schedule,
        'medians_us': medians, 'min_allocated_bytes_per_corpus': allocations,
        'corpus': info, 'corpus_sha256': diagnostic.sha(out / 'corpus.json'),
        'oracle_sha256': diagnostic.sha(diagnostic.fixtures(args.suite) / 'expected.json'),
        'runner_sha256': diagnostic.sha(Path(__file__).resolve()),
        'go': subprocess.check_output(['go', 'version'], env=env, text=True).strip()}
    (out / 'results.json').write_text(json.dumps(result, indent=2) + '\n')
    print(json.dumps({'medians_us': medians, 'allocations': allocations}, indent=2))


if __name__ == '__main__':
    main()

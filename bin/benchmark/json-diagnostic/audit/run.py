#!/usr/bin/env python3
"""Audit runs for the C++ references.

Three configurations per suite:
  stock   - the workspace reference binary, default behaviour;
  audit   - the audit variant, DIAG_C_FULL_CYCLE unset (must match stock);
  freeing - the audit variant with DIAG_C_FULL_CYCLE=1, which releases the
            arena and the payload vector inside the measured pass and reports
            the release time separately (release_us).

Protocol mirrors bin/benchmark/json-diagnostic.py: three processes per
configuration, two warm-up passes and five sampled passes per phase, the
minimum per process and the median across processes. Successful fingerprints
are validated against the frozen oracle.
"""
import argparse
import gzip
import hashlib
import json
import os
import statistics
import subprocess
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[3]
EXTRA_SUCCESSES = {'optional-fields-missing', 'optional-fields-null'}


SUITE_FIXTURE = {'JsonDecoding': ('json-decoding', 'corpus.json'),
                 'JsonTypedAst': ('json-typed-ast', 'tast-corpus.json')}


def sha(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def fixture_corpus(suite: str) -> Path:
    directory, name = SUITE_FIXTURE[suite]
    return ROOT / 'test/fixtures' / directory / name


def corpus_file(suite: str, workdir: Path) -> Path:
    """The runners hand the binary the decompressed corpus; mirror that."""
    if suite == 'JsonTypedAst':
        target = workdir / 'corpus.json'
        target.write_bytes(gzip.decompress((ROOT / 'test/fixtures/json-typed-ast/tast-corpus.json.gz').read_bytes()))
        return target
    return fixture_corpus(suite)


def oracle_for(suite: str):
    directory, _ = SUITE_FIXTURE[suite]
    oracle = json.loads((ROOT / 'test/fixtures' / directory / 'expected.json').read_text())
    corpus = json.loads(fixture_corpus(suite).read_text())
    if suite == 'JsonDecoding':
        expected_success = {case['name'] for case in corpus if case['benchmark']} | EXTRA_SUCCESSES
    else:
        expected_success = None
    return oracle, expected_success


def validate(result: dict, oracle: dict, expected_success, suite: str) -> None:
    if result['modules'] != oracle['modules']:
        raise ValueError(f'{suite}: corpus mismatch')
    if suite == 'JsonTypedAst':
        if result['fingerprints'] != oracle['fingerprints']:
            raise ValueError('JsonTypedAst: payload fingerprint mismatch')
        if result['json_fingerprints'] != oracle['json_fingerprints']:
            raise ValueError('JsonTypedAst: parse fingerprint mismatch')
        return
    if result['timed_cases'] != oracle['timed_cases']:
        raise ValueError(f'{suite}: corpus mismatch')
    if 'names' in result and 'names' in oracle and result['names'] != oracle['names']:
        raise ValueError(f'{suite}: names mismatch')
    if 'json_fingerprints' in result and 'json_fingerprints' in oracle \
            and result['json_fingerprints'] != oracle['json_fingerprints']:
        raise ValueError(f'{suite}: parse fingerprints mismatch')
    for name, got, want in zip(result['names'], result['fingerprints'], oracle['fingerprints']):
        if name in expected_success:
            if got != want:
                raise ValueError(f'{suite}: fingerprint mismatch for {name}')
        elif got not in ('', None):
            raise ValueError(f'{suite}: expected {name} to fail')


def run_config(binary: Path, suite: str, corpus: Path, runs: int, extra_env: dict) -> dict:
    oracle, expected_success = oracle_for(suite)
    processes = []
    env = {**os.environ, 'DIAG_CORPUS': str(corpus), **extra_env}
    for _ in range(runs):
        completed = subprocess.run([str(binary)], capture_output=True, text=True, check=True, env=env)
        result = json.loads(completed.stdout)
        validate(result, oracle, expected_success, suite)
        processes.append(result)
    phases = sorted({phase for result in processes for phase in result['phases']})
    medians = {phase: statistics.median(r['phases'][phase]['time_us'] for r in processes) for phase in phases}
    releases = {phase: statistics.median(r['phases'][phase].get('release_us', 0) for r in processes)
                for phase in phases}
    return {'processes': processes, 'medians_us': medians, 'release_us': releases}


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument('--suite', required=True, choices=['JsonDecoding', 'JsonTypedAst'])
    parser.add_argument('--stock', required=True)
    parser.add_argument('--audit', required=True)
    parser.add_argument('--output', required=True)
    parser.add_argument('--runs', type=int, default=3)
    args = parser.parse_args()

    output = Path(args.output).resolve()
    output.mkdir(parents=True, exist_ok=True)
    corpus = corpus_file(args.suite, output)
    stock = Path(args.stock).resolve()
    audit = Path(args.audit).resolve()
    report = {
        'suite': args.suite,
        'protocol': {'processes_per_configuration': args.runs, 'warmups_per_phase': 2,
                     'samples_per_phase': 5, 'cell': 'median of process minimum times',
                     'GOMAXPROCS': 1, 'GOGC': 100},
        'corpus_sha256': sha(corpus),
        'stock': {'binary': str(stock), 'binary_sha256': sha(stock)},
        'audit': {'binary': str(audit), 'binary_sha256': sha(audit)},
        'configurations': {},
    }
    for label, binary, env in [('stock', stock, {}), ('audit', audit, {}),
                               ('freeing', audit, {'DIAG_C_FULL_CYCLE': '1'})]:
        print(f'== {label} ==', flush=True)
        report['configurations'][label] = run_config(binary, args.suite, corpus, args.runs, env)
        config = report['configurations'][label]
        medians = ', '.join(f'{phase} {value/1000:.3f} ms' for phase, value in config['medians_us'].items())
        releases = ', '.join(f'{phase} free {value/1000:.3f} ms' for phase, value in config['release_us'].items())
        print(f'  {medians} | {releases}', flush=True)
    (output / 'results.json').write_text(json.dumps(report, indent=2) + '\n')
    print(f'written {output / "results.json"}', flush=True)


if __name__ == '__main__':
    main()

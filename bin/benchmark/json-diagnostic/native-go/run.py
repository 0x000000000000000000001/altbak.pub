#!/usr/bin/env python3
"""Audit runner: hand-written Go decoders against the frozen JSON oracle.

Same protocol as bin/benchmark/json-diagnostic.py: three independent
processes, two warm-up passes and five sampled passes per phase, the minimum
per process and the median across processes, GOMAXPROCS=1 and GOGC=100.
Fingerprints are computed outside the timed intervals and validated here
against test/fixtures/json-decoding/expected.json.
"""
import hashlib
import json
import os
import statistics
import subprocess
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[3]
FIXTURE = ROOT / 'test/fixtures/json-decoding'
CORPUS = FIXTURE / 'corpus.json'
ORACLE = json.loads((FIXTURE / 'expected.json').read_text())
EXTRA_SUCCESSES = {'optional-fields-missing', 'optional-fields-null'}
MODES = ['dom', 'dom-clone', 'std', 'std-goccy']


def sha(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def build(out: Path) -> None:
    out.mkdir(parents=True, exist_ok=True)
    env = {**os.environ, 'GOWORK': 'off'}
    subprocess.run(['go', 'mod', 'tidy'], cwd=HERE, check=True, env=env)
    subprocess.run(['go', 'build', '-pgo=off', '-o', str(out / 'native-go'), '.'],
                   cwd=HERE, check=True, env=env)
    subprocess.run(['go', 'build', '-pgo=off', '-tags', 'goccy',
                    '-o', str(out / 'native-go-goccy'), '.'],
                   cwd=HERE, check=True, env=env)


def validate(result: dict, mode: str) -> None:
    if result['names'] != ORACLE['names'] or result['timed_cases'] != ORACLE['timed_cases']:
        raise ValueError(f'{mode}: corpus mismatch')
    if result['json_fingerprints'] != ORACLE['json_fingerprints']:
        for name, got, want in zip(result['names'], result['json_fingerprints'], ORACLE['json_fingerprints']):
            if got != want:
                print(f'  json fingerprint mismatch: {name}: {got[:12]} != {want[:12]}', flush=True)
        raise ValueError(f'{mode}: parse fingerprints do not match the frozen oracle')
    expected_success = {item for item in [c['name'] for c in json.loads(CORPUS.read_text()) if c['benchmark']]} | EXTRA_SUCCESSES
    for name, got, want in zip(result['names'], result['fingerprints'], ORACLE['fingerprints']):
        if name in expected_success:
            if got != want:
                raise ValueError(f'{mode}: fingerprint mismatch for {name}')
        elif got != '':
            raise ValueError(f'{mode}: expected {name} to fail, got a fingerprint')


def run_mode(binary: Path, mode: str, runs: int) -> dict:
    processes = []
    env = {**os.environ, 'DIAG_CORPUS': str(CORPUS), 'GOMAXPROCS': '1', 'GOGC': '100'}
    for _ in range(runs):
        completed = subprocess.run([str(binary), '-mode', mode], capture_output=True,
                                   text=True, check=True, env=env)
        result = json.loads(completed.stdout)
        validate(result, mode)
        processes.append(result)
    phases = sorted({phase for result in processes for phase in result['phases']})
    medians = {phase: statistics.median(r['phases'][phase]['time_us'] for r in processes)
               for phase in phases}
    allocated = {phase: statistics.median(
        sample['allocated_bytes'] for r in processes for sample in r['phases'][phase]['samples'])
        for phase in phases}
    return {
        'mode': mode,
        'binary': binary.name,
        'binary_sha256': sha(binary),
        'processes': processes,
        'medians_us': medians,
        'allocated_bytes_per_corpus': allocated,
    }


def main() -> None:
    output = Path(sys.argv[1]).resolve() if len(sys.argv) > 1 else (ROOT / 'var/benchmark/json-dec-native-go-audit-20260924')
    runs = int(sys.argv[2]) if len(sys.argv) > 2 else 3
    output.mkdir(parents=True, exist_ok=True)
    build(output)
    report = {
        'protocol': {
            'processes_per_mode': runs,
            'warmups_per_phase': 2,
            'samples_per_phase': 5,
            'cell': 'median of process minimum times',
            'GOMAXPROCS': 1,
            'GOGC': 100,
            'fingerprinting_timed': False,
        },
        'corpus_sha256': sha(CORPUS),
        'oracle_sha256': sha(FIXTURE / 'expected.json'),
        'modes': {},
    }
    for mode in MODES:
        binary = output / ('native-go-goccy' if 'goccy' in mode else 'native-go')
        if 'goccy' in mode and not binary.exists():
            continue
        if not binary.exists():
            continue
        print(f'== {mode} ==', flush=True)
        try:
            report['modes'][mode] = run_mode(binary, mode, runs)
        except subprocess.CalledProcessError as error:
            print(f'{mode}: process failed: {error.stderr[-400:] if error.stderr else error}', flush=True)
            continue
        medians = report['modes'][mode]['medians_us']
        print('  ' + ', '.join(f'{phase} {value/1000:.3f} ms' for phase, value in medians.items()), flush=True)
    (output / 'results.json').write_text(json.dumps(report, indent=2) + '\n')
    print(f'written {output / "results.json"}', flush=True)


if __name__ == '__main__':
    main()

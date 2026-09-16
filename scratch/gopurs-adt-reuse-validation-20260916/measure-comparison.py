#!/usr/bin/env python3
"""Measure already built artifacts sequentially; run only with the build slot idle."""
import argparse
import datetime
import hashlib
import importlib.util
import json
import os
from pathlib import Path
import statistics
import subprocess
import sys

sys.dont_write_bytecode = True
BASE = Path(__file__).resolve().parent
ROOT = BASE.parents[1]
parser = argparse.ArgumentParser(description=__doc__)
parser.add_argument('--output', type=Path, default=BASE / 'comparison')
args = parser.parse_args()
destination = args.output.resolve()
destination.mkdir()  # Never overwrite an earlier campaign.
spec = importlib.util.spec_from_file_location('benchmark_validator', ROOT / 'bin/benchmark/validate.py')
validator = importlib.util.module_from_spec(spec)
spec.loader.exec_module(validator)
env = dict(os.environ, GOGC='800', GOWORK='off')
env.pop('PPROF', None)


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def execute(key, repetition, binary, arguments):
    started = datetime.datetime.now(datetime.timezone.utc).isoformat()
    digest = sha(binary)
    result = subprocess.run([str(binary), *arguments], cwd=binary.parent, env=env,
                            capture_output=True, text=True, timeout=120)
    (destination / f'{key}-{repetition}.stdout.log').write_text(result.stdout)
    (destination / f'{key}-{repetition}.stderr.log').write_text(result.stderr)
    if result.returncode:
        raise RuntimeError(f'{key} failed with {result.returncode}; see retained logs')
    if sha(binary) != digest:
        raise RuntimeError('Executable changed while measuring')
    return result.stdout, {'key': key, 'repetition': repetition, 'started_utc': started,
                          'binary': str(binary), 'binary_sha256': digest}


suite_plan = [('before-pure', 'pure'), ('after-pure', 'pure'), ('after-fficc', 'fficc')]
suite_runs = []
for repetition in range(1, 4):
    offset = repetition - 1
    for key, mode in suite_plan[offset:] + suite_plan[:offset]:
        binary = BASE / key / 'benchmark'
        manifest = json.loads((binary.parent / 'manifest.json').read_text())
        if sha(binary) != manifest['binary_sha256']:
            raise RuntimeError('Binary differs from its build manifest: ' + key)
        output, record = execute(key, repetition, binary, [])
        record['validated'] = validator.validate_output(output, mode)
        suite_runs.append(record)
        print(key, repetition, record['validated']['sum_displayed_lines_ms'], flush=True)

probe_plan = [('probe-before-owned', 'pure'), ('probe-after-owned', 'pure'), ('probe-after-owned', 'fficc')]
probe_runs = []
for repetition in range(1, 4):
    offset = repetition - 1
    for directory, mode in probe_plan[offset:] + probe_plan[:offset]:
        key = directory + '-' + mode
        output, record = execute(key, repetition, BASE / directory / 'probe',
                                 ['-mode', mode, '-n', '100000', '-iterations', '5'])
        samples = [json.loads(line) for line in output.splitlines()]
        if len(samples) != 5 or any(sample['depth'] != 22 or sample['input'] != 100000 for sample in samples):
            raise RuntimeError('Invalid probe results: ' + key)
        record['samples'] = samples
        probe_runs.append(record)

summary = {'suites': {}, 'probes': {}}
for key, _ in suite_plan:
    records = [record['validated'] for record in suite_runs if record['key'] == key]
    medians = [statistics.median(column) for column in zip(*(record['times_us'] for record in records))]
    summary['suites'][key] = {'times_us': medians, 'sum_row_medians_ms': sum(medians) / 1000,
                             'process_totals_ms': [record['sum_displayed_lines_ms'] for record in records],
                             'rbtree_us': medians[8]}
for directory, mode in probe_plan:
    key = directory + '-' + mode
    records = [record for record in probe_runs if record['key'] == key]
    samples = [sample for record in records for sample in record['samples']]
    summary['probes'][key] = {field: statistics.median(sample[field] for sample in samples)
                              for field in ['nanoseconds', 'allocated_bytes', 'allocations', 'depth']}
    summary['probes'][key]['sample_count'] = len(samples)
    summary['probes'][key]['entrypoints'] = json.loads((BASE / directory / 'entrypoints.json').read_text())

report = {'environment': {'GOGC': env['GOGC'], 'GOWORK': env['GOWORK'], 'GOMAXPROCS': env.get('GOMAXPROCS')},
          'suite_protocol': '3 processes per variant, rotated order; 3 global warmups, 3/test, best of 10; row medians',
          'probe_protocol': '3 processes per variant, rotated order; 3 warmups, 5 calls/process, GC before each call outside timer',
          'suite_runs': suite_runs, 'probe_runs': probe_runs, 'summary': summary}
(destination / 'results.json').write_text(json.dumps(report, indent=2) + '\n')
print(json.dumps(summary, indent=2))

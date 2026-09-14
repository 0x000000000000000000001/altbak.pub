#!/usr/bin/env python3
"""Re-run already built artifacts sequentially, validate, and summarize 3 runs."""
import argparse
import datetime
import importlib.util
import json
from pathlib import Path
import statistics
import subprocess
import sys
import time

sys.dont_write_bytecode = True
HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
spec = importlib.util.spec_from_file_location('validate', ROOT / 'bin/benchmark/validate.py')
validator = importlib.util.module_from_spec(spec)
spec.loader.exec_module(validator)


def summarize(plan):
    results = {}
    for item in plan:
        runs = [json.loads((HERE / f"measure-{number}-{item['key']}.json").read_text()) for number in range(1, 4)]
        times = [statistics.median(values) for values in zip(*(r['times_us'] for r in runs))]
        results[item['key']] = {'mode': item['mode'], 'command': item['command'],
            'times_us': times, 'total_ms': sum(times) / 1000,
            'run_totals_ms': [r['sum_displayed_lines_ms'] for r in runs],
            'values': runs[0]['values'], 'aggregation': 'row median of 3 independent processes, each best of 10'}
    (HERE / 'results.json').write_text(json.dumps(results, indent=2) + '\n')
    return results


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--summarize-only', action='store_true')
    parser.add_argument('--keys', nargs='*')
    args = parser.parse_args()
    plan = json.loads((HERE / 'measurement-plan.json').read_text())
    if args.keys:
        plan = [item for item in plan if item['key'] in args.keys]
    if not args.summarize_only:
        for number in range(1, 4):
            # Rotate modes in each group across the independent repetitions.
            ordered = []
            for group in dict.fromkeys(item['group'] for item in plan):
                items = [item for item in plan if item['group'] == group]
                offset = (number - 1) % len(items)
                ordered.extend(items[offset:] + items[:offset])
            for item in ordered:
                key = item['key']
                target = HERE / f'measure-{number}-{key}.json'
                if target.exists():
                    raise RuntimeError(f'Refusing to overwrite an existing measurement: {target}')
                started = datetime.datetime.now(datetime.timezone.utc).isoformat()
                print(f'{started}: round {number}/3 {key}', flush=True)
                begin = time.monotonic()
                result = subprocess.run(item['command'], cwd=ROOT, capture_output=True, text=True, timeout=1200)
                (HERE / f'measure-{number}-{key}.stdout.log').write_text(result.stdout)
                (HERE / f'measure-{number}-{key}.stderr.log').write_text(result.stderr)
                if result.returncode:
                    raise RuntimeError(f'{key} exited {result.returncode}; see stderr log')
                checked = validator.validate_output(result.stdout, item['mode'])
                checked.update({'started_utc': started, 'process_seconds': time.monotonic() - begin,
                                'command': item['command'], 'key': key})
                target.write_text(json.dumps(checked, indent=2) + '\n')
                print(f"Validated 14 results: {key}, sum {checked['sum_displayed_lines_ms']:.5f} ms", flush=True)
                time.sleep(1)
    for key, result in summarize(plan).items():
        print(f"{key}: {result['total_ms']:.5f} ms")


if __name__ == '__main__':
    main()

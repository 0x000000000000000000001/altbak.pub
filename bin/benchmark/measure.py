#!/usr/bin/env python3
"""Measure a frozen build plan sequentially in three independent processes.

Each plan entry has key, mode, command (argv), and artifact. Numeric executables
also provide executable_sha256 from their build manifest. Build commands do
not belong in this plan: complete all builds before starting the campaign.
"""
import argparse
from datetime import datetime, timezone
import hashlib
import json
import math
import os
from pathlib import Path
import platform
import re
import statistics
import subprocess
import sys

sys.dont_write_bytecode = True
from validate import validate_output

ROOT = Path(__file__).resolve().parents[2]


def digest(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def fingerprints():
    return {str(p.relative_to(ROOT)): digest(p)
            for folder in ['src', 'srx', 'bin', 'test/native', 'tmp']
            for p in sorted((ROOT / folder).rglob('*'))
            if p.is_file() and 'node_modules' not in p.parts and p.suffix in
            {'.purs', '.js', '.go', '.rs', '.fs', '.java', '.cc', '.ss', '.erl', '.php', '.py', '.mjs', '.c', '.hs', '.kk', '.ml'}}


def check(output, mode):
    result = validate_output(output, mode)
    if not result['values_validated']:
        raise ValueError('A timing without validated results cannot be published')
    if any(not math.isfinite(t) or t <= 0 for t in result['times_us']):
        raise ValueError('Nonpositive or nonfinite per-call timing')
    if mode != 'x':
        counts = [int(n) for n in re.findall(r'^Batch iterations: (\d+)$', output, re.M)]
        if len(counts) != 14 or any(n < 1 or n > 16777216 or n & (n - 1) for n in counts):
            raise ValueError('Expected 14 calibrated power-of-two batch sizes')
        result['batch_iterations'] = counts
    return result


def aggregate(samples):
    if len(samples) != 3:
        raise ValueError('Three independent processes are required')
    if any(s['values'] != samples[0]['values'] or s['labels'] != samples[0]['labels'] for s in samples):
        raise ValueError('Benchmark outputs changed across processes')
    result = dict(samples[0])
    result.pop('batch_iterations', None)
    result['times_us'] = [statistics.median(s['times_us'][i] for s in samples)
                          for i in range(len(samples[0]['times_us']))]
    result['total_ms'] = result['sum_displayed_lines_ms'] = sum(result['times_us']) / 1000
    result['process_runs'] = samples
    result['aggregation'] = 'median per row across three processes; total sums median rows'
    return result


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--plan', type=Path, required=True)
    parser.add_argument('--output', type=Path, required=True)
    parser.add_argument('--resume', action='store_true')
    parser.add_argument('--timeout', type=float, default=600)
    args = parser.parse_args()
    plan = json.loads(args.plan.read_text())
    if not plan or len({p['key'] for p in plan}) != len(plan):
        raise ValueError('Plan keys must be nonempty and unique')
    for item in plan:
        if item['mode'] not in ['pure', 'ffi', 'fficc', 'x']:
            raise ValueError('Unsupported plan mode')
        if '--run-only' not in item['command'] and not item.get('executable_sha256'):
            raise ValueError('Use --run-only or a numeric executable with its build SHA')
        if not Path(item['artifact']).is_dir():
            raise ValueError('Missing frozen build: ' + item['artifact'])
        if not item.get('manifest') or digest(Path(item['manifest'])) != item.get('manifest_sha256'):
            raise ValueError('Missing or changed build manifest: ' + item['key'])
        if not item.get('payload_sha256'):
            raise ValueError('Missing executable/runtime payload hashes: ' + item['key'])
    directory = args.output.resolve()
    directory.mkdir(parents=True, exist_ok=True)
    manifest_path = directory / 'manifest.json'
    current = fingerprints()
    if args.resume:
        manifest = json.loads(manifest_path.read_text())
        if manifest['sources'] != current or manifest['plan'] != plan:
            raise ValueError('Inputs or plan changed; start a new campaign')
    else:
        if manifest_path.exists():
            raise ValueError('Output already contains a campaign; use --resume')
        manifest = {'created_at_utc': datetime.now(timezone.utc).isoformat(),
                    'sources': current, 'plan': plan, 'runs': [],
                    'host': {'system': platform.platform(), 'machine': platform.machine(),
                             'logical_cpus': os.cpu_count()},
                    'protocol': {'processes': 3, 'samples': 10, 'core_batch_target_ms': 10,
                                 'numeric_results': True, 'extended_batches': False}}
        if platform.system() == 'Darwin':
            for name in ['machdep.cpu.brand_string', 'hw.physicalcpu', 'hw.perflevel0.physicalcpu', 'hw.perflevel1.physicalcpu']:
                value = subprocess.run(['sysctl', '-n', name], capture_output=True, text=True)
                if value.returncode == 0:
                    manifest['host'][name] = value.stdout.strip()
        (directory / 'README.before.md').write_text((ROOT / 'README.md').read_text())
    def save():
        manifest_path.write_text(json.dumps(manifest, indent=2) + '\n')
    save()
    samples = {item['key']: [] for item in plan}
    for repetition in range(3):
        ordered = plan[repetition % len(plan):] + plan[:repetition % len(plan)]
        for item in ordered:
            if digest(Path(item['manifest'])) != item['manifest_sha256']:
                raise ValueError('Build manifest changed: ' + item['key'])
            for filename, expected in item['payload_sha256'].items():
                if digest(Path(filename)) != expected:
                    raise ValueError('Executable/runtime payload changed: ' + filename)
            if item.get('executable_sha256') and digest(Path(item['command'][0])) != item['executable_sha256']:
                raise ValueError('Numeric executable changed: ' + item['key'])
            key = item['key']
            stem = f'{repetition + 1}-{key}'
            output = directory / (stem + '.stdout')
            errors = directory / (stem + '.stderr')
            result_path = directory / (stem + '.json')
            if args.resume and result_path.exists():
                record = next((r for r in manifest['runs'] if r['key'] == key and r['process'] == repetition + 1), None)
                if not record or digest(output) != record['stdout_sha256']:
                    raise ValueError('Unverified saved process: ' + stem)
                checked = check(output.read_text(), item['mode'])
            else:
                if fingerprints() != current:
                    raise ValueError('Sources changed during the campaign; no publication')
                print(f'Measuring {key}: process {repetition + 1}/3', flush=True)
                with output.open('w') as out, errors.open('w') as err:
                    process = subprocess.run(item['command'], cwd=item.get('cwd', ROOT), stdout=out, stderr=err,
                                             timeout=args.timeout)
                if process.returncode:
                    raise RuntimeError(f'{key} failed ({process.returncode}); see {errors}')
                for filename, expected in item['payload_sha256'].items():
                    if digest(Path(filename)) != expected:
                        raise ValueError('Executable/runtime payload changed while running: ' + filename)
                checked = check(output.read_text(), item['mode'])
                result_path.write_text(json.dumps(checked, indent=2) + '\n')
                manifest['runs'].append({'key': key, 'process': repetition + 1,
                                         'stdout_sha256': digest(output), 'stderr_sha256': digest(errors),
                                         'result_sha256': digest(result_path), 'command': item['command']})
                save()
                print(f'{key}: {len(checked["values"])} outputs verified, {checked["total_ms"]:.6f} ms', flush=True)
            samples[key].append(checked)
    if fingerprints() != current:
        raise ValueError('Sources changed during the campaign; no publication')
    results = {key: aggregate(runs) for key, runs in samples.items()}
    (directory / 'results.json').write_text(json.dumps(results, indent=2) + '\n')
    manifest['results_sha256'] = digest(directory / 'results.json')
    manifest['completed_at_utc'] = datetime.now(timezone.utc).isoformat()
    save()
    print('Validated measurements: ' + str(directory / 'results.json'), flush=True)


if __name__ == '__main__':
    main()

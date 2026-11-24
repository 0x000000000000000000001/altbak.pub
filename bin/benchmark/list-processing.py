#!/usr/bin/env python3
"""Compare two frozen Go benchmark executables with the ordinary altbak harness."""
import argparse
from datetime import datetime, timezone
import hashlib
import json
import os
from pathlib import Path
import platform
import statistics
import subprocess
import sys

sys.dont_write_bytecode = True
from validate import CASES, validate_output


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def inspect(directory, test='ListOps'):
    manifest = json.loads((directory / 'manifest.json').read_text())
    if (manifest['language'], manifest['mode'], manifest['test']) != ('go', 'test', test):
        raise ValueError(f'Expected an isolated Go {test} build: {directory}')
    if manifest['profile']['backend_profile'] != 'release' or manifest['profile']['pgo']:
        raise ValueError('Both executables must use Release without PGO')
    binary = directory / 'benchmark'
    if sha(binary) != manifest['binary_sha256']:
        raise ValueError(f'Executable differs from its manifest: {binary}')
    source_key = f'src/Test/{test}/{test}.purs'
    source = directory / source_key
    if sha(source) != manifest['inputs'][source_key]:
        raise ValueError(f'{test} source differs from its manifest: {source}')
    return manifest


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--before', type=Path, required=True)
    parser.add_argument('--after', type=Path, required=True)
    parser.add_argument('--output', type=Path, required=True)
    parser.add_argument('--test', choices=CASES, default='ListOps',
                        help='Core benchmark to compare (default: ListOps)')
    parser.add_argument('--core', action='store_true',
                        help='Use core-benchmark executables built from the same generated Go with App as entry')
    args = parser.parse_args()
    directories = {'before': args.before.resolve(), 'after': args.after.resolve()}
    manifests = {name: inspect(directory, args.test) for name, directory in directories.items()}
    for key in ('profile', 'toolchain'):
        if manifests['before'][key] != manifests['after'][key]:
            raise ValueError(f'The configurations differ: {key}')
    sources = {name: {key: value for key, value in manifest['inputs'].items()
                     if key.startswith(('src/', 'srx/'))}
               for name, manifest in manifests.items()}
    if sources['before'] != sources['after']:
        raise ValueError('The benchmark sources differ')
    binaries = {name: directory / ('core-benchmark' if args.core else 'benchmark')
                for name, directory in directories.items()}
    hashes = {name: sha(binary) for name, binary in binaries.items()}
    args.output.mkdir(parents=True, exist_ok=False)
    env = dict(os.environ, GOGC=manifests['before']['profile']['GOGC'])
    env.pop('PPROF', None)
    samples = {'before': [], 'after': []}
    case_index = list(CASES).index(args.test) if args.core else 0
    order = ['before', 'after', 'after', 'before', 'before', 'after']
    for index, name in enumerate(order):
        raw = subprocess.check_output([str(binaries[name])],
                                      cwd=directories[name], env=env, text=True)
        (args.output / f'{index + 1}-{name}.log').write_text(raw)
        result = validate_output(raw, 'pure') if args.core else validate_output(raw, 'test', args.test)
        elapsed = result['times_us'][case_index]
        samples[name].append(elapsed)
        print(f'{name}: {elapsed:.6f} us', flush=True)
    medians = {name: statistics.median(values) for name, values in samples.items()}
    report = {
        'measured_at_utc': datetime.now(timezone.utc).isoformat(),
        'host': platform.platform(), 'machine': platform.machine(),
        'protocol': 'Three independent processes per variant, alternating AB/BA/AB; '
                    'each process: three local warmups, calibration, minimum of ten batches',
        'order': order, 'expected_result': int(CASES[args.test]),
        'scope': f'{args.test} in the {len(CASES)}-case core suite' if args.core else f'isolated {args.test}',
        'global_warmup_suites': 3 if args.core else 0,
        'executable_sha256': hashes,
        'GOGC': env['GOGC'], 'GOMAXPROCS': env.get('GOMAXPROCS', 'default'),
        'samples_us': samples, 'medians_us': medians,
        'speedup': medians['before'] / medians['after'],
        'reduction_percent': 100 * (1 - medians['after'] / medians['before']),
        'workspaces': {name: str(path) for name, path in directories.items()},
        'builds': manifests,
    }
    for directory in directories.values():
        inspect(directory, args.test)  # Reject a binary/source change during measurement.
    if {name: sha(binary) for name, binary in binaries.items()} != hashes:
        raise ValueError('An executable changed during measurement')
    (args.output / 'results.json').write_text(json.dumps(report, indent=2) + '\n')
    print(json.dumps({key: report[key] for key in ('medians_us', 'speedup', 'reduction_percent')}, indent=2))


if __name__ == '__main__':
    main()

#!/usr/bin/env python3
"""Measure frozen AppAll builds against the exact published JavaScript references.

Without --after, record a three-process baseline. With --after, compare the
two complete suites in AB/BA/AB order. Benchmark sources and harnesses are not
rewritten; this runner only validates and executes the frozen artifacts.
"""
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
from library_suite import library_cases
from measure import check
from publish import ROWS
from validate import CASES, validate_output


ROOT = Path(__file__).resolve().parents[2]
CATALOGUE = ROOT / 'bin/benchmark/ratio-measurements.json'
LISTED = (
    'RowToList', 'string/indexOf-hit-10000', 'State Monad',
    'record/set-large-1000', 'record/set-medium-1000', 'effect/ref-write-10000',
    'string/concat-1000', 'AST Evaluation', 'array/map-10000',
)


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def inspect(directory):
    path = directory / 'manifest.json'
    manifest = json.loads(path.read_text())
    expected = {'language': 'go', 'mode': 'pure', 'test': None,
                'entry': 'AppAll', 'include_library': True}
    if any(manifest.get(key) != value for key, value in expected.items()):
        raise ValueError(f'Expected a frozen Go AppAll build: {directory}')
    profile = manifest['profile']
    if profile['backend_profile'] != 'release' or profile['pgo']:
        raise ValueError('Only Release builds without PGO can be compared')
    if sha(directory / 'benchmark') != manifest['binary_sha256']:
        raise ValueError(f'Executable differs from its manifest: {directory}')
    # The native driver records extra diagnostic sources not copied into an
    # AppAll workspace. Verify every recorded live source, every copy present,
    # and explicitly require the complete 37-case suite and its harness.
    for name, digest in manifest['inputs'].items():
        if not name.startswith(('src/', 'srx/')):
            continue
        if sha(ROOT / name) != digest:
            raise ValueError(f'Benchmark source changed since build: {name}')
        copied = directory / name
        if copied.is_file() and sha(copied) != digest:
            raise ValueError(f'Workspace source differs from manifest: {copied}')
    required = ['src/App.purs', 'src/AppAll.purs', 'src/AppLibrary.purs',
                'src/Bench.purs', 'src/Bench.go']
    names = list(CASES) + [case['module'].removeprefix('Test.') for case in library_cases()]
    required += [f'src/Test/{name}/{name}.purs' for name in names]
    for name in required:
        if not (directory / name).is_file():
            raise ValueError(f'Missing suite input: {directory / name}')
        if sha(directory / name) != manifest['inputs'][name]:
            raise ValueError(f'Suite input differs from manifest: {name}')
    oracle = 'bin/benchmark/library-cases.json'
    if sha(ROOT / oracle) != manifest['inputs'][oracle]:
        raise ValueError('Library oracles changed since build')
    return manifest


def definitions():
    if len(CASES) != len(ROWS):
        raise ValueError('Core row names and oracle contract disagree')
    core = [{'name': label, 'module': 'Test.' + module, 'category': 'Stable',
             'expected': int(value)}
            for label, (module, value) in zip(ROWS, CASES.items())]
    library = [{'name': case['label'], 'module': case['module'],
                'category': 'Incl WIP', 'expected': int(case['value'])}
               for case in library_cases()]
    return core + library


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--before', type=Path, required=True)
    parser.add_argument('--after', type=Path, help='omit for a baseline-only campaign')
    parser.add_argument('--output', type=Path, required=True)
    parser.add_argument('--timeout', type=float, default=600)
    args = parser.parse_args()
    directories = {'before': args.before.resolve()}
    if args.after:
        directories['after'] = args.after.resolve()
        if directories['before'] == directories['after']:
            raise ValueError('Before and after must be distinct frozen builds')
    manifests = {name: inspect(directory) for name, directory in directories.items()}
    manifest_hashes = {name: sha(directory / 'manifest.json') for name, directory in directories.items()}
    if args.after:
        for key in ('profile', 'toolchain'):
            if manifests['before'][key] != manifests['after'][key]:
                raise ValueError(f'Build configurations differ: {key}')
        sources = {name: {key: value for key, value in manifest['inputs'].items()
                          if key.startswith(('src/', 'srx/'))}
                   for name, manifest in manifests.items()}
        if sources['before'] != sources['after']:
            raise ValueError('The benchmark sources differ between builds')
    catalogue_hash = sha(CATALOGUE)
    references = json.loads(CATALOGUE.read_text())['columns']['js-pure']
    cases = definitions()
    if len(cases) != 37 or not set(LISTED).issubset(case['name'] for case in cases):
        raise ValueError('The expected 14 core + 23 library suite changed')
    for case in cases:
        if references[case['name']]['time_us'] <= 0:
            raise ValueError('A positive exact JavaScript reference is required')
    args.output.mkdir(parents=True, exist_ok=False)
    env = dict(os.environ, GOGC=manifests['before']['profile']['GOGC'])
    env.pop('PPROF', None)
    order = ['before', 'after', 'after', 'before', 'before', 'after'] if args.after else ['before'] * 3
    samples = {name: [] for name in directories}
    raw_output, raw_errors, checked_runs = {}, {}, []
    for index, variant in enumerate(order, 1):
        directory = directories[variant]
        if sha(directory / 'manifest.json') != manifest_hashes[variant]:
            raise ValueError('Build manifest changed during measurement')
        if sha(directory / 'benchmark') != manifests[variant]['binary_sha256']:
            raise ValueError('Executable changed during measurement')
        process = subprocess.run([str(directory / 'benchmark')], cwd=directory,
                                 env=env, capture_output=True, text=True, timeout=args.timeout)
        stem = f'{index}-{variant}'
        (args.output / f'{stem}.log').write_text(process.stdout)
        (args.output / f'{stem}.stderr').write_text(process.stderr)
        process.check_returncode()
        # Require the appended library section, then also validate calibrated
        # power-of-two batch counts through the established campaign checker.
        validate_output(process.stdout, 'pure', include_library=True)
        checked = check(process.stdout, 'pure')
        times = checked['times_us'] + checked['library']['times_us']
        samples[variant].append(times)
        checked_runs.append({'variant': variant, 'result': checked})
        raw_output[stem] = process.stdout
        raw_errors[stem] = process.stderr
        row = dict(zip((case['name'] for case in cases), times))
        print(f'{stem}: 37 oracles verified; RowToList={row["RowToList"]:.6f} us; '
              f'indexOf-hit={row["string/indexOf-hit-10000"]:.6f} us', flush=True)
    all_tests = {}
    for index, case in enumerate(cases):
        name = case['name']
        reference = references[name]
        values = {variant: [run[index] for run in runs] for variant, runs in samples.items()}
        medians = {variant: statistics.median(runs) for variant, runs in values.items()}
        ratios = {variant: elapsed / reference['time_us'] for variant, elapsed in medians.items()}
        item = {**case, 'listed_js_target': name in LISTED,
                'js_reference': reference, 'target_us_strictly_below': 2 * reference['time_us'],
                'samples_us': values, 'medians_us': medians, 'ratios_js': ratios,
                'all_samples_below_2': {variant: max(runs) < 2 * reference['time_us']
                                        for variant, runs in values.items()}}
        if args.after:
            item['speedup'] = medians['before'] / medians['after']
            item['change_percent'] = 100 * (medians['after'] / medians['before'] - 1)
        all_tests[name] = item
    final = 'after' if args.after else 'before'
    listed = sorted(LISTED, key=lambda name: all_tests[name]['ratios_js'][final], reverse=True)
    remaining = [name for name in listed if all_tests[name]['ratios_js'][final] >= 2]
    report = {
        'measured_at_utc': datetime.now(timezone.utc).isoformat(),
        'host': platform.platform(), 'machine': platform.machine(),
        'mode': 'comparison' if args.after else 'baseline',
        'scope': {'core': 14, 'library': 23, 'listed_js_targets': 9},
        'protocol': 'Three independent processes per variant; AB/BA/AB when comparing; '
                    'three global warmups per suite, three local warmups, calibration, '
                    'minimum of ten batches; median per row of process minima.',
        'order': order, 'GOGC': env['GOGC'], 'GOMAXPROCS': env.get('GOMAXPROCS', 'default'),
        'js_reference_remeasured': False,
        'catalogue': {'path': str(CATALOGUE), 'sha256': catalogue_hash},
        'runner_sha256': sha(Path(__file__)),
        'workspaces': {name: str(path) for name, path in directories.items()},
        'manifest_sha256': manifest_hashes, 'builds': manifests,
        'executable_sha256': {name: manifest['binary_sha256'] for name, manifest in manifests.items()},
        'all_tests': all_tests, 'listed_tests_descending': listed,
        'remaining_listed_at_or_above_2': remaining,
        'remaining_all_at_or_above_2': sorted(
            (name for name, item in all_tests.items() if item['ratios_js'][final] >= 2),
            key=lambda name: all_tests[name]['ratios_js'][final], reverse=True),
        'goal_listed_under_2': not remaining,
        'checked_runs': checked_runs, 'raw_process_output': raw_output,
        'raw_process_stderr': raw_errors,
    }
    for variant, directory in directories.items():
        inspect(directory)
        if sha(directory / 'manifest.json') != manifest_hashes[variant]:
            raise ValueError('Build manifest changed during measurement')
    if sha(CATALOGUE) != catalogue_hash:
        raise ValueError('JavaScript references changed during measurement')
    (args.output / 'results.json').write_text(json.dumps(report, indent=2) + '\n')
    print(json.dumps({name: {'us': all_tests[name]['medians_us'][final],
                            'ratio_js': all_tests[name]['ratios_js'][final]}
                      for name in listed}, indent=2))


if __name__ == '__main__':
    try:
        main()
    except (ValueError, OSError, KeyError, subprocess.SubprocessError) as error:
        print(f'ERROR: {error}', file=sys.stderr)
        sys.exit(1)

#!/usr/bin/env python3
"""Compare two frozen Release JsonDecoding executables in AB/BA/AB order.

Compiler and library-port sources may differ between the two builds. Benchmark
sources, the diagnostic harness, fixtures and oracle must still match their
recorded hashes and the live repository. No executable is rebuilt by this tool.
"""
import argparse
from datetime import datetime, timezone
import hashlib
import importlib.util
import json
import math
from pathlib import Path
import platform
import statistics
import subprocess
import sys

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[2]
SUITE = 'JsonDecoding'
PHASES = ('parse', 'decode', 'combined')
ORDER = ('before', 'after', 'after', 'before', 'before', 'after')
REFERENCES_US = {'c': 657.69, 'js': 9283.94, 'go_published': 2237.333}


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def load_diagnostic():
    spec = importlib.util.spec_from_file_location(
        'json_diagnostic', Path(__file__).with_name('json-diagnostic.py'))
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module


def within(path, root):
    try:
        path.relative_to(root)
        return True
    except ValueError:
        return False


def implementation_source(path, diagnostic):
    """Only compiler/backend and port trees are allowed to differ or be stale."""
    if within(path, diagnostic.COMPILER.parent):
        first = path.relative_to(diagnostic.COMPILER.parent).parts[0]
        return first == diagnostic.COMPILER.name or first.startswith('gopurs-')
    return any(within(path, root) for root in (
        ROOT.parent / 'purescript', diagnostic.PBO, diagnostic.PBO_PURUST,
        ROOT.parent / 'purust',
    ))


def protected_sources(manifest, diagnostic):
    return {name: digest for name, digest in manifest['sources'].items()
            if not implementation_source(Path(name), diagnostic)}


def inspect(workspace, diagnostic):
    manifest_path = workspace / 'manifest.json'
    manifest = json.loads(manifest_path.read_text())
    diagnostic.check_suite(manifest, SUITE)
    if manifest.get('go_backend_profile') != 'release':
        raise ValueError(f'Expected an explicit Release build: {workspace}')
    if manifest.get('go_build_flags') != diagnostic.go_build_flags(False):
        raise ValueError(f'Unexpected Go build flags: {workspace}')
    if not isinstance(manifest.get('go'), str) or not manifest['go'].startswith('go version '):
        raise ValueError(f'Missing Go compiler version: {workspace}')
    if sha(workspace / 'benchmark') != manifest['binary_sha256']:
        raise ValueError(f'Executable differs from its manifest: {workspace}')
    protected = protected_sources(manifest, diagnostic)
    required = diagnostic.source_files(SUITE) + [
        Path(diagnostic.__file__).resolve(), ROOT / 'bin/benchmark/go_profile.py',
    ] + [path for path in diagnostic.fixtures(SUITE).iterdir() if path.is_file()]
    for path in required:
        if str(path) not in protected:
            raise ValueError(f'Manifest omits a benchmark, harness or fixture input: {path}')
    for name, digest in protected.items():
        if sha(Path(name)) != digest:
            raise ValueError(f'Benchmark/harness/fixture source changed since build: {name}')
    # The diagnostic driver copies purs/go/js into its native source tree,
    # and purs/js into its JS oracle tree. Rust FFI is not copied for Go.
    for source in diagnostic.source_files(SUITE):
        for prefix, suffixes in ((workspace, {'.purs', '.go', '.js'}),
                                 (workspace / 'js', {'.purs', '.js'})):
            if source.suffix not in suffixes:
                continue
            copied = prefix / source.relative_to(ROOT)
            if sha(copied) != protected[str(source)]:
                raise ValueError(f'Copied benchmark source differs from manifest: {copied}')
    for name, digest in manifest.get('generated_tast', {}).items():
        path = workspace / name
        if not within(path.resolve(), workspace.resolve()) or sha(path) != digest:
            raise ValueError(f'Generated TAST differs from manifest: {path}')
    return manifest


def validate_run(result, oracle, manifest, diagnostic):
    diagnostic.validate_result(result, oracle, SUITE)
    if result.get('backend') != 'go' or result.get('gomaxprocs') != 1:
        raise ValueError('Expected the Go diagnostic running with GOMAXPROCS=1')
    if result.get('go') != manifest['go'].split()[2]:
        raise ValueError('Reported Go runtime version differs from the build')
    if set(result.get('phases', {})) != set(PHASES) or 'profile_phase' in result:
        raise ValueError('Expected all three ordinary phases, without profiling')
    for phase in PHASES:
        data = result['phases'][phase]
        samples = data.get('samples', [])
        if len(samples) != 5:
            raise ValueError(f'Expected five measured samples after warmup: {phase}')
        for sample in samples:
            for key in ('time_us', 'allocated_bytes'):
                value = sample[key]
                if type(value) not in (int, float) or not math.isfinite(value) or value < 0:
                    raise ValueError(f'Invalid {phase} {key} sample')
            if sample['time_us'] == 0:
                raise ValueError(f'Nonpositive {phase} timing')
        if data['time_us'] != min(sample['time_us'] for sample in samples):
            raise ValueError(f'Published phase time is not the minimum of its five samples: {phase}')


def aggregate(runs):
    if len(runs) != 3:
        raise ValueError('Three independent processes are required per variant')
    return {
        'medians_us': {phase: statistics.median(run['phases'][phase]['time_us'] for run in runs)
                       for phase in PHASES},
        'allocated_bytes_per_corpus': {
            phase: statistics.median(sample['allocated_bytes'] for run in runs
                                     for sample in run['phases'][phase]['samples'])
            for phase in PHASES},
    }


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--before', type=Path, required=True)
    parser.add_argument('--after', type=Path, required=True)
    parser.add_argument('--output', type=Path, required=True)
    parser.add_argument('--timeout', type=float, default=600)
    args = parser.parse_args()
    diagnostic = load_diagnostic()
    workspaces = {'before': args.before.resolve(), 'after': args.after.resolve()}
    if workspaces['before'] == workspaces['after']:
        raise ValueError('Before and after must be distinct frozen workspaces')
    manifests = {variant: inspect(path, diagnostic) for variant, path in workspaces.items()}
    for key in ('go', 'go_backend_profile', 'go_build_flags'):
        if manifests['before'][key] != manifests['after'][key]:
            raise ValueError(f'Build configurations differ: {key}')
    protected = {variant: protected_sources(manifest, diagnostic)
                 for variant, manifest in manifests.items()}
    if protected['before'] != protected['after']:
        raise ValueError('Benchmark, harness or fixture sources differ between builds')
    manifest_hashes = {variant: sha(path / 'manifest.json') for variant, path in workspaces.items()}
    runner = Path(__file__).resolve()
    runner_hash, runner_source = sha(runner), runner.read_text()
    output = args.output.resolve()
    output.mkdir(parents=True, exist_ok=False)
    (output / 'runner.py').write_text(runner_source)
    for variant, workspace in workspaces.items():
        (output / f'{variant}.manifest.json').write_bytes((workspace / 'manifest.json').read_bytes())
    corpus, corpus_info = diagnostic.corpus_data(SUITE, 'fixture12', None, None)
    corpus_path = output / 'corpus.json'
    corpus_path.write_bytes(corpus)
    oracle_path = output / 'oracle.json'
    oracle_path.write_bytes((diagnostic.fixtures(SUITE) / 'expected.json').read_bytes())
    oracle = json.loads(oracle_path.read_text())
    corpus_hash, oracle_hash = sha(corpus_path), sha(oracle_path)
    env = diagnostic.environment()
    env['DIAG_CORPUS'] = str(corpus_path)
    results = {variant: [] for variant in workspaces}
    raw_output, raw_errors, raw_hashes, commands = {}, {}, {}, []
    for index, variant in enumerate(ORDER, 1):
        workspace = workspaces[variant]
        if sha(workspace / 'manifest.json') != manifest_hashes[variant]:
            raise ValueError('Build manifest changed during measurement')
        if sha(workspace / 'benchmark') != manifests[variant]['binary_sha256']:
            raise ValueError('Executable changed during measurement')
        if sha(corpus_path) != corpus_hash or sha(oracle_path) != oracle_hash:
            raise ValueError('Archived corpus or oracle changed during measurement')
        command = [str(workspace / 'benchmark')]
        process = subprocess.run(command, cwd=workspace, env=env, capture_output=True,
                                 text=True, timeout=args.timeout)
        stem = f'{index}-{variant}'
        (output / f'{stem}.log').write_text(process.stdout)
        (output / f'{stem}.stderr').write_text(process.stderr)
        raw_hashes[stem] = {'stdout': sha(output / f'{stem}.log'),
                            'stderr': sha(output / f'{stem}.stderr')}
        process.check_returncode()
        result = json.loads(process.stdout)
        validate_run(result, oracle, manifests[variant], diagnostic)
        results[variant].append(result)
        raw_output[stem], raw_errors[stem] = process.stdout, process.stderr
        commands.append({'variant': variant, 'command': command, 'cwd': str(workspace),
                         'environment': {key: env[key] for key in
                                         ('GOMAXPROCS', 'GOGC', 'GOWORK', 'DIAG_CORPUS')}})
        print(f'{stem}: {oracle["modules"]} oracles verified; '
              f'combined={result["phases"]["combined"]["time_us"]:.6f} us', flush=True)
    for variant, workspace in workspaces.items():
        inspect(workspace, diagnostic)
        if sha(workspace / 'manifest.json') != manifest_hashes[variant]:
            raise ValueError('Build manifest changed during measurement')
    if sha(corpus_path) != corpus_hash or sha(oracle_path) != oracle_hash or sha(runner) != runner_hash:
        raise ValueError('Corpus, oracle or comparison runner changed during measurement')
    summaries = {variant: aggregate(runs) for variant, runs in results.items()}
    medians = {variant: data['medians_us'] for variant, data in summaries.items()}
    ratios = {reference: {variant: values['combined'] / elapsed for variant, values in medians.items()}
              for reference, elapsed in REFERENCES_US.items()}
    all_sources = manifests['before']['sources'].keys() | manifests['after']['sources'].keys()
    implementation_delta = {name: {variant: manifest['sources'].get(name)
                                   for variant, manifest in manifests.items()}
                            for name in sorted(all_sources)
                            if manifests['before']['sources'].get(name) != manifests['after']['sources'].get(name)}
    report = {
        'suite': SUITE, 'measured_at_utc': datetime.now(timezone.utc).isoformat(),
        'host': platform.platform(), 'machine': platform.machine(),
        'corpus': {**corpus_info, 'kind': 'fixture12', 'path': str(corpus_path),
                   'bytes': len(corpus), 'sha256': corpus_hash},
        'oracle': {'path': str(oracle_path), 'sha256': oracle_hash, 'values': oracle},
        'protocol': {'order': list(ORDER), 'processes_per_variant': 3,
                     'warmups_per_phase': 2, 'samples_per_phase_per_process': 5,
                     'timing_cell': 'median of three process minimum times',
                     'allocation_cell': 'median of all fifteen samples per variant and phase',
                     'GOMAXPROCS': 1, 'GOGC': 100, 'GOWORK': 'off',
                     'file_io_timed': False, 'fingerprinting_timed': False,
                     'oracle_cases': oracle['modules'], 'timed_cases': oracle['timed_cases']},
        'references_us': REFERENCES_US,
        'references_source': 'Historical README JsonDecoding references; C and JS are not remeasured',
        'medians_us': medians,
        'allocated_bytes_per_corpus': {variant: data['allocated_bytes_per_corpus']
                                       for variant, data in summaries.items()},
        'ratios_combined': ratios,
        'speedup': {phase: medians['before'][phase] / medians['after'][phase] for phase in PHASES},
        'change_percent': {phase: 100 * (medians['after'][phase] / medians['before'][phase] - 1)
                           for phase in PHASES},
        'goal_combined_under_2c': ratios['c']['after'] < 2,
        'all_after_process_minima_under_2c': all(run['phases']['combined']['time_us'] < 2 * REFERENCES_US['c']
                                                for run in results['after']),
        'results': results, 'raw_process_output': raw_output, 'raw_process_stderr': raw_errors,
        'raw_process_sha256': raw_hashes,
        'invocation': {'argv': [sys.executable, *sys.argv], 'cwd': str(Path.cwd())},
        'commands': commands, 'builds': manifests, 'manifest_sha256': manifest_hashes,
        'workspaces': {variant: str(path) for variant, path in workspaces.items()},
        'protected_sources': protected['before'], 'implementation_source_delta': implementation_delta,
        'runner': {'path': str(runner), 'sha256': runner_hash, 'source': runner_source},
    }
    (output / 'results.json').write_text(json.dumps(report, indent=2) + '\n')
    print(json.dumps({'medians_us': medians, 'combined_ratios': ratios}, indent=2))


if __name__ == '__main__':
    try:
        main()
    except (ValueError, OSError, KeyError, subprocess.SubprocessError) as error:
        print(f'ERROR: {error}', file=sys.stderr)
        sys.exit(1)

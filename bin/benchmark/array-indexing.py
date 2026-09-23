#!/usr/bin/env python3
"""Build or measure the isolated PureScript array-indexing diagnostic (Go + JS)."""
import argparse
import hashlib
import json
import os
from pathlib import Path
import re
import shutil
import statistics
import subprocess
import sys

ROOT = Path(__file__).resolve().parents[2]
DRIVERS = ROOT / 'bin/benchmark/array-indexing'
KERNEL = ROOT / 'src/Test/ArrayIndexing.purs'
BACKEND = ROOT.parent / 'gopurs/gopurs/bin/gopurs'


def digest(path):
    return hashlib.sha256(Path(path).read_bytes()).hexdigest()


def fingerprints(paths):
    return {str(path.resolve()): digest(path) for path in sorted(set(paths)) if path.is_file()}


def write_json(path, value):
    path.write_text(json.dumps(value, indent=2) + '\n')


def run_logged(command, cwd, log, env):
    command = list(map(str, command))
    print(' '.join(command) if len(command) < 20 else f'{command[0]} {command[1]} ({len(command) - 6} modules; full command in manifest)', flush=True)
    with log.open('w') as output:
        result = subprocess.run(command, cwd=cwd, env=env, stdout=output, stderr=subprocess.STDOUT)
    if result.returncode:
        raise RuntimeError(f'Failed: {command}\n{log.read_text()[-6000:]}')
    return command


def dependency_sources():
    # Resolve only the imported modules from the already installed Go package
    # set. No registry update, root symlink mutation or compiler rebuild occurs.
    config = ROOT / 'run/bak/go/spago.go.yaml'
    candidates = list((ROOT / 'run/bak/go/spago/p').glob('*/src/**/*.purs'))
    for relative in re.findall(r'path: "([^"\n]+)"', config.read_text()):
        candidates.extend((ROOT / relative / 'src').glob('**/*.purs'))
    candidates.append(KERNEL)
    modules = {}
    for path in candidates:
        match = re.search(r'^module\s+([A-Z][\w.]*)', path.read_text(), re.MULTILINE)
        if match:
            modules[match[1]] = path.resolve()
    selected, pending = {}, ['Test.ArrayIndexing']
    while pending:
        name = pending.pop()
        if name in selected or name == 'Prim' or name.startswith('Prim.'):
            continue
        if name not in modules:
            raise ValueError(f'Missing installed module {name}; prepare the normal Go dependency set first')
        path = modules[name]
        selected[name] = path
        pending.extend(re.findall(r'^import\s+([A-Z][\w.]*)', path.read_text(), re.MULTILINE))
    return sorted(selected.values())


def audit_generated(path):
    text = path.read_text()
    native = text.split('func Call_Test_ArrayIndexing_nativeReads(', 1)[1].split('\nfunc ', 1)[0]
    boxed = text.split('func Call_Test_ArrayIndexing_boxedReads(', 1)[1].split('\nfunc ', 1)[0]
    if '[]int64' not in native or 'arrayUnsafe_source_' not in native:
        raise ValueError('Native indexing no longer has the audited representation')
    if 'gopurs_runtime.Value' not in boxed or '(*[]gopurs_runtime.Value)' not in boxed:
        raise ValueError('Boxed indexing no longer has the audited representation')
    for name, body in [('native', native), ('boxed', boxed)]:
        if 'make(' in body or 'CoerceTo' in body or 'arrayUnsafe_index_' not in body or 'for {' not in body:
            raise ValueError(f'{name} generated loop needs review: conversion/copy or indexing changed')
    if not re.search(r'\[arrayUnsafe_index_\d+\]\.IntVal', boxed):
        raise ValueError('Boxed source must extract the selected Int after indexing')
    return {'native': '[]int64 indexed inside generated loop',
            'boxed': 'Value -> *[]Value -> selected element.IntVal inside generated loop',
            'whole_array_copy_in_kernel': False, 'sha256': digest(path)}


def build(directory, env):
    directory.mkdir(parents=True, exist_ok=True)
    marker = directory / '.array-indexing-workspace'
    if any(directory.iterdir()) and not marker.exists():
        raise ValueError('Nonempty build directory has no array-indexing workspace marker')
    if marker.exists() and marker.read_text() != str(ROOT):
        raise ValueError('Build directory belongs to another checkout')
    marker.write_text(str(ROOT))
    (directory / 'manifest.json').unlink(missing_ok=True)
    logs = directory / 'logs'
    logs.mkdir(exist_ok=True)
    deps = dependency_sources()
    purs = shutil.which('purs', path=env['PATH'])
    if not purs:
        raise ValueError('The TAST PureScript fork must be available on PATH')
    source_paths = [Path(__file__), ROOT / 'run/bak/go/spago.go.yaml', *DRIVERS.glob('*'), *deps]
    for path in deps:
        source_paths.extend([path.with_suffix('.js'), path.with_suffix('.go')])
    source_paths += [Path(purs), BACKEND, BACKEND.with_suffix('.js')]
    before = fingerprints(source_paths)
    commands = []
    commands.append(run_logged([purs, 'compile', *deps, '--codegen', 'corefn,js', '--output', directory / 'output'],
                               ROOT, logs / 'purs.log', env))
    tast = json.loads((directory / 'output/Test.ArrayIndexing/corefn.json').read_text())
    if 'dataDecls' not in tast or 'classDecls' not in tast:
        raise ValueError('The selected purs does not produce the required TAST; use the local fork')
    commands.append(run_logged([BACKEND, '--main', 'Test.ArrayIndexing'], directory, logs / 'gopurs.log', env))
    generated = directory / 'output/purescript/Test_ArrayIndexing.go'
    audit = audit_generated(generated)
    driver = directory / 'output/array-driver'
    driver.mkdir(exist_ok=True)
    shutil.copy2(DRIVERS / 'driver.go', driver / 'main.go')
    shutil.copy2(DRIVERS / 'driver.mjs', directory / 'driver.mjs')
    commands.append(run_logged(['go', 'build', '-pgo=off', '-o', directory / 'array-benchmark', './array-driver'],
                               directory / 'output', logs / 'go-build.log', env))
    compiler = env.get('CC', 'clang')
    commands.append(run_logged([compiler, '-O3', '-o', directory / 'array-benchmark-c', DRIVERS / 'driver.c'],
                               directory, logs / 'clang.log', env))
    if fingerprints(source_paths) != before:
        raise ValueError('Inputs changed during build')
    artifacts = [directory / 'array-benchmark', directory / 'array-benchmark-c', directory / 'driver.mjs', generated,
                 *sorted((directory / 'output').rglob('*.js'))]
    versions = {name: subprocess.check_output(command, env=env, text=True).strip() for name, command in {
        'go': ['go', 'version'], 'node': ['node', '--version'], 'purs': [purs, '--version'],
        'clang': [compiler, '--version']}.items()}
    manifest = {'schema': 1, 'sources': before, 'artifacts': fingerprints(artifacts),
                'commands': commands, 'versions': versions, 'generated_audit': audit,
                'profile': {key: env.get(key, '') for key in PROFILE_KEYS}}
    write_json(directory / 'manifest.json', manifest)
    print(f'Built and audited {directory}; run with --run-only after all compilation has stopped.')


def oracle(size, count, seed):
    values = [(i * 17 + seed * 31) % 251 + 1 for i in range(size)]
    start = seed % size
    return sum(values) * (count // size) + sum(values[(start + i) % size] for i in range(count % size))


def validate_rows(text, runtime, accesses, batches, seed):
    rows = [json.loads(line) for line in text.splitlines()]
    expected = {(mode, size) for mode in ['native', 'boxed'] for size in [16, 1024, 16384]}
    if len(rows) != 6 or {(row['representation'], row['size']) for row in rows} != expected:
        raise ValueError('Incomplete or duplicate diagnostic cases')
    for row in rows:
        if any(row[key] != value for key, value in {'runtime': runtime, 'accesses': accesses,
               'seed': seed, 'warmups': 3, 'checksum': oracle(row['size'], accesses, seed)}.items()):
            raise ValueError(f'Invalid output metadata/checksum: {row}')
        if len(row['ns_per_access']) != batches or any(not 0 < n < 1e9 for n in row['ns_per_access']):
            raise ValueError('Invalid timing batches')
        if runtime == 'go':
            values = row['bytes_per_access']
            if len(values) != batches or any(not 0 <= n <= 256 for n in values):
                raise ValueError('Go allocation regression')
            if not 0 <= row['probe_bytes_per_access'] <= 256:
                raise ValueError('Go allocation preflight regression')
        elif row['bytes_per_access'] is not None:
            raise ValueError('JS allocation data is not available')
    if runtime == 'go':
        for mode in ['native', 'boxed']:
            costs = [r['probe_bytes_per_access'] for r in rows if r['representation'] == mode]
            if max(costs) - min(costs) > 128:
                raise ValueError('Allocation bytes grow with input array length')
    return rows


def measure(directory, output, args, env):
    manifest = json.loads((directory / 'manifest.json').read_text())
    def check_frozen():
        for section in ['sources', 'artifacts']:
            for path, expected in manifest[section].items():
                if digest(path) != expected:
                    raise ValueError(f'Stale build: {path}')
    check_frozen()
    if manifest['profile'] != {key: env.get(key, '') for key in PROFILE_KEYS}:
        raise ValueError('Runtime/compiler profile changed since build')
    output.mkdir(parents=True, exist_ok=False)
    write_json(output / 'build-manifest.json', manifest)
    shutil.copy2(directory / 'output/purescript/Test_ArrayIndexing.go', output / 'ArrayIndexing.generated.go')
    runtimes = [args.runtime] if args.runtime else ['go', 'js', 'c']
    processes = []
    for process, seed in enumerate([5, 13, 29], 1):
        for runtime in runtimes:
            command = [directory / 'array-benchmark', '-accesses', args.accesses, '-batches', args.batches, '-seed', seed] if runtime == 'go' else [
                directory / 'array-benchmark-c', '-accesses', args.accesses, '-batches', args.batches, '-seed', seed] if runtime == 'c' else [
                'node', directory / 'driver.mjs', args.accesses, args.batches, seed]
            command = list(map(str, command))
            result = subprocess.run(command, cwd=directory, env=env, text=True, capture_output=True, timeout=120)
            (output / f'{runtime}-{process}.stdout.jsonl').write_text(result.stdout)
            (output / f'{runtime}-{process}.stderr.log').write_text(result.stderr)
            if result.returncode:
                raise ValueError(f'{runtime} process {process} failed: {result.stderr}')
            rows = validate_rows(result.stdout, runtime, args.accesses, args.batches, seed)
            processes.append({'runtime': runtime, 'process': process, 'seed': seed, 'command': command, 'cases': rows})
            print(f'Validated {runtime} process {process}: six checksums and {6 * args.batches} batches', flush=True)
    check_frozen()
    summary = []
    for runtime in runtimes:
        for mode in ['native', 'boxed']:
            for size in [16, 1024, 16384]:
                cases = [r for p in processes if p['runtime'] == runtime for r in p['cases']
                         if r['representation'] == mode and r['size'] == size]
                minima = [min(r['ns_per_access']) for r in cases]
                allocated = [b for r in cases for b in (r['bytes_per_access'] or [])]
                summary.append({'runtime': runtime, 'representation': mode, 'size': size,
                    'ns_per_access': statistics.median(minima), 'process_minima_ns': minima,
                    'bytes_per_access': statistics.median(allocated) if allocated else None,
                    'max_bytes_per_access': max(allocated) if allocated else None})
    write_json(output / 'results.json', {'schema': 1, 'diagnostic': 'array-indexing',
        'aggregation': 'median of three independent process minima; ten batches by default',
        'allocation_boundary': 'TotalAlloc around each generated PureScript loop invocation; inputs excluded',
        'runtimes': runtimes, 'build': manifest, 'processes': processes, 'summary': summary})
    print(json.dumps(summary, indent=2))


PROFILE_KEYS = ['GOGC', 'GOMAXPROCS', 'GOWORK', 'GOFLAGS', 'GOEXPERIMENT', 'GOPURS_JS', 'NODE_OPTIONS']


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    action = parser.add_mutually_exclusive_group(required=True)
    action.add_argument('--build-only', action='store_true')
    action.add_argument('--run-only', action='store_true')
    parser.add_argument('--runtime', choices=['go', 'js', 'c'], help='measure only this runtime; a build prepares all of them')
    parser.add_argument('--build-dir', type=Path)
    parser.add_argument('--output', type=Path, help='new campaign directory, required with --run-only')
    parser.add_argument('--accesses', type=int, default=1 << 23)
    parser.add_argument('--batches', type=int, default=10)
    args = parser.parse_args()
    if not 16384 <= args.accesses <= 1 << 23 or args.batches < 3:
        parser.error('accesses must be 16384..8388608, batches must be at least 3')
    if args.run_only and not args.output:
        parser.error('--run-only requires --output')
    args.build_dir = args.build_dir or ROOT / f'run/bak/{args.runtime or "go"}/modes/test-ArrayIndexing'
    directory = args.build_dir.resolve()
    if directory == ROOT or directory in ROOT.parents or args.build_dir.is_symlink():
        parser.error('unsafe build directory')
    env = dict(os.environ, GOGC='800', GOMAXPROCS='1', GOWORK='off', GOPURS_JS='1')
    env['PATH'] = str(Path.home() / '.local/bin') + os.pathsep + env['PATH']
    if args.build_only:
        build(directory, env)
    else:
        measure(directory, args.output.resolve(), args, env)


if __name__ == '__main__':
    try:
        main()
    except (OSError, ValueError, RuntimeError, KeyError, IndexError, subprocess.SubprocessError) as error:
        print(f'ERROR: {error}', file=sys.stderr)
        sys.exit(1)

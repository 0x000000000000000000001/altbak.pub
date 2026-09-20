"""Build isolated, fresh C++11/O3 runners and validate their observable results."""
import argparse
import hashlib
import json
import math
import os
from pathlib import Path
import re
import shutil
import subprocess
import sys
import time

sys.dont_write_bytecode = True
HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
CASES = dict(zip(
    ['AstTree', 'Fib', 'ListOps', 'TCO', 'Records', 'Ackermann', 'Church',
     'Primes', 'RBTree', 'Polymorphism', 'StateMonad', 'LazyEvaluation', 'ArrayOps', 'RowToList'],
    ['7', '55', '202950', '100000', '20000', '125', '100000',
     '21536', '22', '10000000', '1200', '1000000', '202950', '5']))
PROFILE = ['-std=c++11', '-O3', '-DNDEBUG']
ROW = re.compile(r'\(Test\)\s+([^\n]+)\s+\(Output & Warm-up\)\s+(\S+)\s+'
                 r'\(Execution time - best of 10\)\s+([0-9.]+) μs')


def fail(message):
    raise RuntimeError(message)


def digest(path):
    return hashlib.sha256(Path(path).read_bytes()).hexdigest()


def write_json(path, value):
    Path(path).write_text(json.dumps(value, indent=2) + '\n')


def logged(command, directory, log):
    command = list(map(str, command))
    print(f'C++: {log.name}', flush=True)
    with log.open('w') as output:
        output.write(json.dumps(command) + '\n')
        output.flush()
        result = subprocess.run(command, cwd=directory, stdout=output, stderr=subprocess.STDOUT)
    if result.returncode:
        print('\n'.join(log.read_text(errors='replace').splitlines()[-30:]), file=sys.stderr)
        fail(f'Command failed ({result.returncode}); see {log}')
    return command


def version(executable):
    return subprocess.check_output([str(executable), '--version'], text=True,
                                   stderr=subprocess.STDOUT).strip()


def expectation(module, override=None):
    source = ROOT / 'src' / (module.replace('.', '/') + '.purs')
    if not source.is_file():
        fail(f'No source module: {source}')
    label = re.search(r'(?m)^describe\s*=\s*log\s+"([^"\n]+)"', source.read_text())
    name = module.removeprefix('Test.').removesuffix('FFICheatcode').removesuffix('FFI')
    expected = override if override is not None else CASES.get(name)
    if expected is None or label is None:
        fail(f'{module} needs a literal describe label and --expected VALUE')
    return {'module': module, 'label': label.group(1), 'value': expected}


def select(args):
    if args.test:
        module = args.test if args.test.startswith('Test.') else 'Test.' + args.test
        if not re.fullmatch(r'Test\.[A-Z][A-Za-z0-9_.]*', module):
            fail('Invalid --test module name')
        if args.expected and len(args.expected) != 1:
            fail('--test accepts exactly one --expected value')
        return 'test-' + module.replace('.', '-'), None, [expectation(module, args.expected[0] if args.expected else None)]
    if args.x:
        if not (ROOT / 'src/AppX.purs').is_file():
            fail('--x requires an existing src/AppX.purs; it is never overwritten')
        if not args.expected:
            fail('--x requires one --expected VALUE per displayed benchmark result')
        if args.smoke:
            fail('--smoke needs the standard suite or --test; --x is an arbitrary application')
        return 'x', 'AppX', [{'module': None, 'label': None, 'value': x} for x in args.expected]
    if args.expected:
        fail('--expected is only for --test or --x')
    mode, suffix, entry = ('ffi', 'FFI', 'AppCppFFI') if args.ffi else (
        ('fficc', 'FFICheatcode', 'AppCppFFICheatcode') if args.fficc else ('pure', '', 'App'))
    return mode, entry, [expectation('Test.' + name + suffix) for name in CASES]


def fresh_workspace(directory):
    directory = directory.resolve()
    if directory == ROOT or directory in ROOT.parents or directory.is_symlink():
        fail(f'Refusing workspace location: {directory}')
    marker = directory / '.altbak-cpp-build'
    if directory.exists() and any(directory.iterdir()) and not marker.is_file():
        fail(f'Nonempty directory is not an isolated C++ workspace: {directory}')
    directory.mkdir(parents=True, exist_ok=True)
    if marker.exists() and marker.read_text() != str(ROOT) + '\n':
        fail('Workspace belongs to a different checkout')
    marker.write_text(str(ROOT) + '\n')
    # Only our generated artifacts are replaced; dependency caches stay usable.
    for name in ['output', 'ffi', 'sources']:
        child = directory / name
        if child.is_symlink():
            fail(f'Refusing generated-directory symlink: {child}')
        if child.exists():
            shutil.rmtree(child)
    for name in ['Makefile', 'build.json']:
        (directory / name).unlink(missing_ok=True)
    return directory


def main_source(entry, cases):
    lines = ['module Main where', 'import Prelude', 'import Effect (Effect)',
             'import Effect.Console (log)', 'import Bench (runBench)']
    if entry:
        lines.append(f'import {entry} as Entry')
    if entry == 'AppX':
        return '\n'.join(lines + ['main :: Effect Unit', 'main = Entry.main', ''])
    lines += [f'import {case["module"]} as Case{i}' for i, case in enumerate(cases)]
    lines += ['foreign import smokeRequested :: Effect Boolean', 'main :: Effect Unit',
              'main = do', '  smoke <- smokeRequested', '  if smoke then do']
    for i, case in enumerate(cases):
        lines += [f'    result{i} <- Case{i}.act',
                  f'    log ("CPP_RESULT {case["module"]} " <> show result{i})']
    lines.append('  else ' + ('Entry.main' if entry else 'void (runBench Case0.describe Case0.act)'))
    return '\n'.join(lines) + '\n'


def reachable(output):
    result = {}
    pending = ['Main']
    while pending:
        name = pending.pop()
        if name in result or name == 'Prim' or name.startswith('Prim.'):
            continue
        path = output / name / 'corefn.json'
        if not path.is_file():
            fail(f'Missing imported CoreFn module {name}')
        module = json.loads(path.read_text())
        result[name] = path
        pending.extend('.'.join(item['moduleName']) for item in module.get('imports', []))
    return result


def install_ffi(directory, selected, ffi_root, entry):
    namespaces = {name.replace('.', '_') for name in selected}
    providers = {}
    candidates = sorted(p for p in ffi_root.rglob('*') if p.suffix in {'.cpp', '.cc', '.c', '.mm'})
    candidates += sorted((ROOT / 'src/Test').glob('*.cc'))
    candidates += [HERE / 'bench_ffi.cc']
    if entry != 'AppX':
        candidates += [HERE / 'main_ffi.cc']
    for source in candidates:
        text = source.read_text()
        supplied = set(re.findall(r'FOREIGN_BEGIN\s*\(\s*(\w+)\s*\)', text)) & namespaces
        if not supplied:
            continue
        for name in supplied:
            if name in providers:
                fail(f'Duplicate FFI provider for {name}: {source} and {providers[name]}')
            providers[name] = str(source)
        if ffi_root in source.parents:
            target = directory / 'ffi/standard' / source.relative_to(ffi_root)
        elif ROOT / 'src/Test' in source.parents:
            target = directory / 'ffi/Test' / source.name
        else:
            target = directory / 'ffi/BenchRunner' / source.name
        target.parent.mkdir(parents=True, exist_ok=True)
        if 'Data_Array' in supplied and 'exports["rangeImpl"]' not in text:
            text = '#include <functional>\n' + text
            text = text.replace('FOREIGN_END', (HERE / 'array_compat.inc').read_text() + '\nFOREIGN_END')
        target.write_text(text)
    # Preserve provider-relative includes such as strings/utf8.h.
    for source in ffi_root.rglob('*'):
        if source.suffix in {'.h', '.hpp', '.inc'}:
            target = directory / 'ffi/standard' / source.relative_to(ffi_root)
            target.parent.mkdir(parents=True, exist_ok=True)
            shutil.copy2(source, target)
    missing = {}
    for name, source in selected.items():
        foreigns = json.loads(source.read_text()).get('foreign')
        if foreigns and name.replace('.', '_') not in providers:
            if name.startswith('Test.') or name in {'Main', 'Bench'}:
                fail(f'No real FFI provider for {name}; no stubs are generated')
            missing[name] = foreigns
    if missing:
        print('C++: imported library namespaces without providers (no stubs; an executed lookup must fail):\n  '
              + '\n  '.join(sorted(missing)), flush=True)
    return providers, missing


def build(args, mode, entry, cases, directory):
    directory = fresh_workspace(directory)
    shutil.copy2(HERE / 'spago.yaml', directory / 'spago.yaml')
    commands = []
    dependency_root = args.dependency_root or ROOT / 'run/bak/cpp/spago/p'
    if not dependency_root.is_dir():
        if args.dependency_root:
            fail(f'Missing dependency root: {dependency_root}')
        commands.append(logged([args.spago, 'fetch'], directory, directory / 'fetch.log'))
        dependency_root = directory / '.spago/p'
    project = directory / 'sources/project'
    project.mkdir(parents=True)
    for source in sorted((ROOT / 'src').rglob('*.purs')):
        if source.name == 'Main.purs' or (source.name == 'AppX.purs' and entry != 'AppX'):
            continue
        target = project / source.relative_to(ROOT / 'src')
        target.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(source, target)
    main = directory / 'sources/Main.purs'
    main.write_text(main_source(entry, cases))
    dependency_sources = sorted(dependency_root.glob('*/src/**/*.purs'))
    if not dependency_sources:
        fail(f'No PureScript dependencies under {dependency_root}')
    dependencies = []
    for source in dependency_sources:
        target = directory / 'sources/dependencies' / source.relative_to(dependency_root)
        target.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(source, target)
        dependencies.append(target)
    inputs = [main] + sorted(project.rglob('*.purs')) + dependencies
    purs_version = version(args.purs)
    if purs_version != '0.14.4':
        fail(f'pscpp requires the existing PureScript 0.14.4 frontend, got {purs_version}')
    commands.append(logged([args.purs, 'compile', '-g', 'corefn', '-o', directory / 'output', *inputs],
                           directory, directory / 'corefn.log'))
    selected = reachable(directory / 'output')
    commands.append(logged([args.pscpp, '--makefile', *[selected[name] for name in sorted(selected)]],
                           directory, directory / 'pscpp.log'))
    providers, missing = install_ffi(directory, selected, args.ffi_root, entry)
    makefile = directory / 'Makefile'
    # Record the actual compiler invocations, not only "Creating ..." messages.
    makefile.write_text(makefile.read_text().replace('@$(CXX)', '$(CXX)').replace('@$(CC)', '$(CC)'))
    commands.append(logged(['make', '-j' + str(args.jobs), 'main', 'CXX=' + str(args.cxx),
                           'CXXFLAGS=' + ' '.join(PROFILE), 'CFLAGS=-O3 -DNDEBUG'],
                          directory, directory / 'make.log'))
    binary = directory / 'output/bin/main'
    if not binary.is_file():
        fail(f'Build did not produce {binary}')
    write_json(directory / 'build.json', {
        'mode': mode, 'entry': entry or cases[0]['module'], 'profile': PROFILE,
        'fresh_objects': True, 'stubs': False, 'timer_unit': 'microseconds',
        'timer_clock': 'steady_clock, relative origin', 'validation': 'stdout after timing; separate untimed smoke path',
        'binary': str(binary), 'binary_sha256': digest(binary), 'expected': cases,
        'purs_version': purs_version, 'cxx_version': version(args.cxx), 'commands': commands,
        'corefn_modules': sorted(selected), 'ffi_providers': providers,
        'unprovided_foreign_namespaces': missing,
        'dependency_source': str(dependency_root.resolve()),
        'input_hashes': {str(path.relative_to(directory)): digest(path) for path in inputs},
        'native_source_hashes': {str(path.relative_to(directory)): digest(path)
                                 for base in ['ffi', 'output/src'] for path in (directory / base).rglob('*')
                                 if path.suffix in {'.cc', '.cpp', '.h', '.c', '.mm'}},
        'harness_hashes': {p.name: digest(p) for p in HERE.iterdir() if p.is_file()
                           and p.name in ['run', 'driver.py', 'bench_ffi.cc', 'main_ffi.cc', 'array_compat.inc', 'spago.yaml']},
    })
    print(f'Binary: {binary}\nManifest: {directory / "build.json"}', flush=True)


def validate_output(output, expected, smoke, require_total=True):
    if smoke:
        rows = re.findall(r'(?m)^CPP_RESULT (\S+) (\S+)\s*$', output)
        if len(rows) != len(expected) or len(re.findall(r'(?m)^CPP_RESULT\b', output)) != len(expected):
            fail(f'Expected {len(expected)} smoke values, found {len(rows)}')
        for row, case in zip(rows, expected):
            if row != (case['module'], case['value']):
                fail(f'Unexpected smoke result: {row}; expected {case}')
        return {'values': [value for _, value in rows], 'smoke': True}
    rows = ROW.findall(output)
    if (len(rows) != len(expected) or output.count('(Output & Warm-up)') != len(expected)
            or output.count('(Execution time - best of 10)') != len(expected)):
        fail(f'Expected exactly {len(expected)} complete results; found {len(rows)}')
    for (label, value, elapsed), case in zip(rows, expected):
        if value != case['value'] or (case['label'] is not None and label.strip() != case['label']):
            fail(f'Unexpected benchmark result: {(label, value)}; expected {case}')
        if not math.isfinite(float(elapsed)) or float(elapsed) < 0:
            fail('Invalid measured duration')
    totals = re.findall(r'Total exec time: ([0-9.]+) ms', output)
    if require_total and len(totals) != 1:
        fail(f'Expected exactly one total, found {len(totals)}')
    if not require_total and len(totals) > 1:
        fail('Unexpected repeated total')
    summed_ms = sum(float(row[2]) for row in rows) / 1000
    # Both each µs line and the ms total are rounded to two decimal places.
    tolerance = 0.005 + len(rows) * 0.005 / 1000 + 1e-9
    if totals and (not math.isfinite(float(totals[0])) or abs(float(totals[0]) - summed_ms) > tolerance):
        fail(f'Total {totals[0]} ms disagrees with line sum {summed_ms} ms')
    return {'values': [row[1] for row in rows], 'times_us': [float(row[2]) for row in rows],
            'total_ms': float(totals[0]) if totals else None, 'sum_displayed_lines_ms': summed_ms,
            'rounding_tolerance_ms': tolerance, 'smoke': False}


def execute(args, mode, cases, directory):
    metadata = json.loads((directory / 'build.json').read_text())
    if metadata['mode'] != mode:
        fail(f'Build mode is {metadata["mode"]}, requested {mode}')
    if metadata['expected'] != cases:
        fail('Requested expected results differ from the compiled manifest; rebuild or use its original expectations')
    binary = Path(metadata['binary'])
    if digest(binary) != metadata['binary_sha256']:
        fail('Compiled binary changed since build')
    env = os.environ.copy()
    env['ALTBAK_CPP_SMOKE'] = '1' if args.smoke else '0'
    log = directory / (('smoke-' if args.smoke else 'run-') + str(time.time_ns()) + '.log')
    output = []
    with log.open('w') as stream:
        process = subprocess.Popen([str(binary)], cwd=directory, env=env, stdout=subprocess.PIPE,
                                   stderr=subprocess.STDOUT, text=True)
        for line in process.stdout:
            output.append(line)
            stream.write(line)
            stream.flush()
            print(line, end='', flush=True)
        status = process.wait()
    if status:
        fail(f'Runner exited {status}; see {log}')
    result = validate_output(''.join(output), metadata['expected'], args.smoke,
                             require_total=not mode.startswith('test-') and mode != 'x')
    result.update({'binary_sha256': metadata['binary_sha256'], 'build_sha256': digest(directory / 'build.json'),
                   'log': str(log), 'mode': mode})
    write_json(log.with_suffix('.json'), result)
    print(f'C++: {len(result["values"])} expected results verified; {log.with_suffix(".json")}', flush=True)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    mode = parser.add_mutually_exclusive_group()
    mode.add_argument('--ffi', action='store_true')
    mode.add_argument('--fficc', action='store_true')
    mode.add_argument('--test', metavar='MODULE')
    mode.add_argument('--x', action='store_true')
    execution = parser.add_mutually_exclusive_group()
    execution.add_argument('--build-only', action='store_true')
    execution.add_argument('--run-only', action='store_true')
    parser.add_argument('--smoke', action='store_true', help='one untimed act per kernel, validate every value')
    parser.add_argument('--build-dir', type=Path, help='isolated workspace; default run/bak/cpp/modes/MODE')
    parser.add_argument('--expected', action='append', help='expected value for a custom --test or --x')
    parser.add_argument('-c', '--clean', action='store_true', help='compatibility flag: native objects are always fresh')
    parser.add_argument('--jobs', type=int, default=min(8, os.cpu_count() or 1))
    parser.add_argument('--purs', type=Path, default=ROOT.parent / 'node_modules/.bin/purs')
    parser.add_argument('--pscpp', type=Path, default=ROOT.parent / 'pscpp-bin/pscpp')
    parser.add_argument('--cxx', default=shutil.which('clang++') or 'c++')
    parser.add_argument('--spago', default=shutil.which('spago') or 'spago')
    parser.add_argument('--ffi-root', type=Path, default=ROOT.parent / 'cpp-ffi')
    parser.add_argument('--dependency-root', type=Path,
                        help='cached Spago p directory to snapshot; existing C++ cache is preferred')
    args = parser.parse_args()
    if args.jobs < 1 or (args.build_only and args.smoke):
        parser.error('jobs must be positive; --smoke cannot be combined with --build-only')
    selected_mode, entry, cases = select(args)
    directory = (args.build_dir or ROOT / 'run/bak/cpp/modes' / selected_mode).resolve()
    if not args.run_only:
        if not args.ffi_root.is_dir():
            fail('Missing real cpp-ffi checkout; pass --ffi-root PATH')
        build(args, selected_mode, entry, cases, directory)
    if not args.build_only:
        execute(args, selected_mode, cases, directory)


if __name__ == '__main__':
    try:
        main()
    except (RuntimeError, OSError, subprocess.SubprocessError, ValueError) as error:
        print('C++ runner error: ' + str(error), file=sys.stderr)
        sys.exit(1)

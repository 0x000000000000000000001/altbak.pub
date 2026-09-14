#!/usr/bin/env python3
"""Build stock-CoreFn psgo in isolation; reject wrong or incomplete results."""
import argparse
import hashlib
import importlib.util
import json
import os
from pathlib import Path
import shutil
import subprocess
import sys
import tempfile

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[2]


def run(command, work, log, env):
    print(' '.join(map(str, command)), flush=True)
    with (work / log).open('w') as out:
        result = subprocess.run(list(map(str, command)), cwd=work, env=env, stdout=out, stderr=subprocess.STDOUT)
    if result.returncode:
        raise RuntimeError((work / log).read_text()[-10000:] + f'\nSee {work / log}')


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--test')
    parser.add_argument('--expected')
    parser.add_argument('--artifact', type=Path, help='Frozen build directory; requires --run-only')
    parser.add_argument('-c', '--clean', action='store_true', help='Every build already uses a fresh directory')
    action = parser.add_mutually_exclusive_group()
    action.add_argument('--build-only', action='store_true')
    action.add_argument('--run-only', action='store_true')
    args = parser.parse_args()
    if args.artifact and not args.run_only:
        parser.error('--artifact requires --run-only')
    if args.run_only and args.clean:
        parser.error('--run-only conflicts with --clean')
    mode = 'test' if args.test else 'pure'
    spec = importlib.util.spec_from_file_location('validator', ROOT / 'bin/benchmark/validate.py')
    validator = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(validator)
    validator.expected_cases(mode, args.test, args.expected)
    base = ROOT / 'run/bak/psgo'
    store = base / 'output-runs'
    store.mkdir(parents=True, exist_ok=True)
    pointer = store / ('latest-' + mode + ('-' + args.test if args.test else '') + '.json')
    env = os.environ.copy()
    env.setdefault('GOGC', '800')
    # Keep module and compiler caches writable inside the benchmark checkout.
    env.setdefault('GOCACHE', str(base / 'output-build-cache'))
    if args.run_only:
        work = args.artifact.resolve() if args.artifact else Path(json.loads(pointer.read_text())['artifact'])
    else:
        work = Path(tempfile.mkdtemp(prefix=mode + '-', dir=store))
        print(f'Artifact: {work}', flush=True)
        (work / 'src/Test').mkdir(parents=True)
        for name in ['App.purs', 'Bench.purs', 'Bench.js']:
            shutil.copyfile(ROOT / 'src' / name, work / 'src' / name)
        modules = list(validator.CASES) if not args.test else [args.test.removeprefix('Test.')]
        # Compile only core modules: native .go files are not JS foreigns.
        for name in modules:
            source = ROOT / 'src/Test' / (name.replace('.', '/') + '.purs')
            target = work / 'src/Test' / (name.replace('.', '/') + '.purs')
            target.parent.mkdir(parents=True, exist_ok=True)
            shutil.copyfile(source, target)
        if args.test:
            (work / 'src/App.purs').write_text('module App where\nimport Prelude\nimport Effect (Effect)\nimport Bench (runBench)\n' +
                f'import Test.{modules[0]} as Selected\nmain :: Effect Unit\nmain = void (runBench Selected.describe Selected.act)\n')
        shutil.copyfile(base / 'spago.psgo.yaml', work / 'spago.yaml')
        if (base / 'spago').is_dir():
            shutil.copytree(base / 'spago', work / '.spago')
        purs = Path(env.get('PSGO_PURS', str(ROOT / 'run/bak/js/node_modules/purescript/purs.bin'))).resolve()
        (work / 'bin').mkdir()
        (work / 'bin/purs').symlink_to(purs)
        env['PATH'] = str(work / 'bin') + os.pathsep + env['PATH']
        # Stock CoreFn is required by this upstream compiler; TAST is for the
        # separate gopurs backend, not an interchangeable psgo input.
        run(['spago', 'build'], work, 'corefn.log', env)
        for attempt in range(1, 4):
            try:
                run(['psgo', '--no-build'], work, f'psgo-{attempt}.log', env)
                break
            except RuntimeError as error:
                # This upstream generator edits one go.mod concurrently while
                # initializing packages. Retry only that observed startup race.
                if 'go.mod changed during editing' not in str(error) or attempt == 3:
                    raise
                print('Retrying psgo go.mod initialization race', flush=True)
        loader = work / 'purescript-native'
        for name in ['ffi_loader.go', 'go.mod']:
            path = loader / name
            path.write_text(path.read_text().replace('github.com/purescript-native/go-ffi', 'github.com/i-am-the-slime/go-ffi'))
        shutil.copyfile(ROOT / 'bin/psgo/bench_ffi.go', loader / 'bench_ffi.go')
        run(['go', 'mod', 'edit', '-dropreplace', 'github.com/purescript-native/go-ffi'], work, 'drop-replace.log', env)
        local_ffi = ROOT.parent / 'go-ffi'
        if local_ffi.is_dir():
            shutil.copytree(local_ffi, work / 'go-ffi', ignore=shutil.ignore_patterns('.git'))
            run(['go', 'mod', 'edit', '-replace=github.com/i-am-the-slime/go-ffi=./go-ffi'], work, 'ffi-replace.log', env)
        shutil.copyfile(ROOT / 'bin/psgo/main.go', work / 'main.go')
        run(['go', 'mod', 'tidy'], work, 'go-tidy.log', env)
        run(['go', 'build', '-pgo=off', '-o', 'Main', './main.go'], work, 'go-build.log', env)
        hashes = {str(p.relative_to(work)): hashlib.sha256(p.read_bytes()).hexdigest()
                  for folder in ['src', 'output', 'purescript-native', 'go-ffi'] for p in (work / folder).rglob('*')
                  if p.is_file()}
        for name in ['main.go', 'go.mod', 'go.sum']:
            if (work / name).is_file():
                hashes[name] = hashlib.sha256((work / name).read_bytes()).hexdigest()
        manifest = {'mode': mode, 'test': args.test, 'expected': args.expected,
                    'binary_sha256': hashlib.sha256((work / 'Main').read_bytes()).hexdigest(),
                    'source_sha256': hashes, 'gogc': env['GOGC'], 'pgo': 'off',
                    'purs': subprocess.check_output([purs, '--version'], text=True).strip()}
        (work / 'manifest.json').write_text(json.dumps(manifest, indent=2) + '\n')
        pointer.write_text(json.dumps({'artifact': str(work)}) + '\n')
    if args.build_only:
        return
    manifest = json.loads((work / 'manifest.json').read_text())
    if [manifest[k] for k in ['mode', 'test', 'expected']] != [mode, args.test, args.expected]:
        raise RuntimeError('Artifact mode does not match requested benchmark')
    if hashlib.sha256((work / 'Main').read_bytes()).hexdigest() != manifest['binary_sha256']:
        raise RuntimeError('Benchmark binary changed since build')
    if env['GOGC'] != manifest['gogc']:
        raise RuntimeError('GOGC differs from build manifest')
    result = subprocess.run([str(work / 'Main')], cwd=work, env=env, capture_output=True, text=True)
    print(result.stdout, end='')
    print(result.stderr, end='', file=sys.stderr)
    (work / 'run.stdout').write_text(result.stdout)
    if result.returncode:
        raise RuntimeError(f'Benchmark exited {result.returncode}')
    checked = validator.validate_output(result.stdout, mode, args.test, args.expected)
    (work / 'validated.json').write_text(json.dumps(checked, indent=2) + '\n')


if __name__ == '__main__':
    try:
        main()
    except (ValueError, OSError, RuntimeError, subprocess.SubprocessError) as error:
        print(error, file=sys.stderr)
        sys.exit(1)

#!/usr/bin/env python3
"""Build and validate isolated Go/Rust benchmark executables."""
import argparse
import hashlib
import json
import os
from pathlib import Path
import re
import shutil
import subprocess
import sys

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[2]
sys.path.insert(0, str(ROOT / 'bin/benchmark'))
from validate import expected_cases, validate_output


def digest(path):
    return hashlib.sha256(Path(path).read_bytes()).hexdigest()


def write_json(path, value):
    temporary = path.with_suffix(path.suffix + '.tmp')
    temporary.write_text(json.dumps(value, indent=2) + '\n')
    temporary.replace(path)


def inputs(language):
    result = {}
    for base in ['src', 'srx']:
        for path in sorted((ROOT / base).rglob('*')):
            if path.is_file() and path.suffix in {'.purs', '.' + ('go' if language == 'go' else 'rs')}:
                result[str(path.relative_to(ROOT))] = digest(path)
    backend = ROOT.parent / ('gopurs/gopurs/bin' if language == 'go' else 'purust/purust/bin')
    backend_name = 'gopurs' if language == 'go' else 'purust'
    for path in [backend / backend_name, backend / (backend_name + '.js'), backend / (backend_name + '-native')]:
        if path.is_file():
            result[str(path)] = digest(path)
    for path in [ROOT / f'run/bak/{language}/spago.{language}.yaml',
                 Path(__file__), ROOT / f'bin/{language}/run', ROOT / 'bin/benchmark/validate.py']:
        result[str(path.relative_to(ROOT))] = digest(path)
    return result


def logged(command, directory, log, env):
    command = list(map(str, command))
    print(' '.join(command), flush=True)
    with log.open('w') as output:
        output.write(json.dumps(command) + '\n')
        output.flush()
        result = subprocess.run(command, cwd=directory, env=env, stdout=output, stderr=subprocess.STDOUT)
    if result.returncode:
        print('\n'.join(log.read_text(errors='replace').splitlines()[-40:]), file=sys.stderr)
        raise RuntimeError(f'Command failed ({result.returncode}); full log: {log}')
    return command


def profile(language, args, env):
    if language == 'go':
        env['GOGC'] = env.get('GOGC', '800')
        env['GOWORK'] = 'off'
        env.pop('PPROF', None)
        return {'GOGC': env['GOGC'], 'pgo': args.pgo,
                'GOFLAGS': env.get('GOFLAGS', ''), 'GOEXPERIMENT': env.get('GOEXPERIMENT', '')}
    # Override the generated O1/debug=true release profile without editing Purust.
    env.setdefault('CARGO_PROFILE_RELEASE_OPT_LEVEL', '3')
    env.setdefault('CARGO_PROFILE_RELEASE_DEBUG', 'false')
    if 'CARGO_BUILD_TARGET' in env:
        raise ValueError('CARGO_BUILD_TARGET is unsupported: this runner executes a host binary')
    return {key: value for key, value in sorted(env.items()) if key.startswith('CARGO_PROFILE_RELEASE_')
            or key in {'RUSTFLAGS', 'CARGO_ENCODED_RUSTFLAGS'}}


def prepare(directory, language, args):
    marker = directory / '.altbak-native-workspace'
    identity = f'{ROOT}\n{language}\n'
    if directory.is_symlink() or directory == ROOT or directory in ROOT.parents:
        raise ValueError(f'Unsafe workspace: {directory}')
    if directory.exists() and any(directory.iterdir()) and (not marker.is_file() or marker.read_text() != identity):
        raise ValueError(f'Nonempty directory is not our isolated workspace: {directory}')
    if args.clean and directory.exists():
        shutil.rmtree(directory)
    directory.mkdir(parents=True, exist_ok=True)
    marker.write_text(identity)
    # Invalidate the previous build before touching its inputs; retain dependency caches.
    (directory / 'manifest.json').unlink(missing_ok=True)
    for name in ['src', 'output']:
        target = directory / name
        if target.is_symlink():
            target.unlink()
        elif target.exists():
            shutil.rmtree(target)
    shutil.copytree(ROOT / 'src', directory / 'src')
    entry = {'pure': 'App', 'ffi': 'AppFFI', 'fficc': 'AppFFICheatcode', 'test': 'AppX', 'x': 'AppX'}[args.mode]
    if args.mode == 'test':
        # Only the workspace copy is replaced; the user's AppX never changes.
        (directory / 'src/AppX.purs').write_text(f'''module AppX where
import Prelude
import Effect (Effect)
import Bench (runBench)
import Test.{args.test} as Case
main :: Effect Unit
main = void $ runBench Case.describe Case.act
''')
    elif args.mode == 'x':
        for path in sorted((ROOT / 'srx').rglob('*')):
            if path.is_file():
                target = directory / 'src' / path.relative_to(ROOT / 'srx')
                target.parent.mkdir(parents=True, exist_ok=True)
                shutil.copy2(path, target)
        (directory / 'var').mkdir(exist_ok=True)
    source = ROOT / f'run/bak/{language}/spago.{language}.yaml'
    config = re.sub(r'(?m)^(\s*path:\s*)"([^"\n]+)"',
                    lambda match: match[1] + json.dumps(str((ROOT / match[2]).resolve())), source.read_text())
    (directory / 'spago.yaml').write_text(config)
    # An independent lock/cache is kept per mode; no root links or config files are changed.
    return entry


def execute(binary, directory, env, args, name='results'):
    if args.mode == 'x':
        (directory / 'var').mkdir(exist_ok=True)
    raw = directory / (name + '.log')
    with raw.open('w') as output:
        process = subprocess.Popen([str(binary)], cwd=directory, env=env, stdout=subprocess.PIPE,
                                   stderr=subprocess.STDOUT, text=True, bufsize=1)
        chunks = []
        assert process.stdout is not None
        for line in process.stdout:
            print(line, end='', flush=True)
            output.write(line)
            output.flush()
            chunks.append(line)
        code = process.wait()
    if code:
        raise RuntimeError(f'Benchmark exited {code}; see {raw}')
    result = validate_output(''.join(chunks), args.mode, args.test, args.expected)
    write_json(directory / (name + '.json'), result)
    return result


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('language', choices=['go', 'rust'])
    modes = parser.add_mutually_exclusive_group()
    modes.add_argument('--ffi', action='store_true')
    modes.add_argument('--fficc', action='store_true')
    modes.add_argument('--test')
    modes.add_argument('--x', action='store_true')
    parser.add_argument('--expected')
    actions = parser.add_mutually_exclusive_group()
    actions.add_argument('--build-only', action='store_true')
    actions.add_argument('--run-only', action='store_true')
    parser.add_argument('-c', '--clean', action='store_true',
                        help='rebuild the backend compiler and clean this isolated workspace')
    parser.add_argument('--pgo', action='store_true', help='Go: train and rebuild, using the same GOGC')
    parser.add_argument('--pgo-profile', type=Path, help='Go: use an existing profile; avoids a training run')
    parser.add_argument('--build-dir', type=Path, help='override the isolated workspace directory')
    parser.add_argument('-n', '--native', action='store_true', help='use the native compiler binary')
    args = parser.parse_args()
    if args.test:
        args.test = args.test.removeprefix('Test.')
        if not re.fullmatch(r'[A-Z][A-Za-z0-9_]*(?:\.[A-Z][A-Za-z0-9_]*)*', args.test):
            parser.error('invalid --test module')
    args.mode = 'test' if args.test else 'x' if args.x else 'ffi' if args.ffi else 'fficc' if args.fficc else 'pure'
    expected_cases(args.mode, args.test, args.expected)  # Fail before writes/builds for unknown contracts.
    if args.run_only and args.clean:
        parser.error('--run-only is incompatible with --clean (requires a compiler and executable rebuild)')
    if args.language != 'go' and (args.pgo or args.pgo_profile):
        parser.error('PGO options are Go-only')
    if args.pgo_profile:
        args.pgo = True
        args.pgo_profile = args.pgo_profile.resolve(strict=True)
    if args.pgo and args.build_only and not args.pgo_profile:
        parser.error('--build-only --pgo requires --pgo-profile; training executes the benchmark')
    key = args.mode + ('-' + args.test.replace('.', '-') if args.test else '')
    if args.pgo:
        key += '-pgo'
    selected_directory = args.build_dir or ROOT / f'run/bak/{args.language}/modes/{key}'
    if selected_directory.is_symlink():
        raise ValueError('Build directory must not be a symlink')
    directory = selected_directory.resolve()
    env = os.environ.copy()
    env['PATH'] = str(ROOT / 'run/bak/js/node_modules/.bin') + os.pathsep + env['PATH']
    config = profile(args.language, args, env)
    print(f'{args.language}: mode={args.mode}, profile={json.dumps(config, sort_keys=True)}', flush=True)
    binary = directory / 'benchmark'
    manifest_path = directory / 'manifest.json'
    if args.run_only:
        fingerprint = inputs(args.language)
        manifest = json.loads(manifest_path.read_text())
        required = {'language': args.language, 'mode': args.mode, 'test': args.test,
                    'expected': args.expected, 'profile': config, 'inputs': fingerprint,
                    'native': args.native}
        for name, value in required.items():
            if manifest.get(name) != value:
                raise ValueError(f'Stale or incompatible build ({name}); rebuild without --run-only')
        if args.pgo_profile and digest(args.pgo_profile) != manifest.get('pgo_profile_sha256'):
            raise ValueError('Requested PGO profile differs from the compiled profile')
        if digest(binary) != manifest['binary_sha256']:
            raise ValueError('Executable SHA differs from its build manifest')
        execute(binary, directory, env, args)
        return
    entry = prepare(directory, args.language, args)
    logs = directory / 'logs'
    logs.mkdir(exist_ok=True)
    commands = []
    if args.clean:
        backend_name = 'gopurs' if args.language == 'go' else 'purust'
        backend_directory = ROOT.parent / backend_name / backend_name
        commands.append(logged(['npm', 'run', 'build', '--silent'], backend_directory,
                               logs / 'compiler-build.log', env))
    # --clean changes the compiler bundle; record the compiler actually used below.
    fingerprint = inputs(args.language)
    commands.append(logged(['spago', 'build'], directory, logs / 'spago.log', env))
    if args.language == 'go':
        backend = ROOT.parent / 'gopurs/gopurs/bin/gopurs-native' if args.native else ROOT.parent / 'gopurs/gopurs/bin/gopurs'
        commands.append(logged([backend, '--main', entry], directory, logs / 'backend.log', env))
        project = directory / 'output'
        if not (project / 'main/main.go').is_file():
            raise RuntimeError('Gopurs did not generate main/main.go')
        commands.append(logged(['go', 'mod', 'tidy'], project, logs / 'modules.log', env))
        profile_path = args.pgo_profile
        if args.pgo and profile_path is None:
            commands.append(logged(['go', 'build', '-pgo=off', '-o', directory / 'training', './main'], project,
                                   logs / 'training-build.log', env))
            profile_path = directory / 'cpu.prof'
            profile_path.unlink(missing_ok=True)
            training_env = dict(env, PPROF='1')
            execute(directory / 'training', directory, training_env, args, 'training-results')
            if not profile_path.is_file() or not profile_path.stat().st_size:
                raise RuntimeError('PGO training produced no profile')
        commands.append(logged(['go', 'build', '-pgo=' + (str(profile_path) if profile_path else 'off'),
                                '-o', binary, './main'], project, logs / 'build.log', env))
        toolchain = subprocess.check_output(['go', 'version'], text=True, env=env).strip()
    else:
        backend = ROOT.parent / 'purust/purust/bin/purust-native' if args.native else ROOT.parent / 'purust/purust/bin/purust'
        command = [backend, '--main', entry, '--source', 'output', '--out', 'output/purust_output']
        if args.mode == 'x':
            command.append('--threaded')
        commands.append(logged(command, directory, logs / 'backend.log', env))
        project = directory / 'output/purust_output'
        if not (project / 'Cargo.toml').is_file() or not (project / 'src/main.rs').is_file():
            raise RuntimeError('Purust did not generate a Rust executable project')
        target = directory / 'target'
        commands.append(logged(['cargo', 'build', '--release', '--target-dir', target, '--bin', 'purust_output'],
                               project, logs / 'build.log', env))
        shutil.copy2(target / 'release/purust_output', binary)
        toolchain = subprocess.check_output(['rustc', '--version'], text=True, env=env).strip()
    if fingerprint != inputs(args.language):
        raise RuntimeError('Sources changed during build; no executable manifest was accepted')
    manifest = {'language': args.language, 'mode': args.mode, 'test': args.test, 'expected': args.expected,
                'native': args.native, 'entry': entry, 'profile': config, 'inputs': fingerprint, 'binary_sha256': digest(binary),
                'toolchain': toolchain, 'commands': commands}
    if args.language == 'go' and args.pgo:
        manifest['pgo_profile_sha256'] = digest(profile_path)
    write_json(manifest_path, manifest)
    print(f'Built {binary}; manifest {manifest_path}', flush=True)
    if not args.build_only:
        execute(binary, directory, env, args)


if __name__ == '__main__':
    try:
        main()
    except (OSError, ValueError, RuntimeError, subprocess.SubprocessError) as error:
        print(f'ERROR: {error}', file=sys.stderr)
        sys.exit(1)

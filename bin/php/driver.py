#!/usr/bin/env python3
"""Build and validate PHP benchmarks in independent, reusable workspaces."""
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

PHP_FLAGS = ['-d', 'xdebug.mode=off', '-d', 'opcache.enable_cli=1',
             '-d', 'opcache.file_update_protection=0',
             '-d', 'opcache.jit_buffer_size=128M', '-d', 'opcache.jit=1255']
COMPOSER_FLAGS = ['--no-plugins', '--no-scripts', '--no-interaction',
                  '--no-progress', '--prefer-dist', '--no-dev']


def digest(path):
    return hashlib.sha256(Path(path).read_bytes()).hexdigest()


def write_json(path, value):
    temporary = path.with_suffix(path.suffix + '.tmp')
    temporary.write_text(json.dumps(value, indent=2) + '\n')
    temporary.replace(path)


def inputs():
    paths = [p for base in ['src', 'srx'] for p in (ROOT / base).rglob('*')
             if p.is_file() and p.suffix in {'.purs', '.php'}]
    backend = ROOT.parent / 'phpurs/phpurs/bin'
    paths += [backend / 'phpurs', backend / 'phpurs.js', Path(__file__),
              ROOT / 'bin/php/run', ROOT / 'bin/php/composer.phar', ROOT / 'bin/php/composer-state.php',
              ROOT / 'bin/benchmark/validate.py',
              ROOT / 'test/native/php/generated.py', ROOT / 'test/native/oracle.py',
              ROOT / 'run/bak/php/spago.php.yaml', ROOT / 'run/bak/php/composer.json',
              ROOT / 'run/bak/php/composer.lock']
    return {str(p): digest(p) for p in sorted(paths)}


def artifacts(directory):
    paths = [p for base in ['output', 'vendor'] for p in (directory / base).rglob('*')
             if p.is_file() and (base == 'vendor' or p.suffix == '.php' or p.name == 'composer.json')]
    paths += [directory / 'composer.json', directory / 'composer.lock']
    return {str(p.relative_to(directory)): digest(p) for p in sorted(paths)}


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


def composer_inputs(directory, php, env):
    lock_path = directory / 'composer.lock'
    if not lock_path.is_file():
        raise ValueError('Missing composer.lock; prepare and review the benchmark lock before building')
    lock = json.loads(lock_path.read_text())
    generated = json.loads((directory / 'output/composer.json').read_text())
    packages = [p for p in lock.get('packages', []) if p.get('name') == 'phpurs/lib-deps']
    if (generated.get('name') != 'phpurs/lib-deps' or len(packages) != 1
            or packages[0].get('dist', {}).get('type') != 'path'
            or packages[0]['dist'].get('url') != 'output'):
        raise ValueError('composer.lock must pin the generated phpurs/lib-deps path package at output')
    # Composer install checks the root lock, but does not refresh a path
    # package's requirements. Refuse that stale lock instead of solving here.
    defaults = {key: {} for key in ['require', 'require-dev', 'conflict', 'replace',
                                   'provide', 'autoload', 'autoload-dev', 'extra']}
    defaults.update({'type': 'library', 'bin': [], 'include-path': [], 'target-dir': None})
    for key, default in defaults.items():
        if generated.get(key, default) != packages[0].get(key, default):
            raise ValueError(f'Generated Composer {key} differs from composer.lock; '
                             'refresh and review the benchmark lock explicitly')
    phar = ROOT / 'bin/php/composer.phar'
    probe = ROOT / 'bin/php/composer-state.php'
    context = json.loads(subprocess.check_output(
        [php, '-d', 'xdebug.mode=off', str(probe), str(phar), str(directory / 'composer.json')],
        cwd=directory, env=env, text=True))
    if lock.get('content-hash') != context['content-hash']:
        raise ValueError('composer.lock is stale for composer.json; refresh and review the lock explicitly')
    # Hash environment/configuration rather than recording credentials. The
    # PHP platform includes extensions and library versions as seen by Composer.
    settings = {key: value for key, value in env.items()
                if key.startswith('COMPOSER_') or key in {'COMPOSER', 'PHPRC', 'PHP_INI_SCAN_DIR'}}
    return {'version': 1, 'context': context, 'flags': COMPOSER_FLAGS,
            'environment': hashlib.sha256(json.dumps(settings, sort_keys=True).encode()).hexdigest(),
            'files': {str(p): digest(p) for p in [directory / 'composer.json', lock_path,
                      directory / 'output/composer.json', Path(php), phar, probe, Path(__file__)]}}


def vendor_artifacts(directory):
    vendor = directory / 'vendor'
    package = vendor / 'phpurs/lib-deps'
    if (vendor.is_symlink() or not (vendor / 'autoload.php').is_file()
            or not (vendor / 'composer/installed.json').is_file()
            or not package.is_symlink() or package.resolve() != directory / 'output'):
        return None
    files = {}
    # The generated path package is a link to the current output, not a cached
    # copy of PHP. Do not follow directory links while inventorying dependencies.
    for root, dirs, names in os.walk(vendor, followlinks=False):
        for name in sorted(dirs + names):
            path = Path(root) / name
            relative = str(path.relative_to(vendor))
            if path.is_symlink():
                value = {'link': os.readlink(path)}
                if path != package:
                    if not path.is_file():
                        return None
                    value['sha256'] = digest(path)
                files[relative] = value
            elif path.is_file():
                files[relative] = {'sha256': digest(path)}
    return files


def prepare_composer(directory, php, env, log):
    env = dict(env, COMPOSER='composer.json', COMPOSER_VENDOR_DIR='vendor',
               COMPOSER_MIRROR_PATH_REPOS='0')
    required = composer_inputs(directory, php, env)
    stamp = directory / '.composer-prepared.json'
    try:
        previous = json.loads(stamp.read_text())
    except (OSError, ValueError):
        previous = None
    current = vendor_artifacts(directory)
    if current is not None and previous == {'inputs': required, 'vendor': current}:
        print('Composer: reusing verified locked dependencies', flush=True)
        return {'status': 'reused', 'commands': []}
    stamp.unlink(missing_ok=True)
    vendor = directory / 'vendor'
    # install alone does not repair edited package files. Recreate this isolated
    # workspace's dependency tree on a miss, using the committed lock only.
    if vendor.is_symlink():
        vendor.unlink()
    elif vendor.exists():
        shutil.rmtree(vendor)
    command = logged([php, '-d', 'xdebug.mode=off', ROOT / 'bin/php/composer.phar',
                      'install', *COMPOSER_FLAGS], directory, log, env)
    if required != composer_inputs(directory, php, env):
        raise RuntimeError('Composer inputs changed during preparation; no reusable state was accepted')
    current = vendor_artifacts(directory)
    if current is None:
        raise RuntimeError('Composer must install vendor with phpurs/lib-deps linked to output')
    write_json(stamp, {'inputs': required, 'vendor': current})
    return {'status': 'installed', 'commands': [command]}


def prepare(directory, args):
    marker = directory / '.altbak-php-workspace'
    identity = str(ROOT) + '\n'
    if directory == ROOT or directory in ROOT.parents:
        raise ValueError(f'Unsafe workspace: {directory}')
    if directory.exists() and any(directory.iterdir()) and (not marker.is_file() or marker.read_text() != identity):
        raise ValueError(f'Nonempty directory is not our isolated workspace: {directory}')
    if args.clean and directory.exists():
        shutil.rmtree(directory)
    directory.mkdir(parents=True, exist_ok=True)
    marker.write_text(identity)
    (directory / 'manifest.json').unlink(missing_ok=True)
    for name in ['src', 'output', 'opcache']:
        target = directory / name
        if target.is_symlink():
            target.unlink()
        elif target.exists():
            shutil.rmtree(target)
    # JSON diagnostics use independent workspaces and optional dependencies.
    shutil.copytree(ROOT / 'src', directory / 'src',
                    ignore=shutil.ignore_patterns('JsonTypedAst', 'JsonTypedAst.purs',
                                                 'JsonTypedAst.go', 'JsonTypedAst.js',
                                                 'JsonDecoding', 'JsonDecoding.purs',
                                                 'JsonDecoding.go', 'JsonDecoding.js'))
    entry = {'pure': 'App', 'ffi': 'AppFFI', 'fficc': 'AppFFICheatcode', 'test': 'AppX', 'x': 'AppX'}[args.mode]
    if args.mode == 'test':
        # Replace only the isolated copy, preserving the user's AppX.
        (directory / 'src/AppX.purs').write_text(f'''module AppX where
import Prelude
import Effect (Effect)
import Bench (runBench)
import Test.{args.test} as Case
main :: Effect Unit
main = void $ runBench Case.describe Case.act
''')
    elif args.mode == 'x':
        shutil.copytree(ROOT / 'srx', directory / 'src', dirs_exist_ok=True)
        (directory / 'var').mkdir(exist_ok=True)
    source = ROOT / 'run/bak/php/spago.php.yaml'
    config = re.sub(r'(?m)^(\s*(?:path|cmd):\s*)"([^"\n]+)"',
                    lambda m: m[1] + json.dumps(str((ROOT / m[2]).resolve())), source.read_text())
    (directory / 'spago.yaml').write_text(config)
    for name in ['composer.json', 'composer.lock']:
        shutil.copy2(ROOT / 'run/bak/php' / name, directory / name)
    (directory / 'opcache').mkdir(exist_ok=True)
    return entry


def execute(php, entry, directory, env, args):
    command = [php, *PHP_FLAGS, '-d', f'opcache.file_cache={directory / "opcache"}',
               str(directory / 'output' / entry / 'main.mod.php')]
    chunks = []
    with (directory / 'results.log').open('w') as output:
        process = subprocess.Popen(command, cwd=directory, env=env, stdout=subprocess.PIPE,
                                   stderr=subprocess.STDOUT, text=True, bufsize=1)
        assert process.stdout is not None
        for line in process.stdout:
            print(line, end='', flush=True)
            output.write(line)
            output.flush()
            chunks.append(line)
        code = process.wait()
    if code:
        raise RuntimeError(f'PHP exited {code}; see {directory / "results.log"}')
    write_json(directory / 'results.json', validate_output(''.join(chunks), args.mode, args.test, args.expected))


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    modes = parser.add_mutually_exclusive_group()
    modes.add_argument('--ffi', action='store_true')
    modes.add_argument('--fficc', action='store_true')
    modes.add_argument('--test')
    modes.add_argument('--x', action='store_true')
    parser.add_argument('--expected')
    actions = parser.add_mutually_exclusive_group()
    actions.add_argument('--build-only', action='store_true')
    actions.add_argument('--run-only', action='store_true')
    parser.add_argument('-c', '--clean', action='store_true', help='clean only this isolated workspace')
    parser.add_argument('--build-dir', type=Path)
    args = parser.parse_args()
    if args.test:
        args.test = args.test.removeprefix('Test.')
        if not re.fullmatch(r'[A-Z][A-Za-z0-9_]*(?:\.[A-Z][A-Za-z0-9_]*)*', args.test):
            parser.error('invalid --test module')
    args.mode = 'test' if args.test else 'x' if args.x else 'ffi' if args.ffi else 'fficc' if args.fficc else 'pure'
    expected_cases(args.mode, args.test, args.expected)
    if args.run_only and args.clean:
        parser.error('--run-only cannot clean its artifacts')
    key = args.mode + ('-' + args.test.replace('.', '-') if args.test else '')
    selected = args.build_dir or ROOT / 'run/bak/php/modes' / key
    if selected.is_symlink():
        raise ValueError('Build directory must not be a symlink')
    directory = selected.resolve()
    php = shutil.which(os.environ.get('PHP', 'php'))
    if not php:
        raise ValueError('PHP executable not found; set PHP to its path')
    env = os.environ.copy()
    env['PATH'] = str(ROOT / 'run/bak/js/node_modules/.bin') + os.pathsep + env['PATH']
    version = subprocess.check_output([php, '--version'], text=True).strip()
    profile = {'php': php, 'version': version, 'flags': PHP_FLAGS,
               'opcache.file_cache': str(directory / 'opcache'),
               'ini': subprocess.check_output([php, '--ini'], text=True).strip()}
    fingerprint = inputs()
    manifest_path = directory / 'manifest.json'
    required = {'language': 'php', 'mode': args.mode, 'test': args.test, 'expected': args.expected,
                'profile': profile, 'inputs': fingerprint}
    if args.run_only:
        manifest = json.loads(manifest_path.read_text())
        for name, value in required.items():
            if manifest.get(name) != value:
                raise ValueError(f'Stale or incompatible build ({name}); rebuild without --run-only')
        if manifest['artifacts'] != artifacts(directory):
            raise ValueError('PHP/dependency artifacts differ from the build manifest')
        vendor = vendor_artifacts(directory)
        if vendor is None or manifest.get('composer_vendor') != vendor:
            raise ValueError('Composer dependency files or links differ from the build manifest')
        execute(php, manifest['entry'], directory, env, args)
        return
    entry = prepare(directory, args)
    logs = directory / 'logs'
    logs.mkdir(exist_ok=True)
    commands = [logged(['spago', 'build', '--backend-args', '--main ' + entry],
                       directory, logs / 'spago.log', env)]
    if not (directory / 'output' / entry / 'main.mod.php').is_file():
        raise RuntimeError(f'Phpurs did not generate {entry}/main.mod.php')
    composer = prepare_composer(directory, php, env, logs / 'composer.log')
    commands.extend(composer['commands'])
    if args.mode in {'pure', 'ffi', 'fficc', 'x'}:
        commands.append(logged([sys.executable, '-B', ROOT / 'test/native/php/generated.py', directory,
                                '--mode', args.mode, '--php', php], directory,
                               logs / 'generated-contracts.log', env))
    if fingerprint != inputs():
        raise RuntimeError('Sources changed during build; no manifest was accepted')
    write_json(manifest_path, dict(required, entry=entry, commands=commands,
                                   composer=composer, composer_vendor=vendor_artifacts(directory),
                                   artifacts=artifacts(directory)))
    print(f'Built PHP {args.mode}; manifest {manifest_path}', flush=True)
    if not args.build_only:
        execute(php, entry, directory, env, args)


if __name__ == '__main__':
    try:
        main()
    except (OSError, ValueError, RuntimeError, subprocess.SubprocessError) as error:
        print(f'ERROR: {error}', file=sys.stderr)
        sys.exit(1)

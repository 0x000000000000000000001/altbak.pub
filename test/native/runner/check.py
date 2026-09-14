#!/usr/bin/env python3
"""Runner contracts with simulated tools: no compiler or benchmark is launched."""
import argparse
import importlib.util
import json
import os
from pathlib import Path
import sys
import tempfile
import unittest
from unittest.mock import patch

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[3]
spec = importlib.util.spec_from_file_location('native_driver', ROOT / 'bin/native/driver.py')
driver = importlib.util.module_from_spec(spec)
spec.loader.exec_module(driver)


class RunnerContracts(unittest.TestCase):
    def setUp(self):
        self.temporary = tempfile.TemporaryDirectory()
        self.root = Path(self.temporary.name) / 'project'
        self.root.mkdir()
        for path, contents in {
            'src/AppX.purs': 'user-owned AppX', 'src/Test/Fib.purs': 'source',
            'srx/AppX.purs': 'extended entry', 'srx/Test/Extended.purs': 'extended source',
            'run/bak/go/spago.go.yaml': 'path: "../gopurs/gopurs-prelude"\n',
            'run/bak/rust/spago.rust.yaml': 'path: "../purust/purust-prelude"\n',
        }.items():
            target = self.root / path
            target.parent.mkdir(parents=True, exist_ok=True)
            target.write_text(contents)
        self.directory = self.root / 'work'
        self.root_patch = patch.object(driver, 'ROOT', self.root)
        self.root_patch.start()
        self.addCleanup(self.root_patch.stop)
        self.addCleanup(self.temporary.cleanup)

    def args(self, **kwargs):
        return argparse.Namespace(clean=False, mode='test', test='Fib', **kwargs)

    def test_workspace_preserves_appx_and_absolutizes_package_paths(self):
        driver.prepare(self.directory, 'go', self.args())
        self.assertEqual((self.root / 'src/AppX.purs').read_text(), 'user-owned AppX')
        self.assertIn('import Test.Fib as Case', (self.directory / 'src/AppX.purs').read_text())
        self.assertIn(str(self.root.parent / 'gopurs/gopurs-prelude'), (self.directory / 'spago.yaml').read_text())
        args = self.args(); args.mode='x'
        driver.prepare(self.directory, 'go', args)
        self.assertEqual((self.directory / 'src/AppX.purs').read_text(), 'extended entry')
        self.assertEqual((self.root / 'src/AppX.purs').read_text(), 'user-owned AppX')

    def test_unowned_directory_rejected_and_manifest_invalidated(self):
        self.directory.mkdir(); (self.directory / 'user.txt').write_text('keep')
        with self.assertRaises(ValueError):
            driver.prepare(self.directory, 'go', self.args())
        self.assertEqual((self.directory / 'user.txt').read_text(), 'keep')
        (self.directory / 'user.txt').unlink()
        driver.prepare(self.directory, 'go', self.args())
        (self.directory / 'manifest.json').write_text('old build')
        driver.prepare(self.directory, 'go', self.args())
        self.assertFalse((self.directory / 'manifest.json').exists())

    def test_go_gc_and_rust_explicit_profile(self):
        args = self.args(pgo=False)
        normal = {'PPROF':'1'}; pgo = {'PPROF':'1'}
        first = driver.profile('go', args, normal)
        args.pgo=True
        second = driver.profile('go', args, pgo)
        self.assertEqual(first['GOGC'], second['GOGC'])
        self.assertEqual(first['GOGC'], '800')
        self.assertNotIn('PPROF', normal)
        env = {}; result = driver.profile('rust', args, env)
        self.assertEqual(result['CARGO_PROFILE_RELEASE_OPT_LEVEL'], '3')
        self.assertEqual(result['CARGO_PROFILE_RELEASE_DEBUG'], 'false')
        env = {'CARGO_PROFILE_RELEASE_OPT_LEVEL':'1'}
        self.assertEqual(driver.profile('rust', args, env)['CARGO_PROFILE_RELEASE_OPT_LEVEL'], '1')

    def test_build_only_run_only_and_stale_manifest(self):
        # Fake only tool effects; the production orchestration and manifest code run unchanged.
        def tools(command, directory, log, env):
            command=list(map(str, command))
            if 'gopurs' in command[0]:
                target=directory/'output/main/main.go'; target.parent.mkdir(parents=True); target.write_text('generated')
            if command[:2] == ['go','build']:
                Path(command[command.index('-o')+1]).write_text('compiled')
            return command
        argv=['driver.py','go','--ffi','--build-only','--build-dir',str(self.directory)]
        with patch.object(sys,'argv',argv), patch.object(driver,'expected_cases'), \
             patch.object(driver,'inputs',return_value={'source':'abc'}), \
             patch.object(driver,'logged',side_effect=tools), \
             patch.object(driver.subprocess,'check_output',return_value='go test'), \
             patch.object(driver,'execute') as execute:
            driver.main(); execute.assert_not_called()
            argv[3]='--run-only'; driver.main(); execute.assert_called_once()
            argv[2]='--fficc'
            with self.assertRaisesRegex(ValueError,'mode'):
                driver.main()
            argv[2]='--ffi'; (self.directory/'benchmark').write_text('other executable')
            with self.assertRaisesRegex(ValueError,'SHA'):
                driver.main()

    def test_cli_rejects_incompatible_modes_before_build(self):
        for arguments in [['go','--ffi','--fficc'], ['go','--test','bad/name'],
                          ['rust','--pgo'], ['go','--pgo','--build-only'],
                          ['go','--run-only','--clean'], ['rust','--run-only','--clean']]:
            with self.subTest(arguments=arguments), patch.object(sys,'argv',['driver.py']+arguments), \
                 patch.object(driver,'expected_cases'), patch.object(driver,'prepare') as prepare, \
                 patch.object(driver,'logged') as logged, patch.object(driver,'inputs') as inputs:
                with self.assertRaises(SystemExit): driver.main()
                prepare.assert_not_called()
                logged.assert_not_called()
                inputs.assert_not_called()

    def test_clean_rebuilds_compiler_before_snapshot_and_executable(self):
        for language in ['go', 'rust']:
            for clean in [False, True]:
                with self.subTest(language=language, clean=clean):
                    directory = self.root / f'{language}-{clean}'
                    driver.prepare(directory, language, self.args())
                    (directory / 'old-artifact').write_text('cached')
                    version = ['old bundle']
                    events = []
                    backend_name = 'gopurs' if language == 'go' else 'purust'
                    backend_directory = self.root.parent / backend_name / backend_name

                    def inputs(selected):
                        self.assertEqual(selected, language)
                        events.append(('snapshot', version[0]))
                        return {'compiler': version[0], 'source': 'unchanged'}

                    def tools(command, working_directory, log, env):
                        command = list(map(str, command))
                        events.append(('command', command))
                        if command[0] == 'npm':
                            self.assertEqual(command, ['npm', 'run', 'build', '--silent'])
                            self.assertEqual(working_directory, backend_directory)
                            self.assertEqual(log, (directory / 'logs/compiler-build.log').resolve())
                            self.assertFalse((directory / 'old-artifact').exists())
                            version[0] = 'rebuilt bundle'
                        elif Path(command[0]).name == 'gopurs':
                            target = directory / 'output/main/main.go'
                            target.parent.mkdir(parents=True)
                            target.write_text('generated')
                        elif Path(command[0]).name == 'purust':
                            project = directory / 'output/purust_output'
                            (project / 'src').mkdir(parents=True)
                            (project / 'Cargo.toml').write_text('generated')
                            (project / 'src/main.rs').write_text('generated')
                        elif command[:2] == ['go', 'build']:
                            Path(command[command.index('-o') + 1]).write_text('compiled')
                        elif command[:2] == ['cargo', 'build']:
                            target = directory / 'target/release/purust_output'
                            target.parent.mkdir(parents=True)
                            target.write_text('compiled')
                        return command

                    argv = ['driver.py', language, '--build-only', '--build-dir', str(directory)]
                    if clean:
                        argv.append('--clean')
                    with patch.object(sys, 'argv', argv), patch.object(driver, 'expected_cases'), \
                         patch.object(driver, 'inputs', side_effect=inputs), \
                         patch.object(driver, 'logged', side_effect=tools), \
                         patch.object(driver.subprocess, 'check_output', return_value='simulated toolchain'), \
                         patch.object(driver, 'execute') as execute:
                        driver.main()
                        execute.assert_not_called()
                    manifest = json.loads((directory / 'manifest.json').read_text())
                    commands = manifest['commands']
                    self.assertEqual(manifest['inputs']['compiler'], version[0])
                    self.assertEqual(sum(command[0] == 'npm' for command in commands), int(clean))
                    self.assertEqual((directory / 'old-artifact').exists(), not clean)
                    if clean:
                        self.assertEqual(events[0], ('command', ['npm', 'run', 'build', '--silent']))
                        self.assertEqual(events[1], ('snapshot', 'rebuilt bundle'))
                        self.assertEqual(commands[1], ['spago', 'build'])
                    else:
                        self.assertEqual(commands[0], ['spago', 'build'])

    def test_failed_compiler_rebuild_stops_before_frontend_and_manifest(self):
        for language in ['go', 'rust']:
            with self.subTest(language=language):
                directory = self.root / f'failed-{language}'
                driver.prepare(directory, language, self.args())
                (directory / 'manifest.json').write_text('previous build')
                commands = []

                def tools(command, *unused):
                    commands.append(list(map(str, command)))
                    raise RuntimeError('compiler failed')

                argv = ['driver.py', language, '--clean', '--build-only', '--build-dir', str(directory)]
                with patch.object(sys, 'argv', argv), patch.object(driver, 'expected_cases'), \
                     patch.object(driver, 'logged', side_effect=tools), \
                     patch.object(driver, 'inputs', return_value={}) as inputs, \
                     patch.object(driver, 'execute') as execute:
                    with self.assertRaisesRegex(RuntimeError, 'compiler failed'):
                        driver.main()
                    self.assertEqual(commands, [['npm', 'run', 'build', '--silent']])
                    inputs.assert_not_called()
                    execute.assert_not_called()
                self.assertFalse((directory / 'manifest.json').exists())


if __name__ == '__main__':
    unittest.main()

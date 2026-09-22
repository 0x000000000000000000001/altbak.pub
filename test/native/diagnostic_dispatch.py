#!/usr/bin/env python3
"""Check diagnostic routing and optional-source isolation without compiling."""
import argparse
from contextlib import redirect_stderr
import importlib.util
import io
import itertools
import json
from pathlib import Path
import shutil
import subprocess
import sys
import tempfile
import unittest
from unittest.mock import patch

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[2]


def load(name, path):
    spec = importlib.util.spec_from_file_location(name, ROOT / path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module


dispatch = load("diagnostic_dispatch", "bin/benchmark/diagnostic-dispatch.py")
native = load("native_driver", "bin/native/driver.py")
php = load("php_driver", "bin/php/driver.py")
js = load("js_driver", "bin/js/driver.py")
json_diagnostic = load("json_diagnostic", "bin/benchmark/json-diagnostic.py")


class Routing(unittest.TestCase):
    def test_detection_preserves_regular_tests(self):
        self.assertIsNone(dispatch.diagnostic_name(["--test", "Fib", "--pgo"]))
        self.assertEqual(dispatch.diagnostic_name(["--test=Test.JsonTypedAst"]), "JsonTypedAst")
        self.assertEqual(dispatch.diagnostic_name(["--test", "Test.JsonDecoding"]), "JsonDecoding")
        with patch.object(sys, "argv", ["dispatch", "go", "--test", "Fib", "--pgo"]), \
                patch.object(dispatch.os, "execv", side_effect=SystemExit) as launch:
            with self.assertRaises(SystemExit):
                dispatch.main()
        self.assertEqual(launch.call_args.args[1][2:], ["go", "--test", "Fib", "--pgo"])

    def test_incompatible_options_fail_before_dispatch(self):
        combinations = [["--ffi"], ["--build-only", "--run-only"],
                        ["--build-only", "--output", "/tmp/unused-diagnostic-output"], ["--pgo"]]
        for flags in combinations:
            with self.subTest(flags=flags), patch.object(dispatch, "execute") as execute, \
                    patch.object(sys, "argv", ["dispatch", "go", "--test", "ArrayIndexing", *flags]), \
                    redirect_stderr(io.StringIO()):
                with self.assertRaises(SystemExit):
                    dispatch.main()
                execute.assert_not_called()

    def test_array_actions_select_one_runtime(self):
        with tempfile.TemporaryDirectory() as temporary:
            base = Path(temporary).resolve() / "array"
            for runtime in ["go", "js"]:
                for flags, actions in [([], ["--build-only", "--run-only"]),
                                       (["--build-only"], ["--build-only"]),
                                       (["--run-only"], ["--run-only"])]:
                    args = dispatch.options(runtime, ["--test", "ArrayIndexing", "--build-dir", str(base), *flags])
                    with patch.object(dispatch, "run") as run:
                        dispatch.execute(runtime, args)
                    commands = [call.args[0] for call in run.call_args_list]
                    self.assertEqual(len(commands), len(actions))
                    for command, action in zip(commands, actions):
                        self.assertIn(action, command)
                        self.assertEqual(command[command.index("--runtime") + 1], runtime)
                        self.assertEqual(command[command.index("--build-dir") + 1], base)

    def test_json_builds_are_new_and_run_only_reuses_last_success(self):
        with tempfile.TemporaryDirectory() as temporary:
            base = Path(temporary).resolve() / "json"
            counter = itertools.count()
            commands = []

            def run(command):
                commands.append(command)
                if "build" in command:
                    work = command[command.index("--workspace") + 1]
                    work.mkdir(parents=True)
                    (work / "manifest.json").write_text("{}")

            with patch.object(dispatch, "run", side_effect=run), \
                    patch.object(dispatch, "stamp", side_effect=lambda: str(next(counter))):
                args = dispatch.options("js", ["--test", "JsonTypedAst", "--build-dir", str(base), "--build-only"])
                dispatch.execute("js", args)
                first = dispatch.json_workspace(base, "js", True)
                dispatch.execute("js", args)
                second = dispatch.json_workspace(base, "js", True)
                self.assertNotEqual(first, second)
                self.assertTrue(first.is_dir())
                args = dispatch.options("js", ["--test", "JsonTypedAst", "--build-dir", str(base), "--run-only"])
                dispatch.execute("js", args)
                self.assertEqual(commands[-1][commands[-1].index("--workspace") + 1], second)
                self.assertIn("measure", commands[-1])
                self.assertEqual(sum("build" in command for command in commands), 2)
                with patch.object(dispatch, "run", side_effect=RuntimeError("build failed")):
                    args.run_only = False
                    with self.assertRaises(RuntimeError):
                        dispatch.execute("js", args)
                self.assertEqual(dispatch.json_workspace(base, "js", True), second)
                self.assertEqual(dispatch.json_workspace(second, "js", True), second)

    def test_general_json_selects_suite_and_rejects_other_suite_pointer(self):
        with tempfile.TemporaryDirectory() as temporary:
            base = Path(temporary).resolve()
            pointer = base / "latest-build.json"
            pointer.write_text(json.dumps({'runtime': 'go', 'suite': 'JsonDecoding', 'workspace': 'builds/first'}))
            args = dispatch.options('go', ['--test', 'JsonDecoding', '--build-dir', str(base), '--run-only'])
            with patch.object(dispatch, 'run') as run:
                dispatch.execute('go', args)
            command = run.call_args.args[0]
            self.assertEqual(command[command.index('--suite') + 1], 'JsonDecoding')
            self.assertEqual(command[command.index('--workspace') + 1], base / 'builds/first')
            with self.assertRaises(ValueError):
                dispatch.json_workspace(base, 'go', True, 'JsonTypedAst')
            pointer.write_text(json.dumps({'runtime': 'go', 'workspace': 'builds/legacy'}))
            self.assertEqual(dispatch.json_workspace(base, 'go', True), base / 'builds/legacy')
            with self.assertRaises(ValueError):
                dispatch.json_workspace(base, 'go', True, 'JsonDecoding')

    def test_root_runner_selects_only_go_and_js(self):
        with tempfile.TemporaryDirectory() as temporary:
            root = Path(temporary)
            (root / "bin").mkdir()
            shutil.copyfile(ROOT / "bin/run", root / "bin/run")
            for runtime in ["go", "js"]:
                (root / "bin" / runtime).mkdir()
                run = root / "bin" / runtime / "run"
                run.write_text(f'#!/bin/sh\nprintf "{runtime}: %s\\n" "$*"\n')
                run.chmod(0o755)
            for test in ["ArrayIndexing", "Test.JsonTypedAst", "JsonDecoding"]:
                result = subprocess.run(["bash", root / "bin/run", "--test", test, "--run-only"],
                                        capture_output=True, text=True, check=True)
                self.assertIn(f"go: --test {test} --run-only", result.stdout)
                self.assertIn(f"js: --test {test} --run-only", result.stdout)


class OptionalSources(unittest.TestCase):
    def test_regular_snapshots_exclude_only_separate_json_diagnostics(self):
        with tempfile.TemporaryDirectory() as temporary:
            root = Path(temporary) / "repo"
            paths = ["src/App.purs", "srx/AppX.purs", "srx/Test/STArray.purs",
                     "src/Test/JsonTypedAst.purs", "src/Test/JsonTypedAst.go",
                     "src/Test/JsonTypedAst.js",
                     "src/Test/JsonDecoding.purs", "src/Test/JsonDecoding.go", "src/Test/JsonDecoding.js"]
            paths += [f"run/bak/{runtime}/spago.{runtime}.yaml" for runtime in ["go", "php"]]
            paths += [f"run/bak/php/{name}" for name in ["composer.json", "composer.lock"]]
            for relative in paths:
                path = root / relative
                path.parent.mkdir(parents=True, exist_ok=True)
                path.write_text("{}")
            before = {path.relative_to(root): path.read_bytes() for path in root.rglob('*') if path.is_file()}
            for mode in ["pure", "test", "x"]:
                args = argparse.Namespace(mode=mode, clean=False, test="Fib")
                for runtime, driver in [("go", native), ("php", php), ("js", js)]:
                    with self.subTest(runtime=runtime, mode=mode):
                        directory = Path(temporary) / runtime / mode
                        with patch.object(driver, "ROOT", root):
                            if runtime == "go":
                                driver.prepare(directory, runtime, args)
                            elif runtime == "js":
                                driver.snapshot_sources(directory, args)
                            else:
                                driver.prepare(directory, args)
                        self.assertTrue((directory / "src/App.purs").is_file())
                        self.assertEqual((directory / "src/Test/STArray.purs").is_file(), mode == "x")
                        self.assertFalse(list((directory / "src").rglob("JsonTypedAst*")))
                        self.assertFalse(list((directory / "src").rglob("JsonDecoding*")))
            self.assertEqual(before, {path.relative_to(root): path.read_bytes()
                                      for path in root.rglob('*') if path.is_file()})

    def test_legacy_go_snapshots_exclude_json_without_moving_sources(self):
        with tempfile.TemporaryDirectory() as temporary:
            root = Path(temporary)
            runner = root / 'bin/go/test'
            runner.parent.mkdir(parents=True)
            shutil.copyfile(ROOT / 'bin/go/test', runner)
            sources = {
                'src/Test/Fib.purs': 'module Test.Fib where\n',
                'src/Test/JsonTypedAst.purs': 'module Test.JsonTypedAst where\n',
                'src/Test/JsonDecoding.purs': 'module Test.JsonDecoding where\n',
                'src/Test/JsonTypedAst.go': 'typed foreign',
                'src/Test/JsonDecoding.js': 'general foreign',
                'srx/Test/STArray.purs': 'module Helper where\n',
                'src/AppX.purs': 'user AppX',
                'run/bak/go/spago.go.yaml': '{}',
            }
            for relative, contents in sources.items():
                path = root / relative
                path.parent.mkdir(parents=True, exist_ok=True)
                path.write_text(contents)
            spago = root / 'run/bak/js/node_modules/.bin/spago'
            spago.parent.mkdir(parents=True)
            # Record compiler arguments and stop before invoking a real backend.
            spago.write_text('#!/bin/sh\nprintf "%s\\n" "$@" >> spago-arguments\nexit 1\n')
            spago.chmod(0o755)
            result = subprocess.run(['bash', runner], capture_output=True, text=True)
            self.assertEqual(result.returncode, 1, result.stderr)
            self.assertEqual([line for line in result.stdout.splitlines() if line.startswith('=== Testing')],
                             ['=== Testing Test.Fib ==='])
            arguments = (root / 'spago-arguments').read_text().splitlines()
            self.assertEqual(arguments[:3], ['build', '-q', '--purs-args'])
            self.assertEqual(arguments[3].split(), [
                '--exclude-files', 'src/Test/JsonTypedAst.purs',
                '--exclude-files', 'src/Test/JsonTypedAst/**/*.purs',
                '--exclude-files', 'src/Test/JsonDecoding.purs'])
            for relative, contents in sources.items():
                self.assertEqual((root / relative).read_text(), contents)



class JsonSuites(unittest.TestCase):
    def test_general_suite_has_no_pbo_dependency_or_source_fingerprint(self):
        with tempfile.TemporaryDirectory() as temporary:
            base = Path(temporary)
            root, compiler, pbo = base / 'repo', base / 'gopurs/gopurs', base / 'pbo'
            sources = root / 'src/Test'
            paths = [compiler / 'bin/gopurs-native', pbo / 'src/Resolver.purs',
                     compiler.parent / 'gopurs-argonaut-core/spago.yaml',
                     compiler.parent / 'gopurs-argonaut-core/src/Json.go']
            typed_sources = [sources / ('JsonTypedAst.' + ext) for ext in ['purs', 'go', 'js']]
            paths += typed_sources + [sources / ('JsonDecoding.' + ext) for ext in ['purs', 'go', 'js']]
            paths += [root / f'test/fixtures/{suite}/expected.json' for suite in ['json-decoding', 'json-typed-ast']]
            for path in paths:
                path.parent.mkdir(parents=True, exist_ok=True)
                path.write_text('{}')
            with patch.multiple(json_diagnostic, ROOT=root, COMPILER=compiler, PBO=pbo,
                                SOURCES=sources, SOURCE_FILES=typed_sources,
                                FIXTURES=root / 'test/fixtures/json-typed-ast'):
                general = json_diagnostic.fingerprint('JsonDecoding')
                typed = json_diagnostic.fingerprint()
                self.assertFalse(any(str(pbo) in path or 'JsonTypedAst' in path for path in general))
                self.assertNotIn('backend-optimizer', json_diagnostic.dependencies('JsonDecoding'))
                self.assertIn('backend-optimizer', json_diagnostic.dependencies('JsonTypedAst'))
                (pbo / 'src/Resolver.purs').write_text('changed')
                self.assertEqual(general, json_diagnostic.fingerprint('JsonDecoding'))
                self.assertNotEqual(typed, json_diagnostic.fingerprint())
                (compiler.parent / 'gopurs-argonaut-core/src/Json.go').write_text('changed')
                self.assertNotEqual(general, json_diagnostic.fingerprint('JsonDecoding'))

    def test_js_general_build_copies_only_its_sources_and_omits_pbo(self):
        with tempfile.TemporaryDirectory() as temporary:
            root = Path(temporary)
            sources, work = root / 'src/Test', root / 'work'
            sources.mkdir(parents=True)
            work.mkdir()
            for name in ['JsonDecoding', 'JsonTypedAst']:
                for ext in ['purs', 'go', 'js']:
                    (sources / f'{name}.{ext}').write_text('source')
            (work / 'spago.yaml').write_text('package:\n  name: test\n  dependencies:\n    - argonaut-codecs\nworkspace:\n')
            with patch.object(json_diagnostic, 'SOURCES', sources), patch.object(json_diagnostic, 'call'):
                json_diagnostic.build_js(work, {}, 'JsonDecoding')
            self.assertTrue((work / 'js/src/Test/JsonDecoding.purs').is_file())
            self.assertTrue((work / 'js/src/Test/JsonDecoding.js').is_file())
            self.assertFalse((work / 'js/src/Test/JsonDecoding.go').exists())
            self.assertFalse(list((work / 'js/src').rglob('JsonTypedAst*')))
            self.assertNotIn('backend-optimizer', (work / 'js/spago.yaml').read_text())
            self.assertIn('Test.JsonDecoding/index.js', (work / 'js/entry.mjs').read_text())

    def test_wrong_suite_build_is_rejected_and_legacy_tast_manifest_supported(self):
        json_diagnostic.check_suite({}, 'JsonTypedAst')
        json_diagnostic.check_suite({'suite': 'JsonDecoding'}, 'JsonDecoding')
        for manifest, suite in [({}, 'JsonDecoding'), ({'suite': 'JsonDecoding'}, 'JsonTypedAst')]:
            with self.subTest(manifest=manifest, suite=suite), self.assertRaises(ValueError):
                json_diagnostic.check_suite(manifest, suite)

    def test_general_results_validate_case_identity_and_correctness_before_timing(self):
        oracle = {'modules': 2, 'names': ['valid', 'wrong-field'], 'timed_cases': 1,
                  'fingerprints': [123, 456], 'json_fingerprints': [789, 101]}
        json_diagnostic.validate_result(dict(oracle), oracle, 'JsonDecoding')
        for key, wrong in [('modules', 1), ('names', ['wrong-field', 'valid']),
                           ('timed_cases', 2), ('fingerprints', [123, 0]),
                           ('json_fingerprints', [0, 101])]:
            with self.subTest(key=key), self.assertRaises(ValueError):
                json_diagnostic.validate_result(dict(oracle, **{key: wrong}), oracle, 'JsonDecoding')
        legacy = {key: value for key, value in oracle.items() if key not in ['names', 'timed_cases']}
        json_diagnostic.validate_result(legacy, legacy, 'JsonTypedAst')


if __name__ == "__main__":
    unittest.main()

#!/usr/bin/env python3
"""Check diagnostic routing and optional-source isolation without compiling."""
import argparse
from contextlib import redirect_stderr
import importlib.util
import io
import itertools
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


class Routing(unittest.TestCase):
    def test_detection_preserves_regular_tests(self):
        self.assertIsNone(dispatch.diagnostic_name(["--test", "Fib", "--pgo"]))
        self.assertEqual(dispatch.diagnostic_name(["--test=Test.JsonTypedAst"]), "JsonTypedAst")
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
            for test in ["ArrayIndexing", "Test.JsonTypedAst"]:
                result = subprocess.run(["bash", root / "bin/run", "--test", test, "--run-only"],
                                        capture_output=True, text=True, check=True)
                self.assertIn(f"go: --test {test} --run-only", result.stdout)
                self.assertIn(f"js: --test {test} --run-only", result.stdout)


class OptionalSources(unittest.TestCase):
    def test_extended_snapshots_exclude_only_pbo_diagnostic(self):
        with tempfile.TemporaryDirectory() as temporary:
            root = Path(temporary) / "repo"
            paths = ["src/App.purs", "srx/AppX.purs", "srx/Test/STArray.purs",
                     "srx/Test/JsonTypedAst.purs", "srx/Test/JsonTypedAst.go",
                     "srx/Test/JsonTypedAst.js", "srx/Test/JsonTypedAst/Fingerprint.purs"]
            paths += [f"run/bak/{runtime}/spago.{runtime}.yaml" for runtime in ["go", "php"]]
            paths += [f"run/bak/php/{name}" for name in ["composer.json", "composer.lock"]]
            for relative in paths:
                path = root / relative
                path.parent.mkdir(parents=True, exist_ok=True)
                path.write_text("{}")
            args = argparse.Namespace(mode="x", clean=False)
            for runtime, driver in [("go", native), ("php", php)]:
                directory = Path(temporary) / runtime
                with patch.object(driver, "ROOT", root):
                    if runtime == "go":
                        driver.prepare(directory, runtime, args)
                    else:
                        driver.prepare(directory, args)
                self.assertTrue((directory / "src/Test/STArray.purs").is_file())
                self.assertFalse(list((directory / "src").rglob("JsonTypedAst*")))


if __name__ == "__main__":
    unittest.main()

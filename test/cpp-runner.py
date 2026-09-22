"""Regression checks for C++ benchmark output validation; no benchmark is executed."""
import importlib.util
import hashlib
import json
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest
from types import SimpleNamespace
from unittest.mock import patch

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[1]
spec = importlib.util.spec_from_file_location("cpp_driver", ROOT / "bin/cpp/driver.py")
driver = importlib.util.module_from_spec(spec)
spec.loader.exec_module(driver)


class OutputValidation(unittest.TestCase):
    def setUp(self):
        self.expected = [driver.expectation("Test." + name) for name in driver.CASES]
        self.output = "".join(
            "(Test)\n" + case["label"] + "\n(Output & Warm-up)\n" + case["value"] +
            "\n(Execution time - best of 10)\n1000.00 μs\n"
            for case in self.expected) + "Total exec time: 14.00 ms\n"
        self.smoke = "".join("CPP_RESULT " + case["module"] + " " + case["value"] + "\n"
                             for case in self.expected)

    def test_valid_full_suite(self):
        self.assertEqual(driver.validate_output(self.output, self.expected, False)["total_ms"], 14)
        self.assertEqual(len(driver.validate_output(self.smoke, self.expected, True)["values"]), 14)

    def test_project_snapshot_excludes_json_diagnostics(self):
        class StopBeforeCompilation(Exception):
            pass

        with tempfile.TemporaryDirectory(prefix="altbak-cpp-sources-") as temporary:
            root = Path(temporary).resolve()
            paths = ['src/App.purs', 'src/Test/Fibonacci.purs',
                     'src/Test/JsonDecoding.purs', 'src/Test/JsonTypedAst.purs',
                     'src/Test/JsonTypedAst/Fingerprint.purs', 'deps/prelude/src/Prelude.purs']
            for relative in paths:
                path = root / relative
                path.parent.mkdir(parents=True, exist_ok=True)
                path.write_text('-- source snapshot fixture\n')
            work = root / 'build'
            args = SimpleNamespace(dependency_root=root / 'deps', purs='purs')

            def inspect(command, *_):
                inputs = [Path(value) for value in command if str(value).endswith('.purs')]
                self.assertIn(work / 'sources/project/Test/Fibonacci.purs', inputs)
                self.assertFalse(any('JsonTypedAst' in str(path) or 'JsonDecoding' in str(path)
                                     for path in inputs))
                raise StopBeforeCompilation

            with patch.object(driver, 'ROOT', root), patch.object(driver, 'version', return_value='0.14.4'), \
                    patch.object(driver, 'logged', side_effect=inspect), self.assertRaises(StopBeforeCompilation):
                driver.build(args, 'pure', 'App', self.expected, work)

    def test_wrong_result_and_label_are_rejected(self):
        for output in (self.output.replace("\n7\n", "\n999\n", 1),
                       self.output.replace(self.expected[0]["label"], "wrong label", 1)):
            with self.subTest(output=output[:80]), self.assertRaises(RuntimeError):
                driver.validate_output(output, self.expected, False)

    def test_missing_extra_and_malformed_results_are_rejected(self):
        for output in (self.output.replace("(Output & Warm-up)", "missing", 1),
                       self.output + self.output,
                       self.output.replace("1000.00 μs", "nan μs", 1),
                       self.output.replace("1000.00 μs", "-1.00 μs", 1)):
            with self.subTest(output=output[:80]), self.assertRaises(RuntimeError):
                driver.validate_output(output, self.expected, False)

    def test_total_unit_error_is_rejected(self):
        with self.assertRaises(RuntimeError):
            driver.validate_output(self.output.replace("14.00 ms", "0.014 ms"), self.expected, False)

    def test_incorrect_smoke_is_rejected(self):
        for output in (self.smoke.replace("Test.AstTree 7", "Test.AstTree 8"),
                       self.smoke + "CPP_RESULT Test.AstTree 7\n",
                       self.smoke.replace("CPP_RESULT Test.AstTree 7\n", "")):
            with self.subTest(output=output[:80]), self.assertRaises(RuntimeError):
                driver.validate_output(output, self.expected, True)

    def test_cli_exits_nonzero_for_wrong_native_output(self):
        with tempfile.TemporaryDirectory(prefix="altbak-cpp-runner-test-") as directory:
            folder = Path(directory)
            binary = folder / "fake-benchmark"
            wrong_output = self.smoke.replace("Test.AstTree 7", "Test.AstTree 999")
            binary.write_text("#!/usr/bin/env python3\nprint(" + repr(wrong_output) + ")\n")
            binary.chmod(0o700)
            (folder / "build.json").write_text(json.dumps({
                "mode": "pure", "expected": self.expected, "binary": str(binary),
                "binary_sha256": hashlib.sha256(binary.read_bytes()).hexdigest()}))
            result = subprocess.run([str(ROOT / "bin/cpp/run"), "--run-only", "--smoke",
                                     "--build-dir", directory], capture_output=True, text=True)
            self.assertEqual(result.returncode, 1)
            self.assertIn("Unexpected smoke result", result.stderr)


if __name__ == "__main__":
    unittest.main()

"""Regression checks for the shared benchmark validator, without running kernels."""
import importlib.util
from pathlib import Path
import subprocess
import sys
import unittest

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[1]
SPEC = importlib.util.spec_from_file_location("benchmark_validation", ROOT / "bin/benchmark/validate.py")
VALIDATOR = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(VALIDATOR)


def fixture(mode="pure", test=None):
    cases = VALIDATOR.expected_cases(mode, test)
    durations = [max(1000.0, case.get("minimum_us", 0)) for case in cases]
    output = "".join("(Test)\n" + case["label"] + "\n(Output & Warm-up)\n\n" +
                     case["value"] + f"\n\n(Execution time - best of 10)\n\n{duration:.2f} μs\n"
                     for case, duration in zip(cases, durations))
    return output + ("" if mode == "test" else f"Total exec time: {sum(durations) / 1000:.2f} ms\n")


class ValidationTests(unittest.TestCase):
    def test_all_core_modes(self):
        for mode in ["pure", "ffi", "fficc"]:
            result = VALIDATOR.validate_output(fixture(mode), mode)
            self.assertTrue(result["values_validated"])
            self.assertEqual(result["times_us"], [1000] * 14)
            self.assertEqual(result["total_ms"], 14)

    def test_wrong_nominal_result(self):
        with self.assertRaisesRegex(ValueError, "expected '7'"):
            VALIDATOR.validate_output(fixture().replace("\n7\n", "\n8\n", 1))

    def test_wrong_mode_or_order(self):
        with self.assertRaisesRegex(ValueError, "label"):
            VALIDATOR.validate_output(fixture("ffi"), "pure")
        with self.assertRaisesRegex(ValueError, "label"):
            VALIDATOR.validate_output(fixture().replace("AST Evaluation:", "Fibonacci:", 1))

    def test_partial_duplicate_and_malformed_rows(self):
        output = fixture()
        for invalid in [output + output, output.replace("(Test)", "missing", 1),
                        output.replace("(Output & Warm-up)", "missing", 1),
                        output.replace("1000.00 μs", "nan μs", 1),
                        output.replace("1000.00 μs", "-1 μs", 1)]:
            with self.subTest(invalid=invalid[:80]), self.assertRaises(ValueError):
                VALIDATOR.validate_output(invalid)

    def test_incorrect_total_units(self):
        for total in ["0.014", "14000", "nan", "-14"]:
            with self.subTest(total=total), self.assertRaises(ValueError):
                VALIDATOR.validate_output(fixture().replace("14.00 ms", total + " ms"))

    def test_single_core_test(self):
        result = VALIDATOR.validate_output(fixture("test", "StateMonadFFI"), "test", "StateMonadFFI")
        self.assertEqual(result["values"], ["1200"])
        self.assertIsNone(result["total_ms"])

    def test_extended_results_use_their_own_oracles(self):
        result = VALIDATOR.validate_output(fixture("x"), "x")
        self.assertTrue(result["values_validated"])
        self.assertEqual(result["oracle"], "extended")
        self.assertEqual(len(result["times_us"]), 5)

    def test_scientific_notation(self):
        result = VALIDATOR.validate_output(fixture().replace("1000.00 μs", "1e3 μs"))
        self.assertEqual(result["total_ms"], 14)

    def test_cli_propagates_failure_and_success(self):
        command = [sys.executable, str(ROOT / "bin/benchmark/validate.py")]
        bad = subprocess.run(command, input=fixture().replace("\n7\n", "\n8\n", 1),
                             text=True, capture_output=True)
        self.assertEqual(bad.returncode, 1)
        self.assertIn("expected '7'", bad.stderr)
        good = subprocess.run(command, input=fixture(), text=True, capture_output=True)
        self.assertEqual(good.returncode, 0)
        self.assertEqual(good.stdout, fixture())


if __name__ == "__main__":
    unittest.main()

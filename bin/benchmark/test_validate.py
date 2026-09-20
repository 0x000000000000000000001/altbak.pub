"""Regression checks for workload oracles; no benchmark programs are run."""
import unittest

from validate import expected_cases, validate_output


def report(mode="x", replacements=None, durations=None):
    replacements = replacements or {}
    durations = durations or {}
    rows, total = [], 0.0
    for case in expected_cases(mode):
        elapsed = durations.get(case["module"], 10000.0)
        value = replacements.get(case["module"], case["value"])
        rows.append(f"(Test)\n{case['label']}\n(Output & Warm-up)\n{value}\n"
                    f"(Execution time - best of 10)\n{elapsed:.6f} μs\nBatch iterations: 2\n")
        total += elapsed / 1000.0
    return "\n".join(rows) + f"\nTotal exec time: {total:.6f} ms\n"


class BenchmarkValidationTests(unittest.TestCase):
    def test_extended_results_are_validated(self):
        result = validate_output(report(), "x")
        self.assertTrue(result["values_validated"])
        self.assertEqual(result["oracle"], "extended")
        self.assertEqual(result["values"], ["10000", "10", "2000", "10", "Checksum: 679142946"])

    def test_each_extended_result_has_an_oracle(self):
        for case in expected_cases("x"):
            with self.subTest(module=case["module"]), self.assertRaises(ValueError):
                validate_output(report(replacements={case["module"]: "wrong"}), "x")

    def test_old_empty_file_result_is_rejected(self):
        with self.assertRaises(ValueError):
            validate_output(report(replacements={"Test.FileOps": ""}), "x")

    def test_old_overflowing_parallel_sum_is_rejected(self):
        for value in ["Sum of results: 2679142960", "Sum of results: -1615824336"]:
            with self.subTest(value=value), self.assertRaises(ValueError):
                validate_output(report(replacements={"Test.Parallelism": value}), "x")

    def test_no_op_delay_is_rejected_even_with_correct_result(self):
        with self.assertRaises(ValueError):
            validate_output(report(durations={"Test.AffOperations": 1.0}), "x")

    def test_timer_clock_granularity_is_allowed(self):
        validate_output(report(durations={"Test.AffOperations": 9000.0}), "x")

    def test_nonfinite_timing_is_rejected(self):
        with self.assertRaises(ValueError):
            validate_output(report(durations={"Test.StringOps": float("nan")}), "x")

    def test_core_report_and_batch_suffix_remain_supported(self):
        result = validate_output(report("pure"))
        self.assertEqual(len(result["values"]), 14)
        self.assertTrue(result["values_validated"])
        self.assertEqual(result["oracle"], "core")


if __name__ == "__main__":
    unittest.main()

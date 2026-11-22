#!/usr/bin/env python3
"""Validate WIP logs and publication contracts without executing benchmarks."""
from contextlib import redirect_stdout
import copy
import hashlib
import importlib.util
import io
import json
from pathlib import Path
import re
import sys
import tempfile
import unittest
from unittest.mock import patch

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT / "bin/benchmark"))
import validate

SPEC = importlib.util.spec_from_file_location("library_publication", ROOT / "bin/benchmark/publish.py")
publish = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(publish)

COLUMNS = {
    "JavaScript": ["js-pure", "es-pure", "js-ffi", "js-fficc"],
    "Go": ["go-pure", "go-dev", "psgo-pure", "go-ffi", "go-fficc"],
}


def log_section(mode):
    cases = validate.expected_cases(mode)
    rows = [f"(Test)\n{case['label']}\n(Output & Warm-up)\n{case['value']}\n"
            "(Execution time - best of 10)\n1000.000000 μs\nBatch iterations: 32\n"
            for case in cases]
    return "\n".join(rows) + ("" if mode == "wip" else f"Total exec time: {len(cases)}.000000 ms\n")


def aggregate(mode, offset):
    cases = validate.expected_cases(mode)
    times = [offset + index + 1.0 for index in range(len(cases))]
    result = {
        "mode": mode, "labels": [case["label"] for case in cases],
        "values": [case["value"] for case in cases], "values_validated": True,
        "times_us": times, "total_ms": None if mode == "wip" else sum(times) / 1000,
    }
    result["process_runs"] = []
    for factor in [1.5, 0.5, 1.0]:
        sample = {key: copy.deepcopy(value) for key, value in result.items() if key != "process_runs"}
        sample["times_us"] = [value * factor for value in times]
        sample["total_ms"] = None if mode == "wip" else sum(sample["times_us"]) / 1000
        result["process_runs"].append(sample)
    return result


def campaign_results():
    results = {}
    for column, key in enumerate(key for keys in COLUMNS.values() for key in keys):
        mode = "fficc" if key.endswith("-fficc") else "ffi" if key.endswith("-ffi") else "pure"
        results[key] = aggregate(mode, 50.0)
        results[key]["library"] = aggregate("wip", 100.0 * column)
    return results


class LibraryValidation(unittest.TestCase):
    def test_foreign_modes_keep_core_total_and_validate_all_wip_oracles(self):
        for mode in ["ffi", "fficc"]:
            with self.subTest(mode=mode):
                output = log_section(mode) + validate.LIBRARY_MARKER + "\n" + log_section("wip")
                result = validate.validate_output(output, mode, include_library=True)
                self.assertEqual(result["total_ms"], 14.0)
                self.assertEqual(len(result["values"]), 14)
                self.assertEqual(len(result["library"]["values"]), 23)
                self.assertEqual(result["library"]["values"][0], "75025000")
                self.assertIsNone(result["library"]["total_ms"])
                self.assertIsNone(result["library"]["sum_displayed_lines_ms"])

    def test_required_missing_or_corrupt_wip_is_rejected(self):
        for mode in ["ffi", "fficc"]:
            core = log_section(mode)
            with self.subTest(mode=mode, defect="missing"), self.assertRaisesRegex(ValueError, "Missing excluded WIP"):
                validate.validate_output(core, mode, include_library=True)
            for defect in [log_section("wip").replace("\n75025000\n", "\n75025001\n", 1),
                           log_section("wip").replace("\n75025000\n", "\n\n", 1)]:
                with self.subTest(mode=mode, defect="oracle"), self.assertRaises(ValueError):
                    validate.validate_output(core + validate.LIBRARY_MARKER + "\n" + defect,
                                             mode, include_library=True)


class LibraryPublication(unittest.TestCase):
    def publish_fixture(self, directory, results, original):
        campaign = directory / "campaign"
        campaign.mkdir()
        payload = json.dumps(results).encode()
        (campaign / "results.json").write_bytes(payload)
        source_snapshot = {"fixture": "frozen"}
        (campaign / "manifest.json").write_text(json.dumps({
            "completed_at_utc": "fixture-only", "sources": source_snapshot,
            "protocol": {"processes": 3}, "results_sha256": hashlib.sha256(payload).hexdigest(),
        }))
        readme = directory / "README.md"
        readme.write_text(original)
        arguments = ["publish", "--wip-only", "--campaign", str(campaign), "--readme", str(readme)]
        with patch.object(sys, "argv", arguments), patch.object(publish, "fingerprints", return_value=source_snapshot), \
                redirect_stdout(io.StringIO()):
            publish.main()
        return readme.read_text()

    def test_nine_columns_publish_23_rows_and_derived_totals(self):
        original = (ROOT / "README.md").read_text()
        results = campaign_results()
        with tempfile.TemporaryDirectory() as temporary:
            rendered = self.publish_fixture(Path(temporary), results, original)
        labels = [case["label"] for case in validate.expected_cases("wip")]
        original_lines = original.splitlines()
        rendered_lines = rendered.splitlines()
        self.assertEqual(len(original_lines), len(rendered_lines))
        # The derived rows must agree with the displayed cells.
        self.assertEqual(publish.refresh_new_totals(rendered), rendered)
        changed = set()
        for title, keys in COLUMNS.items():
            start = next(i for i, line in enumerate(original_lines)
                         if line.strip() == "#### " + title)
            body = re.search(r"#### " + re.escape(title) + r"\n(.*?)(?=\n#### )",
                             rendered, re.S)[1].splitlines()
            def full(local):
                return start + 1 + local
            for label in labels:
                local = next(i for i, line in enumerate(body)
                             if line.startswith("| [" + label + "]("))
                values = [results[key]["library"]["times_us"][labels.index(label)]
                          for key in keys]
                self.assertEqual(
                    [cell.strip() for cell in rendered_lines[full(local)].split("|")[2:-1]],
                    [f"~ {value:.2f} μs" for value in values])
                self.assertEqual(
                    [cell.strip() for cell in rendered_lines[full(local + 1)].split("|")[2:-1]],
                    ["—"] * len(keys))
                js_time = results["js-pure"]["library"]["times_us"][labels.index(label)]
                expected_js = [
                    "—" if title == "JavaScript" and column == 0
                    else f"{value / js_time:.2f}x"
                    for column, value in enumerate(values)]
                self.assertEqual(
                    [cell.strip() for cell in rendered_lines[full(local + 2)].split("|")[2:-1]],
                    expected_js)
                changed.update({full(local), full(local + 1), full(local + 2)})
            for total in ("**Total Incl WIP**", "**Total Stable + Incl WIP**"):
                local = next(i for i, line in enumerate(body)
                             if line.startswith("| " + total + " |"))
                changed.update({full(local), full(local + 1), full(local + 2), full(local + 3)})
        for index, (old, new) in enumerate(zip(original_lines, rendered_lines)):
            if index not in changed:
                self.assertEqual(new, old, f"unrelated line changed: {index + 1}")

    def test_incomplete_or_invalid_campaign_never_rewrites_readme(self):
        original = (ROOT / "README.md").read_text()
        mutations = {
            "missing column": lambda result: result.pop("go-dev"),
            "wrong median": lambda result: result["js-ffi"]["library"]["times_us"].__setitem__(0, 99999.0),
            "wrong oracle": lambda result: result["go-fficc"]["library"]["values"].__setitem__(0, "0"),
            "missing process": lambda result: result["psgo-pure"]["library"]["process_runs"].pop(),
        }
        for name, mutate in mutations.items():
            with self.subTest(defect=name), tempfile.TemporaryDirectory() as temporary:
                directory = Path(temporary)
                results = campaign_results()
                mutate(results)
                with self.assertRaises((ValueError, KeyError)):
                    self.publish_fixture(directory, results, original)
                self.assertEqual((directory / "README.md").read_text(), original)

    def test_missing_readme_row_is_rejected_without_partial_publication(self):
        original = (ROOT / "README.md").read_text()
        malformed = re.sub(r"(?m)^\| \[array/map-10000\][^\n]*\n", "", original, count=1)
        with tempfile.TemporaryDirectory() as temporary:
            directory = Path(temporary)
            with self.assertRaises(ValueError):
                self.publish_fixture(directory, campaign_results(), malformed)
            self.assertEqual((directory / "README.md").read_text(), malformed)


if __name__ == "__main__":
    unittest.main()

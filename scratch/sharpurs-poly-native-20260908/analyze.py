#!/usr/bin/env python3
"""Compare official App.main and allocation diagnostics without running either."""

import importlib.util
import json
import math
import statistics
from pathlib import Path


DIRECTORY = Path(__file__).resolve().parent
PROJECT = DIRECTORY.parents[1]
SPEC = importlib.util.spec_from_file_location(
    "regression_analysis", DIRECTORY.parent / "sharpurs-regression-check-20260908/analyze.py"
)
COMMON = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(COMMON)
require = COMMON.require
stats = COMMON.stats
POLY = "Polymorphism (10M Type Class Dict Lookups)"


def compare(before, after):
    require(before > 0 and after > 0, "Positive times required for performance ratios")
    return {
        "before": before,
        "after": after,
        "difference": after - before,
        "difference_percent": (after / before - 1) * 100,
        "reduction_percent": (1 - after / before) * 100,
        "speedup_factor": before / after,
    }


def read_diagnostic(path):
    records = [json.loads(line) for line in path.read_text(encoding="utf-8").splitlines() if line.strip()]
    require(
        [record.get("kind") for record in records]
        == ["metadata"] + ["warmup"] * 3 + ["measurement"] * 10 + ["summary"],
        f"{path}: expected metadata, 3 warmups, 10 measurements and summary, in that order",
    )
    metadata, warmups, measurements, summary = records[0], records[1:4], records[4:14], records[14]
    expected_metadata = {
        "globalWarmups": 3,
        "warmups": 3,
        "measurements": 10,
        "serverGc": False,
        "forcedGc": False,
        "allocationScope": "current worker thread",
        "gcScope": "process",
        "action": "Test.Polymorphism.act",
        "input": 10000000,
    }
    for key, expected in expected_metadata.items():
        require(metadata.get(key) == expected, f"{path}: unexpected metadata {key}: {metadata.get(key)!r}")
    require([item["iteration"] for item in warmups] == list(range(1, 4)), f"{path}: warmup iteration mismatch")
    require([item["iteration"] for item in measurements] == list(range(1, 11)), f"{path}: measurement iteration mismatch")
    require(
        all(item.get("output") == "10000000" for item in warmups + measurements + [summary]),
        f"{path}: incorrect Polymorphism output",
    )
    for item in measurements:
        require(math.isfinite(item["elapsedUs"]) and item["elapsedUs"] > 0, f"{path}: invalid time")
        for key in ("allocatedBytes", "gc0", "gc1", "gc2"):
            require(type(item[key]) is int and item[key] >= 0, f"{path}: invalid {key}")
    times = [item["elapsedUs"] for item in measurements]
    allocations = [item["allocatedBytes"] for item in measurements]
    recomputed = {
        "bestUs": min(times),
        "medianUs": statistics.median(times),
        "minAllocatedBytes": min(allocations),
        "maxAllocatedBytes": max(allocations),
        **{key: sum(item[key] for item in measurements) for key in ("gc0", "gc1", "gc2")},
    }
    for key, expected in recomputed.items():
        require(
            key in summary and math.isclose(summary[key], expected, rel_tol=1e-10, abs_tol=1e-9),
            f"{path}: reported {key} does not match samples",
        )
    return {
        "log": str(path),
        "metadata": metadata,
        "warmups": warmups,
        "measurements": measurements,
        "summary": summary,
    }


def diagnostic_stats(runs):
    samples = [sample for run in runs for sample in run["measurements"]]
    return {
        "processes": len(runs),
        "sample_count": len(samples),
        "process_best_us": stats([run["summary"]["bestUs"] for run in runs]),
        "process_median_us": stats([run["summary"]["medianUs"] for run in runs]),
        "all_samples_us": stats([sample["elapsedUs"] for sample in samples]),
        "allocation_bytes_per_action": stats([sample["allocatedBytes"] for sample in samples]),
        "allocation_bytes_process_ranges": [
            {"min": run["summary"]["minAllocatedBytes"], "max": run["summary"]["maxAllocatedBytes"]}
            for run in runs
        ],
        "gc_totals_across_30_samples": {
            key: sum(sample[key] for sample in samples) for key in ("gc0", "gc1", "gc2")
        },
    }


def main():
    historical = [
        COMMON.read_run(PROJECT / "scratch/sharpurs-tco-measure-20260908" / relative)
        for relative in ("official-after.log", "isolated/official-after.log")
    ]
    COMMON.check_outputs(historical[1], historical[0])
    names = [test["name"] for test in historical[0]["tests"]]
    baseline = COMMON.read_readme_baseline(PROJECT.parent / "altbak.pub/README.md", names)
    require(baseline["tests_us"][POLY] == 9689415, "README Polymorphism baseline changed; review assumptions")

    official = {
        variant: [COMMON.read_run(DIRECTORY / f"official-{variant}-{index}.log") for index in range(1, 6)]
        for variant in ("before", "after")
    }
    for runs in official.values():
        for run in runs:
            COMMON.check_outputs(run, historical[0])
    official_stats = {
        variant: {
            "total_ms": stats([run["total_ms"] for run in runs]),
            "tests_us": {
                name: stats([run["tests"][index]["best_of_10_us"] for run in runs])
                for index, name in enumerate(names)
            },
        }
        for variant, runs in official.items()
    }
    official_changes = {
        "total_ms": compare(official_stats["before"]["total_ms"]["median"], official_stats["after"]["total_ms"]["median"]),
        "tests_us": {
            name: compare(official_stats["before"]["tests_us"][name]["median"], official_stats["after"]["tests_us"][name]["median"])
            for name in names
        },
    }

    diagnostics = {
        variant: [read_diagnostic(DIRECTORY / f"diagnostic-{variant}-{index}.jsonl") for index in range(1, 4)]
        for variant in ("before", "after")
    }
    expected_metadata = {key: value for key, value in diagnostics["before"][0]["metadata"].items() if key != "pid"}
    for runs in diagnostics.values():
        for run in runs:
            require(
                {key: value for key, value in run["metadata"].items() if key != "pid"} == expected_metadata,
                f"{run['log']}: process environment or diagnostic protocol differs",
            )
    diagnostic_summaries = {variant: diagnostic_stats(runs) for variant, runs in diagnostics.items()}
    diagnostic_changes = {
        key: compare(diagnostic_summaries["before"][key]["median"], diagnostic_summaries["after"][key]["median"])
        for key in ("process_best_us", "process_median_us", "allocation_bytes_per_action")
        if diagnostic_summaries["after"][key]["median"] > 0
    }
    after = official_stats["after"]
    results = {
        "validation": {
            "passed": True,
            "official_processes_per_variant": 5,
            "official_tests_per_process": 14,
            "official_outputs_and_order_match_both_historical_logs": True,
            "official_sum_tolerance_ms": 0.02,
            "diagnostic_processes_per_variant": 3,
            "diagnostic_global_warmups_per_process": 3,
            "diagnostic_local_warmups_per_process": 3,
            "diagnostic_samples_per_process": 10,
            "diagnostic_outputs_all_equal_10000000": True,
            "diagnostic_summaries_recomputed": True,
            "diagnostic_metadata_equal_except_pid": True,
            "diagnostic_server_gc": False,
            "diagnostic_forced_gc": False,
        },
        "method": {
            "official": "True App.main: best of 10 per test in each fresh process, median across 5 processes.",
            "diagnostic": "True Test.Polymorphism.act after 3 App.warmup and 3 local warmups; 10 actions per fresh process, 3 processes per variant.",
            "diagnostic_time_statistics": "Report median of each process's best time, and median of each process's median time separately.",
            "allocations": "Bytes per action on current worker thread; GC collections process-wide; validation and serialization outside measured snapshots.",
            "deltas": "after minus before; negative percentage means faster. Speedup is before divided by after.",
            "caution": "Descriptive statistics only. Per-test medians need not sum to median total; approximate README historical data is context, not a simultaneous control.",
        },
        "official": {"runs": official, "summaries": official_stats, "changes": official_changes},
        "diagnostic": {"runs": diagnostics, "summaries": diagnostic_summaries, "changes": diagnostic_changes},
        "historical_local_runs": historical,
        "official_readme_baseline": baseline,
        "after_vs_official_readme": {
            "total_ms": compare(baseline["total_ms"], after["total_ms"]["median"]),
            "polymorphism_us": compare(baseline["tests_us"][POLY], after["tests_us"][POLY]["median"]),
        },
    }
    output = DIRECTORY / "comparison.json"
    output.write_text(json.dumps(results, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")
    print("Validated 10 App.main logs (14 outputs each) and 6 diagnostic logs (60 measurements + 18 local warmups).")
    for variant in ("before", "after"):
        total = official_stats[variant]["total_ms"]
        poly = official_stats[variant]["tests_us"][POLY]
        diag = diagnostic_summaries[variant]
        allocations = diag["allocation_bytes_per_action"]
        print(f"{variant}: official total median {total['median']:.2f} ms [{total['min']:.2f}, {total['max']:.2f}], Polymorphism {poly['median'] / 1000:.5f} ms")
        print(f"  diagnostic: median of best {diag['process_best_us']['median'] / 1000:.5f} ms, median of medians {diag['process_median_us']['median'] / 1000:.5f} ms")
        print(f"  allocations/action {allocations['min']}–{allocations['max']} bytes; GC across 30 samples {diag['gc_totals_across_30_samples']}")
    for label, change in (("Total", official_changes["total_ms"]), ("Polymorphism", official_changes["tests_us"][POLY])):
        print(f"Official {label}: {change['difference_percent']:+.2f}%, {change['speedup_factor']:.2f}× speedup")
    print(f"README historical compiled F#/C#: {baseline['total_ms']:.2f} ms total, {baseline['tests_us'][POLY]} µs Polymorphism")
    print(f"Wrote {output}")


if __name__ == "__main__":
    main()

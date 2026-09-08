#!/usr/bin/env python3
"""Validate and compare fresh-process App.main logs; does not run benchmarks."""

import argparse
import json
import re
import statistics
from pathlib import Path


BLOCK = re.compile(
    r"\(Test\)\s+(?P<name>[^\r\n]+):\s+"
    r"\(Output & Warm-up\)\s+(?P<output>.*?)\s+"
    r"\(Execution time - best of 10\)\s+"
    r"(?P<time>\d+(?:\.\d+)?)\s+[μµ]s",
    re.DOTALL,
)
TOTAL = re.compile(r"Total exec time:\s*(\d+(?:\.\d+)?)\s+ms")
FOCUS = {
    "Polymorphism": "Polymorphism (10M Type Class Dict Lookups)",
    "RBTree": "Red-Black Tree (100k Worst-Case Insertions)",
    "Lazy": "Lazy Evaluation (1M Thunks Forced, 1k Depth)",
    "TCO": "Tail Call Optimization (100k calls)",
}


def require(condition, message):
    if not condition:
        raise ValueError(message)


def read_run(path):
    content = path.read_text(encoding="utf-8")
    matches = list(BLOCK.finditer(content))
    require(len(matches) == 14, f"{path}: expected 14 timed tests, found {len(matches)}")
    require(content.count("(Test)") == 14, f"{path}: unexpected test sections")
    totals = TOTAL.findall(content)
    require(len(totals) == 1, f"{path}: expected exactly one total")
    tests = [
        {
            "name": match["name"].strip(),
            "output": match["output"].strip(),
            "best_of_10_us": float(match["time"]),
        }
        for match in matches
    ]
    require(len({test["name"] for test in tests}) == 14, f"{path}: duplicate test names")
    total_ms = float(totals[0])
    sum_ms = sum(test["best_of_10_us"] for test in tests) / 1000
    require(
        abs(total_ms - sum_ms) <= 0.02,
        f"{path}: total {total_ms} ms differs from test sum {sum_ms:.5f} ms",
    )
    require(total_ms > 0, f"{path}: nonpositive total")
    return {
        "log": str(path.resolve()),
        "total_ms": total_ms,
        "sum_of_test_times_ms": sum_ms,
        "rounding_difference_ms": total_ms - sum_ms,
        "tests": tests,
    }


def check_outputs(run, expected):
    signature = lambda record: [(test["name"], test["output"]) for test in record["tests"]]
    require(
        signature(run) == signature(expected),
        f"{run['log']}: test order, names or outputs differ from {expected['log']}",
    )


def stats(values):
    return {
        "n": len(values),
        "values": values,
        "median": statistics.median(values),
        "min": min(values),
        "max": max(values),
    }


def delta(reference, current):
    require(reference > 0, "A reference time must be positive to compute a percentage")
    return {"absolute": current - reference, "percent": (current / reference - 1) * 100}


def read_readme_baseline(path, names):
    text = path.read_text(encoding="utf-8")
    require("#### F#/C#\n" in text, f"{path}: F#/C# baseline section missing")
    section = text.split("#### F#/C#\n", 1)[1].split("\n#### ", 1)[0]
    rows = {}
    for line in section.splitlines():
        cells = [cell.strip() for cell in line.split("|")]
        if len(cells) < 2:
            continue
        measurement = re.fullmatch(r"~\s*([\d.]+)\s*([μµ]s|ms)", cells[1])
        if measurement:
            rows[cells[0].strip("*")] = (float(measurement[1]), measurement[2])
    require("Total Execution Time" in rows, f"{path}: compiled F#/C# total missing")
    total, unit = rows.pop("Total Execution Time")
    require(unit == "ms" and total == 13918.86, f"{path}: baseline changed; review analysis assumptions")
    benchmarks = {}
    for name in names:
        label = name.split(" (", 1)[0]
        require(label in rows, f"{path}: baseline for {name} missing")
        value, unit = rows[label]
        require(unit in ("μs", "µs"), f"{path}: unexpected unit for {label}")
        benchmarks[name] = value
    return {
        "source": str(path.resolve()),
        "section": "F#/C#: Compiled F#/C# (sharpurs)",
        "total_ms": total,
        "tests_us": benchmarks,
        "note": "Official historical figures are approximate and are context, not a simultaneous control.",
    }


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--directory", type=Path, default=Path(__file__).resolve().parent)
    parser.add_argument("--runs", type=int, default=5)
    args = parser.parse_args()
    require(args.runs >= 2, "At least two processes per variant are required")
    directory = args.directory.resolve()
    project = Path(__file__).resolve().parents[2]
    history_directory = project / "scratch/sharpurs-tco-measure-20260908"
    history = [
        read_run(history_directory / "official-after.log"),
        read_run(history_directory / "isolated/official-after.log"),
    ]
    expected = history[0]
    check_outputs(history[1], expected)
    names = [test["name"] for test in expected["tests"]]
    baseline = read_readme_baseline(project.parent / "altbak.pub/README.md", names)
    runs = {
        variant: [read_run(directory / f"{variant}-{index}.log") for index in range(1, args.runs + 1)]
        for variant in ("reference", "current")
    }
    for variant_runs in runs.values():
        for run in variant_runs:
            check_outputs(run, expected)
    summaries = {}
    for variant, variant_runs in runs.items():
        summaries[variant] = {
            "total_ms": stats([run["total_ms"] for run in variant_runs]),
            "tests_us": {
                name: stats([run["tests"][index]["best_of_10_us"] for run in variant_runs])
                for index, name in enumerate(names)
            },
        }
    reference = summaries["reference"]
    current = summaries["current"]
    differences = {
        "total_ms": delta(reference["total_ms"]["median"], current["total_ms"]["median"]),
        "tests_us": {
            name: delta(reference["tests_us"][name]["median"], current["tests_us"][name]["median"])
            for name in names
        },
    }
    current_total = current["total_ms"]["median"]
    comparison = {
        "validation": {
            "passed": True,
            "tests_per_run": 14,
            "processes_per_variant": args.runs,
            "names_order_and_outputs_match_both_historical_logs": True,
            "max_allowed_sum_rounding_difference_ms": 0.02,
        },
        "method": {
            "timing_within_each_process": "App.main execution time, best of 10 per test",
            "between_process_statistic": "Median of fresh-process values; min/max also reported",
            "delta_direction": "current minus reference; positive means slower",
            "caution": "Per-test medians need not sum to the median total. No significance claim is inferred from these descriptive statistics.",
        },
        "runs": runs,
        "summaries": summaries,
        "current_vs_reference": differences,
        "focus": {
            label: {
                "name": name,
                "reference_us": reference["tests_us"][name],
                "current_us": current["tests_us"][name],
                "difference_us": differences["tests_us"][name],
            }
            for label, name in FOCUS.items()
        },
        "historical_local_runs": history,
        "current_median_vs_historical_totals": [
            {"log": run["log"], "historical_total_ms": run["total_ms"], **delta(run["total_ms"], current_total)}
            for run in history
        ],
        "official_readme_baseline": baseline,
        "current_vs_official_readme": {
            "total_ms": delta(baseline["total_ms"], current_total),
            "total_speedup_factor": baseline["total_ms"] / current_total,
            "tests_us": {
                name: delta(baseline["tests_us"][name], current["tests_us"][name]["median"])
                for name in names
            },
        },
    }
    output = directory / "comparison.json"
    output.write_text(json.dumps(comparison, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")
    print(f"Validated {2 * args.runs} fresh-process logs: 14 outputs each; historical outputs match.")
    for variant in ("reference", "current"):
        item = summaries[variant]["total_ms"]
        print(f"{variant}: median {item['median']:.2f} ms, range {item['min']:.2f}–{item['max']:.2f} ms")
    print(f"Current total delta: {differences['total_ms']['percent']:+.2f}%")
    for label, name in FOCUS.items():
        before = reference["tests_us"][name]["median"] / 1000
        after = current["tests_us"][name]["median"] / 1000
        change = differences["tests_us"][name]["percent"]
        print(f"{label}: {before:.5f} → {after:.5f} ms ({change:+.2f}%)")
    print(f"Official README total: {baseline['total_ms']:.2f} ms; current {baseline['total_ms'] / current_total:.2f}× faster")
    print(f"Wrote {output}")


if __name__ == "__main__":
    main()

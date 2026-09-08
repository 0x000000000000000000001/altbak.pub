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
    require(unit == "ms" and total == 294.80, f"{path}: baseline changed; review analysis assumptions")
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


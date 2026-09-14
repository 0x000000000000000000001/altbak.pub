#!/usr/bin/env python3
"""Validate core benchmark results and timings, forwarding stdin unchanged."""
import argparse
import json
import math
from pathlib import Path
import re
import sys

ROOT = Path(__file__).resolve().parents[2]
CASES = dict(zip(
    ["AstTree", "Fib", "ListOps", "TCO", "Records", "Ackermann", "Church",
     "Primes", "RBTree", "Polymorphism", "StateMonad", "LazyEvaluation", "ArrayOps", "RowToList"],
    ["7", "55", "202950", "100000", "20000", "125", "100000", "21536",
     "22", "10000000", "1200", "1000000", "202950", "5"]))
ROW = re.compile(
    r"\(Test\)\s*\n\s*([^\n]+)\s*\n\s*\(Output & Warm-up\)\s*\n"
    r"(.*?)\(Execution time - best of 10\)\s*\n\s*(\S+)\s+μs", re.S)
EXTENDED_LABELS = ["File I/O (10k writes/reads):", "STArray Operations:",
                   "String Operations (1k Regex/Split):", "Aff Operations (Asynchronous Delays)",
                   "Parallelism (10 x Fib 42)"]


def case_for(module, expected=None):
    name = module.removeprefix("Test.")
    base = name.removesuffix("FFICheatcode").removesuffix("FFI")
    value = CASES.get(base) if expected is None else expected
    if value is None:
        raise ValueError(f"No core oracle for {module}; supply --expected VALUE for a custom --test")
    source = ROOT / "src/Test" / (name.replace(".", "/") + ".purs")
    if not source.is_file():
        raise ValueError(f"Missing source for {module}: {source}")
    label = re.search(r'(?m)^describe\s*=\s*(?:\w+\.)?log\s+"([^"\n]+)"', source.read_text())
    if not label:
        raise ValueError(f"Cannot find literal benchmark label in {source}")
    return {"module": "Test." + name, "label": label.group(1), "value": str(value)}


def expected_cases(mode, test=None, expected=None):
    if mode == "test":
        if not test or not re.fullmatch(r"(?:Test\.)?[A-Z][A-Za-z0-9_.]*", test):
            raise ValueError("--mode test requires a valid --test MODULE")
        return [case_for(test, expected)]
    if expected is not None or test:
        raise ValueError("--test and --expected require --mode test")
    if mode == "x":
        # Extended effects have backend/runtime-dependent results. Do not claim
        # that the core oracle validates these; only their timing structure is checked.
        return [{"label": label, "value": None} for label in EXTENDED_LABELS]
    if mode not in {"pure", "ffi", "fficc"}:
        raise ValueError(f"Unknown benchmark mode {mode}")
    suffix = {"pure": "", "ffi": "FFI", "fficc": "FFICheatcode"}[mode]
    return [case_for(name + suffix) for name in CASES]


def validate_output(output, mode="pure", test=None, expected=None):
    cases = expected_cases(mode, test, expected)
    rows = ROW.findall(output)
    if (len(rows) != len(cases) or output.count("(Test)") != len(cases)
            or output.count("(Output & Warm-up)") != len(cases)
            or output.count("(Execution time - best of 10)") != len(cases)):
        raise ValueError(f"Expected exactly {len(cases)} complete benchmark rows; found {len(rows)}")
    times, values, labels = [], [], []
    for (label, value, elapsed), case in zip(rows, cases):
        label, value = label.strip(), value.strip()
        if label != case["label"]:
            raise ValueError(f"Unexpected benchmark label {label!r}; expected {case['label']!r}")
        if case["value"] is not None and value != case["value"]:
            raise ValueError(f"{label} returned {value!r}; expected {case['value']!r}")
        duration = float(elapsed)
        if not math.isfinite(duration) or duration < 0:
            raise ValueError(f"Invalid duration for {label}: {elapsed}")
        labels.append(label)
        values.append(value)
        times.append(duration)
    totals = re.findall(r"Total exec time:\s*(\S+)\s+ms", output)
    if len(totals) != (0 if mode == "test" else 1):
        raise ValueError("Missing or repeated benchmark total")
    summed = sum(times) / 1000.0
    # At most .005 μs rounding per row and .005 ms for the printed total.
    tolerance = .005 + len(rows) * .005 / 1000 + 1e-8
    total = float(totals[0]) if totals else None
    if total is not None and (not math.isfinite(total) or total < 0 or abs(total - summed) > tolerance):
        raise ValueError(f"Printed total {total} ms differs from row sum {summed} ms")
    return {"mode": mode, "labels": labels, "values": values, "times_us": times,
            "total_ms": total, "sum_displayed_lines_ms": summed,
            "values_validated": mode != "x", "timer_unit": "microseconds"}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--mode", choices=["pure", "ffi", "fficc", "test", "x"], default="pure")
    parser.add_argument("--test")
    parser.add_argument("--expected")
    parser.add_argument("--input", type=Path, help="read an existing log instead of stdin")
    parser.add_argument("--output", type=Path, help="write validated results as JSON")
    args = parser.parse_args()
    chunks = []
    source = args.input.open() if args.input else sys.stdin
    try:
        for line in source:
            chunks.append(line)
            print(line, end="", flush=True)
        result = validate_output("".join(chunks), args.mode, args.test, args.expected)
        if args.output:
            args.output.parent.mkdir(parents=True, exist_ok=True)
            args.output.write_text(json.dumps(result, indent=2) + "\n")
        message = (f"{len(result['values'])} expected core results and timings verified"
                   if result["values_validated"] else "extended timing structure checked; core result oracle not applicable")
        print("Benchmark validation: " + message, file=sys.stderr)
    except (ValueError, OSError) as error:
        print("Benchmark validation failed: " + str(error), file=sys.stderr)
        return 1
    finally:
        if args.input:
            source.close()
    return 0


if __name__ == "__main__":
    sys.exit(main())

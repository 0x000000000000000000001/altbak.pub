#!/usr/bin/env python3
"""Measure prebuilt StateMonad variants in alternating AB/BA pairs.

No compiler/build is invoked. --output pilot.jsonl isolates pilot measurements,
raw logs and the checked result reference from the primary measures.jsonl set.
"""
import argparse
import datetime
import json
import math
import os
from pathlib import Path
import re
import subprocess
import time

ROOT = Path(__file__).resolve().parent
ENV = dict(os.environ, GOGC="800", GOMAXPROCS="14", GOMEMLIMIT="off", PPROF="0")
NAMES = ["AstTree", "Fib", "ListOps", "TCO", "Records", "Ackermann", "Church",
         "Primes", "RBTree", "Polymorphism", "StateMonad", "LazyEvaluation", "ArrayOps", "RowToList"]
SIZES = ["n20"]
CORE_RESULT = 1200
RSS_ITERATIONS = 10000
BENCH = re.compile(
    r"^Benchmark(StateMonad|Act)/([^\s]+)-14\s+(\d+)\s+([\d.eE+-]+)\s+ns/op"
    r"\s+([\d.eE+-]+)\s+B/op\s+([\d.eE+-]+)\s+allocs/op$", re.M)


def require(condition, message):
    if not condition:
        raise RuntimeError(message)


class Measurements:
    def __init__(self, output):
        self.output = output.resolve()
        primary = self.output.name == "measures.jsonl"
        self.logs = self.output.parent / ("logs" if primary else self.output.stem + "-logs")
        self.reference = self.output.parent / (
            "expected-results.json" if primary else self.output.stem + "-expected-results.json")
        self.keys = set()
        if self.output.exists():
            for line_no, line in enumerate(self.output.read_text().splitlines(), 1):
                if not line.strip():
                    continue
                row = json.loads(line)
                key = self.key(row)
                require(key not in self.keys,
                        f"{self.output}:{line_no}: duplicate measurement {key}")
                self.keys.add(key)

    @staticmethod
    def key(row):
        return (row["kind"], row["variant"], row["pair"], row.get("test"))

    def run(self, args, label, extra_env=None):
        self.logs.mkdir(parents=True, exist_ok=True)
        log = self.logs / (label + ".log")
        # Reserve before starting the process; do not overwrite a previous run.
        with log.open("x") as out:
            started = time.monotonic()
            result = subprocess.run(args, cwd=ROOT, env=dict(ENV, **(extra_env or {})),
                                    stdout=subprocess.PIPE, stderr=subprocess.STDOUT, text=True)
            out.write(result.stdout)
        require(result.returncode == 0,
                f"{label} failed ({result.returncode}); see {log}")
        print(f"{label}: {time.monotonic() - started:.1f}s", flush=True)
        return result.stdout

    def emit(self, row):
        key = self.key(row)
        require(key not in self.keys, f"duplicate measurement {key}; use another --output")
        row["recorded_utc"] = datetime.datetime.now(datetime.timezone.utc).isoformat()
        self.output.parent.mkdir(parents=True, exist_ok=True)
        with self.output.open("a") as out:
            out.write(json.dumps(row, allow_nan=False) + "\n")
        self.keys.add(key)

    def ensure_phase_new(self, kind):
        require(not any(key[0] == kind for key in self.keys),
                f"{kind} already exists in {self.output}; use another --output")


def order(pair):
    return ["before", "after"] if pair % 2 else ["after", "before"]


def check_results(measurements):
    reference = None
    for variant in ["before", "after"]:
        text = measurements.run(
            [str(ROOT / "bin" / (variant + "-bench")), "-test.run=^TestResults$", "-test.v"],
            variant + "-results")
        lines = [line for line in text.splitlines() if line.startswith("RESULT ")]
        require(len(lines) == 15, f"{variant}: expected 15 RESULT lines, got {len(lines)}")
        require(len(set(lines)) == 15, f"{variant}: duplicate RESULT lines")
        require(lines[0] == f"RESULT core n20 {CORE_RESULT}",
                f"{variant}: unexpected StateMonad core result: {lines[0]!r}")
        if reference is None:
            reference = lines
        else:
            require(lines == reference, "A/B outputs differ")
    acts = re.findall(r'^RESULT act (\w+) (".*")$', "\n".join(reference), re.M)
    require([name for name, _ in acts] == NAMES, "unexpected or duplicate act result inventory")
    expected = {name: json.loads(value) for name, value in acts}
    require(all(isinstance(value, str) for value in expected.values()), "act results must be strings")
    measurements.reference.write_text(json.dumps(expected, indent=2) + "\n")
    return expected


def benchmarks(measurements, kind, pairs, benchtime):
    name = "StateMonad" if kind == "core" else "Act"
    expected = SIZES if kind == "core" else NAMES
    measurements.ensure_phase_new(kind)
    for pair in range(1, pairs + 1):
        for variant in order(pair):
            text = measurements.run(
                [str(ROOT / "bin" / (variant + "-bench")), "-test.run=^$",
                 "-test.bench=^Benchmark" + name + "$", "-test.benchtime=" + benchtime,
                 "-test.count=1", "-test.cpu=14", "-test.benchmem"],
                f"{kind}-{pair:02}-{variant}")
            matches = BENCH.findall(text)
            require([(m[0], m[1]) for m in matches] == [(name, test) for test in expected],
                    f"{kind}/{variant}: unexpected or duplicate benchmark inventory; see raw log")
            for _, test, iterations, ns, bytes_, allocs in matches:
                values = [float(ns), float(bytes_), float(allocs)]
                require(int(iterations) > 0 and all(math.isfinite(value) for value in values)
                        and values[0] > 0 and all(value >= 0 for value in values[1:]),
                        f"{kind}/{variant}/{test}: invalid benchmark values")
                measurements.emit(dict(kind=kind, variant=variant, pair=pair, test=test,
                                       iterations=int(iterations), benchtime=benchtime,
                                       ns_per_op=values[0], bytes_per_op=values[1],
                                       allocs_per_op=values[2]))


def apps(measurements, pairs, expected):
    require(list(expected) == NAMES, "checked expected-results inventory is invalid")
    measurements.ensure_phase_new("app")
    for pair in range(1, pairs + 1):
        for variant in order(pair):
            text = measurements.run([str(ROOT / "bin" / (variant + "-app"))],
                                    f"app-{pair:02}-{variant}")
            chunks = text.split("(Test)")[1:]
            require(len(chunks) == len(NAMES), "application must output exactly 14 tests")
            tests = []
            for name, chunk in zip(NAMES, chunks):
                match = re.search(
                    r"\(Output & Warm-up\)\s*\n(.*?)\n\s*\(Execution time - best of 10\)"
                    r"\s*\n([\d.]+) μs", chunk, re.S)
                require(match is not None, f"{name}: missing application output/time")
                output, us = match.groups()
                require(output.strip() == expected[name],
                        f"{name}: output differs from checked reference")
                tests.append(dict(test=name, result=output.strip(), time_us=float(us)))
            total = re.findall(r"Total exec time: ([\d.]+) ms", text)
            require(len(total) == 1, "application total missing or duplicated")
            measurements.emit(dict(kind="app", variant=variant, pair=pair, tests=tests,
                                   total_ms=float(total[0])))


def rss(measurements, pairs):
    measurements.ensure_phase_new("rss")
    for pair in range(1, pairs + 1):
        for variant in order(pair):
            text = measurements.run(
                ["/usr/bin/time", "-l", str(ROOT / "bin" / (variant + "-bench")),
                 "-test.run=^TestFixedRSS$", "-test.v"], f"rss-n20-{pair:02}-{variant}",
                {"GOPURS_RSS_SIZE": "n20", "GOPURS_RSS_ITERATIONS": str(RSS_ITERATIONS)})
            values = re.findall(r"(\d+)\s+maximum resident set size", text)
            require(len(values) == 1 and int(values[0]) > 0,
                    "RSS measurement missing, duplicated or invalid")
            workloads = re.findall(r"^RSS workload (\w+) iterations=(\d+) result=(-?\d+)$",
                                   text, re.M)
            require(workloads == [("n20", str(RSS_ITERATIONS), str(CORE_RESULT))],
                    "RSS workload, iterations or result differs from the fixed contract")
            measurements.emit(dict(kind="rss", variant=variant, pair=pair, test="n20",
                                   iterations=RSS_ITERATIONS, result=CORE_RESULT,
                                   rss_bytes=int(values[0])))


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("phase", choices=["check", "core", "act", "app", "rss"])
    parser.add_argument("--pairs", type=int, default=10)
    parser.add_argument("--benchtime", help="core/act only; defaults: core=1s, act=200ms")
    parser.add_argument("--output", type=Path, default=ROOT / "measures.jsonl",
                        help="JSONL path; another name also isolates raw logs and expected results")
    args = parser.parse_args()
    if args.pairs < 1:
        parser.error("--pairs must be positive")
    if args.benchtime and args.phase not in ["core", "act"]:
        parser.error("--benchtime applies only to core/act")
    measurements = Measurements(args.output)
    if args.phase == "check":
        check_results(measurements)
    elif args.phase in ["core", "act"]:
        benchmarks(measurements, args.phase, args.pairs,
                   args.benchtime or ("1s" if args.phase == "core" else "200ms"))
    elif args.phase == "app":
        apps(measurements, args.pairs, json.loads(measurements.reference.read_text()))
    else:
        rss(measurements, args.pairs)


if __name__ == "__main__":
    main()

"""Run prebuilt variants sequentially, alternating AB/BA; retain raw logs."""
import argparse
import datetime
import json
import os
from pathlib import Path
import re
import subprocess
import time

ROOT = Path(__file__).resolve().parent
ENV = dict(os.environ, GOGC="800", GOMAXPROCS="14", GOMEMLIMIT="off", PPROF="0")
NAMES = ["AstTree", "Fib", "ListOps", "TCO", "Records", "Ackermann", "Church",
         "Primes", "RBTree", "Polymorphism", "StateMonad", "LazyEvaluation", "ArrayOps", "RowToList"]
BENCH = re.compile(r"^Benchmark(?:SumEvens|Act)/([^\s]+)-14\s+(\d+)\s+([\d.]+)\s+ns/op\s+([\d.]+)\s+B/op\s+([\d.]+)\s+allocs/op$", re.M)


def run(args, label, extra_env=None):
    started = time.monotonic()
    result = subprocess.run(args, cwd=ROOT, env=dict(ENV, **(extra_env or {})),
                            stdout=subprocess.PIPE, stderr=subprocess.STDOUT, text=True)
    (ROOT / "logs" / (label + ".log")).write_text(result.stdout)
    if result.returncode:
        raise RuntimeError(f"{label} failed ({result.returncode}); see raw log")
    print(f"{label}: {time.monotonic() - started:.1f}s", flush=True)
    return result.stdout


def emit(row):
    row["recorded_utc"] = datetime.datetime.now(datetime.timezone.utc).isoformat()
    with (ROOT / "measures.jsonl").open("a") as out:
        out.write(json.dumps(row) + "\n")


def order(pair):
    return ["before", "after"] if pair % 2 else ["after", "before"]


def check_results():
    reference = None
    for variant in ["before", "after"]:
        text = run([str(ROOT / "bin" / (variant + "-bench")), "-test.run=^TestResults$", "-test.v"], variant + "-results")
        lines = [line for line in text.splitlines() if line.startswith("RESULT ")]
        assert len(lines) == 16
        if reference is None:
            reference = lines
        else:
            assert lines == reference, "A/B outputs differ"
    expected = dict((m[0], json.loads(m[1])) for m in re.findall(r'^RESULT act (\w+) (".*")$', "\n".join(reference), re.M))
    assert list(expected) == NAMES
    (ROOT / "expected-results.json").write_text(json.dumps(expected, indent=2) + "\n")
    return expected


def benchmarks(kind, pairs, benchtime):
    name = "SumEvens" if kind == "core" else "Act"
    expected = ["n900", "n90000"] if kind == "core" else NAMES
    for pair in range(1, pairs + 1):
        for variant in order(pair):
            text = run([str(ROOT / "bin" / (variant + "-bench")), "-test.run=^$",
                        "-test.bench=^Benchmark" + name + "$", "-test.benchtime=" + benchtime,
                        "-test.count=1", "-test.cpu=14", "-test.benchmem"], f"{kind}-{pair:02}-{variant}")
            matches = BENCH.findall(text)
            assert [m[0] for m in matches] == expected, text
            for test, iterations, ns, bytes_, allocs in matches:
                emit(dict(kind=kind, variant=variant, pair=pair, test=test, iterations=int(iterations),
                          ns_per_op=float(ns), bytes_per_op=float(bytes_), allocs_per_op=float(allocs)))


def apps(pairs, expected):
    for pair in range(1, pairs + 1):
        for variant in order(pair):
            text = run([str(ROOT / "bin" / (variant + "-app"))], f"app-{pair:02}-{variant}")
            chunks = text.split("(Test)")[1:]
            assert len(chunks) == 14
            tests = []
            for name, chunk in zip(NAMES, chunks):
                match = re.search(r"\(Output & Warm-up\)\s*\n(.*?)\n\s*\(Execution time - best of 10\)\s*\n([\d.]+) μs", chunk, re.S)
                assert match, chunk
                output, us = match.groups()
                assert output.strip() == expected[name], (name, output, expected[name])
                tests.append(dict(test=name, result=output.strip(), time_us=float(us)))
            total = re.search(r"Total exec time: ([\d.]+) ms", text)
            assert total
            emit(dict(kind="app", variant=variant, pair=pair, tests=tests, total_ms=float(total[1])))


def rss(pairs):
    for pair in range(1, pairs + 1):
        for variant in order(pair):
            for size, iterations in [("n900", 10000), ("n90000", 100)]:
                text = run(["/usr/bin/time", "-l", str(ROOT / "bin" / (variant + "-bench")),
                            "-test.run=^TestFixedRSS$", "-test.v"], f"rss-{size}-{pair:02}-{variant}",
                           {"GOPURS_RSS_SIZE": size, "GOPURS_RSS_ITERATIONS": str(iterations)})
                match = re.search(r"(\d+)\s+maximum resident set size", text)
                assert match, text
                emit(dict(kind="rss", variant=variant, pair=pair, test=size, iterations=iterations,
                          rss_bytes=int(match[1])))


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("phase", choices=["check", "core", "act", "app", "rss"])
    parser.add_argument("--pairs", type=int, default=10)
    args = parser.parse_args()
    if args.phase == "check":
        check_results()
    elif args.phase in ["core", "act"]:
        benchmarks(args.phase, args.pairs, "1s" if args.phase == "core" else "200ms")
    elif args.phase == "app":
        apps(args.pairs, json.loads((ROOT / "expected-results.json").read_text()))
    else:
        rss(args.pairs)


if __name__ == "__main__":
    main()

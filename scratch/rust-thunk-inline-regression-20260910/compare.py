"""Compare the saved regression and corrected full Rust runners."""
from pathlib import Path
import hashlib
import json
import re
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
BUILD = HERE / "build"
EXPECTED = ["7", "55", "202950", "100000", "20000", "125", "100000", "21536",
            "22", "10000000", "1200", "1000000", "202950", "5"]
PATTERN = re.compile(r"\(Test\)\s+([^\n]+)\s+\(Output & Warm-up\)\s+(\S+)\s+"
                     r"\(Execution time - best of 10\)\s+([0-9.]+) μs")


def digest(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def measure():
    runs = {"before": [], "after": []}
    hashes = {name: digest(BUILD / f"runner-{name}") for name in runs}
    assert hashes["before"] != hashes["after"]
    for block in range(5):
        order = ["before", "after"] if block % 2 == 0 else ["after", "before"]
        for name in order:
            result = subprocess.run([str(BUILD / f"runner-{name}")], text=True,
                                    capture_output=True, check=True, timeout=60)
            assert not result.stderr, result.stderr
            rows = PATTERN.findall(result.stdout)
            assert [value for _, value, _ in rows] == EXPECTED, rows
            (BUILD / f"timing-{block + 1}-{name}.log").write_text(result.stdout)
            values = {title: float(us) for title, _, us in rows}
            runs[name].append(values)
            print(block + 1, name, round(sum(values.values()) / 1000, 3), "ms",
                  "Lazy:", values["Lazy Evaluation (1M Thunks Forced, 1k Depth):"], "us", flush=True)
    medians = {name: {bench: statistics.median(row[bench] for row in rows)
                      for bench in rows[0]} for name, rows in runs.items()}
    result = {
        "method": "Five alternating pairs, no concurrent builds/tests; unchanged warmup and best of ten; "
                  "all fourteen outputs checked in every process. Totals sum per-benchmark medians.",
        "binary_sha256": hashes,
        "runs_us": runs,
        "median_us": medians,
        "total_us": {name: sum(values.values()) for name, values in medians.items()},
    }
    (HERE / "runner-results.json").write_text(json.dumps(result, indent=2) + "\n")
    print(json.dumps({"total_us": result["total_us"], "median_us": medians}, indent=2))


if __name__ == "__main__":
    measure()

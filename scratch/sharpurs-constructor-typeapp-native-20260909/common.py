"""Validation shared by the isolated comparison experiment; no benchmarks run here."""
import hashlib
import math
import re
import statistics
from pathlib import Path

AUDIT = Path(__file__).resolve().parent
RBT = "Red-Black Tree (100k Worst-Case Insertions)"
LAZY = "Lazy Evaluation (1M Thunks Forced, 1k Depth)"
PRIMES = "Prime Sieve (sum primes up to 500)"
LIST = "List Processing (900 elements)"
ORDER = ["before", "after", "after", "before", "before", "after", "after", "before", "before", "after"]
BLOCK = re.compile(
    r"\(Test\)\s+(?P<name>[^\r\n]+):\s+"
    r"\(Output & Warm-up\)\s+(?P<output>.*?)\s+"
    r"\(Execution time - best of 10\)\s+(?P<time>\d+(?:\.\d+)?)\s+[μµ]s",
    re.DOTALL,
)


def require(condition, message):
    if not condition:
        raise ValueError(message)


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def read_run(path):
    content = path.read_text(encoding="utf-8")
    matches = list(BLOCK.finditer(content))
    totals = re.findall(r"Total exec time:\s*(\d+(?:\.\d+)?)\s+ms", content)
    require(len(matches) == content.count("(Test)") == 14, f"{path}: expected 14 timed tests")
    require(len(totals) == 1, f"{path}: expected exactly one total")
    tests = [{"name": item["name"].strip(), "output": item["output"].strip(),
              "best_of_10_us": float(item["time"])} for item in matches]
    require(len({test["name"] for test in tests}) == 14, f"{path}: duplicate test names")
    require(all(math.isfinite(test["best_of_10_us"]) and test["best_of_10_us"] >= 0
                for test in tests), f"{path}: invalid test time")
    total = float(totals[0])
    require(math.isfinite(total) and total > 0, f"{path}: invalid total")
    require(abs(total - sum(test["best_of_10_us"] for test in tests) / 1000) <= 0.02,
            f"{path}: total differs from the sum of the displayed tests")
    return {"log": path.name, "total_ms": total, "tests": tests}


def check_outputs(run, reference):
    signature = lambda record: [(test["name"], test["output"]) for test in record["tests"]]
    require(signature(run) == signature(reference),
            f"{run['log']}: outputs, names or order differ from the historical reference")


def baseline():
    text = (AUDIT / "README.baseline.md").read_text(encoding="utf-8")
    section = text.split("#### F#/C#\n", 1)[1].split("\n#### ", 1)[0]
    rows = {line.split("|", 1)[0].strip().strip("*"): line.split("|")[1:]
            for line in section.splitlines() if "|" in line}
    value = lambda label, column: float(re.fullmatch(
        r"\s*~\s*([\d.]+)\s*(?:[μµ]s|ms)\s*", rows[label][column])[1])
    actual = {"compiled_total_ms": value("Total Execution Time", 0),
              "compiled_primes_us": value("Prime Sieve", 0),
              "native_optimized_primes_us": value("Prime Sieve", 2),
              "compiled_list_us": value("List Processing", 0),
              "native_optimized_list_us": value("List Processing", 2),
              "compiled_lazy_us": value("Lazy Evaluation", 0),
              "native_optimized_lazy_us": value("Lazy Evaluation", 2),
              "compiled_rbtree_us": value("Red-Black Tree", 0)}
    require(actual == {"compiled_total_ms": 60.41,
                       "compiled_primes_us": 3861.79, "native_optimized_primes_us": 119,
                       "compiled_list_us": 747.63, "native_optimized_list_us": 77,
                       "compiled_lazy_us": 5805.67,
                       "native_optimized_lazy_us": 75, "compiled_rbtree_us": 38925.08},
            "Official README baseline changed; review the analysis assumptions")
    return {**actual, "source": "README.baseline.md", "sha256": sha(AUDIT / "README.baseline.md"),
            "note": "Historical rounded figures, not a simultaneous control. Primes and List Processing are the affected benchmark sites. Native optimized Primes uses a different algorithm, so its timing does not isolate constructor overhead. Lazy and RBTree are unchanged controls; native optimized Lazy performs 1,000 increments and returns 1,000, while compiled Lazy forces 1,000,000 thunks and returns 1,000,000."}


def stats(values):
    return {"n": len(values), "values": values, "median": statistics.median(values),
            "min": min(values), "max": max(values)}


def compare(before, after):
    require(before > 0 and after > 0, "Expected positive comparison times")
    return {"before": before, "after": after, "difference": after - before,
            "difference_percent": (after / before - 1) * 100, "speedup_factor": before / after}

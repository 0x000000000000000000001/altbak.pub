#!/usr/bin/env python3
"""Independent result oracles for the 28 JS, Chez, or Erlang foreign kernels.

No backend build, shared output, or timing is involved. Each native process uses
a temporary directory. Include each published input and small discriminating
inputs so that hardcoded answers and accidentally equal formulas fail.
"""
import argparse
import json
import math
from pathlib import Path
import re
import shutil
import subprocess
import tempfile

ROOT = Path(__file__).resolve().parents[3]


def ast(n):
    result = 1
    for depth in range(1, n + 1):
        result = (depth + 1) * result - 1
    return result


def fib(n):
    a, b = 0, 1
    for _ in range(n):
        a, b = b, a + b
    return a


def ack(m):
    # Closed forms for A(m, 4), independent of the recursive implementations.
    return [5, 6, 11, 125][m]


CASES = {
    "AstTree": [(n, ast(n)) for n in [0, 1, 2, 3, 5]],
    "Fib": [(n, fib(n)) for n in [0, 1, 2, 5, 10]],
    "ListOps": [(n, sum(range(2, n + 1, 2))) for n in [1, 2, 9, 900]],
    "TCO": [(n, sum(x % 3 for x in range(1, n + 1))) for n in [0, 1, 2, 5, 100000]],
    "Records": [(n, sum(x % 5 for x in range(1, n + 1))) for n in [0, 1, 2, 7, 10000]],
    "Ackermann": [(m, ack(m)) for m in range(4)],
    "Church": [(n, n ** 5) for n in [0, 1, 2, 3, 10]],
    "Primes": [(n, sum(x for x in range(2, n + 1)
                       if all(x % d for d in range(2, math.isqrt(x) + 1))))
               for n in [0, 1, 2, 10, 500]],
    # Depths of the specified Okasaki tree after descending insertions.
    "RBTree": [(0, 0), (1, 1), (2, 2), (3, 2), (7, 3), (100000, 22)],
    "Polymorphism": [(n, n) for n in [0, 1, 2, 17, 10000000]],
    # The FFI API takes chain depth; the PureScript act takes repetitions.
    "StateMonad": [(n, n * 20) for n in [0, 1, 2, 7, 60]],
    "LazyEvaluation": [(n, n * 1000) for n in [0, 1, 2, 7, 1000]],
    "ArrayOps": [(n, sum(range(2, n + 1, 2))) for n in [1, 2, 9, 900]],
    "RowToList": [(n, 5) for n in [0, 1, 9]],
}


def run(command, cwd):
    return subprocess.run(command, cwd=cwd, text=True, capture_output=True, timeout=90)


def check(language):
    failures = []
    count = 0
    with tempfile.TemporaryDirectory(prefix="altbak-native-") as directory:
        tmp = Path(directory)
        for name, cases in CASES.items():
            for suffix in ["FFI", "FFICheatcode"]:
                stem = name + suffix
                function = "run" + stem
                if language == "js":
                    shutil.copyfile(ROOT / "src/Test" / (stem + ".js"), tmp / "kernel.mjs")
                    (tmp / "check.mjs").write_text(
                        'import * as kernel from "./kernel.mjs";\n'
                        f'for (const [n, expected] of {json.dumps(cases)}) {{\n'
                        f' const actual = kernel.{function}(n);\n'
                        ' if (actual !== expected) throw new Error(`${n}: ${actual} != ${expected}`);\n'
                        '}\n')
                    result = run(["node", "check.mjs"], tmp)
                elif language == "scheme":
                    source = ROOT / "src/Test" / (stem + ".ss")
                    assertions = "\n".join(
                        f'(unless (= ({function} {n}) {expected}) '
                        f'(error "{stem}" "wrong result" {n} ({function} {n}) {expected}))'
                        for n, expected in cases)
                    (tmp / "check.ss").write_text(
                        f'(load {json.dumps(str(source))})\n'
                        f'(import (chezscheme) (Test.{stem} foreign))\n{assertions}\n')
                    result = run(["chez", "--optimize-level", "3", "--script", "check.ss"], tmp)
                else:
                    source = (ROOT / "src/Test" / (stem + ".erl")).read_text()
                    module = re.search(r"-module\(([^)]+)\)", source)[1]
                    (tmp / (module + ".erl")).write_text(source)
                    result = run(["erlc", "+warnings_as_errors", module + ".erl"], tmp)
                    if result.returncode == 0:
                        assertions = ",".join(
                            f"{expected} = {module}:{function}({n})" for n, expected in cases)
                        result = run(["erl", "+S", "1:1", "-noshell", "-pa", str(tmp),
                                      "-eval", assertions + ",halt()."], tmp)
                if result.returncode:
                    detail = (result.stderr or result.stdout).strip().splitlines()
                    failures.append(f"{stem}: " + " / ".join(detail[:4]))
                else:
                    count += len(cases)
        print(f"{language}: {count} assertions passed, {len(failures)} kernels failed", flush=True)
        for failure in failures:
            print(failure, flush=True)
    return not failures


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("language", choices=["js", "scheme", "erlang"])
    raise SystemExit(0 if check(parser.parse_args().language) else 1)

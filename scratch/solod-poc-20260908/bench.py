#!/usr/bin/env python3
"""Compare three extracted numeric kernels; this is not the full altbak suite.

Wall-clock batches include process startup and the driver, so they do not isolate
function-call overhead. There are no allocations in these kernels: results say
nothing about the Go GC or the compatibility/performance of the complete runtime.
"""

from __future__ import annotations

import argparse
import hashlib
import json
import os
from pathlib import Path
import platform
import random
import re
import statistics
import subprocess
import sys
import time
from datetime import datetime, timezone


ROOT = Path(__file__).resolve().parent
BINARY_NAMES = ("go-original", "go-adapted", "solod-o3")
WORKLOADS = (("fib", 10, 1000), ("ackermann", 4, 10), ("tco", 100000, 10))
TARGET_SECONDS = 0.15
REPETITIONS = 7
RANDOM_SEED = 20260908
INTEGER = re.compile(r"-?[0-9]+")


def oracle(kernel: str, n: int) -> int:
    if kernel == "fib":
        a, b = 0, 1
        for _ in range(n):
            a, b = b, a + b
        return a
    if kernel == "ackermann":
        return 2 ** (n + 3) - 3
    if kernel == "tco":
        groups, remainder = divmod(n, 3)
        return groups * 3 + remainder * (remainder + 1) // 2
    raise ValueError(f"Unknown kernel: {kernel}")


def expected_checksum(kernel: str, iterations: int, n: int) -> int:
    odd_calls = iterations // 2
    even_calls = iterations - odd_calls
    return even_calls * oracle(kernel, n) + odd_calls * oracle(kernel, n + 1)


def run(binary: Path, kernel: str, iterations: int, n: int) -> dict:
    command = [str(binary), kernel, str(iterations), str(n)]
    expected = expected_checksum(kernel, iterations, n)
    start = time.perf_counter_ns()
    result = subprocess.run(
        command, cwd=ROOT, capture_output=True, text=True, timeout=60, check=False
    )
    elapsed_ns = time.perf_counter_ns() - start
    if result.returncode:
        raise RuntimeError(
            f"Command failed ({result.returncode}): {' '.join(command)}\n"
            f"stdout: {result.stdout}\nstderr: {result.stderr}"
        )
    output = result.stdout.strip()
    if not INTEGER.fullmatch(output):
        raise RuntimeError(f"Expected a single integer from {command!r}, got {output!r}")
    checksum = int(output)
    if checksum != expected:
        raise RuntimeError(
            f"Checksum mismatch for {command!r}: expected {expected}, got {checksum}"
        )
    return {
        "binary": binary.name,
        "kernel": kernel,
        "iterations": iterations,
        "n": n,
        "elapsed_ns": elapsed_ns,
        "ns_per_call": elapsed_ns / iterations,
        "checksum": checksum,
        "stderr": result.stderr,
    }


def verify(binaries: list[Path]) -> dict:
    inputs = {
        "fib": list(range(26)),
        "ackermann": list(range(6)),
        "tco": [0, 1, 2, 3, 10, 100000],
    }
    checks = []
    for binary in binaries:
        count = 0
        for kernel, values in inputs.items():
            for n in values:
                # Odd batch size checks both alternating inputs and the sum.
                observation = run(binary, kernel, 3, n)
                checks.append(observation)
                count += 1
        print(f"Verified {binary.name}: {count} oracle checks", flush=True)
    return {"inputs": inputs, "iterations_per_check": 3, "checks": checks}


def calibrate(binary: Path, kernel: str, n: int, minimum: int) -> tuple[int, list[dict]]:
    iterations = minimum
    target_ns = int(TARGET_SECONDS * 1_000_000_000)
    observations = []
    for _ in range(12):
        observation = run(binary, kernel, iterations, n)
        observations.append(observation)
        elapsed = observation["elapsed_ns"]
        if elapsed >= target_ns * 0.85:
            return iterations, observations
        scale = min(20.0, max(1.2, target_ns / max(1, elapsed)))
        iterations = max(iterations + 1, int(iterations * scale))
    raise RuntimeError(f"Could not calibrate {kernel} to approximately {TARGET_SECONDS}s")


def metadata(binaries: list[Path]) -> dict:
    return {
        "utc": datetime.now(timezone.utc).isoformat(),
        "platform": platform.platform(),
        "machine": platform.machine(),
        "processor": platform.processor(),
        "cpu_count": os.cpu_count(),
        "python": sys.version,
        "toolchain": (ROOT / "versions.txt").read_text() if (ROOT / "versions.txt").exists() else None,
        "clock": "time.perf_counter_ns; external subprocess wall time",
        "environment": {
            name: os.environ.get(name)
            for name in ("GOGC", "GOMAXPROCS", "GODEBUG", "ASAN_OPTIONS", "UBSAN_OPTIONS")
        },
        "binaries": {
            binary.name: {
                "path": str(binary),
                "size_bytes": binary.stat().st_size,
                "sha256": hashlib.sha256(binary.read_bytes()).hexdigest(),
            }
            for binary in binaries
        },
    }


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--verify-only", action="store_true", help="check oracles without timing batches")
    args = parser.parse_args()
    binaries = [ROOT / "bin" / name for name in BINARY_NAMES]
    for binary in binaries:
        if not binary.is_file() or not os.access(binary, os.X_OK):
            parser.error(f"Missing executable: {binary}")
    verification_binaries = list(binaries)
    sanitizer = ROOT / "bin" / "solod-sanitize"
    if sanitizer.is_file():
        if not os.access(sanitizer, os.X_OK):
            parser.error(f"Not executable: {sanitizer}")
        verification_binaries.append(sanitizer)

    verification = verify(verification_binaries)
    if args.verify_only:
        (ROOT / "verification.json").write_text(json.dumps({
            "metadata": metadata(verification_binaries),
            "verification": verification,
        }, indent=2) + "\n", encoding="utf-8")
        print("All checksum checks passed.", flush=True)
        return 0

    results = {
        "metadata": metadata(verification_binaries),
        "methodology": {
            "target_batch_seconds": TARGET_SECONDS,
            "repetitions": REPETITIONS,
            "random_seed": RANDOM_SEED,
            "statistic": "median elapsed nanoseconds divided by calls per batch",
            "inputs": "n + (i & 1); Ackermann fixes its first argument to 3",
            "warmup": "one full batch for each binary before each workload",
            "limitations": [
                "Only three numeric kernels, not the complete altbak program or its runtime.",
                "Process startup and driver overhead are included in every sample.",
                "CPU scheduling, thermal state and concurrent work can affect measurements.",
                "No allocations in kernels: no comparison of GC or allocation-heavy workloads.",
            ],
        },
        "verification": verification,
        "workloads": {},
    }
    rng = random.Random(RANDOM_SEED)
    for kernel, n, minimum in WORKLOADS:
        iterations, calibration = calibrate(binaries[0], kernel, n, minimum)
        print(f"Benchmark {kernel} n={n}: {iterations} calls/batch", flush=True)
        warmups = [run(binary, kernel, iterations, n) for binary in binaries]
        samples = {binary.name: [] for binary in binaries}
        execution_order = []
        for repetition in range(REPETITIONS):
            order = list(binaries)
            rng.shuffle(order)
            execution_order.append([binary.name for binary in order])
            for binary in order:
                observation = run(binary, kernel, iterations, n)
                observation["repetition"] = repetition + 1
                samples[binary.name].append(observation)
        summary = {}
        for name, observations in samples.items():
            values = [observation["ns_per_call"] for observation in observations]
            summary[name] = {
                "median_ns_per_call": statistics.median(values),
                "min_ns_per_call": min(values),
                "max_ns_per_call": max(values),
            }
        baseline = summary["go-original"]["median_ns_per_call"]
        for name, stats in summary.items():
            stats["speedup_vs_go_original"] = baseline / stats["median_ns_per_call"]
            print(
                f"  {name:12s} {stats['median_ns_per_call']:12.2f} ns/call"
                f"  {stats['speedup_vs_go_original']:.3f}x vs Go original",
                flush=True,
            )
        results["workloads"][kernel] = {
            "n": n,
            "iterations": iterations,
            "calibration": calibration,
            "warmups": warmups,
            "execution_order": execution_order,
            "samples": samples,
            "summary": summary,
        }

    destination = ROOT / "results.json"
    destination.write_text(json.dumps(results, indent=2) + "\n", encoding="utf-8")
    print(f"Saved raw timings and metadata: {destination}", flush=True)
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except (RuntimeError, subprocess.TimeoutExpired) as error:
        print(f"ERROR: {error}", file=sys.stderr)
        raise SystemExit(1)

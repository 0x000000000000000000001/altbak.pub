#!/usr/bin/env python3
"""Verify and compare the extracted RBTree kernels, including allocation costs.

Each timed process builds and traverses several trees. Go's original heap/GC
strategy and the adapted arena strategy are reported separately. Pool creation,
process startup and the driver are included, amortized over each batch.
"""

from __future__ import annotations

import argparse
from datetime import datetime, timezone
import hashlib
import json
import math
import os
from pathlib import Path
import platform
import random
import re
import statistics
import subprocess
import sys
import time


ROOT = Path(__file__).resolve().parent
TARGET_SECONDS = 0.25
REPETITIONS = 7
RANDOM_SEED = 20260908
VERIFY_SIZES = (0, 1, 2, 3, 7, 10, 100, 1000, 10000, 100000, 100001)
SANITIZER_SIZES = (0, 10, 1000, 100000)
WORKLOADS = (10000, 100000)
INTEGER = re.compile(r"-?[0-9]+")
RSS = re.compile(r"^POC_RSS_BYTES: (\d+)\s*$", re.MULTILINE)
ENVIRONMENT_KEYS = (
    "GOGC", "GOMAXPROCS", "GOMEMLIMIT", "GODEBUG", "GOEXPERIMENT",
    "ASAN_OPTIONS", "UBSAN_OPTIONS", "MallocNanoZone",
)

# Names distinguish the two GC settings of the same original executable.
VARIANTS = (
    {"name": "go-original-gogc100", "binary": "go-original", "env": {"GOGC": "100"}, "arena": False},
    {"name": "go-original-gogc800", "binary": "go-original", "env": {"GOGC": "800"}, "arena": False},
    {"name": "go-syntax-gogc800", "binary": "go-syntax", "env": {"GOGC": "800"}, "arena": False},
    {"name": "go-arena-gogc800", "binary": "go-arena", "env": {"GOGC": "800"}, "arena": True},
    {"name": "solod-arena", "binary": "solod-arena", "env": {}, "arena": True},
)
CALIBRATION_VARIANT = VARIANTS[1]


def effective_environment(variant: dict) -> dict[str, str]:
    environment = os.environ.copy()
    environment.update(variant["env"])
    return environment


def command_for(variant: dict, mode: str, count: int, n: int) -> list[str]:
    return [str(ROOT / "bin" / variant["binary"]), mode, str(count), str(n)]


def invoke(command: list[str], variant: dict) -> dict:
    start = time.perf_counter_ns()
    process = subprocess.run(
        command, cwd=ROOT, env=effective_environment(variant),
        capture_output=True, text=True, timeout=60, check=False,
    )
    elapsed_ns = time.perf_counter_ns() - start
    if process.returncode:
        raise RuntimeError(
            f"Command failed ({process.returncode}): {command!r}\n"
            f"stdout: {process.stdout}\nstderr: {process.stderr}"
        )
    return {
        "variant": variant["name"],
        "command": command,
        "environment_overrides": dict(variant["env"]),
        "elapsed_ns": elapsed_ns,
        "stdout": process.stdout,
        "stderr": process.stderr,
    }


def integer_fields(output: str, count: int) -> list[int]:
    fields = output.split()
    if len(fields) != count or any(not INTEGER.fullmatch(field) for field in fields):
        raise RuntimeError(f"Expected {count} integer fields, got {output!r}")
    return [int(field) for field in fields]


def check_tree(variant: dict, n: int, reference: dict | None = None) -> dict:
    observation = invoke(command_for(variant, "verify", 1, n), variant)
    count, total, depth, black_height, allocations = integer_fields(observation["stdout"], 5)
    values = [count, total, depth, black_height]
    if count != n or total != n * (n + 1) // 2:
        raise RuntimeError(f"Wrong keys in {variant['name']} n={n}: {values}")
    if n == 100000 and values != [100000, 5000050000, 22, 17]:
        raise RuntimeError(f"Independent 100k oracle failed: {variant['name']}: {values}")
    if reference is not None and values != reference["validation"]:
        raise RuntimeError(
            f"Tree mismatch for {variant['name']} n={n}: "
            f"{values} != {reference['validation']}"
        )
    if variant["arena"]:
        if allocations < n:
            raise RuntimeError(f"Invalid arena allocation count: {variant['name']} n={n}: {allocations}")
        if n == 100000 and allocations != 2483948:
            raise RuntimeError(f"100k allocation oracle failed: {variant['name']}: {allocations}")
    elif allocations != -1:
        raise RuntimeError(f"Expected allocations=-1 for {variant['name']}, got {allocations}")
    observation.update({
        "n": n, "count": 1, "validation": values, "allocations": allocations,
        "timing_excluded": True,
    })
    return observation


def sanitizer_variant() -> dict:
    # Append explicit settings so they take precedence over inherited options.
    asan = os.environ.get("ASAN_OPTIONS", "")
    ubsan = os.environ.get("UBSAN_OPTIONS", "")
    return {
        "name": "solod-sanitize", "binary": "solod-sanitize", "arena": True,
        "env": {
            "ASAN_OPTIONS": ":".join(filter(None, (asan, "halt_on_error=1", "detect_stack_use_after_return=1"))),
            "UBSAN_OPTIONS": ":".join(filter(None, (ubsan, "halt_on_error=1"))),
        },
    }


def verify(sanitizer: dict) -> tuple[dict, dict[int, dict]]:
    checks = []
    references = {}
    arena_allocations = {}
    for variant in VARIANTS:
        for n in VERIFY_SIZES:
            observation = check_tree(variant, n, references.get(n))
            references.setdefault(n, observation)
            if variant["arena"]:
                previous = arena_allocations.setdefault(n, observation["allocations"])
                if observation["allocations"] != previous:
                    raise RuntimeError(f"Arena allocation mismatch for {variant['name']} n={n}")
            checks.append(observation)
        print(f"Verified {variant['name']}: {len(VERIFY_SIZES)} tree checks", flush=True)

    stack_detection_supported = True
    for n in SANITIZER_SIZES:
        observation = check_tree(sanitizer, n, references[n])
        # Some ASan runtimes warn about unrecognized options and continue.
        # Preserve that warning and explicitly report the unsupported setting.
        diagnostic = observation["stderr"].lower()
        if "unrecognized" in diagnostic and "detect_stack_use_after_return" in diagnostic:
            stack_detection_supported = False
        if observation["allocations"] != arena_allocations[n]:
            raise RuntimeError(f"Sanitizer allocation mismatch at n={n}")
        checks.append(observation)
    print(f"Verified solod-sanitize: {len(SANITIZER_SIZES)} tree checks", flush=True)
    return {
        "status": "passed", "checks_count": len(checks),
        "sizes": list(VERIFY_SIZES), "sanitizer_sizes": list(SANITIZER_SIZES),
        "independent_oracle_100000": {
            "count": 100000, "sum": 5000050000, "depth": 22,
            "black_height_including_nil": 17, "arena_allocations": 2483948,
        },
        "sanitizer_stack_use_after_return_requested": True,
        "sanitizer_stack_option_accepted_without_warning": stack_detection_supported,
        "checks": checks,
    }, references


def check_checksum(observation: dict, count: int, n: int, depth: int) -> None:
    checksum = integer_fields(observation["stdout"], 1)[0]
    expected = count * depth
    if checksum != expected:
        raise RuntimeError(
            f"Checksum mismatch for {observation['variant']} n={n} count={count}: "
            f"{checksum} != {expected}"
        )
    observation.update({"count": count, "n": n, "checksum": checksum})


def timed_batch(variant: dict, count: int, n: int, depth: int) -> dict:
    observation = invoke(command_for(variant, "bench", count, n), variant)
    check_checksum(observation, count, n, depth)
    observation["ns_per_tree"] = observation["elapsed_ns"] / count
    return observation


def calibrate(n: int, depth: int) -> tuple[int, list[dict]]:
    count = 3
    target_ns = int(TARGET_SECONDS * 1_000_000_000)
    observations = []
    for _ in range(12):
        observation = timed_batch(CALIBRATION_VARIANT, count, n, depth)
        observations.append(observation)
        if observation["elapsed_ns"] >= target_ns:
            return count, observations
        factor = min(8.0, max(1.2, 1.05 * target_ns / max(1, observation["elapsed_ns"])))
        count = max(count + 1, math.ceil(count * factor))
    raise RuntimeError(f"Failed to calibrate n={n} to {TARGET_SECONDS} seconds")


def binary_metadata(variants: tuple | list) -> dict:
    binaries = {}
    for variant in variants:
        name = variant["binary"]
        if name in binaries:
            continue
        path = ROOT / "bin" / name
        if not path.is_file() or not os.access(path, os.X_OK):
            raise RuntimeError(f"Missing executable: {path}")
        binaries[name] = {
            "path": str(path), "size_bytes": path.stat().st_size,
            "sha256": hashlib.sha256(path.read_bytes()).hexdigest(),
        }
    return binaries


def metadata(sanitizer: dict) -> dict:
    versions = ROOT / "versions.txt"
    variants = [*VARIANTS, sanitizer]
    return {
        "utc": datetime.now(timezone.utc).isoformat(),
        "platform": platform.platform(), "machine": platform.machine(),
        "processor": platform.processor(), "cpu_count": os.cpu_count(),
        "python": sys.version,
        "toolchain": versions.read_text(encoding="utf-8") if versions.is_file() else None,
        "clock": "time.perf_counter_ns; external subprocess wall time",
        "inherited_environment": {key: os.environ.get(key) for key in ENVIRONMENT_KEYS},
        "variants": {
            variant["name"]: {
                "binary": variant["binary"], "arena": variant["arena"],
                "effective_environment": {
                    key: effective_environment(variant).get(key) for key in ENVIRONMENT_KEYS
                },
            }
            for variant in variants
        },
        "binaries": binary_metadata(variants),
    }


def memory_runs(count: int, depth: int, enabled: bool) -> dict:
    if not enabled:
        return {"status": "skipped", "reason": "Disabled with --no-memory"}
    if sys.platform != "darwin":
        return {"status": "skipped", "reason": "RSS wrapper specifies macOS byte units only"}
    observations = []
    for variant in VARIANTS:
        command = [sys.executable, str(ROOT / "memory.py"), *command_for(variant, "bench", count, 100000)]
        observation = invoke(command, variant)
        check_checksum(observation, count, 100000, depth)
        match = RSS.search(observation["stderr"])
        if match is None:
            raise RuntimeError(f"Could not parse macOS maximum RSS: {observation['stderr']!r}")
        observation.pop("elapsed_ns")
        observation.update({
            "maximum_resident_set_size_bytes": int(match.group(1)),
            "timing_excluded": True,
        })
        observations.append(observation)
        print(f"RSS {variant['name']}: {int(match.group(1)) / 1048576:.2f} MiB", flush=True)
    return {
        "status": "measured", "units": "bytes, macOS resource.getrusage(RUSAGE_CHILDREN).ru_maxrss in a fresh parent per child",
        "note": "Separate processes after all timing runs; includes runtime and allocator memory.",
        "observations": observations,
    }


def write_json(path: Path, value: dict) -> None:
    path.write_text(json.dumps(value, indent=2) + "\n", encoding="utf-8")


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--verify-only", action="store_true", help="verify all executables without benchmark batches")
    parser.add_argument("--no-memory", action="store_true", help="skip separate macOS maximum-RSS measurements")
    args = parser.parse_args()
    sanitizer = sanitizer_variant()
    initial_metadata = metadata(sanitizer)
    verification, references = verify(sanitizer)
    if args.verify_only:
        destination = ROOT / "verification.json"
        write_json(destination, {"metadata": initial_metadata, "verification": verification})
        print(f"All {verification['checks_count']} tree checks passed; saved {destination}", flush=True)
        return 0

    results = {
        "metadata": initial_metadata,
        "methodology": {
            "target_batch_seconds": TARGET_SECONDS, "minimum_trees_per_batch": 3,
            "calibration_variant": CALIBRATION_VARIANT["name"],
            "repetitions": REPETITIONS, "random_seed": RANDOM_SEED,
            "statistic": "median elapsed nanoseconds divided by trees per batch",
            "inputs": "Each tree inserts n,n-1,...,1 and then traverses both branches to compute depth.",
            "warmup": "One calibrated full batch per variant before each workload.",
            "memory_strategy": {
                "go-original": "Generated persistent nodes with ordinary Go allocation and GC.",
                "go-syntax": "Syntax-only Solod adaptations, with ordinary Go allocation and GC.",
                "go-arena": "Explicit node pool allocated once per process, reset between completed trees.",
                "solod-arena": "Same explicit node pool strategy, transpiled with Solod and compiled with Clang.",
            },
            "limitations": [
                "Extraction of RBTree only, not the complete generated program or runtime.",
                "C versus original Go includes changing the memory strategy; C versus Go arena helps separate it.",
                "All wall times include process startup, initial pool allocation and the benchmark driver.",
                "The arena retains temporary nodes until the completed tree has been traversed.",
                "The fixed input is externally supplied; C driver and kernel are separate units without LTO.",
                "Scheduling, concurrent processes, CPU frequency and thermal state can affect this single session.",
            ],
        },
        "verification": verification,
        "workloads": {},
    }
    rng = random.Random(RANDOM_SEED)
    for n in WORKLOADS:
        depth = references[n]["validation"][2]
        count, calibration = calibrate(n, depth)
        print(f"Benchmark n={n}: {count} trees per batch", flush=True)
        warmups = [timed_batch(variant, count, n, depth) for variant in VARIANTS]
        samples = {variant["name"]: [] for variant in VARIANTS}
        execution_order = []
        for repetition in range(REPETITIONS):
            order = list(VARIANTS)
            rng.shuffle(order)
            execution_order.append([variant["name"] for variant in order])
            for variant in order:
                observation = timed_batch(variant, count, n, depth)
                observation["repetition"] = repetition + 1
                samples[variant["name"]].append(observation)
        summary = {}
        for name, observations in samples.items():
            values = [observation["ns_per_tree"] for observation in observations]
            summary[name] = {
                "median_ns_per_tree": statistics.median(values),
                "min_ns_per_tree": min(values), "max_ns_per_tree": max(values),
                "mean_ns_per_tree": statistics.mean(values),
                "stdev_ns_per_tree": statistics.stdev(values),
            }
        for name, statistics_for_variant in summary.items():
            median = statistics_for_variant["median_ns_per_tree"]
            for baseline in ("go-original-gogc100", "go-original-gogc800", "go-arena-gogc800"):
                statistics_for_variant[f"speedup_vs_{baseline}"] = summary[baseline]["median_ns_per_tree"] / median
            print(
                f"  {name:24s} {median / 1000000:9.3f} ms/tree  "
                f"{statistics_for_variant['speedup_vs_go-original-gogc800']:.3f}x vs Go original GOGC800",
                flush=True,
            )
        results["workloads"][str(n)] = {
            "n": n, "count": count, "depth": depth,
            "calibration": calibration, "warmups": warmups,
            "execution_order": execution_order, "samples": samples, "summary": summary,
        }

    big_workload = results["workloads"]["100000"]
    # Preserve completed timings even if an optional OS memory metric fails.
    write_json(ROOT / "results.json", results)
    results["memory"] = memory_runs(big_workload["count"], big_workload["depth"], not args.no_memory)
    final_binaries = binary_metadata([*VARIANTS, sanitizer])
    if final_binaries != initial_metadata["binaries"]:
        raise RuntimeError("An executable changed during verification or benchmarking")
    results["metadata"]["binary_hashes_verified_after_run"] = True
    results["metadata"]["completed_utc"] = datetime.now(timezone.utc).isoformat()
    destination = ROOT / "results.json"
    write_json(destination, results)
    print(f"Saved raw measurements, validation and metadata: {destination}", flush=True)
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except (RuntimeError, subprocess.TimeoutExpired, OSError) as error:
        print(f"ERROR: {error}", file=sys.stderr)
        raise SystemExit(1)

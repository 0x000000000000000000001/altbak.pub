#!/usr/bin/env python3
"""Screen progressive and capped Go node pools against the generated Go heap.

All timings include process startup, pool growth, full tree construction and
depth traversal, and any reset performed by the runner. RSS is collected in
separate fresh processes. Diagnostic runtime statistics never enter timings.
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
DESTINATION = ROOT / "results-screen.json"
MIB = 1024 * 1024
N = 100000
TARGET_SECONDS = 0.2
REPETITIONS = 5
MEMORY_REPETITIONS = 3
SEED = 20260908
VERIFY_SIZES = (0, 1, 1000, N)
RSS = re.compile(r"^POC_RSS_BYTES: (\d+)\s*$", re.MULTILINE)
INTEGER = re.compile(r"-?[0-9]+")
ENVIRONMENT_KEYS = (
    "GOGC", "GOMEMLIMIT", "GOMAXPROCS", "GODEBUG", "GOEXPERIMENT",
    "GOTOOLCHAIN", "GOFLAGS", "MallocNanoZone",
)


def variant(name: str, cap_mib: int, gogc: int, limit: str = "off", *, pool: bool = True) -> dict:
    return {
        "name": name,
        "binary": "go-pool" if pool else "go-original",
        "pool": pool,
        "cap_mib": cap_mib,
        "env": {"GOGC": str(gogc), "GOMEMLIMIT": limit},
    }


VARIANTS = (
    variant("go-original-gogc100", 0, 100, pool=False),
    variant("go-original-gogc800", 0, 800, pool=False),
    variant("go-original-gogc800-limit64", 0, 800, "64MiB", pool=False),
    variant("go-pool8-gogc800", 8, 800),
    variant("go-pool16-gogc800", 16, 800),
    variant("go-pool32-gogc800", 32, 800),
    variant("go-pool96-gogc800", 96, 800),
    variant("go-pool8-gogc100", 8, 100),
    variant("go-pool16-gogc100", 16, 100),
    variant("go-pool8-gogc800-limit64", 8, 800, "64MiB"),
    variant("go-pool16-gogc800-limit64", 16, 800, "64MiB"),
)
BASELINE = VARIANTS[1]


def environment(item: dict) -> dict[str, str]:
    result = os.environ.copy()
    result.update(item["env"])
    return result


def command(item: dict, mode: str, count: int, n: int) -> list[str]:
    return [str(ROOT / "bin" / item["binary"]), mode, str(count), str(n), str(item["cap_mib"])]


def invoke(argv: list[str], item: dict, *, timed: bool = False) -> dict:
    started = time.perf_counter_ns()
    process = subprocess.run(
        argv, cwd=ROOT, env=environment(item), text=True,
        capture_output=True, check=False, timeout=120,
    )
    elapsed = time.perf_counter_ns() - started
    if process.returncode:
        raise RuntimeError(
            f"Command failed ({process.returncode}): {argv!r}\n"
            f"stdout: {process.stdout}\nstderr: {process.stderr}"
        )
    observation = {
        "variant": item["name"], "command": argv,
        "environment_overrides": item["env"],
        "stdout": process.stdout, "stderr": process.stderr,
    }
    if timed:
        observation["elapsed_ns"] = elapsed
    else:
        observation["timing_excluded"] = True
    return observation


def json_output(observation: dict) -> dict:
    try:
        result = json.loads(observation["stdout"])
    except json.JSONDecodeError as error:
        raise RuntimeError(f"Invalid JSON for {observation['variant']}: {observation['stdout']!r}") from error
    if not isinstance(result, dict):
        raise RuntimeError(f"Expected diagnostic JSON object for {observation['variant']}, got {result!r}")
    return result


def require_integer(value: dict, key: str) -> int:
    result = value.get(key)
    if type(result) is not int:
        raise RuntimeError(f"Expected integer field {key!r}, got {value!r}")
    return result


def verify() -> tuple[dict, dict[int, list[int]]]:
    observations = []
    references = {}
    for item in VARIANTS:
        for n in VERIFY_SIZES:
            observation = invoke(command(item, "verify", 1, n), item)
            result = json_output(observation)
            values = [require_integer(result, key) for key in ("count", "sum", "depth", "black_height")]
            if values[0] != n or values[1] != n * (n + 1) // 2:
                raise RuntimeError(f"Invalid tree keys for {item['name']} n={n}: {result}")
            if n == 0 and values != [0, 0, 0, 1]:
                raise RuntimeError(f"Empty tree oracle failed for {item['name']}: {result}")
            if n == N and values != [100000, 5000050000, 22, 17]:
                raise RuntimeError(f"Independent 100k tree oracle failed for {item['name']}: {result}")
            previous = references.setdefault(n, values)
            if values != previous:
                raise RuntimeError(f"Tree mismatch for {item['name']} n={n}: {values} != {previous}")
            if item["pool"]:
                allocations = require_integer(result, "allocations")
                pool_bytes = require_integer(result, "pool_bytes")
                fallback = require_integer(result, "fallback_allocations")
                if allocations < n or not 0 <= fallback <= allocations:
                    raise RuntimeError(f"Invalid allocation counters for {item['name']} n={n}: {result}")
                if not 0 <= pool_bytes <= item["cap_mib"] * MIB:
                    raise RuntimeError(f"Pool cap violated for {item['name']} n={n}: {result}")
                if n == N and allocations != 2483948:
                    raise RuntimeError(f"Independent 100k allocation oracle failed for {item['name']}: {result}")
                if item["cap_mib"] == 96 and fallback != 0:
                    raise RuntimeError(f"Unexpected fallback within full pool for {item['name']} n={n}: {result}")
            observation.update({"n": n, "count": 1, "validation": values, "result": result})
            observations.append(observation)
        print(f"Verified {item['name']}: {len(VERIFY_SIZES)} trees", flush=True)
    return {
        "status": "passed", "checks_count": len(observations),
        "sizes": list(VERIFY_SIZES),
        "oracle_100000": {
            "count": 100000, "sum": 5000050000, "depth": 22,
            "black_height": 17, "pool_total_allocations": 2483948,
        },
        "checks": observations,
    }, references


def checksum(observation: dict, count: int, n: int, depth: int) -> None:
    output = observation["stdout"].strip()
    if not INTEGER.fullmatch(output):
        raise RuntimeError(f"Invalid benchmark checksum for {observation['variant']}: {output!r}")
    value = int(output)
    if value != count * depth:
        raise RuntimeError(f"Checksum mismatch for {observation['variant']}: {value} != {count * depth}")
    observation.update({"count": count, "n": n, "checksum": value})


def timed_batch(item: dict, count: int, depth: int) -> dict:
    observation = invoke(command(item, "bench", count, N), item, timed=True)
    checksum(observation, count, N, depth)
    observation["ns_per_tree"] = observation["elapsed_ns"] / count
    return observation


def calibrate(depth: int) -> tuple[int, list[dict]]:
    count = 3
    observations = []
    target_ns = int(TARGET_SECONDS * 1_000_000_000)
    for _ in range(12):
        observation = timed_batch(BASELINE, count, depth)
        observations.append(observation)
        if observation["elapsed_ns"] >= target_ns:
            return count, observations
        factor = min(8.0, max(1.2, 1.05 * target_ns / max(1, observation["elapsed_ns"])))
        count = max(count + 1, math.ceil(count * factor))
    raise RuntimeError("Could not calibrate a 0.2 second baseline batch")


def binaries() -> dict:
    result = {}
    for name in sorted({item["binary"] for item in VARIANTS}):
        path = ROOT / "bin" / name
        if not path.is_file() or not os.access(path, os.X_OK):
            raise RuntimeError(f"Missing executable: {path}")
        result[name] = {
            "path": str(path), "size_bytes": path.stat().st_size,
            "sha256": hashlib.sha256(path.read_bytes()).hexdigest(),
        }
    return result


def metadata() -> dict:
    version_files = {}
    for path in (ROOT / "versions.txt", ROOT.parent / "versions.txt"):
        if path.is_file():
            version_files[str(path)] = path.read_text()
    try:
        go = subprocess.run(["go", "version"], capture_output=True, text=True, timeout=15, check=False)
        go_version = {"returncode": go.returncode, "stdout": go.stdout, "stderr": go.stderr}
    except (OSError, subprocess.TimeoutExpired) as error:
        go_version = {"unavailable": str(error)}
    return {
        "started_utc": datetime.now(timezone.utc).isoformat(),
        "platform": platform.platform(), "machine": platform.machine(),
        "cpu_count": os.cpu_count(), "python": sys.version,
        "toolchain_files": version_files, "go_version": go_version,
        "inherited_environment": {key: os.environ.get(key) for key in ENVIRONMENT_KEYS},
        "variants": {
            item["name"]: {
                **item,
                "effective_environment": {key: environment(item).get(key) for key in ENVIRONMENT_KEYS},
            }
            for item in VARIANTS
        },
        "binaries": binaries(),
        "benchmark_script_sha256": hashlib.sha256(Path(__file__).read_bytes()).hexdigest(),
    }


def save(results: dict) -> None:
    results["updated_utc"] = datetime.now(timezone.utc).isoformat()
    temporary = DESTINATION.with_suffix(".json.tmp")
    temporary.write_text(json.dumps(results, indent=2) + "\n")
    temporary.replace(DESTINATION)


def summary(samples: dict[str, list[dict]]) -> dict:
    result = {}
    for name, observations in samples.items():
        values = [item["ns_per_tree"] for item in observations]
        result[name] = {
            "median_ns_per_tree": statistics.median(values),
            "min_ns_per_tree": min(values), "max_ns_per_tree": max(values),
            "mean_ns_per_tree": statistics.mean(values),
            "stdev_ns_per_tree": statistics.stdev(values),
        }
    baseline_median = result[BASELINE["name"]]["median_ns_per_tree"]
    for value in result.values():
        value["speedup_vs_go_original_gogc800"] = baseline_median / value["median_ns_per_tree"]
    return result


def memory_runs(results: dict, count: int, depth: int) -> None:
    if sys.platform != "darwin":
        results["memory"] = {"status": "skipped", "reason": "RSS helper uses macOS byte units"}
        save(results)
        return
    memory = {
        "status": "running", "repetitions": MEMORY_REPETITIONS,
        "units": "bytes; macOS ru_maxrss in a fresh Python parent per child",
        "note": "Separate from timings; peak includes Go runtime, pool and fallback heap.",
        "execution_order": [], "observations": [], "summary": {},
    }
    results["memory"] = memory
    rng = random.Random(SEED + 1)
    for repetition in range(MEMORY_REPETITIONS):
        order = list(VARIANTS)
        rng.shuffle(order)
        memory["execution_order"].append([item["name"] for item in order])
        for item in order:
            argv = [sys.executable, str(ROOT.parent / "memory.py"), *command(item, "bench", count, N)]
            observation = invoke(argv, item)
            checksum(observation, count, N, depth)
            matches = RSS.findall(observation["stderr"])
            if len(matches) != 1:
                raise RuntimeError(f"Expected one RSS reading for {item['name']}: {observation['stderr']!r}")
            observation.update({
                "repetition": repetition + 1,
                "maximum_resident_set_size_bytes": int(matches[0]),
            })
            memory["observations"].append(observation)
        save(results)
        print(f"RSS repetition {repetition + 1}/{MEMORY_REPETITIONS} complete", flush=True)
    for item in VARIANTS:
        values = [
            observation["maximum_resident_set_size_bytes"]
            for observation in memory["observations"] if observation["variant"] == item["name"]
        ]
        memory["summary"][item["name"]] = {
            "median_peak_rss_bytes": statistics.median(values),
            "min_peak_rss_bytes": min(values), "max_peak_rss_bytes": max(values),
        }
        print(f"RSS {item['name']}: median {statistics.median(values) / MIB:.2f} MiB", flush=True)
    memory["status"] = "measured"
    save(results)


def run(args: argparse.Namespace, results: dict) -> None:
    results.update({
        "status": "running", "metadata": metadata(),
        "methodology": {
            "n": N, "target_batch_seconds": TARGET_SECONDS,
            "minimum_trees_per_batch": 3, "repetitions": REPETITIONS,
            "memory_repetitions": MEMORY_REPETITIONS, "random_seed": SEED,
            "calibration_variant": BASELINE["name"],
            "clock": "external time.perf_counter_ns around subprocess.run",
            "statistic": "median external elapsed nanoseconds divided by equal trees per batch",
            "warmup": "one full calibrated batch per variant",
            "memory_limit": "GOMEMLIMIT is explicit off or 64MiB; it is a Go soft memory limit, not an RSS cap",
            "pool_cap": "upper bound for progressively allocated typed blocks, not process memory",
            "workload": "build persistent RBTree by inserting n,n-1,...,1; consume full depth; reset between trees",
            "limitations": [
                "Screening of one extracted generated Go kernel at 100000 insertions, on this machine only.",
                "Arena lifetime must enclose every use of all tree versions; capped pools still use Go GC for overflow.",
                "A retained pointer-containing pool affects Go live heap accounting and GC pacing.",
                "Process startup, initial allocation, resets and output are included in external batch timings.",
                "Diagnostic GC statistics are collected separately and do not describe the timed processes exactly.",
            ],
        },
    })
    save(results)
    verification, references = verify()
    results["verification"] = verification
    results["phase"] = "verified"
    save(results)
    if args.verify_only:
        results["status"] = "verified-only"
    else:
        depth = references[N][2]
        count, calibration = calibrate(depth)
        workload = {
            "n": N, "count": count, "depth": depth, "calibration": calibration,
            "warmups": [], "execution_order": [],
            "samples": {item["name"]: [] for item in VARIANTS},
        }
        results["workload"] = workload
        results["phase"] = "calibrated"
        save(results)
        print(f"Screen n={N}: {count} trees per batch", flush=True)
        workload["warmups"] = [timed_batch(item, count, depth) for item in VARIANTS]
        results["phase"] = "warmed-up"
        save(results)
        rng = random.Random(SEED)
        for repetition in range(REPETITIONS):
            order = list(VARIANTS)
            rng.shuffle(order)
            workload["execution_order"].append([item["name"] for item in order])
            for item in order:
                observation = timed_batch(item, count, depth)
                observation["repetition"] = repetition + 1
                workload["samples"][item["name"]].append(observation)
            save(results)
            print(f"Timing repetition {repetition + 1}/{REPETITIONS} complete", flush=True)
        workload["summary"] = summary(workload["samples"])
        results["phase"] = "timed"
        save(results)
        for name, value in workload["summary"].items():
            print(
                f"{name:36s} {value['median_ns_per_tree'] / 1e6:9.3f} ms/tree "
                f"{value['speedup_vs_go_original_gogc800']:.3f}x vs original GOGC800",
                flush=True,
            )
        memory_runs(results, count, depth)
        results["phase"] = "memory-measured"
        save(results)
        diagnostics = []
        results["diagnostics"] = {
            "status": "running", "mode": "stats", "count": 3, "n": N,
            "timing_excluded": True, "observations": diagnostics,
        }
        for item in VARIANTS:
            observation = invoke(command(item, "stats", 3, N), item)
            observation.update({"count": 3, "n": N, "result": json_output(observation)})
            diagnostics.append(observation)
            save(results)
        results["diagnostics"]["status"] = "collected"
        results["phase"] = "diagnostics-collected"
        results["status"] = "complete"
    if binaries() != results["metadata"]["binaries"]:
        raise RuntimeError("An executable changed during verification or measurements")
    results["metadata"]["binary_hashes_verified_after_run"] = True
    results["metadata"]["completed_utc"] = datetime.now(timezone.utc).isoformat()
    save(results)
    print(f"Saved {DESTINATION}", flush=True)


def main() -> int:
    global DESTINATION
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--verify-only", action="store_true", help="verify every variant without timings, RSS or stats")
    args = parser.parse_args()
    if args.verify_only:
        DESTINATION = ROOT / "verification.json"
    results = {}
    try:
        run(args, results)
    except (RuntimeError, OSError, subprocess.TimeoutExpired) as error:
        results.update({"status": "failed", "error": str(error)})
        save(results)
        print(f"ERROR: {error}", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

#!/usr/bin/env python3
"""Rebuild Purust's numeric kernels and compare them with C using timed batches."""

import argparse
from datetime import datetime, timezone
import importlib.util
import json
import os
from pathlib import Path
import re
import shutil
import statistics
import subprocess
import sys
import tempfile
from types import SimpleNamespace

sys.dont_write_bytecode = True
import run_benchmarks as references

ROOT = references.ROOT
spec = importlib.util.spec_from_file_location("native_driver", ROOT / "bin/native/driver.py")
native = importlib.util.module_from_spec(spec)
spec.loader.exec_module(native)


def source_fingerprints(env):
    result = native.inputs("rust")
    for filename in ["BenchPurust.purs", "bench_purust.rs", "bench_native.c",
                     "run_purust_benchmark.py", "run_benchmarks.py"]:
        path = ROOT / "tmp" / filename
        result[str(path)] = references.sha256(path)
    for tool in ["purs", "spago"]:
        path = Path(shutil.which(tool, path=env["PATH"])).resolve()
        result[str(path)] = references.sha256(path)
    config = (ROOT / "run/bak/rust/spago.rust.yaml").read_text()
    for relative in re.findall(r'(?m)^\s*path:\s*"([^"\n]+)"', config):
        package = (ROOT / relative).resolve()
        for path in sorted(package.rglob("*")):
            if path.is_file() and path.suffix in {".purs", ".rs"}:
                result[str(path)] = references.sha256(path)
    return result


def checked_result(output):
    result = native.validate_output(output)
    counts = [int(n) for n in re.findall(r"^Batch iterations: (\d+)$", output, re.M)]
    if len(counts) != 14 or any(n < 1 or n > 16777216 or n & (n - 1) for n in counts):
        raise ValueError("Invalid batch sizes")
    if any(t <= 0 for t in result["times_us"]):
        raise ValueError("Unresolved zero duration")
    result["batch_iterations"] = counts
    return result


def render_readme(original, purust, c):
    updated = references.render_readme(original, {"C (reference)": c})
    match = re.search(r"(#### Rust\n)(.*?)(?=\n#### )", updated, re.S)
    section = match[2]
    pattern = r"^([^|\n]+)(\|\s*~\s*)([0-9.]+)( μs)"
    labels = re.findall(pattern, section, re.M)
    if [row[0].strip() for row in labels] != references.TABLE_ROWS:
        raise ValueError("Unexpected Rust README rows")
    times = iter(purust["times_us"])
    section = re.sub(pattern, lambda row: row[1] + row[2] + f"{next(times):.3f}" + row[4],
                     section, flags=re.M)
    section, count = re.subn(r"(\*\*Total Execution Time\*\*\s*\|\s*~\s*)[0-9.]+( ms)",
                             lambda row: row[1] + f"{purust['sum_displayed_lines_ms']:.2f}" + row[2], section)
    if count != 1:
        raise ValueError("Expected one Rust total")
    return updated[:match.start(2)] + section + updated[match.end(2):]


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--update-readme", action="store_true")
    parser.add_argument("--build-only", action="store_true")
    args = parser.parse_args()
    if args.build_only and args.update_readme:
        parser.error('--build-only cannot update measurements')
    original = (ROOT / "README.md").read_text()
    parent = ROOT / "var/benchmark/purust-reference"
    parent.mkdir(parents=True, exist_ok=True)
    stamp = datetime.now(timezone.utc).strftime("%Y%m%dT%H%M%SZ-")
    directory = Path(tempfile.mkdtemp(prefix=stamp, dir=parent))
    print(f"Builds and results: {directory}", flush=True)
    build = directory / "purust"
    native.prepare(build, "rust", SimpleNamespace(clean=False, mode="pure"))
    shutil.copy2(ROOT / "tmp/BenchPurust.purs", build / "src/BenchPurust.purs")
    # Reuse downloaded dependencies, never compiled PureScript or Rust output.
    cached = ROOT / "run/bak/rust/modes/pure"
    if (cached / ".spago/p").is_dir():
        shutil.copytree(cached / ".spago/p", build / ".spago/p")
    if (cached / "spago.lock").is_file():
        shutil.copy2(cached / "spago.lock", build / "spago.lock")
    env = os.environ.copy()
    env["PATH"] = str(ROOT / "run/bak/js/node_modules/.bin") + os.pathsep + env["PATH"]
    env["CARGO_PROFILE_RELEASE_OPT_LEVEL"] = "3"
    env["CARGO_PROFILE_RELEASE_DEBUG"] = "false"
    profile = native.profile("rust", args, env)
    fingerprints = source_fingerprints(env)
    commands = []
    def logged(command, cwd, filename):
        commands.append(native.logged(command, cwd, directory / filename, env))
    logged(["spago", "build", "--offline"], build, "spago.log")
    backend = ROOT.parent / "purust/purust/bin/purust"
    logged([backend, "--main", "BenchPurust", "--source", "output", "--out", "output/purust_output"],
           build, "purust.log")
    project = build / "output/purust_output"
    shutil.copy2(project / "src/main.rs", directory / "generated-main.rs")
    shutil.copy2(ROOT / "tmp/bench_purust.rs", project / "src/main.rs")
    logged(["cargo", "build", "--offline", "--release", "--target-dir", build / "target",
            "--bin", "purust_output"], project, "cargo.log")
    binary = build / "target/release/purust_output"
    c_binary = directory / "benchmark_c"
    shutil.copy2(ROOT / "tmp/bench_native.c", directory / "bench_native.c")
    logged(["clang", "-O3", "-o", c_binary, directory / "bench_native.c"], directory, "clang.log")
    if fingerprints != source_fingerprints(env):
        raise RuntimeError("Inputs changed during compilation; rebuild before measuring")
    manifest = {
        "created_at_utc": datetime.now(timezone.utc).isoformat(), "profile": profile,
        "sources": fingerprints, "commands": commands,
        "rustc": subprocess.check_output(["rustc", "--version"], text=True, env=env).strip(),
        "purs": subprocess.check_output(["purs", "--version"], text=True, env=env).strip(),
        "clang": subprocess.check_output(["clang", "--version"], text=True).strip(),
        "binary_sha256": references.sha256(binary), "c_binary_sha256": references.sha256(c_binary),
        "spago_lock_sha256": references.sha256(build / "spago.lock"),
        "cargo_lock_sha256": references.sha256(project / "Cargo.lock"),
        "protocol": {"samples": 10, "global_warmups": 3, "local_warmups": 3,
                     "calibration_target_ms": 10, "maximum_batch_iterations": 16777216,
                     "process_runs_per_language": 3,
                     "statistic": "median of three per-process minimum batch durations divided by invocation count",
                     "numeric_kernels": True},
    }
    native.write_json(directory / "manifest.json", manifest)
    if args.build_only:
        print('Build manifest: ' + str(directory / 'manifest.json'), flush=True)
        return
    runs = {"Purust": [], "C (reference)": []}
    pair = [("Purust", binary), ("C (reference)", c_binary)]
    for repetition in range(3):
        for name, executable in (pair if repetition % 2 == 0 else list(reversed(pair))):
            print(f"Measuring {name}, process {repetition + 1}/3...", flush=True)
            prefix = "purust" if name == "Purust" else "c"
            output = references.run_logged([str(executable)], directory, f"{prefix}-results-{repetition + 1}.log")
            result = checked_result(output)
            runs[name].append(result)
            print(f"{name}: 14/14 results verified; total {result['total_ms']:.6f} ms", flush=True)
    results = {}
    for name, samples in runs.items():
        result = dict(samples[0])
        result.pop("batch_iterations")
        result["times_us"] = [statistics.median(sample["times_us"][i] for sample in samples) for i in range(14)]
        result["total_ms"] = result["sum_displayed_lines_ms"] = sum(result["times_us"]) / 1000.0
        result["process_runs"] = samples
        result["aggregation"] = "median per benchmark across three processes"
        results[name] = result
        print(f"{name}: sum of median rows {result['total_ms']:.6f} ms", flush=True)
    native.write_json(directory / "results.json", results)
    if args.update_readme:
        if (ROOT / "README.md").read_text() != original:
            raise RuntimeError("README changed during measurement; results retained without overwriting it")
        (ROOT / "README.md").write_text(render_readme(original, results["Purust"], results["C (reference)"]))
        print("README: compiled Purust column and C reference updated.", flush=True)


if __name__ == "__main__":
    main()

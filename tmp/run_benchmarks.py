#!/usr/bin/env python3
"""Rebuild native references and measure three validated processes per language."""

import argparse
from datetime import datetime, timezone
import hashlib
import json
import math
from pathlib import Path
import platform
import re
import shutil
import stat
import statistics
import subprocess
import sys
import tempfile

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT / "bin/benchmark"))
from validate import CASES, validate_output

CONFIG = {
    "Koka": ("koka", ["bench_koka.kk"],
             ["koka", "-O3", "-j1", "--builddir=build", "-o", "benchmark", "bench_koka.kk"]),
    "Haskell": ("ghc", ["bench_haskell.hs"],
                ["ghc", "-O2", "-outputdir", ".", "-o", "benchmark", "bench_haskell.hs"]),
    "OCaml": ("ocamlopt", ["bench_ocaml.ml", "bench_clock.c"],
              ["ocamlopt", "-O3", "-o", "benchmark", "bench_clock.c", "bench_ocaml.ml"]),
    "C (reference)": ("clang", ["bench_native.c"],
                      ["clang", "-O3", "-o", "benchmark", "bench_native.c"]),
}
TABLE_ROWS = ["AST Evaluation", "Fibonacci", "List Processing", "Tail Call Optimization",
              "Deep Record Updates", "Ackermann", "Church Numerals", "Prime Sieve",
              "Red-Black Tree", "Polymorphism", "State Monad", "Lazy Evaluation",
              "Array Processing", "RowToList"]
PROCESS_COUNT = 3
FP_LANGUAGES = {"Haskell", "Koka", "OCaml"}


def sha256(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def run_logged(command, directory, name, timeout=180):
    log = directory / name
    with log.open("w") as output:
        result = subprocess.run(command, cwd=directory, stdout=output,
                                stderr=subprocess.STDOUT, timeout=timeout)
    if result.returncode:
        raise RuntimeError(f"Command failed: {command!r}; see {log}")
    return log.read_text()


def checked_result(output):
    result = validate_output(output)
    counts = [int(n) for n in re.findall(r"^Batch iterations: (\d+)$", output, re.M)]
    if len(counts) != 14 or any(n < 1 or n > 16777216 or n & (n - 1) for n in counts):
        raise ValueError(f"Invalid batch sizes: {counts}")
    if any(not math.isfinite(t) or t <= 0 for t in result["times_us"]):
        raise ValueError("Nonpositive or nonfinite duration")
    result["batch_iterations"] = counts
    return result


def aggregate(samples):
    if len(samples) != PROCESS_COUNT:
        raise ValueError("Three validated processes are required")
    result = dict(samples[0])
    result.pop("batch_iterations")
    result["times_us"] = [statistics.median(s["times_us"][i] for s in samples) for i in range(14)]
    result["total_ms"] = result["sum_displayed_lines_ms"] = sum(result["times_us"]) / 1000.0
    result["process_runs"] = samples
    result["aggregation"] = "median per benchmark across three processes; total sums median cells"
    return result


def render_readme(original, results):
    """Replace only the numeric cells of the four validated reference tables."""
    updated = original
    for name, result in results.items():
        pattern = r"(#### " + re.escape(name) + r"\n)(.*?)(?=\n#### |\n> |\Z)"
        matches = list(re.finditer(pattern, updated, re.S))
        if len(matches) != 1:
            raise ValueError(f"Expected one README section for {name}")
        match = matches[0]
        section = match[2]
        labels = re.findall(r"^([^|\n]+)\|\s*~\s*[0-9.]+ μs", section, re.M)
        if [label.strip() for label in labels] != TABLE_ROWS:
            raise ValueError(f"Unexpected timing rows or order in README section {name}")
        times = iter(result["times_us"])
        section = re.sub(r"\|\s*~\s*[0-9.]+ μs", lambda _: f"| ~ {next(times):.3f} μs", section)
        section, count = re.subn(r"\|\s*~\s*[0-9.]+ ms",
                                 f"| ~ {result['sum_displayed_lines_ms']:.2f} ms", section)
        if count != 1:
            raise ValueError(f"Expected one total in README section {name}")
        updated = updated[:match.start(2)] + section + updated[match.end(2):]
    return updated


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--update-readme", action="store_true",
                        help="replace only the selected reference tables after all results validate")
    parser.add_argument("--build-only", action="store_true", help="build and check kernels without timing them")
    parser.add_argument("--languages", nargs="+", choices=list(CONFIG), default=list(CONFIG),
                        help="references to rebuild and measure (default: all four)")
    parser.add_argument("--timeout", type=float, default=180,
                        help="seconds allowed for each build or benchmark process")
    args = parser.parse_args()
    if args.build_only and args.update_readme:
        parser.error('--build-only cannot update measurements')
    if len(set(args.languages)) != len(args.languages):
        parser.error("--languages must not contain duplicates")
    if not math.isfinite(args.timeout) or args.timeout <= 0:
        parser.error("--timeout must be a positive finite number")
    selected = {name: CONFIG[name] for name in args.languages}
    original_readme = (ROOT / "README.md").read_text()
    parent = ROOT / "var/benchmark/native-references"
    parent.mkdir(parents=True, exist_ok=True)
    stamp = datetime.now(timezone.utc).strftime("%Y%m%dT%H%M%SZ-")
    directory = Path(tempfile.mkdtemp(prefix=stamp, dir=parent))
    (directory / "README.before.md").write_text(original_readme)
    print(f"Logs and build manifest: {directory}", flush=True)
    manifest = {
        "started_at_utc": datetime.now(timezone.utc).isoformat(), "status": "building",
        "languages": args.languages,
        "system": platform.platform(), "machine": platform.machine(),
        "protocol": {"global_warmups": 3, "local_warmups": 3, "samples": 10,
                     "calibration_target_ms": 10, "maximum_batch_iterations": 16777216,
                     "process_runs_per_language": PROCESS_COUNT,
                     "statistic": "median of three per-process minimum per-call batch durations",
                     "process_order": "selected order, reverse order, selected order",
                     "clock": "monotonic", "unit": "microseconds"},
        "sources": {}, "builds": {}, "processes": [],
    }
    def save_manifest():
        (directory / "manifest.json").write_text(json.dumps(manifest, indent=2) + "\n")
    if platform.system() == "Darwin":
        cpu = subprocess.run(["sysctl", "-n", "machdep.cpu.brand_string"],
                             capture_output=True, text=True)
        manifest["cpu"] = cpu.stdout.strip() if cpu.returncode == 0 else "unavailable"
    paths = [Path(__file__), ROOT / "bin/benchmark/validate.py"]
    paths += [ROOT / "src/Test" / (name + ".purs") for name in CASES]
    paths += [ROOT / "tmp" / filename for _, files, _ in selected.values() for filename in files]
    for path in paths:
        manifest["sources"][str(path.relative_to(ROOT))] = sha256(path)
    def verify_sources():
        for name, expected in manifest["sources"].items():
            if sha256(ROOT / name) != expected:
                raise RuntimeError(f"Input changed during run: {name}; results retained, no README update")
    pure_snapshot = directory / "purescript"
    pure_snapshot.mkdir()
    for name in CASES:
        shutil.copy2(ROOT / "src/Test" / (name + ".purs"), pure_snapshot)
    save_manifest()
    # Compile every executable before any measurements, using fresh directories.
    for name, (compiler, files, command) in selected.items():
        verify_sources()
        print(f"Building {name}...", flush=True)
        build = directory / compiler
        build.mkdir()
        for filename in files:
            source = ROOT / "tmp" / filename
            shutil.copy2(source, build / filename)
            if sha256(build / filename) != manifest["sources"][str(source.relative_to(ROOT))]:
                raise RuntimeError(f"Source changed while copying: {filename}")
        executable = shutil.which(compiler)
        if executable is None:
            raise RuntimeError(f"Compiler unavailable: {compiler}")
        command = [executable] + command[1:]
        version_flag = "-version" if compiler == "ocamlopt" else "--version"
        version = subprocess.check_output([executable, version_flag], text=True).strip()
        run_logged(command, build, "build.log", args.timeout)
        binary = build / "benchmark"
        binary.chmod(binary.stat().st_mode | stat.S_IXUSR)
        manifest["builds"][name] = {"command": command, "version": version,
                                    "directory": str(build), "binary_sha256": sha256(binary)}
        if name in FP_LANGUAGES:
            checked = run_logged([str(binary), "--check-only"], build, "check.log", args.timeout)
            expected = [f"{key}={value}" for key, value in CASES.items()]
            if checked.splitlines() != expected:
                raise ValueError(f"Incorrect --check-only outputs for {name}; see {build / 'check.log'}")
            manifest["builds"][name]["correctness_validated_cases"] = 14
        save_manifest()
    verify_sources()
    if args.build_only:
        manifest['status'] = 'built'
        save_manifest()
        print('Build manifest: ' + str(directory / 'manifest.json'), flush=True)
        return
    manifest["status"] = "measuring"
    save_manifest()
    samples = {name: [] for name in selected}
    # One process at a time; a failed check prevents every README update.
    for repetition in range(PROCESS_COUNT):
        order = list(selected) if repetition % 2 == 0 else list(reversed(selected))
        for name in order:
            verify_sources()
            compiler = selected[name][0]
            build = directory / compiler
            binary = build / "benchmark"
            if sha256(binary) != manifest["builds"][name]["binary_sha256"]:
                raise RuntimeError(f"Binary changed before measurement: {binary}")
            print(f"Measuring {name}, process {repetition + 1}/{PROCESS_COUNT}...", flush=True)
            log_name = f"results-{repetition + 1}.log"
            output = run_logged([str(binary)], build, log_name, args.timeout)
            result = checked_result(output)
            samples[name].append(result)
            manifest["processes"].append({"language": name, "repetition": repetition + 1,
                                           "log": str(build / log_name),
                                           "sha256": sha256(build / log_name), "validated_cases": 14})
            (build / "validated-processes.json").write_text(json.dumps(samples[name], indent=2) + "\n")
            save_manifest()
            print(f"{name}: 14/14 results verified, total {result['total_ms']:.6f} ms", flush=True)
    results = {name: aggregate(runs) for name, runs in samples.items()}
    for name, result in results.items():
        build = directory / selected[name][0]
        (build / "results.json").write_text(json.dumps(result, indent=2) + "\n")
        print(f"{name}: sum of median rows {result['total_ms']:.6f} ms", flush=True)
    (directory / "results.json").write_text(json.dumps(results, indent=2) + "\n")
    if args.update_readme:
        verify_sources()
        if (ROOT / "README.md").read_text() != original_readme:
            raise RuntimeError("README changed during measurement; results retained without overwriting it")
        (ROOT / "README.md").write_text(render_readme(original_readme, results))
        print("README: only selected reference tables updated.", flush=True)
    manifest.update({"status": "complete", "completed_at_utc": datetime.now(timezone.utc).isoformat(),
                     "readme_updated": args.update_readme})
    save_manifest()


if __name__ == "__main__":
    main()

#!/usr/bin/env python3
"""Rebuild, measure and validate the four handwritten native references."""

import argparse
from datetime import datetime, timezone
import hashlib
import json
from pathlib import Path
import platform
import re
import shutil
import stat
import subprocess
import sys
import tempfile

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT / "bin/benchmark"))
from validate import validate_output

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


def sha256(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def run_logged(command, directory, name):
    log = directory / name
    with log.open("w") as output:
        result = subprocess.run(command, cwd=directory, stdout=output,
                                stderr=subprocess.STDOUT, timeout=180)
    if result.returncode:
        raise RuntimeError(f"Command failed: {command!r}; see {log}")
    return log.read_text()


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
                        help="replace the four reference tables after all results validate")
    args = parser.parse_args()
    original_readme = (ROOT / "README.md").read_text()
    parent = ROOT / "var/benchmark/native-references"
    parent.mkdir(parents=True, exist_ok=True)
    stamp = datetime.now(timezone.utc).strftime("%Y%m%dT%H%M%SZ-")
    directory = Path(tempfile.mkdtemp(prefix=stamp, dir=parent))
    print(f"Logs and build manifest: {directory}", flush=True)
    manifest = {
        "started_at_utc": datetime.now(timezone.utc).isoformat(),
        "system": platform.platform(), "machine": platform.machine(),
        "protocol": {"global_warmups": 3, "local_warmups": 3, "samples": 10,
                     "calibration_target_ms": 10, "maximum_batch_iterations": 16777216,
                     "statistic": "minimum batch duration divided by batch iterations",
                     "clock": "monotonic", "unit": "microseconds"},
        "sources": {}, "builds": {},
    }
    if platform.system() == "Darwin":
        cpu = subprocess.run(["sysctl", "-n", "machdep.cpu.brand_string"],
                             capture_output=True, text=True)
        manifest["cpu"] = cpu.stdout.strip() if cpu.returncode == 0 else "unavailable"
    for path in [Path(__file__), ROOT / "bin/benchmark/validate.py", ROOT / "README.md"]:
        manifest["sources"][str(path.relative_to(ROOT))] = sha256(path)
    # Compile every executable before any measurements, using fresh directories.
    for name, (compiler, files, command) in CONFIG.items():
        print(f"Building {name}...", flush=True)
        build = directory / compiler
        build.mkdir()
        for filename in files:
            source = ROOT / "tmp" / filename
            shutil.copy2(source, build / filename)
            manifest["sources"][str(source.relative_to(ROOT))] = sha256(source)
        executable = shutil.which(compiler)
        if executable is None:
            raise RuntimeError(f"Compiler unavailable: {compiler}")
        command = [executable] + command[1:]
        version_flag = "-version" if compiler == "ocamlopt" else "--version"
        version = subprocess.check_output([executable, version_flag], text=True).strip()
        run_logged(command, build, "build.log")
        binary = build / "benchmark"
        binary.chmod(binary.stat().st_mode | stat.S_IXUSR)
        manifest["builds"][name] = {"command": command, "version": version,
                                    "directory": str(build), "binary_sha256": sha256(binary)}
    (directory / "manifest.json").write_text(json.dumps(manifest, indent=2) + "\n")
    results = {}
    # One process at a time; a failed check prevents every README update.
    for name, (compiler, _, _) in CONFIG.items():
        print(f"Measuring {name}...", flush=True)
        build = directory / compiler
        output = run_logged([str(build / "benchmark")], build, "results.log")
        result = validate_output(output)
        counts = [int(n) for n in re.findall(r"^Batch iterations: (\d+)$", output, re.M)]
        if len(counts) != 14 or any(n < 1 or n > 16777216 or n & (n - 1) for n in counts):
            raise ValueError(f"Invalid batch sizes for {name}: {counts}")
        if any(t <= 0 for t in result["times_us"]):
            raise ValueError(f"Unresolved zero duration for {name}")
        result["batch_iterations"] = counts
        results[name] = result
        (build / "results.json").write_text(json.dumps(result, indent=2) + "\n")
        print(f"{name}: 14/14 results verified, total {result['total_ms']:.6f} ms", flush=True)
    (directory / "results.json").write_text(json.dumps(results, indent=2) + "\n")
    if args.update_readme:
        if (ROOT / "README.md").read_text() != original_readme:
            raise RuntimeError("README changed during measurement; results retained without overwriting it")
        (ROOT / "README.md").write_text(render_readme(original_readme, results))
        print("README: four reference tables updated.", flush=True)


if __name__ == "__main__":
    main()

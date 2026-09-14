#!/usr/bin/env python3
"""Measure the three prebuilt C++ modes sequentially, in three balanced orders."""
import hashlib
import json
from pathlib import Path
import shutil
import statistics
import subprocess
import time

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
ORDERS = [("pure", "ffi", "fficc"), ("ffi", "fficc", "pure"), ("fficc", "pure", "ffi")]
result = {"orders": ORDERS, "aggregation": "median of three best-of-ten times per row; total is sum of row medians",
          "runs": [], "columns": {}}
if (HERE / "results.json").exists():
    raise SystemExit("Results already exist; preserve this dated measurement before rerunning.")
for mode in ORDERS[0]:
    source = ROOT / "run/bak/cpp/modes" / mode / "build.json"
    metadata = json.loads(source.read_text())
    if metadata["mode"] != mode or metadata["profile"] != ["-std=c++11", "-O3", "-DNDEBUG"]:
        raise SystemExit("Unexpected C++ build profile")
    if hashlib.sha256(Path(metadata["binary"]).read_bytes()).hexdigest() != metadata["binary_sha256"]:
        raise SystemExit("Binary differs from manifest")
    shutil.copy2(source, HERE / (mode + "-build.json"))
for round_index, order in enumerate(ORDERS, 1):
    for mode in order:
        directory = ROOT / "run/bak/cpp/modes" / mode
        before = set(directory.glob("run-*.json"))
        command = [str(ROOT / "bin/cpp/run"), "--run-only"] + ([] if mode == "pure" else ["--" + mode])
        log = HERE / (str(round_index) + "-" + mode + ".log")
        print(f"Round {round_index}/3: {mode}", flush=True)
        start = time.monotonic()
        with log.open("w") as stream:
            subprocess.run(command, cwd=ROOT, stdout=stream, stderr=subprocess.STDOUT, check=True)
        outputs = set(directory.glob("run-*.json")) - before
        if len(outputs) != 1:
            raise SystemExit("Expected one validated result file")
        run = json.loads(outputs.pop().read_text())
        run.update({"round": round_index, "wall_seconds": time.monotonic() - start, "archived_log": log.name})
        result["runs"].append(run)
        (HERE / "in-progress.json").write_text(json.dumps(result, indent=2) + "\n")
        print(f"  validated: {run['sum_displayed_lines_ms']:.5f} ms; wall {run['wall_seconds']:.2f}s", flush=True)
for mode in ORDERS[0]:
    runs = [run for run in result["runs"] if run["mode"] == mode]
    rows = [statistics.median(values) for values in zip(*(run["times_us"] for run in runs))]
    result["columns"][mode] = {"times_us": rows, "total_ms": sum(rows) / 1000,
                               "run_totals_ms": [run["sum_displayed_lines_ms"] for run in runs]}
(HERE / "results.json").write_text(json.dumps(result, indent=2) + "\n")
(HERE / "in-progress.json").unlink()
print(json.dumps(result["columns"], indent=2), flush=True)

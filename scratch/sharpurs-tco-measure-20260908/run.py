"""Run already-built variants sequentially, each observation in a fresh process."""
from pathlib import Path
import hashlib
import json
import os
import subprocess
import time

audit = Path(__file__).resolve().parent
dotnet = os.environ.get("DOTNET", "/Users/0x1/.dotnet/dotnet")
order = ["before", "after", "after", "before", "before", "after"]
counts = {"before": 0, "after": 0}
executions = []
for variant in order:
    counts[variant] += 1
    label = f"{variant}-{counts[variant]}"
    assembly = audit / variant / "bin/Release/net8.0/Program.dll"
    start = time.perf_counter()
    result = subprocess.run([dotnet, str(assembly)], capture_output=True, text=True, check=True, timeout=60)
    elapsed = time.perf_counter() - start
    (audit / f"{label}.jsonl").write_text(result.stdout)
    (audit / f"{label}.stderr").write_text(result.stderr)
    rows = [json.loads(line) for line in result.stdout.splitlines()]
    metadata = next(row for row in rows if row["kind"] == "metadata")
    measurements = [row for row in rows if row["kind"] == "measurement"]
    warmups = [row for row in rows if row["kind"] == "warmup"]
    assert len(measurements) == 10 and len(warmups) == 3
    assert metadata["globalWarmups"] == 3
    assert all(row["output"] == "100000" for row in measurements + warmups)
    assert metadata["serverGc"] is False, "GC mode differs from the original program"
    summary = next(row for row in rows if row["kind"] == "summary")
    executions.append({"label": label, "variant": variant, "wallSeconds": elapsed,
                       "assemblySha256": hashlib.sha256(assembly.read_bytes()).hexdigest()})
    print(json.dumps({"label": label, **summary}), flush=True)

# One full original App run per variant confirms the official benchmark context.
# It retains all global warmups, local warmups, and the best-of-ten harness.
for variant in ["before", "after"]:
    assembly = audit / variant / "bin/Release/net8.0/Program.dll"
    start = time.perf_counter()
    result = subprocess.run([dotnet, str(assembly), "--official"], capture_output=True, text=True, check=True, timeout=90)
    elapsed = time.perf_counter() - start
    (audit / f"official-{variant}.log").write_text(result.stdout)
    (audit / f"official-{variant}.stderr").write_text(result.stderr)
    executions.append({"label": f"official-{variant}", "variant": variant, "wallSeconds": elapsed})
    print(f"official-{variant}: completed in {elapsed:.2f} s", flush=True)

(audit / "executions.json").write_text(json.dumps(executions, indent=2) + "\n")

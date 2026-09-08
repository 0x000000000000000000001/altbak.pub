"""Measure already-built programs sequentially, with fresh processes per run."""
from datetime import datetime, timezone
from pathlib import Path
import hashlib
import json
import os
import re
import subprocess
import time

audit = Path(__file__).resolve().parent
root = audit.parent.parent
dotnet = os.environ.get("DOTNET", "/Users/0x1/.dotnet/dotnet")
sha = lambda path: hashlib.sha256(path.read_bytes()).hexdigest()
inputs = json.loads((audit / "inputs.json").read_text())
configurations = {}
for variant in ["before", "after", "diagnostic-before", "diagnostic-after"]:
    configurations[variant] = json.loads((audit / variant / "bin/Release/net8.0/Program.runtimeconfig.json").read_text())
    expected = inputs[("before" if variant.endswith("before") else "after") + "Sha256"]
    for name, value in expected.items():
        if variant.startswith("diagnostic-") and name == "EntryPoint.fs":
            value = inputs["diagnosticEntryPointSha256"]
        assert sha(audit / variant / name) == value, (variant, name)
assert all(config == configurations["before"] for config in configurations.values())

metadata = {
    "startedUtc": datetime.now(timezone.utc).isoformat(),
    "sdk": subprocess.check_output([dotnet, "--version"], text=True).strip(),
    "runtimes": subprocess.check_output([dotnet, "--list-runtimes"], text=True).splitlines(),
    "runtimeConfigs": configurations,
    "environmentOverrides": {key: os.environ.get(key) for key in [
        "DOTNET_TieredCompilation", "DOTNET_TieredPGO", "DOTNET_gcServer",
        "COMPlus_TieredCompilation", "COMPlus_TieredPGO", "COMPlus_gcServer"]},
}
executions = []
(audit / "metadata.json").write_text(json.dumps(metadata, indent=2) + "\n")

def execute(kind, variant, number):
    label = f"{kind}-{variant}-{number}"
    directory = variant if kind == "official" else f"diagnostic-{variant}"
    assembly = audit / directory / "bin/Release/net8.0/Program.dll"
    started = datetime.now(timezone.utc).isoformat()
    load_before = os.getloadavg()
    start = time.perf_counter()
    result = subprocess.run([dotnet, str(assembly)], capture_output=True, text=True, timeout=90)
    elapsed = time.perf_counter() - start
    suffix = ".log" if kind == "official" else ".jsonl"
    (audit / (label + suffix)).write_text(result.stdout)
    (audit / (label + ".stderr")).write_text(result.stderr)
    result.check_returncode()
    if kind == "official":
        total = re.findall(r"Total exec time:\s*([\d.]+) ms", result.stdout)
        assert len(total) == 1
        summary = {"totalMs": float(total[0])}
    else:
        rows = [json.loads(line) for line in result.stdout.splitlines()]
        measured = [row for row in rows if row["kind"] == "measurement"]
        warmups = [row for row in rows if row["kind"] == "warmup"]
        meta = next(row for row in rows if row["kind"] == "metadata")
        assert len(measured) == 10 and len(warmups) == 3
        assert all(row["output"] == "10000000" for row in measured + warmups)
        assert meta["globalWarmups"] == 3 and meta["serverGc"] is False
        summary = next(row for row in rows if row["kind"] == "summary")
    executions.append({"label": label, "kind": kind, "variant": variant,
        "startedUtc": started, "wallSeconds": elapsed, "loadBefore": load_before,
        "loadAfter": os.getloadavg(), "assemblySha256": sha(assembly), "summary": summary})
    (audit / "executions.json").write_text(json.dumps(executions, indent=2) + "\n")
    print(json.dumps({"label": label, "wallSeconds": elapsed, **summary}), flush=True)

for kind, order in [
    ("official", ["before", "after", "after", "before", "before", "after", "after", "before", "before", "after"]),
    ("diagnostic", ["before", "after", "after", "before", "before", "after"]),
]:
    counts = {"before": 0, "after": 0}
    for variant in order:
        counts[variant] += 1
        execute(kind, variant, counts[variant])

assert all(sha(root / "output/Main" / name) == value for name, value in inputs["afterSha256"].items())
metadata["generatedSourcesUnchangedDuringMeasurements"] = True
metadata["finishedUtc"] = datetime.now(timezone.utc).isoformat()
(audit / "metadata.json").write_text(json.dumps(metadata, indent=2) + "\n")
print("Complete: ten official runs, six diagnostics, generated sources unchanged during measurements.", flush=True)

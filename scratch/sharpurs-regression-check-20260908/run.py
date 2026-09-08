"""Run original App.main in ten fresh processes; builds must be finished first."""
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
order = ["reference", "current", "current", "reference", "reference",
         "current", "current", "reference", "reference", "current"]
counts = {variant: 0 for variant in set(order)}
sha = lambda path: hashlib.sha256(path.read_bytes()).hexdigest()
inputs = json.loads((audit / "inputs.json").read_text())
executions = []
runtime_configs = {variant: json.loads((audit / variant / "bin/Release/net8.0/Program.runtimeconfig.json").read_text())
                   for variant in counts}
assert runtime_configs["reference"] == runtime_configs["current"]
for variant in counts:
    expected = inputs[variant + "Sha256"]
    assert all(sha(audit / variant / name) == value for name, value in expected.items())

metadata = {
    "startedUtc": datetime.now(timezone.utc).isoformat(),
    "sdk": subprocess.check_output([dotnet, "--version"], text=True).strip(),
    "runtimes": subprocess.check_output([dotnet, "--list-runtimes"], text=True).splitlines(),
    "runtimeConfigs": runtime_configs, "order": order,
    "environmentOverrides": {name: os.environ.get(name) for name in [
        "DOTNET_TieredCompilation", "DOTNET_TieredPGO", "DOTNET_gcServer",
        "COMPlus_TieredCompilation", "COMPlus_TieredPGO", "COMPlus_gcServer"]},
    "outputValidation": "analyze.py checks the 14 displayed outputs in every process; the original harness does not print each timed result",
}
(audit / "metadata.json").write_text(json.dumps(metadata, indent=2) + "\n")
for variant in order:
    counts[variant] += 1
    label = f"{variant}-{counts[variant]}"
    assembly = audit / variant / "bin/Release/net8.0/Program.dll"
    started = datetime.now(timezone.utc).isoformat()
    load_before = os.getloadavg()
    start = time.perf_counter()
    result = subprocess.run([dotnet, str(assembly)], capture_output=True, text=True, timeout=90)
    elapsed = time.perf_counter() - start
    (audit / f"{label}.log").write_text(result.stdout)
    (audit / f"{label}.stderr").write_text(result.stderr)
    result.check_returncode()
    total = re.findall(r"Total exec time:\s*([\d,.]+) ms", result.stdout)
    assert len(total) == 1, f"Missing completion for {label}"
    row = {"label": label, "variant": variant, "startedUtc": started,
           "wallSeconds": elapsed, "assemblySha256": sha(assembly),
           "loadBefore": load_before, "loadAfter": os.getloadavg(),
           "totalMs": float(total[0].replace(",", ""))}
    executions.append(row)
    (audit / "executions.json").write_text(json.dumps(executions, indent=2) + "\n")
    print(f"{label}: {row['totalMs']:.2f} ms (wall {elapsed:.2f}s)", flush=True)

assert all(sha(root / "output/Main" / name) == value for name, value in inputs["currentSha256"].items())
metadata["generatedSourcesUnchanged"] = True
metadata["finishedUtc"] = datetime.now(timezone.utc).isoformat()
(audit / "metadata.json").write_text(json.dumps(metadata, indent=2) + "\n")
print("Complete: ten original App.main runs; generated working sources unchanged.", flush=True)

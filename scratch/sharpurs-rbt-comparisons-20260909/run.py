"""Run already-built App.main copies, sequentially in fresh processes; never builds."""
import argparse
from datetime import datetime, timezone
import json
import os
from pathlib import Path
import subprocess
import time

from common import AUDIT, ORDER, RBT, baseline, check_outputs, read_run, require, sha

parser = argparse.ArgumentParser(description=__doc__)
parser.add_argument("--probe", action="store_true", help="One before/after pair instead of five each")
args = parser.parse_args()
prefix = "probe-" if args.probe else ""
order = ["before", "after"] if args.probe else ORDER
dotnet = os.environ.get("DOTNET", "/Users/0x1/.dotnet/dotnet")
inputs_path = AUDIT / "inputs.json"
inputs = json.loads(inputs_path.read_text())
source = Path(inputs.get("sourceGeneratedDirectory", inputs.get(
    "sourceGeneratedPath", str(AUDIT.parents[1] / "run/bak/sharp/output/Main"))))
reference = read_run(AUDIT / "expected-output.log")
official = baseline()
configs, assemblies = {}, {}


def verify_sources():
    for directory, expected in [(AUDIT / "before", inputs["beforeSha256"]),
                                (AUDIT / "after", inputs["afterSha256"]),
                                (source, inputs["sourceGeneratedSha256"])]:
        require(expected, f"Empty source manifest for {directory}")
        for name, expected_hash in expected.items():
            require(sha(directory / name) == expected_hash, f"Source changed: {directory / name}")
    require(sha(inputs_path) == manifest_sha, "inputs.json changed during measurements")


manifest_sha = sha(inputs_path)
verify_sources()
for variant in ("before", "after"):
    directory = AUDIT / variant / "bin/Release/net8.0"
    configs[variant] = json.loads((directory / "Program.runtimeconfig.json").read_text())
    assemblies[variant] = {"path": str(directory / "Program.dll"), "sha256": sha(directory / "Program.dll")}
require(configs["before"] == configs["after"], "Runtime configurations differ")
require(not (AUDIT / f"{prefix}executions.json").exists(), "Existing measurement series: preserve or move it before running again")
utc = lambda: datetime.now(timezone.utc).isoformat()
metadata = {"startedUtc": utc(), "probe": args.probe, "order": order,
            "sdk": subprocess.check_output([dotnet, "--version"], text=True).strip(),
            "runtimes": subprocess.check_output([dotnet, "--list-runtimes"], text=True).splitlines(),
            "runtimeConfigs": configs, "assemblies": assemblies, "inputsSha256": manifest_sha,
            "historicalOutputSha256": sha(AUDIT / "expected-output.log"), "officialReadme": official,
            "sourceGeneratedPath": str(source),
            "environmentOverrides": {name: os.environ.get(name) for name in (
                "DOTNET_TieredCompilation", "DOTNET_TieredPGO", "DOTNET_gcServer",
                "COMPlus_TieredCompilation", "COMPlus_TieredPGO", "COMPlus_gcServer")}}


def save(name, value):
    (AUDIT / f"{prefix}{name}.json").write_text(json.dumps(value, indent=2, ensure_ascii=False) + "\n")


save("metadata", metadata)
executions, counts = [], {"before": 0, "after": 0}
for variant in order:
    verify_sources()
    assembly = Path(assemblies[variant]["path"])
    require(sha(assembly) == assemblies[variant]["sha256"], f"Assembly changed: {assembly}")
    counts[variant] += 1
    label = f"{prefix}official-{variant}-{counts[variant]}"
    started, load, clock = utc(), os.getloadavg(), time.perf_counter()
    result = subprocess.run([dotnet, str(assembly)], capture_output=True, text=True, timeout=120)
    wall = time.perf_counter() - clock
    log = AUDIT / f"{label}.log"
    log.write_text(result.stdout)
    (AUDIT / f"{label}.stderr").write_text(result.stderr)
    result.check_returncode()
    run = read_run(log)
    check_outputs(run, reference)
    summary = {"total_ms": run["total_ms"], "rbtree_ms": next(
        test["best_of_10_us"] / 1000 for test in run["tests"] if test["name"] == RBT)}
    executions.append({"label": label, "variant": variant, "startedUtc": started,
                       "wallSeconds": wall, "loadBefore": load, "loadAfter": os.getloadavg(),
                       "assemblySha256": sha(assembly), "logSha256": sha(log), "summary": summary})
    save("executions", executions)
    print(json.dumps({"label": label, "wallSeconds": wall, **summary}), flush=True)
    verify_sources()
verify_sources()
metadata.update({"finishedUtc": utc(), "sourceManifestsUnchanged": True})
save("metadata", metadata)
print(f"Validated {len(executions)} fresh App.main processes; normal generated sources unchanged.", flush=True)

"""Sample a frozen compiler on a copied Aff corpus; all timings are diagnostic."""
import argparse
import collections
import hashlib
import json
import os
from pathlib import Path
import re
import shutil
import subprocess

parser = argparse.ArgumentParser()
parser.add_argument("binary", type=Path)
parser.add_argument("workspace", type=Path, help="compare-aff.mjs *-runs directory")
parser.add_argument("out", type=Path)
parser.add_argument("--phase", choices=["all", "optimize"], default="all")
args = parser.parse_args()
args.binary, args.workspace, args.out = (p.resolve() for p in (args.binary, args.workspace, args.out))
args.out.mkdir(parents=True, exist_ok=False)
cwd = args.workspace / "inputs/gopurs-aff"
for name in [".purmeta", ".cache"]:
    shutil.rmtree(cwd / name, ignore_errors=True)
env = {k: v for k, v in os.environ.items()
       if not re.match(r"^(PURUST_|GOPURS_|NODE_|RUST|CARGO|GOGC$|GOMEMLIMIT$|GOMAXPROCS$|GODEBUG$)", k)}
command = [str(args.binary), "--main", "Test.Main", "--threaded", "--source", "output",
           "--out", str(args.out / "generated"), "--ffi-dir", os.path.relpath(args.workspace / "rust-ffi", cwd)]
with (args.out / "stdout").open("w") as stdout, (args.out / "stderr").open("w") as stderr:
    child = subprocess.Popen(command, cwd=cwd, env=env, stdout=stdout, stderr=subprocess.PIPE, text=True)
    if args.phase == "optimize":
        for line in child.stderr:
            stderr.write(line)
            if line.startswith("[purust] prepare:"):
                break
    sample = subprocess.Popen(["sample", str(child.pid), "10", "1", "-file", str(args.out / "sample.txt")],
                              stdout=subprocess.PIPE, stderr=subprocess.STDOUT, text=True)
    stderr.write(child.stderr.read())
    status = child.wait()
    sample_log, _ = sample.communicate()
    (args.out / "sample.log").write_text(sample_log)
assert status == 0, f"compiler failed: {status}"

stack = []
leaves, families, allocation_callers = (collections.Counter() for _ in range(3))
patterns = {
    "decoding": r"Optimizer_CoreFn_Json|Optimizer_CoreFn_TypeTable|Purs_Data_Argonaut",
    "Map/Set": r"Purs_Data_Map_Internal::|Purs_Data_Set::",
    "PBO": r"Purs_PureScript_Backend_Optimizer_",
    "type substitution": r"Optimizer_TypeSubstitution::|substituteNeutralTypes|instantiateNeutralType",
    "codegen": r"Purs_Purust_CodeGen::",
    "allocation": r"mi_malloc|mi_free|malloc|__rust_alloc|__rust_dealloc",
    "String clone": r"5alloc6string.*5clone",
    "Value clone": r"purust_core..Value.*Clone.*clone",
}


def finish():
    _, count, name, children = stack.pop()
    samples = count - children
    path = "\n".join(frame[2] for frame in stack) + "\n" + name
    if samples <= 0 or "Purs_" not in path:
        return
    if re.search(r"__ulock_wait|__psynch_cvwait|kevent|mach_msg.*trap|thread_switch", name):
        return
    leaves[name] += samples
    for label, pattern in patterns.items():
        if re.search(pattern, path):
            families[label] += samples
    if re.search(patterns["allocation"], name):
        for frame in reversed(stack):
            if "Purs_" in frame[2]:
                allocation_callers[frame[2]] += samples
                break


for line in (args.out / "sample.txt").read_text().splitlines():
    if line.startswith("Total number in stack"):
        break
    match = re.match(r"^([ +!:|]*)(\d+) (.+)$", line)
    if not match:
        continue
    prefix, count, name = match.groups()
    depth, count, name = len(prefix), int(count), name.split("  (in ")[0]
    while stack and stack[-1][0] >= depth:
        finish()
    if stack:
        stack[-1][3] += count
    stack.append([depth, count, name, 0])
while stack:
    finish()
result = {
    "command": command, "cwd": str(cwd), "phase": args.phase,
    "binary_sha256": hashlib.sha256(args.binary.read_bytes()).hexdigest(),
    "metric": "all-thread wall-clock stack samples with PureScript frames, known blocking leaves excluded; inclusive families overlap, not CPU percentages",
    "samples": sum(leaves.values()), "inclusive_families": dict(families),
    "top_exclusive": leaves.most_common(30), "allocation_callers": allocation_callers.most_common(40),
}
(args.out / "summary.json").write_text(json.dumps(result, indent=2) + "\n")
print(json.dumps(result, indent=2))

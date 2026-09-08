"""Reconstruct the recorded fast program and copy the current generated program."""
from pathlib import Path
import difflib
import hashlib
import json
import shutil

audit = Path(__file__).resolve().parent
root = audit.parent.parent
source = root / "output/Main"
historical = audit.parent / "sharpurs-tco-measure-20260908"
recorded = json.loads((historical / "inputs.json").read_text())["sourceSha256"]
sha = lambda path: hashlib.sha256(path.read_bytes()).hexdigest()
current_project = (source / "Program.fsproj").read_text()
old_project = current_project
for module in ["AppJavaFFI.fs", "AppJavaFFICheatcode.fs"]:
    line = f'    <Compile Include="{module}" />\n'
    assert old_project.count(line) == 1
    old_project = old_project.replace(line, "")
assert hashlib.sha256(old_project.encode()).hexdigest() == recorded["Program.fsproj"]
files = sorted(p for p in source.iterdir() if p.is_file() and p.suffix in {
    ".fs", ".cs", ".fsproj", ".csproj", ".props"
})
current_hashes = {p.name: sha(p) for p in files}
changed = [name for name, value in recorded.items() if current_hashes.get(name) != value]
extra = sorted(set(current_hashes) - set(recorded))
assert changed == ["Program.fsproj"], changed
assert extra == ["AppJavaFFI.fs", "AppJavaFFICheatcode.fs"], extra

for variant in ["reference", "current"]:
    target = audit / variant
    target.mkdir(exist_ok=True)
    for path in files:
        if variant == "reference" and path.name not in recorded:
            continue
        if variant == "reference" and path.name == "Program.fsproj":
            (target / path.name).write_text(old_project)
        else:
            shutil.copy2(path, target / path.name)
    if variant == "reference":
        assert all(sha(target / name) == value for name, value in recorded.items())

(audit / "project.diff").write_text("".join(difflib.unified_diff(
    old_project.splitlines(True),
    current_project.splitlines(True),
    fromfile="reference/Program.fsproj", tofile="current/Program.fsproj")))
(audit / "inputs.json").write_text(json.dumps({
    "referenceFileCount": len(recorded), "currentFileCount": len(files),
    "referenceSha256": recorded, "currentSha256": current_hashes,
    "changedCommonFiles": changed, "extraCurrentFiles": extra,
    "historicalInputsSha256": sha(historical / "inputs.json"),
    "officialReadmeSha256": sha(root.parent / "altbak.pub/README.md"),
    "entryPointUnchanged": True, "harness": "Original App.main, global warmup x3, local warmup x3, best of 10",
}, indent=2) + "\n")
print(f"Reference: all {len(recorded)} recorded input hashes match exactly.")
print(f"Current: {len(files)} inputs; only project includes and the two additional modules differ.")

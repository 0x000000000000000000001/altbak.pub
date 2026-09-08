"""Capture the regenerated sources and prepare official and diagnostic copies."""
from pathlib import Path
import difflib, hashlib, json, shutil

audit = Path(__file__).resolve().parent
root = audit.parent.parent
source = root / "run/bak/sharp/output/Main"
sha = lambda path: hashlib.sha256(path.read_bytes()).hexdigest()
before = json.loads((audit / "before-inputs.json").read_text())
for name, value in before.items():
    assert sha(audit / "before" / name) == value, f"Before snapshot changed: {name}"
names = sorted(path.name for path in source.iterdir() if path.is_file() and path.suffix in {".fs", ".cs", ".fsproj", ".csproj", ".props"})
assert names == sorted(before), "Generated input file list changed; review the experiment"
after = {name: sha(source / name) for name in names}
changed = [name for name in names if before[name] != after[name]]
assert "Test.RBTree.fs" in changed, changed
assert "EntryPoint.fs" not in changed and "Program.fsproj" not in changed and "Directory.Build.props" not in changed, changed
for name in changed:
    old = (audit / "before" / name).read_text()
    new = (source / name).read_text()
    (audit / (name + ".diff")).write_text("".join(difflib.unified_diff(old.splitlines(True), new.splitlines(True), fromfile="before/" + name, tofile="after/" + name)))
    (audit / (name + ".generated")).write_text(new)
for variant in ["after", "diagnostic-before", "diagnostic-after"]:
    target = audit / variant
    target.mkdir(exist_ok=True)
    origin = audit / "before" if variant == "diagnostic-before" else source
    for name in names:
        shutil.copy2(origin / name, target / name)
    if variant.startswith("diagnostic-"):
        shutil.copy2(audit / "MeasureEntryPoint.fs", target / "EntryPoint.fs")
(audit / "inputs.json").write_text(json.dumps({
    "changedFiles": changed, "beforeSha256": before, "afterSha256": after,
    "diagnosticEntryPointSha256": sha(audit / "MeasureEntryPoint.fs"),
    "officialReadmeSha256": sha(audit / "README.baseline.md"),
}, indent=2) + "\n")
print(f"Captured {len(names)} generated inputs; changed {changed}.")

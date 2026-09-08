"""Prepare two isolated generated programs differing only in the TCO binding."""
from pathlib import Path
import difflib
import hashlib
import json
import shutil
import subprocess

audit = Path(__file__).resolve().parent
root = audit.parent.parent
source = root / "output/Main"
previous = audit.parent / "sharpurs-int-emission-20260908"
sha = lambda path: hashlib.sha256(path.read_bytes()).hexdigest()
recorded = json.loads((previous / "metadata.json").read_text())["sha256"]
for name in ["Test.TCO.fs", "Sharpurs_Prelude.fs"]:
    expected = recorded[f"altbak.pub-sharpurs/output/Main/{name}"]
    assert sha(source / name) == expected, f"Inspect changed generated input: {name}"

files = sorted(p for p in source.iterdir() if p.is_file() and p.suffix in {
    ".fs", ".cs", ".fsproj", ".csproj", ".props"
})
for variant in ["before", "after"]:
    target = audit / variant
    target.mkdir(exist_ok=True)
    for path in files:
        shutil.copy2(path, target / path.name)
    shutil.copy2(audit / "MeasureEntryPoint.fs", target / "EntryPoint.fs")

subprocess.run(["patch", "--reverse", "--input", str(previous / "Test.TCO.diff"),
                str(audit / "before/Test.TCO.fs")], check=True)
changed = [p.name for p in files if sha(audit / "before" / p.name) != sha(audit / "after" / p.name)]
assert changed == ["Test.TCO.fs"], changed
(audit / "variants.diff").write_text("".join(difflib.unified_diff(
    (audit / "before/Test.TCO.fs").read_text().splitlines(True),
    (audit / "after/Test.TCO.fs").read_text().splitlines(True),
    fromfile="before/Test.TCO.fs", tofile="after/Test.TCO.fs")))
metadata = {
    "inputFileCount": len(files), "variantDifferences": changed,
    "sourceSha256": {p.name: sha(p) for p in files},
    "entryPointSha256": sha(audit / "MeasureEntryPoint.fs"),
    "beforeTcoSha256": sha(audit / "before/Test.TCO.fs"),
    "afterTcoSha256": sha(audit / "after/Test.TCO.fs"),
    "officialReadmeSha256": sha(root.parent / "altbak.pub/README.md"),
    "tastSha256": sha(root / "output/Test.TCO/corefn.json"),
}
(audit / "inputs.json").write_text(json.dumps(metadata, indent=2) + "\n")
print(f"Prepared {len(files)} files per variant; only Test.TCO.fs differs.")

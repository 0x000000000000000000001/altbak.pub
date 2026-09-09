"""Capture normal generated output after backend integration; never patch it."""
from datetime import datetime, timezone
import difflib
import json
from pathlib import Path
import shutil

from common import AUDIT, require, sha

before = json.loads((AUDIT / "before-inputs.json").read_text())
source = Path(before["sourceGeneratedDirectory"])
names = sorted(p.name for p in source.iterdir()
               if p.is_file() and p.suffix in {".fs", ".cs", ".fsproj", ".csproj", ".props"})
require(names == sorted(before["beforeSha256"]), "Build input file list changed")
require(len(names) == 355, "Expected the 355 previously captured build inputs")
require(all(sha(AUDIT / "before" / name) == expected
            for name, expected in before["beforeSha256"].items()), "Before snapshot changed")
hashes = {name: sha(source / name) for name in names}
changed = [name for name in names if hashes[name] != before["beforeSha256"][name]]
require(changed, "No generated change to measure")
protected = [name for name in names if Path(name).suffix != ".fs"
             or "FFI" in name or name in {"EntryPoint.fs", "App.fs", "Bench.fs"}]
require(not set(changed).intersection(protected), "Harness, FFI or build configuration changed")
require(sha(AUDIT / "README.baseline.md") == before["officialReadmeSha256"], "README snapshot changed")
target = AUDIT / "after"
target.mkdir(exist_ok=False)
for name in names:
    shutil.copy2(source / name, target / name)
require(hashes == {name: sha(target / name) for name in names}, "After snapshot copy mismatch")
require(hashes == {name: sha(source / name) for name in names}, "Generation changed during capture")
diffs = AUDIT / "diffs"
diffs.mkdir(exist_ok=False)
for name in changed:
    original, updated = (AUDIT / "before" / name).read_text(), (target / name).read_text()
    (diffs / f"{name}.diff").write_text("".join(difflib.unified_diff(
        original.splitlines(True), updated.splitlines(True),
        fromfile=f"before/{name}", tofile=f"after/{name}")))
    (diffs / f"{name}.generated").write_text(updated)
(AUDIT / "inputs.json").write_text(json.dumps({
    "capturedBeforeUtc": before["capturedUtc"],
    "capturedAfterUtc": datetime.now(timezone.utc).isoformat(),
    "beforeSha256": before["beforeSha256"], "afterSha256": hashes,
    "sourceGeneratedSha256": hashes,
    "sourceGeneratedDirectory": str(source), "changedFiles": changed,
    "unchangedProtectedFiles": protected,
    "officialReadmeSha256": before["officialReadmeSha256"],
}, indent=2) + "\n")
print(json.dumps({"inputsPerVariant": len(names), "changedFiles": changed,
                  "unchangedProtectedFiles": len(protected)}, indent=2))

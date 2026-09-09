"""Snapshot normal compiler output after integration; never patch generated F#."""
from datetime import datetime, timezone
import difflib
import json
from pathlib import Path
import shutil

from common import AUDIT, require, sha

manifest = AUDIT / "inputs.json"
inputs = json.loads(manifest.read_text())
source = Path(inputs["sourceGeneratedDirectory"])
target = AUDIT / "after"
require(not target.exists(), "After snapshot already exists; preserve the recorded inputs")
names = sorted(path.name for path in source.iterdir()
               if path.is_file() and path.suffix in {".fs", ".cs", ".fsproj", ".csproj", ".props"})
hashes = {name: sha(source / name) for name in names}
require(hashes, "Normal generated sources absent")
target.mkdir()
for name in names:
    shutil.copy2(source / name, target / name)
require(hashes == {name: sha(target / name) for name in names}, "Copy differs from generated sources")
require(hashes == {name: sha(source / name) for name in names}, "Normal sources changed during capture")
before = inputs["beforeSha256"]
changed = [name for name in sorted(set(before) | set(hashes)) if before.get(name) != hashes.get(name)]
inputs.update({"afterSha256": hashes, "sourceGeneratedSha256": hashes,
               "changedFiles": changed, "afterCapturedUtc": datetime.now(timezone.utc).isoformat(),
               "afterOrigin": "Unmodified output of the integrated Sharpurs backend"})
manifest.write_text(json.dumps(inputs, indent=2) + "\n")
diffs = []
for name in changed:
    old = (AUDIT / "before" / name).read_text().splitlines(True) if name in before else []
    new = (target / name).read_text().splitlines(True) if name in hashes else []
    diffs.extend(difflib.unified_diff(old, new, fromfile="before/" + name, tofile="after/" + name))
(AUDIT / "generated.diff").write_text("".join(diffs))
shutil.copy2(target / "Test.RBTree.fs", AUDIT / "Test.RBTree.fs.generated")
print(f"Captured {len(hashes)} normal generated inputs; {len(changed)} changed files, no manual F# edits.")

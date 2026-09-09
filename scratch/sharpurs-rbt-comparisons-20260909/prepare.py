"""Isolate two native Int comparisons in generated RBTree.ins only."""
from pathlib import Path
import difflib
import hashlib
import json
import re
import shutil

audit = Path(__file__).resolve().parent
project = audit.parent.parent
source = project / "run/bak/sharp/output/Main"
sha = lambda path: hashlib.sha256(path.read_bytes()).hexdigest()
names = sorted(p.name for p in source.iterdir()
               if p.is_file() and p.suffix in {".fs", ".cs", ".fsproj", ".csproj", ".props"})
assert "Test.RBTree.fs" in names and "EntryPoint.fs" in names
hashes = {name: sha(source / name) for name in names}
for variant in ["before", "after"]:
    target = audit / variant
    target.mkdir(exist_ok=False)
    for name in names:
        shutil.copy2(source / name, target / name)

path = audit / "after/Test.RBTree.fs"
original = path.read_text()
assert "Test_RBTree_Tusd_Ctor of Test_RBTree_Color * Test_RBTree_Tree * int * Test_RBTree_Tree" in original
lines = original.splitlines(keepends=True)
indices = [i for i, line in enumerate(lines) if line.startswith("let rec Test_RBTree_ins_tco ")]
assert len(indices) == 1
index = indices[0]
pattern = re.compile(r"(?<=\(match )\(\(unbox \(\(sharpurs_apply .*?(?= with \| LitBool true \(\))")
matches = list(pattern.finditer(lines[index]))
assert len(matches) == 2
replacements = []
for match, name, operator in zip(matches, ["lessThan", "greaterThan"], ["<", ">"]):
    old = match.group()
    assert f"Data_Ord_{name}" in old and "Data_Ord_ordInt" in old
    assert old.count("sharpurs_apply") == 3
    assert "((box x))" in old and "((box y))" in old
    replacements.append({"operation": name, "before": old,
                         "after": f"(box ((unbox<int> x) {operator} y))"})
line = lines[index]
for replacement in replacements:
    assert line.count(replacement["before"]) == 1
    line = line.replace(replacement["before"], replacement["after"])
assert "Data_Ord_" not in line
lines[index] = line
changed = "".join(lines)
path.write_text(changed)
assert sum(a != b for a, b in zip(original.splitlines(), changed.splitlines())) == 1
after = {name: sha(audit / "after" / name) for name in names}
assert [name for name in names if hashes[name] != after[name]] == ["Test.RBTree.fs"]
assert all(sha(source / name) == expected for name, expected in hashes.items())
shutil.copy2(project.parent / "altbak.pub/README.md", audit / "README.baseline.md")
(audit / "Test.RBTree.fs.diff").write_text("".join(difflib.unified_diff(
    original.splitlines(True), changed.splitlines(True),
    fromfile="before/Test.RBTree.fs", tofile="after/Test.RBTree.fs")))
(audit / "Test.RBTree.fs.generated").write_text(changed)
(audit / "replacements.json").write_text(json.dumps(replacements, indent=2) + "\n")
(audit / "inputs.json").write_text(json.dumps({
    "beforeSha256": hashes, "afterSha256": after,
    "sourceGeneratedSha256": hashes,
    "sourceGeneratedDirectory": str(source),
    "changedFiles": ["Test.RBTree.fs"],
    "officialReadmeSha256": sha(audit / "README.baseline.md"),
}, indent=2) + "\n")
print(f"Prepared {len(names)} inputs per variant; only two comparisons in ins changed; normal generation unchanged.")

"""Capture the generated change and construct diagnostic copies of both programs."""
from pathlib import Path
import difflib
import hashlib
import json
import shutil

audit = Path(__file__).resolve().parent
root = audit.parent.parent
source = root / "output/Main"
sha = lambda path: hashlib.sha256(path.read_bytes()).hexdigest()
before = json.loads((audit / "before-inputs.json").read_text())
(audit / "before").mkdir(exist_ok=True)
for name, value in before.items():
    origin = audit / "Test.Polymorphism.before.fs" if name == "Test.Polymorphism.fs" else source / name
    assert sha(origin) == value, f"Cannot reconstruct recorded input: {name}"
    shutil.copy2(origin, audit / "before" / name)
    assert sha(audit / "before" / name) == value, name
after = {name: sha(source / name) for name in before}
changed = [name for name in before if before[name] != after[name]]
assert changed == ["Test.Polymorphism.fs"], changed
old = (audit / "before/Test.Polymorphism.fs").read_text()
new = (source / "Test.Polymorphism.fs").read_text()
assert old.split("let Test_Polymorphism_act ")[0] == new.split("let Test_Polymorphism_act ")[0]
assert "let rec sharpurs_int_kernel" in new
diff = "".join(difflib.unified_diff(old.splitlines(True), new.splitlines(True),
    fromfile="before/Test.Polymorphism.fs", tofile="after/Test.Polymorphism.fs"))
(audit / "Test.Polymorphism.diff").write_text(diff)
(audit / "Test.Polymorphism.generated.fs").write_text(new)
for variant in ["after", "diagnostic-before", "diagnostic-after"]:
    target = audit / variant
    target.mkdir(exist_ok=True)
    origin = audit / "before" if variant == "diagnostic-before" else source
    for name in before:
        shutil.copy2(origin / name, target / name)
    if variant.startswith("diagnostic-"):
        shutil.copy2(audit / "MeasureEntryPoint.fs", target / "EntryPoint.fs")
(audit / "inputs.json").write_text(json.dumps({
    "changedFiles": changed, "beforeSha256": before, "afterSha256": after,
    "diagnosticEntryPointSha256": sha(audit / "MeasureEntryPoint.fs"),
    "officialReadmeSha256": sha(root.parent / "altbak.pub/README.md"),
}, indent=2) + "\n")
print(f"Only {changed[0]} differs among {len(before)} generated files; generic declarations unchanged.")

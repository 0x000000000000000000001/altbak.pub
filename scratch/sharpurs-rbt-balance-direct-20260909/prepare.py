"""Isolate direct four-argument calls to balance, keeping its object body."""
from pathlib import Path
import difflib
import hashlib
import json
import shutil

audit = Path(__file__).resolve().parent
project = audit.parent.parent
source = project / "run/bak/sharp/output/Main"
sha = lambda path: hashlib.sha256(path.read_bytes()).hexdigest()


def group_end(text, start=0):
    assert text[start] == "("
    depth = 0
    for i in range(start, len(text)):
        if text[i] == "(": depth += 1
        if text[i] == ")": depth -= 1
        if depth == 0: return i + 1
    raise ValueError("Unbalanced generated expression")


def strip_groups(text):
    text = text.strip()
    while text.startswith("(") and group_end(text) == len(text):
        text = text[1:-1].strip()
    return text


def strip_box(text):
    text = strip_groups(text)
    while text.startswith("box "):
        text = strip_groups(text[4:])
    return text


def split_apply(text):
    text = strip_groups(text)
    assert text.startswith("sharpurs_apply ")
    text = text[len("sharpurs_apply "):]
    end = group_end(text)
    function = text[:end]
    argument = text[end:].strip()
    assert argument.startswith("(") and group_end(argument) == len(argument)
    return function, argument


names = sorted(p.name for p in source.iterdir()
               if p.is_file() and p.suffix in {".fs", ".cs", ".fsproj", ".csproj", ".props"})
before = {name: sha(source / name) for name in names}
for variant in ["before", "after"]:
    target = audit / variant
    target.mkdir(exist_ok=False)
    for name in names: shutil.copy2(source / name, target / name)

original = (audit / "before/Test.RBTree.fs").read_text()
lines = original.splitlines(keepends=True)
balance_index, = [i for i, line in enumerate(lines) if line.startswith("let Test_RBTree_balance ")]
balance = lines[balance_index].rstrip("\n")
prefix = "let Test_RBTree_balance  = " + "".join(
    f"(box (fun ({name}: obj) -> " for name in ["v", "v1", "v2", "v3"])
assert balance.startswith(prefix) and balance.endswith(")" * 8)
body = balance[len(prefix):-8]
assert group_end(body) == len(body), "Body must be preserved verbatim"
direct = "let Test_RBTree_balance_direct (v: obj) (v1: obj) (v2: obj) (v3: obj) : obj = " + body
wrapper = prefix + "(Test_RBTree_balance_direct v v1 v2 v3)" + ")" * 8
lines[balance_index] = direct + "\n\n" + wrapper + "\n"

ins_index, = [i for i, line in enumerate(lines) if line.startswith("let rec Test_RBTree_ins_tco ")]
line = lines[ins_index]
calls = []
for start in range(len(line)):
    if not line.startswith("(sharpurs_apply ", start): continue
    end = group_end(line, start)
    expression = line[start:end]
    if "Test_RBTree_balance" not in expression or expression.count("sharpurs_apply") != 4: continue
    current, arguments = expression, []
    while strip_groups(current).startswith("sharpurs_apply "):
        function, argument = split_apply(current)
        arguments.insert(0, argument)
        current = strip_box(function)
    assert current == "Test_RBTree_balance" and len(arguments) == 4
    assert sum("Test_RBTree_ins_tco" in argument for argument in arguments) == 1
    replacement = "(Test_RBTree_balance_direct " + " ".join(arguments) + ")"
    calls.append({"before": expression, "after": replacement, "arguments": arguments})
assert len(calls) == 2
for call in calls:
    assert line.count(call["before"]) == 1
    line = line.replace(call["before"], call["after"])
assert "sharpurs_apply" not in line and line.count("Test_RBTree_balance_direct") == 2
lines[ins_index] = line
changed = "".join(lines)
assert changed.count(body) == 1
(audit / "after/Test.RBTree.fs").write_text(changed)
after = {name: sha(audit / "after" / name) for name in names}
assert [name for name in names if before[name] != after[name]] == ["Test.RBTree.fs"]
assert all(sha(source / name) == value for name, value in before.items())
shutil.copy2(project.parent / "altbak.pub/README.md", audit / "README.baseline.md")
(audit / "balance-body.fs.txt").write_text(body + "\n")
(audit / "replacements.json").write_text(json.dumps(calls, indent=2) + "\n")
(audit / "Test.RBTree.fs.diff").write_text("".join(difflib.unified_diff(
    original.splitlines(True), changed.splitlines(True), fromfile="before/Test.RBTree.fs", tofile="after/Test.RBTree.fs")))
(audit / "Test.RBTree.fs.generated").write_text(changed)
(audit / "inputs.json").write_text(json.dumps({
    "beforeSha256": before, "afterSha256": after,
    "sourceGeneratedSha256": before, "sourceGeneratedDirectory": str(source),
    "changedFiles": ["Test.RBTree.fs"],
    "balanceBodySha256": hashlib.sha256(body.encode()).hexdigest(),
    "officialReadmeSha256": sha(audit / "README.baseline.md"),
}, indent=2) + "\n")
print(f"Prepared {len(names)} inputs; balance body byte-identical, two direct calls, normal generation unchanged.")

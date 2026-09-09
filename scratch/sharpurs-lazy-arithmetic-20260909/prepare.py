"""Isolate two Int operations in buildThunks; preserve normal generated code."""
from pathlib import Path
import difflib
import hashlib
import json
import shutil
import subprocess

audit = Path(__file__).resolve().parent
project = audit.parent.parent
source = project / "run/bak/sharp/output/Main"
sha = lambda path: hashlib.sha256(path.read_bytes()).hexdigest()


def group_end(text, start=0):
    assert text[start] == "("
    depth = 0
    for index in range(start, len(text)):
        if text[index] == "(": depth += 1
        if text[index] == ")": depth -= 1
        if depth == 0: return index + 1
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


def applications(expression):
    current, arguments = strip_groups(expression), []
    while current.startswith("sharpurs_apply "):
        remaining = current[len("sharpurs_apply "):]
        end = group_end(remaining)
        argument = remaining[end:].strip()
        assert argument.startswith("(") and group_end(argument) == len(argument)
        arguments.insert(0, argument)
        current = strip_box(remaining[:end])
    return current, arguments


assert (audit / "README.baseline.md").exists(), "Capture official README first"
names = sorted(p.name for p in source.iterdir()
               if p.is_file() and p.suffix in {".fs", ".cs", ".fsproj", ".csproj", ".props"})
assert names and not (audit / "before").exists() and not (audit / "after").exists()
before = {name: sha(source / name) for name in names}
for variant in ["before", "after"]:
    target = audit / variant
    target.mkdir()
    for name in names: shutil.copy2(source / name, target / name)

filename = "Test.LazyEvaluation.fs"
original = (audit / "before" / filename).read_text()
lines = original.splitlines(keepends=True)
line_index, = [i for i, line in enumerate(lines)
               if line.startswith("let rec Test_LazyEvaluation_buildThunks_tco ")]
line = lines[line_index]
assert '"' not in line, "This small parser only handles the known string-free expression"
operations = {
    "Data_Ring_sub": ("Data_Ring_ringInt", "-"),
    "Data_Semiring_add": ("Data_Semiring_semiringInt", "+"),
}
replacements = []
for start in range(len(line)):
    if not line.startswith("(sharpurs_apply ", start): continue
    expression = line[start:group_end(line, start)]
    function, arguments = applications(expression)
    if function not in operations or len(arguments) != 3: continue
    dictionary, operator = operations[function]
    assert strip_box(arguments[0]) == dictionary
    assert strip_box(arguments[2]) == "1"
    if operator == "-":
        assert strip_box(arguments[1]) == "n"
    else:
        force, force_args = applications(strip_box(arguments[1]))
        assert force == "Test_LazyEvaluation_force" and len(force_args) == 1
        assert strip_box(force_args[0]) == "acc"
    left, right = arguments[1:]
    replacement = f"(box ((unbox<int> (box {left})) {operator} (unbox<int> (box {right}))))"
    replacements.append({"function": function, "dictionary": dictionary,
                         "before": expression, "after": replacement,
                         "left": left, "right": right})

assert len(replacements) == 2
assert {item["function"] for item in replacements} == set(operations)
patched_line = line
for item in replacements:
    assert patched_line.count(item["before"]) == 1
    assert item["after"].count(item["left"]) == 1
    patched_line = patched_line.replace(item["before"], item["after"])
assert patched_line.count("sharpurs_apply") == line.count("sharpurs_apply") - 6
for token in ["fun (", "Test_LazyEvaluation_force", "Test_LazyEvaluation_defer",
              "Test_LazyEvaluation_buildThunks_tco"]:
    assert patched_line.count(token) == line.count(token), token
assert not any(name in patched_line for name in operations)
assert patched_line != line
lines[line_index] = patched_line
changed = "".join(lines)
assert sum(a != b for a, b in zip(original.splitlines(), changed.splitlines())) == 1
(audit / "after" / filename).write_text(changed)
after = {name: sha(audit / "after" / name) for name in names}
assert [name for name in names if before[name] != after[name]] == [filename]
assert all(sha(source / name) == expected for name, expected in before.items())
(audit / (filename + ".before")).write_text(original)
(audit / (filename + ".generated")).write_text(changed)
(audit / (filename + ".diff")).write_text("".join(difflib.unified_diff(
    original.splitlines(True), changed.splitlines(True),
    fromfile="before/" + filename, tofile="after/" + filename)))
(audit / "replacements.json").write_text(json.dumps(replacements, indent=2) + "\n")
backend = project.parent / "sharpurs/sharpurs"
(audit / "inputs.json").write_text(json.dumps({
    "beforeSha256": before, "afterSha256": after,
    "sourceGeneratedSha256": before, "sourceGeneratedDirectory": str(source),
    "changedFiles": [filename], "changedLine": line_index + 1,
    "removedGenericCallsPerThunk": 6,
    "backendCommit": subprocess.check_output(["git", "rev-parse", "HEAD"], cwd=backend, text=True).strip(),
    "backendBundleSha256": sha(backend / "bin/sharpurs.js"),
    "officialReadmeSha256": sha(audit / "README.baseline.md"),
}, indent=2) + "\n")
print(f"Prepared {len(names)} inputs; two Int expressions on one line replaced; normal generation unchanged.")

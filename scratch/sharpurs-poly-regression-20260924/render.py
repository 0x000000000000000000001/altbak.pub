"""Render the captured optimized IR and locate unannotated Abs nodes."""
from collections import Counter
from pathlib import Path
import json

audit = Path(__file__).resolve().parent
read = lambda name: json.loads((audit / name).read_text())
tag = lambda value: value.get("$tag", "").split("$")[-1] if isinstance(value, dict) else ""

def walk(value):
    if isinstance(value, dict):
        yield value
        for child in value.values(): yield from walk(child)
    elif isinstance(value, list):
        for child in value: yield from walk(child)

def strip_typed(value):
    while tag(value) == "Typed": value = value["value1"]
    return value

def compact(value):
    if isinstance(value, list): return "[" + ", ".join(map(compact, value)) + "]"
    if not isinstance(value, dict): return str(value)
    name = tag(value)
    args = [compact(child) for key, child in value.items() if key != "$tag"]
    if name == "Just": return args[0]
    if name == "Qualified": return ".".join(args)
    if name == "Func": return args[0] + " -> " + args[1]
    return name + ("(" + ", ".join(args) + ")" if args else "")

def tree(value, indent=0):
    if isinstance(value, list):
        return [line for child in value for line in tree(child, indent)]
    name = tag(value)
    prefix = " " * indent
    if name == "Typed":
        return [prefix + "Typed " + compact(value["value0"])] + tree(value["value1"], indent + 2)
    if name == "Abs":
        return [prefix + "Abs " + compact(value["value0"])] + tree(value["value1"], indent + 2)
    if name in ["Local", "Var", "Lit"] or not isinstance(value, dict):
        return [prefix + compact(value)]
    if name == "LetRec":
        return [prefix + "LetRec level=" + compact(value["value0"])] + \
            [line for binding in value["value1"]
             for line in [prefix + "  binding " + str(binding["value0"])] + tree(binding["value1"], indent + 4)] + \
            tree(value["value2"], indent + 2)
    if name == "Op2":
        return [prefix + compact(value["value0"])] + tree(value["value1"], indent + 2) + tree(value["value2"], indent + 2)
    return [prefix + name] + [line for key, child in value.items() if key != "$tag" for line in tree(child, indent + 2)]

summary = {}
for path in sorted(audit.glob("*.optimized.json")):
    binding = read(path.name)
    expression = binding["expression"]
    nodes = list(walk(expression))
    counts = Counter(node["$tag"] for node in nodes if "$tag" in node)
    summary[binding["ident"]] = {"recursive": binding["recursive"], "nodes": dict(sorted(counts.items()))}
    (audit / (binding["ident"] + ".tree.txt")).write_text("\n".join(tree(expression)) + "\n")

# Find Abs nodes that are not wrapped in a Typed annotation and would be visited
# with expected=Nothing by the Sharpurs.Optimized envelope.
def abs_without_typed(value, parent_was_typed=False, path="root"):
    results = []
    if isinstance(value, list):
        for index, child in enumerate(value):
            results += abs_without_typed(child, parent_was_typed, f"{path}[{index}]")
        return results
    if not isinstance(value, dict):
        return results
    name = tag(value)
    if name == "Typed":
        return abs_without_typed(value["value1"], True, path + "/Typed")
    if name == "Abs":
        if not parent_was_typed:
            results.append((path, compact(value["value0"])))
        for key, child in value.items():
            if key != "$tag":
                results += abs_without_typed(child, False, f"{path}/{key}")
        return results
    for key, child in value.items():
        if key != "$tag":
            results += abs_without_typed(child, False, f"{path}/{key}")
    return results

for name in ["act", "polyLoop"]:
    expression = read(name + ".optimized.json")["expression"]
    report = abs_without_typed(expression)
    summary[name]["absWithoutTyped"] = report
    print(name, "Abs without Typed:", len(report))
    for path, binders in report[:8]:
        print("   ", path, binders)

(audit / "summary.json").write_text(json.dumps(summary, indent=2) + "\n")
print("Trees written.")

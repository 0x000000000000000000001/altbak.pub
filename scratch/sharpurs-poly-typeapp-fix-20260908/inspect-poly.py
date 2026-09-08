"""Inventory the captured nodes and assert the observed specialized loop shape."""
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
    if name == "Op2":
        return [prefix + compact(value["value0"])] + tree(value["value1"], indent + 2) + tree(value["value2"], indent + 2)
    return [prefix + name] + [line for key, child in value.items() if key != "$tag" for line in tree(child, indent + 2)]

original = read("input.parsed.json")
poly_type_apps = [node for node in walk(original) if tag(node) == "ExprTypeApp"
                  and tag(node["value1"]) == "ExprVar"
                  and compact(node["value1"]["value1"]) == "Test.Polymorphism.polyLoop"]
assert len(poly_type_apps) == 1
poly_type_app = poly_type_apps[0]
assert compact(poly_type_app["value2"]) == "Int"
dictionary_apps = [node for node in walk(original) if tag(node) == "ExprApp" and node["value1"] == poly_type_app]
assert len(dictionary_apps) == 1
dictionary = dictionary_apps[0]["value2"]
assert tag(dictionary) == "ExprVar" and compact(dictionary["value1"]) == "Test.Polymorphism.intMonoidish"
class_decl = original["classDecls"][0]
assert class_decl["name"] == "Monoidish" and class_decl["vars"] == ["a"]
(audit / "call.input.json").write_text(json.dumps({"typeApp": poly_type_app,
    "dictionaryArgument": dictionary, "classDecl": class_decl}, indent=2) + "\n")
summary = {"inputCall": {"typeArgument": "Int", "dictionary": "Test.Polymorphism.intMonoidish",
    "annotation": compact(poly_type_app["value0"]["type"]["value0"])}}
for path in sorted(audit.glob("*.optimized.json")):
    binding = read(path.name)
    expression = binding["expression"]
    nodes = list(walk(expression))
    counts = Counter(node["$tag"] for node in nodes if "$tag" in node)
    refs = sorted({compact(node["value0"]) for node in nodes if tag(node) == "Var"})
    summary[binding["ident"]] = {"recursive": binding["recursive"], "nodes": dict(sorted(counts.items())), "references": refs}
    (audit / (binding["ident"] + ".tree.txt")).write_text("\n".join(tree(expression)) + "\n")

act = read("act.optimized.json")["expression"]
loops = [node for node in walk(act) if tag(node) == "LetRec"]
assert len(loops) == 1
loop = loops[0]
assert loop["value0"] == 1 and len(loop["value1"]) == 1
go = loop["value1"][0]
assert go["value0"] == "go"
signature = go["value1"]["value0"]
assert compact(signature) == "[Int, Int] -> Int"
assert not any(tag(node) == "TypeVar" for node in walk(loop))
body = go["value1"]
parameters = []
while tag(strip_typed(body)) == "Abs":
    abstraction = strip_typed(body)
    parameters.extend(abstraction["value0"])
    body = abstraction["value1"]
assert [parameter["value1"] for parameter in parameters] == [2, 3]
branch = strip_typed(body)
assert tag(branch) == "Branch" and len(branch["value0"]) == 1
assert compact(strip_typed(branch["value0"][0]["value1"])) == "Local(v1, 3)"
recursive_call = strip_typed(branch["value1"])
assert tag(recursive_call) == "App"
assert compact(recursive_call["value0"]) == "Local(go, 1)"
assert len(recursive_call["value1"]) == 2
add = strip_typed(recursive_call["value1"][1])
assert tag(add) == "PrimOp" and compact(add["value0"]["value0"]) == "OpIntNum(OpAdd)"
assert compact(add["value0"]["value1"]) == "Local(v1, 3)"
assert compact(add["value0"]["value2"]) == "Lit(LitInt(1))"
entry = strip_typed(loop["value2"])
assert compact(entry["value0"]) == "Local(go, 1)"
assert [compact(strip_typed(arg)) for arg in entry["value1"]] == ["Local(dummy, 0)", "Lit(LitInt(0))"]
assert summary["act"]["references"] == ["Bench.opaque", "Data.Show.showIntImpl", "Effect.bindE", "Effect.pureE"]
assert summary["act"]["nodes"].get("Syntax$Accessor", 0) == 0
type_apps = [node for node in walk(act) if node.get("$tag") == "Syntax$TypeApp"]
assert len(type_apps) == 1
assert compact(type_apps[0]["value0"]) == "Var(Bench.opaque)"
assert compact(type_apps[0]["value1"]) == "Int"
assert summary["polyLoop"]["nodes"]["Syntax$Accessor"] == 2

summary["inlinedGo"] = {"recursiveLevel": 1, "parameterLevels": [2, 3], "signature": compact(signature),
                        "entry": "go dummy 0", "body": "if n == 0 then acc else go (n - 1) (acc + 1)",
                        "dictionaryAccesses": 0, "residualTypeVariables": dict(Counter(
                            node["value0"] for node in walk(go["value1"]) if tag(node) == "TypeVar"))}
(audit / "go.inlined.json").write_text(json.dumps(go["value1"], indent=2) + "\n")
(audit / "summary.json").write_text(json.dumps(summary, indent=2) + "\n")
print("Verified: generic dictionary accesses retained; specialized act has one local recursive loop, Int addition of 1, no dictionary access.")
print("Verified: inlined go carries [Int, Int] -> Int and no residual type variable.")
print("Verified: the opaque FFI keeps its explicit @Int; the generic loop keeps its dictionary accesses.")

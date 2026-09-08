"""Check the desired type contract; --expect-current verifies the known regression."""
from pathlib import Path
import json
import sys

audit = Path(__file__).resolve().parent
read = lambda name: json.loads((audit / name).read_text())
tag = lambda value: value.get("$tag", "").split("$")[-1] if isinstance(value, dict) else ""

def walk(value):
    if isinstance(value, dict):
        yield value
        for child in value.values(): yield from walk(child)
    elif isinstance(value, list):
        for child in value: yield from walk(child)

def type_text(value):
    name = tag(value)
    if name == "TypeVar": return "TypeVar(" + value["value0"] + ")"
    if name == "Func":
        return "[" + ", ".join(map(type_text, value["value0"])) + "] -> " + type_text(value["value1"])
    if name == "ForAll": return "forall " + " ".join(value["value0"]) + ". " + type_text(value["value1"])
    return name

def typed_chain(value):
    annotations = []
    while tag(value) == "Typed":
        annotations.append(value["value0"])
        value = value["value1"]
    return annotations, value

original = read("input.parsed.json")
assert len(original["decls"]) == 4
inputs = {}
for group in original["decls"]:
    assert tag(group) == "NonRec"
    binding = group["value0"]
    inputs[binding["value1"]] = binding

expected = {"identity": "TypeVar(a)", "useInt": "Int", "useString": "String", "generic": "TypeVar(b)"}
observations = {}
failures = []
for name, wanted in expected.items():
    binding = inputs[name]
    source_type_apps = [node for node in walk(binding["value2"]) if tag(node) == "ExprTypeApp"]
    if name == "identity":
        assert source_type_apps == []
    else:
        assert len(source_type_apps) == 1
        application = source_type_apps[0]
        assert type_text(application["value2"]) == wanted
        assert tag(application["value1"]) == "ExprVar"
        reference = application["value1"]["value1"]
        assert reference["value0"]["value0"] == "TypeAppFixture" and reference["value1"] == "identity"
    source_signature = binding["value0"]["type"]["value0"]
    if name in ["identity", "generic"]:
        assert tag(source_signature) == "ForAll"
        assert source_signature["value0"] == (["a"] if name == "identity" else ["b"])

    expression = read(name + ".optimized.json")["expression"]
    signatures, abstraction = typed_chain(expression)
    assert signatures and tag(abstraction) == "Abs"
    assert len(abstraction["value0"]) == 1
    parameter_level = abstraction["value0"][0]["value1"]
    body_types, body = typed_chain(abstraction["value1"])
    assert body_types and tag(body) == "Local" and body["value1"] == parameter_level
    assert not any(tag(node) in ["App", "Var", "TypeApp"] for node in walk(expression))

    # A future fix may retain ForAll wrappers. Check their body without requiring
    # today's quantifier omission, or any particular number of Typed wrappers.
    signature_bodies = []
    for signature in signatures:
        while tag(signature) == "ForAll": signature = signature["value1"]
        signature_bodies.append(type_text(signature))
    actual_types = list(map(type_text, body_types))
    passed = all(signature == f"[{wanted}] -> {wanted}" for signature in signature_bodies)
    passed = passed and all(actual == wanted for actual in actual_types)
    if not passed: failures.append(name)
    observations[name] = {
        "inputSignature": type_text(source_signature),
        "inputTypeArguments": [type_text(node["value2"]) for node in source_type_apps],
        "optimizedSignatures": list(map(type_text, signatures)),
        "optimizedBodyAnnotations": actual_types, "expectedBodyType": wanted,
        "valueShape": "one parameter, returned unchanged", "typeContractPassed": passed,
    }
    print(f"{name}: {'PASS' if passed else 'FAIL'}; body = {actual_types}; expected only {wanted}")

summary = {"observations": observations, "typeContractPassed": not failures, "failingBindings": failures,
           "inputAndValueShapeChecksPassed": True}
(audit / "summary.json").write_text(json.dumps(summary, indent=2) + "\n")
if sys.argv[1:] == ["--expect-current"]:
    assert failures == ["useInt", "useString", "generic"]
    assert {name: item["optimizedBodyAnnotations"] for name, item in observations.items()} == {
        "identity": ["TypeVar(a)"], "useInt": ["Int", "TypeVar(a)"],
        "useString": ["String", "TypeVar(a)"], "generic": ["TypeVar(b)", "TypeVar(a)"],
    }
    print("Known regression reproduced in all three instantiations; original identity remains generic.")
elif sys.argv[1:]:
    raise SystemExit("Usage: python3 check-fixture.py [--expect-current]")
else:
    print(f"Desired type contract: {len(expected) - len(failures)}/{len(expected)} bindings pass.")
    raise SystemExit(1 if failures else 0)

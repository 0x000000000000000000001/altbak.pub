from pathlib import Path
import json
audit = Path(__file__).resolve().parent
read = lambda name: json.loads((audit / name).read_text())
tag = lambda v: v.get("$tag","").split("$")[-1] if isinstance(v, dict) else ""
def compact(value):
    if isinstance(value, list): return "[" + ", ".join(map(compact, value)) + "]"
    if not isinstance(value, dict): return str(value)
    name = tag(value); args = [compact(c) for k, c in value.items() if k != "$tag"]
    if name == "Just": return args[0]
    if name == "Qualified": return ".".join(args)
    if name == "Func": return args[0] + " -> " + args[1]
    if name == "ADT": return args[-2] + " " + args[-1] if len(args)>=2 else name
    return name + ("(" + ", ".join(args) + ")" if args else "")
def tree(value, indent=0):
    if isinstance(value, list):
        return [l for c in value for l in tree(c, indent)]
    prefix = " " * indent
    name = tag(value)
    if name == "Typed":
        return [prefix + "Typed " + compact(value["value0"])] + tree(value["value1"], indent + 2)
    if name == "Branch":
        lines = [prefix + "Branch"]
        for pair in value["value0"]:
            lines.append(prefix + "  if " + compact(pair["value0"]))
            lines += tree(pair["value1"], indent + 4)
        lines.append(prefix + "  else")
        lines += tree(value["value1"], indent + 4)
        return lines
    if name == "App":
        return [prefix + "App " + compact(value["value0"])] + tree(value["value1"], indent + 2)
    if name == "PrimOp":
        return [prefix + "PrimOp " + compact(value["value0"])] + tree(value["value1"], indent + 2) + (tree(value["value2"], indent + 2) if "value2" in value else [])
    if name == "Let":
        return [prefix + "Let level=" + compact(value["value0"])] + tree(value["value1"], indent + 2) + tree(value["value2"], indent + 2)
    if name in ("Local","Var","Lit") or not isinstance(value, dict):
        return [prefix + compact(value)]
    return [prefix + name] + [l for k, c in value.items() if k != "$tag" for l in tree(c, indent + 2)]
for name in ["balance", "ins"]:
    expr = read(name + ".optimized.json")["expression"]
    (audit / (name + ".tree.txt")).write_text("\n".join(tree(expr)) + "\n")
    print(name, "written")

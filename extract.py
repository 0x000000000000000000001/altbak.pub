import json

with open("execRWST.json", "r") as f:
    data = json.load(f)

def find_case(node):
    if isinstance(node, dict):
        if node.get("type") == "Case":
            print("Found Case:")
            print(json.dumps(node.get("caseAlternatives"), indent=2)[:500])
        for k, v in node.items():
            find_case(v)
    elif isinstance(node, list):
        for item in node:
            find_case(item)

find_case(data)

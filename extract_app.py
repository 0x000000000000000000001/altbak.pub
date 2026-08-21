import json

with open("execRWST.json", "r") as f:
    data = json.load(f)

def find_app(node):
    if isinstance(node, dict):
        if node.get("type") == "App":
            typ = node.get("annotation", {}).get("type")
            if typ:
                print("Found App Type:", json.dumps(typ))
        for k, v in node.items():
            find_app(v)
    elif isinstance(node, list):
        for item in node:
            find_app(item)

find_app(data)

import sys

with open('/Users/0x1/Documents/htdocs/altbak.pub/output/Test.RBTree/Test_RBTree.go', 'r') as f:
    content = f.read()

arena_code = """
var arena_T = make([]Data_Test_RBTree_T, 10000000)
var arenaIdx_T = 0

func alloc_T(v0, v1 gopurs_runtime.Value, v2 int64, v3 gopurs_runtime.Value) unsafe.Pointer {
	idx := arenaIdx_T
	arenaIdx_T++
	node := &arena_T[idx]
	node.V0 = v0
	node.V1 = v1
	node.V2 = v2
	node.V3 = v3
	return unsafe.Pointer(node)
}
"""

content = content.replace("func Is_Data_Test_RBTree_T(v gopurs_runtime.Value) bool {", arena_code + "\nfunc Is_Data_Test_RBTree_T(v gopurs_runtime.Value) bool {")

def patch_braces(text):
    start = 0
    while True:
        idx = text.find('&Data_Test_RBTree_T{', start)
        if idx == -1:
            break
        # find matching brace
        depth = 0
        end = -1
        for i in range(idx + len('&Data_Test_RBTree_T{'), len(text)):
            if text[i] == '{':
                depth += 1
            elif text[i] == '}':
                if depth == 0:
                    end = i
                    break
                depth -= 1
        
        if end != -1:
            inner = text[idx + len('&Data_Test_RBTree_T{'):end]
            # Replace inner commas to split args
            # Actually, just parse it simply since it's always v0, v1, v2, v3
            # Or we can just leave it as is if we define alloc_T to take the struct directly:
            # wait, if we define alloc_T to take a value:
            # We must pass the arguments. 
            pass
        start = idx + 1
    return text

# Actually, the simplest way to avoid heap allocation while keeping the syntax `&Data_Test_RBTree_T{...}` is to redefine the `Data_Test_RBTree_T` struct locally? No.

# Let's write a simple nested brace matcher to extract the 4 arguments.
def extract_args(inner):
    args = []
    depth = 0
    current = ""
    for c in inner:
        if c == '{': depth += 1
        elif c == '}': depth -= 1
        elif c == ',' and depth == 0:
            args.append(current.strip())
            current = ""
            continue
        current += c
    args.append(current.strip())
    return args

def replace_all(text):
    while True:
        idx = text.find('&Data_Test_RBTree_T{')
        if idx == -1: break
        depth = 0
        end = -1
        for i in range(idx + len('&Data_Test_RBTree_T{'), len(text)):
            if text[i] == '{': depth += 1
            elif text[i] == '}':
                if depth == 0:
                    end = i
                    break
                depth -= 1
        if end != -1:
            inner = text[idx + len('&Data_Test_RBTree_T{'):end]
            args = extract_args(inner)
            if len(args) == 4:
                replacement = f"alloc_T({args[0]}, {args[1]}, {args[2]}, {args[3]})"
                text = text[:idx] + replacement + text[end+1:]
            else:
                print("Args length is not 4:", args)
                break
        else:
            break
    return text

content = replace_all(content)

with open('/Users/0x1/Documents/htdocs/altbak.pub/output/Test.RBTree/Test_RBTree.go', 'w') as f:
    f.write(content)


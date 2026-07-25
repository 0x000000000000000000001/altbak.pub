import re

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

# Replace &Data_Test_RBTree_T{a, b, c, d} with alloc_T(a, b, c, d)
content = re.sub(r'&Data_Test_RBTree_T\{([^,]+),\s*([^,]+),\s*([^,]+),\s*([^\}]+)\}', r'alloc_T(\1, \2, \3, \4)', content)

with open('/Users/0x1/Documents/htdocs/altbak.pub/output/Test.RBTree/Test_RBTree.go', 'w') as f:
    f.write(content)

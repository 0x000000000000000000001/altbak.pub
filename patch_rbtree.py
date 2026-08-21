import re

with open('run/bak/go/output/purescript/Data_Map_Internal.go', 'r') as f:
    content = f.read()

patch = """func Call_Data_Map_Internal_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var go_insert func(node *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node
	go_insert = func(node *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
		if node == nil {
			return &Constructor_Data_Map_Internal_Node{1, 1, 1, k_1_loop, v_2_loop, nil, nil}
		}
		
		// MONOMORPHIZED INT COMPARE
		k1 := k_1_loop.IntVal
		k2 := node.V2.IntVal
		
		if k1 < k2 { // LT
			left := go_insert(node.V4)
			return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](Call_Data_Map_Internal_unsafeBalancedNode(
				node.V2, node.V3,
				gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(left)},
				gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(node.V5)},
			))
		} else if k1 > k2 { // GT
			right := go_insert(node.V5)
			return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](Call_Data_Map_Internal_unsafeBalancedNode(
				node.V2, node.V3,
				gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(node.V4)},
				gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(right)},
			))
		} else { // EQ
			return &Constructor_Data_Map_Internal_Node{1, 1, 1, k_1_loop, v_2_loop, node.V4, node.V5}
		}
	}
	
	return gopurs_runtime.Func(func(m gopurs_runtime.Value) gopurs_runtime.Value {
		var root *Constructor_Data_Map_Internal_Node
		if m.UnsafePtr != nil {
			root = (*Constructor_Data_Map_Internal_Node)(m.UnsafePtr)
		}
		newRoot := go_insert(root)
		return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(newRoot)}
	})
}
"""

start_idx = content.find("func Call_Data_Map_Internal_insert(dictOrd_0_loop")
end_idx = content.find("func Call_Data_Map_Internal_foldSubmapBy(dictOrd_0_loop", start_idx)

if start_idx != -1 and end_idx != -1:
    new_content = content[:start_idx] + patch + content[end_idx:]
    with open('run/bak/go/output/purescript/Data_Map_Internal.go', 'w') as f:
        f.write(new_content)
    print("Patched successfully.")
else:
    print("Could not find insert function bounds.")

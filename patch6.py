import re

with open('output/purescript/Data_Map_Internal.go', 'r') as f:
    content = f.read()

# Replace Call_Data_Map_Internal_insert__4289641298 body!
insert_code = """
func Call_Data_Map_Internal_insert_go(dictOrd_0 *Constructor_Data_Ord_Ord, k_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr == nil) {
__t3 = (&Constructor_Data_Map_Internal_Node{1, 1, 1, k_1, v_2, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)})
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
v2_5_1 := uint32(gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2).IntVal)
_ = v2_5_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (v2_5_1 == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](Call_Data_Map_Internal_insert_go(dictOrd_0, k_1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (v2_5_1 == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](Call_Data_Map_Internal_insert_go(dictOrd_0, k_1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (v2_5_1 == 103233856) {
__t2 = (&Constructor_Data_Map_Internal_Node{(*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V1, k_1, v_2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
}

func Call_Data_Map_Internal_insert__4289641298(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
var k_1 gopurs_runtime.Value = k_1_loop
var v_2 gopurs_runtime.Value = v_2_loop
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
    return Call_Data_Map_Internal_insert_go(dictOrd_0, k_1, v_2, v1_4)
})
}
"""

start_idx = content.find("func Call_Data_Map_Internal_insert__4289641298(")
end_idx = content.find("return go__go_3_0_73\n}", start_idx) + len("return go__go_3_0_73\n}")

new_content = content[:start_idx] + insert_code + content[end_idx:]

with open('output/purescript/Data_Map_Internal.go', 'w') as f:
    f.write(new_content)

print("Patched successfully")

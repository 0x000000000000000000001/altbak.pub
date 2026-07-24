package Test_RBTree

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	unsafe "unsafe"
)

var R gopurs_runtime.Value
var once_R sync.Once
func Get_R() gopurs_runtime.Value {
	once_R.Do(func() {
		R = gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}
	})
	return R
}

var B gopurs_runtime.Value
var once_B sync.Once
func Get_B() gopurs_runtime.Value {
	once_B.Do(func() {
		B = gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}
	})
	return B
}

var E gopurs_runtime.Value
var once_E sync.Once
func Get_E() gopurs_runtime.Value {
	once_E.Do(func() {
		E = gopurs_runtime.Value{Type: 9, IntVal: 1548554223, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_E{})}
	})
	return E
}

var T gopurs_runtime.Value
var once_T sync.Once
func Get_T() gopurs_runtime.Value {
	once_T.Do(func() {
		T = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{value0, value1, value2, value3})}
})
})
})
})
	})
	return T
}

var max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		max = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_max(x_0_box, y_1_box)
})
	})
	return max
}

var makeBlack gopurs_runtime.Value
var once_makeBlack sync.Once
func Get_makeBlack() gopurs_runtime.Value {
	once_makeBlack.Do(func() {
		makeBlack = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3983586014) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)(v_0.UnsafePtr).V1, (*Data_Test_RBTree_T)(v_0.UnsafePtr).V2, (*Data_Test_RBTree_T)(v_0.UnsafePtr).V3})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1548554223) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1548554223, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_E{})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}()
})
	})
	return makeBlack
}

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Red-Black Tree (100k Worst-Case Insertions):"))
	})
	return describe
}

var depth gopurs_runtime.Value
var once_depth sync.Once
func Get_depth() gopurs_runtime.Value {
	once_depth.Do(func() {
		depth = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
depth:
for {
if false { continue depth }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1548554223) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3983586014) {
__local_var_1_1 := gopurs_runtime.Apply(Get_depth(), (*Data_Test_RBTree_T)(v_0.UnsafePtr).V1)
_ = __local_var_1_1
__local_var_2_2 := gopurs_runtime.Apply(Get_depth(), (*Data_Test_RBTree_T)(v_0.UnsafePtr).V3)
_ = __local_var_2_2
var __t3 gopurs_runtime.Value
{
if __local_var_1_1.IntVal > __local_var_2_2.IntVal {
__t3 = gopurs_runtime.Int(1 + __local_var_1_1.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Int(1 + __local_var_2_2.IntVal)
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}()
})
	})
	return depth
}

var balance gopurs_runtime.Value
var once_balance sync.Once
func Get_balance() gopurs_runtime.Value {
	once_balance.Do(func() {
		balance = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_balance(v_0_box, v1_1_box, v2_2_box, v3_3_box)
})
	})
	return balance
}

var ins gopurs_runtime.Value
var once_ins sync.Once
func Get_ins() gopurs_runtime.Value {
	once_ins.Do(func() {
		ins = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ins(v_0_box, v1_1_box)
})
	})
	return ins
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insert(x_0_box, s_1_box)
})
	})
	return insert
}

var buildTree gopurs_runtime.Value
var once_buildTree sync.Once
func Get_buildTree() gopurs_runtime.Value {
	once_buildTree.Do(func() {
		buildTree = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_buildTree(v_0_box, v1_1_box)
})
	})
	return buildTree
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(100000))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = dummy_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply(Get_depth(), Call_buildTree(dummy_1_1, gopurs_runtime.Value{Type: 9, IntVal: 1548554223, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_E{})})))), gopurs_runtime.Value{})
})
}()
	})
	return act
}

type Data_Test_RBTree_R struct {
	
}
func Is_Data_Test_RBTree_R(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3668501016
}

type Data_Test_RBTree_B struct {
	
}
func Is_Data_Test_RBTree_B(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1583507464
}

type Data_Test_RBTree_E struct {
	
}
func Is_Data_Test_RBTree_E(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1548554223
}

type Data_Test_RBTree_T struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}
func Is_Data_Test_RBTree_T(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3983586014
}

func Call_max(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
var __t0 gopurs_runtime.Value
{
if x_0.IntVal > y_1.IntVal {
__t0 = x_0
goto end_branch_0
} else {

}
}
{
__t0 = y_1
}
end_branch_0:
return __t0
}

func Call_balance(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value, v3_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var v3_3 gopurs_runtime.Value = v3_3_loop
_ = v3_3
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1583507464) {
var __t1 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3983586014) {
var __t2 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V0.IntVal == 3668501016) {
var __t3 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1.Type == 9 && (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1.IntVal == 3983586014) {
var __t4 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1.UnsafePtr).V0.IntVal == 3668501016) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1.UnsafePtr).V1, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1.UnsafePtr).V3})}, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3, v2_2, v3_3})}})}
goto end_branch_4
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.IntVal == 3983586014) {
var __t5 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V3, v2_2, v3_3})}})}
goto end_branch_5
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0.IntVal == 3668501016) {
var __t6 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.IntVal == 3983586014) {
var __t7 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0.IntVal == 3668501016) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V1})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V3, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3})}})}
goto end_branch_7
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_7:
__t6 = __t7
goto end_branch_6
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0.IntVal == 3668501016) {
var __t8 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.IntVal == 3983586014) {
var __t9 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0.IntVal == 3668501016) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V1})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V3, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3})}})}
goto end_branch_9
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_9:
__t8 = __t9
goto end_branch_8
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_8:
__t4 = __t8
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.IntVal == 3983586014) {
var __t10 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3.UnsafePtr).V3, v2_2, v3_3})}})}
goto end_branch_10
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0.IntVal == 3668501016) {
var __t11 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.IntVal == 3983586014) {
var __t12 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0.IntVal == 3668501016) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V1})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V3, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3})}})}
goto end_branch_12
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_12:
__t11 = __t12
goto end_branch_11
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_11:
__t10 = __t11
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_10:
__t3 = __t10
goto end_branch_3
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0.IntVal == 3668501016) {
var __t13 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.IntVal == 3983586014) {
var __t14 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0.IntVal == 3668501016) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V1})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V3, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3})}})}
goto end_branch_14
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_14:
__t13 = __t14
goto end_branch_13
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_13:
__t3 = __t13
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0.IntVal == 3668501016) {
var __t15 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.IntVal == 3983586014) {
var __t16 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0.IntVal == 3668501016) {
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V1})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V3, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3})}})}
goto end_branch_16
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_16:
__t15 = __t16
goto end_branch_15
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_15:
__t2 = __t15
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V0.IntVal == 3668501016) {
var __t17 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.IntVal == 3983586014) {
var __t18 gopurs_runtime.Value
{
if ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V0.IntVal == 3668501016) {
__t18 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V1})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1.UnsafePtr).V3, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3})}})}
goto end_branch_18
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t18 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_18:
__t17 = __t18
goto end_branch_17
} else {

}
}
{
if ((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.Type == 9 && (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.IntVal == 3983586014) && ((*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.Type == 9 && (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V0.IntVal == 3668501016) {
__t17 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, v1_1, v2_2, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V1})}, (*Data_Test_RBTree_T)(v3_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V1, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V2, (*Data_Test_RBTree_T)((*Data_Test_RBTree_T)(v3_3.UnsafePtr).V3.UnsafePtr).V3})}})}
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_17:
__t1 = __t17
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{v_0, v1_1, v2_2, v3_3})}
}
end_branch_0:
return __t0
}

func Call_ins(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
ins:
for {
if false { continue ins }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 1548554223) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_R{})}, gopurs_runtime.Value{Type: 9, IntVal: 1548554223, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_E{})}, v_0, gopurs_runtime.Value{Type: 9, IntVal: 1548554223, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_E{})}})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 3983586014) {
var __t1 gopurs_runtime.Value
{
if v_0.IntVal < (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2.IntVal {
__t1 = Call_balance((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V0, Call_ins(v_0, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1), (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3)
goto end_branch_1
} else {

}
}
{
if v_0.IntVal > (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2.IntVal {
__t1 = Call_balance((*Data_Test_RBTree_T)(v1_1.UnsafePtr).V0, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2, Call_ins(v_0, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{(*Data_Test_RBTree_T)(v1_1.UnsafePtr).V0, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V1, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V2, (*Data_Test_RBTree_T)(v1_1.UnsafePtr).V3})}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_insert(x_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
__local_var_2_0 := Call_ins(x_0, s_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3983586014) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_T{gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_B{})}, (*Data_Test_RBTree_T)(__local_var_2_0.UnsafePtr).V1, (*Data_Test_RBTree_T)(__local_var_2_0.UnsafePtr).V2, (*Data_Test_RBTree_T)(__local_var_2_0.UnsafePtr).V3})}
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 1548554223) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1548554223, UnsafePtr: unsafe.Pointer(&Data_Test_RBTree_E{})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_buildTree(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
buildTree:
for {
if false { continue buildTree }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if v_0.IntVal == 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
v_0_loop = gopurs_runtime.Int(v_0.IntVal - 1)
v1_1_loop = Call_insert(v_0, v1_1)
continue buildTree
__t0 = gopurs_runtime.Value{}
}
end_branch_0:
return __t0
}
}



package Test_RBTree

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	unsafe "unsafe"
)

var cache_R gopurs_runtime.Value
var once_R sync.Once
func Get_R() gopurs_runtime.Value {
	once_R.Do(func() {
		cache_R = gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}
	})
	return cache_R
}

var cache_B gopurs_runtime.Value
var once_B sync.Once
func Get_B() gopurs_runtime.Value {
	once_B.Do(func() {
		cache_B = gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}
	})
	return cache_B
}

var cache_E gopurs_runtime.Value
var once_E sync.Once
func Get_E() gopurs_runtime.Value {
	once_E.Do(func() {
		cache_E = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: nil}
	})
	return cache_E
}

var cache_T gopurs_runtime.Value
var once_T sync.Once
func Get_T() gopurs_runtime.Value {
	once_T.Do(func() {
		cache_T = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, value0, (*Constructor_T)(value1.UnsafePtr), value2.IntVal, (*Constructor_T)(value3.UnsafePtr)})}
})
})
})
})
	})
	return cache_T
}

var cache_max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		cache_max = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_max(x_0_box.IntVal, y_1_box.IntVal))
})
	})
	return cache_max
}

var cache_makeBlack gopurs_runtime.Value
var once_makeBlack sync.Once
func Get_makeBlack() gopurs_runtime.Value {
	once_makeBlack.Do(func() {
		cache_makeBlack = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_makeBlack((*Constructor_T)(v_0_box.UnsafePtr)))}
})
	})
	return cache_makeBlack
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Red-Black Tree (100k Worst-Case Insertions):"))
	})
	return cache_describe
}

var cache_depth gopurs_runtime.Value
var once_depth sync.Once
func Get_depth() gopurs_runtime.Value {
	once_depth.Do(func() {
		cache_depth = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_depth((*Constructor_T)(v_0_box.UnsafePtr)))
})
	})
	return cache_depth
}

var cache_balance gopurs_runtime.Value
var once_balance sync.Once
func Get_balance() gopurs_runtime.Value {
	once_balance.Do(func() {
		cache_balance = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_balance(v_0_box, (*Constructor_T)(v1_1_box.UnsafePtr), v2_2_box.IntVal, (*Constructor_T)(v3_3_box.UnsafePtr)))}
})
	})
	return cache_balance
}

var cache_ins gopurs_runtime.Value
var once_ins sync.Once
func Get_ins() gopurs_runtime.Value {
	once_ins.Do(func() {
		cache_ins = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_ins(v_0_box.IntVal, (*Constructor_T)(v1_1_box.UnsafePtr)))}
})
	})
	return cache_ins
}

var cache_insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		cache_insert = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_insert(x_0_box.IntVal, (*Constructor_T)(s_1_box.UnsafePtr)))}
})
	})
	return cache_insert
}

var cache_buildTree gopurs_runtime.Value
var once_buildTree sync.Once
func Get_buildTree() gopurs_runtime.Value {
	once_buildTree.Do(func() {
		cache_buildTree = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_buildTree(v_0_box.IntVal, (*Constructor_T)(v1_1_box.UnsafePtr)))}
})
	})
	return cache_buildTree
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(100000)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_depth((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_buildTree(dummy_0.IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: nil}.UnsafePtr)))}.UnsafePtr)))))
}))
	})
	return cache_act
}

type Constructor_R struct {
	Rc uint32
}


type Constructor_B struct {
	Rc uint32
}


type Constructor_E struct {
	Rc uint32
}


type Constructor_T struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 *Constructor_T
	V2 int64
	V3 *Constructor_T
}


func Call_max(x_0_loop int64, y_1_loop int64) int64 {
var x_0 int64 = x_0_loop
_ = x_0
var y_1 int64 = y_1_loop
_ = y_1
var __t0 gopurs_runtime.Value
{
if (x_0) > (y_1) {
__t0 = gopurs_runtime.Int(x_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Int(y_1)
}
end_branch_0:
return __t0.IntVal
}

func Call_makeBlack(v_0_loop *Constructor_T) *Constructor_T {
var v_0 *Constructor_T = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)}.UnsafePtr)})}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (*Constructor_T)(__t0.UnsafePtr)
}

func Call_depth(v_0_loop *Constructor_T) int64 {
depth:
for {
if false { continue depth }
var v_0 *Constructor_T = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__local_var_1_1 := gopurs_runtime.Int(Call_depth((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr)))
_ = __local_var_1_1
__local_var_2_2 := gopurs_runtime.Int(Call_depth((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)}.UnsafePtr)))
_ = __local_var_2_2
var __t3 gopurs_runtime.Value
{
if (__local_var_1_1.IntVal) > (__local_var_2_2.IntVal) {
__t3 = gopurs_runtime.Int((1) + (__local_var_1_1.IntVal))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Int((1) + (__local_var_2_2.IntVal))
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
return __t0.IntVal
}
}

func Call_balance(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_T, v2_2_loop int64, v3_3_loop *Constructor_T) *Constructor_T {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_T = v1_1_loop
_ = v1_1
var v2_2 int64 = v2_2_loop
_ = v2_2
var v3_3 *Constructor_T = v3_3_loop
_ = v3_3
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1583507464) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 3668501016) {
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3983586014 && __t_tag_5.UnsafePtr != nil) {
var __t6 gopurs_runtime.Value
{
var __t_tag_7 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 3668501016) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr), v2_2, v3_3})}.UnsafePtr)})}
goto end_branch_6
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 3983586014 && __t_tag_8.UnsafePtr != nil) {
var __t9 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V0
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 3668501016) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr), v2_2, v3_3})}.UnsafePtr)})}
goto end_branch_9
} else {

}
}
{
var __t_and_12 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr != nil) {

var __t_tag_11 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V0
__t_and_12 = (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 3668501016)
}
if __t_and_12 {
var __t13 gopurs_runtime.Value
{
var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 3983586014 && __t_tag_14.UnsafePtr != nil) {
var __t15 gopurs_runtime.Value
{
var __t_tag_16 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V0
if (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 3668501016) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V3)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_15
} else {

}
}
{
var __t_tag_17 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_19 bool = false
if (__t_tag_17.Type == 9 && __t_tag_17.IntVal == 3983586014 && __t_tag_17.UnsafePtr != nil) {

var __t_tag_18 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0
__t_and_19 = (__t_tag_18.Type == 9 && __t_tag_18.IntVal == 3668501016)
}
if __t_and_19 {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_15:
__t13 = __t15
goto end_branch_13
} else {

}
}
{
var __t_tag_20 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_22 bool = false
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 3983586014 && __t_tag_20.UnsafePtr != nil) {

var __t_tag_21 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0
__t_and_22 = (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 3668501016)
}
if __t_and_22 {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_13:
__t9 = __t13
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_9:
__t6 = __t9
goto end_branch_6
} else {

}
}
{
var __t_and_24 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr != nil) {

var __t_tag_23 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V0
__t_and_24 = (__t_tag_23.Type == 9 && __t_tag_23.IntVal == 3668501016)
}
if __t_and_24 {
var __t25 gopurs_runtime.Value
{
var __t_tag_26 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}
if (__t_tag_26.Type == 9 && __t_tag_26.IntVal == 3983586014 && __t_tag_26.UnsafePtr != nil) {
var __t27 gopurs_runtime.Value
{
var __t_tag_28 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V0
if (__t_tag_28.Type == 9 && __t_tag_28.IntVal == 3668501016) {
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V3)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_27
} else {

}
}
{
var __t_tag_29 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_31 bool = false
if (__t_tag_29.Type == 9 && __t_tag_29.IntVal == 3983586014 && __t_tag_29.UnsafePtr != nil) {

var __t_tag_30 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0
__t_and_31 = (__t_tag_30.Type == 9 && __t_tag_30.IntVal == 3668501016)
}
if __t_and_31 {
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_27
} else {

}
}
{
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_27:
__t25 = __t27
goto end_branch_25
} else {

}
}
{
var __t_tag_32 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_34 bool = false
if (__t_tag_32.Type == 9 && __t_tag_32.IntVal == 3983586014 && __t_tag_32.UnsafePtr != nil) {

var __t_tag_33 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0
__t_and_34 = (__t_tag_33.Type == 9 && __t_tag_33.IntVal == 3668501016)
}
if __t_and_34 {
__t25 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_25:
__t6 = __t25
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_6:
__t4 = __t6
goto end_branch_4
} else {

}
}
{
var __t_tag_35 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}
if (__t_tag_35.Type == 9 && __t_tag_35.IntVal == 3983586014 && __t_tag_35.UnsafePtr != nil) {
var __t36 gopurs_runtime.Value
{
var __t_tag_37 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V0
if (__t_tag_37.Type == 9 && __t_tag_37.IntVal == 3668501016) {
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr), v2_2, v3_3})}.UnsafePtr)})}
goto end_branch_36
} else {

}
}
{
var __t_and_39 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr != nil) {

var __t_tag_38 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V0
__t_and_39 = (__t_tag_38.Type == 9 && __t_tag_38.IntVal == 3668501016)
}
if __t_and_39 {
var __t40 gopurs_runtime.Value
{
var __t_tag_41 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}
if (__t_tag_41.Type == 9 && __t_tag_41.IntVal == 3983586014 && __t_tag_41.UnsafePtr != nil) {
var __t42 gopurs_runtime.Value
{
var __t_tag_43 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V0
if (__t_tag_43.Type == 9 && __t_tag_43.IntVal == 3668501016) {
__t42 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V3)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_42
} else {

}
}
{
var __t_tag_44 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_46 bool = false
if (__t_tag_44.Type == 9 && __t_tag_44.IntVal == 3983586014 && __t_tag_44.UnsafePtr != nil) {

var __t_tag_45 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0
__t_and_46 = (__t_tag_45.Type == 9 && __t_tag_45.IntVal == 3668501016)
}
if __t_and_46 {
__t42 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_42
} else {

}
}
{
__t42 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_42:
__t40 = __t42
goto end_branch_40
} else {

}
}
{
var __t_tag_47 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_49 bool = false
if (__t_tag_47.Type == 9 && __t_tag_47.IntVal == 3983586014 && __t_tag_47.UnsafePtr != nil) {

var __t_tag_48 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0
__t_and_49 = (__t_tag_48.Type == 9 && __t_tag_48.IntVal == 3668501016)
}
if __t_and_49 {
__t40 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_40
} else {

}
}
{
__t40 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_40:
__t36 = __t40
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_36:
__t4 = __t36
goto end_branch_4
} else {

}
}
{
var __t_and_51 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr != nil) {

var __t_tag_50 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V0
__t_and_51 = (__t_tag_50.Type == 9 && __t_tag_50.IntVal == 3668501016)
}
if __t_and_51 {
var __t52 gopurs_runtime.Value
{
var __t_tag_53 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}
if (__t_tag_53.Type == 9 && __t_tag_53.IntVal == 3983586014 && __t_tag_53.UnsafePtr != nil) {
var __t54 gopurs_runtime.Value
{
var __t_tag_55 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V0
if (__t_tag_55.Type == 9 && __t_tag_55.IntVal == 3668501016) {
__t54 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V3)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_54
} else {

}
}
{
var __t_tag_56 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_58 bool = false
if (__t_tag_56.Type == 9 && __t_tag_56.IntVal == 3983586014 && __t_tag_56.UnsafePtr != nil) {

var __t_tag_57 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0
__t_and_58 = (__t_tag_57.Type == 9 && __t_tag_57.IntVal == 3668501016)
}
if __t_and_58 {
__t54 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_54
} else {

}
}
{
__t54 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_54:
__t52 = __t54
goto end_branch_52
} else {

}
}
{
var __t_tag_59 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_61 bool = false
if (__t_tag_59.Type == 9 && __t_tag_59.IntVal == 3983586014 && __t_tag_59.UnsafePtr != nil) {

var __t_tag_60 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0
__t_and_61 = (__t_tag_60.Type == 9 && __t_tag_60.IntVal == 3668501016)
}
if __t_and_61 {
__t52 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_52
} else {

}
}
{
__t52 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_52:
__t4 = __t52
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_4:
__t2 = __t4
goto end_branch_2
} else {

}
}
{
var __t_and_63 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr != nil) {

var __t_tag_62 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V0
__t_and_63 = (__t_tag_62.Type == 9 && __t_tag_62.IntVal == 3668501016)
}
if __t_and_63 {
var __t64 gopurs_runtime.Value
{
var __t_tag_65 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}
if (__t_tag_65.Type == 9 && __t_tag_65.IntVal == 3983586014 && __t_tag_65.UnsafePtr != nil) {
var __t66 gopurs_runtime.Value
{
var __t_tag_67 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V0
if (__t_tag_67.Type == 9 && __t_tag_67.IntVal == 3668501016) {
__t66 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V3)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_66
} else {

}
}
{
var __t_tag_68 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_70 bool = false
if (__t_tag_68.Type == 9 && __t_tag_68.IntVal == 3983586014 && __t_tag_68.UnsafePtr != nil) {

var __t_tag_69 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0
__t_and_70 = (__t_tag_69.Type == 9 && __t_tag_69.IntVal == 3668501016)
}
if __t_and_70 {
__t66 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_66
} else {

}
}
{
__t66 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_66:
__t64 = __t66
goto end_branch_64
} else {

}
}
{
var __t_tag_71 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_73 bool = false
if (__t_tag_71.Type == 9 && __t_tag_71.IntVal == 3983586014 && __t_tag_71.UnsafePtr != nil) {

var __t_tag_72 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0
__t_and_73 = (__t_tag_72.Type == 9 && __t_tag_72.IntVal == 3668501016)
}
if __t_and_73 {
__t64 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_64
} else {

}
}
{
__t64 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_64:
__t2 = __t64
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
var __t_and_75 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr != nil) {

var __t_tag_74 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V0
__t_and_75 = (__t_tag_74.Type == 9 && __t_tag_74.IntVal == 3668501016)
}
if __t_and_75 {
var __t76 gopurs_runtime.Value
{
var __t_tag_77 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}
if (__t_tag_77.Type == 9 && __t_tag_77.IntVal == 3983586014 && __t_tag_77.UnsafePtr != nil) {
var __t78 gopurs_runtime.Value
{
var __t_tag_79 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V0
if (__t_tag_79.Type == 9 && __t_tag_79.IntVal == 3668501016) {
__t78 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr).V3)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_78
} else {

}
}
{
var __t_tag_80 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_82 bool = false
if (__t_tag_80.Type == 9 && __t_tag_80.IntVal == 3983586014 && __t_tag_80.UnsafePtr != nil) {

var __t_tag_81 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0
__t_and_82 = (__t_tag_81.Type == 9 && __t_tag_81.IntVal == 3668501016)
}
if __t_and_82 {
__t78 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_78
} else {

}
}
{
__t78 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_78:
__t76 = __t78
goto end_branch_76
} else {

}
}
{
var __t_tag_83 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}
var __t_and_85 bool = false
if (__t_tag_83.Type == 9 && __t_tag_83.IntVal == 3983586014 && __t_tag_83.UnsafePtr != nil) {

var __t_tag_84 gopurs_runtime.Value = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V0
__t_and_85 = (__t_tag_84.Type == 9 && __t_tag_84.IntVal == 3668501016)
}
if __t_and_85 {
__t76 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, v1_1, v2_2, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V1)}.UnsafePtr)})}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v3_3)}.UnsafePtr).V3)}.UnsafePtr).V3)}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_76
} else {

}
}
{
__t76 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_76:
__t1 = __t76
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, v_0, v1_1, v2_2, v3_3})}
}
end_branch_0:
return (*Constructor_T)(__t0.UnsafePtr)
}

func Call_ins(v_0_loop int64, v1_1_loop *Constructor_T) *Constructor_T {
ins:
for {
if false { continue ins }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *Constructor_T = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 3668501016, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: nil}.UnsafePtr), v_0, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: nil}.UnsafePtr)})}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 3983586014 && gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (v_0) < (gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2).IntVal) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_balance((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_ins(v_0, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}.UnsafePtr)))}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr)))}
goto end_branch_1
} else {

}
}
{
if (v_0) > (gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2).IntVal) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_balance((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_ins(v_0, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr)))}.UnsafePtr)))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V3)}.UnsafePtr)})}
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
return (*Constructor_T)(__t0.UnsafePtr)
}
}

func Call_insert(x_0_loop int64, s_1_loop *Constructor_T) *Constructor_T {
var x_0 int64 = x_0_loop
_ = x_0
var s_1 *Constructor_T = s_1_loop
_ = s_1
__local_var_2_0 := gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_ins(x_0, s_1))}
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3983586014 && __local_var_2_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(&Constructor_T{1, gopurs_runtime.Value{Type: 9, IntVal: 1583507464, UnsafePtr: nil}, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(__local_var_2_0.UnsafePtr).V1)}.UnsafePtr), gopurs_runtime.Int((*Constructor_T)(__local_var_2_0.UnsafePtr).V2).IntVal, (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer((*Constructor_T)(__local_var_2_0.UnsafePtr).V3)}.UnsafePtr)})}
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3983586014 && __local_var_2_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (*Constructor_T)(__t1.UnsafePtr)
}

func Call_buildTree(v_0_loop int64, v1_1_loop *Constructor_T) *Constructor_T {
buildTree:
for {
if false { continue buildTree }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *Constructor_T = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(v1_1)}
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (1)
v1_1_loop = (*Constructor_T)(gopurs_runtime.Value{Type: 9, IntVal: 3983586014, UnsafePtr: unsafe.Pointer(Call_insert(v_0, v1_1))}.UnsafePtr)
continue buildTree
__t0 = gopurs_runtime.Value{}
}
end_branch_0:
return (*Constructor_T)(__t0.UnsafePtr)
}
}



package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Test_ListOps_add gopurs_runtime.Value
var once_Test_ListOps_add sync.Once
func Get_Test_ListOps_add() gopurs_runtime.Value {
	once_Test_ListOps_add.Do(func() {
		cache_Test_ListOps_add = Get_Data_Semiring_intAdd()
	})
	return cache_Test_ListOps_add
}

var cache_Test_ListOps_Nil gopurs_runtime.Value
var once_Test_ListOps_Nil sync.Once
func Get_Test_ListOps_Nil() gopurs_runtime.Value {
	once_Test_ListOps_Nil.Do(func() {
		cache_Test_ListOps_Nil = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Test_ListOps_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Test_ListOps_Nil
}

var cache_Test_ListOps_Cons gopurs_runtime.Value
var once_Test_ListOps_Cons sync.Once
func Get_Test_ListOps_Cons() gopurs_runtime.Value {
	once_Test_ListOps_Cons.Do(func() {
		cache_Test_ListOps_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((&Constructor_Test_ListOps_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](value1)}))}
})
})
	})
	return cache_Test_ListOps_Cons
}

var cache_Test_ListOps_Cons__2036046248 gopurs_runtime.Value
var once_Test_ListOps_Cons__2036046248 sync.Once
func Get_Test_ListOps_Cons__2036046248() gopurs_runtime.Value {
	once_Test_ListOps_Cons__2036046248.Do(func() {
		cache_Test_ListOps_Cons__2036046248 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Rebox_Test_ListOps_2722731916_840821239(Call_Test_ListOps_Cons__2036046248(__eta_norm_1_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](__eta_norm_0_1_box))))}
})
	})
	return cache_Test_ListOps_Cons__2036046248
}

var cache_Test_ListOps_go__range gopurs_runtime.Value
var once_Test_ListOps_go__range sync.Once
func Get_Test_ListOps_go__range() gopurs_runtime.Value {
	once_Test_ListOps_go__range.Do(func() {
		cache_Test_ListOps_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Rebox_Test_ListOps_2722731916_840821239(Call_Test_ListOps_go__range(start_0_box.IntVal, end_1_box.IntVal)))}
})
	})
	return cache_Test_ListOps_go__range
}

var cache_Test_ListOps_foldl gopurs_runtime.Value
var once_Test_ListOps_foldl sync.Once
func Get_Test_ListOps_foldl() gopurs_runtime.Value {
	once_Test_ListOps_foldl.Do(func() {
		cache_Test_ListOps_foldl = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_ListOps_foldl(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_Test_ListOps_foldl
}

var cache_Test_ListOps_foldl__3114708910 gopurs_runtime.Value
var once_Test_ListOps_foldl__3114708910 sync.Once
func Get_Test_ListOps_foldl__3114708910() gopurs_runtime.Value {
	once_Test_ListOps_foldl__3114708910.Do(func() {
		cache_Test_ListOps_foldl__3114708910 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_ListOps_foldl__3114708910(v_0_box, v1_1_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](v2_2_box)))
})
	})
	return cache_Test_ListOps_foldl__3114708910
}

var cache_Test_ListOps_foldl__3865670492 gopurs_runtime.Value
var once_Test_ListOps_foldl__3865670492 sync.Once
func Get_Test_ListOps_foldl__3865670492() gopurs_runtime.Value {
	once_Test_ListOps_foldl__3865670492.Do(func() {
		cache_Test_ListOps_foldl__3865670492 = gopurs_runtime.Func3(func(v_unused_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_ListOps_foldl__3865670492(v_unused_0_box, v1_1_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](v2_2_box)))
})
	})
	return cache_Test_ListOps_foldl__3865670492
}

var cache_Test_ListOps_filterEvens gopurs_runtime.Value
var once_Test_ListOps_filterEvens sync.Once
func Get_Test_ListOps_filterEvens() gopurs_runtime.Value {
	once_Test_ListOps_filterEvens.Do(func() {
		cache_Test_ListOps_filterEvens = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Rebox_Test_ListOps_2722731916_840821239(Call_Test_ListOps_filterEvens(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](lst_0_box))))}
})
	})
	return cache_Test_ListOps_filterEvens
}

var cache_Test_ListOps_sumEvens gopurs_runtime.Value
var once_Test_ListOps_sumEvens sync.Once
func Get_Test_ListOps_sumEvens() gopurs_runtime.Value {
	once_Test_ListOps_sumEvens.Do(func() {
		cache_Test_ListOps_sumEvens = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_ListOps_sumEvens(n_0_box.IntVal))
})
	})
	return cache_Test_ListOps_sumEvens
}

var cache_Test_ListOps_describe gopurs_runtime.Value
var once_Test_ListOps_describe sync.Once
func Get_Test_ListOps_describe() gopurs_runtime.Value {
	once_Test_ListOps_describe.Do(func() {
		cache_Test_ListOps_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("List Processing (900 elements):"))
	})
	return cache_Test_ListOps_describe
}

var cache_Test_ListOps_act gopurs_runtime.Value
var once_Test_ListOps_act sync.Once
func Get_Test_ListOps_act() gopurs_runtime.Value {
	once_Test_ListOps_act.Do(func() {
		cache_Test_ListOps_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(int64(900)))
_ = __local_var_0_0
__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = __local_var_1_1
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_ListOps_sumEvens(__local_var_1_1.IntVal))).StrVal())
})
	})
	return cache_Test_ListOps_act
}

type Constructor_Test_ListOps_Nil[T_a any] struct {
	Rc uint32
}


type Constructor_Test_ListOps_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Test_ListOps_Cons[T_a]
}


func Call_Test_ListOps_Cons__2036046248(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop *Constructor_Test_ListOps_Cons[int64]) *Constructor_Test_ListOps_Cons[int64] {
Cons__2036046248:
for {
if false { continue Cons__2036046248 }
var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Test_ListOps_Cons[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return Rebox_Test_ListOps_840821239_2722731916(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((&Constructor_Test_ListOps_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](value1)}))}
})
}), gopurs_runtime.Int(__eta_norm_1_0), gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Rebox_Test_ListOps_2722731916_840821239(__eta_norm_0_1))})))
}
}

func Call_Test_ListOps_go__range(start_0_loop int64, end_1_loop int64) *Constructor_Test_ListOps_Cons[int64] {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var Call_local_Test_ListOps_go__1739844119_2_0_0 func(int64, *Constructor_Test_ListOps_Cons[int64]) *Constructor_Test_ListOps_Cons[int64]
_ = Call_local_Test_ListOps_go__1739844119_2_0_0
var go__1739844119_2_0_0 gopurs_runtime.Value
_ = go__1739844119_2_0_0
Call_local_Test_ListOps_go__1739844119_2_0_0 = func(curr_3_loop int64, acc_4_loop *Constructor_Test_ListOps_Cons[int64]) *Constructor_Test_ListOps_Cons[int64] {
go__1739844119_2_0_0:
for {
if false { continue go__1739844119_2_0_0 }
var curr_3 int64 = curr_3_loop
_ = curr_3
var acc_4 *Constructor_Test_ListOps_Cons[int64] = acc_4_loop
_ = acc_4
var __t1 *Constructor_Test_ListOps_Cons[int64]
{
if (curr_3) < (start_0) {
__t1 = acc_4
goto end_branch_1
} else {

}
}
{
curr_3_loop = (curr_3) - (int64(1))
acc_4_loop = (&Constructor_Test_ListOps_Cons[int64]{1, curr_3, acc_4})
continue go__1739844119_2_0_0
__t1 = func() *Constructor_Test_ListOps_Cons[int64] { panic("unreachable") }()
}
end_branch_1:
return __t1
}
}
go__1739844119_2_0_0 = gopurs_runtime.Func(func(curr_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Rebox_Test_ListOps_2722731916_840821239(Call_local_Test_ListOps_go__1739844119_2_0_0(curr_3_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](acc_4_loop_val))))}
})
})
var go__go_3_2_1 gopurs_runtime.Value
_ = go__go_3_2_1
// FALLBACK TCO: isLoop=false len=1
go__go_3_2_1 = gopurs_runtime.Func2(func(curr_4 gopurs_runtime.Value, acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Test_ListOps_Cons[int64]
{
if (curr_4.IntVal) < (start_0) {
__t3 = Rebox_Test_ListOps_840821239_2722731916(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](acc_5))
goto end_branch_3
} else {

}
}
{
__t3 = Call_local_Test_ListOps_go__1739844119_2_0_0((curr_4.IntVal) - (int64(1)), (&Constructor_Test_ListOps_Cons[int64]{1, curr_4.IntVal, Rebox_Test_ListOps_840821239_2722731916(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](acc_5))}))
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Rebox_Test_ListOps_2722731916_840821239(__t3))}
})
return Call_local_Test_ListOps_go__1739844119_2_0_0(end_1, (*Constructor_Test_ListOps_Cons[int64])(nil))
}

func Call_Test_ListOps_foldl(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Test_ListOps_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
foldl:
for {
if false { continue foldl }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Test_ListOps_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
v_0_loop = v_0
v1_1_loop = gopurs_runtime.Apply2(v_0, v1_1, (v2_2).V0)
v2_2_loop = (v2_2).V1
continue foldl
__t0 = func() gopurs_runtime.Value { panic("unreachable") }()
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

func Call_Test_ListOps_foldl__3114708910(v_0_loop gopurs_runtime.Value, v1_1_loop int64, v2_2_loop *Constructor_Test_ListOps_Cons[int64]) int64 {
foldl__3114708910:
for {
if false { continue foldl__3114708910 }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Test_ListOps_Cons[int64] = v2_2_loop
_ = v2_2
var __t0 int64
{
if (v2_2 == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
v_0_loop = v_0
v1_1_loop = gopurs_runtime.Apply2(v_0, gopurs_runtime.Int(v1_1), gopurs_runtime.Int((v2_2).V0)).IntVal
v2_2_loop = (v2_2).V1
continue foldl__3114708910
__t0 = func() int64 { panic("unreachable") }()
goto end_branch_0
} else {

}
}
{
__t0 = func() int64 { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Test_ListOps_foldl__3865670492(v_unused_0_loop gopurs_runtime.Value, v1_1_loop int64, v2_2_loop *Constructor_Test_ListOps_Cons[int64]) int64 {
foldl__3865670492:
for {
if false { continue foldl__3865670492 }
var v_unused_0 gopurs_runtime.Value = v_unused_0_loop
_ = v_unused_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Test_ListOps_Cons[int64] = v2_2_loop
_ = v2_2
var __t0 int64
{
if (v2_2 == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
v_unused_0_loop = Get_Data_Semiring_intAdd()
v1_1_loop = (v1_1) + ((v2_2).V0)
v2_2_loop = (v2_2).V1
continue foldl__3865670492
__t0 = func() int64 { panic("unreachable") }()
goto end_branch_0
} else {

}
}
{
__t0 = func() int64 { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Test_ListOps_filterEvens(lst_0_loop *Constructor_Test_ListOps_Cons[int64]) *Constructor_Test_ListOps_Cons[int64] {
var lst_0 *Constructor_Test_ListOps_Cons[int64] = lst_0_loop
_ = lst_0
var Call_local_Test_ListOps_go__717521980_1_0_2 func(*Constructor_Test_ListOps_Cons[int64], *Constructor_Test_ListOps_Cons[int64]) *Constructor_Test_ListOps_Cons[int64]
_ = Call_local_Test_ListOps_go__717521980_1_0_2
var go__717521980_1_0_2 gopurs_runtime.Value
_ = go__717521980_1_0_2
Call_local_Test_ListOps_go__717521980_1_0_2 = func(v_2_loop *Constructor_Test_ListOps_Cons[int64], v1_3_loop *Constructor_Test_ListOps_Cons[int64]) *Constructor_Test_ListOps_Cons[int64] {
go__717521980_1_0_2:
for {
if false { continue go__717521980_1_0_2 }
var v_2 *Constructor_Test_ListOps_Cons[int64] = v_2_loop
_ = v_2
var v1_3 *Constructor_Test_ListOps_Cons[int64] = v1_3_loop
_ = v1_3
var __t2 *Constructor_Test_ListOps_Cons[int64]
{
if (v_2 == nil) {
__t2 = v1_3
goto end_branch_2
} else {

}
}
{
if (v_2 != nil) {
var __t1 *Constructor_Test_ListOps_Cons[int64]
{
if (((v_2).V0) % (int64(2))) == (int64(0)) {
v_2_loop = (v_2).V1
v1_3_loop = (&Constructor_Test_ListOps_Cons[int64]{1, (v_2).V0, v1_3})
continue go__717521980_1_0_2
__t1 = func() *Constructor_Test_ListOps_Cons[int64] { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
v_2_loop = (v_2).V1
v1_3_loop = v1_3
continue go__717521980_1_0_2
__t1 = func() *Constructor_Test_ListOps_Cons[int64] { panic("unreachable") }()
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Test_ListOps_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__717521980_1_0_2 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Rebox_Test_ListOps_2722731916_840821239(Call_local_Test_ListOps_go__717521980_1_0_2(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](v1_3_loop_val))))}
})
})
var go__go_2_3_3 gopurs_runtime.Value
_ = go__go_2_3_3
// FALLBACK TCO: isLoop=false len=1
go__go_2_3_3 = gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Test_ListOps_Cons[int64]
{
var __t_tag_4 *Constructor_Test_ListOps_Cons[int64] = Rebox_Test_ListOps_840821239_2722731916(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](v_3))
if (__t_tag_4 == nil) {
__t7 = Rebox_Test_ListOps_840821239_2722731916(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](v1_4))
goto end_branch_7
} else {

}
}
{
var __t_tag_5 *Constructor_Test_ListOps_Cons[int64] = Rebox_Test_ListOps_840821239_2722731916(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](v_3))
if (__t_tag_5 != nil) {
var __t6 *Constructor_Test_ListOps_Cons[int64]
{
if (((*Constructor_Test_ListOps_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0.IntVal) % (int64(2))) == (int64(0)) {
__t6 = Call_local_Test_ListOps_go__717521980_1_0_2(Rebox_Test_ListOps_840821239_2722731916((*Constructor_Test_ListOps_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1), (&Constructor_Test_ListOps_Cons[int64]{1, (*Constructor_Test_ListOps_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0.IntVal, Rebox_Test_ListOps_840821239_2722731916(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](v1_4))}))
goto end_branch_6
} else {

}
}
{
__t6 = Call_local_Test_ListOps_go__717521980_1_0_2(Rebox_Test_ListOps_840821239_2722731916((*Constructor_Test_ListOps_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1), Rebox_Test_ListOps_840821239_2722731916(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](v1_4)))
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Test_ListOps_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Rebox_Test_ListOps_2722731916_840821239(__t7))}
})
return Call_local_Test_ListOps_go__717521980_1_0_2(lst_0, (*Constructor_Test_ListOps_Cons[int64])(nil))
}

func Call_Test_ListOps_sumEvens(n_0_loop int64) int64 {
var n_0 int64 = n_0_loop
_ = n_0
var Call_local_Test_ListOps_go__1739844119_1_0_4 func(int64, *Constructor_Test_ListOps_Cons[int64]) *Constructor_Test_ListOps_Cons[int64]
_ = Call_local_Test_ListOps_go__1739844119_1_0_4
var go__1739844119_1_0_4 gopurs_runtime.Value
_ = go__1739844119_1_0_4
Call_local_Test_ListOps_go__1739844119_1_0_4 = func(curr_2_loop int64, acc_3_loop *Constructor_Test_ListOps_Cons[int64]) *Constructor_Test_ListOps_Cons[int64] {
go__1739844119_1_0_4:
for {
if false { continue go__1739844119_1_0_4 }
var curr_2 int64 = curr_2_loop
_ = curr_2
var acc_3 *Constructor_Test_ListOps_Cons[int64] = acc_3_loop
_ = acc_3
var __t1 *Constructor_Test_ListOps_Cons[int64]
{
if (curr_2) < (int64(1)) {
__t1 = acc_3
goto end_branch_1
} else {

}
}
{
curr_2_loop = (curr_2) - (int64(1))
acc_3_loop = (&Constructor_Test_ListOps_Cons[int64]{1, curr_2, acc_3})
continue go__1739844119_1_0_4
__t1 = func() *Constructor_Test_ListOps_Cons[int64] { panic("unreachable") }()
}
end_branch_1:
return __t1
}
}
go__1739844119_1_0_4 = gopurs_runtime.Func(func(curr_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Rebox_Test_ListOps_2722731916_840821239(Call_local_Test_ListOps_go__1739844119_1_0_4(curr_2_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[int64]](acc_3_loop_val))))}
})
})
var go__go_2_2_5 gopurs_runtime.Value
_ = go__go_2_2_5
// FALLBACK TCO: isLoop=false len=1
go__go_2_2_5 = gopurs_runtime.Func2(func(curr_3 gopurs_runtime.Value, acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Test_ListOps_Cons[int64]
{
if (curr_3.IntVal) < (int64(1)) {
__t3 = Rebox_Test_ListOps_840821239_2722731916(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](acc_4))
goto end_branch_3
} else {

}
}
{
__t3 = Call_local_Test_ListOps_go__1739844119_1_0_4((curr_3.IntVal) - (int64(1)), (&Constructor_Test_ListOps_Cons[int64]{1, curr_3.IntVal, Rebox_Test_ListOps_840821239_2722731916(gopurs_runtime.CoerceToStruct[Constructor_Test_ListOps_Cons[gopurs_runtime.Value]](acc_4))}))
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Rebox_Test_ListOps_2722731916_840821239(__t3))}
})
return Call_Test_ListOps_foldl__3865670492(Get_Data_Semiring_intAdd(), int64(0), Call_Test_ListOps_filterEvens(Call_local_Test_ListOps_go__1739844119_1_0_4(n_0, (*Constructor_Test_ListOps_Cons[int64])(nil))))
}

func Rebox_Test_ListOps_2722731916_840821239(in *Constructor_Test_ListOps_Cons[int64]) *Constructor_Test_ListOps_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Test_ListOps_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = Rebox_Test_ListOps_2722731916_840821239(in.V1)
	return out
}

func Rebox_Test_ListOps_840821239_2722731916(in *Constructor_Test_ListOps_Cons[gopurs_runtime.Value]) *Constructor_Test_ListOps_Cons[int64] {
	if in == nil { return nil }
	out := &Constructor_Test_ListOps_Cons[int64]{}
		out.V0 = in.V0.IntVal
		out.V1 = Rebox_Test_ListOps_840821239_2722731916(in.V1)
	return out
}



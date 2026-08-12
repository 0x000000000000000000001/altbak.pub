package Test_ListOps

import (
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		cache_Nil = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: nil}
	})
	return cache_Nil
}

var cache_Cons gopurs_runtime.Value
var once_Cons sync.Once
func Get_Cons() gopurs_runtime.Value {
	once_Cons.Do(func() {
		cache_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](value1)})}
})
})
	})
	return cache_Cons
}

var cache_go__range gopurs_runtime.Value
var once_go__range sync.Once
func Get_go__range() gopurs_runtime.Value {
	once_go__range.Do(func() {
		cache_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Call_go__range(start_0_box.IntVal, end_1_box.IntVal))}
})
	})
	return cache_go__range
}

var cache_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_ADT_Test_ListOps_List_Any_Any gopurs_runtime.Value
var once_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_ADT_Test_ListOps_List_Any_Any sync.Once
func Get_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_ADT_Test_ListOps_List_Any_Any() gopurs_runtime.Value {
	once_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_ADT_Test_ListOps_List_Any_Any.Do(func() {
		cache_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_ADT_Test_ListOps_List_Any_Any = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_ADT_Test_ListOps_List_Any_Any(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_ADT_Test_ListOps_List_Any_Any
}

var cache_foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		cache_foldl = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_foldl
}

var cache_foldl__gopurs_runtime_Value_1255354935 gopurs_runtime.Value
var once_foldl__gopurs_runtime_Value_1255354935 sync.Once
func Get_foldl__gopurs_runtime_Value_1255354935() gopurs_runtime.Value {
	once_foldl__gopurs_runtime_Value_1255354935.Do(func() {
		cache_foldl__gopurs_runtime_Value_1255354935 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__gopurs_runtime_Value_1255354935(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_foldl__gopurs_runtime_Value_1255354935
}

var cache_filterEvens gopurs_runtime.Value
var once_filterEvens sync.Once
func Get_filterEvens() gopurs_runtime.Value {
	once_filterEvens.Do(func() {
		cache_filterEvens = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Call_filterEvens(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](lst_0_box)))}
})
	})
	return cache_filterEvens
}

var cache_sumEvens gopurs_runtime.Value
var once_sumEvens sync.Once
func Get_sumEvens() gopurs_runtime.Value {
	once_sumEvens.Do(func() {
		cache_sumEvens = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_sumEvens(n_0_box.IntVal))
})
	})
	return cache_sumEvens
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("List Processing (900 elements):"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(900)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_sumEvens(dummy_0.IntVal))))
}))
	})
	return cache_act
}

type Constructor_Nil[T_a any] struct {
	Rc uint32
}


type Constructor_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Cons[gopurs_runtime.Value]
}


func Call_go__range(start_0_loop int64, end_1_loop int64) *Constructor_Cons[int64] {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var go__go_2_0_0 gopurs_runtime.Value
go__go_2_0_0 = gopurs_runtime.Func(func(curr_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var curr_3_loop gopurs_runtime.Value = curr_3_loop_val
var acc_4_loop *Constructor_Cons[int64] = gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](acc_4_loop_val)
go__go_2_0_0:
for {
if false { continue go__go_2_0_0 }
var curr_3 gopurs_runtime.Value = curr_3_loop
_ = curr_3
var acc_4 *Constructor_Cons[int64] = acc_4_loop
_ = acc_4
var __t1 gopurs_runtime.Value
{
var __t2 gopurs_runtime.Value
{
if (curr_3.IntVal) < (start_0) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(false)
}
end_branch_2:
if (__t2.IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(acc_4)}
goto end_branch_1
} else {

}
}
{
curr_3_loop = gopurs_runtime.Int((curr_3.IntVal) - (1))
acc_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, curr_3, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(acc_4)})})})
continue go__go_2_0_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{}))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](__t1))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Apply2(go__go_2_0_0, gopurs_runtime.Int(end_1), gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: nil}))
}

func Call_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_ADT_Test_ListOps_List_Any_Any(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = Call_foldl(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)}))
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

func Call_foldl(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = Call_foldl(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)}))
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

func Call_foldl__gopurs_runtime_Value_1255354935(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = Call_foldl(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)}))
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

func Call_filterEvens(lst_0_loop *Constructor_Cons[int64]) *Constructor_Cons[int64] {
var lst_0 *Constructor_Cons[int64] = lst_0_loop
_ = lst_0
var go__go_1_0_1 gopurs_runtime.Value
go__go_1_0_1 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *Constructor_Cons[int64] = gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](v1_3_loop_val)
go__go_1_0_1:
for {
if false { continue go__go_1_0_1 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *Constructor_Cons[int64] = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 1127792131 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v1_3)}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1127792131 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Int(2)).IntVal) == (0) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v1_3)})})})
continue go__go_1_0_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
v1_3_loop = v1_3
continue go__go_1_0_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{}))}
}
end_branch_2:
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](__t2))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](__t1))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Apply2(go__go_1_0_1, gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: nil}))
}

func Call_sumEvens(n_0_loop int64) int64 {
var n_0 int64 = n_0_loop
_ = n_0
var go__go_1_0_2 gopurs_runtime.Value
go__go_1_0_2 = gopurs_runtime.Func(func(curr_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var curr_2_loop gopurs_runtime.Value = curr_2_loop_val
var acc_3_loop *Constructor_Cons[int64] = gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](acc_3_loop_val)
go__go_1_0_2:
for {
if false { continue go__go_1_0_2 }
var curr_2 gopurs_runtime.Value = curr_2_loop
_ = curr_2
var acc_3 *Constructor_Cons[int64] = acc_3_loop
_ = acc_3
var __t1 gopurs_runtime.Value
{
var __t2 gopurs_runtime.Value
{
if (curr_2.IntVal) < (1) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(false)
}
end_branch_2:
if (__t2.IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(acc_3)}
goto end_branch_1
} else {

}
}
{
curr_2_loop = gopurs_runtime.Int((curr_2.IntVal) - (1))
acc_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, curr_2, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(acc_3)})})})
continue go__go_1_0_2
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Value{}))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](__t1))}
}
}()
})
})
return Call_foldl(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Call_filterEvens(gopurs_runtime.CoerceToStruct[Constructor_Cons[int64]](gopurs_runtime.Apply2(go__go_1_0_2, gopurs_runtime.Int(n_0), gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: nil}))))})).IntVal
}



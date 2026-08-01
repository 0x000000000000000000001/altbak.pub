package Test_ListOps

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	unsafe "unsafe"
)

var cache_lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		cache_lessThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420))
})
}()
	})
	return cache_lessThan
}

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
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, value0, (*Constructor_Cons[gopurs_runtime.Value])(value1.UnsafePtr)})}
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

var cache_foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		cache_foldl = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl(v_0_box, v1_1_box, (*Constructor_Cons[gopurs_runtime.Value])(v2_2_box.UnsafePtr))
})
	})
	return cache_foldl
}

var cache_foldl__gopurs_runtime_Value_469417201 gopurs_runtime.Value
var once_foldl__gopurs_runtime_Value_469417201 sync.Once
func Get_foldl__gopurs_runtime_Value_469417201() gopurs_runtime.Value {
	once_foldl__gopurs_runtime_Value_469417201.Do(func() {
		cache_foldl__gopurs_runtime_Value_469417201 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_foldl__gopurs_runtime_Value_469417201(v_0_box, v1_1_box.IntVal, (*Constructor_Cons[gopurs_runtime.Value])(v2_2_box.UnsafePtr)))
})
	})
	return cache_foldl__gopurs_runtime_Value_469417201
}

var cache_foldl__gopurs_runtime_Value_2606687793 gopurs_runtime.Value
var once_foldl__gopurs_runtime_Value_2606687793 sync.Once
func Get_foldl__gopurs_runtime_Value_2606687793() gopurs_runtime.Value {
	once_foldl__gopurs_runtime_Value_2606687793.Do(func() {
		cache_foldl__gopurs_runtime_Value_2606687793 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__gopurs_runtime_Value_2606687793(v_0_box, v1_1_box, (*Constructor_Cons[gopurs_runtime.Value])(v2_2_box.UnsafePtr))
})
	})
	return cache_foldl__gopurs_runtime_Value_2606687793
}

var cache_filterEvens gopurs_runtime.Value
var once_filterEvens sync.Once
func Get_filterEvens() gopurs_runtime.Value {
	once_filterEvens.Do(func() {
		cache_filterEvens = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Call_filterEvens((*Constructor_Cons[int64])(lst_0_box.UnsafePtr)))}
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
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), Call_foldl(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(0), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Call_filterEvens((*Constructor_Cons[int64])(gopurs_runtime.Apply2(Get_go__range(), gopurs_runtime.Int(1), dummy_0).UnsafePtr)))}.UnsafePtr))))
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
var go__go_2_0 gopurs_runtime.Value
go__go_2_0 = gopurs_runtime.Func(func(curr_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var curr_3_loop gopurs_runtime.Value = curr_3_loop_val
var acc_4_loop gopurs_runtime.Value = acc_4_loop_val
go__go_2_0:
for {
if false { continue go__go_2_0 }
var curr_3 gopurs_runtime.Value = curr_3_loop
_ = curr_3
var acc_4 gopurs_runtime.Value = acc_4_loop
_ = acc_4
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), curr_3, gopurs_runtime.Int(start_0)).IntVal) != (0) {
__t1 = acc_4
goto end_branch_1
} else {

}
}
{
curr_3_loop = gopurs_runtime.Int((curr_3.IntVal) - (1))
acc_4_loop = func() gopurs_runtime.Value {
if ((go__go_2_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(go__go_2_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V0 = curr_3
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V1 = (*Constructor_Cons[gopurs_runtime.Value])(acc_4.UnsafePtr)
return go__go_2_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, curr_3, (*Constructor_Cons[gopurs_runtime.Value])(acc_4.UnsafePtr)})}
}
}()
continue go__go_2_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return (*Constructor_Cons[int64])(gopurs_runtime.Apply2(go__go_2_0, gopurs_runtime.Int(end_1), gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: nil}).UnsafePtr)
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
__t0 = Call_foldl(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)}.UnsafePtr))
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

func Call_foldl__gopurs_runtime_Value_469417201(v_0_loop gopurs_runtime.Value, v1_1_loop int64, v2_2_loop *Constructor_Cons[gopurs_runtime.Value]) int64 {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(v1_1)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = Call_foldl(v_0, gopurs_runtime.Apply2(v_0, gopurs_runtime.Int(v1_1), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)}.UnsafePtr))
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

func Call_foldl__gopurs_runtime_Value_2606687793(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
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
__t0 = Call_foldl(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)}.UnsafePtr))
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
var go__go_1_0 gopurs_runtime.Value
go__go_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0:
for {
if false { continue go__go_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 1127792131 && v_2.UnsafePtr == nil) {
__t1 = v1_3
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
v1_3_loop = func() gopurs_runtime.Value {
if ((go__go_1_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(go__go_1_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Cons[gopurs_runtime.Value])(go__go_1_0.UnsafePtr).V0 = (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0
(*Constructor_Cons[gopurs_runtime.Value])(go__go_1_0.UnsafePtr).V1 = (*Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr)
return go__go_1_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr)})}
}
}()
continue go__go_1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
v1_3_loop = v1_3
continue go__go_1_0
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
__t1 = __t2
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
}()
})
})
return (*Constructor_Cons[int64])(gopurs_runtime.Apply2(go__go_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: nil}).UnsafePtr)
}

func Call_sumEvens(n_0_loop int64) int64 {
var n_0 int64 = n_0_loop
_ = n_0
return Call_foldl(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(0), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(Call_filterEvens((*Constructor_Cons[int64])(gopurs_runtime.Apply2(Get_go__range(), gopurs_runtime.Int(1), gopurs_runtime.Int(n_0)).UnsafePtr)))}.UnsafePtr)).IntVal
}



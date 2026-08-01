package Test_Primes

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Eq "gopurs/output/Data.Eq"
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
		cache_Nil = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: nil}
	})
	return cache_Nil
}

var cache_Cons gopurs_runtime.Value
var once_Cons sync.Once
func Get_Cons() gopurs_runtime.Value {
	once_Cons.Do(func() {
		cache_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, value0, (*Constructor_Cons[gopurs_runtime.Value])(value1.UnsafePtr)})}
})
})
	})
	return cache_Cons
}

var cache_sumList gopurs_runtime.Value
var once_sumList sync.Once
func Get_sumList() gopurs_runtime.Value {
	once_sumList.Do(func() {
		cache_sumList = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_sumList((*Constructor_Cons[int64])(lst_0_box.UnsafePtr)))
})
	})
	return cache_sumList
}

var cache_reverse gopurs_runtime.Value
var once_reverse sync.Once
func Get_reverse() gopurs_runtime.Value {
	once_reverse.Do(func() {
		cache_reverse = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_reverse((*Constructor_Cons[gopurs_runtime.Value])(lst_0_box.UnsafePtr)))}
})
	})
	return cache_reverse
}

var cache_go__range gopurs_runtime.Value
var once_go__range sync.Once
func Get_go__range() gopurs_runtime.Value {
	once_go__range.Do(func() {
		cache_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_go__range(start_0_box.IntVal, end_1_box.IntVal))}
})
	})
	return cache_go__range
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, lst_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_filter(p_0_box, (*Constructor_Cons[gopurs_runtime.Value])(lst_1_box.UnsafePtr)))}
})
	})
	return cache_filter
}

var cache_filter__gopurs_runtime_Value_4130189293 gopurs_runtime.Value
var once_filter__gopurs_runtime_Value_4130189293 sync.Once
func Get_filter__gopurs_runtime_Value_4130189293() gopurs_runtime.Value {
	once_filter__gopurs_runtime_Value_4130189293.Do(func() {
		cache_filter__gopurs_runtime_Value_4130189293 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, lst_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_filter__gopurs_runtime_Value_4130189293(p_0_box, (*Constructor_Cons[gopurs_runtime.Value])(lst_1_box.UnsafePtr)))}
})
	})
	return cache_filter__gopurs_runtime_Value_4130189293
}

var cache_filter__gopurs_runtime_Value_4150915752 gopurs_runtime.Value
var once_filter__gopurs_runtime_Value_4150915752 sync.Once
func Get_filter__gopurs_runtime_Value_4150915752() gopurs_runtime.Value {
	once_filter__gopurs_runtime_Value_4150915752.Do(func() {
		cache_filter__gopurs_runtime_Value_4150915752 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, lst_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_filter__gopurs_runtime_Value_4150915752(p_0_box, (*Constructor_Cons[gopurs_runtime.Value])(lst_1_box.UnsafePtr)))}
})
	})
	return cache_filter__gopurs_runtime_Value_4150915752
}

var cache_sieve gopurs_runtime.Value
var once_sieve sync.Once
func Get_sieve() gopurs_runtime.Value {
	once_sieve.Do(func() {
		cache_sieve = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_sieve((*Constructor_Cons[int64])(v_0_box.UnsafePtr)))}
})
	})
	return cache_sieve
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Prime Sieve (sum primes up to 500):"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(500)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr == nil) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
v1_3_loop = gopurs_runtime.Int((v1_3.IntVal) + ((*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0.IntVal))
continue go__go_1_0
__t1 = gopurs_runtime.Value{}
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
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Apply2(go__go_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_sieve((*Constructor_Cons[int64])(gopurs_runtime.Apply2(Get_go__range(), gopurs_runtime.Int(2), dummy_0).UnsafePtr)))}, gopurs_runtime.Int(0))))
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


func Call_sumList(lst_0_loop *Constructor_Cons[int64]) int64 {
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
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr == nil) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
v1_3_loop = gopurs_runtime.Int((v1_3.IntVal) + ((*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0.IntVal))
continue go__go_1_0
__t1 = gopurs_runtime.Value{}
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
return gopurs_runtime.Apply2(go__go_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Int(0)).IntVal
}

func Call_reverse(lst_0_loop *Constructor_Cons[gopurs_runtime.Value]) *Constructor_Cons[gopurs_runtime.Value] {
var lst_0 *Constructor_Cons[gopurs_runtime.Value] = lst_0_loop
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
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr == nil) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
v1_3_loop = func() gopurs_runtime.Value {
if ((go__go_1_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(go__go_1_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Cons[gopurs_runtime.Value])(go__go_1_0.UnsafePtr).V0 = (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0
(*Constructor_Cons[gopurs_runtime.Value])(go__go_1_0.UnsafePtr).V1 = (*Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr)
return go__go_1_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr)})}
}
}()
continue go__go_1_0
__t1 = gopurs_runtime.Value{}
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
return (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply2(go__go_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: nil}).UnsafePtr)
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
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, curr_3, (*Constructor_Cons[gopurs_runtime.Value])(acc_4.UnsafePtr)})}
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
return (*Constructor_Cons[int64])(gopurs_runtime.Apply2(go__go_2_0, gopurs_runtime.Int(end_1), gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: nil}).UnsafePtr)
}

func Call_filter(p_0_loop gopurs_runtime.Value, lst_1_loop *Constructor_Cons[gopurs_runtime.Value]) *Constructor_Cons[gopurs_runtime.Value] {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var lst_1 *Constructor_Cons[gopurs_runtime.Value] = lst_1_loop
_ = lst_1
var go__go_2_0 gopurs_runtime.Value
go__go_2_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_0:
for {
if false { continue go__go_2_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr == nil) {
var go__go_5_2 gopurs_runtime.Value
go__go_5_2 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_2:
for {
if false { continue go__go_5_2 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t3 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr == nil) {
__t3 = v1_7
goto end_branch_3
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V1)}
v1_7_loop = func() gopurs_runtime.Value {
if ((go__go_2_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(go__go_2_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V0 = (*Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V0
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V1 = (*Constructor_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr)
return go__go_2_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr)})}
}
}()
continue go__go_5_2
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}()
})
})
__t1 = gopurs_runtime.Apply2(go__go_5_2, v1_4, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: nil})
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr != nil) {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
v1_4_loop = func() gopurs_runtime.Value {
if ((go__go_2_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(go__go_2_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V0 = (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V1 = (*Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr)
return go__go_2_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr)})}
}
}()
continue go__go_2_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
v1_4_loop = v1_4
continue go__go_2_0
__t4 = gopurs_runtime.Value{}
}
end_branch_4:
__t1 = __t4
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
return (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply2(go__go_2_0, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: nil}).UnsafePtr)
}

func Call_filter__gopurs_runtime_Value_4130189293(p_0_loop gopurs_runtime.Value, lst_1_loop *Constructor_Cons[gopurs_runtime.Value]) *Constructor_Cons[gopurs_runtime.Value] {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var lst_1 *Constructor_Cons[gopurs_runtime.Value] = lst_1_loop
_ = lst_1
var go__go_2_0 gopurs_runtime.Value
go__go_2_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_0:
for {
if false { continue go__go_2_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr == nil) {
var go__go_5_2 gopurs_runtime.Value
go__go_5_2 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_2:
for {
if false { continue go__go_5_2 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t3 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr == nil) {
__t3 = v1_7
goto end_branch_3
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V1)}
v1_7_loop = func() gopurs_runtime.Value {
if ((go__go_2_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(go__go_2_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V0 = (*Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V0
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V1 = (*Constructor_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr)
return go__go_2_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr)})}
}
}()
continue go__go_5_2
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}()
})
})
__t1 = gopurs_runtime.Apply2(go__go_5_2, v1_4, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: nil})
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr != nil) {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
v1_4_loop = func() gopurs_runtime.Value {
if ((go__go_2_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(go__go_2_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V0 = (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V1 = (*Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr)
return go__go_2_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr)})}
}
}()
continue go__go_2_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
v1_4_loop = v1_4
continue go__go_2_0
__t4 = gopurs_runtime.Value{}
}
end_branch_4:
__t1 = __t4
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
return (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply2(go__go_2_0, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: nil}).UnsafePtr)
}

func Call_filter__gopurs_runtime_Value_4150915752(p_0_loop gopurs_runtime.Value, lst_1_loop *Constructor_Cons[gopurs_runtime.Value]) *Constructor_Cons[gopurs_runtime.Value] {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var lst_1 *Constructor_Cons[gopurs_runtime.Value] = lst_1_loop
_ = lst_1
var go__go_2_0 gopurs_runtime.Value
go__go_2_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_0:
for {
if false { continue go__go_2_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr == nil) {
var go__go_5_2 gopurs_runtime.Value
go__go_5_2 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_2:
for {
if false { continue go__go_5_2 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t3 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr == nil) {
__t3 = v1_7
goto end_branch_3
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V1)}
v1_7_loop = func() gopurs_runtime.Value {
if ((go__go_2_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(go__go_2_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V0 = (*Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V0
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V1 = (*Constructor_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr)
return go__go_2_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr)})}
}
}()
continue go__go_5_2
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}()
})
})
__t1 = gopurs_runtime.Apply2(go__go_5_2, v1_4, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: nil})
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr != nil) {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
v1_4_loop = func() gopurs_runtime.Value {
if ((go__go_2_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(go__go_2_0.UnsafePtr).Rc) == (1)) {
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V0 = (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0
(*Constructor_Cons[gopurs_runtime.Value])(go__go_2_0.UnsafePtr).V1 = (*Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr)
return go__go_2_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr)})}
}
}()
continue go__go_2_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
v1_4_loop = v1_4
continue go__go_2_0
__t4 = gopurs_runtime.Value{}
}
end_branch_4:
__t1 = __t4
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
return (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply2(go__go_2_0, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: nil}).UnsafePtr)
}

func Call_sieve(v_0_loop *Constructor_Cons[int64]) *Constructor_Cons[int64] {
sieve:
for {
if false { continue sieve }
var v_0 *Constructor_Cons[int64] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 2390177629 && gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 2390177629 && gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__local_var_1_1 := (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
_ = __local_var_1_1
var go__go_2_2 gopurs_runtime.Value
go__go_2_2 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_2:
for {
if false { continue go__go_2_2 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr == nil) {
var go__go_5_4 gopurs_runtime.Value
go__go_5_4 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_4:
for {
if false { continue go__go_5_4 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t5 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr == nil) {
__t5 = v1_7
goto end_branch_5
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V1)}
v1_7_loop = func() gopurs_runtime.Value {
if ((__local_var_1_1.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_1_1.UnsafePtr).Rc) == (1)) {
(*Constructor_Cons[gopurs_runtime.Value])(__local_var_1_1.UnsafePtr).V0 = (*Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V0
(*Constructor_Cons[gopurs_runtime.Value])(__local_var_1_1.UnsafePtr).V1 = (*Constructor_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr)
return __local_var_1_1
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr)})}
}
}()
continue go__go_5_4
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
__t3 = gopurs_runtime.Apply2(go__go_5_4, v1_4, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: nil})
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr != nil) {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, __local_var_1_1).IntVal) == (0)), gopurs_runtime.Bool(false)).IntVal) != (0) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
v1_4_loop = func() gopurs_runtime.Value {
if ((__local_var_1_1.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_1_1.UnsafePtr).Rc) == (1)) {
(*Constructor_Cons[gopurs_runtime.Value])(__local_var_1_1.UnsafePtr).V0 = (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0
(*Constructor_Cons[gopurs_runtime.Value])(__local_var_1_1.UnsafePtr).V1 = (*Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr)
return __local_var_1_1
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr)})}
}
}()
continue go__go_2_2
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
v1_4_loop = v1_4
continue go__go_2_2
__t6 = gopurs_runtime.Value{}
}
end_branch_6:
__t3 = __t6
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}()
})
})
__t0 = func() gopurs_runtime.Value {
if ((__local_var_1_1.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_1_1.UnsafePtr).Rc) == (1)) {
(*Constructor_Cons[gopurs_runtime.Value])(__local_var_1_1.UnsafePtr).V0 = __local_var_1_1
(*Constructor_Cons[gopurs_runtime.Value])(__local_var_1_1.UnsafePtr).V1 = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_sieve((*Constructor_Cons[int64])(gopurs_runtime.Apply2(go__go_2_2, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: nil}).UnsafePtr)))}.UnsafePtr)
return __local_var_1_1
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, __local_var_1_1, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_sieve((*Constructor_Cons[int64])(gopurs_runtime.Apply2(go__go_2_2, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: nil}).UnsafePtr)))}.UnsafePtr)})}
}
}()
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (*Constructor_Cons[int64])(__t0.UnsafePtr)
}
}



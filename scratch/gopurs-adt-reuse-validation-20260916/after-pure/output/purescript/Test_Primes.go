package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Test_Primes_Nil gopurs_runtime.Value
var once_Test_Primes_Nil sync.Once
func Get_Test_Primes_Nil() gopurs_runtime.Value {
	once_Test_Primes_Nil.Do(func() {
		cache_Test_Primes_Nil = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Test_Primes_Nil
}

var cache_Test_Primes_Cons gopurs_runtime.Value
var once_Test_Primes_Cons sync.Once
func Get_Test_Primes_Cons() gopurs_runtime.Value {
	once_Test_Primes_Cons.Do(func() {
		cache_Test_Primes_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((&Constructor_Test_Primes_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](value1)}))}
})
})
	})
	return cache_Test_Primes_Cons
}

var cache_Test_Primes_Cons__1140510884 gopurs_runtime.Value
var once_Test_Primes_Cons__1140510884 sync.Once
func Get_Test_Primes_Cons__1140510884() gopurs_runtime.Value {
	once_Test_Primes_Cons__1140510884.Do(func() {
		cache_Test_Primes_Cons__1140510884 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_Test_Primes_Cons__1140510884(__eta_norm_1_0_box.IntVal, Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](__eta_norm_0_1_box)))))}
})
	})
	return cache_Test_Primes_Cons__1140510884
}

var cache_Test_Primes_sumList gopurs_runtime.Value
var once_Test_Primes_sumList sync.Once
func Get_Test_Primes_sumList() gopurs_runtime.Value {
	once_Test_Primes_sumList.Do(func() {
		cache_Test_Primes_sumList = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_Primes_sumList(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](lst_0_box))))
})
	})
	return cache_Test_Primes_sumList
}

var cache_Test_Primes_reverse gopurs_runtime.Value
var once_Test_Primes_reverse sync.Once
func Get_Test_Primes_reverse() gopurs_runtime.Value {
	once_Test_Primes_reverse.Do(func() {
		cache_Test_Primes_reverse = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_Test_Primes_reverse(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](lst_0_box)))}
})
	})
	return cache_Test_Primes_reverse
}

var cache_Test_Primes_reverse__670871290 gopurs_runtime.Value
var once_Test_Primes_reverse__670871290 sync.Once
func Get_Test_Primes_reverse__670871290() gopurs_runtime.Value {
	once_Test_Primes_reverse__670871290.Do(func() {
		cache_Test_Primes_reverse__670871290 = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_Test_Primes_reverse__670871290(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](lst_0_box)))))}
})
	})
	return cache_Test_Primes_reverse__670871290
}

var cache_Test_Primes_go__range gopurs_runtime.Value
var once_Test_Primes_go__range sync.Once
func Get_Test_Primes_go__range() gopurs_runtime.Value {
	once_Test_Primes_go__range.Do(func() {
		cache_Test_Primes_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_Test_Primes_go__range(start_0_box.IntVal, end_1_box.IntVal)))}
})
	})
	return cache_Test_Primes_go__range
}

var cache_Test_Primes_filter gopurs_runtime.Value
var once_Test_Primes_filter sync.Once
func Get_Test_Primes_filter() gopurs_runtime.Value {
	once_Test_Primes_filter.Do(func() {
		cache_Test_Primes_filter = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, lst_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_Test_Primes_filter(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](lst_1_box)))}
})
	})
	return cache_Test_Primes_filter
}

var cache_Test_Primes_filter__394169086 gopurs_runtime.Value
var once_Test_Primes_filter__394169086 sync.Once
func Get_Test_Primes_filter__394169086() gopurs_runtime.Value {
	once_Test_Primes_filter__394169086.Do(func() {
		cache_Test_Primes_filter__394169086 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, lst_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_Test_Primes_filter__394169086(p_0_box, Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](lst_1_box)))))}
})
	})
	return cache_Test_Primes_filter__394169086
}

var cache_Test_Primes_sieve gopurs_runtime.Value
var once_Test_Primes_sieve sync.Once
func Get_Test_Primes_sieve() gopurs_runtime.Value {
	once_Test_Primes_sieve.Do(func() {
		cache_Test_Primes_sieve = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_Test_Primes_sieve(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_0_box)))))}
})
	})
	return cache_Test_Primes_sieve
}

var cache_Test_Primes_describe gopurs_runtime.Value
var once_Test_Primes_describe sync.Once
func Get_Test_Primes_describe() gopurs_runtime.Value {
	once_Test_Primes_describe.Do(func() {
		cache_Test_Primes_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Prime Sieve (sum primes up to 500):"))
	})
	return cache_Test_Primes_describe
}

var cache_Test_Primes_act gopurs_runtime.Value
var once_Test_Primes_act sync.Once
func Get_Test_Primes_act() gopurs_runtime.Value {
	once_Test_Primes_act.Do(func() {
		cache_Test_Primes_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(int64(500)))
_ = __local_var_0_0
__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = __local_var_1_1
var Call_local_Test_Primes_go__1466046018_2_2_15 func(*Constructor_Test_Primes_Cons[int64], int64) int64
_ = Call_local_Test_Primes_go__1466046018_2_2_15
var go__1466046018_2_2_15 gopurs_runtime.Value
_ = go__1466046018_2_2_15
var Call_local_Test_Primes_go__go_2_3_16 func(*Constructor_Test_Primes_Cons[int64], int64) int64
_ = Call_local_Test_Primes_go__go_2_3_16
var go__go_2_3_16 gopurs_runtime.Value
_ = go__go_2_3_16
Call_local_Test_Primes_go__1466046018_2_2_15 = func(v_3_loop *Constructor_Test_Primes_Cons[int64], v1_4_loop int64) int64 {
go__1466046018_2_2_15:
for {
if false { continue go__1466046018_2_2_15 }
var v_3 *Constructor_Test_Primes_Cons[int64] = v_3_loop
_ = v_3
var v1_4 int64 = v1_4_loop
_ = v1_4
var __t4 int64
{
if (v_3 == nil) {
__t4 = v1_4
goto end_branch_4
} else {

}
}
{
if (v_3 != nil) {
v_3_loop = (v_3).V1
v1_4_loop = (v1_4) + ((v_3).V0)
continue go__1466046018_2_2_15
__t4 = func() int64 { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
__t4 = func() int64 { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__1466046018_2_2_15 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Test_Primes_go__1466046018_2_2_15(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_3_loop_val)), v1_4_loop_val.IntVal))
})
})
Call_local_Test_Primes_go__go_2_3_16 = func(v_3_loop *Constructor_Test_Primes_Cons[int64], v1_4_loop int64) int64 {
go__go_2_3_16:
for {
if false { continue go__go_2_3_16 }
var v_3 *Constructor_Test_Primes_Cons[int64] = v_3_loop
_ = v_3
var v1_4 int64 = v1_4_loop
_ = v1_4
var __t5 int64
{
if (v_3 == nil) {
__t5 = v1_4
goto end_branch_5
} else {

}
}
{
if (v_3 != nil) {
__t5 = Call_local_Test_Primes_go__1466046018_2_2_15((v_3).V1, (v1_4) + ((v_3).V0))
goto end_branch_5
} else {

}
}
{
__t5 = func() int64 { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_2_3_16 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Test_Primes_go__go_2_3_16(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_3_loop_val)), v1_4_loop_val.IntVal))
})
})
var Call_local_Test_Primes_go__29212279_3_6_17 func(int64, *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__29212279_3_6_17
var go__29212279_3_6_17 gopurs_runtime.Value
_ = go__29212279_3_6_17
var Call_local_Test_Primes_go__go_3_7_18 func(int64, *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__go_3_7_18
var go__go_3_7_18 gopurs_runtime.Value
_ = go__go_3_7_18
Call_local_Test_Primes_go__29212279_3_6_17 = func(curr_4_loop int64, acc_5_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__29212279_3_6_17:
for {
if false { continue go__29212279_3_6_17 }
var curr_4 int64 = curr_4_loop
_ = curr_4
var acc_5 *Constructor_Test_Primes_Cons[int64] = acc_5_loop
_ = acc_5
var __t8 *Constructor_Test_Primes_Cons[int64]
{
if (curr_4) < (int64(2)) {
__t8 = acc_5
goto end_branch_8
} else {

}
}
{
curr_4_loop = (curr_4) - (int64(1))
acc_5_loop = (&Constructor_Test_Primes_Cons[int64]{1, curr_4, acc_5})
continue go__29212279_3_6_17
__t8 = func() *Constructor_Test_Primes_Cons[int64] { panic("unreachable") }()
}
end_branch_8:
return __t8
}
}
go__29212279_3_6_17 = gopurs_runtime.Func(func(curr_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__29212279_3_6_17(curr_4_loop_val.IntVal, Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](acc_5_loop_val)))))}
})
})
Call_local_Test_Primes_go__go_3_7_18 = func(curr_4_loop int64, acc_5_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__go_3_7_18:
for {
if false { continue go__go_3_7_18 }
var curr_4 int64 = curr_4_loop
_ = curr_4
var acc_5 *Constructor_Test_Primes_Cons[int64] = acc_5_loop
_ = acc_5
var __t9 *Constructor_Test_Primes_Cons[int64]
{
if (curr_4) < (int64(2)) {
__t9 = acc_5
goto end_branch_9
} else {

}
}
{
__t9 = Call_local_Test_Primes_go__29212279_3_6_17((curr_4) - (int64(1)), (&Constructor_Test_Primes_Cons[int64]{1, curr_4, acc_5}))
}
end_branch_9:
return __t9
}
}
go__go_3_7_18 = gopurs_runtime.Func(func(curr_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__go_3_7_18(curr_4_loop_val.IntVal, Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](acc_5_loop_val)))))}
})
})
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_local_Test_Primes_go__1466046018_2_2_15(Call_Test_Primes_sieve(Call_local_Test_Primes_go__29212279_3_6_17(__local_var_1_1.IntVal, (*Constructor_Test_Primes_Cons[int64])(nil))), int64(0)))).StrVal())
})
	})
	return cache_Test_Primes_act
}

type Constructor_Test_Primes_Nil[T_a any] struct {
	Rc uint32
}


type Constructor_Test_Primes_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Test_Primes_Cons[T_a]
}


func Call_Test_Primes_Cons__1140510884(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
Cons__1140510884:
for {
if false { continue Cons__1140510884 }
var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Test_Primes_Cons[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return (&Constructor_Test_Primes_Cons[int64]{1, __eta_norm_1_0, __eta_norm_0_1})
}
}

func Call_Test_Primes_sumList(lst_0_loop *Constructor_Test_Primes_Cons[int64]) int64 {
var lst_0 *Constructor_Test_Primes_Cons[int64] = lst_0_loop
_ = lst_0
var Call_local_Test_Primes_go__1466046018_1_0_0 func(*Constructor_Test_Primes_Cons[int64], int64) int64
_ = Call_local_Test_Primes_go__1466046018_1_0_0
var go__1466046018_1_0_0 gopurs_runtime.Value
_ = go__1466046018_1_0_0
var Call_local_Test_Primes_go__go_1_1_1 func(*Constructor_Test_Primes_Cons[int64], int64) int64
_ = Call_local_Test_Primes_go__go_1_1_1
var go__go_1_1_1 gopurs_runtime.Value
_ = go__go_1_1_1
Call_local_Test_Primes_go__1466046018_1_0_0 = func(v_2_loop *Constructor_Test_Primes_Cons[int64], v1_3_loop int64) int64 {
go__1466046018_1_0_0:
for {
if false { continue go__1466046018_1_0_0 }
var v_2 *Constructor_Test_Primes_Cons[int64] = v_2_loop
_ = v_2
var v1_3 int64 = v1_3_loop
_ = v1_3
var __t2 int64
{
if (v_2 == nil) {
__t2 = v1_3
goto end_branch_2
} else {

}
}
{
if (v_2 != nil) {
v_2_loop = (v_2).V1
v1_3_loop = (v1_3) + ((v_2).V0)
continue go__1466046018_1_0_0
__t2 = func() int64 { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
__t2 = func() int64 { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__1466046018_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Test_Primes_go__1466046018_1_0_0(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_2_loop_val)), v1_3_loop_val.IntVal))
})
})
Call_local_Test_Primes_go__go_1_1_1 = func(v_2_loop *Constructor_Test_Primes_Cons[int64], v1_3_loop int64) int64 {
go__go_1_1_1:
for {
if false { continue go__go_1_1_1 }
var v_2 *Constructor_Test_Primes_Cons[int64] = v_2_loop
_ = v_2
var v1_3 int64 = v1_3_loop
_ = v1_3
var __t3 int64
{
if (v_2 == nil) {
__t3 = v1_3
goto end_branch_3
} else {

}
}
{
if (v_2 != nil) {
__t3 = Call_local_Test_Primes_go__1466046018_1_0_0((v_2).V1, (v1_3) + ((v_2).V0))
goto end_branch_3
} else {

}
}
{
__t3 = func() int64 { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_1_1_1 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Test_Primes_go__go_1_1_1(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_2_loop_val)), v1_3_loop_val.IntVal))
})
})
return Call_local_Test_Primes_go__1466046018_1_0_0(lst_0, int64(0))
}

func Call_Test_Primes_reverse(lst_0_loop *Constructor_Test_Primes_Cons[gopurs_runtime.Value]) *Constructor_Test_Primes_Cons[gopurs_runtime.Value] {
var lst_0 *Constructor_Test_Primes_Cons[gopurs_runtime.Value] = lst_0_loop
_ = lst_0
var Call_local_Test_Primes_go__go_1_0_2 func(*Constructor_Test_Primes_Cons[gopurs_runtime.Value], *Constructor_Test_Primes_Cons[gopurs_runtime.Value]) *Constructor_Test_Primes_Cons[gopurs_runtime.Value]
_ = Call_local_Test_Primes_go__go_1_0_2
var go__go_1_0_2 gopurs_runtime.Value
_ = go__go_1_0_2
Call_local_Test_Primes_go__go_1_0_2 = func(v_2_loop *Constructor_Test_Primes_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Test_Primes_Cons[gopurs_runtime.Value]) *Constructor_Test_Primes_Cons[gopurs_runtime.Value] {
go__go_1_0_2:
for {
if false { continue go__go_1_0_2 }
var v_2 *Constructor_Test_Primes_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *Constructor_Test_Primes_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t1 *Constructor_Test_Primes_Cons[gopurs_runtime.Value]
{
if (v_2 == nil) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v_2 != nil) {
v_2_loop = (v_2).V1
v1_3_loop = (&Constructor_Test_Primes_Cons[gopurs_runtime.Value]{1, (v_2).V0, v1_3})
continue go__go_1_0_2
__t1 = func() *Constructor_Test_Primes_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Test_Primes_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}
go__go_1_0_2 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_local_Test_Primes_go__go_1_0_2(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
return Call_local_Test_Primes_go__go_1_0_2(lst_0, (*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(nil))
}

func Call_Test_Primes_reverse__670871290(lst_0_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
reverse__670871290:
for {
if false { continue reverse__670871290 }
var lst_0 *Constructor_Test_Primes_Cons[int64] = lst_0_loop
_ = lst_0
var Call_local_Test_Primes_go__302824162_1_0_3 func(*Constructor_Test_Primes_Cons[int64], *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__302824162_1_0_3
var go__302824162_1_0_3 gopurs_runtime.Value
_ = go__302824162_1_0_3
var Call_local_Test_Primes_go__go_1_1_4 func(*Constructor_Test_Primes_Cons[int64], *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__go_1_1_4
var go__go_1_1_4 gopurs_runtime.Value
_ = go__go_1_1_4
Call_local_Test_Primes_go__302824162_1_0_3 = func(v_2_loop *Constructor_Test_Primes_Cons[int64], v1_3_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__302824162_1_0_3:
for {
if false { continue go__302824162_1_0_3 }
var v_2 *Constructor_Test_Primes_Cons[int64] = v_2_loop
_ = v_2
var v1_3 *Constructor_Test_Primes_Cons[int64] = v1_3_loop
_ = v1_3
var __t2 *Constructor_Test_Primes_Cons[int64]
{
if (v_2 == nil) {
__t2 = v1_3
goto end_branch_2
} else {

}
}
{
if (v_2 != nil) {
v_2_loop = (v_2).V1
v1_3_loop = (&Constructor_Test_Primes_Cons[int64]{1, (v_2).V0, v1_3})
continue go__302824162_1_0_3
__t2 = func() *Constructor_Test_Primes_Cons[int64] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__302824162_1_0_3 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__302824162_1_0_3(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_2_loop_val)), Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_3_loop_val)))))}
})
})
Call_local_Test_Primes_go__go_1_1_4 = func(v_2_loop *Constructor_Test_Primes_Cons[int64], v1_3_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__go_1_1_4:
for {
if false { continue go__go_1_1_4 }
var v_2 *Constructor_Test_Primes_Cons[int64] = v_2_loop
_ = v_2
var v1_3 *Constructor_Test_Primes_Cons[int64] = v1_3_loop
_ = v1_3
var __t3 *Constructor_Test_Primes_Cons[int64]
{
if (v_2 == nil) {
__t3 = v1_3
goto end_branch_3
} else {

}
}
{
if (v_2 != nil) {
__t3 = Call_local_Test_Primes_go__302824162_1_0_3((v_2).V1, (&Constructor_Test_Primes_Cons[int64]{1, (v_2).V0, v1_3}))
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_1_1_4 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__go_1_1_4(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_2_loop_val)), Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_3_loop_val)))))}
})
})
return Call_local_Test_Primes_go__302824162_1_0_3(lst_0, (*Constructor_Test_Primes_Cons[int64])(nil))
}
}

func Call_Test_Primes_go__range(start_0_loop int64, end_1_loop int64) *Constructor_Test_Primes_Cons[int64] {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var Call_local_Test_Primes_go__29212279_2_0_5 func(int64, *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__29212279_2_0_5
var go__29212279_2_0_5 gopurs_runtime.Value
_ = go__29212279_2_0_5
var Call_local_Test_Primes_go__go_2_1_6 func(int64, *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__go_2_1_6
var go__go_2_1_6 gopurs_runtime.Value
_ = go__go_2_1_6
Call_local_Test_Primes_go__29212279_2_0_5 = func(curr_3_loop int64, acc_4_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__29212279_2_0_5:
for {
if false { continue go__29212279_2_0_5 }
var curr_3 int64 = curr_3_loop
_ = curr_3
var acc_4 *Constructor_Test_Primes_Cons[int64] = acc_4_loop
_ = acc_4
var __t2 *Constructor_Test_Primes_Cons[int64]
{
if (curr_3) < (start_0) {
__t2 = acc_4
goto end_branch_2
} else {

}
}
{
curr_3_loop = (curr_3) - (int64(1))
acc_4_loop = (&Constructor_Test_Primes_Cons[int64]{1, curr_3, acc_4})
continue go__29212279_2_0_5
__t2 = func() *Constructor_Test_Primes_Cons[int64] { panic("unreachable") }()
}
end_branch_2:
return __t2
}
}
go__29212279_2_0_5 = gopurs_runtime.Func(func(curr_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__29212279_2_0_5(curr_3_loop_val.IntVal, Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](acc_4_loop_val)))))}
})
})
Call_local_Test_Primes_go__go_2_1_6 = func(curr_3_loop int64, acc_4_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__go_2_1_6:
for {
if false { continue go__go_2_1_6 }
var curr_3 int64 = curr_3_loop
_ = curr_3
var acc_4 *Constructor_Test_Primes_Cons[int64] = acc_4_loop
_ = acc_4
var __t3 *Constructor_Test_Primes_Cons[int64]
{
if (curr_3) < (start_0) {
__t3 = acc_4
goto end_branch_3
} else {

}
}
{
__t3 = Call_local_Test_Primes_go__29212279_2_0_5((curr_3) - (int64(1)), (&Constructor_Test_Primes_Cons[int64]{1, curr_3, acc_4}))
}
end_branch_3:
return __t3
}
}
go__go_2_1_6 = gopurs_runtime.Func(func(curr_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__go_2_1_6(curr_3_loop_val.IntVal, Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](acc_4_loop_val)))))}
})
})
return Call_local_Test_Primes_go__29212279_2_0_5(end_1, (*Constructor_Test_Primes_Cons[int64])(nil))
}

func Call_Test_Primes_filter(p_0_loop gopurs_runtime.Value, lst_1_loop *Constructor_Test_Primes_Cons[gopurs_runtime.Value]) *Constructor_Test_Primes_Cons[gopurs_runtime.Value] {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var lst_1 *Constructor_Test_Primes_Cons[gopurs_runtime.Value] = lst_1_loop
_ = lst_1
var Call_local_Test_Primes_go__go_2_0_7 func(*Constructor_Test_Primes_Cons[gopurs_runtime.Value], *Constructor_Test_Primes_Cons[gopurs_runtime.Value]) *Constructor_Test_Primes_Cons[gopurs_runtime.Value]
_ = Call_local_Test_Primes_go__go_2_0_7
var go__go_2_0_7 gopurs_runtime.Value
_ = go__go_2_0_7
Call_local_Test_Primes_go__go_2_0_7 = func(v_3_loop *Constructor_Test_Primes_Cons[gopurs_runtime.Value], v1_4_loop *Constructor_Test_Primes_Cons[gopurs_runtime.Value]) *Constructor_Test_Primes_Cons[gopurs_runtime.Value] {
go__go_2_0_7:
for {
if false { continue go__go_2_0_7 }
var v_3 *Constructor_Test_Primes_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 *Constructor_Test_Primes_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t4 *Constructor_Test_Primes_Cons[gopurs_runtime.Value]
{
if (v_3 == nil) {
var Call_local_Test_Primes_go__go_5_1_8 func(*Constructor_Test_Primes_Cons[gopurs_runtime.Value], *Constructor_Test_Primes_Cons[gopurs_runtime.Value]) *Constructor_Test_Primes_Cons[gopurs_runtime.Value]
_ = Call_local_Test_Primes_go__go_5_1_8
var go__go_5_1_8 gopurs_runtime.Value
_ = go__go_5_1_8
Call_local_Test_Primes_go__go_5_1_8 = func(v_6_loop *Constructor_Test_Primes_Cons[gopurs_runtime.Value], v1_7_loop *Constructor_Test_Primes_Cons[gopurs_runtime.Value]) *Constructor_Test_Primes_Cons[gopurs_runtime.Value] {
go__go_5_1_8:
for {
if false { continue go__go_5_1_8 }
var v_6 *Constructor_Test_Primes_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var v1_7 *Constructor_Test_Primes_Cons[gopurs_runtime.Value] = v1_7_loop
_ = v1_7
var __t2 *Constructor_Test_Primes_Cons[gopurs_runtime.Value]
{
if (v_6 == nil) {
__t2 = v1_7
goto end_branch_2
} else {

}
}
{
if (v_6 != nil) {
v_6_loop = (v_6).V1
v1_7_loop = (&Constructor_Test_Primes_Cons[gopurs_runtime.Value]{1, (v_6).V0, v1_7})
continue go__go_5_1_8
__t2 = func() *Constructor_Test_Primes_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Test_Primes_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__go_5_1_8 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_local_Test_Primes_go__go_5_1_8(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_6_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_7_loop_val)))}
})
})
__t4 = Call_local_Test_Primes_go__go_5_1_8(v1_4, (*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(nil))
goto end_branch_4
} else {

}
}
{
if (v_3 != nil) {
var __t3 *Constructor_Test_Primes_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(p_0, (v_3).V0).IntVal) != (0) {
v_3_loop = (v_3).V1
v1_4_loop = (&Constructor_Test_Primes_Cons[gopurs_runtime.Value]{1, (v_3).V0, v1_4})
continue go__go_2_0_7
__t3 = func() *Constructor_Test_Primes_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
v_3_loop = (v_3).V1
v1_4_loop = v1_4
continue go__go_2_0_7
__t3 = func() *Constructor_Test_Primes_Cons[gopurs_runtime.Value] { panic("unreachable") }()
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Test_Primes_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_2_0_7 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_local_Test_Primes_go__go_2_0_7(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_4_loop_val)))}
})
})
return Call_local_Test_Primes_go__go_2_0_7(lst_1, (*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(nil))
}

func Call_Test_Primes_filter__394169086(p_0_loop gopurs_runtime.Value, lst_1_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
filter__394169086:
for {
if false { continue filter__394169086 }
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var lst_1 *Constructor_Test_Primes_Cons[int64] = lst_1_loop
_ = lst_1
var Call_local_Test_Primes_go__302824162_2_0_9 func(*Constructor_Test_Primes_Cons[int64], *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__302824162_2_0_9
var go__302824162_2_0_9 gopurs_runtime.Value
_ = go__302824162_2_0_9
var Call_local_Test_Primes_go__go_2_1_10 func(*Constructor_Test_Primes_Cons[int64], *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__go_2_1_10
var go__go_2_1_10 gopurs_runtime.Value
_ = go__go_2_1_10
Call_local_Test_Primes_go__302824162_2_0_9 = func(v_3_loop *Constructor_Test_Primes_Cons[int64], v1_4_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__302824162_2_0_9:
for {
if false { continue go__302824162_2_0_9 }
var v_3 *Constructor_Test_Primes_Cons[int64] = v_3_loop
_ = v_3
var v1_4 *Constructor_Test_Primes_Cons[int64] = v1_4_loop
_ = v1_4
var __t7 *Constructor_Test_Primes_Cons[int64]
{
if (v_3 == nil) {
var Call_local_Test_Primes_go__302824162_5_2_11 func(*Constructor_Test_Primes_Cons[int64], *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__302824162_5_2_11
var go__302824162_5_2_11 gopurs_runtime.Value
_ = go__302824162_5_2_11
var Call_local_Test_Primes_go__go_5_3_12 func(*Constructor_Test_Primes_Cons[int64], *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__go_5_3_12
var go__go_5_3_12 gopurs_runtime.Value
_ = go__go_5_3_12
Call_local_Test_Primes_go__302824162_5_2_11 = func(v_6_loop *Constructor_Test_Primes_Cons[int64], v1_7_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__302824162_5_2_11:
for {
if false { continue go__302824162_5_2_11 }
var v_6 *Constructor_Test_Primes_Cons[int64] = v_6_loop
_ = v_6
var v1_7 *Constructor_Test_Primes_Cons[int64] = v1_7_loop
_ = v1_7
var __t4 *Constructor_Test_Primes_Cons[int64]
{
if (v_6 == nil) {
__t4 = v1_7
goto end_branch_4
} else {

}
}
{
if (v_6 != nil) {
v_6_loop = (v_6).V1
v1_7_loop = (&Constructor_Test_Primes_Cons[int64]{1, (v_6).V0, v1_7})
continue go__302824162_5_2_11
__t4 = func() *Constructor_Test_Primes_Cons[int64] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__302824162_5_2_11 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__302824162_5_2_11(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_6_loop_val)), Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_7_loop_val)))))}
})
})
Call_local_Test_Primes_go__go_5_3_12 = func(v_6_loop *Constructor_Test_Primes_Cons[int64], v1_7_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__go_5_3_12:
for {
if false { continue go__go_5_3_12 }
var v_6 *Constructor_Test_Primes_Cons[int64] = v_6_loop
_ = v_6
var v1_7 *Constructor_Test_Primes_Cons[int64] = v1_7_loop
_ = v1_7
var __t5 *Constructor_Test_Primes_Cons[int64]
{
if (v_6 == nil) {
__t5 = v1_7
goto end_branch_5
} else {

}
}
{
if (v_6 != nil) {
__t5 = Call_local_Test_Primes_go__302824162_5_2_11((v_6).V1, (&Constructor_Test_Primes_Cons[int64]{1, (v_6).V0, v1_7}))
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_5_3_12 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__go_5_3_12(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_6_loop_val)), Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_7_loop_val)))))}
})
})
__t7 = Call_local_Test_Primes_go__302824162_5_2_11(v1_4, (*Constructor_Test_Primes_Cons[int64])(nil))
goto end_branch_7
} else {

}
}
{
if (v_3 != nil) {
var __t6 *Constructor_Test_Primes_Cons[int64]
{
if (gopurs_runtime.Apply(p_0, gopurs_runtime.Int((v_3).V0)).IntVal) != (0) {
v_3_loop = (v_3).V1
v1_4_loop = (&Constructor_Test_Primes_Cons[int64]{1, (v_3).V0, v1_4})
continue go__302824162_2_0_9
__t6 = func() *Constructor_Test_Primes_Cons[int64] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
v_3_loop = (v_3).V1
v1_4_loop = v1_4
continue go__302824162_2_0_9
__t6 = func() *Constructor_Test_Primes_Cons[int64] { panic("unreachable") }()
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}
go__302824162_2_0_9 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__302824162_2_0_9(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_3_loop_val)), Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_4_loop_val)))))}
})
})
Call_local_Test_Primes_go__go_2_1_10 = func(v_3_loop *Constructor_Test_Primes_Cons[int64], v1_4_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__go_2_1_10:
for {
if false { continue go__go_2_1_10 }
var v_3 *Constructor_Test_Primes_Cons[int64] = v_3_loop
_ = v_3
var v1_4 *Constructor_Test_Primes_Cons[int64] = v1_4_loop
_ = v1_4
var __t13 *Constructor_Test_Primes_Cons[int64]
{
if (v_3 == nil) {
var Call_local_Test_Primes_go__302824162_5_8_13 func(*Constructor_Test_Primes_Cons[int64], *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__302824162_5_8_13
var go__302824162_5_8_13 gopurs_runtime.Value
_ = go__302824162_5_8_13
var Call_local_Test_Primes_go__go_5_9_14 func(*Constructor_Test_Primes_Cons[int64], *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__go_5_9_14
var go__go_5_9_14 gopurs_runtime.Value
_ = go__go_5_9_14
Call_local_Test_Primes_go__302824162_5_8_13 = func(v_6_loop *Constructor_Test_Primes_Cons[int64], v1_7_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__302824162_5_8_13:
for {
if false { continue go__302824162_5_8_13 }
var v_6 *Constructor_Test_Primes_Cons[int64] = v_6_loop
_ = v_6
var v1_7 *Constructor_Test_Primes_Cons[int64] = v1_7_loop
_ = v1_7
var __t10 *Constructor_Test_Primes_Cons[int64]
{
if (v_6 == nil) {
__t10 = v1_7
goto end_branch_10
} else {

}
}
{
if (v_6 != nil) {
v_6_loop = (v_6).V1
v1_7_loop = (&Constructor_Test_Primes_Cons[int64]{1, (v_6).V0, v1_7})
continue go__302824162_5_8_13
__t10 = func() *Constructor_Test_Primes_Cons[int64] { panic("unreachable") }()
goto end_branch_10
} else {

}
}
{
__t10 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
}
}
go__302824162_5_8_13 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__302824162_5_8_13(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_6_loop_val)), Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_7_loop_val)))))}
})
})
Call_local_Test_Primes_go__go_5_9_14 = func(v_6_loop *Constructor_Test_Primes_Cons[int64], v1_7_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__go_5_9_14:
for {
if false { continue go__go_5_9_14 }
var v_6 *Constructor_Test_Primes_Cons[int64] = v_6_loop
_ = v_6
var v1_7 *Constructor_Test_Primes_Cons[int64] = v1_7_loop
_ = v1_7
var __t11 *Constructor_Test_Primes_Cons[int64]
{
if (v_6 == nil) {
__t11 = v1_7
goto end_branch_11
} else {

}
}
{
if (v_6 != nil) {
__t11 = Call_local_Test_Primes_go__302824162_5_8_13((v_6).V1, (&Constructor_Test_Primes_Cons[int64]{1, (v_6).V0, v1_7}))
goto end_branch_11
} else {

}
}
{
__t11 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}
}
go__go_5_9_14 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__go_5_9_14(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_6_loop_val)), Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_7_loop_val)))))}
})
})
__t13 = Call_local_Test_Primes_go__302824162_5_8_13(v1_4, (*Constructor_Test_Primes_Cons[int64])(nil))
goto end_branch_13
} else {

}
}
{
if (v_3 != nil) {
var __t12 *Constructor_Test_Primes_Cons[int64]
{
if (gopurs_runtime.Apply(p_0, gopurs_runtime.Int((v_3).V0)).IntVal) != (0) {
__t12 = Call_local_Test_Primes_go__302824162_2_0_9((v_3).V1, (&Constructor_Test_Primes_Cons[int64]{1, (v_3).V0, v1_4}))
goto end_branch_12
} else {

}
}
{
__t12 = Call_local_Test_Primes_go__302824162_2_0_9((v_3).V1, v1_4)
}
end_branch_12:
__t13 = __t12
goto end_branch_13
} else {

}
}
{
__t13 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}
}
go__go_2_1_10 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__go_2_1_10(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_3_loop_val)), Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_4_loop_val)))))}
})
})
return Call_local_Test_Primes_go__302824162_2_0_9(lst_1, (*Constructor_Test_Primes_Cons[int64])(nil))
}
}

func Call_Test_Primes_sieve(v_0_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
sieve:
for {
if false { continue sieve }
var v_0 *Constructor_Test_Primes_Cons[int64] = v_0_loop
_ = v_0
var __t1 *Constructor_Test_Primes_Cons[int64]
{
if (v_0 == nil) {
__t1 = (*Constructor_Test_Primes_Cons[int64])(nil)
goto end_branch_1
} else {

}
}
{
if (v_0 != nil) {
// TAST (Let): __local_var_1_0 shape=Other bindingType=Any
__local_var_1_0 := (v_0).V0
_ = __local_var_1_0
__t1 = (&Constructor_Test_Primes_Cons[int64]{1, __local_var_1_0, Call_Test_Primes_sieve(Call_Test_Primes_filter__394169086(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.IntMod(x_2.IntVal, __local_var_1_0)) == (int64(0))) != (true))
}), (v_0).V1))})
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}

func Rebox_Test_Primes_359351273_3637802162(in *Constructor_Test_Primes_Cons[gopurs_runtime.Value]) *Constructor_Test_Primes_Cons[int64] {
	if in == nil { return nil }
	out := &Constructor_Test_Primes_Cons[int64]{}
		out.V0 = in.V0.IntVal
		out.V1 = Rebox_Test_Primes_359351273_3637802162(in.V1)
	return out
}

func Rebox_Test_Primes_3637802162_359351273(in *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Test_Primes_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = Rebox_Test_Primes_3637802162_359351273(in.V1)
	return out
}



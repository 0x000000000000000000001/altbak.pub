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

var cache_Test_Primes_Cons__1975624008 gopurs_runtime.Value
var once_Test_Primes_Cons__1975624008 sync.Once
func Get_Test_Primes_Cons__1975624008() gopurs_runtime.Value {
	once_Test_Primes_Cons__1975624008.Do(func() {
		cache_Test_Primes_Cons__1975624008 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_Test_Primes_Cons__1975624008(__eta_norm_1_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](__eta_norm_0_1_box))))}
})
	})
	return cache_Test_Primes_Cons__1975624008
}

var cache_Test_Primes_sumList gopurs_runtime.Value
var once_Test_Primes_sumList sync.Once
func Get_Test_Primes_sumList() gopurs_runtime.Value {
	once_Test_Primes_sumList.Do(func() {
		cache_Test_Primes_sumList = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_Primes_sumList(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](lst_0_box)))
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

var cache_Test_Primes_reverse__2265911012 gopurs_runtime.Value
var once_Test_Primes_reverse__2265911012 sync.Once
func Get_Test_Primes_reverse__2265911012() gopurs_runtime.Value {
	once_Test_Primes_reverse__2265911012.Do(func() {
		cache_Test_Primes_reverse__2265911012 = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_Test_Primes_reverse__2265911012(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](lst_0_box))))}
})
	})
	return cache_Test_Primes_reverse__2265911012
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

var cache_Test_Primes_filter__3878167378 gopurs_runtime.Value
var once_Test_Primes_filter__3878167378 sync.Once
func Get_Test_Primes_filter__3878167378() gopurs_runtime.Value {
	once_Test_Primes_filter__3878167378.Do(func() {
		cache_Test_Primes_filter__3878167378 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, lst_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_Test_Primes_filter__3878167378(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](lst_1_box))))}
})
	})
	return cache_Test_Primes_filter__3878167378
}

var cache_Test_Primes_sieve gopurs_runtime.Value
var once_Test_Primes_sieve sync.Once
func Get_Test_Primes_sieve() gopurs_runtime.Value {
	once_Test_Primes_sieve.Do(func() {
		cache_Test_Primes_sieve = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_Test_Primes_sieve(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](v_0_box))))}
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
Call_local_Test_Primes_go__1466046018_2_2_15 = func(v_3_loop *Constructor_Test_Primes_Cons[int64], v1_4_loop int64) int64 {
go__1466046018_2_2_15:
for {
if false { continue go__1466046018_2_2_15 }
var v_3 *Constructor_Test_Primes_Cons[int64] = v_3_loop
_ = v_3
var v1_4 int64 = v1_4_loop
_ = v1_4
var __t3 int64
{
if (v_3 == nil) {
__t3 = v1_4
goto end_branch_3
} else {

}
}
{
if (v_3 != nil) {
v_3_loop = (v_3).V1
v1_4_loop = (v1_4) + ((v_3).V0)
continue go__1466046018_2_2_15
__t3 = func() int64 { panic("unreachable") }()
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
go__1466046018_2_2_15 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Test_Primes_go__1466046018_2_2_15(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](v_3_loop_val), v1_4_loop_val.IntVal))
})
})
var go__go_3_4_16 gopurs_runtime.Value
_ = go__go_3_4_16
// FALLBACK TCO: isLoop=false len=1
go__go_3_4_16 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 int64
{
var __t_tag_5 *Constructor_Test_Primes_Cons[int64] = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_4))
if (__t_tag_5 == nil) {
__t7 = v1_5.IntVal
goto end_branch_7
} else {

}
}
{
var __t_tag_6 *Constructor_Test_Primes_Cons[int64] = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_4))
if (__t_tag_6 != nil) {
__t7 = Call_local_Test_Primes_go__1466046018_2_2_15(Rebox_Test_Primes_359351273_3637802162((*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1), (v1_5.IntVal) + ((*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.IntVal))
goto end_branch_7
} else {

}
}
{
__t7 = func() int64 { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Int(__t7)
})
var Call_local_Test_Primes_go__29212279_4_8_17 func(int64, *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__29212279_4_8_17
var go__29212279_4_8_17 gopurs_runtime.Value
_ = go__29212279_4_8_17
Call_local_Test_Primes_go__29212279_4_8_17 = func(curr_5_loop int64, acc_6_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__29212279_4_8_17:
for {
if false { continue go__29212279_4_8_17 }
var curr_5 int64 = curr_5_loop
_ = curr_5
var acc_6 *Constructor_Test_Primes_Cons[int64] = acc_6_loop
_ = acc_6
var __t9 *Constructor_Test_Primes_Cons[int64]
{
if (curr_5) < (int64(2)) {
__t9 = acc_6
goto end_branch_9
} else {

}
}
{
curr_5_loop = (curr_5) - (int64(1))
acc_6_loop = (&Constructor_Test_Primes_Cons[int64]{1, curr_5, acc_6})
continue go__29212279_4_8_17
__t9 = func() *Constructor_Test_Primes_Cons[int64] { panic("unreachable") }()
}
end_branch_9:
return __t9
}
}
go__29212279_4_8_17 = gopurs_runtime.Func(func(curr_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__29212279_4_8_17(curr_5_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](acc_6_loop_val))))}
})
})
var go__go_5_10_18 gopurs_runtime.Value
_ = go__go_5_10_18
// FALLBACK TCO: isLoop=false len=1
go__go_5_10_18 = gopurs_runtime.Func2(func(curr_6 gopurs_runtime.Value, acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 *Constructor_Test_Primes_Cons[int64]
{
if (curr_6.IntVal) < (int64(2)) {
__t11 = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](acc_7))
goto end_branch_11
} else {

}
}
{
__t11 = Call_local_Test_Primes_go__29212279_4_8_17((curr_6.IntVal) - (int64(1)), (&Constructor_Test_Primes_Cons[int64]{1, curr_6.IntVal, Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](acc_7))}))
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(__t11))}
})
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_local_Test_Primes_go__1466046018_2_2_15(Call_Test_Primes_sieve(Call_local_Test_Primes_go__29212279_4_8_17(__local_var_1_1.IntVal, (*Constructor_Test_Primes_Cons[int64])(nil))), int64(0)))).StrVal())
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


func Call_Test_Primes_Cons__1975624008(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
Cons__1975624008:
for {
if false { continue Cons__1975624008 }
var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Test_Primes_Cons[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((&Constructor_Test_Primes_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](value1)}))}
})
}), gopurs_runtime.Int(__eta_norm_1_0), gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(__eta_norm_0_1))})))
}
}

func Call_Test_Primes_sumList(lst_0_loop *Constructor_Test_Primes_Cons[int64]) int64 {
var lst_0 *Constructor_Test_Primes_Cons[int64] = lst_0_loop
_ = lst_0
var Call_local_Test_Primes_go__1466046018_1_0_0 func(*Constructor_Test_Primes_Cons[int64], int64) int64
_ = Call_local_Test_Primes_go__1466046018_1_0_0
var go__1466046018_1_0_0 gopurs_runtime.Value
_ = go__1466046018_1_0_0
Call_local_Test_Primes_go__1466046018_1_0_0 = func(v_2_loop *Constructor_Test_Primes_Cons[int64], v1_3_loop int64) int64 {
go__1466046018_1_0_0:
for {
if false { continue go__1466046018_1_0_0 }
var v_2 *Constructor_Test_Primes_Cons[int64] = v_2_loop
_ = v_2
var v1_3 int64 = v1_3_loop
_ = v1_3
var __t1 int64
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
v1_3_loop = (v1_3) + ((v_2).V0)
continue go__1466046018_1_0_0
__t1 = func() int64 { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
__t1 = func() int64 { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}
go__1466046018_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_local_Test_Primes_go__1466046018_1_0_0(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](v_2_loop_val), v1_3_loop_val.IntVal))
})
})
var go__go_2_2_1 gopurs_runtime.Value
_ = go__go_2_2_1
// FALLBACK TCO: isLoop=false len=1
go__go_2_2_1 = gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 int64
{
var __t_tag_3 *Constructor_Test_Primes_Cons[int64] = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_3))
if (__t_tag_3 == nil) {
__t5 = v1_4.IntVal
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Test_Primes_Cons[int64] = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_3))
if (__t_tag_4 != nil) {
__t5 = Call_local_Test_Primes_go__1466046018_1_0_0(Rebox_Test_Primes_359351273_3637802162((*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1), (v1_4.IntVal) + ((*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0.IntVal))
goto end_branch_5
} else {

}
}
{
__t5 = func() int64 { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Int(__t5)
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

func Call_Test_Primes_reverse__2265911012(lst_0_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
reverse__2265911012:
for {
if false { continue reverse__2265911012 }
var lst_0 *Constructor_Test_Primes_Cons[int64] = lst_0_loop
_ = lst_0
var Call_local_Test_Primes_go__302824162_1_0_3 func(*Constructor_Test_Primes_Cons[int64], *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__302824162_1_0_3
var go__302824162_1_0_3 gopurs_runtime.Value
_ = go__302824162_1_0_3
Call_local_Test_Primes_go__302824162_1_0_3 = func(v_2_loop *Constructor_Test_Primes_Cons[int64], v1_3_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__302824162_1_0_3:
for {
if false { continue go__302824162_1_0_3 }
var v_2 *Constructor_Test_Primes_Cons[int64] = v_2_loop
_ = v_2
var v1_3 *Constructor_Test_Primes_Cons[int64] = v1_3_loop
_ = v1_3
var __t1 *Constructor_Test_Primes_Cons[int64]
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
v1_3_loop = (&Constructor_Test_Primes_Cons[int64]{1, gopurs_runtime.Int((v_2).V0).IntVal, v1_3})
continue go__302824162_1_0_3
__t1 = func() *Constructor_Test_Primes_Cons[int64] { panic("unreachable") }()
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
go__302824162_1_0_3 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__302824162_1_0_3(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](v1_3_loop_val))))}
})
})
var go__go_2_2_4 gopurs_runtime.Value
_ = go__go_2_2_4
// FALLBACK TCO: isLoop=false len=1
go__go_2_2_4 = gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Test_Primes_Cons[int64]
{
var __t_tag_3 *Constructor_Test_Primes_Cons[int64] = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_3))
if (__t_tag_3 == nil) {
__t5 = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_4))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Test_Primes_Cons[int64] = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_3))
if (__t_tag_4 != nil) {
__t5 = Call_local_Test_Primes_go__302824162_1_0_3(Rebox_Test_Primes_359351273_3637802162((*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1), (&Constructor_Test_Primes_Cons[int64]{1, (*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0.IntVal, Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_4))}))
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(__t5))}
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
Call_local_Test_Primes_go__29212279_2_0_5 = func(curr_3_loop int64, acc_4_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__29212279_2_0_5:
for {
if false { continue go__29212279_2_0_5 }
var curr_3 int64 = curr_3_loop
_ = curr_3
var acc_4 *Constructor_Test_Primes_Cons[int64] = acc_4_loop
_ = acc_4
var __t1 *Constructor_Test_Primes_Cons[int64]
{
if (curr_3) < (start_0) {
__t1 = acc_4
goto end_branch_1
} else {

}
}
{
curr_3_loop = (curr_3) - (int64(1))
acc_4_loop = (&Constructor_Test_Primes_Cons[int64]{1, curr_3, acc_4})
continue go__29212279_2_0_5
__t1 = func() *Constructor_Test_Primes_Cons[int64] { panic("unreachable") }()
}
end_branch_1:
return __t1
}
}
go__29212279_2_0_5 = gopurs_runtime.Func(func(curr_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__29212279_2_0_5(curr_3_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](acc_4_loop_val))))}
})
})
var go__go_3_2_6 gopurs_runtime.Value
_ = go__go_3_2_6
// FALLBACK TCO: isLoop=false len=1
go__go_3_2_6 = gopurs_runtime.Func2(func(curr_4 gopurs_runtime.Value, acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Test_Primes_Cons[int64]
{
if (curr_4.IntVal) < (start_0) {
__t3 = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](acc_5))
goto end_branch_3
} else {

}
}
{
__t3 = Call_local_Test_Primes_go__29212279_2_0_5((curr_4.IntVal) - (int64(1)), (&Constructor_Test_Primes_Cons[int64]{1, curr_4.IntVal, Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](acc_5))}))
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(__t3))}
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

func Call_Test_Primes_filter__3878167378(p_0_loop gopurs_runtime.Value, lst_1_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
filter__3878167378:
for {
if false { continue filter__3878167378 }
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var lst_1 *Constructor_Test_Primes_Cons[int64] = lst_1_loop
_ = lst_1
var Call_local_Test_Primes_go__302824162_2_0_9 func(*Constructor_Test_Primes_Cons[int64], *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__302824162_2_0_9
var go__302824162_2_0_9 gopurs_runtime.Value
_ = go__302824162_2_0_9
Call_local_Test_Primes_go__302824162_2_0_9 = func(v_3_loop *Constructor_Test_Primes_Cons[int64], v1_4_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__302824162_2_0_9:
for {
if false { continue go__302824162_2_0_9 }
var v_3 *Constructor_Test_Primes_Cons[int64] = v_3_loop
_ = v_3
var v1_4 *Constructor_Test_Primes_Cons[int64] = v1_4_loop
_ = v1_4
var __t8 *Constructor_Test_Primes_Cons[int64]
{
if (v_3 == nil) {
var Call_local_Test_Primes_go__302824162_5_1_10 func(*Constructor_Test_Primes_Cons[int64], *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__302824162_5_1_10
var go__302824162_5_1_10 gopurs_runtime.Value
_ = go__302824162_5_1_10
Call_local_Test_Primes_go__302824162_5_1_10 = func(v_6_loop *Constructor_Test_Primes_Cons[int64], v1_7_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__302824162_5_1_10:
for {
if false { continue go__302824162_5_1_10 }
var v_6 *Constructor_Test_Primes_Cons[int64] = v_6_loop
_ = v_6
var v1_7 *Constructor_Test_Primes_Cons[int64] = v1_7_loop
_ = v1_7
var __t2 *Constructor_Test_Primes_Cons[int64]
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
v1_7_loop = (&Constructor_Test_Primes_Cons[int64]{1, gopurs_runtime.Int((v_6).V0).IntVal, v1_7})
continue go__302824162_5_1_10
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
go__302824162_5_1_10 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__302824162_5_1_10(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](v_6_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](v1_7_loop_val))))}
})
})
var go__go_6_3_11 gopurs_runtime.Value
_ = go__go_6_3_11
// FALLBACK TCO: isLoop=false len=1
go__go_6_3_11 = gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Test_Primes_Cons[int64]
{
var __t_tag_4 *Constructor_Test_Primes_Cons[int64] = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_7))
if (__t_tag_4 == nil) {
__t6 = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_8))
goto end_branch_6
} else {

}
}
{
var __t_tag_5 *Constructor_Test_Primes_Cons[int64] = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_7))
if (__t_tag_5 != nil) {
__t6 = Call_local_Test_Primes_go__302824162_5_1_10(Rebox_Test_Primes_359351273_3637802162((*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_7.UnsafePtr).V1), (&Constructor_Test_Primes_Cons[int64]{1, (*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_7.UnsafePtr).V0.IntVal, Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_8))}))
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(__t6))}
})
__t8 = Call_local_Test_Primes_go__302824162_5_1_10(v1_4, (*Constructor_Test_Primes_Cons[int64])(nil))
goto end_branch_8
} else {

}
}
{
if (v_3 != nil) {
var __t7 *Constructor_Test_Primes_Cons[int64]
{
if (gopurs_runtime.Apply(p_0, gopurs_runtime.Int((v_3).V0)).IntVal) != (0) {
v_3_loop = (v_3).V1
v1_4_loop = (&Constructor_Test_Primes_Cons[int64]{1, gopurs_runtime.Int((v_3).V0).IntVal, v1_4})
continue go__302824162_2_0_9
__t7 = func() *Constructor_Test_Primes_Cons[int64] { panic("unreachable") }()
goto end_branch_7
} else {

}
}
{
v_3_loop = (v_3).V1
v1_4_loop = v1_4
continue go__302824162_2_0_9
__t7 = func() *Constructor_Test_Primes_Cons[int64] { panic("unreachable") }()
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
__t8 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}
}
go__302824162_2_0_9 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__302824162_2_0_9(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](v_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](v1_4_loop_val))))}
})
})
var go__go_3_9_12 gopurs_runtime.Value
_ = go__go_3_9_12
// FALLBACK TCO: isLoop=false len=1
go__go_3_9_12 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 *Constructor_Test_Primes_Cons[int64]
{
var __t_tag_10 *Constructor_Test_Primes_Cons[int64] = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_4))
if (__t_tag_10 == nil) {
var Call_local_Test_Primes_go__302824162_6_11_13 func(*Constructor_Test_Primes_Cons[int64], *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64]
_ = Call_local_Test_Primes_go__302824162_6_11_13
var go__302824162_6_11_13 gopurs_runtime.Value
_ = go__302824162_6_11_13
Call_local_Test_Primes_go__302824162_6_11_13 = func(v_7_loop *Constructor_Test_Primes_Cons[int64], v1_8_loop *Constructor_Test_Primes_Cons[int64]) *Constructor_Test_Primes_Cons[int64] {
go__302824162_6_11_13:
for {
if false { continue go__302824162_6_11_13 }
var v_7 *Constructor_Test_Primes_Cons[int64] = v_7_loop
_ = v_7
var v1_8 *Constructor_Test_Primes_Cons[int64] = v1_8_loop
_ = v1_8
var __t12 *Constructor_Test_Primes_Cons[int64]
{
if (v_7 == nil) {
__t12 = v1_8
goto end_branch_12
} else {

}
}
{
if (v_7 != nil) {
v_7_loop = (v_7).V1
v1_8_loop = (&Constructor_Test_Primes_Cons[int64]{1, gopurs_runtime.Int((v_7).V0).IntVal, v1_8})
continue go__302824162_6_11_13
__t12 = func() *Constructor_Test_Primes_Cons[int64] { panic("unreachable") }()
goto end_branch_12
} else {

}
}
{
__t12 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}
}
go__302824162_6_11_13 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(Call_local_Test_Primes_go__302824162_6_11_13(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](v_7_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[int64]](v1_8_loop_val))))}
})
})
var go__go_7_13_14 gopurs_runtime.Value
_ = go__go_7_13_14
// FALLBACK TCO: isLoop=false len=1
go__go_7_13_14 = gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 *Constructor_Test_Primes_Cons[int64]
{
var __t_tag_14 *Constructor_Test_Primes_Cons[int64] = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_8))
if (__t_tag_14 == nil) {
__t16 = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_9))
goto end_branch_16
} else {

}
}
{
var __t_tag_15 *Constructor_Test_Primes_Cons[int64] = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_8))
if (__t_tag_15 != nil) {
__t16 = Call_local_Test_Primes_go__302824162_6_11_13(Rebox_Test_Primes_359351273_3637802162((*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_8.UnsafePtr).V1), (&Constructor_Test_Primes_Cons[int64]{1, (*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_8.UnsafePtr).V0.IntVal, Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_9))}))
goto end_branch_16
} else {

}
}
{
__t16 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_16:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(__t16))}
})
__t19 = Call_local_Test_Primes_go__302824162_6_11_13(Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_5)), (*Constructor_Test_Primes_Cons[int64])(nil))
goto end_branch_19
} else {

}
}
{
var __t_tag_17 *Constructor_Test_Primes_Cons[int64] = Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v_4))
if (__t_tag_17 != nil) {
var __t18 *Constructor_Test_Primes_Cons[int64]
{
if (gopurs_runtime.Apply(p_0, gopurs_runtime.Int((*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.IntVal)).IntVal) != (0) {
__t18 = Call_local_Test_Primes_go__302824162_2_0_9(Rebox_Test_Primes_359351273_3637802162((*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1), (&Constructor_Test_Primes_Cons[int64]{1, (*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.IntVal, Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_5))}))
goto end_branch_18
} else {

}
}
{
__t18 = Call_local_Test_Primes_go__302824162_2_0_9(Rebox_Test_Primes_359351273_3637802162((*Constructor_Test_Primes_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1), Rebox_Test_Primes_359351273_3637802162(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons[gopurs_runtime.Value]](v1_5)))
}
end_branch_18:
__t19 = __t18
goto end_branch_19
} else {

}
}
{
__t19 = func() *Constructor_Test_Primes_Cons[int64] { panic("Failed pattern match") }()
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Rebox_Test_Primes_3637802162_359351273(__t19))}
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
__t1 = (&Constructor_Test_Primes_Cons[int64]{1, gopurs_runtime.Int(__local_var_1_0).IntVal, Call_Test_Primes_sieve(Call_Test_Primes_filter__3878167378(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((((x_2.IntVal) % (__local_var_1_0)) == (int64(0))) != (true))
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



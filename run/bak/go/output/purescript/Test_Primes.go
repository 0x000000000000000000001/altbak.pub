package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Test_Primes_logShow gopurs_runtime.Value
var once_Test_Primes_logShow sync.Once
func Get_Test_Primes_logShow() gopurs_runtime.Value {
	once_Test_Primes_logShow.Do(func() {
		cache_Test_Primes_logShow = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_Primes_logShow(a_0_box.IntVal)
})
	})
	return cache_Test_Primes_logShow
}

var cache_Test_Primes_Nil gopurs_runtime.Value
var once_Test_Primes_Nil sync.Once
func Get_Test_Primes_Nil() gopurs_runtime.Value {
	once_Test_Primes_Nil.Do(func() {
		cache_Test_Primes_Nil = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(nil))}
	})
	return cache_Test_Primes_Nil
}

var cache_Test_Primes_Cons gopurs_runtime.Value
var once_Test_Primes_Cons sync.Once
func Get_Test_Primes_Cons() gopurs_runtime.Value {
	once_Test_Primes_Cons.Do(func() {
		cache_Test_Primes_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Test_Primes_Cons{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](value1)})}
})
})
	})
	return cache_Test_Primes_Cons
}

var cache_Test_Primes_sumList gopurs_runtime.Value
var once_Test_Primes_sumList sync.Once
func Get_Test_Primes_sumList() gopurs_runtime.Value {
	once_Test_Primes_sumList.Do(func() {
		cache_Test_Primes_sumList = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_Primes_sumList(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](lst_0_box)))
})
	})
	return cache_Test_Primes_sumList
}

var cache_Test_Primes_reverse gopurs_runtime.Value
var once_Test_Primes_reverse sync.Once
func Get_Test_Primes_reverse() gopurs_runtime.Value {
	once_Test_Primes_reverse.Do(func() {
		cache_Test_Primes_reverse = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_Test_Primes_reverse(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](lst_0_box)))}
})
	})
	return cache_Test_Primes_reverse
}

var cache_Test_Primes_go__range gopurs_runtime.Value
var once_Test_Primes_go__range sync.Once
func Get_Test_Primes_go__range() gopurs_runtime.Value {
	once_Test_Primes_go__range.Do(func() {
		cache_Test_Primes_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_Test_Primes_go__range(start_0_box.IntVal, end_1_box.IntVal))}
})
	})
	return cache_Test_Primes_go__range
}

var cache_Test_Primes_filter gopurs_runtime.Value
var once_Test_Primes_filter sync.Once
func Get_Test_Primes_filter() gopurs_runtime.Value {
	once_Test_Primes_filter.Do(func() {
		cache_Test_Primes_filter = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, lst_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_Test_Primes_filter(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](lst_1_box)))}
})
	})
	return cache_Test_Primes_filter
}

var cache_Test_Primes_sieve gopurs_runtime.Value
var once_Test_Primes_sieve sync.Once
func Get_Test_Primes_sieve() gopurs_runtime.Value {
	once_Test_Primes_sieve.Do(func() {
		cache_Test_Primes_sieve = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_Test_Primes_sieve(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](v_0_box)))}
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
		cache_Test_Primes_act = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(500))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = __local_var_1_1
var go__go_2_2_5 gopurs_runtime.Value
go__go_2_2_5 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop int64 = v1_4_loop_val.IntVal
go__go_2_2_5:
for {
if false { continue go__go_2_2_5 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 int64 = v1_4_loop
_ = v1_4
var __t3 int64
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr == nil) {
__t3 = v1_4
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr != nil) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V1)}
v1_4_loop = (v1_4) + ((*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V0.IntVal)
continue go__go_2_2_5
__t3 = gopurs_runtime.Value{}.IntVal
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_3:
return gopurs_runtime.Int(__t3)
}
}()
})
})
var go__go_3_4_6 gopurs_runtime.Value
go__go_3_4_6 = gopurs_runtime.Func(func(curr_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var curr_4_loop int64 = curr_4_loop_val.IntVal
var acc_5_loop *Constructor_Test_Primes_Cons = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](acc_5_loop_val)
go__go_3_4_6:
for {
if false { continue go__go_3_4_6 }
var curr_4 int64 = curr_4_loop
_ = curr_4
var acc_5 *Constructor_Test_Primes_Cons = acc_5_loop
_ = acc_5
var __t6 *Constructor_Test_Primes_Cons
{
var __t5 bool
{
if (curr_4) < (2) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
if __t5 {
__t6 = acc_5
goto end_branch_6
} else {

}
}
{
curr_4_loop = (curr_4) - (1)
acc_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Test_Primes_Cons{1, gopurs_runtime.Int(curr_4), gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(acc_5)})})})
continue go__go_3_4_6
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t6)}
}
}()
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(gopurs_runtime.Apply2(go__go_2_2_5, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_Test_Primes_sieve(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Apply2(go__go_3_4_6, gopurs_runtime.Int(__local_var_1_1.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(nil))}))))}, gopurs_runtime.Int(0)).IntVal)).StrVal())), gopurs_runtime.Value{})
})
}()
	})
	return cache_Test_Primes_act
}

var cache_Test_Primes_filter__1481233142 gopurs_runtime.Value
var once_Test_Primes_filter__1481233142 sync.Once
func Get_Test_Primes_filter__1481233142() gopurs_runtime.Value {
	once_Test_Primes_filter__1481233142.Do(func() {
		cache_Test_Primes_filter__1481233142 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, lst_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_Test_Primes_filter__1481233142(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](lst_1_box)))}
})
	})
	return cache_Test_Primes_filter__1481233142
}

var cache_Test_Primes_filter__37320371 gopurs_runtime.Value
var once_Test_Primes_filter__37320371 sync.Once
func Get_Test_Primes_filter__37320371() gopurs_runtime.Value {
	once_Test_Primes_filter__37320371.Do(func() {
		cache_Test_Primes_filter__37320371 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, lst_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_Test_Primes_filter__37320371(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](lst_1_box)))}
})
	})
	return cache_Test_Primes_filter__37320371
}

var cache_Test_Primes_reverse__3030426720 gopurs_runtime.Value
var once_Test_Primes_reverse__3030426720 sync.Once
func Get_Test_Primes_reverse__3030426720() gopurs_runtime.Value {
	once_Test_Primes_reverse__3030426720.Do(func() {
		cache_Test_Primes_reverse__3030426720 = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_Test_Primes_reverse__3030426720(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](lst_0_box)))}
})
	})
	return cache_Test_Primes_reverse__3030426720
}

var cache_Test_Primes_reverse__3119428352 gopurs_runtime.Value
var once_Test_Primes_reverse__3119428352 sync.Once
func Get_Test_Primes_reverse__3119428352() gopurs_runtime.Value {
	once_Test_Primes_reverse__3119428352.Do(func() {
		cache_Test_Primes_reverse__3119428352 = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_Test_Primes_reverse__3119428352(gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](lst_0_box)))}
})
	})
	return cache_Test_Primes_reverse__3119428352
}

type Constructor_Test_Primes_Nil struct {
	Rc uint32
}


type Constructor_Test_Primes_Cons struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 *Constructor_Test_Primes_Cons
}


func Call_Test_Primes_logShow(a_0_loop int64) gopurs_runtime.Value {
var a_0 int64 = a_0_loop
_ = a_0
return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(a_0)).StrVal()))
}

func Call_Test_Primes_sumList(lst_0_loop *Constructor_Test_Primes_Cons) int64 {
var lst_0 *Constructor_Test_Primes_Cons = lst_0_loop
_ = lst_0
var go__go_1_0_0 gopurs_runtime.Value
go__go_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop int64 = v1_3_loop_val.IntVal
go__go_1_0_0:
for {
if false { continue go__go_1_0_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 int64 = v1_3_loop
_ = v1_3
var __t1 int64
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr == nil) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_2.UnsafePtr).V1)}
v1_3_loop = (v1_3) + ((*Constructor_Test_Primes_Cons)(v_2.UnsafePtr).V0.IntVal)
continue go__go_1_0_0
__t1 = gopurs_runtime.Value{}.IntVal
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_1:
return gopurs_runtime.Int(__t1)
}
}()
})
})
return gopurs_runtime.Apply2(go__go_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Int(0)).IntVal
}

func Call_Test_Primes_reverse(lst_0_loop *Constructor_Test_Primes_Cons) *Constructor_Test_Primes_Cons {
var lst_0 *Constructor_Test_Primes_Cons = lst_0_loop
_ = lst_0
var go__go_1_0_1 gopurs_runtime.Value
go__go_1_0_1 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *Constructor_Test_Primes_Cons = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](v1_3_loop_val)
go__go_1_0_1:
for {
if false { continue go__go_1_0_1 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *Constructor_Test_Primes_Cons = v1_3_loop
_ = v1_3
var __t1 *Constructor_Test_Primes_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr == nil) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_2.UnsafePtr).V1)}
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Test_Primes_Cons{1, (*Constructor_Test_Primes_Cons)(v_2.UnsafePtr).V0, v1_3})})
continue go__go_1_0_1
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Apply2(go__go_1_0_1, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(nil))}))
}

func Call_Test_Primes_go__range(start_0_loop int64, end_1_loop int64) *Constructor_Test_Primes_Cons {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var go__go_2_0_2 gopurs_runtime.Value
go__go_2_0_2 = gopurs_runtime.Func(func(curr_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var curr_3_loop int64 = curr_3_loop_val.IntVal
var acc_4_loop *Constructor_Test_Primes_Cons = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](acc_4_loop_val)
go__go_2_0_2:
for {
if false { continue go__go_2_0_2 }
var curr_3 int64 = curr_3_loop
_ = curr_3
var acc_4 *Constructor_Test_Primes_Cons = acc_4_loop
_ = acc_4
var __t2 *Constructor_Test_Primes_Cons
{
var __t1 bool
{
if (curr_3) < (start_0) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = acc_4
goto end_branch_2
} else {

}
}
{
curr_3_loop = (curr_3) - (1)
acc_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Test_Primes_Cons{1, gopurs_runtime.Int(curr_3), gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(acc_4)})})})
continue go__go_2_0_2
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Apply2(go__go_2_0_2, gopurs_runtime.Int(end_1), gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(nil))}))
}

func Call_Test_Primes_filter(p_0_loop gopurs_runtime.Value, lst_1_loop *Constructor_Test_Primes_Cons) *Constructor_Test_Primes_Cons {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var lst_1 *Constructor_Test_Primes_Cons = lst_1_loop
_ = lst_1
var go__go_2_0_3 gopurs_runtime.Value
go__go_2_0_3 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop *Constructor_Test_Primes_Cons = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](v1_4_loop_val)
go__go_2_0_3:
for {
if false { continue go__go_2_0_3 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 *Constructor_Test_Primes_Cons = v1_4_loop
_ = v1_4
var __t4 *Constructor_Test_Primes_Cons
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr == nil) {
var go__go_5_1_4 gopurs_runtime.Value
go__go_5_1_4 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop *Constructor_Test_Primes_Cons = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](v1_7_loop_val)
go__go_5_1_4:
for {
if false { continue go__go_5_1_4 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 *Constructor_Test_Primes_Cons = v1_7_loop
_ = v1_7
var __t2 *Constructor_Test_Primes_Cons
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr == nil) {
__t2 = v1_7
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_6.UnsafePtr).V1)}
v1_7_loop = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Test_Primes_Cons{1, (*Constructor_Test_Primes_Cons)(v_6.UnsafePtr).V0, v1_7})})
continue go__go_5_1_4
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Apply2(go__go_5_1_4, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v1_4)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(nil))}))
goto end_branch_4
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr != nil) {
var __t3 *Constructor_Test_Primes_Cons
{
if (gopurs_runtime.Apply(p_0, (*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V0).IntVal) != (0) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V1)}
v1_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Test_Primes_Cons{1, (*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V0, v1_4})})
continue go__go_2_0_3
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V1)}
v1_4_loop = v1_4
continue go__go_2_0_3
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t4)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Apply2(go__go_2_0_3, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(nil))}))
}

func Call_Test_Primes_sieve(v_0_loop *Constructor_Test_Primes_Cons) *Constructor_Test_Primes_Cons {
sieve:
for {
if false { continue sieve }
var v_0 *Constructor_Test_Primes_Cons = v_0_loop
_ = v_0
var __t1 *Constructor_Test_Primes_Cons
{
if (v_0 == nil) {
__t1 = (*Constructor_Test_Primes_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v_0 != nil) {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := (v_0).V0
_ = __local_var_1_0
__t1 = &Constructor_Test_Primes_Cons{1, gopurs_runtime.Int(__local_var_1_0.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_Test_Primes_sieve(Call_Test_Primes_filter__1481233142(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), gopurs_runtime.Int(x_2.IntVal), gopurs_runtime.Int(__local_var_1_0.IntVal)).IntVal) == (0)) != (true))
}), gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((v_0).V1)}))))})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t1)})
}
}

func Call_Test_Primes_filter__1481233142(p_0_loop gopurs_runtime.Value, lst_1_loop *Constructor_Test_Primes_Cons) *Constructor_Test_Primes_Cons {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var lst_1 *Constructor_Test_Primes_Cons = lst_1_loop
_ = lst_1
var go__go_2_0_7 gopurs_runtime.Value
go__go_2_0_7 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop *Constructor_Test_Primes_Cons = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](v1_4_loop_val)
go__go_2_0_7:
for {
if false { continue go__go_2_0_7 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 *Constructor_Test_Primes_Cons = v1_4_loop
_ = v1_4
var __t4 *Constructor_Test_Primes_Cons
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr == nil) {
var go__go_5_1_8 gopurs_runtime.Value
go__go_5_1_8 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop *Constructor_Test_Primes_Cons = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](v1_7_loop_val)
go__go_5_1_8:
for {
if false { continue go__go_5_1_8 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 *Constructor_Test_Primes_Cons = v1_7_loop
_ = v1_7
var __t2 *Constructor_Test_Primes_Cons
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr == nil) {
__t2 = v1_7
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_6.UnsafePtr).V1)}
v1_7_loop = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Test_Primes_Cons{1, (*Constructor_Test_Primes_Cons)(v_6.UnsafePtr).V0, v1_7})})
continue go__go_5_1_8
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Apply2(go__go_5_1_8, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v1_4)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(nil))}))
goto end_branch_4
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr != nil) {
var __t3 *Constructor_Test_Primes_Cons
{
if (gopurs_runtime.Apply(p_0, gopurs_runtime.Int((*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V0.IntVal)).IntVal) != (0) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V1)}
v1_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Test_Primes_Cons{1, gopurs_runtime.Int((*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V0.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v1_4)})})})
continue go__go_2_0_7
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V1)}
v1_4_loop = v1_4
continue go__go_2_0_7
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t3)})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t4)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Apply2(go__go_2_0_7, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(nil))}))
}

func Call_Test_Primes_filter__37320371(p_0_loop gopurs_runtime.Value, lst_1_loop *Constructor_Test_Primes_Cons) *Constructor_Test_Primes_Cons {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var lst_1 *Constructor_Test_Primes_Cons = lst_1_loop
_ = lst_1
var go__go_2_0_9 gopurs_runtime.Value
go__go_2_0_9 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop *Constructor_Test_Primes_Cons = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](v1_4_loop_val)
go__go_2_0_9:
for {
if false { continue go__go_2_0_9 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 *Constructor_Test_Primes_Cons = v1_4_loop
_ = v1_4
var __t4 *Constructor_Test_Primes_Cons
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr == nil) {
var go__go_5_1_10 gopurs_runtime.Value
go__go_5_1_10 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop *Constructor_Test_Primes_Cons = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](v1_7_loop_val)
go__go_5_1_10:
for {
if false { continue go__go_5_1_10 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 *Constructor_Test_Primes_Cons = v1_7_loop
_ = v1_7
var __t2 *Constructor_Test_Primes_Cons
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr == nil) {
__t2 = v1_7
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_6.UnsafePtr).V1)}
v1_7_loop = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Test_Primes_Cons{1, (*Constructor_Test_Primes_Cons)(v_6.UnsafePtr).V0, v1_7})})
continue go__go_5_1_10
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Apply2(go__go_5_1_10, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v1_4)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(nil))}))
goto end_branch_4
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr != nil) {
var __t3 *Constructor_Test_Primes_Cons
{
if (gopurs_runtime.Apply(p_0, (*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V0).IntVal) != (0) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V1)}
v1_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Test_Primes_Cons{1, (*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V0, v1_4})})
continue go__go_2_0_9
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_3.UnsafePtr).V1)}
v1_4_loop = v1_4
continue go__go_2_0_9
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t4)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Apply2(go__go_2_0_9, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(nil))}))
}

func Call_Test_Primes_reverse__3030426720(lst_0_loop *Constructor_Test_Primes_Cons) *Constructor_Test_Primes_Cons {
var lst_0 *Constructor_Test_Primes_Cons = lst_0_loop
_ = lst_0
var go__go_1_0_11 gopurs_runtime.Value
go__go_1_0_11 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *Constructor_Test_Primes_Cons = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](v1_3_loop_val)
go__go_1_0_11:
for {
if false { continue go__go_1_0_11 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *Constructor_Test_Primes_Cons = v1_3_loop
_ = v1_3
var __t1 *Constructor_Test_Primes_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr == nil) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_2.UnsafePtr).V1)}
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Test_Primes_Cons{1, gopurs_runtime.Int((*Constructor_Test_Primes_Cons)(v_2.UnsafePtr).V0.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v1_3)})})})
continue go__go_1_0_11
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Apply2(go__go_1_0_11, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(nil))}))
}

func Call_Test_Primes_reverse__3119428352(lst_0_loop *Constructor_Test_Primes_Cons) *Constructor_Test_Primes_Cons {
var lst_0 *Constructor_Test_Primes_Cons = lst_0_loop
_ = lst_0
var go__go_1_0_12 gopurs_runtime.Value
go__go_1_0_12 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *Constructor_Test_Primes_Cons = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](v1_3_loop_val)
go__go_1_0_12:
for {
if false { continue go__go_1_0_12 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *Constructor_Test_Primes_Cons = v1_3_loop
_ = v1_3
var __t1 *Constructor_Test_Primes_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr == nil) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(v_2.UnsafePtr).V1)}
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Constructor_Test_Primes_Cons{1, (*Constructor_Test_Primes_Cons)(v_2.UnsafePtr).V0, v1_3})})
continue go__go_1_0_12
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Test_Primes_Cons](gopurs_runtime.Apply2(go__go_1_0_12, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*Constructor_Test_Primes_Cons)(nil))}))
}



package Test_Primes

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	unsafe "unsafe"
)

var Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		Nil = gopurs_runtime.Value{Type: 9, IntVal: 3777797863, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Nil{})}
	})
	return Nil
}

var Cons gopurs_runtime.Value
var once_Cons sync.Once
func Get_Cons() gopurs_runtime.Value {
	once_Cons.Do(func() {
		Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Cons{value0, value1})}
})
})
	})
	return Cons
}

var sumList gopurs_runtime.Value
var once_sumList sync.Once
func Get_sumList() gopurs_runtime.Value {
	once_sumList.Do(func() {
		sumList = gopurs_runtime.Func(func(lst_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var lst_0 gopurs_runtime.Value = lst_0_loop
_ = lst_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3777797863) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629) {
v_2_loop = (*Data_Test_Primes_Cons)(v_2.UnsafePtr).V1
v1_3_loop = gopurs_runtime.Int(v1_3.IntVal + (*Data_Test_Primes_Cons)(v_2.UnsafePtr).V0.IntVal)
continue go__1_0
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
return gopurs_runtime.Apply2(go__1_0, lst_0, gopurs_runtime.Int(0))
}()
})
	})
	return sumList
}

var reverse gopurs_runtime.Value
var once_reverse sync.Once
func Get_reverse() gopurs_runtime.Value {
	once_reverse.Do(func() {
		reverse = gopurs_runtime.Func(func(lst_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var lst_0 gopurs_runtime.Value = lst_0_loop
_ = lst_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3777797863) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629) {
v_2_loop = (*Data_Test_Primes_Cons)(v_2.UnsafePtr).V1
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Cons{(*Data_Test_Primes_Cons)(v_2.UnsafePtr).V0, v1_3})}
continue go__1_0
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
return gopurs_runtime.Apply2(go__1_0, lst_0, gopurs_runtime.Value{Type: 9, IntVal: 3777797863, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Nil{})})
}()
})
	})
	return reverse
}

var range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		range_ = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_range_(start_0_box, end_1_box)
})
	})
	return range_
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, lst_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter(p_0_box, lst_1_box)
})
	})
	return filter
}

var sieve gopurs_runtime.Value
var once_sieve sync.Once
func Get_sieve() gopurs_runtime.Value {
	once_sieve.Do(func() {
		sieve = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
sieve:
for {
if false { continue sieve }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3777797863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3777797863, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Nil{})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2390177629) {
__local_var_1_1 := (*Data_Test_Primes_Cons)(v_0.UnsafePtr).V0
_ = __local_var_1_1
var go__2_2 gopurs_runtime.Value
go__2_2 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_2:
for {
if false { continue go__2_2 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 3777797863) {
var go__5_4 gopurs_runtime.Value
go__5_4 = gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_4:
for {
if false { continue go__5_4 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t5 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 3777797863) {
__t5 = v1_7
goto end_branch_5
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629) {
v_6_loop = (*Data_Test_Primes_Cons)(v_6.UnsafePtr).V1
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Cons{(*Data_Test_Primes_Cons)(v_6.UnsafePtr).V0, v1_7})}
continue go__5_4
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
__t3 = gopurs_runtime.Apply2(go__5_4, v1_4, gopurs_runtime.Value{Type: 9, IntVal: 3777797863, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Nil{})})
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629) {
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), (*Data_Test_Primes_Cons)(v_3.UnsafePtr).V0, __local_var_1_1).IntVal != 0 {
v_3_loop = (*Data_Test_Primes_Cons)(v_3.UnsafePtr).V1
v1_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Cons{(*Data_Test_Primes_Cons)(v_3.UnsafePtr).V0, v1_4})}
continue go__2_2
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
v_3_loop = (*Data_Test_Primes_Cons)(v_3.UnsafePtr).V1
v1_4_loop = v1_4
continue go__2_2
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Cons{__local_var_1_1, gopurs_runtime.Apply(Get_sieve(), gopurs_runtime.Apply2(go__2_2, (*Data_Test_Primes_Cons)(v_0.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 3777797863, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Nil{})}))})}
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
	return sieve
}

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Prime Sieve (sum primes up to 500):"))
	})
	return describe
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(500))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = dummy_1_1
var go__2_2 gopurs_runtime.Value
go__2_2 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_2:
for {
if false { continue go__2_2 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 3777797863) {
__t3 = v1_4
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629) {
v_3_loop = (*Data_Test_Primes_Cons)(v_3.UnsafePtr).V1
v1_4_loop = gopurs_runtime.Int(v1_4.IntVal + (*Data_Test_Primes_Cons)(v_3.UnsafePtr).V0.IntVal)
continue go__2_2
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
var go__3_4 gopurs_runtime.Value
go__3_4 = gopurs_runtime.Func(func(curr_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_4:
for {
if false { continue go__3_4 }
var curr_4 gopurs_runtime.Value = curr_4_loop
_ = curr_4
var acc_5 gopurs_runtime.Value = acc_5_loop
_ = acc_5
var __t5 gopurs_runtime.Value
{
if curr_4.IntVal < 2 {
__t5 = acc_5
goto end_branch_5
} else {

}
}
{
curr_4_loop = gopurs_runtime.Int(curr_4.IntVal - 1)
acc_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Cons{curr_4, acc_5})}
continue go__3_4
__t5 = gopurs_runtime.Value{}
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply2(go__2_2, gopurs_runtime.Apply(Get_sieve(), gopurs_runtime.Apply2(go__3_4, dummy_1_1, gopurs_runtime.Value{Type: 9, IntVal: 3777797863, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Nil{})})), gopurs_runtime.Int(0)))), gopurs_runtime.Value{})
})
}()
	})
	return act
}

type Data_Test_Primes_Nil struct {
	
}
func Is_Data_Test_Primes_Nil(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3777797863
}

type Data_Test_Primes_Cons struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}
func Is_Data_Test_Primes_Cons(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2390177629
}

func Call_range_(start_0_loop gopurs_runtime.Value, end_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var start_0 gopurs_runtime.Value = start_0_loop
_ = start_0
var end_1 gopurs_runtime.Value = end_1_loop
_ = end_1
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(curr_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var curr_3 gopurs_runtime.Value = curr_3_loop
_ = curr_3
var acc_4 gopurs_runtime.Value = acc_4_loop
_ = acc_4
var __t1 gopurs_runtime.Value
{
if curr_3.IntVal < start_0.IntVal {
__t1 = acc_4
goto end_branch_1
} else {

}
}
{
curr_3_loop = gopurs_runtime.Int(curr_3.IntVal - 1)
acc_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Cons{curr_3, acc_4})}
continue go__2_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply2(go__2_0, end_1, gopurs_runtime.Value{Type: 9, IntVal: 3777797863, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Nil{})})
}

func Call_filter(p_0_loop gopurs_runtime.Value, lst_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var lst_1 gopurs_runtime.Value = lst_1_loop
_ = lst_1
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 3777797863) {
var go__5_2 gopurs_runtime.Value
go__5_2 = gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_2:
for {
if false { continue go__5_2 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t3 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 3777797863) {
__t3 = v1_7
goto end_branch_3
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629) {
v_6_loop = (*Data_Test_Primes_Cons)(v_6.UnsafePtr).V1
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Cons{(*Data_Test_Primes_Cons)(v_6.UnsafePtr).V0, v1_7})}
continue go__5_2
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
__t1 = gopurs_runtime.Apply2(go__5_2, v1_4, gopurs_runtime.Value{Type: 9, IntVal: 3777797863, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Nil{})})
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629) {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Apply(p_0, (*Data_Test_Primes_Cons)(v_3.UnsafePtr).V0).IntVal != 0 {
v_3_loop = (*Data_Test_Primes_Cons)(v_3.UnsafePtr).V1
v1_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Cons{(*Data_Test_Primes_Cons)(v_3.UnsafePtr).V0, v1_4})}
continue go__2_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
v_3_loop = (*Data_Test_Primes_Cons)(v_3.UnsafePtr).V1
v1_4_loop = v1_4
continue go__2_0
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
return gopurs_runtime.Apply2(go__2_0, lst_1, gopurs_runtime.Value{Type: 9, IntVal: 3777797863, UnsafePtr: unsafe.Pointer(&Data_Test_Primes_Nil{})})
}



package Test_ListOps

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	unsafe "unsafe"
)

var cache_Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		cache_Nil = gopurs_runtime.Value{Type: 9, IntVal: 63553145, UnsafePtr: unsafe.Pointer(&Data_Test_ListOps_Nil{})}
	})
	return cache_Nil
}

var cache_Cons gopurs_runtime.Value
var once_Cons sync.Once
func Get_Cons() gopurs_runtime.Value {
	once_Cons.Do(func() {
		cache_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Data_Test_ListOps_Cons{value0, value1})}
})
})
	})
	return cache_Cons
}

var cache_range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		cache_range_ = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_range_(start_0_box, end_1_box)
})
	})
	return cache_range_
}

var cache_foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		cache_foldl = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_foldl
}

var cache_filterEvens gopurs_runtime.Value
var once_filterEvens sync.Once
func Get_filterEvens() gopurs_runtime.Value {
	once_filterEvens.Do(func() {
		cache_filterEvens = gopurs_runtime.Func(func(lst_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if (v_2.Type == 9 && v_2.IntVal == 63553145) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1127792131) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), (*Data_Test_ListOps_Cons)(v_2.UnsafePtr).V0, gopurs_runtime.Int(2)).IntVal) == (0) {
v_2_loop = (*Data_Test_ListOps_Cons)(v_2.UnsafePtr).V1
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Data_Test_ListOps_Cons{(*Data_Test_ListOps_Cons)(v_2.UnsafePtr).V0, v1_3})}
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
v_2_loop = (*Data_Test_ListOps_Cons)(v_2.UnsafePtr).V1
v1_3_loop = v1_3
continue go__1_0
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
return gopurs_runtime.Apply2(go__1_0, lst_0, gopurs_runtime.Value{Type: 9, IntVal: 63553145, UnsafePtr: unsafe.Pointer(&Data_Test_ListOps_Nil{})})
}()
})
	})
	return cache_filterEvens
}

var cache_sumEvens gopurs_runtime.Value
var once_sumEvens sync.Once
func Get_sumEvens() gopurs_runtime.Value {
	once_sumEvens.Do(func() {
		cache_sumEvens = gopurs_runtime.Func(func(n_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(curr_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var curr_2 gopurs_runtime.Value = curr_2_loop
_ = curr_2
var acc_3 gopurs_runtime.Value = acc_3_loop
_ = acc_3
var __t1 gopurs_runtime.Value
{
if (curr_2.IntVal) < (1) {
__t1 = acc_3
goto end_branch_1
} else {

}
}
{
curr_2_loop = gopurs_runtime.Int((curr_2.IntVal) - (1))
acc_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Data_Test_ListOps_Cons{curr_2, acc_3})}
continue go__1_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return Call_foldl(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_filterEvens(), gopurs_runtime.Apply2(go__1_0, n_0, gopurs_runtime.Value{Type: 9, IntVal: 63553145, UnsafePtr: unsafe.Pointer(&Data_Test_ListOps_Nil{})})))
}()
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
		cache_act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(900))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = dummy_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply(Get_sumEvens(), dummy_1_1))), gopurs_runtime.Value{})
})
}()
	})
	return cache_act
}

type Data_Test_ListOps_Nil struct {
	
}
func Is_Data_Test_ListOps_Nil(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 63553145
}

type Data_Test_ListOps_Cons struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}
func Is_Data_Test_ListOps_Cons(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1127792131
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
if (curr_3.IntVal) < (start_0.IntVal) {
__t1 = acc_4
goto end_branch_1
} else {

}
}
{
curr_3_loop = gopurs_runtime.Int((curr_3.IntVal) - (1))
acc_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(&Data_Test_ListOps_Cons{curr_3, acc_4})}
continue go__2_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply2(go__2_0, end_1, gopurs_runtime.Value{Type: 9, IntVal: 63553145, UnsafePtr: unsafe.Pointer(&Data_Test_ListOps_Nil{})})
}

func Call_foldl(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
foldl:
for {
if false { continue foldl }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 63553145) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1127792131) {
v_0_loop = v_0
v1_1_loop = gopurs_runtime.Apply2(v_0, v1_1, (*Data_Test_ListOps_Cons)(v2_2.UnsafePtr).V0)
v2_2_loop = (*Data_Test_ListOps_Cons)(v2_2.UnsafePtr).V1
continue foldl
__t0 = gopurs_runtime.Value{}
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



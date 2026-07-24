package Test_ListOps

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		Nil = gopurs_runtime.Constructor0("Nil")
	})
	return Nil
}

var Cons gopurs_runtime.Value
var once_Cons sync.Once
func Get_Cons() gopurs_runtime.Value {
	once_Cons.Do(func() {
		Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", value0, value1)
})
})
	})
	return Cons
}

var range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		range_ = gopurs_runtime.Func2(func(start_0 gopurs_runtime.Value, end_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
acc_4_loop = gopurs_runtime.Constructor2("Cons", curr_3, acc_4)
continue go__2_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply2(go__2_0, end_1, gopurs_runtime.Constructor0("Nil"))
})
	})
	return range_
}

var foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		foldl = gopurs_runtime.Func3(Call_foldl)
	})
	return foldl
}

var filterEvens gopurs_runtime.Value
var once_filterEvens sync.Once
func Get_filterEvens() gopurs_runtime.Value {
	once_filterEvens.Do(func() {
		filterEvens = gopurs_runtime.Func(func(lst_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(v_2.StrVal == "Nil").IntVal != 0 {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_2.StrVal == "Cons").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], gopurs_runtime.Int(2)).IntVal == 0 {
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]
v1_3_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], v1_3)
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]
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
return gopurs_runtime.Apply2(go__1_0, lst_0, gopurs_runtime.Constructor0("Nil"))
})
	})
	return filterEvens
}

var sumEvens gopurs_runtime.Value
var once_sumEvens sync.Once
func Get_sumEvens() gopurs_runtime.Value {
	once_sumEvens.Do(func() {
		sumEvens = gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
if curr_2.IntVal < 1 {
__t1 = acc_3
goto end_branch_1
} else {

}
}
{
curr_2_loop = gopurs_runtime.Int(curr_2.IntVal - 1)
acc_3_loop = gopurs_runtime.Constructor2("Cons", curr_2, acc_3)
continue go__1_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply3(Get_foldl(), pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_filterEvens(), gopurs_runtime.Apply2(go__1_0, n_0, gopurs_runtime.Constructor0("Nil"))))
})
	})
	return sumEvens
}

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("List Processing (900 elements):"))
	})
	return describe
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(900))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = dummy_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply(Get_sumEvens(), dummy_1_1))), gopurs_runtime.Value{})
})
}()
	})
	return act
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
if gopurs_runtime.Bool(v2_2_loop.StrVal == "Nil").IntVal != 0 {
__t0 = v1_1_loop
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_2_loop.StrVal == "Cons").IntVal != 0 {
__t0 = gopurs_runtime.Apply3(Get_foldl(), v_0_loop, gopurs_runtime.Apply2(v_0_loop, v1_1_loop, (*[1024]gopurs_runtime.Value)(v2_2_loop.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v2_2_loop.UnsafePtr)[1])
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



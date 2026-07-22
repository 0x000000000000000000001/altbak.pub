package Test_Primes

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		Nil = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
	})
	return Nil
}

var Cons gopurs_runtime.Value
var once_Cons sync.Once
func Get_Cons() gopurs_runtime.Value {
	once_Cons.Do(func() {
		Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": value0, "value1": value1})
})
})
	})
	return Cons
}

var sumList gopurs_runtime.Value
var once_sumList sync.Once
func Get_sumList() gopurs_runtime.Value {
	once_sumList.Do(func() {
		sumList = gopurs_runtime.Func(func(lst_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var v1_3 = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
v1_3_loop = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), v1_3), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
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
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, lst_0), gopurs_runtime.Int(0))
})
	})
	return sumList
}

var reverse gopurs_runtime.Value
var once_reverse sync.Once
func Get_reverse() gopurs_runtime.Value {
	once_reverse.Do(func() {
		reverse = gopurs_runtime.Func(func(lst_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var v1_3 = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
v_2_loop = v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
v1_3_loop = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_3})
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
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__1_0, lst_0), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
})
	})
	return reverse
}

var range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		range_ = gopurs_runtime.Func(func(start_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(end_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(curr_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var curr_3 = curr_3_loop
_ = curr_3
var acc_4 = acc_4_loop
_ = acc_4
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], curr_3), start_0).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = acc_4
goto end_branch_1
} else {

}
}
{
curr_3_loop = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), curr_3), gopurs_runtime.Int(1))
acc_4_loop = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": curr_3, "value1": acc_4})
continue go__2_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, end_1), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
})
})
	})
	return range_
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lst_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var v_3 = v_3_loop
_ = v_3
var v1_4 = v1_4_loop
_ = v1_4
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
var go__5_2 gopurs_runtime.Value
go__5_2 = gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_2:
for {
if false { continue go__5_2 }
var v_6 = v_6_loop
_ = v_6
var v1_7 = v1_7_loop
_ = v1_7
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t3 = v1_7
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
v_6_loop = v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
v1_7_loop = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_7})
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
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__5_2, v1_4), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])).IntVal != 0 {
v_3_loop = v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
v1_4_loop = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_4})
continue go__2_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
v_3_loop = v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
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
return gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_0, lst_1), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
})
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
var v_0 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__local_var_1_1 := v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
var go__2_2 gopurs_runtime.Value
go__2_2 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_2:
for {
if false { continue go__2_2 }
var v_3 = v_3_loop
_ = v_3
var v1_4 = v1_4_loop
_ = v1_4
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
var go__5_4 gopurs_runtime.Value
go__5_4 = gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_4:
for {
if false { continue go__5_4 }
var v_6 = v_6_loop
_ = v_6
var v1_7 = v1_7_loop
_ = v1_7
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t5 = v1_7
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(v_6.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
v_6_loop = v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
v1_7_loop = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_7})
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
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(go__5_4, v1_4), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")}))
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqBooleanImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_intMod(), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), __local_var_1_1)), gopurs_runtime.Int(0))), gopurs_runtime.Bool(false))).IntVal != 0 {
v_3_loop = v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
v1_4_loop = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_4})
continue go__2_2
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
v_3_loop = v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
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
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Cons"), "value0": __local_var_1_1, "value1": gopurs_runtime.Apply(Get_sieve(), gopurs_runtime.Apply(gopurs_runtime.Apply(go__2_2, v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nil")})))})
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
		act = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply(Get_sumList(), gopurs_runtime.Apply(Get_sieve(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_range_(), gopurs_runtime.Int(2)), gopurs_runtime.Int(500))))))
	})
	return act
}



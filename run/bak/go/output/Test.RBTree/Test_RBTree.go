package Test_RBTree

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var R gopurs_runtime.Value
var once_R sync.Once
func Get_R() gopurs_runtime.Value {
	once_R.Do(func() {
		R = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")})
	})
	return R
}

var B gopurs_runtime.Value
var once_B sync.Once
func Get_B() gopurs_runtime.Value {
	once_B.Do(func() {
		B = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")})
	})
	return B
}

var E gopurs_runtime.Value
var once_E sync.Once
func Get_E() gopurs_runtime.Value {
	once_E.Do(func() {
		E = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("E")})
	})
	return E
}

var T gopurs_runtime.Value
var once_T sync.Once
func Get_T() gopurs_runtime.Value {
	once_T.Do(func() {
		T = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": value0, "value1": value1, "value2": value2, "value3": value3})
})
})
})
})
	})
	return T
}

var max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		max = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0), y_1).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t0 = x_0
goto end_branch_0
} else {

}
}
{
__t0 = y_1
}
end_branch_0:
return __t0
})
})
	})
	return max
}

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Red-Black Tree (100k Worst-Case Insertions):"))
	})
	return describe
}

var depth gopurs_runtime.Value
var once_depth sync.Once
func Get_depth() gopurs_runtime.Value {
	once_depth.Do(func() {
		depth = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
depth:
for {
if false { continue depth }
var v_0 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "E")).IntVal != 0 {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T")).IntVal != 0 {
__local_var_1_1 := gopurs_runtime.Apply(Get_depth(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
__local_var_2_2 := gopurs_runtime.Apply(Get_depth(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value3"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], __local_var_1_1), __local_var_2_2).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = __local_var_1_1
goto end_branch_3
} else {

}
}
{
__t3 = __local_var_2_2
}
end_branch_3:
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(1)), __t3)
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
	return depth
}

var balance gopurs_runtime.Value
var once_balance sync.Once
func Get_balance() gopurs_runtime.Value {
	once_balance.Do(func() {
		balance = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "B")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T")).IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]}), "value2": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": v2_2, "value3": v3_3})})
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T")).IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": v2_2, "value3": v3_3})})
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T")).IntVal != 0 {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R")).IntVal != 0 {
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_7:
__t6 = __t7
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T")).IntVal != 0 {
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R")).IntVal != 0 {
__t9 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_9
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
__t9 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_9:
__t8 = __t9
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_8:
__t4 = __t8
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T")).IntVal != 0 {
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R")).IntVal != 0 {
__t10 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": v2_2, "value3": v3_3})})
goto end_branch_10
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T")).IntVal != 0 {
var __t12 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R")).IntVal != 0 {
__t12 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_12
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
__t12 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_12:
__t11 = __t12
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_11:
__t10 = __t11
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_10:
__t3 = __t10
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
var __t13 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T")).IntVal != 0 {
var __t14 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R")).IntVal != 0 {
__t14 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_14
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
__t14 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_14:
__t13 = __t14
goto end_branch_13
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
__t13 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_13:
__t3 = __t13
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
var __t15 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T")).IntVal != 0 {
var __t16 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R")).IntVal != 0 {
__t16 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_16
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
__t16 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_16:
__t15 = __t16
goto end_branch_15
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
__t15 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_15:
__t2 = __t15
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
var __t17 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T")).IntVal != 0 {
var __t18 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R")).IntVal != 0 {
__t18 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value3"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_18
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
__t18 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_18:
__t17 = __t18
goto end_branch_17
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T").IntVal != 0 && gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "R").IntVal != 0)).IntVal != 0 {
__t17 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v1_1, "value2": v2_2, "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})})
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_17:
__t1 = __t17
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_0, "value1": v1_1, "value2": v2_2, "value3": v3_3})
}
end_branch_0:
return __t0
})
})
})
})
	})
	return balance
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
var ins_2_0 gopurs_runtime.Value
ins_2_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "E")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("R")}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("E")}), "value2": x_0, "value3": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("E")})})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_balance(), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), gopurs_runtime.Apply(ins_2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"])
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_balance(), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), gopurs_runtime.Apply(ins_2_0, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
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
})
__local_var_3_3 := gopurs_runtime.Apply(ins_2_0, s_1)
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "T")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("T"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("B")}), "value1": __local_var_3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": __local_var_3_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": __local_var_3_3.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "E")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("E")})
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
})
	})
	return insert
}

var buildTree gopurs_runtime.Value
var once_buildTree sync.Once
func Get_buildTree() gopurs_runtime.Value {
	once_buildTree.Do(func() {
		buildTree = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
buildTree:
for {
if false { continue buildTree }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(Get_buildTree(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_intSub(), v_0), gopurs_runtime.Int(1))), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_insert(), v_0), v1_1))
}
end_branch_0:
return __t0
}
}()
})
})
	})
	return buildTree
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply(Get_depth(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_buildTree(), gopurs_runtime.Int(100000)), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("E")})))))
	})
	return act
}



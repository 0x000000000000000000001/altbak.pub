package Data_Enum

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var guard gopurs_runtime.Value
var once_guard sync.Once
func Get_guard() gopurs_runtime.Value {
	once_guard.Do(func() {
		guard = gopurs_runtime.Apply(pkg_Control_Alternative.Get_guard(), pkg_Data_Maybe.Get_alternativeMaybe())
	})
	return guard
}

var Cardinality gopurs_runtime.Value
var once_Cardinality sync.Once
func Get_Cardinality() gopurs_runtime.Value {
	once_Cardinality.Do(func() {
		Cardinality = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Cardinality
}

var toEnum gopurs_runtime.Value
var once_toEnum sync.Once
func Get_toEnum() gopurs_runtime.Value {
	once_toEnum.Do(func() {
		toEnum = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "toEnum")
})
	})
	return toEnum
}

var succ gopurs_runtime.Value
var once_succ sync.Once
func Get_succ() gopurs_runtime.Value {
	once_succ.Do(func() {
		succ = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "succ")
})
	})
	return succ
}

var upFromIncluding gopurs_runtime.Value
var once_upFromIncluding sync.Once
func Get_upFromIncluding() gopurs_runtime.Value {
	once_upFromIncluding.Do(func() {
		upFromIncluding = gopurs_runtime.Func2(func(dictEnum_0 gopurs_runtime.Value, dictUnfoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable1_1, "unfoldr1"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", x_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), x_2))
}))
})
	})
	return upFromIncluding
}

var showCardinality gopurs_runtime.Value
var once_showCardinality sync.Once
func Get_showCardinality() gopurs_runtime.Value {
	once_showCardinality.Do(func() {
		showCardinality = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Cardinality ").StrVal + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
	})
	return showCardinality
}

var pred gopurs_runtime.Value
var once_pred sync.Once
func Get_pred() gopurs_runtime.Value {
	once_pred.Do(func() {
		pred = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "pred")
})
	})
	return pred
}

var ordCardinality gopurs_runtime.Value
var once_ordCardinality sync.Once
func Get_ordCardinality() gopurs_runtime.Value {
	once_ordCardinality.Do(func() {
		ordCardinality = pkg_Data_Ord.Get_ordInt()
	})
	return ordCardinality
}

var newtypeCardinality gopurs_runtime.Value
var once_newtypeCardinality sync.Once
func Get_newtypeCardinality() gopurs_runtime.Value {
	once_newtypeCardinality.Do(func() {
		newtypeCardinality = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeCardinality
}

var fromEnum gopurs_runtime.Value
var once_fromEnum sync.Once
func Get_fromEnum() gopurs_runtime.Value {
	once_fromEnum.Do(func() {
		fromEnum = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "fromEnum")
})
	})
	return fromEnum
}

var toEnumWithDefaults gopurs_runtime.Value
var once_toEnumWithDefaults sync.Once
func Get_toEnumWithDefaults() gopurs_runtime.Value {
	once_toEnumWithDefaults.Do(func() {
		toEnumWithDefaults = gopurs_runtime.Func(func(dictBoundedEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func3(func(low_2 gopurs_runtime.Value, high_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "toEnum"), x_4)
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_5_1.StrVal == "Just")).IntVal != 0 {
__t2 = (*[1024]gopurs_runtime.Value)(v_5_1.UnsafePtr)[0]
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_5_1.StrVal == "Nothing")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_4.IntVal < gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "fromEnum"), bottom2_1_0).IntVal)).IntVal != 0 {
__t3 = low_2
goto end_branch_3
} else {

}
}
{
__t3 = high_3
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
})
	})
	return toEnumWithDefaults
}

var eqCardinality gopurs_runtime.Value
var once_eqCardinality sync.Once
func Get_eqCardinality() gopurs_runtime.Value {
	once_eqCardinality.Do(func() {
		eqCardinality = pkg_Data_Eq.Get_eqInt()
	})
	return eqCardinality
}

var enumUnit gopurs_runtime.Value
var once_enumUnit sync.Once
func Get_enumUnit() gopurs_runtime.Value {
	once_enumUnit.Do(func() {
		enumUnit = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordUnit()
}))
	})
	return enumUnit
}

var enumTuple gopurs_runtime.Value
var once_enumTuple sync.Once
func Get_enumTuple() gopurs_runtime.Value {
	once_enumTuple.Do(func() {
		enumTuple = gopurs_runtime.Func(func(dictEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictBoundedEnum_3 gopurs_runtime.Value) gopurs_runtime.Value {
Bounded0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_3, "Bounded0"), gopurs_runtime.Value{})
_ = Bounded0_4_2
bottom2_5_3 := gopurs_runtime.RecordGet(Bounded0_4_2, "bottom")
_ = bottom2_5_3
Enum1_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_3, "Enum1"), gopurs_runtime.Value{})
_ = Enum1_6_4
top2_7_5 := gopurs_runtime.RecordGet(Bounded0_4_2, "top")
_ = top2_7_5
__local_var_8_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_6_4, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_8_6
__local_var_9_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_6, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_9_7
eqTuple2_10_9 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_10 gopurs_runtime.Value, y_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "eq"), (*[1024]gopurs_runtime.Value)(x_10.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_11.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_7, "eq"), (*[1024]gopurs_runtime.Value)(x_10.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_11.UnsafePtr)[1]).IntVal != 0)
}))
_ = eqTuple2_10_9
ordTuple1_10_8 := gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_11 gopurs_runtime.Value, y_12 gopurs_runtime.Value) gopurs_runtime.Value {
v_13_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "compare"), (*[1024]gopurs_runtime.Value)(x_11.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_12.UnsafePtr)[0])
_ = v_13_10
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_13_10.StrVal == "LT")).IntVal != 0 {
__t11 = gopurs_runtime.Constructor0("LT")
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(v_13_10.StrVal == "GT")).IntVal != 0 {
__t11 = gopurs_runtime.Constructor0("GT")
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_6, "compare"), (*[1024]gopurs_runtime.Value)(x_11.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_12.UnsafePtr)[1])
}
end_branch_11:
return __t11
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return eqTuple2_10_9
}))
_ = ordTuple1_10_8
return gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_12_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), (*[1024]gopurs_runtime.Value)(v_11.UnsafePtr)[0])
_ = __local_var_12_12
var __t14 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_12_12.StrVal == "Just")).IntVal != 0 {
__t14 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(__local_var_12_12.UnsafePtr)[0], bottom2_5_3))
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_14:
__local_var_13_13 := __t14
_ = __local_var_13_13
__local_var_14_15 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*[1024]gopurs_runtime.Value)(v_11.UnsafePtr)[0])
_ = __local_var_14_15
__local_var_15_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_6_4, "succ"), (*[1024]gopurs_runtime.Value)(v_11.UnsafePtr)[1])
_ = __local_var_15_16
var __t17 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_15_16.StrVal == "Nothing")).IntVal != 0 {
__t17 = __local_var_13_13
goto end_branch_17
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_15_16.StrVal == "Just")).IntVal != 0 {
__t17 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(__local_var_14_15, (*[1024]gopurs_runtime.Value)(__local_var_15_16.UnsafePtr)[0]))
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_12_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), (*[1024]gopurs_runtime.Value)(v_11.UnsafePtr)[0])
_ = __local_var_12_18
var __t20 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_12_18.StrVal == "Just")).IntVal != 0 {
__t20 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(__local_var_12_18.UnsafePtr)[0], top2_7_5))
goto end_branch_20
} else {

}
}
{
__t20 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_20:
__local_var_13_19 := __t20
_ = __local_var_13_19
__local_var_14_21 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*[1024]gopurs_runtime.Value)(v_11.UnsafePtr)[0])
_ = __local_var_14_21
__local_var_15_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_6_4, "pred"), (*[1024]gopurs_runtime.Value)(v_11.UnsafePtr)[1])
_ = __local_var_15_22
var __t23 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_15_22.StrVal == "Nothing")).IntVal != 0 {
__t23 = __local_var_13_19
goto end_branch_23
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_15_22.StrVal == "Just")).IntVal != 0 {
__t23 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(__local_var_14_21, (*[1024]gopurs_runtime.Value)(__local_var_15_22.UnsafePtr)[0]))
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return ordTuple1_10_8
}))
})
})
	})
	return enumTuple
}

var enumOrdering gopurs_runtime.Value
var once_enumOrdering sync.Once
func Get_enumOrdering() gopurs_runtime.Value {
	once_enumOrdering.Do(func() {
		enumOrdering = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.StrVal == "LT")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("EQ"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "EQ")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("GT"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "GT")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.StrVal == "LT")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "EQ")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("LT"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "GT")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("EQ"))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordOrdering()
}))
	})
	return enumOrdering
}

var enumMaybe gopurs_runtime.Value
var once_enumMaybe sync.Once
func Get_enumMaybe() gopurs_runtime.Value {
	once_enumMaybe.Do(func() {
		enumMaybe = gopurs_runtime.Func(func(dictBoundedEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
Enum1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{})
_ = Enum1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_2_1, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_4_3
eqMaybe1_5_5 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_5.StrVal == "Nothing")).IntVal != 0 {
__t6 = gopurs_runtime.Bool(y_6.StrVal == "Nothing")
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_5.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(y_6.StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "eq"), (*[1024]gopurs_runtime.Value)(x_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0]).IntVal != 0).IntVal != 0)
}
end_branch_6:
return __t6
}))
_ = eqMaybe1_5_5
ordMaybe_5_4 := gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, y_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_6.StrVal == "Nothing")).IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_7.StrVal == "Nothing")).IntVal != 0 {
__t8 = gopurs_runtime.Constructor0("EQ")
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Constructor0("LT")
}
end_branch_8:
__t7 = __t8
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(y_7.StrVal == "Nothing")).IntVal != 0 {
__t7 = gopurs_runtime.Constructor0("GT")
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_6.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_7.StrVal == "Just").IntVal != 0)).IntVal != 0 {
__t7 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "compare"), (*[1024]gopurs_runtime.Value)(x_6.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_7.UnsafePtr)[0])
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_5_5
}))
_ = ordMaybe_5_4
return gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_6.StrVal == "Nothing")).IntVal != 0 {
__t9 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Just", bottom2_1_0))
goto end_branch_9
} else {

}
}
{
if (gopurs_runtime.Bool(v_6.StrVal == "Just")).IntVal != 0 {
__local_var_7_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_2_1, "succ"), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0])
_ = __local_var_7_10
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_10.StrVal == "Just")).IntVal != 0 {
__t11 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(__local_var_7_10.UnsafePtr)[0]))
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_11:
__t9 = __t11
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_6.StrVal == "Nothing")).IntVal != 0 {
__t12 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_12
} else {

}
}
{
if (gopurs_runtime.Bool(v_6.StrVal == "Just")).IntVal != 0 {
__t12 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_2_1, "pred"), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0]))
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return ordMaybe_5_4
}))
})
	})
	return enumMaybe
}

var enumInt gopurs_runtime.Value
var once_enumInt sync.Once
func Get_enumInt() gopurs_runtime.Value {
	once_enumInt.Do(func() {
		enumInt = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(n_0.IntVal < pkg_Data_Bounded.Get_topInt().IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Int(n_0.IntVal + gopurs_runtime.Int(1).IntVal))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(n_0.IntVal > pkg_Data_Bounded.Get_bottomInt().IntVal)).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Int(n_0.IntVal - gopurs_runtime.Int(1).IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordInt()
}))
	})
	return enumInt
}

var enumFromTo gopurs_runtime.Value
var once_enumFromTo sync.Once
func Get_enumFromTo() gopurs_runtime.Value {
	once_enumFromTo.Do(func() {
		enumFromTo = gopurs_runtime.Func(func(dictEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
Ord0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "Ord0"), gopurs_runtime.Value{})
_ = Ord0_1_0
return gopurs_runtime.Func3(func(dictUnfoldable1_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_1_0, "Eq0"), gopurs_runtime.Value{}), "eq"), v_3, v1_4)).IntVal != 0 {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(i_5.IntVal <= gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor2("Tuple", v_3, gopurs_runtime.Constructor0("Nothing"))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Constructor2("Tuple", v_3, gopurs_runtime.Constructor1("Just", gopurs_runtime.Int(i_5.IntVal - gopurs_runtime.Int(1).IntVal)))
}
end_branch_5:
return __t5
}), gopurs_runtime.Int(0))
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), v_3, v1_4).StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), a_5)
_ = __local_var_6_6
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_6.StrVal == "Just")).IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply(Get_guard(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), (*[1024]gopurs_runtime.Value)(__local_var_6_6.UnsafePtr)[0], v1_4).StrVal == "GT").IntVal == 0)).StrVal == "Just")).IntVal != 0 {
__t8 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(__local_var_6_6.UnsafePtr)[0])
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_8:
__t7 = __t8
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_6_6.StrVal == "Nothing")).IntVal != 0 {
__t7 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Constructor2("Tuple", a_5, __t7)
}), v_3)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_2, "unfoldr1"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), a_5)
_ = __local_var_6_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_1.StrVal == "Just")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply(Get_guard(), gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), (*[1024]gopurs_runtime.Value)(__local_var_6_1.UnsafePtr)[0], v1_4).StrVal == "LT").IntVal == 0)).StrVal == "Just")).IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(__local_var_6_1.UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_6_1.StrVal == "Nothing")).IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Constructor2("Tuple", a_5, __t2)
}), v_3)
}
end_branch_4:
return __t4
})
})
	})
	return enumFromTo
}

var enumFromThenTo gopurs_runtime.Value
var once_enumFromThenTo sync.Once
func Get_enumFromThenTo() gopurs_runtime.Value {
	once_enumFromThenTo.Do(func() {
		enumFromThenTo = gopurs_runtime.Func5(func(dictUnfoldable_0 gopurs_runtime.Value, dictFunctor_1 gopurs_runtime.Value, dictBoundedEnum_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_6_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "fromEnum"), a_3)
_ = a_prime_6_0
__local_var_7_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "fromEnum"), b_4).IntVal - a_prime_6_0.IntVal
_ = __local_var_7_3
__local_var_8_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "fromEnum"), c_5)
_ = __local_var_8_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_1, "map"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "toEnum"), x_7)
_ = __local_var_8_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_8_1.StrVal == "Just")).IntVal != 0 {
__t2 = (*[1024]gopurs_runtime.Value)(__local_var_8_1.UnsafePtr)[0]
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(e_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(e_9.IntVal <= __local_var_8_4.IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", e_9, gopurs_runtime.Int(e_9.IntVal + __local_var_7_3.IntVal)))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_5:
return __t5
}), a_prime_6_0))
})
})
	})
	return enumFromThenTo
}

var enumEither gopurs_runtime.Value
var once_enumEither sync.Once
func Get_enumEither() gopurs_runtime.Value {
	once_enumEither.Do(func() {
		enumEither = gopurs_runtime.Func(func(dictBoundedEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
Enum1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Enum1"), gopurs_runtime.Value{})
_ = Enum1_1_0
top2_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_0, "Bounded0"), gopurs_runtime.Value{}), "top")
_ = top2_2_1
ordEither_3_2 := gopurs_runtime.Apply(pkg_Data_Either.Get_ordEither(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_1_0, "Ord0"), gopurs_runtime.Value{}))
_ = ordEither_3_2
return gopurs_runtime.Func(func(dictBoundedEnum1_4 gopurs_runtime.Value) gopurs_runtime.Value {
bottom2_5_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_5_3
Enum11_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum1_4, "Enum1"), gopurs_runtime.Value{})
_ = Enum11_6_4
ordEither1_7_5 := gopurs_runtime.Apply(ordEither_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum11_6_4, "Ord0"), gopurs_runtime.Value{}))
_ = ordEither1_7_5
return gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_8.StrVal == "Left")).IntVal != 0 {
__local_var_9_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_1_0, "succ"), (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0])
_ = __local_var_9_7
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_7.StrVal == "Nothing")).IntVal != 0 {
__t8 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Right", bottom2_5_3))
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_9_7.StrVal == "Just")).IntVal != 0 {
__t8 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(__local_var_9_7.UnsafePtr)[0]))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t6 = __t8
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(v_8.StrVal == "Right")).IntVal != 0 {
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum11_6_4, "succ"), (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0])
_ = __local_var_9_9
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_9.StrVal == "Nothing")).IntVal != 0 {
__t10 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_10
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_9_9.StrVal == "Just")).IntVal != 0 {
__t10 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)(__local_var_9_9.UnsafePtr)[0]))
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t6 = __t10
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_8.StrVal == "Left")).IntVal != 0 {
__local_var_9_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_1_0, "pred"), (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0])
_ = __local_var_9_12
var __t13 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_12.StrVal == "Nothing")).IntVal != 0 {
__t13 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_13
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_9_12.StrVal == "Just")).IntVal != 0 {
__t13 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(__local_var_9_12.UnsafePtr)[0]))
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t11 = __t13
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(v_8.StrVal == "Right")).IntVal != 0 {
__local_var_9_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum11_6_4, "pred"), (*[1024]gopurs_runtime.Value)(v_8.UnsafePtr)[0])
_ = __local_var_9_14
var __t15 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_14.StrVal == "Nothing")).IntVal != 0 {
__t15 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Left", top2_2_1))
goto end_branch_15
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_9_14.StrVal == "Just")).IntVal != 0 {
__t15 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)(__local_var_9_14.UnsafePtr)[0]))
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
__t11 = __t15
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return ordEither1_7_5
}))
})
})
	})
	return enumEither
}

var enumBoolean gopurs_runtime.Value
var once_enumBoolean sync.Once
func Get_enumBoolean() gopurs_runtime.Value {
	once_enumBoolean.Do(func() {
		enumBoolean = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == 0)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Bool(true))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_0).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Bool(false))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordBoolean()
}))
	})
	return enumBoolean
}

var downFromIncluding gopurs_runtime.Value
var once_downFromIncluding sync.Once
func Get_downFromIncluding() gopurs_runtime.Value {
	once_downFromIncluding.Do(func() {
		downFromIncluding = gopurs_runtime.Func2(func(dictEnum_0 gopurs_runtime.Value, dictUnfoldable1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable1_1, "unfoldr1"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", x_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), x_2))
}))
})
	})
	return downFromIncluding
}

var downFrom gopurs_runtime.Value
var once_downFrom sync.Once
func Get_downFrom() gopurs_runtime.Value {
	once_downFrom.Do(func() {
		downFrom = gopurs_runtime.Func2(func(dictEnum_0 gopurs_runtime.Value, dictUnfoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_1, "unfoldr"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), x_2)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}))
})
	})
	return downFrom
}

var upFrom gopurs_runtime.Value
var once_upFrom sync.Once
func Get_upFrom() gopurs_runtime.Value {
	once_upFrom.Do(func() {
		upFrom = gopurs_runtime.Func2(func(dictEnum_0 gopurs_runtime.Value, dictUnfoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_1, "unfoldr"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "succ"), x_2)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}))
})
	})
	return upFrom
}

var defaultToEnum gopurs_runtime.Value
var once_defaultToEnum sync.Once
func Get_defaultToEnum() gopurs_runtime.Value {
	once_defaultToEnum.Do(func() {
		defaultToEnum = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
bottom2_1_0 := gopurs_runtime.RecordGet(dictBounded_0, "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func2(func(dictEnum_2 gopurs_runtime.Value, i_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_1 gopurs_runtime.Value
go__4_1 = gopurs_runtime.Func(func(i_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_1:
for {
if false { continue go__4_1 }
var i_5 = i_5_loop
_ = i_5
var x_6 = x_6_loop
_ = x_6
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(i_5.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t4 = gopurs_runtime.Constructor1("Just", x_6)
goto end_branch_4
} else {

}
}
{
v_7_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_2, "succ"), x_6)
_ = v_7_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_7_2.StrVal == "Just")).IntVal != 0 {
i_5_loop = gopurs_runtime.Int(i_5.IntVal - gopurs_runtime.Int(1).IntVal)
x_6_loop = (*[1024]gopurs_runtime.Value)(v_7_2.UnsafePtr)[0]
continue go__4_1
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v_7_2.StrVal == "Nothing")).IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = __t3
}
end_branch_4:
return __t4
}
}()
})
})
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(i_prime_3.IntVal < gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply2(go__4_1, i_prime_3, bottom2_1_0)
}
end_branch_5:
return __t5
})
})
	})
	return defaultToEnum
}

var defaultSucc gopurs_runtime.Value
var once_defaultSucc sync.Once
func Get_defaultSucc() gopurs_runtime.Value {
	once_defaultSucc.Do(func() {
		defaultSucc = gopurs_runtime.Func3(func(toEnum_prime_0 gopurs_runtime.Value, fromEnum_prime_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int(gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal + gopurs_runtime.Int(1).IntVal))
})
	})
	return defaultSucc
}

var defaultPred gopurs_runtime.Value
var once_defaultPred sync.Once
func Get_defaultPred() gopurs_runtime.Value {
	once_defaultPred.Do(func() {
		defaultPred = gopurs_runtime.Func3(func(toEnum_prime_0 gopurs_runtime.Value, fromEnum_prime_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int(gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal - gopurs_runtime.Int(1).IntVal))
})
	})
	return defaultPred
}

var defaultFromEnum gopurs_runtime.Value
var once_defaultFromEnum sync.Once
func Get_defaultFromEnum() gopurs_runtime.Value {
	once_defaultFromEnum.Do(func() {
		defaultFromEnum = gopurs_runtime.Func(func(dictEnum_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(i_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var i_2 = i_2_loop
_ = i_2
var x_3 = x_3_loop
_ = x_3
v_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_0, "pred"), x_3)
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4_1.StrVal == "Just")).IntVal != 0 {
i_2_loop = gopurs_runtime.Int(i_2.IntVal + gopurs_runtime.Int(1).IntVal)
x_3_loop = (*[1024]gopurs_runtime.Value)(v_4_1.UnsafePtr)[0]
continue go__1_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v_4_1.StrVal == "Nothing")).IntVal != 0 {
__t2 = i_2
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Int(0))
})
	})
	return defaultFromEnum
}

var defaultCardinality gopurs_runtime.Value
var once_defaultCardinality sync.Once
func Get_defaultCardinality() gopurs_runtime.Value {
	once_defaultCardinality.Do(func() {
		defaultCardinality = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
bottom2_1_0 := gopurs_runtime.RecordGet(dictBounded_0, "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func(func(dictEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_1 gopurs_runtime.Value
go__3_1 = gopurs_runtime.Func(func(i_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_1:
for {
if false { continue go__3_1 }
var i_4 = i_4_loop
_ = i_4
var x_5 = x_5_loop
_ = x_5
v_6_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEnum_2, "succ"), x_5)
_ = v_6_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_6_2.StrVal == "Just")).IntVal != 0 {
i_4_loop = gopurs_runtime.Int(i_4.IntVal + gopurs_runtime.Int(1).IntVal)
x_5_loop = (*[1024]gopurs_runtime.Value)(v_6_2.UnsafePtr)[0]
continue go__3_1
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v_6_2.StrVal == "Nothing")).IntVal != 0 {
__t3 = i_4
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
return gopurs_runtime.Apply2(go__3_1, gopurs_runtime.Int(1), bottom2_1_0)
})
})
	})
	return defaultCardinality
}

var charToEnum gopurs_runtime.Value
var once_charToEnum sync.Once
func Get_charToEnum() gopurs_runtime.Value {
	once_charToEnum.Do(func() {
		charToEnum = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_0.IntVal >= gopurs_runtime.Apply(Get_toCharCode(), pkg_Data_Bounded.Get_bottomChar()).IntVal).IntVal != 0 && gopurs_runtime.Bool(v_0.IntVal <= gopurs_runtime.Apply(Get_toCharCode(), pkg_Data_Bounded.Get_topChar()).IntVal).IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(Get_fromCharCode(), v_0))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
})
	})
	return charToEnum
}

var enumChar gopurs_runtime.Value
var once_enumChar sync.Once
func Get_enumChar() gopurs_runtime.Value {
	once_enumChar.Do(func() {
		enumChar = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_charToEnum(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_toCharCode(), a_0).IntVal + gopurs_runtime.Int(1).IntVal))
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_charToEnum(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_toCharCode(), a_0).IntVal - gopurs_runtime.Int(1).IntVal))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordChar()
}))
	})
	return enumChar
}

var cardinality gopurs_runtime.Value
var once_cardinality sync.Once
func Get_cardinality() gopurs_runtime.Value {
	once_cardinality.Do(func() {
		cardinality = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "cardinality")
})
	})
	return cardinality
}

var boundedEnumUnit gopurs_runtime.Value
var once_boundedEnumUnit sync.Once
func Get_boundedEnumUnit() gopurs_runtime.Value {
	once_boundedEnumUnit.Do(func() {
		boundedEnumUnit = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(1), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", pkg_Data_Unit.Get_unit())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bounded.Get_boundedUnit()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumUnit()
}))
	})
	return boundedEnumUnit
}

var boundedEnumOrdering gopurs_runtime.Value
var once_boundedEnumOrdering sync.Once
func Get_boundedEnumOrdering() gopurs_runtime.Value {
	once_boundedEnumOrdering.Do(func() {
		boundedEnumOrdering = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(3), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("LT"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("EQ"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(2).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor0("GT"))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.StrVal == "LT")).IntVal != 0 {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "EQ")).IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.StrVal == "GT")).IntVal != 0 {
__t1 = gopurs_runtime.Int(2)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bounded.Get_boundedOrdering()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumOrdering()
}))
	})
	return boundedEnumOrdering
}

var boundedEnumChar gopurs_runtime.Value
var once_boundedEnumChar sync.Once
func Get_boundedEnumChar() gopurs_runtime.Value {
	once_boundedEnumChar.Do(func() {
		boundedEnumChar = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(gopurs_runtime.Apply(Get_toCharCode(), pkg_Data_Bounded.Get_topChar()).IntVal - gopurs_runtime.Apply(Get_toCharCode(), pkg_Data_Bounded.Get_bottomChar()).IntVal), Get_charToEnum(), Get_toCharCode(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bounded.Get_boundedChar()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumChar()
}))
	})
	return boundedEnumChar
}

var boundedEnumBoolean gopurs_runtime.Value
var once_boundedEnumBoolean sync.Once
func Get_boundedEnumBoolean() gopurs_runtime.Value {
	once_boundedEnumBoolean.Do(func() {
		boundedEnumBoolean = gopurs_runtime.RecordDict5("cardinality", "toEnum", "fromEnum", "Bounded0", "Enum1", gopurs_runtime.Int(2), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Bool(false))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Bool(true))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == 0)).IntVal != 0 {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (v_0).IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bounded.Get_boundedBoolean()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumBoolean()
}))
	})
	return boundedEnumBoolean
}

func Get_fromCharCode() gopurs_runtime.Value {
	return _Gopurs_FromCharCode
}

func Get_toCharCode() gopurs_runtime.Value {
	return _Gopurs_ToCharCode
}

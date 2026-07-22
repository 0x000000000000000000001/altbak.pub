package Data_Interval

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Bifoldable "gopurs/output/Data.Bifoldable"
)

var show gopurs_runtime.Value
var once_show sync.Once
func Get_show() gopurs_runtime.Value {
	once_show.Do(func() {
		show = gopurs_runtime.Apply(pkg_Data_Maybe.Get_showMaybe(), pkg_Data_Show.Get_showInt()).PtrVal.(map[string]gopurs_runtime.Value)["show"]
	})
	return show
}

var eq gopurs_runtime.Value
var once_eq sync.Once
func Get_eq() gopurs_runtime.Value {
	once_eq.Do(func() {
		eq = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
end_branch_0:
return __t0
})
})
	})
	return eq
}

var compare gopurs_runtime.Value
var once_compare sync.Once
func Get_compare() gopurs_runtime.Value {
	once_compare.Do(func() {
		compare = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("EQ")})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just").IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})
	})
	return compare
}

var StartEnd gopurs_runtime.Value
var once_StartEnd sync.Once
func Get_StartEnd() gopurs_runtime.Value {
	once_StartEnd.Do(func() {
		StartEnd = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartEnd"), "value0": value0, "value1": value1})
})
})
	})
	return StartEnd
}

var DurationEnd gopurs_runtime.Value
var once_DurationEnd sync.Once
func Get_DurationEnd() gopurs_runtime.Value {
	once_DurationEnd.Do(func() {
		DurationEnd = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationEnd"), "value0": value0, "value1": value1})
})
})
	})
	return DurationEnd
}

var StartDuration gopurs_runtime.Value
var once_StartDuration sync.Once
func Get_StartDuration() gopurs_runtime.Value {
	once_StartDuration.Do(func() {
		StartDuration = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartDuration"), "value0": value0, "value1": value1})
})
})
	})
	return StartDuration
}

var DurationOnly gopurs_runtime.Value
var once_DurationOnly sync.Once
func Get_DurationOnly() gopurs_runtime.Value {
	once_DurationOnly.Do(func() {
		DurationOnly = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationOnly"), "value0": value0})
})
	})
	return DurationOnly
}

var RecurringInterval gopurs_runtime.Value
var once_RecurringInterval sync.Once
func Get_RecurringInterval() gopurs_runtime.Value {
	once_RecurringInterval.Do(func() {
		RecurringInterval = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RecurringInterval"), "value0": value0, "value1": value1})
})
})
	})
	return RecurringInterval
}

var showInterval gopurs_runtime.Value
var once_showInterval sync.Once
func Get_showInterval() gopurs_runtime.Value {
	once_showInterval.Do(func() {
		showInterval = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(StartEnd ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow1_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow1_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Str(")")))))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(DurationEnd ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow1_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Str(")")))))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(StartDuration ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow1_1.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Str(")")))))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(DurationOnly ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Str(")")))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})})
})
})
	})
	return showInterval
}

var showRecurringInterval gopurs_runtime.Value
var once_showRecurringInterval sync.Once
func Get_showRecurringInterval() gopurs_runtime.Value {
	once_showRecurringInterval.Do(func() {
		showRecurringInterval = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictShow1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(RecurringInterval ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(Get_show(), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_showInterval(), dictShow_0), dictShow1_1).PtrVal.(map[string]gopurs_runtime.Value)["show"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Str(")")))))
})})
})
})
	})
	return showRecurringInterval
}

var over gopurs_runtime.Value
var once_over sync.Once
func Get_over() gopurs_runtime.Value {
	once_over.Do(func() {
		over = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(Get_RecurringInterval(), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(f_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})
	})
	return over
}

var foldableInterval gopurs_runtime.Value
var once_foldableInterval sync.Once
func Get_foldableInterval() gopurs_runtime.Value {
	once_foldableInterval.Do(func() {
		foldableInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldl": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v1_1), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v1_1), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v1_1), v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
goto end_branch_0
} else {

}
}
{
__t0 = v1_1
}
end_branch_0:
return __t0
})
})
}), "foldr": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldrDefault(), Get_foldableInterval()), x_0)
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_1 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableInterval().PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], acc_3), gopurs_runtime.Apply(f_2, x_4))
})
})), mempty_1_1)
})
})})
	})
	return foldableInterval
}

var foldableRecurringInterval gopurs_runtime.Value
var once_foldableRecurringInterval sync.Once
func Get_foldableRecurringInterval() gopurs_runtime.Value {
	once_foldableRecurringInterval.Do(func() {
		foldableRecurringInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, i_1), x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"])), x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, i_1), x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, i_1), x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"])
goto end_branch_0
} else {

}
}
{
__t0 = i_1
}
end_branch_0:
return __t0
})
})
}), "foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldrDefault(), Get_foldableInterval()), f_0), i_1)
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_2 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableRecurringInterval().PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], acc_3), gopurs_runtime.Apply(f_2, x_4))
})
})), mempty_1_2)
})
})})
	})
	return foldableRecurringInterval
}

var eqInterval gopurs_runtime.Value
var once_eqInterval sync.Once
func Get_eqInterval() gopurs_runtime.Value {
	once_eqInterval.Do(func() {
		eqInterval = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(y_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq1_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq1_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])).IntVal != 0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(y_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq1_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])).IntVal != 0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(y_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq1_1.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])).IntVal != 0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(y_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly").IntVal != 0 && gopurs_runtime.Apply(gopurs_runtime.Apply(dictEq_0.PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]).IntVal != 0).IntVal != 0)
}
end_branch_0:
return __t0
})
})})
})
})
	})
	return eqInterval
}

var eqRecurringInterval gopurs_runtime.Value
var once_eqRecurringInterval sync.Once
func Get_eqRecurringInterval() gopurs_runtime.Value {
	once_eqRecurringInterval.Do(func() {
		eqRecurringInterval = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eq(), x_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_eqInterval(), dictEq_0), dictEq1_1).PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})})
})
})
	})
	return eqRecurringInterval
}

var ordInterval gopurs_runtime.Value
var once_ordInterval sync.Once
func Get_ordInterval() gopurs_runtime.Value {
	once_ordInterval.Do(func() {
		ordInterval = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
eqInterval1_1_0 := gopurs_runtime.Apply(Get_eqInterval(), gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
eqInterval2_3_1 := gopurs_runtime.Apply(eqInterval1_1_0, gopurs_runtime.Apply(dictOrd1_2.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
v_6_4 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd1_2.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_6_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(v_6_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd1_2.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
v_6_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_6_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(v_6_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd1_2.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
end_branch_8:
__t6 = __t8
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_6:
__t2 = __t6
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
v_6_10 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd1_2.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_6_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(v_6_10.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t11 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
end_branch_11:
__t9 = __t11
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
}
end_branch_9:
__t2 = __t9
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly").IntVal != 0 && gopurs_runtime.Bool(y_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly").IntVal != 0)).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
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
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqInterval2_3_1
})})
})
})
	})
	return ordInterval
}

var ordRecurringInterval gopurs_runtime.Value
var once_ordRecurringInterval sync.Once
func Get_ordRecurringInterval() gopurs_runtime.Value {
	once_ordRecurringInterval.Do(func() {
		ordRecurringInterval = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
ordInterval1_1_0 := gopurs_runtime.Apply(Get_ordInterval(), dictOrd_0)
eqRecurringInterval1_2_1 := gopurs_runtime.Apply(Get_eqRecurringInterval(), gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
eqRecurringInterval2_4_2 := gopurs_runtime.Apply(eqRecurringInterval1_2_1, gopurs_runtime.Apply(dictOrd1_3.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
v_7_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_compare(), x_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_7_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v_7_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(ordInterval1_1_0, dictOrd1_3).PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
end_branch_4:
return __t4
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRecurringInterval2_4_2
})})
})
})
	})
	return ordRecurringInterval
}

var bifunctorInterval gopurs_runtime.Value
var once_bifunctorInterval sync.Once
func Get_bifunctorInterval() gopurs_runtime.Value {
	once_bifunctorInterval.Do(func() {
		bifunctorInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartEnd"), "value0": gopurs_runtime.Apply(v1_1, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(v1_1, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationEnd"), "value0": gopurs_runtime.Apply(v_0, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(v1_1, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartDuration"), "value0": gopurs_runtime.Apply(v1_1, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(v_0, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v2_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationOnly"), "value0": gopurs_runtime.Apply(v_0, v2_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})
})})
	})
	return bifunctorInterval
}

var bifunctorRecurringInterval gopurs_runtime.Value
var once_bifunctorRecurringInterval sync.Once
func Get_bifunctorRecurringInterval() gopurs_runtime.Value {
	once_bifunctorRecurringInterval.Do(func() {
		bifunctorRecurringInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bimap": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartEnd"), "value0": gopurs_runtime.Apply(g_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(g_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationEnd"), "value0": gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(g_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartDuration"), "value0": gopurs_runtime.Apply(g_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationOnly"), "value0": gopurs_runtime.Apply(f_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RecurringInterval"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __t0})
})
})
})})
	})
	return bifunctorRecurringInterval
}

var functorInterval gopurs_runtime.Value
var once_functorInterval sync.Once
func Get_functorInterval() gopurs_runtime.Value {
	once_functorInterval.Do(func() {
		functorInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(v1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartEnd"), "value0": gopurs_runtime.Apply(v1_0, v2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(v1_0, v2_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationEnd"), "value0": v2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(v1_0, v2_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartDuration"), "value0": gopurs_runtime.Apply(v1_0, v2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v2_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v2_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationOnly"), "value0": v2_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})})
	})
	return functorInterval
}

var extendInterval gopurs_runtime.Value
var once_extendInterval sync.Once
func Get_extendInterval() gopurs_runtime.Value {
	once_extendInterval.Do(func() {
		extendInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartEnd"), "value0": gopurs_runtime.Apply(v_0, v1_1), "value1": gopurs_runtime.Apply(v_0, v1_1)})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationEnd"), "value0": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(v_0, v1_1)})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartDuration"), "value0": gopurs_runtime.Apply(v_0, v1_1), "value1": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v1_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationOnly"), "value0": v1_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorInterval()
})})
	})
	return extendInterval
}

var functorRecurringInterval gopurs_runtime.Value
var once_functorRecurringInterval sync.Once
func Get_functorRecurringInterval() gopurs_runtime.Value {
	once_functorRecurringInterval.Do(func() {
		functorRecurringInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartEnd"), "value0": gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationEnd"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartDuration"), "value0": gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly")).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationOnly"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RecurringInterval"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __t0})
})
})})
	})
	return functorRecurringInterval
}

var extendRecurringInterval gopurs_runtime.Value
var once_extendRecurringInterval sync.Once
func Get_extendRecurringInterval() gopurs_runtime.Value {
	once_extendRecurringInterval.Do(func() {
		extendRecurringInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, v_1)
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartEnd"), "value0": __local_var_2_0, "value1": __local_var_2_0})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationEnd"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_2_0})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartDuration"), "value0": __local_var_2_0, "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationOnly"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"]})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RecurringInterval"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __t1})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorRecurringInterval()
})})
	})
	return extendRecurringInterval
}

var traversableInterval gopurs_runtime.Value
var once_traversableInterval sync.Once
func Get_traversableInterval() gopurs_runtime.Value {
	once_traversableInterval.Do(func() {
		traversableInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverse": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
Functor0_2_1 := gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_StartEnd()), gopurs_runtime.Apply(v_3, v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), gopurs_runtime.Apply(v_3, v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(Get_DurationEnd(), v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(v_3, v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__local_var_5_3 := v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("StartDuration"), "value0": v2_6, "value1": __local_var_5_3})
})), gopurs_runtime.Apply(v_3, v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DurationOnly"), "value0": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]}))
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
}), "sequence": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traversableInterval().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorInterval()
}), "Foldable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableInterval()
})})
	})
	return traversableInterval
}

var traversableRecurringInterval gopurs_runtime.Value
var once_traversableRecurringInterval sync.Once
func Get_traversableRecurringInterval() gopurs_runtime.Value {
	once_traversableRecurringInterval.Do(func() {
		traversableRecurringInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverse": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
over1_1_0 := gopurs_runtime.Apply(Get_over(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}))
traverse1_2_1 := gopurs_runtime.Apply(Get_traversableInterval().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0)
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(over1_1_0, gopurs_runtime.Apply(traverse1_2_1, f_3)), i_4)
})
})
}), "sequence": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_traversableRecurringInterval().PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorRecurringInterval()
}), "Foldable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableRecurringInterval()
})})
	})
	return traversableRecurringInterval
}

var bifoldableInterval gopurs_runtime.Value
var once_bifoldableInterval sync.Once
func Get_bifoldableInterval() gopurs_runtime.Value {
	once_bifoldableInterval.Do(func() {
		bifoldableInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bifoldl": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1, gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1, v2_2), v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1, gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v2_2), v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v1_1, gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v2_2), v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v3_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(v_0, v2_2), v3_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})
})
}), "bifoldr": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Bifoldable.Get_bifoldrDefault(), Get_bifoldableInterval()), x_0)
}), "bifoldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_1 := gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
mempty_2_2 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bifoldableInterval().PtrVal.(map[string]gopurs_runtime.Value)["bifoldl"], gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], m_5), gopurs_runtime.Apply(f_3, a_6))
})
})), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], m_5), gopurs_runtime.Apply(g_4, b_6))
})
})), mempty_2_2)
})
})
})})
	})
	return bifoldableInterval
}

var bifoldableRecurringInterval gopurs_runtime.Value
var once_bifoldableRecurringInterval sync.Once
func Get_bifoldableRecurringInterval() gopurs_runtime.Value {
	once_bifoldableRecurringInterval.Do(func() {
		bifoldableRecurringInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bifoldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(g_1, gopurs_runtime.Apply(gopurs_runtime.Apply(g_1, i_2), x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"])), x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(g_1, gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, i_2), x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"])), x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(g_1, gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, i_2), x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"])), x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"])
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, i_2), x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"])
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})
})
}), "bifoldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Bifoldable.Get_bifoldrDefault(), Get_bifoldableInterval()), f_0), g_1), i_2)
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, x_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
})
}), "bifoldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
mempty_2_3 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bifoldableRecurringInterval().PtrVal.(map[string]gopurs_runtime.Value)["bifoldl"], gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_2.PtrVal.(map[string]gopurs_runtime.Value)["append"], m_5), gopurs_runtime.Apply(f_3, a_6))
})
})), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_2.PtrVal.(map[string]gopurs_runtime.Value)["append"], m_5), gopurs_runtime.Apply(g_4, b_6))
})
})), mempty_2_3)
})
})
})})
	})
	return bifoldableRecurringInterval
}

var bitraversableInterval gopurs_runtime.Value
var once_bitraversableInterval sync.Once
func Get_bitraversableInterval() gopurs_runtime.Value {
	once_bitraversableInterval.Do(func() {
		bitraversableInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bitraverse": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartEnd")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_StartEnd()), gopurs_runtime.Apply(v1_4, v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), gopurs_runtime.Apply(v1_4, v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationEnd")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_DurationEnd()), gopurs_runtime.Apply(v_3, v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), gopurs_runtime.Apply(v1_4, v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "StartDuration")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_StartDuration()), gopurs_runtime.Apply(v1_4, v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), gopurs_runtime.Apply(v_3, v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v2_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "DurationOnly")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], Get_DurationOnly()), gopurs_runtime.Apply(v_3, v2_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
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
}), "bisequence": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bitraversableInterval().PtrVal.(map[string]gopurs_runtime.Value)["bitraverse"], dictApplicative_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bifunctorInterval()
}), "Bifoldable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bifoldableInterval()
})})
	})
	return bitraversableInterval
}

var bitraversableRecurringInterval gopurs_runtime.Value
var once_bitraversableRecurringInterval sync.Once
func Get_bitraversableRecurringInterval() gopurs_runtime.Value {
	once_bitraversableRecurringInterval.Do(func() {
		bitraversableRecurringInterval = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bitraverse": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
over1_1_0 := gopurs_runtime.Apply(Get_over(), gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}))
bitraverse1_2_1 := gopurs_runtime.Apply(Get_bitraversableInterval().PtrVal.(map[string]gopurs_runtime.Value)["bitraverse"], dictApplicative_0)
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(over1_1_0, gopurs_runtime.Apply(gopurs_runtime.Apply(bitraverse1_2_1, l_3), r_4)), i_5)
})
})
})
}), "bisequence": gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_bitraversableRecurringInterval().PtrVal.(map[string]gopurs_runtime.Value)["bitraverse"], dictApplicative_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"]), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
}), "Bifunctor0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bifunctorRecurringInterval()
}), "Bifoldable1": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bifoldableRecurringInterval()
})})
	})
	return bitraversableRecurringInterval
}



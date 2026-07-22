package Data_DateTime

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Time "gopurs/output/Data.Time"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Function_Uncurried "gopurs/output/Data.Function.Uncurried"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Time_Component "gopurs/output/Data.Time.Component"
)

var DateTime gopurs_runtime.Value
var once_DateTime sync.Once
func Get_DateTime() gopurs_runtime.Value {
	once_DateTime.Do(func() {
		DateTime = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": value0, "value1": value1})
})
})
	})
	return DateTime
}

var toRecord gopurs_runtime.Value
var once_toRecord sync.Once
func Get_toRecord() gopurs_runtime.Value {
	once_toRecord.Do(func() {
		toRecord = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "January")).IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "February")).IntVal != 0 {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "March")).IntVal != 0 {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "April")).IntVal != 0 {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "May")).IntVal != 0 {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "June")).IntVal != 0 {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "July")).IntVal != 0 {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "August")).IntVal != 0 {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "September")).IntVal != 0 {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "October")).IntVal != 0 {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "November")).IntVal != 0 {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "December")).IntVal != 0 {
__t0 = gopurs_runtime.Int(12)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"year": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "month": __t0, "day": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "hour": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "minute": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], "second": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value2"], "millisecond": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
})
	})
	return toRecord
}

var time gopurs_runtime.Value
var once_time sync.Once
func Get_time() gopurs_runtime.Value {
	once_time.Do(func() {
		time = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
	})
	return time
}

var showDateTime gopurs_runtime.Value
var once_showDateTime sync.Once
func Get_showDateTime() gopurs_runtime.Value {
	once_showDateTime.Do(func() {
		showDateTime = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(DateTime ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Date.Get_showDate().PtrVal.(map[string]gopurs_runtime.Value)["show"], v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Time.Get_showTime().PtrVal.(map[string]gopurs_runtime.Value)["show"], v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Str(")")))))
})})
	})
	return showDateTime
}

var modifyTimeF gopurs_runtime.Value
var once_modifyTimeF sync.Once
func Get_modifyTimeF() gopurs_runtime.Value {
	once_modifyTimeF.Do(func() {
		modifyTimeF = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(Get_DateTime(), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(f_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
})
	})
	return modifyTimeF
}

var modifyTime gopurs_runtime.Value
var once_modifyTime sync.Once
func Get_modifyTime() gopurs_runtime.Value {
	once_modifyTime.Do(func() {
		modifyTime = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
	})
	return modifyTime
}

var modifyDateF gopurs_runtime.Value
var once_modifyDateF sync.Once
func Get_modifyDateF() gopurs_runtime.Value {
	once_modifyDateF.Do(func() {
		modifyDateF = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": a_4, "value1": __local_var_3_0})
})), gopurs_runtime.Apply(f_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
})
})
})
	})
	return modifyDateF
}

var modifyDate gopurs_runtime.Value
var once_modifyDate sync.Once
func Get_modifyDate() gopurs_runtime.Value {
	once_modifyDate.Do(func() {
		modifyDate = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
})
	})
	return modifyDate
}

var eqDateTime gopurs_runtime.Value
var once_eqDateTime sync.Once
func Get_eqDateTime() gopurs_runtime.Value {
	once_eqDateTime.Do(func() {
		eqDateTime = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Date.Get_eqDate().PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Time.Get_eqTime().PtrVal.(map[string]gopurs_runtime.Value)["eq"], x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})})
	})
	return eqDateTime
}

var ordDateTime gopurs_runtime.Value
var once_ordDateTime sync.Once
func Get_ordDateTime() gopurs_runtime.Value {
	once_ordDateTime.Do(func() {
		ordDateTime = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Date.Get_ordDate().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Time.Get_ordTime().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}
end_branch_1:
return __t1
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqDateTime()
})})
	})
	return ordDateTime
}

var diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		diff = gopurs_runtime.Func(func(dictDuration_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dt1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dt2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictDuration_0.PtrVal.(map[string]gopurs_runtime.Value)["toDuration"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn2(), Get_calcDiff()), gopurs_runtime.Apply(Get_toRecord(), dt1_1)), gopurs_runtime.Apply(Get_toRecord(), dt2_2)))
})
})
})
	})
	return diff
}

var date gopurs_runtime.Value
var once_date sync.Once
func Get_date() gopurs_runtime.Value {
	once_date.Do(func() {
		date = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
})
	})
	return date
}

var boundedDateTime gopurs_runtime.Value
var once_boundedDateTime sync.Once
func Get_boundedDateTime() gopurs_runtime.Value {
	once_boundedDateTime.Do(func() {
		boundedDateTime = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bottom": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Date"), "value0": pkg_Data_Date_Component.Get_boundedYear().PtrVal.(map[string]gopurs_runtime.Value)["bottom"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("January")}), "value2": gopurs_runtime.Int(1)}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": gopurs_runtime.Int(0), "value1": gopurs_runtime.Int(0), "value2": gopurs_runtime.Int(0), "value3": gopurs_runtime.Int(0)})}), "top": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Date"), "value0": gopurs_runtime.Int(275759), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("December")}), "value2": gopurs_runtime.Int(31)}), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": gopurs_runtime.Int(23), "value1": gopurs_runtime.Int(59), "value2": gopurs_runtime.Int(59), "value3": gopurs_runtime.Int(999)})}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordDateTime()
})})
	})
	return boundedDateTime
}

var adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		adjust = gopurs_runtime.Func(func(dictDuration_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dt_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_adjustImpl(), pkg_Data_Maybe.Get_Just()), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})), gopurs_runtime.Apply(dictDuration_0.PtrVal.(map[string]gopurs_runtime.Value)["fromDuration"], d_1)), gopurs_runtime.Apply(Get_toRecord(), dt_2))
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_4_2 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumYear().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["year"])
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Apply(pkg_Data_Date.Get_exactDate(), __local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"])})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_4:
__local_var_5_3 := __t4
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["month"].IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_7 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["day"])
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("January")})), __local_var_6_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_11 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["hour"])
var __t12 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_8_11.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_9_13 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["minute"])
var __t14 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_13.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_10_15 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["second"])
var __t16 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_10_15.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_11_17 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["millisecond"])
var __t18 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_17.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t18 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": __local_var_7_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": __local_var_8_11.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_9_13.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": __local_var_10_15.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": __local_var_11_17.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})})
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_18:
__t16 = __t18
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_16:
__t14 = __t16
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_14:
__t12 = __t14
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_12:
__t10 = __t12
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_10:
__t8 = __t10
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_8:
__t6 = __t8
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["month"].IntVal == gopurs_runtime.Int(2).IntVal)).IntVal != 0 {
var __t19 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_20 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["day"])
var __t21 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_20.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_22 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("February")})), __local_var_6_20.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t23 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_22.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_24 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["hour"])
var __t25 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_8_24.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_9_26 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["minute"])
var __t27 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_26.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_10_28 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["second"])
var __t29 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_10_28.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_11_30 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["millisecond"])
var __t31 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_30.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t31 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": __local_var_7_22.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": __local_var_8_24.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_9_26.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": __local_var_10_28.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": __local_var_11_30.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})})
goto end_branch_31
} else {

}
}
{
__t31 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_31:
__t29 = __t31
goto end_branch_29
} else {

}
}
{
__t29 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_29:
__t27 = __t29
goto end_branch_27
} else {

}
}
{
__t27 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_27:
__t25 = __t27
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_25:
__t23 = __t25
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_23:
__t21 = __t23
goto end_branch_21
} else {

}
}
{
__t21 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_21:
__t19 = __t21
goto end_branch_19
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t19 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
__t5 = __t19
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["month"].IntVal == gopurs_runtime.Int(3).IntVal)).IntVal != 0 {
var __t32 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_33 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["day"])
var __t34 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_33.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_35 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("March")})), __local_var_6_33.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t36 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_35.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_37 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["hour"])
var __t38 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_8_37.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_9_39 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["minute"])
var __t40 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_39.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_10_41 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["second"])
var __t42 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_10_41.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_11_43 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["millisecond"])
var __t44 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_43.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t44 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": __local_var_7_35.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": __local_var_8_37.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_9_39.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": __local_var_10_41.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": __local_var_11_43.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})})
goto end_branch_44
} else {

}
}
{
__t44 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_44:
__t42 = __t44
goto end_branch_42
} else {

}
}
{
__t42 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_42:
__t40 = __t42
goto end_branch_40
} else {

}
}
{
__t40 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_40:
__t38 = __t40
goto end_branch_38
} else {

}
}
{
__t38 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_38:
__t36 = __t38
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_36:
__t34 = __t36
goto end_branch_34
} else {

}
}
{
__t34 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_34:
__t32 = __t34
goto end_branch_32
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t32 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_32:
__t5 = __t32
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["month"].IntVal == gopurs_runtime.Int(4).IntVal)).IntVal != 0 {
var __t45 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_46 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["day"])
var __t47 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_46.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_48 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("April")})), __local_var_6_46.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t49 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_48.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_50 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["hour"])
var __t51 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_8_50.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_9_52 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["minute"])
var __t53 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_52.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_10_54 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["second"])
var __t55 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_10_54.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_11_56 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["millisecond"])
var __t57 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_56.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t57 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": __local_var_7_48.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": __local_var_8_50.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_9_52.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": __local_var_10_54.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": __local_var_11_56.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})})
goto end_branch_57
} else {

}
}
{
__t57 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_57:
__t55 = __t57
goto end_branch_55
} else {

}
}
{
__t55 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_55:
__t53 = __t55
goto end_branch_53
} else {

}
}
{
__t53 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_53:
__t51 = __t53
goto end_branch_51
} else {

}
}
{
__t51 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_51:
__t49 = __t51
goto end_branch_49
} else {

}
}
{
__t49 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_49:
__t47 = __t49
goto end_branch_47
} else {

}
}
{
__t47 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_47:
__t45 = __t47
goto end_branch_45
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t45 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_45
} else {

}
}
{
__t45 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_45:
__t5 = __t45
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["month"].IntVal == gopurs_runtime.Int(5).IntVal)).IntVal != 0 {
var __t58 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_59 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["day"])
var __t60 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_59.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_61 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("May")})), __local_var_6_59.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t62 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_61.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_63 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["hour"])
var __t64 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_8_63.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_9_65 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["minute"])
var __t66 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_65.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_10_67 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["second"])
var __t68 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_10_67.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_11_69 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["millisecond"])
var __t70 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_69.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t70 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": __local_var_7_61.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": __local_var_8_63.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_9_65.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": __local_var_10_67.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": __local_var_11_69.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})})
goto end_branch_70
} else {

}
}
{
__t70 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_70:
__t68 = __t70
goto end_branch_68
} else {

}
}
{
__t68 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_68:
__t66 = __t68
goto end_branch_66
} else {

}
}
{
__t66 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_66:
__t64 = __t66
goto end_branch_64
} else {

}
}
{
__t64 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_64:
__t62 = __t64
goto end_branch_62
} else {

}
}
{
__t62 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_62:
__t60 = __t62
goto end_branch_60
} else {

}
}
{
__t60 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_60:
__t58 = __t60
goto end_branch_58
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t58 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_58
} else {

}
}
{
__t58 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_58:
__t5 = __t58
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["month"].IntVal == gopurs_runtime.Int(6).IntVal)).IntVal != 0 {
var __t71 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_72 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["day"])
var __t73 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_72.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_74 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("June")})), __local_var_6_72.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t75 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_74.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_76 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["hour"])
var __t77 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_8_76.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_9_78 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["minute"])
var __t79 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_78.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_10_80 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["second"])
var __t81 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_10_80.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_11_82 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["millisecond"])
var __t83 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_82.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t83 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": __local_var_7_74.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": __local_var_8_76.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_9_78.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": __local_var_10_80.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": __local_var_11_82.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})})
goto end_branch_83
} else {

}
}
{
__t83 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_83:
__t81 = __t83
goto end_branch_81
} else {

}
}
{
__t81 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_81:
__t79 = __t81
goto end_branch_79
} else {

}
}
{
__t79 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_79:
__t77 = __t79
goto end_branch_77
} else {

}
}
{
__t77 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_77:
__t75 = __t77
goto end_branch_75
} else {

}
}
{
__t75 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_75:
__t73 = __t75
goto end_branch_73
} else {

}
}
{
__t73 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_73:
__t71 = __t73
goto end_branch_71
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t71 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_71
} else {

}
}
{
__t71 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_71:
__t5 = __t71
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["month"].IntVal == gopurs_runtime.Int(7).IntVal)).IntVal != 0 {
var __t84 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_85 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["day"])
var __t86 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_85.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_87 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("July")})), __local_var_6_85.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t88 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_87.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_89 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["hour"])
var __t90 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_8_89.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_9_91 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["minute"])
var __t92 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_91.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_10_93 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["second"])
var __t94 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_10_93.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_11_95 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["millisecond"])
var __t96 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_95.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t96 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": __local_var_7_87.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": __local_var_8_89.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_9_91.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": __local_var_10_93.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": __local_var_11_95.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})})
goto end_branch_96
} else {

}
}
{
__t96 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_96:
__t94 = __t96
goto end_branch_94
} else {

}
}
{
__t94 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_94:
__t92 = __t94
goto end_branch_92
} else {

}
}
{
__t92 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_92:
__t90 = __t92
goto end_branch_90
} else {

}
}
{
__t90 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_90:
__t88 = __t90
goto end_branch_88
} else {

}
}
{
__t88 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_88:
__t86 = __t88
goto end_branch_86
} else {

}
}
{
__t86 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_86:
__t84 = __t86
goto end_branch_84
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t84 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_84
} else {

}
}
{
__t84 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_84:
__t5 = __t84
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["month"].IntVal == gopurs_runtime.Int(8).IntVal)).IntVal != 0 {
var __t97 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_98 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["day"])
var __t99 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_98.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_100 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("August")})), __local_var_6_98.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t101 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_100.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_102 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["hour"])
var __t103 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_8_102.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_9_104 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["minute"])
var __t105 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_104.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_10_106 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["second"])
var __t107 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_10_106.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_11_108 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["millisecond"])
var __t109 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_108.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t109 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": __local_var_7_100.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": __local_var_8_102.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_9_104.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": __local_var_10_106.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": __local_var_11_108.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})})
goto end_branch_109
} else {

}
}
{
__t109 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_109:
__t107 = __t109
goto end_branch_107
} else {

}
}
{
__t107 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_107:
__t105 = __t107
goto end_branch_105
} else {

}
}
{
__t105 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_105:
__t103 = __t105
goto end_branch_103
} else {

}
}
{
__t103 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_103:
__t101 = __t103
goto end_branch_101
} else {

}
}
{
__t101 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_101:
__t99 = __t101
goto end_branch_99
} else {

}
}
{
__t99 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_99:
__t97 = __t99
goto end_branch_97
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t97 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_97
} else {

}
}
{
__t97 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_97:
__t5 = __t97
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["month"].IntVal == gopurs_runtime.Int(9).IntVal)).IntVal != 0 {
var __t110 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_111 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["day"])
var __t112 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_111.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_113 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("September")})), __local_var_6_111.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t114 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_113.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_115 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["hour"])
var __t116 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_8_115.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_9_117 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["minute"])
var __t118 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_117.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_10_119 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["second"])
var __t120 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_10_119.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_11_121 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["millisecond"])
var __t122 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_121.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t122 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": __local_var_7_113.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": __local_var_8_115.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_9_117.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": __local_var_10_119.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": __local_var_11_121.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})})
goto end_branch_122
} else {

}
}
{
__t122 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_122:
__t120 = __t122
goto end_branch_120
} else {

}
}
{
__t120 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_120:
__t118 = __t120
goto end_branch_118
} else {

}
}
{
__t118 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_118:
__t116 = __t118
goto end_branch_116
} else {

}
}
{
__t116 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_116:
__t114 = __t116
goto end_branch_114
} else {

}
}
{
__t114 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_114:
__t112 = __t114
goto end_branch_112
} else {

}
}
{
__t112 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_112:
__t110 = __t112
goto end_branch_110
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t110 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_110
} else {

}
}
{
__t110 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_110:
__t5 = __t110
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["month"].IntVal == gopurs_runtime.Int(10).IntVal)).IntVal != 0 {
var __t123 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_124 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["day"])
var __t125 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_124.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_126 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("October")})), __local_var_6_124.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t127 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_126.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_128 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["hour"])
var __t129 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_8_128.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_9_130 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["minute"])
var __t131 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_130.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_10_132 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["second"])
var __t133 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_10_132.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_11_134 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["millisecond"])
var __t135 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_134.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t135 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": __local_var_7_126.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": __local_var_8_128.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_9_130.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": __local_var_10_132.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": __local_var_11_134.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})})
goto end_branch_135
} else {

}
}
{
__t135 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_135:
__t133 = __t135
goto end_branch_133
} else {

}
}
{
__t133 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_133:
__t131 = __t133
goto end_branch_131
} else {

}
}
{
__t131 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_131:
__t129 = __t131
goto end_branch_129
} else {

}
}
{
__t129 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_129:
__t127 = __t129
goto end_branch_127
} else {

}
}
{
__t127 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_127:
__t125 = __t127
goto end_branch_125
} else {

}
}
{
__t125 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_125:
__t123 = __t125
goto end_branch_123
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t123 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_123
} else {

}
}
{
__t123 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_123:
__t5 = __t123
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["month"].IntVal == gopurs_runtime.Int(11).IntVal)).IntVal != 0 {
var __t136 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_137 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["day"])
var __t138 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_137.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_139 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("November")})), __local_var_6_137.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t140 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_139.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_141 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["hour"])
var __t142 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_8_141.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_9_143 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["minute"])
var __t144 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_143.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_10_145 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["second"])
var __t146 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_10_145.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_11_147 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["millisecond"])
var __t148 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_147.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t148 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": __local_var_7_139.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": __local_var_8_141.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_9_143.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": __local_var_10_145.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": __local_var_11_147.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})})
goto end_branch_148
} else {

}
}
{
__t148 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_148:
__t146 = __t148
goto end_branch_146
} else {

}
}
{
__t146 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_146:
__t144 = __t146
goto end_branch_144
} else {

}
}
{
__t144 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_144:
__t142 = __t144
goto end_branch_142
} else {

}
}
{
__t142 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_142:
__t140 = __t142
goto end_branch_140
} else {

}
}
{
__t140 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_140:
__t138 = __t140
goto end_branch_138
} else {

}
}
{
__t138 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_138:
__t136 = __t138
goto end_branch_136
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t136 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_136
} else {

}
}
{
__t136 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_136:
__t5 = __t136
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["month"].IntVal == gopurs_runtime.Int(12).IntVal)).IntVal != 0 {
var __t149 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_150 := gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumDay().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["day"])
var __t151 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_150.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_152 := gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("December")})), __local_var_6_150.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t153 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_152.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_8_154 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["hour"])
var __t155 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_8_154.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_9_156 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["minute"])
var __t157 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_9_156.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_10_158 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["second"])
var __t159 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_10_158.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_11_160 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["millisecond"])
var __t161 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_11_160.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t161 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("DateTime"), "value0": __local_var_7_152.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": __local_var_8_154.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_9_156.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": __local_var_10_158.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": __local_var_11_160.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})})
goto end_branch_161
} else {

}
}
{
__t161 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_161:
__t159 = __t161
goto end_branch_159
} else {

}
}
{
__t159 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_159:
__t157 = __t159
goto end_branch_157
} else {

}
}
{
__t157 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_157:
__t155 = __t157
goto end_branch_155
} else {

}
}
{
__t155 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_155:
__t153 = __t155
goto end_branch_153
} else {

}
}
{
__t153 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_153:
__t151 = __t153
goto end_branch_151
} else {

}
}
{
__t151 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_151:
__t149 = __t151
goto end_branch_149
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t149 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_149
} else {

}
}
{
__t149 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_149:
__t5 = __t149
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t1 = __t5
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
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
})
})
	})
	return adjust
}

func Get_adjustImpl() gopurs_runtime.Value {
	return AdjustImpl
}

func Get_calcDiff() gopurs_runtime.Value {
	return CalcDiff
}

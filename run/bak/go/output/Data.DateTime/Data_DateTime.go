package Data_DateTime

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Time "gopurs/output/Data.Time"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
)

var DateTime gopurs_runtime.Value
var once_DateTime sync.Once
func Get_DateTime() gopurs_runtime.Value {
	once_DateTime.Do(func() {
		DateTime = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("DateTime"), value0, value1)
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
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value1"), "_tag").StrVal == "January")).IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value1"), "_tag").StrVal == "February")).IntVal != 0 {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value1"), "_tag").StrVal == "March")).IntVal != 0 {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value1"), "_tag").StrVal == "April")).IntVal != 0 {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value1"), "_tag").StrVal == "May")).IntVal != 0 {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value1"), "_tag").StrVal == "June")).IntVal != 0 {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value1"), "_tag").StrVal == "July")).IntVal != 0 {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value1"), "_tag").StrVal == "August")).IntVal != 0 {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value1"), "_tag").StrVal == "September")).IntVal != 0 {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value1"), "_tag").StrVal == "October")).IntVal != 0 {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value1"), "_tag").StrVal == "November")).IntVal != 0 {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value1"), "_tag").StrVal == "December")).IntVal != 0 {
__t0 = gopurs_runtime.Int(12)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.RecordDict([]string{"year", "month", "day", "hour", "minute", "second", "millisecond"}, []gopurs_runtime.Value{gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value0"), __t0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value2"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value1"), "value0"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value1"), "value1"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value1"), "value2"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value1"), "value3")})
})
	})
	return toRecord
}

var time gopurs_runtime.Value
var once_time sync.Once
func Get_time() gopurs_runtime.Value {
	once_time.Do(func() {
		time = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v_0, "value1")
})
	})
	return time
}

var showDateTime gopurs_runtime.Value
var once_showDateTime sync.Once
func Get_showDateTime() gopurs_runtime.Value {
	once_showDateTime.Do(func() {
		showDateTime = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(DateTime ").StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date.Get_showDate(), "show"), gopurs_runtime.RecordGet(v_0, "value0")).StrVal).StrVal + gopurs_runtime.Str(" ").StrVal).StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time.Get_showTime(), "show"), gopurs_runtime.RecordGet(v_0, "value1")).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
	})
	return showDateTime
}

var modifyTimeF gopurs_runtime.Value
var once_modifyTimeF sync.Once
func Get_modifyTimeF() gopurs_runtime.Value {
	once_modifyTimeF.Do(func() {
		modifyTimeF = gopurs_runtime.Func3(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(Get_DateTime(), gopurs_runtime.RecordGet(v_2, "value0")), gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(v_2, "value1")))
})
	})
	return modifyTimeF
}

var modifyTime gopurs_runtime.Value
var once_modifyTime sync.Once
func Get_modifyTime() gopurs_runtime.Value {
	once_modifyTime.Do(func() {
		modifyTime = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("DateTime"), gopurs_runtime.RecordGet(v_1, "value0"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_1, "value1")))
})
	})
	return modifyTime
}

var modifyDateF gopurs_runtime.Value
var once_modifyDateF sync.Once
func Get_modifyDateF() gopurs_runtime.Value {
	once_modifyDateF.Do(func() {
		modifyDateF = gopurs_runtime.Func3(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.RecordGet(v_2, "value1")
_ = __local_var_3_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("DateTime"), a_4, __local_var_3_0)
}), gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(v_2, "value0")))
})
	})
	return modifyDateF
}

var modifyDate gopurs_runtime.Value
var once_modifyDate sync.Once
func Get_modifyDate() gopurs_runtime.Value {
	once_modifyDate.Do(func() {
		modifyDate = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("DateTime"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_1, "value0")), gopurs_runtime.RecordGet(v_1, "value1"))
})
	})
	return modifyDate
}

var eqDateTime gopurs_runtime.Value
var once_eqDateTime sync.Once
func Get_eqDateTime() gopurs_runtime.Value {
	once_eqDateTime.Do(func() {
		eqDateTime = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date.Get_eqDate(), "eq"), gopurs_runtime.RecordGet(x_0, "value0"), gopurs_runtime.RecordGet(y_1, "value0")).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_0, "value1"), "value0").IntVal == gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(y_1, "value1"), "value0").IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_0, "value1"), "value1").IntVal == gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(y_1, "value1"), "value1").IntVal).IntVal != 0).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_0, "value1"), "value2").IntVal == gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(y_1, "value1"), "value2").IntVal).IntVal != 0).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(x_0, "value1"), "value3").IntVal == gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(y_1, "value1"), "value3").IntVal).IntVal != 0).IntVal != 0)
}))
	})
	return eqDateTime
}

var ordDateTime gopurs_runtime.Value
var once_ordDateTime sync.Once
func Get_ordDateTime() gopurs_runtime.Value {
	once_ordDateTime.Do(func() {
		ordDateTime = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date.Get_ordDate(), "compare"), gopurs_runtime.RecordGet(x_0, "value0"), gopurs_runtime.RecordGet(y_1, "value0"))
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2_0, "_tag").StrVal == "LT")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("LT"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2_0, "_tag").StrVal == "GT")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("GT"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time.Get_ordTime(), "compare"), gopurs_runtime.RecordGet(x_0, "value1"), gopurs_runtime.RecordGet(y_1, "value1"))
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqDateTime()
}))
	})
	return ordDateTime
}

var diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		diff = gopurs_runtime.Func3(func(dictDuration_0 gopurs_runtime.Value, dt1_1 gopurs_runtime.Value, dt2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "toDuration"), gopurs_runtime.UncurriedApp2(Get_calcDiff(), gopurs_runtime.Apply(Get_toRecord(), dt1_1), gopurs_runtime.Apply(Get_toRecord(), dt2_2)))
})
	})
	return diff
}

var date gopurs_runtime.Value
var once_date sync.Once
func Get_date() gopurs_runtime.Value {
	once_date.Do(func() {
		date = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v_0, "value0")
})
	})
	return date
}

var boundedDateTime gopurs_runtime.Value
var once_boundedDateTime sync.Once
func Get_boundedDateTime() gopurs_runtime.Value {
	once_boundedDateTime.Do(func() {
		boundedDateTime = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("DateTime"), gopurs_runtime.RecordDict4("_tag", "value0", "value1", "value2", gopurs_runtime.Str("Date"), gopurs_runtime.Int(-271820), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("January")), gopurs_runtime.Int(1)), gopurs_runtime.RecordDict5("_tag", "value0", "value1", "value2", "value3", gopurs_runtime.Str("Time"), gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0))), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("DateTime"), gopurs_runtime.RecordDict4("_tag", "value0", "value1", "value2", gopurs_runtime.Str("Date"), gopurs_runtime.Int(275759), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("December")), gopurs_runtime.Int(31)), gopurs_runtime.RecordDict5("_tag", "value0", "value1", "value2", "value3", gopurs_runtime.Str("Time"), gopurs_runtime.Int(23), gopurs_runtime.Int(59), gopurs_runtime.Int(59), gopurs_runtime.Int(999))), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordDateTime()
}))
	})
	return boundedDateTime
}

var adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		adjust = gopurs_runtime.Func3(func(dictDuration_0 gopurs_runtime.Value, d_1 gopurs_runtime.Value, dt_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply4(Get_adjustImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "fromDuration"), d_1), gopurs_runtime.Apply(Get_toRecord(), dt_2))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_3_0, "_tag").StrVal == "Just")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "year").IntVal >= gopurs_runtime.Int(-271820).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "year").IntVal <= gopurs_runtime.Int(275759).IntVal).IntVal != 0)).IntVal != 0 {
__t3 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "year"))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_3:
__local_var_4_2 := __t3
_ = __local_var_4_2
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_2, "_tag").StrVal == "Just")).IntVal != 0 {
__t5 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.Apply(pkg_Data_Date.Get_exactDate(), gopurs_runtime.RecordGet(__local_var_4_2, "value0")))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_5:
__local_var_5_4 := __t5
_ = __local_var_5_4
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "month").IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("January")))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "month").IntVal == gopurs_runtime.Int(2).IntVal)).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("February")))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "month").IntVal == gopurs_runtime.Int(3).IntVal)).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("March")))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "month").IntVal == gopurs_runtime.Int(4).IntVal)).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("April")))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "month").IntVal == gopurs_runtime.Int(5).IntVal)).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("May")))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "month").IntVal == gopurs_runtime.Int(6).IntVal)).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("June")))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "month").IntVal == gopurs_runtime.Int(7).IntVal)).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("July")))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "month").IntVal == gopurs_runtime.Int(8).IntVal)).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("August")))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "month").IntVal == gopurs_runtime.Int(9).IntVal)).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("September")))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "month").IntVal == gopurs_runtime.Int(10).IntVal)).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("October")))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "month").IntVal == gopurs_runtime.Int(11).IntVal)).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("November")))
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "month").IntVal == gopurs_runtime.Int(12).IntVal)).IntVal != 0 {
__t7 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("December")))
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_7:
__local_var_6_6 := __t7
_ = __local_var_6_6
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_5_4, "_tag").StrVal == "Just")).IntVal != 0 {
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_6_6, "_tag").StrVal == "Just")).IntVal != 0 {
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "day").IntVal >= gopurs_runtime.Int(1).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "day").IntVal <= gopurs_runtime.Int(31).IntVal).IntVal != 0)).IntVal != 0 {
__t11 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "value0"), gopurs_runtime.RecordGet(__local_var_6_6, "value0"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "day")))
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_11:
__t10 = __t11
goto end_branch_10
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "day").IntVal >= gopurs_runtime.Int(1).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "day").IntVal <= gopurs_runtime.Int(31).IntVal).IntVal != 0)).IntVal != 0 {
__t10 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_10:
__t9 = __t10
goto end_branch_9
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_5_4, "_tag").StrVal == "Nothing")).IntVal != 0 {
var __t12 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "day").IntVal >= gopurs_runtime.Int(1).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "day").IntVal <= gopurs_runtime.Int(31).IntVal).IntVal != 0)).IntVal != 0 {
__t12 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_12:
__t9 = __t12
goto end_branch_9
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "day").IntVal >= gopurs_runtime.Int(1).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "day").IntVal <= gopurs_runtime.Int(31).IntVal).IntVal != 0)).IntVal != 0 {
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
__local_var_7_8 := __t9
_ = __local_var_7_8
var __t14 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_7_8, "_tag").StrVal == "Just")).IntVal != 0 {
__t14 = gopurs_runtime.RecordGet(__local_var_7_8, "value0")
goto end_branch_14
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_7_8, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t14 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
__local_var_8_13 := __t14
_ = __local_var_8_13
var __t16 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_8_13, "_tag").StrVal == "Just")).IntVal != 0 {
__t16 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.Apply(Get_DateTime(), gopurs_runtime.RecordGet(__local_var_8_13, "value0")))
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_16:
__local_var_9_15 := __t16
_ = __local_var_9_15
var __t17 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "hour").IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "hour").IntVal <= gopurs_runtime.Int(23).IntVal).IntVal != 0)).IntVal != 0 {
var __t18 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "minute").IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "minute").IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
var __t19 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "second").IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "second").IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
var __t20 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "millisecond").IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "millisecond").IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
var __t21 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Just")).IntVal != 0 {
__t21 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_15, "value0"), gopurs_runtime.RecordDict5("_tag", "value0", "value1", "value2", "value3", gopurs_runtime.Str("Time"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "hour"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "minute"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "second"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "millisecond"))))
goto end_branch_21
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t21 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
__t20 = __t21
goto end_branch_20
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Just")).IntVal != 0 {
__t20 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_20
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t20 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
__t19 = __t20
goto end_branch_19
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "millisecond").IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "millisecond").IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
var __t22 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Just")).IntVal != 0 {
__t22 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_22
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t22 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
__t19 = __t22
goto end_branch_19
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Just")).IntVal != 0 {
__t19 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_19
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t19 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
__t18 = __t19
goto end_branch_18
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "second").IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "second").IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
var __t23 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "millisecond").IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "millisecond").IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
var __t24 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Just")).IntVal != 0 {
__t24 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_24
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t24 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
__t23 = __t24
goto end_branch_23
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Just")).IntVal != 0 {
__t23 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_23
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t23 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t18 = __t23
goto end_branch_18
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "millisecond").IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "millisecond").IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
var __t25 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Just")).IntVal != 0 {
__t25 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_25
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t25 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
__t18 = __t25
goto end_branch_18
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Just")).IntVal != 0 {
__t18 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_18
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t18 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
__t17 = __t18
goto end_branch_17
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "millisecond").IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_3_0, "value0"), "millisecond").IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
var __t26 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Just")).IntVal != 0 {
__t26 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_26
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t26 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
__t17 = __t26
goto end_branch_17
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Just")).IntVal != 0 {
__t17 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_17
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_9_15, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t17 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
__t1 = __t17
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_3_0, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
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
	return adjust
}

func Get_adjustImpl() gopurs_runtime.Value {
	return _Gopurs_AdjustImpl
}

func Get_calcDiff() gopurs_runtime.Value {
	return _Gopurs_CalcDiff
}

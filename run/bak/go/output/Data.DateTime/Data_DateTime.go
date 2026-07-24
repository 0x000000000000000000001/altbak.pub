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
return gopurs_runtime.Constructor2("DateTime", value0, value1)
})
})
	})
	return DateTime
}

var toRecord gopurs_runtime.Value
var once_toRecord sync.Once
func Get_toRecord() gopurs_runtime.Value {
	once_toRecord.Do(func() {
		toRecord = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "January").IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "February").IntVal != 0 {
__t0 = gopurs_runtime.Int(2)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "March").IntVal != 0 {
__t0 = gopurs_runtime.Int(3)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "April").IntVal != 0 {
__t0 = gopurs_runtime.Int(4)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "May").IntVal != 0 {
__t0 = gopurs_runtime.Int(5)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "June").IntVal != 0 {
__t0 = gopurs_runtime.Int(6)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "July").IntVal != 0 {
__t0 = gopurs_runtime.Int(7)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "August").IntVal != 0 {
__t0 = gopurs_runtime.Int(8)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "September").IntVal != 0 {
__t0 = gopurs_runtime.Int(9)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "October").IntVal != 0 {
__t0 = gopurs_runtime.Int(10)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "November").IntVal != 0 {
__t0 = gopurs_runtime.Int(11)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[1].StrVal == "December").IntVal != 0 {
__t0 = gopurs_runtime.Int(12)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.RecordDict([]string{"year", "month", "day", "hour", "minute", "second", "millisecond"}, []gopurs_runtime.Value{(*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[0], __t0, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0].UnsafePtr)[2], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].UnsafePtr)[0], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].UnsafePtr)[1], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].UnsafePtr)[2], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].UnsafePtr)[3]})
}()
})
	})
	return toRecord
}

var time gopurs_runtime.Value
var once_time sync.Once
func Get_time() gopurs_runtime.Value {
	once_time.Do(func() {
		time = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]
}()
})
	})
	return time
}

var showDateTime gopurs_runtime.Value
var once_showDateTime sync.Once
func Get_showDateTime() gopurs_runtime.Value {
	once_showDateTime.Do(func() {
		showDateTime = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(DateTime " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Date.Get_showDate(), "show"), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]).StrVal + " " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time.Get_showTime(), "show"), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]).StrVal + ")")
}))
	})
	return showDateTime
}

var modifyTimeF gopurs_runtime.Value
var once_modifyTimeF sync.Once
func Get_modifyTimeF() gopurs_runtime.Value {
	once_modifyTimeF.Do(func() {
		modifyTimeF = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modifyTimeF(dictFunctor_0_box, f_1_box, v_2_box)
})
	})
	return modifyTimeF
}

var modifyTime gopurs_runtime.Value
var once_modifyTime sync.Once
func Get_modifyTime() gopurs_runtime.Value {
	once_modifyTime.Do(func() {
		modifyTime = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modifyTime(f_0_box, v_1_box)
})
	})
	return modifyTime
}

var modifyDateF gopurs_runtime.Value
var once_modifyDateF sync.Once
func Get_modifyDateF() gopurs_runtime.Value {
	once_modifyDateF.Do(func() {
		modifyDateF = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modifyDateF(dictFunctor_0_box, f_1_box, v_2_box)
})
	})
	return modifyDateF
}

var modifyDate gopurs_runtime.Value
var once_modifyDate sync.Once
func Get_modifyDate() gopurs_runtime.Value {
	once_modifyDate.Do(func() {
		modifyDate = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modifyDate(f_0_box, v_1_box)
})
	})
	return modifyDate
}

var eqDateTime gopurs_runtime.Value
var once_eqDateTime sync.Once
func Get_eqDateTime() gopurs_runtime.Value {
	once_eqDateTime.Do(func() {
		eqDateTime = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date.Get_eqDate(), "eq"), (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0]).IntVal != 0 && (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].UnsafePtr)[0].IntVal == (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].UnsafePtr)[0].IntVal && (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].UnsafePtr)[1].IntVal == (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].UnsafePtr)[1].IntVal && (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].UnsafePtr)[2].IntVal == (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].UnsafePtr)[2].IntVal && (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].UnsafePtr)[3].IntVal == (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].UnsafePtr)[3].IntVal)
}))
	})
	return eqDateTime
}

var ordDateTime gopurs_runtime.Value
var once_ordDateTime sync.Once
func Get_ordDateTime() gopurs_runtime.Value {
	once_ordDateTime.Do(func() {
		ordDateTime = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Date.Get_ordDate(), "compare"), (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0])
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2_0.StrVal == "LT").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("LT")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_2_0.StrVal == "GT").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("GT")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time.Get_ordTime(), "compare"), (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1])
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
		diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, dt1_1_box gopurs_runtime.Value, dt2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_diff(dictDuration_0_box, dt1_1_box, dt2_2_box)
})
	})
	return diff
}

var date gopurs_runtime.Value
var once_date sync.Once
func Get_date() gopurs_runtime.Value {
	once_date.Do(func() {
		date = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]
}()
})
	})
	return date
}

var boundedDateTime gopurs_runtime.Value
var once_boundedDateTime sync.Once
func Get_boundedDateTime() gopurs_runtime.Value {
	once_boundedDateTime.Do(func() {
		boundedDateTime = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Constructor2("DateTime", gopurs_runtime.Constructor3("Date", gopurs_runtime.Int(-271820), gopurs_runtime.Constructor0("January"), gopurs_runtime.Int(1)), gopurs_runtime.Constructor4("Time", gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0))), gopurs_runtime.Constructor2("DateTime", gopurs_runtime.Constructor3("Date", gopurs_runtime.Int(275759), gopurs_runtime.Constructor0("December"), gopurs_runtime.Int(31)), gopurs_runtime.Constructor4("Time", gopurs_runtime.Int(23), gopurs_runtime.Int(59), gopurs_runtime.Int(59), gopurs_runtime.Int(999))), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordDateTime()
}))
	})
	return boundedDateTime
}

var adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		adjust = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, d_1_box gopurs_runtime.Value, dt_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_adjust(dictDuration_0_box, d_1_box, dt_2_box)
})
	})
	return adjust
}

func Call_modifyTimeF(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(Get_DateTime(), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]), gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]))
}

func Call_modifyTime(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Constructor2("DateTime", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1]))
}

func Call_modifyDateF(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
__local_var_3_0 := (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]
_ = __local_var_3_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("DateTime", a_4, __local_var_3_0)
}), gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]))
}

func Call_modifyDate(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Constructor2("DateTime", gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1])
}

func Call_diff(dictDuration_0_loop gopurs_runtime.Value, dt1_1_loop gopurs_runtime.Value, dt2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 gopurs_runtime.Value = dictDuration_0_loop
_ = dictDuration_0
var dt1_1 gopurs_runtime.Value = dt1_1_loop
_ = dt1_1
var dt2_2 gopurs_runtime.Value = dt2_2_loop
_ = dt2_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "toDuration"), gopurs_runtime.UncurriedApp2(Get_calcDiff(), gopurs_runtime.Apply(Get_toRecord(), dt1_1), gopurs_runtime.Apply(Get_toRecord(), dt2_2)))
}

func Call_adjust(dictDuration_0_loop gopurs_runtime.Value, d_1_loop gopurs_runtime.Value, dt_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 gopurs_runtime.Value = dictDuration_0_loop
_ = dictDuration_0
var d_1 gopurs_runtime.Value = d_1_loop
_ = d_1
var dt_2 gopurs_runtime.Value = dt_2_loop
_ = dt_2
__local_var_3_0 := gopurs_runtime.Apply4(Get_adjustImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "fromDuration"), d_1), gopurs_runtime.Apply(Get_toRecord(), dt_2))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just").IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "year").IntVal >= -271820 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "year").IntVal <= 275759 {
__t3 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "year"))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
__local_var_4_2 := __t3
_ = __local_var_4_2
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_4_2.StrVal == "Just").IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(pkg_Data_Date.Get_exactDate(), (*[1024]gopurs_runtime.Value)(__local_var_4_2.UnsafePtr)[0]))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_5:
__local_var_5_4 := __t5
_ = __local_var_5_4
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "month").IntVal == 1 {
var __t7 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
var __t8 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal >= 1 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal <= 31 {
__local_var_6_9 := gopurs_runtime.Apply2((*[1024]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)[0], gopurs_runtime.Constructor0("January"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day"))
_ = __local_var_6_9
var __t10 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_9.StrVal == "Just").IntVal != 0 {
__t10 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7.StrVal == "Just").IntVal != 0 {
__t11 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("DateTime", (*[1024]gopurs_runtime.Value)(__local_var_6_9.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0]))
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_11:
return __t11
})
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
})
}
end_branch_10:
var __t12 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal <= 23 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal <= 59 {
var __t13 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal <= 59 {
var __t14 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t14 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor4("Time", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond")))
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_14:
__t13 = __t14
goto end_branch_13
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t13 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_13:
__t12 = __t13
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_12:
__t8 = gopurs_runtime.Apply(__t10, __t12)
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
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Nothing").IntVal != 0 {
__t7 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t6 = __t7
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "month").IntVal == 2 {
var __t15 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
var __t16 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal >= 1 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal <= 31 {
__local_var_6_17 := gopurs_runtime.Apply2((*[1024]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)[0], gopurs_runtime.Constructor0("February"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day"))
_ = __local_var_6_17
var __t18 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_17.StrVal == "Just").IntVal != 0 {
__t18 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7.StrVal == "Just").IntVal != 0 {
__t19 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("DateTime", (*[1024]gopurs_runtime.Value)(__local_var_6_17.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0]))
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_19:
return __t19
})
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
})
}
end_branch_18:
var __t20 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal <= 23 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal <= 59 {
var __t21 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal <= 59 {
var __t22 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t22 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor4("Time", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond")))
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_22:
__t21 = __t22
goto end_branch_21
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t21 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_21
} else {

}
}
{
__t21 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_21:
__t20 = __t21
goto end_branch_20
} else {

}
}
{
__t20 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_20:
__t16 = gopurs_runtime.Apply(__t18, __t20)
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_16:
__t15 = __t16
goto end_branch_15
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Nothing").IntVal != 0 {
__t15 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
__t6 = __t15
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "month").IntVal == 3 {
var __t23 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
var __t24 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal >= 1 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal <= 31 {
__local_var_6_25 := gopurs_runtime.Apply2((*[1024]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)[0], gopurs_runtime.Constructor0("March"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day"))
_ = __local_var_6_25
var __t26 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_25.StrVal == "Just").IntVal != 0 {
__t26 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7.StrVal == "Just").IntVal != 0 {
__t27 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("DateTime", (*[1024]gopurs_runtime.Value)(__local_var_6_25.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0]))
goto end_branch_27
} else {

}
}
{
__t27 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_27:
return __t27
})
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
})
}
end_branch_26:
var __t28 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal <= 23 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal <= 59 {
var __t29 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal <= 59 {
var __t30 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t30 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor4("Time", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond")))
goto end_branch_30
} else {

}
}
{
__t30 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_30:
__t29 = __t30
goto end_branch_29
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t29 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_29
} else {

}
}
{
__t29 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_29:
__t28 = __t29
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_28:
__t24 = gopurs_runtime.Apply(__t26, __t28)
goto end_branch_24
} else {

}
}
{
__t24 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_24:
__t23 = __t24
goto end_branch_23
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Nothing").IntVal != 0 {
__t23 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t6 = __t23
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "month").IntVal == 4 {
var __t31 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
var __t32 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal >= 1 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal <= 31 {
__local_var_6_33 := gopurs_runtime.Apply2((*[1024]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)[0], gopurs_runtime.Constructor0("April"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day"))
_ = __local_var_6_33
var __t34 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_33.StrVal == "Just").IntVal != 0 {
__t34 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7.StrVal == "Just").IntVal != 0 {
__t35 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("DateTime", (*[1024]gopurs_runtime.Value)(__local_var_6_33.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0]))
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_35:
return __t35
})
goto end_branch_34
} else {

}
}
{
__t34 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
})
}
end_branch_34:
var __t36 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal <= 23 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal <= 59 {
var __t37 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal <= 59 {
var __t38 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t38 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor4("Time", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond")))
goto end_branch_38
} else {

}
}
{
__t38 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_38:
__t37 = __t38
goto end_branch_37
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t37 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_37
} else {

}
}
{
__t37 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_37:
__t36 = __t37
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_36:
__t32 = gopurs_runtime.Apply(__t34, __t36)
goto end_branch_32
} else {

}
}
{
__t32 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_32:
__t31 = __t32
goto end_branch_31
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Nothing").IntVal != 0 {
__t31 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
__t6 = __t31
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "month").IntVal == 5 {
var __t39 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
var __t40 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal >= 1 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal <= 31 {
__local_var_6_41 := gopurs_runtime.Apply2((*[1024]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)[0], gopurs_runtime.Constructor0("May"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day"))
_ = __local_var_6_41
var __t42 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_41.StrVal == "Just").IntVal != 0 {
__t42 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t43 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7.StrVal == "Just").IntVal != 0 {
__t43 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("DateTime", (*[1024]gopurs_runtime.Value)(__local_var_6_41.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0]))
goto end_branch_43
} else {

}
}
{
__t43 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_43:
return __t43
})
goto end_branch_42
} else {

}
}
{
__t42 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
})
}
end_branch_42:
var __t44 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal <= 23 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal <= 59 {
var __t45 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal <= 59 {
var __t46 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t46 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor4("Time", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond")))
goto end_branch_46
} else {

}
}
{
__t46 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_46:
__t45 = __t46
goto end_branch_45
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t45 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_45
} else {

}
}
{
__t45 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_45:
__t44 = __t45
goto end_branch_44
} else {

}
}
{
__t44 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_44:
__t40 = gopurs_runtime.Apply(__t42, __t44)
goto end_branch_40
} else {

}
}
{
__t40 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_40:
__t39 = __t40
goto end_branch_39
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Nothing").IntVal != 0 {
__t39 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_39
} else {

}
}
{
__t39 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_39:
__t6 = __t39
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "month").IntVal == 6 {
var __t47 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
var __t48 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal >= 1 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal <= 31 {
__local_var_6_49 := gopurs_runtime.Apply2((*[1024]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)[0], gopurs_runtime.Constructor0("June"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day"))
_ = __local_var_6_49
var __t50 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_49.StrVal == "Just").IntVal != 0 {
__t50 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t51 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7.StrVal == "Just").IntVal != 0 {
__t51 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("DateTime", (*[1024]gopurs_runtime.Value)(__local_var_6_49.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0]))
goto end_branch_51
} else {

}
}
{
__t51 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_51:
return __t51
})
goto end_branch_50
} else {

}
}
{
__t50 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
})
}
end_branch_50:
var __t52 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal <= 23 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal <= 59 {
var __t53 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal <= 59 {
var __t54 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t54 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor4("Time", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond")))
goto end_branch_54
} else {

}
}
{
__t54 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_54:
__t53 = __t54
goto end_branch_53
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t53 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_53
} else {

}
}
{
__t53 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_53:
__t52 = __t53
goto end_branch_52
} else {

}
}
{
__t52 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_52:
__t48 = gopurs_runtime.Apply(__t50, __t52)
goto end_branch_48
} else {

}
}
{
__t48 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_48:
__t47 = __t48
goto end_branch_47
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Nothing").IntVal != 0 {
__t47 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_47
} else {

}
}
{
__t47 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_47:
__t6 = __t47
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "month").IntVal == 7 {
var __t55 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
var __t56 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal >= 1 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal <= 31 {
__local_var_6_57 := gopurs_runtime.Apply2((*[1024]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)[0], gopurs_runtime.Constructor0("July"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day"))
_ = __local_var_6_57
var __t58 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_57.StrVal == "Just").IntVal != 0 {
__t58 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t59 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7.StrVal == "Just").IntVal != 0 {
__t59 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("DateTime", (*[1024]gopurs_runtime.Value)(__local_var_6_57.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0]))
goto end_branch_59
} else {

}
}
{
__t59 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_59:
return __t59
})
goto end_branch_58
} else {

}
}
{
__t58 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
})
}
end_branch_58:
var __t60 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal <= 23 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal <= 59 {
var __t61 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal <= 59 {
var __t62 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t62 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor4("Time", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond")))
goto end_branch_62
} else {

}
}
{
__t62 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_62:
__t61 = __t62
goto end_branch_61
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t61 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_61
} else {

}
}
{
__t61 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_61:
__t60 = __t61
goto end_branch_60
} else {

}
}
{
__t60 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_60:
__t56 = gopurs_runtime.Apply(__t58, __t60)
goto end_branch_56
} else {

}
}
{
__t56 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_56:
__t55 = __t56
goto end_branch_55
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Nothing").IntVal != 0 {
__t55 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_55
} else {

}
}
{
__t55 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_55:
__t6 = __t55
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "month").IntVal == 8 {
var __t63 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
var __t64 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal >= 1 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal <= 31 {
__local_var_6_65 := gopurs_runtime.Apply2((*[1024]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)[0], gopurs_runtime.Constructor0("August"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day"))
_ = __local_var_6_65
var __t66 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_65.StrVal == "Just").IntVal != 0 {
__t66 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t67 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7.StrVal == "Just").IntVal != 0 {
__t67 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("DateTime", (*[1024]gopurs_runtime.Value)(__local_var_6_65.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0]))
goto end_branch_67
} else {

}
}
{
__t67 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_67:
return __t67
})
goto end_branch_66
} else {

}
}
{
__t66 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
})
}
end_branch_66:
var __t68 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal <= 23 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal <= 59 {
var __t69 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal <= 59 {
var __t70 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t70 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor4("Time", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond")))
goto end_branch_70
} else {

}
}
{
__t70 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_70:
__t69 = __t70
goto end_branch_69
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t69 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_69
} else {

}
}
{
__t69 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_69:
__t68 = __t69
goto end_branch_68
} else {

}
}
{
__t68 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_68:
__t64 = gopurs_runtime.Apply(__t66, __t68)
goto end_branch_64
} else {

}
}
{
__t64 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_64:
__t63 = __t64
goto end_branch_63
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Nothing").IntVal != 0 {
__t63 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_63
} else {

}
}
{
__t63 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_63:
__t6 = __t63
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "month").IntVal == 9 {
var __t71 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
var __t72 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal >= 1 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal <= 31 {
__local_var_6_73 := gopurs_runtime.Apply2((*[1024]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)[0], gopurs_runtime.Constructor0("September"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day"))
_ = __local_var_6_73
var __t74 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_73.StrVal == "Just").IntVal != 0 {
__t74 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t75 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7.StrVal == "Just").IntVal != 0 {
__t75 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("DateTime", (*[1024]gopurs_runtime.Value)(__local_var_6_73.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0]))
goto end_branch_75
} else {

}
}
{
__t75 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_75:
return __t75
})
goto end_branch_74
} else {

}
}
{
__t74 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
})
}
end_branch_74:
var __t76 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal <= 23 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal <= 59 {
var __t77 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal <= 59 {
var __t78 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t78 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor4("Time", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond")))
goto end_branch_78
} else {

}
}
{
__t78 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_78:
__t77 = __t78
goto end_branch_77
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t77 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_77
} else {

}
}
{
__t77 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_77:
__t76 = __t77
goto end_branch_76
} else {

}
}
{
__t76 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_76:
__t72 = gopurs_runtime.Apply(__t74, __t76)
goto end_branch_72
} else {

}
}
{
__t72 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_72:
__t71 = __t72
goto end_branch_71
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Nothing").IntVal != 0 {
__t71 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_71
} else {

}
}
{
__t71 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_71:
__t6 = __t71
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "month").IntVal == 10 {
var __t79 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
var __t80 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal >= 1 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal <= 31 {
__local_var_6_81 := gopurs_runtime.Apply2((*[1024]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)[0], gopurs_runtime.Constructor0("October"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day"))
_ = __local_var_6_81
var __t82 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_81.StrVal == "Just").IntVal != 0 {
__t82 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t83 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7.StrVal == "Just").IntVal != 0 {
__t83 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("DateTime", (*[1024]gopurs_runtime.Value)(__local_var_6_81.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0]))
goto end_branch_83
} else {

}
}
{
__t83 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_83:
return __t83
})
goto end_branch_82
} else {

}
}
{
__t82 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
})
}
end_branch_82:
var __t84 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal <= 23 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal <= 59 {
var __t85 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal <= 59 {
var __t86 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t86 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor4("Time", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond")))
goto end_branch_86
} else {

}
}
{
__t86 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_86:
__t85 = __t86
goto end_branch_85
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t85 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_85
} else {

}
}
{
__t85 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_85:
__t84 = __t85
goto end_branch_84
} else {

}
}
{
__t84 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_84:
__t80 = gopurs_runtime.Apply(__t82, __t84)
goto end_branch_80
} else {

}
}
{
__t80 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_80:
__t79 = __t80
goto end_branch_79
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Nothing").IntVal != 0 {
__t79 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_79
} else {

}
}
{
__t79 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_79:
__t6 = __t79
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "month").IntVal == 11 {
var __t87 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
var __t88 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal >= 1 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal <= 31 {
__local_var_6_89 := gopurs_runtime.Apply2((*[1024]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)[0], gopurs_runtime.Constructor0("November"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day"))
_ = __local_var_6_89
var __t90 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_89.StrVal == "Just").IntVal != 0 {
__t90 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t91 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7.StrVal == "Just").IntVal != 0 {
__t91 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("DateTime", (*[1024]gopurs_runtime.Value)(__local_var_6_89.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0]))
goto end_branch_91
} else {

}
}
{
__t91 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_91:
return __t91
})
goto end_branch_90
} else {

}
}
{
__t90 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
})
}
end_branch_90:
var __t92 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal <= 23 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal <= 59 {
var __t93 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal <= 59 {
var __t94 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t94 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor4("Time", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond")))
goto end_branch_94
} else {

}
}
{
__t94 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_94:
__t93 = __t94
goto end_branch_93
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t93 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_93
} else {

}
}
{
__t93 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_93:
__t92 = __t93
goto end_branch_92
} else {

}
}
{
__t92 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_92:
__t88 = gopurs_runtime.Apply(__t90, __t92)
goto end_branch_88
} else {

}
}
{
__t88 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_88:
__t87 = __t88
goto end_branch_87
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Nothing").IntVal != 0 {
__t87 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_87
} else {

}
}
{
__t87 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_87:
__t6 = __t87
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "month").IntVal == 12 {
var __t95 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
var __t96 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal >= 1 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day").IntVal <= 31 {
__local_var_6_97 := gopurs_runtime.Apply2((*[1024]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)[0], gopurs_runtime.Constructor0("December"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "day"))
_ = __local_var_6_97
var __t98 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_97.StrVal == "Just").IntVal != 0 {
__t98 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t99 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_7.StrVal == "Just").IntVal != 0 {
__t99 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("DateTime", (*[1024]gopurs_runtime.Value)(__local_var_6_97.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_7.UnsafePtr)[0]))
goto end_branch_99
} else {

}
}
{
__t99 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_99:
return __t99
})
goto end_branch_98
} else {

}
}
{
__t98 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
})
}
end_branch_98:
var __t100 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour").IntVal <= 23 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute").IntVal <= 59 {
var __t101 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second").IntVal <= 59 {
var __t102 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t102 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor4("Time", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "hour"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "minute"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "second"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond")))
goto end_branch_102
} else {

}
}
{
__t102 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_102:
__t101 = __t102
goto end_branch_101
} else {

}
}
{
if gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal >= 0 && gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], "millisecond").IntVal <= 999 {
__t101 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_101
} else {

}
}
{
__t101 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_101:
__t100 = __t101
goto end_branch_100
} else {

}
}
{
__t100 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_100:
__t96 = gopurs_runtime.Apply(__t98, __t100)
goto end_branch_96
} else {

}
}
{
__t96 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_96:
__t95 = __t96
goto end_branch_95
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Nothing").IntVal != 0 {
__t95 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_95
} else {

}
}
{
__t95 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_95:
__t6 = __t95
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
__t6 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_6
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Nothing").IntVal != 0 {
__t6 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t1 = __t6
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
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

func Get_adjustImpl() gopurs_runtime.Value {
	return _Gopurs_AdjustImpl
}

func Get_calcDiff() gopurs_runtime.Value {
	return _Gopurs_CalcDiff
}

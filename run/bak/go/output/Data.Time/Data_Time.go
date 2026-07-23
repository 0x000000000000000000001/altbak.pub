package Data_Time

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_Ord "gopurs/output/Data.Ord"
)

var Time gopurs_runtime.Value
var once_Time sync.Once
func Get_Time() gopurs_runtime.Value {
	once_Time.Do(func() {
		Time = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor4("Time", value0, value1, value2, value3)
})
})
})
})
	})
	return Time
}

var showTime gopurs_runtime.Value
var once_showTime sync.Once
func Get_showTime() gopurs_runtime.Value {
	once_showTime.Do(func() {
		showTime = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Time (Hour ").StrVal + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.ConstructorGet(v_0, 0)).StrVal).StrVal + gopurs_runtime.Str(") (Minute ").StrVal).StrVal + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.ConstructorGet(v_0, 1)).StrVal).StrVal + gopurs_runtime.Str(") (Second ").StrVal).StrVal + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.ConstructorGet(v_0, 2)).StrVal).StrVal + gopurs_runtime.Str(") (Millisecond ").StrVal).StrVal + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.ConstructorGet(v_0, 3)).StrVal).StrVal + gopurs_runtime.Str("))").StrVal)
}))
	})
	return showTime
}

var setSecond gopurs_runtime.Value
var once_setSecond sync.Once
func Get_setSecond() gopurs_runtime.Value {
	once_setSecond.Do(func() {
		setSecond = gopurs_runtime.Func2(func(s_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor4("Time", gopurs_runtime.ConstructorGet(v_1, 0), gopurs_runtime.ConstructorGet(v_1, 1), s_0, gopurs_runtime.ConstructorGet(v_1, 3))
})
	})
	return setSecond
}

var setMinute gopurs_runtime.Value
var once_setMinute sync.Once
func Get_setMinute() gopurs_runtime.Value {
	once_setMinute.Do(func() {
		setMinute = gopurs_runtime.Func2(func(m_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor4("Time", gopurs_runtime.ConstructorGet(v_1, 0), m_0, gopurs_runtime.ConstructorGet(v_1, 2), gopurs_runtime.ConstructorGet(v_1, 3))
})
	})
	return setMinute
}

var setMillisecond gopurs_runtime.Value
var once_setMillisecond sync.Once
func Get_setMillisecond() gopurs_runtime.Value {
	once_setMillisecond.Do(func() {
		setMillisecond = gopurs_runtime.Func2(func(ms_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor4("Time", gopurs_runtime.ConstructorGet(v_1, 0), gopurs_runtime.ConstructorGet(v_1, 1), gopurs_runtime.ConstructorGet(v_1, 2), ms_0)
})
	})
	return setMillisecond
}

var setHour gopurs_runtime.Value
var once_setHour sync.Once
func Get_setHour() gopurs_runtime.Value {
	once_setHour.Do(func() {
		setHour = gopurs_runtime.Func2(func(h_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor4("Time", h_0, gopurs_runtime.ConstructorGet(v_1, 1), gopurs_runtime.ConstructorGet(v_1, 2), gopurs_runtime.ConstructorGet(v_1, 3))
})
	})
	return setHour
}

var second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		second = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.ConstructorGet(v_0, 2)
})
	})
	return second
}

var minute gopurs_runtime.Value
var once_minute sync.Once
func Get_minute() gopurs_runtime.Value {
	once_minute.Do(func() {
		minute = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.ConstructorGet(v_0, 1)
})
	})
	return minute
}

var millisecond gopurs_runtime.Value
var once_millisecond sync.Once
func Get_millisecond() gopurs_runtime.Value {
	once_millisecond.Do(func() {
		millisecond = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.ConstructorGet(v_0, 3)
})
	})
	return millisecond
}

var millisToTime gopurs_runtime.Value
var once_millisToTime sync.Once
func Get_millisToTime() gopurs_runtime.Value {
	once_millisToTime.Do(func() {
		millisToTime = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
hours_1_0 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.FloatDiv(v_0, gopurs_runtime.Float(3600000.0)))
_ = hours_1_0
minutes_2_1 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.FloatDiv(gopurs_runtime.FloatSub(v_0, gopurs_runtime.FloatMul(hours_1_0, gopurs_runtime.Float(3600000.0))), gopurs_runtime.Float(60000.0)))
_ = minutes_2_1
seconds_3_2 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.FloatDiv(gopurs_runtime.FloatSub(v_0, gopurs_runtime.FloatAdd(gopurs_runtime.FloatMul(hours_1_0, gopurs_runtime.Float(3600000.0)), gopurs_runtime.FloatMul(minutes_2_1, gopurs_runtime.Float(60000.0)))), gopurs_runtime.Float(1000.0)))
_ = seconds_3_2
__local_var_4_3 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), hours_1_0))
_ = __local_var_4_3
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_4_3.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_4_3.IntVal <= gopurs_runtime.Int(23).IntVal).IntVal != 0)).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", __local_var_4_3)
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
var __t20 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just")).IntVal != 0 {
__local_var_6_21 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), minutes_2_1))
_ = __local_var_6_21
var __t28 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_6_21.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_6_21.IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
__local_var_7_29 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), seconds_3_2))
_ = __local_var_7_29
var __t32 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_7_29.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_7_29.IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
__local_var_8_33 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.FloatSub(v_0, gopurs_runtime.FloatAdd(gopurs_runtime.FloatAdd(gopurs_runtime.FloatMul(hours_1_0, gopurs_runtime.Float(3600000.0)), gopurs_runtime.FloatMul(minutes_2_1, gopurs_runtime.Float(60000.0))), gopurs_runtime.FloatMul(seconds_3_2, gopurs_runtime.Float(1000.0))))))
_ = __local_var_8_33
var __t34 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_8_33.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_8_33.IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
__t34 = gopurs_runtime.Constructor4("Time", gopurs_runtime.ConstructorGet(__local_var_5_4, 0), __local_var_6_21, __local_var_7_29, __local_var_8_33)
goto end_branch_34
} else {

}
}
{
__t34 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_34:
__t32 = __t34
goto end_branch_32
} else {

}
}
{
__local_var_8_30 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.FloatSub(v_0, gopurs_runtime.FloatAdd(gopurs_runtime.FloatAdd(gopurs_runtime.FloatMul(hours_1_0, gopurs_runtime.Float(3600000.0)), gopurs_runtime.FloatMul(minutes_2_1, gopurs_runtime.Float(60000.0))), gopurs_runtime.FloatMul(seconds_3_2, gopurs_runtime.Float(1000.0))))))
_ = __local_var_8_30
var __t31 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_8_30.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_8_30.IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
__t32 = __t31
}
end_branch_32:
__t28 = __t32
goto end_branch_28
} else {

}
}
{
__local_var_7_22 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), seconds_3_2))
_ = __local_var_7_22
var __t25 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_7_22.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_7_22.IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
__local_var_8_26 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.FloatSub(v_0, gopurs_runtime.FloatAdd(gopurs_runtime.FloatAdd(gopurs_runtime.FloatMul(hours_1_0, gopurs_runtime.Float(3600000.0)), gopurs_runtime.FloatMul(minutes_2_1, gopurs_runtime.Float(60000.0))), gopurs_runtime.FloatMul(seconds_3_2, gopurs_runtime.Float(1000.0))))))
_ = __local_var_8_26
var __t27 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_8_26.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_8_26.IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
__t25 = __t27
goto end_branch_25
} else {

}
}
{
__local_var_8_23 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.FloatSub(v_0, gopurs_runtime.FloatAdd(gopurs_runtime.FloatAdd(gopurs_runtime.FloatMul(hours_1_0, gopurs_runtime.Float(3600000.0)), gopurs_runtime.FloatMul(minutes_2_1, gopurs_runtime.Float(60000.0))), gopurs_runtime.FloatMul(seconds_3_2, gopurs_runtime.Float(1000.0))))))
_ = __local_var_8_23
var __t24 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_8_23.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_8_23.IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
__t25 = __t24
}
end_branch_25:
__t28 = __t25
}
end_branch_28:
__t20 = __t28
goto end_branch_20
} else {

}
}
{
__local_var_6_6 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), minutes_2_1))
_ = __local_var_6_6
var __t13 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_6_6.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_6_6.IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
__local_var_7_14 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), seconds_3_2))
_ = __local_var_7_14
var __t17 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_7_14.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_7_14.IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
__local_var_8_18 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.FloatSub(v_0, gopurs_runtime.FloatAdd(gopurs_runtime.FloatAdd(gopurs_runtime.FloatMul(hours_1_0, gopurs_runtime.Float(3600000.0)), gopurs_runtime.FloatMul(minutes_2_1, gopurs_runtime.Float(60000.0))), gopurs_runtime.FloatMul(seconds_3_2, gopurs_runtime.Float(1000.0))))))
_ = __local_var_8_18
var __t19 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_8_18.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_8_18.IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
__t17 = __t19
goto end_branch_17
} else {

}
}
{
__local_var_8_15 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.FloatSub(v_0, gopurs_runtime.FloatAdd(gopurs_runtime.FloatAdd(gopurs_runtime.FloatMul(hours_1_0, gopurs_runtime.Float(3600000.0)), gopurs_runtime.FloatMul(minutes_2_1, gopurs_runtime.Float(60000.0))), gopurs_runtime.FloatMul(seconds_3_2, gopurs_runtime.Float(1000.0))))))
_ = __local_var_8_15
var __t16 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_8_15.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_8_15.IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t17 = __t16
}
end_branch_17:
__t13 = __t17
goto end_branch_13
} else {

}
}
{
__local_var_7_7 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), seconds_3_2))
_ = __local_var_7_7
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_7_7.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_7_7.IntVal <= gopurs_runtime.Int(59).IntVal).IntVal != 0)).IntVal != 0 {
__local_var_8_11 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.FloatSub(v_0, gopurs_runtime.FloatAdd(gopurs_runtime.FloatAdd(gopurs_runtime.FloatMul(hours_1_0, gopurs_runtime.Float(3600000.0)), gopurs_runtime.FloatMul(minutes_2_1, gopurs_runtime.Float(60000.0))), gopurs_runtime.FloatMul(seconds_3_2, gopurs_runtime.Float(1000.0))))))
_ = __local_var_8_11
var __t12 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_8_11.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_8_11.IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
__t10 = __t12
goto end_branch_10
} else {

}
}
{
__local_var_8_8 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.FloatSub(v_0, gopurs_runtime.FloatAdd(gopurs_runtime.FloatAdd(gopurs_runtime.FloatMul(hours_1_0, gopurs_runtime.Float(3600000.0)), gopurs_runtime.FloatMul(minutes_2_1, gopurs_runtime.Float(60000.0))), gopurs_runtime.FloatMul(seconds_3_2, gopurs_runtime.Float(1000.0))))))
_ = __local_var_8_8
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_8_8.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_8_8.IntVal <= gopurs_runtime.Int(999).IntVal).IntVal != 0)).IntVal != 0 {
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
__t10 = __t9
}
end_branch_10:
__t13 = __t10
}
end_branch_13:
__t20 = __t13
}
end_branch_20:
return __t20
})
	})
	return millisToTime
}

var hour gopurs_runtime.Value
var once_hour sync.Once
func Get_hour() gopurs_runtime.Value {
	once_hour.Do(func() {
		hour = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.ConstructorGet(v_0, 0)
})
	})
	return hour
}

var timeToMillis gopurs_runtime.Value
var once_timeToMillis sync.Once
func Get_timeToMillis() gopurs_runtime.Value {
	once_timeToMillis.Do(func() {
		timeToMillis = gopurs_runtime.Func(func(t_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.FloatAdd(gopurs_runtime.FloatAdd(gopurs_runtime.FloatAdd(gopurs_runtime.FloatMul(gopurs_runtime.Float(3600000.0), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.ConstructorGet(t_0, 0))), gopurs_runtime.FloatMul(gopurs_runtime.Float(60000.0), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.ConstructorGet(t_0, 1)))), gopurs_runtime.FloatMul(gopurs_runtime.Float(1000.0), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.ConstructorGet(t_0, 2)))), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.ConstructorGet(t_0, 3)))
})
	})
	return timeToMillis
}

var eqTime gopurs_runtime.Value
var once_eqTime sync.Once
func Get_eqTime() gopurs_runtime.Value {
	once_eqTime.Do(func() {
		eqTime = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(x_0, 0).IntVal == gopurs_runtime.ConstructorGet(y_1, 0).IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(x_0, 1).IntVal == gopurs_runtime.ConstructorGet(y_1, 1).IntVal).IntVal != 0).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(x_0, 2).IntVal == gopurs_runtime.ConstructorGet(y_1, 2).IntVal).IntVal != 0).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.ConstructorGet(x_0, 3).IntVal == gopurs_runtime.ConstructorGet(y_1, 3).IntVal).IntVal != 0)
}))
	})
	return eqTime
}

var ordTime gopurs_runtime.Value
var once_ordTime sync.Once
func Get_ordTime() gopurs_runtime.Value {
	once_ordTime.Do(func() {
		ordTime = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), gopurs_runtime.ConstructorGet(x_0, 0), gopurs_runtime.ConstructorGet(y_1, 0))
_ = v_2_0
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2_0.StrVal == "LT")).IntVal != 0 {
__t5 = gopurs_runtime.Constructor0("LT")
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(v_2_0.StrVal == "GT")).IntVal != 0 {
__t5 = gopurs_runtime.Constructor0("GT")
goto end_branch_5
} else {

}
}
{
v1_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), gopurs_runtime.ConstructorGet(x_0, 1), gopurs_runtime.ConstructorGet(y_1, 1))
_ = v1_3_1
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_3_1.StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("LT")
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v1_3_1.StrVal == "GT")).IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("GT")
goto end_branch_4
} else {

}
}
{
v2_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), gopurs_runtime.ConstructorGet(x_0, 2), gopurs_runtime.ConstructorGet(y_1, 2))
_ = v2_4_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_4_2.StrVal == "LT")).IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("LT")
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v2_4_2.StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("GT")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), gopurs_runtime.ConstructorGet(x_0, 3), gopurs_runtime.ConstructorGet(y_1, 3))
}
end_branch_3:
__t4 = __t3
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
return __t5
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqTime()
}))
	})
	return ordTime
}

var diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		diff = gopurs_runtime.Func3(func(dictDuration_0 gopurs_runtime.Value, t1_1 gopurs_runtime.Value, t2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "toDuration"), gopurs_runtime.FloatAdd(gopurs_runtime.Apply(Get_timeToMillis(), t1_1), gopurs_runtime.FloatNeg(gopurs_runtime.Apply(Get_timeToMillis(), t2_2))))
})
	})
	return diff
}

var boundedTime gopurs_runtime.Value
var once_boundedTime sync.Once
func Get_boundedTime() gopurs_runtime.Value {
	once_boundedTime.Do(func() {
		boundedTime = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Constructor4("Time", gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0)), gopurs_runtime.Constructor4("Time", gopurs_runtime.Int(23), gopurs_runtime.Int(59), gopurs_runtime.Int(59), gopurs_runtime.Int(999)), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordTime()
}))
	})
	return boundedTime
}

var maxTime gopurs_runtime.Value
var once_maxTime sync.Once
func Get_maxTime() gopurs_runtime.Value {
	once_maxTime.Do(func() {
		maxTime = gopurs_runtime.Apply(Get_timeToMillis(), gopurs_runtime.Constructor4("Time", gopurs_runtime.Int(23), gopurs_runtime.Int(59), gopurs_runtime.Int(59), gopurs_runtime.Int(999)))
	})
	return maxTime
}

var minTime gopurs_runtime.Value
var once_minTime sync.Once
func Get_minTime() gopurs_runtime.Value {
	once_minTime.Do(func() {
		minTime = gopurs_runtime.Apply(Get_timeToMillis(), gopurs_runtime.Constructor4("Time", gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0), gopurs_runtime.Int(0)))
	})
	return minTime
}

var adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		adjust = gopurs_runtime.Func3(func(dictDuration_0 gopurs_runtime.Value, d_1 gopurs_runtime.Value, t_2 gopurs_runtime.Value) gopurs_runtime.Value {
d_prime_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "fromDuration"), d_1)
_ = d_prime_3_0
wholeDays_4_1 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.FloatDiv(d_prime_3_0, gopurs_runtime.Float(86400000.0)))
_ = wholeDays_4_1
msAdjusted_5_2 := gopurs_runtime.FloatAdd(gopurs_runtime.FloatAdd(gopurs_runtime.Apply(Get_timeToMillis(), t_2), d_prime_3_0), gopurs_runtime.FloatNeg(gopurs_runtime.FloatMul(wholeDays_4_1, gopurs_runtime.Float(86400000.0))))
_ = msAdjusted_5_2
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.FloatGt(msAdjusted_5_2, Get_maxTime())).IntVal != 0 {
__t4 = gopurs_runtime.Float(1.0)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.FloatLt(msAdjusted_5_2, Get_minTime())).IntVal != 0 {
__t4 = gopurs_runtime.Float(-1.0)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Float(0.0)
}
end_branch_4:
wrap_6_3 := __t4
_ = wrap_6_3
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.FloatAdd(wholeDays_4_1, wrap_6_3), gopurs_runtime.Apply(Get_millisToTime(), gopurs_runtime.FloatAdd(msAdjusted_5_2, gopurs_runtime.FloatMul(gopurs_runtime.Float(86400000.0), gopurs_runtime.FloatNeg(wrap_6_3)))))
})
	})
	return adjust
}



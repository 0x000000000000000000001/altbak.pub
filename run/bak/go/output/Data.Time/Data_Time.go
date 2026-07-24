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
return gopurs_runtime.Str("(Time (Hour " + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]).StrVal + ") (Minute " + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]).StrVal + ") (Second " + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[2]).StrVal + ") (Millisecond " + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[3]).StrVal + "))")
}))
	})
	return showTime
}

var setSecond gopurs_runtime.Value
var once_setSecond sync.Once
func Get_setSecond() gopurs_runtime.Value {
	once_setSecond.Do(func() {
		setSecond = gopurs_runtime.Func2(func(s_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor4("Time", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1], s_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[3])
})
	})
	return setSecond
}

var setMinute gopurs_runtime.Value
var once_setMinute sync.Once
func Get_setMinute() gopurs_runtime.Value {
	once_setMinute.Do(func() {
		setMinute = gopurs_runtime.Func2(func(m_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor4("Time", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], m_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[3])
})
	})
	return setMinute
}

var setMillisecond gopurs_runtime.Value
var once_setMillisecond sync.Once
func Get_setMillisecond() gopurs_runtime.Value {
	once_setMillisecond.Do(func() {
		setMillisecond = gopurs_runtime.Func2(func(ms_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor4("Time", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[2], ms_0)
})
	})
	return setMillisecond
}

var setHour gopurs_runtime.Value
var once_setHour sync.Once
func Get_setHour() gopurs_runtime.Value {
	once_setHour.Do(func() {
		setHour = gopurs_runtime.Func2(func(h_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor4("Time", h_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[3])
})
	})
	return setHour
}

var second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		second = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[2]
})
	})
	return second
}

var minute gopurs_runtime.Value
var once_minute sync.Once
func Get_minute() gopurs_runtime.Value {
	once_minute.Do(func() {
		minute = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]
})
	})
	return minute
}

var millisecond gopurs_runtime.Value
var once_millisecond sync.Once
func Get_millisecond() gopurs_runtime.Value {
	once_millisecond.Do(func() {
		millisecond = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[3]
})
	})
	return millisecond
}

var millisToTime gopurs_runtime.Value
var once_millisToTime sync.Once
func Get_millisToTime() gopurs_runtime.Value {
	once_millisToTime.Do(func() {
		millisToTime = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
hours_1_0 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(v_0.FloatVal() / 3600000.0))
_ = hours_1_0
minutes_2_1 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(v_0.FloatVal() - hours_1_0.FloatVal() * 3600000.0 / 60000.0))
_ = minutes_2_1
seconds_3_2 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(v_0.FloatVal() - hours_1_0.FloatVal() * 3600000.0 + minutes_2_1.FloatVal() * 60000.0 / 1000.0))
_ = seconds_3_2
__local_var_4_3 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), hours_1_0))
_ = __local_var_4_3
var __t5 gopurs_runtime.Value
{
if __local_var_4_3.IntVal >= 0 && __local_var_4_3.IntVal <= 23 {
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
var __t7 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_5_4.StrVal == "Just").IntVal != 0 {
__t7 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(Get_Time(), (*[1024]gopurs_runtime.Value)(__local_var_5_4.UnsafePtr)[0]))
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_7:
__local_var_6_6 := __t7
_ = __local_var_6_6
__local_var_7_8 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), minutes_2_1))
_ = __local_var_7_8
var __t9 gopurs_runtime.Value
{
if __local_var_7_8.IntVal >= 0 && __local_var_7_8.IntVal <= 59 {
var __t10 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_6_6.StrVal == "Just").IntVal != 0 {
__local_var_8_11 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), seconds_3_2))
_ = __local_var_8_11
var __t14 gopurs_runtime.Value
{
if __local_var_8_11.IntVal >= 0 && __local_var_8_11.IntVal <= 59 {
__local_var_9_15 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(v_0.FloatVal() - hours_1_0.FloatVal() * 3600000.0 + minutes_2_1.FloatVal() * 60000.0 + seconds_3_2.FloatVal() * 1000.0)))
_ = __local_var_9_15
var __t16 gopurs_runtime.Value
{
if __local_var_9_15.IntVal >= 0 && __local_var_9_15.IntVal <= 999 {
__t16 = gopurs_runtime.Apply3((*[1024]gopurs_runtime.Value)(__local_var_6_6.UnsafePtr)[0], __local_var_7_8, __local_var_8_11, __local_var_9_15)
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t14 = __t16
goto end_branch_14
} else {

}
}
{
__local_var_9_12 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(v_0.FloatVal() - hours_1_0.FloatVal() * 3600000.0 + minutes_2_1.FloatVal() * 60000.0 + seconds_3_2.FloatVal() * 1000.0)))
_ = __local_var_9_12
var __t13 gopurs_runtime.Value
{
if __local_var_9_12.IntVal >= 0 && __local_var_9_12.IntVal <= 999 {
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
__t14 = __t13
}
end_branch_14:
__t10 = __t14
goto end_branch_10
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_6_6.StrVal == "Nothing").IntVal != 0 {
__local_var_8_17 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(v_0.FloatVal() - hours_1_0.FloatVal() * 3600000.0 + minutes_2_1.FloatVal() * 60000.0 + seconds_3_2.FloatVal() * 1000.0)))
_ = __local_var_8_17
var __t18 gopurs_runtime.Value
{
if __local_var_8_17.IntVal >= 0 && __local_var_8_17.IntVal <= 999 {
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
__t10 = __t18
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t9 = __t10
goto end_branch_9
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_6_6.StrVal == "Just").IntVal != 0 {
__local_var_8_19 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(v_0.FloatVal() - hours_1_0.FloatVal() * 3600000.0 + minutes_2_1.FloatVal() * 60000.0 + seconds_3_2.FloatVal() * 1000.0)))
_ = __local_var_8_19
var __t20 gopurs_runtime.Value
{
if __local_var_8_19.IntVal >= 0 && __local_var_8_19.IntVal <= 999 {
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
__t9 = __t20
goto end_branch_9
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_6_6.StrVal == "Nothing").IntVal != 0 {
__local_var_8_21 := gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(v_0.FloatVal() - hours_1_0.FloatVal() * 3600000.0 + minutes_2_1.FloatVal() * 60000.0 + seconds_3_2.FloatVal() * 1000.0)))
_ = __local_var_8_21
var __t22 gopurs_runtime.Value
{
if __local_var_8_21.IntVal >= 0 && __local_var_8_21.IntVal <= 999 {
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
__t9 = __t22
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})
	})
	return millisToTime
}

var hour gopurs_runtime.Value
var once_hour sync.Once
func Get_hour() gopurs_runtime.Value {
	once_hour.Do(func() {
		hour = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]
})
	})
	return hour
}

var timeToMillis gopurs_runtime.Value
var once_timeToMillis sync.Once
func Get_timeToMillis() gopurs_runtime.Value {
	once_timeToMillis.Do(func() {
		timeToMillis = gopurs_runtime.Func(func(t_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(3600000.0 * gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), (*[1024]gopurs_runtime.Value)(t_0.UnsafePtr)[0]).FloatVal() + 60000.0 * gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), (*[1024]gopurs_runtime.Value)(t_0.UnsafePtr)[1]).FloatVal() + 1000.0 * gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), (*[1024]gopurs_runtime.Value)(t_0.UnsafePtr)[2]).FloatVal() + gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), (*[1024]gopurs_runtime.Value)(t_0.UnsafePtr)[3]).FloatVal())
})
	})
	return timeToMillis
}

var eqTime gopurs_runtime.Value
var once_eqTime sync.Once
func Get_eqTime() gopurs_runtime.Value {
	once_eqTime.Do(func() {
		eqTime = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0].IntVal == (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0].IntVal && (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1].IntVal == (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1].IntVal && (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[2].IntVal == (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[2].IntVal && (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[3].IntVal == (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[3].IntVal)
}))
	})
	return eqTime
}

var ordTime gopurs_runtime.Value
var once_ordTime sync.Once
func Get_ordTime() gopurs_runtime.Value {
	once_ordTime.Do(func() {
		ordTime = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[0])
_ = v_2_0
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2_0.StrVal == "LT").IntVal != 0 {
__t5 = gopurs_runtime.Constructor0("LT")
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v_2_0.StrVal == "GT").IntVal != 0 {
__t5 = gopurs_runtime.Constructor0("GT")
goto end_branch_5
} else {

}
}
{
v1_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[1])
_ = v1_3_1
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3_1.StrVal == "LT").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("LT")
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(v1_3_1.StrVal == "GT").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("GT")
goto end_branch_4
} else {

}
}
{
v2_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[2], (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[2])
_ = v2_4_2
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_4_2.StrVal == "LT").IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("LT")
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v2_4_2.StrVal == "GT").IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("GT")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[3], (*[1024]gopurs_runtime.Value)(y_1.UnsafePtr)[3])
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
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "toDuration"), gopurs_runtime.Float(gopurs_runtime.Apply(Get_timeToMillis(), t1_1).FloatVal() + gopurs_runtime.FloatNeg(gopurs_runtime.Apply(Get_timeToMillis(), t2_2)).FloatVal()))
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
wholeDays_4_1 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(d_prime_3_0.FloatVal() / 86400000.0))
_ = wholeDays_4_1
msAdjusted_5_2 := gopurs_runtime.Apply(Get_timeToMillis(), t_2).FloatVal() + d_prime_3_0.FloatVal() + gopurs_runtime.FloatNeg(wholeDays_4_1.FloatVal() * 86400000.0).FloatVal()
_ = msAdjusted_5_2
var __t4 gopurs_runtime.Value
{
if msAdjusted_5_2 > Get_maxTime().FloatVal() {
__t4 = gopurs_runtime.Float(1.0)
goto end_branch_4
} else {

}
}
{
if msAdjusted_5_2 < Get_minTime().FloatVal() {
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
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Float(wholeDays_4_1.FloatVal() + wrap_6_3.FloatVal()), gopurs_runtime.Apply(Get_millisToTime(), gopurs_runtime.Float(msAdjusted_5_2 + 86400000.0 * gopurs_runtime.FloatNeg(wrap_6_3).FloatVal())))
})
	})
	return adjust
}





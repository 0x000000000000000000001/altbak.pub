package Data_Time

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Time_Duration "gopurs/output/Data.Time.Duration"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Time_Component "gopurs/output/Data.Time.Component"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var fromJust gopurs_runtime.Value
var once_fromJust sync.Once
func Get_fromJust() gopurs_runtime.Value {
	once_fromJust.Do(func() {
		fromJust = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
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
	return fromJust
}

var negateDuration gopurs_runtime.Value
var once_negateDuration sync.Once
func Get_negateDuration() gopurs_runtime.Value {
	once_negateDuration.Do(func() {
		negateDuration = gopurs_runtime.Apply(pkg_Data_Time_Duration.Get_negateDuration(), pkg_Data_Time_Duration.Get_durationMilliseconds())
	})
	return negateDuration
}

var Time gopurs_runtime.Value
var once_Time sync.Once
func Get_Time() gopurs_runtime.Value {
	once_Time.Do(func() {
		Time = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": value0, "value1": value1, "value2": value2, "value3": value3})
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
		showTime = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Time ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Hour ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Str(")")))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Minute ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Str(")")))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Second ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"])), gopurs_runtime.Str(")")))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str(" ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(Millisecond ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value3"])), gopurs_runtime.Str(")")))), gopurs_runtime.Str(")")))))))))
})})
	})
	return showTime
}

var setSecond gopurs_runtime.Value
var once_setSecond sync.Once
func Get_setSecond() gopurs_runtime.Value {
	once_setSecond.Do(func() {
		setSecond = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": s_0, "value3": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
})
})
	})
	return setSecond
}

var setMinute gopurs_runtime.Value
var once_setMinute sync.Once
func Get_setMinute() gopurs_runtime.Value {
	once_setMinute.Do(func() {
		setMinute = gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": m_0, "value2": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
})
})
	})
	return setMinute
}

var setMillisecond gopurs_runtime.Value
var once_setMillisecond sync.Once
func Get_setMillisecond() gopurs_runtime.Value {
	once_setMillisecond.Do(func() {
		setMillisecond = gopurs_runtime.Func(func(ms_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": ms_0})
})
})
	})
	return setMillisecond
}

var setHour gopurs_runtime.Value
var once_setHour sync.Once
func Get_setHour() gopurs_runtime.Value {
	once_setHour.Do(func() {
		setHour = gopurs_runtime.Func(func(h_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": h_0, "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"], "value3": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"]})
})
})
	})
	return setHour
}

var second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		second = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"]
})
	})
	return second
}

var minute gopurs_runtime.Value
var once_minute sync.Once
func Get_minute() gopurs_runtime.Value {
	once_minute.Do(func() {
		minute = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
	})
	return minute
}

var millisecond gopurs_runtime.Value
var once_millisecond sync.Once
func Get_millisecond() gopurs_runtime.Value {
	once_millisecond.Do(func() {
		millisecond = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value3"]
})
	})
	return millisecond
}

var millisToTime gopurs_runtime.Value
var once_millisToTime sync.Once
func Get_millisToTime() gopurs_runtime.Value {
	once_millisToTime.Do(func() {
		millisToTime = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
hours_1_0 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_numDiv(), v_0), gopurs_runtime.Float(3600000.0)))
minutes_2_1 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_numDiv(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_numSub(), v_0), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numMul(), hours_1_0), gopurs_runtime.Float(3600000.0)))), gopurs_runtime.Float(60000.0)))
seconds_3_2 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_numDiv(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_numSub(), v_0), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numMul(), hours_1_0), gopurs_runtime.Float(3600000.0))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numMul(), minutes_2_1), gopurs_runtime.Float(60000.0))))), gopurs_runtime.Float(1000.0)))
__local_var_4_3 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumHour().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), hours_1_0)))
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_4_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_5_5 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMinute().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), minutes_2_1)))
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_5_5.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_6_7 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumSecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), seconds_3_2)))
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_7.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__local_var_7_9 := gopurs_runtime.Apply(pkg_Data_Time_Component.Get_boundedEnumMillisecond().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_numSub(), v_0), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numMul(), hours_1_0), gopurs_runtime.Float(3600000.0))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numMul(), minutes_2_1), gopurs_runtime.Float(60000.0)))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numMul(), seconds_3_2), gopurs_runtime.Float(1000.0)))))))
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_9.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t10 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": __local_var_4_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_5_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": __local_var_6_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value3": __local_var_7_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"]})})
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
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_6:
__t4 = __t6
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_4:
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust()
})), __t4)
})
	})
	return millisToTime
}

var hour gopurs_runtime.Value
var once_hour sync.Once
func Get_hour() gopurs_runtime.Value {
	once_hour.Do(func() {
		hour = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
})
	})
	return hour
}

var timeToMillis gopurs_runtime.Value
var once_timeToMillis sync.Once
func Get_timeToMillis() gopurs_runtime.Value {
	once_timeToMillis.Do(func() {
		timeToMillis = gopurs_runtime.Func(func(t_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numMul(), gopurs_runtime.Float(3600000.0)), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), t_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numMul(), gopurs_runtime.Float(60000.0)), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), t_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numMul(), gopurs_runtime.Float(1000.0)), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), t_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"])))), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), t_0.PtrVal.(map[string]gopurs_runtime.Value)["value3"]))
})
	})
	return timeToMillis
}

var eqTime gopurs_runtime.Value
var once_eqTime sync.Once
func Get_eqTime() gopurs_runtime.Value {
	once_eqTime.Do(func() {
		eqTime = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), x_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"]))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), x_0.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"]))
})
})})
	})
	return eqTime
}

var ordTime gopurs_runtime.Value
var once_ordTime sync.Once
func Get_ordTime() gopurs_runtime.Value {
	once_ordTime.Do(func() {
		ordTime = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(v_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_5
} else {

}
}
{
v1_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v1_3_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t4 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_4
} else {

}
}
{
v2_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("LT")})
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v2_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("GT")})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordInt().PtrVal.(map[string]gopurs_runtime.Value)["compare"], x_0.PtrVal.(map[string]gopurs_runtime.Value)["value3"]), y_1.PtrVal.(map[string]gopurs_runtime.Value)["value3"])
}
end_branch_3:
__t4 = __t3
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
return __t5
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqTime()
})})
	})
	return ordTime
}

var diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		diff = gopurs_runtime.Func(func(dictDuration_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictDuration_0.PtrVal.(map[string]gopurs_runtime.Value)["toDuration"], gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), gopurs_runtime.Apply(Get_timeToMillis(), t1_1)), gopurs_runtime.Apply(Get_negateDuration(), gopurs_runtime.Apply(Get_timeToMillis(), t2_2))))
})
})
})
	})
	return diff
}

var boundedTime gopurs_runtime.Value
var once_boundedTime sync.Once
func Get_boundedTime() gopurs_runtime.Value {
	once_boundedTime.Do(func() {
		boundedTime = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bottom": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": gopurs_runtime.Int(0), "value1": gopurs_runtime.Int(0), "value2": gopurs_runtime.Int(0), "value3": gopurs_runtime.Int(0)}), "top": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": gopurs_runtime.Int(23), "value1": gopurs_runtime.Int(59), "value2": gopurs_runtime.Int(59), "value3": gopurs_runtime.Int(999)}), "Ord0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordTime()
})})
	})
	return boundedTime
}

var maxTime gopurs_runtime.Value
var once_maxTime sync.Once
func Get_maxTime() gopurs_runtime.Value {
	once_maxTime.Do(func() {
		maxTime = gopurs_runtime.Apply(Get_timeToMillis(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": gopurs_runtime.Int(23), "value1": gopurs_runtime.Int(59), "value2": gopurs_runtime.Int(59), "value3": gopurs_runtime.Int(999)}))
	})
	return maxTime
}

var minTime gopurs_runtime.Value
var once_minTime sync.Once
func Get_minTime() gopurs_runtime.Value {
	once_minTime.Do(func() {
		minTime = gopurs_runtime.Apply(Get_timeToMillis(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Time"), "value0": gopurs_runtime.Int(0), "value1": gopurs_runtime.Int(0), "value2": gopurs_runtime.Int(0), "value3": gopurs_runtime.Int(0)}))
	})
	return minTime
}

var adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		adjust = gopurs_runtime.Func(func(dictDuration_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_2 gopurs_runtime.Value) gopurs_runtime.Value {
d_prime_3_0 := gopurs_runtime.Apply(dictDuration_0.PtrVal.(map[string]gopurs_runtime.Value)["fromDuration"], d_1)
wholeDays_4_1 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_EuclideanRing.Get_numDiv(), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), d_prime_3_0)), gopurs_runtime.Float(86400000.0)))
msAdjusted_5_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), gopurs_runtime.Apply(Get_timeToMillis(), t_2)), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), d_prime_3_0), gopurs_runtime.Apply(Get_negateDuration(), gopurs_runtime.Apply(pkg_Data_Time_Duration.Get_durationDays().PtrVal.(map[string]gopurs_runtime.Value)["fromDuration"], wholeDays_4_1))))
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordNumber().PtrVal.(map[string]gopurs_runtime.Value)["compare"], msAdjusted_5_2), Get_maxTime()).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
__t4 = gopurs_runtime.Float(1.0)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordNumber().PtrVal.(map[string]gopurs_runtime.Value)["compare"], msAdjusted_5_2), Get_minTime()).PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_numSub(), gopurs_runtime.Float(0.0)), gopurs_runtime.Float(1.0))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Float(0.0)
}
end_branch_4:
wrap_6_3 := __t4
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), wholeDays_4_1), wrap_6_3), "value1": gopurs_runtime.Apply(Get_millisToTime(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numAdd(), msAdjusted_5_2), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semiring.Get_numMul(), gopurs_runtime.Float(86400000.0)), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Ring.Get_numSub(), gopurs_runtime.Float(0.0)), wrap_6_3))))})
})
})
})
	})
	return adjust
}



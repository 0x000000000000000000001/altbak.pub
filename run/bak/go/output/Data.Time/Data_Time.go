package Data_Time

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Time_Component "gopurs/output/Data.Time.Component"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_Time_Duration "gopurs/output/Data.Time.Duration"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	unsafe "unsafe"
)

var cache_negate gopurs_runtime.Value
var once_negate sync.Once
func Get_negate() gopurs_runtime.Value {
	once_negate.Do(func() {
		cache_negate = func() gopurs_runtime.Value {
zero_0_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_0_0
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "sub"), zero_0_0, a_1)
})
}()
	})
	return cache_negate
}

var cache_Time gopurs_runtime.Value
var once_Time sync.Once
func Get_Time() gopurs_runtime.Value {
	once_Time.Do(func() {
		cache_Time = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{value0, value1, value2, value3})}
})
})
})
})
	})
	return cache_Time
}

var cache_showTime gopurs_runtime.Value
var once_showTime sync.Once
func Get_showTime() gopurs_runtime.Value {
	once_showTime.Do(func() {
		cache_showTime = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Time "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_showHour(), "show"), (*Constructor_Time)(v_0.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_showMinute(), "show"), (*Constructor_Time)(v_0.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_showSecond(), "show"), (*Constructor_Time)(v_0.UnsafePtr).V2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_showMillisecond(), "show"), (*Constructor_Time)(v_0.UnsafePtr).V3), gopurs_runtime.Str(")")))))))))
}))
	})
	return cache_showTime
}

var cache_setSecond gopurs_runtime.Value
var once_setSecond sync.Once
func Get_setSecond() gopurs_runtime.Value {
	once_setSecond.Do(func() {
		cache_setSecond = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_setSecond(s_0_box, v_1_box)
})
	})
	return cache_setSecond
}

var cache_setMinute gopurs_runtime.Value
var once_setMinute sync.Once
func Get_setMinute() gopurs_runtime.Value {
	once_setMinute.Do(func() {
		cache_setMinute = gopurs_runtime.Func2(func(m_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_setMinute(m_0_box, v_1_box)
})
	})
	return cache_setMinute
}

var cache_setMillisecond gopurs_runtime.Value
var once_setMillisecond sync.Once
func Get_setMillisecond() gopurs_runtime.Value {
	once_setMillisecond.Do(func() {
		cache_setMillisecond = gopurs_runtime.Func2(func(ms_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_setMillisecond(ms_0_box, v_1_box)
})
	})
	return cache_setMillisecond
}

var cache_setHour gopurs_runtime.Value
var once_setHour sync.Once
func Get_setHour() gopurs_runtime.Value {
	once_setHour.Do(func() {
		cache_setHour = gopurs_runtime.Func2(func(h_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_setHour(h_0_box, v_1_box)
})
	})
	return cache_setHour
}

var cache_second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		cache_second = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_second(v_0_box)
})
	})
	return cache_second
}

var cache_minute gopurs_runtime.Value
var once_minute sync.Once
func Get_minute() gopurs_runtime.Value {
	once_minute.Do(func() {
		cache_minute = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minute(v_0_box)
})
	})
	return cache_minute
}

var cache_millisecond gopurs_runtime.Value
var once_millisecond sync.Once
func Get_millisecond() gopurs_runtime.Value {
	once_millisecond.Do(func() {
		cache_millisecond = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_millisecond(v_0_box)
})
	})
	return cache_millisecond
}

var cache_millisToTime gopurs_runtime.Value
var once_millisToTime sync.Once
func Get_millisToTime() gopurs_runtime.Value {
	once_millisToTime.Do(func() {
		cache_millisToTime = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_millisToTime(v_0_box)
})
	})
	return cache_millisToTime
}

var cache_hour gopurs_runtime.Value
var once_hour sync.Once
func Get_hour() gopurs_runtime.Value {
	once_hour.Do(func() {
		cache_hour = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_hour(v_0_box)
})
	})
	return cache_hour
}

var cache_timeToMillis gopurs_runtime.Value
var once_timeToMillis sync.Once
func Get_timeToMillis() gopurs_runtime.Value {
	once_timeToMillis.Do(func() {
		cache_timeToMillis = gopurs_runtime.Func(func(t_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_timeToMillis(t_0_box)
})
	})
	return cache_timeToMillis
}

var cache_eqTime gopurs_runtime.Value
var once_eqTime sync.Once
func Get_eqTime() gopurs_runtime.Value {
	once_eqTime.Do(func() {
		cache_eqTime = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_eqHour(), "eq"), (*Constructor_Time)(x_0.UnsafePtr).V0, (*Constructor_Time)(y_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_eqMinute(), "eq"), (*Constructor_Time)(x_0.UnsafePtr).V1, (*Constructor_Time)(y_1.UnsafePtr).V1)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_eqSecond(), "eq"), (*Constructor_Time)(x_0.UnsafePtr).V2, (*Constructor_Time)(y_1.UnsafePtr).V2)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_eqMillisecond(), "eq"), (*Constructor_Time)(x_0.UnsafePtr).V3, (*Constructor_Time)(y_1.UnsafePtr).V3))
}))
	})
	return cache_eqTime
}

var cache_ordTime gopurs_runtime.Value
var once_ordTime sync.Once
func Get_ordTime() gopurs_runtime.Value {
	once_ordTime.Do(func() {
		cache_ordTime = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqTime()
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_ordHour(), "compare"), (*Constructor_Time)(x_0.UnsafePtr).V0, (*Constructor_Time)(y_1.UnsafePtr).V0)
_ = v_2_0
var __t5 gopurs_runtime.Value
{
if (v_2_0.Type == 9 && v_2_0.IntVal == 1527465420) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
if (v_2_0.Type == 9 && v_2_0.IntVal == 380165415) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
v1_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_ordMinute(), "compare"), (*Constructor_Time)(x_0.UnsafePtr).V1, (*Constructor_Time)(y_1.UnsafePtr).V1)
_ = v1_3_1
var __t4 gopurs_runtime.Value
{
if (v1_3_1.Type == 9 && v1_3_1.IntVal == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if (v1_3_1.Type == 9 && v1_3_1.IntVal == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
v2_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_ordSecond(), "compare"), (*Constructor_Time)(x_0.UnsafePtr).V2, (*Constructor_Time)(y_1.UnsafePtr).V2)
_ = v2_4_2
var __t3 gopurs_runtime.Value
{
if (v2_4_2.Type == 9 && v2_4_2.IntVal == 1527465420) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
if (v2_4_2.Type == 9 && v2_4_2.IntVal == 380165415) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_ordMillisecond(), "compare"), (*Constructor_Time)(x_0.UnsafePtr).V3, (*Constructor_Time)(y_1.UnsafePtr).V3)
}
end_branch_3:
__t4 = __t3
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
return __t5
}))
	})
	return cache_ordTime
}

var cache_diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		cache_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, t1_1_box gopurs_runtime.Value, t2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_diff(dictDuration_0_box, t1_1_box, t2_2_box)
})
	})
	return cache_diff
}

var cache_boundedTime gopurs_runtime.Value
var once_boundedTime sync.Once
func Get_boundedTime() gopurs_runtime.Value {
	once_boundedTime.Do(func() {
		cache_boundedTime = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordTime()
}), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedHour(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMinute(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedSecond(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMillisecond(), "bottom")})}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedHour(), "top"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMinute(), "top"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedSecond(), "top"), gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMillisecond(), "top")})})
	})
	return cache_boundedTime
}

var cache_maxTime gopurs_runtime.Value
var once_maxTime sync.Once
func Get_maxTime() gopurs_runtime.Value {
	once_maxTime.Do(func() {
		cache_maxTime = gopurs_runtime.Apply(Get_timeToMillis(), gopurs_runtime.RecordGet(Get_boundedTime(), "top"))
	})
	return cache_maxTime
}

var cache_minTime gopurs_runtime.Value
var once_minTime sync.Once
func Get_minTime() gopurs_runtime.Value {
	once_minTime.Do(func() {
		cache_minTime = gopurs_runtime.Apply(Get_timeToMillis(), gopurs_runtime.RecordGet(Get_boundedTime(), "bottom"))
	})
	return cache_minTime
}

var cache_adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		cache_adjust = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, d_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_adjust(dictDuration_0_box, d_1_box, t_2_box)
})
	})
	return cache_adjust
}

type Constructor_Time struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func Call_setSecond(s_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{(*Constructor_Time)(v_1.UnsafePtr).V0, (*Constructor_Time)(v_1.UnsafePtr).V1, s_0, (*Constructor_Time)(v_1.UnsafePtr).V3})}
}

func Call_setMinute(m_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{(*Constructor_Time)(v_1.UnsafePtr).V0, m_0, (*Constructor_Time)(v_1.UnsafePtr).V2, (*Constructor_Time)(v_1.UnsafePtr).V3})}
}

func Call_setMillisecond(ms_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ms_0 gopurs_runtime.Value = ms_0_loop
_ = ms_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{(*Constructor_Time)(v_1.UnsafePtr).V0, (*Constructor_Time)(v_1.UnsafePtr).V1, (*Constructor_Time)(v_1.UnsafePtr).V2, ms_0})}
}

func Call_setHour(h_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var h_0 gopurs_runtime.Value = h_0_loop
_ = h_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{h_0, (*Constructor_Time)(v_1.UnsafePtr).V1, (*Constructor_Time)(v_1.UnsafePtr).V2, (*Constructor_Time)(v_1.UnsafePtr).V3})}
}

func Call_second(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Constructor_Time)(v_0.UnsafePtr).V2
}

func Call_minute(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Constructor_Time)(v_0.UnsafePtr).V1
}

func Call_millisecond(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Constructor_Time)(v_0.UnsafePtr).V3
}

func Call_millisToTime(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
hours_1_0 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), v_0, gopurs_runtime.Float(3600000.0)))
_ = hours_1_0
minutes_2_1 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "sub"), v_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), hours_1_0, gopurs_runtime.Float(3600000.0))), gopurs_runtime.Float(60000.0)))
_ = minutes_2_1
seconds_3_2 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "sub"), v_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), hours_1_0, gopurs_runtime.Float(3600000.0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), minutes_2_1, gopurs_runtime.Float(60000.0)))), gopurs_runtime.Float(1000.0)))
_ = seconds_3_2
__local_var_4_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_Time(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumHour(), "toEnum"), gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), hours_1_0)))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumMinute(), "toEnum"), gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), minutes_2_1)))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumSecond(), "toEnum"), gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), seconds_3_2)))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumMillisecond(), "toEnum"), gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "sub"), v_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), hours_1_0, gopurs_runtime.Float(3600000.0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), minutes_2_1, gopurs_runtime.Float(60000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), seconds_3_2, gopurs_runtime.Float(1000.0))))))))
_ = __local_var_4_3
var __t4 gopurs_runtime.Value
{
if (__local_var_4_3.Type == 9 && __local_var_4_3.IntVal == 930809136 && __local_var_4_3.UnsafePtr != nil) {
__t4 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_4_3.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}

func Call_hour(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Constructor_Time)(v_0.UnsafePtr).V0
}

func Call_timeToMillis(t_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var t_0 gopurs_runtime.Value = t_0_loop
_ = t_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), gopurs_runtime.Float(3600000.0), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumHour(), "fromEnum"), (*Constructor_Time)(t_0.UnsafePtr).V0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), gopurs_runtime.Float(60000.0), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumMinute(), "fromEnum"), (*Constructor_Time)(t_0.UnsafePtr).V1)))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), gopurs_runtime.Float(1000.0), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumSecond(), "fromEnum"), (*Constructor_Time)(t_0.UnsafePtr).V2)))), gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedEnumMillisecond(), "fromEnum"), (*Constructor_Time)(t_0.UnsafePtr).V3)))
}

func Call_diff(dictDuration_0_loop gopurs_runtime.Value, t1_1_loop gopurs_runtime.Value, t2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 gopurs_runtime.Value = dictDuration_0_loop
_ = dictDuration_0
var t1_1 gopurs_runtime.Value = t1_1_loop
_ = t1_1
var t2_2 gopurs_runtime.Value = t2_2_loop
_ = t2_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "toDuration"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_semigroupMilliseconds(), "append"), gopurs_runtime.Apply(Get_timeToMillis(), t1_1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "toDuration"), gopurs_runtime.Apply(pkg_Data_Time_Duration.Get_negate(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "fromDuration"), gopurs_runtime.Apply(Get_timeToMillis(), t2_2))))))
}

func Call_adjust(dictDuration_0_loop gopurs_runtime.Value, d_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 gopurs_runtime.Value = dictDuration_0_loop
_ = dictDuration_0
var d_1 gopurs_runtime.Value = d_1_loop
_ = d_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
d_prime_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDuration_0, "fromDuration"), d_1)
_ = d_prime_3_0
wholeDays_4_1 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), d_prime_3_0, gopurs_runtime.Float(86400000.0)))
_ = wholeDays_4_1
msAdjusted_5_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_semigroupMilliseconds(), "append"), gopurs_runtime.Apply(Get_timeToMillis(), t_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_semigroupMilliseconds(), "append"), d_prime_3_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "toDuration"), gopurs_runtime.Apply(pkg_Data_Time_Duration.Get_negate(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "fromDuration"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationDays(), "fromDuration"), wholeDays_4_1))))))
_ = msAdjusted_5_2
var __t4 gopurs_runtime.Value
{
if (msAdjusted_5_2.FloatVal()) > (Get_maxTime().FloatVal()) {
__t4 = gopurs_runtime.Float(1.0)
goto end_branch_4
} else {

}
}
{
if (msAdjusted_5_2.FloatVal()) < (Get_minTime().FloatVal()) {
__t4 = gopurs_runtime.Apply(Get_negate(), gopurs_runtime.Float(1.0))
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_semigroupDays(), "append"), wholeDays_4_1, wrap_6_3), gopurs_runtime.Apply(Get_millisToTime(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_semigroupMilliseconds(), "append"), msAdjusted_5_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "mul"), gopurs_runtime.Float(86400000.0), gopurs_runtime.Apply(Get_negate(), wrap_6_3))))})}
}



package Data_Time

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Time_Component "gopurs/output/Data.Time.Component"
	pkg_Data_Time_Duration "gopurs/output/Data.Time.Duration"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Time gopurs_runtime.Value
var once_Time sync.Once
func Get_Time() gopurs_runtime.Value {
	once_Time.Do(func() {
		cache_Time = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{1, value0.IntVal, value1.IntVal, value2.IntVal, value3.IntVal})}
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
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Time "), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_show__1488465650(), gopurs_runtime.Int((*Constructor_Time)(v_0.UnsafePtr).V0)).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(" "), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_show__1488465650(), gopurs_runtime.Int((*Constructor_Time)(v_0.UnsafePtr).V1)).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(" "), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_show__1488465650(), gopurs_runtime.Int((*Constructor_Time)(v_0.UnsafePtr).V2)).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(" "), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_show__1488465650(), gopurs_runtime.Int((*Constructor_Time)(v_0.UnsafePtr).V3)).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())).StrVal())).StrVal())).StrVal())).StrVal())).StrVal())).StrVal())
}))
	})
	return cache_showTime
}

var cache_setSecond gopurs_runtime.Value
var once_setSecond sync.Once
func Get_setSecond() gopurs_runtime.Value {
	once_setSecond.Do(func() {
		cache_setSecond = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_setSecond(s_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Time](v_1_box)))}
})
	})
	return cache_setSecond
}

var cache_setMinute gopurs_runtime.Value
var once_setMinute sync.Once
func Get_setMinute() gopurs_runtime.Value {
	once_setMinute.Do(func() {
		cache_setMinute = gopurs_runtime.Func2(func(m_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_setMinute(m_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Time](v_1_box)))}
})
	})
	return cache_setMinute
}

var cache_setMillisecond gopurs_runtime.Value
var once_setMillisecond sync.Once
func Get_setMillisecond() gopurs_runtime.Value {
	once_setMillisecond.Do(func() {
		cache_setMillisecond = gopurs_runtime.Func2(func(ms_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_setMillisecond(ms_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Time](v_1_box)))}
})
	})
	return cache_setMillisecond
}

var cache_setHour gopurs_runtime.Value
var once_setHour sync.Once
func Get_setHour() gopurs_runtime.Value {
	once_setHour.Do(func() {
		cache_setHour = gopurs_runtime.Func2(func(h_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_setHour(h_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Time](v_1_box)))}
})
	})
	return cache_setHour
}

var cache_second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		cache_second = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_second(gopurs_runtime.CoerceToStruct[Constructor_Time](v_0_box)))
})
	})
	return cache_second
}

var cache_minute gopurs_runtime.Value
var once_minute sync.Once
func Get_minute() gopurs_runtime.Value {
	once_minute.Do(func() {
		cache_minute = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_minute(gopurs_runtime.CoerceToStruct[Constructor_Time](v_0_box)))
})
	})
	return cache_minute
}

var cache_millisecond gopurs_runtime.Value
var once_millisecond sync.Once
func Get_millisecond() gopurs_runtime.Value {
	once_millisecond.Do(func() {
		cache_millisecond = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_millisecond(gopurs_runtime.CoerceToStruct[Constructor_Time](v_0_box)))
})
	})
	return cache_millisecond
}

var cache_millisToTime gopurs_runtime.Value
var once_millisToTime sync.Once
func Get_millisToTime() gopurs_runtime.Value {
	once_millisToTime.Do(func() {
		cache_millisToTime = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_millisToTime(v_0_box.FloatVal()))}
})
	})
	return cache_millisToTime
}

var cache_hour gopurs_runtime.Value
var once_hour sync.Once
func Get_hour() gopurs_runtime.Value {
	once_hour.Do(func() {
		cache_hour = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_hour(gopurs_runtime.CoerceToStruct[Constructor_Time](v_0_box)))
})
	})
	return cache_hour
}

var cache_timeToMillis gopurs_runtime.Value
var once_timeToMillis sync.Once
func Get_timeToMillis() gopurs_runtime.Value {
	once_timeToMillis.Do(func() {
		cache_timeToMillis = gopurs_runtime.Func(func(t_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_timeToMillis(gopurs_runtime.CoerceToStruct[Constructor_Time](t_0_box)))
})
	})
	return cache_timeToMillis
}

var cache_eqTime gopurs_runtime.Value
var once_eqTime sync.Once
func Get_eqTime() gopurs_runtime.Value {
	once_eqTime.Do(func() {
		cache_eqTime = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int((*Constructor_Time)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Time)(y_1.UnsafePtr).V0)).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int((*Constructor_Time)(x_0.UnsafePtr).V1), gopurs_runtime.Int((*Constructor_Time)(y_1.UnsafePtr).V1)).IntVal) != (0))).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int((*Constructor_Time)(x_0.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Time)(y_1.UnsafePtr).V2)).IntVal) != (0))).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int((*Constructor_Time)(x_0.UnsafePtr).V3), gopurs_runtime.Int((*Constructor_Time)(y_1.UnsafePtr).V3)).IntVal) != (0))).IntVal) != (0))
})
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
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
v_2_0 := gopurs_runtime.Apply2(Get_compare__372254389(), gopurs_runtime.Int((*Constructor_Time)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Time)(y_1.UnsafePtr).V0))
_ = v_2_0
var __t5 uint32
{
if (uint32(v_2_0.IntVal) == 1527465420) {
__t5 = 1527465420
goto end_branch_5
} else {

}
}
{
if (uint32(v_2_0.IntVal) == 380165415) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
v1_3_1 := gopurs_runtime.Apply2(Get_compare__372254389(), gopurs_runtime.Int((*Constructor_Time)(x_0.UnsafePtr).V1), gopurs_runtime.Int((*Constructor_Time)(y_1.UnsafePtr).V1))
_ = v1_3_1
var __t4 uint32
{
if (uint32(v1_3_1.IntVal) == 1527465420) {
__t4 = 1527465420
goto end_branch_4
} else {

}
}
{
if (uint32(v1_3_1.IntVal) == 380165415) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
v2_4_2 := gopurs_runtime.Apply2(Get_compare__372254389(), gopurs_runtime.Int((*Constructor_Time)(x_0.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Time)(y_1.UnsafePtr).V2))
_ = v2_4_2
var __t3 uint32
{
if (uint32(v2_4_2.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(v2_4_2.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
__t3 = uint32(gopurs_runtime.Apply2(Get_compare__372254389(), gopurs_runtime.Int((*Constructor_Time)(x_0.UnsafePtr).V3), gopurs_runtime.Int((*Constructor_Time)(y_1.UnsafePtr).V3)).IntVal)
}
end_branch_3:
__t4 = __t3
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
})
}))
	})
	return cache_ordTime
}

var cache_diff gopurs_runtime.Value
var once_diff sync.Once
func Get_diff() gopurs_runtime.Value {
	once_diff.Do(func() {
		cache_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, t1_1_box gopurs_runtime.Value, t2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_diff(gopurs_runtime.CoerceToStruct[pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]](dictDuration_0_box), gopurs_runtime.CoerceToStruct[Constructor_Time](t1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Time](t2_2_box))
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
}), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{1, gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedHour(), "bottom").IntVal, gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMinute(), "bottom").IntVal, gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedSecond(), "bottom").IntVal, gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMillisecond(), "bottom").IntVal})}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{1, gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedHour(), "top").IntVal, gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMinute(), "top").IntVal, gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedSecond(), "top").IntVal, gopurs_runtime.RecordGet(pkg_Data_Time_Component.Get_boundedMillisecond(), "top").IntVal})})
	})
	return cache_boundedTime
}

var cache_maxTime gopurs_runtime.Value
var once_maxTime sync.Once
func Get_maxTime() gopurs_runtime.Value {
	once_maxTime.Do(func() {
		cache_maxTime = gopurs_runtime.Float(gopurs_runtime.Float(Call_timeToMillis(gopurs_runtime.CoerceToStruct[Constructor_Time](gopurs_runtime.RecordGet(Get_boundedTime(), "top")))).FloatVal())
	})
	return cache_maxTime
}

var cache_minTime gopurs_runtime.Value
var once_minTime sync.Once
func Get_minTime() gopurs_runtime.Value {
	once_minTime.Do(func() {
		cache_minTime = gopurs_runtime.Float(gopurs_runtime.Float(Call_timeToMillis(gopurs_runtime.CoerceToStruct[Constructor_Time](gopurs_runtime.RecordGet(Get_boundedTime(), "bottom")))).FloatVal())
	})
	return cache_minTime
}

var cache_adjust gopurs_runtime.Value
var once_adjust sync.Once
func Get_adjust() gopurs_runtime.Value {
	once_adjust.Do(func() {
		cache_adjust = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, d_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_adjust(gopurs_runtime.CoerceToStruct[pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]](dictDuration_0_box), d_1_box, gopurs_runtime.CoerceToStruct[Constructor_Time](t_2_box)))}
})
	})
	return cache_adjust
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_apply__3882563466 gopurs_runtime.Value
var once_apply__3882563466 sync.Once
func Get_apply__3882563466() gopurs_runtime.Value {
	once_apply__3882563466.Do(func() {
		cache_apply__3882563466 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply")
	})
	return cache_apply__3882563466
}

var cache_apply__3867059818 gopurs_runtime.Value
var once_apply__3867059818 sync.Once
func Get_apply__3867059818() gopurs_runtime.Value {
	once_apply__3867059818.Do(func() {
		cache_apply__3867059818 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply")
	})
	return cache_apply__3867059818
}

var cache_apply__3489442218 gopurs_runtime.Value
var once_apply__3489442218 sync.Once
func Get_apply__3489442218() gopurs_runtime.Value {
	once_apply__3489442218.Do(func() {
		cache_apply__3489442218 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply")
	})
	return cache_apply__3489442218
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_fromEnum__3599151655 gopurs_runtime.Value
var once_fromEnum__3599151655 sync.Once
func Get_fromEnum__3599151655() gopurs_runtime.Value {
	once_fromEnum__3599151655.Do(func() {
		cache_fromEnum__3599151655 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumYear(), "fromEnum")
	})
	return cache_fromEnum__3599151655
}

var cache_fromEnum__1637084359 gopurs_runtime.Value
var once_fromEnum__1637084359 sync.Once
func Get_fromEnum__1637084359() gopurs_runtime.Value {
	once_fromEnum__1637084359.Do(func() {
		cache_fromEnum__1637084359 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEnum__1637084359(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_fromEnum__1637084359
}

var cache_toEnum__2099864294 gopurs_runtime.Value
var once_toEnum__2099864294 sync.Once
func Get_toEnum__2099864294() gopurs_runtime.Value {
	once_toEnum__2099864294.Do(func() {
		cache_toEnum__2099864294 = gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumDay(), "toEnum")
	})
	return cache_toEnum__2099864294
}

var cache_toEnum__3317293286 gopurs_runtime.Value
var once_toEnum__3317293286 sync.Once
func Get_toEnum__3317293286() gopurs_runtime.Value {
	once_toEnum__3317293286.Do(func() {
		cache_toEnum__3317293286 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnum__3317293286(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_toEnum__3317293286
}

var cache_eq__2843686287 gopurs_runtime.Value
var once_eq__2843686287 sync.Once
func Get_eq__2843686287() gopurs_runtime.Value {
	once_eq__2843686287.Do(func() {
		cache_eq__2843686287 = pkg_Data_Eq.Get_eqIntImpl()
	})
	return cache_eq__2843686287
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_div__1002719800 gopurs_runtime.Value
var once_div__1002719800 sync.Once
func Get_div__1002719800() gopurs_runtime.Value {
	once_div__1002719800.Do(func() {
		cache_div__1002719800 = gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div")
	})
	return cache_div__1002719800
}

var cache_div__2579358968 gopurs_runtime.Value
var once_div__2579358968 sync.Once
func Get_div__2579358968() gopurs_runtime.Value {
	once_div__2579358968.Do(func() {
		cache_div__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_div__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_div__2579358968
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__2165919164 gopurs_runtime.Value
var once_map__2165919164 sync.Once
func Get_map__2165919164() gopurs_runtime.Value {
	once_map__2165919164.Do(func() {
		cache_map__2165919164 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__2165919164
}

var cache_map__901270812 gopurs_runtime.Value
var once_map__901270812 sync.Once
func Get_map__901270812() gopurs_runtime.Value {
	once_map__901270812.Do(func() {
		cache_map__901270812 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__901270812
}

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj")
	})
	return cache_conj__3676519832
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj")
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not")
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_applyMaybe__3561700045 gopurs_runtime.Value
var once_applyMaybe__3561700045 sync.Once
func Get_applyMaybe__3561700045() gopurs_runtime.Value {
	once_applyMaybe__3561700045.Do(func() {
		cache_applyMaybe__3561700045 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1))})))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyMaybe__3561700045
}

var cache_fromJust__1791383420 gopurs_runtime.Value
var once_fromJust__1791383420 sync.Once
func Get_fromJust__1791383420() gopurs_runtime.Value {
	once_fromJust__1791383420.Do(func() {
		cache_fromJust__1791383420 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__1791383420(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_fromJust__1791383420
}

var cache_fromJust__755886620 gopurs_runtime.Value
var once_fromJust__755886620 sync.Once
func Get_fromJust__755886620() gopurs_runtime.Value {
	once_fromJust__755886620.Do(func() {
		cache_fromJust__755886620 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__755886620(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fromJust__755886620
}

var cache_fromMaybe__1972796397 gopurs_runtime.Value
var once_fromMaybe__1972796397 sync.Once
func Get_fromMaybe__1972796397() gopurs_runtime.Value {
	once_fromMaybe__1972796397.Do(func() {
		cache_fromMaybe__1972796397 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__1972796397(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__1972796397
}

var cache_fromMaybe__430429096 gopurs_runtime.Value
var once_fromMaybe__430429096 sync.Once
func Get_fromMaybe__430429096() gopurs_runtime.Value {
	once_fromMaybe__430429096.Do(func() {
		cache_fromMaybe__430429096 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__430429096(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__430429096
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_functorMaybe__2097654001 gopurs_runtime.Value
var once_functorMaybe__2097654001 sync.Once
func Get_functorMaybe__2097654001() gopurs_runtime.Value {
	once_functorMaybe__2097654001.Do(func() {
		cache_functorMaybe__2097654001 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2097654001
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
}

var cache_over__1710753389 gopurs_runtime.Value
var once_over__1710753389 sync.Once
func Get_over__1710753389() gopurs_runtime.Value {
	once_over__1710753389.Do(func() {
		cache_over__1710753389 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over__1710753389(v_0_box)
})
	})
	return cache_over__1710753389
}

var cache_over__3209389325 gopurs_runtime.Value
var once_over__3209389325 sync.Once
func Get_over__3209389325() gopurs_runtime.Value {
	once_over__3209389325.Do(func() {
		cache_over__3209389325 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over__3209389325(v_0_box)
})
	})
	return cache_over__3209389325
}

var cache_over__1462660013 gopurs_runtime.Value
var once_over__1462660013 sync.Once
func Get_over__1462660013() gopurs_runtime.Value {
	once_over__1462660013.Do(func() {
		cache_over__1462660013 = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over__1462660013(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_1_box), v_2_box)
})
	})
	return cache_over__1462660013
}

var cache_unwrap__777744115 gopurs_runtime.Value
var once_unwrap__777744115 sync.Once
func Get_unwrap__777744115() gopurs_runtime.Value {
	once_unwrap__777744115.Do(func() {
		cache_unwrap__777744115 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_unwrap__777744115
}

var cache_unwrap__3267718003 gopurs_runtime.Value
var once_unwrap__3267718003 sync.Once
func Get_unwrap__3267718003() gopurs_runtime.Value {
	once_unwrap__3267718003.Do(func() {
		cache_unwrap__3267718003 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__3267718003(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__3267718003
}

var cache_compare__372254389 gopurs_runtime.Value
var once_compare__372254389 sync.Once
func Get_compare__372254389() gopurs_runtime.Value {
	once_compare__372254389.Do(func() {
		cache_compare__372254389 = gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})
	})
	return cache_compare__372254389
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_greaterThan__1061005983 gopurs_runtime.Value
var once_greaterThan__1061005983 sync.Once
func Get_greaterThan__1061005983() gopurs_runtime.Value {
	once_greaterThan__1061005983.Do(func() {
		cache_greaterThan__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__1061005983(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThan__1061005983
}

var cache_greaterThan__1409282474 gopurs_runtime.Value
var once_greaterThan__1409282474 sync.Once
func Get_greaterThan__1409282474() gopurs_runtime.Value {
	once_greaterThan__1409282474.Do(func() {
		cache_greaterThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan__1409282474
}

var cache_greaterThanOrEq__4087042607 gopurs_runtime.Value
var once_greaterThanOrEq__4087042607 sync.Once
func Get_greaterThanOrEq__4087042607() gopurs_runtime.Value {
	once_greaterThanOrEq__4087042607.Do(func() {
		cache_greaterThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThanOrEq__4087042607
}

var cache_greaterThanOrEq__1061005983 gopurs_runtime.Value
var once_greaterThanOrEq__1061005983 sync.Once
func Get_greaterThanOrEq__1061005983() gopurs_runtime.Value {
	once_greaterThanOrEq__1061005983.Do(func() {
		cache_greaterThanOrEq__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1061005983(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThanOrEq__1061005983
}

var cache_greaterThanOrEq__1409282474 gopurs_runtime.Value
var once_greaterThanOrEq__1409282474 sync.Once
func Get_greaterThanOrEq__1409282474() gopurs_runtime.Value {
	once_greaterThanOrEq__1409282474.Do(func() {
		cache_greaterThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq__1409282474
}

var cache_lessThan__1061005983 gopurs_runtime.Value
var once_lessThan__1061005983 sync.Once
func Get_lessThan__1061005983() gopurs_runtime.Value {
	once_lessThan__1061005983.Do(func() {
		cache_lessThan__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1061005983(a1_0_box, a2_1_box))
})
	})
	return cache_lessThan__1061005983
}

var cache_lessThan__1409282474 gopurs_runtime.Value
var once_lessThan__1409282474 sync.Once
func Get_lessThan__1409282474() gopurs_runtime.Value {
	once_lessThan__1409282474.Do(func() {
		cache_lessThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan__1409282474
}

var cache_lessThanOrEq__4087042607 gopurs_runtime.Value
var once_lessThanOrEq__4087042607 sync.Once
func Get_lessThanOrEq__4087042607() gopurs_runtime.Value {
	once_lessThanOrEq__4087042607.Do(func() {
		cache_lessThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__4087042607
}

var cache_lessThanOrEq__1061005983 gopurs_runtime.Value
var once_lessThanOrEq__1061005983 sync.Once
func Get_lessThanOrEq__1061005983() gopurs_runtime.Value {
	once_lessThanOrEq__1061005983.Do(func() {
		cache_lessThanOrEq__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1061005983(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__1061005983
}

var cache_lessThanOrEq__1409282474 gopurs_runtime.Value
var once_lessThanOrEq__1409282474 sync.Once
func Get_lessThanOrEq__1409282474() gopurs_runtime.Value {
	once_lessThanOrEq__1409282474.Do(func() {
		cache_lessThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThanOrEq__1409282474
}

var cache_negate__2635823316 gopurs_runtime.Value
var once_negate__2635823316 sync.Once
func Get_negate__2635823316() gopurs_runtime.Value {
	once_negate__2635823316.Do(func() {
		cache_negate__2635823316 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__2635823316(a_0_box)
})
	})
	return cache_negate__2635823316
}

var cache_negate__2151342916 gopurs_runtime.Value
var once_negate__2151342916 sync.Once
func Get_negate__2151342916() gopurs_runtime.Value {
	once_negate__2151342916.Do(func() {
		cache_negate__2151342916 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__2151342916(a_0_box)
})
	})
	return cache_negate__2151342916
}

var cache_negate__1364373265 gopurs_runtime.Value
var once_negate__1364373265 sync.Once
func Get_negate__1364373265() gopurs_runtime.Value {
	once_negate__1364373265.Do(func() {
		cache_negate__1364373265 = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__1364373265(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dictRing_0_box))
})
	})
	return cache_negate__1364373265
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = pkg_Data_Ring.Get_intSub()
	})
	return cache_sub__1043827704
}

var cache_sub__1135378904 gopurs_runtime.Value
var once_sub__1135378904 sync.Once
func Get_sub__1135378904() gopurs_runtime.Value {
	once_sub__1135378904.Do(func() {
		cache_sub__1135378904 = pkg_Data_Ring.Get_numSub()
	})
	return cache_sub__1135378904
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_append__3678571768 gopurs_runtime.Value
var once_append__3678571768 sync.Once
func Get_append__3678571768() gopurs_runtime.Value {
	once_append__3678571768.Do(func() {
		cache_append__3678571768 = gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_semigroupMilliseconds(), "append")
	})
	return cache_append__3678571768
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = pkg_Data_Semiring.Get_intAdd()
	})
	return cache_add__560788792
}

var cache_add__137136408 gopurs_runtime.Value
var once_add__137136408 sync.Once
func Get_add__137136408() gopurs_runtime.Value {
	once_add__137136408.Do(func() {
		cache_add__137136408 = pkg_Data_Semiring.Get_numAdd()
	})
	return cache_add__137136408
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_mul__137136408 gopurs_runtime.Value
var once_mul__137136408 sync.Once
func Get_mul__137136408() gopurs_runtime.Value {
	once_mul__137136408.Do(func() {
		cache_mul__137136408 = pkg_Data_Semiring.Get_numMul()
	})
	return cache_mul__137136408
}

var cache_mul__1614463960 gopurs_runtime.Value
var once_mul__1614463960 sync.Once
func Get_mul__1614463960() gopurs_runtime.Value {
	once_mul__1614463960.Do(func() {
		cache_mul__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mul__1614463960
}

var cache_zero__1204848985 gopurs_runtime.Value
var once_zero__1204848985 sync.Once
func Get_zero__1204848985() gopurs_runtime.Value {
	once_zero__1204848985.Do(func() {
		cache_zero__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1204848985(dict_0_box)
})
	})
	return cache_zero__1204848985
}

var cache_show__1488465650 gopurs_runtime.Value
var once_show__1488465650 sync.Once
func Get_show__1488465650() gopurs_runtime.Value {
	once_show__1488465650.Do(func() {
		cache_show__1488465650 = gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show")
	})
	return cache_show__1488465650
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_fromDuration__4195558286 gopurs_runtime.Value
var once_fromDuration__4195558286 sync.Once
func Get_fromDuration__4195558286() gopurs_runtime.Value {
	once_fromDuration__4195558286.Do(func() {
		cache_fromDuration__4195558286 = gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationDays(), "fromDuration")
	})
	return cache_fromDuration__4195558286
}

var cache_fromDuration__1721614606 gopurs_runtime.Value
var once_fromDuration__1721614606 sync.Once
func Get_fromDuration__1721614606() gopurs_runtime.Value {
	once_fromDuration__1721614606.Do(func() {
		cache_fromDuration__1721614606 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromDuration__1721614606(gopurs_runtime.CoerceToStruct[pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_fromDuration__1721614606
}

var cache_negateDuration__4195558286 gopurs_runtime.Value
var once_negateDuration__4195558286 sync.Once
func Get_negateDuration__4195558286() gopurs_runtime.Value {
	once_negateDuration__4195558286.Do(func() {
		cache_negateDuration__4195558286 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negateDuration__4195558286(x_0_box)
})
	})
	return cache_negateDuration__4195558286
}

var cache_negateDuration__3870190523 gopurs_runtime.Value
var once_negateDuration__3870190523 sync.Once
func Get_negateDuration__3870190523() gopurs_runtime.Value {
	once_negateDuration__3870190523.Do(func() {
		cache_negateDuration__3870190523 = gopurs_runtime.Func2(func(dictDuration_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negateDuration__3870190523(gopurs_runtime.CoerceToStruct[pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]](dictDuration_0_box), x_1_box)
})
	})
	return cache_negateDuration__3870190523
}

var cache_toDuration__2440169646 gopurs_runtime.Value
var once_toDuration__2440169646 sync.Once
func Get_toDuration__2440169646() gopurs_runtime.Value {
	once_toDuration__2440169646.Do(func() {
		cache_toDuration__2440169646 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toDuration__2440169646(gopurs_runtime.CoerceToStruct[pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_toDuration__2440169646
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_unsafePartial__4044492477 gopurs_runtime.Value
var once_unsafePartial__4044492477 sync.Once
func Get_unsafePartial__4044492477() gopurs_runtime.Value {
	once_unsafePartial__4044492477.Do(func() {
		cache_unsafePartial__4044492477 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__4044492477
}

type Constructor_Time struct {
	Rc uint32
	V0 int64
	V1 int64
	V2 int64
	V3 int64
}


func Call_setSecond(s_0_loop int64, v_1_loop *Constructor_Time) *Constructor_Time {
var s_0 int64 = s_0_loop
_ = s_0
var v_1 *Constructor_Time = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Time](gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{1, (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1, s_0, (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3})})
}

func Call_setMinute(m_0_loop int64, v_1_loop *Constructor_Time) *Constructor_Time {
var m_0 int64 = m_0_loop
_ = m_0
var v_1 *Constructor_Time = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Time](gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{1, (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, m_0, (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3})})
}

func Call_setMillisecond(ms_0_loop int64, v_1_loop *Constructor_Time) *Constructor_Time {
var ms_0 int64 = ms_0_loop
_ = ms_0
var v_1 *Constructor_Time = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Time](gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{1, (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1, (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, ms_0})})
}

func Call_setHour(h_0_loop int64, v_1_loop *Constructor_Time) *Constructor_Time {
var h_0 int64 = h_0_loop
_ = h_0
var v_1 *Constructor_Time = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Time](gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Time{1, h_0, (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1, (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V2, (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V3})})
}

func Call_second(v_0_loop *Constructor_Time) int64 {
var v_0 *Constructor_Time = v_0_loop
_ = v_0
return (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2
}

func Call_minute(v_0_loop *Constructor_Time) int64 {
var v_0 *Constructor_Time = v_0_loop
_ = v_0
return (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_millisecond(v_0_loop *Constructor_Time) int64 {
var v_0 *Constructor_Time = v_0_loop
_ = v_0
return (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3
}

func Call_millisToTime(v_0_loop float64) *Constructor_Time {
var v_0 float64 = v_0_loop
_ = v_0
hours_1_0 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_div__1002719800(), gopurs_runtime.Float(v_0), gopurs_runtime.Float(3600000.0)).FloatVal())).FloatVal()
_ = hours_1_0
minutes_2_1 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_div__1002719800(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_sub__1135378904(), gopurs_runtime.Float(v_0), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_mul__137136408(), gopurs_runtime.Float(hours_1_0), gopurs_runtime.Float(3600000.0)).FloatVal())).FloatVal()), gopurs_runtime.Float(60000.0)).FloatVal())).FloatVal()
_ = minutes_2_1
seconds_3_2 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_div__1002719800(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_sub__1135378904(), gopurs_runtime.Float(v_0), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_add__137136408(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_mul__137136408(), gopurs_runtime.Float(hours_1_0), gopurs_runtime.Float(3600000.0)).FloatVal()), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_mul__137136408(), gopurs_runtime.Float(minutes_2_1), gopurs_runtime.Float(60000.0)).FloatVal())).FloatVal())).FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal()
_ = seconds_3_2
return gopurs_runtime.CoerceToStruct[Constructor_Time](gopurs_runtime.Apply2(Get_unsafePartial__4044492477(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust__755886620()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*Constructor_Time]](gopurs_runtime.Apply2(Get_apply__3882563466(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_apply__3867059818(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_apply__3489442218(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_map__2165919164(), Get_Time(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(Get_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(hours_1_0))).IntVal))))})))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(Get_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(minutes_2_1))).IntVal))))})))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(Get_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(seconds_3_2))).IntVal))))})))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(Get_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_Int.Get_unsafeClamp(), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_sub__1135378904(), gopurs_runtime.Float(v_0), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_add__137136408(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_add__137136408(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_mul__137136408(), gopurs_runtime.Float(hours_1_0), gopurs_runtime.Float(3600000.0)).FloatVal()), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_mul__137136408(), gopurs_runtime.Float(minutes_2_1), gopurs_runtime.Float(60000.0)).FloatVal())).FloatVal()), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_mul__137136408(), gopurs_runtime.Float(seconds_3_2), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal())).FloatVal()))).IntVal))))})))}))
}

func Call_hour(v_0_loop *Constructor_Time) int64 {
var v_0 *Constructor_Time = v_0_loop
_ = v_0
return (*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_timeToMillis(t_0_loop *Constructor_Time) float64 {
var t_0 *Constructor_Time = t_0_loop
_ = t_0
return gopurs_runtime.Apply2(Get_add__137136408(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_add__137136408(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_add__137136408(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_mul__137136408(), gopurs_runtime.Float(3600000.0), gopurs_runtime.Float(gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_fromEnum__3599151655(), gopurs_runtime.Int((*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(t_0)}.UnsafePtr).V0)).IntVal)).FloatVal())).FloatVal()), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_mul__137136408(), gopurs_runtime.Float(60000.0), gopurs_runtime.Float(gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_fromEnum__3599151655(), gopurs_runtime.Int((*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(t_0)}.UnsafePtr).V1)).IntVal)).FloatVal())).FloatVal())).FloatVal()), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_mul__137136408(), gopurs_runtime.Float(1000.0), gopurs_runtime.Float(gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_fromEnum__3599151655(), gopurs_runtime.Int((*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(t_0)}.UnsafePtr).V2)).IntVal)).FloatVal())).FloatVal())).FloatVal()), gopurs_runtime.Float(gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_fromEnum__3599151655(), gopurs_runtime.Int((*Constructor_Time)(gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(t_0)}.UnsafePtr).V3)).IntVal)).FloatVal())).FloatVal()
}

func Call_diff(dictDuration_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value], t1_1_loop *Constructor_Time, t2_2_loop *Constructor_Time) gopurs_runtime.Value {
var dictDuration_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var t1_1 *Constructor_Time = t1_1_loop
_ = t1_1
var t2_2 *Constructor_Time = t2_2_loop
_ = t2_2
return gopurs_runtime.Apply(dictDuration_0.V1, gopurs_runtime.Float(gopurs_runtime.Apply2(Get_append__3678571768(), gopurs_runtime.Float(gopurs_runtime.Float(Call_timeToMillis(t1_1)).FloatVal()), gopurs_runtime.Float(Call_negateDuration__4195558286(gopurs_runtime.Float(gopurs_runtime.Float(Call_timeToMillis(t2_2)).FloatVal())).FloatVal())).FloatVal()))
}

func Call_adjust(dictDuration_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value], d_1_loop gopurs_runtime.Value, t_2_loop *Constructor_Time) *pkg_Data_Tuple.Constructor_Tuple[float64, *Constructor_Time] {
var dictDuration_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var d_1 gopurs_runtime.Value = d_1_loop
_ = d_1
var t_2 *Constructor_Time = t_2_loop
_ = t_2
d_prime_3_0 := gopurs_runtime.Apply(dictDuration_0.V0, d_1)
_ = d_prime_3_0
wholeDays_4_1 := gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_div__1002719800(), gopurs_runtime.Float(gopurs_runtime.Apply(Get_unwrap__777744115(), gopurs_runtime.Float(d_prime_3_0.FloatVal())).FloatVal()), gopurs_runtime.Float(86400000.0)).FloatVal())).FloatVal()
_ = wholeDays_4_1
msAdjusted_5_2 := gopurs_runtime.Apply2(Get_append__3678571768(), gopurs_runtime.Float(gopurs_runtime.Float(Call_timeToMillis(t_2)).FloatVal()), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_append__3678571768(), gopurs_runtime.Float(d_prime_3_0.FloatVal()), gopurs_runtime.Float(Call_negateDuration__4195558286(gopurs_runtime.Float(gopurs_runtime.Apply(Get_fromDuration__4195558286(), gopurs_runtime.Float(wholeDays_4_1)).FloatVal())).FloatVal())).FloatVal()))
_ = msAdjusted_5_2
var __t5 float64
{
if (gopurs_runtime.Bool(Call_greaterThan__1061005983(gopurs_runtime.Float(msAdjusted_5_2.FloatVal()), gopurs_runtime.Float(Get_maxTime().FloatVal()))).IntVal) != (0) {
__t5 = 1.0
goto end_branch_5
} else {

}
}
{
var __t4 float64
{
if (gopurs_runtime.Bool(Call_lessThan__1061005983(gopurs_runtime.Float(msAdjusted_5_2.FloatVal()), gopurs_runtime.Float(Get_minTime().FloatVal()))).IntVal) != (0) {
__t4 = Call_negate__2151342916(gopurs_runtime.Float(1.0)).FloatVal()
goto end_branch_4
} else {

}
}
{
__t4 = 0.0
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
wrap_6_3 := __t5
_ = wrap_6_3
return gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[float64, *Constructor_Time]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Float(gopurs_runtime.Apply2(Get_append__3678571768(), gopurs_runtime.Float(wholeDays_4_1), gopurs_runtime.Float(wrap_6_3)).FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Time](gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_millisToTime(gopurs_runtime.Apply2(Get_append__3678571768(), gopurs_runtime.Float(msAdjusted_5_2.FloatVal()), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_mul__137136408(), gopurs_runtime.Float(86400000.0), gopurs_runtime.Float(Call_negate__2151342916(gopurs_runtime.Float(wrap_6_3)).FloatVal())).FloatVal())).FloatVal()))}))}})})
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fromEnum__1637084359(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_toEnum__3317293286(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_div__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_fromJust__1791383420(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
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

func Call_fromJust__755886620(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
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

func Call_fromMaybe__1972796397(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
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

func Call_fromMaybe__430429096(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
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

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
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

func Call_over__1710753389(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_over__3209389325(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_over__1462660013(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], _dollar__unused_1_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_1_loop
_ = _dollar__unused_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__3267718003(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_greaterThan__1061005983(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.FloatVal()) > (a2_1.FloatVal()) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_greaterThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_greaterThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_greaterThanOrEq__1061005983(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.FloatVal()) < (a2_1.FloatVal()) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_greaterThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_lessThan__1061005983(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.FloatVal()) < (a2_1.FloatVal()) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_lessThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_lessThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_lessThanOrEq__1061005983(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.FloatVal()) > (a2_1.FloatVal()) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_lessThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_negate__2635823316(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Int(-(a_0.IntVal))
}

func Call_negate__2151342916(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Float(-(a_0.FloatVal()))
}

func Call_negate__1364373265(dictRing_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictRing_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dictRing_0_loop
_ = dictRing_0
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(dictRing_0.V0, gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictRing_0.V1, Semiring0_1_0.V3, a_2)
})
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mul__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fromDuration__1721614606(dict_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_negateDuration__4195558286(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "toDuration"), gopurs_runtime.Apply(pkg_Data_Time_Duration.Get_negate(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Time_Duration.Get_durationMilliseconds(), "fromDuration"), x_0)))
}

func Call_negateDuration__3870190523(dictDuration_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictDuration_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictDuration_0.V1, gopurs_runtime.Apply(pkg_Data_Time_Duration.Get_negate(), gopurs_runtime.Apply(dictDuration_0.V0, x_1)))
}

func Call_toDuration__2440169646(dict_0_loop *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Time_Duration.Constructor_Duration[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}



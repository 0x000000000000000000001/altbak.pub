package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Time_Time gopurs_runtime.Value
var once_Data_Time_Time sync.Once
func Get_Data_Time_Time() gopurs_runtime.Value {
	once_Data_Time_Time.Do(func() {
		cache_Data_Time_Time = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((&Constructor_Data_Time_Time{1, value0.IntVal, value1.IntVal, value2.IntVal, value3.IntVal}))}
})
})
})
})
	})
	return cache_Data_Time_Time
}

var cache_Data_Time_showTime gopurs_runtime.Value
var once_Data_Time_showTime sync.Once
func Get_Data_Time_showTime() gopurs_runtime.Value {
	once_Data_Time_showTime.Do(func() {
		cache_Data_Time_showTime = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_4160946377_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Time_Time]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((((((("(Time (Hour ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Time_Time)(v_0.UnsafePtr).V0)).StrVal())) + (") (Minute ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Time_Time)(v_0.UnsafePtr).V1)).StrVal())) + (") (Second ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Time_Time)(v_0.UnsafePtr).V2)).StrVal())) + (") (Millisecond ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Time_Time)(v_0.UnsafePtr).V3)).StrVal())) + ("))"))
})})))}
	})
	return cache_Data_Time_showTime
}

var cache_Data_Time_setSecond gopurs_runtime.Value
var once_Data_Time_setSecond sync.Once
func Get_Data_Time_setSecond() gopurs_runtime.Value {
	once_Data_Time_setSecond.Do(func() {
		cache_Data_Time_setSecond = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_Time_setSecond(s_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](v_1_box)))}
})
	})
	return cache_Data_Time_setSecond
}

var cache_Data_Time_setMinute gopurs_runtime.Value
var once_Data_Time_setMinute sync.Once
func Get_Data_Time_setMinute() gopurs_runtime.Value {
	once_Data_Time_setMinute.Do(func() {
		cache_Data_Time_setMinute = gopurs_runtime.Func2(func(m_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_Time_setMinute(m_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](v_1_box)))}
})
	})
	return cache_Data_Time_setMinute
}

var cache_Data_Time_setMillisecond gopurs_runtime.Value
var once_Data_Time_setMillisecond sync.Once
func Get_Data_Time_setMillisecond() gopurs_runtime.Value {
	once_Data_Time_setMillisecond.Do(func() {
		cache_Data_Time_setMillisecond = gopurs_runtime.Func2(func(ms_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_Time_setMillisecond(ms_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](v_1_box)))}
})
	})
	return cache_Data_Time_setMillisecond
}

var cache_Data_Time_setHour gopurs_runtime.Value
var once_Data_Time_setHour sync.Once
func Get_Data_Time_setHour() gopurs_runtime.Value {
	once_Data_Time_setHour.Do(func() {
		cache_Data_Time_setHour = gopurs_runtime.Func2(func(h_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_Time_setHour(h_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](v_1_box)))}
})
	})
	return cache_Data_Time_setHour
}

var cache_Data_Time_second gopurs_runtime.Value
var once_Data_Time_second sync.Once
func Get_Data_Time_second() gopurs_runtime.Value {
	once_Data_Time_second.Do(func() {
		cache_Data_Time_second = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Time_second(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](v_0_box)))
})
	})
	return cache_Data_Time_second
}

var cache_Data_Time_minute gopurs_runtime.Value
var once_Data_Time_minute sync.Once
func Get_Data_Time_minute() gopurs_runtime.Value {
	once_Data_Time_minute.Do(func() {
		cache_Data_Time_minute = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Time_minute(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](v_0_box)))
})
	})
	return cache_Data_Time_minute
}

var cache_Data_Time_millisecond gopurs_runtime.Value
var once_Data_Time_millisecond sync.Once
func Get_Data_Time_millisecond() gopurs_runtime.Value {
	once_Data_Time_millisecond.Do(func() {
		cache_Data_Time_millisecond = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Time_millisecond(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](v_0_box)))
})
	})
	return cache_Data_Time_millisecond
}

var cache_Data_Time_millisToTime gopurs_runtime.Value
var once_Data_Time_millisToTime sync.Once
func Get_Data_Time_millisToTime() gopurs_runtime.Value {
	once_Data_Time_millisToTime.Do(func() {
		cache_Data_Time_millisToTime = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_Time_millisToTime(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Time_millisToTime
}

var cache_Data_Time_hour gopurs_runtime.Value
var once_Data_Time_hour sync.Once
func Get_Data_Time_hour() gopurs_runtime.Value {
	once_Data_Time_hour.Do(func() {
		cache_Data_Time_hour = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Time_hour(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](v_0_box)))
})
	})
	return cache_Data_Time_hour
}

var cache_Data_Time_timeToMillis gopurs_runtime.Value
var once_Data_Time_timeToMillis sync.Once
func Get_Data_Time_timeToMillis() gopurs_runtime.Value {
	once_Data_Time_timeToMillis.Do(func() {
		cache_Data_Time_timeToMillis = gopurs_runtime.Func(func(t_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Time_timeToMillis(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](t_0_box)))
})
	})
	return cache_Data_Time_timeToMillis
}

var cache_Data_Time_eqTime gopurs_runtime.Value
var once_Data_Time_eqTime sync.Once
func Get_Data_Time_eqTime() gopurs_runtime.Value {
	once_Data_Time_eqTime.Do(func() {
		cache_Data_Time_eqTime = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_2330623017_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Time_Time]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((((((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V0) == ((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V0)) && (((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V1) == ((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V1))) && (((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V2) == ((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V2))) && (((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V3) == ((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V3)))
})})))}
	})
	return cache_Data_Time_eqTime
}

var cache_Data_Time_ordTime gopurs_runtime.Value
var once_Data_Time_ordTime sync.Once
func Get_Data_Time_ordTime() gopurs_runtime.Value {
	once_Data_Time_ordTime.Do(func() {
		cache_Data_Time_ordTime = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_3383090825_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Time_Time]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_2330623017_3790796878(Rebox_Data_Time_3790796878_2330623017(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Time_eqTime()))))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_2_0 := uint32(gopurs_runtime.Apply2(Rebox_Data_Time_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V0)).IntVal)
_ = v_2_0
var __t5 uint32
{
if (v_2_0 == 1527465420) {
__t5 = 1527465420
goto end_branch_5
} else {

}
}
{
if (v_2_0 == 380165415) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
// TAST (Let): v1_3_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_3_1 := uint32(gopurs_runtime.Apply2(Rebox_Data_Time_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V1), gopurs_runtime.Int((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V1)).IntVal)
_ = v1_3_1
var __t4 uint32
{
if (v1_3_1 == 1527465420) {
__t4 = 1527465420
goto end_branch_4
} else {

}
}
{
if (v1_3_1 == 380165415) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
// TAST (Let): v2_4_2 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v2_4_2 := uint32(gopurs_runtime.Apply2(Rebox_Data_Time_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V2)).IntVal)
_ = v2_4_2
var __t3 uint32
{
if (v2_4_2 == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (v2_4_2 == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
__t3 = uint32(gopurs_runtime.Apply2(Rebox_Data_Time_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V3), gopurs_runtime.Int((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V3)).IntVal)
}
end_branch_3:
__t4 = __t3
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Time_ordTime
}

var cache_Data_Time_diff gopurs_runtime.Value
var once_Data_Time_diff sync.Once
func Get_Data_Time_diff() gopurs_runtime.Value {
	once_Data_Time_diff.Do(func() {
		cache_Data_Time_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, t1_1_box gopurs_runtime.Value, t2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_diff(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dictDuration_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](t1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](t2_2_box))
})
	})
	return cache_Data_Time_diff
}

var cache_Data_Time_boundedTime gopurs_runtime.Value
var once_Data_Time_boundedTime sync.Once
func Get_Data_Time_boundedTime() gopurs_runtime.Value {
	once_Data_Time_boundedTime.Do(func() {
		cache_Data_Time_boundedTime = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_4136451977_2094947566((&Constructor_Data_Bounded_Bounded[*Constructor_Data_Time_Time]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_3383090825_4177771502(Rebox_Data_Time_4177771502_3383090825(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Time_ordTime()))))}
}), (&Constructor_Data_Time_Time{1, Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_3764732725_2094947566(Rebox_Data_Time_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_Component_boundedHour()))))}).IntVal, Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_3764732725_2094947566(Rebox_Data_Time_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_Component_boundedMinute()))))}).IntVal, Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_3764732725_2094947566(Rebox_Data_Time_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_Component_boundedSecond()))))}).IntVal, Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_3764732725_2094947566(Rebox_Data_Time_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_Component_boundedMillisecond()))))}).IntVal}), (&Constructor_Data_Time_Time{1, Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_3764732725_2094947566(Rebox_Data_Time_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_Component_boundedHour()))))}).IntVal, Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_3764732725_2094947566(Rebox_Data_Time_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_Component_boundedMinute()))))}).IntVal, Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_3764732725_2094947566(Rebox_Data_Time_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_Component_boundedSecond()))))}).IntVal, Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_3764732725_2094947566(Rebox_Data_Time_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_Component_boundedMillisecond()))))}).IntVal})})))}
	})
	return cache_Data_Time_boundedTime
}

var cache_Data_Time_maxTime gopurs_runtime.Value
var once_Data_Time_maxTime sync.Once
func Get_Data_Time_maxTime() gopurs_runtime.Value {
	once_Data_Time_maxTime.Do(func() {
		cache_Data_Time_maxTime = gopurs_runtime.Float(Call_Data_Time_timeToMillis(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_4136451977_2094947566(Rebox_Data_Time_2094947566_4136451977(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_boundedTime()))))}))))
	})
	return cache_Data_Time_maxTime
}

var cache_Data_Time_minTime gopurs_runtime.Value
var once_Data_Time_minTime sync.Once
func Get_Data_Time_minTime() gopurs_runtime.Value {
	once_Data_Time_minTime.Do(func() {
		cache_Data_Time_minTime = gopurs_runtime.Float(Call_Data_Time_timeToMillis(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_4136451977_2094947566(Rebox_Data_Time_2094947566_4136451977(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_boundedTime()))))}))))
	})
	return cache_Data_Time_minTime
}

var cache_Data_Time_adjust gopurs_runtime.Value
var once_Data_Time_adjust sync.Once
func Get_Data_Time_adjust() gopurs_runtime.Value {
	once_Data_Time_adjust.Do(func() {
		cache_Data_Time_adjust = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, d_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Time_adjust(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dictDuration_0_box), d_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](t_2_box))
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()
})
	})
	return cache_Data_Time_adjust
}

type Constructor_Data_Time_Time struct {
	Rc uint32
	V0 int64
	V1 int64
	V2 int64
	V3 int64
}


func Call_Data_Time_setSecond(s_0_loop int64, v_1_loop *Constructor_Data_Time_Time) *Constructor_Data_Time_Time {
var s_0 int64 = s_0_loop
_ = s_0
var v_1 *Constructor_Data_Time_Time = v_1_loop
_ = v_1
return (&Constructor_Data_Time_Time{1, (v_1).V0, (v_1).V1, s_0, (v_1).V3})
}

func Call_Data_Time_setMinute(m_0_loop int64, v_1_loop *Constructor_Data_Time_Time) *Constructor_Data_Time_Time {
var m_0 int64 = m_0_loop
_ = m_0
var v_1 *Constructor_Data_Time_Time = v_1_loop
_ = v_1
return (&Constructor_Data_Time_Time{1, (v_1).V0, m_0, (v_1).V2, (v_1).V3})
}

func Call_Data_Time_setMillisecond(ms_0_loop int64, v_1_loop *Constructor_Data_Time_Time) *Constructor_Data_Time_Time {
var ms_0 int64 = ms_0_loop
_ = ms_0
var v_1 *Constructor_Data_Time_Time = v_1_loop
_ = v_1
return (&Constructor_Data_Time_Time{1, (v_1).V0, (v_1).V1, (v_1).V2, ms_0})
}

func Call_Data_Time_setHour(h_0_loop int64, v_1_loop *Constructor_Data_Time_Time) *Constructor_Data_Time_Time {
var h_0 int64 = h_0_loop
_ = h_0
var v_1 *Constructor_Data_Time_Time = v_1_loop
_ = v_1
return (&Constructor_Data_Time_Time{1, h_0, (v_1).V1, (v_1).V2, (v_1).V3})
}

func Call_Data_Time_second(v_0_loop *Constructor_Data_Time_Time) int64 {
var v_0 *Constructor_Data_Time_Time = v_0_loop
_ = v_0
return (v_0).V2
}

func Call_Data_Time_minute(v_0_loop *Constructor_Data_Time_Time) int64 {
var v_0 *Constructor_Data_Time_Time = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Time_millisecond(v_0_loop *Constructor_Data_Time_Time) int64 {
var v_0 *Constructor_Data_Time_Time = v_0_loop
_ = v_0
return (v_0).V3
}

func Call_Data_Time_millisToTime(v_0_loop float64) *Constructor_Data_Time_Time {
var v_0 float64 = v_0_loop
_ = v_0
// TAST (Let): hours_1_0 shape=App(Var) bindingType=Number
hours_1_0 := gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float((v_0) / (3600000.0))).FloatVal()
_ = hours_1_0
// TAST (Let): minutes_2_1 shape=App(Var) bindingType=Number
minutes_2_1 := gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float(((v_0) - ((hours_1_0) * (3600000.0))) / (60000.0))).FloatVal()
_ = minutes_2_1
// TAST (Let): seconds_3_2 shape=App(Var) bindingType=Number
seconds_3_2 := gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float(((v_0) - (((hours_1_0) * (3600000.0)) + ((minutes_2_1) * (60000.0)))) / (1000.0))).FloatVal()
_ = seconds_3_2
// TAST (Let): __local_var_4_5 shape=App(Var) bindingType=Any
__local_var_4_5 := gopurs_runtime.Apply(Get_Data_Int_floor(), gopurs_runtime.Float(hours_1_0))
_ = __local_var_4_5
var __t6 *Constructor_Data_Maybe_Just[int64]
{
if ((__local_var_4_5.IntVal) >= (int64(0))) && ((__local_var_4_5.IntVal) <= (int64(23))) {
__t6 = Rebox_Data_Time_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(__local_var_4_5.IntVal), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_6
} else {

}
}
{
__t6 = Rebox_Data_Time_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_6:
// TAST (Let): __local_var_4_4 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_4_4 := __t6
_ = __local_var_4_4
var __t24 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_4_4 != nil) {
// TAST (Let): __local_var_5_12 shape=App(Var) bindingType=Any
__local_var_5_12 := gopurs_runtime.Apply(Get_Data_Int_floor(), gopurs_runtime.Float(minutes_2_1))
_ = __local_var_5_12
var __t13 *Constructor_Data_Maybe_Just[int64]
{
if ((__local_var_5_12.IntVal) >= (int64(0))) && ((__local_var_5_12.IntVal) <= (int64(59))) {
__t13 = Rebox_Data_Time_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(__local_var_5_12.IntVal), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_13
} else {

}
}
{
__t13 = Rebox_Data_Time_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_13:
// TAST (Let): __local_var_5_11 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_5_11 := __t13
_ = __local_var_5_11
var __t14 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_11 != nil) {
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Time_Time(), gopurs_runtime.Int((__local_var_4_4).V0), gopurs_runtime.Int((__local_var_5_11).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_14:
// TAST (Let): __local_var_5_10 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [(Func [Int, Int] (ADT ["Data","Time","Time"] []))])
__local_var_5_10 := __t14
_ = __local_var_5_10
// TAST (Let): __local_var_6_16 shape=App(Var) bindingType=Any
__local_var_6_16 := gopurs_runtime.Apply(Get_Data_Int_floor(), gopurs_runtime.Float(seconds_3_2))
_ = __local_var_6_16
var __t17 *Constructor_Data_Maybe_Just[int64]
{
if ((__local_var_6_16.IntVal) >= (int64(0))) && ((__local_var_6_16.IntVal) <= (int64(59))) {
__t17 = Rebox_Data_Time_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(__local_var_6_16.IntVal), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_17
} else {

}
}
{
__t17 = Rebox_Data_Time_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_17:
// TAST (Let): __local_var_6_15 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_6_15 := __t17
_ = __local_var_6_15
var __t19 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_10 != nil) {
var __t18 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_6_15 != nil) {
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_10).V0, gopurs_runtime.Int((__local_var_6_15).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_18:
__t19 = __t18
goto end_branch_19
} else {

}
}
{
if (__local_var_5_10 == nil) {
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_19
} else {

}
}
{
__t19 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_19:
// TAST (Let): __local_var_5_9 shape=Let(Let(Branch(Branch(Other, def=Other), Other, def=Other))) bindingType=(ADT ["Data","Maybe","Maybe"] [(Func [Int] (ADT ["Data","Time","Time"] []))])
__local_var_5_9 := __t19
_ = __local_var_5_9
// TAST (Let): __local_var_6_20 shape=App(Var) bindingType=Any
__local_var_6_20 := gopurs_runtime.Apply(Get_Data_Int_floor(), gopurs_runtime.Float((v_0) - ((((hours_1_0) * (3600000.0)) + ((minutes_2_1) * (60000.0))) + ((seconds_3_2) * (1000.0)))))
_ = __local_var_6_20
var __t23 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((__local_var_6_20.IntVal) >= (int64(0))) && ((__local_var_6_20.IntVal) <= (int64(999))) {
var __t22 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_9 != nil) {
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_9).V0, gopurs_runtime.Int(__local_var_6_20.IntVal)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_22
} else {

}
}
{
if (__local_var_5_9 == nil) {
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_22
} else {

}
}
{
__t22 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_22:
__t23 = __t22
goto end_branch_23
} else {

}
}
{
var __t21 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_9 != nil) {
__t21 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_21
} else {

}
}
{
if (__local_var_5_9 == nil) {
__t21 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_21
} else {

}
}
{
__t21 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_21:
__t23 = __t21
}
end_branch_23:
__t24 = __t23
goto end_branch_24
} else {

}
}
{
// TAST (Let): __local_var_5_7 shape=App(Var) bindingType=Any
__local_var_5_7 := gopurs_runtime.Apply(Get_Data_Int_floor(), gopurs_runtime.Float((v_0) - ((((hours_1_0) * (3600000.0)) + ((minutes_2_1) * (60000.0))) + ((seconds_3_2) * (1000.0)))))
_ = __local_var_5_7
var __t8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((__local_var_5_7.IntVal) >= (int64(0))) && ((__local_var_5_7.IntVal) <= (int64(999))) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_8:
__t24 = __t8
}
end_branch_24:
// TAST (Let): __local_var_4_3 shape=Let(Branch(Let(Let(Branch(Branch(Other, Other, def=Other), def=Branch(Other, Other, def=Other)))), def=Let(Branch(Other, def=Other)))) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Time","Time"] [])])
var __local_var_4_3 *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] = Rebox_Data_Time_3094389156_3839235747(__t24)
var __t25 *Constructor_Data_Time_Time
{
if (__local_var_4_3 != nil) {
__t25 = (__local_var_4_3).V0
goto end_branch_25
} else {

}
}
{
__t25 = func() *Constructor_Data_Time_Time { panic("Failed pattern match") }()
}
end_branch_25:
return __t25
}

func Call_Data_Time_hour(v_0_loop *Constructor_Data_Time_Time) int64 {
var v_0 *Constructor_Data_Time_Time = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_Time_timeToMillis(t_0_loop *Constructor_Data_Time_Time) float64 {
var t_0 *Constructor_Data_Time_Time = t_0_loop
_ = t_0
return ((((3600000.0) * (gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int((t_0).V0)).FloatVal())) + ((60000.0) * (gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int((t_0).V1)).FloatVal()))) + ((1000.0) * (gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int((t_0).V2)).FloatVal()))) + (gopurs_runtime.Apply(Get_Data_Int_toNumber(), gopurs_runtime.Int((t_0).V3)).FloatVal())
}

func Call_Data_Time_diff(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value], t1_1_loop *Constructor_Data_Time_Time, t2_2_loop *Constructor_Data_Time_Time) gopurs_runtime.Value {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var t1_1 *Constructor_Data_Time_Time = t1_1_loop
_ = t1_1
var t2_2 *Constructor_Data_Time_Time = t2_2_loop
_ = t2_2
return gopurs_runtime.Apply(dictDuration_0.V1, gopurs_runtime.Float((Call_Data_Time_timeToMillis(t1_1)) + (Call_Data_Time_Duration_negateDuration__2130721996(Call_Data_Time_timeToMillis(t2_2)))))
}

func Call_Data_Time_adjust(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value], d_1_loop gopurs_runtime.Value, t_2_loop *Constructor_Data_Time_Time) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var d_1 gopurs_runtime.Value = d_1_loop
_ = d_1
var t_2 *Constructor_Data_Time_Time = t_2_loop
_ = t_2
// TAST (Let): d_prime__3_0 shape=App(Other) bindingType=Number
d_prime__3_0 := gopurs_runtime.Apply(dictDuration_0.V0, d_1).FloatVal()
_ = d_prime__3_0
// TAST (Let): wholeDays_4_1 shape=App(Var) bindingType=Number
wholeDays_4_1 := gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float((gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Float(d_prime__3_0)).FloatVal()) / (86400000.0))).FloatVal()
_ = wholeDays_4_1
// TAST (Let): msAdjusted_5_2 shape=Other bindingType=Number
msAdjusted_5_2 := ((Call_Data_Time_timeToMillis(t_2)) + (d_prime__3_0)) + (Call_Data_Time_Duration_negateDuration__2130721996(gopurs_runtime.Apply2(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_5.FloatVal()) * (86400000.0))
}), gopurs_runtime.Float(wholeDays_4_1)).FloatVal()))
_ = msAdjusted_5_2
var __t5 float64
{
if (msAdjusted_5_2) > (Get_Data_Time_maxTime().FloatVal()) {
__t5 = 1.0
goto end_branch_5
} else {

}
}
{
var __t4 float64
{
if (msAdjusted_5_2) < (Get_Data_Time_minTime().FloatVal()) {
__t4 = -1.0
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
// TAST (Let): wrap_6_3 shape=Branch(LitNumber, def=Branch(LitNumber, def=LitNumber)) bindingType=Any
wrap_6_3 := __t5
_ = wrap_6_3
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Time_4002750455_138441832(Rebox_Data_Time_138441832_4002750455(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Float((wholeDays_4_1) + (wrap_6_3)), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_Time_millisToTime((msAdjusted_5_2) + ((86400000.0) * (-(wrap_6_3)))))}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Rebox_Data_Time_138441832_4002750455(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[float64, *Constructor_Data_Time_Time] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[float64, *Constructor_Data_Time_Time]{}
		out.V0 = in.V0.FloatVal()
		out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](in.V1)
	return out
}

func Rebox_Data_Time_2094947566_3764732725(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1.IntVal
		out.V2 = in.V2.IntVal
	return out
}

func Rebox_Data_Time_2094947566_4136451977(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[*Constructor_Data_Time_Time] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[*Constructor_Data_Time_Time]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](in.V1)
		out.V2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](in.V2)
	return out
}

func Rebox_Data_Time_2330623017_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Time_Time]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Time_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Time_3094389156_3839235747(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](in.V0)
	return out
}

func Rebox_Data_Time_3383090825_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Time_Time]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Time_3764732725_2094947566(in *Constructor_Data_Bounded_Bounded[int64]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Int(in.V1)
		out.V2 = gopurs_runtime.Int(in.V2)
	return out
}

func Rebox_Data_Time_3790796878_2330623017(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Time_Time] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Time_Time]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Time_4002750455_138441832(in *Constructor_Data_Tuple_Tuple[float64, *Constructor_Data_Time_Time]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Float(in.V0)
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Time_4136451977_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_Time_Time]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(in.V1)}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_Time_4160946377_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Time_Time]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Time_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Time_4177771502_3383090825(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Time_Time] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_Time_Time]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}



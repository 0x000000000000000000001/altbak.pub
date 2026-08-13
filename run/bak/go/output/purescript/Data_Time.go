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
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Data_Time_Time{1, value0.IntVal, value1.IntVal, value2.IntVal, value3.IntVal})}
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
		cache_Data_Time_showTime = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((((((("(Time ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Time_Time)(v_0.UnsafePtr).V0)).StrVal())) + (" ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Time_Time)(v_0.UnsafePtr).V1)).StrVal())) + (" ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Time_Time)(v_0.UnsafePtr).V2)).StrVal())) + (" ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Time_Time)(v_0.UnsafePtr).V3)).StrVal())) + (")"))
}))
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
		cache_Data_Time_eqTime = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((((((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V0) == ((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V0)) && (((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V1) == ((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V1))) && (((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V2) == ((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V2))) && (((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V3) == ((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V3)))
})
}))
	})
	return cache_Data_Time_eqTime
}

var cache_Data_Time_ordTime gopurs_runtime.Value
var once_Data_Time_ordTime sync.Once
func Get_Data_Time_ordTime() gopurs_runtime.Value {
	once_Data_Time_ordTime.Do(func() {
		cache_Data_Time_ordTime = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((((((*Constructor_Data_Time_Time)(x_1.UnsafePtr).V0) == ((*Constructor_Data_Time_Time)(y_2.UnsafePtr).V0)) && (((*Constructor_Data_Time_Time)(x_1.UnsafePtr).V1) == ((*Constructor_Data_Time_Time)(y_2.UnsafePtr).V1))) && (((*Constructor_Data_Time_Time)(x_1.UnsafePtr).V2) == ((*Constructor_Data_Time_Time)(y_2.UnsafePtr).V2))) && (((*Constructor_Data_Time_Time)(x_1.UnsafePtr).V3) == ((*Constructor_Data_Time_Time)(y_2.UnsafePtr).V3)))
})
}))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 -> gopurs_runtime.Value
v_2_0 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V0))
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
// TAST (Let): v1_3_1 -> gopurs_runtime.Value
v1_3_1 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V1), gopurs_runtime.Int((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V1))
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
// TAST (Let): v2_4_2 -> gopurs_runtime.Value
v2_4_2 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V2))
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
__t3 = uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Time_Time)(x_0.UnsafePtr).V3), gopurs_runtime.Int((*Constructor_Data_Time_Time)(y_1.UnsafePtr).V3)).IntVal)
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
	return cache_Data_Time_ordTime
}

var cache_Data_Time_diff gopurs_runtime.Value
var once_Data_Time_diff sync.Once
func Get_Data_Time_diff() gopurs_runtime.Value {
	once_Data_Time_diff.Do(func() {
		cache_Data_Time_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, t1_1_box gopurs_runtime.Value, t2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_diff(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dictDuration_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](t1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](t2_2_box))
})
	})
	return cache_Data_Time_diff
}

var cache_Data_Time_boundedTime gopurs_runtime.Value
var once_Data_Time_boundedTime sync.Once
func Get_Data_Time_boundedTime() gopurs_runtime.Value {
	once_Data_Time_boundedTime.Do(func() {
		cache_Data_Time_boundedTime = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((((((*Constructor_Data_Time_Time)(x_2.UnsafePtr).V0) == ((*Constructor_Data_Time_Time)(y_3.UnsafePtr).V0)) && (((*Constructor_Data_Time_Time)(x_2.UnsafePtr).V1) == ((*Constructor_Data_Time_Time)(y_3.UnsafePtr).V1))) && (((*Constructor_Data_Time_Time)(x_2.UnsafePtr).V2) == ((*Constructor_Data_Time_Time)(y_3.UnsafePtr).V2))) && (((*Constructor_Data_Time_Time)(x_2.UnsafePtr).V3) == ((*Constructor_Data_Time_Time)(y_3.UnsafePtr).V3)))
})
}))
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Time_Time)(x_1.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Data_Time_Time)(y_2.UnsafePtr).V0))
_ = v_3_0
var __t5 uint32
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t5 = 1527465420
goto end_branch_5
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Time_Time)(x_1.UnsafePtr).V1), gopurs_runtime.Int((*Constructor_Data_Time_Time)(y_2.UnsafePtr).V1))
_ = v1_4_1
var __t4 uint32
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t4 = 1527465420
goto end_branch_4
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
// TAST (Let): v2_5_2 -> gopurs_runtime.Value
v2_5_2 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Time_Time)(x_1.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Data_Time_Time)(y_2.UnsafePtr).V2))
_ = v2_5_2
var __t3 uint32
{
if (uint32(v2_5_2.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(v2_5_2.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
__t3 = uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Time_Time)(x_1.UnsafePtr).V3), gopurs_runtime.Int((*Constructor_Data_Time_Time)(y_2.UnsafePtr).V3)).IntVal)
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
}), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Data_Time_Time{1, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedHour(), "bottom").IntVal, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedMinute(), "bottom").IntVal, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedSecond(), "bottom").IntVal, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedMillisecond(), "bottom").IntVal})}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Data_Time_Time{1, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedHour(), "top").IntVal, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedMinute(), "top").IntVal, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedSecond(), "top").IntVal, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedMillisecond(), "top").IntVal})})
	})
	return cache_Data_Time_boundedTime
}

var cache_Data_Time_maxTime gopurs_runtime.Value
var once_Data_Time_maxTime sync.Once
func Get_Data_Time_maxTime() gopurs_runtime.Value {
	once_Data_Time_maxTime.Do(func() {
		cache_Data_Time_maxTime = gopurs_runtime.Float(Call_Data_Time_timeToMillis(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Data_Time_Time{1, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedHour(), "top").IntVal, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedMinute(), "top").IntVal, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedSecond(), "top").IntVal, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedMillisecond(), "top").IntVal})})))
	})
	return cache_Data_Time_maxTime
}

var cache_Data_Time_minTime gopurs_runtime.Value
var once_Data_Time_minTime sync.Once
func Get_Data_Time_minTime() gopurs_runtime.Value {
	once_Data_Time_minTime.Do(func() {
		cache_Data_Time_minTime = gopurs_runtime.Float(Call_Data_Time_timeToMillis(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(&Constructor_Data_Time_Time{1, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedHour(), "bottom").IntVal, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedMinute(), "bottom").IntVal, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedSecond(), "bottom").IntVal, gopurs_runtime.RecordGet(Get_Data_Time_Component_boundedMillisecond(), "bottom").IntVal})})))
	})
	return cache_Data_Time_minTime
}

var cache_Data_Time_adjust gopurs_runtime.Value
var once_Data_Time_adjust sync.Once
func Get_Data_Time_adjust() gopurs_runtime.Value {
	once_Data_Time_adjust.Do(func() {
		cache_Data_Time_adjust = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, d_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Time_adjust(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dictDuration_0_box), d_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](t_2_box)))}
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
return &Constructor_Data_Time_Time{1, (v_1).V0, (v_1).V1, s_0, (v_1).V3}
}

func Call_Data_Time_setMinute(m_0_loop int64, v_1_loop *Constructor_Data_Time_Time) *Constructor_Data_Time_Time {
var m_0 int64 = m_0_loop
_ = m_0
var v_1 *Constructor_Data_Time_Time = v_1_loop
_ = v_1
return &Constructor_Data_Time_Time{1, (v_1).V0, m_0, (v_1).V2, (v_1).V3}
}

func Call_Data_Time_setMillisecond(ms_0_loop int64, v_1_loop *Constructor_Data_Time_Time) *Constructor_Data_Time_Time {
var ms_0 int64 = ms_0_loop
_ = ms_0
var v_1 *Constructor_Data_Time_Time = v_1_loop
_ = v_1
return &Constructor_Data_Time_Time{1, (v_1).V0, (v_1).V1, (v_1).V2, ms_0}
}

func Call_Data_Time_setHour(h_0_loop int64, v_1_loop *Constructor_Data_Time_Time) *Constructor_Data_Time_Time {
var h_0 int64 = h_0_loop
_ = h_0
var v_1 *Constructor_Data_Time_Time = v_1_loop
_ = v_1
return &Constructor_Data_Time_Time{1, h_0, (v_1).V1, (v_1).V2, (v_1).V3}
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
// TAST (Let): hours_1_0 -> float64
hours_1_0 := gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float((v_0) / (3600000.0))).FloatVal()
_ = hours_1_0
// TAST (Let): minutes_2_1 -> float64
minutes_2_1 := gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float(((v_0) - ((hours_1_0) * (3600000.0))) / (60000.0))).FloatVal()
_ = minutes_2_1
// TAST (Let): seconds_3_2 -> float64
seconds_3_2 := gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float(((v_0) - (((hours_1_0) * (3600000.0)) + ((minutes_2_1) * (60000.0)))) / (1000.0))).FloatVal()
_ = seconds_3_2
// TAST (Let): __local_var_4_6 -> *Constructor_Data_Maybe_Just
__local_var_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_Int_unsafeClamp(), gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float(hours_1_0))).IntVal)))
_ = __local_var_4_6
var __t7 *Constructor_Data_Maybe_Just
{
if (__local_var_4_6 != nil) {
__t7 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Time_Time(), gopurs_runtime.Int((__local_var_4_6).V0.IntVal))}
goto end_branch_7
} else {

}
}
{
__t7 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_7:
// TAST (Let): __local_var_4_5 -> *Constructor_Data_Maybe_Just
var __local_var_4_5 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t7)})
// TAST (Let): __local_var_5_8 -> *Constructor_Data_Maybe_Just
__local_var_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_Int_unsafeClamp(), gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float(minutes_2_1))).IntVal)))
_ = __local_var_5_8
var __t9 *Constructor_Data_Maybe_Just
{
if (__local_var_4_5 != nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_4_5).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_8)}))
goto end_branch_9
} else {

}
}
{
if (__local_var_4_5 == nil) {
__t9 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
// TAST (Let): __local_var_4_4 -> *Constructor_Data_Maybe_Just
__local_var_4_4 := __t9
_ = __local_var_4_4
// TAST (Let): __local_var_5_10 -> *Constructor_Data_Maybe_Just
__local_var_5_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_Int_unsafeClamp(), gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float(seconds_3_2))).IntVal)))
_ = __local_var_5_10
var __t11 *Constructor_Data_Maybe_Just
{
if (__local_var_4_4 != nil) {
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_4_4).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_10)}))
goto end_branch_11
} else {

}
}
{
if (__local_var_4_4 == nil) {
__t11 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
// TAST (Let): __local_var_4_3 -> *Constructor_Data_Maybe_Just
__local_var_4_3 := __t11
_ = __local_var_4_3
// TAST (Let): __local_var_5_12 -> *Constructor_Data_Maybe_Just
__local_var_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_Int_unsafeClamp(), gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float((v_0) - ((((hours_1_0) * (3600000.0)) + ((minutes_2_1) * (60000.0))) + ((seconds_3_2) * (1000.0)))))).IntVal)))
_ = __local_var_5_12
var __t13 *Constructor_Data_Maybe_Just
{
if (__local_var_4_3 != nil) {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_4_3).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_12)}))
goto end_branch_13
} else {

}
}
{
if (__local_var_4_3 == nil) {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_13:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](gopurs_runtime.Apply2(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_fromJust__755886620()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t13)}))
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

func Call_Data_Time_diff(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration, t1_1_loop *Constructor_Data_Time_Time, t2_2_loop *Constructor_Data_Time_Time) gopurs_runtime.Value {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration = dictDuration_0_loop
_ = dictDuration_0
var t1_1 *Constructor_Data_Time_Time = t1_1_loop
_ = t1_1
var t2_2 *Constructor_Data_Time_Time = t2_2_loop
_ = t2_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration_0.V1), gopurs_runtime.Float((Call_Data_Time_timeToMillis(t1_1)) + (gopurs_runtime.Apply(Get_Data_Time_Duration_negateDuration__4195558286(), gopurs_runtime.Float(Call_Data_Time_timeToMillis(t2_2))).FloatVal())))
}

func Call_Data_Time_adjust(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration, d_1_loop gopurs_runtime.Value, t_2_loop *Constructor_Data_Time_Time) *Constructor_Data_Tuple_Tuple {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration = dictDuration_0_loop
_ = dictDuration_0
var d_1 gopurs_runtime.Value = d_1_loop
_ = d_1
var t_2 *Constructor_Data_Time_Time = t_2_loop
_ = t_2
// TAST (Let): d_prime_3_0 -> gopurs_runtime.Value
d_prime_3_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration_0.V0), d_1)
_ = d_prime_3_0
// TAST (Let): wholeDays_4_1 -> float64
wholeDays_4_1 := gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float((d_prime_3_0.FloatVal()) / (86400000.0))).FloatVal()
_ = wholeDays_4_1
// TAST (Let): msAdjusted_5_2 -> float64
msAdjusted_5_2 := ((Call_Data_Time_timeToMillis(t_2)) + (d_prime_3_0.FloatVal())) + (gopurs_runtime.Apply(Get_Data_Time_Duration_negateDuration__4195558286(), gopurs_runtime.Float((wholeDays_4_1) * (86400000.0))).FloatVal())
_ = msAdjusted_5_2
var __t8 float64
{
var __t7 bool
{
if (msAdjusted_5_2) > (Get_Data_Time_maxTime().FloatVal()) {
__t7 = true
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
if __t7 {
__t8 = 1.0
goto end_branch_8
} else {

}
}
{
var __t6 float64
{
var __t5 bool
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(msAdjusted_5_2), gopurs_runtime.Float(Get_Data_Time_minTime().FloatVal()))
if (uint32(__t_tag_4.IntVal) == 1527465420) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
if __t5 {
__t6 = -1.0
goto end_branch_6
} else {

}
}
{
__t6 = 0.0
}
end_branch_6:
__t8 = __t6
}
end_branch_8:
// TAST (Let): wrap_6_3 -> gopurs_runtime.Value
var wrap_6_3 gopurs_runtime.Value = gopurs_runtime.Float(__t8)
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Float((wholeDays_4_1) + (wrap_6_3.FloatVal())), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_Time_millisToTime((msAdjusted_5_2) + ((86400000.0) * (-(wrap_6_3.FloatVal())))))}})})
}



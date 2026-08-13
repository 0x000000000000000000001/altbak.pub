package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_DateTime_DateTime gopurs_runtime.Value
var once_Data_DateTime_DateTime sync.Once
func Get_Data_DateTime_DateTime() gopurs_runtime.Value {
	once_Data_DateTime_DateTime.Do(func() {
		cache_Data_DateTime_DateTime = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_Data_DateTime_DateTime{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](value0), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](value1)})}
})
})
	})
	return cache_Data_DateTime_DateTime
}

var cache_Data_DateTime_toRecord gopurs_runtime.Value
var once_Data_DateTime_toRecord sync.Once
func Get_Data_DateTime_toRecord() gopurs_runtime.Value {
	once_Data_DateTime_toRecord.Do(func() {
		cache_Data_DateTime_toRecord = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DateTime_toRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](v_0_box))
})
	})
	return cache_Data_DateTime_toRecord
}

var cache_Data_DateTime_time gopurs_runtime.Value
var once_Data_DateTime_time sync.Once
func Get_Data_DateTime_time() gopurs_runtime.Value {
	once_Data_DateTime_time.Do(func() {
		cache_Data_DateTime_time = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_DateTime_time(gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](v_0_box)))}
})
	})
	return cache_Data_DateTime_time
}

var cache_Data_DateTime_showDateTime gopurs_runtime.Value
var once_Data_DateTime_showDateTime sync.Once
func Get_Data_DateTime_showDateTime() gopurs_runtime.Value {
	once_Data_DateTime_showDateTime.Do(func() {
		cache_Data_DateTime_showDateTime = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(DateTime ") + (gopurs_runtime.Apply(Get_Data_Show_show__1723386194(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V0)}).StrVal())) + (" ")) + (gopurs_runtime.Apply(Get_Data_Show_show__1073032466(), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V1)}).StrVal())) + (")"))
}))
	})
	return cache_Data_DateTime_showDateTime
}

var cache_Data_DateTime_modifyTimeF gopurs_runtime.Value
var once_Data_DateTime_modifyTimeF sync.Once
func Get_Data_DateTime_modifyTimeF() gopurs_runtime.Value {
	once_Data_DateTime_modifyTimeF.Do(func() {
		cache_Data_DateTime_modifyTimeF = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DateTime_modifyTimeF(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](v_2_box))
})
	})
	return cache_Data_DateTime_modifyTimeF
}

var cache_Data_DateTime_modifyTime gopurs_runtime.Value
var once_Data_DateTime_modifyTime sync.Once
func Get_Data_DateTime_modifyTime() gopurs_runtime.Value {
	once_Data_DateTime_modifyTime.Do(func() {
		cache_Data_DateTime_modifyTime = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(Call_Data_DateTime_modifyTime(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](v_1_box)))}
})
	})
	return cache_Data_DateTime_modifyTime
}

var cache_Data_DateTime_modifyDateF gopurs_runtime.Value
var once_Data_DateTime_modifyDateF sync.Once
func Get_Data_DateTime_modifyDateF() gopurs_runtime.Value {
	once_Data_DateTime_modifyDateF.Do(func() {
		cache_Data_DateTime_modifyDateF = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DateTime_modifyDateF(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](v_2_box))
})
	})
	return cache_Data_DateTime_modifyDateF
}

var cache_Data_DateTime_modifyDate gopurs_runtime.Value
var once_Data_DateTime_modifyDate sync.Once
func Get_Data_DateTime_modifyDate() gopurs_runtime.Value {
	once_Data_DateTime_modifyDate.Do(func() {
		cache_Data_DateTime_modifyDate = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(Call_Data_DateTime_modifyDate(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](v_1_box)))}
})
	})
	return cache_Data_DateTime_modifyDate
}

var cache_Data_DateTime_eqDateTime gopurs_runtime.Value
var once_Data_DateTime_eqDateTime sync.Once
func Get_Data_DateTime_eqDateTime() gopurs_runtime.Value {
	once_Data_DateTime_eqDateTime.Do(func() {
		cache_Data_DateTime_eqDateTime = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(Get_Data_Eq_eq__1204755874(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V0)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V0)}).IntVal) != (0)) && ((gopurs_runtime.Apply2(Get_Data_Eq_eq__1287514754(), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1)}).IntVal) != (0)))
})
}))
	})
	return cache_Data_DateTime_eqDateTime
}

var cache_Data_DateTime_ordDateTime gopurs_runtime.Value
var once_Data_DateTime_ordDateTime sync.Once
func Get_Data_DateTime_ordDateTime() gopurs_runtime.Value {
	once_Data_DateTime_ordDateTime.Do(func() {
		cache_Data_DateTime_ordDateTime = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_DateTime_eqDateTime()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 -> gopurs_runtime.Value
v_2_0 := gopurs_runtime.Apply2(Get_Data_Ord_compare__146529112(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V0)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V0)})
_ = v_2_0
var __t1 uint32
{
if (uint32(v_2_0.IntVal) == 1527465420) {
__t1 = 1527465420
goto end_branch_1
} else {

}
}
{
if (uint32(v_2_0.IntVal) == 380165415) {
__t1 = 380165415
goto end_branch_1
} else {

}
}
{
__t1 = uint32(gopurs_runtime.Apply2(Get_Data_Ord_compare__463614392(), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1)}).IntVal)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
})
}))
	})
	return cache_Data_DateTime_ordDateTime
}

var cache_Data_DateTime_diff gopurs_runtime.Value
var once_Data_DateTime_diff sync.Once
func Get_Data_DateTime_diff() gopurs_runtime.Value {
	once_Data_DateTime_diff.Do(func() {
		cache_Data_DateTime_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, dt1_1_box gopurs_runtime.Value, dt2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DateTime_diff(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dictDuration_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](dt1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](dt2_2_box))
})
	})
	return cache_Data_DateTime_diff
}

var cache_Data_DateTime_date gopurs_runtime.Value
var once_Data_DateTime_date sync.Once
func Get_Data_DateTime_date() gopurs_runtime.Value {
	once_Data_DateTime_date.Do(func() {
		cache_Data_DateTime_date = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_DateTime_date(gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](v_0_box)))}
})
	})
	return cache_Data_DateTime_date
}

var cache_Data_DateTime_boundedDateTime gopurs_runtime.Value
var once_Data_DateTime_boundedDateTime sync.Once
func Get_Data_DateTime_boundedDateTime() gopurs_runtime.Value {
	once_Data_DateTime_boundedDateTime.Do(func() {
		cache_Data_DateTime_boundedDateTime = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_DateTime_ordDateTime()
}), gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_Data_DateTime_DateTime{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.RecordGet(Get_Data_Date_boundedDate(), "bottom")), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](gopurs_runtime.RecordGet(Get_Data_Time_boundedTime(), "bottom"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_Data_DateTime_DateTime{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.RecordGet(Get_Data_Date_boundedDate(), "top")), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](gopurs_runtime.RecordGet(Get_Data_Time_boundedTime(), "top"))})})
	})
	return cache_Data_DateTime_boundedDateTime
}

var cache_Data_DateTime_adjust gopurs_runtime.Value
var once_Data_DateTime_adjust sync.Once
func Get_Data_DateTime_adjust() gopurs_runtime.Value {
	once_Data_DateTime_adjust.Do(func() {
		cache_Data_DateTime_adjust = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, d_1_box gopurs_runtime.Value, dt_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_DateTime_adjust(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dictDuration_0_box), d_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](dt_2_box)))}
})
	})
	return cache_Data_DateTime_adjust
}

type Constructor_Data_DateTime_DateTime struct {
	Rc uint32
	V0 *Constructor_Data_Date_Date
	V1 *Constructor_Data_Time_Time
}


func Call_Data_DateTime_toRecord(v_0_loop *Constructor_Data_DateTime_DateTime) gopurs_runtime.Value {
var v_0 *Constructor_Data_DateTime_DateTime = v_0_loop
_ = v_0
var __t12 gopurs_runtime.Value
{
var __t_tag_0 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_0) == 1908470532) {
__t12 = gopurs_runtime.Int(1)
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_1) == 2455627378) {
__t12 = gopurs_runtime.Int(2)
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_2) == 4162469099) {
__t12 = gopurs_runtime.Int(3)
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_3) == 1692989816) {
__t12 = gopurs_runtime.Int(4)
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_4) == 330658827) {
__t12 = gopurs_runtime.Int(5)
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_5) == 4067355978) {
__t12 = gopurs_runtime.Int(6)
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_6) == 2276710548) {
__t12 = gopurs_runtime.Int(7)
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_7) == 243771071) {
__t12 = gopurs_runtime.Int(8)
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_8) == 215731793) {
__t12 = gopurs_runtime.Int(9)
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_9) == 8639228) {
__t12 = gopurs_runtime.Int(10)
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_10) == 49471444) {
__t12 = gopurs_runtime.Int(11)
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_11) == 3889233761) {
__t12 = gopurs_runtime.Int(12)
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return gopurs_runtime.RecordDict([]string{"day", "hour", "millisecond", "minute", "month", "second", "year"}, []gopurs_runtime.Value{gopurs_runtime.Int(((v_0).V0).V2), gopurs_runtime.Int(((v_0).V1).V0), gopurs_runtime.Int(((v_0).V1).V3), gopurs_runtime.Int(((v_0).V1).V1), gopurs_runtime.Int(__t12.IntVal), gopurs_runtime.Int(((v_0).V1).V2), gopurs_runtime.Int(((v_0).V0).V0)})
}

func Call_Data_DateTime_time(v_0_loop *Constructor_Data_DateTime_DateTime) *Constructor_Data_Time_Time {
var v_0 *Constructor_Data_DateTime_DateTime = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_DateTime_modifyTimeF(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_DateTime_DateTime) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_DateTime_DateTime = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Apply(Get_Data_DateTime_DateTime(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((v_2).V0)}), gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((v_2).V1)}))
}

func Call_Data_DateTime_modifyTime(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_DateTime_DateTime) *Constructor_Data_DateTime_DateTime {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_DateTime_DateTime = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_Data_DateTime_DateTime{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((v_1).V1)}))})})
}

func Call_Data_DateTime_modifyDateF(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_DateTime_DateTime) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_DateTime_DateTime = v_2_loop
_ = v_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Time_Time
__local_var_3_0 := (v_2).V1
_ = __local_var_3_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_Data_DateTime_DateTime{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](a_4), __local_var_3_0})}
}), gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((v_2).V0)}))
}

func Call_Data_DateTime_modifyDate(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_DateTime_DateTime) *Constructor_Data_DateTime_DateTime {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_DateTime_DateTime = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_Data_DateTime_DateTime{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((v_1).V0)})), (v_1).V1})})
}

func Call_Data_DateTime_diff(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration, dt1_1_loop *Constructor_Data_DateTime_DateTime, dt2_2_loop *Constructor_Data_DateTime_DateTime) gopurs_runtime.Value {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration = dictDuration_0_loop
_ = dictDuration_0
var dt1_1 *Constructor_Data_DateTime_DateTime = dt1_1_loop
_ = dt1_1
var dt2_2 *Constructor_Data_DateTime_DateTime = dt2_2_loop
_ = dt2_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration_0.V1), gopurs_runtime.Float(gopurs_runtime.UncurriedApp2(Get_Data_DateTime_calcDiff(), Call_Data_DateTime_toRecord(dt1_1), Call_Data_DateTime_toRecord(dt2_2)).FloatVal()))
}

func Call_Data_DateTime_date(v_0_loop *Constructor_Data_DateTime_DateTime) *Constructor_Data_Date_Date {
var v_0 *Constructor_Data_DateTime_DateTime = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_DateTime_adjust(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration, d_1_loop gopurs_runtime.Value, dt_2_loop *Constructor_Data_DateTime_DateTime) *Constructor_Data_Maybe_Just {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration = dictDuration_0_loop
_ = dictDuration_0
var d_1 gopurs_runtime.Value = d_1_loop
_ = d_1
var dt_2 *Constructor_Data_DateTime_DateTime = dt_2_loop
_ = dt_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply4(Get_Data_DateTime_adjustImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Float(gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration_0.V0), d_1).FloatVal()), Call_Data_DateTime_toRecord(dt_2)))
_ = __local_var_3_0
var __t25 gopurs_runtime.Value
{
if (__local_var_3_0 != nil) {
// TAST (Let): __local_var_4_5 -> *Constructor_Data_Maybe_Just
__local_var_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "year").IntVal)))
_ = __local_var_4_5
var __t6 gopurs_runtime.Value
{
if (__local_var_4_5 != nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_5).V0.IntVal))})}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}
end_branch_6:
// TAST (Let): __local_var_4_4 -> *Constructor_Data_Maybe_Just
__local_var_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t6)
_ = __local_var_4_4
// TAST (Let): __local_var_5_7 -> *Constructor_Data_Maybe_Just
__local_var_5_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2309750950(), gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "month").IntVal)))
_ = __local_var_5_7
var __t8 gopurs_runtime.Value
{
if (__local_var_4_4 != nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_4_4).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_7)})))}
goto end_branch_8
} else {

}
}
{
if (__local_var_4_4 == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
// TAST (Let): __local_var_4_3 -> *Constructor_Data_Maybe_Just
__local_var_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t8)
_ = __local_var_4_3
// TAST (Let): __local_var_5_9 -> *Constructor_Data_Maybe_Just
__local_var_5_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "day").IntVal)))
_ = __local_var_5_9
var __t10 gopurs_runtime.Value
{
if (__local_var_4_3 != nil) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_4_3).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_9)})))}
goto end_branch_10
} else {

}
}
{
if (__local_var_4_3 == nil) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
// TAST (Let): __local_var_4_2 -> *Constructor_Data_Maybe_Just
__local_var_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_bindMaybe(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t10))}, Get_Control_Bind_identity()))
_ = __local_var_4_2
var __t11 gopurs_runtime.Value
{
if (__local_var_4_2 != nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_DateTime_DateTime(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((__local_var_4_2).V0))})})}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}
end_branch_11:
// TAST (Let): __local_var_4_1 -> *Constructor_Data_Maybe_Just
__local_var_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t11)
_ = __local_var_4_1
// TAST (Let): __local_var_5_16 -> *Constructor_Data_Maybe_Just
__local_var_5_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "hour").IntVal)))
_ = __local_var_5_16
var __t17 gopurs_runtime.Value
{
if (__local_var_5_16 != nil) {
__t17 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Time_Time(), gopurs_runtime.Int((__local_var_5_16).V0.IntVal))})}
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}
end_branch_17:
// TAST (Let): __local_var_5_15 -> *Constructor_Data_Maybe_Just
__local_var_5_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t17)
_ = __local_var_5_15
// TAST (Let): __local_var_6_18 -> *Constructor_Data_Maybe_Just
__local_var_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "minute").IntVal)))
_ = __local_var_6_18
var __t19 gopurs_runtime.Value
{
if (__local_var_5_15 != nil) {
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_5_15).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_18)})))}
goto end_branch_19
} else {

}
}
{
if (__local_var_5_15 == nil) {
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
// TAST (Let): __local_var_5_14 -> *Constructor_Data_Maybe_Just
__local_var_5_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t19)
_ = __local_var_5_14
// TAST (Let): __local_var_6_20 -> *Constructor_Data_Maybe_Just
__local_var_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "second").IntVal)))
_ = __local_var_6_20
var __t21 gopurs_runtime.Value
{
if (__local_var_5_14 != nil) {
__t21 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_5_14).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_20)})))}
goto end_branch_21
} else {

}
}
{
if (__local_var_5_14 == nil) {
__t21 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
// TAST (Let): __local_var_5_13 -> *Constructor_Data_Maybe_Just
__local_var_5_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t21)
_ = __local_var_5_13
// TAST (Let): __local_var_6_22 -> *Constructor_Data_Maybe_Just
__local_var_6_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2099864294(), gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "millisecond").IntVal)))
_ = __local_var_6_22
var __t23 gopurs_runtime.Value
{
if (__local_var_5_13 != nil) {
__t23 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_5_13).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_22)})))}
goto end_branch_23
} else {

}
}
{
if (__local_var_5_13 == nil) {
__t23 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
// TAST (Let): __local_var_5_12 -> *Constructor_Data_Maybe_Just
__local_var_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t23)
_ = __local_var_5_12
var __t24 gopurs_runtime.Value
{
if (__local_var_4_1 != nil) {
__t24 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_4_1).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_12)})))}
goto end_branch_24
} else {

}
}
{
if (__local_var_4_1 == nil) {
__t24 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
__t25 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t24))}
goto end_branch_25
} else {

}
}
{
if (__local_var_3_0 == nil) {
__t25 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t25)
}

func Get_Data_DateTime_adjustImpl() gopurs_runtime.Value {
	return _Gopurs_Data_DateTime_AdjustImpl
}

func Get_Data_DateTime_calcDiff() gopurs_runtime.Value {
	return _Gopurs_Data_DateTime_CalcDiff
}

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
		cache_Data_DateTime_showDateTime = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(DateTime ") + (gopurs_runtime.Apply(Get_Data_Show_show__1723386194(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V0)}).StrVal())) + (" ")) + (gopurs_runtime.Apply(Get_Data_Show_show__1073032466(), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V1)}).StrVal())) + (")"))
})})}
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
		cache_Data_DateTime_eqDateTime = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(Get_Data_Eq_eq__1204755874(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V0)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V0)}).IntVal) != (0)) && ((gopurs_runtime.Apply2(Get_Data_Eq_eq__1287514754(), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1)}).IntVal) != (0)))
})
})})}
	})
	return cache_Data_DateTime_eqDateTime
}

var cache_Data_DateTime_ordDateTime gopurs_runtime.Value
var once_Data_DateTime_ordDateTime sync.Once
func Get_Data_DateTime_ordDateTime() gopurs_runtime.Value {
	once_Data_DateTime_ordDateTime.Do(func() {
		cache_Data_DateTime_ordDateTime = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_DateTime_eqDateTime()))}
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
})})}
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
		cache_Data_DateTime_boundedDateTime = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(&Constructor_Data_Bounded_Bounded{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_DateTime_ordDateTime()))}
}), gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_Data_DateTime_DateTime{1, &Constructor_Data_Date_Date{1, -271820, 1908470532, 1}, &Constructor_Data_Time_Time{1, 0, 0, 0, 0}})}, gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(&Constructor_Data_DateTime_DateTime{1, &Constructor_Data_Date_Date{1, 275759, 3889233761, 31}, &Constructor_Data_Time_Time{1, 23, 59, 59, 999}})}})}
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
var __t12 int64
{
var __t_tag_0 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_0) == 1908470532) {
__t12 = 1
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_1) == 2455627378) {
__t12 = 2
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_2) == 4162469099) {
__t12 = 3
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_3) == 1692989816) {
__t12 = 4
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_4) == 330658827) {
__t12 = 5
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_5) == 4067355978) {
__t12 = 6
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_6) == 2276710548) {
__t12 = 7
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_7) == 243771071) {
__t12 = 8
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_8) == 215731793) {
__t12 = 9
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_9) == 8639228) {
__t12 = 10
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_10) == 49471444) {
__t12 = 11
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_11) == 3889233761) {
__t12 = 12
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_12:
return gopurs_runtime.RecordDict([]string{"day", "hour", "millisecond", "minute", "month", "second", "year"}, []gopurs_runtime.Value{gopurs_runtime.Int(((v_0).V0).V2), gopurs_runtime.Int(((v_0).V1).V0), gopurs_runtime.Int(((v_0).V1).V3), gopurs_runtime.Int(((v_0).V1).V1), gopurs_runtime.Int(__t12), gopurs_runtime.Int(((v_0).V1).V2), gopurs_runtime.Int(((v_0).V0).V0)})
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
return &Constructor_Data_DateTime_DateTime{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((v_1).V1)}))}
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
return &Constructor_Data_DateTime_DateTime{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((v_1).V0)})), (v_1).V1}
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
var __t57 *Constructor_Data_Maybe_Just
{
if (__local_var_3_0 != nil) {
var __t10 *Constructor_Data_Maybe_Just
{
var __t7 bool
{
if (gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "year").IntVal).IntVal) < (gopurs_runtime.Int(1).IntVal) {
__t7 = false
goto end_branch_7
} else {

}
}
{
__t7 = true
}
end_branch_7:
var __t_and_9 bool = false
if __t7 {

var __t8 bool
{
if (gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "year").IntVal).IntVal) > (gopurs_runtime.Int(31).IntVal) {
__t8 = false
goto end_branch_8
} else {

}
}
{
__t8 = true
}
end_branch_8:
__t_and_9 = __t8
}
if __t_and_9 {
__t10 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "year").IntVal)}
goto end_branch_10
} else {

}
}
{
__t10 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_10:
// TAST (Let): __local_var_4_6 -> *Constructor_Data_Maybe_Just
var __local_var_4_6 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t10)})
var __t11 *Constructor_Data_Maybe_Just
{
if (__local_var_4_6 != nil) {
__t11 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_6).V0.IntVal))}
goto end_branch_11
} else {

}
}
{
__t11 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_11:
// TAST (Let): __local_var_4_5 -> *Constructor_Data_Maybe_Just
var __local_var_4_5 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t11)})
// TAST (Let): __local_var_5_12 -> *Constructor_Data_Maybe_Just
__local_var_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2309750950(), gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "month").IntVal)))
_ = __local_var_5_12
var __t14 *Constructor_Data_Maybe_Just
{
if (__local_var_4_5 != nil) {
var __t13 *Constructor_Data_Maybe_Just
{
if (__local_var_5_12 != nil) {
__t13 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_4_5).V0, (__local_var_5_12).V0)}
goto end_branch_13
} else {

}
}
{
__t13 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
if (__local_var_4_5 == nil) {
__t14 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_14:
// TAST (Let): __local_var_4_4 -> *Constructor_Data_Maybe_Just
var __local_var_4_4 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t14)})
var __t19 *Constructor_Data_Maybe_Just
{
var __t16 bool
{
if (gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "day").IntVal).IntVal) < (gopurs_runtime.Int(1).IntVal) {
__t16 = false
goto end_branch_16
} else {

}
}
{
__t16 = true
}
end_branch_16:
var __t_and_18 bool = false
if __t16 {

var __t17 bool
{
if (gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "day").IntVal).IntVal) > (gopurs_runtime.Int(31).IntVal) {
__t17 = false
goto end_branch_17
} else {

}
}
{
__t17 = true
}
end_branch_17:
__t_and_18 = __t17
}
if __t_and_18 {
__t19 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "day").IntVal)}
goto end_branch_19
} else {

}
}
{
__t19 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_19:
// TAST (Let): __local_var_5_15 -> *Constructor_Data_Maybe_Just
var __local_var_5_15 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t19)})
var __t21 *Constructor_Data_Maybe_Just
{
if (__local_var_4_4 != nil) {
var __t20 *Constructor_Data_Maybe_Just
{
if (__local_var_5_15 != nil) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_4_4).V0, (__local_var_5_15).V0)}
goto end_branch_20
} else {

}
}
{
__t20 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_20:
__t21 = __t20
goto end_branch_21
} else {

}
}
{
if (__local_var_4_4 == nil) {
__t21 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_21
} else {

}
}
{
__t21 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_21:
// TAST (Let): __local_var_4_3 -> *Constructor_Data_Maybe_Just
var __local_var_4_3 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t21)})
var __t22 *Constructor_Data_Maybe_Just
{
if (__local_var_4_3 != nil) {
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((__local_var_4_3).V0)
goto end_branch_22
} else {

}
}
{
if (__local_var_4_3 == nil) {
__t22 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_22:
// TAST (Let): __local_var_4_2 -> *Constructor_Data_Maybe_Just
var __local_var_4_2 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t22)})
var __t23 *Constructor_Data_Maybe_Just
{
if (__local_var_4_2 != nil) {
__t23 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_DateTime_DateTime(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((__local_var_4_2).V0))})}
goto end_branch_23
} else {

}
}
{
__t23 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_23:
// TAST (Let): __local_var_4_1 -> *Constructor_Data_Maybe_Just
var __local_var_4_1 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t23)})
var __t32 *Constructor_Data_Maybe_Just
{
var __t29 bool
{
if (gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "hour").IntVal).IntVal) < (gopurs_runtime.Int(1).IntVal) {
__t29 = false
goto end_branch_29
} else {

}
}
{
__t29 = true
}
end_branch_29:
var __t_and_31 bool = false
if __t29 {

var __t30 bool
{
if (gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "hour").IntVal).IntVal) > (gopurs_runtime.Int(31).IntVal) {
__t30 = false
goto end_branch_30
} else {

}
}
{
__t30 = true
}
end_branch_30:
__t_and_31 = __t30
}
if __t_and_31 {
__t32 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "hour").IntVal)}
goto end_branch_32
} else {

}
}
{
__t32 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_32:
// TAST (Let): __local_var_5_28 -> *Constructor_Data_Maybe_Just
var __local_var_5_28 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t32)})
var __t33 *Constructor_Data_Maybe_Just
{
if (__local_var_5_28 != nil) {
__t33 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Time_Time(), gopurs_runtime.Int((__local_var_5_28).V0.IntVal))}
goto end_branch_33
} else {

}
}
{
__t33 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_33:
// TAST (Let): __local_var_5_27 -> *Constructor_Data_Maybe_Just
var __local_var_5_27 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t33)})
var __t38 *Constructor_Data_Maybe_Just
{
var __t35 bool
{
if (gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "minute").IntVal).IntVal) < (gopurs_runtime.Int(1).IntVal) {
__t35 = false
goto end_branch_35
} else {

}
}
{
__t35 = true
}
end_branch_35:
var __t_and_37 bool = false
if __t35 {

var __t36 bool
{
if (gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "minute").IntVal).IntVal) > (gopurs_runtime.Int(31).IntVal) {
__t36 = false
goto end_branch_36
} else {

}
}
{
__t36 = true
}
end_branch_36:
__t_and_37 = __t36
}
if __t_and_37 {
__t38 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "minute").IntVal)}
goto end_branch_38
} else {

}
}
{
__t38 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_38:
// TAST (Let): __local_var_6_34 -> *Constructor_Data_Maybe_Just
var __local_var_6_34 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t38)})
var __t40 *Constructor_Data_Maybe_Just
{
if (__local_var_5_27 != nil) {
var __t39 *Constructor_Data_Maybe_Just
{
if (__local_var_6_34 != nil) {
__t39 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_5_27).V0, (__local_var_6_34).V0)}
goto end_branch_39
} else {

}
}
{
__t39 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_39:
__t40 = __t39
goto end_branch_40
} else {

}
}
{
if (__local_var_5_27 == nil) {
__t40 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_40
} else {

}
}
{
__t40 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_40:
// TAST (Let): __local_var_5_26 -> *Constructor_Data_Maybe_Just
var __local_var_5_26 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t40)})
var __t45 *Constructor_Data_Maybe_Just
{
var __t42 bool
{
if (gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "second").IntVal).IntVal) < (gopurs_runtime.Int(1).IntVal) {
__t42 = false
goto end_branch_42
} else {

}
}
{
__t42 = true
}
end_branch_42:
var __t_and_44 bool = false
if __t42 {

var __t43 bool
{
if (gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "second").IntVal).IntVal) > (gopurs_runtime.Int(31).IntVal) {
__t43 = false
goto end_branch_43
} else {

}
}
{
__t43 = true
}
end_branch_43:
__t_and_44 = __t43
}
if __t_and_44 {
__t45 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "second").IntVal)}
goto end_branch_45
} else {

}
}
{
__t45 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_45:
// TAST (Let): __local_var_6_41 -> *Constructor_Data_Maybe_Just
var __local_var_6_41 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t45)})
var __t47 *Constructor_Data_Maybe_Just
{
if (__local_var_5_26 != nil) {
var __t46 *Constructor_Data_Maybe_Just
{
if (__local_var_6_41 != nil) {
__t46 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_5_26).V0, (__local_var_6_41).V0)}
goto end_branch_46
} else {

}
}
{
__t46 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_46:
__t47 = __t46
goto end_branch_47
} else {

}
}
{
if (__local_var_5_26 == nil) {
__t47 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_47
} else {

}
}
{
__t47 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_47:
// TAST (Let): __local_var_5_25 -> *Constructor_Data_Maybe_Just
var __local_var_5_25 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t47)})
var __t52 *Constructor_Data_Maybe_Just
{
var __t49 bool
{
if (gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "millisecond").IntVal).IntVal) < (gopurs_runtime.Int(1).IntVal) {
__t49 = false
goto end_branch_49
} else {

}
}
{
__t49 = true
}
end_branch_49:
var __t_and_51 bool = false
if __t49 {

var __t50 bool
{
if (gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "millisecond").IntVal).IntVal) > (gopurs_runtime.Int(31).IntVal) {
__t50 = false
goto end_branch_50
} else {

}
}
{
__t50 = true
}
end_branch_50:
__t_and_51 = __t50
}
if __t_and_51 {
__t52 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_3_0).V0, "millisecond").IntVal)}
goto end_branch_52
} else {

}
}
{
__t52 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_52:
// TAST (Let): __local_var_6_48 -> *Constructor_Data_Maybe_Just
var __local_var_6_48 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t52)})
var __t54 *Constructor_Data_Maybe_Just
{
if (__local_var_5_25 != nil) {
var __t53 *Constructor_Data_Maybe_Just
{
if (__local_var_6_48 != nil) {
__t53 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_5_25).V0, (__local_var_6_48).V0)}
goto end_branch_53
} else {

}
}
{
__t53 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_53:
__t54 = __t53
goto end_branch_54
} else {

}
}
{
if (__local_var_5_25 == nil) {
__t54 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_54
} else {

}
}
{
__t54 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_54:
// TAST (Let): __local_var_5_24 -> *Constructor_Data_Maybe_Just
var __local_var_5_24 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t54)})
var __t56 *Constructor_Data_Maybe_Just
{
if (__local_var_4_1 != nil) {
var __t55 *Constructor_Data_Maybe_Just
{
if (__local_var_5_24 != nil) {
__t55 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_4_1).V0, (__local_var_5_24).V0)}
goto end_branch_55
} else {

}
}
{
__t55 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_55:
__t56 = __t55
goto end_branch_56
} else {

}
}
{
if (__local_var_4_1 == nil) {
__t56 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_56
} else {

}
}
{
__t56 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_56:
__t57 = __t56
goto end_branch_57
} else {

}
}
{
if (__local_var_3_0 == nil) {
__t57 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_57
} else {

}
}
{
__t57 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_57:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t57)})
}

func Get_Data_DateTime_adjustImpl() gopurs_runtime.Value {
	return _Gopurs_Data_DateTime_AdjustImpl
}

func Get_Data_DateTime_calcDiff() gopurs_runtime.Value {
	return _Gopurs_Data_DateTime_CalcDiff
}

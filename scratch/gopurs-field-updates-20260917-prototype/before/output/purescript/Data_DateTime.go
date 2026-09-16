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
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer((&Constructor_Data_DateTime_DateTime{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](value0), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](value1)}))}
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
return func() gopurs_runtime.Value {
				orig := Call_Data_DateTime_toRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](v_0_box))
				_ = orig
				return gopurs_runtime.RecordDict([]string{"day", "hour", "millisecond", "minute", "month", "second", "year"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.day), gopurs_runtime.Int(orig.hour), gopurs_runtime.Int(orig.millisecond), gopurs_runtime.Int(orig.minute), gopurs_runtime.Int(orig.month), gopurs_runtime.Int(orig.second), gopurs_runtime.Int(orig.year)})
				}()
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
		cache_Data_DateTime_showDateTime = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_2380273161_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_DateTime_DateTime]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((((((((("(DateTime ") + (gopurs_runtime.Apply(Rebox_Data_DateTime_1386611502_1726078761(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Date_showDate())).V0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V0)}).StrVal())) + (" (Time (Hour ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V1).V0)).StrVal())) + (") (Minute ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V1).V1)).StrVal())) + (") (Second ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V1).V2)).StrVal())) + (") (Millisecond ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V1).V3)).StrVal())) + (")))"))
})})))}
	})
	return cache_Data_DateTime_showDateTime
}

var cache_Data_DateTime_modifyTimeF gopurs_runtime.Value
var once_Data_DateTime_modifyTimeF sync.Once
func Get_Data_DateTime_modifyTimeF() gopurs_runtime.Value {
	once_Data_DateTime_modifyTimeF.Do(func() {
		cache_Data_DateTime_modifyTimeF = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DateTime_modifyTimeF(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](v_2_box))
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
return Call_Data_DateTime_modifyDateF(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](v_2_box))
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
		cache_Data_DateTime_eqDateTime = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_503123561_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_DateTime_DateTime]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(Rebox_Data_DateTime_3790796878_1957390985(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Date_eqDate())).V0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V0)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V0)}).IntVal) != (0)) && (((((((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V0) == (((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V0)) && ((((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V1) == (((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V1))) && ((((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V2) == (((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V2))) && ((((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V3) == (((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V3))))
})})))}
	})
	return cache_Data_DateTime_eqDateTime
}

var cache_Data_DateTime_ordDateTime gopurs_runtime.Value
var once_Data_DateTime_ordDateTime sync.Once
func Get_Data_DateTime_ordDateTime() gopurs_runtime.Value {
	once_Data_DateTime_ordDateTime.Do(func() {
		cache_Data_DateTime_ordDateTime = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_1888627657_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_DateTime_DateTime]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_503123561_3790796878(Rebox_Data_DateTime_3790796878_503123561(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_DateTime_eqDateTime()))))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_2_0 := uint32(gopurs_runtime.Apply2(Rebox_Data_DateTime_4177771502_758368489(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Date_ordDate())).V1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V0)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V0)}).IntVal)
_ = v_2_0
var __t7 uint32
{
if (v_2_0 == 1527465420) {
__t7 = 1527465420
goto end_branch_7
} else {

}
}
{
if (v_2_0 == 380165415) {
__t7 = 380165415
goto end_branch_7
} else {

}
}
{
// TAST (Let): v_3_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_3_1 := uint32(gopurs_runtime.Apply2(Rebox_Data_DateTime_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V0), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V0)).IntVal)
_ = v_3_1
var __t6 uint32
{
if (v_3_1 == 1527465420) {
__t6 = 1527465420
goto end_branch_6
} else {

}
}
{
if (v_3_1 == 380165415) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
// TAST (Let): v1_4_2 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_4_2 := uint32(gopurs_runtime.Apply2(Rebox_Data_DateTime_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V1), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V1)).IntVal)
_ = v1_4_2
var __t5 uint32
{
if (v1_4_2 == 1527465420) {
__t5 = 1527465420
goto end_branch_5
} else {

}
}
{
if (v1_4_2 == 380165415) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
// TAST (Let): v2_5_3 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v2_5_3 := uint32(gopurs_runtime.Apply2(Rebox_Data_DateTime_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V2), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V2)).IntVal)
_ = v2_5_3
var __t4 uint32
{
if (v2_5_3 == 1527465420) {
__t4 = 1527465420
goto end_branch_4
} else {

}
}
{
if (v2_5_3 == 380165415) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
__t4 = uint32(gopurs_runtime.Apply2(Rebox_Data_DateTime_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V3), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V3)).IntVal)
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
__t6 = __t5
}
end_branch_6:
__t7 = __t6
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t7), UnsafePtr: nil}
})})))}
	})
	return cache_Data_DateTime_ordDateTime
}

var cache_Data_DateTime_diff gopurs_runtime.Value
var once_Data_DateTime_diff sync.Once
func Get_Data_DateTime_diff() gopurs_runtime.Value {
	once_Data_DateTime_diff.Do(func() {
		cache_Data_DateTime_diff = gopurs_runtime.Func(func(dictDuration_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DateTime_diff(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dictDuration_0_box))
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
		cache_Data_DateTime_boundedDateTime = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_3247149001_2094947566((&Constructor_Data_Bounded_Bounded[*Constructor_Data_DateTime_DateTime]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_1888627657_4177771502(Rebox_Data_DateTime_4177771502_1888627657(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_DateTime_ordDateTime()))))}
}), (&Constructor_Data_DateTime_DateTime{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_3523628265_2094947566(Rebox_Data_DateTime_2094947566_3523628265(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Date_boundedDate()))))})), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_4136451977_2094947566(Rebox_Data_DateTime_2094947566_4136451977(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_boundedTime()))))}))}), (&Constructor_Data_DateTime_DateTime{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_3523628265_2094947566(Rebox_Data_DateTime_2094947566_3523628265(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Date_boundedDate()))))})), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_4136451977_2094947566(Rebox_Data_DateTime_2094947566_4136451977(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Time_boundedTime()))))}))})})))}
	})
	return cache_Data_DateTime_boundedDateTime
}

var cache_Data_DateTime_adjust gopurs_runtime.Value
var once_Data_DateTime_adjust sync.Once
func Get_Data_DateTime_adjust() gopurs_runtime.Value {
	once_Data_DateTime_adjust.Do(func() {
		cache_Data_DateTime_adjust = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, d_1_box gopurs_runtime.Value, dt_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_DateTime_adjust(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dictDuration_0_box), d_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](dt_2_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_DateTime_adjust
}

type Constructor_Data_DateTime_DateTime struct {
	Rc uint32
	V0 *Constructor_Data_Date_Date
	V1 *Constructor_Data_Time_Time
}


func Call_Data_DateTime_toRecord(v_0_loop *Constructor_Data_DateTime_DateTime) struct{
	day int64
	hour int64
	millisecond int64
	minute int64
	month int64
	second int64
	year int64
} {
var v_0 *Constructor_Data_DateTime_DateTime = v_0_loop
_ = v_0
var __t12 int64
{
var __t_tag_0 uint32 = ((v_0).V0).V1
_ = __t_tag_0
if (uint32(__t_tag_0) == 1908470532) {
__t12 = int64(1)
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = ((v_0).V0).V1
_ = __t_tag_1
if (uint32(__t_tag_1) == 2455627378) {
__t12 = int64(2)
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = ((v_0).V0).V1
_ = __t_tag_2
if (uint32(__t_tag_2) == 4162469099) {
__t12 = int64(3)
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = ((v_0).V0).V1
_ = __t_tag_3
if (uint32(__t_tag_3) == 1692989816) {
__t12 = int64(4)
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = ((v_0).V0).V1
_ = __t_tag_4
if (uint32(__t_tag_4) == 330658827) {
__t12 = int64(5)
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = ((v_0).V0).V1
_ = __t_tag_5
if (uint32(__t_tag_5) == 4067355978) {
__t12 = int64(6)
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = ((v_0).V0).V1
_ = __t_tag_6
if (uint32(__t_tag_6) == 2276710548) {
__t12 = int64(7)
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = ((v_0).V0).V1
_ = __t_tag_7
if (uint32(__t_tag_7) == 243771071) {
__t12 = int64(8)
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = ((v_0).V0).V1
_ = __t_tag_8
if (uint32(__t_tag_8) == 215731793) {
__t12 = int64(9)
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = ((v_0).V0).V1
_ = __t_tag_9
if (uint32(__t_tag_9) == 8639228) {
__t12 = int64(10)
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = ((v_0).V0).V1
_ = __t_tag_10
if (uint32(__t_tag_10) == 49471444) {
__t12 = int64(11)
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = ((v_0).V0).V1
_ = __t_tag_11
if (uint32(__t_tag_11) == 3889233761) {
__t12 = int64(12)
goto end_branch_12
} else {

}
}
{
__t12 = func() int64 { panic("Failed pattern match") }()
}
end_branch_12:
return struct{
	day int64
	hour int64
	millisecond int64
	minute int64
	month int64
	second int64
	year int64
}{((v_0).V0).V2, ((v_0).V1).V0, ((v_0).V1).V3, ((v_0).V1).V1, __t12, ((v_0).V1).V2, ((v_0).V0).V0}
}

func Call_Data_DateTime_time(v_0_loop *Constructor_Data_DateTime_DateTime) *Constructor_Data_Time_Time {
var v_0 *Constructor_Data_DateTime_DateTime = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_DateTime_modifyTimeF(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_DateTime_DateTime) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_DateTime_DateTime = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Apply(Get_Data_DateTime_DateTime(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((v_2).V0)}), gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((v_2).V1)}))
}

func Call_Data_DateTime_modifyTime(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_DateTime_DateTime) *Constructor_Data_DateTime_DateTime {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_DateTime_DateTime = v_1_loop
_ = v_1
return (&Constructor_Data_DateTime_DateTime{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((v_1).V1)}))})
}

func Call_Data_DateTime_modifyDateF(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_DateTime_DateTime) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_DateTime_DateTime = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Apply2(Get_Data_Function_flip__1501157206(), Get_Data_DateTime_DateTime(), gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((v_2).V1)}), gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((v_2).V0)}))
}

func Call_Data_DateTime_modifyDate(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_DateTime_DateTime) *Constructor_Data_DateTime_DateTime {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_DateTime_DateTime = v_1_loop
_ = v_1
return (&Constructor_Data_DateTime_DateTime{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((v_1).V0)})), (v_1).V1})
}

func Call_Data_DateTime_diff(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
// TAST (Let): toDuration_1_0 shape=App(Var) bindingType=(Func [Number] (TypeVar d$scope8))
toDuration_1_0 := Call_Data_Time_Duration_toDuration(dictDuration_0)
_ = toDuration_1_0
return gopurs_runtime.Func2(func(dt1_2 gopurs_runtime.Value, dt2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(toDuration_1_0, gopurs_runtime.Float(gopurs_runtime.UncurriedApp2(Get_Data_DateTime_calcDiff(), func() gopurs_runtime.Value {
				orig := Call_Data_DateTime_toRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](dt1_2))
				_ = orig
				return gopurs_runtime.RecordDict([]string{"day", "hour", "millisecond", "minute", "month", "second", "year"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.day), gopurs_runtime.Int(orig.hour), gopurs_runtime.Int(orig.millisecond), gopurs_runtime.Int(orig.minute), gopurs_runtime.Int(orig.month), gopurs_runtime.Int(orig.second), gopurs_runtime.Int(orig.year)})
				}(), func() gopurs_runtime.Value {
				orig := Call_Data_DateTime_toRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](dt2_3))
				_ = orig
				return gopurs_runtime.RecordDict([]string{"day", "hour", "millisecond", "minute", "month", "second", "year"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.day), gopurs_runtime.Int(orig.hour), gopurs_runtime.Int(orig.millisecond), gopurs_runtime.Int(orig.minute), gopurs_runtime.Int(orig.month), gopurs_runtime.Int(orig.second), gopurs_runtime.Int(orig.year)})
				}()).FloatVal()))
})
}

func Call_Data_DateTime_date(v_0_loop *Constructor_Data_DateTime_DateTime) *Constructor_Data_Date_Date {
var v_0 *Constructor_Data_DateTime_DateTime = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_DateTime_adjust(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value], d_1_loop gopurs_runtime.Value, dt_2_loop *Constructor_Data_DateTime_DateTime) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var d_1 gopurs_runtime.Value = d_1_loop
_ = d_1
var dt_2 *Constructor_Data_DateTime_DateTime = dt_2_loop
_ = dt_2
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
__local_var_3_0 := gopurs_runtime.Apply4(Get_Data_DateTime_adjustImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))}, gopurs_runtime.Float(gopurs_runtime.Apply(dictDuration_0.V0, d_1).FloatVal()), func() gopurs_runtime.Value {
				orig := Call_Data_DateTime_toRecord(dt_2)
				_ = orig
				return gopurs_runtime.RecordDict([]string{"day", "hour", "millisecond", "minute", "month", "second", "year"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.day), gopurs_runtime.Int(orig.hour), gopurs_runtime.Int(orig.millisecond), gopurs_runtime.Int(orig.minute), gopurs_runtime.Int(orig.month), gopurs_runtime.Int(orig.second), gopurs_runtime.Int(orig.year)})
				}())
_ = __local_var_3_0
var __t338 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
var __t3 *Constructor_Data_Maybe_Just[int64]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "year").IntVal) >= (int64(-271820))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "year").IntVal) <= (int64(275759))) {
__t3 = Rebox_Data_DateTime_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "year").IntVal), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_3
} else {

}
}
{
__t3 = Rebox_Data_DateTime_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_3:
// TAST (Let): __local_var_4_2 shape=Branch(Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_4_2 := __t3
_ = __local_var_4_2
var __t5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_4_2 != nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_2).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_5:
// TAST (Let): __local_var_5_4 shape=Branch(Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope70)])
__local_var_5_4 := __t5
_ = __local_var_5_4
var __t317 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(1)) {
var __t40 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
// TAST (Let): __local_var_6_28 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_28 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_28
var __t39 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t31 bool
{
var __t29 bool
{
if (__local_var_5_4 != nil) {
__t29 = true
goto end_branch_29
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t29 = false
goto end_branch_29
} else {

}
}
{
__t29 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_29:
if __t29 {
__t31 = true
goto end_branch_31
} else {

}
}
{
var __t30 bool
{
if (__local_var_5_4 != nil) {
__t30 = false
goto end_branch_30
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t30 = true
goto end_branch_30
} else {

}
}
{
__t30 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_30:
if __t30 {
__t31 = false
goto end_branch_31
} else {

}
}
{
__t31 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_31:
if __t31 {
var __t35 gopurs_runtime.Value
{
var __t32 bool
{
if (__local_var_5_4 != nil) {
__t32 = true
goto end_branch_32
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t32 = false
goto end_branch_32
} else {

}
}
{
__t32 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_32:
if __t32 {
var __t33 gopurs_runtime.Value
{
if (__local_var_5_4 != nil) {
__t33 = gopurs_runtime.Apply((__local_var_5_4).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil})
goto end_branch_33
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t33 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_33
} else {

}
}
{
__t33 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_33:
__t35 = gopurs_runtime.Apply(__t33, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal))
goto end_branch_35
} else {

}
}
{
var __t34 bool
{
if (__local_var_5_4 != nil) {
__t34 = false
goto end_branch_34
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t34 = true
goto end_branch_34
} else {

}
}
{
__t34 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_34:
if __t34 {
__t35 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_35
} else {

}
}
{
__t35 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_35:
__t39 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_28, __t35))
goto end_branch_39
} else {

}
}
{
var __t38 bool
{
var __t36 bool
{
if (__local_var_5_4 != nil) {
__t36 = true
goto end_branch_36
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t36 = false
goto end_branch_36
} else {

}
}
{
__t36 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_36:
if __t36 {
__t38 = false
goto end_branch_38
} else {

}
}
{
var __t37 bool
{
if (__local_var_5_4 != nil) {
__t37 = false
goto end_branch_37
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t37 = true
goto end_branch_37
} else {

}
}
{
__t37 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_37:
if __t37 {
__t38 = true
goto end_branch_38
} else {

}
}
{
__t38 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_38:
if __t38 {
__t39 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_39
} else {

}
}
{
__t39 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_39:
__t40 = __t39
goto end_branch_40
} else {

}
}
{
// TAST (Let): __local_var_6_17 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_17 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_17
var __t27 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t20 bool
{
var __t18 bool
{
if (__local_var_5_4 != nil) {
__t18 = true
goto end_branch_18
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t18 = false
goto end_branch_18
} else {

}
}
{
__t18 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_18:
if __t18 {
__t20 = false
goto end_branch_20
} else {

}
}
{
var __t19 bool
{
if (__local_var_5_4 != nil) {
__t19 = false
goto end_branch_19
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t19 = true
goto end_branch_19
} else {

}
}
{
__t19 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_19:
if __t19 {
__t20 = false
goto end_branch_20
} else {

}
}
{
__t20 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_20:
if __t20 {
var __t23 gopurs_runtime.Value
{
var __t21 bool
{
if (__local_var_5_4 != nil) {
__t21 = true
goto end_branch_21
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t21 = false
goto end_branch_21
} else {

}
}
{
__t21 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_21:
if __t21 {
__t23 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_23
} else {

}
}
{
var __t22 bool
{
if (__local_var_5_4 != nil) {
__t22 = false
goto end_branch_22
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t22 = true
goto end_branch_22
} else {

}
}
{
__t22 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_22:
if __t22 {
__t23 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t27 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_17, __t23))
goto end_branch_27
} else {

}
}
{
var __t26 bool
{
var __t24 bool
{
if (__local_var_5_4 != nil) {
__t24 = true
goto end_branch_24
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t24 = false
goto end_branch_24
} else {

}
}
{
__t24 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_24:
if __t24 {
__t26 = true
goto end_branch_26
} else {

}
}
{
var __t25 bool
{
if (__local_var_5_4 != nil) {
__t25 = false
goto end_branch_25
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t25 = true
goto end_branch_25
} else {

}
}
{
__t25 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_25:
if __t25 {
__t26 = true
goto end_branch_26
} else {

}
}
{
__t26 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_26:
if __t26 {
__t27 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_27
} else {

}
}
{
__t27 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_27:
__t40 = __t27
}
end_branch_40:
__t317 = __t40
goto end_branch_317
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(2)) {
var __t64 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
// TAST (Let): __local_var_6_52 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_52 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_52
var __t63 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t55 bool
{
var __t53 bool
{
if (__local_var_5_4 != nil) {
__t53 = true
goto end_branch_53
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t53 = false
goto end_branch_53
} else {

}
}
{
__t53 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_53:
if __t53 {
__t55 = true
goto end_branch_55
} else {

}
}
{
var __t54 bool
{
if (__local_var_5_4 != nil) {
__t54 = false
goto end_branch_54
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t54 = true
goto end_branch_54
} else {

}
}
{
__t54 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_54:
if __t54 {
__t55 = false
goto end_branch_55
} else {

}
}
{
__t55 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_55:
if __t55 {
var __t59 gopurs_runtime.Value
{
var __t56 bool
{
if (__local_var_5_4 != nil) {
__t56 = true
goto end_branch_56
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t56 = false
goto end_branch_56
} else {

}
}
{
__t56 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_56:
if __t56 {
var __t57 gopurs_runtime.Value
{
if (__local_var_5_4 != nil) {
__t57 = gopurs_runtime.Apply((__local_var_5_4).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil})
goto end_branch_57
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t57 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_57
} else {

}
}
{
__t57 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_57:
__t59 = gopurs_runtime.Apply(__t57, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal))
goto end_branch_59
} else {

}
}
{
var __t58 bool
{
if (__local_var_5_4 != nil) {
__t58 = false
goto end_branch_58
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t58 = true
goto end_branch_58
} else {

}
}
{
__t58 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_58:
if __t58 {
__t59 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_59
} else {

}
}
{
__t59 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_59:
__t63 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_52, __t59))
goto end_branch_63
} else {

}
}
{
var __t62 bool
{
var __t60 bool
{
if (__local_var_5_4 != nil) {
__t60 = true
goto end_branch_60
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t60 = false
goto end_branch_60
} else {

}
}
{
__t60 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_60:
if __t60 {
__t62 = false
goto end_branch_62
} else {

}
}
{
var __t61 bool
{
if (__local_var_5_4 != nil) {
__t61 = false
goto end_branch_61
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t61 = true
goto end_branch_61
} else {

}
}
{
__t61 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_61:
if __t61 {
__t62 = true
goto end_branch_62
} else {

}
}
{
__t62 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_62:
if __t62 {
__t63 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_63
} else {

}
}
{
__t63 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_63:
__t64 = __t63
goto end_branch_64
} else {

}
}
{
// TAST (Let): __local_var_6_41 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_41 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_41
var __t51 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t44 bool
{
var __t42 bool
{
if (__local_var_5_4 != nil) {
__t42 = true
goto end_branch_42
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t42 = false
goto end_branch_42
} else {

}
}
{
__t42 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_42:
if __t42 {
__t44 = false
goto end_branch_44
} else {

}
}
{
var __t43 bool
{
if (__local_var_5_4 != nil) {
__t43 = false
goto end_branch_43
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t43 = true
goto end_branch_43
} else {

}
}
{
__t43 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_43:
if __t43 {
__t44 = false
goto end_branch_44
} else {

}
}
{
__t44 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_44:
if __t44 {
var __t47 gopurs_runtime.Value
{
var __t45 bool
{
if (__local_var_5_4 != nil) {
__t45 = true
goto end_branch_45
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t45 = false
goto end_branch_45
} else {

}
}
{
__t45 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_45:
if __t45 {
__t47 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_47
} else {

}
}
{
var __t46 bool
{
if (__local_var_5_4 != nil) {
__t46 = false
goto end_branch_46
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t46 = true
goto end_branch_46
} else {

}
}
{
__t46 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_46:
if __t46 {
__t47 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_47
} else {

}
}
{
__t47 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_47:
__t51 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_41, __t47))
goto end_branch_51
} else {

}
}
{
var __t50 bool
{
var __t48 bool
{
if (__local_var_5_4 != nil) {
__t48 = true
goto end_branch_48
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t48 = false
goto end_branch_48
} else {

}
}
{
__t48 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_48:
if __t48 {
__t50 = true
goto end_branch_50
} else {

}
}
{
var __t49 bool
{
if (__local_var_5_4 != nil) {
__t49 = false
goto end_branch_49
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t49 = true
goto end_branch_49
} else {

}
}
{
__t49 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_49:
if __t49 {
__t50 = true
goto end_branch_50
} else {

}
}
{
__t50 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_50:
if __t50 {
__t51 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_51
} else {

}
}
{
__t51 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_51:
__t64 = __t51
}
end_branch_64:
__t317 = __t64
goto end_branch_317
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(3)) {
var __t88 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
// TAST (Let): __local_var_6_76 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_76 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_76
var __t87 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t79 bool
{
var __t77 bool
{
if (__local_var_5_4 != nil) {
__t77 = true
goto end_branch_77
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t77 = false
goto end_branch_77
} else {

}
}
{
__t77 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_77:
if __t77 {
__t79 = true
goto end_branch_79
} else {

}
}
{
var __t78 bool
{
if (__local_var_5_4 != nil) {
__t78 = false
goto end_branch_78
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t78 = true
goto end_branch_78
} else {

}
}
{
__t78 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_78:
if __t78 {
__t79 = false
goto end_branch_79
} else {

}
}
{
__t79 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_79:
if __t79 {
var __t83 gopurs_runtime.Value
{
var __t80 bool
{
if (__local_var_5_4 != nil) {
__t80 = true
goto end_branch_80
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t80 = false
goto end_branch_80
} else {

}
}
{
__t80 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_80:
if __t80 {
var __t81 gopurs_runtime.Value
{
if (__local_var_5_4 != nil) {
__t81 = gopurs_runtime.Apply((__local_var_5_4).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil})
goto end_branch_81
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t81 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_81
} else {

}
}
{
__t81 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_81:
__t83 = gopurs_runtime.Apply(__t81, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal))
goto end_branch_83
} else {

}
}
{
var __t82 bool
{
if (__local_var_5_4 != nil) {
__t82 = false
goto end_branch_82
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t82 = true
goto end_branch_82
} else {

}
}
{
__t82 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_82:
if __t82 {
__t83 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_83
} else {

}
}
{
__t83 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_83:
__t87 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_76, __t83))
goto end_branch_87
} else {

}
}
{
var __t86 bool
{
var __t84 bool
{
if (__local_var_5_4 != nil) {
__t84 = true
goto end_branch_84
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t84 = false
goto end_branch_84
} else {

}
}
{
__t84 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_84:
if __t84 {
__t86 = false
goto end_branch_86
} else {

}
}
{
var __t85 bool
{
if (__local_var_5_4 != nil) {
__t85 = false
goto end_branch_85
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t85 = true
goto end_branch_85
} else {

}
}
{
__t85 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_85:
if __t85 {
__t86 = true
goto end_branch_86
} else {

}
}
{
__t86 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_86:
if __t86 {
__t87 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_87
} else {

}
}
{
__t87 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_87:
__t88 = __t87
goto end_branch_88
} else {

}
}
{
// TAST (Let): __local_var_6_65 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_65 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_65
var __t75 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t68 bool
{
var __t66 bool
{
if (__local_var_5_4 != nil) {
__t66 = true
goto end_branch_66
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t66 = false
goto end_branch_66
} else {

}
}
{
__t66 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_66:
if __t66 {
__t68 = false
goto end_branch_68
} else {

}
}
{
var __t67 bool
{
if (__local_var_5_4 != nil) {
__t67 = false
goto end_branch_67
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t67 = true
goto end_branch_67
} else {

}
}
{
__t67 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_67:
if __t67 {
__t68 = false
goto end_branch_68
} else {

}
}
{
__t68 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_68:
if __t68 {
var __t71 gopurs_runtime.Value
{
var __t69 bool
{
if (__local_var_5_4 != nil) {
__t69 = true
goto end_branch_69
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t69 = false
goto end_branch_69
} else {

}
}
{
__t69 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_69:
if __t69 {
__t71 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_71
} else {

}
}
{
var __t70 bool
{
if (__local_var_5_4 != nil) {
__t70 = false
goto end_branch_70
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t70 = true
goto end_branch_70
} else {

}
}
{
__t70 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_70:
if __t70 {
__t71 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_71
} else {

}
}
{
__t71 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_71:
__t75 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_65, __t71))
goto end_branch_75
} else {

}
}
{
var __t74 bool
{
var __t72 bool
{
if (__local_var_5_4 != nil) {
__t72 = true
goto end_branch_72
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t72 = false
goto end_branch_72
} else {

}
}
{
__t72 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_72:
if __t72 {
__t74 = true
goto end_branch_74
} else {

}
}
{
var __t73 bool
{
if (__local_var_5_4 != nil) {
__t73 = false
goto end_branch_73
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t73 = true
goto end_branch_73
} else {

}
}
{
__t73 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_73:
if __t73 {
__t74 = true
goto end_branch_74
} else {

}
}
{
__t74 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_74:
if __t74 {
__t75 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_75
} else {

}
}
{
__t75 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_75:
__t88 = __t75
}
end_branch_88:
__t317 = __t88
goto end_branch_317
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(4)) {
var __t112 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
// TAST (Let): __local_var_6_100 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_100 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_100
var __t111 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t103 bool
{
var __t101 bool
{
if (__local_var_5_4 != nil) {
__t101 = true
goto end_branch_101
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t101 = false
goto end_branch_101
} else {

}
}
{
__t101 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_101:
if __t101 {
__t103 = true
goto end_branch_103
} else {

}
}
{
var __t102 bool
{
if (__local_var_5_4 != nil) {
__t102 = false
goto end_branch_102
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t102 = true
goto end_branch_102
} else {

}
}
{
__t102 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_102:
if __t102 {
__t103 = false
goto end_branch_103
} else {

}
}
{
__t103 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_103:
if __t103 {
var __t107 gopurs_runtime.Value
{
var __t104 bool
{
if (__local_var_5_4 != nil) {
__t104 = true
goto end_branch_104
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t104 = false
goto end_branch_104
} else {

}
}
{
__t104 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_104:
if __t104 {
var __t105 gopurs_runtime.Value
{
if (__local_var_5_4 != nil) {
__t105 = gopurs_runtime.Apply((__local_var_5_4).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil})
goto end_branch_105
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t105 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_105
} else {

}
}
{
__t105 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_105:
__t107 = gopurs_runtime.Apply(__t105, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal))
goto end_branch_107
} else {

}
}
{
var __t106 bool
{
if (__local_var_5_4 != nil) {
__t106 = false
goto end_branch_106
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t106 = true
goto end_branch_106
} else {

}
}
{
__t106 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_106:
if __t106 {
__t107 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_107
} else {

}
}
{
__t107 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_107:
__t111 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_100, __t107))
goto end_branch_111
} else {

}
}
{
var __t110 bool
{
var __t108 bool
{
if (__local_var_5_4 != nil) {
__t108 = true
goto end_branch_108
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t108 = false
goto end_branch_108
} else {

}
}
{
__t108 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_108:
if __t108 {
__t110 = false
goto end_branch_110
} else {

}
}
{
var __t109 bool
{
if (__local_var_5_4 != nil) {
__t109 = false
goto end_branch_109
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t109 = true
goto end_branch_109
} else {

}
}
{
__t109 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_109:
if __t109 {
__t110 = true
goto end_branch_110
} else {

}
}
{
__t110 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_110:
if __t110 {
__t111 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_111
} else {

}
}
{
__t111 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_111:
__t112 = __t111
goto end_branch_112
} else {

}
}
{
// TAST (Let): __local_var_6_89 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_89 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_89
var __t99 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t92 bool
{
var __t90 bool
{
if (__local_var_5_4 != nil) {
__t90 = true
goto end_branch_90
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t90 = false
goto end_branch_90
} else {

}
}
{
__t90 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_90:
if __t90 {
__t92 = false
goto end_branch_92
} else {

}
}
{
var __t91 bool
{
if (__local_var_5_4 != nil) {
__t91 = false
goto end_branch_91
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t91 = true
goto end_branch_91
} else {

}
}
{
__t91 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_91:
if __t91 {
__t92 = false
goto end_branch_92
} else {

}
}
{
__t92 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_92:
if __t92 {
var __t95 gopurs_runtime.Value
{
var __t93 bool
{
if (__local_var_5_4 != nil) {
__t93 = true
goto end_branch_93
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t93 = false
goto end_branch_93
} else {

}
}
{
__t93 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_93:
if __t93 {
__t95 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_95
} else {

}
}
{
var __t94 bool
{
if (__local_var_5_4 != nil) {
__t94 = false
goto end_branch_94
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t94 = true
goto end_branch_94
} else {

}
}
{
__t94 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_94:
if __t94 {
__t95 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_95
} else {

}
}
{
__t95 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_95:
__t99 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_89, __t95))
goto end_branch_99
} else {

}
}
{
var __t98 bool
{
var __t96 bool
{
if (__local_var_5_4 != nil) {
__t96 = true
goto end_branch_96
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t96 = false
goto end_branch_96
} else {

}
}
{
__t96 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_96:
if __t96 {
__t98 = true
goto end_branch_98
} else {

}
}
{
var __t97 bool
{
if (__local_var_5_4 != nil) {
__t97 = false
goto end_branch_97
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t97 = true
goto end_branch_97
} else {

}
}
{
__t97 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_97:
if __t97 {
__t98 = true
goto end_branch_98
} else {

}
}
{
__t98 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_98:
if __t98 {
__t99 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_99
} else {

}
}
{
__t99 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_99:
__t112 = __t99
}
end_branch_112:
__t317 = __t112
goto end_branch_317
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(5)) {
var __t136 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
// TAST (Let): __local_var_6_124 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_124 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_124
var __t135 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t127 bool
{
var __t125 bool
{
if (__local_var_5_4 != nil) {
__t125 = true
goto end_branch_125
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t125 = false
goto end_branch_125
} else {

}
}
{
__t125 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_125:
if __t125 {
__t127 = true
goto end_branch_127
} else {

}
}
{
var __t126 bool
{
if (__local_var_5_4 != nil) {
__t126 = false
goto end_branch_126
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t126 = true
goto end_branch_126
} else {

}
}
{
__t126 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_126:
if __t126 {
__t127 = false
goto end_branch_127
} else {

}
}
{
__t127 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_127:
if __t127 {
var __t131 gopurs_runtime.Value
{
var __t128 bool
{
if (__local_var_5_4 != nil) {
__t128 = true
goto end_branch_128
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t128 = false
goto end_branch_128
} else {

}
}
{
__t128 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_128:
if __t128 {
var __t129 gopurs_runtime.Value
{
if (__local_var_5_4 != nil) {
__t129 = gopurs_runtime.Apply((__local_var_5_4).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil})
goto end_branch_129
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t129 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_129
} else {

}
}
{
__t129 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_129:
__t131 = gopurs_runtime.Apply(__t129, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal))
goto end_branch_131
} else {

}
}
{
var __t130 bool
{
if (__local_var_5_4 != nil) {
__t130 = false
goto end_branch_130
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t130 = true
goto end_branch_130
} else {

}
}
{
__t130 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_130:
if __t130 {
__t131 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_131
} else {

}
}
{
__t131 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_131:
__t135 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_124, __t131))
goto end_branch_135
} else {

}
}
{
var __t134 bool
{
var __t132 bool
{
if (__local_var_5_4 != nil) {
__t132 = true
goto end_branch_132
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t132 = false
goto end_branch_132
} else {

}
}
{
__t132 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_132:
if __t132 {
__t134 = false
goto end_branch_134
} else {

}
}
{
var __t133 bool
{
if (__local_var_5_4 != nil) {
__t133 = false
goto end_branch_133
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t133 = true
goto end_branch_133
} else {

}
}
{
__t133 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_133:
if __t133 {
__t134 = true
goto end_branch_134
} else {

}
}
{
__t134 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_134:
if __t134 {
__t135 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_135
} else {

}
}
{
__t135 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_135:
__t136 = __t135
goto end_branch_136
} else {

}
}
{
// TAST (Let): __local_var_6_113 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_113 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_113
var __t123 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t116 bool
{
var __t114 bool
{
if (__local_var_5_4 != nil) {
__t114 = true
goto end_branch_114
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t114 = false
goto end_branch_114
} else {

}
}
{
__t114 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_114:
if __t114 {
__t116 = false
goto end_branch_116
} else {

}
}
{
var __t115 bool
{
if (__local_var_5_4 != nil) {
__t115 = false
goto end_branch_115
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t115 = true
goto end_branch_115
} else {

}
}
{
__t115 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_115:
if __t115 {
__t116 = false
goto end_branch_116
} else {

}
}
{
__t116 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_116:
if __t116 {
var __t119 gopurs_runtime.Value
{
var __t117 bool
{
if (__local_var_5_4 != nil) {
__t117 = true
goto end_branch_117
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t117 = false
goto end_branch_117
} else {

}
}
{
__t117 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_117:
if __t117 {
__t119 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_119
} else {

}
}
{
var __t118 bool
{
if (__local_var_5_4 != nil) {
__t118 = false
goto end_branch_118
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t118 = true
goto end_branch_118
} else {

}
}
{
__t118 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_118:
if __t118 {
__t119 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_119
} else {

}
}
{
__t119 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_119:
__t123 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_113, __t119))
goto end_branch_123
} else {

}
}
{
var __t122 bool
{
var __t120 bool
{
if (__local_var_5_4 != nil) {
__t120 = true
goto end_branch_120
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t120 = false
goto end_branch_120
} else {

}
}
{
__t120 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_120:
if __t120 {
__t122 = true
goto end_branch_122
} else {

}
}
{
var __t121 bool
{
if (__local_var_5_4 != nil) {
__t121 = false
goto end_branch_121
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t121 = true
goto end_branch_121
} else {

}
}
{
__t121 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_121:
if __t121 {
__t122 = true
goto end_branch_122
} else {

}
}
{
__t122 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_122:
if __t122 {
__t123 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_123
} else {

}
}
{
__t123 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_123:
__t136 = __t123
}
end_branch_136:
__t317 = __t136
goto end_branch_317
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(6)) {
var __t160 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
// TAST (Let): __local_var_6_148 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_148 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_148
var __t159 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t151 bool
{
var __t149 bool
{
if (__local_var_5_4 != nil) {
__t149 = true
goto end_branch_149
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t149 = false
goto end_branch_149
} else {

}
}
{
__t149 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_149:
if __t149 {
__t151 = true
goto end_branch_151
} else {

}
}
{
var __t150 bool
{
if (__local_var_5_4 != nil) {
__t150 = false
goto end_branch_150
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t150 = true
goto end_branch_150
} else {

}
}
{
__t150 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_150:
if __t150 {
__t151 = false
goto end_branch_151
} else {

}
}
{
__t151 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_151:
if __t151 {
var __t155 gopurs_runtime.Value
{
var __t152 bool
{
if (__local_var_5_4 != nil) {
__t152 = true
goto end_branch_152
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t152 = false
goto end_branch_152
} else {

}
}
{
__t152 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_152:
if __t152 {
var __t153 gopurs_runtime.Value
{
if (__local_var_5_4 != nil) {
__t153 = gopurs_runtime.Apply((__local_var_5_4).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil})
goto end_branch_153
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t153 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_153
} else {

}
}
{
__t153 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_153:
__t155 = gopurs_runtime.Apply(__t153, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal))
goto end_branch_155
} else {

}
}
{
var __t154 bool
{
if (__local_var_5_4 != nil) {
__t154 = false
goto end_branch_154
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t154 = true
goto end_branch_154
} else {

}
}
{
__t154 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_154:
if __t154 {
__t155 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_155
} else {

}
}
{
__t155 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_155:
__t159 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_148, __t155))
goto end_branch_159
} else {

}
}
{
var __t158 bool
{
var __t156 bool
{
if (__local_var_5_4 != nil) {
__t156 = true
goto end_branch_156
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t156 = false
goto end_branch_156
} else {

}
}
{
__t156 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_156:
if __t156 {
__t158 = false
goto end_branch_158
} else {

}
}
{
var __t157 bool
{
if (__local_var_5_4 != nil) {
__t157 = false
goto end_branch_157
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t157 = true
goto end_branch_157
} else {

}
}
{
__t157 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_157:
if __t157 {
__t158 = true
goto end_branch_158
} else {

}
}
{
__t158 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_158:
if __t158 {
__t159 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_159
} else {

}
}
{
__t159 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_159:
__t160 = __t159
goto end_branch_160
} else {

}
}
{
// TAST (Let): __local_var_6_137 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_137 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_137
var __t147 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t140 bool
{
var __t138 bool
{
if (__local_var_5_4 != nil) {
__t138 = true
goto end_branch_138
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t138 = false
goto end_branch_138
} else {

}
}
{
__t138 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_138:
if __t138 {
__t140 = false
goto end_branch_140
} else {

}
}
{
var __t139 bool
{
if (__local_var_5_4 != nil) {
__t139 = false
goto end_branch_139
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t139 = true
goto end_branch_139
} else {

}
}
{
__t139 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_139:
if __t139 {
__t140 = false
goto end_branch_140
} else {

}
}
{
__t140 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_140:
if __t140 {
var __t143 gopurs_runtime.Value
{
var __t141 bool
{
if (__local_var_5_4 != nil) {
__t141 = true
goto end_branch_141
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t141 = false
goto end_branch_141
} else {

}
}
{
__t141 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_141:
if __t141 {
__t143 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_143
} else {

}
}
{
var __t142 bool
{
if (__local_var_5_4 != nil) {
__t142 = false
goto end_branch_142
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t142 = true
goto end_branch_142
} else {

}
}
{
__t142 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_142:
if __t142 {
__t143 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_143
} else {

}
}
{
__t143 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_143:
__t147 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_137, __t143))
goto end_branch_147
} else {

}
}
{
var __t146 bool
{
var __t144 bool
{
if (__local_var_5_4 != nil) {
__t144 = true
goto end_branch_144
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t144 = false
goto end_branch_144
} else {

}
}
{
__t144 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_144:
if __t144 {
__t146 = true
goto end_branch_146
} else {

}
}
{
var __t145 bool
{
if (__local_var_5_4 != nil) {
__t145 = false
goto end_branch_145
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t145 = true
goto end_branch_145
} else {

}
}
{
__t145 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_145:
if __t145 {
__t146 = true
goto end_branch_146
} else {

}
}
{
__t146 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_146:
if __t146 {
__t147 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_147
} else {

}
}
{
__t147 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_147:
__t160 = __t147
}
end_branch_160:
__t317 = __t160
goto end_branch_317
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(7)) {
var __t184 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
// TAST (Let): __local_var_6_172 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_172 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_172
var __t183 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t175 bool
{
var __t173 bool
{
if (__local_var_5_4 != nil) {
__t173 = true
goto end_branch_173
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t173 = false
goto end_branch_173
} else {

}
}
{
__t173 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_173:
if __t173 {
__t175 = true
goto end_branch_175
} else {

}
}
{
var __t174 bool
{
if (__local_var_5_4 != nil) {
__t174 = false
goto end_branch_174
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t174 = true
goto end_branch_174
} else {

}
}
{
__t174 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_174:
if __t174 {
__t175 = false
goto end_branch_175
} else {

}
}
{
__t175 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_175:
if __t175 {
var __t179 gopurs_runtime.Value
{
var __t176 bool
{
if (__local_var_5_4 != nil) {
__t176 = true
goto end_branch_176
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t176 = false
goto end_branch_176
} else {

}
}
{
__t176 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_176:
if __t176 {
var __t177 gopurs_runtime.Value
{
if (__local_var_5_4 != nil) {
__t177 = gopurs_runtime.Apply((__local_var_5_4).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil})
goto end_branch_177
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t177 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_177
} else {

}
}
{
__t177 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_177:
__t179 = gopurs_runtime.Apply(__t177, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal))
goto end_branch_179
} else {

}
}
{
var __t178 bool
{
if (__local_var_5_4 != nil) {
__t178 = false
goto end_branch_178
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t178 = true
goto end_branch_178
} else {

}
}
{
__t178 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_178:
if __t178 {
__t179 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_179
} else {

}
}
{
__t179 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_179:
__t183 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_172, __t179))
goto end_branch_183
} else {

}
}
{
var __t182 bool
{
var __t180 bool
{
if (__local_var_5_4 != nil) {
__t180 = true
goto end_branch_180
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t180 = false
goto end_branch_180
} else {

}
}
{
__t180 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_180:
if __t180 {
__t182 = false
goto end_branch_182
} else {

}
}
{
var __t181 bool
{
if (__local_var_5_4 != nil) {
__t181 = false
goto end_branch_181
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t181 = true
goto end_branch_181
} else {

}
}
{
__t181 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_181:
if __t181 {
__t182 = true
goto end_branch_182
} else {

}
}
{
__t182 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_182:
if __t182 {
__t183 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_183
} else {

}
}
{
__t183 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_183:
__t184 = __t183
goto end_branch_184
} else {

}
}
{
// TAST (Let): __local_var_6_161 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_161 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_161
var __t171 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t164 bool
{
var __t162 bool
{
if (__local_var_5_4 != nil) {
__t162 = true
goto end_branch_162
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t162 = false
goto end_branch_162
} else {

}
}
{
__t162 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_162:
if __t162 {
__t164 = false
goto end_branch_164
} else {

}
}
{
var __t163 bool
{
if (__local_var_5_4 != nil) {
__t163 = false
goto end_branch_163
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t163 = true
goto end_branch_163
} else {

}
}
{
__t163 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_163:
if __t163 {
__t164 = false
goto end_branch_164
} else {

}
}
{
__t164 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_164:
if __t164 {
var __t167 gopurs_runtime.Value
{
var __t165 bool
{
if (__local_var_5_4 != nil) {
__t165 = true
goto end_branch_165
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t165 = false
goto end_branch_165
} else {

}
}
{
__t165 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_165:
if __t165 {
__t167 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_167
} else {

}
}
{
var __t166 bool
{
if (__local_var_5_4 != nil) {
__t166 = false
goto end_branch_166
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t166 = true
goto end_branch_166
} else {

}
}
{
__t166 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_166:
if __t166 {
__t167 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_167
} else {

}
}
{
__t167 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_167:
__t171 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_161, __t167))
goto end_branch_171
} else {

}
}
{
var __t170 bool
{
var __t168 bool
{
if (__local_var_5_4 != nil) {
__t168 = true
goto end_branch_168
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t168 = false
goto end_branch_168
} else {

}
}
{
__t168 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_168:
if __t168 {
__t170 = true
goto end_branch_170
} else {

}
}
{
var __t169 bool
{
if (__local_var_5_4 != nil) {
__t169 = false
goto end_branch_169
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t169 = true
goto end_branch_169
} else {

}
}
{
__t169 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_169:
if __t169 {
__t170 = true
goto end_branch_170
} else {

}
}
{
__t170 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_170:
if __t170 {
__t171 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_171
} else {

}
}
{
__t171 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_171:
__t184 = __t171
}
end_branch_184:
__t317 = __t184
goto end_branch_317
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(8)) {
var __t208 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
// TAST (Let): __local_var_6_196 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_196 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_196
var __t207 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t199 bool
{
var __t197 bool
{
if (__local_var_5_4 != nil) {
__t197 = true
goto end_branch_197
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t197 = false
goto end_branch_197
} else {

}
}
{
__t197 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_197:
if __t197 {
__t199 = true
goto end_branch_199
} else {

}
}
{
var __t198 bool
{
if (__local_var_5_4 != nil) {
__t198 = false
goto end_branch_198
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t198 = true
goto end_branch_198
} else {

}
}
{
__t198 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_198:
if __t198 {
__t199 = false
goto end_branch_199
} else {

}
}
{
__t199 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_199:
if __t199 {
var __t203 gopurs_runtime.Value
{
var __t200 bool
{
if (__local_var_5_4 != nil) {
__t200 = true
goto end_branch_200
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t200 = false
goto end_branch_200
} else {

}
}
{
__t200 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_200:
if __t200 {
var __t201 gopurs_runtime.Value
{
if (__local_var_5_4 != nil) {
__t201 = gopurs_runtime.Apply((__local_var_5_4).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil})
goto end_branch_201
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t201 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_201
} else {

}
}
{
__t201 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_201:
__t203 = gopurs_runtime.Apply(__t201, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal))
goto end_branch_203
} else {

}
}
{
var __t202 bool
{
if (__local_var_5_4 != nil) {
__t202 = false
goto end_branch_202
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t202 = true
goto end_branch_202
} else {

}
}
{
__t202 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_202:
if __t202 {
__t203 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_203
} else {

}
}
{
__t203 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_203:
__t207 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_196, __t203))
goto end_branch_207
} else {

}
}
{
var __t206 bool
{
var __t204 bool
{
if (__local_var_5_4 != nil) {
__t204 = true
goto end_branch_204
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t204 = false
goto end_branch_204
} else {

}
}
{
__t204 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_204:
if __t204 {
__t206 = false
goto end_branch_206
} else {

}
}
{
var __t205 bool
{
if (__local_var_5_4 != nil) {
__t205 = false
goto end_branch_205
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t205 = true
goto end_branch_205
} else {

}
}
{
__t205 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_205:
if __t205 {
__t206 = true
goto end_branch_206
} else {

}
}
{
__t206 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_206:
if __t206 {
__t207 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_207
} else {

}
}
{
__t207 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_207:
__t208 = __t207
goto end_branch_208
} else {

}
}
{
// TAST (Let): __local_var_6_185 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_185 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_185
var __t195 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t188 bool
{
var __t186 bool
{
if (__local_var_5_4 != nil) {
__t186 = true
goto end_branch_186
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t186 = false
goto end_branch_186
} else {

}
}
{
__t186 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_186:
if __t186 {
__t188 = false
goto end_branch_188
} else {

}
}
{
var __t187 bool
{
if (__local_var_5_4 != nil) {
__t187 = false
goto end_branch_187
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t187 = true
goto end_branch_187
} else {

}
}
{
__t187 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_187:
if __t187 {
__t188 = false
goto end_branch_188
} else {

}
}
{
__t188 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_188:
if __t188 {
var __t191 gopurs_runtime.Value
{
var __t189 bool
{
if (__local_var_5_4 != nil) {
__t189 = true
goto end_branch_189
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t189 = false
goto end_branch_189
} else {

}
}
{
__t189 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_189:
if __t189 {
__t191 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_191
} else {

}
}
{
var __t190 bool
{
if (__local_var_5_4 != nil) {
__t190 = false
goto end_branch_190
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t190 = true
goto end_branch_190
} else {

}
}
{
__t190 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_190:
if __t190 {
__t191 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_191
} else {

}
}
{
__t191 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_191:
__t195 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_185, __t191))
goto end_branch_195
} else {

}
}
{
var __t194 bool
{
var __t192 bool
{
if (__local_var_5_4 != nil) {
__t192 = true
goto end_branch_192
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t192 = false
goto end_branch_192
} else {

}
}
{
__t192 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_192:
if __t192 {
__t194 = true
goto end_branch_194
} else {

}
}
{
var __t193 bool
{
if (__local_var_5_4 != nil) {
__t193 = false
goto end_branch_193
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t193 = true
goto end_branch_193
} else {

}
}
{
__t193 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_193:
if __t193 {
__t194 = true
goto end_branch_194
} else {

}
}
{
__t194 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_194:
if __t194 {
__t195 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_195
} else {

}
}
{
__t195 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_195:
__t208 = __t195
}
end_branch_208:
__t317 = __t208
goto end_branch_317
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(9)) {
var __t232 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
// TAST (Let): __local_var_6_220 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_220 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_220
var __t231 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t223 bool
{
var __t221 bool
{
if (__local_var_5_4 != nil) {
__t221 = true
goto end_branch_221
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t221 = false
goto end_branch_221
} else {

}
}
{
__t221 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_221:
if __t221 {
__t223 = true
goto end_branch_223
} else {

}
}
{
var __t222 bool
{
if (__local_var_5_4 != nil) {
__t222 = false
goto end_branch_222
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t222 = true
goto end_branch_222
} else {

}
}
{
__t222 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_222:
if __t222 {
__t223 = false
goto end_branch_223
} else {

}
}
{
__t223 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_223:
if __t223 {
var __t227 gopurs_runtime.Value
{
var __t224 bool
{
if (__local_var_5_4 != nil) {
__t224 = true
goto end_branch_224
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t224 = false
goto end_branch_224
} else {

}
}
{
__t224 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_224:
if __t224 {
var __t225 gopurs_runtime.Value
{
if (__local_var_5_4 != nil) {
__t225 = gopurs_runtime.Apply((__local_var_5_4).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil})
goto end_branch_225
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t225 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_225
} else {

}
}
{
__t225 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_225:
__t227 = gopurs_runtime.Apply(__t225, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal))
goto end_branch_227
} else {

}
}
{
var __t226 bool
{
if (__local_var_5_4 != nil) {
__t226 = false
goto end_branch_226
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t226 = true
goto end_branch_226
} else {

}
}
{
__t226 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_226:
if __t226 {
__t227 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_227
} else {

}
}
{
__t227 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_227:
__t231 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_220, __t227))
goto end_branch_231
} else {

}
}
{
var __t230 bool
{
var __t228 bool
{
if (__local_var_5_4 != nil) {
__t228 = true
goto end_branch_228
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t228 = false
goto end_branch_228
} else {

}
}
{
__t228 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_228:
if __t228 {
__t230 = false
goto end_branch_230
} else {

}
}
{
var __t229 bool
{
if (__local_var_5_4 != nil) {
__t229 = false
goto end_branch_229
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t229 = true
goto end_branch_229
} else {

}
}
{
__t229 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_229:
if __t229 {
__t230 = true
goto end_branch_230
} else {

}
}
{
__t230 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_230:
if __t230 {
__t231 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_231
} else {

}
}
{
__t231 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_231:
__t232 = __t231
goto end_branch_232
} else {

}
}
{
// TAST (Let): __local_var_6_209 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_209 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_209
var __t219 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t212 bool
{
var __t210 bool
{
if (__local_var_5_4 != nil) {
__t210 = true
goto end_branch_210
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t210 = false
goto end_branch_210
} else {

}
}
{
__t210 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_210:
if __t210 {
__t212 = false
goto end_branch_212
} else {

}
}
{
var __t211 bool
{
if (__local_var_5_4 != nil) {
__t211 = false
goto end_branch_211
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t211 = true
goto end_branch_211
} else {

}
}
{
__t211 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_211:
if __t211 {
__t212 = false
goto end_branch_212
} else {

}
}
{
__t212 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_212:
if __t212 {
var __t215 gopurs_runtime.Value
{
var __t213 bool
{
if (__local_var_5_4 != nil) {
__t213 = true
goto end_branch_213
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t213 = false
goto end_branch_213
} else {

}
}
{
__t213 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_213:
if __t213 {
__t215 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_215
} else {

}
}
{
var __t214 bool
{
if (__local_var_5_4 != nil) {
__t214 = false
goto end_branch_214
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t214 = true
goto end_branch_214
} else {

}
}
{
__t214 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_214:
if __t214 {
__t215 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_215
} else {

}
}
{
__t215 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_215:
__t219 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_209, __t215))
goto end_branch_219
} else {

}
}
{
var __t218 bool
{
var __t216 bool
{
if (__local_var_5_4 != nil) {
__t216 = true
goto end_branch_216
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t216 = false
goto end_branch_216
} else {

}
}
{
__t216 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_216:
if __t216 {
__t218 = true
goto end_branch_218
} else {

}
}
{
var __t217 bool
{
if (__local_var_5_4 != nil) {
__t217 = false
goto end_branch_217
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t217 = true
goto end_branch_217
} else {

}
}
{
__t217 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_217:
if __t217 {
__t218 = true
goto end_branch_218
} else {

}
}
{
__t218 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_218:
if __t218 {
__t219 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_219
} else {

}
}
{
__t219 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_219:
__t232 = __t219
}
end_branch_232:
__t317 = __t232
goto end_branch_317
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(10)) {
var __t256 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
// TAST (Let): __local_var_6_244 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_244 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_244
var __t255 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t247 bool
{
var __t245 bool
{
if (__local_var_5_4 != nil) {
__t245 = true
goto end_branch_245
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t245 = false
goto end_branch_245
} else {

}
}
{
__t245 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_245:
if __t245 {
__t247 = true
goto end_branch_247
} else {

}
}
{
var __t246 bool
{
if (__local_var_5_4 != nil) {
__t246 = false
goto end_branch_246
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t246 = true
goto end_branch_246
} else {

}
}
{
__t246 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_246:
if __t246 {
__t247 = false
goto end_branch_247
} else {

}
}
{
__t247 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_247:
if __t247 {
var __t251 gopurs_runtime.Value
{
var __t248 bool
{
if (__local_var_5_4 != nil) {
__t248 = true
goto end_branch_248
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t248 = false
goto end_branch_248
} else {

}
}
{
__t248 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_248:
if __t248 {
var __t249 gopurs_runtime.Value
{
if (__local_var_5_4 != nil) {
__t249 = gopurs_runtime.Apply((__local_var_5_4).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil})
goto end_branch_249
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t249 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_249
} else {

}
}
{
__t249 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_249:
__t251 = gopurs_runtime.Apply(__t249, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal))
goto end_branch_251
} else {

}
}
{
var __t250 bool
{
if (__local_var_5_4 != nil) {
__t250 = false
goto end_branch_250
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t250 = true
goto end_branch_250
} else {

}
}
{
__t250 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_250:
if __t250 {
__t251 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_251
} else {

}
}
{
__t251 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_251:
__t255 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_244, __t251))
goto end_branch_255
} else {

}
}
{
var __t254 bool
{
var __t252 bool
{
if (__local_var_5_4 != nil) {
__t252 = true
goto end_branch_252
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t252 = false
goto end_branch_252
} else {

}
}
{
__t252 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_252:
if __t252 {
__t254 = false
goto end_branch_254
} else {

}
}
{
var __t253 bool
{
if (__local_var_5_4 != nil) {
__t253 = false
goto end_branch_253
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t253 = true
goto end_branch_253
} else {

}
}
{
__t253 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_253:
if __t253 {
__t254 = true
goto end_branch_254
} else {

}
}
{
__t254 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_254:
if __t254 {
__t255 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_255
} else {

}
}
{
__t255 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_255:
__t256 = __t255
goto end_branch_256
} else {

}
}
{
// TAST (Let): __local_var_6_233 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_233 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_233
var __t243 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t236 bool
{
var __t234 bool
{
if (__local_var_5_4 != nil) {
__t234 = true
goto end_branch_234
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t234 = false
goto end_branch_234
} else {

}
}
{
__t234 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_234:
if __t234 {
__t236 = false
goto end_branch_236
} else {

}
}
{
var __t235 bool
{
if (__local_var_5_4 != nil) {
__t235 = false
goto end_branch_235
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t235 = true
goto end_branch_235
} else {

}
}
{
__t235 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_235:
if __t235 {
__t236 = false
goto end_branch_236
} else {

}
}
{
__t236 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_236:
if __t236 {
var __t239 gopurs_runtime.Value
{
var __t237 bool
{
if (__local_var_5_4 != nil) {
__t237 = true
goto end_branch_237
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t237 = false
goto end_branch_237
} else {

}
}
{
__t237 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_237:
if __t237 {
__t239 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_239
} else {

}
}
{
var __t238 bool
{
if (__local_var_5_4 != nil) {
__t238 = false
goto end_branch_238
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t238 = true
goto end_branch_238
} else {

}
}
{
__t238 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_238:
if __t238 {
__t239 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_239
} else {

}
}
{
__t239 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_239:
__t243 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_233, __t239))
goto end_branch_243
} else {

}
}
{
var __t242 bool
{
var __t240 bool
{
if (__local_var_5_4 != nil) {
__t240 = true
goto end_branch_240
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t240 = false
goto end_branch_240
} else {

}
}
{
__t240 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_240:
if __t240 {
__t242 = true
goto end_branch_242
} else {

}
}
{
var __t241 bool
{
if (__local_var_5_4 != nil) {
__t241 = false
goto end_branch_241
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t241 = true
goto end_branch_241
} else {

}
}
{
__t241 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_241:
if __t241 {
__t242 = true
goto end_branch_242
} else {

}
}
{
__t242 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_242:
if __t242 {
__t243 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_243
} else {

}
}
{
__t243 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_243:
__t256 = __t243
}
end_branch_256:
__t317 = __t256
goto end_branch_317
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(11)) {
var __t280 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
// TAST (Let): __local_var_6_268 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_268 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_268
var __t279 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t271 bool
{
var __t269 bool
{
if (__local_var_5_4 != nil) {
__t269 = true
goto end_branch_269
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t269 = false
goto end_branch_269
} else {

}
}
{
__t269 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_269:
if __t269 {
__t271 = true
goto end_branch_271
} else {

}
}
{
var __t270 bool
{
if (__local_var_5_4 != nil) {
__t270 = false
goto end_branch_270
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t270 = true
goto end_branch_270
} else {

}
}
{
__t270 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_270:
if __t270 {
__t271 = false
goto end_branch_271
} else {

}
}
{
__t271 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_271:
if __t271 {
var __t275 gopurs_runtime.Value
{
var __t272 bool
{
if (__local_var_5_4 != nil) {
__t272 = true
goto end_branch_272
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t272 = false
goto end_branch_272
} else {

}
}
{
__t272 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_272:
if __t272 {
var __t273 gopurs_runtime.Value
{
if (__local_var_5_4 != nil) {
__t273 = gopurs_runtime.Apply((__local_var_5_4).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil})
goto end_branch_273
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t273 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_273
} else {

}
}
{
__t273 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_273:
__t275 = gopurs_runtime.Apply(__t273, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal))
goto end_branch_275
} else {

}
}
{
var __t274 bool
{
if (__local_var_5_4 != nil) {
__t274 = false
goto end_branch_274
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t274 = true
goto end_branch_274
} else {

}
}
{
__t274 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_274:
if __t274 {
__t275 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_275
} else {

}
}
{
__t275 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_275:
__t279 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_268, __t275))
goto end_branch_279
} else {

}
}
{
var __t278 bool
{
var __t276 bool
{
if (__local_var_5_4 != nil) {
__t276 = true
goto end_branch_276
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t276 = false
goto end_branch_276
} else {

}
}
{
__t276 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_276:
if __t276 {
__t278 = false
goto end_branch_278
} else {

}
}
{
var __t277 bool
{
if (__local_var_5_4 != nil) {
__t277 = false
goto end_branch_277
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t277 = true
goto end_branch_277
} else {

}
}
{
__t277 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_277:
if __t277 {
__t278 = true
goto end_branch_278
} else {

}
}
{
__t278 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_278:
if __t278 {
__t279 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_279
} else {

}
}
{
__t279 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_279:
__t280 = __t279
goto end_branch_280
} else {

}
}
{
// TAST (Let): __local_var_6_257 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_257 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_257
var __t267 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t260 bool
{
var __t258 bool
{
if (__local_var_5_4 != nil) {
__t258 = true
goto end_branch_258
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t258 = false
goto end_branch_258
} else {

}
}
{
__t258 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_258:
if __t258 {
__t260 = false
goto end_branch_260
} else {

}
}
{
var __t259 bool
{
if (__local_var_5_4 != nil) {
__t259 = false
goto end_branch_259
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t259 = true
goto end_branch_259
} else {

}
}
{
__t259 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_259:
if __t259 {
__t260 = false
goto end_branch_260
} else {

}
}
{
__t260 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_260:
if __t260 {
var __t263 gopurs_runtime.Value
{
var __t261 bool
{
if (__local_var_5_4 != nil) {
__t261 = true
goto end_branch_261
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t261 = false
goto end_branch_261
} else {

}
}
{
__t261 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_261:
if __t261 {
__t263 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_263
} else {

}
}
{
var __t262 bool
{
if (__local_var_5_4 != nil) {
__t262 = false
goto end_branch_262
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t262 = true
goto end_branch_262
} else {

}
}
{
__t262 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_262:
if __t262 {
__t263 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_263
} else {

}
}
{
__t263 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_263:
__t267 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_257, __t263))
goto end_branch_267
} else {

}
}
{
var __t266 bool
{
var __t264 bool
{
if (__local_var_5_4 != nil) {
__t264 = true
goto end_branch_264
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t264 = false
goto end_branch_264
} else {

}
}
{
__t264 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_264:
if __t264 {
__t266 = true
goto end_branch_266
} else {

}
}
{
var __t265 bool
{
if (__local_var_5_4 != nil) {
__t265 = false
goto end_branch_265
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t265 = true
goto end_branch_265
} else {

}
}
{
__t265 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_265:
if __t265 {
__t266 = true
goto end_branch_266
} else {

}
}
{
__t266 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_266:
if __t266 {
__t267 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_267
} else {

}
}
{
__t267 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_267:
__t280 = __t267
}
end_branch_280:
__t317 = __t280
goto end_branch_317
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(12)) {
var __t304 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
// TAST (Let): __local_var_6_292 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_292 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_292
var __t303 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t295 bool
{
var __t293 bool
{
if (__local_var_5_4 != nil) {
__t293 = true
goto end_branch_293
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t293 = false
goto end_branch_293
} else {

}
}
{
__t293 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_293:
if __t293 {
__t295 = true
goto end_branch_295
} else {

}
}
{
var __t294 bool
{
if (__local_var_5_4 != nil) {
__t294 = false
goto end_branch_294
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t294 = true
goto end_branch_294
} else {

}
}
{
__t294 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_294:
if __t294 {
__t295 = false
goto end_branch_295
} else {

}
}
{
__t295 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_295:
if __t295 {
var __t299 gopurs_runtime.Value
{
var __t296 bool
{
if (__local_var_5_4 != nil) {
__t296 = true
goto end_branch_296
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t296 = false
goto end_branch_296
} else {

}
}
{
__t296 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_296:
if __t296 {
var __t297 gopurs_runtime.Value
{
if (__local_var_5_4 != nil) {
__t297 = gopurs_runtime.Apply((__local_var_5_4).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil})
goto end_branch_297
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t297 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_297
} else {

}
}
{
__t297 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_297:
__t299 = gopurs_runtime.Apply(__t297, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal))
goto end_branch_299
} else {

}
}
{
var __t298 bool
{
if (__local_var_5_4 != nil) {
__t298 = false
goto end_branch_298
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t298 = true
goto end_branch_298
} else {

}
}
{
__t298 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_298:
if __t298 {
__t299 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_299
} else {

}
}
{
__t299 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_299:
__t303 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_292, __t299))
goto end_branch_303
} else {

}
}
{
var __t302 bool
{
var __t300 bool
{
if (__local_var_5_4 != nil) {
__t300 = true
goto end_branch_300
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t300 = false
goto end_branch_300
} else {

}
}
{
__t300 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_300:
if __t300 {
__t302 = false
goto end_branch_302
} else {

}
}
{
var __t301 bool
{
if (__local_var_5_4 != nil) {
__t301 = false
goto end_branch_301
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t301 = true
goto end_branch_301
} else {

}
}
{
__t301 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_301:
if __t301 {
__t302 = true
goto end_branch_302
} else {

}
}
{
__t302 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_302:
if __t302 {
__t303 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_303
} else {

}
}
{
__t303 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_303:
__t304 = __t303
goto end_branch_304
} else {

}
}
{
// TAST (Let): __local_var_6_281 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_281 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_281
var __t291 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t284 bool
{
var __t282 bool
{
if (__local_var_5_4 != nil) {
__t282 = true
goto end_branch_282
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t282 = false
goto end_branch_282
} else {

}
}
{
__t282 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_282:
if __t282 {
__t284 = false
goto end_branch_284
} else {

}
}
{
var __t283 bool
{
if (__local_var_5_4 != nil) {
__t283 = false
goto end_branch_283
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t283 = true
goto end_branch_283
} else {

}
}
{
__t283 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_283:
if __t283 {
__t284 = false
goto end_branch_284
} else {

}
}
{
__t284 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_284:
if __t284 {
var __t287 gopurs_runtime.Value
{
var __t285 bool
{
if (__local_var_5_4 != nil) {
__t285 = true
goto end_branch_285
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t285 = false
goto end_branch_285
} else {

}
}
{
__t285 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_285:
if __t285 {
__t287 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_287
} else {

}
}
{
var __t286 bool
{
if (__local_var_5_4 != nil) {
__t286 = false
goto end_branch_286
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t286 = true
goto end_branch_286
} else {

}
}
{
__t286 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_286:
if __t286 {
__t287 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_287
} else {

}
}
{
__t287 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_287:
__t291 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_281, __t287))
goto end_branch_291
} else {

}
}
{
var __t290 bool
{
var __t288 bool
{
if (__local_var_5_4 != nil) {
__t288 = true
goto end_branch_288
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t288 = false
goto end_branch_288
} else {

}
}
{
__t288 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_288:
if __t288 {
__t290 = true
goto end_branch_290
} else {

}
}
{
var __t289 bool
{
if (__local_var_5_4 != nil) {
__t289 = false
goto end_branch_289
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t289 = true
goto end_branch_289
} else {

}
}
{
__t289 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_289:
if __t289 {
__t290 = true
goto end_branch_290
} else {

}
}
{
__t290 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_290:
if __t290 {
__t291 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_291
} else {

}
}
{
__t291 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_291:
__t304 = __t291
}
end_branch_304:
__t317 = __t304
goto end_branch_317
} else {

}
}
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
// TAST (Let): __local_var_6_305 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_305 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_305
var __t316 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t308 bool
{
var __t306 bool
{
if (__local_var_5_4 != nil) {
__t306 = false
goto end_branch_306
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t306 = false
goto end_branch_306
} else {

}
}
{
__t306 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_306:
if __t306 {
__t308 = true
goto end_branch_308
} else {

}
}
{
var __t307 bool
{
if (__local_var_5_4 != nil) {
__t307 = true
goto end_branch_307
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t307 = true
goto end_branch_307
} else {

}
}
{
__t307 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_307:
if __t307 {
__t308 = false
goto end_branch_308
} else {

}
}
{
__t308 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_308:
if __t308 {
var __t312 gopurs_runtime.Value
{
var __t309 bool
{
if (__local_var_5_4 != nil) {
__t309 = false
goto end_branch_309
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t309 = false
goto end_branch_309
} else {

}
}
{
__t309 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_309:
if __t309 {
var __t310 gopurs_runtime.Value
{
if (__local_var_5_4 != nil) {
__t310 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_310
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t310 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_310
} else {

}
}
{
__t310 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_310:
__t312 = gopurs_runtime.Apply(__t310, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal))
goto end_branch_312
} else {

}
}
{
var __t311 bool
{
if (__local_var_5_4 != nil) {
__t311 = true
goto end_branch_311
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t311 = true
goto end_branch_311
} else {

}
}
{
__t311 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_311:
if __t311 {
__t312 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_312
} else {

}
}
{
__t312 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_312:
__t316 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_305, __t312))
goto end_branch_316
} else {

}
}
{
var __t315 bool
{
var __t313 bool
{
if (__local_var_5_4 != nil) {
__t313 = false
goto end_branch_313
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t313 = false
goto end_branch_313
} else {

}
}
{
__t313 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_313:
if __t313 {
__t315 = false
goto end_branch_315
} else {

}
}
{
var __t314 bool
{
if (__local_var_5_4 != nil) {
__t314 = true
goto end_branch_314
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t314 = true
goto end_branch_314
} else {

}
}
{
__t314 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_314:
if __t314 {
__t315 = true
goto end_branch_315
} else {

}
}
{
__t315 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_315:
if __t315 {
__t316 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_316
} else {

}
}
{
__t316 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_316:
__t317 = __t316
goto end_branch_317
} else {

}
}
{
// TAST (Let): __local_var_6_6 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)])] (TypeApp (TypeVar m$scope68) [(TypeVar a$scope67)]))
__local_var_6_6 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_6
var __t16 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t9 bool
{
var __t7 bool
{
if (__local_var_5_4 != nil) {
__t7 = false
goto end_branch_7
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t7 = false
goto end_branch_7
} else {

}
}
{
__t7 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_7:
if __t7 {
__t9 = false
goto end_branch_9
} else {

}
}
{
var __t8 bool
{
if (__local_var_5_4 != nil) {
__t8 = true
goto end_branch_8
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t8 = true
goto end_branch_8
} else {

}
}
{
__t8 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_8:
if __t8 {
__t9 = false
goto end_branch_9
} else {

}
}
{
__t9 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_9:
if __t9 {
var __t12 gopurs_runtime.Value
{
var __t10 bool
{
if (__local_var_5_4 != nil) {
__t10 = false
goto end_branch_10
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t10 = false
goto end_branch_10
} else {

}
}
{
__t10 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_10:
if __t10 {
__t12 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_12
} else {

}
}
{
var __t11 bool
{
if (__local_var_5_4 != nil) {
__t11 = true
goto end_branch_11
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t11 = true
goto end_branch_11
} else {

}
}
{
__t11 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_11:
if __t11 {
__t12 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_6_6, __t12))
goto end_branch_16
} else {

}
}
{
var __t15 bool
{
var __t13 bool
{
if (__local_var_5_4 != nil) {
__t13 = false
goto end_branch_13
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t13 = false
goto end_branch_13
} else {

}
}
{
__t13 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_13:
if __t13 {
__t15 = true
goto end_branch_15
} else {

}
}
{
var __t14 bool
{
if (__local_var_5_4 != nil) {
__t14 = true
goto end_branch_14
} else {

}
}
{
if (__local_var_5_4 == nil) {
__t14 = true
goto end_branch_14
} else {

}
}
{
__t14 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_14:
if __t14 {
__t15 = true
goto end_branch_15
} else {

}
}
{
__t15 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_15:
if __t15 {
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_16
} else {

}
}
{
__t16 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_16:
__t317 = __t16
}
end_branch_317:
// TAST (Let): __local_var_4_1 shape=Let(Let(Branch(Branch(Let(Branch(App(Other), Other, def=Other)), def=Let(Branch(App(Other), Other, def=Other))), Branch(Let(Branch(App(Other), Other, def=Other)), def=Let(Branch(App(Other), Other, def=Other))), Branch(Let(Branch(App(Other), Other, def=Other)), def=Let(Branch(App(Other), Other, def=Other))), Branch(Let(Branch(App(Other), Other, def=Other)), def=Let(Branch(App(Other), Other, def=Other))), Branch(Let(Branch(App(Other), Other, def=Other)), def=Let(Branch(App(Other), Other, def=Other))), Branch(Let(Branch(App(Other), Other, def=Other)), def=Let(Branch(App(Other), Other, def=Other))), Branch(Let(Branch(App(Other), Other, def=Other)), def=Let(Branch(App(Other), Other, def=Other))), Branch(Let(Branch(App(Other), Other, def=Other)), def=Let(Branch(App(Other), Other, def=Other))), Branch(Let(Branch(App(Other), Other, def=Other)), def=Let(Branch(App(Other), Other, def=Other))), Branch(Let(Branch(App(Other), Other, def=Other)), def=Let(Branch(App(Other), Other, def=Other))), Branch(Let(Branch(App(Other), Other, def=Other)), def=Let(Branch(App(Other), Other, def=Other))), Branch(Let(Branch(App(Other), Other, def=Other)), def=Let(Branch(App(Other), Other, def=Other))), Let(Branch(App(Other), Other, def=Other)), def=Let(Branch(App(Other), Other, def=Other))))) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Date"] [])])
var __local_var_4_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = Rebox_Data_DateTime_3094389156_2280409795(__t317)
var __t319 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_4_1 != nil) {
__t319 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(Get_Data_DateTime_DateTime(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((__local_var_4_1).V0)}), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_319
} else {

}
}
{
__t319 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_319:
// TAST (Let): __local_var_5_318 shape=Branch(Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope70)])
__local_var_5_318 := __t319
_ = __local_var_5_318
var __t321 *Constructor_Data_Maybe_Just[int64]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "hour").IntVal) >= (int64(0))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "hour").IntVal) <= (int64(23))) {
__t321 = Rebox_Data_DateTime_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "hour").IntVal), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_321
} else {

}
}
{
__t321 = Rebox_Data_DateTime_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_321:
// TAST (Let): __local_var_6_320 shape=Branch(Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_6_320 := __t321
_ = __local_var_6_320
var __t337 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_6_320 != nil) {
var __t336 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "minute").IntVal) >= (int64(0))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "minute").IntVal) <= (int64(59))) {
var __t332 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "second").IntVal) >= (int64(0))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "second").IntVal) <= (int64(59))) {
var __t331 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_318 != nil) {
var __t330 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) >= (int64(0))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) <= (int64(999))) {
var __t329 gopurs_runtime.Value
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) >= (int64(0))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) <= (int64(999))) {
__t329 = gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((&Constructor_Data_Time_Time{1, (__local_var_6_320).V0, gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "minute").IntVal, gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "second").IntVal, gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal}))}
goto end_branch_329
} else {

}
}
{
__t329 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_329:
__t330 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_318).V0, __t329), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_330
} else {

}
}
{
__t330 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_330:
__t331 = __t330
goto end_branch_331
} else {

}
}
{
if (__local_var_5_318 == nil) {
__t331 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_331
} else {

}
}
{
__t331 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_331:
__t332 = __t331
goto end_branch_332
} else {

}
}
{
var __t328 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_318 != nil) {
var __t327 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if false {
var __t326 gopurs_runtime.Value
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) >= (int64(0))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) <= (int64(999))) {
__t326 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_326
} else {

}
}
{
__t326 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_326:
__t327 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_318).V0, __t326), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_327
} else {

}
}
{
__t327 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_327:
__t328 = __t327
goto end_branch_328
} else {

}
}
{
if (__local_var_5_318 == nil) {
__t328 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_328
} else {

}
}
{
__t328 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_328:
__t332 = __t328
}
end_branch_332:
__t336 = __t332
goto end_branch_336
} else {

}
}
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "second").IntVal) >= (int64(0))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "second").IntVal) <= (int64(59))) {
var __t335 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_318 != nil) {
var __t334 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if false {
var __t333 gopurs_runtime.Value
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) >= (int64(0))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) <= (int64(999))) {
__t333 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_333
} else {

}
}
{
__t333 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_333:
__t334 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_318).V0, __t333), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_334
} else {

}
}
{
__t334 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_334:
__t335 = __t334
goto end_branch_335
} else {

}
}
{
if (__local_var_5_318 == nil) {
__t335 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_335
} else {

}
}
{
__t335 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_335:
__t336 = __t335
goto end_branch_336
} else {

}
}
{
var __t325 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_318 != nil) {
var __t324 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if false {
var __t323 gopurs_runtime.Value
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) >= (int64(0))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) <= (int64(999))) {
__t323 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
goto end_branch_323
} else {

}
}
{
__t323 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_323:
__t324 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_318).V0, __t323), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_324
} else {

}
}
{
__t324 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_324:
__t325 = __t324
goto end_branch_325
} else {

}
}
{
if (__local_var_5_318 == nil) {
__t325 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_325
} else {

}
}
{
__t325 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_325:
__t336 = __t325
}
end_branch_336:
__t337 = __t336
goto end_branch_337
} else {

}
}
{
var __t322 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_318 != nil) {
__t322 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_322
} else {

}
}
{
if (__local_var_5_318 == nil) {
__t322 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_322
} else {

}
}
{
__t322 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_322:
__t337 = __t322
}
end_branch_337:
__t338 = __t337
goto end_branch_338
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr == nil) {
__t338 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_338
} else {

}
}
{
__t338 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_338:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t338)}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Rebox_Data_DateTime_1386611502_1726078761(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[*Constructor_Data_Date_Date]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_DateTime_1888627657_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_DateTime_DateTime]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_DateTime_2094947566_3523628265(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[*Constructor_Data_Date_Date]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V1)
		out.V2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V2)
	return out
}

func Rebox_Data_DateTime_2094947566_4136451977(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[*Constructor_Data_Time_Time] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[*Constructor_Data_Time_Time]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](in.V1)
		out.V2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](in.V2)
	return out
}

func Rebox_Data_DateTime_2380273161_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_DateTime_DateTime]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_DateTime_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_DateTime_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}

func Rebox_Data_DateTime_3247149001_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_DateTime_DateTime]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(in.V1)}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_DateTime_3523628265_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_Date_Date]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V1)}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_DateTime_3790796878_1957390985(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Date_Date]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_DateTime_3790796878_503123561(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_DateTime_DateTime] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_DateTime_DateTime]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_DateTime_4136451977_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_Time_Time]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(in.V1)}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_DateTime_4177771502_1888627657(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_DateTime_DateTime] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_DateTime_DateTime]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_DateTime_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_DateTime_4177771502_758368489(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_DateTime_503123561_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_DateTime_DateTime]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Get_Data_DateTime_adjustImpl() gopurs_runtime.Value {
	return _Gopurs_Data_DateTime_AdjustImpl
}

func Get_Data_DateTime_calcDiff() gopurs_runtime.Value {
	return _Gopurs_Data_DateTime_CalcDiff
}

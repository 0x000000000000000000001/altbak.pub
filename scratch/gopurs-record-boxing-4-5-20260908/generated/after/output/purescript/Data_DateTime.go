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
return gopurs_runtime.Str((((((((((("(DateTime ") + (gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_Date_Date]](Get_Data_Date_showDate()).V0), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V0)}).StrVal())) + (" (Time ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V1).V0)).StrVal())) + (" ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V1).V1)).StrVal())) + (" ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V1).V2)).StrVal())) + (" ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(v_0.UnsafePtr).V1).V3)).StrVal())) + ("))"))
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
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[*Constructor_Data_Date_Date]](Get_Data_Date_eqDate()).V0), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V0)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V0)}).IntVal) != (0)) && (((((((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V0) == (((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V0)) && ((((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V1) == (((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V1))) && ((((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V2) == (((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V2))) && ((((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V3) == (((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V3))))
})})))}
	})
	return cache_Data_DateTime_eqDateTime
}

var cache_Data_DateTime_ordDateTime gopurs_runtime.Value
var once_Data_DateTime_ordDateTime sync.Once
func Get_Data_DateTime_ordDateTime() gopurs_runtime.Value {
	once_Data_DateTime_ordDateTime.Do(func() {
		cache_Data_DateTime_ordDateTime = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_1888627657_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_DateTime_DateTime]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_503123561_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[*Constructor_Data_DateTime_DateTime]](Get_Data_DateTime_eqDateTime())))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_2_0 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date]](Get_Data_Date_ordDate()).V1), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V0)}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V0)}).IntVal)
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
// TAST (Let): v_3_1 shape=App(Var) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_3_1 := uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V0), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V0)).IntVal)
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
// TAST (Let): v1_4_2 shape=App(Var) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_4_2 := uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V1), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V1)).IntVal)
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
// TAST (Let): v2_5_3 shape=App(Var) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v2_5_3 := uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V2), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V2)).IntVal)
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
__t4 = uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(x_0.UnsafePtr).V1).V3), gopurs_runtime.Int(((*Constructor_Data_DateTime_DateTime)(y_1.UnsafePtr).V1).V3)).IntVal)
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
		cache_Data_DateTime_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, dt1_1_box gopurs_runtime.Value, dt2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DateTime_diff(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dictDuration_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](dt1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_DateTime_DateTime](dt2_2_box))
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
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_1888627657_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_DateTime_DateTime]](Get_Data_DateTime_ordDateTime())))}
}), (&Constructor_Data_DateTime_DateTime{1, (&Constructor_Data_Date_Date{1, int64(-271820), 1908470532, int64(1)}), (&Constructor_Data_Time_Time{1, int64(0), int64(0), int64(0), int64(0)})}), (&Constructor_Data_DateTime_DateTime{1, (&Constructor_Data_Date_Date{1, int64(275759), 3889233761, int64(31)}), (&Constructor_Data_Time_Time{1, int64(23), int64(59), int64(59), int64(999)})})})))}
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
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
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
if (uint32(__t_tag_0) == 1908470532) {
__t12 = int64(1)
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_1) == 2455627378) {
__t12 = int64(2)
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_2) == 4162469099) {
__t12 = int64(3)
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_3) == 1692989816) {
__t12 = int64(4)
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_4) == 330658827) {
__t12 = int64(5)
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_5) == 4067355978) {
__t12 = int64(6)
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_6) == 2276710548) {
__t12 = int64(7)
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_7) == 243771071) {
__t12 = int64(8)
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_8) == 215731793) {
__t12 = int64(9)
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_9) == 8639228) {
__t12 = int64(10)
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = ((v_0).V0).V1
if (uint32(__t_tag_10) == 49471444) {
__t12 = int64(11)
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = ((v_0).V0).V1
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Apply(Get_Data_DateTime_DateTime(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((v_2).V0)}), gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((v_2).V1)}))
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer((&Constructor_Data_DateTime_DateTime{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](a_3), (v_2).V1}))}
}), gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((v_2).V0)}))
}

func Call_Data_DateTime_modifyDate(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_DateTime_DateTime) *Constructor_Data_DateTime_DateTime {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_DateTime_DateTime = v_1_loop
_ = v_1
return (&Constructor_Data_DateTime_DateTime{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((v_1).V0)})), (v_1).V1})
}

func Call_Data_DateTime_diff(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value], dt1_1_loop *Constructor_Data_DateTime_DateTime, dt2_2_loop *Constructor_Data_DateTime_DateTime) gopurs_runtime.Value {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var dt1_1 *Constructor_Data_DateTime_DateTime = dt1_1_loop
_ = dt1_1
var dt2_2 *Constructor_Data_DateTime_DateTime = dt2_2_loop
_ = dt2_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration_0.V1), gopurs_runtime.Float(gopurs_runtime.UncurriedApp2(Get_Data_DateTime_calcDiff(), func() gopurs_runtime.Value {
				orig := Call_Data_DateTime_toRecord(dt1_1)
				_ = orig
				return gopurs_runtime.RecordDict([]string{"day", "hour", "millisecond", "minute", "month", "second", "year"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.day), gopurs_runtime.Int(orig.hour), gopurs_runtime.Int(orig.millisecond), gopurs_runtime.Int(orig.minute), gopurs_runtime.Int(orig.month), gopurs_runtime.Int(orig.second), gopurs_runtime.Int(orig.year)})
				}(), func() gopurs_runtime.Value {
				orig := Call_Data_DateTime_toRecord(dt2_2)
				_ = orig
				return gopurs_runtime.RecordDict([]string{"day", "hour", "millisecond", "minute", "month", "second", "year"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.day), gopurs_runtime.Int(orig.hour), gopurs_runtime.Int(orig.millisecond), gopurs_runtime.Int(orig.minute), gopurs_runtime.Int(orig.month), gopurs_runtime.Int(orig.second), gopurs_runtime.Int(orig.year)})
				}()).FloatVal()))
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
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Float(gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration_0.V0), d_1).FloatVal()), func() gopurs_runtime.Value {
				orig := Call_Data_DateTime_toRecord(dt_2)
				_ = orig
				return gopurs_runtime.RecordDict([]string{"day", "hour", "millisecond", "minute", "month", "second", "year"}, []gopurs_runtime.Value{gopurs_runtime.Int(orig.day), gopurs_runtime.Int(orig.hour), gopurs_runtime.Int(orig.millisecond), gopurs_runtime.Int(orig.minute), gopurs_runtime.Int(orig.month), gopurs_runtime.Int(orig.second), gopurs_runtime.Int(orig.year)})
				}())
_ = __local_var_3_0
var __t41 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "year").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "year").IntVal) <= (int64(31))) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "year").IntVal)}))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_3:
// TAST (Let): __local_var_4_2 shape=Branch(Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_4_2 := Rebox_Data_DateTime_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))
_ = __local_var_4_2
var __t29 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_4_2 != nil) {
var __t16 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(1)) {
var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_2).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal)))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_4:
__t16 = __t4
goto end_branch_16
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(2)) {
var __t5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_2).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal)))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_5:
__t16 = __t5
goto end_branch_16
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(3)) {
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_2).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal)))
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_6:
__t16 = __t6
goto end_branch_16
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(4)) {
var __t7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_2).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal)))
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_7:
__t16 = __t7
goto end_branch_16
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(5)) {
var __t8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_2).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal)))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_8:
__t16 = __t8
goto end_branch_16
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(6)) {
var __t9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_2).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal)))
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_9:
__t16 = __t9
goto end_branch_16
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(7)) {
var __t10 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_2).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal)))
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_10:
__t16 = __t10
goto end_branch_16
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(8)) {
var __t11 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_2).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal)))
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_11:
__t16 = __t11
goto end_branch_16
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(9)) {
var __t12 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_2).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal)))
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_12:
__t16 = __t12
goto end_branch_16
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(10)) {
var __t13 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_2).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal)))
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_13:
__t16 = __t13
goto end_branch_16
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(11)) {
var __t14 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_2).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal)))
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_14:
__t16 = __t14
goto end_branch_16
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(12)) {
var __t15 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int((__local_var_4_2).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal)))
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_15:
__t16 = __t15
goto end_branch_16
} else {

}
}
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_16:
__t29 = __t16
goto end_branch_29
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(1)) {
var __t17 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_17:
__t29 = __t17
goto end_branch_29
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(2)) {
var __t18 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_18:
__t29 = __t18
goto end_branch_29
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(3)) {
var __t19 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_19:
__t29 = __t19
goto end_branch_29
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(4)) {
var __t20 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t20 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_20
} else {

}
}
{
__t20 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_20:
__t29 = __t20
goto end_branch_29
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(5)) {
var __t21 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t21 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_21
} else {

}
}
{
__t21 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_21:
__t29 = __t21
goto end_branch_29
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(6)) {
var __t22 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_22:
__t29 = __t22
goto end_branch_29
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(7)) {
var __t23 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_23:
__t29 = __t23
goto end_branch_29
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(8)) {
var __t24 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t24 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_24
} else {

}
}
{
__t24 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_24:
__t29 = __t24
goto end_branch_29
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(9)) {
var __t25 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_25:
__t29 = __t25
goto end_branch_29
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(10)) {
var __t26 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_26:
__t29 = __t26
goto end_branch_29
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(11)) {
var __t27 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t27 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_27
} else {

}
}
{
__t27 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_27:
__t29 = __t27
goto end_branch_29
} else {

}
}
{
if (gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "month").IntVal) == (int64(12)) {
var __t28 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t28 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_28:
__t29 = __t28
goto end_branch_29
} else {

}
}
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "day").IntVal) <= (int64(31))) {
__t29 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_29
} else {

}
}
{
__t29 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_29:
// TAST (Let): __local_var_4_1 shape=Let(Branch(Branch(Branch(App(Var), def=Other), Branch(App(Var), def=Other), Branch(App(Var), def=Other), Branch(App(Var), def=Other), Branch(App(Var), def=Other), Branch(App(Var), def=Other), Branch(App(Var), def=Other), Branch(App(Var), def=Other), Branch(App(Var), def=Other), Branch(App(Var), def=Other), Branch(App(Var), def=Other), Branch(App(Var), def=Other), Other, def=Other), Branch(Other, def=Other), Branch(Other, def=Other), Branch(Other, def=Other), Branch(Other, def=Other), Branch(Other, def=Other), Branch(Other, def=Other), Branch(Other, def=Other), Branch(Other, def=Other), Branch(Other, def=Other), Branch(Other, def=Other), Branch(Other, def=Other), Branch(Other, def=Other), Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Date"] [])])
var __local_var_4_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t29)})
var __t40 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_4_1 != nil) {
var __t38 gopurs_runtime.Value
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "hour").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "hour").IntVal) <= (int64(31))) {
var __t37 gopurs_runtime.Value
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "minute").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "minute").IntVal) <= (int64(31))) {
var __t35 gopurs_runtime.Value
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "second").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "second").IntVal) <= (int64(31))) {
__t35 = gopurs_runtime.Func(func(__eta_norm_0_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_33 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_5)
if (__t_tag_33 != nil) {
__t34 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((&Constructor_Data_Time_Time{1, gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "hour").IntVal, gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "minute").IntVal, gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "second").IntVal, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__eta_norm_0_5.UnsafePtr).V0.IntVal}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_34
} else {

}
}
{
__t34 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_34:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t34)}
})
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.Func(func(__eta_norm_0_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
})
}
end_branch_35:
var __t36 gopurs_runtime.Value
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) <= (int64(31))) {
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal)}))}
goto end_branch_36
} else {

}
}
{
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_36:
__t37 = gopurs_runtime.Apply(__t35, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_1170268447_3094389156(Rebox_Data_DateTime_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t36))))})
goto end_branch_37
} else {

}
}
{
var __t31 gopurs_runtime.Value
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "second").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "second").IntVal) <= (int64(31))) {
__t31 = gopurs_runtime.Func(func(__eta_norm_0_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
})
goto end_branch_31
} else {

}
}
{
__t31 = gopurs_runtime.Func(func(__eta_norm_0_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
})
}
end_branch_31:
var __t32 gopurs_runtime.Value
{
if ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) >= (int64(1))) && ((gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal) <= (int64(31))) {
__t32 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, "millisecond").IntVal)}))}
goto end_branch_32
} else {

}
}
{
__t32 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_32:
__t37 = gopurs_runtime.Apply(__t31, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_DateTime_1170268447_3094389156(Rebox_Data_DateTime_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t32))))})
}
end_branch_37:
__t38 = __t37
goto end_branch_38
} else {

}
}
{
__t38 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
}
end_branch_38:
// TAST (Let): __local_var_5_30 shape=Branch(Branch(App(Branch(Abs(Branch(Other, def=Other)), def=Abs(Other))), def=App(Branch(Abs(Other), def=Abs(Other)))), def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Time","Time"] [])])
__local_var_5_30 := Rebox_Data_DateTime_3094389156_3839235747(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t38))
_ = __local_var_5_30
var __t39 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_30 != nil) {
__t39 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1665554298, UnsafePtr: unsafe.Pointer((&Constructor_Data_DateTime_DateTime{1, (__local_var_4_1).V0, (__local_var_5_30).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_39
} else {

}
}
{
__t39 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_39:
__t40 = __t39
goto end_branch_40
} else {

}
}
{
__t40 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_40:
__t41 = __t40
goto end_branch_41
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr == nil) {
__t41 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_41
} else {

}
}
{
__t41 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_41:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t41)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Rebox_Data_DateTime_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_DateTime_1888627657_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_DateTime_DateTime]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
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

func Rebox_Data_DateTime_3094389156_3839235747(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](in.V0)
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

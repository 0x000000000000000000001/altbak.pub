package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Date_toEnum gopurs_runtime.Value
var once_Data_Date_toEnum sync.Once
func Get_Data_Date_toEnum() gopurs_runtime.Value {
	once_Data_Date_toEnum.Do(func() {
		cache_Data_Date_toEnum = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[int64]](Get_Data_Date_Component_boundedEnumDay()).V4)
	})
	return cache_Data_Date_toEnum
}

var cache_Data_Date_ordMaybe gopurs_runtime.Value
var once_Data_Date_ordMaybe sync.Once
func Get_Data_Date_ordMaybe() gopurs_runtime.Value {
	once_Data_Date_ordMaybe.Do(func() {
		cache_Data_Date_ordMaybe = func() gopurs_runtime.Value {
// TAST (Let): eqMaybe1_0_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])
eqMaybe1_0_0 := (&Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_0)
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_1)
__t6 = (__t_tag_5 == nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_0)
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_1)
__t_and_3 = ((__t_tag_2 != nil)) && (((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_0.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_1.UnsafePtr).V0.IntVal))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})})
_ = eqMaybe1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3092443796_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[int64]]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3508461103_3790796878(eqMaybe1_0_0))}
}), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 uint32
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1)
if (__t_tag_7 == nil) {
var __t9 uint32
{
var __t_tag_8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_2)
if (__t_tag_8 == nil) {
__t9 = 902936544
goto end_branch_9
} else {

}
}
{
__t9 = 1527465420
}
end_branch_9:
__t14 = __t9
goto end_branch_14
} else {

}
}
{
var __t_tag_10 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_2)
if (__t_tag_10 == nil) {
__t14 = 380165415
goto end_branch_14
} else {

}
}
{
var __t_tag_11 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1)
var __t_and_13 bool = false
if (__t_tag_11 != nil) {

var __t_tag_12 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_2)
__t_and_13 = (__t_tag_12 != nil)
}
if __t_and_13 {
__t14 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[int64]](Get_Data_Ord_ordInt()).V1), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal)
goto end_branch_14
} else {

}
}
{
__t14 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t14), UnsafePtr: nil}
})})))}
}()
	})
	return cache_Data_Date_ordMaybe
}

var cache_Data_Date_Date gopurs_runtime.Value
var once_Data_Date_Date sync.Once
func Get_Data_Date_Date() gopurs_runtime.Value {
	once_Data_Date_Date.Do(func() {
		cache_Data_Date_Date = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, value0.IntVal, uint32(value1.IntVal), value2.IntVal}))}
})
})
})
	})
	return cache_Data_Date_Date
}

var cache_Data_Date_year gopurs_runtime.Value
var once_Data_Date_year sync.Once
func Get_Data_Date_year() gopurs_runtime.Value {
	once_Data_Date_year.Do(func() {
		cache_Data_Date_year = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Date_year(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_0_box)))
})
	})
	return cache_Data_Date_year
}

var cache_Data_Date_weekday gopurs_runtime.Value
var once_Data_Date_weekday sync.Once
func Get_Data_Date_weekday() gopurs_runtime.Value {
	once_Data_Date_weekday.Do(func() {
		cache_Data_Date_weekday = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Date_weekday(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_0_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Date_weekday
}

var cache_Data_Date_showDate gopurs_runtime.Value
var once_Data_Date_showDate sync.Once
func Get_Data_Date_showDate() gopurs_runtime.Value {
	once_Data_Date_showDate.Do(func() {
		cache_Data_Date_showDate = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1726078761_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Date_Date]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 string
{
var __t_tag_0 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_0) == 1908470532) {
__t12 = "January"
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_1) == 2455627378) {
__t12 = "February"
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_2) == 4162469099) {
__t12 = "March"
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_3) == 1692989816) {
__t12 = "April"
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_4) == 330658827) {
__t12 = "May"
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_5) == 4067355978) {
__t12 = "June"
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_6) == 2276710548) {
__t12 = "July"
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_7) == 243771071) {
__t12 = "August"
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_8) == 215731793) {
__t12 = "September"
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_9) == 8639228) {
__t12 = "October"
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_10) == 49471444) {
__t12 = "November"
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_11) == 3889233761) {
__t12 = "December"
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_12:
return gopurs_runtime.Str((((((("(Date ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0)).StrVal())) + (" ")) + (__t12)) + (" ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V2)).StrVal())) + (")"))
})})))}
	})
	return cache_Data_Date_showDate
}

var cache_Data_Date_month gopurs_runtime.Value
var once_Data_Date_month sync.Once
func Get_Data_Date_month() gopurs_runtime.Value {
	once_Data_Date_month.Do(func() {
		cache_Data_Date_month = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Date_month(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_0_box))), UnsafePtr: nil}
})
	})
	return cache_Data_Date_month
}

var cache_Data_Date_isLeapYear gopurs_runtime.Value
var once_Data_Date_isLeapYear sync.Once
func Get_Data_Date_isLeapYear() gopurs_runtime.Value {
	once_Data_Date_isLeapYear.Do(func() {
		cache_Data_Date_isLeapYear = gopurs_runtime.Func(func(y_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Date_isLeapYear(y_0_box.IntVal))
})
	})
	return cache_Data_Date_isLeapYear
}

var cache_Data_Date_lastDayOfMonth gopurs_runtime.Value
var once_Data_Date_lastDayOfMonth sync.Once
func Get_Data_Date_lastDayOfMonth() gopurs_runtime.Value {
	once_Data_Date_lastDayOfMonth.Do(func() {
		cache_Data_Date_lastDayOfMonth = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth(y_0_box.IntVal, uint32(m_1_box.IntVal)))
})
	})
	return cache_Data_Date_lastDayOfMonth
}

var cache_Data_Date_eqDate gopurs_runtime.Value
var once_Data_Date_eqDate sync.Once
func Get_Data_Date_eqDate() gopurs_runtime.Value {
	once_Data_Date_eqDate.Do(func() {
		cache_Data_Date_eqDate = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1957390985_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Date_Date]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_26 bool = false
if ((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V0) == ((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V0) {

var __t25 bool
{
var __t_tag_3 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
if (uint32(__t_tag_3) == 1908470532) {
var __t_tag_4 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
__t25 = (uint32(__t_tag_4) == 1908470532)
goto end_branch_25
} else {

}
}
{
var __t_tag_5 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
if (uint32(__t_tag_5) == 2455627378) {
var __t_tag_6 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
__t25 = (uint32(__t_tag_6) == 2455627378)
goto end_branch_25
} else {

}
}
{
var __t_tag_7 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
if (uint32(__t_tag_7) == 4162469099) {
var __t_tag_8 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
__t25 = (uint32(__t_tag_8) == 4162469099)
goto end_branch_25
} else {

}
}
{
var __t_tag_9 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
if (uint32(__t_tag_9) == 1692989816) {
var __t_tag_10 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
__t25 = (uint32(__t_tag_10) == 1692989816)
goto end_branch_25
} else {

}
}
{
var __t_tag_11 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
if (uint32(__t_tag_11) == 330658827) {
var __t_tag_12 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
__t25 = (uint32(__t_tag_12) == 330658827)
goto end_branch_25
} else {

}
}
{
var __t_tag_13 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
if (uint32(__t_tag_13) == 4067355978) {
var __t_tag_14 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
__t25 = (uint32(__t_tag_14) == 4067355978)
goto end_branch_25
} else {

}
}
{
var __t_tag_15 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
if (uint32(__t_tag_15) == 2276710548) {
var __t_tag_16 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
__t25 = (uint32(__t_tag_16) == 2276710548)
goto end_branch_25
} else {

}
}
{
var __t_tag_17 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
if (uint32(__t_tag_17) == 243771071) {
var __t_tag_18 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
__t25 = (uint32(__t_tag_18) == 243771071)
goto end_branch_25
} else {

}
}
{
var __t_tag_19 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
if (uint32(__t_tag_19) == 215731793) {
var __t_tag_20 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
__t25 = (uint32(__t_tag_20) == 215731793)
goto end_branch_25
} else {

}
}
{
var __t_tag_21 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
if (uint32(__t_tag_21) == 8639228) {
var __t_tag_22 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
__t25 = (uint32(__t_tag_22) == 8639228)
goto end_branch_25
} else {

}
}
{
var __t_tag_23 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
if (uint32(__t_tag_23) == 49471444) {
var __t_tag_24 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
__t25 = (uint32(__t_tag_24) == 49471444)
goto end_branch_25
} else {

}
}
{
var __t_tag_0 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
var __t_and_2 bool = false
if (uint32(__t_tag_0) == 3889233761) {

var __t_tag_1 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
__t_and_2 = (uint32(__t_tag_1) == 3889233761)
}
__t25 = __t_and_2
}
end_branch_25:
__t_and_26 = __t25
}
return gopurs_runtime.Bool((__t_and_26) && (((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V2) == ((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V2)))
})})))}
	})
	return cache_Data_Date_eqDate
}

var cache_Data_Date_ordDate gopurs_runtime.Value
var once_Data_Date_ordDate sync.Once
func Get_Data_Date_ordDate() gopurs_runtime.Value {
	once_Data_Date_ordDate.Do(func() {
		cache_Data_Date_ordDate = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_758368489_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1957390985_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[*Constructor_Data_Date_Date]](Get_Data_Date_eqDate())))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_2_0 := uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V0)).IntVal)
_ = v_2_0
var __t3 uint32
{
if (v_2_0 == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (v_2_0 == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
// TAST (Let): v1_3_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_3_1 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](Get_Data_Date_Component_ordMonth()).V1), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1), UnsafePtr: nil}).IntVal)
_ = v1_3_1
var __t2 uint32
{
if (v1_3_1 == 1527465420) {
__t2 = 1527465420
goto end_branch_2
} else {

}
}
{
if (v1_3_1 == 380165415) {
__t2 = 380165415
goto end_branch_2
} else {

}
}
{
__t2 = uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V2)).IntVal)
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Date_ordDate
}

var cache_Data_Date_enumDate gopurs_runtime.Value
var once_Data_Date_enumDate sync.Once
func Get_Data_Date_enumDate() gopurs_runtime.Value {
	once_Data_Date_enumDate.Do(func() {
		cache_Data_Date_enumDate = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3494563625_556578094((&Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_758368489_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date]](Get_Data_Date_ordDate())))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pm_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
pm_1_0 := Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_pred__4151667621(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil})))
_ = pm_1_0
// TAST (Let): __local_var_2_2 shape=Other bindingType=Int
__local_var_2_2 := (gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V2).IntVal) - (int64(1))
_ = __local_var_2_2
var __t3 gopurs_runtime.Value
{
if ((__local_var_2_2) >= (int64(1))) && ((__local_var_2_2) <= (int64(31))) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_2_2)}))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_3:
// TAST (Let): pd_2_1 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
pd_2_1 := Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))
_ = pd_2_1
var __t7 uint32
{
var __t5 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t5 = gopurs_runtime.Bool(true)
goto end_branch_5
} else {

}
}
{
if (pd_2_1 != nil) {
__t5 = gopurs_runtime.Bool(false)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
if (__t5.IntVal) != (0) {
var __t6 gopurs_runtime.Value
{
if (pm_1_0 == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (pm_1_0 != nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64((pm_1_0).V0), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = uint32(__t6.IntVal)
goto end_branch_7
} else {

}
}
{
__t7 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
}
end_branch_7:
// TAST (Let): m_prime__3_4 shape=Branch(Branch(Other, Other, def=Other), def=Other) bindingType=Any
m_prime__3_4 := __t7
_ = m_prime__3_4
var __t10 int64
{
if (m_prime__3_4 == 1908470532) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 2455627378) {
var __t9 int64
{
if ((((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0) % (int64(4))) == (int64(0))) && (((((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0) % (int64(400))) == (int64(0))) || (((((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0) % (int64(100))) == (int64(0))) != (true))) {
__t9 = int64(29)
goto end_branch_9
} else {

}
}
{
__t9 = int64(28)
}
end_branch_9:
__t10 = __t9
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 4162469099) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 1692989816) {
__t10 = int64(30)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 330658827) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 4067355978) {
__t10 = int64(30)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 2276710548) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 243771071) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 215731793) {
__t10 = int64(30)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 8639228) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 49471444) {
__t10 = int64(30)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 3889233761) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
__t10 = func() int64 { panic("Failed pattern match") }()
}
end_branch_10:
// TAST (Let): l_4_8 shape=Branch(LitInt, Branch(LitInt, def=LitInt), LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, def=Other) bindingType=Int
l_4_8 := __t10
_ = l_4_8
var __t28 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t15 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t15 = gopurs_runtime.Bool(true)
goto end_branch_15
} else {

}
}
{
if (pd_2_1 != nil) {
__t15 = gopurs_runtime.Bool(false)
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
var __t_and_17 bool = false
if (__t15.IntVal) != (0) {

var __t16 gopurs_runtime.Value
{
if (pm_1_0 == nil) {
__t16 = gopurs_runtime.Bool(true)
goto end_branch_16
} else {

}
}
{
if (pm_1_0 != nil) {
__t16 = gopurs_runtime.Bool(false)
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t_and_17 = (__t16.IntVal) != (0)
}
if __t_and_17 {
var __t27 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
// TAST (Let): __local_var_5_18 shape=Other bindingType=Int
__local_var_5_18 := (gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0).IntVal) - (int64(1))
_ = __local_var_5_18
var __t19 gopurs_runtime.Value
{
if ((__local_var_5_18) >= (int64(1))) && ((__local_var_5_18) <= (int64(31))) {
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_5_18)}))}
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_19:
var __t_tag_20 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t19))
if (__t_tag_20 != nil) {
// TAST (Let): __local_var_5_22 shape=Other bindingType=Int
__local_var_5_22 := (gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0).IntVal) - (int64(1))
_ = __local_var_5_22
var __t23 gopurs_runtime.Value
{
if ((__local_var_5_22) >= (int64(1))) && ((__local_var_5_22) <= (int64(31))) {
__t23 = gopurs_runtime.Int(__local_var_5_22)
goto end_branch_23
} else {

}
}
{
__t23 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_23:
// TAST (Let): __local_var_5_21 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_5_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), __t23, gopurs_runtime.Value{Type: 9, IntVal: int64(m_prime__3_4), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_5_21
var __t26 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t25 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t25 = gopurs_runtime.Bool(true)
goto end_branch_25
} else {

}
}
{
if (pd_2_1 != nil) {
__t25 = gopurs_runtime.Bool(false)
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
if (__t25.IntVal) != (0) {
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_21).V0, gopurs_runtime.Int(l_4_8)), true}
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
var __t24 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (pd_2_1 != nil) {
__t24 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_21).V0, gopurs_runtime.Int((pd_2_1).V0)), true}
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
__t26 = __t24
}
end_branch_26:
__t27 = __t26
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
__t28 = __t27
goto end_branch_28
} else {

}
}
{
// TAST (Let): __local_var_5_11 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_5_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(m_prime__3_4), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_5_11
var __t14 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t13 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t13 = gopurs_runtime.Bool(true)
goto end_branch_13
} else {

}
}
{
if (pd_2_1 != nil) {
__t13 = gopurs_runtime.Bool(false)
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
if (__t13.IntVal) != (0) {
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_11).V0, gopurs_runtime.Int(l_4_8)), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_14
} else {

}
}
{
var __t12 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (pd_2_1 != nil) {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_11).V0, gopurs_runtime.Int((pd_2_1).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
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
__t14 = __t12
}
end_branch_14:
__t28 = __t14
}
end_branch_28:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t28)}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): sm_1_29 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
sm_1_29 := Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_succ__4151667621(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil})))
_ = sm_1_29
// TAST (Let): __local_var_2_31 shape=Other bindingType=Int
__local_var_2_31 := ((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V2) + (int64(1))
_ = __local_var_2_31
var __t32 gopurs_runtime.Value
{
if ((__local_var_2_31) >= (int64(1))) && ((__local_var_2_31) <= (int64(31))) {
__t32 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_2_31)}))}
goto end_branch_32
} else {

}
}
{
__t32 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_32:
// TAST (Let): v1_2_30 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
v1_2_30 := Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t32))
_ = v1_2_30
var __t49 *Constructor_Data_Maybe_Just[int64]
{
var __t47 int64
{
var __t_tag_34 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_34) == 1908470532) {
__t47 = int64(31)
goto end_branch_47
} else {

}
}
{
var __t_tag_35 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_35) == 2455627378) {
var __t36 int64
{
if ((((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0) % (int64(4))) == (int64(0))) && (((((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0) % (int64(400))) == (int64(0))) || (((((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0) % (int64(100))) == (int64(0))) != (true))) {
__t36 = int64(29)
goto end_branch_36
} else {

}
}
{
__t36 = int64(28)
}
end_branch_36:
__t47 = __t36
goto end_branch_47
} else {

}
}
{
var __t_tag_37 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_37) == 4162469099) {
__t47 = int64(31)
goto end_branch_47
} else {

}
}
{
var __t_tag_38 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_38) == 1692989816) {
__t47 = int64(30)
goto end_branch_47
} else {

}
}
{
var __t_tag_39 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_39) == 330658827) {
__t47 = int64(31)
goto end_branch_47
} else {

}
}
{
var __t_tag_40 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_40) == 4067355978) {
__t47 = int64(30)
goto end_branch_47
} else {

}
}
{
var __t_tag_41 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_41) == 2276710548) {
__t47 = int64(31)
goto end_branch_47
} else {

}
}
{
var __t_tag_42 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_42) == 243771071) {
__t47 = int64(31)
goto end_branch_47
} else {

}
}
{
var __t_tag_43 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_43) == 215731793) {
__t47 = int64(30)
goto end_branch_47
} else {

}
}
{
var __t_tag_44 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_44) == 8639228) {
__t47 = int64(31)
goto end_branch_47
} else {

}
}
{
var __t_tag_45 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_45) == 49471444) {
__t47 = int64(30)
goto end_branch_47
} else {

}
}
{
var __t_tag_46 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
if (uint32(__t_tag_46) == 3889233761) {
__t47 = int64(31)
goto end_branch_47
} else {

}
}
{
__t47 = func() int64 { panic("Failed pattern match") }()
}
end_branch_47:
var __t_tag_48 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[int64]]](Get_Data_Interval_ordMaybe()).V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(v1_2_30))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__t47)}))})
if (uint32(__t_tag_48.IntVal) == 380165415) {
__t49 = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
goto end_branch_49
} else {

}
}
{
__t49 = v1_2_30
}
end_branch_49:
// TAST (Let): sd_3_33 shape=Branch(Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
sd_3_33 := __t49
_ = sd_3_33
var __t73 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t57 gopurs_runtime.Value
{
if (sd_3_33 == nil) {
__t57 = gopurs_runtime.Bool(true)
goto end_branch_57
} else {

}
}
{
if (sd_3_33 != nil) {
__t57 = gopurs_runtime.Bool(false)
goto end_branch_57
} else {

}
}
{
__t57 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_57:
var __t_and_59 bool = false
if (__t57.IntVal) != (0) {

var __t58 gopurs_runtime.Value
{
if (sm_1_29 == nil) {
__t58 = gopurs_runtime.Bool(true)
goto end_branch_58
} else {

}
}
{
if (sm_1_29 != nil) {
__t58 = gopurs_runtime.Bool(false)
goto end_branch_58
} else {

}
}
{
__t58 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_58:
__t_and_59 = (__t58.IntVal) != (0)
}
if __t_and_59 {
var __t72 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
// TAST (Let): __local_var_4_60 shape=Other bindingType=Int
__local_var_4_60 := ((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0) + (int64(1))
_ = __local_var_4_60
var __t61 gopurs_runtime.Value
{
if ((__local_var_4_60) >= (int64(1))) && ((__local_var_4_60) <= (int64(31))) {
__t61 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_4_60)}))}
goto end_branch_61
} else {

}
}
{
__t61 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_61:
var __t_tag_62 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t61))
if (__t_tag_62 != nil) {
// TAST (Let): __local_var_4_64 shape=Other bindingType=Int
__local_var_4_64 := ((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0) + (int64(1))
_ = __local_var_4_64
var __t65 gopurs_runtime.Value
{
if ((__local_var_4_64) >= (int64(1))) && ((__local_var_4_64) <= (int64(31))) {
__t65 = gopurs_runtime.Int(__local_var_4_64)
goto end_branch_65
} else {

}
}
{
__t65 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_65:
var __t68 uint32
{
var __t66 gopurs_runtime.Value
{
if (sd_3_33 == nil) {
__t66 = gopurs_runtime.Bool(true)
goto end_branch_66
} else {

}
}
{
if (sd_3_33 != nil) {
__t66 = gopurs_runtime.Bool(false)
goto end_branch_66
} else {

}
}
{
__t66 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_66:
if (__t66.IntVal) != (0) {
var __t67 gopurs_runtime.Value
{
if (sm_1_29 == nil) {
__t67 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_67
} else {

}
}
{
if (sm_1_29 != nil) {
__t67 = gopurs_runtime.Value{Type: 9, IntVal: int64((sm_1_29).V0), UnsafePtr: nil}
goto end_branch_67
} else {

}
}
{
__t67 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_67:
__t68 = uint32(__t67.IntVal)
goto end_branch_68
} else {

}
}
{
__t68 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
}
end_branch_68:
// TAST (Let): __local_var_4_63 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_4_63 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), __t65, gopurs_runtime.Value{Type: 9, IntVal: int64(__t68), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_4_63
var __t71 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t70 gopurs_runtime.Value
{
if (sd_3_33 == nil) {
__t70 = gopurs_runtime.Bool(true)
goto end_branch_70
} else {

}
}
{
if (sd_3_33 != nil) {
__t70 = gopurs_runtime.Bool(false)
goto end_branch_70
} else {

}
}
{
__t70 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_70:
if (__t70.IntVal) != (0) {
__t71 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_4_63).V0, gopurs_runtime.Int(int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_71
} else {

}
}
{
var __t69 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (sd_3_33 != nil) {
__t69 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_4_63).V0, gopurs_runtime.Int((sd_3_33).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_69
} else {

}
}
{
__t69 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_69:
__t71 = __t69
}
end_branch_71:
__t72 = __t71
goto end_branch_72
} else {

}
}
{
__t72 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_72:
__t73 = __t72
goto end_branch_73
} else {

}
}
{
var __t53 uint32
{
var __t51 gopurs_runtime.Value
{
if (sd_3_33 == nil) {
__t51 = gopurs_runtime.Bool(true)
goto end_branch_51
} else {

}
}
{
if (sd_3_33 != nil) {
__t51 = gopurs_runtime.Bool(false)
goto end_branch_51
} else {

}
}
{
__t51 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_51:
if (__t51.IntVal) != (0) {
var __t52 gopurs_runtime.Value
{
if (sm_1_29 == nil) {
__t52 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_52
} else {

}
}
{
if (sm_1_29 != nil) {
__t52 = gopurs_runtime.Value{Type: 9, IntVal: int64((sm_1_29).V0), UnsafePtr: nil}
goto end_branch_52
} else {

}
}
{
__t52 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_52:
__t53 = uint32(__t52.IntVal)
goto end_branch_53
} else {

}
}
{
__t53 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
}
end_branch_53:
// TAST (Let): __local_var_4_50 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_4_50 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(__t53), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_4_50
var __t56 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t55 gopurs_runtime.Value
{
if (sd_3_33 == nil) {
__t55 = gopurs_runtime.Bool(true)
goto end_branch_55
} else {

}
}
{
if (sd_3_33 != nil) {
__t55 = gopurs_runtime.Bool(false)
goto end_branch_55
} else {

}
}
{
__t55 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_55:
if (__t55.IntVal) != (0) {
__t56 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_4_50).V0, gopurs_runtime.Int(int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_56
} else {

}
}
{
var __t54 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (sd_3_33 != nil) {
__t54 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_4_50).V0, gopurs_runtime.Int((sd_3_33).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_54
} else {

}
}
{
__t54 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_54:
__t56 = __t54
}
end_branch_56:
__t73 = __t56
}
end_branch_73:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t73)}
})})))}
	})
	return cache_Data_Date_enumDate
}

var cache_Data_Date_pred gopurs_runtime.Value
var once_Data_Date_pred sync.Once
func Get_Data_Date_pred() gopurs_runtime.Value {
	once_Data_Date_pred.Do(func() {
		cache_Data_Date_pred = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Date_pred(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Date_pred
}

var cache_Data_Date_diff gopurs_runtime.Value
var once_Data_Date_diff sync.Once
func Get_Data_Date_diff() gopurs_runtime.Value {
	once_Data_Date_diff.Do(func() {
		cache_Data_Date_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Date_diff(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dictDuration_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v1_2_box))
})
	})
	return cache_Data_Date_diff
}

var cache_Data_Date_day gopurs_runtime.Value
var once_Data_Date_day sync.Once
func Get_Data_Date_day() gopurs_runtime.Value {
	once_Data_Date_day.Do(func() {
		cache_Data_Date_day = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Date_day(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_0_box)))
})
	})
	return cache_Data_Date_day
}

var cache_Data_Date_canonicalDate gopurs_runtime.Value
var once_Data_Date_canonicalDate sync.Once
func Get_Data_Date_canonicalDate() gopurs_runtime.Value {
	once_Data_Date_canonicalDate.Do(func() {
		cache_Data_Date_canonicalDate = gopurs_runtime.Func3(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Date_canonicalDate(y_0_box.IntVal, uint32(m_1_box.IntVal), d_2_box.IntVal))}
})
	})
	return cache_Data_Date_canonicalDate
}

var cache_Data_Date_exactDate gopurs_runtime.Value
var once_Data_Date_exactDate sync.Once
func Get_Data_Date_exactDate() gopurs_runtime.Value {
	once_Data_Date_exactDate.Do(func() {
		cache_Data_Date_exactDate = gopurs_runtime.Func3(func(y_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Date_exactDate(y_0_box.IntVal, uint32(m_1_box.IntVal), d_2_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Date_exactDate
}

var cache_Data_Date_boundedDate gopurs_runtime.Value
var once_Data_Date_boundedDate sync.Once
func Get_Data_Date_boundedDate() gopurs_runtime.Value {
	once_Data_Date_boundedDate.Do(func() {
		cache_Data_Date_boundedDate = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3523628265_2094947566((&Constructor_Data_Bounded_Bounded[*Constructor_Data_Date_Date]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_758368489_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date]](Get_Data_Date_ordDate())))}
}), (&Constructor_Data_Date_Date{1, int64(-271820), 1908470532, int64(1)}), (&Constructor_Data_Date_Date{1, int64(275759), 3889233761, int64(31)})})))}
	})
	return cache_Data_Date_boundedDate
}

var cache_Data_Date_adjust gopurs_runtime.Value
var once_Data_Date_adjust sync.Once
func Get_Data_Date_adjust() gopurs_runtime.Value {
	once_Data_Date_adjust.Do(func() {
		cache_Data_Date_adjust = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, date_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Date_adjust(v_0_box.FloatVal(), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](date_1_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Date_adjust
}

type Constructor_Data_Date_Date struct {
	Rc uint32
	V0 int64
	V1 uint32
	V2 int64
}


func Call_Data_Date_year(v_0_loop *Constructor_Data_Date_Date) int64 {
var v_0 *Constructor_Data_Date_Date = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_Date_weekday(v_0_loop *Constructor_Data_Date_Date) uint32 {
var v_0 *Constructor_Data_Date_Date = v_0_loop
_ = v_0
var __t13 int64
{
var __t_tag_1 uint32 = (v_0).V1
if (uint32(__t_tag_1) == 1908470532) {
__t13 = int64(1)
goto end_branch_13
} else {

}
}
{
var __t_tag_2 uint32 = (v_0).V1
if (uint32(__t_tag_2) == 2455627378) {
__t13 = int64(2)
goto end_branch_13
} else {

}
}
{
var __t_tag_3 uint32 = (v_0).V1
if (uint32(__t_tag_3) == 4162469099) {
__t13 = int64(3)
goto end_branch_13
} else {

}
}
{
var __t_tag_4 uint32 = (v_0).V1
if (uint32(__t_tag_4) == 1692989816) {
__t13 = int64(4)
goto end_branch_13
} else {

}
}
{
var __t_tag_5 uint32 = (v_0).V1
if (uint32(__t_tag_5) == 330658827) {
__t13 = int64(5)
goto end_branch_13
} else {

}
}
{
var __t_tag_6 uint32 = (v_0).V1
if (uint32(__t_tag_6) == 4067355978) {
__t13 = int64(6)
goto end_branch_13
} else {

}
}
{
var __t_tag_7 uint32 = (v_0).V1
if (uint32(__t_tag_7) == 2276710548) {
__t13 = int64(7)
goto end_branch_13
} else {

}
}
{
var __t_tag_8 uint32 = (v_0).V1
if (uint32(__t_tag_8) == 243771071) {
__t13 = int64(8)
goto end_branch_13
} else {

}
}
{
var __t_tag_9 uint32 = (v_0).V1
if (uint32(__t_tag_9) == 215731793) {
__t13 = int64(9)
goto end_branch_13
} else {

}
}
{
var __t_tag_10 uint32 = (v_0).V1
if (uint32(__t_tag_10) == 8639228) {
__t13 = int64(10)
goto end_branch_13
} else {

}
}
{
var __t_tag_11 uint32 = (v_0).V1
if (uint32(__t_tag_11) == 49471444) {
__t13 = int64(11)
goto end_branch_13
} else {

}
}
{
var __t_tag_12 uint32 = (v_0).V1
if (uint32(__t_tag_12) == 3889233761) {
__t13 = int64(12)
goto end_branch_13
} else {

}
}
{
__t13 = func() int64 { panic("Failed pattern match") }()
}
end_branch_13:
// TAST (Let): n_1_0 shape=UncurriedApp(Var) bindingType=Any
n_1_0 := gopurs_runtime.UncurriedApp3(Get_Data_Date_calcWeekday(), gopurs_runtime.Int((v_0).V0), gopurs_runtime.Int(__t13), gopurs_runtime.Int((v_0).V2))
_ = n_1_0
var __t15 uint32
{
if (n_1_0.IntVal) == (int64(0)) {
__t15 = 1326716170
goto end_branch_15
} else {

}
}
{
var __t14 uint32
{
if (n_1_0.IntVal) == (int64(1)) {
__t14 = 2900196686
goto end_branch_14
} else {

}
}
{
if (n_1_0.IntVal) == (int64(2)) {
__t14 = 20457557
goto end_branch_14
} else {

}
}
{
if (n_1_0.IntVal) == (int64(3)) {
__t14 = 4227105004
goto end_branch_14
} else {

}
}
{
if (n_1_0.IntVal) == (int64(4)) {
__t14 = 3818857258
goto end_branch_14
} else {

}
}
{
if (n_1_0.IntVal) == (int64(5)) {
__t14 = 2946274527
goto end_branch_14
} else {

}
}
{
if (n_1_0.IntVal) == (int64(6)) {
__t14 = 1070786179
goto end_branch_14
} else {

}
}
{
if (n_1_0.IntVal) == (int64(7)) {
__t14 = 1326716170
goto end_branch_14
} else {

}
}
{
__t14 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_14:
__t15 = __t14
}
end_branch_15:
return __t15
}

func Call_Data_Date_month(v_0_loop *Constructor_Data_Date_Date) uint32 {
var v_0 *Constructor_Data_Date_Date = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Date_isLeapYear(y_0_loop int64) bool {
var y_0 int64 = y_0_loop
_ = y_0
return (((y_0) % (int64(4))) == (int64(0))) && ((((y_0) % (int64(400))) == (int64(0))) || ((((y_0) % (int64(100))) == (int64(0))) != (true)))
}

func Call_Data_Date_lastDayOfMonth(y_0_loop int64, m_1_loop uint32) int64 {
var y_0 int64 = y_0_loop
_ = y_0
var m_1 uint32 = m_1_loop
_ = m_1
var __t1 int64
{
if (m_1 == 1908470532) {
__t1 = int64(31)
goto end_branch_1
} else {

}
}
{
if (m_1 == 2455627378) {
var __t0 int64
{
if (((y_0) % (int64(4))) == (int64(0))) && ((((y_0) % (int64(400))) == (int64(0))) || ((((y_0) % (int64(100))) == (int64(0))) != (true))) {
__t0 = int64(29)
goto end_branch_0
} else {

}
}
{
__t0 = int64(28)
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if (m_1 == 4162469099) {
__t1 = int64(31)
goto end_branch_1
} else {

}
}
{
if (m_1 == 1692989816) {
__t1 = int64(30)
goto end_branch_1
} else {

}
}
{
if (m_1 == 330658827) {
__t1 = int64(31)
goto end_branch_1
} else {

}
}
{
if (m_1 == 4067355978) {
__t1 = int64(30)
goto end_branch_1
} else {

}
}
{
if (m_1 == 2276710548) {
__t1 = int64(31)
goto end_branch_1
} else {

}
}
{
if (m_1 == 243771071) {
__t1 = int64(31)
goto end_branch_1
} else {

}
}
{
if (m_1 == 215731793) {
__t1 = int64(30)
goto end_branch_1
} else {

}
}
{
if (m_1 == 8639228) {
__t1 = int64(31)
goto end_branch_1
} else {

}
}
{
if (m_1 == 49471444) {
__t1 = int64(30)
goto end_branch_1
} else {

}
}
{
if (m_1 == 3889233761) {
__t1 = int64(31)
goto end_branch_1
} else {

}
}
{
__t1 = func() int64 { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_Data_Date_pred(v_0_loop *Constructor_Data_Date_Date) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 *Constructor_Data_Date_Date = v_0_loop
_ = v_0
// TAST (Let): pm_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
pm_1_0 := Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_pred__4151667621(), gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil})))
_ = pm_1_0
// TAST (Let): __local_var_2_2 shape=Other bindingType=Int
__local_var_2_2 := (gopurs_runtime.Int((v_0).V2).IntVal) - (int64(1))
_ = __local_var_2_2
var __t3 gopurs_runtime.Value
{
if ((__local_var_2_2) >= (int64(1))) && ((__local_var_2_2) <= (int64(31))) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_2_2)}))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_3:
// TAST (Let): pd_2_1 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
pd_2_1 := Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))
_ = pd_2_1
var __t7 uint32
{
var __t5 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t5 = gopurs_runtime.Bool(true)
goto end_branch_5
} else {

}
}
{
if (pd_2_1 != nil) {
__t5 = gopurs_runtime.Bool(false)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
if (__t5.IntVal) != (0) {
var __t6 gopurs_runtime.Value
{
if (pm_1_0 == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (pm_1_0 != nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64((pm_1_0).V0), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = uint32(__t6.IntVal)
goto end_branch_7
} else {

}
}
{
__t7 = (v_0).V1
}
end_branch_7:
// TAST (Let): m_prime__3_4 shape=Branch(Branch(Other, Other, def=Other), def=Other) bindingType=Any
m_prime__3_4 := __t7
_ = m_prime__3_4
var __t10 int64
{
if (m_prime__3_4 == 1908470532) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 2455627378) {
var __t9 int64
{
if ((((v_0).V0) % (int64(4))) == (int64(0))) && (((((v_0).V0) % (int64(400))) == (int64(0))) || (((((v_0).V0) % (int64(100))) == (int64(0))) != (true))) {
__t9 = int64(29)
goto end_branch_9
} else {

}
}
{
__t9 = int64(28)
}
end_branch_9:
__t10 = __t9
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 4162469099) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 1692989816) {
__t10 = int64(30)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 330658827) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 4067355978) {
__t10 = int64(30)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 2276710548) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 243771071) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 215731793) {
__t10 = int64(30)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 8639228) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 49471444) {
__t10 = int64(30)
goto end_branch_10
} else {

}
}
{
if (m_prime__3_4 == 3889233761) {
__t10 = int64(31)
goto end_branch_10
} else {

}
}
{
__t10 = func() int64 { panic("Failed pattern match") }()
}
end_branch_10:
// TAST (Let): l_4_8 shape=Branch(LitInt, Branch(LitInt, def=LitInt), LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, def=Other) bindingType=Int
l_4_8 := __t10
_ = l_4_8
var __t28 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t15 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t15 = gopurs_runtime.Bool(true)
goto end_branch_15
} else {

}
}
{
if (pd_2_1 != nil) {
__t15 = gopurs_runtime.Bool(false)
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
var __t_and_17 bool = false
if (__t15.IntVal) != (0) {

var __t16 gopurs_runtime.Value
{
if (pm_1_0 == nil) {
__t16 = gopurs_runtime.Bool(true)
goto end_branch_16
} else {

}
}
{
if (pm_1_0 != nil) {
__t16 = gopurs_runtime.Bool(false)
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t_and_17 = (__t16.IntVal) != (0)
}
if __t_and_17 {
var __t27 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
// TAST (Let): __local_var_5_18 shape=Other bindingType=Int
__local_var_5_18 := (gopurs_runtime.Int((v_0).V0).IntVal) - (int64(1))
_ = __local_var_5_18
var __t19 gopurs_runtime.Value
{
if ((__local_var_5_18) >= (int64(1))) && ((__local_var_5_18) <= (int64(31))) {
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_5_18)}))}
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_19:
var __t_tag_20 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t19))
if (__t_tag_20 != nil) {
// TAST (Let): __local_var_5_22 shape=Other bindingType=Int
__local_var_5_22 := (gopurs_runtime.Int((v_0).V0).IntVal) - (int64(1))
_ = __local_var_5_22
var __t23 gopurs_runtime.Value
{
if ((__local_var_5_22) >= (int64(1))) && ((__local_var_5_22) <= (int64(31))) {
__t23 = gopurs_runtime.Int(__local_var_5_22)
goto end_branch_23
} else {

}
}
{
__t23 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_23:
// TAST (Let): __local_var_5_21 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_5_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), __t23, gopurs_runtime.Value{Type: 9, IntVal: int64(m_prime__3_4), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_5_21
var __t26 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t25 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t25 = gopurs_runtime.Bool(true)
goto end_branch_25
} else {

}
}
{
if (pd_2_1 != nil) {
__t25 = gopurs_runtime.Bool(false)
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
if (__t25.IntVal) != (0) {
__t26 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_21).V0, gopurs_runtime.Int(l_4_8)), true}
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
var __t24 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (pd_2_1 != nil) {
__t24 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_21).V0, gopurs_runtime.Int((pd_2_1).V0)), true}
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
__t26 = __t24
}
end_branch_26:
__t27 = __t26
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
__t28 = __t27
goto end_branch_28
} else {

}
}
{
// TAST (Let): __local_var_5_11 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_5_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((v_0).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(m_prime__3_4), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_5_11
var __t14 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t13 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t13 = gopurs_runtime.Bool(true)
goto end_branch_13
} else {

}
}
{
if (pd_2_1 != nil) {
__t13 = gopurs_runtime.Bool(false)
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
if (__t13.IntVal) != (0) {
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_11).V0, gopurs_runtime.Int(l_4_8)), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_14
} else {

}
}
{
var __t12 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (pd_2_1 != nil) {
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_5_11).V0, gopurs_runtime.Int((pd_2_1).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
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
__t14 = __t12
}
end_branch_14:
__t28 = __t14
}
end_branch_28:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t28)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Date_diff(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value], v_1_loop *Constructor_Data_Date_Date, v1_2_loop *Constructor_Data_Date_Date) gopurs_runtime.Value {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
var v_1 *Constructor_Data_Date_Date = v_1_loop
_ = v_1
var v1_2 *Constructor_Data_Date_Date = v1_2_loop
_ = v1_2
var __t12 int64
{
var __t_tag_0 uint32 = (v_1).V1
if (uint32(__t_tag_0) == 1908470532) {
__t12 = int64(1)
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = (v_1).V1
if (uint32(__t_tag_1) == 2455627378) {
__t12 = int64(2)
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = (v_1).V1
if (uint32(__t_tag_2) == 4162469099) {
__t12 = int64(3)
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = (v_1).V1
if (uint32(__t_tag_3) == 1692989816) {
__t12 = int64(4)
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = (v_1).V1
if (uint32(__t_tag_4) == 330658827) {
__t12 = int64(5)
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = (v_1).V1
if (uint32(__t_tag_5) == 4067355978) {
__t12 = int64(6)
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = (v_1).V1
if (uint32(__t_tag_6) == 2276710548) {
__t12 = int64(7)
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = (v_1).V1
if (uint32(__t_tag_7) == 243771071) {
__t12 = int64(8)
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = (v_1).V1
if (uint32(__t_tag_8) == 215731793) {
__t12 = int64(9)
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = (v_1).V1
if (uint32(__t_tag_9) == 8639228) {
__t12 = int64(10)
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = (v_1).V1
if (uint32(__t_tag_10) == 49471444) {
__t12 = int64(11)
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = (v_1).V1
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
var __t25 int64
{
var __t_tag_13 uint32 = (v1_2).V1
if (uint32(__t_tag_13) == 1908470532) {
__t25 = int64(1)
goto end_branch_25
} else {

}
}
{
var __t_tag_14 uint32 = (v1_2).V1
if (uint32(__t_tag_14) == 2455627378) {
__t25 = int64(2)
goto end_branch_25
} else {

}
}
{
var __t_tag_15 uint32 = (v1_2).V1
if (uint32(__t_tag_15) == 4162469099) {
__t25 = int64(3)
goto end_branch_25
} else {

}
}
{
var __t_tag_16 uint32 = (v1_2).V1
if (uint32(__t_tag_16) == 1692989816) {
__t25 = int64(4)
goto end_branch_25
} else {

}
}
{
var __t_tag_17 uint32 = (v1_2).V1
if (uint32(__t_tag_17) == 330658827) {
__t25 = int64(5)
goto end_branch_25
} else {

}
}
{
var __t_tag_18 uint32 = (v1_2).V1
if (uint32(__t_tag_18) == 4067355978) {
__t25 = int64(6)
goto end_branch_25
} else {

}
}
{
var __t_tag_19 uint32 = (v1_2).V1
if (uint32(__t_tag_19) == 2276710548) {
__t25 = int64(7)
goto end_branch_25
} else {

}
}
{
var __t_tag_20 uint32 = (v1_2).V1
if (uint32(__t_tag_20) == 243771071) {
__t25 = int64(8)
goto end_branch_25
} else {

}
}
{
var __t_tag_21 uint32 = (v1_2).V1
if (uint32(__t_tag_21) == 215731793) {
__t25 = int64(9)
goto end_branch_25
} else {

}
}
{
var __t_tag_22 uint32 = (v1_2).V1
if (uint32(__t_tag_22) == 8639228) {
__t25 = int64(10)
goto end_branch_25
} else {

}
}
{
var __t_tag_23 uint32 = (v1_2).V1
if (uint32(__t_tag_23) == 49471444) {
__t25 = int64(11)
goto end_branch_25
} else {

}
}
{
var __t_tag_24 uint32 = (v1_2).V1
if (uint32(__t_tag_24) == 3889233761) {
__t25 = int64(12)
goto end_branch_25
} else {

}
}
{
__t25 = func() int64 { panic("Failed pattern match") }()
}
end_branch_25:
return gopurs_runtime.Apply(gopurs_runtime.Box(dictDuration_0.V1), gopurs_runtime.Float(gopurs_runtime.UncurriedApp6(Get_Data_Date_calcDiff(), gopurs_runtime.Int((v_1).V0), gopurs_runtime.Int(__t12), gopurs_runtime.Int((v_1).V2), gopurs_runtime.Int((v1_2).V0), gopurs_runtime.Int(__t25), gopurs_runtime.Int((v1_2).V2)).FloatVal()))
}

func Call_Data_Date_day(v_0_loop *Constructor_Data_Date_Date) int64 {
var v_0 *Constructor_Data_Date_Date = v_0_loop
_ = v_0
return (v_0).V2
}

func Call_Data_Date_canonicalDate(y_0_loop int64, m_1_loop uint32, d_2_loop int64) *Constructor_Data_Date_Date {
var y_0 int64 = y_0_loop
_ = y_0
var m_1 uint32 = m_1_loop
_ = m_1
var d_2 int64 = d_2_loop
_ = d_2
var __t1 int64
{
if (m_1 == 1908470532) {
__t1 = int64(1)
goto end_branch_1
} else {

}
}
{
if (m_1 == 2455627378) {
__t1 = int64(2)
goto end_branch_1
} else {

}
}
{
if (m_1 == 4162469099) {
__t1 = int64(3)
goto end_branch_1
} else {

}
}
{
if (m_1 == 1692989816) {
__t1 = int64(4)
goto end_branch_1
} else {

}
}
{
if (m_1 == 330658827) {
__t1 = int64(5)
goto end_branch_1
} else {

}
}
{
if (m_1 == 4067355978) {
__t1 = int64(6)
goto end_branch_1
} else {

}
}
{
if (m_1 == 2276710548) {
__t1 = int64(7)
goto end_branch_1
} else {

}
}
{
if (m_1 == 243771071) {
__t1 = int64(8)
goto end_branch_1
} else {

}
}
{
if (m_1 == 215731793) {
__t1 = int64(9)
goto end_branch_1
} else {

}
}
{
if (m_1 == 8639228) {
__t1 = int64(10)
goto end_branch_1
} else {

}
}
{
if (m_1 == 49471444) {
__t1 = int64(11)
goto end_branch_1
} else {

}
}
{
if (m_1 == 3889233761) {
__t1 = int64(12)
goto end_branch_1
} else {

}
}
{
__t1 = func() int64 { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.UncurriedApp4(Get_Data_Date_canonicalDateImpl(), gopurs_runtime.Func3(func(y_prime__3 gopurs_runtime.Value, m_prime__4 gopurs_runtime.Value, d_prime__5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 uint32
{
if (m_prime__4.IntVal) == (int64(1)) {
__t0 = 1908470532
goto end_branch_0
} else {

}
}
{
if (m_prime__4.IntVal) == (int64(2)) {
__t0 = 2455627378
goto end_branch_0
} else {

}
}
{
if (m_prime__4.IntVal) == (int64(3)) {
__t0 = 4162469099
goto end_branch_0
} else {

}
}
{
if (m_prime__4.IntVal) == (int64(4)) {
__t0 = 1692989816
goto end_branch_0
} else {

}
}
{
if (m_prime__4.IntVal) == (int64(5)) {
__t0 = 330658827
goto end_branch_0
} else {

}
}
{
if (m_prime__4.IntVal) == (int64(6)) {
__t0 = 4067355978
goto end_branch_0
} else {

}
}
{
if (m_prime__4.IntVal) == (int64(7)) {
__t0 = 2276710548
goto end_branch_0
} else {

}
}
{
if (m_prime__4.IntVal) == (int64(8)) {
__t0 = 243771071
goto end_branch_0
} else {

}
}
{
if (m_prime__4.IntVal) == (int64(9)) {
__t0 = 215731793
goto end_branch_0
} else {

}
}
{
if (m_prime__4.IntVal) == (int64(10)) {
__t0 = 8639228
goto end_branch_0
} else {

}
}
{
if (m_prime__4.IntVal) == (int64(11)) {
__t0 = 49471444
goto end_branch_0
} else {

}
}
{
if (m_prime__4.IntVal) == (int64(12)) {
__t0 = 3889233761
goto end_branch_0
} else {

}
}
{
__t0 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, y_prime__3.IntVal, __t0, d_prime__5.IntVal}))}
}), gopurs_runtime.Int(y_0), gopurs_runtime.Int(__t1), gopurs_runtime.Int(d_2)))
}

func Call_Data_Date_exactDate(y_0_loop int64, m_1_loop uint32, d_2_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
var y_0 int64 = y_0_loop
_ = y_0
var m_1 uint32 = m_1_loop
_ = m_1
var d_2 int64 = d_2_loop
_ = d_2
var __t15 gopurs_runtime.Value
{
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
__local_var_3_0 := Call_Data_Date_canonicalDate(y_0, m_1, d_2)
_ = __local_var_3_0
var __t_and_14 bool = false
if ((__local_var_3_0).V0) == (y_0) {

var __t13 bool
{
var __t_tag_2 uint32 = (__local_var_3_0).V1
if (uint32(__t_tag_2) == 1908470532) {
__t13 = (m_1 == 1908470532)
goto end_branch_13
} else {

}
}
{
var __t_tag_3 uint32 = (__local_var_3_0).V1
if (uint32(__t_tag_3) == 2455627378) {
__t13 = (m_1 == 2455627378)
goto end_branch_13
} else {

}
}
{
var __t_tag_4 uint32 = (__local_var_3_0).V1
if (uint32(__t_tag_4) == 4162469099) {
__t13 = (m_1 == 4162469099)
goto end_branch_13
} else {

}
}
{
var __t_tag_5 uint32 = (__local_var_3_0).V1
if (uint32(__t_tag_5) == 1692989816) {
__t13 = (m_1 == 1692989816)
goto end_branch_13
} else {

}
}
{
var __t_tag_6 uint32 = (__local_var_3_0).V1
if (uint32(__t_tag_6) == 330658827) {
__t13 = (m_1 == 330658827)
goto end_branch_13
} else {

}
}
{
var __t_tag_7 uint32 = (__local_var_3_0).V1
if (uint32(__t_tag_7) == 4067355978) {
__t13 = (m_1 == 4067355978)
goto end_branch_13
} else {

}
}
{
var __t_tag_8 uint32 = (__local_var_3_0).V1
if (uint32(__t_tag_8) == 2276710548) {
__t13 = (m_1 == 2276710548)
goto end_branch_13
} else {

}
}
{
var __t_tag_9 uint32 = (__local_var_3_0).V1
if (uint32(__t_tag_9) == 243771071) {
__t13 = (m_1 == 243771071)
goto end_branch_13
} else {

}
}
{
var __t_tag_10 uint32 = (__local_var_3_0).V1
if (uint32(__t_tag_10) == 215731793) {
__t13 = (m_1 == 215731793)
goto end_branch_13
} else {

}
}
{
var __t_tag_11 uint32 = (__local_var_3_0).V1
if (uint32(__t_tag_11) == 8639228) {
__t13 = (m_1 == 8639228)
goto end_branch_13
} else {

}
}
{
var __t_tag_12 uint32 = (__local_var_3_0).V1
if (uint32(__t_tag_12) == 49471444) {
__t13 = (m_1 == 49471444)
goto end_branch_13
} else {

}
}
{
var __t_tag_1 uint32 = (__local_var_3_0).V1
__t13 = ((uint32(__t_tag_1) == 3889233761)) && ((m_1 == 3889233761))
}
end_branch_13:
__t_and_14 = __t13
}
if (__t_and_14) && (((__local_var_3_0).V2) == (d_2)) {
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, y_0, m_1, d_2}))}}))}
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_2280409795_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_15:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_2280409795_3094389156(Rebox_Data_Date_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t15))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Date_adjust(v_0_loop float64, date_1_loop *Constructor_Data_Date_Date) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 float64 = v_0_loop
_ = v_0
var date_1 *Constructor_Data_Date_Date = date_1_loop
_ = date_1
var adj_2_0_0 gopurs_runtime.Value
_ = adj_2_0_0
// FALLBACK TCO: isLoop=false len=1
adj_2_0_0 = gopurs_runtime.Func2(func(v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t201 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_3.IntVal) == (int64(0)) {
__t201 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v2_4))}})
goto end_branch_201
} else {

}
}
{
// TAST (Let): j_5_1 shape=Other bindingType=Int
j_5_1 := (v1_3.IntVal) + ((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V2)
_ = j_5_1
// TAST (Let): low_6_2 shape=Other bindingType=Boolean
low_6_2 := (j_5_1) < (int64(1))
_ = low_6_2
var __t7 uint32
{
if low_6_2 {
// TAST (Let): __local_var_7_5 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
__local_var_7_5 := Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_pred__4151667621(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil})))
_ = __local_var_7_5
var __t6 gopurs_runtime.Value
{
if (__local_var_7_5 == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (__local_var_7_5 != nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64((__local_var_7_5).V0), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = uint32(__t6.IntVal)
goto end_branch_7
} else {

}
}
{
__t7 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
}
end_branch_7:
// TAST (Let): __local_var_7_4 shape=Branch(Let(Branch(Other, Other, def=Other)), def=Other) bindingType=(ADT ["Data","Date","Component","Month"] [])
__local_var_7_4 := __t7
_ = __local_var_7_4
var __t9 int64
{
if (__local_var_7_4 == 1908470532) {
__t9 = int64(31)
goto end_branch_9
} else {

}
}
{
if (__local_var_7_4 == 2455627378) {
var __t8 int64
{
if ((((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) % (int64(4))) == (int64(0))) && (((((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) % (int64(400))) == (int64(0))) || (((((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) % (int64(100))) == (int64(0))) != (true))) {
__t8 = int64(29)
goto end_branch_8
} else {

}
}
{
__t8 = int64(28)
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
if (__local_var_7_4 == 4162469099) {
__t9 = int64(31)
goto end_branch_9
} else {

}
}
{
if (__local_var_7_4 == 1692989816) {
__t9 = int64(30)
goto end_branch_9
} else {

}
}
{
if (__local_var_7_4 == 330658827) {
__t9 = int64(31)
goto end_branch_9
} else {

}
}
{
if (__local_var_7_4 == 4067355978) {
__t9 = int64(30)
goto end_branch_9
} else {

}
}
{
if (__local_var_7_4 == 2276710548) {
__t9 = int64(31)
goto end_branch_9
} else {

}
}
{
if (__local_var_7_4 == 243771071) {
__t9 = int64(31)
goto end_branch_9
} else {

}
}
{
if (__local_var_7_4 == 215731793) {
__t9 = int64(30)
goto end_branch_9
} else {

}
}
{
if (__local_var_7_4 == 8639228) {
__t9 = int64(31)
goto end_branch_9
} else {

}
}
{
if (__local_var_7_4 == 49471444) {
__t9 = int64(30)
goto end_branch_9
} else {

}
}
{
if (__local_var_7_4 == 3889233761) {
__t9 = int64(31)
goto end_branch_9
} else {

}
}
{
__t9 = func() int64 { panic("Failed pattern match") }()
}
end_branch_9:
// TAST (Let): l_7_3 shape=Let(Branch(LitInt, Branch(LitInt, def=LitInt), LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, def=Other)) bindingType=Int
l_7_3 := __t9
_ = l_7_3
// TAST (Let): hi_8_10 shape=Other bindingType=Boolean
hi_8_10 := (j_5_1) > (l_7_3)
_ = hi_8_10
var __t12 int64
{
if low_6_2 {
__t12 = j_5_1
goto end_branch_12
} else {

}
}
{
if hi_8_10 {
__t12 = ((j_5_1) - (l_7_3)) - (int64(1))
goto end_branch_12
} else {

}
}
{
__t12 = int64(0)
}
end_branch_12:
// TAST (Let): __local_var_9_11 shape=App(Other) bindingType=Any
__local_var_9_11 := gopurs_runtime.Apply(adj_2_0_0, gopurs_runtime.Int(__t12))
_ = __local_var_9_11
var __t200 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if low_6_2 {
var __t65 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
// TAST (Let): __local_var_10_20 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_10_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0, (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1, int64(1)}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_10_20
// TAST (Let): pm_11_21 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
pm_11_21 := Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_pred__4151667621(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_10_20).V0.UnsafePtr).V1), UnsafePtr: nil})))
_ = pm_11_21
var __t23 gopurs_runtime.Value
{
if (pm_11_21 == nil) {
__t23 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_23
} else {

}
}
{
if (pm_11_21 != nil) {
__t23 = gopurs_runtime.Value{Type: 9, IntVal: int64((pm_11_21).V0), UnsafePtr: nil}
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
// TAST (Let): m_prime__12_22 shape=Branch(Other, Other, def=Other) bindingType=(ADT ["Data","Date","Component","Month"] [])
m_prime__12_22 := uint32(__t23.IntVal)
_ = m_prime__12_22
var __t26 int64
{
if (m_prime__12_22 == 1908470532) {
__t26 = int64(31)
goto end_branch_26
} else {

}
}
{
if (m_prime__12_22 == 2455627378) {
var __t25 int64
{
if ((((*Constructor_Data_Date_Date)((__local_var_10_20).V0.UnsafePtr).V0) % (int64(4))) == (int64(0))) && (((((*Constructor_Data_Date_Date)((__local_var_10_20).V0.UnsafePtr).V0) % (int64(400))) == (int64(0))) || (((((*Constructor_Data_Date_Date)((__local_var_10_20).V0.UnsafePtr).V0) % (int64(100))) == (int64(0))) != (true))) {
__t25 = int64(29)
goto end_branch_25
} else {

}
}
{
__t25 = int64(28)
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
if (m_prime__12_22 == 4162469099) {
__t26 = int64(31)
goto end_branch_26
} else {

}
}
{
if (m_prime__12_22 == 1692989816) {
__t26 = int64(30)
goto end_branch_26
} else {

}
}
{
if (m_prime__12_22 == 330658827) {
__t26 = int64(31)
goto end_branch_26
} else {

}
}
{
if (m_prime__12_22 == 4067355978) {
__t26 = int64(30)
goto end_branch_26
} else {

}
}
{
if (m_prime__12_22 == 2276710548) {
__t26 = int64(31)
goto end_branch_26
} else {

}
}
{
if (m_prime__12_22 == 243771071) {
__t26 = int64(31)
goto end_branch_26
} else {

}
}
{
if (m_prime__12_22 == 215731793) {
__t26 = int64(30)
goto end_branch_26
} else {

}
}
{
if (m_prime__12_22 == 8639228) {
__t26 = int64(31)
goto end_branch_26
} else {

}
}
{
if (m_prime__12_22 == 49471444) {
__t26 = int64(30)
goto end_branch_26
} else {

}
}
{
if (m_prime__12_22 == 3889233761) {
__t26 = int64(31)
goto end_branch_26
} else {

}
}
{
__t26 = func() int64 { panic("Failed pattern match") }()
}
end_branch_26:
// TAST (Let): l_13_24 shape=Branch(LitInt, Branch(LitInt, def=LitInt), LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, def=Other) bindingType=Int
l_13_24 := __t26
_ = l_13_24
var __t34 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t27 gopurs_runtime.Value
{
if (pm_11_21 == nil) {
__t27 = gopurs_runtime.Bool(true)
goto end_branch_27
} else {

}
}
{
if (pm_11_21 != nil) {
__t27 = gopurs_runtime.Bool(false)
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
if (__t27.IntVal) != (0) {
var __t33 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
// TAST (Let): __local_var_14_28 shape=Other bindingType=Int
__local_var_14_28 := (gopurs_runtime.Int((*Constructor_Data_Date_Date)((__local_var_10_20).V0.UnsafePtr).V0).IntVal) - (int64(1))
_ = __local_var_14_28
var __t29 gopurs_runtime.Value
{
if ((__local_var_14_28) >= (int64(1))) && ((__local_var_14_28) <= (int64(31))) {
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_14_28)}))}
goto end_branch_29
} else {

}
}
{
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_29:
var __t_tag_30 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t29))
if (__t_tag_30 != nil) {
// TAST (Let): __local_var_14_31 shape=Other bindingType=Int
__local_var_14_31 := (gopurs_runtime.Int((*Constructor_Data_Date_Date)((__local_var_10_20).V0.UnsafePtr).V0).IntVal) - (int64(1))
_ = __local_var_14_31
var __t32 gopurs_runtime.Value
{
if ((__local_var_14_31) >= (int64(1))) && ((__local_var_14_31) <= (int64(31))) {
__t32 = gopurs_runtime.Int(__local_var_14_31)
goto end_branch_32
} else {

}
}
{
__t32 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_32:
__t33 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, __t32.IntVal, m_prime__12_22, l_13_24}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_33
} else {

}
}
{
__t33 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_33:
__t34 = __t33
goto end_branch_34
} else {

}
}
{
__t34 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (*Constructor_Data_Date_Date)((__local_var_10_20).V0.UnsafePtr).V0, m_prime__12_22, l_13_24}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_34:
if (__t34 != nil) {
// TAST (Let): __local_var_10_35 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_10_35 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0, (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1, int64(1)}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_10_35
// TAST (Let): pm_11_36 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
pm_11_36 := Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_pred__4151667621(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_10_35).V0.UnsafePtr).V1), UnsafePtr: nil})))
_ = pm_11_36
var __t38 gopurs_runtime.Value
{
if (pm_11_36 == nil) {
__t38 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_38
} else {

}
}
{
if (pm_11_36 != nil) {
__t38 = gopurs_runtime.Value{Type: 9, IntVal: int64((pm_11_36).V0), UnsafePtr: nil}
goto end_branch_38
} else {

}
}
{
__t38 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_38:
// TAST (Let): m_prime__12_37 shape=Branch(Other, Other, def=Other) bindingType=(ADT ["Data","Date","Component","Month"] [])
m_prime__12_37 := uint32(__t38.IntVal)
_ = m_prime__12_37
var __t41 int64
{
if (m_prime__12_37 == 1908470532) {
__t41 = int64(31)
goto end_branch_41
} else {

}
}
{
if (m_prime__12_37 == 2455627378) {
var __t40 int64
{
if ((((*Constructor_Data_Date_Date)((__local_var_10_35).V0.UnsafePtr).V0) % (int64(4))) == (int64(0))) && (((((*Constructor_Data_Date_Date)((__local_var_10_35).V0.UnsafePtr).V0) % (int64(400))) == (int64(0))) || (((((*Constructor_Data_Date_Date)((__local_var_10_35).V0.UnsafePtr).V0) % (int64(100))) == (int64(0))) != (true))) {
__t40 = int64(29)
goto end_branch_40
} else {

}
}
{
__t40 = int64(28)
}
end_branch_40:
__t41 = __t40
goto end_branch_41
} else {

}
}
{
if (m_prime__12_37 == 4162469099) {
__t41 = int64(31)
goto end_branch_41
} else {

}
}
{
if (m_prime__12_37 == 1692989816) {
__t41 = int64(30)
goto end_branch_41
} else {

}
}
{
if (m_prime__12_37 == 330658827) {
__t41 = int64(31)
goto end_branch_41
} else {

}
}
{
if (m_prime__12_37 == 4067355978) {
__t41 = int64(30)
goto end_branch_41
} else {

}
}
{
if (m_prime__12_37 == 2276710548) {
__t41 = int64(31)
goto end_branch_41
} else {

}
}
{
if (m_prime__12_37 == 243771071) {
__t41 = int64(31)
goto end_branch_41
} else {

}
}
{
if (m_prime__12_37 == 215731793) {
__t41 = int64(30)
goto end_branch_41
} else {

}
}
{
if (m_prime__12_37 == 8639228) {
__t41 = int64(31)
goto end_branch_41
} else {

}
}
{
if (m_prime__12_37 == 49471444) {
__t41 = int64(30)
goto end_branch_41
} else {

}
}
{
if (m_prime__12_37 == 3889233761) {
__t41 = int64(31)
goto end_branch_41
} else {

}
}
{
__t41 = func() int64 { panic("Failed pattern match") }()
}
end_branch_41:
// TAST (Let): l_13_39 shape=Branch(LitInt, Branch(LitInt, def=LitInt), LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, def=Other) bindingType=Int
l_13_39 := __t41
_ = l_13_39
var __t49 gopurs_runtime.Value
{
var __t42 gopurs_runtime.Value
{
if (pm_11_36 == nil) {
__t42 = gopurs_runtime.Bool(true)
goto end_branch_42
} else {

}
}
{
if (pm_11_36 != nil) {
__t42 = gopurs_runtime.Bool(false)
goto end_branch_42
} else {

}
}
{
__t42 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_42:
if (__t42.IntVal) != (0) {
var __t48 gopurs_runtime.Value
{
// TAST (Let): __local_var_14_43 shape=Other bindingType=Int
__local_var_14_43 := (gopurs_runtime.Int((*Constructor_Data_Date_Date)((__local_var_10_35).V0.UnsafePtr).V0).IntVal) - (int64(1))
_ = __local_var_14_43
var __t44 gopurs_runtime.Value
{
if ((__local_var_14_43) >= (int64(1))) && ((__local_var_14_43) <= (int64(31))) {
__t44 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_14_43)}))}
goto end_branch_44
} else {

}
}
{
__t44 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_44:
var __t_tag_45 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t44))
if (__t_tag_45 != nil) {
// TAST (Let): __local_var_14_46 shape=Other bindingType=Int
__local_var_14_46 := (gopurs_runtime.Int((*Constructor_Data_Date_Date)((__local_var_10_35).V0.UnsafePtr).V0).IntVal) - (int64(1))
_ = __local_var_14_46
var __t47 gopurs_runtime.Value
{
if ((__local_var_14_46) >= (int64(1))) && ((__local_var_14_46) <= (int64(31))) {
__t47 = gopurs_runtime.Int(__local_var_14_46)
goto end_branch_47
} else {

}
}
{
__t47 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_47:
__t48 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, __t47.IntVal, m_prime__12_37, l_13_39}))}
goto end_branch_48
} else {

}
}
{
__t48 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_48:
__t49 = __t48
goto end_branch_49
} else {

}
}
{
__t49 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (*Constructor_Data_Date_Date)((__local_var_10_35).V0.UnsafePtr).V0, m_prime__12_37, l_13_39}))}
}
end_branch_49:
__t65 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_9_11, __t49))
goto end_branch_65
} else {

}
}
{
// TAST (Let): __local_var_10_50 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_10_50 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0, (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1, int64(1)}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_10_50
// TAST (Let): pm_11_51 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
pm_11_51 := Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_pred__4151667621(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_10_50).V0.UnsafePtr).V1), UnsafePtr: nil})))
_ = pm_11_51
var __t53 gopurs_runtime.Value
{
if (pm_11_51 == nil) {
__t53 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_53
} else {

}
}
{
if (pm_11_51 != nil) {
__t53 = gopurs_runtime.Value{Type: 9, IntVal: int64((pm_11_51).V0), UnsafePtr: nil}
goto end_branch_53
} else {

}
}
{
__t53 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_53:
// TAST (Let): m_prime__12_52 shape=Branch(Other, Other, def=Other) bindingType=(ADT ["Data","Date","Component","Month"] [])
m_prime__12_52 := uint32(__t53.IntVal)
_ = m_prime__12_52
var __t56 int64
{
if (m_prime__12_52 == 1908470532) {
__t56 = int64(31)
goto end_branch_56
} else {

}
}
{
if (m_prime__12_52 == 2455627378) {
var __t55 int64
{
if ((((*Constructor_Data_Date_Date)((__local_var_10_50).V0.UnsafePtr).V0) % (int64(4))) == (int64(0))) && (((((*Constructor_Data_Date_Date)((__local_var_10_50).V0.UnsafePtr).V0) % (int64(400))) == (int64(0))) || (((((*Constructor_Data_Date_Date)((__local_var_10_50).V0.UnsafePtr).V0) % (int64(100))) == (int64(0))) != (true))) {
__t55 = int64(29)
goto end_branch_55
} else {

}
}
{
__t55 = int64(28)
}
end_branch_55:
__t56 = __t55
goto end_branch_56
} else {

}
}
{
if (m_prime__12_52 == 4162469099) {
__t56 = int64(31)
goto end_branch_56
} else {

}
}
{
if (m_prime__12_52 == 1692989816) {
__t56 = int64(30)
goto end_branch_56
} else {

}
}
{
if (m_prime__12_52 == 330658827) {
__t56 = int64(31)
goto end_branch_56
} else {

}
}
{
if (m_prime__12_52 == 4067355978) {
__t56 = int64(30)
goto end_branch_56
} else {

}
}
{
if (m_prime__12_52 == 2276710548) {
__t56 = int64(31)
goto end_branch_56
} else {

}
}
{
if (m_prime__12_52 == 243771071) {
__t56 = int64(31)
goto end_branch_56
} else {

}
}
{
if (m_prime__12_52 == 215731793) {
__t56 = int64(30)
goto end_branch_56
} else {

}
}
{
if (m_prime__12_52 == 8639228) {
__t56 = int64(31)
goto end_branch_56
} else {

}
}
{
if (m_prime__12_52 == 49471444) {
__t56 = int64(30)
goto end_branch_56
} else {

}
}
{
if (m_prime__12_52 == 3889233761) {
__t56 = int64(31)
goto end_branch_56
} else {

}
}
{
__t56 = func() int64 { panic("Failed pattern match") }()
}
end_branch_56:
// TAST (Let): l_13_54 shape=Branch(LitInt, Branch(LitInt, def=LitInt), LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, LitInt, def=Other) bindingType=Int
l_13_54 := __t56
_ = l_13_54
var __t64 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t57 gopurs_runtime.Value
{
if (pm_11_51 == nil) {
__t57 = gopurs_runtime.Bool(true)
goto end_branch_57
} else {

}
}
{
if (pm_11_51 != nil) {
__t57 = gopurs_runtime.Bool(false)
goto end_branch_57
} else {

}
}
{
__t57 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_57:
if (__t57.IntVal) != (0) {
var __t63 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
// TAST (Let): __local_var_14_58 shape=Other bindingType=Int
__local_var_14_58 := (gopurs_runtime.Int((*Constructor_Data_Date_Date)((__local_var_10_50).V0.UnsafePtr).V0).IntVal) - (int64(1))
_ = __local_var_14_58
var __t59 gopurs_runtime.Value
{
if ((__local_var_14_58) >= (int64(1))) && ((__local_var_14_58) <= (int64(31))) {
__t59 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_14_58)}))}
goto end_branch_59
} else {

}
}
{
__t59 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_59:
var __t_tag_60 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t59))
if (__t_tag_60 != nil) {
// TAST (Let): __local_var_14_61 shape=Other bindingType=Int
__local_var_14_61 := (gopurs_runtime.Int((*Constructor_Data_Date_Date)((__local_var_10_50).V0.UnsafePtr).V0).IntVal) - (int64(1))
_ = __local_var_14_61
var __t62 gopurs_runtime.Value
{
if ((__local_var_14_61) >= (int64(1))) && ((__local_var_14_61) <= (int64(31))) {
__t62 = gopurs_runtime.Int(__local_var_14_61)
goto end_branch_62
} else {

}
}
{
__t62 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_62:
__t63 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, __t62.IntVal, m_prime__12_52, l_13_54}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_63
} else {

}
}
{
__t63 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_63:
__t64 = __t63
goto end_branch_64
} else {

}
}
{
__t64 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (*Constructor_Data_Date_Date)((__local_var_10_50).V0.UnsafePtr).V0, m_prime__12_52, l_13_54}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_64:
if (__t64 == nil) {
__t65 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_65
} else {

}
}
{
__t65 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_65:
__t200 = __t65
goto end_branch_200
} else {

}
}
{
if hi_8_10 {
var __t199 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
// TAST (Let): sm_10_66 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
sm_10_66 := Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_succ__4151667621(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil})))
_ = sm_10_66
// TAST (Let): __local_var_11_68 shape=Other bindingType=Int
__local_var_11_68 := (l_7_3) + (int64(1))
_ = __local_var_11_68
var __t69 gopurs_runtime.Value
{
if ((__local_var_11_68) >= (int64(1))) && ((__local_var_11_68) <= (int64(31))) {
__t69 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_11_68)}))}
goto end_branch_69
} else {

}
}
{
__t69 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_69:
// TAST (Let): v1_11_67 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
v1_11_67 := Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t69))
_ = v1_11_67
var __t86 *Constructor_Data_Maybe_Just[int64]
{
var __t84 int64
{
var __t_tag_71 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_71) == 1908470532) {
__t84 = int64(31)
goto end_branch_84
} else {

}
}
{
var __t_tag_72 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_72) == 2455627378) {
var __t73 int64
{
if ((((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) % (int64(4))) == (int64(0))) && (((((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) % (int64(400))) == (int64(0))) || (((((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) % (int64(100))) == (int64(0))) != (true))) {
__t73 = int64(29)
goto end_branch_73
} else {

}
}
{
__t73 = int64(28)
}
end_branch_73:
__t84 = __t73
goto end_branch_84
} else {

}
}
{
var __t_tag_74 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_74) == 4162469099) {
__t84 = int64(31)
goto end_branch_84
} else {

}
}
{
var __t_tag_75 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_75) == 1692989816) {
__t84 = int64(30)
goto end_branch_84
} else {

}
}
{
var __t_tag_76 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_76) == 330658827) {
__t84 = int64(31)
goto end_branch_84
} else {

}
}
{
var __t_tag_77 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_77) == 4067355978) {
__t84 = int64(30)
goto end_branch_84
} else {

}
}
{
var __t_tag_78 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_78) == 2276710548) {
__t84 = int64(31)
goto end_branch_84
} else {

}
}
{
var __t_tag_79 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_79) == 243771071) {
__t84 = int64(31)
goto end_branch_84
} else {

}
}
{
var __t_tag_80 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_80) == 215731793) {
__t84 = int64(30)
goto end_branch_84
} else {

}
}
{
var __t_tag_81 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_81) == 8639228) {
__t84 = int64(31)
goto end_branch_84
} else {

}
}
{
var __t_tag_82 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_82) == 49471444) {
__t84 = int64(30)
goto end_branch_84
} else {

}
}
{
var __t_tag_83 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_83) == 3889233761) {
__t84 = int64(31)
goto end_branch_84
} else {

}
}
{
__t84 = func() int64 { panic("Failed pattern match") }()
}
end_branch_84:
var __t_tag_85 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[int64]]](Get_Data_Interval_ordMaybe()).V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(v1_11_67))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__t84)}))})
if (uint32(__t_tag_85.IntVal) == 380165415) {
__t86 = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
goto end_branch_86
} else {

}
}
{
__t86 = v1_11_67
}
end_branch_86:
// TAST (Let): sd_12_70 shape=Branch(Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
sd_12_70 := __t86
_ = sd_12_70
var __t110 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t94 gopurs_runtime.Value
{
if (sd_12_70 == nil) {
__t94 = gopurs_runtime.Bool(true)
goto end_branch_94
} else {

}
}
{
if (sd_12_70 != nil) {
__t94 = gopurs_runtime.Bool(false)
goto end_branch_94
} else {

}
}
{
__t94 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_94:
var __t_and_96 bool = false
if (__t94.IntVal) != (0) {

var __t95 gopurs_runtime.Value
{
if (sm_10_66 == nil) {
__t95 = gopurs_runtime.Bool(true)
goto end_branch_95
} else {

}
}
{
if (sm_10_66 != nil) {
__t95 = gopurs_runtime.Bool(false)
goto end_branch_95
} else {

}
}
{
__t95 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_95:
__t_and_96 = (__t95.IntVal) != (0)
}
if __t_and_96 {
var __t109 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
// TAST (Let): __local_var_13_97 shape=Other bindingType=Int
__local_var_13_97 := ((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) + (int64(1))
_ = __local_var_13_97
var __t98 gopurs_runtime.Value
{
if ((__local_var_13_97) >= (int64(1))) && ((__local_var_13_97) <= (int64(31))) {
__t98 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_13_97)}))}
goto end_branch_98
} else {

}
}
{
__t98 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_98:
var __t_tag_99 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t98))
if (__t_tag_99 != nil) {
// TAST (Let): __local_var_13_101 shape=Other bindingType=Int
__local_var_13_101 := ((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) + (int64(1))
_ = __local_var_13_101
var __t102 gopurs_runtime.Value
{
if ((__local_var_13_101) >= (int64(1))) && ((__local_var_13_101) <= (int64(31))) {
__t102 = gopurs_runtime.Int(__local_var_13_101)
goto end_branch_102
} else {

}
}
{
__t102 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_102:
var __t105 uint32
{
var __t103 gopurs_runtime.Value
{
if (sd_12_70 == nil) {
__t103 = gopurs_runtime.Bool(true)
goto end_branch_103
} else {

}
}
{
if (sd_12_70 != nil) {
__t103 = gopurs_runtime.Bool(false)
goto end_branch_103
} else {

}
}
{
__t103 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_103:
if (__t103.IntVal) != (0) {
var __t104 gopurs_runtime.Value
{
if (sm_10_66 == nil) {
__t104 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_104
} else {

}
}
{
if (sm_10_66 != nil) {
__t104 = gopurs_runtime.Value{Type: 9, IntVal: int64((sm_10_66).V0), UnsafePtr: nil}
goto end_branch_104
} else {

}
}
{
__t104 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_104:
__t105 = uint32(__t104.IntVal)
goto end_branch_105
} else {

}
}
{
__t105 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
}
end_branch_105:
// TAST (Let): __local_var_13_100 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_13_100 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), __t102, gopurs_runtime.Value{Type: 9, IntVal: int64(__t105), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_13_100
var __t108 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t107 gopurs_runtime.Value
{
if (sd_12_70 == nil) {
__t107 = gopurs_runtime.Bool(true)
goto end_branch_107
} else {

}
}
{
if (sd_12_70 != nil) {
__t107 = gopurs_runtime.Bool(false)
goto end_branch_107
} else {

}
}
{
__t107 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_107:
if (__t107.IntVal) != (0) {
__t108 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_13_100).V0, gopurs_runtime.Int(int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_108
} else {

}
}
{
var __t106 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (sd_12_70 != nil) {
__t106 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_13_100).V0, gopurs_runtime.Int((sd_12_70).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_106
} else {

}
}
{
__t106 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_106:
__t108 = __t106
}
end_branch_108:
__t109 = __t108
goto end_branch_109
} else {

}
}
{
__t109 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_109:
__t110 = __t109
goto end_branch_110
} else {

}
}
{
var __t90 uint32
{
var __t88 gopurs_runtime.Value
{
if (sd_12_70 == nil) {
__t88 = gopurs_runtime.Bool(true)
goto end_branch_88
} else {

}
}
{
if (sd_12_70 != nil) {
__t88 = gopurs_runtime.Bool(false)
goto end_branch_88
} else {

}
}
{
__t88 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_88:
if (__t88.IntVal) != (0) {
var __t89 gopurs_runtime.Value
{
if (sm_10_66 == nil) {
__t89 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_89
} else {

}
}
{
if (sm_10_66 != nil) {
__t89 = gopurs_runtime.Value{Type: 9, IntVal: int64((sm_10_66).V0), UnsafePtr: nil}
goto end_branch_89
} else {

}
}
{
__t89 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_89:
__t90 = uint32(__t89.IntVal)
goto end_branch_90
} else {

}
}
{
__t90 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
}
end_branch_90:
// TAST (Let): __local_var_13_87 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_13_87 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(__t90), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_13_87
var __t93 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t92 gopurs_runtime.Value
{
if (sd_12_70 == nil) {
__t92 = gopurs_runtime.Bool(true)
goto end_branch_92
} else {

}
}
{
if (sd_12_70 != nil) {
__t92 = gopurs_runtime.Bool(false)
goto end_branch_92
} else {

}
}
{
__t92 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_92:
if (__t92.IntVal) != (0) {
__t93 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_13_87).V0, gopurs_runtime.Int(int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_93
} else {

}
}
{
var __t91 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (sd_12_70 != nil) {
__t91 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_13_87).V0, gopurs_runtime.Int((sd_12_70).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_91
} else {

}
}
{
__t91 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_91:
__t93 = __t91
}
end_branch_93:
__t110 = __t93
}
end_branch_110:
if (__t110 != nil) {
// TAST (Let): sm_10_111 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
sm_10_111 := Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_succ__4151667621(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil})))
_ = sm_10_111
// TAST (Let): __local_var_11_113 shape=Other bindingType=Int
__local_var_11_113 := (l_7_3) + (int64(1))
_ = __local_var_11_113
var __t114 gopurs_runtime.Value
{
if ((__local_var_11_113) >= (int64(1))) && ((__local_var_11_113) <= (int64(31))) {
__t114 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_11_113)}))}
goto end_branch_114
} else {

}
}
{
__t114 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_114:
// TAST (Let): v1_11_112 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
v1_11_112 := Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t114))
_ = v1_11_112
var __t131 *Constructor_Data_Maybe_Just[int64]
{
var __t129 int64
{
var __t_tag_116 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_116) == 1908470532) {
__t129 = int64(31)
goto end_branch_129
} else {

}
}
{
var __t_tag_117 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_117) == 2455627378) {
var __t118 int64
{
if ((((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) % (int64(4))) == (int64(0))) && (((((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) % (int64(400))) == (int64(0))) || (((((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) % (int64(100))) == (int64(0))) != (true))) {
__t118 = int64(29)
goto end_branch_118
} else {

}
}
{
__t118 = int64(28)
}
end_branch_118:
__t129 = __t118
goto end_branch_129
} else {

}
}
{
var __t_tag_119 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_119) == 4162469099) {
__t129 = int64(31)
goto end_branch_129
} else {

}
}
{
var __t_tag_120 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_120) == 1692989816) {
__t129 = int64(30)
goto end_branch_129
} else {

}
}
{
var __t_tag_121 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_121) == 330658827) {
__t129 = int64(31)
goto end_branch_129
} else {

}
}
{
var __t_tag_122 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_122) == 4067355978) {
__t129 = int64(30)
goto end_branch_129
} else {

}
}
{
var __t_tag_123 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_123) == 2276710548) {
__t129 = int64(31)
goto end_branch_129
} else {

}
}
{
var __t_tag_124 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_124) == 243771071) {
__t129 = int64(31)
goto end_branch_129
} else {

}
}
{
var __t_tag_125 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_125) == 215731793) {
__t129 = int64(30)
goto end_branch_129
} else {

}
}
{
var __t_tag_126 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_126) == 8639228) {
__t129 = int64(31)
goto end_branch_129
} else {

}
}
{
var __t_tag_127 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_127) == 49471444) {
__t129 = int64(30)
goto end_branch_129
} else {

}
}
{
var __t_tag_128 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_128) == 3889233761) {
__t129 = int64(31)
goto end_branch_129
} else {

}
}
{
__t129 = func() int64 { panic("Failed pattern match") }()
}
end_branch_129:
var __t_tag_130 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[int64]]](Get_Data_Interval_ordMaybe()).V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(v1_11_112))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__t129)}))})
if (uint32(__t_tag_130.IntVal) == 380165415) {
__t131 = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
goto end_branch_131
} else {

}
}
{
__t131 = v1_11_112
}
end_branch_131:
// TAST (Let): sd_12_115 shape=Branch(Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
sd_12_115 := __t131
_ = sd_12_115
var __t153 gopurs_runtime.Value
{
var __t138 gopurs_runtime.Value
{
if (sd_12_115 == nil) {
__t138 = gopurs_runtime.Bool(true)
goto end_branch_138
} else {

}
}
{
if (sd_12_115 != nil) {
__t138 = gopurs_runtime.Bool(false)
goto end_branch_138
} else {

}
}
{
__t138 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_138:
var __t_and_140 bool = false
if (__t138.IntVal) != (0) {

var __t139 gopurs_runtime.Value
{
if (sm_10_111 == nil) {
__t139 = gopurs_runtime.Bool(true)
goto end_branch_139
} else {

}
}
{
if (sm_10_111 != nil) {
__t139 = gopurs_runtime.Bool(false)
goto end_branch_139
} else {

}
}
{
__t139 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_139:
__t_and_140 = (__t139.IntVal) != (0)
}
if __t_and_140 {
var __t152 gopurs_runtime.Value
{
// TAST (Let): __local_var_13_141 shape=Other bindingType=Int
__local_var_13_141 := ((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) + (int64(1))
_ = __local_var_13_141
var __t142 gopurs_runtime.Value
{
if ((__local_var_13_141) >= (int64(1))) && ((__local_var_13_141) <= (int64(31))) {
__t142 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_13_141)}))}
goto end_branch_142
} else {

}
}
{
__t142 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_142:
var __t_tag_143 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t142))
if (__t_tag_143 != nil) {
// TAST (Let): __local_var_13_145 shape=Other bindingType=Int
__local_var_13_145 := ((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) + (int64(1))
_ = __local_var_13_145
var __t146 gopurs_runtime.Value
{
if ((__local_var_13_145) >= (int64(1))) && ((__local_var_13_145) <= (int64(31))) {
__t146 = gopurs_runtime.Int(__local_var_13_145)
goto end_branch_146
} else {

}
}
{
__t146 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_146:
var __t149 uint32
{
var __t147 gopurs_runtime.Value
{
if (sd_12_115 == nil) {
__t147 = gopurs_runtime.Bool(true)
goto end_branch_147
} else {

}
}
{
if (sd_12_115 != nil) {
__t147 = gopurs_runtime.Bool(false)
goto end_branch_147
} else {

}
}
{
__t147 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_147:
if (__t147.IntVal) != (0) {
var __t148 gopurs_runtime.Value
{
if (sm_10_111 == nil) {
__t148 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_148
} else {

}
}
{
if (sm_10_111 != nil) {
__t148 = gopurs_runtime.Value{Type: 9, IntVal: int64((sm_10_111).V0), UnsafePtr: nil}
goto end_branch_148
} else {

}
}
{
__t148 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_148:
__t149 = uint32(__t148.IntVal)
goto end_branch_149
} else {

}
}
{
__t149 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
}
end_branch_149:
// TAST (Let): __local_var_13_144 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_13_144 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), __t146, gopurs_runtime.Value{Type: 9, IntVal: int64(__t149), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_13_144
var __t151 gopurs_runtime.Value
{
var __t150 gopurs_runtime.Value
{
if (sd_12_115 == nil) {
__t150 = gopurs_runtime.Bool(true)
goto end_branch_150
} else {

}
}
{
if (sd_12_115 != nil) {
__t150 = gopurs_runtime.Bool(false)
goto end_branch_150
} else {

}
}
{
__t150 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_150:
if (__t150.IntVal) != (0) {
__t151 = gopurs_runtime.Apply((__local_var_13_144).V0, gopurs_runtime.Int(int64(1)))
goto end_branch_151
} else {

}
}
{
if (sd_12_115 != nil) {
__t151 = gopurs_runtime.Apply((__local_var_13_144).V0, gopurs_runtime.Int((sd_12_115).V0))
goto end_branch_151
} else {

}
}
{
__t151 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_151:
__t152 = __t151
goto end_branch_152
} else {

}
}
{
__t152 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_152:
__t153 = __t152
goto end_branch_153
} else {

}
}
{
var __t135 uint32
{
var __t133 gopurs_runtime.Value
{
if (sd_12_115 == nil) {
__t133 = gopurs_runtime.Bool(true)
goto end_branch_133
} else {

}
}
{
if (sd_12_115 != nil) {
__t133 = gopurs_runtime.Bool(false)
goto end_branch_133
} else {

}
}
{
__t133 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_133:
if (__t133.IntVal) != (0) {
var __t134 gopurs_runtime.Value
{
if (sm_10_111 == nil) {
__t134 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_134
} else {

}
}
{
if (sm_10_111 != nil) {
__t134 = gopurs_runtime.Value{Type: 9, IntVal: int64((sm_10_111).V0), UnsafePtr: nil}
goto end_branch_134
} else {

}
}
{
__t134 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_134:
__t135 = uint32(__t134.IntVal)
goto end_branch_135
} else {

}
}
{
__t135 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
}
end_branch_135:
// TAST (Let): __local_var_13_132 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_13_132 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(__t135), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_13_132
var __t137 gopurs_runtime.Value
{
var __t136 gopurs_runtime.Value
{
if (sd_12_115 == nil) {
__t136 = gopurs_runtime.Bool(true)
goto end_branch_136
} else {

}
}
{
if (sd_12_115 != nil) {
__t136 = gopurs_runtime.Bool(false)
goto end_branch_136
} else {

}
}
{
__t136 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_136:
if (__t136.IntVal) != (0) {
__t137 = gopurs_runtime.Apply((__local_var_13_132).V0, gopurs_runtime.Int(int64(1)))
goto end_branch_137
} else {

}
}
{
if (sd_12_115 != nil) {
__t137 = gopurs_runtime.Apply((__local_var_13_132).V0, gopurs_runtime.Int((sd_12_115).V0))
goto end_branch_137
} else {

}
}
{
__t137 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_137:
__t153 = __t137
}
end_branch_153:
__t199 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_9_11, __t153))
goto end_branch_199
} else {

}
}
{
// TAST (Let): sm_10_154 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
sm_10_154 := Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_succ__4151667621(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil})))
_ = sm_10_154
// TAST (Let): __local_var_11_156 shape=Other bindingType=Int
__local_var_11_156 := (l_7_3) + (int64(1))
_ = __local_var_11_156
var __t157 gopurs_runtime.Value
{
if ((__local_var_11_156) >= (int64(1))) && ((__local_var_11_156) <= (int64(31))) {
__t157 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_11_156)}))}
goto end_branch_157
} else {

}
}
{
__t157 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_157:
// TAST (Let): v1_11_155 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
v1_11_155 := Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t157))
_ = v1_11_155
var __t174 *Constructor_Data_Maybe_Just[int64]
{
var __t172 int64
{
var __t_tag_159 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_159) == 1908470532) {
__t172 = int64(31)
goto end_branch_172
} else {

}
}
{
var __t_tag_160 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_160) == 2455627378) {
var __t161 int64
{
if ((((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) % (int64(4))) == (int64(0))) && (((((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) % (int64(400))) == (int64(0))) || (((((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) % (int64(100))) == (int64(0))) != (true))) {
__t161 = int64(29)
goto end_branch_161
} else {

}
}
{
__t161 = int64(28)
}
end_branch_161:
__t172 = __t161
goto end_branch_172
} else {

}
}
{
var __t_tag_162 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_162) == 4162469099) {
__t172 = int64(31)
goto end_branch_172
} else {

}
}
{
var __t_tag_163 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_163) == 1692989816) {
__t172 = int64(30)
goto end_branch_172
} else {

}
}
{
var __t_tag_164 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_164) == 330658827) {
__t172 = int64(31)
goto end_branch_172
} else {

}
}
{
var __t_tag_165 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_165) == 4067355978) {
__t172 = int64(30)
goto end_branch_172
} else {

}
}
{
var __t_tag_166 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_166) == 2276710548) {
__t172 = int64(31)
goto end_branch_172
} else {

}
}
{
var __t_tag_167 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_167) == 243771071) {
__t172 = int64(31)
goto end_branch_172
} else {

}
}
{
var __t_tag_168 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_168) == 215731793) {
__t172 = int64(30)
goto end_branch_172
} else {

}
}
{
var __t_tag_169 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_169) == 8639228) {
__t172 = int64(31)
goto end_branch_172
} else {

}
}
{
var __t_tag_170 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_170) == 49471444) {
__t172 = int64(30)
goto end_branch_172
} else {

}
}
{
var __t_tag_171 uint32 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
if (uint32(__t_tag_171) == 3889233761) {
__t172 = int64(31)
goto end_branch_172
} else {

}
}
{
__t172 = func() int64 { panic("Failed pattern match") }()
}
end_branch_172:
var __t_tag_173 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[int64]]](Get_Data_Interval_ordMaybe()).V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(v1_11_155))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__t172)}))})
if (uint32(__t_tag_173.IntVal) == 380165415) {
__t174 = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
goto end_branch_174
} else {

}
}
{
__t174 = v1_11_155
}
end_branch_174:
// TAST (Let): sd_12_158 shape=Branch(Other, def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
sd_12_158 := __t174
_ = sd_12_158
var __t198 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t182 gopurs_runtime.Value
{
if (sd_12_158 == nil) {
__t182 = gopurs_runtime.Bool(true)
goto end_branch_182
} else {

}
}
{
if (sd_12_158 != nil) {
__t182 = gopurs_runtime.Bool(false)
goto end_branch_182
} else {

}
}
{
__t182 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_182:
var __t_and_184 bool = false
if (__t182.IntVal) != (0) {

var __t183 gopurs_runtime.Value
{
if (sm_10_154 == nil) {
__t183 = gopurs_runtime.Bool(true)
goto end_branch_183
} else {

}
}
{
if (sm_10_154 != nil) {
__t183 = gopurs_runtime.Bool(false)
goto end_branch_183
} else {

}
}
{
__t183 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_183:
__t_and_184 = (__t183.IntVal) != (0)
}
if __t_and_184 {
var __t197 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
// TAST (Let): __local_var_13_185 shape=Other bindingType=Int
__local_var_13_185 := ((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) + (int64(1))
_ = __local_var_13_185
var __t186 gopurs_runtime.Value
{
if ((__local_var_13_185) >= (int64(1))) && ((__local_var_13_185) <= (int64(31))) {
__t186 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(__local_var_13_185)}))}
goto end_branch_186
} else {

}
}
{
__t186 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_186:
var __t_tag_187 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t186))
if (__t_tag_187 != nil) {
// TAST (Let): __local_var_13_189 shape=Other bindingType=Int
__local_var_13_189 := ((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0) + (int64(1))
_ = __local_var_13_189
var __t190 gopurs_runtime.Value
{
if ((__local_var_13_189) >= (int64(1))) && ((__local_var_13_189) <= (int64(31))) {
__t190 = gopurs_runtime.Int(__local_var_13_189)
goto end_branch_190
} else {

}
}
{
__t190 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_190:
var __t193 uint32
{
var __t191 gopurs_runtime.Value
{
if (sd_12_158 == nil) {
__t191 = gopurs_runtime.Bool(true)
goto end_branch_191
} else {

}
}
{
if (sd_12_158 != nil) {
__t191 = gopurs_runtime.Bool(false)
goto end_branch_191
} else {

}
}
{
__t191 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_191:
if (__t191.IntVal) != (0) {
var __t192 gopurs_runtime.Value
{
if (sm_10_154 == nil) {
__t192 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_192
} else {

}
}
{
if (sm_10_154 != nil) {
__t192 = gopurs_runtime.Value{Type: 9, IntVal: int64((sm_10_154).V0), UnsafePtr: nil}
goto end_branch_192
} else {

}
}
{
__t192 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_192:
__t193 = uint32(__t192.IntVal)
goto end_branch_193
} else {

}
}
{
__t193 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
}
end_branch_193:
// TAST (Let): __local_var_13_188 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_13_188 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), __t190, gopurs_runtime.Value{Type: 9, IntVal: int64(__t193), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_13_188
var __t196 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t195 gopurs_runtime.Value
{
if (sd_12_158 == nil) {
__t195 = gopurs_runtime.Bool(true)
goto end_branch_195
} else {

}
}
{
if (sd_12_158 != nil) {
__t195 = gopurs_runtime.Bool(false)
goto end_branch_195
} else {

}
}
{
__t195 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_195:
if (__t195.IntVal) != (0) {
__t196 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_13_188).V0, gopurs_runtime.Int(int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_196
} else {

}
}
{
var __t194 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (sd_12_158 != nil) {
__t194 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_13_188).V0, gopurs_runtime.Int((sd_12_158).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_194
} else {

}
}
{
__t194 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_194:
__t196 = __t194
}
end_branch_196:
__t197 = __t196
goto end_branch_197
} else {

}
}
{
__t197 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_197:
__t198 = __t197
goto end_branch_198
} else {

}
}
{
var __t178 uint32
{
var __t176 gopurs_runtime.Value
{
if (sd_12_158 == nil) {
__t176 = gopurs_runtime.Bool(true)
goto end_branch_176
} else {

}
}
{
if (sd_12_158 != nil) {
__t176 = gopurs_runtime.Bool(false)
goto end_branch_176
} else {

}
}
{
__t176 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_176:
if (__t176.IntVal) != (0) {
var __t177 gopurs_runtime.Value
{
if (sm_10_154 == nil) {
__t177 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_177
} else {

}
}
{
if (sm_10_154 != nil) {
__t177 = gopurs_runtime.Value{Type: 9, IntVal: int64((sm_10_154).V0), UnsafePtr: nil}
goto end_branch_177
} else {

}
}
{
__t177 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_177:
__t178 = uint32(__t177.IntVal)
goto end_branch_178
} else {

}
}
{
__t178 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
}
end_branch_178:
// TAST (Let): __local_var_13_175 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_13_175 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(__t178), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_13_175
var __t181 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t180 gopurs_runtime.Value
{
if (sd_12_158 == nil) {
__t180 = gopurs_runtime.Bool(true)
goto end_branch_180
} else {

}
}
{
if (sd_12_158 != nil) {
__t180 = gopurs_runtime.Bool(false)
goto end_branch_180
} else {

}
}
{
__t180 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_180:
if (__t180.IntVal) != (0) {
__t181 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_13_175).V0, gopurs_runtime.Int(int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_181
} else {

}
}
{
var __t179 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (sd_12_158 != nil) {
__t179 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_13_175).V0, gopurs_runtime.Int((sd_12_158).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_179
} else {

}
}
{
__t179 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_179:
__t181 = __t179
}
end_branch_181:
__t198 = __t181
}
end_branch_198:
if (__t198 == nil) {
__t199 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_199
} else {

}
}
{
__t199 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_199:
__t200 = __t199
goto end_branch_200
} else {

}
}
{
var __t19 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
// TAST (Let): __local_var_10_13 shape=App(Var) bindingType=Any
__local_var_10_13 := gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil})
_ = __local_var_10_13
var __t14 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((j_5_1) >= (int64(1))) && ((j_5_1) <= (int64(31))) {
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(__local_var_10_13, gopurs_runtime.Int(j_5_1)), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
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
if (__t14 != nil) {
// TAST (Let): __local_var_10_15 shape=App(Var) bindingType=Any
__local_var_10_15 := gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil})
_ = __local_var_10_15
var __t16 gopurs_runtime.Value
{
if ((j_5_1) >= (int64(1))) && ((j_5_1) <= (int64(31))) {
__t16 = gopurs_runtime.Apply(__local_var_10_15, gopurs_runtime.Int(j_5_1))
goto end_branch_16
} else {

}
}
{
__t16 = ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0
}
end_branch_16:
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_9_11, __t16))
goto end_branch_19
} else {

}
}
{
// TAST (Let): __local_var_10_17 shape=App(Var) bindingType=Any
__local_var_10_17 := gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil})
_ = __local_var_10_17
var __t18 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((j_5_1) >= (int64(1))) && ((j_5_1) <= (int64(31))) {
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(__local_var_10_17, gopurs_runtime.Int(j_5_1)), true}
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
if (__t18 == nil) {
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
__t19 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_19:
__t200 = __t19
}
end_branch_200:
__t201 = __t200
}
end_branch_201:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t201)}
})
// TAST (Let): __local_var_3_202 shape=App(Var) bindingType=Any
__local_var_3_202 := gopurs_runtime.Apply(Get_Data_Int_fromNumber(), gopurs_runtime.Float(v_0))
_ = __local_var_3_202
var __t203 gopurs_runtime.Value
{
if (__local_var_3_202.Type == 9 && __local_var_3_202.IntVal == 930809136 && __local_var_3_202.UnsafePtr != nil) {
__t203 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_2280409795_3094389156(Rebox_Data_Date_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(adj_2_0_0, gopurs_runtime.Int((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_202.UnsafePtr).V0.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(date_1)})))))}
goto end_branch_203
} else {

}
}
{
if (__local_var_3_202.Type == 9 && __local_var_3_202.IntVal == 930809136 && __local_var_3_202.UnsafePtr == nil) {
__t203 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
goto end_branch_203
} else {

}
}
{
__t203 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
}
end_branch_203:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t203))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Rebox_Data_Date_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Date_1726078761_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Date_Date]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Date_1957390985_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Date_Date]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Date_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Date_3092443796_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[int64]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Date_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Date_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}

func Rebox_Data_Date_3094389156_622082505(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[uint32]{}
		out.V0 = uint32(in.V0.IntVal)
	return out
}

func Rebox_Data_Date_3494563625_556578094(in *Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Date_3508461103_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Date_3523628265_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_Date_Date]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V1)}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_Date_758368489_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Get_Data_Date_calcDiff() gopurs_runtime.Value {
	return _Gopurs_Data_Date_CalcDiff
}

func Get_Data_Date_calcWeekday() gopurs_runtime.Value {
	return _Gopurs_Data_Date_CalcWeekday
}

func Get_Data_Date_canonicalDateImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Date_CanonicalDateImpl
}

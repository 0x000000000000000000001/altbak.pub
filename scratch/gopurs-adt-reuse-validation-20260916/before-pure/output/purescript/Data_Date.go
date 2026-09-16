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
		cache_Data_Date_toEnum = Call_Data_Enum_toEnum(Rebox_Data_Date_1306125126_123048125(Rebox_Data_Date_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumDay()))))
	})
	return cache_Data_Date_toEnum
}

var cache_Data_Date_ordMaybe gopurs_runtime.Value
var once_Data_Date_ordMaybe sync.Once
func Get_Data_Date_ordMaybe() gopurs_runtime.Value {
	once_Data_Date_ordMaybe.Do(func() {
		cache_Data_Date_ordMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3092443796_4177771502(Rebox_Data_Date_4177771502_3092443796(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Maybe_ordMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3308271157_4177771502(Rebox_Data_Date_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))})))))}
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
_ = __t_tag_0
if (uint32(__t_tag_0) == 1908470532) {
__t12 = "January"
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
_ = __t_tag_1
if (uint32(__t_tag_1) == 2455627378) {
__t12 = "February"
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
_ = __t_tag_2
if (uint32(__t_tag_2) == 4162469099) {
__t12 = "March"
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
_ = __t_tag_3
if (uint32(__t_tag_3) == 1692989816) {
__t12 = "April"
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
_ = __t_tag_4
if (uint32(__t_tag_4) == 330658827) {
__t12 = "May"
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
_ = __t_tag_5
if (uint32(__t_tag_5) == 4067355978) {
__t12 = "June"
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
_ = __t_tag_6
if (uint32(__t_tag_6) == 2276710548) {
__t12 = "July"
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
_ = __t_tag_7
if (uint32(__t_tag_7) == 243771071) {
__t12 = "August"
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
_ = __t_tag_8
if (uint32(__t_tag_8) == 215731793) {
__t12 = "September"
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
_ = __t_tag_9
if (uint32(__t_tag_9) == 8639228) {
__t12 = "October"
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
_ = __t_tag_10
if (uint32(__t_tag_10) == 49471444) {
__t12 = "November"
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
_ = __t_tag_11
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
return gopurs_runtime.Str((((((("(Date (Year ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0)).StrVal())) + (") ")) + (__t12)) + (" (Day ")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V2)).StrVal())) + ("))"))
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
_ = __t_tag_3
if (uint32(__t_tag_3) == 1908470532) {
var __t_tag_4 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
_ = __t_tag_4
__t25 = (uint32(__t_tag_4) == 1908470532)
goto end_branch_25
} else {

}
}
{
var __t_tag_5 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
_ = __t_tag_5
if (uint32(__t_tag_5) == 2455627378) {
var __t_tag_6 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
_ = __t_tag_6
__t25 = (uint32(__t_tag_6) == 2455627378)
goto end_branch_25
} else {

}
}
{
var __t_tag_7 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
_ = __t_tag_7
if (uint32(__t_tag_7) == 4162469099) {
var __t_tag_8 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
_ = __t_tag_8
__t25 = (uint32(__t_tag_8) == 4162469099)
goto end_branch_25
} else {

}
}
{
var __t_tag_9 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
_ = __t_tag_9
if (uint32(__t_tag_9) == 1692989816) {
var __t_tag_10 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
_ = __t_tag_10
__t25 = (uint32(__t_tag_10) == 1692989816)
goto end_branch_25
} else {

}
}
{
var __t_tag_11 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
_ = __t_tag_11
if (uint32(__t_tag_11) == 330658827) {
var __t_tag_12 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
_ = __t_tag_12
__t25 = (uint32(__t_tag_12) == 330658827)
goto end_branch_25
} else {

}
}
{
var __t_tag_13 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
_ = __t_tag_13
if (uint32(__t_tag_13) == 4067355978) {
var __t_tag_14 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
_ = __t_tag_14
__t25 = (uint32(__t_tag_14) == 4067355978)
goto end_branch_25
} else {

}
}
{
var __t_tag_15 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
_ = __t_tag_15
if (uint32(__t_tag_15) == 2276710548) {
var __t_tag_16 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
_ = __t_tag_16
__t25 = (uint32(__t_tag_16) == 2276710548)
goto end_branch_25
} else {

}
}
{
var __t_tag_17 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
_ = __t_tag_17
if (uint32(__t_tag_17) == 243771071) {
var __t_tag_18 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
_ = __t_tag_18
__t25 = (uint32(__t_tag_18) == 243771071)
goto end_branch_25
} else {

}
}
{
var __t_tag_19 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
_ = __t_tag_19
if (uint32(__t_tag_19) == 215731793) {
var __t_tag_20 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
_ = __t_tag_20
__t25 = (uint32(__t_tag_20) == 215731793)
goto end_branch_25
} else {

}
}
{
var __t_tag_21 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
_ = __t_tag_21
if (uint32(__t_tag_21) == 8639228) {
var __t_tag_22 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
_ = __t_tag_22
__t25 = (uint32(__t_tag_22) == 8639228)
goto end_branch_25
} else {

}
}
{
var __t_tag_23 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
_ = __t_tag_23
if (uint32(__t_tag_23) == 49471444) {
var __t_tag_24 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
_ = __t_tag_24
__t25 = (uint32(__t_tag_24) == 49471444)
goto end_branch_25
} else {

}
}
{
var __t_tag_0 uint32 = (*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1
_ = __t_tag_0
var __t_and_2 bool = false
if (uint32(__t_tag_0) == 3889233761) {

var __t_tag_1 uint32 = (*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1
_ = __t_tag_1
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
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1957390985_3790796878(Rebox_Data_Date_3790796878_1957390985(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Date_eqDate()))))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_2_0 := uint32(gopurs_runtime.Apply2(Rebox_Data_Date_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V0)).IntVal)
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
v1_3_1 := uint32(gopurs_runtime.Apply2(Rebox_Data_Date_4177771502_3730953251(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Date_Component_ordMonth())).V1, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1), UnsafePtr: nil}).IntVal)
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
__t2 = uint32(gopurs_runtime.Apply2(Rebox_Data_Date_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt())).V1, gopurs_runtime.Int((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V2)).IntVal)
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
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_758368489_4177771502(Rebox_Data_Date_4177771502_758368489(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Date_ordDate()))))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pm_1_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
var pm_1_0 *Constructor_Data_Maybe_Just[uint32] = Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_2359585123(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumMonth())).V1, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil})))
// TAST (Let): pd_2_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
var pd_2_1 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumDay())).V1, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V2))))
var __t5 uint32
{
var __t3 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t3 = gopurs_runtime.Bool(true)
goto end_branch_3
} else {

}
}
{
if (pd_2_1 != nil) {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
if (__t3.IntVal) != (0) {
var __t4 gopurs_runtime.Value
{
if (pm_1_0 == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if (pm_1_0 != nil) {
__t4 = gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: int64((pm_1_0).V0), UnsafePtr: nil})
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = uint32(__t4.IntVal)
goto end_branch_5
} else {

}
}
{
__t5 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
}
end_branch_5:
// TAST (Let): m_prime__3_2 shape=Branch(Branch(Other, App(Var), def=Other), def=Other) bindingType=Any
m_prime__3_2 := __t5
_ = m_prime__3_2
// TAST (Let): l_4_6 shape=App(Var) bindingType=Int
l_4_6 := Call_Data_Date_lastDayOfMonth((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0, m_prime__3_2)
_ = l_4_6
var __t11 gopurs_runtime.Value
{
var __t8 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t8 = gopurs_runtime.Bool(true)
goto end_branch_8
} else {

}
}
{
if (pd_2_1 != nil) {
__t8 = gopurs_runtime.Bool(false)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
var __t_and_10 bool = false
if (__t8.IntVal) != (0) {

var __t9 gopurs_runtime.Value
{
if (pm_1_0 == nil) {
__t9 = gopurs_runtime.Bool(true)
goto end_branch_9
} else {

}
}
{
if (pm_1_0 != nil) {
__t9 = gopurs_runtime.Bool(false)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
__t_and_10 = (__t9.IntVal) != (0)
}
if __t_and_10 {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumYear())).V1, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0))))}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
}
end_branch_11:
// TAST (Let): __local_var_5_7 shape=Branch(App(Other), def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_5_7 := Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t11))
_ = __local_var_5_7
var __t18 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_5_7 != nil) {
// TAST (Let): __local_var_6_14 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(Func [Int] (ADT ["Data","Date","Date"] []))])
__local_var_6_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((__local_var_5_7).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(m_prime__3_2), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
_ = __local_var_6_14
var __t17 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t16 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t16 = gopurs_runtime.Bool(true)
goto end_branch_16
} else {

}
}
{
if (pd_2_1 != nil) {
__t16 = gopurs_runtime.Bool(false)
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
if (__t16.IntVal) != (0) {
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_6_14).V0, gopurs_runtime.Int(l_4_6)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_17
} else {

}
}
{
var __t15 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (pd_2_1 != nil) {
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_6_14).V0, gopurs_runtime.Int((pd_2_1).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_15:
__t17 = __t15
}
end_branch_17:
__t18 = __t17
goto end_branch_18
} else {

}
}
{
var __t13 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t12 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t12 = gopurs_runtime.Bool(true)
goto end_branch_12
} else {

}
}
{
if (pd_2_1 != nil) {
__t12 = gopurs_runtime.Bool(false)
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
if (__t12.IntVal) != (0) {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_13:
__t18 = __t13
}
end_branch_18:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t18)}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): sm_1_19 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
var sm_1_19 *Constructor_Data_Maybe_Just[uint32] = Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_2359585123(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumMonth())).V2, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil})))
// TAST (Let): v1_2_20 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
var v1_2_20 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumDay())).V2, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V2))))
var __t61 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_36 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Maybe_ordMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3308271157_4177771502(Rebox_Data_Date_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(v1_2_20))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0, (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))})
_ = __t_tag_36
if (uint32(__t_tag_36.IntVal) == 380165415) {
var __t43 gopurs_runtime.Value
{
var __t40 gopurs_runtime.Value
{
var __t_tag_38 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_38
if (__t_tag_38 == nil) {
__t40 = gopurs_runtime.Bool(true)
goto end_branch_40
} else {

}
}
{
var __t_tag_39 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_39
if (__t_tag_39 != nil) {
__t40 = gopurs_runtime.Bool(false)
goto end_branch_40
} else {

}
}
{
__t40 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_40:
var __t_and_42 bool = false
if (__t40.IntVal) != (0) {

var __t41 gopurs_runtime.Value
{
if (sm_1_19 == nil) {
__t41 = gopurs_runtime.Bool(true)
goto end_branch_41
} else {

}
}
{
if (sm_1_19 != nil) {
__t41 = gopurs_runtime.Bool(false)
goto end_branch_41
} else {

}
}
{
__t41 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_41:
__t_and_42 = (__t41.IntVal) != (0)
}
if __t_and_42 {
__t43 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumYear())).V2, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0))))}
goto end_branch_43
} else {

}
}
{
__t43 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
}
end_branch_43:
// TAST (Let): __local_var_3_37 shape=Branch(App(Other), def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_3_37 := Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t43))
_ = __local_var_3_37
var __t60 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_37 != nil) {
var __t53 uint32
{
var __t51 gopurs_runtime.Value
{
var __t_tag_49 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_49
if (__t_tag_49 == nil) {
__t51 = gopurs_runtime.Bool(true)
goto end_branch_51
} else {

}
}
{
var __t_tag_50 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_50
if (__t_tag_50 != nil) {
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
if (sm_1_19 == nil) {
__t52 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_52
} else {

}
}
{
if (sm_1_19 != nil) {
__t52 = gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: int64((sm_1_19).V0), UnsafePtr: nil})
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
// TAST (Let): __local_var_4_48 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(Func [Int] (ADT ["Data","Date","Date"] []))])
__local_var_4_48 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((__local_var_3_37).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(__t53), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
_ = __local_var_4_48
var __t59 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t58 gopurs_runtime.Value
{
var __t_tag_56 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_56
if (__t_tag_56 == nil) {
__t58 = gopurs_runtime.Bool(true)
goto end_branch_58
} else {

}
}
{
var __t_tag_57 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_57
if (__t_tag_57 != nil) {
__t58 = gopurs_runtime.Bool(false)
goto end_branch_58
} else {

}
}
{
__t58 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_58:
if (__t58.IntVal) != (0) {
__t59 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_4_48).V0, gopurs_runtime.Int(int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_59
} else {

}
}
{
var __t55 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_54 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_54
if (__t_tag_54 != nil) {
__t55 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_4_48).V0, ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_55
} else {

}
}
{
__t55 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_55:
__t59 = __t55
}
end_branch_59:
__t60 = __t59
goto end_branch_60
} else {

}
}
{
var __t47 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t46 gopurs_runtime.Value
{
var __t_tag_44 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_44
if (__t_tag_44 == nil) {
__t46 = gopurs_runtime.Bool(true)
goto end_branch_46
} else {

}
}
{
var __t_tag_45 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_45
if (__t_tag_45 != nil) {
__t46 = gopurs_runtime.Bool(false)
goto end_branch_46
} else {

}
}
{
__t46 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_46:
if (__t46.IntVal) != (0) {
__t47 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_47
} else {

}
}
{
__t47 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_47:
__t60 = __t47
}
end_branch_60:
__t61 = __t60
goto end_branch_61
} else {

}
}
{
var __t25 gopurs_runtime.Value
{
var __t22 gopurs_runtime.Value
{
if (v1_2_20 == nil) {
__t22 = gopurs_runtime.Bool(true)
goto end_branch_22
} else {

}
}
{
if (v1_2_20 != nil) {
__t22 = gopurs_runtime.Bool(false)
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
var __t_and_24 bool = false
if (__t22.IntVal) != (0) {

var __t23 gopurs_runtime.Value
{
if (sm_1_19 == nil) {
__t23 = gopurs_runtime.Bool(true)
goto end_branch_23
} else {

}
}
{
if (sm_1_19 != nil) {
__t23 = gopurs_runtime.Bool(false)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t_and_24 = (__t23.IntVal) != (0)
}
if __t_and_24 {
__t25 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumYear())).V2, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0))))}
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
}
end_branch_25:
// TAST (Let): __local_var_3_21 shape=Branch(App(Other), def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_3_21 := Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t25))
_ = __local_var_3_21
var __t35 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_21 != nil) {
var __t31 uint32
{
var __t29 gopurs_runtime.Value
{
if (v1_2_20 == nil) {
__t29 = gopurs_runtime.Bool(true)
goto end_branch_29
} else {

}
}
{
if (v1_2_20 != nil) {
__t29 = gopurs_runtime.Bool(false)
goto end_branch_29
} else {

}
}
{
__t29 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_29:
if (__t29.IntVal) != (0) {
var __t30 gopurs_runtime.Value
{
if (sm_1_19 == nil) {
__t30 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_30
} else {

}
}
{
if (sm_1_19 != nil) {
__t30 = gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: int64((sm_1_19).V0), UnsafePtr: nil})
goto end_branch_30
} else {

}
}
{
__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_30:
__t31 = uint32(__t30.IntVal)
goto end_branch_31
} else {

}
}
{
__t31 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
}
end_branch_31:
// TAST (Let): __local_var_4_28 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(Func [Int] (ADT ["Data","Date","Date"] []))])
__local_var_4_28 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((__local_var_3_21).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(__t31), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
_ = __local_var_4_28
var __t34 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t33 gopurs_runtime.Value
{
if (v1_2_20 == nil) {
__t33 = gopurs_runtime.Bool(true)
goto end_branch_33
} else {

}
}
{
if (v1_2_20 != nil) {
__t33 = gopurs_runtime.Bool(false)
goto end_branch_33
} else {

}
}
{
__t33 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_33:
if (__t33.IntVal) != (0) {
__t34 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_4_28).V0, gopurs_runtime.Int(int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_34
} else {

}
}
{
var __t32 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_2_20 != nil) {
__t32 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_4_28).V0, gopurs_runtime.Int((v1_2_20).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_32
} else {

}
}
{
__t32 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_32:
__t34 = __t32
}
end_branch_34:
__t35 = __t34
goto end_branch_35
} else {

}
}
{
var __t27 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t26 gopurs_runtime.Value
{
if (v1_2_20 == nil) {
__t26 = gopurs_runtime.Bool(true)
goto end_branch_26
} else {

}
}
{
if (v1_2_20 != nil) {
__t26 = gopurs_runtime.Bool(false)
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
if (__t26.IntVal) != (0) {
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
__t27 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_27:
__t35 = __t27
}
end_branch_35:
__t61 = __t35
}
end_branch_61:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t61)}
})})))}
	})
	return cache_Data_Date_enumDate
}

var cache_Data_Date_pred gopurs_runtime.Value
var once_Data_Date_pred sync.Once
func Get_Data_Date_pred() gopurs_runtime.Value {
	once_Data_Date_pred.Do(func() {
		cache_Data_Date_pred = Call_Data_Enum_pred(Rebox_Data_Date_3494563625_556578094(Rebox_Data_Date_556578094_3494563625(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_enumDate()))))
	})
	return cache_Data_Date_pred
}

var cache_Data_Date_diff gopurs_runtime.Value
var once_Data_Date_diff sync.Once
func Get_Data_Date_diff() gopurs_runtime.Value {
	once_Data_Date_diff.Do(func() {
		cache_Data_Date_diff = gopurs_runtime.Func(func(dictDuration_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Date_diff(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]](dictDuration_0_box))
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
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_758368489_4177771502(Rebox_Data_Date_4177771502_758368489(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Date_ordDate()))))}
}), (&Constructor_Data_Date_Date{1, Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3764732725_2094947566(Rebox_Data_Date_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Date_Component_boundedYear()))))}).IntVal, uint32(Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_832288803_2094947566(Rebox_Data_Date_2094947566_832288803(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Date_Component_boundedMonth()))))}).IntVal), Call_Data_Bounded_bottom(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3764732725_2094947566(Rebox_Data_Date_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Date_Component_boundedDay()))))}).IntVal}), (&Constructor_Data_Date_Date{1, Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3764732725_2094947566(Rebox_Data_Date_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Date_Component_boundedYear()))))}).IntVal, uint32(Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_832288803_2094947566(Rebox_Data_Date_2094947566_832288803(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Date_Component_boundedMonth()))))}).IntVal), Call_Data_Bounded_top(gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3764732725_2094947566(Rebox_Data_Date_2094947566_3764732725(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](Get_Data_Date_Component_boundedDay()))))}).IntVal})})))}
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
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
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
_ = __t_tag_1
if (uint32(__t_tag_1) == 1908470532) {
__t13 = int64(1)
goto end_branch_13
} else {

}
}
{
var __t_tag_2 uint32 = (v_0).V1
_ = __t_tag_2
if (uint32(__t_tag_2) == 2455627378) {
__t13 = int64(2)
goto end_branch_13
} else {

}
}
{
var __t_tag_3 uint32 = (v_0).V1
_ = __t_tag_3
if (uint32(__t_tag_3) == 4162469099) {
__t13 = int64(3)
goto end_branch_13
} else {

}
}
{
var __t_tag_4 uint32 = (v_0).V1
_ = __t_tag_4
if (uint32(__t_tag_4) == 1692989816) {
__t13 = int64(4)
goto end_branch_13
} else {

}
}
{
var __t_tag_5 uint32 = (v_0).V1
_ = __t_tag_5
if (uint32(__t_tag_5) == 330658827) {
__t13 = int64(5)
goto end_branch_13
} else {

}
}
{
var __t_tag_6 uint32 = (v_0).V1
_ = __t_tag_6
if (uint32(__t_tag_6) == 4067355978) {
__t13 = int64(6)
goto end_branch_13
} else {

}
}
{
var __t_tag_7 uint32 = (v_0).V1
_ = __t_tag_7
if (uint32(__t_tag_7) == 2276710548) {
__t13 = int64(7)
goto end_branch_13
} else {

}
}
{
var __t_tag_8 uint32 = (v_0).V1
_ = __t_tag_8
if (uint32(__t_tag_8) == 243771071) {
__t13 = int64(8)
goto end_branch_13
} else {

}
}
{
var __t_tag_9 uint32 = (v_0).V1
_ = __t_tag_9
if (uint32(__t_tag_9) == 215731793) {
__t13 = int64(9)
goto end_branch_13
} else {

}
}
{
var __t_tag_10 uint32 = (v_0).V1
_ = __t_tag_10
if (uint32(__t_tag_10) == 8639228) {
__t13 = int64(10)
goto end_branch_13
} else {

}
}
{
var __t_tag_11 uint32 = (v_0).V1
_ = __t_tag_11
if (uint32(__t_tag_11) == 49471444) {
__t13 = int64(11)
goto end_branch_13
} else {

}
}
{
var __t_tag_12 uint32 = (v_0).V1
_ = __t_tag_12
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
// TAST (Let): n_1_0 shape=UncurriedApp(Var) bindingType=Int
n_1_0 := gopurs_runtime.UncurriedApp3(Get_Data_Date_calcWeekday(), gopurs_runtime.Int((v_0).V0), gopurs_runtime.Int(__t13), gopurs_runtime.Int((v_0).V2)).IntVal
_ = n_1_0
var __t15 uint32
{
if (n_1_0) == (int64(0)) {
__t15 = 1326716170
goto end_branch_15
} else {

}
}
{
var __t14 uint32
{
if (n_1_0) == (int64(1)) {
__t14 = 2900196686
goto end_branch_14
} else {

}
}
{
if (n_1_0) == (int64(2)) {
__t14 = 20457557
goto end_branch_14
} else {

}
}
{
if (n_1_0) == (int64(3)) {
__t14 = 4227105004
goto end_branch_14
} else {

}
}
{
if (n_1_0) == (int64(4)) {
__t14 = 3818857258
goto end_branch_14
} else {

}
}
{
if (n_1_0) == (int64(5)) {
__t14 = 2946274527
goto end_branch_14
} else {

}
}
{
if (n_1_0) == (int64(6)) {
__t14 = 1070786179
goto end_branch_14
} else {

}
}
{
if (n_1_0) == (int64(7)) {
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
return ((gopurs_runtime.IntMod(y_0, int64(4))) == (int64(0))) && (((gopurs_runtime.IntMod(y_0, int64(400))) == (int64(0))) || (((gopurs_runtime.IntMod(y_0, int64(100))) == (int64(0))) != (true)))
}

func Call_Data_Date_lastDayOfMonth(y_0_loop int64, m_1_loop uint32) int64 {
var y_0 int64 = y_0_loop
_ = y_0
var m_1 uint32 = m_1_loop
_ = m_1
// TAST (Let): unsafeDay__3466805691_2_0 shape=App(Var) bindingType=(Func [Int] Int)
unsafeDay__3466805691_2_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_fromJust__218925574(), Call_Data_Enum_toEnum(Rebox_Data_Date_1306125126_123048125(Rebox_Data_Date_123048125_1306125126(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Date_Component_boundedEnumDay())))))
_ = unsafeDay__3466805691_2_0
var __t2 int64
{
if (m_1 == 1908470532) {
__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 2455627378) {
var __t1 int64
{
if ((gopurs_runtime.IntMod(y_0, int64(4))) == (int64(0))) && (((gopurs_runtime.IntMod(y_0, int64(400))) == (int64(0))) || (((gopurs_runtime.IntMod(y_0, int64(100))) == (int64(0))) != (true))) {
__t1 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(29))).IntVal
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(28))).IntVal
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (m_1 == 4162469099) {
__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 1692989816) {
__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(30))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 330658827) {
__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 4067355978) {
__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(30))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 2276710548) {
__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 243771071) {
__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 215731793) {
__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(30))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 8639228) {
__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 49471444) {
__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(30))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 3889233761) {
__t2 = gopurs_runtime.Apply(unsafeDay__3466805691_2_0, gopurs_runtime.Int(int64(31))).IntVal
goto end_branch_2
} else {

}
}
{
__t2 = func() int64 { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}

func Call_Data_Date_diff(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration[gopurs_runtime.Value] = dictDuration_0_loop
_ = dictDuration_0
// TAST (Let): toDuration_1_0 shape=App(Var) bindingType=(Func [Number] (TypeVar d$scope15))
toDuration_1_0 := Call_Data_Time_Duration_toDuration(dictDuration_0)
_ = toDuration_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 int64
{
var __t_tag_1 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
_ = __t_tag_1
if (uint32(__t_tag_1) == 1908470532) {
__t13 = int64(1)
goto end_branch_13
} else {

}
}
{
var __t_tag_2 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
_ = __t_tag_2
if (uint32(__t_tag_2) == 2455627378) {
__t13 = int64(2)
goto end_branch_13
} else {

}
}
{
var __t_tag_3 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
_ = __t_tag_3
if (uint32(__t_tag_3) == 4162469099) {
__t13 = int64(3)
goto end_branch_13
} else {

}
}
{
var __t_tag_4 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
_ = __t_tag_4
if (uint32(__t_tag_4) == 1692989816) {
__t13 = int64(4)
goto end_branch_13
} else {

}
}
{
var __t_tag_5 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
_ = __t_tag_5
if (uint32(__t_tag_5) == 330658827) {
__t13 = int64(5)
goto end_branch_13
} else {

}
}
{
var __t_tag_6 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
_ = __t_tag_6
if (uint32(__t_tag_6) == 4067355978) {
__t13 = int64(6)
goto end_branch_13
} else {

}
}
{
var __t_tag_7 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
_ = __t_tag_7
if (uint32(__t_tag_7) == 2276710548) {
__t13 = int64(7)
goto end_branch_13
} else {

}
}
{
var __t_tag_8 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
_ = __t_tag_8
if (uint32(__t_tag_8) == 243771071) {
__t13 = int64(8)
goto end_branch_13
} else {

}
}
{
var __t_tag_9 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
_ = __t_tag_9
if (uint32(__t_tag_9) == 215731793) {
__t13 = int64(9)
goto end_branch_13
} else {

}
}
{
var __t_tag_10 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
_ = __t_tag_10
if (uint32(__t_tag_10) == 8639228) {
__t13 = int64(10)
goto end_branch_13
} else {

}
}
{
var __t_tag_11 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
_ = __t_tag_11
if (uint32(__t_tag_11) == 49471444) {
__t13 = int64(11)
goto end_branch_13
} else {

}
}
{
var __t_tag_12 uint32 = (*Constructor_Data_Date_Date)(v_2.UnsafePtr).V1
_ = __t_tag_12
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
var __t26 int64
{
var __t_tag_14 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
_ = __t_tag_14
if (uint32(__t_tag_14) == 1908470532) {
__t26 = int64(1)
goto end_branch_26
} else {

}
}
{
var __t_tag_15 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
_ = __t_tag_15
if (uint32(__t_tag_15) == 2455627378) {
__t26 = int64(2)
goto end_branch_26
} else {

}
}
{
var __t_tag_16 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
_ = __t_tag_16
if (uint32(__t_tag_16) == 4162469099) {
__t26 = int64(3)
goto end_branch_26
} else {

}
}
{
var __t_tag_17 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
_ = __t_tag_17
if (uint32(__t_tag_17) == 1692989816) {
__t26 = int64(4)
goto end_branch_26
} else {

}
}
{
var __t_tag_18 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
_ = __t_tag_18
if (uint32(__t_tag_18) == 330658827) {
__t26 = int64(5)
goto end_branch_26
} else {

}
}
{
var __t_tag_19 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
_ = __t_tag_19
if (uint32(__t_tag_19) == 4067355978) {
__t26 = int64(6)
goto end_branch_26
} else {

}
}
{
var __t_tag_20 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
_ = __t_tag_20
if (uint32(__t_tag_20) == 2276710548) {
__t26 = int64(7)
goto end_branch_26
} else {

}
}
{
var __t_tag_21 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
_ = __t_tag_21
if (uint32(__t_tag_21) == 243771071) {
__t26 = int64(8)
goto end_branch_26
} else {

}
}
{
var __t_tag_22 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
_ = __t_tag_22
if (uint32(__t_tag_22) == 215731793) {
__t26 = int64(9)
goto end_branch_26
} else {

}
}
{
var __t_tag_23 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
_ = __t_tag_23
if (uint32(__t_tag_23) == 8639228) {
__t26 = int64(10)
goto end_branch_26
} else {

}
}
{
var __t_tag_24 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
_ = __t_tag_24
if (uint32(__t_tag_24) == 49471444) {
__t26 = int64(11)
goto end_branch_26
} else {

}
}
{
var __t_tag_25 uint32 = (*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V1
_ = __t_tag_25
if (uint32(__t_tag_25) == 3889233761) {
__t26 = int64(12)
goto end_branch_26
} else {

}
}
{
__t26 = func() int64 { panic("Failed pattern match") }()
}
end_branch_26:
return gopurs_runtime.Apply(toDuration_1_0, gopurs_runtime.Float(gopurs_runtime.UncurriedApp6(Get_Data_Date_calcDiff(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_2.UnsafePtr).V0), gopurs_runtime.Int(__t13), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_2.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V0), gopurs_runtime.Int(__t26), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v1_3.UnsafePtr).V2)).FloatVal()))
})
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
var __t15 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]
{
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
__local_var_3_0 := Call_Data_Date_canonicalDate(y_0, m_1, d_2)
_ = __local_var_3_0
var __t_and_14 bool = false
if ((__local_var_3_0).V0) == (y_0) {

var __t13 bool
{
var __t_tag_2 uint32 = (__local_var_3_0).V1
_ = __t_tag_2
if (uint32(__t_tag_2) == 1908470532) {
__t13 = (m_1 == 1908470532)
goto end_branch_13
} else {

}
}
{
var __t_tag_3 uint32 = (__local_var_3_0).V1
_ = __t_tag_3
if (uint32(__t_tag_3) == 2455627378) {
__t13 = (m_1 == 2455627378)
goto end_branch_13
} else {

}
}
{
var __t_tag_4 uint32 = (__local_var_3_0).V1
_ = __t_tag_4
if (uint32(__t_tag_4) == 4162469099) {
__t13 = (m_1 == 4162469099)
goto end_branch_13
} else {

}
}
{
var __t_tag_5 uint32 = (__local_var_3_0).V1
_ = __t_tag_5
if (uint32(__t_tag_5) == 1692989816) {
__t13 = (m_1 == 1692989816)
goto end_branch_13
} else {

}
}
{
var __t_tag_6 uint32 = (__local_var_3_0).V1
_ = __t_tag_6
if (uint32(__t_tag_6) == 330658827) {
__t13 = (m_1 == 330658827)
goto end_branch_13
} else {

}
}
{
var __t_tag_7 uint32 = (__local_var_3_0).V1
_ = __t_tag_7
if (uint32(__t_tag_7) == 4067355978) {
__t13 = (m_1 == 4067355978)
goto end_branch_13
} else {

}
}
{
var __t_tag_8 uint32 = (__local_var_3_0).V1
_ = __t_tag_8
if (uint32(__t_tag_8) == 2276710548) {
__t13 = (m_1 == 2276710548)
goto end_branch_13
} else {

}
}
{
var __t_tag_9 uint32 = (__local_var_3_0).V1
_ = __t_tag_9
if (uint32(__t_tag_9) == 243771071) {
__t13 = (m_1 == 243771071)
goto end_branch_13
} else {

}
}
{
var __t_tag_10 uint32 = (__local_var_3_0).V1
_ = __t_tag_10
if (uint32(__t_tag_10) == 215731793) {
__t13 = (m_1 == 215731793)
goto end_branch_13
} else {

}
}
{
var __t_tag_11 uint32 = (__local_var_3_0).V1
_ = __t_tag_11
if (uint32(__t_tag_11) == 8639228) {
__t13 = (m_1 == 8639228)
goto end_branch_13
} else {

}
}
{
var __t_tag_12 uint32 = (__local_var_3_0).V1
_ = __t_tag_12
if (uint32(__t_tag_12) == 49471444) {
__t13 = (m_1 == 49471444)
goto end_branch_13
} else {

}
}
{
var __t_tag_1 uint32 = (__local_var_3_0).V1
_ = __t_tag_1
__t13 = ((uint32(__t_tag_1) == 3889233761)) && ((m_1 == 3889233761))
}
end_branch_13:
__t_and_14 = __t13
}
if (__t_and_14) && (((__local_var_3_0).V2) == (d_2)) {
__t15 = Rebox_Data_Date_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, y_0, m_1, d_2}))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_15
} else {

}
}
{
__t15 = Rebox_Data_Date_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_15:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_2280409795_3094389156(__t15))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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
var adj_2_0_0_cell *gopurs_runtime.Value
_ = adj_2_0_0_cell
// FALLBACK TCO: isLoop=false len=1
adj_2_0_0 = gopurs_runtime.Func2(func(v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t55 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]
{
if (v1_3.IntVal) == (int64(0)) {
__t55 = Rebox_Data_Date_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v2_4))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_55
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
var __t6 uint32
{
if low_6_2 {
// TAST (Let): __local_var_7_4 shape=App(Other) bindingType=Any
__local_var_7_4 := gopurs_runtime.Apply(Rebox_Data_Date_556578094_2359585123(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumMonth())).V1, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil})
_ = __local_var_7_4
var __t5 gopurs_runtime.Value
{
if (__local_var_7_4.Type == 9 && __local_var_7_4.IntVal == 930809136 && __local_var_7_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
if (__local_var_7_4.Type == 9 && __local_var_7_4.IntVal == 930809136 && __local_var_7_4.UnsafePtr != nil) {
__t5 = gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_7_4.UnsafePtr).V0.IntVal)), UnsafePtr: nil})
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = uint32(__t5.IntVal)
goto end_branch_6
} else {

}
}
{
__t6 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
}
end_branch_6:
// TAST (Let): l_7_3 shape=App(Var) bindingType=Int
l_7_3 := Call_Data_Date_lastDayOfMonth((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0, __t6)
_ = l_7_3
// TAST (Let): hi_8_7 shape=Other bindingType=Boolean
hi_8_7 := (j_5_1) > (l_7_3)
_ = hi_8_7
var __t53 gopurs_runtime.Value
{
if low_6_2 {
__t53 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_2280409795_3094389156(Rebox_Data_Date_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Call_Control_Bind_bind(Rebox_Data_Date_3887487416_2748095225(Rebox_Data_Date_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())))), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((&Constructor_Data_Date_Date{1, (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0, (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1, int64(1)}))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))}, Call_Data_Enum_pred(Rebox_Data_Date_3494563625_556578094(Rebox_Data_Date_556578094_3494563625(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_enumDate())))))))))}
goto end_branch_53
} else {

}
}
{
if hi_8_7 {
// TAST (Let): sm_9_10 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Date","Component","Month"] [])])
var sm_9_10 *Constructor_Data_Maybe_Just[uint32] = Rebox_Data_Date_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_2359585123(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumMonth())).V2, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil})))
// TAST (Let): v1_10_11 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
var v1_10_11 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumDay())).V2, gopurs_runtime.Int(l_7_3))))
var __t52 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_27 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_Maybe_ordMaybe(gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_3308271157_4177771502(Rebox_Data_Date_4177771502_3308271157(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordInt()))))}), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(v1_10_11))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0, (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))})
_ = __t_tag_27
if (uint32(__t_tag_27.IntVal) == 380165415) {
var __t34 gopurs_runtime.Value
{
var __t31 gopurs_runtime.Value
{
var __t_tag_29 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_29
if (__t_tag_29 == nil) {
__t31 = gopurs_runtime.Bool(true)
goto end_branch_31
} else {

}
}
{
var __t_tag_30 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_30
if (__t_tag_30 != nil) {
__t31 = gopurs_runtime.Bool(false)
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
var __t_and_33 bool = false
if (__t31.IntVal) != (0) {

var __t32 gopurs_runtime.Value
{
if (sm_9_10 == nil) {
__t32 = gopurs_runtime.Bool(true)
goto end_branch_32
} else {

}
}
{
if (sm_9_10 != nil) {
__t32 = gopurs_runtime.Bool(false)
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_32:
__t_and_33 = (__t32.IntVal) != (0)
}
if __t_and_33 {
__t34 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumYear())).V2, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0))))}
goto end_branch_34
} else {

}
}
{
__t34 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
}
end_branch_34:
// TAST (Let): __local_var_11_28 shape=Branch(App(Other), def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_11_28 := Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t34))
_ = __local_var_11_28
var __t51 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_11_28 != nil) {
var __t44 uint32
{
var __t42 gopurs_runtime.Value
{
var __t_tag_40 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_40
if (__t_tag_40 == nil) {
__t42 = gopurs_runtime.Bool(true)
goto end_branch_42
} else {

}
}
{
var __t_tag_41 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_41
if (__t_tag_41 != nil) {
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
var __t43 gopurs_runtime.Value
{
if (sm_9_10 == nil) {
__t43 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_43
} else {

}
}
{
if (sm_9_10 != nil) {
__t43 = gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: int64((sm_9_10).V0), UnsafePtr: nil})
goto end_branch_43
} else {

}
}
{
__t43 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_43:
__t44 = uint32(__t43.IntVal)
goto end_branch_44
} else {

}
}
{
__t44 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
}
end_branch_44:
// TAST (Let): __local_var_12_39 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(Func [Int] (ADT ["Data","Date","Date"] []))])
__local_var_12_39 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((__local_var_11_28).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(__t44), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
_ = __local_var_12_39
var __t50 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t49 gopurs_runtime.Value
{
var __t_tag_47 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_47
if (__t_tag_47 == nil) {
__t49 = gopurs_runtime.Bool(true)
goto end_branch_49
} else {

}
}
{
var __t_tag_48 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_48
if (__t_tag_48 != nil) {
__t49 = gopurs_runtime.Bool(false)
goto end_branch_49
} else {

}
}
{
__t49 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_49:
if (__t49.IntVal) != (0) {
__t50 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_12_39).V0, gopurs_runtime.Int(int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_50
} else {

}
}
{
var __t46 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_45 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_45
if (__t_tag_45 != nil) {
__t46 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_12_39).V0, ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_46
} else {

}
}
{
__t46 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_46:
__t50 = __t46
}
end_branch_50:
__t51 = __t50
goto end_branch_51
} else {

}
}
{
var __t38 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t37 gopurs_runtime.Value
{
var __t_tag_35 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_35
if (__t_tag_35 == nil) {
__t37 = gopurs_runtime.Bool(true)
goto end_branch_37
} else {

}
}
{
var __t_tag_36 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
_ = __t_tag_36
if (__t_tag_36 != nil) {
__t37 = gopurs_runtime.Bool(false)
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
if (__t37.IntVal) != (0) {
__t38 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_38
} else {

}
}
{
__t38 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_38:
__t51 = __t38
}
end_branch_51:
__t52 = __t51
goto end_branch_52
} else {

}
}
{
var __t16 gopurs_runtime.Value
{
var __t13 gopurs_runtime.Value
{
if (v1_10_11 == nil) {
__t13 = gopurs_runtime.Bool(true)
goto end_branch_13
} else {

}
}
{
if (v1_10_11 != nil) {
__t13 = gopurs_runtime.Bool(false)
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
var __t_and_15 bool = false
if (__t13.IntVal) != (0) {

var __t14 gopurs_runtime.Value
{
if (sm_9_10 == nil) {
__t14 = gopurs_runtime.Bool(true)
goto end_branch_14
} else {

}
}
{
if (sm_9_10 != nil) {
__t14 = gopurs_runtime.Bool(false)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
__t_and_15 = (__t14.IntVal) != (0)
}
if __t_and_15 {
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Rebox_Data_Date_556578094_4060049525(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[gopurs_runtime.Value]](Get_Data_Date_Component_enumYear())).V2, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0))))}
goto end_branch_16
} else {

}
}
{
__t16 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_1170268447_3094389156(Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
}
end_branch_16:
// TAST (Let): __local_var_11_12 shape=Branch(App(Other), def=Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_11_12 := Rebox_Data_Date_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t16))
_ = __local_var_11_12
var __t26 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_11_12 != nil) {
var __t22 uint32
{
var __t20 gopurs_runtime.Value
{
if (v1_10_11 == nil) {
__t20 = gopurs_runtime.Bool(true)
goto end_branch_20
} else {

}
}
{
if (v1_10_11 != nil) {
__t20 = gopurs_runtime.Bool(false)
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
if (__t20.IntVal) != (0) {
var __t21 gopurs_runtime.Value
{
if (sm_9_10 == nil) {
__t21 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_21
} else {

}
}
{
if (sm_9_10 != nil) {
__t21 = gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: int64((sm_9_10).V0), UnsafePtr: nil})
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
__t22 = uint32(__t21.IntVal)
goto end_branch_22
} else {

}
}
{
__t22 = (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1
}
end_branch_22:
// TAST (Let): __local_var_12_19 shape=Other bindingType=(ADT ["Data","Maybe","Maybe"] [(Func [Int] (ADT ["Data","Date","Date"] []))])
__local_var_12_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((__local_var_11_12).V0), gopurs_runtime.Value{Type: 9, IntVal: int64(__t22), UnsafePtr: nil}), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
_ = __local_var_12_19
var __t25 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t24 gopurs_runtime.Value
{
if (v1_10_11 == nil) {
__t24 = gopurs_runtime.Bool(true)
goto end_branch_24
} else {

}
}
{
if (v1_10_11 != nil) {
__t24 = gopurs_runtime.Bool(false)
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
if (__t24.IntVal) != (0) {
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_12_19).V0, gopurs_runtime.Int(int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_25
} else {

}
}
{
var __t23 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_10_11 != nil) {
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((__local_var_12_19).V0, gopurs_runtime.Int((v1_10_11).V0)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_23:
__t25 = __t23
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
var __t18 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t17 gopurs_runtime.Value
{
if (v1_10_11 == nil) {
__t17 = gopurs_runtime.Bool(true)
goto end_branch_17
} else {

}
}
{
if (v1_10_11 != nil) {
__t17 = gopurs_runtime.Bool(false)
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
if (__t17.IntVal) != (0) {
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
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
__t26 = __t18
}
end_branch_26:
__t52 = __t26
}
end_branch_52:
__t53 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t52)}
goto end_branch_53
} else {

}
}
{
// TAST (Let): __local_var_9_8 shape=App(Var) bindingType=Any
__local_var_9_8 := gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil})
_ = __local_var_9_8
var __t9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((j_5_1) >= (int64(1))) && ((j_5_1) <= (int64(31))) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(__local_var_9_8, gopurs_runtime.Int(j_5_1)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_9:
__t53 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
}
end_branch_53:
var __t54 int64
{
if low_6_2 {
__t54 = j_5_1
goto end_branch_54
} else {

}
}
{
if hi_8_7 {
__t54 = ((j_5_1) - (l_7_3)) - (int64(1))
goto end_branch_54
} else {

}
}
{
__t54 = int64(0)
}
end_branch_54:
__t55 = Rebox_Data_Date_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Call_Control_Bind_bind(Rebox_Data_Date_3887487416_2748095225(Rebox_Data_Date_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe())))), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t53))}, gopurs_runtime.Apply((*adj_2_0_0_cell), gopurs_runtime.Int(__t54)))))
}
end_branch_55:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_2280409795_3094389156(__t55))}
})
adj_2_0_0_cell = &adj_2_0_0
// TAST (Let): __local_var_3_56 shape=App(Var) bindingType=Any
__local_var_3_56 := gopurs_runtime.Apply(Get_Data_Int_fromNumber(), gopurs_runtime.Float(v_0))
_ = __local_var_3_56
var __t57 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_56.Type == 9 && __local_var_3_56.IntVal == 930809136 && __local_var_3_56.UnsafePtr != nil) {
__t57 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(adj_2_0_0, gopurs_runtime.Int((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_3_56.UnsafePtr).V0.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(date_1)}))
goto end_branch_57
} else {

}
}
{
if (__local_var_3_56.Type == 9 && __local_var_3_56.IntVal == 930809136 && __local_var_3_56.UnsafePtr == nil) {
__t57 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_57
} else {

}
}
{
__t57 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_57:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t57)}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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

func Rebox_Data_Date_123048125_1306125126(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Date_1306125126_123048125(in *Constructor_Data_Enum_BoundedEnum[int64]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
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

func Rebox_Data_Date_2094947566_3764732725(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1.IntVal
		out.V2 = in.V2.IntVal
	return out
}

func Rebox_Data_Date_2094947566_832288803(in *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]) *Constructor_Data_Bounded_Bounded[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[uint32]{}
		out.V0 = in.V0
		out.V1 = uint32(in.V1.IntVal)
		out.V2 = uint32(in.V2.IntVal)
	return out
}

func Rebox_Data_Date_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Date_2748095225_3887487416(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
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

func Rebox_Data_Date_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
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

func Rebox_Data_Date_3523628265_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_Date_Date]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V1)}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_Date_3764732725_2094947566(in *Constructor_Data_Bounded_Bounded[int64]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Int(in.V1)
		out.V2 = gopurs_runtime.Int(in.V2)
	return out
}

func Rebox_Data_Date_3790796878_1957390985(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Date_Date]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Date_3887487416_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Date_4177771502_3092443796(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[int64]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[int64]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Date_4177771502_3308271157(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Date_4177771502_3730953251(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[uint32]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Date_4177771502_758368489(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Date_556578094_2359585123(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[uint32]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Date_556578094_3494563625(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[*Constructor_Data_Date_Date]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Date_556578094_4060049525(in *Constructor_Data_Enum_Enum[gopurs_runtime.Value]) *Constructor_Data_Enum_Enum[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Date_758368489_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Date_Date]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Date_832288803_2094947566(in *Constructor_Data_Bounded_Bounded[uint32]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
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

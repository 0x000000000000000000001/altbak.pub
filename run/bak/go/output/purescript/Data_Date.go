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
		cache_Data_Date_toEnum = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Date_toEnum(n_0_box.IntVal))}
})
	})
	return cache_Data_Date_toEnum
}

var cache_Data_Date_ordMaybe gopurs_runtime.Value
var once_Data_Date_ordMaybe sync.Once
func Get_Data_Date_ordMaybe() gopurs_runtime.Value {
	once_Data_Date_ordMaybe.Do(func() {
		cache_Data_Date_ordMaybe = func() gopurs_runtime.Value {
// TAST (Let): eqMaybe1_0_0 -> *Constructor_Data_Eq_Eq
eqMaybe1_0_0 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 bool
{
if (x_0.Type == 9 && x_0.IntVal == 930809136 && x_0.UnsafePtr == nil) {
var __t1 bool
{
if (y_1.Type == 9 && y_1.IntVal == 930809136 && y_1.UnsafePtr == nil) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 930809136 && x_0.UnsafePtr != nil)) && ((y_1.Type == 9 && y_1.IntVal == 930809136 && y_1.UnsafePtr != nil)) {
__t2 = ((*Constructor_Data_Maybe_Just)(x_0.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Maybe_Just)(y_1.UnsafePtr).V0.IntVal)
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
})
})}
_ = eqMaybe1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMaybe1_0_0)}
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 uint32
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t3 uint32
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t3 = 902936544
goto end_branch_3
} else {

}
}
{
__t3 = 1527465420
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t4 = uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, (*Constructor_Data_Maybe_Just)(x_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_2.UnsafePtr).V0).IntVal)
goto end_branch_4
} else {

}
}
{
__t4 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
})
})})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Data_Date_Date{1, value0.IntVal, uint32(value1.IntVal), value2.IntVal})}
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
		cache_Data_Date_weekday = gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 int64
{
var __t_tag_1 uint32 = (*Constructor_Data_Date_Date)(v_1.UnsafePtr).V1
if (uint32(__t_tag_1) == 1908470532) {
__t13 = 1
goto end_branch_13
} else {

}
}
{
var __t_tag_2 uint32 = (*Constructor_Data_Date_Date)(v_1.UnsafePtr).V1
if (uint32(__t_tag_2) == 2455627378) {
__t13 = 2
goto end_branch_13
} else {

}
}
{
var __t_tag_3 uint32 = (*Constructor_Data_Date_Date)(v_1.UnsafePtr).V1
if (uint32(__t_tag_3) == 4162469099) {
__t13 = 3
goto end_branch_13
} else {

}
}
{
var __t_tag_4 uint32 = (*Constructor_Data_Date_Date)(v_1.UnsafePtr).V1
if (uint32(__t_tag_4) == 1692989816) {
__t13 = 4
goto end_branch_13
} else {

}
}
{
var __t_tag_5 uint32 = (*Constructor_Data_Date_Date)(v_1.UnsafePtr).V1
if (uint32(__t_tag_5) == 330658827) {
__t13 = 5
goto end_branch_13
} else {

}
}
{
var __t_tag_6 uint32 = (*Constructor_Data_Date_Date)(v_1.UnsafePtr).V1
if (uint32(__t_tag_6) == 4067355978) {
__t13 = 6
goto end_branch_13
} else {

}
}
{
var __t_tag_7 uint32 = (*Constructor_Data_Date_Date)(v_1.UnsafePtr).V1
if (uint32(__t_tag_7) == 2276710548) {
__t13 = 7
goto end_branch_13
} else {

}
}
{
var __t_tag_8 uint32 = (*Constructor_Data_Date_Date)(v_1.UnsafePtr).V1
if (uint32(__t_tag_8) == 243771071) {
__t13 = 8
goto end_branch_13
} else {

}
}
{
var __t_tag_9 uint32 = (*Constructor_Data_Date_Date)(v_1.UnsafePtr).V1
if (uint32(__t_tag_9) == 215731793) {
__t13 = 9
goto end_branch_13
} else {

}
}
{
var __t_tag_10 uint32 = (*Constructor_Data_Date_Date)(v_1.UnsafePtr).V1
if (uint32(__t_tag_10) == 8639228) {
__t13 = 10
goto end_branch_13
} else {

}
}
{
var __t_tag_11 uint32 = (*Constructor_Data_Date_Date)(v_1.UnsafePtr).V1
if (uint32(__t_tag_11) == 49471444) {
__t13 = 11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 uint32 = (*Constructor_Data_Date_Date)(v_1.UnsafePtr).V1
if (uint32(__t_tag_12) == 3889233761) {
__t13 = 12
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_13:
// TAST (Let): n_2_0 -> gopurs_runtime.Value
n_2_0 := gopurs_runtime.UncurriedApp3(Get_Data_Date_calcWeekday(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_1.UnsafePtr).V0), gopurs_runtime.Int(__t13), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_1.UnsafePtr).V2))
_ = n_2_0
var __t18 uint32
{
if (n_2_0.IntVal) == (0) {
// TAST (Let): __local_var_3_16 -> *Constructor_Data_Maybe_Just
__local_var_3_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2793813158(), gopurs_runtime.Int(7)))
_ = __local_var_3_16
var __t17 uint32
{
if (__local_var_3_16 != nil) {
__t17 = uint32((__local_var_3_16).V0.IntVal)
goto end_branch_17
} else {

}
}
{
__t17 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_17:
__t18 = __t17
goto end_branch_18
} else {

}
}
{
// TAST (Let): __local_var_3_14 -> *Constructor_Data_Maybe_Just
__local_var_3_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2793813158(), gopurs_runtime.Int(n_2_0.IntVal)))
_ = __local_var_3_14
var __t15 uint32
{
if (__local_var_3_14 != nil) {
__t15 = uint32((__local_var_3_14).V0.IntVal)
goto end_branch_15
} else {

}
}
{
__t15 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_15:
__t18 = __t15
}
end_branch_18:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t18), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Date_weekday
}

var cache_Data_Date_showDate gopurs_runtime.Value
var once_Data_Date_showDate sync.Once
func Get_Data_Date_showDate() gopurs_runtime.Value {
	once_Data_Date_showDate.Do(func() {
		cache_Data_Date_showDate = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})})}
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
		cache_Data_Date_eqDate = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V0) == ((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V0)) && ((gopurs_runtime.Apply2(Get_Data_Eq_eq__3887832182(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1), UnsafePtr: nil}).IntVal) != (0))) && (((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V2) == ((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V2)))
})
})})}
	})
	return cache_Data_Date_eqDate
}

var cache_Data_Date_ordDate gopurs_runtime.Value
var once_Data_Date_ordDate sync.Once
func Get_Data_Date_ordDate() gopurs_runtime.Value {
	once_Data_Date_ordDate.Do(func() {
		cache_Data_Date_ordDate = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Date_eqDate()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 -> gopurs_runtime.Value
v_2_0 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V0))
_ = v_2_0
var __t3 uint32
{
if (uint32(v_2_0.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(v_2_0.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
// TAST (Let): v1_3_1 -> gopurs_runtime.Value
v1_3_1 := gopurs_runtime.Apply2(Get_Data_Ord_compare__696857420(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1), UnsafePtr: nil})
_ = v1_3_1
var __t2 uint32
{
if (uint32(v1_3_1.IntVal) == 1527465420) {
__t2 = 1527465420
goto end_branch_2
} else {

}
}
{
if (uint32(v1_3_1.IntVal) == 380165415) {
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
})
})})}
	})
	return cache_Data_Date_ordDate
}

var cache_Data_Date_enumDate gopurs_runtime.Value
var once_Data_Date_enumDate sync.Once
func Get_Data_Date_enumDate() gopurs_runtime.Value {
	once_Data_Date_enumDate.Do(func() {
		cache_Data_Date_enumDate = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Enum_Enum{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Date_ordDate()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 int64
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_2.IntVal) == 1908470532) {
__t14 = 1
goto end_branch_14
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_3.IntVal) == 2455627378) {
__t14 = 2
goto end_branch_14
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_4.IntVal) == 4162469099) {
__t14 = 3
goto end_branch_14
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_5.IntVal) == 1692989816) {
__t14 = 4
goto end_branch_14
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_6.IntVal) == 330658827) {
__t14 = 5
goto end_branch_14
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_7.IntVal) == 4067355978) {
__t14 = 6
goto end_branch_14
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_8.IntVal) == 2276710548) {
__t14 = 7
goto end_branch_14
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_9.IntVal) == 243771071) {
__t14 = 8
goto end_branch_14
} else {

}
}
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_10.IntVal) == 215731793) {
__t14 = 9
goto end_branch_14
} else {

}
}
{
var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_11.IntVal) == 8639228) {
__t14 = 10
goto end_branch_14
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_12.IntVal) == 49471444) {
__t14 = 11
goto end_branch_14
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_13.IntVal) == 3889233761) {
__t14 = 12
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_14:
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Int((gopurs_runtime.Int(__t14).IntVal) - (1))
_ = __local_var_1_1
var __t15 *Constructor_Data_Maybe_Just
{
if (__local_var_1_1.IntVal) == (1) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (2) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (3) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (4) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (5) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (6) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (7) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (8) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (9) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (10) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (11) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (12) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
__t15 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_15:
// TAST (Let): pm_1_0 -> *Constructor_Data_Maybe_Just
var pm_1_0 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t15)})
// TAST (Let): __local_var_2_17 -> gopurs_runtime.Value
__local_var_2_17 := gopurs_runtime.Int((gopurs_runtime.Int(gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V2).IntVal).IntVal) - (1))
_ = __local_var_2_17
var __t21 *Constructor_Data_Maybe_Just
{
var __t18 bool
{
if (__local_var_2_17.IntVal) < (1) {
__t18 = false
goto end_branch_18
} else {

}
}
{
__t18 = true
}
end_branch_18:
var __t_and_20 bool = false
if __t18 {

var __t19 bool
{
if (__local_var_2_17.IntVal) > (31) {
__t19 = false
goto end_branch_19
} else {

}
}
{
__t19 = true
}
end_branch_19:
__t_and_20 = __t19
}
if __t_and_20 {
__t21 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_2_17.IntVal)}
goto end_branch_21
} else {

}
}
{
__t21 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_21:
// TAST (Let): pd_2_16 -> *Constructor_Data_Maybe_Just
var pd_2_16 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t21)})
var __t25 uint32
{
var __t23 gopurs_runtime.Value
{
if (pd_2_16 == nil) {
__t23 = gopurs_runtime.Bool(true)
goto end_branch_23
} else {

}
}
{
if (pd_2_16 != nil) {
__t23 = gopurs_runtime.Bool(false)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
if (__t23.IntVal) != (0) {
var __t24 gopurs_runtime.Value
{
if (pm_1_0 == nil) {
__t24 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_24
} else {

}
}
{
if (pm_1_0 != nil) {
__t24 = (pm_1_0).V0
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
__t25 = uint32(__t24.IntVal)
goto end_branch_25
} else {

}
}
{
__t25 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
}
end_branch_25:
// TAST (Let): m_prime_3_22 -> gopurs_runtime.Value
var m_prime_3_22 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64(__t25), UnsafePtr: nil}
// TAST (Let): l_4_26 -> gopurs_runtime.Value
var l_4_26 gopurs_runtime.Value = gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0, uint32(m_prime_3_22.IntVal)))
var __t38 *Constructor_Data_Maybe_Just
{
var __t30 gopurs_runtime.Value
{
if (pd_2_16 == nil) {
__t30 = gopurs_runtime.Bool(true)
goto end_branch_30
} else {

}
}
{
if (pd_2_16 != nil) {
__t30 = gopurs_runtime.Bool(false)
goto end_branch_30
} else {

}
}
{
__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_30:
var __t_and_32 bool = false
if (__t30.IntVal) != (0) {

var __t31 gopurs_runtime.Value
{
if (pm_1_0 == nil) {
__t31 = gopurs_runtime.Bool(true)
goto end_branch_31
} else {

}
}
{
if (pm_1_0 != nil) {
__t31 = gopurs_runtime.Bool(false)
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
__t_and_32 = (__t31.IntVal) != (0)
}
if __t_and_32 {
// TAST (Let): __local_var_5_33 -> gopurs_runtime.Value
__local_var_5_33 := gopurs_runtime.Int((gopurs_runtime.Int(gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0).IntVal).IntVal) - (1))
_ = __local_var_5_33
var __t37 *Constructor_Data_Maybe_Just
{
var __t34 bool
{
if (__local_var_5_33.IntVal) < (1) {
__t34 = false
goto end_branch_34
} else {

}
}
{
__t34 = true
}
end_branch_34:
var __t_and_36 bool = false
if __t34 {

var __t35 bool
{
if (__local_var_5_33.IntVal) > (31) {
__t35 = false
goto end_branch_35
} else {

}
}
{
__t35 = true
}
end_branch_35:
__t_and_36 = __t35
}
if __t_and_36 {
__t37 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_5_33.IntVal)}
goto end_branch_37
} else {

}
}
{
__t37 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_37:
__t38 = __t37
goto end_branch_38
} else {

}
}
{
__t38 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0)}
}
end_branch_38:
// TAST (Let): __local_var_5_29 -> *Constructor_Data_Maybe_Just
var __local_var_5_29 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t38)})
var __t39 *Constructor_Data_Maybe_Just
{
if (__local_var_5_29 != nil) {
__t39 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Date_Date(), gopurs_runtime.Int((__local_var_5_29).V0.IntVal))}
goto end_branch_39
} else {

}
}
{
__t39 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_39:
// TAST (Let): __local_var_5_28 -> *Constructor_Data_Maybe_Just
var __local_var_5_28 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t39)})
var __t40 *Constructor_Data_Maybe_Just
{
if (__local_var_5_28 != nil) {
__t40 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_5_28).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(m_prime_3_22.IntVal)), UnsafePtr: nil})}
goto end_branch_40
} else {

}
}
{
if (__local_var_5_28 == nil) {
__t40 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_40
} else {

}
}
{
__t40 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_40:
// TAST (Let): __local_var_5_27 -> *Constructor_Data_Maybe_Just
var __local_var_5_27 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t40)})
var __t43 *Constructor_Data_Maybe_Just
{
var __t42 gopurs_runtime.Value
{
if (pd_2_16 == nil) {
__t42 = gopurs_runtime.Bool(true)
goto end_branch_42
} else {

}
}
{
if (pd_2_16 != nil) {
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
__t43 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(l_4_26.IntVal)})})
goto end_branch_43
} else {

}
}
{
__t43 = pd_2_16
}
end_branch_43:
// TAST (Let): __local_var_6_41 -> *Constructor_Data_Maybe_Just
__local_var_6_41 := __t43
_ = __local_var_6_41
var __t45 *Constructor_Data_Maybe_Just
{
if (__local_var_5_27 != nil) {
var __t44 *Constructor_Data_Maybe_Just
{
if (__local_var_6_41 != nil) {
__t44 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_5_27).V0, (__local_var_6_41).V0)}
goto end_branch_44
} else {

}
}
{
__t44 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_44:
__t45 = __t44
goto end_branch_45
} else {

}
}
{
if (__local_var_5_27 == nil) {
__t45 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_45
} else {

}
}
{
__t45 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_45:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t45)}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t60 int64
{
var __t_tag_48 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_48.IntVal) == 1908470532) {
__t60 = 2
goto end_branch_60
} else {

}
}
{
var __t_tag_49 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_49.IntVal) == 2455627378) {
__t60 = 3
goto end_branch_60
} else {

}
}
{
var __t_tag_50 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_50.IntVal) == 4162469099) {
__t60 = 4
goto end_branch_60
} else {

}
}
{
var __t_tag_51 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_51.IntVal) == 1692989816) {
__t60 = 5
goto end_branch_60
} else {

}
}
{
var __t_tag_52 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_52.IntVal) == 330658827) {
__t60 = 6
goto end_branch_60
} else {

}
}
{
var __t_tag_53 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_53.IntVal) == 4067355978) {
__t60 = 7
goto end_branch_60
} else {

}
}
{
var __t_tag_54 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_54.IntVal) == 2276710548) {
__t60 = 8
goto end_branch_60
} else {

}
}
{
var __t_tag_55 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_55.IntVal) == 243771071) {
__t60 = 9
goto end_branch_60
} else {

}
}
{
var __t_tag_56 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_56.IntVal) == 215731793) {
__t60 = 10
goto end_branch_60
} else {

}
}
{
var __t_tag_57 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_57.IntVal) == 8639228) {
__t60 = 11
goto end_branch_60
} else {

}
}
{
var __t_tag_58 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_58.IntVal) == 49471444) {
__t60 = 12
goto end_branch_60
} else {

}
}
{
var __t_tag_59 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_59.IntVal) == 3889233761) {
__t60 = 13
goto end_branch_60
} else {

}
}
{
__t60 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_60:
// TAST (Let): __local_var_1_47 -> gopurs_runtime.Value
__local_var_1_47 := gopurs_runtime.Int(__t60)
_ = __local_var_1_47
var __t61 *Constructor_Data_Maybe_Just
{
if (__local_var_1_47.IntVal) == (1) {
__t61 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}}
goto end_branch_61
} else {

}
}
{
if (__local_var_1_47.IntVal) == (2) {
__t61 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}}
goto end_branch_61
} else {

}
}
{
if (__local_var_1_47.IntVal) == (3) {
__t61 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}}
goto end_branch_61
} else {

}
}
{
if (__local_var_1_47.IntVal) == (4) {
__t61 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}}
goto end_branch_61
} else {

}
}
{
if (__local_var_1_47.IntVal) == (5) {
__t61 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}}
goto end_branch_61
} else {

}
}
{
if (__local_var_1_47.IntVal) == (6) {
__t61 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}}
goto end_branch_61
} else {

}
}
{
if (__local_var_1_47.IntVal) == (7) {
__t61 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}}
goto end_branch_61
} else {

}
}
{
if (__local_var_1_47.IntVal) == (8) {
__t61 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}}
goto end_branch_61
} else {

}
}
{
if (__local_var_1_47.IntVal) == (9) {
__t61 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}}
goto end_branch_61
} else {

}
}
{
if (__local_var_1_47.IntVal) == (10) {
__t61 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}}
goto end_branch_61
} else {

}
}
{
if (__local_var_1_47.IntVal) == (11) {
__t61 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}}
goto end_branch_61
} else {

}
}
{
if (__local_var_1_47.IntVal) == (12) {
__t61 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}}
goto end_branch_61
} else {

}
}
{
__t61 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_61:
// TAST (Let): sm_1_46 -> *Constructor_Data_Maybe_Just
var sm_1_46 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t61)})
// TAST (Let): __local_var_2_63 -> gopurs_runtime.Value
__local_var_2_63 := gopurs_runtime.Int(((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V2) + (1))
_ = __local_var_2_63
var __t67 *Constructor_Data_Maybe_Just
{
var __t64 bool
{
if (__local_var_2_63.IntVal) < (1) {
__t64 = false
goto end_branch_64
} else {

}
}
{
__t64 = true
}
end_branch_64:
var __t_and_66 bool = false
if __t64 {

var __t65 bool
{
if (__local_var_2_63.IntVal) > (31) {
__t65 = false
goto end_branch_65
} else {

}
}
{
__t65 = true
}
end_branch_65:
__t_and_66 = __t65
}
if __t_and_66 {
__t67 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_2_63.IntVal)}
goto end_branch_67
} else {

}
}
{
__t67 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_67:
// TAST (Let): v1_2_62 -> *Constructor_Data_Maybe_Just
var v1_2_62 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t67)})
var __t72 *Constructor_Data_Maybe_Just
{
var __t71 bool
{
// TAST (Let): __local_var_3_69 -> *Constructor_Data_Maybe_Just
var __local_var_3_69 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0, (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1))})})
var __t70 uint32
{
if (v1_2_62 == nil) {
__t70 = 1527465420
goto end_branch_70
} else {

}
}
{
if (v1_2_62 != nil) {
__t70 = uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, (v1_2_62).V0, (__local_var_3_69).V0).IntVal)
goto end_branch_70
} else {

}
}
{
__t70 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_70:
if (__t70 == 380165415) {
__t71 = true
goto end_branch_71
} else {

}
}
{
__t71 = false
}
end_branch_71:
if __t71 {
__t72 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_72
} else {

}
}
{
__t72 = v1_2_62
}
end_branch_72:
// TAST (Let): sd_3_68 -> *Constructor_Data_Maybe_Just
sd_3_68 := __t72
_ = sd_3_68
var __t84 *Constructor_Data_Maybe_Just
{
var __t76 gopurs_runtime.Value
{
if (sd_3_68 == nil) {
__t76 = gopurs_runtime.Bool(true)
goto end_branch_76
} else {

}
}
{
if (sd_3_68 != nil) {
__t76 = gopurs_runtime.Bool(false)
goto end_branch_76
} else {

}
}
{
__t76 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_76:
var __t_and_78 bool = false
if (__t76.IntVal) != (0) {

var __t77 gopurs_runtime.Value
{
if (sm_1_46 == nil) {
__t77 = gopurs_runtime.Bool(true)
goto end_branch_77
} else {

}
}
{
if (sm_1_46 != nil) {
__t77 = gopurs_runtime.Bool(false)
goto end_branch_77
} else {

}
}
{
__t77 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_77:
__t_and_78 = (__t77.IntVal) != (0)
}
if __t_and_78 {
// TAST (Let): __local_var_4_79 -> gopurs_runtime.Value
__local_var_4_79 := gopurs_runtime.Int(((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0) + (1))
_ = __local_var_4_79
var __t83 *Constructor_Data_Maybe_Just
{
var __t80 bool
{
if (__local_var_4_79.IntVal) < (1) {
__t80 = false
goto end_branch_80
} else {

}
}
{
__t80 = true
}
end_branch_80:
var __t_and_82 bool = false
if __t80 {

var __t81 bool
{
if (__local_var_4_79.IntVal) > (31) {
__t81 = false
goto end_branch_81
} else {

}
}
{
__t81 = true
}
end_branch_81:
__t_and_82 = __t81
}
if __t_and_82 {
__t83 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_4_79.IntVal)}
goto end_branch_83
} else {

}
}
{
__t83 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_83:
__t84 = __t83
goto end_branch_84
} else {

}
}
{
__t84 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0)}
}
end_branch_84:
// TAST (Let): __local_var_4_75 -> *Constructor_Data_Maybe_Just
var __local_var_4_75 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t84)})
var __t85 *Constructor_Data_Maybe_Just
{
if (__local_var_4_75 != nil) {
__t85 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Date_Date(), gopurs_runtime.Int((__local_var_4_75).V0.IntVal))}
goto end_branch_85
} else {

}
}
{
__t85 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_85:
// TAST (Let): __local_var_4_74 -> *Constructor_Data_Maybe_Just
var __local_var_4_74 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t85)})
var __t89 uint32
{
var __t87 gopurs_runtime.Value
{
if (sd_3_68 == nil) {
__t87 = gopurs_runtime.Bool(true)
goto end_branch_87
} else {

}
}
{
if (sd_3_68 != nil) {
__t87 = gopurs_runtime.Bool(false)
goto end_branch_87
} else {

}
}
{
__t87 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_87:
if (__t87.IntVal) != (0) {
var __t88 gopurs_runtime.Value
{
if (sm_1_46 == nil) {
__t88 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_88
} else {

}
}
{
if (sm_1_46 != nil) {
__t88 = (sm_1_46).V0
goto end_branch_88
} else {

}
}
{
__t88 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_88:
__t89 = uint32(__t88.IntVal)
goto end_branch_89
} else {

}
}
{
__t89 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
}
end_branch_89:
// TAST (Let): __local_var_5_86 -> *Constructor_Data_Maybe_Just
var __local_var_5_86 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(__t89), UnsafePtr: nil}})})
var __t90 *Constructor_Data_Maybe_Just
{
if (__local_var_4_74 != nil) {
__t90 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_4_74).V0, (__local_var_5_86).V0)}
goto end_branch_90
} else {

}
}
{
if (__local_var_4_74 == nil) {
__t90 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_90
} else {

}
}
{
__t90 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_90:
// TAST (Let): __local_var_4_73 -> *Constructor_Data_Maybe_Just
var __local_var_4_73 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t90)})
var __t93 *Constructor_Data_Maybe_Just
{
var __t92 gopurs_runtime.Value
{
if (sd_3_68 == nil) {
__t92 = gopurs_runtime.Bool(true)
goto end_branch_92
} else {

}
}
{
if (sd_3_68 != nil) {
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
__t93 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(1)})})
goto end_branch_93
} else {

}
}
{
__t93 = sd_3_68
}
end_branch_93:
// TAST (Let): __local_var_5_91 -> *Constructor_Data_Maybe_Just
__local_var_5_91 := __t93
_ = __local_var_5_91
var __t95 *Constructor_Data_Maybe_Just
{
if (__local_var_4_73 != nil) {
var __t94 *Constructor_Data_Maybe_Just
{
if (__local_var_5_91 != nil) {
__t94 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_4_73).V0, (__local_var_5_91).V0)}
goto end_branch_94
} else {

}
}
{
__t94 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_94:
__t95 = __t94
goto end_branch_95
} else {

}
}
{
if (__local_var_4_73 == nil) {
__t95 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_95
} else {

}
}
{
__t95 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_95:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t95)}
})})}
	})
	return cache_Data_Date_enumDate
}

var cache_Data_Date_pred gopurs_runtime.Value
var once_Data_Date_pred sync.Once
func Get_Data_Date_pred() gopurs_runtime.Value {
	once_Data_Date_pred.Do(func() {
		cache_Data_Date_pred = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Date_pred(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_0_box)))}
})
	})
	return cache_Data_Date_pred
}

var cache_Data_Date_diff gopurs_runtime.Value
var once_Data_Date_diff sync.Once
func Get_Data_Date_diff() gopurs_runtime.Value {
	once_Data_Date_diff.Do(func() {
		cache_Data_Date_diff = gopurs_runtime.Func3(func(dictDuration_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Date_diff(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Duration_Duration](dictDuration_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v1_2_box))
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Date_exactDate(y_0_box.IntVal, uint32(m_1_box.IntVal), d_2_box.IntVal))}
})
	})
	return cache_Data_Date_exactDate
}

var cache_Data_Date_boundedDate gopurs_runtime.Value
var once_Data_Date_boundedDate sync.Once
func Get_Data_Date_boundedDate() gopurs_runtime.Value {
	once_Data_Date_boundedDate.Do(func() {
		cache_Data_Date_boundedDate = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(&Constructor_Data_Bounded_Bounded{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Date_ordDate()))}
}), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Data_Date_Date{1, -271820, 1908470532, 1})}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Data_Date_Date{1, 275759, 3889233761, 31})}})}
	})
	return cache_Data_Date_boundedDate
}

var cache_Data_Date_adjust gopurs_runtime.Value
var once_Data_Date_adjust sync.Once
func Get_Data_Date_adjust() gopurs_runtime.Value {
	once_Data_Date_adjust.Do(func() {
		cache_Data_Date_adjust = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, date_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Date_adjust(v_0_box.FloatVal(), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](date_1_box)))}
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


func Call_Data_Date_toEnum(n_0_loop int64) *Constructor_Data_Maybe_Just {
var n_0 int64 = n_0_loop
_ = n_0
var __t3 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (n_0) < (1) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
var __t_and_2 bool = false
if __t0 {

var __t1 bool
{
if (n_0) > (31) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
__t_and_2 = __t1
}
if __t_and_2 {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(n_0)}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)})
}

func Call_Data_Date_year(v_0_loop *Constructor_Data_Date_Date) int64 {
var v_0 *Constructor_Data_Date_Date = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_Date_month(v_0_loop *Constructor_Data_Date_Date) uint32 {
var v_0 *Constructor_Data_Date_Date = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Date_isLeapYear(y_0_loop int64) bool {
var y_0 int64 = y_0_loop
_ = y_0
return ((gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), gopurs_runtime.Int(y_0), gopurs_runtime.Int(4)).IntVal) == (0)) && (((gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), gopurs_runtime.Int(y_0), gopurs_runtime.Int(400)).IntVal) == (0)) || (((gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), gopurs_runtime.Int(y_0), gopurs_runtime.Int(100)).IntVal) == (0)) != (true)))
}

func Call_Data_Date_lastDayOfMonth(y_0_loop int64, m_1_loop uint32) int64 {
var y_0 int64 = y_0_loop
_ = y_0
var m_1 uint32 = m_1_loop
_ = m_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_fromJust__1577979644()
}))
_ = __local_var_2_0
// TAST (Let): unsafeDay_3_1 -> gopurs_runtime.Value
unsafeDay_3_1 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
var __t2 bool
{
if (x_3.IntVal) < (1) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
var __t_and_4 bool = false
if __t2 {

var __t3 bool
{
if (x_3.IntVal) > (31) {
__t3 = false
goto end_branch_3
} else {

}
}
{
__t3 = true
}
end_branch_3:
__t_and_4 = __t3
}
if __t_and_4 {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(x_3.IntVal)}
goto end_branch_5
} else {

}
}
{
__t5 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_5:
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)})
})
_ = unsafeDay_3_1
var __t7 int64
{
if (m_1 == 1908470532) {
__t7 = gopurs_runtime.Apply(unsafeDay_3_1, gopurs_runtime.Int(31)).IntVal
goto end_branch_7
} else {

}
}
{
if (m_1 == 2455627378) {
var __t6 int64
{
if Call_Data_Date_isLeapYear(y_0) {
__t6 = gopurs_runtime.Apply(unsafeDay_3_1, gopurs_runtime.Int(29)).IntVal
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Apply(unsafeDay_3_1, gopurs_runtime.Int(28)).IntVal
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
if (m_1 == 4162469099) {
__t7 = gopurs_runtime.Apply(unsafeDay_3_1, gopurs_runtime.Int(31)).IntVal
goto end_branch_7
} else {

}
}
{
if (m_1 == 1692989816) {
__t7 = gopurs_runtime.Apply(unsafeDay_3_1, gopurs_runtime.Int(30)).IntVal
goto end_branch_7
} else {

}
}
{
if (m_1 == 330658827) {
__t7 = gopurs_runtime.Apply(unsafeDay_3_1, gopurs_runtime.Int(31)).IntVal
goto end_branch_7
} else {

}
}
{
if (m_1 == 4067355978) {
__t7 = gopurs_runtime.Apply(unsafeDay_3_1, gopurs_runtime.Int(30)).IntVal
goto end_branch_7
} else {

}
}
{
if (m_1 == 2276710548) {
__t7 = gopurs_runtime.Apply(unsafeDay_3_1, gopurs_runtime.Int(31)).IntVal
goto end_branch_7
} else {

}
}
{
if (m_1 == 243771071) {
__t7 = gopurs_runtime.Apply(unsafeDay_3_1, gopurs_runtime.Int(31)).IntVal
goto end_branch_7
} else {

}
}
{
if (m_1 == 215731793) {
__t7 = gopurs_runtime.Apply(unsafeDay_3_1, gopurs_runtime.Int(30)).IntVal
goto end_branch_7
} else {

}
}
{
if (m_1 == 8639228) {
__t7 = gopurs_runtime.Apply(unsafeDay_3_1, gopurs_runtime.Int(31)).IntVal
goto end_branch_7
} else {

}
}
{
if (m_1 == 49471444) {
__t7 = gopurs_runtime.Apply(unsafeDay_3_1, gopurs_runtime.Int(30)).IntVal
goto end_branch_7
} else {

}
}
{
if (m_1 == 3889233761) {
__t7 = gopurs_runtime.Apply(unsafeDay_3_1, gopurs_runtime.Int(31)).IntVal
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_7:
return __t7
}

func Call_Data_Date_pred(v_0_loop *Constructor_Data_Date_Date) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Date_Date = v_0_loop
_ = v_0
var __t14 int64
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}
if (uint32(__t_tag_2.IntVal) == 1908470532) {
__t14 = 1
goto end_branch_14
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}
if (uint32(__t_tag_3.IntVal) == 2455627378) {
__t14 = 2
goto end_branch_14
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}
if (uint32(__t_tag_4.IntVal) == 4162469099) {
__t14 = 3
goto end_branch_14
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}
if (uint32(__t_tag_5.IntVal) == 1692989816) {
__t14 = 4
goto end_branch_14
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}
if (uint32(__t_tag_6.IntVal) == 330658827) {
__t14 = 5
goto end_branch_14
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}
if (uint32(__t_tag_7.IntVal) == 4067355978) {
__t14 = 6
goto end_branch_14
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}
if (uint32(__t_tag_8.IntVal) == 2276710548) {
__t14 = 7
goto end_branch_14
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}
if (uint32(__t_tag_9.IntVal) == 243771071) {
__t14 = 8
goto end_branch_14
} else {

}
}
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}
if (uint32(__t_tag_10.IntVal) == 215731793) {
__t14 = 9
goto end_branch_14
} else {

}
}
{
var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}
if (uint32(__t_tag_11.IntVal) == 8639228) {
__t14 = 10
goto end_branch_14
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}
if (uint32(__t_tag_12.IntVal) == 49471444) {
__t14 = 11
goto end_branch_14
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}
if (uint32(__t_tag_13.IntVal) == 3889233761) {
__t14 = 12
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_14:
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Int((gopurs_runtime.Int(__t14).IntVal) - (1))
_ = __local_var_1_1
var __t15 *Constructor_Data_Maybe_Just
{
if (__local_var_1_1.IntVal) == (1) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (2) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (3) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (4) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (5) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (6) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (7) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (8) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (9) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (10) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (11) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
if (__local_var_1_1.IntVal) == (12) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}}
goto end_branch_15
} else {

}
}
{
__t15 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_15:
// TAST (Let): pm_1_0 -> *Constructor_Data_Maybe_Just
var pm_1_0 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t15)})
// TAST (Let): __local_var_2_17 -> gopurs_runtime.Value
__local_var_2_17 := gopurs_runtime.Int((gopurs_runtime.Int(gopurs_runtime.Int((v_0).V2).IntVal).IntVal) - (1))
_ = __local_var_2_17
var __t21 *Constructor_Data_Maybe_Just
{
var __t18 bool
{
if (__local_var_2_17.IntVal) < (1) {
__t18 = false
goto end_branch_18
} else {

}
}
{
__t18 = true
}
end_branch_18:
var __t_and_20 bool = false
if __t18 {

var __t19 bool
{
if (__local_var_2_17.IntVal) > (31) {
__t19 = false
goto end_branch_19
} else {

}
}
{
__t19 = true
}
end_branch_19:
__t_and_20 = __t19
}
if __t_and_20 {
__t21 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_2_17.IntVal)}
goto end_branch_21
} else {

}
}
{
__t21 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_21:
// TAST (Let): pd_2_16 -> *Constructor_Data_Maybe_Just
var pd_2_16 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t21)})
var __t25 uint32
{
var __t23 gopurs_runtime.Value
{
if (pd_2_16 == nil) {
__t23 = gopurs_runtime.Bool(true)
goto end_branch_23
} else {

}
}
{
if (pd_2_16 != nil) {
__t23 = gopurs_runtime.Bool(false)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
if (__t23.IntVal) != (0) {
var __t24 gopurs_runtime.Value
{
if (pm_1_0 == nil) {
__t24 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_24
} else {

}
}
{
if (pm_1_0 != nil) {
__t24 = (pm_1_0).V0
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
__t25 = uint32(__t24.IntVal)
goto end_branch_25
} else {

}
}
{
__t25 = (v_0).V1
}
end_branch_25:
// TAST (Let): m_prime_3_22 -> gopurs_runtime.Value
var m_prime_3_22 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64(__t25), UnsafePtr: nil}
// TAST (Let): l_4_26 -> gopurs_runtime.Value
var l_4_26 gopurs_runtime.Value = gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth((v_0).V0, uint32(m_prime_3_22.IntVal)))
var __t38 *Constructor_Data_Maybe_Just
{
var __t30 gopurs_runtime.Value
{
if (pd_2_16 == nil) {
__t30 = gopurs_runtime.Bool(true)
goto end_branch_30
} else {

}
}
{
if (pd_2_16 != nil) {
__t30 = gopurs_runtime.Bool(false)
goto end_branch_30
} else {

}
}
{
__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_30:
var __t_and_32 bool = false
if (__t30.IntVal) != (0) {

var __t31 gopurs_runtime.Value
{
if (pm_1_0 == nil) {
__t31 = gopurs_runtime.Bool(true)
goto end_branch_31
} else {

}
}
{
if (pm_1_0 != nil) {
__t31 = gopurs_runtime.Bool(false)
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
__t_and_32 = (__t31.IntVal) != (0)
}
if __t_and_32 {
// TAST (Let): __local_var_5_33 -> gopurs_runtime.Value
__local_var_5_33 := gopurs_runtime.Int((gopurs_runtime.Int(gopurs_runtime.Int((v_0).V0).IntVal).IntVal) - (1))
_ = __local_var_5_33
var __t37 *Constructor_Data_Maybe_Just
{
var __t34 bool
{
if (__local_var_5_33.IntVal) < (1) {
__t34 = false
goto end_branch_34
} else {

}
}
{
__t34 = true
}
end_branch_34:
var __t_and_36 bool = false
if __t34 {

var __t35 bool
{
if (__local_var_5_33.IntVal) > (31) {
__t35 = false
goto end_branch_35
} else {

}
}
{
__t35 = true
}
end_branch_35:
__t_and_36 = __t35
}
if __t_and_36 {
__t37 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_5_33.IntVal)}
goto end_branch_37
} else {

}
}
{
__t37 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_37:
__t38 = __t37
goto end_branch_38
} else {

}
}
{
__t38 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((v_0).V0)}
}
end_branch_38:
// TAST (Let): __local_var_5_29 -> *Constructor_Data_Maybe_Just
var __local_var_5_29 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t38)})
var __t39 *Constructor_Data_Maybe_Just
{
if (__local_var_5_29 != nil) {
__t39 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Date_Date(), gopurs_runtime.Int((__local_var_5_29).V0.IntVal))}
goto end_branch_39
} else {

}
}
{
__t39 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_39:
// TAST (Let): __local_var_5_28 -> *Constructor_Data_Maybe_Just
var __local_var_5_28 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t39)})
var __t40 *Constructor_Data_Maybe_Just
{
if (__local_var_5_28 != nil) {
__t40 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_5_28).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(m_prime_3_22.IntVal)), UnsafePtr: nil})}
goto end_branch_40
} else {

}
}
{
if (__local_var_5_28 == nil) {
__t40 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_40
} else {

}
}
{
__t40 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_40:
// TAST (Let): __local_var_5_27 -> *Constructor_Data_Maybe_Just
var __local_var_5_27 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t40)})
var __t43 *Constructor_Data_Maybe_Just
{
var __t42 gopurs_runtime.Value
{
if (pd_2_16 == nil) {
__t42 = gopurs_runtime.Bool(true)
goto end_branch_42
} else {

}
}
{
if (pd_2_16 != nil) {
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
__t43 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(l_4_26.IntVal)})})
goto end_branch_43
} else {

}
}
{
__t43 = pd_2_16
}
end_branch_43:
// TAST (Let): __local_var_6_41 -> *Constructor_Data_Maybe_Just
__local_var_6_41 := __t43
_ = __local_var_6_41
var __t45 *Constructor_Data_Maybe_Just
{
if (__local_var_5_27 != nil) {
var __t44 *Constructor_Data_Maybe_Just
{
if (__local_var_6_41 != nil) {
__t44 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_5_27).V0, (__local_var_6_41).V0)}
goto end_branch_44
} else {

}
}
{
__t44 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_44:
__t45 = __t44
goto end_branch_45
} else {

}
}
{
if (__local_var_5_27 == nil) {
__t45 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_45
} else {

}
}
{
__t45 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_45:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t45)})
}

func Call_Data_Date_diff(dictDuration_0_loop *Constructor_Data_Time_Duration_Duration, v_1_loop *Constructor_Data_Date_Date, v1_2_loop *Constructor_Data_Date_Date) gopurs_runtime.Value {
var dictDuration_0 *Constructor_Data_Time_Duration_Duration = dictDuration_0_loop
_ = dictDuration_0
var v_1 *Constructor_Data_Date_Date = v_1_loop
_ = v_1
var v1_2 *Constructor_Data_Date_Date = v1_2_loop
_ = v1_2
var __t12 int64
{
var __t_tag_0 uint32 = (v_1).V1
if (uint32(__t_tag_0) == 1908470532) {
__t12 = 1
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = (v_1).V1
if (uint32(__t_tag_1) == 2455627378) {
__t12 = 2
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = (v_1).V1
if (uint32(__t_tag_2) == 4162469099) {
__t12 = 3
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = (v_1).V1
if (uint32(__t_tag_3) == 1692989816) {
__t12 = 4
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = (v_1).V1
if (uint32(__t_tag_4) == 330658827) {
__t12 = 5
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = (v_1).V1
if (uint32(__t_tag_5) == 4067355978) {
__t12 = 6
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = (v_1).V1
if (uint32(__t_tag_6) == 2276710548) {
__t12 = 7
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = (v_1).V1
if (uint32(__t_tag_7) == 243771071) {
__t12 = 8
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = (v_1).V1
if (uint32(__t_tag_8) == 215731793) {
__t12 = 9
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = (v_1).V1
if (uint32(__t_tag_9) == 8639228) {
__t12 = 10
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = (v_1).V1
if (uint32(__t_tag_10) == 49471444) {
__t12 = 11
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = (v_1).V1
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
var __t25 int64
{
var __t_tag_13 uint32 = (v1_2).V1
if (uint32(__t_tag_13) == 1908470532) {
__t25 = 1
goto end_branch_25
} else {

}
}
{
var __t_tag_14 uint32 = (v1_2).V1
if (uint32(__t_tag_14) == 2455627378) {
__t25 = 2
goto end_branch_25
} else {

}
}
{
var __t_tag_15 uint32 = (v1_2).V1
if (uint32(__t_tag_15) == 4162469099) {
__t25 = 3
goto end_branch_25
} else {

}
}
{
var __t_tag_16 uint32 = (v1_2).V1
if (uint32(__t_tag_16) == 1692989816) {
__t25 = 4
goto end_branch_25
} else {

}
}
{
var __t_tag_17 uint32 = (v1_2).V1
if (uint32(__t_tag_17) == 330658827) {
__t25 = 5
goto end_branch_25
} else {

}
}
{
var __t_tag_18 uint32 = (v1_2).V1
if (uint32(__t_tag_18) == 4067355978) {
__t25 = 6
goto end_branch_25
} else {

}
}
{
var __t_tag_19 uint32 = (v1_2).V1
if (uint32(__t_tag_19) == 2276710548) {
__t25 = 7
goto end_branch_25
} else {

}
}
{
var __t_tag_20 uint32 = (v1_2).V1
if (uint32(__t_tag_20) == 243771071) {
__t25 = 8
goto end_branch_25
} else {

}
}
{
var __t_tag_21 uint32 = (v1_2).V1
if (uint32(__t_tag_21) == 215731793) {
__t25 = 9
goto end_branch_25
} else {

}
}
{
var __t_tag_22 uint32 = (v1_2).V1
if (uint32(__t_tag_22) == 8639228) {
__t25 = 10
goto end_branch_25
} else {

}
}
{
var __t_tag_23 uint32 = (v1_2).V1
if (uint32(__t_tag_23) == 49471444) {
__t25 = 11
goto end_branch_25
} else {

}
}
{
var __t_tag_24 uint32 = (v1_2).V1
if (uint32(__t_tag_24) == 3889233761) {
__t25 = 12
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
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
var __t2 int64
{
if (m_1 == 1908470532) {
__t2 = 1
goto end_branch_2
} else {

}
}
{
if (m_1 == 2455627378) {
__t2 = 2
goto end_branch_2
} else {

}
}
{
if (m_1 == 4162469099) {
__t2 = 3
goto end_branch_2
} else {

}
}
{
if (m_1 == 1692989816) {
__t2 = 4
goto end_branch_2
} else {

}
}
{
if (m_1 == 330658827) {
__t2 = 5
goto end_branch_2
} else {

}
}
{
if (m_1 == 4067355978) {
__t2 = 6
goto end_branch_2
} else {

}
}
{
if (m_1 == 2276710548) {
__t2 = 7
goto end_branch_2
} else {

}
}
{
if (m_1 == 243771071) {
__t2 = 8
goto end_branch_2
} else {

}
}
{
if (m_1 == 215731793) {
__t2 = 9
goto end_branch_2
} else {

}
}
{
if (m_1 == 8639228) {
__t2 = 10
goto end_branch_2
} else {

}
}
{
if (m_1 == 49471444) {
__t2 = 11
goto end_branch_2
} else {

}
}
{
if (m_1 == 3889233761) {
__t2 = 12
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.UncurriedApp4(Get_Data_Date_canonicalDateImpl(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_0 -> *Constructor_Data_Maybe_Just
__local_var_7_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2309750950(), gopurs_runtime.Int(m_prime_5.IntVal)))
_ = __local_var_7_0
var __t1 uint32
{
if (__local_var_7_0 != nil) {
__t1 = uint32((__local_var_7_0).V0.IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Data_Date_Date{1, y_prime_4.IntVal, __t1, d_prime_6.IntVal})}
})
})
})
})), gopurs_runtime.Int(y_0), gopurs_runtime.Int(__t2), gopurs_runtime.Int(d_2)))
}

func Call_Data_Date_exactDate(y_0_loop int64, m_1_loop uint32, d_2_loop int64) *Constructor_Data_Maybe_Just {
var y_0 int64 = y_0_loop
_ = y_0
var m_1 uint32 = m_1_loop
_ = m_1
var d_2 int64 = d_2_loop
_ = d_2
var __t0 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply2(Get_Data_Eq_eq__1204755874(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Date_canonicalDate(y_0, m_1, d_2))}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Data_Date_Date{1, y_0, m_1, d_2})}).IntVal) != (0) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Data_Date_Date{1, y_0, m_1, d_2})}}
goto end_branch_0
} else {

}
}
{
__t0 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_Date_adjust(v_0_loop float64, date_1_loop *Constructor_Data_Date_Date) *Constructor_Data_Maybe_Just {
var v_0 float64 = v_0_loop
_ = v_0
var date_1 *Constructor_Data_Date_Date = date_1_loop
_ = date_1
var adj_2_0_0 gopurs_runtime.Value
adj_2_0_0 = gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
var v2_4_loop *Constructor_Data_Date_Date = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v2_4_loop_val)
adj_2_0_0:
for {
if false { continue adj_2_0_0 }
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var v2_4 *Constructor_Data_Date_Date = v2_4_loop
_ = v2_4
var __t68 *Constructor_Data_Maybe_Just
{
if (v1_3.IntVal) == (0) {
__t68 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(v2_4)}}
goto end_branch_68
} else {

}
}
{
// TAST (Let): j_5_1 -> int64
j_5_1 := (v1_3.IntVal) + ((v2_4).V2)
_ = j_5_1
var __t3 bool
{
if (j_5_1) < (1) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
// TAST (Let): low_6_2 -> bool
low_6_2 := __t3
_ = low_6_2
var __t22 uint32
{
if low_6_2 {
var __t19 int64
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil}
if (uint32(__t_tag_7.IntVal) == 1908470532) {
__t19 = 1
goto end_branch_19
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil}
if (uint32(__t_tag_8.IntVal) == 2455627378) {
__t19 = 2
goto end_branch_19
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil}
if (uint32(__t_tag_9.IntVal) == 4162469099) {
__t19 = 3
goto end_branch_19
} else {

}
}
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil}
if (uint32(__t_tag_10.IntVal) == 1692989816) {
__t19 = 4
goto end_branch_19
} else {

}
}
{
var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil}
if (uint32(__t_tag_11.IntVal) == 330658827) {
__t19 = 5
goto end_branch_19
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil}
if (uint32(__t_tag_12.IntVal) == 4067355978) {
__t19 = 6
goto end_branch_19
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil}
if (uint32(__t_tag_13.IntVal) == 2276710548) {
__t19 = 7
goto end_branch_19
} else {

}
}
{
var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil}
if (uint32(__t_tag_14.IntVal) == 243771071) {
__t19 = 8
goto end_branch_19
} else {

}
}
{
var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil}
if (uint32(__t_tag_15.IntVal) == 215731793) {
__t19 = 9
goto end_branch_19
} else {

}
}
{
var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil}
if (uint32(__t_tag_16.IntVal) == 8639228) {
__t19 = 10
goto end_branch_19
} else {

}
}
{
var __t_tag_17 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil}
if (uint32(__t_tag_17.IntVal) == 49471444) {
__t19 = 11
goto end_branch_19
} else {

}
}
{
var __t_tag_18 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil}
if (uint32(__t_tag_18.IntVal) == 3889233761) {
__t19 = 12
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_19:
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.Int((gopurs_runtime.Int(__t19).IntVal) - (1))
_ = __local_var_7_6
var __t20 *Constructor_Data_Maybe_Just
{
if (__local_var_7_6.IntVal) == (1) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}}
goto end_branch_20
} else {

}
}
{
if (__local_var_7_6.IntVal) == (2) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}}
goto end_branch_20
} else {

}
}
{
if (__local_var_7_6.IntVal) == (3) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}}
goto end_branch_20
} else {

}
}
{
if (__local_var_7_6.IntVal) == (4) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}}
goto end_branch_20
} else {

}
}
{
if (__local_var_7_6.IntVal) == (5) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}}
goto end_branch_20
} else {

}
}
{
if (__local_var_7_6.IntVal) == (6) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}}
goto end_branch_20
} else {

}
}
{
if (__local_var_7_6.IntVal) == (7) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}}
goto end_branch_20
} else {

}
}
{
if (__local_var_7_6.IntVal) == (8) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}}
goto end_branch_20
} else {

}
}
{
if (__local_var_7_6.IntVal) == (9) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}}
goto end_branch_20
} else {

}
}
{
if (__local_var_7_6.IntVal) == (10) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}}
goto end_branch_20
} else {

}
}
{
if (__local_var_7_6.IntVal) == (11) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}}
goto end_branch_20
} else {

}
}
{
if (__local_var_7_6.IntVal) == (12) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}}
goto end_branch_20
} else {

}
}
{
__t20 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_20:
// TAST (Let): __local_var_7_5 -> *Constructor_Data_Maybe_Just
var __local_var_7_5 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t20)})
var __t21 gopurs_runtime.Value
{
if (__local_var_7_5 == nil) {
__t21 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_21
} else {

}
}
{
if (__local_var_7_5 != nil) {
__t21 = (__local_var_7_5).V0
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
__t22 = (v2_4).V1
}
end_branch_22:
// TAST (Let): l_7_4 -> gopurs_runtime.Value
var l_7_4 gopurs_runtime.Value = gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth((v2_4).V0, __t22))
var __t24 bool
{
if (j_5_1) > (l_7_4.IntVal) {
__t24 = true
goto end_branch_24
} else {

}
}
{
__t24 = false
}
end_branch_24:
// TAST (Let): hi_8_23 -> bool
hi_8_23 := __t24
_ = hi_8_23
var __t65 *Constructor_Data_Maybe_Just
{
if low_6_2 {
// TAST (Let): __local_var_9_33 -> *Constructor_Data_Maybe_Just
var __local_var_9_33 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Data_Date_Date{1, (v2_4).V0, (v2_4).V1, 1})}})})
var __t48 int64
{
var __t_tag_36 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_36.IntVal) == 1908470532) {
__t48 = 1
goto end_branch_48
} else {

}
}
{
var __t_tag_37 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_37.IntVal) == 2455627378) {
__t48 = 2
goto end_branch_48
} else {

}
}
{
var __t_tag_38 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_38.IntVal) == 4162469099) {
__t48 = 3
goto end_branch_48
} else {

}
}
{
var __t_tag_39 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_39.IntVal) == 1692989816) {
__t48 = 4
goto end_branch_48
} else {

}
}
{
var __t_tag_40 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_40.IntVal) == 330658827) {
__t48 = 5
goto end_branch_48
} else {

}
}
{
var __t_tag_41 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_41.IntVal) == 4067355978) {
__t48 = 6
goto end_branch_48
} else {

}
}
{
var __t_tag_42 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_42.IntVal) == 2276710548) {
__t48 = 7
goto end_branch_48
} else {

}
}
{
var __t_tag_43 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_43.IntVal) == 243771071) {
__t48 = 8
goto end_branch_48
} else {

}
}
{
var __t_tag_44 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_44.IntVal) == 215731793) {
__t48 = 9
goto end_branch_48
} else {

}
}
{
var __t_tag_45 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_45.IntVal) == 8639228) {
__t48 = 10
goto end_branch_48
} else {

}
}
{
var __t_tag_46 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_46.IntVal) == 49471444) {
__t48 = 11
goto end_branch_48
} else {

}
}
{
var __t_tag_47 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V1), UnsafePtr: nil}
if (uint32(__t_tag_47.IntVal) == 3889233761) {
__t48 = 12
goto end_branch_48
} else {

}
}
{
__t48 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_48:
// TAST (Let): __local_var_10_35 -> gopurs_runtime.Value
__local_var_10_35 := gopurs_runtime.Int((gopurs_runtime.Int(__t48).IntVal) - (1))
_ = __local_var_10_35
var __t49 *Constructor_Data_Maybe_Just
{
if (__local_var_10_35.IntVal) == (1) {
__t49 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}}
goto end_branch_49
} else {

}
}
{
if (__local_var_10_35.IntVal) == (2) {
__t49 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}}
goto end_branch_49
} else {

}
}
{
if (__local_var_10_35.IntVal) == (3) {
__t49 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}}
goto end_branch_49
} else {

}
}
{
if (__local_var_10_35.IntVal) == (4) {
__t49 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}}
goto end_branch_49
} else {

}
}
{
if (__local_var_10_35.IntVal) == (5) {
__t49 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}}
goto end_branch_49
} else {

}
}
{
if (__local_var_10_35.IntVal) == (6) {
__t49 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}}
goto end_branch_49
} else {

}
}
{
if (__local_var_10_35.IntVal) == (7) {
__t49 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}}
goto end_branch_49
} else {

}
}
{
if (__local_var_10_35.IntVal) == (8) {
__t49 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}}
goto end_branch_49
} else {

}
}
{
if (__local_var_10_35.IntVal) == (9) {
__t49 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}}
goto end_branch_49
} else {

}
}
{
if (__local_var_10_35.IntVal) == (10) {
__t49 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}}
goto end_branch_49
} else {

}
}
{
if (__local_var_10_35.IntVal) == (11) {
__t49 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}}
goto end_branch_49
} else {

}
}
{
if (__local_var_10_35.IntVal) == (12) {
__t49 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}}
goto end_branch_49
} else {

}
}
{
__t49 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_49:
// TAST (Let): pm_10_34 -> *Constructor_Data_Maybe_Just
var pm_10_34 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t49)})
var __t51 gopurs_runtime.Value
{
if (pm_10_34 == nil) {
__t51 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_51
} else {

}
}
{
if (pm_10_34 != nil) {
__t51 = (pm_10_34).V0
goto end_branch_51
} else {

}
}
{
__t51 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_51:
// TAST (Let): m_prime_11_50 -> uint32
m_prime_11_50 := uint32(__t51.IntVal)
_ = m_prime_11_50
var __t61 *Constructor_Data_Maybe_Just
{
var __t55 gopurs_runtime.Value
{
if (pm_10_34 == nil) {
__t55 = gopurs_runtime.Bool(true)
goto end_branch_55
} else {

}
}
{
if (pm_10_34 != nil) {
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
// TAST (Let): __local_var_12_56 -> gopurs_runtime.Value
__local_var_12_56 := gopurs_runtime.Int((gopurs_runtime.Int(gopurs_runtime.Int((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V0).IntVal).IntVal) - (1))
_ = __local_var_12_56
var __t60 *Constructor_Data_Maybe_Just
{
var __t57 bool
{
if (__local_var_12_56.IntVal) < (1) {
__t57 = false
goto end_branch_57
} else {

}
}
{
__t57 = true
}
end_branch_57:
var __t_and_59 bool = false
if __t57 {

var __t58 bool
{
if (__local_var_12_56.IntVal) > (31) {
__t58 = false
goto end_branch_58
} else {

}
}
{
__t58 = true
}
end_branch_58:
__t_and_59 = __t58
}
if __t_and_59 {
__t60 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_12_56.IntVal)}
goto end_branch_60
} else {

}
}
{
__t60 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_60:
__t61 = __t60
goto end_branch_61
} else {

}
}
{
__t61 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V0)}
}
end_branch_61:
// TAST (Let): __local_var_12_54 -> *Constructor_Data_Maybe_Just
var __local_var_12_54 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t61)})
var __t62 *Constructor_Data_Maybe_Just
{
if (__local_var_12_54 != nil) {
__t62 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Date_Date(), gopurs_runtime.Int((__local_var_12_54).V0.IntVal))}
goto end_branch_62
} else {

}
}
{
__t62 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_62:
// TAST (Let): __local_var_12_53 -> *Constructor_Data_Maybe_Just
var __local_var_12_53 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t62)})
var __t63 *Constructor_Data_Maybe_Just
{
if (__local_var_12_53 != nil) {
__t63 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_12_53).V0, gopurs_runtime.Value{Type: 9, IntVal: int64(m_prime_11_50), UnsafePtr: nil})}
goto end_branch_63
} else {

}
}
{
if (__local_var_12_53 == nil) {
__t63 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_63
} else {

}
}
{
__t63 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_63:
// TAST (Let): __local_var_12_52 -> *Constructor_Data_Maybe_Just
var __local_var_12_52 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t63)})
var __t64 *Constructor_Data_Maybe_Just
{
if (__local_var_12_52 != nil) {
__t64 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply((__local_var_12_52).V0, gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth((*Constructor_Data_Date_Date)((__local_var_9_33).V0.UnsafePtr).V0, m_prime_11_50)))}
goto end_branch_64
} else {

}
}
{
if (__local_var_12_52 == nil) {
__t64 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_64
} else {

}
}
{
__t64 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_64:
__t65 = __t64
goto end_branch_65
} else {

}
}
{
if hi_8_23 {
__t65 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_succ__2858180024(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Data_Date_Date{1, (v2_4).V0, (v2_4).V1, l_7_4.IntVal})})))})
goto end_branch_65
} else {

}
}
{
// TAST (Let): __local_var_9_26 -> gopurs_runtime.Value
__local_var_9_26 := gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((v2_4).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((v2_4).V1), UnsafePtr: nil})
_ = __local_var_9_26
var __t31 *Constructor_Data_Maybe_Just
{
var __t28 bool
{
if (gopurs_runtime.Int(j_5_1).IntVal) < (gopurs_runtime.Int(1).IntVal) {
__t28 = false
goto end_branch_28
} else {

}
}
{
__t28 = true
}
end_branch_28:
var __t_and_30 bool = false
if __t28 {

var __t29 bool
{
if (gopurs_runtime.Int(j_5_1).IntVal) > (gopurs_runtime.Int(31).IntVal) {
__t29 = false
goto end_branch_29
} else {

}
}
{
__t29 = true
}
end_branch_29:
__t_and_30 = __t29
}
if __t_and_30 {
__t31 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(j_5_1)}
goto end_branch_31
} else {

}
}
{
__t31 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_31:
// TAST (Let): __local_var_10_27 -> *Constructor_Data_Maybe_Just
var __local_var_10_27 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t31)})
var __t32 *Constructor_Data_Maybe_Just
{
if (__local_var_10_27 != nil) {
__t32 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(__local_var_9_26, gopurs_runtime.Int((__local_var_10_27).V0.IntVal))))}}
goto end_branch_32
} else {

}
}
{
__t32 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_32:
__t65 = __t32
}
end_branch_65:
// TAST (Let): __local_var_9_25 -> *Constructor_Data_Maybe_Just
var __local_var_9_25 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t65)})
var __t67 *Constructor_Data_Maybe_Just
{
if (__local_var_9_25 != nil) {
var __t66 int64
{
if low_6_2 {
__t66 = j_5_1
goto end_branch_66
} else {

}
}
{
if hi_8_23 {
__t66 = ((j_5_1) - (l_7_4.IntVal)) - (1)
goto end_branch_66
} else {

}
}
{
__t66 = 0
}
end_branch_66:
v1_3_loop = gopurs_runtime.Int(__t66)
v2_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((__local_var_9_25).V0)
continue adj_2_0_0
__t67 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
goto end_branch_67
} else {

}
}
{
if (__local_var_9_25 == nil) {
__t67 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_67
} else {

}
}
{
__t67 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_67:
__t68 = __t67
}
end_branch_68:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t68)}
}
}()
})
})
// TAST (Let): __local_var_3_69 -> *Constructor_Data_Maybe_Just
__local_var_3_69 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Int_fromNumber(), gopurs_runtime.Float(v_0)))
_ = __local_var_3_69
var __t70 *Constructor_Data_Maybe_Just
{
if (__local_var_3_69 != nil) {
__t70 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(adj_2_0_0, gopurs_runtime.Int((__local_var_3_69).V0.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(date_1)}))
goto end_branch_70
} else {

}
}
{
if (__local_var_3_69 == nil) {
__t70 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_70
} else {

}
}
{
__t70 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_70:
return __t70
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

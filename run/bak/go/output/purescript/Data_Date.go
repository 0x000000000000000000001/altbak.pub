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
		cache_Data_Date_toEnum = gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum")
	})
	return cache_Data_Date_toEnum
}

var cache_Data_Date_ordMaybe gopurs_runtime.Value
var once_Data_Date_ordMaybe sync.Once
func Get_Data_Date_ordMaybe() gopurs_runtime.Value {
	once_Data_Date_ordMaybe.Do(func() {
		cache_Data_Date_ordMaybe = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_1 -> gopurs_runtime.Value
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_ordDay(), "Eq0"), gopurs_runtime.Value{})
_ = __local_var_0_1
// TAST (Let): eqMaybe1_0_0 -> gopurs_runtime.Value
eqMaybe1_0_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t2 bool
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t3 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "eq"), (*Constructor_Data_Maybe_Just)(x_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_2.UnsafePtr).V0).IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
})
}))
_ = eqMaybe1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_0_0
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 uint32
{
if (x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr == nil) {
var __t4 uint32
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t4 = 902936544
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr == nil) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136 && x_1.UnsafePtr != nil)) && ((y_2.Type == 9 && y_2.IntVal == 930809136 && y_2.UnsafePtr != nil)) {
__t5 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Date_Component_ordDay(), "compare"), (*Constructor_Data_Maybe_Just)(x_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_2.UnsafePtr).V0).IntVal)
goto end_branch_5
} else {

}
}
{
__t5 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
})
}))))}
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
		cache_Data_Date_showDate = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
}))
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
		cache_Data_Date_eqDate = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V0) == ((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V0)) && ((gopurs_runtime.Apply2(Get_Data_Eq_eq__3887832182(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V1), UnsafePtr: nil}).IntVal) != (0))) && (((*Constructor_Data_Date_Date)(x_0.UnsafePtr).V2) == ((*Constructor_Data_Date_Date)(y_1.UnsafePtr).V2)))
})
}))
	})
	return cache_Data_Date_eqDate
}

var cache_Data_Date_ordDate gopurs_runtime.Value
var once_Data_Date_ordDate sync.Once
func Get_Data_Date_ordDate() gopurs_runtime.Value {
	once_Data_Date_ordDate.Do(func() {
		cache_Data_Date_ordDate = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((((*Constructor_Data_Date_Date)(x_1.UnsafePtr).V0) == ((*Constructor_Data_Date_Date)(y_2.UnsafePtr).V0)) && ((gopurs_runtime.Apply2(Get_Data_Eq_eq__3887832182(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(x_1.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(y_2.UnsafePtr).V1), UnsafePtr: nil}).IntVal) != (0))) && (((*Constructor_Data_Date_Date)(x_1.UnsafePtr).V2) == ((*Constructor_Data_Date_Date)(y_2.UnsafePtr).V2)))
})
}))
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
}))
	})
	return cache_Data_Date_ordDate
}

var cache_Data_Date_enumDate gopurs_runtime.Value
var once_Data_Date_enumDate sync.Once
func Get_Data_Date_enumDate() gopurs_runtime.Value {
	once_Data_Date_enumDate.Do(func() {
		cache_Data_Date_enumDate = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((((*Constructor_Data_Date_Date)(x_2.UnsafePtr).V0) == ((*Constructor_Data_Date_Date)(y_3.UnsafePtr).V0)) && ((gopurs_runtime.Apply2(Get_Data_Eq_eq__3887832182(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(x_2.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(y_3.UnsafePtr).V1), UnsafePtr: nil}).IntVal) != (0))) && (((*Constructor_Data_Date_Date)(x_2.UnsafePtr).V2) == ((*Constructor_Data_Date_Date)(y_3.UnsafePtr).V2)))
})
}))
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Date_Date)(x_1.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Data_Date_Date)(y_2.UnsafePtr).V0))
_ = v_3_0
var __t3 uint32
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(Get_Data_Ord_compare__696857420(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(x_1.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(y_2.UnsafePtr).V1), UnsafePtr: nil})
_ = v1_4_1
var __t2 uint32
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t2 = 1527465420
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t2 = 380165415
goto end_branch_2
} else {

}
}
{
__t2 = uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Date_Date)(x_1.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Data_Date_Date)(y_2.UnsafePtr).V2)).IntVal)
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pm_1_4 -> *Constructor_Data_Maybe_Just
pm_1_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "fromEnum"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}).IntVal) - (1))))
_ = pm_1_4
// TAST (Let): pd_2_5 -> *Constructor_Data_Maybe_Just
pd_2_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "fromEnum"), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V2)).IntVal) - (1))))
_ = pd_2_5
var __t9 uint32
{
var __t7 gopurs_runtime.Value
{
if (pd_2_5 == nil) {
__t7 = gopurs_runtime.Bool(true)
goto end_branch_7
} else {

}
}
{
if (pd_2_5 != nil) {
__t7 = gopurs_runtime.Bool(false)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
if (__t7.IntVal) != (0) {
var __t8 gopurs_runtime.Value
{
if (pm_1_4 == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_8
} else {

}
}
{
if (pm_1_4 != nil) {
__t8 = (pm_1_4).V0
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t9 = uint32(__t8.IntVal)
goto end_branch_9
} else {

}
}
{
__t9 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
}
end_branch_9:
// TAST (Let): m_prime_3_6 -> gopurs_runtime.Value
var m_prime_3_6 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64(__t9), UnsafePtr: nil}
// TAST (Let): l_4_10 -> gopurs_runtime.Value
var l_4_10 gopurs_runtime.Value = gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0, uint32(m_prime_3_6.IntVal)))
var __t17 *Constructor_Data_Maybe_Just
{
var __t14 gopurs_runtime.Value
{
if (pd_2_5 == nil) {
__t14 = gopurs_runtime.Bool(true)
goto end_branch_14
} else {

}
}
{
if (pd_2_5 != nil) {
__t14 = gopurs_runtime.Bool(false)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
var __t_and_16 bool = false
if (__t14.IntVal) != (0) {

var __t15 gopurs_runtime.Value
{
if (pm_1_4 == nil) {
__t15 = gopurs_runtime.Bool(true)
goto end_branch_15
} else {

}
}
{
if (pm_1_4 != nil) {
__t15 = gopurs_runtime.Bool(false)
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
__t_and_16 = (__t15.IntVal) != (0)
}
if __t_and_16 {
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "fromEnum"), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0)).IntVal) - (1)))))})
goto end_branch_17
} else {

}
}
{
__t17 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0)}
}
end_branch_17:
// TAST (Let): __local_var_5_13 -> *Constructor_Data_Maybe_Just
var __local_var_5_13 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t17)})
var __t18 *Constructor_Data_Maybe_Just
{
if (__local_var_5_13 != nil) {
__t18 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Date_Date(), gopurs_runtime.Int((__local_var_5_13).V0.IntVal))}
goto end_branch_18
} else {

}
}
{
__t18 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_18:
// TAST (Let): __local_var_5_12 -> *Constructor_Data_Maybe_Just
var __local_var_5_12 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t18)})
var __t19 *Constructor_Data_Maybe_Just
{
if (__local_var_5_12 != nil) {
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_5_12).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(m_prime_3_6.IntVal)), UnsafePtr: nil}})}))
goto end_branch_19
} else {

}
}
{
if (__local_var_5_12 == nil) {
__t19 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_19:
// TAST (Let): __local_var_5_11 -> *Constructor_Data_Maybe_Just
__local_var_5_11 := __t19
_ = __local_var_5_11
var __t22 *Constructor_Data_Maybe_Just
{
var __t21 gopurs_runtime.Value
{
if (pd_2_5 == nil) {
__t21 = gopurs_runtime.Bool(true)
goto end_branch_21
} else {

}
}
{
if (pd_2_5 != nil) {
__t21 = gopurs_runtime.Bool(false)
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
if (__t21.IntVal) != (0) {
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(l_4_10.IntVal)})})
goto end_branch_22
} else {

}
}
{
__t22 = pd_2_5
}
end_branch_22:
// TAST (Let): __local_var_6_20 -> *Constructor_Data_Maybe_Just
__local_var_6_20 := __t22
_ = __local_var_6_20
var __t23 *Constructor_Data_Maybe_Just
{
if (__local_var_5_11 != nil) {
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_5_11).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_20)}))
goto end_branch_23
} else {

}
}
{
if (__local_var_5_11 == nil) {
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_23:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t23)}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): sm_1_24 -> *Constructor_Data_Maybe_Just
sm_1_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "fromEnum"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1), UnsafePtr: nil}).IntVal) + (1))))
_ = sm_1_24
// TAST (Let): v1_2_25 -> *Constructor_Data_Maybe_Just
v1_2_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "fromEnum"), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V2)).IntVal) + (1))))
_ = v1_2_25
var __t30 *Constructor_Data_Maybe_Just
{
var __t29 bool
{
// TAST (Let): __local_var_3_27 -> *Constructor_Data_Maybe_Just
var __local_var_3_27 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0, (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1))})})
var __t28 uint32
{
if (v1_2_25 == nil) {
__t28 = 1527465420
goto end_branch_28
} else {

}
}
{
if (v1_2_25 != nil) {
__t28 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Date_Component_ordDay(), "compare"), (v1_2_25).V0, (__local_var_3_27).V0).IntVal)
goto end_branch_28
} else {

}
}
{
__t28 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_28:
if (__t28 == 380165415) {
__t29 = true
goto end_branch_29
} else {

}
}
{
__t29 = false
}
end_branch_29:
if __t29 {
__t30 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_30
} else {

}
}
{
__t30 = v1_2_25
}
end_branch_30:
// TAST (Let): sd_3_26 -> *Constructor_Data_Maybe_Just
sd_3_26 := __t30
_ = sd_3_26
var __t37 *Constructor_Data_Maybe_Just
{
var __t34 gopurs_runtime.Value
{
if (sd_3_26 == nil) {
__t34 = gopurs_runtime.Bool(true)
goto end_branch_34
} else {

}
}
{
if (sd_3_26 != nil) {
__t34 = gopurs_runtime.Bool(false)
goto end_branch_34
} else {

}
}
{
__t34 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_34:
var __t_and_36 bool = false
if (__t34.IntVal) != (0) {

var __t35 gopurs_runtime.Value
{
if (sm_1_24 == nil) {
__t35 = gopurs_runtime.Bool(true)
goto end_branch_35
} else {

}
}
{
if (sm_1_24 != nil) {
__t35 = gopurs_runtime.Bool(false)
goto end_branch_35
} else {

}
}
{
__t35 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_35:
__t_and_36 = (__t35.IntVal) != (0)
}
if __t_and_36 {
__t37 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "fromEnum"), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0)).IntVal) + (1)))))})
goto end_branch_37
} else {

}
}
{
__t37 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_0.UnsafePtr).V0)}
}
end_branch_37:
// TAST (Let): __local_var_4_33 -> *Constructor_Data_Maybe_Just
var __local_var_4_33 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t37)})
var __t38 *Constructor_Data_Maybe_Just
{
if (__local_var_4_33 != nil) {
__t38 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Date_Date(), gopurs_runtime.Int((__local_var_4_33).V0.IntVal))}
goto end_branch_38
} else {

}
}
{
__t38 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_38:
// TAST (Let): __local_var_4_32 -> *Constructor_Data_Maybe_Just
var __local_var_4_32 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t38)})
var __t42 uint32
{
var __t40 gopurs_runtime.Value
{
if (sd_3_26 == nil) {
__t40 = gopurs_runtime.Bool(true)
goto end_branch_40
} else {

}
}
{
if (sd_3_26 != nil) {
__t40 = gopurs_runtime.Bool(false)
goto end_branch_40
} else {

}
}
{
__t40 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_40:
if (__t40.IntVal) != (0) {
var __t41 gopurs_runtime.Value
{
if (sm_1_24 == nil) {
__t41 = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
goto end_branch_41
} else {

}
}
{
if (sm_1_24 != nil) {
__t41 = (sm_1_24).V0
goto end_branch_41
} else {

}
}
{
__t41 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_41:
__t42 = uint32(__t41.IntVal)
goto end_branch_42
} else {

}
}
{
__t42 = (*Constructor_Data_Date_Date)(v_0.UnsafePtr).V1
}
end_branch_42:
// TAST (Let): __local_var_5_39 -> *Constructor_Data_Maybe_Just
var __local_var_5_39 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(__t42), UnsafePtr: nil}})})
var __t43 *Constructor_Data_Maybe_Just
{
if (__local_var_4_32 != nil) {
__t43 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_4_32).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_39)}))
goto end_branch_43
} else {

}
}
{
if (__local_var_4_32 == nil) {
__t43 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_43
} else {

}
}
{
__t43 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_43:
// TAST (Let): __local_var_4_31 -> *Constructor_Data_Maybe_Just
__local_var_4_31 := __t43
_ = __local_var_4_31
var __t46 *Constructor_Data_Maybe_Just
{
var __t45 gopurs_runtime.Value
{
if (sd_3_26 == nil) {
__t45 = gopurs_runtime.Bool(true)
goto end_branch_45
} else {

}
}
{
if (sd_3_26 != nil) {
__t45 = gopurs_runtime.Bool(false)
goto end_branch_45
} else {

}
}
{
__t45 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_45:
if (__t45.IntVal) != (0) {
__t46 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2099864294(), gopurs_runtime.Int(1)))
goto end_branch_46
} else {

}
}
{
__t46 = sd_3_26
}
end_branch_46:
// TAST (Let): __local_var_5_44 -> *Constructor_Data_Maybe_Just
__local_var_5_44 := __t46
_ = __local_var_5_44
var __t47 *Constructor_Data_Maybe_Just
{
if (__local_var_4_31 != nil) {
__t47 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_4_31).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_44)}))
goto end_branch_47
} else {

}
}
{
if (__local_var_4_31 == nil) {
__t47 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_47
} else {

}
}
{
__t47 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_47:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t47)}
}))
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
		cache_Data_Date_boundedDate = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((((*Constructor_Data_Date_Date)(x_2.UnsafePtr).V0) == ((*Constructor_Data_Date_Date)(y_3.UnsafePtr).V0)) && ((gopurs_runtime.Apply2(Get_Data_Eq_eq__3887832182(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(x_2.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(y_3.UnsafePtr).V1), UnsafePtr: nil}).IntVal) != (0))) && (((*Constructor_Data_Date_Date)(x_2.UnsafePtr).V2) == ((*Constructor_Data_Date_Date)(y_3.UnsafePtr).V2)))
})
}))
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Date_Date)(x_1.UnsafePtr).V0), gopurs_runtime.Int((*Constructor_Data_Date_Date)(y_2.UnsafePtr).V0))
_ = v_3_0
var __t3 uint32
{
if (uint32(v_3_0.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(Get_Data_Ord_compare__696857420(), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(x_1.UnsafePtr).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(y_2.UnsafePtr).V1), UnsafePtr: nil})
_ = v1_4_1
var __t2 uint32
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t2 = 1527465420
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t2 = 380165415
goto end_branch_2
} else {

}
}
{
__t2 = uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int((*Constructor_Data_Date_Date)(x_1.UnsafePtr).V2), gopurs_runtime.Int((*Constructor_Data_Date_Date)(y_2.UnsafePtr).V2)).IntVal)
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Data_Date_Date{1, gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedYear(), "bottom").IntVal, uint32(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedMonth(), "bottom").IntVal), gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedDay(), "bottom").IntVal})}, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Data_Date_Date{1, gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedYear(), "top").IntVal, uint32(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedMonth(), "top").IntVal), gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedDay(), "top").IntVal})})
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
var __t2 int64
{
if (m_1 == 1908470532) {
__t2 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 2455627378) {
var __t1 int64
{
if Call_Data_Date_isLeapYear(y_0) {
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(29))).IntVal
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(28))).IntVal
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (m_1 == 4162469099) {
__t2 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 1692989816) {
__t2 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 330658827) {
__t2 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 4067355978) {
__t2 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 2276710548) {
__t2 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 243771071) {
__t2 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 215731793) {
__t2 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 8639228) {
__t2 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 49471444) {
__t2 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(30))).IntVal
goto end_branch_2
} else {

}
}
{
if (m_1 == 3889233761) {
__t2 = gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int(31))).IntVal
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_2:
return __t2
}

func Call_Data_Date_pred(v_0_loop *Constructor_Data_Date_Date) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Date_Date = v_0_loop
_ = v_0
// TAST (Let): pm_1_0 -> *Constructor_Data_Maybe_Just
pm_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "fromEnum"), gopurs_runtime.Value{Type: 9, IntVal: int64((v_0).V1), UnsafePtr: nil}).IntVal) - (1))))
_ = pm_1_0
// TAST (Let): pd_2_1 -> *Constructor_Data_Maybe_Just
pd_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "fromEnum"), gopurs_runtime.Int((v_0).V2)).IntVal) - (1))))
_ = pd_2_1
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
__t4 = (pm_1_0).V0
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
__t5 = (v_0).V1
}
end_branch_5:
// TAST (Let): m_prime_3_2 -> gopurs_runtime.Value
var m_prime_3_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
// TAST (Let): l_4_6 -> gopurs_runtime.Value
var l_4_6 gopurs_runtime.Value = gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth((v_0).V0, uint32(m_prime_3_2.IntVal)))
var __t13 *Constructor_Data_Maybe_Just
{
var __t10 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t10 = gopurs_runtime.Bool(true)
goto end_branch_10
} else {

}
}
{
if (pd_2_1 != nil) {
__t10 = gopurs_runtime.Bool(false)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
var __t_and_12 bool = false
if (__t10.IntVal) != (0) {

var __t11 gopurs_runtime.Value
{
if (pm_1_0 == nil) {
__t11 = gopurs_runtime.Bool(true)
goto end_branch_11
} else {

}
}
{
if (pm_1_0 != nil) {
__t11 = gopurs_runtime.Bool(false)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
__t_and_12 = (__t11.IntVal) != (0)
}
if __t_and_12 {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "fromEnum"), gopurs_runtime.Int((v_0).V0)).IntVal) - (1)))))})
goto end_branch_13
} else {

}
}
{
__t13 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((v_0).V0)}
}
end_branch_13:
// TAST (Let): __local_var_5_9 -> *Constructor_Data_Maybe_Just
var __local_var_5_9 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t13)})
var __t14 *Constructor_Data_Maybe_Just
{
if (__local_var_5_9 != nil) {
__t14 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Date_Date(), gopurs_runtime.Int((__local_var_5_9).V0.IntVal))}
goto end_branch_14
} else {

}
}
{
__t14 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_14:
// TAST (Let): __local_var_5_8 -> *Constructor_Data_Maybe_Just
var __local_var_5_8 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t14)})
var __t15 *Constructor_Data_Maybe_Just
{
if (__local_var_5_8 != nil) {
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_5_8).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(m_prime_3_2.IntVal)), UnsafePtr: nil}})}))
goto end_branch_15
} else {

}
}
{
if (__local_var_5_8 == nil) {
__t15 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_15:
// TAST (Let): __local_var_5_7 -> *Constructor_Data_Maybe_Just
__local_var_5_7 := __t15
_ = __local_var_5_7
var __t18 *Constructor_Data_Maybe_Just
{
var __t17 gopurs_runtime.Value
{
if (pd_2_1 == nil) {
__t17 = gopurs_runtime.Bool(true)
goto end_branch_17
} else {

}
}
{
if (pd_2_1 != nil) {
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
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(l_4_6.IntVal)})})
goto end_branch_18
} else {

}
}
{
__t18 = pd_2_1
}
end_branch_18:
// TAST (Let): __local_var_6_16 -> *Constructor_Data_Maybe_Just
__local_var_6_16 := __t18
_ = __local_var_6_16
var __t19 *Constructor_Data_Maybe_Just
{
if (__local_var_5_7 != nil) {
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_5_7).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_16)}))
goto end_branch_19
} else {

}
}
{
if (__local_var_5_7 == nil) {
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_19:
return __t19
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
_ = adj_2_0_0
adj_2_0_0 = gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t38 *Constructor_Data_Maybe_Just
{
if (v1_3.IntVal) == (0) {
__t38 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](v2_4))}})})
goto end_branch_38
} else {

}
}
{
// TAST (Let): j_5_1 -> int64
j_5_1 := (v1_3.IntVal) + ((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V2)
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
var __t7 uint32
{
if low_6_2 {
// TAST (Let): __local_var_7_5 -> *Constructor_Data_Maybe_Just
__local_var_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "fromEnum"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil}).IntVal) - (1))))
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
__t6 = (__local_var_7_5).V0
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
// TAST (Let): l_7_4 -> gopurs_runtime.Value
var l_7_4 gopurs_runtime.Value = gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0, __t7))
var __t9 bool
{
if (j_5_1) > (l_7_4.IntVal) {
__t9 = true
goto end_branch_9
} else {

}
}
{
__t9 = false
}
end_branch_9:
// TAST (Let): hi_8_8 -> bool
hi_8_8 := __t9
_ = hi_8_8
var __t36 *Constructor_Data_Maybe_Just
{
if low_6_2 {
// TAST (Let): __local_var_9_13 -> gopurs_runtime.Value
__local_var_9_13 := gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil})
_ = __local_var_9_13
// TAST (Let): __local_var_10_14 -> *Constructor_Data_Maybe_Just
__local_var_10_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2099864294(), gopurs_runtime.Int(1)))
_ = __local_var_10_14
var __t15 *Constructor_Data_Maybe_Just
{
if (__local_var_10_14 != nil) {
__t15 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(__local_var_9_13, gopurs_runtime.Int((__local_var_10_14).V0.IntVal))))}}
goto end_branch_15
} else {

}
}
{
__t15 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_15:
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_bindMaybe(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t15)}, gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pm_10_16 -> *Constructor_Data_Maybe_Just
pm_10_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumMonth(), "fromEnum"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v_9.UnsafePtr).V1), UnsafePtr: nil}).IntVal) - (1))))
_ = pm_10_16
// TAST (Let): pd_11_17 -> *Constructor_Data_Maybe_Just
pd_11_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "fromEnum"), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_9.UnsafePtr).V2)).IntVal) - (1))))
_ = pd_11_17
var __t21 uint32
{
var __t19 gopurs_runtime.Value
{
if (pd_11_17 == nil) {
__t19 = gopurs_runtime.Bool(true)
goto end_branch_19
} else {

}
}
{
if (pd_11_17 != nil) {
__t19 = gopurs_runtime.Bool(false)
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
if (__t19.IntVal) != (0) {
var __t20 gopurs_runtime.Value
{
if (pm_10_16 == nil) {
__t20 = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
goto end_branch_20
} else {

}
}
{
if (pm_10_16 != nil) {
__t20 = (pm_10_16).V0
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
__t21 = uint32(__t20.IntVal)
goto end_branch_21
} else {

}
}
{
__t21 = (*Constructor_Data_Date_Date)(v_9.UnsafePtr).V1
}
end_branch_21:
// TAST (Let): m_prime_12_18 -> gopurs_runtime.Value
var m_prime_12_18 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: int64(__t21), UnsafePtr: nil}
// TAST (Let): l_13_22 -> gopurs_runtime.Value
var l_13_22 gopurs_runtime.Value = gopurs_runtime.Int(Call_Data_Date_lastDayOfMonth((*Constructor_Data_Date_Date)(v_9.UnsafePtr).V0, uint32(m_prime_12_18.IntVal)))
var __t29 *Constructor_Data_Maybe_Just
{
var __t26 gopurs_runtime.Value
{
if (pd_11_17 == nil) {
__t26 = gopurs_runtime.Bool(true)
goto end_branch_26
} else {

}
}
{
if (pd_11_17 != nil) {
__t26 = gopurs_runtime.Bool(false)
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
var __t_and_28 bool = false
if (__t26.IntVal) != (0) {

var __t27 gopurs_runtime.Value
{
if (pm_10_16 == nil) {
__t27 = gopurs_runtime.Bool(true)
goto end_branch_27
} else {

}
}
{
if (pm_10_16 != nil) {
__t27 = gopurs_runtime.Bool(false)
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
__t_and_28 = (__t27.IntVal) != (0)
}
if __t_and_28 {
__t29 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumDay(), "fromEnum"), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_9.UnsafePtr).V0)).IntVal) - (1)))))})
goto end_branch_29
} else {

}
}
{
__t29 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Date_Date)(v_9.UnsafePtr).V0)}
}
end_branch_29:
// TAST (Let): __local_var_14_25 -> *Constructor_Data_Maybe_Just
var __local_var_14_25 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t29)})
var __t30 *Constructor_Data_Maybe_Just
{
if (__local_var_14_25 != nil) {
__t30 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Date_Date(), gopurs_runtime.Int((__local_var_14_25).V0.IntVal))}
goto end_branch_30
} else {

}
}
{
__t30 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_30:
// TAST (Let): __local_var_14_24 -> *Constructor_Data_Maybe_Just
var __local_var_14_24 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t30)})
var __t31 *Constructor_Data_Maybe_Just
{
if (__local_var_14_24 != nil) {
__t31 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_14_24).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(m_prime_12_18.IntVal)), UnsafePtr: nil}})}))
goto end_branch_31
} else {

}
}
{
if (__local_var_14_24 == nil) {
__t31 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_31
} else {

}
}
{
__t31 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_31:
// TAST (Let): __local_var_14_23 -> *Constructor_Data_Maybe_Just
__local_var_14_23 := __t31
_ = __local_var_14_23
var __t34 *Constructor_Data_Maybe_Just
{
var __t33 gopurs_runtime.Value
{
if (pd_11_17 == nil) {
__t33 = gopurs_runtime.Bool(true)
goto end_branch_33
} else {

}
}
{
if (pd_11_17 != nil) {
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
__t34 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(l_13_22.IntVal)})})
goto end_branch_34
} else {

}
}
{
__t34 = pd_11_17
}
end_branch_34:
// TAST (Let): __local_var_15_32 -> *Constructor_Data_Maybe_Just
__local_var_15_32 := __t34
_ = __local_var_15_32
var __t35 *Constructor_Data_Maybe_Just
{
if (__local_var_14_23 != nil) {
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), (__local_var_14_23).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_15_32)}))
goto end_branch_35
} else {

}
}
{
if (__local_var_14_23 == nil) {
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_35:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t35)}
}))))})
goto end_branch_36
} else {

}
}
{
if hi_8_8 {
__t36 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_succ__2858180024(), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(&Constructor_Data_Date_Date{1, (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0, (*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1, l_7_4.IntVal})})))})
goto end_branch_36
} else {

}
}
{
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.Apply2(Get_Data_Date_Date(), gopurs_runtime.Int((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_Data_Date_Date)(v2_4.UnsafePtr).V1), UnsafePtr: nil})
_ = __local_var_9_10
// TAST (Let): __local_var_10_11 -> *Constructor_Data_Maybe_Just
__local_var_10_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Enum_toEnum__2099864294(), gopurs_runtime.Int(j_5_1)))
_ = __local_var_10_11
var __t12 *Constructor_Data_Maybe_Just
{
if (__local_var_10_11 != nil) {
__t12 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(__local_var_9_10, gopurs_runtime.Int((__local_var_10_11).V0.IntVal))))}}
goto end_branch_12
} else {

}
}
{
__t12 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_12:
__t36 = __t12
}
end_branch_36:
var __t37 int64
{
if low_6_2 {
__t37 = j_5_1
goto end_branch_37
} else {

}
}
{
if hi_8_8 {
__t37 = ((j_5_1) - (l_7_4.IntVal)) - (1)
goto end_branch_37
} else {

}
}
{
__t37 = 0
}
end_branch_37:
__t38 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_bindMaybe(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t36)}, gopurs_runtime.Apply(adj_2_0_0, gopurs_runtime.Int(__t37))))
}
end_branch_38:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t38)}
})
})
// TAST (Let): __local_var_3_39 -> *Constructor_Data_Maybe_Just
__local_var_3_39 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Int_fromNumber(), gopurs_runtime.Float(v_0)))
_ = __local_var_3_39
var __t40 *Constructor_Data_Maybe_Just
{
if (__local_var_3_39 != nil) {
__t40 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(adj_2_0_0, gopurs_runtime.Int((__local_var_3_39).V0.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(date_1)}))
goto end_branch_40
} else {

}
}
{
if (__local_var_3_39 == nil) {
__t40 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_40
} else {

}
}
{
__t40 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_40:
return __t40
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

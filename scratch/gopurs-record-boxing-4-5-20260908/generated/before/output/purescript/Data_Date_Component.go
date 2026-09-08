package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Date_Component_Year gopurs_runtime.Value
var once_Data_Date_Component_Year sync.Once
func Get_Data_Date_Component_Year() gopurs_runtime.Value {
	once_Data_Date_Component_Year.Do(func() {
		cache_Data_Date_Component_Year = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Date_Component_Year(x_0_box.IntVal))
})
	})
	return cache_Data_Date_Component_Year
}

var cache_Data_Date_Component_Monday gopurs_runtime.Value
var once_Data_Date_Component_Monday sync.Once
func Get_Data_Date_Component_Monday() gopurs_runtime.Value {
	once_Data_Date_Component_Monday.Do(func() {
		cache_Data_Date_Component_Monday = gopurs_runtime.Value{Type: 9, IntVal: int64(2900196686), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_Monday
}

var cache_Data_Date_Component_Tuesday gopurs_runtime.Value
var once_Data_Date_Component_Tuesday sync.Once
func Get_Data_Date_Component_Tuesday() gopurs_runtime.Value {
	once_Data_Date_Component_Tuesday.Do(func() {
		cache_Data_Date_Component_Tuesday = gopurs_runtime.Value{Type: 9, IntVal: int64(20457557), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_Tuesday
}

var cache_Data_Date_Component_Wednesday gopurs_runtime.Value
var once_Data_Date_Component_Wednesday sync.Once
func Get_Data_Date_Component_Wednesday() gopurs_runtime.Value {
	once_Data_Date_Component_Wednesday.Do(func() {
		cache_Data_Date_Component_Wednesday = gopurs_runtime.Value{Type: 9, IntVal: int64(4227105004), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_Wednesday
}

var cache_Data_Date_Component_Thursday gopurs_runtime.Value
var once_Data_Date_Component_Thursday sync.Once
func Get_Data_Date_Component_Thursday() gopurs_runtime.Value {
	once_Data_Date_Component_Thursday.Do(func() {
		cache_Data_Date_Component_Thursday = gopurs_runtime.Value{Type: 9, IntVal: int64(3818857258), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_Thursday
}

var cache_Data_Date_Component_Friday gopurs_runtime.Value
var once_Data_Date_Component_Friday sync.Once
func Get_Data_Date_Component_Friday() gopurs_runtime.Value {
	once_Data_Date_Component_Friday.Do(func() {
		cache_Data_Date_Component_Friday = gopurs_runtime.Value{Type: 9, IntVal: int64(2946274527), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_Friday
}

var cache_Data_Date_Component_Saturday gopurs_runtime.Value
var once_Data_Date_Component_Saturday sync.Once
func Get_Data_Date_Component_Saturday() gopurs_runtime.Value {
	once_Data_Date_Component_Saturday.Do(func() {
		cache_Data_Date_Component_Saturday = gopurs_runtime.Value{Type: 9, IntVal: int64(1070786179), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_Saturday
}

var cache_Data_Date_Component_Sunday gopurs_runtime.Value
var once_Data_Date_Component_Sunday sync.Once
func Get_Data_Date_Component_Sunday() gopurs_runtime.Value {
	once_Data_Date_Component_Sunday.Do(func() {
		cache_Data_Date_Component_Sunday = gopurs_runtime.Value{Type: 9, IntVal: int64(1326716170), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_Sunday
}

var cache_Data_Date_Component_January gopurs_runtime.Value
var once_Data_Date_Component_January sync.Once
func Get_Data_Date_Component_January() gopurs_runtime.Value {
	once_Data_Date_Component_January.Do(func() {
		cache_Data_Date_Component_January = gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_January
}

var cache_Data_Date_Component_February gopurs_runtime.Value
var once_Data_Date_Component_February sync.Once
func Get_Data_Date_Component_February() gopurs_runtime.Value {
	once_Data_Date_Component_February.Do(func() {
		cache_Data_Date_Component_February = gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_February
}

var cache_Data_Date_Component_March gopurs_runtime.Value
var once_Data_Date_Component_March sync.Once
func Get_Data_Date_Component_March() gopurs_runtime.Value {
	once_Data_Date_Component_March.Do(func() {
		cache_Data_Date_Component_March = gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_March
}

var cache_Data_Date_Component_April gopurs_runtime.Value
var once_Data_Date_Component_April sync.Once
func Get_Data_Date_Component_April() gopurs_runtime.Value {
	once_Data_Date_Component_April.Do(func() {
		cache_Data_Date_Component_April = gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_April
}

var cache_Data_Date_Component_May gopurs_runtime.Value
var once_Data_Date_Component_May sync.Once
func Get_Data_Date_Component_May() gopurs_runtime.Value {
	once_Data_Date_Component_May.Do(func() {
		cache_Data_Date_Component_May = gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_May
}

var cache_Data_Date_Component_June gopurs_runtime.Value
var once_Data_Date_Component_June sync.Once
func Get_Data_Date_Component_June() gopurs_runtime.Value {
	once_Data_Date_Component_June.Do(func() {
		cache_Data_Date_Component_June = gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_June
}

var cache_Data_Date_Component_July gopurs_runtime.Value
var once_Data_Date_Component_July sync.Once
func Get_Data_Date_Component_July() gopurs_runtime.Value {
	once_Data_Date_Component_July.Do(func() {
		cache_Data_Date_Component_July = gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_July
}

var cache_Data_Date_Component_August gopurs_runtime.Value
var once_Data_Date_Component_August sync.Once
func Get_Data_Date_Component_August() gopurs_runtime.Value {
	once_Data_Date_Component_August.Do(func() {
		cache_Data_Date_Component_August = gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_August
}

var cache_Data_Date_Component_September gopurs_runtime.Value
var once_Data_Date_Component_September sync.Once
func Get_Data_Date_Component_September() gopurs_runtime.Value {
	once_Data_Date_Component_September.Do(func() {
		cache_Data_Date_Component_September = gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_September
}

var cache_Data_Date_Component_October gopurs_runtime.Value
var once_Data_Date_Component_October sync.Once
func Get_Data_Date_Component_October() gopurs_runtime.Value {
	once_Data_Date_Component_October.Do(func() {
		cache_Data_Date_Component_October = gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_October
}

var cache_Data_Date_Component_November gopurs_runtime.Value
var once_Data_Date_Component_November sync.Once
func Get_Data_Date_Component_November() gopurs_runtime.Value {
	once_Data_Date_Component_November.Do(func() {
		cache_Data_Date_Component_November = gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_November
}

var cache_Data_Date_Component_December gopurs_runtime.Value
var once_Data_Date_Component_December sync.Once
func Get_Data_Date_Component_December() gopurs_runtime.Value {
	once_Data_Date_Component_December.Do(func() {
		cache_Data_Date_Component_December = gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}
	})
	return cache_Data_Date_Component_December
}

var cache_Data_Date_Component_Day gopurs_runtime.Value
var once_Data_Date_Component_Day sync.Once
func Get_Data_Date_Component_Day() gopurs_runtime.Value {
	once_Data_Date_Component_Day.Do(func() {
		cache_Data_Date_Component_Day = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Date_Component_Day(x_0_box.IntVal))
})
	})
	return cache_Data_Date_Component_Day
}

var cache_Data_Date_Component_showYear gopurs_runtime.Value
var once_Data_Date_Component_showYear sync.Once
func Get_Data_Date_Component_showYear() gopurs_runtime.Value {
	once_Data_Date_Component_showYear.Do(func() {
		cache_Data_Date_Component_showYear = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_1636311157_1386611502((&Constructor_Data_Show_Show[int64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Year ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(v_0.IntVal)).StrVal())) + (")"))
})})))}
	})
	return cache_Data_Date_Component_showYear
}

var cache_Data_Date_Component_showWeekday gopurs_runtime.Value
var once_Data_Date_Component_showWeekday sync.Once
func Get_Data_Date_Component_showWeekday() gopurs_runtime.Value {
	once_Data_Date_Component_showWeekday.Do(func() {
		cache_Data_Date_Component_showWeekday = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_1209612131_1386611502((&Constructor_Data_Show_Show[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 string
{
var __t_tag_0 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_0) == 2900196686) {
__t7 = "Monday"
goto end_branch_7
} else {

}
}
{
var __t_tag_1 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_1) == 20457557) {
__t7 = "Tuesday"
goto end_branch_7
} else {

}
}
{
var __t_tag_2 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_2) == 4227105004) {
__t7 = "Wednesday"
goto end_branch_7
} else {

}
}
{
var __t_tag_3 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_3) == 3818857258) {
__t7 = "Thursday"
goto end_branch_7
} else {

}
}
{
var __t_tag_4 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_4) == 2946274527) {
__t7 = "Friday"
goto end_branch_7
} else {

}
}
{
var __t_tag_5 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_5) == 1070786179) {
__t7 = "Saturday"
goto end_branch_7
} else {

}
}
{
var __t_tag_6 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_6) == 1326716170) {
__t7 = "Sunday"
goto end_branch_7
} else {

}
}
{
__t7 = func() string { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Str(__t7)
})})))}
	})
	return cache_Data_Date_Component_showWeekday
}

var cache_Data_Date_Component_showMonth gopurs_runtime.Value
var once_Data_Date_Component_showMonth sync.Once
func Get_Data_Date_Component_showMonth() gopurs_runtime.Value {
	once_Data_Date_Component_showMonth.Do(func() {
		cache_Data_Date_Component_showMonth = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_1209612131_1386611502((&Constructor_Data_Show_Show[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 string
{
var __t_tag_0 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_0) == 1908470532) {
__t12 = "January"
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_1) == 2455627378) {
__t12 = "February"
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_2) == 4162469099) {
__t12 = "March"
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_3) == 1692989816) {
__t12 = "April"
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_4) == 330658827) {
__t12 = "May"
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_5) == 4067355978) {
__t12 = "June"
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_6) == 2276710548) {
__t12 = "July"
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_7) == 243771071) {
__t12 = "August"
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_8) == 215731793) {
__t12 = "September"
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_9) == 8639228) {
__t12 = "October"
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_10) == 49471444) {
__t12 = "November"
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_11) == 3889233761) {
__t12 = "December"
goto end_branch_12
} else {

}
}
{
__t12 = func() string { panic("Failed pattern match") }()
}
end_branch_12:
return gopurs_runtime.Str(__t12)
})})))}
	})
	return cache_Data_Date_Component_showMonth
}

var cache_Data_Date_Component_showDay gopurs_runtime.Value
var once_Data_Date_Component_showDay sync.Once
func Get_Data_Date_Component_showDay() gopurs_runtime.Value {
	once_Data_Date_Component_showDay.Do(func() {
		cache_Data_Date_Component_showDay = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_1636311157_1386611502((&Constructor_Data_Show_Show[int64]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Day ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(v_0.IntVal)).StrVal())) + (")"))
})})))}
	})
	return cache_Data_Date_Component_showDay
}

var cache_Data_Date_Component_ordYear gopurs_runtime.Value
var once_Data_Date_Component_ordYear sync.Once
func Get_Data_Date_Component_ordYear() gopurs_runtime.Value {
	once_Data_Date_Component_ordYear.Do(func() {
		cache_Data_Date_Component_ordYear = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3308271157_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[int64]](Get_Data_Ord_ordInt())))}
	})
	return cache_Data_Date_Component_ordYear
}

var cache_Data_Date_Component_ordDay gopurs_runtime.Value
var once_Data_Date_Component_ordDay sync.Once
func Get_Data_Date_Component_ordDay() gopurs_runtime.Value {
	once_Data_Date_Component_ordDay.Do(func() {
		cache_Data_Date_Component_ordDay = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3308271157_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[int64]](Get_Data_Ord_ordInt())))}
	})
	return cache_Data_Date_Component_ordDay
}

var cache_Data_Date_Component_eqYear gopurs_runtime.Value
var once_Data_Date_Component_eqYear sync.Once
func Get_Data_Date_Component_eqYear() gopurs_runtime.Value {
	once_Data_Date_Component_eqYear.Do(func() {
		cache_Data_Date_Component_eqYear = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}
	})
	return cache_Data_Date_Component_eqYear
}

var cache_Data_Date_Component_eqWeekday gopurs_runtime.Value
var once_Data_Date_Component_eqWeekday sync.Once
func Get_Data_Date_Component_eqWeekday() gopurs_runtime.Value {
	once_Data_Date_Component_eqWeekday.Do(func() {
		cache_Data_Date_Component_eqWeekday = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 bool
{
var __t_tag_3 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_3) == 2900196686) {
var __t_tag_4 uint32 = uint32(y_1.IntVal)
__t15 = (uint32(__t_tag_4) == 2900196686)
goto end_branch_15
} else {

}
}
{
var __t_tag_5 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_5) == 20457557) {
var __t_tag_6 uint32 = uint32(y_1.IntVal)
__t15 = (uint32(__t_tag_6) == 20457557)
goto end_branch_15
} else {

}
}
{
var __t_tag_7 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_7) == 4227105004) {
var __t_tag_8 uint32 = uint32(y_1.IntVal)
__t15 = (uint32(__t_tag_8) == 4227105004)
goto end_branch_15
} else {

}
}
{
var __t_tag_9 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_9) == 3818857258) {
var __t_tag_10 uint32 = uint32(y_1.IntVal)
__t15 = (uint32(__t_tag_10) == 3818857258)
goto end_branch_15
} else {

}
}
{
var __t_tag_11 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_11) == 2946274527) {
var __t_tag_12 uint32 = uint32(y_1.IntVal)
__t15 = (uint32(__t_tag_12) == 2946274527)
goto end_branch_15
} else {

}
}
{
var __t_tag_13 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_13) == 1070786179) {
var __t_tag_14 uint32 = uint32(y_1.IntVal)
__t15 = (uint32(__t_tag_14) == 1070786179)
goto end_branch_15
} else {

}
}
{
var __t_tag_0 uint32 = uint32(x_0.IntVal)
var __t_and_2 bool = false
if (uint32(__t_tag_0) == 1326716170) {

var __t_tag_1 uint32 = uint32(y_1.IntVal)
__t_and_2 = (uint32(__t_tag_1) == 1326716170)
}
__t15 = __t_and_2
}
end_branch_15:
return gopurs_runtime.Bool(__t15)
})})))}
	})
	return cache_Data_Date_Component_eqWeekday
}

var cache_Data_Date_Component_ordWeekday gopurs_runtime.Value
var once_Data_Date_Component_ordWeekday sync.Once
func Get_Data_Date_Component_ordWeekday() gopurs_runtime.Value {
	once_Data_Date_Component_ordWeekday.Do(func() {
		cache_Data_Date_Component_ordWeekday = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3730953251_4177771502((&Constructor_Data_Ord_Ord[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3768443459_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[uint32]](Get_Data_Date_Component_eqWeekday())))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 uint32
{
var __t_tag_0 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_0) == 2900196686) {
var __t2 uint32
{
var __t_tag_1 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_1) == 2900196686) {
__t2 = 902936544
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t27 = __t2
goto end_branch_27
} else {

}
}
{
var __t_tag_3 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_3) == 2900196686) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_4 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_4) == 20457557) {
var __t6 uint32
{
var __t_tag_5 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_5) == 20457557) {
__t6 = 902936544
goto end_branch_6
} else {

}
}
{
__t6 = 1527465420
}
end_branch_6:
__t27 = __t6
goto end_branch_27
} else {

}
}
{
var __t_tag_7 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_7) == 20457557) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_8 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_8) == 4227105004) {
var __t10 uint32
{
var __t_tag_9 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_9) == 4227105004) {
__t10 = 902936544
goto end_branch_10
} else {

}
}
{
__t10 = 1527465420
}
end_branch_10:
__t27 = __t10
goto end_branch_27
} else {

}
}
{
var __t_tag_11 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_11) == 4227105004) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_12 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_12) == 3818857258) {
var __t14 uint32
{
var __t_tag_13 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_13) == 3818857258) {
__t14 = 902936544
goto end_branch_14
} else {

}
}
{
__t14 = 1527465420
}
end_branch_14:
__t27 = __t14
goto end_branch_27
} else {

}
}
{
var __t_tag_15 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_15) == 3818857258) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_16 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_16) == 2946274527) {
var __t18 uint32
{
var __t_tag_17 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_17) == 2946274527) {
__t18 = 902936544
goto end_branch_18
} else {

}
}
{
__t18 = 1527465420
}
end_branch_18:
__t27 = __t18
goto end_branch_27
} else {

}
}
{
var __t_tag_19 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_19) == 2946274527) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_20 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_20) == 1070786179) {
var __t22 uint32
{
var __t_tag_21 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_21) == 1070786179) {
__t22 = 902936544
goto end_branch_22
} else {

}
}
{
__t22 = 1527465420
}
end_branch_22:
__t27 = __t22
goto end_branch_27
} else {

}
}
{
var __t_tag_23 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_23) == 1070786179) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_24 uint32 = uint32(x_0.IntVal)
var __t_and_26 bool = false
if (uint32(__t_tag_24) == 1326716170) {

var __t_tag_25 uint32 = uint32(y_1.IntVal)
__t_and_26 = (uint32(__t_tag_25) == 1326716170)
}
if __t_and_26 {
__t27 = 902936544
goto end_branch_27
} else {

}
}
{
__t27 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t27), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Date_Component_ordWeekday
}

var cache_Data_Date_Component_eqMonth gopurs_runtime.Value
var once_Data_Date_Component_eqMonth sync.Once
func Get_Data_Date_Component_eqMonth() gopurs_runtime.Value {
	once_Data_Date_Component_eqMonth.Do(func() {
		cache_Data_Date_Component_eqMonth = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t25 bool
{
var __t_tag_3 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_3) == 1908470532) {
var __t_tag_4 uint32 = uint32(y_1.IntVal)
__t25 = (uint32(__t_tag_4) == 1908470532)
goto end_branch_25
} else {

}
}
{
var __t_tag_5 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_5) == 2455627378) {
var __t_tag_6 uint32 = uint32(y_1.IntVal)
__t25 = (uint32(__t_tag_6) == 2455627378)
goto end_branch_25
} else {

}
}
{
var __t_tag_7 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_7) == 4162469099) {
var __t_tag_8 uint32 = uint32(y_1.IntVal)
__t25 = (uint32(__t_tag_8) == 4162469099)
goto end_branch_25
} else {

}
}
{
var __t_tag_9 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_9) == 1692989816) {
var __t_tag_10 uint32 = uint32(y_1.IntVal)
__t25 = (uint32(__t_tag_10) == 1692989816)
goto end_branch_25
} else {

}
}
{
var __t_tag_11 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_11) == 330658827) {
var __t_tag_12 uint32 = uint32(y_1.IntVal)
__t25 = (uint32(__t_tag_12) == 330658827)
goto end_branch_25
} else {

}
}
{
var __t_tag_13 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_13) == 4067355978) {
var __t_tag_14 uint32 = uint32(y_1.IntVal)
__t25 = (uint32(__t_tag_14) == 4067355978)
goto end_branch_25
} else {

}
}
{
var __t_tag_15 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_15) == 2276710548) {
var __t_tag_16 uint32 = uint32(y_1.IntVal)
__t25 = (uint32(__t_tag_16) == 2276710548)
goto end_branch_25
} else {

}
}
{
var __t_tag_17 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_17) == 243771071) {
var __t_tag_18 uint32 = uint32(y_1.IntVal)
__t25 = (uint32(__t_tag_18) == 243771071)
goto end_branch_25
} else {

}
}
{
var __t_tag_19 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_19) == 215731793) {
var __t_tag_20 uint32 = uint32(y_1.IntVal)
__t25 = (uint32(__t_tag_20) == 215731793)
goto end_branch_25
} else {

}
}
{
var __t_tag_21 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_21) == 8639228) {
var __t_tag_22 uint32 = uint32(y_1.IntVal)
__t25 = (uint32(__t_tag_22) == 8639228)
goto end_branch_25
} else {

}
}
{
var __t_tag_23 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_23) == 49471444) {
var __t_tag_24 uint32 = uint32(y_1.IntVal)
__t25 = (uint32(__t_tag_24) == 49471444)
goto end_branch_25
} else {

}
}
{
var __t_tag_0 uint32 = uint32(x_0.IntVal)
var __t_and_2 bool = false
if (uint32(__t_tag_0) == 3889233761) {

var __t_tag_1 uint32 = uint32(y_1.IntVal)
__t_and_2 = (uint32(__t_tag_1) == 3889233761)
}
__t25 = __t_and_2
}
end_branch_25:
return gopurs_runtime.Bool(__t25)
})})))}
	})
	return cache_Data_Date_Component_eqMonth
}

var cache_Data_Date_Component_ordMonth gopurs_runtime.Value
var once_Data_Date_Component_ordMonth sync.Once
func Get_Data_Date_Component_ordMonth() gopurs_runtime.Value {
	once_Data_Date_Component_ordMonth.Do(func() {
		cache_Data_Date_Component_ordMonth = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3730953251_4177771502((&Constructor_Data_Ord_Ord[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3768443459_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[uint32]](Get_Data_Date_Component_eqMonth())))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t47 uint32
{
var __t_tag_0 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_0) == 1908470532) {
var __t2 uint32
{
var __t_tag_1 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_1) == 1908470532) {
__t2 = 902936544
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t47 = __t2
goto end_branch_47
} else {

}
}
{
var __t_tag_3 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_3) == 1908470532) {
__t47 = 380165415
goto end_branch_47
} else {

}
}
{
var __t_tag_4 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_4) == 2455627378) {
var __t6 uint32
{
var __t_tag_5 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_5) == 2455627378) {
__t6 = 902936544
goto end_branch_6
} else {

}
}
{
__t6 = 1527465420
}
end_branch_6:
__t47 = __t6
goto end_branch_47
} else {

}
}
{
var __t_tag_7 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_7) == 2455627378) {
__t47 = 380165415
goto end_branch_47
} else {

}
}
{
var __t_tag_8 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_8) == 4162469099) {
var __t10 uint32
{
var __t_tag_9 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_9) == 4162469099) {
__t10 = 902936544
goto end_branch_10
} else {

}
}
{
__t10 = 1527465420
}
end_branch_10:
__t47 = __t10
goto end_branch_47
} else {

}
}
{
var __t_tag_11 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_11) == 4162469099) {
__t47 = 380165415
goto end_branch_47
} else {

}
}
{
var __t_tag_12 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_12) == 1692989816) {
var __t14 uint32
{
var __t_tag_13 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_13) == 1692989816) {
__t14 = 902936544
goto end_branch_14
} else {

}
}
{
__t14 = 1527465420
}
end_branch_14:
__t47 = __t14
goto end_branch_47
} else {

}
}
{
var __t_tag_15 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_15) == 1692989816) {
__t47 = 380165415
goto end_branch_47
} else {

}
}
{
var __t_tag_16 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_16) == 330658827) {
var __t18 uint32
{
var __t_tag_17 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_17) == 330658827) {
__t18 = 902936544
goto end_branch_18
} else {

}
}
{
__t18 = 1527465420
}
end_branch_18:
__t47 = __t18
goto end_branch_47
} else {

}
}
{
var __t_tag_19 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_19) == 330658827) {
__t47 = 380165415
goto end_branch_47
} else {

}
}
{
var __t_tag_20 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_20) == 4067355978) {
var __t22 uint32
{
var __t_tag_21 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_21) == 4067355978) {
__t22 = 902936544
goto end_branch_22
} else {

}
}
{
__t22 = 1527465420
}
end_branch_22:
__t47 = __t22
goto end_branch_47
} else {

}
}
{
var __t_tag_23 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_23) == 4067355978) {
__t47 = 380165415
goto end_branch_47
} else {

}
}
{
var __t_tag_24 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_24) == 2276710548) {
var __t26 uint32
{
var __t_tag_25 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_25) == 2276710548) {
__t26 = 902936544
goto end_branch_26
} else {

}
}
{
__t26 = 1527465420
}
end_branch_26:
__t47 = __t26
goto end_branch_47
} else {

}
}
{
var __t_tag_27 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_27) == 2276710548) {
__t47 = 380165415
goto end_branch_47
} else {

}
}
{
var __t_tag_28 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_28) == 243771071) {
var __t30 uint32
{
var __t_tag_29 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_29) == 243771071) {
__t30 = 902936544
goto end_branch_30
} else {

}
}
{
__t30 = 1527465420
}
end_branch_30:
__t47 = __t30
goto end_branch_47
} else {

}
}
{
var __t_tag_31 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_31) == 243771071) {
__t47 = 380165415
goto end_branch_47
} else {

}
}
{
var __t_tag_32 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_32) == 215731793) {
var __t34 uint32
{
var __t_tag_33 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_33) == 215731793) {
__t34 = 902936544
goto end_branch_34
} else {

}
}
{
__t34 = 1527465420
}
end_branch_34:
__t47 = __t34
goto end_branch_47
} else {

}
}
{
var __t_tag_35 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_35) == 215731793) {
__t47 = 380165415
goto end_branch_47
} else {

}
}
{
var __t_tag_36 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_36) == 8639228) {
var __t38 uint32
{
var __t_tag_37 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_37) == 8639228) {
__t38 = 902936544
goto end_branch_38
} else {

}
}
{
__t38 = 1527465420
}
end_branch_38:
__t47 = __t38
goto end_branch_47
} else {

}
}
{
var __t_tag_39 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_39) == 8639228) {
__t47 = 380165415
goto end_branch_47
} else {

}
}
{
var __t_tag_40 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_40) == 49471444) {
var __t42 uint32
{
var __t_tag_41 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_41) == 49471444) {
__t42 = 902936544
goto end_branch_42
} else {

}
}
{
__t42 = 1527465420
}
end_branch_42:
__t47 = __t42
goto end_branch_47
} else {

}
}
{
var __t_tag_43 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_43) == 49471444) {
__t47 = 380165415
goto end_branch_47
} else {

}
}
{
var __t_tag_44 uint32 = uint32(x_0.IntVal)
var __t_and_46 bool = false
if (uint32(__t_tag_44) == 3889233761) {

var __t_tag_45 uint32 = uint32(y_1.IntVal)
__t_and_46 = (uint32(__t_tag_45) == 3889233761)
}
if __t_and_46 {
__t47 = 902936544
goto end_branch_47
} else {

}
}
{
__t47 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_47:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t47), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Date_Component_ordMonth
}

var cache_Data_Date_Component_eqDay gopurs_runtime.Value
var once_Data_Date_Component_eqDay sync.Once
func Get_Data_Date_Component_eqDay() gopurs_runtime.Value {
	once_Data_Date_Component_eqDay.Do(func() {
		cache_Data_Date_Component_eqDay = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_1053099733_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[int64]](Get_Data_Eq_eqInt())))}
	})
	return cache_Data_Date_Component_eqDay
}

var cache_Data_Date_Component_boundedYear gopurs_runtime.Value
var once_Data_Date_Component_boundedYear sync.Once
func Get_Data_Date_Component_boundedYear() gopurs_runtime.Value {
	once_Data_Date_Component_boundedYear.Do(func() {
		cache_Data_Date_Component_boundedYear = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3764732725_2094947566((&Constructor_Data_Bounded_Bounded[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3308271157_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[int64]](Get_Data_Ord_ordInt())))}
}), int64(-271820), int64(275759)})))}
	})
	return cache_Data_Date_Component_boundedYear
}

var cache_Data_Date_Component_boundedWeekday gopurs_runtime.Value
var once_Data_Date_Component_boundedWeekday sync.Once
func Get_Data_Date_Component_boundedWeekday() gopurs_runtime.Value {
	once_Data_Date_Component_boundedWeekday.Do(func() {
		cache_Data_Date_Component_boundedWeekday = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_832288803_2094947566((&Constructor_Data_Bounded_Bounded[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3730953251_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](Get_Data_Date_Component_ordWeekday())))}
}), 2900196686, 1326716170})))}
	})
	return cache_Data_Date_Component_boundedWeekday
}

var cache_Data_Date_Component_boundedMonth gopurs_runtime.Value
var once_Data_Date_Component_boundedMonth sync.Once
func Get_Data_Date_Component_boundedMonth() gopurs_runtime.Value {
	once_Data_Date_Component_boundedMonth.Do(func() {
		cache_Data_Date_Component_boundedMonth = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_832288803_2094947566((&Constructor_Data_Bounded_Bounded[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3730953251_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](Get_Data_Date_Component_ordMonth())))}
}), 1908470532, 3889233761})))}
	})
	return cache_Data_Date_Component_boundedMonth
}

var cache_Data_Date_Component_boundedEnumYear gopurs_runtime.Value
var once_Data_Date_Component_boundedEnumYear sync.Once
func Get_Data_Date_Component_boundedEnumYear() gopurs_runtime.Value {
	once_Data_Date_Component_boundedEnumYear.Do(func() {
		cache_Data_Date_Component_boundedEnumYear = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_1306125126_123048125((&Constructor_Data_Enum_BoundedEnum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3764732725_2094947566(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[int64]](Get_Data_Date_Component_boundedYear())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_4060049525_556578094(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[int64]](Get_Data_Date_Component_enumYear())))}
}), gopurs_runtime.Int(int64(547580)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v_0.IntVal)
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((n_0.IntVal) >= (int64(-271820))) && ((n_0.IntVal) <= (int64(275759))) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(n_0.IntVal)}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_1170268447_3094389156(Rebox_Data_Date_Component_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
})})))}
	})
	return cache_Data_Date_Component_boundedEnumYear
}

var cache_Data_Date_Component_enumYear gopurs_runtime.Value
var once_Data_Date_Component_enumYear sync.Once
func Get_Data_Date_Component_enumYear() gopurs_runtime.Value {
	once_Data_Date_Component_enumYear.Do(func() {
		cache_Data_Date_Component_enumYear = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_4060049525_556578094((&Constructor_Data_Enum_Enum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3308271157_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[int64]](Get_Data_Ord_ordInt())))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[int64]](Get_Data_Date_Component_boundedEnumYear()).V4), gopurs_runtime.Int((gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[int64]](Get_Data_Date_Component_boundedEnumYear()).V3), x_0).IntVal).IntVal) - (int64(1))))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[int64]](Get_Data_Date_Component_boundedEnumYear()).V4), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[int64]](Get_Data_Date_Component_boundedEnumYear()).V3), x_0).IntVal) + (int64(1))))
})})))}
	})
	return cache_Data_Date_Component_enumYear
}

var cache_Data_Date_Component_boundedEnumWeekday gopurs_runtime.Value
var once_Data_Date_Component_boundedEnumWeekday sync.Once
func Get_Data_Date_Component_boundedEnumWeekday() gopurs_runtime.Value {
	once_Data_Date_Component_boundedEnumWeekday.Do(func() {
		cache_Data_Date_Component_boundedEnumWeekday = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_4021906832_123048125((&Constructor_Data_Enum_BoundedEnum[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_832288803_2094947566(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[uint32]](Get_Data_Date_Component_boundedWeekday())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_2359585123_556578094(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[uint32]](Get_Data_Date_Component_enumWeekday())))}
}), gopurs_runtime.Int(int64(7)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 int64
{
var __t_tag_0 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_0) == 2900196686) {
__t7 = int64(1)
goto end_branch_7
} else {

}
}
{
var __t_tag_1 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_1) == 20457557) {
__t7 = int64(2)
goto end_branch_7
} else {

}
}
{
var __t_tag_2 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_2) == 4227105004) {
__t7 = int64(3)
goto end_branch_7
} else {

}
}
{
var __t_tag_3 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_3) == 3818857258) {
__t7 = int64(4)
goto end_branch_7
} else {

}
}
{
var __t_tag_4 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_4) == 2946274527) {
__t7 = int64(5)
goto end_branch_7
} else {

}
}
{
var __t_tag_5 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_5) == 1070786179) {
__t7 = int64(6)
goto end_branch_7
} else {

}
}
{
var __t_tag_6 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_6) == 1326716170) {
__t7 = int64(7)
goto end_branch_7
} else {

}
}
{
__t7 = func() int64 { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Int(__t7)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v_0.IntVal) == (int64(1)) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2900196686), UnsafePtr: nil}}))}
goto end_branch_8
} else {

}
}
{
if (v_0.IntVal) == (int64(2)) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(20457557), UnsafePtr: nil}}))}
goto end_branch_8
} else {

}
}
{
if (v_0.IntVal) == (int64(3)) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4227105004), UnsafePtr: nil}}))}
goto end_branch_8
} else {

}
}
{
if (v_0.IntVal) == (int64(4)) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3818857258), UnsafePtr: nil}}))}
goto end_branch_8
} else {

}
}
{
if (v_0.IntVal) == (int64(5)) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2946274527), UnsafePtr: nil}}))}
goto end_branch_8
} else {

}
}
{
if (v_0.IntVal) == (int64(6)) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1070786179), UnsafePtr: nil}}))}
goto end_branch_8
} else {

}
}
{
if (v_0.IntVal) == (int64(7)) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1326716170), UnsafePtr: nil}}))}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_622082505_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_622082505_3094389156(Rebox_Data_Date_Component_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t8))))}
})})))}
	})
	return cache_Data_Date_Component_boundedEnumWeekday
}

var cache_Data_Date_Component_enumWeekday gopurs_runtime.Value
var once_Data_Date_Component_enumWeekday sync.Once
func Get_Data_Date_Component_enumWeekday() gopurs_runtime.Value {
	once_Data_Date_Component_enumWeekday.Do(func() {
		cache_Data_Date_Component_enumWeekday = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_2359585123_556578094((&Constructor_Data_Enum_Enum[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3730953251_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](Get_Data_Date_Component_ordWeekday())))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](Get_Data_Date_Component_boundedEnumWeekday()).V4), gopurs_runtime.Int((gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](Get_Data_Date_Component_boundedEnumWeekday()).V3), x_0).IntVal).IntVal) - (int64(1))))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](Get_Data_Date_Component_boundedEnumWeekday()).V4), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](Get_Data_Date_Component_boundedEnumWeekday()).V3), x_0).IntVal) + (int64(1))))
})})))}
	})
	return cache_Data_Date_Component_enumWeekday
}

var cache_Data_Date_Component_boundedEnumMonth gopurs_runtime.Value
var once_Data_Date_Component_boundedEnumMonth sync.Once
func Get_Data_Date_Component_boundedEnumMonth() gopurs_runtime.Value {
	once_Data_Date_Component_boundedEnumMonth.Do(func() {
		cache_Data_Date_Component_boundedEnumMonth = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_4021906832_123048125((&Constructor_Data_Enum_BoundedEnum[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_832288803_2094947566(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[uint32]](Get_Data_Date_Component_boundedMonth())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_2359585123_556578094(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[uint32]](Get_Data_Date_Component_enumMonth())))}
}), gopurs_runtime.Int(int64(12)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 int64
{
var __t_tag_0 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_0) == 1908470532) {
__t12 = int64(1)
goto end_branch_12
} else {

}
}
{
var __t_tag_1 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_1) == 2455627378) {
__t12 = int64(2)
goto end_branch_12
} else {

}
}
{
var __t_tag_2 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_2) == 4162469099) {
__t12 = int64(3)
goto end_branch_12
} else {

}
}
{
var __t_tag_3 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_3) == 1692989816) {
__t12 = int64(4)
goto end_branch_12
} else {

}
}
{
var __t_tag_4 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_4) == 330658827) {
__t12 = int64(5)
goto end_branch_12
} else {

}
}
{
var __t_tag_5 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_5) == 4067355978) {
__t12 = int64(6)
goto end_branch_12
} else {

}
}
{
var __t_tag_6 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_6) == 2276710548) {
__t12 = int64(7)
goto end_branch_12
} else {

}
}
{
var __t_tag_7 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_7) == 243771071) {
__t12 = int64(8)
goto end_branch_12
} else {

}
}
{
var __t_tag_8 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_8) == 215731793) {
__t12 = int64(9)
goto end_branch_12
} else {

}
}
{
var __t_tag_9 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_9) == 8639228) {
__t12 = int64(10)
goto end_branch_12
} else {

}
}
{
var __t_tag_10 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_10) == 49471444) {
__t12 = int64(11)
goto end_branch_12
} else {

}
}
{
var __t_tag_11 uint32 = uint32(v_0.IntVal)
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
return gopurs_runtime.Int(__t12)
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v_0.IntVal) == (int64(1)) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}}))}
goto end_branch_13
} else {

}
}
{
if (v_0.IntVal) == (int64(2)) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2455627378), UnsafePtr: nil}}))}
goto end_branch_13
} else {

}
}
{
if (v_0.IntVal) == (int64(3)) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4162469099), UnsafePtr: nil}}))}
goto end_branch_13
} else {

}
}
{
if (v_0.IntVal) == (int64(4)) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1692989816), UnsafePtr: nil}}))}
goto end_branch_13
} else {

}
}
{
if (v_0.IntVal) == (int64(5)) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(330658827), UnsafePtr: nil}}))}
goto end_branch_13
} else {

}
}
{
if (v_0.IntVal) == (int64(6)) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(4067355978), UnsafePtr: nil}}))}
goto end_branch_13
} else {

}
}
{
if (v_0.IntVal) == (int64(7)) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2276710548), UnsafePtr: nil}}))}
goto end_branch_13
} else {

}
}
{
if (v_0.IntVal) == (int64(8)) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(243771071), UnsafePtr: nil}}))}
goto end_branch_13
} else {

}
}
{
if (v_0.IntVal) == (int64(9)) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(215731793), UnsafePtr: nil}}))}
goto end_branch_13
} else {

}
}
{
if (v_0.IntVal) == (int64(10)) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(8639228), UnsafePtr: nil}}))}
goto end_branch_13
} else {

}
}
{
if (v_0.IntVal) == (int64(11)) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(49471444), UnsafePtr: nil}}))}
goto end_branch_13
} else {

}
}
{
if (v_0.IntVal) == (int64(12)) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3889233761), UnsafePtr: nil}}))}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_622082505_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[uint32]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_622082505_3094389156(Rebox_Data_Date_Component_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t13))))}
})})))}
	})
	return cache_Data_Date_Component_boundedEnumMonth
}

var cache_Data_Date_Component_enumMonth gopurs_runtime.Value
var once_Data_Date_Component_enumMonth sync.Once
func Get_Data_Date_Component_enumMonth() gopurs_runtime.Value {
	once_Data_Date_Component_enumMonth.Do(func() {
		cache_Data_Date_Component_enumMonth = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_2359585123_556578094((&Constructor_Data_Enum_Enum[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3730953251_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](Get_Data_Date_Component_ordMonth())))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](Get_Data_Date_Component_boundedEnumMonth()).V4), gopurs_runtime.Int((gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](Get_Data_Date_Component_boundedEnumMonth()).V3), x_0).IntVal).IntVal) - (int64(1))))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](Get_Data_Date_Component_boundedEnumMonth()).V4), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[uint32]](Get_Data_Date_Component_boundedEnumMonth()).V3), x_0).IntVal) + (int64(1))))
})})))}
	})
	return cache_Data_Date_Component_enumMonth
}

var cache_Data_Date_Component_boundedDay gopurs_runtime.Value
var once_Data_Date_Component_boundedDay sync.Once
func Get_Data_Date_Component_boundedDay() gopurs_runtime.Value {
	once_Data_Date_Component_boundedDay.Do(func() {
		cache_Data_Date_Component_boundedDay = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3764732725_2094947566((&Constructor_Data_Bounded_Bounded[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3308271157_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[int64]](Get_Data_Ord_ordInt())))}
}), int64(1), int64(31)})))}
	})
	return cache_Data_Date_Component_boundedDay
}

var cache_Data_Date_Component_boundedEnumDay gopurs_runtime.Value
var once_Data_Date_Component_boundedEnumDay sync.Once
func Get_Data_Date_Component_boundedEnumDay() gopurs_runtime.Value {
	once_Data_Date_Component_boundedEnumDay.Do(func() {
		cache_Data_Date_Component_boundedEnumDay = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_1306125126_123048125((&Constructor_Data_Enum_BoundedEnum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3764732725_2094947566(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[int64]](Get_Data_Date_Component_boundedDay())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_4060049525_556578094(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum[int64]](Get_Data_Date_Component_enumDay())))}
}), gopurs_runtime.Int(int64(31)).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v_0.IntVal)
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((n_0.IntVal) >= (int64(1))) && ((n_0.IntVal) <= (int64(31))) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(n_0.IntVal)}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_1170268447_3094389156(Rebox_Data_Date_Component_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
})})))}
	})
	return cache_Data_Date_Component_boundedEnumDay
}

var cache_Data_Date_Component_enumDay gopurs_runtime.Value
var once_Data_Date_Component_enumDay sync.Once
func Get_Data_Date_Component_enumDay() gopurs_runtime.Value {
	once_Data_Date_Component_enumDay.Do(func() {
		cache_Data_Date_Component_enumDay = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_4060049525_556578094((&Constructor_Data_Enum_Enum[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Date_Component_3308271157_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[int64]](Get_Data_Ord_ordInt())))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[int64]](Get_Data_Date_Component_boundedEnumDay()).V4), gopurs_runtime.Int((gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[int64]](Get_Data_Date_Component_boundedEnumDay()).V3), x_0).IntVal).IntVal) - (int64(1))))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[int64]](Get_Data_Date_Component_boundedEnumDay()).V4), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[int64]](Get_Data_Date_Component_boundedEnumDay()).V3), x_0).IntVal) + (int64(1))))
})})))}
	})
	return cache_Data_Date_Component_enumDay
}

type Constructor_Data_Date_Component_Monday struct {
	Rc uint32
}


type Constructor_Data_Date_Component_Tuesday struct {
	Rc uint32
}


type Constructor_Data_Date_Component_Wednesday struct {
	Rc uint32
}


type Constructor_Data_Date_Component_Thursday struct {
	Rc uint32
}


type Constructor_Data_Date_Component_Friday struct {
	Rc uint32
}


type Constructor_Data_Date_Component_Saturday struct {
	Rc uint32
}


type Constructor_Data_Date_Component_Sunday struct {
	Rc uint32
}


type Constructor_Data_Date_Component_January struct {
	Rc uint32
}


type Constructor_Data_Date_Component_February struct {
	Rc uint32
}


type Constructor_Data_Date_Component_March struct {
	Rc uint32
}


type Constructor_Data_Date_Component_April struct {
	Rc uint32
}


type Constructor_Data_Date_Component_May struct {
	Rc uint32
}


type Constructor_Data_Date_Component_June struct {
	Rc uint32
}


type Constructor_Data_Date_Component_July struct {
	Rc uint32
}


type Constructor_Data_Date_Component_August struct {
	Rc uint32
}


type Constructor_Data_Date_Component_September struct {
	Rc uint32
}


type Constructor_Data_Date_Component_October struct {
	Rc uint32
}


type Constructor_Data_Date_Component_November struct {
	Rc uint32
}


type Constructor_Data_Date_Component_December struct {
	Rc uint32
}


func Call_Data_Date_Component_Year(x_0_loop int64) int64 {
var x_0 int64 = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Date_Component_Day(x_0_loop int64) int64 {
var x_0 int64 = x_0_loop
_ = x_0
return x_0
}

func Rebox_Data_Date_Component_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Date_Component_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Date_Component_1209612131_1386611502(in *Constructor_Data_Show_Show[uint32]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Date_Component_1306125126_123048125(in *Constructor_Data_Enum_BoundedEnum[int64]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Date_Component_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Date_Component_2359585123_556578094(in *Constructor_Data_Enum_Enum[uint32]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Date_Component_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Date_Component_3094389156_622082505(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[uint32]{}
		out.V0 = uint32(in.V0.IntVal)
	return out
}

func Rebox_Data_Date_Component_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Date_Component_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Date_Component_3764732725_2094947566(in *Constructor_Data_Bounded_Bounded[int64]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Int(in.V1)
		out.V2 = gopurs_runtime.Int(in.V2)
	return out
}

func Rebox_Data_Date_Component_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Date_Component_4021906832_123048125(in *Constructor_Data_Enum_BoundedEnum[uint32]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Date_Component_4060049525_556578094(in *Constructor_Data_Enum_Enum[int64]) *Constructor_Data_Enum_Enum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_Enum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Date_Component_622082505_3094389156(in *Constructor_Data_Maybe_Just[uint32]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Data_Date_Component_832288803_2094947566(in *Constructor_Data_Bounded_Bounded[uint32]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
	return out
}



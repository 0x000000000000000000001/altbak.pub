package Data_Interval_Duration

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Second gopurs_runtime.Value
var once_Second sync.Once
func Get_Second() gopurs_runtime.Value {
	once_Second.Do(func() {
		cache_Second = gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil}
	})
	return cache_Second
}

var cache_Minute gopurs_runtime.Value
var once_Minute sync.Once
func Get_Minute() gopurs_runtime.Value {
	once_Minute.Do(func() {
		cache_Minute = gopurs_runtime.Value{Type: 9, IntVal: int64(217821258), UnsafePtr: nil}
	})
	return cache_Minute
}

var cache_Hour gopurs_runtime.Value
var once_Hour sync.Once
func Get_Hour() gopurs_runtime.Value {
	once_Hour.Do(func() {
		cache_Hour = gopurs_runtime.Value{Type: 9, IntVal: int64(1292308612), UnsafePtr: nil}
	})
	return cache_Hour
}

var cache_Day gopurs_runtime.Value
var once_Day sync.Once
func Get_Day() gopurs_runtime.Value {
	once_Day.Do(func() {
		cache_Day = gopurs_runtime.Value{Type: 9, IntVal: int64(2311060696), UnsafePtr: nil}
	})
	return cache_Day
}

var cache_Week gopurs_runtime.Value
var once_Week sync.Once
func Get_Week() gopurs_runtime.Value {
	once_Week.Do(func() {
		cache_Week = gopurs_runtime.Value{Type: 9, IntVal: int64(401302776), UnsafePtr: nil}
	})
	return cache_Week
}

var cache_Month gopurs_runtime.Value
var once_Month sync.Once
func Get_Month() gopurs_runtime.Value {
	once_Month.Do(func() {
		cache_Month = gopurs_runtime.Value{Type: 9, IntVal: int64(3327533908), UnsafePtr: nil}
	})
	return cache_Month
}

var cache_Year gopurs_runtime.Value
var once_Year sync.Once
func Get_Year() gopurs_runtime.Value {
	once_Year.Do(func() {
		cache_Year = gopurs_runtime.Value{Type: 9, IntVal: int64(3631736139), UnsafePtr: nil}
	})
	return cache_Year
}

var cache_Duration gopurs_runtime.Value
var once_Duration sync.Once
func Get_Duration() gopurs_runtime.Value {
	once_Duration.Do(func() {
		cache_Duration = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Duration(x_0_box)
})
	})
	return cache_Duration
}

var cache_showDurationComponent gopurs_runtime.Value
var once_showDurationComponent sync.Once
func Get_showDurationComponent() gopurs_runtime.Value {
	once_showDurationComponent.Do(func() {
		cache_showDurationComponent = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (uint32(v_0.IntVal) == 217821258) {
__t0 = gopurs_runtime.Str("Minute")
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 3908053364) {
__t0 = gopurs_runtime.Str("Second")
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 1292308612) {
__t0 = gopurs_runtime.Str("Hour")
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 2311060696) {
__t0 = gopurs_runtime.Str("Day")
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 401302776) {
__t0 = gopurs_runtime.Str("Week")
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 3327533908) {
__t0 = gopurs_runtime.Str("Month")
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 3631736139) {
__t0 = gopurs_runtime.Str("Year")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str(__t0.StrVal())
}))
	})
	return cache_showDurationComponent
}

var cache_showMap gopurs_runtime.Value
var once_showMap sync.Once
func Get_showMap() gopurs_runtime.Value {
	once_showMap.Do(func() {
		cache_showMap = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[*pkg_Data_Map_Internal.Constructor_Node[uint32, float64]]](gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_showMap(), Get_showDurationComponent(), pkg_Data_Show.Get_showNumber())))}
	})
	return cache_showMap
}

var cache_showDuration gopurs_runtime.Value
var once_showDuration sync.Once
func Get_showDuration() gopurs_runtime.Value {
	once_showDuration.Do(func() {
		cache_showDuration = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Duration "), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_show__2896747026(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](v_0))}).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
	})
	return cache_showDuration
}

var cache_newtypeDuration gopurs_runtime.Value
var once_newtypeDuration sync.Once
func Get_newtypeDuration() gopurs_runtime.Value {
	once_newtypeDuration.Do(func() {
		cache_newtypeDuration = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeDuration
}

var cache_eqDurationComponent gopurs_runtime.Value
var once_eqDurationComponent sync.Once
func Get_eqDurationComponent() gopurs_runtime.Value {
	once_eqDurationComponent.Do(func() {
		cache_eqDurationComponent = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
if (uint32(x_0.IntVal) == 3908053364) {
var __t0 bool
{
if (uint32(y_1.IntVal) == 3908053364) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t6 = __t0
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 217821258) {
var __t1 bool
{
if (uint32(y_1.IntVal) == 217821258) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t6 = __t1
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 1292308612) {
var __t2 bool
{
if (uint32(y_1.IntVal) == 1292308612) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t6 = __t2
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 2311060696) {
var __t3 bool
{
if (uint32(y_1.IntVal) == 2311060696) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t6 = __t3
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 401302776) {
var __t4 bool
{
if (uint32(y_1.IntVal) == 401302776) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t6 = __t4
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 3327533908) {
var __t5 bool
{
if (uint32(y_1.IntVal) == 3327533908) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if ((uint32(x_0.IntVal) == 3631736139)) && ((uint32(y_1.IntVal) == 3631736139)) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
}))
	})
	return cache_eqDurationComponent
}

var cache_eqMap gopurs_runtime.Value
var once_eqMap sync.Once
func Get_eqMap() gopurs_runtime.Value {
	once_eqMap.Do(func() {
		cache_eqMap = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[*pkg_Data_Map_Internal.Constructor_Node[uint32, float64]]](gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_eqMap(), Get_eqDurationComponent(), pkg_Data_Eq.Get_eqNumber())))}
	})
	return cache_eqMap
}

var cache_ordDurationComponent gopurs_runtime.Value
var once_ordDurationComponent sync.Once
func Get_ordDurationComponent() gopurs_runtime.Value {
	once_ordDurationComponent.Do(func() {
		cache_ordDurationComponent = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqDurationComponent()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (uint32(x_0.IntVal) == 3908053364) {
var __t0 uint32
{
if (uint32(y_1.IntVal) == 3908053364) {
__t0 = 902936544
goto end_branch_0
} else {

}
}
{
__t0 = 1527465420
}
end_branch_0:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t0), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (uint32(y_1.IntVal) == 3908053364) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 217821258) {
var __t1 uint32
{
if (uint32(y_1.IntVal) == 217821258) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
__t1 = 1527465420
}
end_branch_1:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (uint32(y_1.IntVal) == 217821258) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 1292308612) {
var __t2 uint32
{
if (uint32(y_1.IntVal) == 1292308612) {
__t2 = 902936544
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t2), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (uint32(y_1.IntVal) == 1292308612) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 2311060696) {
var __t3 uint32
{
if (uint32(y_1.IntVal) == 2311060696) {
__t3 = 902936544
goto end_branch_3
} else {

}
}
{
__t3 = 1527465420
}
end_branch_3:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (uint32(y_1.IntVal) == 2311060696) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 401302776) {
var __t4 uint32
{
if (uint32(y_1.IntVal) == 401302776) {
__t4 = 902936544
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (uint32(y_1.IntVal) == 401302776) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (uint32(x_0.IntVal) == 3327533908) {
var __t5 uint32
{
if (uint32(y_1.IntVal) == 3327533908) {
__t5 = 902936544
goto end_branch_5
} else {

}
}
{
__t5 = 1527465420
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (uint32(y_1.IntVal) == 3327533908) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if ((uint32(x_0.IntVal) == 3631736139)) && ((uint32(y_1.IntVal) == 3631736139)) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t6.IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_ordDurationComponent
}

var cache_ordMap gopurs_runtime.Value
var once_ordMap sync.Once
func Get_ordMap() gopurs_runtime.Value {
	once_ordMap.Do(func() {
		cache_ordMap = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_Map_Internal.Constructor_Node[uint32, float64]]](gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_ordMap(), Get_ordDurationComponent(), pkg_Data_Ord.Get_ordNumber())))}
	})
	return cache_ordMap
}

var cache_semigroupDuration gopurs_runtime.Value
var once_semigroupDuration sync.Once
func Get_semigroupDuration() gopurs_runtime.Value {
	once_semigroupDuration.Do(func() {
		cache_semigroupDuration = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_unionWith__952079555(pkg_Data_Semiring.Get_numAdd(), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](v_0))}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](v1_1))})))}))}
})
}))
	})
	return cache_semigroupDuration
}

var cache_monoidDuration gopurs_runtime.Value
var once_monoidDuration sync.Once
func Get_monoidDuration() gopurs_runtime.Value {
	once_monoidDuration.Do(func() {
		cache_monoidDuration = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupDuration()
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
	})
	return cache_monoidDuration
}

var cache_eqDuration gopurs_runtime.Value
var once_eqDuration sync.Once
func Get_eqDuration() gopurs_runtime.Value {
	once_eqDuration.Do(func() {
		cache_eqDuration = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_eq__2224314568(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](x_0))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](y_1))}).IntVal) != (0))
})
}))
	})
	return cache_eqDuration
}

var cache_ordDuration gopurs_runtime.Value
var once_ordDuration sync.Once
func Get_ordDuration() gopurs_runtime.Value {
	once_ordDuration.Do(func() {
		cache_ordDuration = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqDuration()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(Get_compare__231252914(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](x_0))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](y_1))}).IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_ordDuration
}

var cache_durationFromComponent gopurs_runtime.Value
var once_durationFromComponent sync.Once
func Get_durationFromComponent() gopurs_runtime.Value {
	once_durationFromComponent.Do(func() {
		cache_durationFromComponent = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_durationFromComponent(uint32(k_0_box.IntVal), v_1_box.FloatVal()))}
})
	})
	return cache_durationFromComponent
}

var cache_hour gopurs_runtime.Value
var once_hour sync.Once
func Get_hour() gopurs_runtime.Value {
	once_hour.Do(func() {
		cache_hour = gopurs_runtime.Apply(Get_durationFromComponent(), gopurs_runtime.Value{Type: 9, IntVal: int64(1292308612), UnsafePtr: nil})
	})
	return cache_hour
}

var cache_millisecond gopurs_runtime.Value
var once_millisecond sync.Once
func Get_millisecond() gopurs_runtime.Value {
	once_millisecond.Do(func() {
		cache_millisecond = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_millisecond(x_0_box.FloatVal()))}
})
	})
	return cache_millisecond
}

var cache_minute gopurs_runtime.Value
var once_minute sync.Once
func Get_minute() gopurs_runtime.Value {
	once_minute.Do(func() {
		cache_minute = gopurs_runtime.Apply(Get_durationFromComponent(), gopurs_runtime.Value{Type: 9, IntVal: int64(217821258), UnsafePtr: nil})
	})
	return cache_minute
}

var cache_month gopurs_runtime.Value
var once_month sync.Once
func Get_month() gopurs_runtime.Value {
	once_month.Do(func() {
		cache_month = gopurs_runtime.Apply(Get_durationFromComponent(), gopurs_runtime.Value{Type: 9, IntVal: int64(3327533908), UnsafePtr: nil})
	})
	return cache_month
}

var cache_second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		cache_second = gopurs_runtime.Apply(Get_durationFromComponent(), gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil})
	})
	return cache_second
}

var cache_week gopurs_runtime.Value
var once_week sync.Once
func Get_week() gopurs_runtime.Value {
	once_week.Do(func() {
		cache_week = gopurs_runtime.Apply(Get_durationFromComponent(), gopurs_runtime.Value{Type: 9, IntVal: int64(401302776), UnsafePtr: nil})
	})
	return cache_week
}

var cache_year gopurs_runtime.Value
var once_year sync.Once
func Get_year() gopurs_runtime.Value {
	once_year.Do(func() {
		cache_year = gopurs_runtime.Apply(Get_durationFromComponent(), gopurs_runtime.Value{Type: 9, IntVal: int64(3631736139), UnsafePtr: nil})
	})
	return cache_year
}

var cache_day gopurs_runtime.Value
var once_day sync.Once
func Get_day() gopurs_runtime.Value {
	once_day.Do(func() {
		cache_day = gopurs_runtime.Apply(Get_durationFromComponent(), gopurs_runtime.Value{Type: 9, IntVal: int64(2311060696), UnsafePtr: nil})
	})
	return cache_day
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_eq__2224314568 gopurs_runtime.Value
var once_eq__2224314568 sync.Once
func Get_eq__2224314568() gopurs_runtime.Value {
	once_eq__2224314568.Do(func() {
		cache_eq__2224314568 = gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[*pkg_Data_Map_Internal.Constructor_Node[uint32, float64]]](Get_eqMap()).V0
	})
	return cache_eq__2224314568
}

var cache_div__1002719800 gopurs_runtime.Value
var once_div__1002719800 sync.Once
func Get_div__1002719800() gopurs_runtime.Value {
	once_div__1002719800.Do(func() {
		cache_div__1002719800 = gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div")
	})
	return cache_div__1002719800
}

var cache_div__2579358968 gopurs_runtime.Value
var once_div__2579358968 sync.Once
func Get_div__2579358968() gopurs_runtime.Value {
	once_div__2579358968.Do(func() {
		cache_div__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_div__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_div__2579358968
}

var cache_empty__1299254065 gopurs_runtime.Value
var once_empty__1299254065 sync.Once
func Get_empty__1299254065() gopurs_runtime.Value {
	once_empty__1299254065.Do(func() {
		cache_empty__1299254065 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
	})
	return cache_empty__1299254065
}

var cache_singleton__2450056090 gopurs_runtime.Value
var once_singleton__2450056090 sync.Once
func Get_singleton__2450056090() gopurs_runtime.Value {
	once_singleton__2450056090.Do(func() {
		cache_singleton__2450056090 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_singleton__2450056090(k_0_box, v_1_box))}
})
	})
	return cache_singleton__2450056090
}

var cache_singleton__1518627866 gopurs_runtime.Value
var once_singleton__1518627866 sync.Once
func Get_singleton__1518627866() gopurs_runtime.Value {
	once_singleton__1518627866.Do(func() {
		cache_singleton__1518627866 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_singleton__1518627866(k_0_box, v_1_box))}
})
	})
	return cache_singleton__1518627866
}

var cache_unionWith__2507192643 gopurs_runtime.Value
var once_unionWith__2507192643 sync.Once
func Get_unionWith__2507192643() gopurs_runtime.Value {
	once_unionWith__2507192643.Do(func() {
		cache_unionWith__2507192643 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unionWith__2507192643(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_unionWith__2507192643
}

var cache_unionWith__952079555 gopurs_runtime.Value
var once_unionWith__952079555 sync.Once
func Get_unionWith__952079555() gopurs_runtime.Value {
	once_unionWith__952079555.Do(func() {
		cache_unionWith__952079555 = gopurs_runtime.Func3(func(app_0_box gopurs_runtime.Value, m1_1_box gopurs_runtime.Value, m2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_unionWith__952079555(app_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_2_box)))}
})
	})
	return cache_unionWith__952079555
}

var cache_unsafeBalancedNode__1305301638 gopurs_runtime.Value
var once_unsafeBalancedNode__1305301638 sync.Once
func Get_unsafeBalancedNode__1305301638() gopurs_runtime.Value {
	once_unsafeBalancedNode__1305301638.Do(func() {
		cache_unsafeBalancedNode__1305301638 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeBalancedNode__1305301638(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_unsafeBalancedNode__1305301638
}

var cache_unsafeNode__1305301638 gopurs_runtime.Value
var once_unsafeNode__1305301638 sync.Once
func Get_unsafeNode__1305301638() gopurs_runtime.Value {
	once_unsafeNode__1305301638.Do(func() {
		cache_unsafeNode__1305301638 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeNode__1305301638(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_unsafeNode__1305301638
}

var cache_unsafeSplit__4154869695 gopurs_runtime.Value
var once_unsafeSplit__4154869695 sync.Once
func Get_unsafeSplit__4154869695() gopurs_runtime.Value {
	once_unsafeSplit__4154869695.Do(func() {
		cache_unsafeSplit__4154869695 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeSplit__4154869695(__local_var_0_box, __local_var_1_box, __local_var_2_box)
})
	})
	return cache_unsafeSplit__4154869695
}

var cache_unsafeUnionWith__4109280494 gopurs_runtime.Value
var once_unsafeUnionWith__4109280494 sync.Once
func Get_unsafeUnionWith__4109280494() gopurs_runtime.Value {
	once_unsafeUnionWith__4109280494.Do(func() {
		cache_unsafeUnionWith__4109280494 = gopurs_runtime.Func4(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeUnionWith__4109280494(__local_var_0_box, __local_var_1_box, __local_var_2_box, __local_var_3_box)
})
	})
	return cache_unsafeUnionWith__4109280494
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_compare__231252914 gopurs_runtime.Value
var once_compare__231252914 sync.Once
func Get_compare__231252914() gopurs_runtime.Value {
	once_compare__231252914.Do(func() {
		cache_compare__231252914 = gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_Map_Internal.Constructor_Node[uint32, float64]]](Get_ordMap()).V1
	})
	return cache_compare__231252914
}

var cache_greaterThan__4087042607 gopurs_runtime.Value
var once_greaterThan__4087042607 sync.Once
func Get_greaterThan__4087042607() gopurs_runtime.Value {
	once_greaterThan__4087042607.Do(func() {
		cache_greaterThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThan__4087042607
}

var cache_greaterThan__1409282474 gopurs_runtime.Value
var once_greaterThan__1409282474 sync.Once
func Get_greaterThan__1409282474() gopurs_runtime.Value {
	once_greaterThan__1409282474.Do(func() {
		cache_greaterThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan__1409282474
}

var cache_lessThanOrEq__4087042607 gopurs_runtime.Value
var once_lessThanOrEq__4087042607 sync.Once
func Get_lessThanOrEq__4087042607() gopurs_runtime.Value {
	once_lessThanOrEq__4087042607.Do(func() {
		cache_lessThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__4087042607
}

var cache_lessThanOrEq__1409282474 gopurs_runtime.Value
var once_lessThanOrEq__1409282474 sync.Once
func Get_lessThanOrEq__1409282474() gopurs_runtime.Value {
	once_lessThanOrEq__1409282474.Do(func() {
		cache_lessThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThanOrEq__1409282474
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = pkg_Data_Semiring.Get_intAdd()
	})
	return cache_add__560788792
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_show__2896747026 gopurs_runtime.Value
var once_show__2896747026 sync.Once
func Get_show__2896747026() gopurs_runtime.Value {
	once_show__2896747026.Do(func() {
		cache_show__2896747026 = gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[*pkg_Data_Map_Internal.Constructor_Node[uint32, float64]]](Get_showMap()).V0
	})
	return cache_show__2896747026
}

type Constructor_Second struct {
	Rc uint32
}


type Constructor_Minute struct {
	Rc uint32
}


type Constructor_Hour struct {
	Rc uint32
}


type Constructor_Day struct {
	Rc uint32
}


type Constructor_Week struct {
	Rc uint32
}


type Constructor_Month struct {
	Rc uint32
}


type Constructor_Year struct {
	Rc uint32
}


func Call_Duration(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_durationFromComponent(k_0_loop uint32, v_1_loop float64) *pkg_Data_Map_Internal.Constructor_Node[uint32, float64] {
var k_0 uint32 = k_0_loop
_ = k_0
var v_1 float64 = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_singleton__1518627866(gopurs_runtime.Value{Type: 9, IntVal: int64(k_0), UnsafePtr: nil}, gopurs_runtime.Float(v_1)))})
}

func Call_millisecond(x_0_loop float64) *pkg_Data_Map_Internal.Constructor_Node[uint32, float64] {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_singleton__1518627866(gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil}, gopurs_runtime.Float(gopurs_runtime.Float(gopurs_runtime.Apply2(Get_div__1002719800(), gopurs_runtime.Float(gopurs_runtime.Float(x_0).FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal()).FloatVal())))})
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_div__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_singleton__2450056090(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_0, v_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_singleton__1518627866(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_0, v_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_unionWith__2507192643(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(app_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, app_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
})
})
})
}

func Call_unionWith__952079555(app_0_loop gopurs_runtime.Value, m1_1_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value], m2_2_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var app_0 gopurs_runtime.Value = app_0_loop
_ = app_0
var m1_1 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = m1_1_loop
_ = m1_1
var m2_2 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = m2_2_loop
_ = m2_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), gopurs_runtime.RecordGet(Get_ordDurationComponent(), "compare"), app_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(m1_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(m2_2)}))
}

func Call_unsafeBalancedNode__1305301638(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t27 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t6 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_6
} else {

}
}
{
if ((__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil)) && ((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), gopurs_runtime.Int(1))).IntVal) != (0)) {
var __t5 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_4 bool = false
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr != nil) {

var __t3 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t_and_4 = (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0), gopurs_runtime.Int(__t3.IntVal))).IntVal) != (0)
}
if __t_and_4 {
__t5 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t5)}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}
}
end_branch_6:
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
goto end_branch_27
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t26 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t19 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int(1)).IntVal))).IntVal) != (0) {
var __t12 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}
var __t_and_11 bool = false
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 324739070 && __t_tag_7.UnsafePtr != nil) {

var __t10 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 324739070 && __t_tag_8.UnsafePtr == nil) {
__t10 = gopurs_runtime.Int(0)
goto end_branch_10
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 324739070 && __t_tag_9.UnsafePtr != nil) {
__t10 = gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t_and_11 = (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V0), gopurs_runtime.Int(__t10.IntVal))).IntVal) != (0)
}
if __t_and_11 {
__t12 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})))}))
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}))
}
end_branch_12:
__t19 = __t12
goto end_branch_19
} else {

}
}
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0), gopurs_runtime.Int(1)).IntVal))).IntVal) != (0) {
var __t18 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_17 bool = false
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 324739070 && __t_tag_13.UnsafePtr != nil) {

var __t16 gopurs_runtime.Value
{
var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 324739070 && __t_tag_14.UnsafePtr == nil) {
__t16 = gopurs_runtime.Int(0)
goto end_branch_16
} else {

}
}
{
var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 324739070 && __t_tag_15.UnsafePtr != nil) {
__t16 = gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
__t_and_17 = (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(__t16.IntVal), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0))).IntVal) != (0)
}
if __t_and_17 {
__t18 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_18:
__t19 = __t18
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_19:
__t26 = __t19
goto end_branch_26
} else {

}
}
{
if ((__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil)) && ((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int(1))).IntVal) != (0)) {
var __t25 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t_tag_20 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}
var __t_and_24 bool = false
if (__t_tag_20.Type == 9 && __t_tag_20.IntVal == 324739070 && __t_tag_20.UnsafePtr != nil) {

var __t23 gopurs_runtime.Value
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_21.Type == 9 && __t_tag_21.IntVal == 324739070 && __t_tag_21.UnsafePtr == nil) {
__t23 = gopurs_runtime.Int(0)
goto end_branch_23
} else {

}
}
{
var __t_tag_22 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}
if (__t_tag_22.Type == 9 && __t_tag_22.IntVal == 324739070 && __t_tag_22.UnsafePtr != nil) {
__t23 = gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}.UnsafePtr).V0)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
__t_and_24 = (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(__t23.IntVal), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V0))).IntVal) != (0)
}
if __t_and_24 {
__t25 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))})))}))
}
end_branch_25:
__t26 = __t25
goto end_branch_26
} else {

}
}
{
__t26 = gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeNode(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}))
}
end_branch_26:
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t26)}
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t27))}
}

func Call_unsafeNode__1305301638(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t3 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
var __t0 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(1), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0)).IntVal, gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(1), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1)).IntVal, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t0))}
goto end_branch_3
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(1), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0)).IntVal, gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(1), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)).IntVal, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_2
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
var __t1 int64
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0))).IntVal) != (0) {
__t1 = (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V0
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(1), gopurs_runtime.Int(__t1)).IntVal, gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(1), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V1)).IntVal), gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V1)).IntVal, __local_var_0, __local_var_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3)})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
}

func Call_unsafeSplit__4154869695(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __t4 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})}
goto end_branch_4
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
v_3_0 := gopurs_runtime.Apply2(__local_var_0, __local_var_1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2)
_ = v_3_0
var __t3 gopurs_runtime.Value
{
if (uint32(v_3_0.IntVal) == 1527465420) {
v1_4_1 := gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)})
_ = v1_4_1
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V0, (*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeBalancedNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_1.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)}))})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 380165415) {
v1_4_2 := gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeSplit(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5)})
_ = v1_4_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeBalancedNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V1)})), (*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4_2.UnsafePtr).V2})}
goto end_branch_3
} else {

}
}
{
if (uint32(v_3_0.IntVal) == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V3})}), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V4, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2.UnsafePtr).V5})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t3))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 3373277644, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value]](__t4))}
}

func Call_unsafeUnionWith__4109280494(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop gopurs_runtime.Value, __local_var_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
var __local_var_3 gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
var __t6 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))}
goto end_branch_6
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
v_4_0 := gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeSplit(), __local_var_0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_2))})
_ = v_4_0
l_prime_5_1 := gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)})
_ = l_prime_5_1
r_prime_6_2 := gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), __local_var_0, __local_var_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V2)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)})
_ = r_prime_6_2
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeBalancedNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Apply2(__local_var_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}.UnsafePtr).V0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Split[gopurs_runtime.Value, gopurs_runtime.Value])(v_4_0.UnsafePtr).V0)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeBalancedNode(), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](l_prime_5_1))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](r_prime_6_2))})))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t5))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t6))}
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_greaterThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_greaterThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_lessThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_lessThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}



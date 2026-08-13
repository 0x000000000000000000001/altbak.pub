package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Interval_Duration_add gopurs_runtime.Value
var once_Data_Interval_Duration_add sync.Once
func Get_Data_Interval_Duration_add() gopurs_runtime.Value {
	once_Data_Interval_Duration_add.Do(func() {
		cache_Data_Interval_Duration_add = Get_Data_Semiring_numAdd()
	})
	return cache_Data_Interval_Duration_add
}

var cache_Data_Interval_Duration_Second gopurs_runtime.Value
var once_Data_Interval_Duration_Second sync.Once
func Get_Data_Interval_Duration_Second() gopurs_runtime.Value {
	once_Data_Interval_Duration_Second.Do(func() {
		cache_Data_Interval_Duration_Second = gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Second
}

var cache_Data_Interval_Duration_Minute gopurs_runtime.Value
var once_Data_Interval_Duration_Minute sync.Once
func Get_Data_Interval_Duration_Minute() gopurs_runtime.Value {
	once_Data_Interval_Duration_Minute.Do(func() {
		cache_Data_Interval_Duration_Minute = gopurs_runtime.Value{Type: 9, IntVal: int64(217821258), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Minute
}

var cache_Data_Interval_Duration_Hour gopurs_runtime.Value
var once_Data_Interval_Duration_Hour sync.Once
func Get_Data_Interval_Duration_Hour() gopurs_runtime.Value {
	once_Data_Interval_Duration_Hour.Do(func() {
		cache_Data_Interval_Duration_Hour = gopurs_runtime.Value{Type: 9, IntVal: int64(1292308612), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Hour
}

var cache_Data_Interval_Duration_Day gopurs_runtime.Value
var once_Data_Interval_Duration_Day sync.Once
func Get_Data_Interval_Duration_Day() gopurs_runtime.Value {
	once_Data_Interval_Duration_Day.Do(func() {
		cache_Data_Interval_Duration_Day = gopurs_runtime.Value{Type: 9, IntVal: int64(2311060696), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Day
}

var cache_Data_Interval_Duration_Week gopurs_runtime.Value
var once_Data_Interval_Duration_Week sync.Once
func Get_Data_Interval_Duration_Week() gopurs_runtime.Value {
	once_Data_Interval_Duration_Week.Do(func() {
		cache_Data_Interval_Duration_Week = gopurs_runtime.Value{Type: 9, IntVal: int64(401302776), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Week
}

var cache_Data_Interval_Duration_Month gopurs_runtime.Value
var once_Data_Interval_Duration_Month sync.Once
func Get_Data_Interval_Duration_Month() gopurs_runtime.Value {
	once_Data_Interval_Duration_Month.Do(func() {
		cache_Data_Interval_Duration_Month = gopurs_runtime.Value{Type: 9, IntVal: int64(3327533908), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Month
}

var cache_Data_Interval_Duration_Year gopurs_runtime.Value
var once_Data_Interval_Duration_Year sync.Once
func Get_Data_Interval_Duration_Year() gopurs_runtime.Value {
	once_Data_Interval_Duration_Year.Do(func() {
		cache_Data_Interval_Duration_Year = gopurs_runtime.Value{Type: 9, IntVal: int64(3631736139), UnsafePtr: nil}
	})
	return cache_Data_Interval_Duration_Year
}

var cache_Data_Interval_Duration_Duration gopurs_runtime.Value
var once_Data_Interval_Duration_Duration sync.Once
func Get_Data_Interval_Duration_Duration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Duration.Do(func() {
		cache_Data_Interval_Duration_Duration = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_Duration_Duration(x_0_box)
})
	})
	return cache_Data_Interval_Duration_Duration
}

var cache_Data_Interval_Duration_showDurationComponent gopurs_runtime.Value
var once_Data_Interval_Duration_showDurationComponent sync.Once
func Get_Data_Interval_Duration_showDurationComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_showDurationComponent.Do(func() {
		cache_Data_Interval_Duration_showDurationComponent = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_Interval_Duration_showDurationComponent
}

var cache_Data_Interval_Duration_showMap gopurs_runtime.Value
var once_Data_Interval_Duration_showMap sync.Once
func Get_Data_Interval_Duration_showMap() gopurs_runtime.Value {
	once_Data_Interval_Duration_showMap.Do(func() {
		cache_Data_Interval_Duration_showMap = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](gopurs_runtime.Apply2(Get_Data_Map_Internal_showMap(), Get_Data_Interval_Duration_showDurationComponent(), Get_Data_Show_showNumber())))}
	})
	return cache_Data_Interval_Duration_showMap
}

var cache_Data_Interval_Duration_showDuration gopurs_runtime.Value
var once_Data_Interval_Duration_showDuration sync.Once
func Get_Data_Interval_Duration_showDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_showDuration.Do(func() {
		cache_Data_Interval_Duration_showDuration = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Duration ") + (gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Data_Interval_Duration_showMap()).V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0))}).StrVal())) + (")"))
}))
	})
	return cache_Data_Interval_Duration_showDuration
}

var cache_Data_Interval_Duration_newtypeDuration gopurs_runtime.Value
var once_Data_Interval_Duration_newtypeDuration sync.Once
func Get_Data_Interval_Duration_newtypeDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_newtypeDuration.Do(func() {
		cache_Data_Interval_Duration_newtypeDuration = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Interval_Duration_newtypeDuration
}

var cache_Data_Interval_Duration_eqDurationComponent gopurs_runtime.Value
var once_Data_Interval_Duration_eqDurationComponent sync.Once
func Get_Data_Interval_Duration_eqDurationComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_eqDurationComponent.Do(func() {
		cache_Data_Interval_Duration_eqDurationComponent = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_Interval_Duration_eqDurationComponent
}

var cache_Data_Interval_Duration_eqMap gopurs_runtime.Value
var once_Data_Interval_Duration_eqMap sync.Once
func Get_Data_Interval_Duration_eqMap() gopurs_runtime.Value {
	once_Data_Interval_Duration_eqMap.Do(func() {
		cache_Data_Interval_Duration_eqMap = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.Apply2(Get_Data_Map_Internal_eqMap(), Get_Data_Interval_Duration_eqDurationComponent(), Get_Data_Eq_eqNumber())))}
	})
	return cache_Data_Interval_Duration_eqMap
}

var cache_Data_Interval_Duration_ordDurationComponent gopurs_runtime.Value
var once_Data_Interval_Duration_ordDurationComponent sync.Once
func Get_Data_Interval_Duration_ordDurationComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_ordDurationComponent.Do(func() {
		cache_Data_Interval_Duration_ordDurationComponent = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_Duration_eqDurationComponent()
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
	return cache_Data_Interval_Duration_ordDurationComponent
}

var cache_Data_Interval_Duration_ordMap gopurs_runtime.Value
var once_Data_Interval_Duration_ordMap sync.Once
func Get_Data_Interval_Duration_ordMap() gopurs_runtime.Value {
	once_Data_Interval_Duration_ordMap.Do(func() {
		cache_Data_Interval_Duration_ordMap = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.Apply2(Get_Data_Map_Internal_ordMap(), Get_Data_Interval_Duration_ordDurationComponent(), Get_Data_Ord_ordNumber())))}
	})
	return cache_Data_Interval_Duration_ordMap
}

var cache_Data_Interval_Duration_semigroupDuration gopurs_runtime.Value
var once_Data_Interval_Duration_semigroupDuration sync.Once
func Get_Data_Interval_Duration_semigroupDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_semigroupDuration.Do(func() {
		cache_Data_Interval_Duration_semigroupDuration = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), gopurs_runtime.RecordGet(Get_Data_Interval_Duration_ordDurationComponent(), "compare"), Get_Data_Semiring_numAdd(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_1))})))}
})
}))
	})
	return cache_Data_Interval_Duration_semigroupDuration
}

var cache_Data_Interval_Duration_monoidDuration gopurs_runtime.Value
var once_Data_Interval_Duration_monoidDuration sync.Once
func Get_Data_Interval_Duration_monoidDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_monoidDuration.Do(func() {
		cache_Data_Interval_Duration_monoidDuration = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_Duration_semigroupDuration()
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
	})
	return cache_Data_Interval_Duration_monoidDuration
}

var cache_Data_Interval_Duration_eqDuration gopurs_runtime.Value
var once_Data_Interval_Duration_eqDuration sync.Once
func Get_Data_Interval_Duration_eqDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_eqDuration.Do(func() {
		cache_Data_Interval_Duration_eqDuration = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Interval_Duration_eqMap()).V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1))}).IntVal) != (0))
})
}))
	})
	return cache_Data_Interval_Duration_eqDuration
}

var cache_Data_Interval_Duration_ordDuration gopurs_runtime.Value
var once_Data_Interval_Duration_ordDuration sync.Once
func Get_Data_Interval_Duration_ordDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_ordDuration.Do(func() {
		cache_Data_Interval_Duration_ordDuration = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_Duration_eqDuration()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Interval_Duration_ordMap()).V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1))}).IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Interval_Duration_ordDuration
}

var cache_Data_Interval_Duration_durationFromComponent gopurs_runtime.Value
var once_Data_Interval_Duration_durationFromComponent sync.Once
func Get_Data_Interval_Duration_durationFromComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_durationFromComponent.Do(func() {
		cache_Data_Interval_Duration_durationFromComponent = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_durationFromComponent(uint32(k_0_box.IntVal), v_1_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_durationFromComponent
}

var cache_Data_Interval_Duration_hour gopurs_runtime.Value
var once_Data_Interval_Duration_hour sync.Once
func Get_Data_Interval_Duration_hour() gopurs_runtime.Value {
	once_Data_Interval_Duration_hour.Do(func() {
		cache_Data_Interval_Duration_hour = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_hour(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_hour
}

var cache_Data_Interval_Duration_millisecond gopurs_runtime.Value
var once_Data_Interval_Duration_millisecond sync.Once
func Get_Data_Interval_Duration_millisecond() gopurs_runtime.Value {
	once_Data_Interval_Duration_millisecond.Do(func() {
		cache_Data_Interval_Duration_millisecond = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_millisecond(x_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_millisecond
}

var cache_Data_Interval_Duration_minute gopurs_runtime.Value
var once_Data_Interval_Duration_minute sync.Once
func Get_Data_Interval_Duration_minute() gopurs_runtime.Value {
	once_Data_Interval_Duration_minute.Do(func() {
		cache_Data_Interval_Duration_minute = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_minute(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_minute
}

var cache_Data_Interval_Duration_month gopurs_runtime.Value
var once_Data_Interval_Duration_month sync.Once
func Get_Data_Interval_Duration_month() gopurs_runtime.Value {
	once_Data_Interval_Duration_month.Do(func() {
		cache_Data_Interval_Duration_month = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_month(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_month
}

var cache_Data_Interval_Duration_second gopurs_runtime.Value
var once_Data_Interval_Duration_second sync.Once
func Get_Data_Interval_Duration_second() gopurs_runtime.Value {
	once_Data_Interval_Duration_second.Do(func() {
		cache_Data_Interval_Duration_second = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_second(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_second
}

var cache_Data_Interval_Duration_week gopurs_runtime.Value
var once_Data_Interval_Duration_week sync.Once
func Get_Data_Interval_Duration_week() gopurs_runtime.Value {
	once_Data_Interval_Duration_week.Do(func() {
		cache_Data_Interval_Duration_week = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_week(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_week
}

var cache_Data_Interval_Duration_year gopurs_runtime.Value
var once_Data_Interval_Duration_year sync.Once
func Get_Data_Interval_Duration_year() gopurs_runtime.Value {
	once_Data_Interval_Duration_year.Do(func() {
		cache_Data_Interval_Duration_year = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_year(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_year
}

var cache_Data_Interval_Duration_day gopurs_runtime.Value
var once_Data_Interval_Duration_day sync.Once
func Get_Data_Interval_Duration_day() gopurs_runtime.Value {
	once_Data_Interval_Duration_day.Do(func() {
		cache_Data_Interval_Duration_day = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_day(v_0_box.FloatVal()))}
})
	})
	return cache_Data_Interval_Duration_day
}

type Constructor_Data_Interval_Duration_Second struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Minute struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Hour struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Day struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Week struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Month struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Year struct {
	Rc uint32
}


func Call_Data_Interval_Duration_Duration(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Interval_Duration_durationFromComponent(k_0_loop uint32, v_1_loop float64) *Constructor_Data_Map_Internal_Node {
var k_0 uint32 = k_0_loop
_ = k_0
var v_1 float64 = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(k_0), UnsafePtr: nil}, gopurs_runtime.Float(v_1), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})})})
}

func Call_Data_Interval_Duration_hour(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(1292308612), UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})})})
}

func Call_Data_Interval_Duration_millisecond(x_0_loop float64) *Constructor_Data_Map_Internal_Node {
var x_0 float64 = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil}, gopurs_runtime.Float(gopurs_runtime.Float((gopurs_runtime.Float(x_0).FloatVal()) / (1000.0)).FloatVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})})})
}

func Call_Data_Interval_Duration_minute(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(217821258), UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})})})
}

func Call_Data_Interval_Duration_month(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(3327533908), UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})})})
}

func Call_Data_Interval_Duration_second(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})})})
}

func Call_Data_Interval_Duration_week(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(401302776), UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})})})
}

func Call_Data_Interval_Duration_year(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(3631736139), UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})})})
}

func Call_Data_Interval_Duration_day(v_0_loop float64) *Constructor_Data_Map_Internal_Node {
var v_0 float64 = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: int64(2311060696), UnsafePtr: nil}, gopurs_runtime.Float(v_0), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})})})
}



package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Time_Component_Second gopurs_runtime.Value
var once_Data_Time_Component_Second sync.Once
func Get_Data_Time_Component_Second() gopurs_runtime.Value {
	once_Data_Time_Component_Second.Do(func() {
		cache_Data_Time_Component_Second = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Component_Second(x_0_box)
})
	})
	return cache_Data_Time_Component_Second
}

var cache_Data_Time_Component_Minute gopurs_runtime.Value
var once_Data_Time_Component_Minute sync.Once
func Get_Data_Time_Component_Minute() gopurs_runtime.Value {
	once_Data_Time_Component_Minute.Do(func() {
		cache_Data_Time_Component_Minute = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Component_Minute(x_0_box)
})
	})
	return cache_Data_Time_Component_Minute
}

var cache_Data_Time_Component_Millisecond gopurs_runtime.Value
var once_Data_Time_Component_Millisecond sync.Once
func Get_Data_Time_Component_Millisecond() gopurs_runtime.Value {
	once_Data_Time_Component_Millisecond.Do(func() {
		cache_Data_Time_Component_Millisecond = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Component_Millisecond(x_0_box)
})
	})
	return cache_Data_Time_Component_Millisecond
}

var cache_Data_Time_Component_Hour gopurs_runtime.Value
var once_Data_Time_Component_Hour sync.Once
func Get_Data_Time_Component_Hour() gopurs_runtime.Value {
	once_Data_Time_Component_Hour.Do(func() {
		cache_Data_Time_Component_Hour = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Component_Hour(x_0_box)
})
	})
	return cache_Data_Time_Component_Hour
}

var cache_Data_Time_Component_showSecond gopurs_runtime.Value
var once_Data_Time_Component_showSecond sync.Once
func Get_Data_Time_Component_showSecond() gopurs_runtime.Value {
	once_Data_Time_Component_showSecond.Do(func() {
		cache_Data_Time_Component_showSecond = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Second ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(v_0.IntVal)).StrVal())) + (")"))
})})}
	})
	return cache_Data_Time_Component_showSecond
}

var cache_Data_Time_Component_showMinute gopurs_runtime.Value
var once_Data_Time_Component_showMinute sync.Once
func Get_Data_Time_Component_showMinute() gopurs_runtime.Value {
	once_Data_Time_Component_showMinute.Do(func() {
		cache_Data_Time_Component_showMinute = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Minute ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(v_0.IntVal)).StrVal())) + (")"))
})})}
	})
	return cache_Data_Time_Component_showMinute
}

var cache_Data_Time_Component_showMillisecond gopurs_runtime.Value
var once_Data_Time_Component_showMillisecond sync.Once
func Get_Data_Time_Component_showMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_showMillisecond.Do(func() {
		cache_Data_Time_Component_showMillisecond = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Millisecond ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(v_0.IntVal)).StrVal())) + (")"))
})})}
	})
	return cache_Data_Time_Component_showMillisecond
}

var cache_Data_Time_Component_showHour gopurs_runtime.Value
var once_Data_Time_Component_showHour sync.Once
func Get_Data_Time_Component_showHour() gopurs_runtime.Value {
	once_Data_Time_Component_showHour.Do(func() {
		cache_Data_Time_Component_showHour = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Hour ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(v_0.IntVal)).StrVal())) + (")"))
})})}
	})
	return cache_Data_Time_Component_showHour
}

var cache_Data_Time_Component_ordSecond gopurs_runtime.Value
var once_Data_Time_Component_ordSecond sync.Once
func Get_Data_Time_Component_ordSecond() gopurs_runtime.Value {
	once_Data_Time_Component_ordSecond.Do(func() {
		cache_Data_Time_Component_ordSecond = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordInt()))}
	})
	return cache_Data_Time_Component_ordSecond
}

var cache_Data_Time_Component_ordMinute gopurs_runtime.Value
var once_Data_Time_Component_ordMinute sync.Once
func Get_Data_Time_Component_ordMinute() gopurs_runtime.Value {
	once_Data_Time_Component_ordMinute.Do(func() {
		cache_Data_Time_Component_ordMinute = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordInt()))}
	})
	return cache_Data_Time_Component_ordMinute
}

var cache_Data_Time_Component_ordMillisecond gopurs_runtime.Value
var once_Data_Time_Component_ordMillisecond sync.Once
func Get_Data_Time_Component_ordMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_ordMillisecond.Do(func() {
		cache_Data_Time_Component_ordMillisecond = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordInt()))}
	})
	return cache_Data_Time_Component_ordMillisecond
}

var cache_Data_Time_Component_ordHour gopurs_runtime.Value
var once_Data_Time_Component_ordHour sync.Once
func Get_Data_Time_Component_ordHour() gopurs_runtime.Value {
	once_Data_Time_Component_ordHour.Do(func() {
		cache_Data_Time_Component_ordHour = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordInt()))}
	})
	return cache_Data_Time_Component_ordHour
}

var cache_Data_Time_Component_eqSecond gopurs_runtime.Value
var once_Data_Time_Component_eqSecond sync.Once
func Get_Data_Time_Component_eqSecond() gopurs_runtime.Value {
	once_Data_Time_Component_eqSecond.Do(func() {
		cache_Data_Time_Component_eqSecond = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqInt()))}
	})
	return cache_Data_Time_Component_eqSecond
}

var cache_Data_Time_Component_eqMinute gopurs_runtime.Value
var once_Data_Time_Component_eqMinute sync.Once
func Get_Data_Time_Component_eqMinute() gopurs_runtime.Value {
	once_Data_Time_Component_eqMinute.Do(func() {
		cache_Data_Time_Component_eqMinute = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqInt()))}
	})
	return cache_Data_Time_Component_eqMinute
}

var cache_Data_Time_Component_eqMillisecond gopurs_runtime.Value
var once_Data_Time_Component_eqMillisecond sync.Once
func Get_Data_Time_Component_eqMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_eqMillisecond.Do(func() {
		cache_Data_Time_Component_eqMillisecond = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqInt()))}
	})
	return cache_Data_Time_Component_eqMillisecond
}

var cache_Data_Time_Component_eqHour gopurs_runtime.Value
var once_Data_Time_Component_eqHour sync.Once
func Get_Data_Time_Component_eqHour() gopurs_runtime.Value {
	once_Data_Time_Component_eqHour.Do(func() {
		cache_Data_Time_Component_eqHour = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Eq_eqInt()))}
	})
	return cache_Data_Time_Component_eqHour
}

var cache_Data_Time_Component_boundedSecond gopurs_runtime.Value
var once_Data_Time_Component_boundedSecond sync.Once
func Get_Data_Time_Component_boundedSecond() gopurs_runtime.Value {
	once_Data_Time_Component_boundedSecond.Do(func() {
		cache_Data_Time_Component_boundedSecond = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(&Constructor_Data_Bounded_Bounded{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordInt()))}
}), gopurs_runtime.Int(0), gopurs_runtime.Int(59)})}
	})
	return cache_Data_Time_Component_boundedSecond
}

var cache_Data_Time_Component_boundedMinute gopurs_runtime.Value
var once_Data_Time_Component_boundedMinute sync.Once
func Get_Data_Time_Component_boundedMinute() gopurs_runtime.Value {
	once_Data_Time_Component_boundedMinute.Do(func() {
		cache_Data_Time_Component_boundedMinute = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(&Constructor_Data_Bounded_Bounded{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordInt()))}
}), gopurs_runtime.Int(0), gopurs_runtime.Int(59)})}
	})
	return cache_Data_Time_Component_boundedMinute
}

var cache_Data_Time_Component_boundedMillisecond gopurs_runtime.Value
var once_Data_Time_Component_boundedMillisecond sync.Once
func Get_Data_Time_Component_boundedMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_boundedMillisecond.Do(func() {
		cache_Data_Time_Component_boundedMillisecond = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(&Constructor_Data_Bounded_Bounded{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordInt()))}
}), gopurs_runtime.Int(0), gopurs_runtime.Int(999)})}
	})
	return cache_Data_Time_Component_boundedMillisecond
}

var cache_Data_Time_Component_boundedHour gopurs_runtime.Value
var once_Data_Time_Component_boundedHour sync.Once
func Get_Data_Time_Component_boundedHour() gopurs_runtime.Value {
	once_Data_Time_Component_boundedHour.Do(func() {
		cache_Data_Time_Component_boundedHour = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(&Constructor_Data_Bounded_Bounded{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordInt()))}
}), gopurs_runtime.Int(0), gopurs_runtime.Int(23)})}
	})
	return cache_Data_Time_Component_boundedHour
}

var cache_Data_Time_Component_boundedEnumSecond gopurs_runtime.Value
var once_Data_Time_Component_boundedEnumSecond sync.Once
func Get_Data_Time_Component_boundedEnumSecond() gopurs_runtime.Value {
	once_Data_Time_Component_boundedEnumSecond.Do(func() {
		cache_Data_Time_Component_boundedEnumSecond = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(&Constructor_Data_Enum_BoundedEnum{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](Get_Data_Time_Component_boundedSecond()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](Get_Data_Time_Component_enumSecond()))}
}), gopurs_runtime.Int(60).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v_0.IntVal)
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (n_0.IntVal) < (0) {
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
if (n_0.IntVal) > (59) {
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
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(n_0.IntVal)}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})})}
	})
	return cache_Data_Time_Component_boundedEnumSecond
}

var cache_Data_Time_Component_enumSecond gopurs_runtime.Value
var once_Data_Time_Component_enumSecond sync.Once
func Get_Data_Time_Component_enumSecond() gopurs_runtime.Value {
	once_Data_Time_Component_enumSecond.Do(func() {
		cache_Data_Time_Component_enumSecond = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Enum_Enum{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordInt()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Int((gopurs_runtime.Int(x_0.IntVal).IntVal) - (1))
_ = __local_var_1_0
var __t4 *Constructor_Data_Maybe_Just
{
var __t1 bool
{
if (__local_var_1_0.IntVal) < (0) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
var __t_and_3 bool = false
if __t1 {

var __t2 bool
{
if (__local_var_1_0.IntVal) > (59) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
__t_and_3 = __t2
}
if __t_and_3 {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_1_0.IntVal)}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_5 -> gopurs_runtime.Value
__local_var_1_5 := gopurs_runtime.Int((x_0.IntVal) + (1))
_ = __local_var_1_5
var __t9 *Constructor_Data_Maybe_Just
{
var __t6 bool
{
if (__local_var_1_5.IntVal) < (0) {
__t6 = false
goto end_branch_6
} else {

}
}
{
__t6 = true
}
end_branch_6:
var __t_and_8 bool = false
if __t6 {

var __t7 bool
{
if (__local_var_1_5.IntVal) > (59) {
__t7 = false
goto end_branch_7
} else {

}
}
{
__t7 = true
}
end_branch_7:
__t_and_8 = __t7
}
if __t_and_8 {
__t9 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_1_5.IntVal)}
goto end_branch_9
} else {

}
}
{
__t9 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
})})}
	})
	return cache_Data_Time_Component_enumSecond
}

var cache_Data_Time_Component_boundedEnumMinute gopurs_runtime.Value
var once_Data_Time_Component_boundedEnumMinute sync.Once
func Get_Data_Time_Component_boundedEnumMinute() gopurs_runtime.Value {
	once_Data_Time_Component_boundedEnumMinute.Do(func() {
		cache_Data_Time_Component_boundedEnumMinute = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(&Constructor_Data_Enum_BoundedEnum{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](Get_Data_Time_Component_boundedMinute()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](Get_Data_Time_Component_enumMinute()))}
}), gopurs_runtime.Int(60).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v_0.IntVal)
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (n_0.IntVal) < (0) {
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
if (n_0.IntVal) > (59) {
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
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(n_0.IntVal)}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})})}
	})
	return cache_Data_Time_Component_boundedEnumMinute
}

var cache_Data_Time_Component_enumMinute gopurs_runtime.Value
var once_Data_Time_Component_enumMinute sync.Once
func Get_Data_Time_Component_enumMinute() gopurs_runtime.Value {
	once_Data_Time_Component_enumMinute.Do(func() {
		cache_Data_Time_Component_enumMinute = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Enum_Enum{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordInt()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Int((gopurs_runtime.Int(x_0.IntVal).IntVal) - (1))
_ = __local_var_1_0
var __t4 *Constructor_Data_Maybe_Just
{
var __t1 bool
{
if (__local_var_1_0.IntVal) < (0) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
var __t_and_3 bool = false
if __t1 {

var __t2 bool
{
if (__local_var_1_0.IntVal) > (59) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
__t_and_3 = __t2
}
if __t_and_3 {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_1_0.IntVal)}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_5 -> gopurs_runtime.Value
__local_var_1_5 := gopurs_runtime.Int((x_0.IntVal) + (1))
_ = __local_var_1_5
var __t9 *Constructor_Data_Maybe_Just
{
var __t6 bool
{
if (__local_var_1_5.IntVal) < (0) {
__t6 = false
goto end_branch_6
} else {

}
}
{
__t6 = true
}
end_branch_6:
var __t_and_8 bool = false
if __t6 {

var __t7 bool
{
if (__local_var_1_5.IntVal) > (59) {
__t7 = false
goto end_branch_7
} else {

}
}
{
__t7 = true
}
end_branch_7:
__t_and_8 = __t7
}
if __t_and_8 {
__t9 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_1_5.IntVal)}
goto end_branch_9
} else {

}
}
{
__t9 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
})})}
	})
	return cache_Data_Time_Component_enumMinute
}

var cache_Data_Time_Component_boundedEnumMillisecond gopurs_runtime.Value
var once_Data_Time_Component_boundedEnumMillisecond sync.Once
func Get_Data_Time_Component_boundedEnumMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_boundedEnumMillisecond.Do(func() {
		cache_Data_Time_Component_boundedEnumMillisecond = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(&Constructor_Data_Enum_BoundedEnum{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](Get_Data_Time_Component_boundedMillisecond()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](Get_Data_Time_Component_enumMillisecond()))}
}), gopurs_runtime.Int(1000).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v_0.IntVal)
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (n_0.IntVal) < (0) {
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
if (n_0.IntVal) > (999) {
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
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(n_0.IntVal)}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})})}
	})
	return cache_Data_Time_Component_boundedEnumMillisecond
}

var cache_Data_Time_Component_enumMillisecond gopurs_runtime.Value
var once_Data_Time_Component_enumMillisecond sync.Once
func Get_Data_Time_Component_enumMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_enumMillisecond.Do(func() {
		cache_Data_Time_Component_enumMillisecond = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Enum_Enum{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordInt()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Int((gopurs_runtime.Int(x_0.IntVal).IntVal) - (1))
_ = __local_var_1_0
var __t4 *Constructor_Data_Maybe_Just
{
var __t1 bool
{
if (__local_var_1_0.IntVal) < (0) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
var __t_and_3 bool = false
if __t1 {

var __t2 bool
{
if (__local_var_1_0.IntVal) > (999) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
__t_and_3 = __t2
}
if __t_and_3 {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_1_0.IntVal)}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_5 -> gopurs_runtime.Value
__local_var_1_5 := gopurs_runtime.Int((x_0.IntVal) + (1))
_ = __local_var_1_5
var __t9 *Constructor_Data_Maybe_Just
{
var __t6 bool
{
if (__local_var_1_5.IntVal) < (0) {
__t6 = false
goto end_branch_6
} else {

}
}
{
__t6 = true
}
end_branch_6:
var __t_and_8 bool = false
if __t6 {

var __t7 bool
{
if (__local_var_1_5.IntVal) > (999) {
__t7 = false
goto end_branch_7
} else {

}
}
{
__t7 = true
}
end_branch_7:
__t_and_8 = __t7
}
if __t_and_8 {
__t9 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_1_5.IntVal)}
goto end_branch_9
} else {

}
}
{
__t9 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
})})}
	})
	return cache_Data_Time_Component_enumMillisecond
}

var cache_Data_Time_Component_boundedEnumHour gopurs_runtime.Value
var once_Data_Time_Component_boundedEnumHour sync.Once
func Get_Data_Time_Component_boundedEnumHour() gopurs_runtime.Value {
	once_Data_Time_Component_boundedEnumHour.Do(func() {
		cache_Data_Time_Component_boundedEnumHour = gopurs_runtime.Value{Type: 9, IntVal: 287434377, UnsafePtr: unsafe.Pointer(&Constructor_Data_Enum_BoundedEnum{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](Get_Data_Time_Component_boundedHour()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](Get_Data_Time_Component_enumHour()))}
}), gopurs_runtime.Int(24).IntVal, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v_0.IntVal)
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (n_0.IntVal) < (0) {
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
if (n_0.IntVal) > (23) {
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
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(n_0.IntVal)}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})})}
	})
	return cache_Data_Time_Component_boundedEnumHour
}

var cache_Data_Time_Component_enumHour gopurs_runtime.Value
var once_Data_Time_Component_enumHour sync.Once
func Get_Data_Time_Component_enumHour() gopurs_runtime.Value {
	once_Data_Time_Component_enumHour.Do(func() {
		cache_Data_Time_Component_enumHour = gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Enum_Enum{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Ord_ordInt()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Int((gopurs_runtime.Int(x_0.IntVal).IntVal) - (1))
_ = __local_var_1_0
var __t4 *Constructor_Data_Maybe_Just
{
var __t1 bool
{
if (__local_var_1_0.IntVal) < (0) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
var __t_and_3 bool = false
if __t1 {

var __t2 bool
{
if (__local_var_1_0.IntVal) > (23) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
__t_and_3 = __t2
}
if __t_and_3 {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_1_0.IntVal)}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_5 -> gopurs_runtime.Value
__local_var_1_5 := gopurs_runtime.Int((x_0.IntVal) + (1))
_ = __local_var_1_5
var __t9 *Constructor_Data_Maybe_Just
{
var __t6 bool
{
if (__local_var_1_5.IntVal) < (0) {
__t6 = false
goto end_branch_6
} else {

}
}
{
__t6 = true
}
end_branch_6:
var __t_and_8 bool = false
if __t6 {

var __t7 bool
{
if (__local_var_1_5.IntVal) > (23) {
__t7 = false
goto end_branch_7
} else {

}
}
{
__t7 = true
}
end_branch_7:
__t_and_8 = __t7
}
if __t_and_8 {
__t9 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_1_5.IntVal)}
goto end_branch_9
} else {

}
}
{
__t9 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
})})}
	})
	return cache_Data_Time_Component_enumHour
}

func Call_Data_Time_Component_Second(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Time_Component_Minute(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Time_Component_Millisecond(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Time_Component_Hour(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}



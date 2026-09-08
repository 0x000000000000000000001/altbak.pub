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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Call_Data_Interval_Duration_Duration(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](x_0_box))))}
})
	})
	return cache_Data_Interval_Duration_Duration
}

var cache_Data_Interval_Duration_showDurationComponent gopurs_runtime.Value
var once_Data_Interval_Duration_showDurationComponent sync.Once
func Get_Data_Interval_Duration_showDurationComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_showDurationComponent.Do(func() {
		cache_Data_Interval_Duration_showDurationComponent = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_1209612131_1386611502((&Constructor_Data_Show_Show[uint32]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 string
{
var __t_tag_0 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_0) == 217821258) {
__t7 = "Minute"
goto end_branch_7
} else {

}
}
{
var __t_tag_1 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_1) == 3908053364) {
__t7 = "Second"
goto end_branch_7
} else {

}
}
{
var __t_tag_2 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_2) == 1292308612) {
__t7 = "Hour"
goto end_branch_7
} else {

}
}
{
var __t_tag_3 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_3) == 2311060696) {
__t7 = "Day"
goto end_branch_7
} else {

}
}
{
var __t_tag_4 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_4) == 401302776) {
__t7 = "Week"
goto end_branch_7
} else {

}
}
{
var __t_tag_5 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_5) == 3327533908) {
__t7 = "Month"
goto end_branch_7
} else {

}
}
{
var __t_tag_6 uint32 = uint32(v_0.IntVal)
if (uint32(__t_tag_6) == 3631736139) {
__t7 = "Year"
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
	return cache_Data_Interval_Duration_showDurationComponent
}

var cache_Data_Interval_Duration_showMap gopurs_runtime.Value
var once_Data_Interval_Duration_showMap sync.Once
func Get_Data_Interval_Duration_showMap() gopurs_runtime.Value {
	once_Data_Interval_Duration_showMap.Do(func() {
		cache_Data_Interval_Duration_showMap = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_983382226_1386611502(Rebox_Data_Interval_Duration_1386611502_983382226(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_Map_Internal_showMap(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_1209612131_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[uint32]](Get_Data_Interval_Duration_showDurationComponent())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3263178038_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[float64]](Get_Data_Show_showNumber())))})))))}
	})
	return cache_Data_Interval_Duration_showMap
}

var cache_Data_Interval_Duration_showDuration gopurs_runtime.Value
var once_Data_Interval_Duration_showDuration sync.Once
func Get_Data_Interval_Duration_showDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_showDuration.Do(func() {
		cache_Data_Interval_Duration_showDuration = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_983382226_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[uint32, float64]]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Duration ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Data_Map_Internal_showMap(), gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_1209612131_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[uint32]](Get_Data_Interval_Duration_showDurationComponent())))}, gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3263178038_1386611502(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[float64]](Get_Data_Show_showNumber())))}), "show"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Rebox_Data_Interval_Duration_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0))))}).StrVal())) + (")"))
})})))}
	})
	return cache_Data_Interval_Duration_showDuration
}

var cache_Data_Interval_Duration_newtypeDuration gopurs_runtime.Value
var once_Data_Interval_Duration_newtypeDuration sync.Once
func Get_Data_Interval_Duration_newtypeDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_newtypeDuration.Do(func() {
		cache_Data_Interval_Duration_newtypeDuration = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_1236362600_385277032((&Constructor_Data_Newtype_Newtype[*Constructor_Data_Map_Internal_Node[uint32, float64], *Constructor_Data_Map_Internal_Node[uint32, float64]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})))}
	})
	return cache_Data_Interval_Duration_newtypeDuration
}

var cache_Data_Interval_Duration_eqDurationComponent gopurs_runtime.Value
var once_Data_Interval_Duration_eqDurationComponent sync.Once
func Get_Data_Interval_Duration_eqDurationComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_eqDurationComponent.Do(func() {
		cache_Data_Interval_Duration_eqDurationComponent = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 bool
{
var __t_tag_3 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_3) == 3908053364) {
var __t_tag_4 uint32 = uint32(y_1.IntVal)
__t15 = (uint32(__t_tag_4) == 3908053364)
goto end_branch_15
} else {

}
}
{
var __t_tag_5 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_5) == 217821258) {
var __t_tag_6 uint32 = uint32(y_1.IntVal)
__t15 = (uint32(__t_tag_6) == 217821258)
goto end_branch_15
} else {

}
}
{
var __t_tag_7 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_7) == 1292308612) {
var __t_tag_8 uint32 = uint32(y_1.IntVal)
__t15 = (uint32(__t_tag_8) == 1292308612)
goto end_branch_15
} else {

}
}
{
var __t_tag_9 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_9) == 2311060696) {
var __t_tag_10 uint32 = uint32(y_1.IntVal)
__t15 = (uint32(__t_tag_10) == 2311060696)
goto end_branch_15
} else {

}
}
{
var __t_tag_11 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_11) == 401302776) {
var __t_tag_12 uint32 = uint32(y_1.IntVal)
__t15 = (uint32(__t_tag_12) == 401302776)
goto end_branch_15
} else {

}
}
{
var __t_tag_13 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_13) == 3327533908) {
var __t_tag_14 uint32 = uint32(y_1.IntVal)
__t15 = (uint32(__t_tag_14) == 3327533908)
goto end_branch_15
} else {

}
}
{
var __t_tag_0 uint32 = uint32(x_0.IntVal)
var __t_and_2 bool = false
if (uint32(__t_tag_0) == 3631736139) {

var __t_tag_1 uint32 = uint32(y_1.IntVal)
__t_and_2 = (uint32(__t_tag_1) == 3631736139)
}
__t15 = __t_and_2
}
end_branch_15:
return gopurs_runtime.Bool(__t15)
})})))}
	})
	return cache_Data_Interval_Duration_eqDurationComponent
}

var cache_Data_Interval_Duration_eqMap gopurs_runtime.Value
var once_Data_Interval_Duration_eqMap sync.Once
func Get_Data_Interval_Duration_eqMap() gopurs_runtime.Value {
	once_Data_Interval_Duration_eqMap.Do(func() {
		cache_Data_Interval_Duration_eqMap = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_2958538738_3790796878(Rebox_Data_Interval_Duration_3790796878_2958538738(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_Map_Internal_eqMap(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3768443459_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[uint32]](Get_Data_Interval_Duration_eqDurationComponent())))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_687527510_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[float64]](Get_Data_Eq_eqNumber())))})))))}
	})
	return cache_Data_Interval_Duration_eqMap
}

var cache_Data_Interval_Duration_ordDurationComponent gopurs_runtime.Value
var once_Data_Interval_Duration_ordDurationComponent sync.Once
func Get_Data_Interval_Duration_ordDurationComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_ordDurationComponent.Do(func() {
		cache_Data_Interval_Duration_ordDurationComponent = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3730953251_4177771502((&Constructor_Data_Ord_Ord[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3768443459_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[uint32]](Get_Data_Interval_Duration_eqDurationComponent())))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 uint32
{
var __t_tag_0 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_0) == 3908053364) {
var __t2 uint32
{
var __t_tag_1 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_1) == 3908053364) {
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
if (uint32(__t_tag_3) == 3908053364) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_4 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_4) == 217821258) {
var __t6 uint32
{
var __t_tag_5 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_5) == 217821258) {
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
if (uint32(__t_tag_7) == 217821258) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_8 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_8) == 1292308612) {
var __t10 uint32
{
var __t_tag_9 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_9) == 1292308612) {
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
if (uint32(__t_tag_11) == 1292308612) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_12 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_12) == 2311060696) {
var __t14 uint32
{
var __t_tag_13 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_13) == 2311060696) {
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
if (uint32(__t_tag_15) == 2311060696) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_16 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_16) == 401302776) {
var __t18 uint32
{
var __t_tag_17 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_17) == 401302776) {
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
if (uint32(__t_tag_19) == 401302776) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_20 uint32 = uint32(x_0.IntVal)
if (uint32(__t_tag_20) == 3327533908) {
var __t22 uint32
{
var __t_tag_21 uint32 = uint32(y_1.IntVal)
if (uint32(__t_tag_21) == 3327533908) {
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
if (uint32(__t_tag_23) == 3327533908) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_24 uint32 = uint32(x_0.IntVal)
var __t_and_26 bool = false
if (uint32(__t_tag_24) == 3631736139) {

var __t_tag_25 uint32 = uint32(y_1.IntVal)
__t_and_26 = (uint32(__t_tag_25) == 3631736139)
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
	return cache_Data_Interval_Duration_ordDurationComponent
}

var cache_Data_Interval_Duration_ordMap gopurs_runtime.Value
var once_Data_Interval_Duration_ordMap sync.Once
func Get_Data_Interval_Duration_ordMap() gopurs_runtime.Value {
	once_Data_Interval_Duration_ordMap.Do(func() {
		cache_Data_Interval_Duration_ordMap = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3498527378_4177771502(Rebox_Data_Interval_Duration_4177771502_3498527378(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_Map_Internal_ordMap(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3730953251_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](Get_Data_Interval_Duration_ordDurationComponent())))}, gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3047586294_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[float64]](Get_Data_Ord_ordNumber())))})))))}
	})
	return cache_Data_Interval_Duration_ordMap
}

var cache_Data_Interval_Duration_semigroupDuration gopurs_runtime.Value
var once_Data_Interval_Duration_semigroupDuration sync.Once
func Get_Data_Interval_Duration_semigroupDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_semigroupDuration.Do(func() {
		cache_Data_Interval_Duration_semigroupDuration = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_1201208274_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[uint32, float64]]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Rebox_Data_Interval_Duration_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 uint32
{
var __t_tag_0 uint32 = uint32(x_2.IntVal)
if (uint32(__t_tag_0) == 3908053364) {
var __t2 uint32
{
var __t_tag_1 uint32 = uint32(y_3.IntVal)
if (uint32(__t_tag_1) == 3908053364) {
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
var __t_tag_3 uint32 = uint32(y_3.IntVal)
if (uint32(__t_tag_3) == 3908053364) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_4 uint32 = uint32(x_2.IntVal)
if (uint32(__t_tag_4) == 217821258) {
var __t6 uint32
{
var __t_tag_5 uint32 = uint32(y_3.IntVal)
if (uint32(__t_tag_5) == 217821258) {
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
var __t_tag_7 uint32 = uint32(y_3.IntVal)
if (uint32(__t_tag_7) == 217821258) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_8 uint32 = uint32(x_2.IntVal)
if (uint32(__t_tag_8) == 1292308612) {
var __t10 uint32
{
var __t_tag_9 uint32 = uint32(y_3.IntVal)
if (uint32(__t_tag_9) == 1292308612) {
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
var __t_tag_11 uint32 = uint32(y_3.IntVal)
if (uint32(__t_tag_11) == 1292308612) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_12 uint32 = uint32(x_2.IntVal)
if (uint32(__t_tag_12) == 2311060696) {
var __t14 uint32
{
var __t_tag_13 uint32 = uint32(y_3.IntVal)
if (uint32(__t_tag_13) == 2311060696) {
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
var __t_tag_15 uint32 = uint32(y_3.IntVal)
if (uint32(__t_tag_15) == 2311060696) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_16 uint32 = uint32(x_2.IntVal)
if (uint32(__t_tag_16) == 401302776) {
var __t18 uint32
{
var __t_tag_17 uint32 = uint32(y_3.IntVal)
if (uint32(__t_tag_17) == 401302776) {
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
var __t_tag_19 uint32 = uint32(y_3.IntVal)
if (uint32(__t_tag_19) == 401302776) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_20 uint32 = uint32(x_2.IntVal)
if (uint32(__t_tag_20) == 3327533908) {
var __t22 uint32
{
var __t_tag_21 uint32 = uint32(y_3.IntVal)
if (uint32(__t_tag_21) == 3327533908) {
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
var __t_tag_23 uint32 = uint32(y_3.IntVal)
if (uint32(__t_tag_23) == 3327533908) {
__t27 = 380165415
goto end_branch_27
} else {

}
}
{
var __t_tag_24 uint32 = uint32(x_2.IntVal)
var __t_and_26 bool = false
if (uint32(__t_tag_24) == 3631736139) {

var __t_tag_25 uint32 = uint32(y_3.IntVal)
__t_and_26 = (uint32(__t_tag_25) == 3631736139)
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
}), Get_Data_Semiring_numAdd(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Rebox_Data_Interval_Duration_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0))))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Rebox_Data_Interval_Duration_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_1))))})))))}
})})))}
	})
	return cache_Data_Interval_Duration_semigroupDuration
}

var cache_Data_Interval_Duration_monoidDuration gopurs_runtime.Value
var once_Data_Interval_Duration_monoidDuration sync.Once
func Get_Data_Interval_Duration_monoidDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_monoidDuration.Do(func() {
		cache_Data_Interval_Duration_monoidDuration = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_4248054386_1201789390((&Constructor_Data_Monoid_Monoid[*Constructor_Data_Map_Internal_Node[uint32, float64]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_1201208274_4179793454(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[uint32, float64]]](Get_Data_Interval_Duration_semigroupDuration())))}
}), Rebox_Data_Interval_Duration_2487766124_261879545((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})))}
	})
	return cache_Data_Interval_Duration_monoidDuration
}

var cache_Data_Interval_Duration_eqDuration gopurs_runtime.Value
var once_Data_Interval_Duration_eqDuration sync.Once
func Get_Data_Interval_Duration_eqDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_eqDuration.Do(func() {
		cache_Data_Interval_Duration_eqDuration = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_2958538738_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[uint32, float64]]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Data_Map_Internal_eqMap(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3768443459_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[uint32]](Get_Data_Interval_Duration_eqDurationComponent())))}, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_687527510_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[float64]](Get_Data_Eq_eqNumber())))}), "eq"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Rebox_Data_Interval_Duration_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_0))))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Rebox_Data_Interval_Duration_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](y_1))))}).IntVal) != (0))
})})))}
	})
	return cache_Data_Interval_Duration_eqDuration
}

var cache_Data_Interval_Duration_ordDuration gopurs_runtime.Value
var once_Data_Interval_Duration_ordDuration sync.Once
func Get_Data_Interval_Duration_ordDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_ordDuration.Do(func() {
		cache_Data_Interval_Duration_ordDuration = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3498527378_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[uint32, float64]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_2958538738_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[uint32, float64]]](Get_Data_Interval_Duration_eqDuration())))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Data_Map_Internal_ordMap(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3730953251_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](Get_Data_Interval_Duration_ordDurationComponent())))}, gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_3047586294_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[float64]](Get_Data_Ord_ordNumber())))}), "compare"), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Rebox_Data_Interval_Duration_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_0))))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Rebox_Data_Interval_Duration_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](y_1))))}).IntVal)), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Interval_Duration_ordDuration
}

var cache_Data_Interval_Duration_durationFromComponent gopurs_runtime.Value
var once_Data_Interval_Duration_durationFromComponent sync.Once
func Get_Data_Interval_Duration_durationFromComponent() gopurs_runtime.Value {
	once_Data_Interval_Duration_durationFromComponent.Do(func() {
		cache_Data_Interval_Duration_durationFromComponent = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Call_Data_Interval_Duration_durationFromComponent(uint32(k_0_box.IntVal), v_1_box.FloatVal())))}
})
	})
	return cache_Data_Interval_Duration_durationFromComponent
}

var cache_Data_Interval_Duration_hour gopurs_runtime.Value
var once_Data_Interval_Duration_hour sync.Once
func Get_Data_Interval_Duration_hour() gopurs_runtime.Value {
	once_Data_Interval_Duration_hour.Do(func() {
		cache_Data_Interval_Duration_hour = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Call_Data_Interval_Duration_hour(v_0_box.FloatVal())))}
})
	})
	return cache_Data_Interval_Duration_hour
}

var cache_Data_Interval_Duration_millisecond gopurs_runtime.Value
var once_Data_Interval_Duration_millisecond sync.Once
func Get_Data_Interval_Duration_millisecond() gopurs_runtime.Value {
	once_Data_Interval_Duration_millisecond.Do(func() {
		cache_Data_Interval_Duration_millisecond = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Call_Data_Interval_Duration_millisecond(x_0_box.FloatVal())))}
})
	})
	return cache_Data_Interval_Duration_millisecond
}

var cache_Data_Interval_Duration_minute gopurs_runtime.Value
var once_Data_Interval_Duration_minute sync.Once
func Get_Data_Interval_Duration_minute() gopurs_runtime.Value {
	once_Data_Interval_Duration_minute.Do(func() {
		cache_Data_Interval_Duration_minute = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Call_Data_Interval_Duration_minute(v_0_box.FloatVal())))}
})
	})
	return cache_Data_Interval_Duration_minute
}

var cache_Data_Interval_Duration_month gopurs_runtime.Value
var once_Data_Interval_Duration_month sync.Once
func Get_Data_Interval_Duration_month() gopurs_runtime.Value {
	once_Data_Interval_Duration_month.Do(func() {
		cache_Data_Interval_Duration_month = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Call_Data_Interval_Duration_month(v_0_box.FloatVal())))}
})
	})
	return cache_Data_Interval_Duration_month
}

var cache_Data_Interval_Duration_second gopurs_runtime.Value
var once_Data_Interval_Duration_second sync.Once
func Get_Data_Interval_Duration_second() gopurs_runtime.Value {
	once_Data_Interval_Duration_second.Do(func() {
		cache_Data_Interval_Duration_second = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Call_Data_Interval_Duration_second(v_0_box.FloatVal())))}
})
	})
	return cache_Data_Interval_Duration_second
}

var cache_Data_Interval_Duration_week gopurs_runtime.Value
var once_Data_Interval_Duration_week sync.Once
func Get_Data_Interval_Duration_week() gopurs_runtime.Value {
	once_Data_Interval_Duration_week.Do(func() {
		cache_Data_Interval_Duration_week = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Call_Data_Interval_Duration_week(v_0_box.FloatVal())))}
})
	})
	return cache_Data_Interval_Duration_week
}

var cache_Data_Interval_Duration_year gopurs_runtime.Value
var once_Data_Interval_Duration_year sync.Once
func Get_Data_Interval_Duration_year() gopurs_runtime.Value {
	once_Data_Interval_Duration_year.Do(func() {
		cache_Data_Interval_Duration_year = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Call_Data_Interval_Duration_year(v_0_box.FloatVal())))}
})
	})
	return cache_Data_Interval_Duration_year
}

var cache_Data_Interval_Duration_day gopurs_runtime.Value
var once_Data_Interval_Duration_day sync.Once
func Get_Data_Interval_Duration_day() gopurs_runtime.Value {
	once_Data_Interval_Duration_day.Do(func() {
		cache_Data_Interval_Duration_day = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(Call_Data_Interval_Duration_day(v_0_box.FloatVal())))}
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


func Call_Data_Interval_Duration_Duration(x_0_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_Map_Internal_Node[uint32, float64] {
var x_0 *Constructor_Data_Map_Internal_Node[uint32, float64] = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Interval_Duration_durationFromComponent(k_0_loop uint32, v_1_loop float64) *Constructor_Data_Map_Internal_Node[uint32, float64] {
var k_0 uint32 = k_0_loop
_ = k_0
var v_1 float64 = v_1_loop
_ = v_1
return (&Constructor_Data_Map_Internal_Node[uint32, float64]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(k_0), UnsafePtr: nil}.IntVal, gopurs_runtime.Float(v_1).IntVal, uint32(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.FloatVal(), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)})
}

func Call_Data_Interval_Duration_hour(v_0_loop float64) *Constructor_Data_Map_Internal_Node[uint32, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (&Constructor_Data_Map_Internal_Node[uint32, float64]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1292308612), UnsafePtr: nil}.IntVal, gopurs_runtime.Float(v_0).IntVal, uint32(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.FloatVal(), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)})
}

func Call_Data_Interval_Duration_millisecond(x_0_loop float64) *Constructor_Data_Map_Internal_Node[uint32, float64] {
var x_0 float64 = x_0_loop
_ = x_0
return (&Constructor_Data_Map_Internal_Node[uint32, float64]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil}.IntVal, gopurs_runtime.Float(gopurs_runtime.Float((gopurs_runtime.Float(x_0).FloatVal()) / (1000.0)).FloatVal()).IntVal, uint32(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.FloatVal(), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)})
}

func Call_Data_Interval_Duration_minute(v_0_loop float64) *Constructor_Data_Map_Internal_Node[uint32, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (&Constructor_Data_Map_Internal_Node[uint32, float64]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(217821258), UnsafePtr: nil}.IntVal, gopurs_runtime.Float(v_0).IntVal, uint32(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.FloatVal(), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)})
}

func Call_Data_Interval_Duration_month(v_0_loop float64) *Constructor_Data_Map_Internal_Node[uint32, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (&Constructor_Data_Map_Internal_Node[uint32, float64]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3327533908), UnsafePtr: nil}.IntVal, gopurs_runtime.Float(v_0).IntVal, uint32(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.FloatVal(), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)})
}

func Call_Data_Interval_Duration_second(v_0_loop float64) *Constructor_Data_Map_Internal_Node[uint32, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (&Constructor_Data_Map_Internal_Node[uint32, float64]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3908053364), UnsafePtr: nil}.IntVal, gopurs_runtime.Float(v_0).IntVal, uint32(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.FloatVal(), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)})
}

func Call_Data_Interval_Duration_week(v_0_loop float64) *Constructor_Data_Map_Internal_Node[uint32, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (&Constructor_Data_Map_Internal_Node[uint32, float64]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(401302776), UnsafePtr: nil}.IntVal, gopurs_runtime.Float(v_0).IntVal, uint32(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.FloatVal(), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)})
}

func Call_Data_Interval_Duration_year(v_0_loop float64) *Constructor_Data_Map_Internal_Node[uint32, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (&Constructor_Data_Map_Internal_Node[uint32, float64]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(3631736139), UnsafePtr: nil}.IntVal, gopurs_runtime.Float(v_0).IntVal, uint32(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.FloatVal(), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)})
}

func Call_Data_Interval_Duration_day(v_0_loop float64) *Constructor_Data_Map_Internal_Node[uint32, float64] {
var v_0 float64 = v_0_loop
_ = v_0
return (&Constructor_Data_Map_Internal_Node[uint32, float64]{1, gopurs_runtime.Value{Type: 9, IntVal: int64(2311060696), UnsafePtr: nil}.IntVal, gopurs_runtime.Float(v_0).IntVal, uint32(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124((*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)))}.FloatVal(), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil), (*Constructor_Data_Map_Internal_Node[uint32, float64])(nil)})
}

func Rebox_Data_Interval_Duration_1201208274_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[uint32, float64]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_Duration_1209612131_1386611502(in *Constructor_Data_Show_Show[uint32]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_Duration_1236362600_385277032(in *Constructor_Data_Newtype_Newtype[*Constructor_Data_Map_Internal_Node[uint32, float64], *Constructor_Data_Map_Internal_Node[uint32, float64]]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_Duration_1386611502_983382226(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[uint32, float64]] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[uint32, float64]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_Duration_2487766124_261879545(in *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Map_Internal_Node[uint32, float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Map_Internal_Node[uint32, float64]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = uint32(in.V2.IntVal)
		out.V3 = in.V3.FloatVal()
		out.V4 = Rebox_Data_Interval_Duration_2487766124_261879545(in.V4)
		out.V5 = Rebox_Data_Interval_Duration_2487766124_261879545(in.V5)
	return out
}

func Rebox_Data_Interval_Duration_261879545_2487766124(in *Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
		out.V3 = gopurs_runtime.Float(in.V3)
		out.V4 = Rebox_Data_Interval_Duration_261879545_2487766124(in.V4)
		out.V5 = Rebox_Data_Interval_Duration_261879545_2487766124(in.V5)
	return out
}

func Rebox_Data_Interval_Duration_2958538738_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[uint32, float64]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_Duration_3047586294_4177771502(in *Constructor_Data_Ord_Ord[float64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Interval_Duration_3263178038_1386611502(in *Constructor_Data_Show_Show[float64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_Duration_3498527378_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[uint32, float64]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Interval_Duration_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Interval_Duration_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_Duration_3790796878_2958538738(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[uint32, float64]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[uint32, float64]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_Duration_4177771502_3498527378(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[uint32, float64]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[uint32, float64]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Interval_Duration_4248054386_1201789390(in *Constructor_Data_Monoid_Monoid[*Constructor_Data_Map_Internal_Node[uint32, float64]]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_261879545_2487766124(in.V1))}
	return out
}

func Rebox_Data_Interval_Duration_687527510_3790796878(in *Constructor_Data_Eq_Eq[float64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_Duration_983382226_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[uint32, float64]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}



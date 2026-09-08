package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Interval_showMaybe gopurs_runtime.Value
var once_Data_Interval_showMaybe sync.Once
func Get_Data_Interval_showMaybe() gopurs_runtime.Value {
	once_Data_Interval_showMaybe.Do(func() {
		cache_Data_Interval_showMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_2818770644_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 string
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
if (__t_tag_0 != nil) {
__t2 = (("(Just ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
if (__t_tag_1 == nil) {
__t2 = "Nothing"
goto end_branch_2
} else {

}
}
{
__t2 = func() string { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Str(__t2)
})})))}
	})
	return cache_Data_Interval_showMaybe
}

var cache_Data_Interval_eqMaybe gopurs_runtime.Value
var once_Data_Interval_eqMaybe sync.Once
func Get_Data_Interval_eqMaybe() gopurs_runtime.Value {
	once_Data_Interval_eqMaybe.Do(func() {
		cache_Data_Interval_eqMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_3960201844_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[int64]]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 bool
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_0)
if (__t_tag_3 == nil) {
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_1)
__t5 = (__t_tag_4 == nil)
goto end_branch_5
} else {

}
}
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_0)
var __t_and_2 bool = false
if (__t_tag_0 != nil) {

var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_1)
__t_and_2 = ((__t_tag_1 != nil)) && (((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_0.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_1.UnsafePtr).V0.IntVal))
}
__t5 = __t_and_2
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
})})))}
	})
	return cache_Data_Interval_eqMaybe
}

var cache_Data_Interval_ordMaybe gopurs_runtime.Value
var once_Data_Interval_ordMaybe sync.Once
func Get_Data_Interval_ordMaybe() gopurs_runtime.Value {
	once_Data_Interval_ordMaybe.Do(func() {
		cache_Data_Interval_ordMaybe = func() gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_3092443796_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[int64]]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_3508461103_3790796878(eqMaybe1_0_0))}
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
	return cache_Data_Interval_ordMaybe
}

var cache_Data_Interval_StartEnd gopurs_runtime.Value
var once_Data_Interval_StartEnd sync.Once
func Get_Data_Interval_StartEnd() gopurs_runtime.Value {
	once_Data_Interval_StartEnd.Do(func() {
		cache_Data_Interval_StartEnd = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
})
})
	})
	return cache_Data_Interval_StartEnd
}

var cache_Data_Interval_DurationEnd gopurs_runtime.Value
var once_Data_Interval_DurationEnd sync.Once
func Get_Data_Interval_DurationEnd() gopurs_runtime.Value {
	once_Data_Interval_DurationEnd.Do(func() {
		cache_Data_Interval_DurationEnd = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
})
})
	})
	return cache_Data_Interval_DurationEnd
}

var cache_Data_Interval_StartDuration gopurs_runtime.Value
var once_Data_Interval_StartDuration sync.Once
func Get_Data_Interval_StartDuration() gopurs_runtime.Value {
	once_Data_Interval_StartDuration.Do(func() {
		cache_Data_Interval_StartDuration = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1}))}
})
})
	})
	return cache_Data_Interval_StartDuration
}

var cache_Data_Interval_DurationOnly gopurs_runtime.Value
var once_Data_Interval_DurationOnly sync.Once
func Get_Data_Interval_DurationOnly() gopurs_runtime.Value {
	once_Data_Interval_DurationOnly.Do(func() {
		cache_Data_Interval_DurationOnly = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
})
	})
	return cache_Data_Interval_DurationOnly
}

var cache_Data_Interval_RecurringInterval gopurs_runtime.Value
var once_Data_Interval_RecurringInterval sync.Once
func Get_Data_Interval_RecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_RecurringInterval.Do(func() {
		cache_Data_Interval_RecurringInterval = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, Rebox_Data_Interval_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](value0)), value1}))}
})
})
	})
	return cache_Data_Interval_RecurringInterval
}

var cache_Data_Interval_showInterval gopurs_runtime.Value
var once_Data_Interval_showInterval sync.Once
func Get_Data_Interval_showInterval() gopurs_runtime.Value {
	once_Data_Interval_showInterval.Do(func() {
		cache_Data_Interval_showInterval = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_showInterval(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_Data_Interval_showInterval
}

var cache_Data_Interval_showRecurringInterval gopurs_runtime.Value
var once_Data_Interval_showRecurringInterval sync.Once
func Get_Data_Interval_showRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_showRecurringInterval.Do(func() {
		cache_Data_Interval_showRecurringInterval = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_showRecurringInterval(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_Data_Interval_showRecurringInterval
}

var cache_Data_Interval_over gopurs_runtime.Value
var once_Data_Interval_over sync.Once
func Get_Data_Interval_over() gopurs_runtime.Value {
	once_Data_Interval_over.Do(func() {
		cache_Data_Interval_over = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_over(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box))
})
	})
	return cache_Data_Interval_over
}

var cache_Data_Interval_interval gopurs_runtime.Value
var once_Data_Interval_interval sync.Once
func Get_Data_Interval_interval() gopurs_runtime.Value {
	once_Data_Interval_interval.Do(func() {
		cache_Data_Interval_interval = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_interval(gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_Data_Interval_interval
}

var cache_Data_Interval_foldableInterval gopurs_runtime.Value
var once_Data_Interval_foldableInterval sync.Once
func Get_Data_Interval_foldableInterval() gopurs_runtime.Value {
	once_Data_Interval_foldableInterval.Do(func() {
		cache_Data_Interval_foldableInterval = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=Other bindingType=(TypeVar m)
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Interval_foldableInterval(), "foldl"), gopurs_runtime.Func2(func(acc_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), acc_4, gopurs_runtime.Apply(f_3, x_5))
}), mempty_2_1)
})
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = v1_1
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Foldable_foldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Interval_foldableInterval()))}, x_0)
})}))}
	})
	return cache_Data_Interval_foldableInterval
}

var cache_Data_Interval_foldableRecurringInterval gopurs_runtime.Value
var once_Data_Interval_foldableRecurringInterval sync.Once
func Get_Data_Interval_foldableRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_foldableRecurringInterval.Do(func() {
		cache_Data_Interval_foldableRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_170409538_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=Other bindingType=(TypeVar m)
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Interval_foldableRecurringInterval(), "foldl"), gopurs_runtime.Func2(func(acc_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), acc_4, gopurs_runtime.Apply(f_3, x_5))
}), mempty_2_1)
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, i_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 237113226) {
__t5 = gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply2(f_0, i_1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1.UnsafePtr).V1)
goto end_branch_5
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1992629780) {
__t5 = gopurs_runtime.Apply2(f_0, i_1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1.UnsafePtr).V1)
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2020675835) {
__t5 = gopurs_runtime.Apply2(f_0, i_1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = i_1
}
end_branch_5:
return __t5
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, i_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Data_Foldable_foldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Interval_foldableInterval()))}, f_0, i_1, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1)
})})))}
	})
	return cache_Data_Interval_foldableRecurringInterval
}

var cache_Data_Interval_eqInterval gopurs_runtime.Value
var once_Data_Interval_eqInterval sync.Once
func Get_Data_Interval_eqInterval() gopurs_runtime.Value {
	once_Data_Interval_eqInterval.Do(func() {
		cache_Data_Interval_eqInterval = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_eqInterval(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_Data_Interval_eqInterval
}

var cache_Data_Interval_eqRecurringInterval gopurs_runtime.Value
var once_Data_Interval_eqRecurringInterval sync.Once
func Get_Data_Interval_eqRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_eqRecurringInterval.Do(func() {
		cache_Data_Interval_eqRecurringInterval = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_eqRecurringInterval(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_Data_Interval_eqRecurringInterval
}

var cache_Data_Interval_ordInterval gopurs_runtime.Value
var once_Data_Interval_ordInterval sync.Once
func Get_Data_Interval_ordInterval() gopurs_runtime.Value {
	once_Data_Interval_ordInterval.Do(func() {
		cache_Data_Interval_ordInterval = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_ordInterval(dictOrd_0_box)
})
	})
	return cache_Data_Interval_ordInterval
}

var cache_Data_Interval_ordRecurringInterval gopurs_runtime.Value
var once_Data_Interval_ordRecurringInterval sync.Once
func Get_Data_Interval_ordRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_ordRecurringInterval.Do(func() {
		cache_Data_Interval_ordRecurringInterval = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_ordRecurringInterval(dictOrd_0_box)
})
	})
	return cache_Data_Interval_ordRecurringInterval
}

var cache_Data_Interval_bifunctorInterval gopurs_runtime.Value
var once_Data_Interval_bifunctorInterval sync.Once
func Get_Data_Interval_bifunctorInterval() gopurs_runtime.Value {
	once_Data_Interval_bifunctorInterval.Do(func() {
		cache_Data_Interval_bifunctorInterval = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)}))}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)}))}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)}))}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})}))}
	})
	return cache_Data_Interval_bifunctorInterval
}

var cache_Data_Interval_bifunctorRecurringInterval gopurs_runtime.Value
var once_Data_Interval_bifunctorRecurringInterval sync.Once
func Get_Data_Interval_bifunctorRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_bifunctorRecurringInterval.Do(func() {
		cache_Data_Interval_bifunctorRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_2718763394_1688994542((&Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 237113226) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V1)}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1992629780) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V1)}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2020675835) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V1)}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2281256335) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1.UnsafePtr).V0)}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, __t4}))}
})})))}
	})
	return cache_Data_Interval_bifunctorRecurringInterval
}

var cache_Data_Interval_functorInterval gopurs_runtime.Value
var once_Data_Interval_functorInterval sync.Once
func Get_Data_Interval_functorInterval() gopurs_runtime.Value {
	once_Data_Interval_functorInterval.Do(func() {
		cache_Data_Interval_functorInterval = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v1_0 gopurs_runtime.Value, v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_1.Type == 9 && v2_1.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0), gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V1)}))}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V1)}))}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V1}))}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_1.UnsafePtr).V0}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})}))}
	})
	return cache_Data_Interval_functorInterval
}

var cache_Data_Interval_extendInterval gopurs_runtime.Value
var once_Data_Interval_extendInterval sync.Once
func Get_Data_Interval_extendInterval() gopurs_runtime.Value {
	once_Data_Interval_extendInterval.Do(func() {
		cache_Data_Interval_extendInterval = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Interval_functorInterval()))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Apply(v_0, v1_1)}))}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(v_0, v1_1)}))}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1}))}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})}))}
	})
	return cache_Data_Interval_extendInterval
}

var cache_Data_Interval_functorRecurringInterval gopurs_runtime.Value
var once_Data_Interval_functorRecurringInterval sync.Once
func Get_Data_Interval_functorRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_functorRecurringInterval.Do(func() {
		cache_Data_Interval_functorRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_749713730_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 237113226) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V1)}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1992629780) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V1)}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2020675835) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V0), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V1}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2281256335) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V0}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, __t4}))}
})})))}
	})
	return cache_Data_Interval_functorRecurringInterval
}

var cache_Data_Interval_extendRecurringInterval gopurs_runtime.Value
var once_Data_Interval_extendRecurringInterval sync.Once
func Get_Data_Interval_extendRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_extendRecurringInterval.Do(func() {
		cache_Data_Interval_extendRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_844300469_3290176857((&Constructor_Control_Extend_Extend[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_749713730_2812149806(Rebox_Data_Interval_2812149806_749713730(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Interval_functorRecurringInterval()))))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=(TypeVar b)
__local_var_2_1 := gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](v_1))})
_ = __local_var_2_1
// TAST (Let): __local_var_2_0 shape=Let(Abs(Other)) bindingType=(Func [(ADT ["Data","Interval","Interval"] [(TypeVar d), (TypeVar a)])] (TypeVar b))
__local_var_2_0 := gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
})
_ = __local_var_2_0
var __t6 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 237113226) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1), gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)}))}
goto end_branch_6
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1992629780) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)}))}
goto end_branch_6
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2020675835) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V1}))}
goto end_branch_6
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2281256335) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr).V0}))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, __t6}))}
})})))}
	})
	return cache_Data_Interval_extendRecurringInterval
}

var cache_Data_Interval_traversableInterval gopurs_runtime.Value
var once_Data_Interval_traversableInterval sync.Once
func Get_Data_Interval_traversableInterval() gopurs_runtime.Value {
	once_Data_Interval_traversableInterval.Do(func() {
		cache_Data_Interval_traversableInterval = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Interval_foldableInterval()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Interval_functorInterval()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Interval_traversableInterval(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 237113226) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1992629780) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Apply(Get_Data_Interval_DurationEnd(), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2020675835) {
// TAST (Let): __local_var_5_2 shape=Other bindingType=Any
__local_var_5_2 := (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_2
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, v2_6, __local_var_5_2}))}
}), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2281256335) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0}))})
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
})}))}
	})
	return cache_Data_Interval_traversableInterval
}

var cache_Data_Interval_traversableRecurringInterval gopurs_runtime.Value
var once_Data_Interval_traversableRecurringInterval sync.Once
func Get_Data_Interval_traversableRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_traversableRecurringInterval.Do(func() {
		cache_Data_Interval_traversableRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_185619522_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_170409538_1680800814(Rebox_Data_Interval_1680800814_170409538(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Interval_foldableRecurringInterval()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_749713730_2812149806(Rebox_Data_Interval_2812149806_749713730(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Interval_functorRecurringInterval()))))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Interval_traversableRecurringInterval(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, i_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_1 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
Apply0_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_1
// TAST (Let): Functor0_5_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_2
return Call_Data_Interval_over(Functor0_1_0, gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 237113226) {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_2.V0), Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(f_2, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0)), gopurs_runtime.Apply(f_2, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1))
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1992629780) {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_2.V0), gopurs_runtime.Apply(Get_Data_Interval_DurationEnd(), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1))
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 2020675835) {
// TAST (Let): __local_var_7_3 shape=Other bindingType=Any
__local_var_7_3 := (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1
_ = __local_var_7_3
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_2.V0), gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, v2_8, __local_var_7_3}))}
}), gopurs_runtime.Apply(f_2, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0))
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 2281256335) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0}))})
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}), gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](i_3))
})
})})))}
	})
	return cache_Data_Interval_traversableRecurringInterval
}

var cache_Data_Interval_bifoldableInterval gopurs_runtime.Value
var once_Data_Interval_bifoldableInterval sync.Once
func Get_Data_Interval_bifoldableInterval() gopurs_runtime.Value {
	once_Data_Interval_bifoldableInterval.Do(func() {
		cache_Data_Interval_bifoldableInterval = gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=Other bindingType=(TypeVar m)
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Interval_bifoldableInterval()).V1), gopurs_runtime.Func2(func(m_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(f_3, a_6))
}), gopurs_runtime.Func2(func(m_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(g_4, b_6))
}), mempty_2_1)
})
}), gopurs_runtime.Func4(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value, v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v3_3.Type == 9 && v3_3.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v1_1, v2_2, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Bifoldable_bifoldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Interval_bifoldableInterval()))}, x_0)
})}))}
	})
	return cache_Data_Interval_bifoldableInterval
}

var cache_Data_Interval_bifoldableRecurringInterval gopurs_runtime.Value
var once_Data_Interval_bifoldableRecurringInterval sync.Once
func Get_Data_Interval_bifoldableRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_bifoldableRecurringInterval.Do(func() {
		cache_Data_Interval_bifoldableRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1868148450_3566843086((&Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=Other bindingType=(TypeVar m)
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Interval_bifoldableRecurringInterval()).V1), gopurs_runtime.Func2(func(m_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(f_3, a_6))
}), gopurs_runtime.Func2(func(m_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(g_4, b_6))
}), mempty_2_1)
})
}), gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, i_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 237113226) {
__t6 = gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(g_1, i_2, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1.UnsafePtr).V1)
goto end_branch_6
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1992629780) {
__t6 = gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, i_2, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1.UnsafePtr).V0), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1.UnsafePtr).V1)
goto end_branch_6
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2020675835) {
__t6 = gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, i_2, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1.UnsafePtr).V1), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2281256335) {
__t6 = gopurs_runtime.Apply2(f_0, i_2, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, i_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply5(Get_Data_Bifoldable_bifoldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Interval_bifoldableInterval()))}, f_0, g_1, i_2, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1)
})})))}
	})
	return cache_Data_Interval_bifoldableRecurringInterval
}

var cache_Data_Interval_bitraversableInterval gopurs_runtime.Value
var once_Data_Interval_bitraversableInterval sync.Once
func Get_Data_Interval_bitraversableInterval() gopurs_runtime.Value {
	once_Data_Interval_bitraversableInterval.Do(func() {
		cache_Data_Interval_bitraversableInterval = gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Interval_bifoldableInterval()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Interval_bifunctorInterval()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](Get_Data_Interval_bitraversableInterval()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_DurationEnd(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_StartDuration(), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_DurationOnly(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
})}))}
	})
	return cache_Data_Interval_bitraversableInterval
}

var cache_Data_Interval_bitraversableRecurringInterval gopurs_runtime.Value
var once_Data_Interval_bitraversableRecurringInterval sync.Once
func Get_Data_Interval_bitraversableRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_bitraversableRecurringInterval.Do(func() {
		cache_Data_Interval_bitraversableRecurringInterval = gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1937664898_3561684974((&Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1868148450_3566843086(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Interval_bifoldableRecurringInterval())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_2718763394_1688994542(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Interval_bifunctorRecurringInterval())))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Interval_bitraversableRecurringInterval()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func3(func(l_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value, i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_1 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f)])
Apply0_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_1
// TAST (Let): Functor0_6_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_2
return Call_Data_Interval_over(Functor0_1_0, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 237113226) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_2.V0), Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(r_3, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)), gopurs_runtime.Apply(r_3, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 1992629780) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_2.V0), Get_Data_Interval_DurationEnd(), gopurs_runtime.Apply(l_2, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)), gopurs_runtime.Apply(r_3, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2020675835) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_2.V0), Get_Data_Interval_StartDuration(), gopurs_runtime.Apply(r_3, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)), gopurs_runtime.Apply(l_2, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2281256335) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_2.V0), Get_Data_Interval_DurationOnly(), gopurs_runtime.Apply(l_2, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](i_4))
})
})})))}
	})
	return cache_Data_Interval_bitraversableRecurringInterval
}

type Constructor_Data_Interval_StartEnd[T_d any, T_a any] struct {
	Rc uint32
	V0 T_a
	V1 T_a
}


type Constructor_Data_Interval_DurationEnd[T_d any, T_a any] struct {
	Rc uint32
	V0 T_d
	V1 T_a
}


type Constructor_Data_Interval_StartDuration[T_d any, T_a any] struct {
	Rc uint32
	V0 T_a
	V1 T_d
}


type Constructor_Data_Interval_DurationOnly[T_d any, T_a any] struct {
	Rc uint32
	V0 T_d
}


type Constructor_Data_Interval_RecurringInterval[T_d any, T_a any] struct {
	Rc uint32
	V0 *Constructor_Data_Maybe_Just[int64]
	V1 gopurs_runtime.Value
}


func Call_Data_Interval_showInterval(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 string
{
if (v_2.Type == 9 && v_2.IntVal == 237113226) {
__t0 = (((("(StartEnd ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1992629780) {
__t0 = (((("(DurationEnd ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2020675835) {
__t0 = (((("(StartDuration ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2281256335) {
__t0 = (("(DurationOnly ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
__t0 = func() string { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str(__t0)
})}))}
}

func Call_Data_Interval_showRecurringInterval(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
// TAST (Let): showInterval2_2_0 shape=LitRecord bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","Interval","Interval"] [(TypeVar d), (TypeVar a)])])
showInterval2_2_0 := (&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 string
{
if (v_2.Type == 9 && v_2.IntVal == 237113226) {
__t1 = (((("(StartEnd ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1992629780) {
__t1 = (((("(DurationEnd ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2020675835) {
__t1 = (((("(StartDuration ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2281256335) {
__t1 = (("(DurationOnly ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_1
} else {

}
}
{
__t1 = func() string { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Str(__t1)
})})
_ = showInterval2_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_951668290_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(RecurringInterval ") + (gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]]](Get_Data_Interval_showMaybe()).V0), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1170268447_3094389156((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0))}).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(showInterval2_2_0.V0), (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1).StrVal())) + (")"))
})})))}
}

func Call_Data_Interval_over(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Apply(Get_Data_Interval_RecurringInterval(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1170268447_3094389156((v_2).V0))}), gopurs_runtime.Apply(f_1, (v_2).V1))
}

func Call_Data_Interval_interval(v_0_loop *Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Interval_eqInterval(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 bool
{
if (x_2.Type == 9 && x_2.IntVal == 237113226) {
__t0 = ((y_3.Type == 9 && y_3.IntVal == 237113226)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_0
} else {

}
}
{
if (x_2.Type == 9 && x_2.IntVal == 1992629780) {
__t0 = ((y_3.Type == 9 && y_3.IntVal == 1992629780)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_0
} else {

}
}
{
if (x_2.Type == 9 && x_2.IntVal == 2020675835) {
__t0 = ((y_3.Type == 9 && y_3.IntVal == 2020675835)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_0
} else {

}
}
{
__t0 = ((x_2.Type == 9 && x_2.IntVal == 2281256335)) && (((y_3.Type == 9 && y_3.IntVal == 2281256335)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_0:
return gopurs_runtime.Bool(__t0)
})}))}
}

func Call_Data_Interval_eqRecurringInterval(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
// TAST (Let): eqInterval2_2_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Interval","Interval"] [(TypeVar d), (TypeVar a)])])
eqInterval2_2_0 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (x_2.Type == 9 && x_2.IntVal == 237113226) {
__t1 = ((y_3.Type == 9 && y_3.IntVal == 237113226)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_1
} else {

}
}
{
if (x_2.Type == 9 && x_2.IntVal == 1992629780) {
__t1 = ((y_3.Type == 9 && y_3.IntVal == 1992629780)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_1
} else {

}
}
{
if (x_2.Type == 9 && x_2.IntVal == 2020675835) {
__t1 = ((y_3.Type == 9 && y_3.IntVal == 2020675835)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_1
} else {

}
}
{
__t1 = ((x_2.Type == 9 && x_2.IntVal == 2281256335)) && (((y_3.Type == 9 && y_3.IntVal == 2281256335)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})})
_ = eqInterval2_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1142559074_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 bool
{
var __t_tag_5 *Constructor_Data_Maybe_Just[int64] = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0
if (__t_tag_5 == nil) {
var __t_tag_6 *Constructor_Data_Maybe_Just[int64] = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0
__t7 = (__t_tag_6 == nil)
goto end_branch_7
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[int64] = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0
var __t_and_4 bool = false
if (__t_tag_2 != nil) {

var __t_tag_3 *Constructor_Data_Maybe_Just[int64] = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0
__t_and_4 = ((__t_tag_3 != nil)) && ((((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0).V0) == (gopurs_runtime.Int(((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).V0).IntVal))
}
__t7 = __t_and_4
}
end_branch_7:
return gopurs_runtime.Bool((__t7) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqInterval2_2_0.V0), (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V1).IntVal) != (0)))
})})))}
}

func Call_Data_Interval_ordInterval(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqInterval1__193435443_1_0 shape=Let(Abs(LitRecord)) bindingType=Any
eqInterval1__193435443_1_0 := gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 bool
{
if (x_3.Type == 9 && x_3.IntVal == 237113226) {
__t2 = ((y_4.Type == 9 && y_4.IntVal == 237113226)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_2
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 1992629780) {
__t2 = ((y_4.Type == 9 && y_4.IntVal == 1992629780)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_2
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 2020675835) {
__t2 = ((y_4.Type == 9 && y_4.IntVal == 2020675835)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_2
} else {

}
}
{
__t2 = ((x_3.Type == 9 && x_3.IntVal == 2281256335)) && (((y_4.Type == 9 && y_4.IntVal == 2281256335)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
})}))}
})
_ = eqInterval1__193435443_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqInterval2_3_3 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Interval","Interval"] [(TypeVar d), (TypeVar a)])])
eqInterval2_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqInterval1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})))
_ = eqInterval2_3_3
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqInterval2_3_3)}
}), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 uint32
{
if (x_4.Type == 9 && x_4.IntVal == 237113226) {
var __t6 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 237113226) {
// TAST (Let): v_6_4 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_6_4 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
_ = v_6_4
var __t5 uint32
{
if (v_6_4 == 1527465420) {
__t5 = 1527465420
goto end_branch_5
} else {

}
}
{
if (v_6_4 == 380165415) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
__t5 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = 1527465420
}
end_branch_6:
__t13 = __t6
goto end_branch_13
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 237113226) {
__t13 = 380165415
goto end_branch_13
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 1992629780) {
var __t9 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 1992629780) {
// TAST (Let): v_6_7 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_6_7 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
_ = v_6_7
var __t8 uint32
{
if (v_6_7 == 1527465420) {
__t8 = 1527465420
goto end_branch_8
} else {

}
}
{
if (v_6_7 == 380165415) {
__t8 = 380165415
goto end_branch_8
} else {

}
}
{
__t8 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = 1527465420
}
end_branch_9:
__t13 = __t9
goto end_branch_13
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 1992629780) {
__t13 = 380165415
goto end_branch_13
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 2020675835) {
var __t12 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 2020675835) {
// TAST (Let): v_6_10 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_6_10 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
_ = v_6_10
var __t11 uint32
{
if (v_6_10 == 1527465420) {
__t11 = 1527465420
goto end_branch_11
} else {

}
}
{
if (v_6_10 == 380165415) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
__t11 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
}
end_branch_11:
__t12 = __t11
goto end_branch_12
} else {

}
}
{
__t12 = 1527465420
}
end_branch_12:
__t13 = __t12
goto end_branch_13
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 2020675835) {
__t13 = 380165415
goto end_branch_13
} else {

}
}
{
if ((x_4.Type == 9 && x_4.IntVal == 2281256335)) && ((y_5.Type == 9 && y_5.IntVal == 2281256335)) {
__t13 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
goto end_branch_13
} else {

}
}
{
__t13 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t13), UnsafePtr: nil}
})}))}
})
}

func Call_Data_Interval_ordRecurringInterval(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): eqInterval1__193435443_1_1 shape=Let(Abs(LitRecord)) bindingType=Any
eqInterval1__193435443_1_1 := gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
if (x_3.Type == 9 && x_3.IntVal == 237113226) {
__t3 = ((y_4.Type == 9 && y_4.IntVal == 237113226)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_3
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 1992629780) {
__t3 = ((y_4.Type == 9 && y_4.IntVal == 1992629780)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "eq"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_3
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 2020675835) {
__t3 = ((y_4.Type == 9 && y_4.IntVal == 2020675835)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "eq"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_3
} else {

}
}
{
__t3 = ((x_3.Type == 9 && x_3.IntVal == 2281256335)) && (((y_4.Type == 9 && y_4.IntVal == 2281256335)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "eq"), (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
})}))}
})
_ = eqInterval1__193435443_1_1
// TAST (Let): ordInterval1__193435443_1_0 shape=Let(Abs(Let(LitRecord))) bindingType=Any
ordInterval1__193435443_1_0 := gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqInterval2_3_4 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Interval","Interval"] [(TypeVar d), (TypeVar a)])])
eqInterval2_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqInterval1__193435443_1_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})))
_ = eqInterval2_3_4
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqInterval2_3_4)}
}), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 uint32
{
if (x_4.Type == 9 && x_4.IntVal == 237113226) {
var __t7 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 237113226) {
// TAST (Let): v_6_5 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_6_5 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
_ = v_6_5
var __t6 uint32
{
if (v_6_5 == 1527465420) {
__t6 = 1527465420
goto end_branch_6
} else {

}
}
{
if (v_6_5 == 380165415) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
__t6 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = 1527465420
}
end_branch_7:
__t14 = __t7
goto end_branch_14
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 237113226) {
__t14 = 380165415
goto end_branch_14
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 1992629780) {
var __t10 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 1992629780) {
// TAST (Let): v_6_8 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_6_8 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
_ = v_6_8
var __t9 uint32
{
if (v_6_8 == 1527465420) {
__t9 = 1527465420
goto end_branch_9
} else {

}
}
{
if (v_6_8 == 380165415) {
__t9 = 380165415
goto end_branch_9
} else {

}
}
{
__t9 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
}
end_branch_9:
__t10 = __t9
goto end_branch_10
} else {

}
}
{
__t10 = 1527465420
}
end_branch_10:
__t14 = __t10
goto end_branch_14
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 1992629780) {
__t14 = 380165415
goto end_branch_14
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 2020675835) {
var __t13 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 2020675835) {
// TAST (Let): v_6_11 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_6_11 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
_ = v_6_11
var __t12 uint32
{
if (v_6_11 == 1527465420) {
__t12 = 1527465420
goto end_branch_12
} else {

}
}
{
if (v_6_11 == 380165415) {
__t12 = 380165415
goto end_branch_12
} else {

}
}
{
__t12 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
}
end_branch_12:
__t13 = __t12
goto end_branch_13
} else {

}
}
{
__t13 = 1527465420
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 2020675835) {
__t14 = 380165415
goto end_branch_14
} else {

}
}
{
if ((x_4.Type == 9 && x_4.IntVal == 2281256335)) && ((y_5.Type == 9 && y_5.IntVal == 2281256335)) {
__t14 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
goto end_branch_14
} else {

}
}
{
__t14 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t14), UnsafePtr: nil}
})}))}
})
_ = ordInterval1__193435443_1_0
// TAST (Let): __local_var_2_16 shape=App(Other) bindingType=Any
__local_var_2_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_16
// TAST (Let): eqRecurringInterval1__193435443_2_15 shape=Let(Abs(Let(LitRecord))) bindingType=Any
eqRecurringInterval1__193435443_2_15 := gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqInterval2_4_17 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Interval","Interval"] [(TypeVar d), (TypeVar a)])])
eqInterval2_4_17 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 bool
{
if (x_4.Type == 9 && x_4.IntVal == 237113226) {
__t18 = ((y_5.Type == 9 && y_5.IntVal == 237113226)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_18
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 1992629780) {
__t18 = ((y_5.Type == 9 && y_5.IntVal == 1992629780)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_16, "eq"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_18
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 2020675835) {
__t18 = ((y_5.Type == 9 && y_5.IntVal == 2020675835)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_16, "eq"), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal) != (0)))
goto end_branch_18
} else {

}
}
{
__t18 = ((x_4.Type == 9 && x_4.IntVal == 2281256335)) && (((y_5.Type == 9 && y_5.IntVal == 2281256335)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_16, "eq"), (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_18:
return gopurs_runtime.Bool(__t18)
})})
_ = eqInterval2_4_17
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1142559074_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 bool
{
var __t_tag_22 *Constructor_Data_Maybe_Just[int64] = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0
if (__t_tag_22 == nil) {
var __t_tag_23 *Constructor_Data_Maybe_Just[int64] = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0
__t24 = (__t_tag_23 == nil)
goto end_branch_24
} else {

}
}
{
var __t_tag_19 *Constructor_Data_Maybe_Just[int64] = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0
var __t_and_21 bool = false
if (__t_tag_19 != nil) {

var __t_tag_20 *Constructor_Data_Maybe_Just[int64] = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0
__t_and_21 = ((__t_tag_20 != nil)) && ((((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0).V0) == (gopurs_runtime.Int(((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0).V0).IntVal))
}
__t24 = __t_and_21
}
end_branch_24:
return gopurs_runtime.Bool((__t24) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqInterval2_4_17.V0), (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V1, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V1).IntVal) != (0)))
})})))}
})
_ = eqRecurringInterval1__193435443_2_15
return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordInterval2_4_25 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Interval","Interval"] [(TypeVar d), (TypeVar a)])])
ordInterval2_4_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordInterval1__193435443_1_0, dictOrd1_3))
_ = ordInterval2_4_25
// TAST (Let): eqRecurringInterval2_5_26 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Interval","RecurringInterval"] [(TypeVar d), (TypeVar a)])])
eqRecurringInterval2_5_26 := Rebox_Data_Interval_3790796878_1142559074(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqRecurringInterval1__193435443_2_15, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{}))))
_ = eqRecurringInterval2_5_26
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_2740699138_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_1142559074_3790796878(eqRecurringInterval2_5_26))}
}), gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, y_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 uint32
{
var __t_tag_28 *Constructor_Data_Maybe_Just[int64] = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0
if (__t_tag_28 == nil) {
var __t30 uint32
{
var __t_tag_29 *Constructor_Data_Maybe_Just[int64] = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0
if (__t_tag_29 == nil) {
__t30 = 902936544
goto end_branch_30
} else {

}
}
{
__t30 = 1527465420
}
end_branch_30:
__t35 = __t30
goto end_branch_35
} else {

}
}
{
var __t_tag_31 *Constructor_Data_Maybe_Just[int64] = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0
if (__t_tag_31 == nil) {
__t35 = 380165415
goto end_branch_35
} else {

}
}
{
var __t_tag_32 *Constructor_Data_Maybe_Just[int64] = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0
var __t_and_34 bool = false
if (__t_tag_32 != nil) {

var __t_tag_33 *Constructor_Data_Maybe_Just[int64] = (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0
__t_and_34 = (__t_tag_33 != nil)
}
if __t_and_34 {
__t35 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[int64]](Get_Data_Ord_ordInt()).V1), gopurs_runtime.Int(((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0).V0), gopurs_runtime.Int(((*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0).V0)).IntVal)
goto end_branch_35
} else {

}
}
{
__t35 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_35:
// TAST (Let): v_8_27 shape=Branch(Branch(Other, def=Other), Other, App(Other), def=Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_8_27 := __t35
_ = v_8_27
var __t36 uint32
{
if (v_8_27 == 1527465420) {
__t36 = 1527465420
goto end_branch_36
} else {

}
}
{
if (v_8_27 == 380165415) {
__t36 = 380165415
goto end_branch_36
} else {

}
}
{
__t36 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordInterval2_4_25.V1), (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V1, (*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V1).IntVal)
}
end_branch_36:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t36), UnsafePtr: nil}
})})))}
})
}

func Rebox_Data_Interval_1142559074_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Interval_1680800814_170409538(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Interval_170409538_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Interval_185619522_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Interval_1868148450_3566843086(in *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Interval_1937664898_3561684974(in *Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Interval_2718763394_1688994542(in *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_2740699138_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Interval_2812149806_749713730(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_2818770644_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[int64]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_3092443796_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[int64]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Interval_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Interval_3508461103_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_3790796878_1142559074(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_3960201844_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[int64]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_749713730_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_844300469_3290176857(in *Constructor_Control_Extend_Extend[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Extend_Extend[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Extend_Extend[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Interval_951668290_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Interval_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}



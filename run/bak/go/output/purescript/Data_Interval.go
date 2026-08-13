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
		cache_Data_Interval_showMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 string
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = (("(Just ") + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), (*Constructor_Data_Maybe_Just)(v_0.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = "Nothing"
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_0:
return gopurs_runtime.Str(__t0)
})})}
	})
	return cache_Data_Interval_showMaybe
}

var cache_Data_Interval_eqMaybe gopurs_runtime.Value
var once_Data_Interval_eqMaybe sync.Once
func Get_Data_Interval_eqMaybe() gopurs_runtime.Value {
	once_Data_Interval_eqMaybe.Do(func() {
		cache_Data_Interval_eqMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (x_0.Type == 9 && x_0.IntVal == 930809136 && x_0.UnsafePtr == nil) {
var __t0 bool
{
if (y_1.Type == 9 && y_1.IntVal == 930809136 && y_1.UnsafePtr == nil) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 930809136 && x_0.UnsafePtr != nil)) && ((y_1.Type == 9 && y_1.IntVal == 930809136 && y_1.UnsafePtr != nil)) {
__t1 = ((*Constructor_Data_Maybe_Just)(x_0.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Maybe_Just)(y_1.UnsafePtr).V0.IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})
})})}
	})
	return cache_Data_Interval_eqMaybe
}

var cache_Data_Interval_ordMaybe gopurs_runtime.Value
var once_Data_Interval_ordMaybe sync.Once
func Get_Data_Interval_ordMaybe() gopurs_runtime.Value {
	once_Data_Interval_ordMaybe.Do(func() {
		cache_Data_Interval_ordMaybe = func() gopurs_runtime.Value {
// TAST (Let): eqMaybe1_0_0 -> gopurs_runtime.Value
eqMaybe1_0_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
}))
_ = eqMaybe1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_0_0
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
__t4 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Ord_ordInt(), "compare"), (*Constructor_Data_Maybe_Just)(x_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just)(y_2.UnsafePtr).V0).IntVal)
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
}))))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, value0, value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, value0, value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, value0, value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, value0})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_RecurringInterval{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](value0), value1})}
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
return Call_Data_Interval_over(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval](v_2_box))
})
	})
	return cache_Data_Interval_over
}

var cache_Data_Interval_interval gopurs_runtime.Value
var once_Data_Interval_interval sync.Once
func Get_Data_Interval_interval() gopurs_runtime.Value {
	once_Data_Interval_interval.Do(func() {
		cache_Data_Interval_interval = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_interval(gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval](v_0_box))
})
	})
	return cache_Data_Interval_interval
}

var cache_Data_Interval_foldableInterval gopurs_runtime.Value
var once_Data_Interval_foldableInterval sync.Once
func Get_Data_Interval_foldableInterval() gopurs_runtime.Value {
	once_Data_Interval_foldableInterval.Do(func() {
		cache_Data_Interval_foldableInterval = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Interval_foldableInterval(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), acc_4, gopurs_runtime.Apply(f_3, x_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_DurationEnd)(v2_2.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_StartDuration)(v2_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = v1_1
}
end_branch_2:
return __t2
})
})
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Foldable_foldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_Interval_foldableInterval()))}, x_0)
}))
	})
	return cache_Data_Interval_foldableInterval
}

var cache_Data_Interval_foldableRecurringInterval gopurs_runtime.Value
var once_Data_Interval_foldableRecurringInterval sync.Once
func Get_Data_Interval_foldableRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_foldableRecurringInterval.Do(func() {
		cache_Data_Interval_foldableRecurringInterval = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Interval_foldableRecurringInterval(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), acc_4, gopurs_runtime.Apply(f_3, x_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 237113226) {
__t5 = gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply2(f_0, i_1, (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1.UnsafePtr).V1)
goto end_branch_5
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1992629780) {
__t5 = gopurs_runtime.Apply2(f_0, i_1, (*Constructor_Data_Interval_DurationEnd)((*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1.UnsafePtr).V1)
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2020675835) {
__t5 = gopurs_runtime.Apply2(f_0, i_1, (*Constructor_Data_Interval_StartDuration)((*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = i_1
}
end_branch_5:
return __t5
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Data_Foldable_foldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_Interval_foldableInterval()))}, f_0, i_1, (*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1)
})
})
}))
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
		cache_Data_Interval_bifunctorInterval = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_DurationEnd)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_DurationEnd)(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartDuration)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_StartDuration)(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_DurationOnly)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})
}))
	})
	return cache_Data_Interval_bifunctorInterval
}

var cache_Data_Interval_bifunctorRecurringInterval gopurs_runtime.Value
var once_Data_Interval_bifunctorRecurringInterval sync.Once
func Get_Data_Interval_bifunctorRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_bifunctorRecurringInterval.Do(func() {
		cache_Data_Interval_bifunctorRecurringInterval = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 237113226) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V1)})}
goto end_branch_4
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1992629780) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_DurationEnd)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_DurationEnd)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V1)})}
goto end_branch_4
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2020675835) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_StartDuration)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartDuration)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V1)})}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2281256335) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_DurationOnly)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V0)})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_RecurringInterval{1, (*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V0, __t4})}
})
})
}))
	})
	return cache_Data_Interval_bifunctorRecurringInterval
}

var cache_Data_Interval_functorInterval gopurs_runtime.Value
var once_Data_Interval_functorInterval sync.Once
func Get_Data_Interval_functorInterval() gopurs_runtime.Value {
	once_Data_Interval_functorInterval.Do(func() {
		cache_Data_Interval_functorInterval = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_1.Type == 9 && v2_1.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartEnd)(v2_1.UnsafePtr).V0), gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartEnd)(v2_1.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, (*Constructor_Data_Interval_DurationEnd)(v2_1.UnsafePtr).V0, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_DurationEnd)(v2_1.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartDuration)(v2_1.UnsafePtr).V0), (*Constructor_Data_Interval_StartDuration)(v2_1.UnsafePtr).V1})}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, (*Constructor_Data_Interval_DurationOnly)(v2_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_Data_Interval_functorInterval
}

var cache_Data_Interval_extendInterval gopurs_runtime.Value
var once_Data_Interval_extendInterval sync.Once
func Get_Data_Interval_extendInterval() gopurs_runtime.Value {
	once_Data_Interval_extendInterval.Do(func() {
		cache_Data_Interval_extendInterval = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_functorInterval()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Apply(v_0, v1_1)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, (*Constructor_Data_Interval_DurationEnd)(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(v_0, v1_1)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, gopurs_runtime.Apply(v_0, v1_1), (*Constructor_Data_Interval_StartDuration)(v1_1.UnsafePtr).V1})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, (*Constructor_Data_Interval_DurationOnly)(v1_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_Data_Interval_extendInterval
}

var cache_Data_Interval_functorRecurringInterval gopurs_runtime.Value
var once_Data_Interval_functorRecurringInterval sync.Once
func Get_Data_Interval_functorRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_functorRecurringInterval.Do(func() {
		cache_Data_Interval_functorRecurringInterval = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 237113226) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V1)})}
goto end_branch_4
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1992629780) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, (*Constructor_Data_Interval_DurationEnd)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_DurationEnd)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V1)})}
goto end_branch_4
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2020675835) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartDuration)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V0), (*Constructor_Data_Interval_StartDuration)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V1})}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2281256335) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, (*Constructor_Data_Interval_DurationOnly)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V0})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_RecurringInterval{1, (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V0, __t4})}
})
}))
	})
	return cache_Data_Interval_functorRecurringInterval
}

var cache_Data_Interval_extendRecurringInterval gopurs_runtime.Value
var once_Data_Interval_extendRecurringInterval sync.Once
func Get_Data_Interval_extendRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_extendRecurringInterval.Do(func() {
		cache_Data_Interval_extendRecurringInterval = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_functorRecurringInterval()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval](v_1))})
_ = __local_var_2_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
})
_ = __local_var_2_0
var __t6 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 237113226) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1), gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1)})}
goto end_branch_6
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1992629780) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, (*Constructor_Data_Interval_DurationEnd)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1)})}
goto end_branch_6
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2020675835) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1), (*Constructor_Data_Interval_StartDuration)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V1})}
goto end_branch_6
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2281256335) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, (*Constructor_Data_Interval_DurationOnly)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V0})}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_RecurringInterval{1, (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V0, __t6})}
})
}))
	})
	return cache_Data_Interval_extendRecurringInterval
}

var cache_Data_Interval_traversableInterval gopurs_runtime.Value
var once_Data_Interval_traversableInterval sync.Once
func Get_Data_Interval_traversableInterval() gopurs_runtime.Value {
	once_Data_Interval_traversableInterval.Do(func() {
		cache_Data_Interval_traversableInterval = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_foldableInterval()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_functorInterval()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Interval_traversableInterval(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 237113226) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartEnd)(v1_4.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartEnd)(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1992629780) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Apply(Get_Data_Interval_DurationEnd(), (*Constructor_Data_Interval_DurationEnd)(v1_4.UnsafePtr).V0), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationEnd)(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2020675835) {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := (*Constructor_Data_Interval_StartDuration)(v1_4.UnsafePtr).V1
_ = __local_var_5_2
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, v2_6, __local_var_5_2})}
}), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartDuration)(v1_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2281256335) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, (*Constructor_Data_Interval_DurationOnly)(v1_4.UnsafePtr).V0})})
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
})
}))
	})
	return cache_Data_Interval_traversableInterval
}

var cache_Data_Interval_traversableRecurringInterval gopurs_runtime.Value
var once_Data_Interval_traversableRecurringInterval sync.Once
func Get_Data_Interval_traversableRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_traversableRecurringInterval.Do(func() {
		cache_Data_Interval_traversableRecurringInterval = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_foldableRecurringInterval()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_functorRecurringInterval()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Interval_traversableRecurringInterval(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(Get_Data_Interval_RecurringInterval(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Interval_RecurringInterval)(i_3.UnsafePtr).V0)}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Interval_traversableInterval(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, f_2, (*Constructor_Data_Interval_RecurringInterval)(i_3.UnsafePtr).V1))
})
})
}))
	})
	return cache_Data_Interval_traversableRecurringInterval
}

var cache_Data_Interval_bifoldableInterval gopurs_runtime.Value
var once_Data_Interval_bifoldableInterval sync.Once
func Get_Data_Interval_bifoldableInterval() gopurs_runtime.Value {
	once_Data_Interval_bifoldableInterval.Do(func() {
		cache_Data_Interval_bifoldableInterval = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Interval_bifoldableInterval(), "bifoldl"), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(f_3, a_6))
})
}), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(g_4, b_6))
})
}), mempty_2_1)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v3_3.Type == 9 && v3_3.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v1_1, v2_2, (*Constructor_Data_Interval_StartEnd)(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd)(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_DurationEnd)(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_DurationEnd)(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_StartDuration)(v3_3.UnsafePtr).V1), (*Constructor_Data_Interval_StartDuration)(v3_3.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_DurationOnly)(v3_3.UnsafePtr).V0)
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
})
})
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupEndo1_4_3 -> gopurs_runtime.Value
semigroupEndo1_4_3 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Apply(v1_5, x_6))
})
})
}))
_ = semigroupEndo1_4_3
return gopurs_runtime.Apply5(gopurs_runtime.RecordGet(Get_Data_Interval_bifoldableInterval(), "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_4_3
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
}))))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(x_0, x_4)
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, x_4)
}), p_3, z_2)
})
})
})
}))
	})
	return cache_Data_Interval_bifoldableInterval
}

var cache_Data_Interval_bifoldableRecurringInterval gopurs_runtime.Value
var once_Data_Interval_bifoldableRecurringInterval sync.Once
func Get_Data_Interval_bifoldableRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_bifoldableRecurringInterval.Do(func() {
		cache_Data_Interval_bifoldableRecurringInterval = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Interval_bifoldableRecurringInterval(), "bifoldl"), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(f_3, a_6))
})
}), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(g_4, b_6))
})
}), mempty_2_1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(Get_Data_Interval_bifoldableInterval(), "bifoldl"), f_0, g_1, i_2, (*Constructor_Data_Interval_RecurringInterval)(x_3.UnsafePtr).V1)
})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(Get_Data_Interval_bifoldableInterval(), "bifoldr"), f_0, g_1, i_2, (*Constructor_Data_Interval_RecurringInterval)(x_3.UnsafePtr).V1)
})
})
})
}))
	})
	return cache_Data_Interval_bifoldableRecurringInterval
}

var cache_Data_Interval_bitraversableInterval gopurs_runtime.Value
var once_Data_Interval_bitraversableInterval sync.Once
func Get_Data_Interval_bitraversableInterval() gopurs_runtime.Value {
	once_Data_Interval_bitraversableInterval.Do(func() {
		cache_Data_Interval_bitraversableInterval = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_bifoldableInterval()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_bifunctorInterval()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Interval_bitraversableInterval(), "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartEnd)(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartEnd)(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_DurationEnd(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationEnd)(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_DurationEnd)(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_StartDuration(), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartDuration)(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartDuration)(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_DurationOnly(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationOnly)(v2_5.UnsafePtr).V0))
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
})
})
}))
	})
	return cache_Data_Interval_bitraversableInterval
}

var cache_Data_Interval_bitraversableRecurringInterval gopurs_runtime.Value
var once_Data_Interval_bitraversableRecurringInterval sync.Once
func Get_Data_Interval_bitraversableRecurringInterval() gopurs_runtime.Value {
	once_Data_Interval_bitraversableRecurringInterval.Do(func() {
		cache_Data_Interval_bitraversableRecurringInterval = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_bifoldableRecurringInterval()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_bifunctorRecurringInterval()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Interval_bitraversableRecurringInterval(), "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(Get_Data_Interval_RecurringInterval(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Interval_RecurringInterval)(i_4.UnsafePtr).V0)}), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(Get_Data_Interval_bitraversableInterval(), "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, l_2, r_3, (*Constructor_Data_Interval_RecurringInterval)(i_4.UnsafePtr).V1))
})
})
})
}))
	})
	return cache_Data_Interval_bitraversableRecurringInterval
}

var cache_Data_Interval_bifoldableInterval__2998510362 gopurs_runtime.Value
var once_Data_Interval_bifoldableInterval__2998510362 sync.Once
func Get_Data_Interval_bifoldableInterval__2998510362() gopurs_runtime.Value {
	once_Data_Interval_bifoldableInterval__2998510362.Do(func() {
		cache_Data_Interval_bifoldableInterval__2998510362 = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Interval_bifoldableInterval(), "bifoldl"), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(f_3, a_6))
})
}), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(g_4, b_6))
})
}), mempty_2_1)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v3_3.Type == 9 && v3_3.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v1_1, v2_2, (*Constructor_Data_Interval_StartEnd)(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd)(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_DurationEnd)(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_DurationEnd)(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_StartDuration)(v3_3.UnsafePtr).V1), (*Constructor_Data_Interval_StartDuration)(v3_3.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_DurationOnly)(v3_3.UnsafePtr).V0)
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
})
})
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupEndo1_4_3 -> gopurs_runtime.Value
semigroupEndo1_4_3 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Apply(v1_5, x_6))
})
})
}))
_ = semigroupEndo1_4_3
return gopurs_runtime.Apply5(gopurs_runtime.RecordGet(Get_Data_Interval_bifoldableInterval(), "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_4_3
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
}))))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(x_0, x_4)
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, x_4)
}), p_3, z_2)
})
})
})
}))
	})
	return cache_Data_Interval_bifoldableInterval__2998510362
}

var cache_Data_Interval_bifoldableInterval__740659151 gopurs_runtime.Value
var once_Data_Interval_bifoldableInterval__740659151 sync.Once
func Get_Data_Interval_bifoldableInterval__740659151() gopurs_runtime.Value {
	once_Data_Interval_bifoldableInterval__740659151.Do(func() {
		cache_Data_Interval_bifoldableInterval__740659151 = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Interval_bifoldableInterval(), "bifoldl"), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(f_3, a_6))
})
}), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(g_4, b_6))
})
}), mempty_2_1)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v3_3.Type == 9 && v3_3.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v1_1, v2_2, (*Constructor_Data_Interval_StartEnd)(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd)(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_DurationEnd)(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_DurationEnd)(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_StartDuration)(v3_3.UnsafePtr).V1), (*Constructor_Data_Interval_StartDuration)(v3_3.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_DurationOnly)(v3_3.UnsafePtr).V0)
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
})
})
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupEndo1_4_3 -> gopurs_runtime.Value
semigroupEndo1_4_3 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Apply(v1_5, x_6))
})
})
}))
_ = semigroupEndo1_4_3
return gopurs_runtime.Apply5(gopurs_runtime.RecordGet(Get_Data_Interval_bifoldableInterval(), "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_4_3
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
}))))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(x_0, x_4)
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, x_4)
}), p_3, z_2)
})
})
})
}))
	})
	return cache_Data_Interval_bifoldableInterval__740659151
}

var cache_Data_Interval_bifoldableRecurringInterval__4077340974 gopurs_runtime.Value
var once_Data_Interval_bifoldableRecurringInterval__4077340974 sync.Once
func Get_Data_Interval_bifoldableRecurringInterval__4077340974() gopurs_runtime.Value {
	once_Data_Interval_bifoldableRecurringInterval__4077340974.Do(func() {
		cache_Data_Interval_bifoldableRecurringInterval__4077340974 = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Interval_bifoldableRecurringInterval(), "bifoldl"), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(f_3, a_6))
})
}), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), m_5, gopurs_runtime.Apply(g_4, b_6))
})
}), mempty_2_1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(Get_Data_Interval_bifoldableInterval(), "bifoldl"), f_0, g_1, i_2, (*Constructor_Data_Interval_RecurringInterval)(x_3.UnsafePtr).V1)
})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(Get_Data_Interval_bifoldableInterval(), "bifoldr"), f_0, g_1, i_2, (*Constructor_Data_Interval_RecurringInterval)(x_3.UnsafePtr).V1)
})
})
})
}))
	})
	return cache_Data_Interval_bifoldableRecurringInterval__4077340974
}

var cache_Data_Interval_bifunctorInterval__3636391546 gopurs_runtime.Value
var once_Data_Interval_bifunctorInterval__3636391546 sync.Once
func Get_Data_Interval_bifunctorInterval__3636391546() gopurs_runtime.Value {
	once_Data_Interval_bifunctorInterval__3636391546.Do(func() {
		cache_Data_Interval_bifunctorInterval__3636391546 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_DurationEnd)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_DurationEnd)(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartDuration)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_StartDuration)(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_DurationOnly)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})
}))
	})
	return cache_Data_Interval_bifunctorInterval__3636391546
}

var cache_Data_Interval_bifunctorInterval__3665086044 gopurs_runtime.Value
var once_Data_Interval_bifunctorInterval__3665086044 sync.Once
func Get_Data_Interval_bifunctorInterval__3665086044() gopurs_runtime.Value {
	once_Data_Interval_bifunctorInterval__3665086044.Do(func() {
		cache_Data_Interval_bifunctorInterval__3665086044 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_DurationEnd)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_DurationEnd)(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, gopurs_runtime.Apply(v1_1, (*Constructor_Data_Interval_StartDuration)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_StartDuration)(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, gopurs_runtime.Apply(v_0, (*Constructor_Data_Interval_DurationOnly)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})
}))
	})
	return cache_Data_Interval_bifunctorInterval__3665086044
}

var cache_Data_Interval_bifunctorRecurringInterval__261139004 gopurs_runtime.Value
var once_Data_Interval_bifunctorRecurringInterval__261139004 sync.Once
func Get_Data_Interval_bifunctorRecurringInterval__261139004() gopurs_runtime.Value {
	once_Data_Interval_bifunctorRecurringInterval__261139004.Do(func() {
		cache_Data_Interval_bifunctorRecurringInterval__261139004 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 237113226) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V1)})}
goto end_branch_4
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1992629780) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_DurationEnd)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_DurationEnd)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V1)})}
goto end_branch_4
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2020675835) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, gopurs_runtime.Apply(g_1, (*Constructor_Data_Interval_StartDuration)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartDuration)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V1)})}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2281256335) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_DurationOnly)((*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V1.UnsafePtr).V0)})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_RecurringInterval{1, (*Constructor_Data_Interval_RecurringInterval)(v_2.UnsafePtr).V0, __t4})}
})
})
}))
	})
	return cache_Data_Interval_bifunctorRecurringInterval__261139004
}

var cache_Data_Interval_bitraversableInterval__1397501818 gopurs_runtime.Value
var once_Data_Interval_bitraversableInterval__1397501818 sync.Once
func Get_Data_Interval_bitraversableInterval__1397501818() gopurs_runtime.Value {
	once_Data_Interval_bitraversableInterval__1397501818.Do(func() {
		cache_Data_Interval_bitraversableInterval__1397501818 = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_bifoldableInterval()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_bifunctorInterval()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Interval_bitraversableInterval(), "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartEnd)(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartEnd)(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_DurationEnd(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationEnd)(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_DurationEnd)(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_StartDuration(), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartDuration)(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartDuration)(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_DurationOnly(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationOnly)(v2_5.UnsafePtr).V0))
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
})
})
}))
	})
	return cache_Data_Interval_bitraversableInterval__1397501818
}

var cache_Data_Interval_extendInterval__2367002404 gopurs_runtime.Value
var once_Data_Interval_extendInterval__2367002404 sync.Once
func Get_Data_Interval_extendInterval__2367002404() gopurs_runtime.Value {
	once_Data_Interval_extendInterval__2367002404.Do(func() {
		cache_Data_Interval_extendInterval__2367002404 = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_functorInterval()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Apply(v_0, v1_1)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, (*Constructor_Data_Interval_DurationEnd)(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(v_0, v1_1)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, gopurs_runtime.Apply(v_0, v1_1), (*Constructor_Data_Interval_StartDuration)(v1_1.UnsafePtr).V1})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, (*Constructor_Data_Interval_DurationOnly)(v1_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_Data_Interval_extendInterval__2367002404
}

var cache_Data_Interval_foldableInterval__3140210451 gopurs_runtime.Value
var once_Data_Interval_foldableInterval__3140210451 sync.Once
func Get_Data_Interval_foldableInterval__3140210451() gopurs_runtime.Value {
	once_Data_Interval_foldableInterval__3140210451.Do(func() {
		cache_Data_Interval_foldableInterval__3140210451 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), mempty_2_1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Interval_StartEnd)(v2_4.UnsafePtr).V0)), gopurs_runtime.Apply(f_3, (*Constructor_Data_Interval_StartEnd)(v2_4.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), mempty_2_1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Interval_DurationEnd)(v2_4.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), mempty_2_1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Interval_StartDuration)(v2_4.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
__t2 = mempty_2_1
}
end_branch_2:
return __t2
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t3 = gopurs_runtime.Apply2(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V1)
goto end_branch_3
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t3 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_DurationEnd)(v2_2.UnsafePtr).V1)
goto end_branch_3
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t3 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_StartDuration)(v2_2.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
__t3 = v1_1
}
end_branch_3:
return __t3
})
})
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Foldable_foldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_Interval_foldableInterval()))}, x_0)
}))
	})
	return cache_Data_Interval_foldableInterval__3140210451
}

var cache_Data_Interval_foldableInterval__526261656 gopurs_runtime.Value
var once_Data_Interval_foldableInterval__526261656 sync.Once
func Get_Data_Interval_foldableInterval__526261656() gopurs_runtime.Value {
	once_Data_Interval_foldableInterval__526261656.Do(func() {
		cache_Data_Interval_foldableInterval__526261656 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), mempty_2_1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Interval_StartEnd)(v2_4.UnsafePtr).V0)), gopurs_runtime.Apply(f_3, (*Constructor_Data_Interval_StartEnd)(v2_4.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), mempty_2_1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Interval_DurationEnd)(v2_4.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), mempty_2_1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Interval_StartDuration)(v2_4.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
__t2 = mempty_2_1
}
end_branch_2:
return __t2
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t3 = gopurs_runtime.Apply2(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V1)
goto end_branch_3
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t3 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_DurationEnd)(v2_2.UnsafePtr).V1)
goto end_branch_3
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t3 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_StartDuration)(v2_2.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
__t3 = v1_1
}
end_branch_3:
return __t3
})
})
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Foldable_foldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_Interval_foldableInterval()))}, x_0)
}))
	})
	return cache_Data_Interval_foldableInterval__526261656
}

var cache_Data_Interval_foldableRecurringInterval__598519513 gopurs_runtime.Value
var once_Data_Interval_foldableRecurringInterval__598519513 sync.Once
func Get_Data_Interval_foldableRecurringInterval__598519513() gopurs_runtime.Value {
	once_Data_Interval_foldableRecurringInterval__598519513.Do(func() {
		cache_Data_Interval_foldableRecurringInterval__598519513 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(x_4.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 237113226) {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), mempty_2_1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(x_4.UnsafePtr).V1.UnsafePtr).V0)), gopurs_runtime.Apply(f_3, (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(x_4.UnsafePtr).V1.UnsafePtr).V1))
goto end_branch_5
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(x_4.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1992629780) {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), mempty_2_1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Interval_DurationEnd)((*Constructor_Data_Interval_RecurringInterval)(x_4.UnsafePtr).V1.UnsafePtr).V1))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(x_4.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2020675835) {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), mempty_2_1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Interval_StartDuration)((*Constructor_Data_Interval_RecurringInterval)(x_4.UnsafePtr).V1.UnsafePtr).V0))
goto end_branch_5
} else {

}
}
{
__t5 = mempty_2_1
}
end_branch_5:
return __t5
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 237113226) {
__t9 = gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply2(f_0, i_1, (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1.UnsafePtr).V1)
goto end_branch_9
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 1992629780) {
__t9 = gopurs_runtime.Apply2(f_0, i_1, (*Constructor_Data_Interval_DurationEnd)((*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1.UnsafePtr).V1)
goto end_branch_9
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 2020675835) {
__t9 = gopurs_runtime.Apply2(f_0, i_1, (*Constructor_Data_Interval_StartDuration)((*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1.UnsafePtr).V0)
goto end_branch_9
} else {

}
}
{
__t9 = i_1
}
end_branch_9:
return __t9
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Data_Foldable_foldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_Interval_foldableInterval()))}, f_0, i_1, (*Constructor_Data_Interval_RecurringInterval)(x_2.UnsafePtr).V1)
})
})
}))
	})
	return cache_Data_Interval_foldableRecurringInterval__598519513
}

var cache_Data_Interval_functorInterval__3565473811 gopurs_runtime.Value
var once_Data_Interval_functorInterval__3565473811 sync.Once
func Get_Data_Interval_functorInterval__3565473811() gopurs_runtime.Value {
	once_Data_Interval_functorInterval__3565473811.Do(func() {
		cache_Data_Interval_functorInterval__3565473811 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_1.Type == 9 && v2_1.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartEnd)(v2_1.UnsafePtr).V0), gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartEnd)(v2_1.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, (*Constructor_Data_Interval_DurationEnd)(v2_1.UnsafePtr).V0, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_DurationEnd)(v2_1.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartDuration)(v2_1.UnsafePtr).V0), (*Constructor_Data_Interval_StartDuration)(v2_1.UnsafePtr).V1})}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, (*Constructor_Data_Interval_DurationOnly)(v2_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_Data_Interval_functorInterval__3565473811
}

var cache_Data_Interval_functorInterval__3172181073 gopurs_runtime.Value
var once_Data_Interval_functorInterval__3172181073 sync.Once
func Get_Data_Interval_functorInterval__3172181073() gopurs_runtime.Value {
	once_Data_Interval_functorInterval__3172181073.Do(func() {
		cache_Data_Interval_functorInterval__3172181073 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_1.Type == 9 && v2_1.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartEnd)(v2_1.UnsafePtr).V0), gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartEnd)(v2_1.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, (*Constructor_Data_Interval_DurationEnd)(v2_1.UnsafePtr).V0, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_DurationEnd)(v2_1.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, gopurs_runtime.Apply(v1_0, (*Constructor_Data_Interval_StartDuration)(v2_1.UnsafePtr).V0), (*Constructor_Data_Interval_StartDuration)(v2_1.UnsafePtr).V1})}
goto end_branch_0
} else {

}
}
{
if (v2_1.Type == 9 && v2_1.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, (*Constructor_Data_Interval_DurationOnly)(v2_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_Data_Interval_functorInterval__3172181073
}

var cache_Data_Interval_functorRecurringInterval__1167854705 gopurs_runtime.Value
var once_Data_Interval_functorRecurringInterval__1167854705 sync.Once
func Get_Data_Interval_functorRecurringInterval__1167854705() gopurs_runtime.Value {
	once_Data_Interval_functorRecurringInterval__1167854705.Do(func() {
		cache_Data_Interval_functorRecurringInterval__1167854705 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 237113226) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartEnd{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartEnd)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V1)})}
goto end_branch_4
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1992629780) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationEnd{1, (*Constructor_Data_Interval_DurationEnd)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_DurationEnd)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V1)})}
goto end_branch_4
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2020675835) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Interval_StartDuration)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V0), (*Constructor_Data_Interval_StartDuration)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V1})}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2281256335) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, (*Constructor_Data_Interval_DurationOnly)((*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V1.UnsafePtr).V0})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_RecurringInterval{1, (*Constructor_Data_Interval_RecurringInterval)(v_1.UnsafePtr).V0, __t4})}
})
}))
	})
	return cache_Data_Interval_functorRecurringInterval__1167854705
}

var cache_Data_Interval_interval__413767002 gopurs_runtime.Value
var once_Data_Interval_interval__413767002 sync.Once
func Get_Data_Interval_interval__413767002() gopurs_runtime.Value {
	once_Data_Interval_interval__413767002.Do(func() {
		cache_Data_Interval_interval__413767002 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_interval__413767002(gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval](v_0_box))
})
	})
	return cache_Data_Interval_interval__413767002
}

var cache_Data_Interval_over__3140749367 gopurs_runtime.Value
var once_Data_Interval_over__3140749367 sync.Once
func Get_Data_Interval_over__3140749367() gopurs_runtime.Value {
	once_Data_Interval_over__3140749367.Do(func() {
		cache_Data_Interval_over__3140749367 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_over__3140749367(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Interval_RecurringInterval](v_2_box))
})
	})
	return cache_Data_Interval_over__3140749367
}

var cache_Data_Interval_traversableInterval__1002691347 gopurs_runtime.Value
var once_Data_Interval_traversableInterval__1002691347 sync.Once
func Get_Data_Interval_traversableInterval__1002691347() gopurs_runtime.Value {
	once_Data_Interval_traversableInterval__1002691347.Do(func() {
		cache_Data_Interval_traversableInterval__1002691347 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_foldableInterval()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_functorInterval()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Interval_traversableInterval(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 237113226) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartEnd)(v1_4.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartEnd)(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1992629780) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Apply(Get_Data_Interval_DurationEnd(), (*Constructor_Data_Interval_DurationEnd)(v1_4.UnsafePtr).V0), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationEnd)(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2020675835) {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := (*Constructor_Data_Interval_StartDuration)(v1_4.UnsafePtr).V1
_ = __local_var_5_2
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_StartDuration{1, v2_6, __local_var_5_2})}
}), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartDuration)(v1_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2281256335) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_DurationOnly{1, (*Constructor_Data_Interval_DurationOnly)(v1_4.UnsafePtr).V0})})
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
})
}))
	})
	return cache_Data_Interval_traversableInterval__1002691347
}

type Constructor_Data_Interval_StartEnd struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_Data_Interval_DurationEnd struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_Data_Interval_StartDuration struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_Data_Interval_DurationOnly struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


type Constructor_Data_Interval_RecurringInterval struct {
	Rc uint32
	V0 *Constructor_Data_Maybe_Just
	V1 gopurs_runtime.Value
}


func Call_Data_Interval_showInterval(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 string
{
if (v_2.Type == 9 && v_2.IntVal == 237113226) {
__t0 = (((("(StartEnd ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartEnd)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartEnd)(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1992629780) {
__t0 = (((("(DurationEnd ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_DurationEnd)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_DurationEnd)(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2020675835) {
__t0 = (((("(StartDuration ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartDuration)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_StartDuration)(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2281256335) {
__t0 = (("(DurationOnly ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_DurationOnly)(v_2.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_0:
return gopurs_runtime.Str(__t0)
}))
}

func Call_Data_Interval_showRecurringInterval(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
// TAST (Let): showInterval2_2_0 -> *Constructor_Data_Show_Show
showInterval2_2_0 := &Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 string
{
if (v_2.Type == 9 && v_2.IntVal == 237113226) {
__t1 = (((("(StartEnd ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartEnd)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartEnd)(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1992629780) {
__t1 = (((("(DurationEnd ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_DurationEnd)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_DurationEnd)(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2020675835) {
__t1 = (((("(StartDuration ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Interval_StartDuration)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_StartDuration)(v_2.UnsafePtr).V1).StrVal())) + (")")
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2281256335) {
__t1 = (("(DurationOnly ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Interval_DurationOnly)(v_2.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_1:
return gopurs_runtime.Str(__t1)
})}
_ = showInterval2_2_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(RecurringInterval ") + (gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Data_Interval_showMaybe()).V0), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Interval_RecurringInterval)(v_3.UnsafePtr).V0)}).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(showInterval2_2_0.V0), (*Constructor_Data_Interval_RecurringInterval)(v_3.UnsafePtr).V1).StrVal())) + (")"))
}))
}

func Call_Data_Interval_over(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Interval_RecurringInterval) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_Interval_RecurringInterval = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Apply(Get_Data_Interval_RecurringInterval(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((v_2).V0)}), gopurs_runtime.Apply(f_1, (v_2).V1))
}

func Call_Data_Interval_interval(v_0_loop *Constructor_Data_Interval_RecurringInterval) gopurs_runtime.Value {
var v_0 *Constructor_Data_Interval_RecurringInterval = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Interval_eqInterval(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
if (x_2.Type == 9 && x_2.IntVal == 237113226) {
var __t0 bool
{
if (y_3.Type == 9 && y_3.IntVal == 237113226) {
__t0 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartEnd)(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartEnd)(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd)(y_3.UnsafePtr).V1).IntVal) != (0))
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t3 = __t0
goto end_branch_3
} else {

}
}
{
if (x_2.Type == 9 && x_2.IntVal == 1992629780) {
var __t1 bool
{
if (y_3.Type == 9 && y_3.IntVal == 1992629780) {
__t1 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_DurationEnd)(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_DurationEnd)(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd)(y_3.UnsafePtr).V1).IntVal) != (0))
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t3 = __t1
goto end_branch_3
} else {

}
}
{
if (x_2.Type == 9 && x_2.IntVal == 2020675835) {
var __t2 bool
{
if (y_3.Type == 9 && y_3.IntVal == 2020675835) {
__t2 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartDuration)(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_StartDuration)(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration)(y_3.UnsafePtr).V1).IntVal) != (0))
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
if ((x_2.Type == 9 && x_2.IntVal == 2281256335)) && ((y_3.Type == 9 && y_3.IntVal == 2281256335)) {
__t3 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_DurationOnly)(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly)(y_3.UnsafePtr).V0).IntVal) != (0)
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
}

func Call_Data_Interval_eqRecurringInterval(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
// TAST (Let): eqInterval2_2_0 -> *Constructor_Data_Eq_Eq
eqInterval2_2_0 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 bool
{
if (x_2.Type == 9 && x_2.IntVal == 237113226) {
var __t1 bool
{
if (y_3.Type == 9 && y_3.IntVal == 237113226) {
__t1 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartEnd)(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartEnd)(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd)(y_3.UnsafePtr).V1).IntVal) != (0))
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t4 = __t1
goto end_branch_4
} else {

}
}
{
if (x_2.Type == 9 && x_2.IntVal == 1992629780) {
var __t2 bool
{
if (y_3.Type == 9 && y_3.IntVal == 1992629780) {
__t2 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_DurationEnd)(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_DurationEnd)(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd)(y_3.UnsafePtr).V1).IntVal) != (0))
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t4 = __t2
goto end_branch_4
} else {

}
}
{
if (x_2.Type == 9 && x_2.IntVal == 2020675835) {
var __t3 bool
{
if (y_3.Type == 9 && y_3.IntVal == 2020675835) {
__t3 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Interval_StartDuration)(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration)(y_3.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_StartDuration)(x_2.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration)(y_3.UnsafePtr).V1).IntVal) != (0))
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if ((x_2.Type == 9 && x_2.IntVal == 2281256335)) && ((y_3.Type == 9 && y_3.IntVal == 2281256335)) {
__t4 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Interval_DurationOnly)(x_2.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly)(y_3.UnsafePtr).V0).IntVal) != (0)
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
return gopurs_runtime.Bool(__t4)
})
})}
_ = eqInterval2_2_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 bool
{
var __t_tag_5 *Constructor_Data_Maybe_Just = (*Constructor_Data_Interval_RecurringInterval)(x_3.UnsafePtr).V0
if (__t_tag_5 == nil) {
var __t7 bool
{
var __t_tag_6 *Constructor_Data_Maybe_Just = (*Constructor_Data_Interval_RecurringInterval)(y_4.UnsafePtr).V0
if (__t_tag_6 == nil) {
__t7 = true
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
__t11 = __t7
goto end_branch_11
} else {

}
}
{
var __t_tag_8 *Constructor_Data_Maybe_Just = (*Constructor_Data_Interval_RecurringInterval)(x_3.UnsafePtr).V0
var __t_and_10 bool = false
if (__t_tag_8 != nil) {

var __t_tag_9 *Constructor_Data_Maybe_Just = (*Constructor_Data_Interval_RecurringInterval)(y_4.UnsafePtr).V0
__t_and_10 = (__t_tag_9 != nil)
}
if __t_and_10 {
__t11 = (((*Constructor_Data_Interval_RecurringInterval)(x_3.UnsafePtr).V0).V0.IntVal) == (((*Constructor_Data_Interval_RecurringInterval)(y_4.UnsafePtr).V0).V0.IntVal)
goto end_branch_11
} else {

}
}
{
__t11 = false
}
end_branch_11:
return gopurs_runtime.Bool((__t11) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqInterval2_2_0.V0), (*Constructor_Data_Interval_RecurringInterval)(x_3.UnsafePtr).V1, (*Constructor_Data_Interval_RecurringInterval)(y_4.UnsafePtr).V1).IntVal) != (0)))
})
}))
}

func Call_Data_Interval_ordInterval(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): eqInterval2_3_1 -> gopurs_runtime.Value
eqInterval2_3_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
if (x_4.Type == 9 && x_4.IntVal == 237113226) {
var __t3 bool
{
if (y_5.Type == 9 && y_5.IntVal == 237113226) {
__t3 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "eq"), (*Constructor_Data_Interval_StartEnd)(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd)(y_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "eq"), (*Constructor_Data_Interval_StartEnd)(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd)(y_5.UnsafePtr).V1).IntVal) != (0))
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
if (x_4.Type == 9 && x_4.IntVal == 1992629780) {
var __t4 bool
{
if (y_5.Type == 9 && y_5.IntVal == 1992629780) {
__t4 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Constructor_Data_Interval_DurationEnd)(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd)(y_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "eq"), (*Constructor_Data_Interval_DurationEnd)(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd)(y_5.UnsafePtr).V1).IntVal) != (0))
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
if (x_4.Type == 9 && x_4.IntVal == 2020675835) {
var __t5 bool
{
if (y_5.Type == 9 && y_5.IntVal == 2020675835) {
__t5 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "eq"), (*Constructor_Data_Interval_StartDuration)(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration)(y_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Constructor_Data_Interval_StartDuration)(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration)(y_5.UnsafePtr).V1).IntVal) != (0))
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
if ((x_4.Type == 9 && x_4.IntVal == 2281256335)) && ((y_5.Type == 9 && y_5.IntVal == 2281256335)) {
__t6 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Constructor_Data_Interval_DurationOnly)(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly)(y_5.UnsafePtr).V0).IntVal) != (0)
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
_ = eqInterval2_3_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqInterval2_3_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 uint32
{
if (x_4.Type == 9 && x_4.IntVal == 237113226) {
var __t9 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 237113226) {
// TAST (Let): v_6_7 -> gopurs_runtime.Value
v_6_7 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_StartEnd)(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd)(y_5.UnsafePtr).V0)
_ = v_6_7
var __t8 uint32
{
if (uint32(v_6_7.IntVal) == 1527465420) {
__t8 = 1527465420
goto end_branch_8
} else {

}
}
{
if (uint32(v_6_7.IntVal) == 380165415) {
__t8 = 380165415
goto end_branch_8
} else {

}
}
{
__t8 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_StartEnd)(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd)(y_5.UnsafePtr).V1).IntVal)
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
__t16 = __t9
goto end_branch_16
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 237113226) {
__t16 = 380165415
goto end_branch_16
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 1992629780) {
var __t12 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 1992629780) {
// TAST (Let): v_6_10 -> gopurs_runtime.Value
v_6_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_DurationEnd)(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd)(y_5.UnsafePtr).V0)
_ = v_6_10
var __t11 uint32
{
if (uint32(v_6_10.IntVal) == 1527465420) {
__t11 = 1527465420
goto end_branch_11
} else {

}
}
{
if (uint32(v_6_10.IntVal) == 380165415) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
__t11 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_DurationEnd)(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd)(y_5.UnsafePtr).V1).IntVal)
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
__t16 = __t12
goto end_branch_16
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 1992629780) {
__t16 = 380165415
goto end_branch_16
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 2020675835) {
var __t15 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 2020675835) {
// TAST (Let): v_6_13 -> gopurs_runtime.Value
v_6_13 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Interval_StartDuration)(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration)(y_5.UnsafePtr).V0)
_ = v_6_13
var __t14 uint32
{
if (uint32(v_6_13.IntVal) == 1527465420) {
__t14 = 1527465420
goto end_branch_14
} else {

}
}
{
if (uint32(v_6_13.IntVal) == 380165415) {
__t14 = 380165415
goto end_branch_14
} else {

}
}
{
__t14 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_StartDuration)(x_4.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration)(y_5.UnsafePtr).V1).IntVal)
}
end_branch_14:
__t15 = __t14
goto end_branch_15
} else {

}
}
{
__t15 = 1527465420
}
end_branch_15:
__t16 = __t15
goto end_branch_16
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 2020675835) {
__t16 = 380165415
goto end_branch_16
} else {

}
}
{
if ((x_4.Type == 9 && x_4.IntVal == 2281256335)) && ((y_5.Type == 9 && y_5.IntVal == 2281256335)) {
__t16 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_DurationOnly)(x_4.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly)(y_5.UnsafePtr).V0).IntVal)
goto end_branch_16
} else {

}
}
{
__t16 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_16:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t16), UnsafePtr: nil}
})
}))
})
}

func Call_Data_Interval_ordRecurringInterval(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_4_4
// TAST (Let): eqInterval2_4_3 -> gopurs_runtime.Value
eqInterval2_4_3 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 bool
{
if (x_5.Type == 9 && x_5.IntVal == 237113226) {
var __t5 bool
{
if (y_6.Type == 9 && y_6.IntVal == 237113226) {
__t5 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "eq"), (*Constructor_Data_Interval_StartEnd)(x_5.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd)(y_6.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "eq"), (*Constructor_Data_Interval_StartEnd)(x_5.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd)(y_6.UnsafePtr).V1).IntVal) != (0))
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
__t8 = __t5
goto end_branch_8
} else {

}
}
{
if (x_5.Type == 9 && x_5.IntVal == 1992629780) {
var __t6 bool
{
if (y_6.Type == 9 && y_6.IntVal == 1992629780) {
__t6 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Constructor_Data_Interval_DurationEnd)(x_5.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd)(y_6.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "eq"), (*Constructor_Data_Interval_DurationEnd)(x_5.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd)(y_6.UnsafePtr).V1).IntVal) != (0))
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t8 = __t6
goto end_branch_8
} else {

}
}
{
if (x_5.Type == 9 && x_5.IntVal == 2020675835) {
var __t7 bool
{
if (y_6.Type == 9 && y_6.IntVal == 2020675835) {
__t7 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "eq"), (*Constructor_Data_Interval_StartDuration)(x_5.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration)(y_6.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Constructor_Data_Interval_StartDuration)(x_5.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration)(y_6.UnsafePtr).V1).IntVal) != (0))
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
if ((x_5.Type == 9 && x_5.IntVal == 2281256335)) && ((y_6.Type == 9 && y_6.IntVal == 2281256335)) {
__t8 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Constructor_Data_Interval_DurationOnly)(x_5.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly)(y_6.UnsafePtr).V0).IntVal) != (0)
goto end_branch_8
} else {

}
}
{
__t8 = false
}
end_branch_8:
return gopurs_runtime.Bool(__t8)
})
}))
_ = eqInterval2_4_3
// TAST (Let): ordInterval2_4_2 -> *Constructor_Data_Ord_Ord
ordInterval2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqInterval2_4_3
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 uint32
{
if (x_5.Type == 9 && x_5.IntVal == 237113226) {
var __t11 uint32
{
if (y_6.Type == 9 && y_6.IntVal == 237113226) {
// TAST (Let): v_7_9 -> gopurs_runtime.Value
v_7_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_3, "compare"), (*Constructor_Data_Interval_StartEnd)(x_5.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd)(y_6.UnsafePtr).V0)
_ = v_7_9
var __t10 uint32
{
if (uint32(v_7_9.IntVal) == 1527465420) {
__t10 = 1527465420
goto end_branch_10
} else {

}
}
{
if (uint32(v_7_9.IntVal) == 380165415) {
__t10 = 380165415
goto end_branch_10
} else {

}
}
{
__t10 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_3, "compare"), (*Constructor_Data_Interval_StartEnd)(x_5.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd)(y_6.UnsafePtr).V1).IntVal)
}
end_branch_10:
__t11 = __t10
goto end_branch_11
} else {

}
}
{
__t11 = 1527465420
}
end_branch_11:
__t18 = __t11
goto end_branch_18
} else {

}
}
{
if (y_6.Type == 9 && y_6.IntVal == 237113226) {
__t18 = 380165415
goto end_branch_18
} else {

}
}
{
if (x_5.Type == 9 && x_5.IntVal == 1992629780) {
var __t14 uint32
{
if (y_6.Type == 9 && y_6.IntVal == 1992629780) {
// TAST (Let): v_7_12 -> gopurs_runtime.Value
v_7_12 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_DurationEnd)(x_5.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd)(y_6.UnsafePtr).V0)
_ = v_7_12
var __t13 uint32
{
if (uint32(v_7_12.IntVal) == 1527465420) {
__t13 = 1527465420
goto end_branch_13
} else {

}
}
{
if (uint32(v_7_12.IntVal) == 380165415) {
__t13 = 380165415
goto end_branch_13
} else {

}
}
{
__t13 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_3, "compare"), (*Constructor_Data_Interval_DurationEnd)(x_5.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd)(y_6.UnsafePtr).V1).IntVal)
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = 1527465420
}
end_branch_14:
__t18 = __t14
goto end_branch_18
} else {

}
}
{
if (y_6.Type == 9 && y_6.IntVal == 1992629780) {
__t18 = 380165415
goto end_branch_18
} else {

}
}
{
if (x_5.Type == 9 && x_5.IntVal == 2020675835) {
var __t17 uint32
{
if (y_6.Type == 9 && y_6.IntVal == 2020675835) {
// TAST (Let): v_7_15 -> gopurs_runtime.Value
v_7_15 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_3, "compare"), (*Constructor_Data_Interval_StartDuration)(x_5.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration)(y_6.UnsafePtr).V0)
_ = v_7_15
var __t16 uint32
{
if (uint32(v_7_15.IntVal) == 1527465420) {
__t16 = 1527465420
goto end_branch_16
} else {

}
}
{
if (uint32(v_7_15.IntVal) == 380165415) {
__t16 = 380165415
goto end_branch_16
} else {

}
}
{
__t16 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_StartDuration)(x_5.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration)(y_6.UnsafePtr).V1).IntVal)
}
end_branch_16:
__t17 = __t16
goto end_branch_17
} else {

}
}
{
__t17 = 1527465420
}
end_branch_17:
__t18 = __t17
goto end_branch_18
} else {

}
}
{
if (y_6.Type == 9 && y_6.IntVal == 2020675835) {
__t18 = 380165415
goto end_branch_18
} else {

}
}
{
if ((x_5.Type == 9 && x_5.IntVal == 2281256335)) && ((y_6.Type == 9 && y_6.IntVal == 2281256335)) {
__t18 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Interval_DurationOnly)(x_5.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly)(y_6.UnsafePtr).V0).IntVal)
goto end_branch_18
} else {

}
}
{
__t18 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_18:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t18), UnsafePtr: nil}
})
})))
_ = ordInterval2_4_2
// TAST (Let): __local_var_5_20 -> gopurs_runtime.Value
__local_var_5_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_5_20
// TAST (Let): eqInterval2_6_21 -> *Constructor_Data_Eq_Eq
eqInterval2_6_21 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t25 bool
{
if (x_6.Type == 9 && x_6.IntVal == 237113226) {
var __t22 bool
{
if (y_7.Type == 9 && y_7.IntVal == 237113226) {
__t22 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_20, "eq"), (*Constructor_Data_Interval_StartEnd)(x_6.UnsafePtr).V0, (*Constructor_Data_Interval_StartEnd)(y_7.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_20, "eq"), (*Constructor_Data_Interval_StartEnd)(x_6.UnsafePtr).V1, (*Constructor_Data_Interval_StartEnd)(y_7.UnsafePtr).V1).IntVal) != (0))
goto end_branch_22
} else {

}
}
{
__t22 = false
}
end_branch_22:
__t25 = __t22
goto end_branch_25
} else {

}
}
{
if (x_6.Type == 9 && x_6.IntVal == 1992629780) {
var __t23 bool
{
if (y_7.Type == 9 && y_7.IntVal == 1992629780) {
__t23 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "eq"), (*Constructor_Data_Interval_DurationEnd)(x_6.UnsafePtr).V0, (*Constructor_Data_Interval_DurationEnd)(y_7.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_20, "eq"), (*Constructor_Data_Interval_DurationEnd)(x_6.UnsafePtr).V1, (*Constructor_Data_Interval_DurationEnd)(y_7.UnsafePtr).V1).IntVal) != (0))
goto end_branch_23
} else {

}
}
{
__t23 = false
}
end_branch_23:
__t25 = __t23
goto end_branch_25
} else {

}
}
{
if (x_6.Type == 9 && x_6.IntVal == 2020675835) {
var __t24 bool
{
if (y_7.Type == 9 && y_7.IntVal == 2020675835) {
__t24 = ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_20, "eq"), (*Constructor_Data_Interval_StartDuration)(x_6.UnsafePtr).V0, (*Constructor_Data_Interval_StartDuration)(y_7.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "eq"), (*Constructor_Data_Interval_StartDuration)(x_6.UnsafePtr).V1, (*Constructor_Data_Interval_StartDuration)(y_7.UnsafePtr).V1).IntVal) != (0))
goto end_branch_24
} else {

}
}
{
__t24 = false
}
end_branch_24:
__t25 = __t24
goto end_branch_25
} else {

}
}
{
if ((x_6.Type == 9 && x_6.IntVal == 2281256335)) && ((y_7.Type == 9 && y_7.IntVal == 2281256335)) {
__t25 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "eq"), (*Constructor_Data_Interval_DurationOnly)(x_6.UnsafePtr).V0, (*Constructor_Data_Interval_DurationOnly)(y_7.UnsafePtr).V0).IntVal) != (0)
goto end_branch_25
} else {

}
}
{
__t25 = false
}
end_branch_25:
return gopurs_runtime.Bool(__t25)
})
})}
_ = eqInterval2_6_21
// TAST (Let): eqRecurringInterval2_5_19 -> gopurs_runtime.Value
eqRecurringInterval2_5_19 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t32 bool
{
var __t_tag_26 *Constructor_Data_Maybe_Just = (*Constructor_Data_Interval_RecurringInterval)(x_7.UnsafePtr).V0
if (__t_tag_26 == nil) {
var __t28 bool
{
var __t_tag_27 *Constructor_Data_Maybe_Just = (*Constructor_Data_Interval_RecurringInterval)(y_8.UnsafePtr).V0
if (__t_tag_27 == nil) {
__t28 = true
goto end_branch_28
} else {

}
}
{
__t28 = false
}
end_branch_28:
__t32 = __t28
goto end_branch_32
} else {

}
}
{
var __t_tag_29 *Constructor_Data_Maybe_Just = (*Constructor_Data_Interval_RecurringInterval)(x_7.UnsafePtr).V0
var __t_and_31 bool = false
if (__t_tag_29 != nil) {

var __t_tag_30 *Constructor_Data_Maybe_Just = (*Constructor_Data_Interval_RecurringInterval)(y_8.UnsafePtr).V0
__t_and_31 = (__t_tag_30 != nil)
}
if __t_and_31 {
__t32 = (((*Constructor_Data_Interval_RecurringInterval)(x_7.UnsafePtr).V0).V0.IntVal) == (((*Constructor_Data_Interval_RecurringInterval)(y_8.UnsafePtr).V0).V0.IntVal)
goto end_branch_32
} else {

}
}
{
__t32 = false
}
end_branch_32:
return gopurs_runtime.Bool((__t32) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqInterval2_6_21.V0), (*Constructor_Data_Interval_RecurringInterval)(x_7.UnsafePtr).V1, (*Constructor_Data_Interval_RecurringInterval)(y_8.UnsafePtr).V1).IntVal) != (0)))
})
}))
_ = eqRecurringInterval2_5_19
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRecurringInterval2_5_19
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 uint32
{
var __t_tag_34 *Constructor_Data_Maybe_Just = (*Constructor_Data_Interval_RecurringInterval)(x_6.UnsafePtr).V0
if (__t_tag_34 == nil) {
var __t36 uint32
{
var __t_tag_35 *Constructor_Data_Maybe_Just = (*Constructor_Data_Interval_RecurringInterval)(y_7.UnsafePtr).V0
if (__t_tag_35 == nil) {
__t36 = 902936544
goto end_branch_36
} else {

}
}
{
__t36 = 1527465420
}
end_branch_36:
__t41 = __t36
goto end_branch_41
} else {

}
}
{
var __t_tag_37 *Constructor_Data_Maybe_Just = (*Constructor_Data_Interval_RecurringInterval)(y_7.UnsafePtr).V0
if (__t_tag_37 == nil) {
__t41 = 380165415
goto end_branch_41
} else {

}
}
{
var __t_tag_38 *Constructor_Data_Maybe_Just = (*Constructor_Data_Interval_RecurringInterval)(x_6.UnsafePtr).V0
var __t_and_40 bool = false
if (__t_tag_38 != nil) {

var __t_tag_39 *Constructor_Data_Maybe_Just = (*Constructor_Data_Interval_RecurringInterval)(y_7.UnsafePtr).V0
__t_and_40 = (__t_tag_39 != nil)
}
if __t_and_40 {
__t41 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Ord_ordInt(), "compare"), ((*Constructor_Data_Interval_RecurringInterval)(x_6.UnsafePtr).V0).V0, ((*Constructor_Data_Interval_RecurringInterval)(y_7.UnsafePtr).V0).V0).IntVal)
goto end_branch_41
} else {

}
}
{
__t41 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_41:
// TAST (Let): v_8_33 -> uint32
v_8_33 := __t41
_ = v_8_33
var __t42 uint32
{
if (v_8_33 == 1527465420) {
__t42 = 1527465420
goto end_branch_42
} else {

}
}
{
if (v_8_33 == 380165415) {
__t42 = 380165415
goto end_branch_42
} else {

}
}
{
__t42 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordInterval2_4_2.V1), (*Constructor_Data_Interval_RecurringInterval)(x_6.UnsafePtr).V1, (*Constructor_Data_Interval_RecurringInterval)(y_7.UnsafePtr).V1).IntVal)
}
end_branch_42:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t42), UnsafePtr: nil}
})
}))
})
}

func Call_Data_Interval_interval__413767002(v_0_loop *Constructor_Data_Interval_RecurringInterval) gopurs_runtime.Value {
var v_0 *Constructor_Data_Interval_RecurringInterval = v_0_loop
_ = v_0
return (v_0).V1
}

func Call_Data_Interval_over__3140749367(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Interval_RecurringInterval) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_Interval_RecurringInterval = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Apply(Get_Data_Interval_RecurringInterval(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((v_2).V0)}), gopurs_runtime.Apply(f_1, (v_2).V1))
}



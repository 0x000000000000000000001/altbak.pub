package Data_Interval

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Extend "gopurs/output/Control.Extend"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Bifoldable "gopurs/output/Data.Bifoldable"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Bitraversable "gopurs/output/Data.Bitraversable"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_showMaybe gopurs_runtime.Value
var once_showMaybe sync.Once
func Get_showMaybe() gopurs_runtime.Value {
	once_showMaybe.Do(func() {
		cache_showMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[*pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.Apply(pkg_Data_Maybe.Get_showMaybe(), pkg_Data_Show.Get_showInt())))}
	})
	return cache_showMaybe
}

var cache_eqMaybe gopurs_runtime.Value
var once_eqMaybe sync.Once
func Get_eqMaybe() gopurs_runtime.Value {
	once_eqMaybe.Do(func() {
		cache_eqMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&pkg_Data_Eq.Constructor_Eq[*pkg_Data_Maybe.Constructor_Just[int64]]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = ((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_0.UnsafePtr).V0.IntVal) == ((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_1.UnsafePtr).V0.IntVal)
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
	return cache_eqMaybe
}

var cache_ordMaybe gopurs_runtime.Value
var once_ordMaybe sync.Once
func Get_ordMaybe() gopurs_runtime.Value {
	once_ordMaybe.Do(func() {
		cache_ordMaybe = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}))
_ = __local_var_0_0
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_2
eqMaybe1_1_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 bool
{
if (x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr == nil) {
var __t3 bool
{
if (y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr == nil) {
__t3 = true
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
if ((x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr != nil)) && ((y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr != nil)) {
__t4 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "eq"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)
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
}))
_ = eqMaybe1_1_1
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_Maybe.Constructor_Just[int64]]](gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_1_1
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr == nil) {
var __t5 uint32
{
if (y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr == nil) {
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
if (y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if ((x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr != nil)) && ((y_3.Type == 9 && y_3.IntVal == 930809136 && y_3.UnsafePtr != nil)) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "compare"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
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
}))))}
}()
	})
	return cache_ordMaybe
}

var cache_StartEnd gopurs_runtime.Value
var once_StartEnd sync.Once
func Get_StartEnd() gopurs_runtime.Value {
	once_StartEnd.Do(func() {
		cache_StartEnd = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_StartEnd
}

var cache_DurationEnd gopurs_runtime.Value
var once_DurationEnd sync.Once
func Get_DurationEnd() gopurs_runtime.Value {
	once_DurationEnd.Do(func() {
		cache_DurationEnd = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_DurationEnd
}

var cache_StartDuration gopurs_runtime.Value
var once_StartDuration sync.Once
func Get_StartDuration() gopurs_runtime.Value {
	once_StartDuration.Do(func() {
		cache_StartDuration = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_StartDuration
}

var cache_DurationOnly gopurs_runtime.Value
var once_DurationOnly sync.Once
func Get_DurationOnly() gopurs_runtime.Value {
	once_DurationOnly.Do(func() {
		cache_DurationOnly = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0})}
})
	})
	return cache_DurationOnly
}

var cache_RecurringInterval gopurs_runtime.Value
var once_RecurringInterval sync.Once
func Get_RecurringInterval() gopurs_runtime.Value {
	once_RecurringInterval.Do(func() {
		cache_RecurringInterval = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](value0), value1})}
})
})
	})
	return cache_RecurringInterval
}

var cache_showInterval gopurs_runtime.Value
var once_showInterval sync.Once
func Get_showInterval() gopurs_runtime.Value {
	once_showInterval.Do(func() {
		cache_showInterval = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showInterval(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_showInterval
}

var cache_showRecurringInterval gopurs_runtime.Value
var once_showRecurringInterval sync.Once
func Get_showRecurringInterval() gopurs_runtime.Value {
	once_showRecurringInterval.Do(func() {
		cache_showRecurringInterval = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showRecurringInterval(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_showRecurringInterval
}

var cache_over gopurs_runtime.Value
var once_over sync.Once
func Get_over() gopurs_runtime.Value {
	once_over.Do(func() {
		cache_over = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box))
})
	})
	return cache_over
}

var cache_foldableInterval gopurs_runtime.Value
var once_foldableInterval sync.Once
func Get_foldableInterval() gopurs_runtime.Value {
	once_foldableInterval.Do(func() {
		cache_foldableInterval = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableInterval(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, acc_4, gopurs_runtime.Apply(f_3, x_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
return gopurs_runtime.Apply2(pkg_Data_Foldable.Get_foldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](Get_foldableInterval()))}, x_0)
}))
	})
	return cache_foldableInterval
}

var cache_foldableRecurringInterval gopurs_runtime.Value
var once_foldableRecurringInterval sync.Once
func Get_foldableRecurringInterval() gopurs_runtime.Value {
	once_foldableRecurringInterval.Do(func() {
		cache_foldableRecurringInterval = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableRecurringInterval(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, acc_4, gopurs_runtime.Apply(f_3, x_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableInterval(), "foldl"), f_0, i_1)
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableInterval(), "foldr"), f_0, i_1)
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_3, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1)
})
})
}))
	})
	return cache_foldableRecurringInterval
}

var cache_eqInterval gopurs_runtime.Value
var once_eqInterval sync.Once
func Get_eqInterval() gopurs_runtime.Value {
	once_eqInterval.Do(func() {
		cache_eqInterval = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqInterval(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_eqInterval
}

var cache_eqRecurringInterval gopurs_runtime.Value
var once_eqRecurringInterval sync.Once
func Get_eqRecurringInterval() gopurs_runtime.Value {
	once_eqRecurringInterval.Do(func() {
		cache_eqRecurringInterval = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqRecurringInterval(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_eqRecurringInterval
}

var cache_ordInterval gopurs_runtime.Value
var once_ordInterval sync.Once
func Get_ordInterval() gopurs_runtime.Value {
	once_ordInterval.Do(func() {
		cache_ordInterval = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordInterval(dictOrd_0_box)
})
	})
	return cache_ordInterval
}

var cache_ordRecurringInterval gopurs_runtime.Value
var once_ordRecurringInterval sync.Once
func Get_ordRecurringInterval() gopurs_runtime.Value {
	once_ordRecurringInterval.Do(func() {
		cache_ordRecurringInterval = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordRecurringInterval(dictOrd_0_box)
})
	})
	return cache_ordRecurringInterval
}

var cache_bifunctorInterval gopurs_runtime.Value
var once_bifunctorInterval sync.Once
func Get_bifunctorInterval() gopurs_runtime.Value {
	once_bifunctorInterval.Do(func() {
		cache_bifunctorInterval = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v_0, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
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
	return cache_bifunctorInterval
}

var cache_bifunctorRecurringInterval gopurs_runtime.Value
var once_bifunctorRecurringInterval sync.Once
func Get_bifunctorRecurringInterval() gopurs_runtime.Value {
	once_bifunctorRecurringInterval.Do(func() {
		cache_bifunctorRecurringInterval = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifunctorInterval(), "bimap"), f_0, g_1, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
})
})
}))
	})
	return cache_bifunctorRecurringInterval
}

var cache_functorInterval gopurs_runtime.Value
var once_functorInterval sync.Once
func Get_functorInterval() gopurs_runtime.Value {
	once_functorInterval.Do(func() {
		cache_functorInterval = gopurs_runtime.RecordDict1("map", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_bifunctorInterval(), "bimap"), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})))
	})
	return cache_functorInterval
}

var cache_extendInterval gopurs_runtime.Value
var once_extendInterval sync.Once
func Get_extendInterval() gopurs_runtime.Value {
	once_extendInterval.Do(func() {
		cache_extendInterval = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorInterval()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Apply(v_0, v1_1)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(v_0, v1_1)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1), (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0})}
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
	return cache_extendInterval
}

var cache_functorRecurringInterval gopurs_runtime.Value
var once_functorRecurringInterval sync.Once
func Get_functorRecurringInterval() gopurs_runtime.Value {
	once_functorRecurringInterval.Do(func() {
		cache_functorRecurringInterval = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorInterval(), "map"), f_0, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorRecurringInterval
}

var cache_extendRecurringInterval gopurs_runtime.Value
var once_extendRecurringInterval sync.Once
func Get_extendRecurringInterval() gopurs_runtime.Value {
	once_extendRecurringInterval.Do(func() {
		cache_extendRecurringInterval = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorRecurringInterval()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](v_1))})
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_extendInterval(), "extend"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_0
}), (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_extendRecurringInterval
}

var cache_traversableInterval gopurs_runtime.Value
var once_traversableInterval sync.Once
func Get_traversableInterval() gopurs_runtime.Value {
	once_traversableInterval.Do(func() {
		cache_traversableInterval = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableInterval()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorInterval()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableInterval(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, pkg_Data_Traversable.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 237113226) {
__t3 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_StartEnd(), gopurs_runtime.Apply(v_3, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1992629780) {
__t3 = gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Apply(Get_DurationEnd(), (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0), gopurs_runtime.Apply(v_3, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2020675835) {
__local_var_5_2 := (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_2
__t3 = gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, v2_6, __local_var_5_2})}
}), gopurs_runtime.Apply(v_3, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2281256335) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0})})
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
	return cache_traversableInterval
}

var cache_traversableRecurringInterval gopurs_runtime.Value
var once_traversableRecurringInterval sync.Once
func Get_traversableRecurringInterval() gopurs_runtime.Value {
	once_traversableRecurringInterval.Do(func() {
		cache_traversableRecurringInterval = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableRecurringInterval()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorRecurringInterval()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableRecurringInterval(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, pkg_Data_Traversable.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over(Functor0_1_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableInterval(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, f_2), gopurs_runtime.CoerceToStruct[Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](i_3))
})
})
}))
	})
	return cache_traversableRecurringInterval
}

var cache_bifoldableInterval gopurs_runtime.Value
var once_bifoldableInterval sync.Once
func Get_bifoldableInterval() gopurs_runtime.Value {
	once_bifoldableInterval.Do(func() {
		cache_bifoldableInterval = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifoldableInterval(), "bifoldl"), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, m_5, gopurs_runtime.Apply(f_3, a_6))
})
}), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, m_5, gopurs_runtime.Apply(g_4, b_6))
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
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v1_1, v2_2, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1), (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
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
return gopurs_runtime.Apply2(pkg_Data_Bifoldable.Get_bifoldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]](Get_bifoldableInterval()))}, x_0)
}))
	})
	return cache_bifoldableInterval
}

var cache_bifoldableRecurringInterval gopurs_runtime.Value
var once_bifoldableRecurringInterval sync.Once
func Get_bifoldableRecurringInterval() gopurs_runtime.Value {
	once_bifoldableRecurringInterval.Do(func() {
		cache_bifoldableRecurringInterval = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifoldableRecurringInterval(), "bifoldl"), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, m_5, gopurs_runtime.Apply(f_3, a_6))
})
}), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, m_5, gopurs_runtime.Apply(g_4, b_6))
})
}), mempty_2_1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifoldableInterval(), "bifoldl"), f_0, g_1, i_2)
_ = __local_var_3_2
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1)
})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_3 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifoldableInterval(), "bifoldr"), f_0, g_1, i_2)
_ = __local_var_3_3
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1)
})
})
})
}))
	})
	return cache_bifoldableRecurringInterval
}

var cache_bitraversableInterval gopurs_runtime.Value
var once_bitraversableInterval sync.Once
func Get_bitraversableInterval() gopurs_runtime.Value {
	once_bitraversableInterval.Do(func() {
		cache_bitraversableInterval = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bifoldableInterval()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bifunctorInterval()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bitraversableInterval(), "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, pkg_Data_Bitraversable.Get_identity(), pkg_Data_Bitraversable.Get_identity1())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_StartEnd(), gopurs_runtime.Apply(v1_4, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_DurationEnd(), gopurs_runtime.Apply(v_3, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_StartDuration(), gopurs_runtime.Apply(v1_4, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(Functor0_2_1.V0, Get_DurationOnly(), gopurs_runtime.Apply(v_3, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0))
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
	return cache_bitraversableInterval
}

var cache_bitraversableRecurringInterval gopurs_runtime.Value
var once_bitraversableRecurringInterval sync.Once
func Get_bitraversableRecurringInterval() gopurs_runtime.Value {
	once_bitraversableRecurringInterval.Do(func() {
		cache_bitraversableRecurringInterval = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bifoldableRecurringInterval()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bifunctorRecurringInterval()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bitraversableRecurringInterval(), "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, pkg_Data_Bitraversable.Get_identity(), pkg_Data_Bitraversable.Get_identity1())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over(Functor0_1_0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bitraversableInterval(), "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, l_2, r_3), gopurs_runtime.CoerceToStruct[Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](i_4))
})
})
})
}))
	})
	return cache_bitraversableRecurringInterval
}

var cache_pure__189931222 gopurs_runtime.Value
var once_pure__189931222 sync.Once
func Get_pure__189931222() gopurs_runtime.Value {
	once_pure__189931222.Do(func() {
		cache_pure__189931222 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__189931222(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__189931222
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__1253336208 gopurs_runtime.Value
var once_pure__1253336208 sync.Once
func Get_pure__1253336208() gopurs_runtime.Value {
	once_pure__1253336208.Do(func() {
		cache_pure__1253336208 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1253336208(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__1253336208
}

var cache_apply__4203183626 gopurs_runtime.Value
var once_apply__4203183626 sync.Once
func Get_apply__4203183626() gopurs_runtime.Value {
	once_apply__4203183626.Do(func() {
		cache_apply__4203183626 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__4203183626(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__4203183626
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_apply__197474060 gopurs_runtime.Value
var once_apply__197474060 sync.Once
func Get_apply__197474060() gopurs_runtime.Value {
	once_apply__197474060.Do(func() {
		cache_apply__197474060 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__197474060(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__197474060
}

var cache_extend__1264481661 gopurs_runtime.Value
var once_extend__1264481661 sync.Once
func Get_extend__1264481661() gopurs_runtime.Value {
	once_extend__1264481661.Do(func() {
		cache_extend__1264481661 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extend__1264481661(gopurs_runtime.CoerceToStruct[pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_extend__1264481661
}

var cache_extend__4254185051 gopurs_runtime.Value
var once_extend__4254185051 sync.Once
func Get_extend__4254185051() gopurs_runtime.Value {
	once_extend__4254185051.Do(func() {
		cache_extend__4254185051 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extend__4254185051(v_0_box, v1_1_box)
})
	})
	return cache_extend__4254185051
}

var cache_compose__1987728071 gopurs_runtime.Value
var once_compose__1987728071 sync.Once
func Get_compose__1987728071() gopurs_runtime.Value {
	once_compose__1987728071.Do(func() {
		cache_compose__1987728071 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1987728071(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_compose__1987728071
}

var cache_compose__1555187646 gopurs_runtime.Value
var once_compose__1555187646 sync.Once
func Get_compose__1555187646() gopurs_runtime.Value {
	once_compose__1555187646.Do(func() {
		cache_compose__1555187646 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1555187646(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__1555187646
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

var cache_bifoldMap__4006734481 gopurs_runtime.Value
var once_bifoldMap__4006734481 sync.Once
func Get_bifoldMap__4006734481() gopurs_runtime.Value {
	once_bifoldMap__4006734481.Do(func() {
		cache_bifoldMap__4006734481 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMap__4006734481(gopurs_runtime.CoerceToStruct[pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldMap__4006734481
}

var cache_bifoldMap__1302573585 gopurs_runtime.Value
var once_bifoldMap__1302573585 sync.Once
func Get_bifoldMap__1302573585() gopurs_runtime.Value {
	once_bifoldMap__1302573585.Do(func() {
		cache_bifoldMap__1302573585 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMap__1302573585(gopurs_runtime.CoerceToStruct[pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldMap__1302573585
}

var cache_bifoldMap__1245832375 gopurs_runtime.Value
var once_bifoldMap__1245832375 sync.Once
func Get_bifoldMap__1245832375() gopurs_runtime.Value {
	once_bifoldMap__1245832375.Do(func() {
		cache_bifoldMap__1245832375 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMap__1245832375(gopurs_runtime.CoerceToStruct[pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldMap__1245832375
}

var cache_bifoldMap__3662584887 gopurs_runtime.Value
var once_bifoldMap__3662584887 sync.Once
func Get_bifoldMap__3662584887() gopurs_runtime.Value {
	once_bifoldMap__3662584887.Do(func() {
		cache_bifoldMap__3662584887 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMap__3662584887(gopurs_runtime.CoerceToStruct[pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldMap__3662584887
}

var cache_bifoldl__2116322576 gopurs_runtime.Value
var once_bifoldl__2116322576 sync.Once
func Get_bifoldl__2116322576() gopurs_runtime.Value {
	once_bifoldl__2116322576.Do(func() {
		cache_bifoldl__2116322576 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldl__2116322576(gopurs_runtime.CoerceToStruct[pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldl__2116322576
}

var cache_bifoldl__31590006 gopurs_runtime.Value
var once_bifoldl__31590006 sync.Once
func Get_bifoldl__31590006() gopurs_runtime.Value {
	once_bifoldl__31590006.Do(func() {
		cache_bifoldl__31590006 = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldl__31590006(v_0_box, v1_1_box, v2_2_box, v3_3_box)
})
	})
	return cache_bifoldl__31590006
}

var cache_bifoldr__2116322576 gopurs_runtime.Value
var once_bifoldr__2116322576 sync.Once
func Get_bifoldr__2116322576() gopurs_runtime.Value {
	once_bifoldr__2116322576.Do(func() {
		cache_bifoldr__2116322576 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldr__2116322576(gopurs_runtime.CoerceToStruct[pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldr__2116322576
}

var cache_bifoldr__31590006 gopurs_runtime.Value
var once_bifoldr__31590006 sync.Once
func Get_bifoldr__31590006() gopurs_runtime.Value {
	once_bifoldr__31590006.Do(func() {
		cache_bifoldr__31590006 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldr__31590006(x_0_box)
})
	})
	return cache_bifoldr__31590006
}

var cache_bifoldrDefault__1989667951 gopurs_runtime.Value
var once_bifoldrDefault__1989667951 sync.Once
func Get_bifoldrDefault__1989667951() gopurs_runtime.Value {
	once_bifoldrDefault__1989667951.Do(func() {
		cache_bifoldrDefault__1989667951 = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldrDefault__1989667951(gopurs_runtime.CoerceToStruct[pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, p_4_box)
})
	})
	return cache_bifoldrDefault__1989667951
}

var cache_bifoldrDefault__2116322576 gopurs_runtime.Value
var once_bifoldrDefault__2116322576 sync.Once
func Get_bifoldrDefault__2116322576() gopurs_runtime.Value {
	once_bifoldrDefault__2116322576.Do(func() {
		cache_bifoldrDefault__2116322576 = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldrDefault__2116322576(gopurs_runtime.CoerceToStruct[pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, p_4_box)
})
	})
	return cache_bifoldrDefault__2116322576
}

var cache_bifoldrDefault__31590006 gopurs_runtime.Value
var once_bifoldrDefault__31590006 sync.Once
func Get_bifoldrDefault__31590006() gopurs_runtime.Value {
	once_bifoldrDefault__31590006.Do(func() {
		cache_bifoldrDefault__31590006 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value, p_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldrDefault__31590006(f_0_box, g_1_box, z_2_box, p_3_box)
})
	})
	return cache_bifoldrDefault__31590006
}

var cache_bimap__4044928099 gopurs_runtime.Value
var once_bimap__4044928099 sync.Once
func Get_bimap__4044928099() gopurs_runtime.Value {
	once_bimap__4044928099.Do(func() {
		cache_bimap__4044928099 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap__4044928099(gopurs_runtime.CoerceToStruct[pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bimap__4044928099
}

var cache_bimap__132457202 gopurs_runtime.Value
var once_bimap__132457202 sync.Once
func Get_bimap__132457202() gopurs_runtime.Value {
	once_bimap__132457202.Do(func() {
		cache_bimap__132457202 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap__132457202(gopurs_runtime.CoerceToStruct[pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bimap__132457202
}

var cache_bimap__1727657434 gopurs_runtime.Value
var once_bimap__1727657434 sync.Once
func Get_bimap__1727657434() gopurs_runtime.Value {
	once_bimap__1727657434.Do(func() {
		cache_bimap__1727657434 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap__1727657434(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_bimap__1727657434
}

var cache_bitraverse__3884078439 gopurs_runtime.Value
var once_bitraverse__3884078439 sync.Once
func Get_bitraverse__3884078439() gopurs_runtime.Value {
	once_bitraverse__3884078439.Do(func() {
		cache_bitraverse__3884078439 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverse__3884078439(gopurs_runtime.CoerceToStruct[pkg_Data_Bitraversable.Constructor_Bitraversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bitraverse__3884078439
}

var cache_bitraverse__4064111983 gopurs_runtime.Value
var once_bitraverse__4064111983 sync.Once
func Get_bitraverse__4064111983() gopurs_runtime.Value {
	once_bitraverse__4064111983.Do(func() {
		cache_bitraverse__4064111983 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverse__4064111983(dictApplicative_0_box)
})
	})
	return cache_bitraverse__4064111983
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

var cache_eq__501078914 gopurs_runtime.Value
var once_eq__501078914 sync.Once
func Get_eq__501078914() gopurs_runtime.Value {
	once_eq__501078914.Do(func() {
		cache_eq__501078914 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__501078914(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__501078914
}

var cache_eq__3433516078 gopurs_runtime.Value
var once_eq__3433516078 sync.Once
func Get_eq__3433516078() gopurs_runtime.Value {
	once_eq__3433516078.Do(func() {
		cache_eq__3433516078 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__3433516078(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_eq__3433516078
}

var cache_foldMap__4098395794 gopurs_runtime.Value
var once_foldMap__4098395794 sync.Once
func Get_foldMap__4098395794() gopurs_runtime.Value {
	once_foldMap__4098395794.Do(func() {
		cache_foldMap__4098395794 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__4098395794(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap__4098395794
}

var cache_foldMap__2966595236 gopurs_runtime.Value
var once_foldMap__2966595236 sync.Once
func Get_foldMap__2966595236() gopurs_runtime.Value {
	once_foldMap__2966595236.Do(func() {
		cache_foldMap__2966595236 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__2966595236(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap__2966595236
}

var cache_foldMap__2247887508 gopurs_runtime.Value
var once_foldMap__2247887508 sync.Once
func Get_foldMap__2247887508() gopurs_runtime.Value {
	once_foldMap__2247887508.Do(func() {
		cache_foldMap__2247887508 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__2247887508(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap__2247887508
}

var cache_foldMap__1596329762 gopurs_runtime.Value
var once_foldMap__1596329762 sync.Once
func Get_foldMap__1596329762() gopurs_runtime.Value {
	once_foldMap__1596329762.Do(func() {
		cache_foldMap__1596329762 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__1596329762(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap__1596329762
}

var cache_foldableFreeMonoidTree__2832280077 gopurs_runtime.Value
var once_foldableFreeMonoidTree__2832280077 sync.Once
func Get_foldableFreeMonoidTree__2832280077() gopurs_runtime.Value {
	once_foldableFreeMonoidTree__2832280077.Do(func() {
		cache_foldableFreeMonoidTree__2832280077 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableFreeMonoidTree(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_0 gopurs_runtime.Value
go__go_1_2_0 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var acc_2_loop gopurs_runtime.Value = acc_2_loop_val
var lhs_3_loop gopurs_runtime.Value = lhs_3_loop_val
var rhs_4_loop gopurs_runtime.Value = rhs_4_loop_val
go__go_1_2_0:
for {
if false { continue go__go_1_2_0 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t7 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, acc_2, (*pkg_Data_Foldable.Constructor_Node[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0)
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_2_0
__t7 = gopurs_runtime.Value{}
goto end_branch_7
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2812549951) {
var __t5 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = rhs_4
continue go__go_1_2_0
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
var __t3 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1
continue go__go_1_2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value]{1, (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1, rhs_4})}
continue go__go_1_2_0
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t5 = __t3
}
end_branch_5:
__t7 = __t5
goto end_branch_7
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
var __t6 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
__t6 = acc_2
goto end_branch_6
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_2_0
__t6 = gopurs_runtime.Value{}
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}()
})
})
})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__go_1_2_0, a_2, b_3, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)})
})
})
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_8_1 gopurs_runtime.Value
go__go_1_8_1 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var acc_2_loop gopurs_runtime.Value = acc_2_loop_val
var lhs_3_loop gopurs_runtime.Value = lhs_3_loop_val
var rhs_4_loop gopurs_runtime.Value = rhs_4_loop_val
go__go_1_8_1:
for {
if false { continue go__go_1_8_1 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t13 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, (*pkg_Data_Foldable.Constructor_Node[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0, acc_2)
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_8_1
__t13 = gopurs_runtime.Value{}
goto end_branch_13
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2812549951) {
var __t11 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = lhs_3
rhs_4_loop = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_8_1
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
var __t9 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
rhs_4_loop = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_8_1
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value]{1, lhs_3, (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0})}
rhs_4_loop = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_8_1
__t9 = gopurs_runtime.Value{}
}
end_branch_9:
__t11 = __t9
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
var __t12 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
__t12 = acc_2
goto end_branch_12
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_8_1
__t12 = gopurs_runtime.Value{}
}
end_branch_12:
__t13 = __t12
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}
}()
})
})
})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__go_1_8_1, a_2, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}, b_3)
})
})
}))
	})
	return cache_foldableFreeMonoidTree__2832280077
}

var cache_foldl__2151204251 gopurs_runtime.Value
var once_foldl__2151204251 sync.Once
func Get_foldl__2151204251() gopurs_runtime.Value {
	once_foldl__2151204251.Do(func() {
		cache_foldl__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__2151204251
}

var cache_foldl__3288778237 gopurs_runtime.Value
var once_foldl__3288778237 sync.Once
func Get_foldl__3288778237() gopurs_runtime.Value {
	once_foldl__3288778237.Do(func() {
		cache_foldl__3288778237 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3288778237(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_foldl__3288778237
}

var cache_foldr__2111289130 gopurs_runtime.Value
var once_foldr__2111289130 sync.Once
func Get_foldr__2111289130() gopurs_runtime.Value {
	once_foldr__2111289130.Do(func() {
		cache_foldr__2111289130 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2111289130(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2111289130
}

var cache_foldr__2151204251 gopurs_runtime.Value
var once_foldr__2151204251 sync.Once
func Get_foldr__2151204251() gopurs_runtime.Value {
	once_foldr__2151204251.Do(func() {
		cache_foldr__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2151204251
}

var cache_foldr__1459781277 gopurs_runtime.Value
var once_foldr__1459781277 sync.Once
func Get_foldr__1459781277() gopurs_runtime.Value {
	once_foldr__1459781277.Do(func() {
		cache_foldr__1459781277 = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__1459781277(fn_0_box)
})
	})
	return cache_foldr__1459781277
}

var cache_foldr__3288778237 gopurs_runtime.Value
var once_foldr__3288778237 sync.Once
func Get_foldr__3288778237() gopurs_runtime.Value {
	once_foldr__3288778237.Do(func() {
		cache_foldr__3288778237 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__3288778237(x_0_box)
})
	})
	return cache_foldr__3288778237
}

var cache_foldrDefault__2858227716 gopurs_runtime.Value
var once_foldrDefault__2858227716 sync.Once
func Get_foldrDefault__2858227716() gopurs_runtime.Value {
	once_foldrDefault__2858227716.Do(func() {
		cache_foldrDefault__2858227716 = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value, u_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrDefault__2858227716(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), c_1_box, u_2_box, xs_3_box)
})
	})
	return cache_foldrDefault__2858227716
}

var cache_foldrDefault__2151204251 gopurs_runtime.Value
var once_foldrDefault__2151204251 sync.Once
func Get_foldrDefault__2151204251() gopurs_runtime.Value {
	once_foldrDefault__2151204251.Do(func() {
		cache_foldrDefault__2151204251 = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value, u_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrDefault__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), c_1_box, u_2_box, xs_3_box)
})
	})
	return cache_foldrDefault__2151204251
}

var cache_foldrDefault__3288778237 gopurs_runtime.Value
var once_foldrDefault__3288778237 sync.Once
func Get_foldrDefault__3288778237() gopurs_runtime.Value {
	once_foldrDefault__3288778237.Do(func() {
		cache_foldrDefault__3288778237 = gopurs_runtime.Func3(func(c_0_box gopurs_runtime.Value, u_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrDefault__3288778237(c_0_box, u_1_box, xs_2_box)
})
	})
	return cache_foldrDefault__3288778237
}

var cache_monoidFreeMonoidTree__2615096836 gopurs_runtime.Value
var once_monoidFreeMonoidTree__2615096836 sync.Once
func Get_monoidFreeMonoidTree__2615096836() gopurs_runtime.Value {
	once_monoidFreeMonoidTree__2615096836.Do(func() {
		cache_monoidFreeMonoidTree__2615096836 = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_semigroupFreeMonoidTree()
}), gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)})
	})
	return cache_monoidFreeMonoidTree__2615096836
}

var cache_semigroupFreeMonoidTree__2398658907 gopurs_runtime.Value
var once_semigroupFreeMonoidTree__2398658907 sync.Once
func Get_semigroupFreeMonoidTree__2398658907() gopurs_runtime.Value {
	once_semigroupFreeMonoidTree__2398658907.Do(func() {
		cache_semigroupFreeMonoidTree__2398658907 = gopurs_runtime.RecordDict1("append", pkg_Data_Foldable.Get_Append())
	})
	return cache_semigroupFreeMonoidTree__2398658907
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_const__2291477660 gopurs_runtime.Value
var once_const__2291477660 sync.Once
func Get_const__2291477660() gopurs_runtime.Value {
	once_const__2291477660.Do(func() {
		cache_const__2291477660 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__2291477660(a_0_box, v_1_box)
})
	})
	return cache_const__2291477660
}

var cache_map__2665381605 gopurs_runtime.Value
var once_map__2665381605 sync.Once
func Get_map__2665381605() gopurs_runtime.Value {
	once_map__2665381605.Do(func() {
		cache_map__2665381605 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2665381605(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2665381605
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__3634441076 gopurs_runtime.Value
var once_map__3634441076 sync.Once
func Get_map__3634441076() gopurs_runtime.Value {
	once_map__3634441076.Do(func() {
		cache_map__3634441076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3634441076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3634441076
}

var cache_map__3683879988 gopurs_runtime.Value
var once_map__3683879988 sync.Once
func Get_map__3683879988() gopurs_runtime.Value {
	once_map__3683879988.Do(func() {
		cache_map__3683879988 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3683879988(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3683879988
}

var cache_map__3663575028 gopurs_runtime.Value
var once_map__3663575028 sync.Once
func Get_map__3663575028() gopurs_runtime.Value {
	once_map__3663575028.Do(func() {
		cache_map__3663575028 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3663575028(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3663575028
}

var cache_map__2579103836 gopurs_runtime.Value
var once_map__2579103836 sync.Once
func Get_map__2579103836() gopurs_runtime.Value {
	once_map__2579103836.Do(func() {
		cache_map__2579103836 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2579103836(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_map__2579103836
}

var cache_mapFlipped__260821093 gopurs_runtime.Value
var once_mapFlipped__260821093 sync.Once
func Get_mapFlipped__260821093() gopurs_runtime.Value {
	once_mapFlipped__260821093.Do(func() {
		cache_mapFlipped__260821093 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__260821093(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__260821093
}

var cache_mapFlipped__4215217780 gopurs_runtime.Value
var once_mapFlipped__4215217780 sync.Once
func Get_mapFlipped__4215217780() gopurs_runtime.Value {
	once_mapFlipped__4215217780.Do(func() {
		cache_mapFlipped__4215217780 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__4215217780(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__4215217780
}

var cache_mapFlipped__742928244 gopurs_runtime.Value
var once_mapFlipped__742928244 sync.Once
func Get_mapFlipped__742928244() gopurs_runtime.Value {
	once_mapFlipped__742928244.Do(func() {
		cache_mapFlipped__742928244 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__742928244(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__742928244
}

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_conj__3676519832
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__3201284355(__eta0_0_box)
})
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_bifoldableInterval__2998510362 gopurs_runtime.Value
var once_bifoldableInterval__2998510362 sync.Once
func Get_bifoldableInterval__2998510362() gopurs_runtime.Value {
	once_bifoldableInterval__2998510362.Do(func() {
		cache_bifoldableInterval__2998510362 = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifoldableInterval(), "bifoldl"), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, m_5, gopurs_runtime.Apply(f_3, a_6))
})
}), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, m_5, gopurs_runtime.Apply(g_4, b_6))
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
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v1_1, v2_2, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1), (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
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
return gopurs_runtime.Apply2(pkg_Data_Bifoldable.Get_bifoldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]](Get_bifoldableInterval()))}, x_0)
}))
	})
	return cache_bifoldableInterval__2998510362
}

var cache_bifoldableInterval__740659151 gopurs_runtime.Value
var once_bifoldableInterval__740659151 sync.Once
func Get_bifoldableInterval__740659151() gopurs_runtime.Value {
	once_bifoldableInterval__740659151.Do(func() {
		cache_bifoldableInterval__740659151 = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifoldableInterval(), "bifoldl"), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, m_5, gopurs_runtime.Apply(f_3, a_6))
})
}), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, m_5, gopurs_runtime.Apply(g_4, b_6))
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
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v1_1, v2_2, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1), (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
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
return gopurs_runtime.Apply2(pkg_Data_Bifoldable.Get_bifoldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]](Get_bifoldableInterval()))}, x_0)
}))
	})
	return cache_bifoldableInterval__740659151
}

var cache_bifoldableRecurringInterval__4077340974 gopurs_runtime.Value
var once_bifoldableRecurringInterval__4077340974 sync.Once
func Get_bifoldableRecurringInterval__4077340974() gopurs_runtime.Value {
	once_bifoldableRecurringInterval__4077340974.Do(func() {
		cache_bifoldableRecurringInterval__4077340974 = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifoldableRecurringInterval(), "bifoldl"), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, m_5, gopurs_runtime.Apply(f_3, a_6))
})
}), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, m_5, gopurs_runtime.Apply(g_4, b_6))
})
}), mempty_2_1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifoldableInterval(), "bifoldl"), f_0, g_1, i_2)
_ = __local_var_3_2
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1)
})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_3 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifoldableInterval(), "bifoldr"), f_0, g_1, i_2)
_ = __local_var_3_3
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1)
})
})
})
}))
	})
	return cache_bifoldableRecurringInterval__4077340974
}

var cache_bifunctorInterval__3636391546 gopurs_runtime.Value
var once_bifunctorInterval__3636391546 sync.Once
func Get_bifunctorInterval__3636391546() gopurs_runtime.Value {
	once_bifunctorInterval__3636391546.Do(func() {
		cache_bifunctorInterval__3636391546 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v_0, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
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
	return cache_bifunctorInterval__3636391546
}

var cache_bifunctorInterval__3665086044 gopurs_runtime.Value
var once_bifunctorInterval__3665086044 sync.Once
func Get_bifunctorInterval__3665086044() gopurs_runtime.Value {
	once_bifunctorInterval__3665086044.Do(func() {
		cache_bifunctorInterval__3665086044 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v_0, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
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
	return cache_bifunctorInterval__3665086044
}

var cache_bifunctorRecurringInterval__261139004 gopurs_runtime.Value
var once_bifunctorRecurringInterval__261139004 sync.Once
func Get_bifunctorRecurringInterval__261139004() gopurs_runtime.Value {
	once_bifunctorRecurringInterval__261139004.Do(func() {
		cache_bifunctorRecurringInterval__261139004 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifunctorInterval(), "bimap"), f_0, g_1, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
})
})
}))
	})
	return cache_bifunctorRecurringInterval__261139004
}

var cache_bitraversableInterval__1397501818 gopurs_runtime.Value
var once_bitraversableInterval__1397501818 sync.Once
func Get_bitraversableInterval__1397501818() gopurs_runtime.Value {
	once_bitraversableInterval__1397501818.Do(func() {
		cache_bitraversableInterval__1397501818 = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bifoldableInterval()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bifunctorInterval()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bitraversableInterval(), "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, pkg_Data_Bitraversable.Get_identity(), pkg_Data_Bitraversable.Get_identity1())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_StartEnd(), gopurs_runtime.Apply(v1_4, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_DurationEnd(), gopurs_runtime.Apply(v_3, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_StartDuration(), gopurs_runtime.Apply(v1_4, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(Functor0_2_1.V0, Get_DurationOnly(), gopurs_runtime.Apply(v_3, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0))
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
	return cache_bitraversableInterval__1397501818
}

var cache_extendInterval__2367002404 gopurs_runtime.Value
var once_extendInterval__2367002404 sync.Once
func Get_extendInterval__2367002404() gopurs_runtime.Value {
	once_extendInterval__2367002404.Do(func() {
		cache_extendInterval__2367002404 = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorInterval()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Apply(v_0, v1_1)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(v_0, v1_1)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1), (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0})}
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
	return cache_extendInterval__2367002404
}

var cache_foldableInterval__3140210451 gopurs_runtime.Value
var once_foldableInterval__3140210451 sync.Once
func Get_foldableInterval__3140210451() gopurs_runtime.Value {
	once_foldableInterval__3140210451.Do(func() {
		cache_foldableInterval__3140210451 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableInterval(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, acc_4, gopurs_runtime.Apply(f_3, x_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
return gopurs_runtime.Apply2(pkg_Data_Foldable.Get_foldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](Get_foldableInterval()))}, x_0)
}))
	})
	return cache_foldableInterval__3140210451
}

var cache_foldableInterval__526261656 gopurs_runtime.Value
var once_foldableInterval__526261656 sync.Once
func Get_foldableInterval__526261656() gopurs_runtime.Value {
	once_foldableInterval__526261656.Do(func() {
		cache_foldableInterval__526261656 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableInterval(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, acc_4, gopurs_runtime.Apply(f_3, x_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
return gopurs_runtime.Apply2(pkg_Data_Foldable.Get_foldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](Get_foldableInterval()))}, x_0)
}))
	})
	return cache_foldableInterval__526261656
}

var cache_foldableRecurringInterval__598519513 gopurs_runtime.Value
var once_foldableRecurringInterval__598519513 sync.Once
func Get_foldableRecurringInterval__598519513() gopurs_runtime.Value {
	once_foldableRecurringInterval__598519513.Do(func() {
		cache_foldableRecurringInterval__598519513 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableRecurringInterval(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, acc_4, gopurs_runtime.Apply(f_3, x_5))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableInterval(), "foldl"), f_0, i_1)
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableInterval(), "foldr"), f_0, i_1)
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_3, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1)
})
})
}))
	})
	return cache_foldableRecurringInterval__598519513
}

var cache_functorInterval__3565473811 gopurs_runtime.Value
var once_functorInterval__3565473811 sync.Once
func Get_functorInterval__3565473811() gopurs_runtime.Value {
	once_functorInterval__3565473811.Do(func() {
		cache_functorInterval__3565473811 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_bifunctorInterval(), "bimap"), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})))
	})
	return cache_functorInterval__3565473811
}

var cache_functorInterval__3172181073 gopurs_runtime.Value
var once_functorInterval__3172181073 sync.Once
func Get_functorInterval__3172181073() gopurs_runtime.Value {
	once_functorInterval__3172181073.Do(func() {
		cache_functorInterval__3172181073 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_bifunctorInterval(), "bimap"), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})))
	})
	return cache_functorInterval__3172181073
}

var cache_functorRecurringInterval__1167854705 gopurs_runtime.Value
var once_functorRecurringInterval__1167854705 sync.Once
func Get_functorRecurringInterval__1167854705() gopurs_runtime.Value {
	once_functorRecurringInterval__1167854705.Do(func() {
		cache_functorRecurringInterval__1167854705 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorInterval(), "map"), f_0, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorRecurringInterval__1167854705
}

var cache_interval__413767002 gopurs_runtime.Value
var once_interval__413767002 sync.Once
func Get_interval__413767002() gopurs_runtime.Value {
	once_interval__413767002.Do(func() {
		cache_interval__413767002 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_interval__413767002(gopurs_runtime.CoerceToStruct[Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_interval__413767002
}

var cache_over__3140749367 gopurs_runtime.Value
var once_over__3140749367 sync.Once
func Get_over__3140749367() gopurs_runtime.Value {
	once_over__3140749367.Do(func() {
		cache_over__3140749367 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over__3140749367(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box))
})
	})
	return cache_over__3140749367
}

var cache_traversableInterval__1002691347 gopurs_runtime.Value
var once_traversableInterval__1002691347 sync.Once
func Get_traversableInterval__1002691347() gopurs_runtime.Value {
	once_traversableInterval__1002691347.Do(func() {
		cache_traversableInterval__1002691347 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableInterval()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorInterval()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableInterval(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, pkg_Data_Traversable.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 237113226) {
__t3 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_StartEnd(), gopurs_runtime.Apply(v_3, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1992629780) {
__t3 = gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Apply(Get_DurationEnd(), (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0), gopurs_runtime.Apply(v_3, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2020675835) {
__local_var_5_2 := (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_2
__t3 = gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, v2_6, __local_var_5_2})}
}), gopurs_runtime.Apply(v_3, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2281256335) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0})})
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
	return cache_traversableInterval__1002691347
}

var cache_unwrap__1971311275 gopurs_runtime.Value
var once_unwrap__1971311275 sync.Once
func Get_unwrap__1971311275() gopurs_runtime.Value {
	once_unwrap__1971311275.Do(func() {
		cache_unwrap__1971311275 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__1971311275(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__1971311275
}

var cache_unwrap__1763047240 gopurs_runtime.Value
var once_unwrap__1763047240 sync.Once
func Get_unwrap__1763047240() gopurs_runtime.Value {
	once_unwrap__1763047240.Do(func() {
		cache_unwrap__1763047240 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__1763047240(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__1763047240
}

var cache_unwrap__3267718003 gopurs_runtime.Value
var once_unwrap__3267718003 sync.Once
func Get_unwrap__3267718003() gopurs_runtime.Value {
	once_unwrap__3267718003.Do(func() {
		cache_unwrap__3267718003 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__3267718003(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__3267718003
}

var cache_unwrap__2928868755 gopurs_runtime.Value
var once_unwrap__2928868755 sync.Once
func Get_unwrap__2928868755() gopurs_runtime.Value {
	once_unwrap__2928868755.Do(func() {
		cache_unwrap__2928868755 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__2928868755(__eta0_0_box)
})
	})
	return cache_unwrap__2928868755
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

var cache_compare__738396984 gopurs_runtime.Value
var once_compare__738396984 sync.Once
func Get_compare__738396984() gopurs_runtime.Value {
	once_compare__738396984.Do(func() {
		cache_compare__738396984 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__738396984(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__738396984
}

var cache_compare__2740609364 gopurs_runtime.Value
var once_compare__2740609364 sync.Once
func Get_compare__2740609364() gopurs_runtime.Value {
	once_compare__2740609364.Do(func() {
		cache_compare__2740609364 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__2740609364(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_compare__2740609364
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
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

var cache_show__2929403666 gopurs_runtime.Value
var once_show__2929403666 sync.Once
func Get_show__2929403666() gopurs_runtime.Value {
	once_show__2929403666.Do(func() {
		cache_show__2929403666 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2929403666(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2929403666
}

var cache_show__3698026194 gopurs_runtime.Value
var once_show__3698026194 sync.Once
func Get_show__3698026194() gopurs_runtime.Value {
	once_show__3698026194.Do(func() {
		cache_show__3698026194 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__3698026194(__eta0_0_box)
})
	})
	return cache_show__3698026194
}

var cache_traverse__314957093 gopurs_runtime.Value
var once_traverse__314957093 sync.Once
func Get_traverse__314957093() gopurs_runtime.Value {
	once_traverse__314957093.Do(func() {
		cache_traverse__314957093 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__314957093(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse__314957093
}

var cache_traverse__4126651533 gopurs_runtime.Value
var once_traverse__4126651533 sync.Once
func Get_traverse__4126651533() gopurs_runtime.Value {
	once_traverse__4126651533.Do(func() {
		cache_traverse__4126651533 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__4126651533(dictApplicative_0_box)
})
	})
	return cache_traverse__4126651533
}

type Constructor_StartEnd[T_d any, T_a any] struct {
	Rc uint32
	V0 T_a
	V1 T_a
}


type Constructor_DurationEnd[T_d any, T_a any] struct {
	Rc uint32
	V0 T_d
	V1 T_a
}


type Constructor_StartDuration[T_d any, T_a any] struct {
	Rc uint32
	V0 T_a
	V1 T_d
}


type Constructor_DurationOnly[T_d any, T_a any] struct {
	Rc uint32
	V0 T_d
}


type Constructor_RecurringInterval[T_d any, T_a any] struct {
	Rc uint32
	V0 *pkg_Data_Maybe.Constructor_Just[int64]
	V1 gopurs_runtime.Value
}


func Call_showInterval(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(StartEnd "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(" "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())).StrVal())).StrVal())
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(DurationEnd "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(" "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())).StrVal())).StrVal())
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(StartDuration "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(" "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())).StrVal())).StrVal())
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2281256335) {
__t0 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(DurationOnly "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
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
}

func Call_showRecurringInterval(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
showInterval2_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](Call_showInterval(dictShow_0, dictShow1_1))
_ = showInterval2_2_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(RecurringInterval "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(Call_show__3698026194(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)}).StrVal()), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(" "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(showInterval2_2_0.V0, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())).StrVal())).StrVal())
}))
}

func Call_over(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Apply(Get_RecurringInterval(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0)}), gopurs_runtime.Apply(f_1, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1))
}

func Call_eqInterval(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = (Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0))).IntVal) != (0)
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
__t1 = (Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0))).IntVal) != (0)
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
__t2 = (Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0))).IntVal) != (0)
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
__t3 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)
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

func Call_eqRecurringInterval(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
eqInterval2_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](Call_eqInterval(dictEq_0, dictEq1_1))
_ = eqInterval2_2_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((Call_eq__3433516078(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0)}).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(eqInterval2_2_0.V0, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V1).IntVal) != (0))).IntVal) != (0))
})
}))
}

func Call_ordInterval(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqInterval1_1_0 := gopurs_runtime.Apply(Get_eqInterval(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqInterval1_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
eqInterval2_3_1 := gopurs_runtime.Apply(eqInterval1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{}))
_ = eqInterval2_3_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqInterval2_3_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (x_4.Type == 9 && x_4.IntVal == 237113226) {
var __t4 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 237113226) {
v_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
_ = v_6_2
var __t3 uint32
{
if (uint32(v_6_2.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(v_6_2.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
__t3 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 237113226) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 1992629780) {
var __t7 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 1992629780) {
v_6_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
_ = v_6_5
var __t6 uint32
{
if (uint32(v_6_5.IntVal) == 1527465420) {
__t6 = 1527465420
goto end_branch_6
} else {

}
}
{
if (uint32(v_6_5.IntVal) == 380165415) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
__t6 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
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
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t7), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 1992629780) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 2020675835) {
var __t10 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 2020675835) {
v_6_8 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
_ = v_6_8
var __t9 uint32
{
if (uint32(v_6_8.IntVal) == 1527465420) {
__t9 = 1527465420
goto end_branch_9
} else {

}
}
{
if (uint32(v_6_8.IntVal) == 380165415) {
__t9 = 380165415
goto end_branch_9
} else {

}
}
{
__t9 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
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
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t10), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 2020675835) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if ((x_4.Type == 9 && x_4.IntVal == 2281256335)) && ((y_5.Type == 9 && y_5.IntVal == 2281256335)) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t11.IntVal)), UnsafePtr: nil}
})
}))
})
}

func Call_ordRecurringInterval(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordInterval1_1_0 := Call_ordInterval(dictOrd_0)
_ = ordInterval1_1_0
eqRecurringInterval1_2_1 := gopurs_runtime.Apply(Get_eqRecurringInterval(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqRecurringInterval1_2_1
return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
ordInterval2_4_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordInterval1_1_0, dictOrd1_3))
_ = ordInterval2_4_2
eqRecurringInterval2_5_3 := gopurs_runtime.Apply(eqRecurringInterval1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{}))
_ = eqRecurringInterval2_5_3
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRecurringInterval2_5_3
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_7 gopurs_runtime.Value) gopurs_runtime.Value {
v_8_4 := Call_compare__2740609364(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V0)})
_ = v_8_4
var __t5 uint32
{
if (uint32(v_8_4.IntVal) == 1527465420) {
__t5 = 1527465420
goto end_branch_5
} else {

}
}
{
if (uint32(v_8_4.IntVal) == 380165415) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
__t5 = uint32(gopurs_runtime.Apply2(ordInterval2_4_2.V1, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(x_6.UnsafePtr).V1, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(y_7.UnsafePtr).V1).IntVal)
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
})
}))
})
}

func Call_pure__189931222(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1253336208(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__4203183626(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__197474060(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_extend__1264481661(dict_0_loop *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_extend__4254185051(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Apply(v_0, v1_1)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(v_0, v1_1)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1), (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_compose__1987728071(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__1555187646(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bifoldMap__4006734481(dict_0_loop *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bifoldMap__1302573585(dict_0_loop *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(dict_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](pkg_Data_Bifoldable.Get_monoidDual()))})
}

func Call_bifoldMap__1245832375(dict_0_loop *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bifoldMap__3662584887(dict_0_loop *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bifoldl__2116322576(dict_0_loop *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bifoldl__31590006(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value, v3_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var v3_3 gopurs_runtime.Value = v3_3_loop
_ = v3_3
var __t0 gopurs_runtime.Value
{
if (v3_3.Type == 9 && v3_3.IntVal == 237113226) {
__t0 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v1_1, v2_2, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 1992629780) {
__t0 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2020675835) {
__t0 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1), (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2281256335) {
__t0 = gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_bifoldr__2116322576(dict_0_loop *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_bifoldr__31590006(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(pkg_Data_Bifoldable.Get_bifoldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value]](Get_bifoldableInterval()))}, x_0)
}

func Call_bifoldrDefault__1989667951(dictBifoldable_0_loop *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 gopurs_runtime.Value = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(dictBifoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](pkg_Data_Bifoldable.Get_monoidEndo()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, x_5)
}), p_4, z_3)
}

func Call_bifoldrDefault__2116322576(dictBifoldable_0_loop *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *pkg_Data_Bifoldable.Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 gopurs_runtime.Value = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(dictBifoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](pkg_Data_Bifoldable.Get_monoidEndo()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, x_5)
}), p_4, z_3)
}

func Call_bifoldrDefault__31590006(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value, p_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var z_2 gopurs_runtime.Value = z_2_loop
_ = z_2
var p_3 gopurs_runtime.Value = p_3_loop
_ = p_3
return gopurs_runtime.Apply5(gopurs_runtime.RecordGet(Get_bifoldableInterval(), "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](pkg_Data_Bifoldable.Get_monoidEndo()))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, x_4)
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, x_4)
}), p_3, z_2)
}

func Call_bimap__4044928099(dict_0_loop *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bimap__132457202(dict_0_loop *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bimap__1727657434(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v_0, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_bitraverse__3884078439(dict_0_loop *pkg_Data_Bitraversable.Constructor_Bitraversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bitraversable.Constructor_Bitraversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_bitraverse__4064111983(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_StartEnd(), gopurs_runtime.Apply(v1_4, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_DurationEnd(), gopurs_runtime.Apply(v_3, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_StartDuration(), gopurs_runtime.Apply(v1_4, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(Functor0_2_1.V0, Get_DurationOnly(), gopurs_runtime.Apply(v_3, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0))
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
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__501078914(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__3433516078(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[*pkg_Data_Maybe.Constructor_Just[int64]]](Get_eqMaybe()).V0, __eta0_0, __eta1_1)
}

func Call_foldMap__4098395794(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldMap__2966595236(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(dict_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](pkg_Data_Foldable.Get_monoidFreeMonoidTree()))})
}

func Call_foldMap__2247887508(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldMap__1596329762(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__3288778237(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Apply2(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0), (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = v1_1
}
end_branch_0:
return __t0
}

func Call_foldr__2111289130(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__1459781277(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_2 gopurs_runtime.Value
go__go_1_0_2 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var acc_2_loop gopurs_runtime.Value = acc_2_loop_val
var lhs_3_loop gopurs_runtime.Value = lhs_3_loop_val
var rhs_4_loop gopurs_runtime.Value = rhs_4_loop_val
go__go_1_0_2:
for {
if false { continue go__go_1_0_2 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t5 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, (*pkg_Data_Foldable.Constructor_Node[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0, acc_2)
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_0_2
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2812549951) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = lhs_3
rhs_4_loop = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_0_2
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
var __t1 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
rhs_4_loop = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_0_2
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value]{1, lhs_3, (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0})}
rhs_4_loop = (*pkg_Data_Foldable.Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_0_2
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
__t3 = __t1
}
end_branch_3:
__t5 = __t3
goto end_branch_5
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
var __t4 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
__t4 = acc_2
goto end_branch_4
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_0_2
__t4 = gopurs_runtime.Value{}
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__go_1_0_2, a_2, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}, b_3)
})
})
}

func Call_foldr__3288778237(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(pkg_Data_Foldable.Get_foldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](Get_foldableInterval()))}, x_0)
}

func Call_foldrDefault__2858227716(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], c_1_loop gopurs_runtime.Value, u_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
var u_2 gopurs_runtime.Value = u_2_loop
_ = u_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableFreeMonoidTree(), "foldr"), c_1, u_2, gopurs_runtime.Apply3(dictFoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](pkg_Data_Foldable.Get_monoidFreeMonoidTree()))}, pkg_Data_Foldable.Get_Node(), xs_3))
}

func Call_foldrDefault__2151204251(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], c_1_loop gopurs_runtime.Value, u_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
var u_2 gopurs_runtime.Value = u_2_loop
_ = u_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableFreeMonoidTree(), "foldr"), c_1, u_2, gopurs_runtime.Apply3(dictFoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](pkg_Data_Foldable.Get_monoidFreeMonoidTree()))}, pkg_Data_Foldable.Get_Node(), xs_3))
}

func Call_foldrDefault__3288778237(c_0_loop gopurs_runtime.Value, u_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var c_0 gopurs_runtime.Value = c_0_loop
_ = c_0
var u_1 gopurs_runtime.Value = u_1_loop
_ = u_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableFreeMonoidTree(), "foldr"), c_0, u_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableInterval(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](pkg_Data_Foldable.Get_monoidFreeMonoidTree()))}, pkg_Data_Foldable.Get_Node(), xs_2))
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__2291477660(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_map__2665381605(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3634441076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3683879988(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3663575028(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2579103836(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifunctorInterval(), "bimap"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), __eta0_0, __eta1_1)
}

func Call_mapFlipped__260821093(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_mapFlipped__4215217780(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_mapFlipped__742928244(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_conj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) && ((__eta1_1.IntVal) != (0)))
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) || ((__eta1_1.IntVal) != (0)))
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__3201284355(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) != (true))
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_interval__413767002(v_0_loop *Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_over__3140749367(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Apply(Get_RecurringInterval(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0)}), gopurs_runtime.Apply(f_1, (*Constructor_RecurringInterval[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1))
}

func Call_unwrap__1971311275(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__1763047240(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__3267718003(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__2928868755(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return __eta0_0
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__738396984(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__2740609364(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_Maybe.Constructor_Just[int64]]](Get_ordMaybe()).V1, __eta0_0, __eta1_1)
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2929403666(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__3698026194(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[*pkg_Data_Maybe.Constructor_Just[int64]]](Get_showMaybe()).V0, __eta0_0)
}

func Call_traverse__314957093(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverse__4126651533(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 237113226) {
__t3 = gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_StartEnd(), gopurs_runtime.Apply(v_3, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1992629780) {
__t3 = gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Apply(Get_DurationEnd(), (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0), gopurs_runtime.Apply(v_3, (*Constructor_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2020675835) {
__local_var_5_2 := (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_2
__t3 = gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value]{1, v2_6, __local_var_5_2})}
}), gopurs_runtime.Apply(v_3, (*Constructor_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2281256335) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0})})
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
}



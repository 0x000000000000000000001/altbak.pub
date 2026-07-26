package Data_Interval

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Bifoldable "gopurs/output/Data.Bifoldable"
	pkg_Data_Bitraversable "gopurs/output/Data.Bitraversable"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	unsafe "unsafe"
)

var cache_show gopurs_runtime.Value
var once_show sync.Once
func Get_show() gopurs_runtime.Value {
	once_show.Do(func() {
		cache_show = gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_Maybe.Get_showMaybe(), pkg_Data_Show.Get_showInt()), "show")
	})
	return cache_show
}

var cache_compare gopurs_runtime.Value
var once_compare sync.Once
func Get_compare() gopurs_runtime.Value {
	once_compare.Do(func() {
		cache_compare = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (x_1.Type == 9 && x_1.IntVal == 3589588149) {
var __t2 gopurs_runtime.Value
{
if (y_2.Type == 9 && y_2.IntVal == 3589588149) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (y_2.Type == 9 && y_2.IntVal == 3589588149) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if ((x_1.Type == 9 && x_1.IntVal == 930809136)) && ((y_2.Type == 9 && y_2.IntVal == 930809136)) {
__t1 = gopurs_runtime.Apply2(__local_var_0_0, (*pkg_Data_Maybe.Constructor_Just)(x_1.UnsafePtr).V0, (*pkg_Data_Maybe.Constructor_Just)(y_2.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
}()
	})
	return cache_compare
}

var cache_StartEnd gopurs_runtime.Value
var once_StartEnd sync.Once
func Get_StartEnd() gopurs_runtime.Value {
	once_StartEnd.Do(func() {
		cache_StartEnd = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_StartEnd{value0, value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_DurationEnd{value0, value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration{value0, value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly{value0})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_RecurringInterval{value0, value1})}
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
return Call_over(dictFunctor_0_box, f_1_box, v_2_box)
})
	})
	return cache_over
}

var cache_foldableInterval gopurs_runtime.Value
var once_foldableInterval sync.Once
func Get_foldableInterval() gopurs_runtime.Value {
	once_foldableInterval.Do(func() {
		cache_foldableInterval = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableInterval(), "foldl"), gopurs_runtime.Func2(func(acc_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), acc_3, gopurs_runtime.Apply(f_2, x_4))
}), mempty_1_0)
})
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t1 = gopurs_runtime.Apply2(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_StartEnd)(v2_2.UnsafePtr).V0), (*Constructor_StartEnd)(v2_2.UnsafePtr).V1)
goto end_branch_1
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t1 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_DurationEnd)(v2_2.UnsafePtr).V1)
goto end_branch_1
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t1 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_StartDuration)(v2_2.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = v1_1
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Foldable.Get_foldrDefault(), Get_foldableInterval(), x_0)
}))
	})
	return cache_foldableInterval
}

var cache_foldableRecurringInterval gopurs_runtime.Value
var once_foldableRecurringInterval sync.Once
func Get_foldableRecurringInterval() gopurs_runtime.Value {
	once_foldableRecurringInterval.Do(func() {
		cache_foldableRecurringInterval = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableRecurringInterval(), "foldl"), gopurs_runtime.Func2(func(acc_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), acc_3, gopurs_runtime.Apply(f_2, x_4))
}), mempty_1_0)
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, i_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableInterval(), "foldl"), f_0, i_1)
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, (*Constructor_RecurringInterval)(x_3.UnsafePtr).V1)
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, i_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableInterval(), "foldr"), f_0, i_1)
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, (*Constructor_RecurringInterval)(x_3.UnsafePtr).V1)
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
		cache_bifunctorInterval = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_StartEnd{gopurs_runtime.Apply(v1_1, (*Constructor_StartEnd)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_StartEnd)(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_DurationEnd{gopurs_runtime.Apply(v_0, (*Constructor_DurationEnd)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v1_1, (*Constructor_DurationEnd)(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration{gopurs_runtime.Apply(v1_1, (*Constructor_StartDuration)(v2_2.UnsafePtr).V0), gopurs_runtime.Apply(v_0, (*Constructor_StartDuration)(v2_2.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly{gopurs_runtime.Apply(v_0, (*Constructor_DurationOnly)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
	})
	return cache_bifunctorInterval
}

var cache_bifunctorRecurringInterval gopurs_runtime.Value
var once_bifunctorRecurringInterval sync.Once
func Get_bifunctorRecurringInterval() gopurs_runtime.Value {
	once_bifunctorRecurringInterval.Do(func() {
		cache_bifunctorRecurringInterval = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_RecurringInterval{(*Constructor_RecurringInterval)(v_2.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifunctorInterval(), "bimap"), f_0, g_1, (*Constructor_RecurringInterval)(v_2.UnsafePtr).V1)})}
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
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 237113226) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 237113226, UnsafePtr: unsafe.Pointer(&Constructor_StartEnd{gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Apply(v_0, v1_1)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1992629780) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1992629780, UnsafePtr: unsafe.Pointer(&Constructor_DurationEnd{(*Constructor_DurationEnd)(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(v_0, v1_1)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2020675835) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration{gopurs_runtime.Apply(v_0, v1_1), (*Constructor_StartDuration)(v1_1.UnsafePtr).V1})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2281256335) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly{(*Constructor_DurationOnly)(v1_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
	})
	return cache_extendInterval
}

var cache_functorRecurringInterval gopurs_runtime.Value
var once_functorRecurringInterval sync.Once
func Get_functorRecurringInterval() gopurs_runtime.Value {
	once_functorRecurringInterval.Do(func() {
		cache_functorRecurringInterval = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_RecurringInterval{(*Constructor_RecurringInterval)(v_1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorInterval(), "map"), f_0, (*Constructor_RecurringInterval)(v_1.UnsafePtr).V1)})}
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
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, v_1)
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 2355637979, UnsafePtr: unsafe.Pointer(&Constructor_RecurringInterval{(*Constructor_RecurringInterval)(v_1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_extendInterval(), "extend"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_0
}), (*Constructor_RecurringInterval)(v_1.UnsafePtr).V1)})}
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableInterval(), "traverse"), dictApplicative_0, pkg_Data_Traversable.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_2_1
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), Get_StartEnd(), gopurs_runtime.Apply(v_3, (*Constructor_StartEnd)(v1_4.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_StartEnd)(v1_4.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.Apply(Get_DurationEnd(), (*Constructor_DurationEnd)(v1_4.UnsafePtr).V0), gopurs_runtime.Apply(v_3, (*Constructor_DurationEnd)(v1_4.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2020675835) {
__local_var_5_3 := (*Constructor_StartDuration)(v1_4.UnsafePtr).V1
_ = __local_var_5_3
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2020675835, UnsafePtr: unsafe.Pointer(&Constructor_StartDuration{v2_6, __local_var_5_3})}
}), gopurs_runtime.Apply(v_3, (*Constructor_StartDuration)(v1_4.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2281256335, UnsafePtr: unsafe.Pointer(&Constructor_DurationOnly{(*Constructor_DurationOnly)(v1_4.UnsafePtr).V0})})
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableRecurringInterval(), "traverse"), dictApplicative_0, pkg_Data_Traversable.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
over1_1_0 := gopurs_runtime.Apply(Get_over(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = over1_1_0
traverse1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableInterval(), "traverse"), dictApplicative_0)
_ = traverse1_2_1
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(over1_1_0, gopurs_runtime.Apply(traverse1_2_1, f_3), i_4)
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
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifoldableInterval(), "bifoldl"), gopurs_runtime.Func2(func(m_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), m_5, gopurs_runtime.Apply(f_3, a_6))
}), gopurs_runtime.Func2(func(m_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), m_5, gopurs_runtime.Apply(g_4, b_6))
}), mempty_2_1)
})
}), gopurs_runtime.Func4(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value, v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v3_3.Type == 9 && v3_3.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v1_1, v2_2, (*Constructor_StartEnd)(v3_3.UnsafePtr).V0), (*Constructor_StartEnd)(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_DurationEnd)(v3_3.UnsafePtr).V0), (*Constructor_DurationEnd)(v3_3.UnsafePtr).V1)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_StartDuration)(v3_3.UnsafePtr).V1), (*Constructor_StartDuration)(v3_3.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_DurationOnly)(v3_3.UnsafePtr).V0)
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
return gopurs_runtime.Apply2(pkg_Data_Bifoldable.Get_bifoldrDefault(), Get_bifoldableInterval(), x_0)
}))
	})
	return cache_bifoldableInterval
}

var cache_bifoldableRecurringInterval gopurs_runtime.Value
var once_bifoldableRecurringInterval sync.Once
func Get_bifoldableRecurringInterval() gopurs_runtime.Value {
	once_bifoldableRecurringInterval.Do(func() {
		cache_bifoldableRecurringInterval = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifoldableRecurringInterval(), "bifoldl"), gopurs_runtime.Func2(func(m_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), m_5, gopurs_runtime.Apply(f_3, a_6))
}), gopurs_runtime.Func2(func(m_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), m_5, gopurs_runtime.Apply(g_4, b_6))
}), mempty_2_1)
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, i_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifoldableInterval(), "bifoldl"), f_0, g_1, i_2)
_ = __local_var_3_2
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, (*Constructor_RecurringInterval)(x_4.UnsafePtr).V1)
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, i_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_3 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bifoldableInterval(), "bifoldr"), f_0, g_1, i_2)
_ = __local_var_3_3
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, (*Constructor_RecurringInterval)(x_4.UnsafePtr).V1)
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
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bitraversableInterval(), "bitraverse"), dictApplicative_0, pkg_Data_Bitraversable.Get_identity(), pkg_Data_Bitraversable.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), Get_StartEnd(), gopurs_runtime.Apply(v1_4, (*Constructor_StartEnd)(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_StartEnd)(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), Get_DurationEnd(), gopurs_runtime.Apply(v_3, (*Constructor_DurationEnd)(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_DurationEnd)(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), Get_StartDuration(), gopurs_runtime.Apply(v1_4, (*Constructor_StartDuration)(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_StartDuration)(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), Get_DurationOnly(), gopurs_runtime.Apply(v_3, (*Constructor_DurationOnly)(v2_5.UnsafePtr).V0))
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
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_bitraversableRecurringInterval(), "bitraverse"), dictApplicative_0, pkg_Data_Bitraversable.Get_identity(), pkg_Data_Bitraversable.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
over1_1_0 := gopurs_runtime.Apply(Get_over(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = over1_1_0
bitraverse1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_bitraversableInterval(), "bitraverse"), dictApplicative_0)
_ = bitraverse1_2_1
return gopurs_runtime.Func3(func(l_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(over1_1_0, gopurs_runtime.Apply2(bitraverse1_2_1, l_3, r_4), i_5)
})
}))
	})
	return cache_bitraversableRecurringInterval
}

type Constructor_StartEnd struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_DurationEnd struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_StartDuration struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_DurationOnly struct {
	V0 gopurs_runtime.Value
}


type Constructor_RecurringInterval struct {
	V0 gopurs_runtime.Value
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
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(StartEnd "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictShow1_1.UnsafePtr)).V0, (*Constructor_StartEnd)(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictShow1_1.UnsafePtr)).V0, (*Constructor_StartEnd)(v_2.UnsafePtr).V1), gopurs_runtime.Str(")")))))
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(DurationEnd "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictShow_0.UnsafePtr)).V0, (*Constructor_DurationEnd)(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictShow1_1.UnsafePtr)).V0, (*Constructor_DurationEnd)(v_2.UnsafePtr).V1), gopurs_runtime.Str(")")))))
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(StartDuration "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictShow1_1.UnsafePtr)).V0, (*Constructor_StartDuration)(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictShow_0.UnsafePtr)).V0, (*Constructor_StartDuration)(v_2.UnsafePtr).V1), gopurs_runtime.Str(")")))))
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2281256335) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(DurationOnly "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictShow_0.UnsafePtr)).V0, (*Constructor_DurationOnly)(v_2.UnsafePtr).V0), gopurs_runtime.Str(")")))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
}

func Call_showRecurringInterval(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(RecurringInterval "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(Get_show(), (*Constructor_RecurringInterval)(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_showInterval(dictShow_0, dictShow1_1), "show"), (*Constructor_RecurringInterval)(v_2.UnsafePtr).V1), gopurs_runtime.Str(")")))))
}))
}

func Call_over(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, gopurs_runtime.Apply(Get_RecurringInterval(), (*Constructor_RecurringInterval)(v_2.UnsafePtr).V0), gopurs_runtime.Apply(f_1, (*Constructor_RecurringInterval)(v_2.UnsafePtr).V1))
}

func Call_eqInterval(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_2.Type == 9 && x_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Bool(((y_3.Type == 9 && y_3.IntVal == 237113226)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictEq1_1.UnsafePtr)).V0, (*Constructor_StartEnd)(x_2.UnsafePtr).V0, (*Constructor_StartEnd)(y_3.UnsafePtr).V0), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictEq1_1.UnsafePtr)).V0, (*Constructor_StartEnd)(x_2.UnsafePtr).V1, (*Constructor_StartEnd)(y_3.UnsafePtr).V1)).IntVal) != (0)))
goto end_branch_0
} else {

}
}
{
if (x_2.Type == 9 && x_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Bool(((y_3.Type == 9 && y_3.IntVal == 1992629780)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictEq_0.UnsafePtr)).V0, (*Constructor_DurationEnd)(x_2.UnsafePtr).V0, (*Constructor_DurationEnd)(y_3.UnsafePtr).V0), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictEq1_1.UnsafePtr)).V0, (*Constructor_DurationEnd)(x_2.UnsafePtr).V1, (*Constructor_DurationEnd)(y_3.UnsafePtr).V1)).IntVal) != (0)))
goto end_branch_0
} else {

}
}
{
if (x_2.Type == 9 && x_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Bool(((y_3.Type == 9 && y_3.IntVal == 2020675835)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictEq1_1.UnsafePtr)).V0, (*Constructor_StartDuration)(x_2.UnsafePtr).V0, (*Constructor_StartDuration)(y_3.UnsafePtr).V0), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictEq_0.UnsafePtr)).V0, (*Constructor_StartDuration)(x_2.UnsafePtr).V1, (*Constructor_StartDuration)(y_3.UnsafePtr).V1)).IntVal) != (0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(((x_2.Type == 9 && x_2.IntVal == 2281256335)) && (((y_3.Type == 9 && y_3.IntVal == 2281256335)) && ((gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictEq_0.UnsafePtr)).V0, (*Constructor_DurationOnly)(x_2.UnsafePtr).V0, (*Constructor_DurationOnly)(y_3.UnsafePtr).V0).IntVal) != (0))))
}
end_branch_0:
return __t0
}))
}

func Call_eqRecurringInterval(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_RecurringInterval)(x_2.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3589588149) {
var __t_tag_5 gopurs_runtime.Value = (*Constructor_RecurringInterval)(y_3.UnsafePtr).V0
__t3 = gopurs_runtime.Bool((__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3589588149))
goto end_branch_3
} else {

}
}
{
var __t_tag_0 gopurs_runtime.Value = (*Constructor_RecurringInterval)(x_2.UnsafePtr).V0
var __t_and_2 bool = false
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 930809136) {

var __t_tag_1 gopurs_runtime.Value = (*Constructor_RecurringInterval)(y_3.UnsafePtr).V0
__t_and_2 = ((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 930809136)) && (((*pkg_Data_Maybe.Constructor_Just)((*Constructor_RecurringInterval)(x_2.UnsafePtr).V0.UnsafePtr).V0.IntVal) == ((*pkg_Data_Maybe.Constructor_Just)((*Constructor_RecurringInterval)(y_3.UnsafePtr).V0.UnsafePtr).V0.IntVal))
}
__t3 = gopurs_runtime.Bool(__t_and_2)
}
end_branch_3:
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), __t3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_eqInterval(dictEq_0, dictEq1_1), "eq"), (*Constructor_RecurringInterval)(x_2.UnsafePtr).V1, (*Constructor_RecurringInterval)(y_3.UnsafePtr).V1))
}))
}

func Call_ordInterval(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqInterval1_1_0 := gopurs_runtime.Apply(Get_eqInterval(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = eqInterval1_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
eqInterval2_3_1 := gopurs_runtime.Apply(eqInterval1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{}))
_ = eqInterval2_3_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqInterval2_3_1
}), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (x_4.Type == 9 && x_4.IntVal == 237113226) {
var __t3 gopurs_runtime.Value
{
if (y_5.Type == 9 && y_5.IntVal == 237113226) {
v_6_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_StartEnd)(x_4.UnsafePtr).V0, (*Constructor_StartEnd)(y_5.UnsafePtr).V0)
_ = v_6_4
var __t5 gopurs_runtime.Value
{
if (v_6_4.Type == 9 && v_6_4.IntVal == 1527465420) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
if (v_6_4.Type == 9 && v_6_4.IntVal == 380165415) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_StartEnd)(x_4.UnsafePtr).V1, (*Constructor_StartEnd)(y_5.UnsafePtr).V1)
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 237113226) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 1992629780) {
var __t6 gopurs_runtime.Value
{
if (y_5.Type == 9 && y_5.IntVal == 1992629780) {
v_6_7 := gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, (*Constructor_DurationEnd)(x_4.UnsafePtr).V0, (*Constructor_DurationEnd)(y_5.UnsafePtr).V0)
_ = v_6_7
var __t8 gopurs_runtime.Value
{
if (v_6_7.Type == 9 && v_6_7.IntVal == 1527465420) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_8
} else {

}
}
{
if (v_6_7.Type == 9 && v_6_7.IntVal == 380165415) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_DurationEnd)(x_4.UnsafePtr).V1, (*Constructor_DurationEnd)(y_5.UnsafePtr).V1)
}
end_branch_8:
__t6 = __t8
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_6:
__t2 = __t6
goto end_branch_2
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 1992629780) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 2020675835) {
var __t9 gopurs_runtime.Value
{
if (y_5.Type == 9 && y_5.IntVal == 2020675835) {
v_6_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_StartDuration)(x_4.UnsafePtr).V0, (*Constructor_StartDuration)(y_5.UnsafePtr).V0)
_ = v_6_10
var __t11 gopurs_runtime.Value
{
if (v_6_10.Type == 9 && v_6_10.IntVal == 1527465420) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
if (v_6_10.Type == 9 && v_6_10.IntVal == 380165415) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, (*Constructor_StartDuration)(x_4.UnsafePtr).V1, (*Constructor_StartDuration)(y_5.UnsafePtr).V1)
}
end_branch_11:
__t9 = __t11
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_9:
__t2 = __t9
goto end_branch_2
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 2020675835) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
if ((x_4.Type == 9 && x_4.IntVal == 2281256335)) && ((y_5.Type == 9 && y_5.IntVal == 2281256335)) {
__t2 = gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictOrd_0.UnsafePtr)).V0, (*Constructor_DurationOnly)(x_4.UnsafePtr).V0, (*Constructor_DurationOnly)(y_5.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
})
}

func Call_ordRecurringInterval(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordInterval1_1_0 := gopurs_runtime.Apply(Get_ordInterval(), dictOrd_0)
_ = ordInterval1_1_0
eqRecurringInterval1_2_1 := gopurs_runtime.Apply(Get_eqRecurringInterval(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = eqRecurringInterval1_2_1
return gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
eqRecurringInterval2_4_2 := gopurs_runtime.Apply(eqRecurringInterval1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{}))
_ = eqRecurringInterval2_4_2
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRecurringInterval2_4_2
}), gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
v_7_3 := gopurs_runtime.Apply2(Get_compare(), (*Constructor_RecurringInterval)(x_5.UnsafePtr).V0, (*Constructor_RecurringInterval)(y_6.UnsafePtr).V0)
_ = v_7_3
var __t4 gopurs_runtime.Value
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordInterval1_1_0, dictOrd1_3), "compare"), (*Constructor_RecurringInterval)(x_5.UnsafePtr).V1, (*Constructor_RecurringInterval)(y_6.UnsafePtr).V1)
}
end_branch_4:
return __t4
}))
})
}



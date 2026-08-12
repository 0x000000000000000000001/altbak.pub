package Effect

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_monadEffect gopurs_runtime.Value
var once_monadEffect sync.Once
func Get_monadEffect() gopurs_runtime.Value {
	once_monadEffect.Do(func() {
		cache_monadEffect = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindEffect()
}))
	})
	return cache_monadEffect
}

var cache_bindEffect gopurs_runtime.Value
var once_bindEffect sync.Once
func Get_bindEffect() gopurs_runtime.Value {
	once_bindEffect.Do(func() {
		cache_bindEffect = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_bindE())
	})
	return cache_bindEffect
}

var cache_applyEffect gopurs_runtime.Value
var once_applyEffect sync.Once
func Get_applyEffect() gopurs_runtime.Value {
	once_applyEffect.Do(func() {
		cache_applyEffect = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect
}

var cache_applicativeEffect gopurs_runtime.Value
var once_applicativeEffect sync.Once
func Get_applicativeEffect() gopurs_runtime.Value {
	once_applicativeEffect.Do(func() {
		cache_applicativeEffect = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_pureE())
	})
	return cache_applicativeEffect
}

var cache_functorEffect gopurs_runtime.Value
var once_functorEffect sync.Once
func Get_functorEffect() gopurs_runtime.Value {
	once_functorEffect.Do(func() {
		cache_functorEffect = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect
}

var cache_semigroupEffect gopurs_runtime.Value
var once_semigroupEffect sync.Once
func Get_semigroupEffect() gopurs_runtime.Value {
	once_semigroupEffect.Do(func() {
		cache_semigroupEffect = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupEffect(dictSemigroup_0_box)
})
	})
	return cache_semigroupEffect
}

var cache_monoidEffect gopurs_runtime.Value
var once_monoidEffect sync.Once
func Get_monoidEffect() gopurs_runtime.Value {
	once_monoidEffect.Do(func() {
		cache_monoidEffect = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidEffect(dictMonoid_0_box)
})
	})
	return cache_monoidEffect
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

var cache_apply__4277237866 gopurs_runtime.Value
var once_apply__4277237866 sync.Once
func Get_apply__4277237866() gopurs_runtime.Value {
	once_apply__4277237866.Do(func() {
		cache_apply__4277237866 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__4277237866(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__4277237866
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

var cache_lift2__2762258480 gopurs_runtime.Value
var once_lift2__2762258480 sync.Once
func Get_lift2__2762258480() gopurs_runtime.Value {
	once_lift2__2762258480.Do(func() {
		cache_lift2__2762258480 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2762258480(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2762258480
}

var cache_lift2__1465919478 gopurs_runtime.Value
var once_lift2__1465919478 sync.Once
func Get_lift2__1465919478() gopurs_runtime.Value {
	once_lift2__1465919478.Do(func() {
		cache_lift2__1465919478 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__1465919478(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_lift2__1465919478
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

var cache_map__4224113637 gopurs_runtime.Value
var once_map__4224113637 sync.Once
func Get_map__4224113637() gopurs_runtime.Value {
	once_map__4224113637.Do(func() {
		cache_map__4224113637 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__4224113637(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__4224113637
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

var cache_map__1483545076 gopurs_runtime.Value
var once_map__1483545076 sync.Once
func Get_map__1483545076() gopurs_runtime.Value {
	once_map__1483545076.Do(func() {
		cache_map__1483545076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1483545076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1483545076
}

var cache_mempty__2312420373 gopurs_runtime.Value
var once_mempty__2312420373 sync.Once
func Get_mempty__2312420373() gopurs_runtime.Value {
	once_mempty__2312420373.Do(func() {
		cache_mempty__2312420373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty__2312420373(dict_0_box)
})
	})
	return cache_mempty__2312420373
}

var cache_applicativeEffect__1969567048 gopurs_runtime.Value
var once_applicativeEffect__1969567048 sync.Once
func Get_applicativeEffect__1969567048() gopurs_runtime.Value {
	once_applicativeEffect__1969567048.Do(func() {
		cache_applicativeEffect__1969567048 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_pureE())
	})
	return cache_applicativeEffect__1969567048
}

var cache_applyEffect__1723132130 gopurs_runtime.Value
var once_applyEffect__1723132130 sync.Once
func Get_applyEffect__1723132130() gopurs_runtime.Value {
	once_applyEffect__1723132130.Do(func() {
		cache_applyEffect__1723132130 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect__1723132130
}

var cache_applyEffect__2014400020 gopurs_runtime.Value
var once_applyEffect__2014400020 sync.Once
func Get_applyEffect__2014400020() gopurs_runtime.Value {
	once_applyEffect__2014400020.Do(func() {
		cache_applyEffect__2014400020 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect__2014400020
}

var cache_bindEffect__3856311079 gopurs_runtime.Value
var once_bindEffect__3856311079 sync.Once
func Get_bindEffect__3856311079() gopurs_runtime.Value {
	once_bindEffect__3856311079.Do(func() {
		cache_bindEffect__3856311079 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}), Get_bindE())
	})
	return cache_bindEffect__3856311079
}

var cache_functorEffect__3107547953 gopurs_runtime.Value
var once_functorEffect__3107547953 sync.Once
func Get_functorEffect__3107547953() gopurs_runtime.Value {
	once_functorEffect__3107547953.Do(func() {
		cache_functorEffect__3107547953 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect__3107547953
}

func Call_semigroupEffect(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applyEffect(), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
__local_var_2_1 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_2_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
Bind1_5_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_2
Applicative0_6_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_3
return gopurs_runtime.Apply2(Bind1_5_2.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, __local_var_2_1, a_3), gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_5_2.V1, b_4, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_6_3.V1, gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}

func Call_monoidEffect(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
semigroupEffect1_1_0 := Call_semigroupEffect(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupEffect1_1_0
__local_var_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffect1_1_0
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}))
}

func Call_apply__4203183626(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__4277237866(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_lift2__2762258480(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__1465919478(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
Bind1_3_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_0
Applicative0_4_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_1
return gopurs_runtime.Apply2(Bind1_3_0.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorEffect(), "map"), __eta0_0, __eta1_1), gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_3_0.V1, __eta2_2, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_4_1.V1, gopurs_runtime.Apply(f_prime_5, a_prime_6))
}))
}))
}

func Call_map__2665381605(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__4224113637(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1483545076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Get_bindE() gopurs_runtime.Value {
	return _Gopurs_BindE
}

func Get_forE() gopurs_runtime.Value {
	return _Gopurs_ForE
}

func Get_foreachE() gopurs_runtime.Value {
	return _Gopurs_ForeachE
}

func Get_pureE() gopurs_runtime.Value {
	return _Gopurs_PureE
}

func Get_untilE() gopurs_runtime.Value {
	return _Gopurs_UntilE
}

func Get_whileE() gopurs_runtime.Value {
	return _Gopurs_WhileE
}

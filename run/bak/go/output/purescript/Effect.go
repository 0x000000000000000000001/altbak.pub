package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Effect_monadEffect gopurs_runtime.Value
var once_Effect_monadEffect sync.Once
func Get_Effect_monadEffect() gopurs_runtime.Value {
	once_Effect_monadEffect.Do(func() {
		cache_Effect_monadEffect = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_bindEffect()
}))
	})
	return cache_Effect_monadEffect
}

var cache_Effect_bindEffect gopurs_runtime.Value
var once_Effect_bindEffect sync.Once
func Get_Effect_bindEffect() gopurs_runtime.Value {
	once_Effect_bindEffect.Do(func() {
		cache_Effect_bindEffect = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
	})
	return cache_Effect_bindEffect
}

var cache_Effect_applyEffect gopurs_runtime.Value
var once_Effect_applyEffect sync.Once
func Get_Effect_applyEffect() gopurs_runtime.Value {
	once_Effect_applyEffect.Do(func() {
		cache_Effect_applyEffect = func() gopurs_runtime.Value {
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_Effect_applyEffect
}

var cache_Effect_applicativeEffect gopurs_runtime.Value
var once_Effect_applicativeEffect sync.Once
func Get_Effect_applicativeEffect() gopurs_runtime.Value {
	once_Effect_applicativeEffect.Do(func() {
		cache_Effect_applicativeEffect = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_pureE())
	})
	return cache_Effect_applicativeEffect
}

var cache_Effect_functorEffect gopurs_runtime.Value
var once_Effect_functorEffect sync.Once
func Get_Effect_functorEffect() gopurs_runtime.Value {
	once_Effect_functorEffect.Do(func() {
		cache_Effect_functorEffect = func() gopurs_runtime.Value {
// TAST (Let): Apply0_0_0 -> *Constructor_Control_Apply_Apply
Apply0_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_0_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_Effect_functorEffect
}

var cache_Effect_semigroupEffect gopurs_runtime.Value
var once_Effect_semigroupEffect sync.Once
func Get_Effect_semigroupEffect() gopurs_runtime.Value {
	once_Effect_semigroupEffect.Do(func() {
		cache_Effect_semigroupEffect = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_semigroupEffect(dictSemigroup_0_box)
})
	})
	return cache_Effect_semigroupEffect
}

var cache_Effect_monoidEffect gopurs_runtime.Value
var once_Effect_monoidEffect sync.Once
func Get_Effect_monoidEffect() gopurs_runtime.Value {
	once_Effect_monoidEffect.Do(func() {
		cache_Effect_monoidEffect = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_monoidEffect(dictMonoid_0_box)
})
	})
	return cache_Effect_monoidEffect
}

var cache_Effect_applicativeEffect__284161122 gopurs_runtime.Value
var once_Effect_applicativeEffect__284161122 sync.Once
func Get_Effect_applicativeEffect__284161122() gopurs_runtime.Value {
	once_Effect_applicativeEffect__284161122.Do(func() {
		cache_Effect_applicativeEffect__284161122 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_pureE())
	})
	return cache_Effect_applicativeEffect__284161122
}

var cache_Effect_applicativeEffect__1969567048 gopurs_runtime.Value
var once_Effect_applicativeEffect__1969567048 sync.Once
func Get_Effect_applicativeEffect__1969567048() gopurs_runtime.Value {
	once_Effect_applicativeEffect__1969567048.Do(func() {
		cache_Effect_applicativeEffect__1969567048 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_pureE())
	})
	return cache_Effect_applicativeEffect__1969567048
}

var cache_Effect_applyEffect__1723132130 gopurs_runtime.Value
var once_Effect_applyEffect__1723132130 sync.Once
func Get_Effect_applyEffect__1723132130() gopurs_runtime.Value {
	once_Effect_applyEffect__1723132130.Do(func() {
		cache_Effect_applyEffect__1723132130 = func() gopurs_runtime.Value {
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_Effect_applyEffect__1723132130
}

var cache_Effect_applyEffect__2014400020 gopurs_runtime.Value
var once_Effect_applyEffect__2014400020 sync.Once
func Get_Effect_applyEffect__2014400020() gopurs_runtime.Value {
	once_Effect_applyEffect__2014400020.Do(func() {
		cache_Effect_applyEffect__2014400020 = func() gopurs_runtime.Value {
// TAST (Let): Bind1_0_0 -> *Constructor_Control_Bind_Bind
Bind1_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
// TAST (Let): Applicative0_1_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_0_0.V1), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_1.V1), gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_Effect_applyEffect__2014400020
}

var cache_Effect_bindEffect__2113658466 gopurs_runtime.Value
var once_Effect_bindEffect__2113658466 sync.Once
func Get_Effect_bindEffect__2113658466() gopurs_runtime.Value {
	once_Effect_bindEffect__2113658466.Do(func() {
		cache_Effect_bindEffect__2113658466 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
	})
	return cache_Effect_bindEffect__2113658466
}

var cache_Effect_bindEffect__3856311079 gopurs_runtime.Value
var once_Effect_bindEffect__3856311079 sync.Once
func Get_Effect_bindEffect__3856311079() gopurs_runtime.Value {
	once_Effect_bindEffect__3856311079.Do(func() {
		cache_Effect_bindEffect__3856311079 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
	})
	return cache_Effect_bindEffect__3856311079
}

var cache_Effect_functorEffect__347161653 gopurs_runtime.Value
var once_Effect_functorEffect__347161653 sync.Once
func Get_Effect_functorEffect__347161653() gopurs_runtime.Value {
	once_Effect_functorEffect__347161653.Do(func() {
		cache_Effect_functorEffect__347161653 = func() gopurs_runtime.Value {
// TAST (Let): Apply0_0_0 -> *Constructor_Control_Apply_Apply
Apply0_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_0_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_Effect_functorEffect__347161653
}

var cache_Effect_functorEffect__3107547953 gopurs_runtime.Value
var once_Effect_functorEffect__3107547953 sync.Once
func Get_Effect_functorEffect__3107547953() gopurs_runtime.Value {
	once_Effect_functorEffect__3107547953.Do(func() {
		cache_Effect_functorEffect__3107547953 = func() gopurs_runtime.Value {
// TAST (Let): Apply0_0_0 -> *Constructor_Control_Apply_Apply
Apply0_0_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_0_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_Effect_functorEffect__3107547953
}

var cache_Effect_monadEffect__3527935219 gopurs_runtime.Value
var once_Effect_monadEffect__3527935219 sync.Once
func Get_Effect_monadEffect__3527935219() gopurs_runtime.Value {
	once_Effect_monadEffect__3527935219.Do(func() {
		cache_Effect_monadEffect__3527935219 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_bindEffect()
}))
	})
	return cache_Effect_monadEffect__3527935219
}

func Call_Effect_semigroupEffect(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_applyEffect(), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_2_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_2 -> *Constructor_Control_Bind_Bind
Bind1_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_2
// TAST (Let): Applicative0_6_3 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Effect_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), __local_var_2_1, a_3), gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_2.V1), b_4, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_3.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}

func Call_Effect_monoidEffect(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): semigroupEffect1_1_0 -> gopurs_runtime.Value
semigroupEffect1_1_0 := Call_Effect_semigroupEffect(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupEffect1_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffect1_1_0
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}))
}

func Get_Effect_bindE() gopurs_runtime.Value {
	return _Gopurs_Effect_BindE
}

func Get_Effect_forE() gopurs_runtime.Value {
	return _Gopurs_Effect_ForE
}

func Get_Effect_foreachE() gopurs_runtime.Value {
	return _Gopurs_Effect_ForeachE
}

func Get_Effect_pureE() gopurs_runtime.Value {
	return _Gopurs_Effect_PureE
}

func Get_Effect_untilE() gopurs_runtime.Value {
	return _Gopurs_Effect_UntilE
}

func Get_Effect_whileE() gopurs_runtime.Value {
	return _Gopurs_Effect_WhileE
}

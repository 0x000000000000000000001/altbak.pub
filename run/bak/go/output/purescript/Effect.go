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
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_0_0
// TAST (Let): Bind1_1_1 -> *Constructor_Control_Bind_Bind
Bind1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_1
// TAST (Let): Applicative0_2_2 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_2
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), f_3, gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), a_4, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_2.V1), gopurs_runtime.Apply(f_prime_5, a_prime_6))
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
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Applicative0_3_2 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_2
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}))
}), Get_Effect_pureE())
	})
	return cache_Effect_applicativeEffect
}

var cache_Effect_functorEffect gopurs_runtime.Value
var once_Effect_functorEffect sync.Once
func Get_Effect_functorEffect() gopurs_runtime.Value {
	once_Effect_functorEffect.Do(func() {
		cache_Effect_functorEffect = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_3_2
// TAST (Let): Bind1_4_3 -> *Constructor_Control_Bind_Bind
Bind1_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_3
// TAST (Let): Applicative0_5_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_3.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_3.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_1_1
// TAST (Let): Bind1_2_5 -> *Constructor_Control_Bind_Bind
Bind1_2_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_5
// TAST (Let): Applicative0_3_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_6
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_5.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_5.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_6.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_0_0
// TAST (Let): Apply0_1_7 -> *Constructor_Control_Apply_Apply
Apply0_1_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_7
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_7.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "pure"), f_2), a_3)
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
// TAST (Let): __local_var_1_14 -> gopurs_runtime.Value
__local_var_1_14 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_23 -> gopurs_runtime.Value
__local_var_3_23 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_3_23
// TAST (Let): Bind1_4_24 -> *Constructor_Control_Bind_Bind
Bind1_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_23, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_24
// TAST (Let): Applicative0_5_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_23, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_15 -> gopurs_runtime.Value
__local_var_4_15 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_16 -> gopurs_runtime.Value
__local_var_5_16 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_17 -> gopurs_runtime.Value
__local_var_7_17 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_17
// TAST (Let): Bind1_8_18 -> *Constructor_Control_Bind_Bind
Bind1_8_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_18
// TAST (Let): Applicative0_9_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_19
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_18.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_18.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_19.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_16
// TAST (Let): Bind1_6_20 -> *Constructor_Control_Bind_Bind
Bind1_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_16, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_20
// TAST (Let): Applicative0_7_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_16, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_21
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_21.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_4_15
// TAST (Let): Apply0_5_22 -> *Constructor_Control_Apply_Apply
Apply0_5_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_15, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_22
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_22.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_15, "pure"), f_6), a_7)
})
}))
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_24.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_24.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_25.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_1_14
// TAST (Let): Bind1_2_26 -> *Constructor_Control_Bind_Bind
Bind1_2_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_26
// TAST (Let): Applicative0_3_27 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_27
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_5_2
// TAST (Let): Bind1_6_3 -> *Constructor_Control_Bind_Bind
Bind1_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_3
// TAST (Let): Applicative0_7_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_3.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_3.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_4.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_6
// TAST (Let): Bind1_8_7 -> *Constructor_Control_Bind_Bind
Bind1_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_7
// TAST (Let): Applicative0_9_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_8
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_7.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_7.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_8.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_5_5
// TAST (Let): Bind1_6_9 -> *Constructor_Control_Bind_Bind
Bind1_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_9
// TAST (Let): Applicative0_7_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_10
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_9.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_9.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_10.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_3_1
// TAST (Let): Bind1_4_11 -> *Constructor_Control_Bind_Bind
Bind1_4_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_11
// TAST (Let): Applicative0_5_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_11.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_11.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_12.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_2_0
// TAST (Let): Apply0_3_13 -> *Constructor_Control_Apply_Apply
Apply0_3_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_13
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_13.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "pure"), f_4), a_5)
})
}))
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_26.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_26.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_27.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}))
}), Get_Effect_pureE())
	})
	return cache_Effect_applicativeEffect__284161122
}

var cache_Effect_applicativeEffect__1969567048 gopurs_runtime.Value
var once_Effect_applicativeEffect__1969567048 sync.Once
func Get_Effect_applicativeEffect__1969567048() gopurs_runtime.Value {
	once_Effect_applicativeEffect__1969567048.Do(func() {
		cache_Effect_applicativeEffect__1969567048 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_14 -> gopurs_runtime.Value
__local_var_1_14 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_23 -> gopurs_runtime.Value
__local_var_3_23 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_3_23
// TAST (Let): Bind1_4_24 -> *Constructor_Control_Bind_Bind
Bind1_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_23, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_24
// TAST (Let): Applicative0_5_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_23, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_15 -> gopurs_runtime.Value
__local_var_4_15 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_16 -> gopurs_runtime.Value
__local_var_5_16 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_17 -> gopurs_runtime.Value
__local_var_7_17 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_17
// TAST (Let): Bind1_8_18 -> *Constructor_Control_Bind_Bind
Bind1_8_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_18
// TAST (Let): Applicative0_9_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_19
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_18.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_18.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_19.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_16
// TAST (Let): Bind1_6_20 -> *Constructor_Control_Bind_Bind
Bind1_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_16, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_20
// TAST (Let): Applicative0_7_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_16, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_21
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_21.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_4_15
// TAST (Let): Apply0_5_22 -> *Constructor_Control_Apply_Apply
Apply0_5_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_15, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_22
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_22.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_15, "pure"), f_6), a_7)
})
}))
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_24.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_24.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_25.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_1_14
// TAST (Let): Bind1_2_26 -> *Constructor_Control_Bind_Bind
Bind1_2_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_26
// TAST (Let): Applicative0_3_27 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_27
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_5_2
// TAST (Let): Bind1_6_3 -> *Constructor_Control_Bind_Bind
Bind1_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_3
// TAST (Let): Applicative0_7_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_3.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_3.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_4.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_6
// TAST (Let): Bind1_8_7 -> *Constructor_Control_Bind_Bind
Bind1_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_7
// TAST (Let): Applicative0_9_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_8
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_7.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_7.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_8.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_5_5
// TAST (Let): Bind1_6_9 -> *Constructor_Control_Bind_Bind
Bind1_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_9
// TAST (Let): Applicative0_7_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_10
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_9.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_9.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_10.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_3_1
// TAST (Let): Bind1_4_11 -> *Constructor_Control_Bind_Bind
Bind1_4_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_11
// TAST (Let): Applicative0_5_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_11.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_11.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_12.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_2_0
// TAST (Let): Apply0_3_13 -> *Constructor_Control_Apply_Apply
Apply0_3_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_13
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_13.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "pure"), f_4), a_5)
})
}))
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_26.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_26.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_27.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}))
}), Get_Effect_pureE())
	})
	return cache_Effect_applicativeEffect__1969567048
}

var cache_Effect_applyEffect__1723132130 gopurs_runtime.Value
var once_Effect_applyEffect__1723132130 sync.Once
func Get_Effect_applyEffect__1723132130() gopurs_runtime.Value {
	once_Effect_applyEffect__1723132130.Do(func() {
		cache_Effect_applyEffect__1723132130 = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_26 -> gopurs_runtime.Value
__local_var_0_26 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_38 -> gopurs_runtime.Value
__local_var_2_38 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_47 -> gopurs_runtime.Value
__local_var_4_47 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_4_47
// TAST (Let): Bind1_5_48 -> *Constructor_Control_Bind_Bind
Bind1_5_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_47, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_48
// TAST (Let): Applicative0_6_49 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_47, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_49
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_39 -> gopurs_runtime.Value
__local_var_5_39 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_40 -> gopurs_runtime.Value
__local_var_6_40 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_41 -> gopurs_runtime.Value
__local_var_8_41 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_41
// TAST (Let): Bind1_9_42 -> *Constructor_Control_Bind_Bind
Bind1_9_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_41, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_42
// TAST (Let): Applicative0_10_43 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_41, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_43
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_42.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_42.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_43.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_6_40
// TAST (Let): Bind1_7_44 -> *Constructor_Control_Bind_Bind
Bind1_7_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_40, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_44
// TAST (Let): Applicative0_8_45 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_40, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_45
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_44.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_44.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_45.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_5_39
// TAST (Let): Apply0_6_46 -> *Constructor_Control_Apply_Apply
Apply0_6_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_39, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_6_46
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_6_46.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_39, "pure"), f_7), a_8)
})
}))
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_48.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_48.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_49.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_2_38
// TAST (Let): Bind1_3_50 -> *Constructor_Control_Bind_Bind
Bind1_3_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_38, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_50
// TAST (Let): Applicative0_4_51 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_38, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_51
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_27 -> gopurs_runtime.Value
__local_var_3_27 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_28 -> gopurs_runtime.Value
__local_var_4_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_29 -> gopurs_runtime.Value
__local_var_6_29 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_30 -> gopurs_runtime.Value
__local_var_8_30 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_30
// TAST (Let): Bind1_9_31 -> *Constructor_Control_Bind_Bind
Bind1_9_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_30, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_31
// TAST (Let): Applicative0_10_32 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_30, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_32
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_31.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_31.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_32.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_6_29
// TAST (Let): Bind1_7_33 -> *Constructor_Control_Bind_Bind
Bind1_7_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_29, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_33
// TAST (Let): Applicative0_8_34 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_29, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_34
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_33.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_33.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_34.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_4_28
// TAST (Let): Bind1_5_35 -> *Constructor_Control_Bind_Bind
Bind1_5_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_35
// TAST (Let): Applicative0_6_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_36
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_35.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_35.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_36.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_3_27
// TAST (Let): Apply0_4_37 -> *Constructor_Control_Apply_Apply
Apply0_4_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_27, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_37
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_37.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_27, "pure"), f_5), a_6)
})
}))
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_50.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_50.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_51.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_66 -> gopurs_runtime.Value
__local_var_2_66 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_75 -> gopurs_runtime.Value
__local_var_4_75 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_4_75
// TAST (Let): Bind1_5_76 -> *Constructor_Control_Bind_Bind
Bind1_5_76 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_75, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_76
// TAST (Let): Applicative0_6_77 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_75, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_77
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_67 -> gopurs_runtime.Value
__local_var_5_67 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_68 -> gopurs_runtime.Value
__local_var_6_68 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_69 -> gopurs_runtime.Value
__local_var_8_69 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_69
// TAST (Let): Bind1_9_70 -> *Constructor_Control_Bind_Bind
Bind1_9_70 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_69, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_70
// TAST (Let): Applicative0_10_71 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_71 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_69, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_71
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_70.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_70.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_71.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_6_68
// TAST (Let): Bind1_7_72 -> *Constructor_Control_Bind_Bind
Bind1_7_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_68, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_72
// TAST (Let): Applicative0_8_73 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_68, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_73
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_72.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_72.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_73.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_5_67
// TAST (Let): Apply0_6_74 -> *Constructor_Control_Apply_Apply
Apply0_6_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_67, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_6_74
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_6_74.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_67, "pure"), f_7), a_8)
})
}))
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_76.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_76.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_77.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_2_66
// TAST (Let): Bind1_3_78 -> *Constructor_Control_Bind_Bind
Bind1_3_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_66, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_78
// TAST (Let): Applicative0_4_79 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_79 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_66, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_79
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_52 -> gopurs_runtime.Value
__local_var_3_52 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_53 -> gopurs_runtime.Value
__local_var_4_53 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_54 -> gopurs_runtime.Value
__local_var_6_54 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_6_54
// TAST (Let): Bind1_7_55 -> *Constructor_Control_Bind_Bind
Bind1_7_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_54, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_55
// TAST (Let): Applicative0_8_56 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_54, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_56
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_55.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_55.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_56.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_57 -> gopurs_runtime.Value
__local_var_6_57 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_58 -> gopurs_runtime.Value
__local_var_8_58 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_58
// TAST (Let): Bind1_9_59 -> *Constructor_Control_Bind_Bind
Bind1_9_59 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_58, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_59
// TAST (Let): Applicative0_10_60 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_58, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_60
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_59.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_59.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_60.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_6_57
// TAST (Let): Bind1_7_61 -> *Constructor_Control_Bind_Bind
Bind1_7_61 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_57, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_61
// TAST (Let): Applicative0_8_62 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_62 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_57, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_62
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_61.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_61.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_62.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_4_53
// TAST (Let): Bind1_5_63 -> *Constructor_Control_Bind_Bind
Bind1_5_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_53, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_63
// TAST (Let): Applicative0_6_64 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_53, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_64
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_63.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_63.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_64.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_3_52
// TAST (Let): Apply0_4_65 -> *Constructor_Control_Apply_Apply
Apply0_4_65 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_52, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_65
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_65.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_52, "pure"), f_5), a_6)
})
}))
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_78.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_78.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_79.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_0_26
// TAST (Let): Bind1_1_80 -> *Constructor_Control_Bind_Bind
Bind1_1_80 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_26, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_80
// TAST (Let): Applicative0_2_81 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_26, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_81
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_6_3
// TAST (Let): Bind1_7_4 -> *Constructor_Control_Bind_Bind
Bind1_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_4
// TAST (Let): Applicative0_8_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_5
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_4.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_4.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_5.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_4_2
// TAST (Let): Bind1_5_6 -> *Constructor_Control_Bind_Bind
Bind1_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_6
// TAST (Let): Applicative0_6_7 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_7
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_6.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_6.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_7.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_9 -> gopurs_runtime.Value
__local_var_6_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_10
// TAST (Let): Bind1_9_11 -> *Constructor_Control_Bind_Bind
Bind1_9_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_11
// TAST (Let): Applicative0_10_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_11.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_11.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_12.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_6_9
// TAST (Let): Bind1_7_13 -> *Constructor_Control_Bind_Bind
Bind1_7_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_9, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_13
// TAST (Let): Applicative0_8_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_9, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_14.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_15 -> gopurs_runtime.Value
__local_var_6_15 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_16 -> gopurs_runtime.Value
__local_var_8_16 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_16
// TAST (Let): Bind1_9_17 -> *Constructor_Control_Bind_Bind
Bind1_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_16, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_17
// TAST (Let): Applicative0_10_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_16, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_18
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_17.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_17.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_18.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_6_15
// TAST (Let): Bind1_7_19 -> *Constructor_Control_Bind_Bind
Bind1_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_19
// TAST (Let): Applicative0_8_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_20
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_20.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_4_8
// TAST (Let): Bind1_5_21 -> *Constructor_Control_Bind_Bind
Bind1_5_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_21
// TAST (Let): Applicative0_6_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_22
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_21.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_21.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_22.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_2_1
// TAST (Let): Bind1_3_23 -> *Constructor_Control_Bind_Bind
Bind1_3_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_23
// TAST (Let): Applicative0_4_24 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_24
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_23.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_23.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_24.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_1_0
// TAST (Let): Apply0_2_25 -> *Constructor_Control_Apply_Apply
Apply0_2_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_25
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_25.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), f_3), a_4)
})
}))
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_80.V1), f_3, gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_80.V1), a_4, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_81.V1), gopurs_runtime.Apply(f_prime_5, a_prime_6))
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
// TAST (Let): __local_var_0_26 -> gopurs_runtime.Value
__local_var_0_26 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_38 -> gopurs_runtime.Value
__local_var_2_38 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_47 -> gopurs_runtime.Value
__local_var_4_47 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_4_47
// TAST (Let): Bind1_5_48 -> *Constructor_Control_Bind_Bind
Bind1_5_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_47, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_48
// TAST (Let): Applicative0_6_49 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_47, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_49
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_39 -> gopurs_runtime.Value
__local_var_5_39 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_40 -> gopurs_runtime.Value
__local_var_6_40 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_41 -> gopurs_runtime.Value
__local_var_8_41 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_41
// TAST (Let): Bind1_9_42 -> *Constructor_Control_Bind_Bind
Bind1_9_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_41, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_42
// TAST (Let): Applicative0_10_43 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_41, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_43
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_42.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_42.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_43.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_6_40
// TAST (Let): Bind1_7_44 -> *Constructor_Control_Bind_Bind
Bind1_7_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_40, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_44
// TAST (Let): Applicative0_8_45 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_40, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_45
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_44.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_44.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_45.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_5_39
// TAST (Let): Apply0_6_46 -> *Constructor_Control_Apply_Apply
Apply0_6_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_39, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_6_46
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_6_46.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_39, "pure"), f_7), a_8)
})
}))
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_48.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_48.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_49.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_2_38
// TAST (Let): Bind1_3_50 -> *Constructor_Control_Bind_Bind
Bind1_3_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_38, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_50
// TAST (Let): Applicative0_4_51 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_38, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_51
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_27 -> gopurs_runtime.Value
__local_var_3_27 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_28 -> gopurs_runtime.Value
__local_var_4_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_29 -> gopurs_runtime.Value
__local_var_6_29 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_30 -> gopurs_runtime.Value
__local_var_8_30 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_30
// TAST (Let): Bind1_9_31 -> *Constructor_Control_Bind_Bind
Bind1_9_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_30, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_31
// TAST (Let): Applicative0_10_32 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_30, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_32
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_31.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_31.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_32.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_6_29
// TAST (Let): Bind1_7_33 -> *Constructor_Control_Bind_Bind
Bind1_7_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_29, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_33
// TAST (Let): Applicative0_8_34 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_29, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_34
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_33.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_33.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_34.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_4_28
// TAST (Let): Bind1_5_35 -> *Constructor_Control_Bind_Bind
Bind1_5_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_35
// TAST (Let): Applicative0_6_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_36
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_35.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_35.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_36.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_3_27
// TAST (Let): Apply0_4_37 -> *Constructor_Control_Apply_Apply
Apply0_4_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_27, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_37
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_37.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_27, "pure"), f_5), a_6)
})
}))
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_50.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_50.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_51.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_66 -> gopurs_runtime.Value
__local_var_2_66 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_75 -> gopurs_runtime.Value
__local_var_4_75 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_4_75
// TAST (Let): Bind1_5_76 -> *Constructor_Control_Bind_Bind
Bind1_5_76 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_75, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_76
// TAST (Let): Applicative0_6_77 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_75, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_77
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_67 -> gopurs_runtime.Value
__local_var_5_67 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_68 -> gopurs_runtime.Value
__local_var_6_68 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_69 -> gopurs_runtime.Value
__local_var_8_69 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_69
// TAST (Let): Bind1_9_70 -> *Constructor_Control_Bind_Bind
Bind1_9_70 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_69, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_70
// TAST (Let): Applicative0_10_71 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_71 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_69, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_71
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_70.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_70.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_71.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_6_68
// TAST (Let): Bind1_7_72 -> *Constructor_Control_Bind_Bind
Bind1_7_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_68, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_72
// TAST (Let): Applicative0_8_73 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_68, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_73
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_72.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_72.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_73.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_5_67
// TAST (Let): Apply0_6_74 -> *Constructor_Control_Apply_Apply
Apply0_6_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_67, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_6_74
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_6_74.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_67, "pure"), f_7), a_8)
})
}))
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_76.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_76.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_77.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_2_66
// TAST (Let): Bind1_3_78 -> *Constructor_Control_Bind_Bind
Bind1_3_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_66, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_78
// TAST (Let): Applicative0_4_79 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_79 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_66, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_79
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_52 -> gopurs_runtime.Value
__local_var_3_52 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_53 -> gopurs_runtime.Value
__local_var_4_53 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_54 -> gopurs_runtime.Value
__local_var_6_54 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_6_54
// TAST (Let): Bind1_7_55 -> *Constructor_Control_Bind_Bind
Bind1_7_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_54, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_55
// TAST (Let): Applicative0_8_56 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_54, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_56
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_55.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_55.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_56.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_57 -> gopurs_runtime.Value
__local_var_6_57 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_58 -> gopurs_runtime.Value
__local_var_8_58 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_58
// TAST (Let): Bind1_9_59 -> *Constructor_Control_Bind_Bind
Bind1_9_59 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_58, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_59
// TAST (Let): Applicative0_10_60 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_58, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_60
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_59.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_59.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_60.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_6_57
// TAST (Let): Bind1_7_61 -> *Constructor_Control_Bind_Bind
Bind1_7_61 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_57, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_61
// TAST (Let): Applicative0_8_62 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_62 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_57, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_62
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_61.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_61.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_62.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_4_53
// TAST (Let): Bind1_5_63 -> *Constructor_Control_Bind_Bind
Bind1_5_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_53, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_63
// TAST (Let): Applicative0_6_64 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_53, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_64
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_63.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_63.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_64.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_3_52
// TAST (Let): Apply0_4_65 -> *Constructor_Control_Apply_Apply
Apply0_4_65 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_52, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_65
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_65.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_52, "pure"), f_5), a_6)
})
}))
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_78.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_78.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_79.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_0_26
// TAST (Let): Bind1_1_80 -> *Constructor_Control_Bind_Bind
Bind1_1_80 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_26, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_80
// TAST (Let): Applicative0_2_81 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_26, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_81
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_6_3
// TAST (Let): Bind1_7_4 -> *Constructor_Control_Bind_Bind
Bind1_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_4
// TAST (Let): Applicative0_8_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_5
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_4.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_4.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_5.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_4_2
// TAST (Let): Bind1_5_6 -> *Constructor_Control_Bind_Bind
Bind1_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_6
// TAST (Let): Applicative0_6_7 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_7
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_6.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_6.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_7.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_9 -> gopurs_runtime.Value
__local_var_6_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_10
// TAST (Let): Bind1_9_11 -> *Constructor_Control_Bind_Bind
Bind1_9_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_11
// TAST (Let): Applicative0_10_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_11.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_11.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_12.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_6_9
// TAST (Let): Bind1_7_13 -> *Constructor_Control_Bind_Bind
Bind1_7_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_9, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_13
// TAST (Let): Applicative0_8_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_9, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_14.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_15 -> gopurs_runtime.Value
__local_var_6_15 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_16 -> gopurs_runtime.Value
__local_var_8_16 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_16
// TAST (Let): Bind1_9_17 -> *Constructor_Control_Bind_Bind
Bind1_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_16, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_17
// TAST (Let): Applicative0_10_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_16, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_18
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_17.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_17.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_18.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_6_15
// TAST (Let): Bind1_7_19 -> *Constructor_Control_Bind_Bind
Bind1_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_19
// TAST (Let): Applicative0_8_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_20
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_20.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_4_8
// TAST (Let): Bind1_5_21 -> *Constructor_Control_Bind_Bind
Bind1_5_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_21
// TAST (Let): Applicative0_6_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_22
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_21.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_21.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_22.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_2_1
// TAST (Let): Bind1_3_23 -> *Constructor_Control_Bind_Bind
Bind1_3_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_23
// TAST (Let): Applicative0_4_24 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_24
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_23.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_23.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_24.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_1_0
// TAST (Let): Apply0_2_25 -> *Constructor_Control_Apply_Apply
Apply0_2_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_25
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_25.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), f_3), a_4)
})
}))
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_80.V1), f_3, gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_80.V1), a_4, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_81.V1), gopurs_runtime.Apply(f_prime_5, a_prime_6))
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
// TAST (Let): __local_var_1_14 -> gopurs_runtime.Value
__local_var_1_14 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_23 -> gopurs_runtime.Value
__local_var_3_23 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_3_23
// TAST (Let): Bind1_4_24 -> *Constructor_Control_Bind_Bind
Bind1_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_23, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_24
// TAST (Let): Applicative0_5_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_23, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_15 -> gopurs_runtime.Value
__local_var_4_15 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_16 -> gopurs_runtime.Value
__local_var_5_16 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_17 -> gopurs_runtime.Value
__local_var_7_17 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_17
// TAST (Let): Bind1_8_18 -> *Constructor_Control_Bind_Bind
Bind1_8_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_18
// TAST (Let): Applicative0_9_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_19
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_18.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_18.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_19.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_16
// TAST (Let): Bind1_6_20 -> *Constructor_Control_Bind_Bind
Bind1_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_16, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_20
// TAST (Let): Applicative0_7_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_16, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_21
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_21.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_4_15
// TAST (Let): Apply0_5_22 -> *Constructor_Control_Apply_Apply
Apply0_5_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_15, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_22
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_22.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_15, "pure"), f_6), a_7)
})
}))
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_24.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_24.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_25.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_1_14
// TAST (Let): Bind1_2_26 -> *Constructor_Control_Bind_Bind
Bind1_2_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_26
// TAST (Let): Applicative0_3_27 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_27
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_5_2
// TAST (Let): Bind1_6_3 -> *Constructor_Control_Bind_Bind
Bind1_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_3
// TAST (Let): Applicative0_7_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_3.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_3.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_4.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_6
// TAST (Let): Bind1_8_7 -> *Constructor_Control_Bind_Bind
Bind1_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_7
// TAST (Let): Applicative0_9_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_8
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_7.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_7.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_8.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_5_5
// TAST (Let): Bind1_6_9 -> *Constructor_Control_Bind_Bind
Bind1_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_9
// TAST (Let): Applicative0_7_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_10
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_9.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_9.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_10.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_3_1
// TAST (Let): Bind1_4_11 -> *Constructor_Control_Bind_Bind
Bind1_4_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_11
// TAST (Let): Applicative0_5_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_11.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_11.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_12.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_2_0
// TAST (Let): Apply0_3_13 -> *Constructor_Control_Apply_Apply
Apply0_3_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_13
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_13.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "pure"), f_4), a_5)
})
}))
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_26.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_26.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_27.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}))
}), Get_Effect_bindE())
	})
	return cache_Effect_bindEffect__2113658466
}

var cache_Effect_bindEffect__3856311079 gopurs_runtime.Value
var once_Effect_bindEffect__3856311079 sync.Once
func Get_Effect_bindEffect__3856311079() gopurs_runtime.Value {
	once_Effect_bindEffect__3856311079.Do(func() {
		cache_Effect_bindEffect__3856311079 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_14 -> gopurs_runtime.Value
__local_var_1_14 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_23 -> gopurs_runtime.Value
__local_var_3_23 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_3_23
// TAST (Let): Bind1_4_24 -> *Constructor_Control_Bind_Bind
Bind1_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_23, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_24
// TAST (Let): Applicative0_5_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_23, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_15 -> gopurs_runtime.Value
__local_var_4_15 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_16 -> gopurs_runtime.Value
__local_var_5_16 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_17 -> gopurs_runtime.Value
__local_var_7_17 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_17
// TAST (Let): Bind1_8_18 -> *Constructor_Control_Bind_Bind
Bind1_8_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_18
// TAST (Let): Applicative0_9_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_19
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_18.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_18.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_19.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_16
// TAST (Let): Bind1_6_20 -> *Constructor_Control_Bind_Bind
Bind1_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_16, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_20
// TAST (Let): Applicative0_7_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_16, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_21
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_21.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_4_15
// TAST (Let): Apply0_5_22 -> *Constructor_Control_Apply_Apply
Apply0_5_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_15, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_22
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_22.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_15, "pure"), f_6), a_7)
})
}))
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_24.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_24.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_25.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_1_14
// TAST (Let): Bind1_2_26 -> *Constructor_Control_Bind_Bind
Bind1_2_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_26
// TAST (Let): Applicative0_3_27 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_27
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_5_2
// TAST (Let): Bind1_6_3 -> *Constructor_Control_Bind_Bind
Bind1_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_3
// TAST (Let): Applicative0_7_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_3.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_3.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_4.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_6
// TAST (Let): Bind1_8_7 -> *Constructor_Control_Bind_Bind
Bind1_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_7
// TAST (Let): Applicative0_9_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_8
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_7.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_7.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_8.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_5_5
// TAST (Let): Bind1_6_9 -> *Constructor_Control_Bind_Bind
Bind1_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_9
// TAST (Let): Applicative0_7_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_10
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_9.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_9.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_10.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_3_1
// TAST (Let): Bind1_4_11 -> *Constructor_Control_Bind_Bind
Bind1_4_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_11
// TAST (Let): Applicative0_5_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_11.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_11.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_12.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_2_0
// TAST (Let): Apply0_3_13 -> *Constructor_Control_Apply_Apply
Apply0_3_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_13
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_13.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "pure"), f_4), a_5)
})
}))
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_26.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_26.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_27.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}))
}), Get_Effect_bindE())
	})
	return cache_Effect_bindEffect__3856311079
}

var cache_Effect_functorEffect__347161653 gopurs_runtime.Value
var once_Effect_functorEffect__347161653 sync.Once
func Get_Effect_functorEffect__347161653() gopurs_runtime.Value {
	once_Effect_functorEffect__347161653.Do(func() {
		cache_Effect_functorEffect__347161653 = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_12 -> gopurs_runtime.Value
__local_var_1_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_21 -> gopurs_runtime.Value
__local_var_3_21 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_3_21
// TAST (Let): Bind1_4_22 -> *Constructor_Control_Bind_Bind
Bind1_4_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_21, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_22
// TAST (Let): Applicative0_5_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_21, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_23
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_13 -> gopurs_runtime.Value
__local_var_4_13 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_14 -> gopurs_runtime.Value
__local_var_5_14 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_15 -> gopurs_runtime.Value
__local_var_7_15 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_15
// TAST (Let): Bind1_8_16 -> *Constructor_Control_Bind_Bind
Bind1_8_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_16
// TAST (Let): Applicative0_9_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_17
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_16.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_16.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_17.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_14
// TAST (Let): Bind1_6_18 -> *Constructor_Control_Bind_Bind
Bind1_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_18
// TAST (Let): Applicative0_7_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_19
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_19.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_4_13
// TAST (Let): Apply0_5_20 -> *Constructor_Control_Apply_Apply
Apply0_5_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_20
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_20.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "pure"), f_6), a_7)
})
}))
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_22.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_22.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_23.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_1_12
// TAST (Let): Bind1_2_24 -> *Constructor_Control_Bind_Bind
Bind1_2_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_24
// TAST (Let): Applicative0_3_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_4
// TAST (Let): Bind1_8_5 -> *Constructor_Control_Bind_Bind
Bind1_8_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_5
// TAST (Let): Applicative0_9_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_6
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_5.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_5.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_6.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_3
// TAST (Let): Bind1_6_7 -> *Constructor_Control_Bind_Bind
Bind1_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_7
// TAST (Let): Applicative0_7_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_8
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_8.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_3_2
// TAST (Let): Bind1_4_9 -> *Constructor_Control_Bind_Bind
Bind1_4_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_9
// TAST (Let): Applicative0_5_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_10
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_9.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_9.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_10.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_2_1
// TAST (Let): Apply0_3_11 -> *Constructor_Control_Apply_Apply
Apply0_3_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_11
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_11.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), f_4), a_5)
})
}))
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_24.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_24.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_25.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_0_0
// TAST (Let): Apply0_1_26 -> *Constructor_Control_Apply_Apply
Apply0_1_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_26
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_26.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "pure"), f_2), a_3)
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
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_12 -> gopurs_runtime.Value
__local_var_1_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_21 -> gopurs_runtime.Value
__local_var_3_21 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_3_21
// TAST (Let): Bind1_4_22 -> *Constructor_Control_Bind_Bind
Bind1_4_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_21, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_22
// TAST (Let): Applicative0_5_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_21, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_23
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_13 -> gopurs_runtime.Value
__local_var_4_13 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_14 -> gopurs_runtime.Value
__local_var_5_14 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_15 -> gopurs_runtime.Value
__local_var_7_15 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_15
// TAST (Let): Bind1_8_16 -> *Constructor_Control_Bind_Bind
Bind1_8_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_16
// TAST (Let): Applicative0_9_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_17
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_16.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_16.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_17.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_14
// TAST (Let): Bind1_6_18 -> *Constructor_Control_Bind_Bind
Bind1_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_18
// TAST (Let): Applicative0_7_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_19
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_19.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_4_13
// TAST (Let): Apply0_5_20 -> *Constructor_Control_Apply_Apply
Apply0_5_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_20
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_20.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "pure"), f_6), a_7)
})
}))
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_22.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_22.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_23.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_1_12
// TAST (Let): Bind1_2_24 -> *Constructor_Control_Bind_Bind
Bind1_2_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_24
// TAST (Let): Applicative0_3_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_4
// TAST (Let): Bind1_8_5 -> *Constructor_Control_Bind_Bind
Bind1_8_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_5
// TAST (Let): Applicative0_9_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_6
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_5.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_5.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_6.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_3
// TAST (Let): Bind1_6_7 -> *Constructor_Control_Bind_Bind
Bind1_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_7
// TAST (Let): Applicative0_7_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_8
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_8.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_3_2
// TAST (Let): Bind1_4_9 -> *Constructor_Control_Bind_Bind
Bind1_4_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_9
// TAST (Let): Applicative0_5_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_10
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_9.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_9.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_10.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_2_1
// TAST (Let): Apply0_3_11 -> *Constructor_Control_Apply_Apply
Apply0_3_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_11
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_11.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), f_4), a_5)
})
}))
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_24.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_24.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_25.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_0_0
// TAST (Let): Apply0_1_26 -> *Constructor_Control_Apply_Apply
Apply0_1_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_26
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_26.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_0, "pure"), f_2), a_3)
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
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_11 -> gopurs_runtime.Value
__local_var_2_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_20 -> gopurs_runtime.Value
__local_var_4_20 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_4_20
// TAST (Let): Bind1_5_21 -> *Constructor_Control_Bind_Bind
Bind1_5_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_20, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_21
// TAST (Let): Applicative0_6_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_20, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_22
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_12 -> gopurs_runtime.Value
__local_var_5_12 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_13 -> gopurs_runtime.Value
__local_var_6_13 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_14 -> gopurs_runtime.Value
__local_var_8_14 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_14
// TAST (Let): Bind1_9_15 -> *Constructor_Control_Bind_Bind
Bind1_9_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_15
// TAST (Let): Applicative0_10_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_16
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_15.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_15.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_16.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_6_13
// TAST (Let): Bind1_7_17 -> *Constructor_Control_Bind_Bind
Bind1_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_17
// TAST (Let): Applicative0_8_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_18
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_17.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_17.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_18.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_5_12
// TAST (Let): Apply0_6_19 -> *Constructor_Control_Apply_Apply
Apply0_6_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_12, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_6_19
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_6_19.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_12, "pure"), f_7), a_8)
})
}))
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_21.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_21.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_22.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_2_11
// TAST (Let): Bind1_3_23 -> *Constructor_Control_Bind_Bind
Bind1_3_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_23
// TAST (Let): Applicative0_4_24 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_24
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_3 -> gopurs_runtime.Value
__local_var_8_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_3
// TAST (Let): Bind1_9_4 -> *Constructor_Control_Bind_Bind
Bind1_9_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_4
// TAST (Let): Applicative0_10_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_5
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_4.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_4.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_5.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_6_2
// TAST (Let): Bind1_7_6 -> *Constructor_Control_Bind_Bind
Bind1_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_6
// TAST (Let): Applicative0_8_7 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_7
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_6.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_6.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_7.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_4_1
// TAST (Let): Bind1_5_8 -> *Constructor_Control_Bind_Bind
Bind1_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_8
// TAST (Let): Applicative0_6_9 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_9
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_8.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_8.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_9.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_3_0
// TAST (Let): Apply0_4_10 -> *Constructor_Control_Apply_Apply
Apply0_4_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_10
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_10.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_0, "pure"), f_5), a_6)
})
}))
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_23.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_23.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_24.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_39 -> gopurs_runtime.Value
__local_var_2_39 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_48 -> gopurs_runtime.Value
__local_var_4_48 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_4_48
// TAST (Let): Bind1_5_49 -> *Constructor_Control_Bind_Bind
Bind1_5_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_48, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_49
// TAST (Let): Applicative0_6_50 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_48, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_50
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_40 -> gopurs_runtime.Value
__local_var_5_40 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_41 -> gopurs_runtime.Value
__local_var_6_41 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_42 -> gopurs_runtime.Value
__local_var_8_42 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_42
// TAST (Let): Bind1_9_43 -> *Constructor_Control_Bind_Bind
Bind1_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_42, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_43
// TAST (Let): Applicative0_10_44 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_42, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_44
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_43.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_43.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_44.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_6_41
// TAST (Let): Bind1_7_45 -> *Constructor_Control_Bind_Bind
Bind1_7_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_41, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_45
// TAST (Let): Applicative0_8_46 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_41, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_46
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_45.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_45.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_46.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_5_40
// TAST (Let): Apply0_6_47 -> *Constructor_Control_Apply_Apply
Apply0_6_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_40, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_6_47
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_6_47.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_40, "pure"), f_7), a_8)
})
}))
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_49.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_49.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_50.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_2_39
// TAST (Let): Bind1_3_51 -> *Constructor_Control_Bind_Bind
Bind1_3_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_39, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_51
// TAST (Let): Applicative0_4_52 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_52 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_39, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_52
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_25 -> gopurs_runtime.Value
__local_var_3_25 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_26 -> gopurs_runtime.Value
__local_var_4_26 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_27 -> gopurs_runtime.Value
__local_var_6_27 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_6_27
// TAST (Let): Bind1_7_28 -> *Constructor_Control_Bind_Bind
Bind1_7_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_27, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_28
// TAST (Let): Applicative0_8_29 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_27, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_29
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_28.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_28.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_29.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_30 -> gopurs_runtime.Value
__local_var_6_30 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_31 -> gopurs_runtime.Value
__local_var_8_31 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_8_31
// TAST (Let): Bind1_9_32 -> *Constructor_Control_Bind_Bind
Bind1_9_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_31, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_32
// TAST (Let): Applicative0_10_33 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_31, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_33
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_32.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_32.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_33.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_6_30
// TAST (Let): Bind1_7_34 -> *Constructor_Control_Bind_Bind
Bind1_7_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_30, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_34
// TAST (Let): Applicative0_8_35 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_30, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_35
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_34.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_34.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_35.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_4_26
// TAST (Let): Bind1_5_36 -> *Constructor_Control_Bind_Bind
Bind1_5_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_26, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_36
// TAST (Let): Applicative0_6_37 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_26, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_37
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_36.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_36.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_37.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_3_25
// TAST (Let): Apply0_4_38 -> *Constructor_Control_Apply_Apply
Apply0_4_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_25, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_38
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_38.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_25, "pure"), f_5), a_6)
})
}))
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_51.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_51.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_52.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
	})
	return cache_Effect_monadEffect__3527935219
}

func Call_Effect_semigroupEffect(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): __local_var_1_15 -> gopurs_runtime.Value
__local_var_1_15 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_24 -> gopurs_runtime.Value
__local_var_3_24 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_3_24
// TAST (Let): Bind1_4_25 -> *Constructor_Control_Bind_Bind
Bind1_4_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_24, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_25
// TAST (Let): Applicative0_5_26 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_24, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_26
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_16 -> gopurs_runtime.Value
__local_var_4_16 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_17 -> gopurs_runtime.Value
__local_var_5_17 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_18 -> gopurs_runtime.Value
__local_var_7_18 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_18
// TAST (Let): Bind1_8_19 -> *Constructor_Control_Bind_Bind
Bind1_8_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_18, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_19
// TAST (Let): Applicative0_9_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_18, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_20
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_19.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_19.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_20.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_17
// TAST (Let): Bind1_6_21 -> *Constructor_Control_Bind_Bind
Bind1_6_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_21
// TAST (Let): Applicative0_7_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_22
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_21.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_21.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_22.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_4_16
// TAST (Let): Apply0_5_23 -> *Constructor_Control_Apply_Apply
Apply0_5_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_16, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_23
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_23.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_16, "pure"), f_6), a_7)
})
}))
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_25.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_25.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_26.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_1_15
// TAST (Let): Bind1_2_27 -> *Constructor_Control_Bind_Bind
Bind1_2_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_27
// TAST (Let): Applicative0_3_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_28
// TAST (Let): __local_var_1_0 -> *Constructor_Control_Apply_Apply
__local_var_1_0 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_5_3
// TAST (Let): Bind1_6_4 -> *Constructor_Control_Bind_Bind
Bind1_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_4
// TAST (Let): Applicative0_7_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_5
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_4.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_4.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_5.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_7
// TAST (Let): Bind1_8_8 -> *Constructor_Control_Bind_Bind
Bind1_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_8
// TAST (Let): Applicative0_9_9 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_9
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_8.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_8.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_9.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_5_6
// TAST (Let): Bind1_6_10 -> *Constructor_Control_Bind_Bind
Bind1_6_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_10
// TAST (Let): Applicative0_7_11 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_11
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_10.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_10.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_11.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_3_2
// TAST (Let): Bind1_4_12 -> *Constructor_Control_Bind_Bind
Bind1_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_12
// TAST (Let): Applicative0_5_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_13
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_12.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_12.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_13.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_2_1
// TAST (Let): Apply0_3_14 -> *Constructor_Control_Apply_Apply
Apply0_3_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_14
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_14.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), f_4), a_5)
})
}))
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_27.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_27.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_28.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
})}
_ = __local_var_1_0
// TAST (Let): Functor0_2_29 -> *Constructor_Data_Functor_Functor
Functor0_2_29 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_2_29
// TAST (Let): __local_var_3_30 -> gopurs_runtime.Value
__local_var_3_30 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = __local_var_3_30
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_29.V0), __local_var_3_30, a_4), b_5)
})
}))
}

func Call_Effect_monoidEffect(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_254 -> gopurs_runtime.Value
__local_var_1_254 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_393 -> gopurs_runtime.Value
__local_var_3_393 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_405 -> gopurs_runtime.Value
__local_var_5_405 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_414 -> gopurs_runtime.Value
__local_var_7_414 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_414
// TAST (Let): Bind1_8_415 -> *Constructor_Control_Bind_Bind
Bind1_8_415 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_414, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_415
// TAST (Let): Applicative0_9_416 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_416 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_414, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_416
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_406 -> gopurs_runtime.Value
__local_var_8_406 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_407 -> gopurs_runtime.Value
__local_var_9_407 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_408 -> gopurs_runtime.Value
__local_var_11_408 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_408
// TAST (Let): Bind1_12_409 -> *Constructor_Control_Bind_Bind
Bind1_12_409 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_408, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_409
// TAST (Let): Applicative0_13_410 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_410 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_408, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_410
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_409.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_409.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_410.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_407
// TAST (Let): Bind1_10_411 -> *Constructor_Control_Bind_Bind
Bind1_10_411 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_407, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_411
// TAST (Let): Applicative0_11_412 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_412 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_407, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_412
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_411.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_411.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_412.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_8_406
// TAST (Let): Apply0_9_413 -> *Constructor_Control_Apply_Apply
Apply0_9_413 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_406, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_9_413
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_9_413.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_406, "pure"), f_10), a_11)
})
}))
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_415.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_415.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_416.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_405
// TAST (Let): Bind1_6_417 -> *Constructor_Control_Bind_Bind
Bind1_6_417 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_405, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_417
// TAST (Let): Applicative0_7_418 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_418 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_405, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_418
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_394 -> gopurs_runtime.Value
__local_var_6_394 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_395 -> gopurs_runtime.Value
__local_var_7_395 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_396 -> gopurs_runtime.Value
__local_var_9_396 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_397 -> gopurs_runtime.Value
__local_var_11_397 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_397
// TAST (Let): Bind1_12_398 -> *Constructor_Control_Bind_Bind
Bind1_12_398 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_397, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_398
// TAST (Let): Applicative0_13_399 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_399 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_397, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_399
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_398.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_398.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_399.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_396
// TAST (Let): Bind1_10_400 -> *Constructor_Control_Bind_Bind
Bind1_10_400 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_396, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_400
// TAST (Let): Applicative0_11_401 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_401 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_396, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_401
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_400.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_400.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_401.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_7_395
// TAST (Let): Bind1_8_402 -> *Constructor_Control_Bind_Bind
Bind1_8_402 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_395, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_402
// TAST (Let): Applicative0_9_403 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_403 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_395, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_403
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_402.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_402.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_403.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_6_394
// TAST (Let): Apply0_7_404 -> *Constructor_Control_Apply_Apply
Apply0_7_404 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_394, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_404
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_404.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_394, "pure"), f_8), a_9)
})
}))
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_417.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_417.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_418.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_433 -> gopurs_runtime.Value
__local_var_5_433 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_442 -> gopurs_runtime.Value
__local_var_7_442 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_442
// TAST (Let): Bind1_8_443 -> *Constructor_Control_Bind_Bind
Bind1_8_443 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_442, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_443
// TAST (Let): Applicative0_9_444 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_444 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_442, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_444
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_434 -> gopurs_runtime.Value
__local_var_8_434 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_435 -> gopurs_runtime.Value
__local_var_9_435 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_436 -> gopurs_runtime.Value
__local_var_11_436 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_436
// TAST (Let): Bind1_12_437 -> *Constructor_Control_Bind_Bind
Bind1_12_437 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_436, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_437
// TAST (Let): Applicative0_13_438 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_438 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_436, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_438
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_437.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_437.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_438.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_435
// TAST (Let): Bind1_10_439 -> *Constructor_Control_Bind_Bind
Bind1_10_439 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_435, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_439
// TAST (Let): Applicative0_11_440 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_440 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_435, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_440
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_439.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_439.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_440.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_8_434
// TAST (Let): Apply0_9_441 -> *Constructor_Control_Apply_Apply
Apply0_9_441 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_434, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_9_441
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_9_441.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_434, "pure"), f_10), a_11)
})
}))
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_443.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_443.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_444.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_5_433
// TAST (Let): Bind1_6_445 -> *Constructor_Control_Bind_Bind
Bind1_6_445 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_433, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_445
// TAST (Let): Applicative0_7_446 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_446 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_433, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_446
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_419 -> gopurs_runtime.Value
__local_var_6_419 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_420 -> gopurs_runtime.Value
__local_var_7_420 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_421 -> gopurs_runtime.Value
__local_var_9_421 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_9_421
// TAST (Let): Bind1_10_422 -> *Constructor_Control_Bind_Bind
Bind1_10_422 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_421, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_422
// TAST (Let): Applicative0_11_423 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_423 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_421, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_423
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_422.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_422.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_423.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_424 -> gopurs_runtime.Value
__local_var_9_424 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_425 -> gopurs_runtime.Value
__local_var_11_425 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_425
// TAST (Let): Bind1_12_426 -> *Constructor_Control_Bind_Bind
Bind1_12_426 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_425, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_426
// TAST (Let): Applicative0_13_427 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_427 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_425, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_427
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_426.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_426.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_427.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_9_424
// TAST (Let): Bind1_10_428 -> *Constructor_Control_Bind_Bind
Bind1_10_428 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_424, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_428
// TAST (Let): Applicative0_11_429 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_429 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_424, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_429
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_428.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_428.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_429.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_7_420
// TAST (Let): Bind1_8_430 -> *Constructor_Control_Bind_Bind
Bind1_8_430 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_420, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_430
// TAST (Let): Applicative0_9_431 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_431 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_420, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_431
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_430.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_430.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_431.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_6_419
// TAST (Let): Apply0_7_432 -> *Constructor_Control_Apply_Apply
Apply0_7_432 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_419, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_432
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_432.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_419, "pure"), f_8), a_9)
})
}))
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_445.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_445.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_446.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_3_393
// TAST (Let): Bind1_4_447 -> *Constructor_Control_Bind_Bind
Bind1_4_447 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_393, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_447
// TAST (Let): Applicative0_5_448 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_448 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_393, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_448
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_255 -> gopurs_runtime.Value
__local_var_4_255 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_282 -> gopurs_runtime.Value
__local_var_5_282 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_294 -> gopurs_runtime.Value
__local_var_7_294 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_303 -> gopurs_runtime.Value
__local_var_9_303 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_9_303
// TAST (Let): Bind1_10_304 -> *Constructor_Control_Bind_Bind
Bind1_10_304 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_303, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_304
// TAST (Let): Applicative0_11_305 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_305 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_303, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_305
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_295 -> gopurs_runtime.Value
__local_var_10_295 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_296 -> gopurs_runtime.Value
__local_var_11_296 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_297 -> gopurs_runtime.Value
__local_var_13_297 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_297
// TAST (Let): Bind1_14_298 -> *Constructor_Control_Bind_Bind
Bind1_14_298 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_297, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_298
// TAST (Let): Applicative0_15_299 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_299 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_297, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_299
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_298.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_298.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_299.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_296
// TAST (Let): Bind1_12_300 -> *Constructor_Control_Bind_Bind
Bind1_12_300 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_296, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_300
// TAST (Let): Applicative0_13_301 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_301 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_296, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_301
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_300.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_300.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_301.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_10_295
// TAST (Let): Apply0_11_302 -> *Constructor_Control_Apply_Apply
Apply0_11_302 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_295, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_11_302
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_11_302.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_295, "pure"), f_12), a_13)
})
}))
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_304.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_304.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_305.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_7_294
// TAST (Let): Bind1_8_306 -> *Constructor_Control_Bind_Bind
Bind1_8_306 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_294, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_306
// TAST (Let): Applicative0_9_307 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_307 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_294, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_307
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_283 -> gopurs_runtime.Value
__local_var_8_283 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_284 -> gopurs_runtime.Value
__local_var_9_284 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_285 -> gopurs_runtime.Value
__local_var_11_285 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_286 -> gopurs_runtime.Value
__local_var_13_286 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_286
// TAST (Let): Bind1_14_287 -> *Constructor_Control_Bind_Bind
Bind1_14_287 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_286, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_287
// TAST (Let): Applicative0_15_288 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_288 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_286, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_288
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_287.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_287.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_288.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_285
// TAST (Let): Bind1_12_289 -> *Constructor_Control_Bind_Bind
Bind1_12_289 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_285, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_289
// TAST (Let): Applicative0_13_290 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_290 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_285, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_290
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_289.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_289.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_290.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_284
// TAST (Let): Bind1_10_291 -> *Constructor_Control_Bind_Bind
Bind1_10_291 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_284, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_291
// TAST (Let): Applicative0_11_292 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_292 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_284, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_292
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_291.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_291.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_292.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_8_283
// TAST (Let): Apply0_9_293 -> *Constructor_Control_Apply_Apply
Apply0_9_293 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_283, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_9_293
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_9_293.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_283, "pure"), f_10), a_11)
})
}))
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_306.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_306.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_307.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_334 -> gopurs_runtime.Value
__local_var_7_334 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_346 -> gopurs_runtime.Value
__local_var_9_346 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_355 -> gopurs_runtime.Value
__local_var_11_355 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_355
// TAST (Let): Bind1_12_356 -> *Constructor_Control_Bind_Bind
Bind1_12_356 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_355, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_356
// TAST (Let): Applicative0_13_357 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_357 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_355, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_357
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_347 -> gopurs_runtime.Value
__local_var_12_347 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_348 -> gopurs_runtime.Value
__local_var_13_348 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_349 -> gopurs_runtime.Value
__local_var_15_349 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_15_349
// TAST (Let): Bind1_16_350 -> *Constructor_Control_Bind_Bind
Bind1_16_350 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_349, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_350
// TAST (Let): Applicative0_17_351 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_351 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_349, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_351
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_350.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_350.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_351.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_13_348
// TAST (Let): Bind1_14_352 -> *Constructor_Control_Bind_Bind
Bind1_14_352 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_348, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_352
// TAST (Let): Applicative0_15_353 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_353 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_348, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_353
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_352.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_352.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_353.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_12_347
// TAST (Let): Apply0_13_354 -> *Constructor_Control_Apply_Apply
Apply0_13_354 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_347, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_13_354
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_13_354.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_347, "pure"), f_14), a_15)
})
}))
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_356.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_356.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_357.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_346
// TAST (Let): Bind1_10_358 -> *Constructor_Control_Bind_Bind
Bind1_10_358 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_346, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_358
// TAST (Let): Applicative0_11_359 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_359 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_346, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_359
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_335 -> gopurs_runtime.Value
__local_var_10_335 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_336 -> gopurs_runtime.Value
__local_var_11_336 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_337 -> gopurs_runtime.Value
__local_var_13_337 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_338 -> gopurs_runtime.Value
__local_var_15_338 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_15_338
// TAST (Let): Bind1_16_339 -> *Constructor_Control_Bind_Bind
Bind1_16_339 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_338, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_339
// TAST (Let): Applicative0_17_340 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_340 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_338, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_340
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_339.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_339.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_340.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_13_337
// TAST (Let): Bind1_14_341 -> *Constructor_Control_Bind_Bind
Bind1_14_341 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_337, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_341
// TAST (Let): Applicative0_15_342 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_342 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_337, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_342
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_341.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_341.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_342.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_336
// TAST (Let): Bind1_12_343 -> *Constructor_Control_Bind_Bind
Bind1_12_343 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_336, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_343
// TAST (Let): Applicative0_13_344 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_344 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_336, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_344
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_343.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_343.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_344.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_10_335
// TAST (Let): Apply0_11_345 -> *Constructor_Control_Apply_Apply
Apply0_11_345 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_335, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_11_345
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_11_345.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_335, "pure"), f_12), a_13)
})
}))
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_358.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_358.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_359.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_374 -> gopurs_runtime.Value
__local_var_9_374 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_383 -> gopurs_runtime.Value
__local_var_11_383 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_383
// TAST (Let): Bind1_12_384 -> *Constructor_Control_Bind_Bind
Bind1_12_384 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_383, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_384
// TAST (Let): Applicative0_13_385 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_385 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_383, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_385
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_375 -> gopurs_runtime.Value
__local_var_12_375 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_376 -> gopurs_runtime.Value
__local_var_13_376 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_377 -> gopurs_runtime.Value
__local_var_15_377 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_15_377
// TAST (Let): Bind1_16_378 -> *Constructor_Control_Bind_Bind
Bind1_16_378 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_377, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_378
// TAST (Let): Applicative0_17_379 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_379 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_377, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_379
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_378.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_378.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_379.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_13_376
// TAST (Let): Bind1_14_380 -> *Constructor_Control_Bind_Bind
Bind1_14_380 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_376, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_380
// TAST (Let): Applicative0_15_381 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_381 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_376, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_381
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_380.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_380.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_381.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_12_375
// TAST (Let): Apply0_13_382 -> *Constructor_Control_Apply_Apply
Apply0_13_382 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_375, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_13_382
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_13_382.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_375, "pure"), f_14), a_15)
})
}))
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_384.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_384.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_385.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_9_374
// TAST (Let): Bind1_10_386 -> *Constructor_Control_Bind_Bind
Bind1_10_386 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_374, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_386
// TAST (Let): Applicative0_11_387 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_387 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_374, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_387
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_360 -> gopurs_runtime.Value
__local_var_10_360 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_361 -> gopurs_runtime.Value
__local_var_11_361 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_362 -> gopurs_runtime.Value
__local_var_13_362 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_362
// TAST (Let): Bind1_14_363 -> *Constructor_Control_Bind_Bind
Bind1_14_363 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_362, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_363
// TAST (Let): Applicative0_15_364 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_364 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_362, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_364
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_363.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_363.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_364.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_365 -> gopurs_runtime.Value
__local_var_13_365 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_366 -> gopurs_runtime.Value
__local_var_15_366 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_15_366
// TAST (Let): Bind1_16_367 -> *Constructor_Control_Bind_Bind
Bind1_16_367 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_366, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_367
// TAST (Let): Applicative0_17_368 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_368 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_366, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_368
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_367.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_367.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_368.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_365
// TAST (Let): Bind1_14_369 -> *Constructor_Control_Bind_Bind
Bind1_14_369 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_365, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_369
// TAST (Let): Applicative0_15_370 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_370 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_365, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_370
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_369.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_369.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_370.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_361
// TAST (Let): Bind1_12_371 -> *Constructor_Control_Bind_Bind
Bind1_12_371 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_361, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_371
// TAST (Let): Applicative0_13_372 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_372 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_361, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_372
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_371.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_371.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_372.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_10_360
// TAST (Let): Apply0_11_373 -> *Constructor_Control_Apply_Apply
Apply0_11_373 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_360, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_11_373
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_11_373.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_360, "pure"), f_12), a_13)
})
}))
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_386.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_386.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_387.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_7_334
// TAST (Let): Bind1_8_388 -> *Constructor_Control_Bind_Bind
Bind1_8_388 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_334, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_388
// TAST (Let): Applicative0_9_389 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_389 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_334, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_389
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_308 -> gopurs_runtime.Value
__local_var_8_308 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_309 -> gopurs_runtime.Value
__local_var_9_309 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_310 -> gopurs_runtime.Value
__local_var_11_310 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_311 -> gopurs_runtime.Value
__local_var_13_311 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_311
// TAST (Let): Bind1_14_312 -> *Constructor_Control_Bind_Bind
Bind1_14_312 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_311, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_312
// TAST (Let): Applicative0_15_313 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_313 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_311, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_313
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_312.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_312.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_313.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_310
// TAST (Let): Bind1_12_314 -> *Constructor_Control_Bind_Bind
Bind1_12_314 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_310, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_314
// TAST (Let): Applicative0_13_315 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_315 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_310, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_315
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_314.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_314.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_315.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_316 -> gopurs_runtime.Value
__local_var_11_316 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_317 -> gopurs_runtime.Value
__local_var_13_317 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_318 -> gopurs_runtime.Value
__local_var_15_318 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_15_318
// TAST (Let): Bind1_16_319 -> *Constructor_Control_Bind_Bind
Bind1_16_319 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_318, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_319
// TAST (Let): Applicative0_17_320 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_320 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_318, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_320
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_319.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_319.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_320.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_13_317
// TAST (Let): Bind1_14_321 -> *Constructor_Control_Bind_Bind
Bind1_14_321 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_317, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_321
// TAST (Let): Applicative0_15_322 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_322 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_317, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_322
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_321.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_321.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_322.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_323 -> gopurs_runtime.Value
__local_var_13_323 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_324 -> gopurs_runtime.Value
__local_var_15_324 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_15_324
// TAST (Let): Bind1_16_325 -> *Constructor_Control_Bind_Bind
Bind1_16_325 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_324, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_325
// TAST (Let): Applicative0_17_326 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_326 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_324, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_326
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_325.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_325.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_326.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_323
// TAST (Let): Bind1_14_327 -> *Constructor_Control_Bind_Bind
Bind1_14_327 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_323, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_327
// TAST (Let): Applicative0_15_328 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_328 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_323, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_328
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_327.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_327.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_328.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_316
// TAST (Let): Bind1_12_329 -> *Constructor_Control_Bind_Bind
Bind1_12_329 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_316, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_329
// TAST (Let): Applicative0_13_330 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_330 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_316, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_330
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_329.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_329.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_330.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_309
// TAST (Let): Bind1_10_331 -> *Constructor_Control_Bind_Bind
Bind1_10_331 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_309, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_331
// TAST (Let): Applicative0_11_332 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_332 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_309, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_332
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_331.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_331.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_332.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_8_308
// TAST (Let): Apply0_9_333 -> *Constructor_Control_Apply_Apply
Apply0_9_333 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_308, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_9_333
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_9_333.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_308, "pure"), f_10), a_11)
})
}))
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_388.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_388.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_389.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_282
// TAST (Let): Bind1_6_390 -> *Constructor_Control_Bind_Bind
Bind1_6_390 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_282, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_390
// TAST (Let): Applicative0_7_391 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_391 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_282, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_391
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_256 -> gopurs_runtime.Value
__local_var_6_256 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_257 -> gopurs_runtime.Value
__local_var_7_257 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_258 -> gopurs_runtime.Value
__local_var_9_258 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_259 -> gopurs_runtime.Value
__local_var_11_259 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_259
// TAST (Let): Bind1_12_260 -> *Constructor_Control_Bind_Bind
Bind1_12_260 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_259, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_260
// TAST (Let): Applicative0_13_261 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_261 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_259, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_261
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_260.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_260.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_261.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_258
// TAST (Let): Bind1_10_262 -> *Constructor_Control_Bind_Bind
Bind1_10_262 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_258, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_262
// TAST (Let): Applicative0_11_263 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_263 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_258, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_263
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_262.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_262.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_263.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_264 -> gopurs_runtime.Value
__local_var_9_264 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_265 -> gopurs_runtime.Value
__local_var_11_265 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_266 -> gopurs_runtime.Value
__local_var_13_266 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_266
// TAST (Let): Bind1_14_267 -> *Constructor_Control_Bind_Bind
Bind1_14_267 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_266, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_267
// TAST (Let): Applicative0_15_268 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_268 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_266, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_268
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_267.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_267.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_268.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_265
// TAST (Let): Bind1_12_269 -> *Constructor_Control_Bind_Bind
Bind1_12_269 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_265, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_269
// TAST (Let): Applicative0_13_270 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_270 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_265, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_270
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_269.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_269.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_270.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_271 -> gopurs_runtime.Value
__local_var_11_271 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_272 -> gopurs_runtime.Value
__local_var_13_272 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_272
// TAST (Let): Bind1_14_273 -> *Constructor_Control_Bind_Bind
Bind1_14_273 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_272, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_273
// TAST (Let): Applicative0_15_274 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_274 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_272, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_274
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_273.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_273.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_274.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_271
// TAST (Let): Bind1_12_275 -> *Constructor_Control_Bind_Bind
Bind1_12_275 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_271, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_275
// TAST (Let): Applicative0_13_276 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_276 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_271, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_276
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_275.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_275.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_276.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_264
// TAST (Let): Bind1_10_277 -> *Constructor_Control_Bind_Bind
Bind1_10_277 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_264, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_277
// TAST (Let): Applicative0_11_278 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_278 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_264, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_278
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_277.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_277.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_278.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_7_257
// TAST (Let): Bind1_8_279 -> *Constructor_Control_Bind_Bind
Bind1_8_279 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_257, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_279
// TAST (Let): Applicative0_9_280 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_280 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_257, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_280
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_279.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_279.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_280.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_6_256
// TAST (Let): Apply0_7_281 -> *Constructor_Control_Apply_Apply
Apply0_7_281 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_256, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_281
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_281.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_256, "pure"), f_8), a_9)
})
}))
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_390.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_390.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_391.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_4_255
// TAST (Let): Apply0_5_392 -> *Constructor_Control_Apply_Apply
Apply0_5_392 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_255, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_392
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_392.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_255, "pure"), f_6), a_7)
})
}))
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_447.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_447.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_448.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_463 -> gopurs_runtime.Value
__local_var_3_463 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_472 -> gopurs_runtime.Value
__local_var_5_472 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_5_472
// TAST (Let): Bind1_6_473 -> *Constructor_Control_Bind_Bind
Bind1_6_473 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_472, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_473
// TAST (Let): Applicative0_7_474 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_474 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_472, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_474
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_464 -> gopurs_runtime.Value
__local_var_6_464 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_465 -> gopurs_runtime.Value
__local_var_7_465 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_466 -> gopurs_runtime.Value
__local_var_9_466 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_9_466
// TAST (Let): Bind1_10_467 -> *Constructor_Control_Bind_Bind
Bind1_10_467 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_466, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_467
// TAST (Let): Applicative0_11_468 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_468 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_466, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_468
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_467.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_467.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_468.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_7_465
// TAST (Let): Bind1_8_469 -> *Constructor_Control_Bind_Bind
Bind1_8_469 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_465, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_469
// TAST (Let): Applicative0_9_470 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_470 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_465, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_470
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_469.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_469.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_470.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_6_464
// TAST (Let): Apply0_7_471 -> *Constructor_Control_Apply_Apply
Apply0_7_471 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_464, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_471
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_471.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_464, "pure"), f_8), a_9)
})
}))
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_473.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_473.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_474.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_3_463
// TAST (Let): Bind1_4_475 -> *Constructor_Control_Bind_Bind
Bind1_4_475 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_463, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_475
// TAST (Let): Applicative0_5_476 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_476 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_463, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_476
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_449 -> gopurs_runtime.Value
__local_var_4_449 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_450 -> gopurs_runtime.Value
__local_var_5_450 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_451 -> gopurs_runtime.Value
__local_var_7_451 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_451
// TAST (Let): Bind1_8_452 -> *Constructor_Control_Bind_Bind
Bind1_8_452 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_451, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_452
// TAST (Let): Applicative0_9_453 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_453 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_451, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_453
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_452.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_452.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_453.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_454 -> gopurs_runtime.Value
__local_var_7_454 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_455 -> gopurs_runtime.Value
__local_var_9_455 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_9_455
// TAST (Let): Bind1_10_456 -> *Constructor_Control_Bind_Bind
Bind1_10_456 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_455, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_456
// TAST (Let): Applicative0_11_457 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_457 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_455, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_457
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_456.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_456.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_457.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_454
// TAST (Let): Bind1_8_458 -> *Constructor_Control_Bind_Bind
Bind1_8_458 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_454, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_458
// TAST (Let): Applicative0_9_459 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_459 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_454, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_459
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_458.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_458.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_459.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_450
// TAST (Let): Bind1_6_460 -> *Constructor_Control_Bind_Bind
Bind1_6_460 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_450, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_460
// TAST (Let): Applicative0_7_461 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_461 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_450, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_461
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_460.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_460.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_461.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_4_449
// TAST (Let): Apply0_5_462 -> *Constructor_Control_Apply_Apply
Apply0_5_462 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_449, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_462
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_462.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_449, "pure"), f_6), a_7)
})
}))
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_475.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_475.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_476.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_1_254
// TAST (Let): Bind1_2_477 -> *Constructor_Control_Bind_Bind
Bind1_2_477 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_254, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_477
// TAST (Let): Applicative0_3_478 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_478 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_254, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_478
// TAST (Let): __local_var_1_1 -> *Constructor_Control_Apply_Apply
__local_var_1_1 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_29 -> gopurs_runtime.Value
__local_var_3_29 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_56 -> gopurs_runtime.Value
__local_var_5_56 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_68 -> gopurs_runtime.Value
__local_var_7_68 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_77 -> gopurs_runtime.Value
__local_var_9_77 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_9_77
// TAST (Let): Bind1_10_78 -> *Constructor_Control_Bind_Bind
Bind1_10_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_77, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_78
// TAST (Let): Applicative0_11_79 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_79 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_77, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_79
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_69 -> gopurs_runtime.Value
__local_var_10_69 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_70 -> gopurs_runtime.Value
__local_var_11_70 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_71 -> gopurs_runtime.Value
__local_var_13_71 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_71
// TAST (Let): Bind1_14_72 -> *Constructor_Control_Bind_Bind
Bind1_14_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_71, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_72
// TAST (Let): Applicative0_15_73 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_71, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_73
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_72.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_72.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_73.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_70
// TAST (Let): Bind1_12_74 -> *Constructor_Control_Bind_Bind
Bind1_12_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_70, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_74
// TAST (Let): Applicative0_13_75 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_75 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_70, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_75
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_74.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_74.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_75.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_10_69
// TAST (Let): Apply0_11_76 -> *Constructor_Control_Apply_Apply
Apply0_11_76 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_69, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_11_76
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_11_76.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_69, "pure"), f_12), a_13)
})
}))
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_78.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_78.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_79.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_7_68
// TAST (Let): Bind1_8_80 -> *Constructor_Control_Bind_Bind
Bind1_8_80 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_68, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_80
// TAST (Let): Applicative0_9_81 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_68, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_81
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_57 -> gopurs_runtime.Value
__local_var_8_57 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_58 -> gopurs_runtime.Value
__local_var_9_58 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_59 -> gopurs_runtime.Value
__local_var_11_59 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_60 -> gopurs_runtime.Value
__local_var_13_60 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_60
// TAST (Let): Bind1_14_61 -> *Constructor_Control_Bind_Bind
Bind1_14_61 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_60, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_61
// TAST (Let): Applicative0_15_62 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_62 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_60, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_62
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_61.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_61.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_62.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_59
// TAST (Let): Bind1_12_63 -> *Constructor_Control_Bind_Bind
Bind1_12_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_59, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_63
// TAST (Let): Applicative0_13_64 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_59, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_64
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_63.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_63.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_64.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_58
// TAST (Let): Bind1_10_65 -> *Constructor_Control_Bind_Bind
Bind1_10_65 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_58, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_65
// TAST (Let): Applicative0_11_66 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_66 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_58, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_66
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_65.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_65.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_66.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_8_57
// TAST (Let): Apply0_9_67 -> *Constructor_Control_Apply_Apply
Apply0_9_67 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_57, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_9_67
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_9_67.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_57, "pure"), f_10), a_11)
})
}))
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_80.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_80.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_81.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_96 -> gopurs_runtime.Value
__local_var_7_96 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_105 -> gopurs_runtime.Value
__local_var_9_105 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_9_105
// TAST (Let): Bind1_10_106 -> *Constructor_Control_Bind_Bind
Bind1_10_106 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_105, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_106
// TAST (Let): Applicative0_11_107 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_107 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_105, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_107
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_97 -> gopurs_runtime.Value
__local_var_10_97 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_98 -> gopurs_runtime.Value
__local_var_11_98 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_99 -> gopurs_runtime.Value
__local_var_13_99 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_99
// TAST (Let): Bind1_14_100 -> *Constructor_Control_Bind_Bind
Bind1_14_100 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_99, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_100
// TAST (Let): Applicative0_15_101 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_101 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_99, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_101
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_100.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_100.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_101.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_98
// TAST (Let): Bind1_12_102 -> *Constructor_Control_Bind_Bind
Bind1_12_102 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_98, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_102
// TAST (Let): Applicative0_13_103 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_103 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_98, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_103
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_102.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_102.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_103.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_10_97
// TAST (Let): Apply0_11_104 -> *Constructor_Control_Apply_Apply
Apply0_11_104 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_97, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_11_104
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_11_104.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_97, "pure"), f_12), a_13)
})
}))
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_106.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_106.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_107.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_96
// TAST (Let): Bind1_8_108 -> *Constructor_Control_Bind_Bind
Bind1_8_108 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_96, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_108
// TAST (Let): Applicative0_9_109 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_109 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_96, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_109
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_82 -> gopurs_runtime.Value
__local_var_8_82 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_83 -> gopurs_runtime.Value
__local_var_9_83 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_84 -> gopurs_runtime.Value
__local_var_11_84 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_84
// TAST (Let): Bind1_12_85 -> *Constructor_Control_Bind_Bind
Bind1_12_85 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_84, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_85
// TAST (Let): Applicative0_13_86 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_86 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_84, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_86
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_85.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_85.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_86.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_87 -> gopurs_runtime.Value
__local_var_11_87 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_88 -> gopurs_runtime.Value
__local_var_13_88 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_88
// TAST (Let): Bind1_14_89 -> *Constructor_Control_Bind_Bind
Bind1_14_89 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_88, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_89
// TAST (Let): Applicative0_15_90 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_90 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_88, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_90
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_89.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_89.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_90.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_87
// TAST (Let): Bind1_12_91 -> *Constructor_Control_Bind_Bind
Bind1_12_91 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_87, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_91
// TAST (Let): Applicative0_13_92 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_92 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_87, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_92
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_91.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_91.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_92.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_83
// TAST (Let): Bind1_10_93 -> *Constructor_Control_Bind_Bind
Bind1_10_93 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_83, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_93
// TAST (Let): Applicative0_11_94 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_94 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_83, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_94
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_93.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_93.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_94.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_8_82
// TAST (Let): Apply0_9_95 -> *Constructor_Control_Apply_Apply
Apply0_9_95 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_82, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_9_95
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_9_95.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_82, "pure"), f_10), a_11)
})
}))
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_108.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_108.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_109.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_56
// TAST (Let): Bind1_6_110 -> *Constructor_Control_Bind_Bind
Bind1_6_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_56, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_110
// TAST (Let): Applicative0_7_111 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_111 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_56, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_111
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_30 -> gopurs_runtime.Value
__local_var_6_30 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_31 -> gopurs_runtime.Value
__local_var_7_31 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_32 -> gopurs_runtime.Value
__local_var_9_32 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_33 -> gopurs_runtime.Value
__local_var_11_33 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_33
// TAST (Let): Bind1_12_34 -> *Constructor_Control_Bind_Bind
Bind1_12_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_33, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_34
// TAST (Let): Applicative0_13_35 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_33, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_35
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_34.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_34.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_35.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_32
// TAST (Let): Bind1_10_36 -> *Constructor_Control_Bind_Bind
Bind1_10_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_32, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_36
// TAST (Let): Applicative0_11_37 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_32, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_37
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_36.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_36.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_37.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_38 -> gopurs_runtime.Value
__local_var_9_38 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_39 -> gopurs_runtime.Value
__local_var_11_39 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_40 -> gopurs_runtime.Value
__local_var_13_40 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_40
// TAST (Let): Bind1_14_41 -> *Constructor_Control_Bind_Bind
Bind1_14_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_40, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_41
// TAST (Let): Applicative0_15_42 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_40, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_42
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_41.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_41.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_42.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_39
// TAST (Let): Bind1_12_43 -> *Constructor_Control_Bind_Bind
Bind1_12_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_39, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_43
// TAST (Let): Applicative0_13_44 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_39, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_44
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_43.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_43.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_44.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_45 -> gopurs_runtime.Value
__local_var_11_45 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_46 -> gopurs_runtime.Value
__local_var_13_46 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_46
// TAST (Let): Bind1_14_47 -> *Constructor_Control_Bind_Bind
Bind1_14_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_46, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_47
// TAST (Let): Applicative0_15_48 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_46, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_48
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_47.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_47.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_48.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_45
// TAST (Let): Bind1_12_49 -> *Constructor_Control_Bind_Bind
Bind1_12_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_45, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_49
// TAST (Let): Applicative0_13_50 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_45, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_50
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_49.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_49.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_50.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_38
// TAST (Let): Bind1_10_51 -> *Constructor_Control_Bind_Bind
Bind1_10_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_38, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_51
// TAST (Let): Applicative0_11_52 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_52 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_38, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_52
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_51.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_51.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_52.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_7_31
// TAST (Let): Bind1_8_53 -> *Constructor_Control_Bind_Bind
Bind1_8_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_31, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_53
// TAST (Let): Applicative0_9_54 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_31, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_54
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_53.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_53.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_54.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_6_30
// TAST (Let): Apply0_7_55 -> *Constructor_Control_Apply_Apply
Apply0_7_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_30, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_55
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_55.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_30, "pure"), f_8), a_9)
})
}))
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_110.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_110.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_111.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_138 -> gopurs_runtime.Value
__local_var_5_138 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_165 -> gopurs_runtime.Value
__local_var_7_165 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_177 -> gopurs_runtime.Value
__local_var_9_177 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_186 -> gopurs_runtime.Value
__local_var_11_186 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_186
// TAST (Let): Bind1_12_187 -> *Constructor_Control_Bind_Bind
Bind1_12_187 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_186, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_187
// TAST (Let): Applicative0_13_188 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_188 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_186, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_188
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_178 -> gopurs_runtime.Value
__local_var_12_178 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_179 -> gopurs_runtime.Value
__local_var_13_179 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_180 -> gopurs_runtime.Value
__local_var_15_180 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_15_180
// TAST (Let): Bind1_16_181 -> *Constructor_Control_Bind_Bind
Bind1_16_181 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_180, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_181
// TAST (Let): Applicative0_17_182 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_182 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_180, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_182
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_181.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_181.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_182.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_13_179
// TAST (Let): Bind1_14_183 -> *Constructor_Control_Bind_Bind
Bind1_14_183 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_179, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_183
// TAST (Let): Applicative0_15_184 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_184 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_179, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_184
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_183.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_183.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_184.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_12_178
// TAST (Let): Apply0_13_185 -> *Constructor_Control_Apply_Apply
Apply0_13_185 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_178, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_13_185
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_13_185.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_178, "pure"), f_14), a_15)
})
}))
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_187.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_187.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_188.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_177
// TAST (Let): Bind1_10_189 -> *Constructor_Control_Bind_Bind
Bind1_10_189 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_177, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_189
// TAST (Let): Applicative0_11_190 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_190 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_177, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_190
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_166 -> gopurs_runtime.Value
__local_var_10_166 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_167 -> gopurs_runtime.Value
__local_var_11_167 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_168 -> gopurs_runtime.Value
__local_var_13_168 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_169 -> gopurs_runtime.Value
__local_var_15_169 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_15_169
// TAST (Let): Bind1_16_170 -> *Constructor_Control_Bind_Bind
Bind1_16_170 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_169, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_170
// TAST (Let): Applicative0_17_171 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_171 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_169, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_171
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_170.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_170.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_171.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_13_168
// TAST (Let): Bind1_14_172 -> *Constructor_Control_Bind_Bind
Bind1_14_172 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_168, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_172
// TAST (Let): Applicative0_15_173 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_173 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_168, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_173
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_172.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_172.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_173.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_167
// TAST (Let): Bind1_12_174 -> *Constructor_Control_Bind_Bind
Bind1_12_174 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_167, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_174
// TAST (Let): Applicative0_13_175 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_175 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_167, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_175
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_174.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_174.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_175.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_10_166
// TAST (Let): Apply0_11_176 -> *Constructor_Control_Apply_Apply
Apply0_11_176 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_166, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_11_176
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_11_176.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_166, "pure"), f_12), a_13)
})
}))
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_189.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_189.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_190.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_205 -> gopurs_runtime.Value
__local_var_9_205 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_214 -> gopurs_runtime.Value
__local_var_11_214 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_214
// TAST (Let): Bind1_12_215 -> *Constructor_Control_Bind_Bind
Bind1_12_215 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_214, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_215
// TAST (Let): Applicative0_13_216 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_216 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_214, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_216
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_206 -> gopurs_runtime.Value
__local_var_12_206 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_207 -> gopurs_runtime.Value
__local_var_13_207 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_208 -> gopurs_runtime.Value
__local_var_15_208 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_15_208
// TAST (Let): Bind1_16_209 -> *Constructor_Control_Bind_Bind
Bind1_16_209 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_208, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_209
// TAST (Let): Applicative0_17_210 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_210 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_208, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_210
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_209.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_209.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_210.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_13_207
// TAST (Let): Bind1_14_211 -> *Constructor_Control_Bind_Bind
Bind1_14_211 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_207, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_211
// TAST (Let): Applicative0_15_212 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_212 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_207, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_212
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_211.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_211.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_212.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_12_206
// TAST (Let): Apply0_13_213 -> *Constructor_Control_Apply_Apply
Apply0_13_213 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_206, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_13_213
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_13_213.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_206, "pure"), f_14), a_15)
})
}))
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_215.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_215.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_216.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_9_205
// TAST (Let): Bind1_10_217 -> *Constructor_Control_Bind_Bind
Bind1_10_217 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_205, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_217
// TAST (Let): Applicative0_11_218 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_218 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_205, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_218
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_191 -> gopurs_runtime.Value
__local_var_10_191 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_192 -> gopurs_runtime.Value
__local_var_11_192 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_193 -> gopurs_runtime.Value
__local_var_13_193 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_193
// TAST (Let): Bind1_14_194 -> *Constructor_Control_Bind_Bind
Bind1_14_194 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_193, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_194
// TAST (Let): Applicative0_15_195 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_195 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_193, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_195
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_194.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_194.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_195.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_196 -> gopurs_runtime.Value
__local_var_13_196 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_197 -> gopurs_runtime.Value
__local_var_15_197 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_15_197
// TAST (Let): Bind1_16_198 -> *Constructor_Control_Bind_Bind
Bind1_16_198 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_197, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_198
// TAST (Let): Applicative0_17_199 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_199 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_197, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_199
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_198.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_198.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_199.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_196
// TAST (Let): Bind1_14_200 -> *Constructor_Control_Bind_Bind
Bind1_14_200 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_196, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_200
// TAST (Let): Applicative0_15_201 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_201 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_196, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_201
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_200.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_200.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_201.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_192
// TAST (Let): Bind1_12_202 -> *Constructor_Control_Bind_Bind
Bind1_12_202 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_192, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_202
// TAST (Let): Applicative0_13_203 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_203 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_192, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_203
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_202.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_202.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_203.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_10_191
// TAST (Let): Apply0_11_204 -> *Constructor_Control_Apply_Apply
Apply0_11_204 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_191, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_11_204
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_11_204.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_191, "pure"), f_12), a_13)
})
}))
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_217.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_217.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_218.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_7_165
// TAST (Let): Bind1_8_219 -> *Constructor_Control_Bind_Bind
Bind1_8_219 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_165, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_219
// TAST (Let): Applicative0_9_220 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_220 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_165, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_220
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_139 -> gopurs_runtime.Value
__local_var_8_139 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_140 -> gopurs_runtime.Value
__local_var_9_140 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_141 -> gopurs_runtime.Value
__local_var_11_141 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_142 -> gopurs_runtime.Value
__local_var_13_142 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_142
// TAST (Let): Bind1_14_143 -> *Constructor_Control_Bind_Bind
Bind1_14_143 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_142, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_143
// TAST (Let): Applicative0_15_144 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_144 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_142, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_144
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_143.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_143.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_144.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_141
// TAST (Let): Bind1_12_145 -> *Constructor_Control_Bind_Bind
Bind1_12_145 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_141, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_145
// TAST (Let): Applicative0_13_146 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_146 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_141, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_146
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_145.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_145.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_146.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_147 -> gopurs_runtime.Value
__local_var_11_147 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_148 -> gopurs_runtime.Value
__local_var_13_148 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_149 -> gopurs_runtime.Value
__local_var_15_149 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_15_149
// TAST (Let): Bind1_16_150 -> *Constructor_Control_Bind_Bind
Bind1_16_150 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_149, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_150
// TAST (Let): Applicative0_17_151 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_151 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_149, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_151
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_150.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_150.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_151.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_13_148
// TAST (Let): Bind1_14_152 -> *Constructor_Control_Bind_Bind
Bind1_14_152 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_148, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_152
// TAST (Let): Applicative0_15_153 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_153 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_148, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_153
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_152.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_152.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_153.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_154 -> gopurs_runtime.Value
__local_var_13_154 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_155 -> gopurs_runtime.Value
__local_var_15_155 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_15_155
// TAST (Let): Bind1_16_156 -> *Constructor_Control_Bind_Bind
Bind1_16_156 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_155, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_156
// TAST (Let): Applicative0_17_157 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_157 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_155, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_157
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_156.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_156.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_157.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_154
// TAST (Let): Bind1_14_158 -> *Constructor_Control_Bind_Bind
Bind1_14_158 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_154, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_158
// TAST (Let): Applicative0_15_159 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_159 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_154, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_159
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_158.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_158.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_159.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_147
// TAST (Let): Bind1_12_160 -> *Constructor_Control_Bind_Bind
Bind1_12_160 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_147, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_160
// TAST (Let): Applicative0_13_161 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_161 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_147, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_161
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_160.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_160.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_161.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_140
// TAST (Let): Bind1_10_162 -> *Constructor_Control_Bind_Bind
Bind1_10_162 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_140, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_162
// TAST (Let): Applicative0_11_163 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_163 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_140, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_163
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_162.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_162.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_163.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_8_139
// TAST (Let): Apply0_9_164 -> *Constructor_Control_Apply_Apply
Apply0_9_164 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_139, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_9_164
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_9_164.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_139, "pure"), f_10), a_11)
})
}))
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_219.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_219.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_220.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_235 -> gopurs_runtime.Value
__local_var_7_235 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_244 -> gopurs_runtime.Value
__local_var_9_244 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_9_244
// TAST (Let): Bind1_10_245 -> *Constructor_Control_Bind_Bind
Bind1_10_245 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_244, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_245
// TAST (Let): Applicative0_11_246 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_246 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_244, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_246
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_236 -> gopurs_runtime.Value
__local_var_10_236 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_237 -> gopurs_runtime.Value
__local_var_11_237 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_238 -> gopurs_runtime.Value
__local_var_13_238 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_238
// TAST (Let): Bind1_14_239 -> *Constructor_Control_Bind_Bind
Bind1_14_239 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_238, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_239
// TAST (Let): Applicative0_15_240 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_240 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_238, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_240
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_239.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_239.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_240.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_237
// TAST (Let): Bind1_12_241 -> *Constructor_Control_Bind_Bind
Bind1_12_241 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_237, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_241
// TAST (Let): Applicative0_13_242 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_242 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_237, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_242
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_241.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_241.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_242.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_10_236
// TAST (Let): Apply0_11_243 -> *Constructor_Control_Apply_Apply
Apply0_11_243 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_236, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_11_243
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_11_243.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_236, "pure"), f_12), a_13)
})
}))
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_245.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_245.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_246.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_7_235
// TAST (Let): Bind1_8_247 -> *Constructor_Control_Bind_Bind
Bind1_8_247 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_235, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_247
// TAST (Let): Applicative0_9_248 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_248 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_235, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_248
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_221 -> gopurs_runtime.Value
__local_var_8_221 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_222 -> gopurs_runtime.Value
__local_var_9_222 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_223 -> gopurs_runtime.Value
__local_var_11_223 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_223
// TAST (Let): Bind1_12_224 -> *Constructor_Control_Bind_Bind
Bind1_12_224 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_223, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_224
// TAST (Let): Applicative0_13_225 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_225 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_223, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_225
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_224.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_224.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_225.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_226 -> gopurs_runtime.Value
__local_var_11_226 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_227 -> gopurs_runtime.Value
__local_var_13_227 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_227
// TAST (Let): Bind1_14_228 -> *Constructor_Control_Bind_Bind
Bind1_14_228 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_227, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_228
// TAST (Let): Applicative0_15_229 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_229 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_227, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_229
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_228.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_228.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_229.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_226
// TAST (Let): Bind1_12_230 -> *Constructor_Control_Bind_Bind
Bind1_12_230 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_226, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_230
// TAST (Let): Applicative0_13_231 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_231 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_226, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_231
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_230.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_230.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_231.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_222
// TAST (Let): Bind1_10_232 -> *Constructor_Control_Bind_Bind
Bind1_10_232 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_222, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_232
// TAST (Let): Applicative0_11_233 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_233 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_222, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_233
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_232.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_232.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_233.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_8_221
// TAST (Let): Apply0_9_234 -> *Constructor_Control_Apply_Apply
Apply0_9_234 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_221, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_9_234
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_9_234.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_221, "pure"), f_10), a_11)
})
}))
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_247.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_247.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_248.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_138
// TAST (Let): Bind1_6_249 -> *Constructor_Control_Bind_Bind
Bind1_6_249 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_138, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_249
// TAST (Let): Applicative0_7_250 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_250 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_138, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_250
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_112 -> gopurs_runtime.Value
__local_var_6_112 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_113 -> gopurs_runtime.Value
__local_var_7_113 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_114 -> gopurs_runtime.Value
__local_var_9_114 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_115 -> gopurs_runtime.Value
__local_var_11_115 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_115
// TAST (Let): Bind1_12_116 -> *Constructor_Control_Bind_Bind
Bind1_12_116 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_115, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_116
// TAST (Let): Applicative0_13_117 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_117 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_115, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_117
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_116.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_116.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_117.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_114
// TAST (Let): Bind1_10_118 -> *Constructor_Control_Bind_Bind
Bind1_10_118 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_114, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_118
// TAST (Let): Applicative0_11_119 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_119 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_114, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_119
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_118.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_118.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_119.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_120 -> gopurs_runtime.Value
__local_var_9_120 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_121 -> gopurs_runtime.Value
__local_var_11_121 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_122 -> gopurs_runtime.Value
__local_var_13_122 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_122
// TAST (Let): Bind1_14_123 -> *Constructor_Control_Bind_Bind
Bind1_14_123 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_122, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_123
// TAST (Let): Applicative0_15_124 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_124 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_122, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_124
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_123.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_123.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_124.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_11_121
// TAST (Let): Bind1_12_125 -> *Constructor_Control_Bind_Bind
Bind1_12_125 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_121, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_125
// TAST (Let): Applicative0_13_126 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_126 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_121, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_126
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_125.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_125.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_126.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_127 -> gopurs_runtime.Value
__local_var_11_127 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_128 -> gopurs_runtime.Value
__local_var_13_128 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_13_128
// TAST (Let): Bind1_14_129 -> *Constructor_Control_Bind_Bind
Bind1_14_129 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_128, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_129
// TAST (Let): Applicative0_15_130 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_130 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_128, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_130
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_129.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_129.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_130.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_127
// TAST (Let): Bind1_12_131 -> *Constructor_Control_Bind_Bind
Bind1_12_131 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_127, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_131
// TAST (Let): Applicative0_13_132 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_132 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_127, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_132
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_131.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_131.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_132.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_120
// TAST (Let): Bind1_10_133 -> *Constructor_Control_Bind_Bind
Bind1_10_133 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_120, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_133
// TAST (Let): Applicative0_11_134 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_134 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_120, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_134
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_133.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_133.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_134.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_7_113
// TAST (Let): Bind1_8_135 -> *Constructor_Control_Bind_Bind
Bind1_8_135 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_113, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_135
// TAST (Let): Applicative0_9_136 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_136 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_113, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_136
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_135.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_135.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_136.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_6_112
// TAST (Let): Apply0_7_137 -> *Constructor_Control_Apply_Apply
Apply0_7_137 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_112, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_137
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_137.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_112, "pure"), f_8), a_9)
})
}))
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_249.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_249.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_250.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_3_29
// TAST (Let): Bind1_4_251 -> *Constructor_Control_Bind_Bind
Bind1_4_251 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_29, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_251
// TAST (Let): Applicative0_5_252 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_252 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_29, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_252
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_6 -> gopurs_runtime.Value
__local_var_9_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_9_6
// TAST (Let): Bind1_10_7 -> *Constructor_Control_Bind_Bind
Bind1_10_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_7
// TAST (Let): Applicative0_11_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_8
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_7.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_7.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_8.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_7_5
// TAST (Let): Bind1_8_9 -> *Constructor_Control_Bind_Bind
Bind1_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_9
// TAST (Let): Applicative0_9_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_10
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_9.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_9.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_10.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_11 -> gopurs_runtime.Value
__local_var_7_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_12 -> gopurs_runtime.Value
__local_var_9_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_13 -> gopurs_runtime.Value
__local_var_11_13 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_13
// TAST (Let): Bind1_12_14 -> *Constructor_Control_Bind_Bind
Bind1_12_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_14
// TAST (Let): Applicative0_13_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_15
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_14.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_14.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_15.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_9_12
// TAST (Let): Bind1_10_16 -> *Constructor_Control_Bind_Bind
Bind1_10_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_16
// TAST (Let): Applicative0_11_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_17
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_16.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_16.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_17.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_18 -> gopurs_runtime.Value
__local_var_9_18 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_19 -> gopurs_runtime.Value
__local_var_11_19 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_11_19
// TAST (Let): Bind1_12_20 -> *Constructor_Control_Bind_Bind
Bind1_12_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_19, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_20
// TAST (Let): Applicative0_13_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_19, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_21
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_20.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_20.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_21.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), Get_Effect_pureE())
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_applyEffect()
}), Get_Effect_bindE())
}))
_ = __local_var_9_18
// TAST (Let): Bind1_10_22 -> *Constructor_Control_Bind_Bind
Bind1_10_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_18, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_22
// TAST (Let): Applicative0_11_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_18, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_23
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_22.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_22.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_23.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_7_11
// TAST (Let): Bind1_8_24 -> *Constructor_Control_Bind_Bind
Bind1_8_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_24
// TAST (Let): Applicative0_9_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_24.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_24.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_25.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), Get_Effect_bindE())
}))
_ = __local_var_5_4
// TAST (Let): Bind1_6_26 -> *Constructor_Control_Bind_Bind
Bind1_6_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_26
// TAST (Let): Applicative0_7_27 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_27
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_functorEffect()
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_26.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_26.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_27.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_4_3
// TAST (Let): Apply0_5_28 -> *Constructor_Control_Apply_Apply
Apply0_5_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_28
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_28.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), f_6), a_7)
})
}))
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_251.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_251.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_252.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), Get_Effect_pureE())
_ = __local_var_2_2
// TAST (Let): Apply0_3_253 -> *Constructor_Control_Apply_Apply
Apply0_3_253 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_253
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_253.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "pure"), f_4), a_5)
})
}))
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_477.V1), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_477.V1), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_478.V1), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
})}
_ = __local_var_1_1
// TAST (Let): Functor0_2_479 -> *Constructor_Data_Functor_Functor
Functor0_2_479 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_1_1.V0), gopurs_runtime.Value{}))
_ = Functor0_2_479
// TAST (Let): __local_var_3_480 -> gopurs_runtime.Value
__local_var_3_480 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_3_480
// TAST (Let): semigroupEffect1_1_0 -> gopurs_runtime.Value
semigroupEffect1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_1_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_479.V0), __local_var_3_480, a_4), b_5)
})
}))
_ = semigroupEffect1_1_0
// TAST (Let): __local_var_2_481 -> gopurs_runtime.Value
__local_var_2_481 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_2_481
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffect1_1_0
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_481
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

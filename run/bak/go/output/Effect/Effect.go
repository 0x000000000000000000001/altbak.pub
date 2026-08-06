package Effect

import (
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
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Bind1"), gopurs_runtime.Value{})
_ = __local_var_0_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEffect()
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "bind"), f_1, gopurs_runtime.Func(func(f_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "bind"), a_2, gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply(f_prime_3, a_prime_4))
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
		cache_functorEffect = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "pure"), f_0), a_1)
})
}))
	})
	return cache_functorEffect
}

var cache_lift2 gopurs_runtime.Value
var once_lift2 sync.Once
func Get_lift2() gopurs_runtime.Value {
	once_lift2.Do(func() {
		cache_lift2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_lift2
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

func Call_lift2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_0, "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applyEffect(), "Functor0"), gopurs_runtime.Value{}), "map"), f_0, a_1), gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_0, "bind"), b_2, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
}

func Call_semigroupEffect(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_lift2(), gopurs_runtime.RecordGet(dictSemigroup_0, "append")))
}

func Call_monoidEffect(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
semigroupEffect1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_lift2(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append")))
_ = semigroupEffect1_1_0
__local_var_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffect1_1_0
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}))
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

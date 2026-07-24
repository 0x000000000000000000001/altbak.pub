package Effect

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var monadEffect gopurs_runtime.Value
var once_monadEffect sync.Once
func Get_monadEffect() gopurs_runtime.Value {
	once_monadEffect.Do(func() {
		monadEffect = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindEffect()
}))
	})
	return monadEffect
}

var bindEffect gopurs_runtime.Value
var once_bindEffect sync.Once
func Get_bindEffect() gopurs_runtime.Value {
	once_bindEffect.Do(func() {
		bindEffect = gopurs_runtime.RecordDict2("bind", "Apply0", Get_bindE(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}))
	})
	return bindEffect
}

var applyEffect gopurs_runtime.Value
var once_applyEffect sync.Once
func Get_applyEffect() gopurs_runtime.Value {
	once_applyEffect.Do(func() {
		applyEffect = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
f_prime_2_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Value{})
_ = f_prime_2_0
a_prime_3_1 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
_ = a_prime_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeEffect(), "pure"), gopurs_runtime.Apply(f_prime_2_0, a_prime_3_1)), gopurs_runtime.Value{})
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEffect()
}))
	})
	return applyEffect
}

var applicativeEffect gopurs_runtime.Value
var once_applicativeEffect sync.Once
func Get_applicativeEffect() gopurs_runtime.Value {
	once_applicativeEffect.Do(func() {
		applicativeEffect = gopurs_runtime.RecordDict2("pure", "Apply0", Get_pureE(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEffect()
}))
	})
	return applicativeEffect
}

var functorEffect gopurs_runtime.Value
var once_functorEffect sync.Once
func Get_functorEffect() gopurs_runtime.Value {
	once_functorEffect.Do(func() {
		functorEffect = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_2_0 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
_ = a_prime_2_0
return gopurs_runtime.Apply(f_0, a_prime_2_0)
})
}))
	})
	return functorEffect
}

var semigroupEffect gopurs_runtime.Value
var once_semigroupEffect sync.Once
func Get_semigroupEffect() gopurs_runtime.Value {
	once_semigroupEffect.Do(func() {
		semigroupEffect = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_3_0 := gopurs_runtime.Apply(a_1, gopurs_runtime.Value{})
_ = a_prime_3_0
a_prime_4_1 := gopurs_runtime.Apply(b_2, gopurs_runtime.Value{})
_ = a_prime_4_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), a_prime_3_0, a_prime_4_1)
})
}))
}()
})
	})
	return semigroupEffect
}

var monoidEffect gopurs_runtime.Value
var once_monoidEffect sync.Once
func Get_monoidEffect() gopurs_runtime.Value {
	once_monoidEffect.Do(func() {
		monoidEffect = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0_loop, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupEffect1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_4_2 := gopurs_runtime.Apply(a_2, gopurs_runtime.Value{})
_ = a_prime_4_2
a_prime_5_3 := gopurs_runtime.Apply(b_3, gopurs_runtime.Value{})
_ = a_prime_5_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), a_prime_4_2, a_prime_5_3)
})
}))
_ = semigroupEffect1_2_1
__local_var_3_4 := gopurs_runtime.RecordGet(dictMonoid_0_loop, "mempty")
_ = __local_var_3_4
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_4
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffect1_2_1
}))
}()
})
	})
	return monoidEffect
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

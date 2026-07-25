package Data_Monoid_Dual

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Dual gopurs_runtime.Value
var once_Dual sync.Once
func Get_Dual() gopurs_runtime.Value {
	once_Dual.Do(func() {
		cache_Dual = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Dual
}

var cache_showDual gopurs_runtime.Value
var once_showDual sync.Once
func Get_showDual() gopurs_runtime.Value {
	once_showDual.Do(func() {
		cache_showDual = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Dual ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}()
})
	})
	return cache_showDual
}

var cache_semigroupDual gopurs_runtime.Value
var once_semigroupDual sync.Once
func Get_semigroupDual() gopurs_runtime.Value {
	once_semigroupDual.Do(func() {
		cache_semigroupDual = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), v1_2, v_1)
}))
}()
})
	})
	return cache_semigroupDual
}

var cache_ordDual gopurs_runtime.Value
var once_ordDual sync.Once
func Get_ordDual() gopurs_runtime.Value {
	once_ordDual.Do(func() {
		cache_ordDual = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}()
})
	})
	return cache_ordDual
}

var cache_monoidDual gopurs_runtime.Value
var once_monoidDual sync.Once
func Get_monoidDual() gopurs_runtime.Value {
	once_monoidDual.Do(func() {
		cache_monoidDual = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupDual1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), v1_3, v_2)
}))
_ = semigroupDual1_2_1
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDual1_2_1
}))
}()
})
	})
	return cache_monoidDual
}

var cache_functorDual gopurs_runtime.Value
var once_functorDual sync.Once
func Get_functorDual() gopurs_runtime.Value {
	once_functorDual.Do(func() {
		cache_functorDual = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
}))
	})
	return cache_functorDual
}

var cache_eqDual gopurs_runtime.Value
var once_eqDual sync.Once
func Get_eqDual() gopurs_runtime.Value {
	once_eqDual.Do(func() {
		cache_eqDual = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}()
})
	})
	return cache_eqDual
}

var cache_eq1Dual gopurs_runtime.Value
var once_eq1Dual sync.Once
func Get_eq1Dual() gopurs_runtime.Value {
	once_eq1Dual.Do(func() {
		cache_eq1Dual = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1Dual
}

var cache_ord1Dual gopurs_runtime.Value
var once_ord1Dual sync.Once
func Get_ord1Dual() gopurs_runtime.Value {
	once_ord1Dual.Do(func() {
		cache_ord1Dual = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Dual()
}))
	})
	return cache_ord1Dual
}

var cache_boundedDual gopurs_runtime.Value
var once_boundedDual sync.Once
func Get_boundedDual() gopurs_runtime.Value {
	once_boundedDual.Do(func() {
		cache_boundedDual = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}()
})
	})
	return cache_boundedDual
}

var cache_applyDual gopurs_runtime.Value
var once_applyDual sync.Once
func Get_applyDual() gopurs_runtime.Value {
	once_applyDual.Do(func() {
		cache_applyDual = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorDual()
}))
	})
	return cache_applyDual
}

var cache_bindDual gopurs_runtime.Value
var once_bindDual sync.Once
func Get_bindDual() gopurs_runtime.Value {
	once_bindDual.Do(func() {
		cache_bindDual = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDual()
}))
	})
	return cache_bindDual
}

var cache_applicativeDual gopurs_runtime.Value
var once_applicativeDual sync.Once
func Get_applicativeDual() gopurs_runtime.Value {
	once_applicativeDual.Do(func() {
		cache_applicativeDual = gopurs_runtime.RecordDict2("pure", "Apply0", Get_Dual(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDual()
}))
	})
	return cache_applicativeDual
}

var cache_monadDual gopurs_runtime.Value
var once_monadDual sync.Once
func Get_monadDual() gopurs_runtime.Value {
	once_monadDual.Do(func() {
		cache_monadDual = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindDual()
}))
	})
	return cache_monadDual
}





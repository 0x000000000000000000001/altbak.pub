package Data_Monoid_Dual

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Dual gopurs_runtime.Value
var once_Dual sync.Once
func Get_Dual() gopurs_runtime.Value {
	once_Dual.Do(func() {
		Dual = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Dual
}

var showDual gopurs_runtime.Value
var once_showDual sync.Once
func Get_showDual() gopurs_runtime.Value {
	once_showDual.Do(func() {
		showDual = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Dual " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal + ")")
}))
})
	})
	return showDual
}

var semigroupDual gopurs_runtime.Value
var once_semigroupDual sync.Once
func Get_semigroupDual() gopurs_runtime.Value {
	once_semigroupDual.Do(func() {
		semigroupDual = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), v1_2, v_1)
}))
})
	})
	return semigroupDual
}

var ordDual gopurs_runtime.Value
var once_ordDual sync.Once
func Get_ordDual() gopurs_runtime.Value {
	once_ordDual.Do(func() {
		ordDual = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordDual
}

var monoidDual gopurs_runtime.Value
var once_monoidDual sync.Once
func Get_monoidDual() gopurs_runtime.Value {
	once_monoidDual.Do(func() {
		monoidDual = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupDual1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), v1_3, v_2)
}))
_ = semigroupDual1_2_1
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDual1_2_1
}))
})
	})
	return monoidDual
}

var functorDual gopurs_runtime.Value
var once_functorDual sync.Once
func Get_functorDual() gopurs_runtime.Value {
	once_functorDual.Do(func() {
		functorDual = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
}))
	})
	return functorDual
}

var eqDual gopurs_runtime.Value
var once_eqDual sync.Once
func Get_eqDual() gopurs_runtime.Value {
	once_eqDual.Do(func() {
		eqDual = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqDual
}

var eq1Dual gopurs_runtime.Value
var once_eq1Dual sync.Once
func Get_eq1Dual() gopurs_runtime.Value {
	once_eq1Dual.Do(func() {
		eq1Dual = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return eq1Dual
}

var ord1Dual gopurs_runtime.Value
var once_ord1Dual sync.Once
func Get_ord1Dual() gopurs_runtime.Value {
	once_ord1Dual.Do(func() {
		ord1Dual = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Dual()
}))
	})
	return ord1Dual
}

var boundedDual gopurs_runtime.Value
var once_boundedDual sync.Once
func Get_boundedDual() gopurs_runtime.Value {
	once_boundedDual.Do(func() {
		boundedDual = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBounded_0
})
	})
	return boundedDual
}

var applyDual gopurs_runtime.Value
var once_applyDual sync.Once
func Get_applyDual() gopurs_runtime.Value {
	once_applyDual.Do(func() {
		applyDual = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorDual()
}))
	})
	return applyDual
}

var bindDual gopurs_runtime.Value
var once_bindDual sync.Once
func Get_bindDual() gopurs_runtime.Value {
	once_bindDual.Do(func() {
		bindDual = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDual()
}))
	})
	return bindDual
}

var applicativeDual gopurs_runtime.Value
var once_applicativeDual sync.Once
func Get_applicativeDual() gopurs_runtime.Value {
	once_applicativeDual.Do(func() {
		applicativeDual = gopurs_runtime.RecordDict2("pure", "Apply0", Get_Dual(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDual()
}))
	})
	return applicativeDual
}

var monadDual gopurs_runtime.Value
var once_monadDual sync.Once
func Get_monadDual() gopurs_runtime.Value {
	once_monadDual.Do(func() {
		monadDual = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindDual()
}))
	})
	return monadDual
}





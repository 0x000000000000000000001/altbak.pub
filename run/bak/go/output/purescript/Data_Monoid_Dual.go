package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Monoid_Dual_Dual gopurs_runtime.Value
var once_Data_Monoid_Dual_Dual sync.Once
func Get_Data_Monoid_Dual_Dual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_Dual.Do(func() {
		cache_Data_Monoid_Dual_Dual = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Dual_Dual(x_0_box)
})
	})
	return cache_Data_Monoid_Dual_Dual
}

var cache_Data_Monoid_Dual_showDual gopurs_runtime.Value
var once_Data_Monoid_Dual_showDual sync.Once
func Get_Data_Monoid_Dual_showDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_showDual.Do(func() {
		cache_Data_Monoid_Dual_showDual = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Dual_showDual(dictShow_0_box)
})
	})
	return cache_Data_Monoid_Dual_showDual
}

var cache_Data_Monoid_Dual_semigroupDual gopurs_runtime.Value
var once_Data_Monoid_Dual_semigroupDual sync.Once
func Get_Data_Monoid_Dual_semigroupDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_semigroupDual.Do(func() {
		cache_Data_Monoid_Dual_semigroupDual = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Dual_semigroupDual(dictSemigroup_0_box)
})
	})
	return cache_Data_Monoid_Dual_semigroupDual
}

var cache_Data_Monoid_Dual_ordDual gopurs_runtime.Value
var once_Data_Monoid_Dual_ordDual sync.Once
func Get_Data_Monoid_Dual_ordDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_ordDual.Do(func() {
		cache_Data_Monoid_Dual_ordDual = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Dual_ordDual(dictOrd_0_box)
})
	})
	return cache_Data_Monoid_Dual_ordDual
}

var cache_Data_Monoid_Dual_monoidDual gopurs_runtime.Value
var once_Data_Monoid_Dual_monoidDual sync.Once
func Get_Data_Monoid_Dual_monoidDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_monoidDual.Do(func() {
		cache_Data_Monoid_Dual_monoidDual = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Dual_monoidDual(dictMonoid_0_box)
})
	})
	return cache_Data_Monoid_Dual_monoidDual
}

var cache_Data_Monoid_Dual_functorDual gopurs_runtime.Value
var once_Data_Monoid_Dual_functorDual sync.Once
func Get_Data_Monoid_Dual_functorDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_functorDual.Do(func() {
		cache_Data_Monoid_Dual_functorDual = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Monoid_Dual_functorDual
}

var cache_Data_Monoid_Dual_eqDual gopurs_runtime.Value
var once_Data_Monoid_Dual_eqDual sync.Once
func Get_Data_Monoid_Dual_eqDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_eqDual.Do(func() {
		cache_Data_Monoid_Dual_eqDual = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Dual_eqDual(dictEq_0_box)
})
	})
	return cache_Data_Monoid_Dual_eqDual
}

var cache_Data_Monoid_Dual_eq1Dual gopurs_runtime.Value
var once_Data_Monoid_Dual_eq1Dual sync.Once
func Get_Data_Monoid_Dual_eq1Dual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_eq1Dual.Do(func() {
		cache_Data_Monoid_Dual_eq1Dual = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_Data_Monoid_Dual_eq1Dual
}

var cache_Data_Monoid_Dual_ord1Dual gopurs_runtime.Value
var once_Data_Monoid_Dual_ord1Dual sync.Once
func Get_Data_Monoid_Dual_ord1Dual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_ord1Dual.Do(func() {
		cache_Data_Monoid_Dual_ord1Dual = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_1, "eq")
}))
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}))
	})
	return cache_Data_Monoid_Dual_ord1Dual
}

var cache_Data_Monoid_Dual_boundedDual gopurs_runtime.Value
var once_Data_Monoid_Dual_boundedDual sync.Once
func Get_Data_Monoid_Dual_boundedDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_boundedDual.Do(func() {
		cache_Data_Monoid_Dual_boundedDual = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Dual_boundedDual(dictBounded_0_box)
})
	})
	return cache_Data_Monoid_Dual_boundedDual
}

var cache_Data_Monoid_Dual_applyDual gopurs_runtime.Value
var once_Data_Monoid_Dual_applyDual sync.Once
func Get_Data_Monoid_Dual_applyDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_applyDual.Do(func() {
		cache_Data_Monoid_Dual_applyDual = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, m_2)
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_Data_Monoid_Dual_applyDual
}

var cache_Data_Monoid_Dual_bindDual gopurs_runtime.Value
var once_Data_Monoid_Dual_bindDual sync.Once
func Get_Data_Monoid_Dual_bindDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_bindDual.Do(func() {
		cache_Data_Monoid_Dual_bindDual = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, m_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, v1_2)
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_Data_Monoid_Dual_bindDual
}

var cache_Data_Monoid_Dual_applicativeDual gopurs_runtime.Value
var once_Data_Monoid_Dual_applicativeDual sync.Once
func Get_Data_Monoid_Dual_applicativeDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_applicativeDual.Do(func() {
		cache_Data_Monoid_Dual_applicativeDual = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, m_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, v1_2)
})
}))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_Monoid_Dual_applicativeDual
}

var cache_Data_Monoid_Dual_monadDual gopurs_runtime.Value
var once_Data_Monoid_Dual_monadDual sync.Once
func Get_Data_Monoid_Dual_monadDual() gopurs_runtime.Value {
	once_Data_Monoid_Dual_monadDual.Do(func() {
		cache_Data_Monoid_Dual_monadDual = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, m_4)
})
}))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, v1_3)
})
}))
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, m_4)
})
}))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, v1_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, v_1)
})
}))
}))
	})
	return cache_Data_Monoid_Dual_monadDual
}

var cache_Data_Monoid_Dual_applicativeDual__4045440648 gopurs_runtime.Value
var once_Data_Monoid_Dual_applicativeDual__4045440648 sync.Once
func Get_Data_Monoid_Dual_applicativeDual__4045440648() gopurs_runtime.Value {
	once_Data_Monoid_Dual_applicativeDual__4045440648.Do(func() {
		cache_Data_Monoid_Dual_applicativeDual__4045440648 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, m_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, v1_2)
})
}))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_Monoid_Dual_applicativeDual__4045440648
}

var cache_Data_Monoid_Dual_applyDual__3199351098 gopurs_runtime.Value
var once_Data_Monoid_Dual_applyDual__3199351098 sync.Once
func Get_Data_Monoid_Dual_applyDual__3199351098() gopurs_runtime.Value {
	once_Data_Monoid_Dual_applyDual__3199351098.Do(func() {
		cache_Data_Monoid_Dual_applyDual__3199351098 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, m_2)
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_Data_Monoid_Dual_applyDual__3199351098
}

var cache_Data_Monoid_Dual_bindDual__329376103 gopurs_runtime.Value
var once_Data_Monoid_Dual_bindDual__329376103 sync.Once
func Get_Data_Monoid_Dual_bindDual__329376103() gopurs_runtime.Value {
	once_Data_Monoid_Dual_bindDual__329376103.Do(func() {
		cache_Data_Monoid_Dual_bindDual__329376103 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, m_3)
})
}))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, v1_2)
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_Data_Monoid_Dual_bindDual__329376103
}

var cache_Data_Monoid_Dual_eq1Dual__1905950174 gopurs_runtime.Value
var once_Data_Monoid_Dual_eq1Dual__1905950174 sync.Once
func Get_Data_Monoid_Dual_eq1Dual__1905950174() gopurs_runtime.Value {
	once_Data_Monoid_Dual_eq1Dual__1905950174.Do(func() {
		cache_Data_Monoid_Dual_eq1Dual__1905950174 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_Data_Monoid_Dual_eq1Dual__1905950174
}

var cache_Data_Monoid_Dual_functorDual__943655089 gopurs_runtime.Value
var once_Data_Monoid_Dual_functorDual__943655089 sync.Once
func Get_Data_Monoid_Dual_functorDual__943655089() gopurs_runtime.Value {
	once_Data_Monoid_Dual_functorDual__943655089.Do(func() {
		cache_Data_Monoid_Dual_functorDual__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Monoid_Dual_functorDual__943655089
}

func Call_Data_Monoid_Dual_Dual(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Monoid_Dual_showDual(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Dual ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}

func Call_Data_Monoid_Dual_semigroupDual(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), v1_2, v_1)
})
}))
}

func Call_Data_Monoid_Dual_ordDual(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_Data_Monoid_Dual_monoidDual(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupDual1_1_0 -> gopurs_runtime.Value
semigroupDual1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "append"), v1_3, v_2)
})
}))
_ = semigroupDual1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDual1_1_0
}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))
}

func Call_Data_Monoid_Dual_eqDual(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_Data_Monoid_Dual_boundedDual(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}



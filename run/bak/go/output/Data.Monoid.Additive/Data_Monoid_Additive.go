package Data_Monoid_Additive

import (
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Additive gopurs_runtime.Value
var once_Additive sync.Once
func Get_Additive() gopurs_runtime.Value {
	once_Additive.Do(func() {
		cache_Additive = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Additive(x_0_box)
})
	})
	return cache_Additive
}

var cache_showAdditive gopurs_runtime.Value
var once_showAdditive sync.Once
func Get_showAdditive() gopurs_runtime.Value {
	once_showAdditive.Do(func() {
		cache_showAdditive = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showAdditive(dictShow_0_box)
})
	})
	return cache_showAdditive
}

var cache_semigroupAdditive gopurs_runtime.Value
var once_semigroupAdditive sync.Once
func Get_semigroupAdditive() gopurs_runtime.Value {
	once_semigroupAdditive.Do(func() {
		cache_semigroupAdditive = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupAdditive(dictSemiring_0_box)
})
	})
	return cache_semigroupAdditive
}

var cache_ordAdditive gopurs_runtime.Value
var once_ordAdditive sync.Once
func Get_ordAdditive() gopurs_runtime.Value {
	once_ordAdditive.Do(func() {
		cache_ordAdditive = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordAdditive(dictOrd_0_box)
})
	})
	return cache_ordAdditive
}

var cache_monoidAdditive gopurs_runtime.Value
var once_monoidAdditive sync.Once
func Get_monoidAdditive() gopurs_runtime.Value {
	once_monoidAdditive.Do(func() {
		cache_monoidAdditive = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidAdditive(dictSemiring_0_box)
})
	})
	return cache_monoidAdditive
}

var cache_functorAdditive gopurs_runtime.Value
var once_functorAdditive sync.Once
func Get_functorAdditive() gopurs_runtime.Value {
	once_functorAdditive.Do(func() {
		cache_functorAdditive = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorAdditive
}

var cache_functorAdditive__gopurs_runtime_Value_943655089 gopurs_runtime.Value
var once_functorAdditive__gopurs_runtime_Value_943655089 sync.Once
func Get_functorAdditive__gopurs_runtime_Value_943655089() gopurs_runtime.Value {
	once_functorAdditive__gopurs_runtime_Value_943655089.Do(func() {
		cache_functorAdditive__gopurs_runtime_Value_943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorAdditive__gopurs_runtime_Value_943655089
}

var cache_eqAdditive gopurs_runtime.Value
var once_eqAdditive sync.Once
func Get_eqAdditive() gopurs_runtime.Value {
	once_eqAdditive.Do(func() {
		cache_eqAdditive = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqAdditive(dictEq_0_box)
})
	})
	return cache_eqAdditive
}

var cache_eq1Additive gopurs_runtime.Value
var once_eq1Additive sync.Once
func Get_eq1Additive() gopurs_runtime.Value {
	once_eq1Additive.Do(func() {
		cache_eq1Additive = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1Additive
}

var cache_eq1Additive__gopurs_runtime_Value_1905950174 gopurs_runtime.Value
var once_eq1Additive__gopurs_runtime_Value_1905950174 sync.Once
func Get_eq1Additive__gopurs_runtime_Value_1905950174() gopurs_runtime.Value {
	once_eq1Additive__gopurs_runtime_Value_1905950174.Do(func() {
		cache_eq1Additive__gopurs_runtime_Value_1905950174 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1Additive__gopurs_runtime_Value_1905950174
}

var cache_ord1Additive gopurs_runtime.Value
var once_ord1Additive sync.Once
func Get_ord1Additive() gopurs_runtime.Value {
	once_ord1Additive.Do(func() {
		cache_ord1Additive = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Additive()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}))
	})
	return cache_ord1Additive
}

var cache_boundedAdditive gopurs_runtime.Value
var once_boundedAdditive sync.Once
func Get_boundedAdditive() gopurs_runtime.Value {
	once_boundedAdditive.Do(func() {
		cache_boundedAdditive = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedAdditive(dictBounded_0_box)
})
	})
	return cache_boundedAdditive
}

var cache_applyAdditive gopurs_runtime.Value
var once_applyAdditive sync.Once
func Get_applyAdditive() gopurs_runtime.Value {
	once_applyAdditive.Do(func() {
		cache_applyAdditive = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorAdditive()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_applyAdditive
}

var cache_applyAdditive__gopurs_runtime_Value_3199351098 gopurs_runtime.Value
var once_applyAdditive__gopurs_runtime_Value_3199351098 sync.Once
func Get_applyAdditive__gopurs_runtime_Value_3199351098() gopurs_runtime.Value {
	once_applyAdditive__gopurs_runtime_Value_3199351098.Do(func() {
		cache_applyAdditive__gopurs_runtime_Value_3199351098 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorAdditive__gopurs_runtime_Value_943655089()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_applyAdditive__gopurs_runtime_Value_3199351098
}

var cache_bindAdditive gopurs_runtime.Value
var once_bindAdditive sync.Once
func Get_bindAdditive() gopurs_runtime.Value {
	once_bindAdditive.Do(func() {
		cache_bindAdditive = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAdditive()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_bindAdditive
}

var cache_bindAdditive__gopurs_runtime_Value_329376103 gopurs_runtime.Value
var once_bindAdditive__gopurs_runtime_Value_329376103 sync.Once
func Get_bindAdditive__gopurs_runtime_Value_329376103() gopurs_runtime.Value {
	once_bindAdditive__gopurs_runtime_Value_329376103.Do(func() {
		cache_bindAdditive__gopurs_runtime_Value_329376103 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAdditive__gopurs_runtime_Value_3199351098()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_bindAdditive__gopurs_runtime_Value_329376103
}

var cache_applicativeAdditive gopurs_runtime.Value
var once_applicativeAdditive sync.Once
func Get_applicativeAdditive() gopurs_runtime.Value {
	once_applicativeAdditive.Do(func() {
		cache_applicativeAdditive = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAdditive()
}), Get_Additive())
	})
	return cache_applicativeAdditive
}

var cache_applicativeAdditive__gopurs_runtime_Value_4045440648 gopurs_runtime.Value
var once_applicativeAdditive__gopurs_runtime_Value_4045440648 sync.Once
func Get_applicativeAdditive__gopurs_runtime_Value_4045440648() gopurs_runtime.Value {
	once_applicativeAdditive__gopurs_runtime_Value_4045440648.Do(func() {
		cache_applicativeAdditive__gopurs_runtime_Value_4045440648 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAdditive__gopurs_runtime_Value_3199351098()
}), Get_Additive())
	})
	return cache_applicativeAdditive__gopurs_runtime_Value_4045440648
}

var cache_monadAdditive gopurs_runtime.Value
var once_monadAdditive sync.Once
func Get_monadAdditive() gopurs_runtime.Value {
	once_monadAdditive.Do(func() {
		cache_monadAdditive = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeAdditive()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindAdditive()
}))
	})
	return cache_monadAdditive
}

func Call_Additive(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showAdditive(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Additive "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1), gopurs_runtime.Str(")"))).StrVal())
}))
}

func Call_semigroupAdditive(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), v_1, v1_2)
})
}))
}

func Call_ordAdditive(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_monoidAdditive(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
semigroupAdditive1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), v_1, v1_2)
})
}))
_ = semigroupAdditive1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_1_0
}), gopurs_runtime.RecordGet(dictSemiring_0, "zero"))
}

func Call_eqAdditive(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_boundedAdditive(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}



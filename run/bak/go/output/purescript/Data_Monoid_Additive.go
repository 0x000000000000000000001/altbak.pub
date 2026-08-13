package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Monoid_Additive_Additive gopurs_runtime.Value
var once_Data_Monoid_Additive_Additive sync.Once
func Get_Data_Monoid_Additive_Additive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_Additive.Do(func() {
		cache_Data_Monoid_Additive_Additive = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_Additive(x_0_box)
})
	})
	return cache_Data_Monoid_Additive_Additive
}

var cache_Data_Monoid_Additive_showAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_showAdditive sync.Once
func Get_Data_Monoid_Additive_showAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_showAdditive.Do(func() {
		cache_Data_Monoid_Additive_showAdditive = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_showAdditive(dictShow_0_box)
})
	})
	return cache_Data_Monoid_Additive_showAdditive
}

var cache_Data_Monoid_Additive_semigroupAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_semigroupAdditive sync.Once
func Get_Data_Monoid_Additive_semigroupAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_semigroupAdditive.Do(func() {
		cache_Data_Monoid_Additive_semigroupAdditive = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_semigroupAdditive(dictSemiring_0_box)
})
	})
	return cache_Data_Monoid_Additive_semigroupAdditive
}

var cache_Data_Monoid_Additive_ordAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_ordAdditive sync.Once
func Get_Data_Monoid_Additive_ordAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_ordAdditive.Do(func() {
		cache_Data_Monoid_Additive_ordAdditive = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_ordAdditive(dictOrd_0_box)
})
	})
	return cache_Data_Monoid_Additive_ordAdditive
}

var cache_Data_Monoid_Additive_monoidAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_monoidAdditive sync.Once
func Get_Data_Monoid_Additive_monoidAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_monoidAdditive.Do(func() {
		cache_Data_Monoid_Additive_monoidAdditive = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_monoidAdditive(dictSemiring_0_box)
})
	})
	return cache_Data_Monoid_Additive_monoidAdditive
}

var cache_Data_Monoid_Additive_functorAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_functorAdditive sync.Once
func Get_Data_Monoid_Additive_functorAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_functorAdditive.Do(func() {
		cache_Data_Monoid_Additive_functorAdditive = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Monoid_Additive_functorAdditive
}

var cache_Data_Monoid_Additive_eqAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_eqAdditive sync.Once
func Get_Data_Monoid_Additive_eqAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_eqAdditive.Do(func() {
		cache_Data_Monoid_Additive_eqAdditive = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_eqAdditive(dictEq_0_box)
})
	})
	return cache_Data_Monoid_Additive_eqAdditive
}

var cache_Data_Monoid_Additive_eq1Additive gopurs_runtime.Value
var once_Data_Monoid_Additive_eq1Additive sync.Once
func Get_Data_Monoid_Additive_eq1Additive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_eq1Additive.Do(func() {
		cache_Data_Monoid_Additive_eq1Additive = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_Data_Monoid_Additive_eq1Additive
}

var cache_Data_Monoid_Additive_ord1Additive gopurs_runtime.Value
var once_Data_Monoid_Additive_ord1Additive sync.Once
func Get_Data_Monoid_Additive_ord1Additive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_ord1Additive.Do(func() {
		cache_Data_Monoid_Additive_ord1Additive = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_1, "eq")
}))
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}))
	})
	return cache_Data_Monoid_Additive_ord1Additive
}

var cache_Data_Monoid_Additive_boundedAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_boundedAdditive sync.Once
func Get_Data_Monoid_Additive_boundedAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_boundedAdditive.Do(func() {
		cache_Data_Monoid_Additive_boundedAdditive = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Additive_boundedAdditive(dictBounded_0_box)
})
	})
	return cache_Data_Monoid_Additive_boundedAdditive
}

var cache_Data_Monoid_Additive_applyAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_applyAdditive sync.Once
func Get_Data_Monoid_Additive_applyAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_applyAdditive.Do(func() {
		cache_Data_Monoid_Additive_applyAdditive = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_Monoid_Additive_applyAdditive
}

var cache_Data_Monoid_Additive_bindAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_bindAdditive sync.Once
func Get_Data_Monoid_Additive_bindAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_bindAdditive.Do(func() {
		cache_Data_Monoid_Additive_bindAdditive = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_Monoid_Additive_bindAdditive
}

var cache_Data_Monoid_Additive_applicativeAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_applicativeAdditive sync.Once
func Get_Data_Monoid_Additive_applicativeAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_applicativeAdditive.Do(func() {
		cache_Data_Monoid_Additive_applicativeAdditive = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
}), Get_Data_Monoid_Additive_Additive())
	})
	return cache_Data_Monoid_Additive_applicativeAdditive
}

var cache_Data_Monoid_Additive_monadAdditive gopurs_runtime.Value
var once_Data_Monoid_Additive_monadAdditive sync.Once
func Get_Data_Monoid_Additive_monadAdditive() gopurs_runtime.Value {
	once_Data_Monoid_Additive_monadAdditive.Do(func() {
		cache_Data_Monoid_Additive_monadAdditive = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
}), Get_Data_Monoid_Additive_Additive())
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
	return cache_Data_Monoid_Additive_monadAdditive
}

var cache_Data_Monoid_Additive_applicativeAdditive__4045440648 gopurs_runtime.Value
var once_Data_Monoid_Additive_applicativeAdditive__4045440648 sync.Once
func Get_Data_Monoid_Additive_applicativeAdditive__4045440648() gopurs_runtime.Value {
	once_Data_Monoid_Additive_applicativeAdditive__4045440648.Do(func() {
		cache_Data_Monoid_Additive_applicativeAdditive__4045440648 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
}), Get_Data_Monoid_Additive_Additive())
	})
	return cache_Data_Monoid_Additive_applicativeAdditive__4045440648
}

var cache_Data_Monoid_Additive_applyAdditive__3199351098 gopurs_runtime.Value
var once_Data_Monoid_Additive_applyAdditive__3199351098 sync.Once
func Get_Data_Monoid_Additive_applyAdditive__3199351098() gopurs_runtime.Value {
	once_Data_Monoid_Additive_applyAdditive__3199351098.Do(func() {
		cache_Data_Monoid_Additive_applyAdditive__3199351098 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_Monoid_Additive_applyAdditive__3199351098
}

var cache_Data_Monoid_Additive_bindAdditive__329376103 gopurs_runtime.Value
var once_Data_Monoid_Additive_bindAdditive__329376103 sync.Once
func Get_Data_Monoid_Additive_bindAdditive__329376103() gopurs_runtime.Value {
	once_Data_Monoid_Additive_bindAdditive__329376103.Do(func() {
		cache_Data_Monoid_Additive_bindAdditive__329376103 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_Monoid_Additive_bindAdditive__329376103
}

var cache_Data_Monoid_Additive_eq1Additive__1905950174 gopurs_runtime.Value
var once_Data_Monoid_Additive_eq1Additive__1905950174 sync.Once
func Get_Data_Monoid_Additive_eq1Additive__1905950174() gopurs_runtime.Value {
	once_Data_Monoid_Additive_eq1Additive__1905950174.Do(func() {
		cache_Data_Monoid_Additive_eq1Additive__1905950174 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_Data_Monoid_Additive_eq1Additive__1905950174
}

var cache_Data_Monoid_Additive_functorAdditive__943655089 gopurs_runtime.Value
var once_Data_Monoid_Additive_functorAdditive__943655089 sync.Once
func Get_Data_Monoid_Additive_functorAdditive__943655089() gopurs_runtime.Value {
	once_Data_Monoid_Additive_functorAdditive__943655089.Do(func() {
		cache_Data_Monoid_Additive_functorAdditive__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Monoid_Additive_functorAdditive__943655089
}

func Call_Data_Monoid_Additive_Additive(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Monoid_Additive_showAdditive(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Additive ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}

func Call_Data_Monoid_Additive_semigroupAdditive(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), v_1, v1_2)
})
}))
}

func Call_Data_Monoid_Additive_ordAdditive(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_Data_Monoid_Additive_monoidAdditive(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
// TAST (Let): semigroupAdditive1_1_0 -> gopurs_runtime.Value
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

func Call_Data_Monoid_Additive_eqAdditive(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_Data_Monoid_Additive_boundedAdditive(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}



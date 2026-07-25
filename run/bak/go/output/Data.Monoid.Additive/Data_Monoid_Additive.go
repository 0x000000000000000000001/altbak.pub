package Data_Monoid_Additive

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
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
		cache_functorAdditive = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
}))
	})
	return cache_functorAdditive
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
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
}))
	})
	return cache_applyAdditive
}

var cache_bindAdditive gopurs_runtime.Value
var once_bindAdditive sync.Once
func Get_bindAdditive() gopurs_runtime.Value {
	once_bindAdditive.Do(func() {
		cache_bindAdditive = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAdditive()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
}))
	})
	return cache_bindAdditive
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Additive "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictShow_0.UnsafePtr)).V0, v_1), gopurs_runtime.Str(")")))
}))
}

func Call_semigroupAdditive(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData4)(dictSemiring_0.UnsafePtr)).V0, v_1, v1_2)
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
semigroupAdditive1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData4)(dictSemiring_0.UnsafePtr)).V0, v_1, v1_2)
}))
_ = semigroupAdditive1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_1_0
}), ((*gopurs_runtime.RecordData4)(dictSemiring_0.UnsafePtr)).V3)
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



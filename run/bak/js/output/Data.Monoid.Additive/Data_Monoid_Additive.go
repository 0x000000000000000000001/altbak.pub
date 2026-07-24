package Data_Monoid_Additive

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Additive gopurs_runtime.Value
var once_Additive sync.Once
func Get_Additive() gopurs_runtime.Value {
	once_Additive.Do(func() {
		Additive = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Additive
}

var showAdditive gopurs_runtime.Value
var once_showAdditive sync.Once
func Get_showAdditive() gopurs_runtime.Value {
	once_showAdditive.Do(func() {
		showAdditive = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Additive " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal + ")")
}))
})
	})
	return showAdditive
}

var semigroupAdditive gopurs_runtime.Value
var once_semigroupAdditive sync.Once
func Get_semigroupAdditive() gopurs_runtime.Value {
	once_semigroupAdditive.Do(func() {
		semigroupAdditive = gopurs_runtime.Func(func(dictSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), v_1, v1_2)
}))
})
	})
	return semigroupAdditive
}

var ordAdditive gopurs_runtime.Value
var once_ordAdditive sync.Once
func Get_ordAdditive() gopurs_runtime.Value {
	once_ordAdditive.Do(func() {
		ordAdditive = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordAdditive
}

var monoidAdditive gopurs_runtime.Value
var once_monoidAdditive sync.Once
func Get_monoidAdditive() gopurs_runtime.Value {
	once_monoidAdditive.Do(func() {
		monoidAdditive = gopurs_runtime.Func(func(dictSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupAdditive1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), v_1, v1_2)
}))
_ = semigroupAdditive1_1_0
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictSemiring_0, "zero"), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_1_0
}))
})
	})
	return monoidAdditive
}

var functorAdditive gopurs_runtime.Value
var once_functorAdditive sync.Once
func Get_functorAdditive() gopurs_runtime.Value {
	once_functorAdditive.Do(func() {
		functorAdditive = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
}))
	})
	return functorAdditive
}

var eqAdditive gopurs_runtime.Value
var once_eqAdditive sync.Once
func Get_eqAdditive() gopurs_runtime.Value {
	once_eqAdditive.Do(func() {
		eqAdditive = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqAdditive
}

var eq1Additive gopurs_runtime.Value
var once_eq1Additive sync.Once
func Get_eq1Additive() gopurs_runtime.Value {
	once_eq1Additive.Do(func() {
		eq1Additive = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return eq1Additive
}

var ord1Additive gopurs_runtime.Value
var once_ord1Additive sync.Once
func Get_ord1Additive() gopurs_runtime.Value {
	once_ord1Additive.Do(func() {
		ord1Additive = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Additive()
}))
	})
	return ord1Additive
}

var boundedAdditive gopurs_runtime.Value
var once_boundedAdditive sync.Once
func Get_boundedAdditive() gopurs_runtime.Value {
	once_boundedAdditive.Do(func() {
		boundedAdditive = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBounded_0
})
	})
	return boundedAdditive
}

var applyAdditive gopurs_runtime.Value
var once_applyAdditive sync.Once
func Get_applyAdditive() gopurs_runtime.Value {
	once_applyAdditive.Do(func() {
		applyAdditive = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorAdditive()
}))
	})
	return applyAdditive
}

var bindAdditive gopurs_runtime.Value
var once_bindAdditive sync.Once
func Get_bindAdditive() gopurs_runtime.Value {
	once_bindAdditive.Do(func() {
		bindAdditive = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAdditive()
}))
	})
	return bindAdditive
}

var applicativeAdditive gopurs_runtime.Value
var once_applicativeAdditive sync.Once
func Get_applicativeAdditive() gopurs_runtime.Value {
	once_applicativeAdditive.Do(func() {
		applicativeAdditive = gopurs_runtime.RecordDict2("pure", "Apply0", Get_Additive(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyAdditive()
}))
	})
	return applicativeAdditive
}

var monadAdditive gopurs_runtime.Value
var once_monadAdditive sync.Once
func Get_monadAdditive() gopurs_runtime.Value {
	once_monadAdditive.Do(func() {
		monadAdditive = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeAdditive()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindAdditive()
}))
	})
	return monadAdditive
}





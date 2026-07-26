package Data_Monoid_Multiplicative

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var cache_Multiplicative gopurs_runtime.Value
var once_Multiplicative sync.Once
func Get_Multiplicative() gopurs_runtime.Value {
	once_Multiplicative.Do(func() {
		cache_Multiplicative = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Multiplicative(x_0_box)
})
	})
	return cache_Multiplicative
}

var cache_showMultiplicative gopurs_runtime.Value
var once_showMultiplicative sync.Once
func Get_showMultiplicative() gopurs_runtime.Value {
	once_showMultiplicative.Do(func() {
		cache_showMultiplicative = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showMultiplicative(dictShow_0_box)
})
	})
	return cache_showMultiplicative
}

var cache_semigroupMultiplicative gopurs_runtime.Value
var once_semigroupMultiplicative sync.Once
func Get_semigroupMultiplicative() gopurs_runtime.Value {
	once_semigroupMultiplicative.Do(func() {
		cache_semigroupMultiplicative = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupMultiplicative(dictSemiring_0_box)
})
	})
	return cache_semigroupMultiplicative
}

var cache_ordMultiplicative gopurs_runtime.Value
var once_ordMultiplicative sync.Once
func Get_ordMultiplicative() gopurs_runtime.Value {
	once_ordMultiplicative.Do(func() {
		cache_ordMultiplicative = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordMultiplicative(dictOrd_0_box)
})
	})
	return cache_ordMultiplicative
}

var cache_monoidMultiplicative gopurs_runtime.Value
var once_monoidMultiplicative sync.Once
func Get_monoidMultiplicative() gopurs_runtime.Value {
	once_monoidMultiplicative.Do(func() {
		cache_monoidMultiplicative = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidMultiplicative(dictSemiring_0_box)
})
	})
	return cache_monoidMultiplicative
}

var cache_functorMultiplicative gopurs_runtime.Value
var once_functorMultiplicative sync.Once
func Get_functorMultiplicative() gopurs_runtime.Value {
	once_functorMultiplicative.Do(func() {
		cache_functorMultiplicative = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
}))
	})
	return cache_functorMultiplicative
}

var cache_eqMultiplicative gopurs_runtime.Value
var once_eqMultiplicative sync.Once
func Get_eqMultiplicative() gopurs_runtime.Value {
	once_eqMultiplicative.Do(func() {
		cache_eqMultiplicative = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqMultiplicative(dictEq_0_box)
})
	})
	return cache_eqMultiplicative
}

var cache_eq1Multiplicative gopurs_runtime.Value
var once_eq1Multiplicative sync.Once
func Get_eq1Multiplicative() gopurs_runtime.Value {
	once_eq1Multiplicative.Do(func() {
		cache_eq1Multiplicative = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1Multiplicative
}

var cache_ord1Multiplicative gopurs_runtime.Value
var once_ord1Multiplicative sync.Once
func Get_ord1Multiplicative() gopurs_runtime.Value {
	once_ord1Multiplicative.Do(func() {
		cache_ord1Multiplicative = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Multiplicative()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}))
	})
	return cache_ord1Multiplicative
}

var cache_boundedMultiplicative gopurs_runtime.Value
var once_boundedMultiplicative sync.Once
func Get_boundedMultiplicative() gopurs_runtime.Value {
	once_boundedMultiplicative.Do(func() {
		cache_boundedMultiplicative = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedMultiplicative(dictBounded_0_box)
})
	})
	return cache_boundedMultiplicative
}

var cache_applyMultiplicative gopurs_runtime.Value
var once_applyMultiplicative sync.Once
func Get_applyMultiplicative() gopurs_runtime.Value {
	once_applyMultiplicative.Do(func() {
		cache_applyMultiplicative = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMultiplicative()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
}))
	})
	return cache_applyMultiplicative
}

var cache_bindMultiplicative gopurs_runtime.Value
var once_bindMultiplicative sync.Once
func Get_bindMultiplicative() gopurs_runtime.Value {
	once_bindMultiplicative.Do(func() {
		cache_bindMultiplicative = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMultiplicative()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
}))
	})
	return cache_bindMultiplicative
}

var cache_applicativeMultiplicative gopurs_runtime.Value
var once_applicativeMultiplicative sync.Once
func Get_applicativeMultiplicative() gopurs_runtime.Value {
	once_applicativeMultiplicative.Do(func() {
		cache_applicativeMultiplicative = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMultiplicative()
}), Get_Multiplicative())
	})
	return cache_applicativeMultiplicative
}

var cache_monadMultiplicative gopurs_runtime.Value
var once_monadMultiplicative sync.Once
func Get_monadMultiplicative() gopurs_runtime.Value {
	once_monadMultiplicative.Do(func() {
		cache_monadMultiplicative = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeMultiplicative()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindMultiplicative()
}))
	})
	return cache_monadMultiplicative
}

func Call_Multiplicative(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showMultiplicative(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Multiplicative "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1), gopurs_runtime.Str(")")))
}))
}

func Call_semigroupMultiplicative(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), v_1, v1_2)
}))
}

func Call_ordMultiplicative(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_monoidMultiplicative(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
semigroupMultiplicative1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), v_1, v1_2)
}))
_ = semigroupMultiplicative1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMultiplicative1_1_0
}), gopurs_runtime.RecordGet(dictSemiring_0, "one"))
}

func Call_eqMultiplicative(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_boundedMultiplicative(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}



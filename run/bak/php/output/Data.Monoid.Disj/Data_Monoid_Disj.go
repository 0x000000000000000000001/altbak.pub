package Data_Monoid_Disj

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var cache_Disj gopurs_runtime.Value
var once_Disj sync.Once
func Get_Disj() gopurs_runtime.Value {
	once_Disj.Do(func() {
		cache_Disj = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Disj(x_0_box)
})
	})
	return cache_Disj
}

var cache_showDisj gopurs_runtime.Value
var once_showDisj sync.Once
func Get_showDisj() gopurs_runtime.Value {
	once_showDisj.Do(func() {
		cache_showDisj = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showDisj(dictShow_0_box)
})
	})
	return cache_showDisj
}

var cache_semiringDisj gopurs_runtime.Value
var once_semiringDisj sync.Once
func Get_semiringDisj() gopurs_runtime.Value {
	once_semiringDisj.Do(func() {
		cache_semiringDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semiringDisj(dictHeytingAlgebra_0_box)
})
	})
	return cache_semiringDisj
}

var cache_semigroupDisj gopurs_runtime.Value
var once_semigroupDisj sync.Once
func Get_semigroupDisj() gopurs_runtime.Value {
	once_semigroupDisj.Do(func() {
		cache_semigroupDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupDisj(dictHeytingAlgebra_0_box)
})
	})
	return cache_semigroupDisj
}

var cache_ordDisj gopurs_runtime.Value
var once_ordDisj sync.Once
func Get_ordDisj() gopurs_runtime.Value {
	once_ordDisj.Do(func() {
		cache_ordDisj = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordDisj(dictOrd_0_box)
})
	})
	return cache_ordDisj
}

var cache_monoidDisj gopurs_runtime.Value
var once_monoidDisj sync.Once
func Get_monoidDisj() gopurs_runtime.Value {
	once_monoidDisj.Do(func() {
		cache_monoidDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidDisj(dictHeytingAlgebra_0_box)
})
	})
	return cache_monoidDisj
}

var cache_functorDisj gopurs_runtime.Value
var once_functorDisj sync.Once
func Get_functorDisj() gopurs_runtime.Value {
	once_functorDisj.Do(func() {
		cache_functorDisj = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
}))
	})
	return cache_functorDisj
}

var cache_eqDisj gopurs_runtime.Value
var once_eqDisj sync.Once
func Get_eqDisj() gopurs_runtime.Value {
	once_eqDisj.Do(func() {
		cache_eqDisj = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqDisj(dictEq_0_box)
})
	})
	return cache_eqDisj
}

var cache_eq1Disj gopurs_runtime.Value
var once_eq1Disj sync.Once
func Get_eq1Disj() gopurs_runtime.Value {
	once_eq1Disj.Do(func() {
		cache_eq1Disj = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1Disj
}

var cache_ord1Disj gopurs_runtime.Value
var once_ord1Disj sync.Once
func Get_ord1Disj() gopurs_runtime.Value {
	once_ord1Disj.Do(func() {
		cache_ord1Disj = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Disj()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}))
	})
	return cache_ord1Disj
}

var cache_boundedDisj gopurs_runtime.Value
var once_boundedDisj sync.Once
func Get_boundedDisj() gopurs_runtime.Value {
	once_boundedDisj.Do(func() {
		cache_boundedDisj = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedDisj(dictBounded_0_box)
})
	})
	return cache_boundedDisj
}

var cache_applyDisj gopurs_runtime.Value
var once_applyDisj sync.Once
func Get_applyDisj() gopurs_runtime.Value {
	once_applyDisj.Do(func() {
		cache_applyDisj = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorDisj()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
}))
	})
	return cache_applyDisj
}

var cache_bindDisj gopurs_runtime.Value
var once_bindDisj sync.Once
func Get_bindDisj() gopurs_runtime.Value {
	once_bindDisj.Do(func() {
		cache_bindDisj = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDisj()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
}))
	})
	return cache_bindDisj
}

var cache_applicativeDisj gopurs_runtime.Value
var once_applicativeDisj sync.Once
func Get_applicativeDisj() gopurs_runtime.Value {
	once_applicativeDisj.Do(func() {
		cache_applicativeDisj = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDisj()
}), Get_Disj())
	})
	return cache_applicativeDisj
}

var cache_monadDisj gopurs_runtime.Value
var once_monadDisj sync.Once
func Get_monadDisj() gopurs_runtime.Value {
	once_monadDisj.Do(func() {
		cache_monadDisj = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeDisj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindDisj()
}))
	})
	return cache_monadDisj
}

func Call_Disj(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showDisj(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Disj "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictShow_0.UnsafePtr)).V0, v_1), gopurs_runtime.Str(")")))
}))
}

func Call_semiringDisj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData)(dictHeytingAlgebra_0.UnsafePtr)).Vals[1], v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData)(dictHeytingAlgebra_0.UnsafePtr)).Vals[0], v_1, v1_2)
}), ((*gopurs_runtime.RecordData)(dictHeytingAlgebra_0.UnsafePtr)).Vals[5], ((*gopurs_runtime.RecordData)(dictHeytingAlgebra_0.UnsafePtr)).Vals[2])
}

func Call_semigroupDisj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData)(dictHeytingAlgebra_0.UnsafePtr)).Vals[1], v_1, v1_2)
}))
}

func Call_ordDisj(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_monoidDisj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
semigroupDisj1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData)(dictHeytingAlgebra_0.UnsafePtr)).Vals[1], v_1, v1_2)
}))
_ = semigroupDisj1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_1_0
}), ((*gopurs_runtime.RecordData)(dictHeytingAlgebra_0.UnsafePtr)).Vals[2])
}

func Call_eqDisj(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_boundedDisj(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}



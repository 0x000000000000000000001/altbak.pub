package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Monoid_Disj_Disj gopurs_runtime.Value
var once_Data_Monoid_Disj_Disj sync.Once
func Get_Data_Monoid_Disj_Disj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_Disj.Do(func() {
		cache_Data_Monoid_Disj_Disj = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Disj_Disj(x_0_box)
})
	})
	return cache_Data_Monoid_Disj_Disj
}

var cache_Data_Monoid_Disj_showDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_showDisj sync.Once
func Get_Data_Monoid_Disj_showDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_showDisj.Do(func() {
		cache_Data_Monoid_Disj_showDisj = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Disj_showDisj(dictShow_0_box)
})
	})
	return cache_Data_Monoid_Disj_showDisj
}

var cache_Data_Monoid_Disj_semiringDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_semiringDisj sync.Once
func Get_Data_Monoid_Disj_semiringDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_semiringDisj.Do(func() {
		cache_Data_Monoid_Disj_semiringDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Disj_semiringDisj(dictHeytingAlgebra_0_box)
})
	})
	return cache_Data_Monoid_Disj_semiringDisj
}

var cache_Data_Monoid_Disj_semigroupDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_semigroupDisj sync.Once
func Get_Data_Monoid_Disj_semigroupDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_semigroupDisj.Do(func() {
		cache_Data_Monoid_Disj_semigroupDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Disj_semigroupDisj(dictHeytingAlgebra_0_box)
})
	})
	return cache_Data_Monoid_Disj_semigroupDisj
}

var cache_Data_Monoid_Disj_ordDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_ordDisj sync.Once
func Get_Data_Monoid_Disj_ordDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_ordDisj.Do(func() {
		cache_Data_Monoid_Disj_ordDisj = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Disj_ordDisj(dictOrd_0_box)
})
	})
	return cache_Data_Monoid_Disj_ordDisj
}

var cache_Data_Monoid_Disj_monoidDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_monoidDisj sync.Once
func Get_Data_Monoid_Disj_monoidDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_monoidDisj.Do(func() {
		cache_Data_Monoid_Disj_monoidDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Disj_monoidDisj(dictHeytingAlgebra_0_box)
})
	})
	return cache_Data_Monoid_Disj_monoidDisj
}

var cache_Data_Monoid_Disj_functorDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_functorDisj sync.Once
func Get_Data_Monoid_Disj_functorDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_functorDisj.Do(func() {
		cache_Data_Monoid_Disj_functorDisj = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Monoid_Disj_functorDisj
}

var cache_Data_Monoid_Disj_eqDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_eqDisj sync.Once
func Get_Data_Monoid_Disj_eqDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_eqDisj.Do(func() {
		cache_Data_Monoid_Disj_eqDisj = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Disj_eqDisj(dictEq_0_box)
})
	})
	return cache_Data_Monoid_Disj_eqDisj
}

var cache_Data_Monoid_Disj_eq1Disj gopurs_runtime.Value
var once_Data_Monoid_Disj_eq1Disj sync.Once
func Get_Data_Monoid_Disj_eq1Disj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_eq1Disj.Do(func() {
		cache_Data_Monoid_Disj_eq1Disj = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_Data_Monoid_Disj_eq1Disj
}

var cache_Data_Monoid_Disj_ord1Disj gopurs_runtime.Value
var once_Data_Monoid_Disj_ord1Disj sync.Once
func Get_Data_Monoid_Disj_ord1Disj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_ord1Disj.Do(func() {
		cache_Data_Monoid_Disj_ord1Disj = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Disj_eq1Disj()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}))
	})
	return cache_Data_Monoid_Disj_ord1Disj
}

var cache_Data_Monoid_Disj_boundedDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_boundedDisj sync.Once
func Get_Data_Monoid_Disj_boundedDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_boundedDisj.Do(func() {
		cache_Data_Monoid_Disj_boundedDisj = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Disj_boundedDisj(dictBounded_0_box)
})
	})
	return cache_Data_Monoid_Disj_boundedDisj
}

var cache_Data_Monoid_Disj_applyDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_applyDisj sync.Once
func Get_Data_Monoid_Disj_applyDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_applyDisj.Do(func() {
		cache_Data_Monoid_Disj_applyDisj = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Disj_functorDisj()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_Data_Monoid_Disj_applyDisj
}

var cache_Data_Monoid_Disj_bindDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_bindDisj sync.Once
func Get_Data_Monoid_Disj_bindDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_bindDisj.Do(func() {
		cache_Data_Monoid_Disj_bindDisj = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Disj_applyDisj()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_Data_Monoid_Disj_bindDisj
}

var cache_Data_Monoid_Disj_applicativeDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_applicativeDisj sync.Once
func Get_Data_Monoid_Disj_applicativeDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_applicativeDisj.Do(func() {
		cache_Data_Monoid_Disj_applicativeDisj = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Disj_applyDisj()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_Monoid_Disj_applicativeDisj
}

var cache_Data_Monoid_Disj_monadDisj gopurs_runtime.Value
var once_Data_Monoid_Disj_monadDisj sync.Once
func Get_Data_Monoid_Disj_monadDisj() gopurs_runtime.Value {
	once_Data_Monoid_Disj_monadDisj.Do(func() {
		cache_Data_Monoid_Disj_monadDisj = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Disj_applicativeDisj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Disj_bindDisj()
}))
	})
	return cache_Data_Monoid_Disj_monadDisj
}

var cache_Data_Monoid_Disj_applicativeDisj__4045440648 gopurs_runtime.Value
var once_Data_Monoid_Disj_applicativeDisj__4045440648 sync.Once
func Get_Data_Monoid_Disj_applicativeDisj__4045440648() gopurs_runtime.Value {
	once_Data_Monoid_Disj_applicativeDisj__4045440648.Do(func() {
		cache_Data_Monoid_Disj_applicativeDisj__4045440648 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Disj_applyDisj()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_Monoid_Disj_applicativeDisj__4045440648
}

var cache_Data_Monoid_Disj_applyDisj__3199351098 gopurs_runtime.Value
var once_Data_Monoid_Disj_applyDisj__3199351098 sync.Once
func Get_Data_Monoid_Disj_applyDisj__3199351098() gopurs_runtime.Value {
	once_Data_Monoid_Disj_applyDisj__3199351098.Do(func() {
		cache_Data_Monoid_Disj_applyDisj__3199351098 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Disj_functorDisj()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
})
}))
	})
	return cache_Data_Monoid_Disj_applyDisj__3199351098
}

var cache_Data_Monoid_Disj_bindDisj__329376103 gopurs_runtime.Value
var once_Data_Monoid_Disj_bindDisj__329376103 sync.Once
func Get_Data_Monoid_Disj_bindDisj__329376103() gopurs_runtime.Value {
	once_Data_Monoid_Disj_bindDisj__329376103.Do(func() {
		cache_Data_Monoid_Disj_bindDisj__329376103 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Disj_applyDisj()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
})
}))
	})
	return cache_Data_Monoid_Disj_bindDisj__329376103
}

var cache_Data_Monoid_Disj_eq1Disj__1905950174 gopurs_runtime.Value
var once_Data_Monoid_Disj_eq1Disj__1905950174 sync.Once
func Get_Data_Monoid_Disj_eq1Disj__1905950174() gopurs_runtime.Value {
	once_Data_Monoid_Disj_eq1Disj__1905950174.Do(func() {
		cache_Data_Monoid_Disj_eq1Disj__1905950174 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_Data_Monoid_Disj_eq1Disj__1905950174
}

var cache_Data_Monoid_Disj_functorDisj__943655089 gopurs_runtime.Value
var once_Data_Monoid_Disj_functorDisj__943655089 sync.Once
func Get_Data_Monoid_Disj_functorDisj__943655089() gopurs_runtime.Value {
	once_Data_Monoid_Disj_functorDisj__943655089.Do(func() {
		cache_Data_Monoid_Disj_functorDisj__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_Data_Monoid_Disj_functorDisj__943655089
}

func Call_Data_Monoid_Disj_Disj(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Monoid_Disj_showDisj(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Disj ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}

func Call_Data_Monoid_Disj_semiringDisj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), v_1, v1_2)
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
})
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt"), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff"))
}

func Call_Data_Monoid_Disj_semigroupDisj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), v_1, v1_2)
})
}))
}

func Call_Data_Monoid_Disj_ordDisj(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_Data_Monoid_Disj_monoidDisj(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
// TAST (Let): semigroupDisj1_1_0 -> gopurs_runtime.Value
semigroupDisj1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), v_1, v1_2)
})
}))
_ = semigroupDisj1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_1_0
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff"))
}

func Call_Data_Monoid_Disj_eqDisj(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_Data_Monoid_Disj_boundedDisj(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}



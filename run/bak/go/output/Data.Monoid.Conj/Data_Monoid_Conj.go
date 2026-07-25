package Data_Monoid_Conj

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Conj gopurs_runtime.Value
var once_Conj sync.Once
func Get_Conj() gopurs_runtime.Value {
	once_Conj.Do(func() {
		cache_Conj = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Conj
}

var cache_showConj gopurs_runtime.Value
var once_showConj sync.Once
func Get_showConj() gopurs_runtime.Value {
	once_showConj.Do(func() {
		cache_showConj = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Conj ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}()
})
	})
	return cache_showConj
}

var cache_semiringConj gopurs_runtime.Value
var once_semiringConj sync.Once
func Get_semiringConj() gopurs_runtime.Value {
	once_semiringConj.Do(func() {
		cache_semiringConj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.RecordDict4("zero", "one", "add", "mul", gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt"), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff"), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), v_1, v1_2)
}))
}()
})
	})
	return cache_semiringConj
}

var cache_semigroupConj gopurs_runtime.Value
var once_semigroupConj sync.Once
func Get_semigroupConj() gopurs_runtime.Value {
	once_semigroupConj.Do(func() {
		cache_semigroupConj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
}))
}()
})
	})
	return cache_semigroupConj
}

var cache_ordConj gopurs_runtime.Value
var once_ordConj sync.Once
func Get_ordConj() gopurs_runtime.Value {
	once_ordConj.Do(func() {
		cache_ordConj = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}()
})
	})
	return cache_ordConj
}

var cache_monoidConj gopurs_runtime.Value
var once_monoidConj sync.Once
func Get_monoidConj() gopurs_runtime.Value {
	once_monoidConj.Do(func() {
		cache_monoidConj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
semigroupConj1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
}))
_ = semigroupConj1_1_0
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt"), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_1_0
}))
}()
})
	})
	return cache_monoidConj
}

var cache_functorConj gopurs_runtime.Value
var once_functorConj sync.Once
func Get_functorConj() gopurs_runtime.Value {
	once_functorConj.Do(func() {
		cache_functorConj = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
}))
	})
	return cache_functorConj
}

var cache_eqConj gopurs_runtime.Value
var once_eqConj sync.Once
func Get_eqConj() gopurs_runtime.Value {
	once_eqConj.Do(func() {
		cache_eqConj = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}()
})
	})
	return cache_eqConj
}

var cache_eq1Conj gopurs_runtime.Value
var once_eq1Conj sync.Once
func Get_eq1Conj() gopurs_runtime.Value {
	once_eq1Conj.Do(func() {
		cache_eq1Conj = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return cache_eq1Conj
}

var cache_ord1Conj gopurs_runtime.Value
var once_ord1Conj sync.Once
func Get_ord1Conj() gopurs_runtime.Value {
	once_ord1Conj.Do(func() {
		cache_ord1Conj = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Conj()
}))
	})
	return cache_ord1Conj
}

var cache_boundedConj gopurs_runtime.Value
var once_boundedConj sync.Once
func Get_boundedConj() gopurs_runtime.Value {
	once_boundedConj.Do(func() {
		cache_boundedConj = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}()
})
	})
	return cache_boundedConj
}

var cache_applyConj gopurs_runtime.Value
var once_applyConj sync.Once
func Get_applyConj() gopurs_runtime.Value {
	once_applyConj.Do(func() {
		cache_applyConj = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorConj()
}))
	})
	return cache_applyConj
}

var cache_bindConj gopurs_runtime.Value
var once_bindConj sync.Once
func Get_bindConj() gopurs_runtime.Value {
	once_bindConj.Do(func() {
		cache_bindConj = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyConj()
}))
	})
	return cache_bindConj
}

var cache_applicativeConj gopurs_runtime.Value
var once_applicativeConj sync.Once
func Get_applicativeConj() gopurs_runtime.Value {
	once_applicativeConj.Do(func() {
		cache_applicativeConj = gopurs_runtime.RecordDict2("pure", "Apply0", Get_Conj(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyConj()
}))
	})
	return cache_applicativeConj
}

var cache_monadConj gopurs_runtime.Value
var once_monadConj sync.Once
func Get_monadConj() gopurs_runtime.Value {
	once_monadConj.Do(func() {
		cache_monadConj = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeConj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindConj()
}))
	})
	return cache_monadConj
}





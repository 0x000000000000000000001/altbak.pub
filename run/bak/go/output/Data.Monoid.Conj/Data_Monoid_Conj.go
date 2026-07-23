package Data_Monoid_Conj

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Conj gopurs_runtime.Value
var once_Conj sync.Once
func Get_Conj() gopurs_runtime.Value {
	once_Conj.Do(func() {
		Conj = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Conj
}

var showConj gopurs_runtime.Value
var once_showConj sync.Once
func Get_showConj() gopurs_runtime.Value {
	once_showConj.Do(func() {
		showConj = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Conj ").StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
})
	})
	return showConj
}

var semiringConj gopurs_runtime.Value
var once_semiringConj sync.Once
func Get_semiringConj() gopurs_runtime.Value {
	once_semiringConj.Do(func() {
		semiringConj = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("zero", "one", "add", "mul", gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt"), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff"), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), v_1, v1_2)
}))
})
	})
	return semiringConj
}

var semigroupConj gopurs_runtime.Value
var once_semigroupConj sync.Once
func Get_semigroupConj() gopurs_runtime.Value {
	once_semigroupConj.Do(func() {
		semigroupConj = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
}))
})
	})
	return semigroupConj
}

var ordConj gopurs_runtime.Value
var once_ordConj sync.Once
func Get_ordConj() gopurs_runtime.Value {
	once_ordConj.Do(func() {
		ordConj = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictOrd_0
})
	})
	return ordConj
}

var monoidConj gopurs_runtime.Value
var once_monoidConj sync.Once
func Get_monoidConj() gopurs_runtime.Value {
	once_monoidConj.Do(func() {
		monoidConj = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupConj1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
}))
_ = semigroupConj1_1_0
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt"), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_1_0
}))
})
	})
	return monoidConj
}

var functorConj gopurs_runtime.Value
var once_functorConj sync.Once
func Get_functorConj() gopurs_runtime.Value {
	once_functorConj.Do(func() {
		functorConj = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
}))
	})
	return functorConj
}

var eqConj gopurs_runtime.Value
var once_eqConj sync.Once
func Get_eqConj() gopurs_runtime.Value {
	once_eqConj.Do(func() {
		eqConj = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEq_0
})
	})
	return eqConj
}

var eq1Conj gopurs_runtime.Value
var once_eq1Conj sync.Once
func Get_eq1Conj() gopurs_runtime.Value {
	once_eq1Conj.Do(func() {
		eq1Conj = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return eq1Conj
}

var ord1Conj gopurs_runtime.Value
var once_ord1Conj sync.Once
func Get_ord1Conj() gopurs_runtime.Value {
	once_ord1Conj.Do(func() {
		ord1Conj = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Conj()
}))
	})
	return ord1Conj
}

var boundedConj gopurs_runtime.Value
var once_boundedConj sync.Once
func Get_boundedConj() gopurs_runtime.Value {
	once_boundedConj.Do(func() {
		boundedConj = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBounded_0
})
	})
	return boundedConj
}

var applyConj gopurs_runtime.Value
var once_applyConj sync.Once
func Get_applyConj() gopurs_runtime.Value {
	once_applyConj.Do(func() {
		applyConj = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorConj()
}))
	})
	return applyConj
}

var bindConj gopurs_runtime.Value
var once_bindConj sync.Once
func Get_bindConj() gopurs_runtime.Value {
	once_bindConj.Do(func() {
		bindConj = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyConj()
}))
	})
	return bindConj
}

var applicativeConj gopurs_runtime.Value
var once_applicativeConj sync.Once
func Get_applicativeConj() gopurs_runtime.Value {
	once_applicativeConj.Do(func() {
		applicativeConj = gopurs_runtime.RecordDict2("pure", "Apply0", Get_Conj(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyConj()
}))
	})
	return applicativeConj
}

var monadConj gopurs_runtime.Value
var once_monadConj sync.Once
func Get_monadConj() gopurs_runtime.Value {
	once_monadConj.Do(func() {
		monadConj = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeConj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindConj()
}))
	})
	return monadConj
}



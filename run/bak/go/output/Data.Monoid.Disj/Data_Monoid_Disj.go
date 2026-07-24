package Data_Monoid_Disj

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Disj gopurs_runtime.Value
var once_Disj sync.Once
func Get_Disj() gopurs_runtime.Value {
	once_Disj.Do(func() {
		Disj = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return Disj
}

var showDisj gopurs_runtime.Value
var once_showDisj sync.Once
func Get_showDisj() gopurs_runtime.Value {
	once_showDisj.Do(func() {
		showDisj = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Disj " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), v_1).StrVal + ")")
}))
}()
})
	})
	return showDisj
}

var semiringDisj gopurs_runtime.Value
var once_semiringDisj sync.Once
func Get_semiringDisj() gopurs_runtime.Value {
	once_semiringDisj.Do(func() {
		semiringDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.RecordDict4("zero", "one", "add", "mul", gopurs_runtime.RecordGet(dictHeytingAlgebra_0_loop, "ff"), gopurs_runtime.RecordGet(dictHeytingAlgebra_0_loop, "tt"), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0_loop, "disj"), v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0_loop, "conj"), v_1, v1_2)
}))
}()
})
	})
	return semiringDisj
}

var semigroupDisj gopurs_runtime.Value
var once_semigroupDisj sync.Once
func Get_semigroupDisj() gopurs_runtime.Value {
	once_semigroupDisj.Do(func() {
		semigroupDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0_loop, "disj"), v_1, v1_2)
}))
}()
})
	})
	return semigroupDisj
}

var ordDisj gopurs_runtime.Value
var once_ordDisj sync.Once
func Get_ordDisj() gopurs_runtime.Value {
	once_ordDisj.Do(func() {
		ordDisj = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0_loop
}()
})
	})
	return ordDisj
}

var monoidDisj gopurs_runtime.Value
var once_monoidDisj sync.Once
func Get_monoidDisj() gopurs_runtime.Value {
	once_monoidDisj.Do(func() {
		monoidDisj = gopurs_runtime.Func(func(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
semigroupDisj1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0_loop, "disj"), v_1, v1_2)
}))
_ = semigroupDisj1_1_0
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictHeytingAlgebra_0_loop, "ff"), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_1_0
}))
}()
})
	})
	return monoidDisj
}

var functorDisj gopurs_runtime.Value
var once_functorDisj sync.Once
func Get_functorDisj() gopurs_runtime.Value {
	once_functorDisj.Do(func() {
		functorDisj = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
}))
	})
	return functorDisj
}

var eqDisj gopurs_runtime.Value
var once_eqDisj sync.Once
func Get_eqDisj() gopurs_runtime.Value {
	once_eqDisj.Do(func() {
		eqDisj = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0_loop
}()
})
	})
	return eqDisj
}

var eq1Disj gopurs_runtime.Value
var once_eq1Disj sync.Once
func Get_eq1Disj() gopurs_runtime.Value {
	once_eq1Disj.Do(func() {
		eq1Disj = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}))
	})
	return eq1Disj
}

var ord1Disj gopurs_runtime.Value
var once_ord1Disj sync.Once
func Get_ord1Disj() gopurs_runtime.Value {
	once_ord1Disj.Do(func() {
		ord1Disj = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictOrd_0, "compare")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Disj()
}))
	})
	return ord1Disj
}

var boundedDisj gopurs_runtime.Value
var once_boundedDisj sync.Once
func Get_boundedDisj() gopurs_runtime.Value {
	once_boundedDisj.Do(func() {
		boundedDisj = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0_loop
}()
})
	})
	return boundedDisj
}

var applyDisj gopurs_runtime.Value
var once_applyDisj sync.Once
func Get_applyDisj() gopurs_runtime.Value {
	once_applyDisj.Do(func() {
		applyDisj = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, v1_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorDisj()
}))
	})
	return applyDisj
}

var bindDisj gopurs_runtime.Value
var once_bindDisj sync.Once
func Get_bindDisj() gopurs_runtime.Value {
	once_bindDisj.Do(func() {
		bindDisj = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDisj()
}))
	})
	return bindDisj
}

var applicativeDisj gopurs_runtime.Value
var once_applicativeDisj sync.Once
func Get_applicativeDisj() gopurs_runtime.Value {
	once_applicativeDisj.Do(func() {
		applicativeDisj = gopurs_runtime.RecordDict2("pure", "Apply0", Get_Disj(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyDisj()
}))
	})
	return applicativeDisj
}

var monadDisj gopurs_runtime.Value
var once_monadDisj sync.Once
func Get_monadDisj() gopurs_runtime.Value {
	once_monadDisj.Do(func() {
		monadDisj = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeDisj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindDisj()
}))
	})
	return monadDisj
}





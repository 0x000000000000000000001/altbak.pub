package Data_Monoid_Alternate

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Alternate gopurs_runtime.Value
var once_Alternate sync.Once
func Get_Alternate() gopurs_runtime.Value {
	once_Alternate.Do(func() {
		Alternate = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return Alternate
}

var showAlternate gopurs_runtime.Value
var once_showAlternate sync.Once
func Get_showAlternate() gopurs_runtime.Value {
	once_showAlternate.Do(func() {
		showAlternate = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Alternate " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal() + ")")
}))
}()
})
	})
	return showAlternate
}

var semigroupAlternate gopurs_runtime.Value
var once_semigroupAlternate sync.Once
func Get_semigroupAlternate() gopurs_runtime.Value {
	once_semigroupAlternate.Do(func() {
		semigroupAlternate = gopurs_runtime.Func(func(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), v_1, v1_2)
}))
}()
})
	})
	return semigroupAlternate
}

var plusAlternate gopurs_runtime.Value
var once_plusAlternate sync.Once
func Get_plusAlternate() gopurs_runtime.Value {
	once_plusAlternate.Do(func() {
		plusAlternate = gopurs_runtime.Func(func(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
return dictPlus_0
}()
})
	})
	return plusAlternate
}

var ordAlternate gopurs_runtime.Value
var once_ordAlternate sync.Once
func Get_ordAlternate() gopurs_runtime.Value {
	once_ordAlternate.Do(func() {
		ordAlternate = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}()
})
	})
	return ordAlternate
}

var ord1Alternate gopurs_runtime.Value
var once_ord1Alternate sync.Once
func Get_ord1Alternate() gopurs_runtime.Value {
	once_ord1Alternate.Do(func() {
		ord1Alternate = gopurs_runtime.Func(func(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
return dictOrd1_0
}()
})
	})
	return ord1Alternate
}

var newtypeAlternate gopurs_runtime.Value
var once_newtypeAlternate sync.Once
func Get_newtypeAlternate() gopurs_runtime.Value {
	once_newtypeAlternate.Do(func() {
		newtypeAlternate = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeAlternate
}

var monoidAlternate gopurs_runtime.Value
var once_monoidAlternate sync.Once
func Get_monoidAlternate() gopurs_runtime.Value {
	once_monoidAlternate.Do(func() {
		monoidAlternate = gopurs_runtime.Func(func(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupAlternate1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "alt"), v_2, v1_3)
}))
_ = semigroupAlternate1_2_1
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictPlus_0, "empty"), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAlternate1_2_1
}))
}()
})
	})
	return monoidAlternate
}

var monadAlternate gopurs_runtime.Value
var once_monadAlternate sync.Once
func Get_monadAlternate() gopurs_runtime.Value {
	once_monadAlternate.Do(func() {
		monadAlternate = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return dictMonad_0
}()
})
	})
	return monadAlternate
}

var functorAlternate gopurs_runtime.Value
var once_functorAlternate sync.Once
func Get_functorAlternate() gopurs_runtime.Value {
	once_functorAlternate.Do(func() {
		functorAlternate = gopurs_runtime.Func(func(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return dictFunctor_0
}()
})
	})
	return functorAlternate
}

var extendAlternate gopurs_runtime.Value
var once_extendAlternate sync.Once
func Get_extendAlternate() gopurs_runtime.Value {
	once_extendAlternate.Do(func() {
		extendAlternate = gopurs_runtime.Func(func(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
return dictExtend_0
}()
})
	})
	return extendAlternate
}

var eqAlternate gopurs_runtime.Value
var once_eqAlternate sync.Once
func Get_eqAlternate() gopurs_runtime.Value {
	once_eqAlternate.Do(func() {
		eqAlternate = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}()
})
	})
	return eqAlternate
}

var eq1Alternate gopurs_runtime.Value
var once_eq1Alternate sync.Once
func Get_eq1Alternate() gopurs_runtime.Value {
	once_eq1Alternate.Do(func() {
		eq1Alternate = gopurs_runtime.Func(func(dictEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
return dictEq1_0
}()
})
	})
	return eq1Alternate
}

var comonadAlternate gopurs_runtime.Value
var once_comonadAlternate sync.Once
func Get_comonadAlternate() gopurs_runtime.Value {
	once_comonadAlternate.Do(func() {
		comonadAlternate = gopurs_runtime.Func(func(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
return dictComonad_0
}()
})
	})
	return comonadAlternate
}

var boundedAlternate gopurs_runtime.Value
var once_boundedAlternate sync.Once
func Get_boundedAlternate() gopurs_runtime.Value {
	once_boundedAlternate.Do(func() {
		boundedAlternate = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}()
})
	})
	return boundedAlternate
}

var bindAlternate gopurs_runtime.Value
var once_bindAlternate sync.Once
func Get_bindAlternate() gopurs_runtime.Value {
	once_bindAlternate.Do(func() {
		bindAlternate = gopurs_runtime.Func(func(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
return dictBind_0
}()
})
	})
	return bindAlternate
}

var applyAlternate gopurs_runtime.Value
var once_applyAlternate sync.Once
func Get_applyAlternate() gopurs_runtime.Value {
	once_applyAlternate.Do(func() {
		applyAlternate = gopurs_runtime.Func(func(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
return dictApply_0
}()
})
	})
	return applyAlternate
}

var applicativeAlternate gopurs_runtime.Value
var once_applicativeAlternate sync.Once
func Get_applicativeAlternate() gopurs_runtime.Value {
	once_applicativeAlternate.Do(func() {
		applicativeAlternate = gopurs_runtime.Func(func(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return dictApplicative_0
}()
})
	})
	return applicativeAlternate
}

var alternativeAlternate gopurs_runtime.Value
var once_alternativeAlternate sync.Once
func Get_alternativeAlternate() gopurs_runtime.Value {
	once_alternativeAlternate.Do(func() {
		alternativeAlternate = gopurs_runtime.Func(func(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
return dictAlternative_0
}()
})
	})
	return alternativeAlternate
}

var altAlternate gopurs_runtime.Value
var once_altAlternate sync.Once
func Get_altAlternate() gopurs_runtime.Value {
	once_altAlternate.Do(func() {
		altAlternate = gopurs_runtime.Func(func(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
return dictAlt_0
}()
})
	})
	return altAlternate
}





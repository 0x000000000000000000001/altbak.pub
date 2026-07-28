package Data_Monoid_Alternate

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var cache_Alternate gopurs_runtime.Value
var once_Alternate sync.Once
func Get_Alternate() gopurs_runtime.Value {
	once_Alternate.Do(func() {
		cache_Alternate = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Alternate(x_0_box)
})
	})
	return cache_Alternate
}

var cache_showAlternate gopurs_runtime.Value
var once_showAlternate sync.Once
func Get_showAlternate() gopurs_runtime.Value {
	once_showAlternate.Do(func() {
		cache_showAlternate = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showAlternate(dictShow_0_box)
})
	})
	return cache_showAlternate
}

var cache_semigroupAlternate gopurs_runtime.Value
var once_semigroupAlternate sync.Once
func Get_semigroupAlternate() gopurs_runtime.Value {
	once_semigroupAlternate.Do(func() {
		cache_semigroupAlternate = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupAlternate(dictAlt_0_box)
})
	})
	return cache_semigroupAlternate
}

var cache_plusAlternate gopurs_runtime.Value
var once_plusAlternate sync.Once
func Get_plusAlternate() gopurs_runtime.Value {
	once_plusAlternate.Do(func() {
		cache_plusAlternate = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusAlternate(dictPlus_0_box)
})
	})
	return cache_plusAlternate
}

var cache_ordAlternate gopurs_runtime.Value
var once_ordAlternate sync.Once
func Get_ordAlternate() gopurs_runtime.Value {
	once_ordAlternate.Do(func() {
		cache_ordAlternate = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordAlternate(dictOrd_0_box)
})
	})
	return cache_ordAlternate
}

var cache_ord1Alternate gopurs_runtime.Value
var once_ord1Alternate sync.Once
func Get_ord1Alternate() gopurs_runtime.Value {
	once_ord1Alternate.Do(func() {
		cache_ord1Alternate = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1Alternate(dictOrd1_0_box)
})
	})
	return cache_ord1Alternate
}

var cache_newtypeAlternate gopurs_runtime.Value
var once_newtypeAlternate sync.Once
func Get_newtypeAlternate() gopurs_runtime.Value {
	once_newtypeAlternate.Do(func() {
		cache_newtypeAlternate = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeAlternate
}

var cache_monoidAlternate gopurs_runtime.Value
var once_monoidAlternate sync.Once
func Get_monoidAlternate() gopurs_runtime.Value {
	once_monoidAlternate.Do(func() {
		cache_monoidAlternate = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidAlternate(dictPlus_0_box)
})
	})
	return cache_monoidAlternate
}

var cache_monadAlternate gopurs_runtime.Value
var once_monadAlternate sync.Once
func Get_monadAlternate() gopurs_runtime.Value {
	once_monadAlternate.Do(func() {
		cache_monadAlternate = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAlternate(dictMonad_0_box)
})
	})
	return cache_monadAlternate
}

var cache_functorAlternate gopurs_runtime.Value
var once_functorAlternate sync.Once
func Get_functorAlternate() gopurs_runtime.Value {
	once_functorAlternate.Do(func() {
		cache_functorAlternate = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorAlternate(dictFunctor_0_box)
})
	})
	return cache_functorAlternate
}

var cache_extendAlternate gopurs_runtime.Value
var once_extendAlternate sync.Once
func Get_extendAlternate() gopurs_runtime.Value {
	once_extendAlternate.Do(func() {
		cache_extendAlternate = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extendAlternate(dictExtend_0_box)
})
	})
	return cache_extendAlternate
}

var cache_eqAlternate gopurs_runtime.Value
var once_eqAlternate sync.Once
func Get_eqAlternate() gopurs_runtime.Value {
	once_eqAlternate.Do(func() {
		cache_eqAlternate = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqAlternate(dictEq_0_box)
})
	})
	return cache_eqAlternate
}

var cache_eq1Alternate gopurs_runtime.Value
var once_eq1Alternate sync.Once
func Get_eq1Alternate() gopurs_runtime.Value {
	once_eq1Alternate.Do(func() {
		cache_eq1Alternate = gopurs_runtime.Func(func(dictEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1Alternate(dictEq1_0_box)
})
	})
	return cache_eq1Alternate
}

var cache_comonadAlternate gopurs_runtime.Value
var once_comonadAlternate sync.Once
func Get_comonadAlternate() gopurs_runtime.Value {
	once_comonadAlternate.Do(func() {
		cache_comonadAlternate = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadAlternate(dictComonad_0_box)
})
	})
	return cache_comonadAlternate
}

var cache_boundedAlternate gopurs_runtime.Value
var once_boundedAlternate sync.Once
func Get_boundedAlternate() gopurs_runtime.Value {
	once_boundedAlternate.Do(func() {
		cache_boundedAlternate = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedAlternate(dictBounded_0_box)
})
	})
	return cache_boundedAlternate
}

var cache_bindAlternate gopurs_runtime.Value
var once_bindAlternate sync.Once
func Get_bindAlternate() gopurs_runtime.Value {
	once_bindAlternate.Do(func() {
		cache_bindAlternate = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindAlternate(dictBind_0_box)
})
	})
	return cache_bindAlternate
}

var cache_applyAlternate gopurs_runtime.Value
var once_applyAlternate sync.Once
func Get_applyAlternate() gopurs_runtime.Value {
	once_applyAlternate.Do(func() {
		cache_applyAlternate = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyAlternate(dictApply_0_box)
})
	})
	return cache_applyAlternate
}

var cache_applicativeAlternate gopurs_runtime.Value
var once_applicativeAlternate sync.Once
func Get_applicativeAlternate() gopurs_runtime.Value {
	once_applicativeAlternate.Do(func() {
		cache_applicativeAlternate = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeAlternate(dictApplicative_0_box)
})
	})
	return cache_applicativeAlternate
}

var cache_alternativeAlternate gopurs_runtime.Value
var once_alternativeAlternate sync.Once
func Get_alternativeAlternate() gopurs_runtime.Value {
	once_alternativeAlternate.Do(func() {
		cache_alternativeAlternate = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeAlternate(dictAlternative_0_box)
})
	})
	return cache_alternativeAlternate
}

var cache_altAlternate gopurs_runtime.Value
var once_altAlternate sync.Once
func Get_altAlternate() gopurs_runtime.Value {
	once_altAlternate.Do(func() {
		cache_altAlternate = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altAlternate(dictAlt_0_box)
})
	})
	return cache_altAlternate
}

func Call_Alternate(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showAlternate(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Alternate "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1), gopurs_runtime.Str(")")))
}))
}

func Call_semigroupAlternate(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), v_1, v1_2)
}))
}

func Call_plusAlternate(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
return dictPlus_0
}

func Call_ordAlternate(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_ord1Alternate(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
return dictOrd1_0
}

func Call_monoidAlternate(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupAlternate1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "alt"), v_2, v1_3)
}))
_ = semigroupAlternate1_2_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAlternate1_2_1
}), gopurs_runtime.RecordGet(dictPlus_0, "empty"))
}

func Call_monadAlternate(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return dictMonad_0
}

func Call_functorAlternate(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return dictFunctor_0
}

func Call_extendAlternate(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
return dictExtend_0
}

func Call_eqAlternate(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_eq1Alternate(dictEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
return dictEq1_0
}

func Call_comonadAlternate(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
return dictComonad_0
}

func Call_boundedAlternate(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return dictBounded_0
}

func Call_bindAlternate(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
return dictBind_0
}

func Call_applyAlternate(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
return dictApply_0
}

func Call_applicativeAlternate(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return dictApplicative_0
}

func Call_alternativeAlternate(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
return dictAlternative_0
}

func Call_altAlternate(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
return dictAlt_0
}



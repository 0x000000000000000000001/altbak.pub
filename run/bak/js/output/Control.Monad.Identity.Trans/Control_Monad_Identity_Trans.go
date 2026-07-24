package Control_Monad_Identity_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var IdentityT gopurs_runtime.Value
var once_IdentityT sync.Once
func Get_IdentityT() gopurs_runtime.Value {
	once_IdentityT.Do(func() {
		IdentityT = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return IdentityT
}

var monadSTIdentityT gopurs_runtime.Value
var once_monadSTIdentityT sync.Once
func Get_monadSTIdentityT() gopurs_runtime.Value {
	once_monadSTIdentityT.Do(func() {
		monadSTIdentityT = gopurs_runtime.Func(func(dictMonadST_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonadST_0
})
	})
	return monadSTIdentityT
}

var traversableIdentityT gopurs_runtime.Value
var once_traversableIdentityT sync.Once
func Get_traversableIdentityT() gopurs_runtime.Value {
	once_traversableIdentityT.Do(func() {
		traversableIdentityT = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictTraversable_0
})
	})
	return traversableIdentityT
}

var runIdentityT gopurs_runtime.Value
var once_runIdentityT sync.Once
func Get_runIdentityT() gopurs_runtime.Value {
	once_runIdentityT.Do(func() {
		runIdentityT = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return runIdentityT
}

var plusIdentityT gopurs_runtime.Value
var once_plusIdentityT sync.Once
func Get_plusIdentityT() gopurs_runtime.Value {
	once_plusIdentityT.Do(func() {
		plusIdentityT = gopurs_runtime.Func(func(dictPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictPlus_0
})
	})
	return plusIdentityT
}

var newtypeIdentityT gopurs_runtime.Value
var once_newtypeIdentityT sync.Once
func Get_newtypeIdentityT() gopurs_runtime.Value {
	once_newtypeIdentityT.Do(func() {
		newtypeIdentityT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeIdentityT
}

var monadWriterIdentityT gopurs_runtime.Value
var once_monadWriterIdentityT sync.Once
func Get_monadWriterIdentityT() gopurs_runtime.Value {
	once_monadWriterIdentityT.Do(func() {
		monadWriterIdentityT = gopurs_runtime.Func(func(dictMonadWriter_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonadWriter_0
})
	})
	return monadWriterIdentityT
}

var monadTransIdentityT gopurs_runtime.Value
var once_monadTransIdentityT sync.Once
func Get_monadTransIdentityT() gopurs_runtime.Value {
	once_monadTransIdentityT.Do(func() {
		monadTransIdentityT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_IdentityT()
}))
	})
	return monadTransIdentityT
}

var monadThrowIdentityT gopurs_runtime.Value
var once_monadThrowIdentityT sync.Once
func Get_monadThrowIdentityT() gopurs_runtime.Value {
	once_monadThrowIdentityT.Do(func() {
		monadThrowIdentityT = gopurs_runtime.Func(func(dictMonadThrow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonadThrow_0
})
	})
	return monadThrowIdentityT
}

var monadTellIdentityT gopurs_runtime.Value
var once_monadTellIdentityT sync.Once
func Get_monadTellIdentityT() gopurs_runtime.Value {
	once_monadTellIdentityT.Do(func() {
		monadTellIdentityT = gopurs_runtime.Func(func(dictMonadTell_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonadTell_0
})
	})
	return monadTellIdentityT
}

var monadStateIdentityT gopurs_runtime.Value
var once_monadStateIdentityT sync.Once
func Get_monadStateIdentityT() gopurs_runtime.Value {
	once_monadStateIdentityT.Do(func() {
		monadStateIdentityT = gopurs_runtime.Func(func(dictMonadState_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonadState_0
})
	})
	return monadStateIdentityT
}

var monadRecIdentityT gopurs_runtime.Value
var once_monadRecIdentityT sync.Once
func Get_monadRecIdentityT() gopurs_runtime.Value {
	once_monadRecIdentityT.Do(func() {
		monadRecIdentityT = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonadRec_0
})
	})
	return monadRecIdentityT
}

var monadReaderIdentityT gopurs_runtime.Value
var once_monadReaderIdentityT sync.Once
func Get_monadReaderIdentityT() gopurs_runtime.Value {
	once_monadReaderIdentityT.Do(func() {
		monadReaderIdentityT = gopurs_runtime.Func(func(dictMonadReader_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonadReader_0
})
	})
	return monadReaderIdentityT
}

var monadPlusIdentityT gopurs_runtime.Value
var once_monadPlusIdentityT sync.Once
func Get_monadPlusIdentityT() gopurs_runtime.Value {
	once_monadPlusIdentityT.Do(func() {
		monadPlusIdentityT = gopurs_runtime.Func(func(dictMonadPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonadPlus_0
})
	})
	return monadPlusIdentityT
}

var monadIdentityT gopurs_runtime.Value
var once_monadIdentityT sync.Once
func Get_monadIdentityT() gopurs_runtime.Value {
	once_monadIdentityT.Do(func() {
		monadIdentityT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonad_0
})
	})
	return monadIdentityT
}

var monadErrorIdentityT gopurs_runtime.Value
var once_monadErrorIdentityT sync.Once
func Get_monadErrorIdentityT() gopurs_runtime.Value {
	once_monadErrorIdentityT.Do(func() {
		monadErrorIdentityT = gopurs_runtime.Func(func(dictMonadError_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonadError_0
})
	})
	return monadErrorIdentityT
}

var monadEffectIdentityT gopurs_runtime.Value
var once_monadEffectIdentityT sync.Once
func Get_monadEffectIdentityT() gopurs_runtime.Value {
	once_monadEffectIdentityT.Do(func() {
		monadEffectIdentityT = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonadEffect_0
})
	})
	return monadEffectIdentityT
}

var monadContIdentityT gopurs_runtime.Value
var once_monadContIdentityT sync.Once
func Get_monadContIdentityT() gopurs_runtime.Value {
	once_monadContIdentityT.Do(func() {
		monadContIdentityT = gopurs_runtime.Func(func(dictMonadCont_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonadCont_0
})
	})
	return monadContIdentityT
}

var monadAskIdentityT gopurs_runtime.Value
var once_monadAskIdentityT sync.Once
func Get_monadAskIdentityT() gopurs_runtime.Value {
	once_monadAskIdentityT.Do(func() {
		monadAskIdentityT = gopurs_runtime.Func(func(dictMonadAsk_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonadAsk_0
})
	})
	return monadAskIdentityT
}

var mapIdentityT gopurs_runtime.Value
var once_mapIdentityT sync.Once
func Get_mapIdentityT() gopurs_runtime.Value {
	once_mapIdentityT.Do(func() {
		mapIdentityT = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v_1)
})
	})
	return mapIdentityT
}

var functorIdentityT gopurs_runtime.Value
var once_functorIdentityT sync.Once
func Get_functorIdentityT() gopurs_runtime.Value {
	once_functorIdentityT.Do(func() {
		functorIdentityT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictFunctor_0
})
	})
	return functorIdentityT
}

var foldableIdentityT gopurs_runtime.Value
var once_foldableIdentityT sync.Once
func Get_foldableIdentityT() gopurs_runtime.Value {
	once_foldableIdentityT.Do(func() {
		foldableIdentityT = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictFoldable_0
})
	})
	return foldableIdentityT
}

var extendIdentityI gopurs_runtime.Value
var once_extendIdentityI sync.Once
func Get_extendIdentityI() gopurs_runtime.Value {
	once_extendIdentityI.Do(func() {
		extendIdentityI = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, x_4)
}), v_3)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}))
})
	})
	return extendIdentityI
}

var eqIdentityT gopurs_runtime.Value
var once_eqIdentityT sync.Once
func Get_eqIdentityT() gopurs_runtime.Value {
	once_eqIdentityT.Do(func() {
		eqIdentityT = gopurs_runtime.Func2(func(dictEq1_0 gopurs_runtime.Value, dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
eq11_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), dictEq_1)
_ = eq11_2_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_2_0, x_3, y_4)
}))
})
	})
	return eqIdentityT
}

var ordIdentityT gopurs_runtime.Value
var once_ordIdentityT sync.Once
func Get_ordIdentityT() gopurs_runtime.Value {
	once_ordIdentityT.Do(func() {
		ordIdentityT = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), dictOrd_2)
_ = compare11_3_1
eq11_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_2, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_4_2
eqIdentityT2_5_3 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_4_2, x_5, y_6)
}))
_ = eqIdentityT2_5_3
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, y_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_3_1, x_6, y_7)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqIdentityT2_5_3
}))
})
})
	})
	return ordIdentityT
}

var eq1IdentityT gopurs_runtime.Value
var once_eq1IdentityT sync.Once
func Get_eq1IdentityT() gopurs_runtime.Value {
	once_eq1IdentityT.Do(func() {
		eq1IdentityT = gopurs_runtime.Func(func(dictEq1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), dictEq_1)
}))
})
	})
	return eq1IdentityT
}

var ord1IdentityT gopurs_runtime.Value
var once_ord1IdentityT sync.Once
func Get_ord1IdentityT() gopurs_runtime.Value {
	once_ord1IdentityT.Do(func() {
		ord1IdentityT = gopurs_runtime.Func(func(dictOrd1_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_1
eq1IdentityT1_3_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "eq1"), dictEq_3)
}))
_ = eq1IdentityT1_3_2
return gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), dictOrd_4)
_ = compare11_5_3
eq11_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_4, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_6_4
eqIdentityT2_7_5 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_7 gopurs_runtime.Value, y_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_6_4, x_7, y_8)
}))
_ = eqIdentityT2_7_5
return gopurs_runtime.RecordGet(gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_5_3, x_8, y_9)
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return eqIdentityT2_7_5
})), "compare")
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1IdentityT1_3_2
}))
})
	})
	return ord1IdentityT
}

var comonadIdentityT gopurs_runtime.Value
var once_comonadIdentityT sync.Once
func Get_comonadIdentityT() gopurs_runtime.Value {
	once_comonadIdentityT.Do(func() {
		comonadIdentityT = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
extendIdentityI1_3_2 := gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "extend"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, x_5)
}), v_4)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}))
_ = extendIdentityI1_3_2
return gopurs_runtime.RecordDict2("extract", "Extend0", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), x_4)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return extendIdentityI1_3_2
}))
})
	})
	return comonadIdentityT
}

var bindIdentityT gopurs_runtime.Value
var once_bindIdentityT sync.Once
func Get_bindIdentityT() gopurs_runtime.Value {
	once_bindIdentityT.Do(func() {
		bindIdentityT = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictBind_0
})
	})
	return bindIdentityT
}

var applyIdentityT gopurs_runtime.Value
var once_applyIdentityT sync.Once
func Get_applyIdentityT() gopurs_runtime.Value {
	once_applyIdentityT.Do(func() {
		applyIdentityT = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictApply_0
})
	})
	return applyIdentityT
}

var applicativeIdentityT gopurs_runtime.Value
var once_applicativeIdentityT sync.Once
func Get_applicativeIdentityT() gopurs_runtime.Value {
	once_applicativeIdentityT.Do(func() {
		applicativeIdentityT = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictApplicative_0
})
	})
	return applicativeIdentityT
}

var alternativeIdentityT gopurs_runtime.Value
var once_alternativeIdentityT sync.Once
func Get_alternativeIdentityT() gopurs_runtime.Value {
	once_alternativeIdentityT.Do(func() {
		alternativeIdentityT = gopurs_runtime.Func(func(dictAlternative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictAlternative_0
})
	})
	return alternativeIdentityT
}

var altIdentityT gopurs_runtime.Value
var once_altIdentityT sync.Once
func Get_altIdentityT() gopurs_runtime.Value {
	once_altIdentityT.Do(func() {
		altIdentityT = gopurs_runtime.Func(func(dictAlt_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dictAlt_0
})
	})
	return altIdentityT
}





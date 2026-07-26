package Control_Monad_Identity_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_IdentityT gopurs_runtime.Value
var once_IdentityT sync.Once
func Get_IdentityT() gopurs_runtime.Value {
	once_IdentityT.Do(func() {
		cache_IdentityT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_IdentityT(x_0_box)
})
	})
	return cache_IdentityT
}

var cache_monadSTIdentityT gopurs_runtime.Value
var once_monadSTIdentityT sync.Once
func Get_monadSTIdentityT() gopurs_runtime.Value {
	once_monadSTIdentityT.Do(func() {
		cache_monadSTIdentityT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadSTIdentityT(dictMonadST_0_box)
})
	})
	return cache_monadSTIdentityT
}

var cache_traversableIdentityT gopurs_runtime.Value
var once_traversableIdentityT sync.Once
func Get_traversableIdentityT() gopurs_runtime.Value {
	once_traversableIdentityT.Do(func() {
		cache_traversableIdentityT = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traversableIdentityT(dictTraversable_0_box)
})
	})
	return cache_traversableIdentityT
}

var cache_runIdentityT gopurs_runtime.Value
var once_runIdentityT sync.Once
func Get_runIdentityT() gopurs_runtime.Value {
	once_runIdentityT.Do(func() {
		cache_runIdentityT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runIdentityT(v_0_box)
})
	})
	return cache_runIdentityT
}

var cache_plusIdentityT gopurs_runtime.Value
var once_plusIdentityT sync.Once
func Get_plusIdentityT() gopurs_runtime.Value {
	once_plusIdentityT.Do(func() {
		cache_plusIdentityT = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusIdentityT(dictPlus_0_box)
})
	})
	return cache_plusIdentityT
}

var cache_newtypeIdentityT gopurs_runtime.Value
var once_newtypeIdentityT sync.Once
func Get_newtypeIdentityT() gopurs_runtime.Value {
	once_newtypeIdentityT.Do(func() {
		cache_newtypeIdentityT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeIdentityT
}

var cache_monadWriterIdentityT gopurs_runtime.Value
var once_monadWriterIdentityT sync.Once
func Get_monadWriterIdentityT() gopurs_runtime.Value {
	once_monadWriterIdentityT.Do(func() {
		cache_monadWriterIdentityT = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadWriterIdentityT(dictMonadWriter_0_box)
})
	})
	return cache_monadWriterIdentityT
}

var cache_monadTransIdentityT gopurs_runtime.Value
var once_monadTransIdentityT sync.Once
func Get_monadTransIdentityT() gopurs_runtime.Value {
	once_monadTransIdentityT.Do(func() {
		cache_monadTransIdentityT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_IdentityT()
}))
	})
	return cache_monadTransIdentityT
}

var cache_monadThrowIdentityT gopurs_runtime.Value
var once_monadThrowIdentityT sync.Once
func Get_monadThrowIdentityT() gopurs_runtime.Value {
	once_monadThrowIdentityT.Do(func() {
		cache_monadThrowIdentityT = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadThrowIdentityT(dictMonadThrow_0_box)
})
	})
	return cache_monadThrowIdentityT
}

var cache_monadTellIdentityT gopurs_runtime.Value
var once_monadTellIdentityT sync.Once
func Get_monadTellIdentityT() gopurs_runtime.Value {
	once_monadTellIdentityT.Do(func() {
		cache_monadTellIdentityT = gopurs_runtime.Func(func(dictMonadTell_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadTellIdentityT(dictMonadTell_0_box)
})
	})
	return cache_monadTellIdentityT
}

var cache_monadStateIdentityT gopurs_runtime.Value
var once_monadStateIdentityT sync.Once
func Get_monadStateIdentityT() gopurs_runtime.Value {
	once_monadStateIdentityT.Do(func() {
		cache_monadStateIdentityT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadStateIdentityT(dictMonadState_0_box)
})
	})
	return cache_monadStateIdentityT
}

var cache_monadRecIdentityT gopurs_runtime.Value
var once_monadRecIdentityT sync.Once
func Get_monadRecIdentityT() gopurs_runtime.Value {
	once_monadRecIdentityT.Do(func() {
		cache_monadRecIdentityT = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadRecIdentityT(dictMonadRec_0_box)
})
	})
	return cache_monadRecIdentityT
}

var cache_monadReaderIdentityT gopurs_runtime.Value
var once_monadReaderIdentityT sync.Once
func Get_monadReaderIdentityT() gopurs_runtime.Value {
	once_monadReaderIdentityT.Do(func() {
		cache_monadReaderIdentityT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadReaderIdentityT(dictMonadReader_0_box)
})
	})
	return cache_monadReaderIdentityT
}

var cache_monadPlusIdentityT gopurs_runtime.Value
var once_monadPlusIdentityT sync.Once
func Get_monadPlusIdentityT() gopurs_runtime.Value {
	once_monadPlusIdentityT.Do(func() {
		cache_monadPlusIdentityT = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadPlusIdentityT(dictMonadPlus_0_box)
})
	})
	return cache_monadPlusIdentityT
}

var cache_monadIdentityT gopurs_runtime.Value
var once_monadIdentityT sync.Once
func Get_monadIdentityT() gopurs_runtime.Value {
	once_monadIdentityT.Do(func() {
		cache_monadIdentityT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadIdentityT(dictMonad_0_box)
})
	})
	return cache_monadIdentityT
}

var cache_monadErrorIdentityT gopurs_runtime.Value
var once_monadErrorIdentityT sync.Once
func Get_monadErrorIdentityT() gopurs_runtime.Value {
	once_monadErrorIdentityT.Do(func() {
		cache_monadErrorIdentityT = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadErrorIdentityT(dictMonadError_0_box)
})
	})
	return cache_monadErrorIdentityT
}

var cache_monadEffectIdentityT gopurs_runtime.Value
var once_monadEffectIdentityT sync.Once
func Get_monadEffectIdentityT() gopurs_runtime.Value {
	once_monadEffectIdentityT.Do(func() {
		cache_monadEffectIdentityT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadEffectIdentityT(dictMonadEffect_0_box)
})
	})
	return cache_monadEffectIdentityT
}

var cache_monadContIdentityT gopurs_runtime.Value
var once_monadContIdentityT sync.Once
func Get_monadContIdentityT() gopurs_runtime.Value {
	once_monadContIdentityT.Do(func() {
		cache_monadContIdentityT = gopurs_runtime.Func(func(dictMonadCont_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadContIdentityT(dictMonadCont_0_box)
})
	})
	return cache_monadContIdentityT
}

var cache_monadAskIdentityT gopurs_runtime.Value
var once_monadAskIdentityT sync.Once
func Get_monadAskIdentityT() gopurs_runtime.Value {
	once_monadAskIdentityT.Do(func() {
		cache_monadAskIdentityT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAskIdentityT(dictMonadAsk_0_box)
})
	})
	return cache_monadAskIdentityT
}

var cache_mapIdentityT gopurs_runtime.Value
var once_mapIdentityT sync.Once
func Get_mapIdentityT() gopurs_runtime.Value {
	once_mapIdentityT.Do(func() {
		cache_mapIdentityT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapIdentityT(f_0_box, v_1_box)
})
	})
	return cache_mapIdentityT
}

var cache_functorIdentityT gopurs_runtime.Value
var once_functorIdentityT sync.Once
func Get_functorIdentityT() gopurs_runtime.Value {
	once_functorIdentityT.Do(func() {
		cache_functorIdentityT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorIdentityT(dictFunctor_0_box)
})
	})
	return cache_functorIdentityT
}

var cache_foldableIdentityT gopurs_runtime.Value
var once_foldableIdentityT sync.Once
func Get_foldableIdentityT() gopurs_runtime.Value {
	once_foldableIdentityT.Do(func() {
		cache_foldableIdentityT = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableIdentityT(dictFoldable_0_box)
})
	})
	return cache_foldableIdentityT
}

var cache_extendIdentityI gopurs_runtime.Value
var once_extendIdentityI sync.Once
func Get_extendIdentityI() gopurs_runtime.Value {
	once_extendIdentityI.Do(func() {
		cache_extendIdentityI = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extendIdentityI(dictExtend_0_box)
})
	})
	return cache_extendIdentityI
}

var cache_eqIdentityT gopurs_runtime.Value
var once_eqIdentityT sync.Once
func Get_eqIdentityT() gopurs_runtime.Value {
	once_eqIdentityT.Do(func() {
		cache_eqIdentityT = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqIdentityT(dictEq1_0_box, dictEq_1_box)
})
	})
	return cache_eqIdentityT
}

var cache_ordIdentityT gopurs_runtime.Value
var once_ordIdentityT sync.Once
func Get_ordIdentityT() gopurs_runtime.Value {
	once_ordIdentityT.Do(func() {
		cache_ordIdentityT = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordIdentityT(dictOrd1_0_box)
})
	})
	return cache_ordIdentityT
}

var cache_eq1IdentityT gopurs_runtime.Value
var once_eq1IdentityT sync.Once
func Get_eq1IdentityT() gopurs_runtime.Value {
	once_eq1IdentityT.Do(func() {
		cache_eq1IdentityT = gopurs_runtime.Func(func(dictEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1IdentityT(dictEq1_0_box)
})
	})
	return cache_eq1IdentityT
}

var cache_ord1IdentityT gopurs_runtime.Value
var once_ord1IdentityT sync.Once
func Get_ord1IdentityT() gopurs_runtime.Value {
	once_ord1IdentityT.Do(func() {
		cache_ord1IdentityT = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1IdentityT(dictOrd1_0_box)
})
	})
	return cache_ord1IdentityT
}

var cache_comonadIdentityT gopurs_runtime.Value
var once_comonadIdentityT sync.Once
func Get_comonadIdentityT() gopurs_runtime.Value {
	once_comonadIdentityT.Do(func() {
		cache_comonadIdentityT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadIdentityT(dictComonad_0_box)
})
	})
	return cache_comonadIdentityT
}

var cache_bindIdentityT gopurs_runtime.Value
var once_bindIdentityT sync.Once
func Get_bindIdentityT() gopurs_runtime.Value {
	once_bindIdentityT.Do(func() {
		cache_bindIdentityT = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindIdentityT(dictBind_0_box)
})
	})
	return cache_bindIdentityT
}

var cache_applyIdentityT gopurs_runtime.Value
var once_applyIdentityT sync.Once
func Get_applyIdentityT() gopurs_runtime.Value {
	once_applyIdentityT.Do(func() {
		cache_applyIdentityT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyIdentityT(dictApply_0_box)
})
	})
	return cache_applyIdentityT
}

var cache_applicativeIdentityT gopurs_runtime.Value
var once_applicativeIdentityT sync.Once
func Get_applicativeIdentityT() gopurs_runtime.Value {
	once_applicativeIdentityT.Do(func() {
		cache_applicativeIdentityT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeIdentityT(dictApplicative_0_box)
})
	})
	return cache_applicativeIdentityT
}

var cache_alternativeIdentityT gopurs_runtime.Value
var once_alternativeIdentityT sync.Once
func Get_alternativeIdentityT() gopurs_runtime.Value {
	once_alternativeIdentityT.Do(func() {
		cache_alternativeIdentityT = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeIdentityT(dictAlternative_0_box)
})
	})
	return cache_alternativeIdentityT
}

var cache_altIdentityT gopurs_runtime.Value
var once_altIdentityT sync.Once
func Get_altIdentityT() gopurs_runtime.Value {
	once_altIdentityT.Do(func() {
		cache_altIdentityT = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altIdentityT(dictAlt_0_box)
})
	})
	return cache_altIdentityT
}

func Call_IdentityT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_monadSTIdentityT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
return dictMonadST_0
}

func Call_traversableIdentityT(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
return dictTraversable_0
}

func Call_runIdentityT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_plusIdentityT(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
return dictPlus_0
}

func Call_monadWriterIdentityT(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
return dictMonadWriter_0
}

func Call_monadThrowIdentityT(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
return dictMonadThrow_0
}

func Call_monadTellIdentityT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
return dictMonadTell_0
}

func Call_monadStateIdentityT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
return dictMonadState_0
}

func Call_monadRecIdentityT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
return dictMonadRec_0
}

func Call_monadReaderIdentityT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
return dictMonadReader_0
}

func Call_monadPlusIdentityT(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
_ = dictMonadPlus_0
return dictMonadPlus_0
}

func Call_monadIdentityT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return dictMonad_0
}

func Call_monadErrorIdentityT(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
return dictMonadError_0
}

func Call_monadEffectIdentityT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
return dictMonadEffect_0
}

func Call_monadContIdentityT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
return dictMonadCont_0
}

func Call_monadAskIdentityT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
return dictMonadAsk_0
}

func Call_mapIdentityT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_functorIdentityT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return dictFunctor_0
}

func Call_foldableIdentityT(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return dictFoldable_0
}

func Call_extendIdentityI(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictExtend_0.UnsafePtr)).V0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, x_4)
}), v_3)
}))
}

func Call_eqIdentityT(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
_ = dictEq_1
eq11_2_0 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictEq1_0.UnsafePtr)).V0, dictEq_1)
_ = eq11_2_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_2_0, x_3, y_4)
}))
}

func Call_ordIdentityT(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_3_1 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictOrd1_0.UnsafePtr)).V0, dictOrd_2)
_ = compare11_3_1
eq11_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_2, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_4_2
eqIdentityT2_5_3 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_4_2, x_5, y_6)
}))
_ = eqIdentityT2_5_3
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return eqIdentityT2_5_3
}), gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, y_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_3_1, x_6, y_7)
}))
})
}

func Call_eq1IdentityT(dictEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictEq1_0.UnsafePtr)).V0, dictEq_1)
}))
}

func Call_ord1IdentityT(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_1
eq1IdentityT1_3_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "eq1"), dictEq_3)
}))
_ = eq1IdentityT1_3_2
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1IdentityT1_3_2
}), gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_5_3 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictOrd1_0.UnsafePtr)).V0, dictOrd_4)
_ = compare11_5_3
eq11_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_4, "Eq0"), gopurs_runtime.Value{}))
_ = eq11_6_4
eqIdentityT2_7_5 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_7 gopurs_runtime.Value, y_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq11_6_4, x_7, y_8)
}))
_ = eqIdentityT2_7_5
return gopurs_runtime.RecordGet(gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return eqIdentityT2_7_5
}), gopurs_runtime.Func2(func(x_8 gopurs_runtime.Value, y_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(compare11_5_3, x_8, y_9)
})), "compare")
}))
}

func Call_comonadIdentityT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
extendIdentityI1_3_2 := gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "extend"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, x_5)
}), v_4)
}))
_ = extendIdentityI1_3_2
return gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return extendIdentityI1_3_2
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictComonad_0.UnsafePtr)).V0, x_4)
}))
}

func Call_bindIdentityT(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
return dictBind_0
}

func Call_applyIdentityT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
return dictApply_0
}

func Call_applicativeIdentityT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return dictApplicative_0
}

func Call_alternativeIdentityT(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
return dictAlternative_0
}

func Call_altIdentityT(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
return dictAlt_0
}



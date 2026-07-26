package Control_Monad_Reader_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Function "gopurs/output/Data.Function"
)

var cache_ReaderT gopurs_runtime.Value
var once_ReaderT sync.Once
func Get_ReaderT() gopurs_runtime.Value {
	once_ReaderT.Do(func() {
		cache_ReaderT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ReaderT(x_0_box)
})
	})
	return cache_ReaderT
}

var cache_withReaderT gopurs_runtime.Value
var once_withReaderT sync.Once
func Get_withReaderT() gopurs_runtime.Value {
	once_withReaderT.Do(func() {
		cache_withReaderT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withReaderT(f_0_box, v_1_box)
})
	})
	return cache_withReaderT
}

var cache_runReaderT gopurs_runtime.Value
var once_runReaderT sync.Once
func Get_runReaderT() gopurs_runtime.Value {
	once_runReaderT.Do(func() {
		cache_runReaderT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runReaderT(v_0_box)
})
	})
	return cache_runReaderT
}

var cache_newtypeReaderT gopurs_runtime.Value
var once_newtypeReaderT sync.Once
func Get_newtypeReaderT() gopurs_runtime.Value {
	once_newtypeReaderT.Do(func() {
		cache_newtypeReaderT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeReaderT
}

var cache_monadTransReaderT gopurs_runtime.Value
var once_monadTransReaderT sync.Once
func Get_monadTransReaderT() gopurs_runtime.Value {
	once_monadTransReaderT.Do(func() {
		cache_monadTransReaderT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), Get_ReaderT(), pkg_Data_Function.Get_const_())
}))
	})
	return cache_monadTransReaderT
}

var cache_mapReaderT gopurs_runtime.Value
var once_mapReaderT sync.Once
func Get_mapReaderT() gopurs_runtime.Value {
	once_mapReaderT.Do(func() {
		cache_mapReaderT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapReaderT(f_0_box, v_1_box)
})
	})
	return cache_mapReaderT
}

var cache_functorReaderT gopurs_runtime.Value
var once_functorReaderT sync.Once
func Get_functorReaderT() gopurs_runtime.Value {
	once_functorReaderT.Do(func() {
		cache_functorReaderT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorReaderT(dictFunctor_0_box)
})
	})
	return cache_functorReaderT
}

var cache_distributiveReaderT gopurs_runtime.Value
var once_distributiveReaderT sync.Once
func Get_distributiveReaderT() gopurs_runtime.Value {
	once_distributiveReaderT.Do(func() {
		cache_distributiveReaderT = gopurs_runtime.Func(func(dictDistributive_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_distributiveReaderT(dictDistributive_0_box)
})
	})
	return cache_distributiveReaderT
}

var cache_applyReaderT gopurs_runtime.Value
var once_applyReaderT sync.Once
func Get_applyReaderT() gopurs_runtime.Value {
	once_applyReaderT.Do(func() {
		cache_applyReaderT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyReaderT(dictApply_0_box)
})
	})
	return cache_applyReaderT
}

var cache_bindReaderT gopurs_runtime.Value
var once_bindReaderT sync.Once
func Get_bindReaderT() gopurs_runtime.Value {
	once_bindReaderT.Do(func() {
		cache_bindReaderT = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindReaderT(dictBind_0_box)
})
	})
	return cache_bindReaderT
}

var cache_semigroupReaderT gopurs_runtime.Value
var once_semigroupReaderT sync.Once
func Get_semigroupReaderT() gopurs_runtime.Value {
	once_semigroupReaderT.Do(func() {
		cache_semigroupReaderT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupReaderT(dictApply_0_box)
})
	})
	return cache_semigroupReaderT
}

var cache_applicativeReaderT gopurs_runtime.Value
var once_applicativeReaderT sync.Once
func Get_applicativeReaderT() gopurs_runtime.Value {
	once_applicativeReaderT.Do(func() {
		cache_applicativeReaderT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeReaderT(dictApplicative_0_box)
})
	})
	return cache_applicativeReaderT
}

var cache_monadReaderT gopurs_runtime.Value
var once_monadReaderT sync.Once
func Get_monadReaderT() gopurs_runtime.Value {
	once_monadReaderT.Do(func() {
		cache_monadReaderT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadReaderT(dictMonad_0_box)
})
	})
	return cache_monadReaderT
}

var cache_monadAskReaderT gopurs_runtime.Value
var once_monadAskReaderT sync.Once
func Get_monadAskReaderT() gopurs_runtime.Value {
	once_monadAskReaderT.Do(func() {
		cache_monadAskReaderT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAskReaderT(dictMonad_0_box)
})
	})
	return cache_monadAskReaderT
}

var cache_monadReaderReaderT gopurs_runtime.Value
var once_monadReaderReaderT sync.Once
func Get_monadReaderReaderT() gopurs_runtime.Value {
	once_monadReaderReaderT.Do(func() {
		cache_monadReaderReaderT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadReaderReaderT(dictMonad_0_box)
})
	})
	return cache_monadReaderReaderT
}

var cache_monadContReaderT gopurs_runtime.Value
var once_monadContReaderT sync.Once
func Get_monadContReaderT() gopurs_runtime.Value {
	once_monadContReaderT.Do(func() {
		cache_monadContReaderT = gopurs_runtime.Func(func(dictMonadCont_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadContReaderT(dictMonadCont_0_box)
})
	})
	return cache_monadContReaderT
}

var cache_monadEffectReader gopurs_runtime.Value
var once_monadEffectReader sync.Once
func Get_monadEffectReader() gopurs_runtime.Value {
	once_monadEffectReader.Do(func() {
		cache_monadEffectReader = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadEffectReader(dictMonadEffect_0_box)
})
	})
	return cache_monadEffectReader
}

var cache_monadRecReaderT gopurs_runtime.Value
var once_monadRecReaderT sync.Once
func Get_monadRecReaderT() gopurs_runtime.Value {
	once_monadRecReaderT.Do(func() {
		cache_monadRecReaderT = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadRecReaderT(dictMonadRec_0_box)
})
	})
	return cache_monadRecReaderT
}

var cache_monadStateReaderT gopurs_runtime.Value
var once_monadStateReaderT sync.Once
func Get_monadStateReaderT() gopurs_runtime.Value {
	once_monadStateReaderT.Do(func() {
		cache_monadStateReaderT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadStateReaderT(dictMonadState_0_box)
})
	})
	return cache_monadStateReaderT
}

var cache_monadTellReaderT gopurs_runtime.Value
var once_monadTellReaderT sync.Once
func Get_monadTellReaderT() gopurs_runtime.Value {
	once_monadTellReaderT.Do(func() {
		cache_monadTellReaderT = gopurs_runtime.Func(func(dictMonadTell_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadTellReaderT(dictMonadTell_0_box)
})
	})
	return cache_monadTellReaderT
}

var cache_monadWriterReaderT gopurs_runtime.Value
var once_monadWriterReaderT sync.Once
func Get_monadWriterReaderT() gopurs_runtime.Value {
	once_monadWriterReaderT.Do(func() {
		cache_monadWriterReaderT = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadWriterReaderT(dictMonadWriter_0_box)
})
	})
	return cache_monadWriterReaderT
}

var cache_monadThrowReaderT gopurs_runtime.Value
var once_monadThrowReaderT sync.Once
func Get_monadThrowReaderT() gopurs_runtime.Value {
	once_monadThrowReaderT.Do(func() {
		cache_monadThrowReaderT = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadThrowReaderT(dictMonadThrow_0_box)
})
	})
	return cache_monadThrowReaderT
}

var cache_monadErrorReaderT gopurs_runtime.Value
var once_monadErrorReaderT sync.Once
func Get_monadErrorReaderT() gopurs_runtime.Value {
	once_monadErrorReaderT.Do(func() {
		cache_monadErrorReaderT = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadErrorReaderT(dictMonadError_0_box)
})
	})
	return cache_monadErrorReaderT
}

var cache_monadSTReaderT gopurs_runtime.Value
var once_monadSTReaderT sync.Once
func Get_monadSTReaderT() gopurs_runtime.Value {
	once_monadSTReaderT.Do(func() {
		cache_monadSTReaderT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadSTReaderT(dictMonadST_0_box)
})
	})
	return cache_monadSTReaderT
}

var cache_monoidReaderT gopurs_runtime.Value
var once_monoidReaderT sync.Once
func Get_monoidReaderT() gopurs_runtime.Value {
	once_monoidReaderT.Do(func() {
		cache_monoidReaderT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidReaderT(dictApplicative_0_box)
})
	})
	return cache_monoidReaderT
}

var cache_altReaderT gopurs_runtime.Value
var once_altReaderT sync.Once
func Get_altReaderT() gopurs_runtime.Value {
	once_altReaderT.Do(func() {
		cache_altReaderT = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altReaderT(dictAlt_0_box)
})
	})
	return cache_altReaderT
}

var cache_plusReaderT gopurs_runtime.Value
var once_plusReaderT sync.Once
func Get_plusReaderT() gopurs_runtime.Value {
	once_plusReaderT.Do(func() {
		cache_plusReaderT = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusReaderT(dictPlus_0_box)
})
	})
	return cache_plusReaderT
}

var cache_alternativeReaderT gopurs_runtime.Value
var once_alternativeReaderT sync.Once
func Get_alternativeReaderT() gopurs_runtime.Value {
	once_alternativeReaderT.Do(func() {
		cache_alternativeReaderT = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeReaderT(dictAlternative_0_box)
})
	})
	return cache_alternativeReaderT
}

var cache_monadPlusReaderT gopurs_runtime.Value
var once_monadPlusReaderT sync.Once
func Get_monadPlusReaderT() gopurs_runtime.Value {
	once_monadPlusReaderT.Do(func() {
		cache_monadPlusReaderT = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadPlusReaderT(dictMonadPlus_0_box)
})
	})
	return cache_monadPlusReaderT
}

func Call_ReaderT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_withReaderT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), v_1, f_0)
}

func Call_runReaderT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_mapReaderT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), f_0, v_1)
}

func Call_functorReaderT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), Get_mapReaderT(), ((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0))
}

func Call_distributiveReaderT(dictDistributive_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
distributiveReaderT:
for {
if false { continue distributiveReaderT }
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
functorReaderT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), Get_mapReaderT(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{}), "map")))
_ = functorReaderT1_1_0
return gopurs_runtime.RecordDict3("Functor0", "collect", "distribute", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_1_0
}), gopurs_runtime.Func2(func(dictFunctor_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_distributiveReaderT(), dictDistributive_0), "distribute"), dictFunctor_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_2, "map"), f_3))
}), gopurs_runtime.Func(func(dictFunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
collect1_3_1 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictDistributive_0.UnsafePtr)).V0, dictFunctor_2)
_ = collect1_3_1
return gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, e_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(collect1_3_1, gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(r_6, e_5)
}), a_4)
})
}))
}
}

func Call_applyReaderT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
functorReaderT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), Get_mapReaderT(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{}), "map")))
_ = functorReaderT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_1_0
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictApply_0.UnsafePtr)).V0, gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Apply(v1_3, r_4))
}))
}

func Call_bindReaderT(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
applyReaderT1_1_0 := gopurs_runtime.Apply(Get_applyReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = applyReaderT1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_1_0
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictBind_0.UnsafePtr)).V0, gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_3, a_5, r_4)
}))
}))
}

func Call_semigroupReaderT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
__local_var_1_0 := gopurs_runtime.Apply(Get_applyReaderT(), dictApply_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_3_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), __local_var_3_1, a_4), b_5)
}))
})
}

func Call_applicativeReaderT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
applyReaderT1_1_0 := gopurs_runtime.Apply(Get_applyReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = applyReaderT1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_1_0
}), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), Get_ReaderT(), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pkg_Data_Function.Get_const_(), ((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0)))
}

func Call_monadReaderT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applicativeReaderT1_1_0 := gopurs_runtime.Apply(Get_applicativeReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = applicativeReaderT1_1_0
bindReaderT1_2_1 := gopurs_runtime.Apply(Get_bindReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1_NOT_FOUND"), gopurs_runtime.Value{}))
_ = bindReaderT1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeReaderT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bindReaderT1_2_1
}))
}

func Call_monadAskReaderT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), dictMonad_0)
_ = monadReaderT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "pure"))
}

func Call_monadReaderReaderT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), dictMonad_0)
_ = monadReaderT1_1_0
monadAskReaderT1_2_1 := gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "pure"))
_ = monadAskReaderT1_2_1
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskReaderT1_2_1
}), Get_withReaderT())
}

func Call_monadContReaderT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = monadReaderT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictMonadCont_0.UnsafePtr)).V0, gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), Get_ReaderT(), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pkg_Data_Function.Get_const_(), c_4)), r_3)
}))
}))
}

func Call_monadEffectReader(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadReaderT1_2_1 := gopurs_runtime.Apply(Get_monadReaderT(), Monad0_1_0)
_ = monadReaderT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad0_1_0), ((*gopurs_runtime.RecordData1)(dictMonadEffect_0.UnsafePtr)).V0))
}

func Call_monadRecReaderT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
monadReaderT1_4_3 := gopurs_runtime.Apply(Get_monadReaderT(), Monad0_1_0)
_ = monadReaderT1_4_3
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_4_3
}), gopurs_runtime.Func3(func(k_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictMonadRec_0.UnsafePtr)).V0, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply2(k_5, a_prime_8, r_7), pure_3_2)
}), a_6)
}))
}

func Call_monadStateReaderT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadReaderT1_2_1 := gopurs_runtime.Apply(Get_monadReaderT(), Monad0_1_0)
_ = monadReaderT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad0_1_0), ((*gopurs_runtime.RecordData1)(dictMonadState_0.UnsafePtr)).V0))
}

func Call_monadTellReaderT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad1_1_0
Semigroup0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Semigroup0_2_1
monadReaderT1_3_2 := gopurs_runtime.Apply(Get_monadReaderT(), Monad1_1_0)
_ = monadReaderT1_3_2
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad1_1_0), ((*gopurs_runtime.RecordData1)(dictMonadTell_0.UnsafePtr)).V0))
}

func Call_monadWriterReaderT(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
Monoid0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monoid0_1_0
monadTellReaderT1_2_1 := gopurs_runtime.Apply(Get_monadTellReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1_NOT_FOUND"), gopurs_runtime.Value{}))
_ = monadTellReaderT1_2_1
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellReaderT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_1_0
}), gopurs_runtime.Apply(Get_mapReaderT(), ((*gopurs_runtime.RecordData2)(dictMonadWriter_0.UnsafePtr)).V0), gopurs_runtime.Apply(Get_mapReaderT(), ((*gopurs_runtime.RecordData2)(dictMonadWriter_0.UnsafePtr)).V1))
}

func Call_monadThrowReaderT(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadReaderT1_2_1 := gopurs_runtime.Apply(Get_monadReaderT(), Monad0_1_0)
_ = monadReaderT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad0_1_0), ((*gopurs_runtime.RecordData1)(dictMonadThrow_0.UnsafePtr)).V0))
}

func Call_monadErrorReaderT(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
monadThrowReaderT1_1_0 := gopurs_runtime.Apply(Get_monadThrowReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = monadThrowReaderT1_1_0
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowReaderT1_1_0
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, h_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictMonadError_0.UnsafePtr)).V0, gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(h_3, e_5, r_4)
}))
}))
}

func Call_monadSTReaderT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadReaderT1_2_1 := gopurs_runtime.Apply(Get_monadReaderT(), Monad0_1_0)
_ = monadReaderT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad0_1_0), ((*gopurs_runtime.RecordData1)(dictMonadST_0.UnsafePtr)).V0))
}

func Call_monoidReaderT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
semigroupReaderT1_1_0 := gopurs_runtime.Apply(Get_semigroupReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = semigroupReaderT1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupReaderT2_3_1 := gopurs_runtime.Apply(semigroupReaderT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupReaderT2_3_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupReaderT2_3_1
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_applicativeReaderT(), dictApplicative_0), "pure"), gopurs_runtime.RecordGet(dictMonoid_2, "mempty")))
})
}

func Call_altReaderT(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
functorReaderT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), Get_mapReaderT(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{}), "map")))
_ = functorReaderT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_1_0
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictAlt_0.UnsafePtr)).V0, gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Apply(v1_3, r_4))
}))
}

func Call_plusReaderT(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
altReaderT1_1_0 := gopurs_runtime.Apply(Get_altReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = altReaderT1_1_0
__local_var_2_1 := ((*gopurs_runtime.RecordData1)(dictPlus_0.UnsafePtr)).V0
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altReaderT1_1_0
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}))
}

func Call_alternativeReaderT(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
applicativeReaderT1_1_0 := gopurs_runtime.Apply(Get_applicativeReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = applicativeReaderT1_1_0
plusReaderT1_2_1 := gopurs_runtime.Apply(Get_plusReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1_NOT_FOUND"), gopurs_runtime.Value{}))
_ = plusReaderT1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeReaderT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusReaderT1_2_1
}))
}

func Call_monadPlusReaderT(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
_ = dictMonadPlus_0
monadReaderT1_1_0 := gopurs_runtime.Apply(Get_monadReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = monadReaderT1_1_0
alternativeReaderT1_2_1 := gopurs_runtime.Apply(Get_alternativeReaderT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Alternative1_NOT_FOUND"), gopurs_runtime.Value{}))
_ = alternativeReaderT1_2_1
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeReaderT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}))
}



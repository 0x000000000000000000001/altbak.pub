package Control_Monad_Reader_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
		cache_withReaderT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withReaderT(f_0_box, v_1_box, x_2_box)
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
		cache_monadTransReaderT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func3(func(dictMonad_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
	})
	return cache_monadTransReaderT
}

var cache_mapReaderT gopurs_runtime.Value
var once_mapReaderT sync.Once
func Get_mapReaderT() gopurs_runtime.Value {
	once_mapReaderT.Do(func() {
		cache_mapReaderT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapReaderT(f_0_box, v_1_box, x_2_box)
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

func Call_withReaderT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
}

func Call_runReaderT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_mapReaderT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_functorReaderT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), x_1)
_ = __local_var_2_0
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(v_3, x_4))
})
}))
}

func Call_distributiveReaderT(dictDistributive_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
distributiveReaderT:
for {
if false { continue distributiveReaderT }
var dictDistributive_0 gopurs_runtime.Value = dictDistributive_0_loop
_ = dictDistributive_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorReaderT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), x_2)
_ = __local_var_3_2
return gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
}))
_ = functorReaderT1_2_1
return gopurs_runtime.RecordDict3("Functor0", "collect", "distribute", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_1
}), gopurs_runtime.Func2(func(dictFunctor_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_distributiveReaderT(dictDistributive_0), "distribute"), dictFunctor_3)
_ = __local_var_5_3
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_3, "map"), f_4)
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(__local_var_6_4, x_7))
})
}), gopurs_runtime.Func(func(dictFunctor_3 gopurs_runtime.Value) gopurs_runtime.Value {
collect1_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictDistributive_0, "collect"), dictFunctor_3)
_ = collect1_4_5
return gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(collect1_4_5, gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(r_7, e_6)
}), a_5)
})
}))
}
}

func Call_applyReaderT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorReaderT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), x_2)
_ = __local_var_3_2
return gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
}))
_ = functorReaderT1_2_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_1
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
}))
}

func Call_bindReaderT(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorReaderT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
}))
_ = functorReaderT1_3_3
applyReaderT1_3_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_3
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
}))
_ = applyReaderT1_3_2
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_3_2
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_5, a_7, r_6)
}))
}))
}

func Call_semigroupReaderT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_3_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), __local_var_3_1)
_ = __local_var_6_2
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(a_4, r_7)), gopurs_runtime.Apply(b_5, r_7))
})
}))
})
}

func Call_applicativeReaderT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorReaderT1_3_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), x_3)
_ = __local_var_4_3
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(v_5, x_6))
})
}))
_ = functorReaderT1_3_2
applyReaderT1_4_4 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_2
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
}))
_ = applyReaderT1_4_4
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_4_4
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), x_5)
_ = __local_var_6_5
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_5
})
}))
}

func Call_monadReaderT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
functorReaderT1_4_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "map"), x_4)
_ = __local_var_5_5
return gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_5, gopurs_runtime.Apply(v_6, x_7))
})
}))
_ = functorReaderT1_4_4
applyReaderT1_5_6 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_4_4
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "apply"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
}))
_ = applyReaderT1_5_6
applicativeReaderT1_3_2 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_5_6
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), x_6)
_ = __local_var_7_7
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_7
})
}))
_ = applicativeReaderT1_3_2
bindReaderT1_4_8 := Call_bindReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = bindReaderT1_4_8
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeReaderT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bindReaderT1_4_8
}))
}

func Call_monadAskReaderT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadReaderT1_1_0 := Call_monadReaderT(dictMonad_0)
_ = monadReaderT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"))
}

func Call_monadReaderReaderT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadReaderT1_1_0 := Call_monadReaderT(dictMonad_0)
_ = monadReaderT1_1_0
monadAskReaderT1_2_1 := gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"))
_ = monadAskReaderT1_2_1
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskReaderT1_2_1
}), Get_withReaderT())
}

func Call_monadContReaderT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
monadReaderT1_1_0 := Call_monadReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0"), gopurs_runtime.Value{}))
_ = monadReaderT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := gopurs_runtime.Apply(c_4, x_5)
_ = __local_var_6_1
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_1
})
}), r_3)
}))
}))
}

func Call_monadEffectReader(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadReaderT1_2_1 := Call_monadReaderT(Monad0_1_0)
_ = monadReaderT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}))
}

func Call_monadRecReaderT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
monadReaderT1_4_3 := Call_monadReaderT(Monad0_1_0)
_ = monadReaderT1_4_3
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_4_3
}), gopurs_runtime.Func3(func(k_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply2(k_5, a_prime_8, r_7), pure_3_2)
}), a_6)
}))
}

func Call_monadStateReaderT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadReaderT1_2_1 := Call_monadReaderT(Monad0_1_0)
_ = monadReaderT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), x_4))
}))
}

func Call_monadTellReaderT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
Semigroup0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_2_1
monadReaderT1_3_2 := Call_monadReaderT(Monad1_1_0)
_ = monadReaderT1_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad1_1_0)
_ = __local_var_4_3
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_5))
}))
}

func Call_monadWriterReaderT(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
Monoid0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0"), gopurs_runtime.Value{})
_ = Monoid0_1_0
monadTellReaderT1_2_1 := Call_monadTellReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1"), gopurs_runtime.Value{}))
_ = monadTellReaderT1_2_1
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellReaderT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_1_0
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), gopurs_runtime.Apply(v_3, x_4))
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply(v_3, x_4))
}))
}

func Call_monadThrowReaderT(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadReaderT1_2_1 := Call_monadReaderT(Monad0_1_0)
_ = monadReaderT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), x_4))
}))
}

func Call_monadErrorReaderT(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
monadThrowReaderT1_1_0 := Call_monadThrowReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowReaderT1_1_0
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowReaderT1_1_0
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, h_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), gopurs_runtime.Apply(v_2, r_4), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(h_3, e_5, r_4)
}))
}))
}

func Call_monadSTReaderT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadReaderT1_2_1 := Call_monadReaderT(Monad0_1_0)
_ = monadReaderT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransReaderT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_monoidReaderT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_4_3
semigroupReaderT2_4_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), __local_var_4_3)
_ = __local_var_7_4
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Apply(a_5, r_8)), gopurs_runtime.Apply(b_6, r_8))
})
}))
_ = semigroupReaderT2_4_2
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"))
_ = __local_var_5_5
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupReaderT2_4_2
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_5_5
}))
})
}

func Call_altReaderT(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorReaderT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), x_2)
_ = __local_var_3_2
return gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(v_4, x_5))
})
}))
_ = functorReaderT1_2_1
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_1
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
}))
}

func Call_plusReaderT(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorReaderT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
}))
_ = functorReaderT1_3_3
altReaderT1_3_2 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_3
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "alt"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
}))
_ = altReaderT1_3_2
__local_var_4_5 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = __local_var_4_5
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return altReaderT1_3_2
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_4_5
}))
}

func Call_alternativeReaderT(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
functorReaderT1_4_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "map"), x_4)
_ = __local_var_5_5
return gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_5, gopurs_runtime.Apply(v_6, x_7))
})
}))
_ = functorReaderT1_4_4
applyReaderT1_5_6 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_4_4
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "apply"), gopurs_runtime.Apply(v_5, r_7), gopurs_runtime.Apply(v1_6, r_7))
}))
_ = applyReaderT1_5_6
applicativeReaderT1_3_2 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_5_6
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), x_6)
_ = __local_var_7_7
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_7
})
}))
_ = applicativeReaderT1_3_2
__local_var_4_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_4_8
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_8, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_5_9
__local_var_6_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_9, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_11
functorReaderT1_7_12 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_11, "map"), x_7)
_ = __local_var_8_13
return gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_13, gopurs_runtime.Apply(v_9, x_10))
})
}))
_ = functorReaderT1_7_12
altReaderT1_8_14 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_7_12
}), gopurs_runtime.Func3(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_9, "alt"), gopurs_runtime.Apply(v_8, r_10), gopurs_runtime.Apply(v1_9, r_10))
}))
_ = altReaderT1_8_14
__local_var_9_15 := gopurs_runtime.RecordGet(__local_var_4_8, "empty")
_ = __local_var_9_15
plusReaderT1_6_10 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return altReaderT1_8_14
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_9_15
}))
_ = plusReaderT1_6_10
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeReaderT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return plusReaderT1_6_10
}))
}

func Call_monadPlusReaderT(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
_ = dictMonadPlus_0
monadReaderT1_1_0 := Call_monadReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Monad0"), gopurs_runtime.Value{}))
_ = monadReaderT1_1_0
alternativeReaderT1_2_1 := Call_alternativeReaderT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Alternative1"), gopurs_runtime.Value{}))
_ = alternativeReaderT1_2_1
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeReaderT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadReaderT1_1_0
}))
}



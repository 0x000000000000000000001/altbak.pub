package Control_Monad_State_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	unsafe "unsafe"
)

var cache_StateT gopurs_runtime.Value
var once_StateT sync.Once
func Get_StateT() gopurs_runtime.Value {
	once_StateT.Do(func() {
		cache_StateT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_StateT(x_0_box)
})
	})
	return cache_StateT
}

var cache_withStateT gopurs_runtime.Value
var once_withStateT sync.Once
func Get_withStateT() gopurs_runtime.Value {
	once_withStateT.Do(func() {
		cache_withStateT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withStateT(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_withStateT
}

var cache_runStateT gopurs_runtime.Value
var once_runStateT sync.Once
func Get_runStateT() gopurs_runtime.Value {
	once_runStateT.Do(func() {
		cache_runStateT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runStateT(v_0_box)
})
	})
	return cache_runStateT
}

var cache_newtypeStateT gopurs_runtime.Value
var once_newtypeStateT sync.Once
func Get_newtypeStateT() gopurs_runtime.Value {
	once_newtypeStateT.Do(func() {
		cache_newtypeStateT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeStateT
}

var cache_monadTransStateT gopurs_runtime.Value
var once_monadTransStateT sync.Once
func Get_monadTransStateT() gopurs_runtime.Value {
	once_monadTransStateT.Do(func() {
		cache_monadTransStateT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func3(func(dictMonad_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), m_1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{x_3, s_2})})
}))
}))
	})
	return cache_monadTransStateT
}

var cache_mapStateT gopurs_runtime.Value
var once_mapStateT sync.Once
func Get_mapStateT() gopurs_runtime.Value {
	once_mapStateT.Do(func() {
		cache_mapStateT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapStateT(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_mapStateT
}

var cache_lazyStateT gopurs_runtime.Value
var once_lazyStateT sync.Once
func Get_lazyStateT() gopurs_runtime.Value {
	once_lazyStateT.Do(func() {
		cache_lazyStateT = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, pkg_Data_Unit.Get_unit(), s_1)
}))
	})
	return cache_lazyStateT
}

var cache_functorStateT gopurs_runtime.Value
var once_functorStateT sync.Once
func Get_functorStateT() gopurs_runtime.Value {
	once_functorStateT.Do(func() {
		cache_functorStateT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorStateT(dictFunctor_0_box)
})
	})
	return cache_functorStateT
}

var cache_execStateT gopurs_runtime.Value
var once_execStateT sync.Once
func Get_execStateT() gopurs_runtime.Value {
	once_execStateT.Do(func() {
		cache_execStateT = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_execStateT(dictFunctor_0_box, v_1_box, s_2_box)
})
	})
	return cache_execStateT
}

var cache_evalStateT gopurs_runtime.Value
var once_evalStateT sync.Once
func Get_evalStateT() gopurs_runtime.Value {
	once_evalStateT.Do(func() {
		cache_evalStateT = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_evalStateT(dictFunctor_0_box, v_1_box, s_2_box)
})
	})
	return cache_evalStateT
}

var cache_monadStateT gopurs_runtime.Value
var once_monadStateT sync.Once
func Get_monadStateT() gopurs_runtime.Value {
	once_monadStateT.Do(func() {
		cache_monadStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadStateT(dictMonad_0_box)
})
	})
	return cache_monadStateT
}

var cache_bindStateT gopurs_runtime.Value
var once_bindStateT sync.Once
func Get_bindStateT() gopurs_runtime.Value {
	once_bindStateT.Do(func() {
		cache_bindStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindStateT(dictMonad_0_box)
})
	})
	return cache_bindStateT
}

var cache_applyStateT gopurs_runtime.Value
var once_applyStateT sync.Once
func Get_applyStateT() gopurs_runtime.Value {
	once_applyStateT.Do(func() {
		cache_applyStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyStateT(dictMonad_0_box)
})
	})
	return cache_applyStateT
}

var cache_applicativeStateT gopurs_runtime.Value
var once_applicativeStateT sync.Once
func Get_applicativeStateT() gopurs_runtime.Value {
	once_applicativeStateT.Do(func() {
		cache_applicativeStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeStateT(dictMonad_0_box)
})
	})
	return cache_applicativeStateT
}

var cache_semigroupStateT gopurs_runtime.Value
var once_semigroupStateT sync.Once
func Get_semigroupStateT() gopurs_runtime.Value {
	once_semigroupStateT.Do(func() {
		cache_semigroupStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupStateT(dictMonad_0_box)
})
	})
	return cache_semigroupStateT
}

var cache_monadAskStateT gopurs_runtime.Value
var once_monadAskStateT sync.Once
func Get_monadAskStateT() gopurs_runtime.Value {
	once_monadAskStateT.Do(func() {
		cache_monadAskStateT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAskStateT(dictMonadAsk_0_box)
})
	})
	return cache_monadAskStateT
}

var cache_monadReaderStateT gopurs_runtime.Value
var once_monadReaderStateT sync.Once
func Get_monadReaderStateT() gopurs_runtime.Value {
	once_monadReaderStateT.Do(func() {
		cache_monadReaderStateT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadReaderStateT(dictMonadReader_0_box)
})
	})
	return cache_monadReaderStateT
}

var cache_monadContStateT gopurs_runtime.Value
var once_monadContStateT sync.Once
func Get_monadContStateT() gopurs_runtime.Value {
	once_monadContStateT.Do(func() {
		cache_monadContStateT = gopurs_runtime.Func(func(dictMonadCont_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadContStateT(dictMonadCont_0_box)
})
	})
	return cache_monadContStateT
}

var cache_monadEffectState gopurs_runtime.Value
var once_monadEffectState sync.Once
func Get_monadEffectState() gopurs_runtime.Value {
	once_monadEffectState.Do(func() {
		cache_monadEffectState = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadEffectState(dictMonadEffect_0_box)
})
	})
	return cache_monadEffectState
}

var cache_monadRecStateT gopurs_runtime.Value
var once_monadRecStateT sync.Once
func Get_monadRecStateT() gopurs_runtime.Value {
	once_monadRecStateT.Do(func() {
		cache_monadRecStateT = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadRecStateT(dictMonadRec_0_box)
})
	})
	return cache_monadRecStateT
}

var cache_monadStateStateT gopurs_runtime.Value
var once_monadStateStateT sync.Once
func Get_monadStateStateT() gopurs_runtime.Value {
	once_monadStateStateT.Do(func() {
		cache_monadStateStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadStateStateT(dictMonad_0_box)
})
	})
	return cache_monadStateStateT
}

var cache_monadTellStateT gopurs_runtime.Value
var once_monadTellStateT sync.Once
func Get_monadTellStateT() gopurs_runtime.Value {
	once_monadTellStateT.Do(func() {
		cache_monadTellStateT = gopurs_runtime.Func(func(dictMonadTell_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadTellStateT(dictMonadTell_0_box)
})
	})
	return cache_monadTellStateT
}

var cache_monadWriterStateT gopurs_runtime.Value
var once_monadWriterStateT sync.Once
func Get_monadWriterStateT() gopurs_runtime.Value {
	once_monadWriterStateT.Do(func() {
		cache_monadWriterStateT = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadWriterStateT(dictMonadWriter_0_box)
})
	})
	return cache_monadWriterStateT
}

var cache_monadThrowStateT gopurs_runtime.Value
var once_monadThrowStateT sync.Once
func Get_monadThrowStateT() gopurs_runtime.Value {
	once_monadThrowStateT.Do(func() {
		cache_monadThrowStateT = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadThrowStateT(dictMonadThrow_0_box)
})
	})
	return cache_monadThrowStateT
}

var cache_monadErrorStateT gopurs_runtime.Value
var once_monadErrorStateT sync.Once
func Get_monadErrorStateT() gopurs_runtime.Value {
	once_monadErrorStateT.Do(func() {
		cache_monadErrorStateT = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadErrorStateT(dictMonadError_0_box)
})
	})
	return cache_monadErrorStateT
}

var cache_monadSTStateT gopurs_runtime.Value
var once_monadSTStateT sync.Once
func Get_monadSTStateT() gopurs_runtime.Value {
	once_monadSTStateT.Do(func() {
		cache_monadSTStateT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadSTStateT(dictMonadST_0_box)
})
	})
	return cache_monadSTStateT
}

var cache_monoidStateT gopurs_runtime.Value
var once_monoidStateT sync.Once
func Get_monoidStateT() gopurs_runtime.Value {
	once_monoidStateT.Do(func() {
		cache_monoidStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidStateT(dictMonad_0_box)
})
	})
	return cache_monoidStateT
}

var cache_altStateT gopurs_runtime.Value
var once_altStateT sync.Once
func Get_altStateT() gopurs_runtime.Value {
	once_altStateT.Do(func() {
		cache_altStateT = gopurs_runtime.Func2(func(dictMonad_0_box gopurs_runtime.Value, dictAlt_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altStateT(dictMonad_0_box, dictAlt_1_box)
})
	})
	return cache_altStateT
}

var cache_plusStateT gopurs_runtime.Value
var once_plusStateT sync.Once
func Get_plusStateT() gopurs_runtime.Value {
	once_plusStateT.Do(func() {
		cache_plusStateT = gopurs_runtime.Func2(func(dictMonad_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusStateT(dictMonad_0_box, dictPlus_1_box)
})
	})
	return cache_plusStateT
}

var cache_alternativeStateT gopurs_runtime.Value
var once_alternativeStateT sync.Once
func Get_alternativeStateT() gopurs_runtime.Value {
	once_alternativeStateT.Do(func() {
		cache_alternativeStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeStateT(dictMonad_0_box)
})
	})
	return cache_alternativeStateT
}

var cache_monadPlusStateT gopurs_runtime.Value
var once_monadPlusStateT sync.Once
func Get_monadPlusStateT() gopurs_runtime.Value {
	once_monadPlusStateT.Do(func() {
		cache_monadPlusStateT = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadPlusStateT(dictMonadPlus_0_box)
})
	})
	return cache_monadPlusStateT
}

func Call_StateT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_withStateT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
}

func Call_runStateT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_mapStateT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_functorStateT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_2, s_3))
}))
}

func Call_execStateT(dictFunctor_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), pkg_Data_Tuple.Get_snd(), gopurs_runtime.Apply(v_1, s_2))
}

func Call_evalStateT(dictFunctor_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), pkg_Data_Tuple.Get_fst(), gopurs_runtime.Apply(v_1, s_2))
}

func Call_monadStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), dictMonad_0)
}))
}

func Call_bindStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyStateT(), dictMonad_0)
}), gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value, s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(v_1, s_3), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1)
}))
}))
}

func Call_applyStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorStateT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_3, s_4))
}))
_ = functorStateT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(Get_bindStateT(), dictMonad_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_2_1
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_applicativeStateT(), dictMonad_0), "pure"), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
}))
}

func Call_applicativeStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyStateT(), dictMonad_0)
}), gopurs_runtime.Func2(func(a_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_1, s_2})})
}))
}

func Call_semigroupStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(Get_applyStateT(), dictMonad_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_3_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), __local_var_3_1, a_4), b_5)
}))
})
}

func Call_monadAskStateT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadStateT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad0_1_0)
}))
_ = monadStateT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_monadTransStateT(), "lift"), Monad0_1_0, gopurs_runtime.RecordGet(dictMonadAsk_0, "ask")))
}

func Call_monadReaderStateT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
monadAskStateT1_1_0 := gopurs_runtime.Apply(Get_monadAskStateT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{}))
_ = monadAskStateT1_1_0
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskStateT1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), x_2)
_ = __local_var_3_1
return gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(v_4, x_5))
})
}))
}

func Call_monadContStateT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_0
monadStateT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), __local_var_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), __local_var_1_0)
}))
_ = monadStateT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_3, gopurs_runtime.Func2(func(a_6 gopurs_runtime.Value, s_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_6, s_prime_7})})
}), s_4)
}))
}))
}

func Call_monadEffectState(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadStateT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad0_1_0)
}))
_ = monadStateT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransStateT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}))
}

func Call_monadRecStateT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadStateT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad0_1_0)
}))
_ = monadStateT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1), gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 525585346) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V1})}})}
goto end_branch_2
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 60402430) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V1})}})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), __t2)
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_4, s_5})})
}))
}

func Call_monadStateStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadStateT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), dictMonad_0)
}))
_ = monadStateT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_1_0
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply(f_2, x_3))
}))
}

func Call_monadTellStateT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
Semigroup0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_2_1
monadStateT1_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad1_1_0)
}))
_ = monadStateT1_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransStateT(), "lift"), Monad1_1_0)
_ = __local_var_4_3
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_5))
}))
}

func Call_monadWriterStateT(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
MonadTell1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1"), gopurs_runtime.Value{})
_ = MonadTell1_1_0
Monad1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_4_3
Monoid0_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0"), gopurs_runtime.Value{})
_ = Monoid0_5_4
monadTellStateT1_6_5 := gopurs_runtime.Apply(Get_monadTellStateT(), MonadTell1_1_0)
_ = monadTellStateT1_6_5
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellStateT1_6_5
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_5_4
}), gopurs_runtime.Func2(func(m_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), gopurs_runtime.Apply(m_7, s_8)), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1})}, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V1})})
}))
}), gopurs_runtime.Func2(func(m_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply(m_7, s_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1})}, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V1})})
})))
}))
}

func Call_monadThrowStateT(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
lift1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransStateT(), "lift"), Monad0_1_0)
_ = lift1_2_1
monadStateT1_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad0_1_0)
}))
_ = monadStateT1_3_2
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_3_2
}), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(lift1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_4))
}))
}

func Call_monadErrorStateT(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
monadThrowStateT1_1_0 := gopurs_runtime.Apply(Get_monadThrowStateT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowStateT1_1_0
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowStateT1_1_0
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, h_3 gopurs_runtime.Value, s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), gopurs_runtime.Apply(v_2, s_4), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(h_3, e_5, s_4)
}))
}))
}

func Call_monadSTStateT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadStateT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad0_1_0)
}))
_ = monadStateT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransStateT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_monoidStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
semigroupStateT1_1_0 := gopurs_runtime.Apply(Get_semigroupStateT(), dictMonad_0)
_ = semigroupStateT1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupStateT2_3_1 := gopurs_runtime.Apply(semigroupStateT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupStateT2_3_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupStateT2_3_1
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_applicativeStateT(), dictMonad_0), "pure"), gopurs_runtime.RecordGet(dictMonoid_2, "mempty")))
})
}

func Call_altStateT(dictMonad_0_loop gopurs_runtime.Value, dictAlt_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var dictAlt_1 gopurs_runtime.Value = dictAlt_1_loop
_ = dictAlt_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_0
functorStateT1_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_4, s_5))
}))
_ = functorStateT1_3_1
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_3_1
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_1, "alt"), gopurs_runtime.Apply(v_4, s_6), gopurs_runtime.Apply(v1_5, s_6))
}))
}

func Call_plusStateT(dictMonad_0_loop gopurs_runtime.Value, dictPlus_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var dictPlus_1 gopurs_runtime.Value = dictPlus_1_loop
_ = dictPlus_1
empty_2_0 := gopurs_runtime.RecordGet(dictPlus_1, "empty")
_ = empty_2_0
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_1, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_2
functorStateT1_5_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_6, s_7))
}))
_ = functorStateT1_5_4
altStateT2_5_3 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_4
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "alt"), gopurs_runtime.Apply(v_6, s_8), gopurs_runtime.Apply(v1_7, s_8))
}))
_ = altStateT2_5_3
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return altStateT2_5_3
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_2_0
}))
}

func Call_alternativeStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applicativeStateT1_1_0 := gopurs_runtime.Apply(Get_applicativeStateT(), dictMonad_0)
_ = applicativeStateT1_1_0
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_3_1
empty_4_2 := gopurs_runtime.RecordGet(__local_var_3_1, "empty")
_ = empty_4_2
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_5_4
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_5
functorStateT1_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_5, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_7, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_8, s_9))
}))
_ = functorStateT1_7_6
altStateT2_8_7 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_7_6
}), gopurs_runtime.Func3(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "alt"), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Apply(v1_9, s_10))
}))
_ = altStateT2_8_7
plusStateT2_5_3 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return altStateT2_8_7
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_4_2
}))
_ = plusStateT2_5_3
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeStateT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return plusStateT2_5_3
}))
})
}

func Call_monadPlusStateT(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
_ = dictMonadPlus_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadStateT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad0_1_0)
}))
_ = monadStateT1_2_1
alternativeStateT1_3_2 := gopurs_runtime.Apply2(Get_alternativeStateT(), Monad0_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Alternative1"), gopurs_runtime.Value{}))
_ = alternativeStateT1_3_2
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeStateT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}))
}



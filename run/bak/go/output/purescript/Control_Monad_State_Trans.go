package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_State_Trans_StateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_StateT sync.Once
func Get_Control_Monad_State_Trans_StateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_StateT.Do(func() {
		cache_Control_Monad_State_Trans_StateT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_StateT(x_0_box)
})
	})
	return cache_Control_Monad_State_Trans_StateT
}

var cache_Control_Monad_State_Trans_withStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_withStateT sync.Once
func Get_Control_Monad_State_Trans_withStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_withStateT.Do(func() {
		cache_Control_Monad_State_Trans_withStateT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_withStateT(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Control_Monad_State_Trans_withStateT
}

var cache_Control_Monad_State_Trans_runStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_runStateT sync.Once
func Get_Control_Monad_State_Trans_runStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_runStateT.Do(func() {
		cache_Control_Monad_State_Trans_runStateT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_runStateT(v_0_box)
})
	})
	return cache_Control_Monad_State_Trans_runStateT
}

var cache_Control_Monad_State_Trans_newtypeStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_newtypeStateT sync.Once
func Get_Control_Monad_State_Trans_newtypeStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_newtypeStateT.Do(func() {
		cache_Control_Monad_State_Trans_newtypeStateT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Control_Monad_State_Trans_newtypeStateT
}

var cache_Control_Monad_State_Trans_monadTransStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadTransStateT sync.Once
func Get_Control_Monad_State_Trans_monadTransStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadTransStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadTransStateT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), m_3, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_5, s_4})})
}))
})
})
}))
	})
	return cache_Control_Monad_State_Trans_monadTransStateT
}

var cache_Control_Monad_State_Trans_lift gopurs_runtime.Value
var once_Control_Monad_State_Trans_lift sync.Once
func Get_Control_Monad_State_Trans_lift() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_lift.Do(func() {
		cache_Control_Monad_State_Trans_lift = gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift")
	})
	return cache_Control_Monad_State_Trans_lift
}

var cache_Control_Monad_State_Trans_mapStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_mapStateT sync.Once
func Get_Control_Monad_State_Trans_mapStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_mapStateT.Do(func() {
		cache_Control_Monad_State_Trans_mapStateT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_mapStateT(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Control_Monad_State_Trans_mapStateT
}

var cache_Control_Monad_State_Trans_lazyStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_lazyStateT sync.Once
func Get_Control_Monad_State_Trans_lazyStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_lazyStateT.Do(func() {
		cache_Control_Monad_State_Trans_lazyStateT = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, Get_Data_Unit_unit(), s_1)
})
}))
	})
	return cache_Control_Monad_State_Trans_lazyStateT
}

var cache_Control_Monad_State_Trans_functorStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_functorStateT sync.Once
func Get_Control_Monad_State_Trans_functorStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_functorStateT.Do(func() {
		cache_Control_Monad_State_Trans_functorStateT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_functorStateT(dictFunctor_0_box)
})
	})
	return cache_Control_Monad_State_Trans_functorStateT
}

var cache_Control_Monad_State_Trans_execStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_execStateT sync.Once
func Get_Control_Monad_State_Trans_execStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_execStateT.Do(func() {
		cache_Control_Monad_State_Trans_execStateT = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_execStateT(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), v_1_box, s_2_box)
})
	})
	return cache_Control_Monad_State_Trans_execStateT
}

var cache_Control_Monad_State_Trans_evalStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_evalStateT sync.Once
func Get_Control_Monad_State_Trans_evalStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_evalStateT.Do(func() {
		cache_Control_Monad_State_Trans_evalStateT = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_evalStateT(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), v_1_box, s_2_box)
})
	})
	return cache_Control_Monad_State_Trans_evalStateT
}

var cache_Control_Monad_State_Trans_monadStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadStateT sync.Once
func Get_Control_Monad_State_Trans_monadStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monadStateT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monadStateT
}

var cache_Control_Monad_State_Trans_bindStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_bindStateT sync.Once
func Get_Control_Monad_State_Trans_bindStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_bindStateT.Do(func() {
		cache_Control_Monad_State_Trans_bindStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_bindStateT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_State_Trans_bindStateT
}

var cache_Control_Monad_State_Trans_applyStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_applyStateT sync.Once
func Get_Control_Monad_State_Trans_applyStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_applyStateT.Do(func() {
		cache_Control_Monad_State_Trans_applyStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_State_Trans_applyStateT
}

var cache_Control_Monad_State_Trans_applicativeStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_applicativeStateT sync.Once
func Get_Control_Monad_State_Trans_applicativeStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_applicativeStateT.Do(func() {
		cache_Control_Monad_State_Trans_applicativeStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_State_Trans_applicativeStateT
}

var cache_Control_Monad_State_Trans_semigroupStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_semigroupStateT sync.Once
func Get_Control_Monad_State_Trans_semigroupStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_semigroupStateT.Do(func() {
		cache_Control_Monad_State_Trans_semigroupStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_semigroupStateT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_State_Trans_semigroupStateT
}

var cache_Control_Monad_State_Trans_monadAskStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadAskStateT sync.Once
func Get_Control_Monad_State_Trans_monadAskStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadAskStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadAskStateT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monadAskStateT(dictMonadAsk_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monadAskStateT
}

var cache_Control_Monad_State_Trans_monadReaderStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadReaderStateT sync.Once
func Get_Control_Monad_State_Trans_monadReaderStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadReaderStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadReaderStateT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monadReaderStateT(dictMonadReader_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monadReaderStateT
}

var cache_Control_Monad_State_Trans_monadContStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadContStateT sync.Once
func Get_Control_Monad_State_Trans_monadContStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadContStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadContStateT = gopurs_runtime.Func(func(dictMonadCont_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monadContStateT(dictMonadCont_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monadContStateT
}

var cache_Control_Monad_State_Trans_monadEffectState gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadEffectState sync.Once
func Get_Control_Monad_State_Trans_monadEffectState() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadEffectState.Do(func() {
		cache_Control_Monad_State_Trans_monadEffectState = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monadEffectState(dictMonadEffect_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monadEffectState
}

var cache_Control_Monad_State_Trans_monadRecStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadRecStateT sync.Once
func Get_Control_Monad_State_Trans_monadRecStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadRecStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadRecStateT = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monadRecStateT(dictMonadRec_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monadRecStateT
}

var cache_Control_Monad_State_Trans_monadStateStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadStateStateT sync.Once
func Get_Control_Monad_State_Trans_monadStateStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadStateStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadStateStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monadStateStateT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monadStateStateT
}

var cache_Control_Monad_State_Trans_monadTellStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadTellStateT sync.Once
func Get_Control_Monad_State_Trans_monadTellStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadTellStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadTellStateT = gopurs_runtime.Func(func(dictMonadTell_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monadTellStateT(dictMonadTell_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monadTellStateT
}

var cache_Control_Monad_State_Trans_monadWriterStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadWriterStateT sync.Once
func Get_Control_Monad_State_Trans_monadWriterStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadWriterStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadWriterStateT = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monadWriterStateT(dictMonadWriter_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monadWriterStateT
}

var cache_Control_Monad_State_Trans_monadThrowStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadThrowStateT sync.Once
func Get_Control_Monad_State_Trans_monadThrowStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadThrowStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadThrowStateT = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monadThrowStateT(dictMonadThrow_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monadThrowStateT
}

var cache_Control_Monad_State_Trans_monadErrorStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadErrorStateT sync.Once
func Get_Control_Monad_State_Trans_monadErrorStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadErrorStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadErrorStateT = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monadErrorStateT(dictMonadError_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monadErrorStateT
}

var cache_Control_Monad_State_Trans_monadSTStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadSTStateT sync.Once
func Get_Control_Monad_State_Trans_monadSTStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadSTStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadSTStateT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monadSTStateT(dictMonadST_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monadSTStateT
}

var cache_Control_Monad_State_Trans_monoidStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monoidStateT sync.Once
func Get_Control_Monad_State_Trans_monoidStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monoidStateT.Do(func() {
		cache_Control_Monad_State_Trans_monoidStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monoidStateT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monoidStateT
}

var cache_Control_Monad_State_Trans_altStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_altStateT sync.Once
func Get_Control_Monad_State_Trans_altStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_altStateT.Do(func() {
		cache_Control_Monad_State_Trans_altStateT = gopurs_runtime.Func2(func(dictMonad_0_box gopurs_runtime.Value, dictAlt_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_altStateT(dictMonad_0_box, dictAlt_1_box)
})
	})
	return cache_Control_Monad_State_Trans_altStateT
}

var cache_Control_Monad_State_Trans_plusStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_plusStateT sync.Once
func Get_Control_Monad_State_Trans_plusStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_plusStateT.Do(func() {
		cache_Control_Monad_State_Trans_plusStateT = gopurs_runtime.Func2(func(dictMonad_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_plusStateT(dictMonad_0_box, dictPlus_1_box)
})
	})
	return cache_Control_Monad_State_Trans_plusStateT
}

var cache_Control_Monad_State_Trans_alternativeStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_alternativeStateT sync.Once
func Get_Control_Monad_State_Trans_alternativeStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_alternativeStateT.Do(func() {
		cache_Control_Monad_State_Trans_alternativeStateT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_alternativeStateT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_State_Trans_alternativeStateT
}

var cache_Control_Monad_State_Trans_monadPlusStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadPlusStateT sync.Once
func Get_Control_Monad_State_Trans_monadPlusStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadPlusStateT.Do(func() {
		cache_Control_Monad_State_Trans_monadPlusStateT = gopurs_runtime.Func(func(dictMonadPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_monadPlusStateT(dictMonadPlus_0_box)
})
	})
	return cache_Control_Monad_State_Trans_monadPlusStateT
}

var cache_Control_Monad_State_Trans_mapStateT__2276753851 gopurs_runtime.Value
var once_Control_Monad_State_Trans_mapStateT__2276753851 sync.Once
func Get_Control_Monad_State_Trans_mapStateT__2276753851() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_mapStateT__2276753851.Do(func() {
		cache_Control_Monad_State_Trans_mapStateT__2276753851 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_mapStateT__2276753851(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Control_Monad_State_Trans_mapStateT__2276753851
}

var cache_Control_Monad_State_Trans_mapStateT__1987836370 gopurs_runtime.Value
var once_Control_Monad_State_Trans_mapStateT__1987836370 sync.Once
func Get_Control_Monad_State_Trans_mapStateT__1987836370() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_mapStateT__1987836370.Do(func() {
		cache_Control_Monad_State_Trans_mapStateT__1987836370 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_State_Trans_mapStateT__1987836370(f_0_box, v_1_box, x_2_box))}
})
	})
	return cache_Control_Monad_State_Trans_mapStateT__1987836370
}

var cache_Control_Monad_State_Trans_mapStateT__3766252210 gopurs_runtime.Value
var once_Control_Monad_State_Trans_mapStateT__3766252210 sync.Once
func Get_Control_Monad_State_Trans_mapStateT__3766252210() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_mapStateT__3766252210.Do(func() {
		cache_Control_Monad_State_Trans_mapStateT__3766252210 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_mapStateT__3766252210(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Control_Monad_State_Trans_mapStateT__3766252210
}

var cache_Control_Monad_State_Trans_monadTransStateT__2411581572 gopurs_runtime.Value
var once_Control_Monad_State_Trans_monadTransStateT__2411581572 sync.Once
func Get_Control_Monad_State_Trans_monadTransStateT__2411581572() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_monadTransStateT__2411581572.Do(func() {
		cache_Control_Monad_State_Trans_monadTransStateT__2411581572 = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), m_3, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_5, s_4})})
}))
})
})
}))
	})
	return cache_Control_Monad_State_Trans_monadTransStateT__2411581572
}

var cache_Control_Monad_State_Trans_withStateT__2250856667 gopurs_runtime.Value
var once_Control_Monad_State_Trans_withStateT__2250856667 sync.Once
func Get_Control_Monad_State_Trans_withStateT__2250856667() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_withStateT__2250856667.Do(func() {
		cache_Control_Monad_State_Trans_withStateT__2250856667 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_State_Trans_withStateT__2250856667(f_0_box, v_1_box, x_2_box))}
})
	})
	return cache_Control_Monad_State_Trans_withStateT__2250856667
}

func Call_Control_Monad_State_Trans_StateT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_State_Trans_withStateT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
}

func Call_Control_Monad_State_Trans_runStateT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Control_Monad_State_Trans_mapStateT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_Control_Monad_State_Trans_functorStateT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_2, s_3))
})
})
}))
}

func Call_Control_Monad_State_Trans_execStateT(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), Get_Data_Tuple_snd(), gopurs_runtime.Apply(v_1, s_2))
}

func Call_Control_Monad_State_Trans_evalStateT(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), Get_Data_Tuple_fst(), gopurs_runtime.Apply(v_1, s_2))
}

func Call_Control_Monad_State_Trans_monadStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_bindStateT(dictMonad_0)
}))
}

func Call_Control_Monad_State_Trans_bindStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(v_2, s_4), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)
}))
})
})
}))
}

func Call_Control_Monad_State_Trans_applyStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorStateT1_1_0 -> gopurs_runtime.Value
functorStateT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_3, s_4))
})
})
}))
_ = functorStateT1_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_bindStateT(dictMonad_0)
}))
_ = __local_var_2_2
// TAST (Let): Bind1_3_3 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_3
// TAST (Let): Applicative0_4_4 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_1_0
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_3.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_3.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_4.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}

func Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, s_3})})
})
}))
}

func Call_Control_Monad_State_Trans_semigroupStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): applyStateT1_1_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
applyStateT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_applyStateT(dictMonad_0))
_ = applyStateT1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyStateT1_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_2
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyStateT1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), __local_var_4_2, a_5), b_6)
})
}))
})
}

func Call_Control_Monad_State_Trans_monadAskStateT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): monadStateT1_1_0 -> gopurs_runtime.Value
monadStateT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_bindStateT(__local_var_1_1)
}))
_ = monadStateT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_1_0
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})))}, gopurs_runtime.RecordGet(dictMonadAsk_0, "ask")))
}

func Call_Control_Monad_State_Trans_monadReaderStateT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
// TAST (Let): monadAskStateT1_1_0 -> gopurs_runtime.Value
monadAskStateT1_1_0 := Call_Control_Monad_State_Trans_monadAskStateT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{}))
_ = monadAskStateT1_1_0
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskStateT1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), x_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(v_4, x_5))
})
})
}))
}

func Call_Control_Monad_State_Trans_monadContStateT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): monadStateT1_1_0 -> gopurs_runtime.Value
monadStateT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_bindStateT(__local_var_1_1)
}))
_ = monadStateT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_5, s_prime_6})})
})
}), s_3)
}))
})
}))
}

func Call_Control_Monad_State_Trans_monadEffectState(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadStateT1_2_1 -> gopurs_runtime.Value
monadStateT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_bindStateT(Monad0_1_0)
}))
_ = monadStateT1_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}))
}

func Call_Control_Monad_State_Trans_monadRecStateT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Applicative0_3_2 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_2
// TAST (Let): monadStateT1_4_3 -> gopurs_runtime.Value
monadStateT1_4_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_bindStateT(Monad0_1_0)
}))
_ = monadStateT1_4_3
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_4_3
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1), gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_9.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 525585346) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_9.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_9.UnsafePtr).V1})}})}
goto end_branch_6
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_9.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 60402430) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_9.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v2_9.UnsafePtr).V1})}})}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), __t6)
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_6, s_7})})
})
})
}))
}

func Call_Control_Monad_State_Trans_monadStateStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): monadStateT1_2_1 -> gopurs_runtime.Value
monadStateT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_bindStateT(dictMonad_0)
}))
_ = monadStateT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Apply(f_3, x_4))
})
}))
}

func Call_Control_Monad_State_Trans_monadTellStateT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
// TAST (Let): Monad1_1_0 -> gopurs_runtime.Value
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
// TAST (Let): Semigroup0_2_1 -> gopurs_runtime.Value
Semigroup0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_2_1
// TAST (Let): monadStateT1_3_2 -> gopurs_runtime.Value
monadStateT1_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_bindStateT(Monad1_1_0)
}))
_ = monadStateT1_3_2
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift"), Monad1_1_0)
_ = __local_var_4_3
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_5))
}))
}

func Call_Control_Monad_State_Trans_monadWriterStateT(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
// TAST (Let): MonadTell1_1_0 -> gopurs_runtime.Value
MonadTell1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1"), gopurs_runtime.Value{})
_ = MonadTell1_1_0
// TAST (Let): Monad1_2_1 -> gopurs_runtime.Value
Monad1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_2_1
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): Applicative0_4_3 -> gopurs_runtime.Value
Applicative0_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_4_3
// TAST (Let): Monoid0_5_4 -> gopurs_runtime.Value
Monoid0_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0"), gopurs_runtime.Value{})
_ = Monoid0_5_4
// TAST (Let): monadTellStateT1_6_5 -> gopurs_runtime.Value
monadTellStateT1_6_5 := Call_Control_Monad_State_Trans_monadTellStateT(MonadTell1_1_0)
_ = monadTellStateT1_6_5
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellStateT1_6_5
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_5_4
}), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), gopurs_runtime.Apply(m_7, s_8)), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_4_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1})}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V1})})
}))
})
}), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(m_7, s_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_4_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1})}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0.UnsafePtr).V1})})
})))
})
}))
}

func Call_Control_Monad_State_Trans_monadThrowStateT(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
// TAST (Let): Monad0_1_0 -> *Constructor_Control_Monad_Monad[gopurs_runtime.Value]
Monad0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): monadStateT1_2_1 -> gopurs_runtime.Value
monadStateT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_bindStateT(__local_var_2_2)
}))
_ = monadStateT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}), gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_1_0)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_3))
}))
}

func Call_Control_Monad_State_Trans_monadErrorStateT(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): monadThrowStateT1_1_0 -> gopurs_runtime.Value
monadThrowStateT1_1_0 := Call_Control_Monad_State_Trans_monadThrowStateT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowStateT1_1_0
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowStateT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), gopurs_runtime.Apply(v_2, s_4), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(h_3, e_5, s_4)
}))
})
})
}))
}

func Call_Control_Monad_State_Trans_monadSTStateT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadStateT1_2_1 -> gopurs_runtime.Value
monadStateT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_bindStateT(Monad0_1_0)
}))
_ = monadStateT1_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_State_Trans_monadTransStateT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_Control_Monad_State_Trans_monoidStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): applicativeStateT1_1_0 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
applicativeStateT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0))
_ = applicativeStateT1_1_0
// TAST (Let): semigroupStateT1_2_1 -> gopurs_runtime.Value
semigroupStateT1_2_1 := Call_Control_Monad_State_Trans_semigroupStateT(dictMonad_0)
_ = semigroupStateT1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupStateT2_4_2 -> gopurs_runtime.Value
semigroupStateT2_4_2 := gopurs_runtime.Apply(semigroupStateT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupStateT2_4_2
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupStateT2_4_2
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeStateT1_1_0.V1), gopurs_runtime.RecordGet(dictMonoid_3, "mempty")))
})
}

func Call_Control_Monad_State_Trans_altStateT(dictMonad_0_loop gopurs_runtime.Value, dictAlt_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var dictAlt_1 gopurs_runtime.Value = dictAlt_1_loop
_ = dictAlt_1
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): functorStateT1_2_0 -> gopurs_runtime.Value
functorStateT1_2_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_4, s_5))
})
})
}))
_ = functorStateT1_2_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_2_0
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_1, "alt"), gopurs_runtime.Apply(v_3, s_5), gopurs_runtime.Apply(v1_4, s_5))
})
})
}))
}

func Call_Control_Monad_State_Trans_plusStateT(dictMonad_0_loop gopurs_runtime.Value, dictPlus_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var dictPlus_1 gopurs_runtime.Value = dictPlus_1_loop
_ = dictPlus_1
// TAST (Let): empty_2_0 -> gopurs_runtime.Value
empty_2_0 := gopurs_runtime.RecordGet(dictPlus_1, "empty")
_ = empty_2_0
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_1, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_4
// TAST (Let): functorStateT1_4_3 -> gopurs_runtime.Value
functorStateT1_4_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_6, s_7))
})
})
}))
_ = functorStateT1_4_3
// TAST (Let): altStateT2_3_1 -> gopurs_runtime.Value
altStateT2_3_1 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_4_3
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "alt"), gopurs_runtime.Apply(v_5, s_7), gopurs_runtime.Apply(v1_6, s_7))
})
})
}))
_ = altStateT2_3_1
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return altStateT2_3_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_2_0
}))
}

func Call_Control_Monad_State_Trans_alternativeStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): applicativeStateT1_1_0 -> gopurs_runtime.Value
applicativeStateT1_1_0 := Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
_ = applicativeStateT1_1_0
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): empty_4_3 -> gopurs_runtime.Value
empty_4_3 := gopurs_runtime.RecordGet(__local_var_3_2, "empty")
_ = empty_4_3
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): functorStateT1_6_6 -> gopurs_runtime.Value
functorStateT1_6_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_7, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_8, s_9))
})
})
}))
_ = functorStateT1_6_6
// TAST (Let): altStateT2_5_4 -> gopurs_runtime.Value
altStateT2_5_4 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_6_6
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "alt"), gopurs_runtime.Apply(v_7, s_9), gopurs_runtime.Apply(v1_8, s_9))
})
})
}))
_ = altStateT2_5_4
// TAST (Let): plusStateT2_3_1 -> gopurs_runtime.Value
plusStateT2_3_1 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return altStateT2_5_4
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_4_3
}))
_ = plusStateT2_3_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeStateT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return plusStateT2_3_1
}))
})
}

func Call_Control_Monad_State_Trans_monadPlusStateT(dictMonadPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadPlus_0 gopurs_runtime.Value = dictMonadPlus_0_loop
_ = dictMonadPlus_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadStateT1_2_1 -> gopurs_runtime.Value
monadStateT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_bindStateT(Monad0_1_0)
}))
_ = monadStateT1_2_1
// TAST (Let): alternativeStateT1_3_2 -> gopurs_runtime.Value
alternativeStateT1_3_2 := gopurs_runtime.Apply(Call_Control_Monad_State_Trans_alternativeStateT(Monad0_1_0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Alternative1"), gopurs_runtime.Value{}))
_ = alternativeStateT1_3_2
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeStateT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}))
}

func Call_Control_Monad_State_Trans_mapStateT__2276753851(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_Control_Monad_State_Trans_mapStateT__1987836370(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2)))
}

func Call_Control_Monad_State_Trans_mapStateT__3766252210(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_Control_Monad_State_Trans_withStateT__2250856667(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2)))
}



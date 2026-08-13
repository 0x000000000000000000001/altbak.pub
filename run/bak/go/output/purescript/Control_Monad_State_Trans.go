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
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), m_3, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_5, s_4})})
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
return Call_Control_Monad_State_Trans_execStateT(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), v_1_box, s_2_box)
})
	})
	return cache_Control_Monad_State_Trans_execStateT
}

var cache_Control_Monad_State_Trans_evalStateT gopurs_runtime.Value
var once_Control_Monad_State_Trans_evalStateT sync.Once
func Get_Control_Monad_State_Trans_evalStateT() gopurs_runtime.Value {
	once_Control_Monad_State_Trans_evalStateT.Do(func() {
		cache_Control_Monad_State_Trans_evalStateT = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_evalStateT(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), v_1_box, s_2_box)
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
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), m_3, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_5, s_4})})
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple)(v1_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_4.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_2, s_3))
})
})
}))
}

func Call_Control_Monad_State_Trans_execStateT(dictFunctor_0_loop *Constructor_Data_Functor_Functor, v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), Get_Data_Tuple_snd(), gopurs_runtime.Apply(v_1, s_2))
}

func Call_Control_Monad_State_Trans_evalStateT(dictFunctor_0_loop *Constructor_Data_Functor_Functor, v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
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
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(v_2, s_4), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_3, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1)
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_3, s_4))
})
})
}))
_ = functorStateT1_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_3 -> *Constructor_Control_Bind_Bind
Bind1_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_3
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_3.V1), gopurs_runtime.Apply(v_4, s_6), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_2_2
// TAST (Let): Bind1_3_4 -> *Constructor_Control_Bind_Bind
Bind1_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_4
// TAST (Let): Applicative0_4_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_5
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_1_0
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_4.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_4.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_5.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
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
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): functorStateT1_3_1 -> gopurs_runtime.Value
functorStateT1_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_4, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_5, s_6))
})
})
}))
_ = functorStateT1_3_1
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_4 -> *Constructor_Control_Bind_Bind
Bind1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_4
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_4.V1), gopurs_runtime.Apply(v_6, s_8), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_7, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_4_3
// TAST (Let): Bind1_5_5 -> *Constructor_Control_Bind_Bind
Bind1_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_5
// TAST (Let): Applicative0_6_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_6
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_3_1
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_5.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_5.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_6.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_2, s_3})})
})
}))
}

func Call_Control_Monad_State_Trans_semigroupStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): functorStateT1_1_1 -> gopurs_runtime.Value
functorStateT1_1_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_3, s_4))
})
})
}))
_ = functorStateT1_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_3_4 -> gopurs_runtime.Value
pure_3_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_4
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_6
// TAST (Let): functorStateT1_5_5 -> gopurs_runtime.Value
functorStateT1_5_5 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_6, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_8 -> *Constructor_Control_Bind_Bind
Bind1_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_8
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_8.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_7
// TAST (Let): Bind1_7_9 -> *Constructor_Control_Bind_Bind
Bind1_7_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_9
// TAST (Let): Applicative0_8_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_10
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_5
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_9.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_9.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_10.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_4, s_5})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_11 -> *Constructor_Control_Bind_Bind
Bind1_3_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_11
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_11.V1), gopurs_runtime.Apply(v_4, s_6), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_2_3
// TAST (Let): Bind1_3_12 -> *Constructor_Control_Bind_Bind
Bind1_3_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_12
// TAST (Let): Applicative0_4_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_13
// TAST (Let): applyStateT1_1_0 -> *Constructor_Control_Apply_Apply
applyStateT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_1_1
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_12.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_12.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_13.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
})))
_ = applyStateT1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_14 -> *Constructor_Data_Functor_Functor
Functor0_3_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyStateT1_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_3_14
// TAST (Let): __local_var_4_15 -> gopurs_runtime.Value
__local_var_4_15 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_15
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyStateT1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_14.V0), __local_var_4_15, a_5), b_6)
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
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorStateT1_5_3 -> gopurs_runtime.Value
functorStateT1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_3
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_6 -> *Constructor_Control_Bind_Bind
Bind1_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_6
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_8
// TAST (Let): functorStateT1_9_7 -> gopurs_runtime.Value
functorStateT1_9_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_8, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_7
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_10 -> *Constructor_Control_Bind_Bind
Bind1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_10
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_10.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_9
// TAST (Let): Bind1_11_11 -> *Constructor_Control_Bind_Bind
Bind1_11_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_11
// TAST (Let): Applicative0_12_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_7
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_11.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_11.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_12.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_6.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_5
// TAST (Let): Bind1_7_13 -> *Constructor_Control_Bind_Bind
Bind1_7_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_13
// TAST (Let): Applicative0_8_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_3
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_14.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_4, s_5})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_15 -> *Constructor_Control_Bind_Bind
Bind1_3_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_15
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_17 -> gopurs_runtime.Value
__local_var_5_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_17
// TAST (Let): functorStateT1_5_16 -> gopurs_runtime.Value
functorStateT1_5_16 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_17, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_16
// TAST (Let): __local_var_6_18 -> gopurs_runtime.Value
__local_var_6_18 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_7_19 -> gopurs_runtime.Value
pure_7_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_19
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_21 -> gopurs_runtime.Value
__local_var_9_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_21
// TAST (Let): functorStateT1_9_20 -> gopurs_runtime.Value
functorStateT1_9_20 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_21, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_20
// TAST (Let): __local_var_10_22 -> gopurs_runtime.Value
__local_var_10_22 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_23 -> *Constructor_Control_Bind_Bind
Bind1_11_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_23
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_23.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_22
// TAST (Let): Bind1_11_24 -> *Constructor_Control_Bind_Bind
Bind1_11_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_22, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_24
// TAST (Let): Applicative0_12_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_22, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_20
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_24.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_24.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_25.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_19, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_8, s_9})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_26 -> *Constructor_Control_Bind_Bind
Bind1_7_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_26
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_26.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_18
// TAST (Let): Bind1_7_27 -> *Constructor_Control_Bind_Bind
Bind1_7_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_27
// TAST (Let): Applicative0_8_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_28
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_16
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_27.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_27.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_28.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_15.V1), gopurs_runtime.Apply(v_4, s_6), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1)
}))
})
})
}))
}))
_ = monadStateT1_1_0
// TAST (Let): __local_var_2_29 -> *Constructor_Control_Monad_Monad
__local_var_2_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_2_29
// TAST (Let): Bind1_3_30 -> *Constructor_Control_Bind_Bind
Bind1_3_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_29.V1), gopurs_runtime.Value{}))
_ = Bind1_3_30
// TAST (Let): pure_4_31 -> gopurs_runtime.Value
pure_4_31 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_29.V0), gopurs_runtime.Value{}), "pure")
_ = pure_4_31
// TAST (Let): __local_var_5_32 -> gopurs_runtime.Value
__local_var_5_32 := gopurs_runtime.RecordGet(dictMonadAsk_0, "ask")
_ = __local_var_5_32
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_1_0
}), gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_30.V1), __local_var_5_32, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_31, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_7, s_6})})
}))
}))
}

func Call_Control_Monad_State_Trans_monadReaderStateT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): monadStateT1_2_2 -> gopurs_runtime.Value
monadStateT1_2_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_4_4 -> gopurs_runtime.Value
pure_4_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_4
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_6
// TAST (Let): functorStateT1_6_5 -> gopurs_runtime.Value
functorStateT1_6_5 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_6, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_8, s_9))
})
})
}))
_ = functorStateT1_6_5
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_8_8 -> gopurs_runtime.Value
pure_8_8 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_8
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_10
// TAST (Let): functorStateT1_10_9 -> gopurs_runtime.Value
functorStateT1_10_9 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_10, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_12, s_13))
})
})
}))
_ = functorStateT1_10_9
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_12 -> *Constructor_Control_Bind_Bind
Bind1_12_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_12
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_14 -> gopurs_runtime.Value
__local_var_14_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_14
// TAST (Let): functorStateT1_14_13 -> gopurs_runtime.Value
functorStateT1_14_13 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_14, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_16, s_17))
})
})
}))
_ = functorStateT1_14_13
// TAST (Let): __local_var_15_15 -> gopurs_runtime.Value
__local_var_15_15 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_16 -> *Constructor_Control_Bind_Bind
Bind1_16_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_16
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_16.V1), gopurs_runtime.Apply(v_17, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_18, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_15_15
// TAST (Let): Bind1_16_17 -> *Constructor_Control_Bind_Bind
Bind1_16_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_17
// TAST (Let): Applicative0_17_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_18
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_14_13
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_17.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_17.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_18.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_12.V1), gopurs_runtime.Apply(v_13, s_15), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_14, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_11_11
// TAST (Let): Bind1_12_19 -> *Constructor_Control_Bind_Bind
Bind1_12_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_19
// TAST (Let): Applicative0_13_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_20
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_10_9
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_19.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_19.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_20.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_8_8, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, s_10})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_21 -> *Constructor_Control_Bind_Bind
Bind1_8_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_21
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_23 -> gopurs_runtime.Value
__local_var_10_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_23
// TAST (Let): functorStateT1_10_22 -> gopurs_runtime.Value
functorStateT1_10_22 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_23, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_12, s_13))
})
})
}))
_ = functorStateT1_10_22
// TAST (Let): __local_var_11_24 -> gopurs_runtime.Value
__local_var_11_24 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_12_25 -> gopurs_runtime.Value
pure_12_25 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_25
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_27 -> gopurs_runtime.Value
__local_var_14_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_27
// TAST (Let): functorStateT1_14_26 -> gopurs_runtime.Value
functorStateT1_14_26 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_27, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_16, s_17))
})
})
}))
_ = functorStateT1_14_26
// TAST (Let): __local_var_15_28 -> gopurs_runtime.Value
__local_var_15_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_29 -> *Constructor_Control_Bind_Bind
Bind1_16_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_29
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_31 -> gopurs_runtime.Value
__local_var_18_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_31
// TAST (Let): functorStateT1_18_30 -> gopurs_runtime.Value
functorStateT1_18_30 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_31, "map"), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Tuple_Tuple)(v1_22.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_22.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_20, s_21))
})
})
}))
_ = functorStateT1_18_30
// TAST (Let): __local_var_19_32 -> gopurs_runtime.Value
__local_var_19_32 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_33 -> *Constructor_Control_Bind_Bind
Bind1_20_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_33
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_33.V1), gopurs_runtime.Apply(v_21, s_23), gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_22, (*Constructor_Data_Tuple_Tuple)(v1_24.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_24.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_19_32
// TAST (Let): Bind1_20_34 -> *Constructor_Control_Bind_Bind
Bind1_20_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_32, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_34
// TAST (Let): Applicative0_21_35 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_32, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_35
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_18_30
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_34.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_34.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_35.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_29.V1), gopurs_runtime.Apply(v_17, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_18, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_15_28
// TAST (Let): Bind1_16_36 -> *Constructor_Control_Bind_Bind
Bind1_16_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_36
// TAST (Let): Applicative0_17_37 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_37
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_14_26
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_36.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_36.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_37.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_12_25, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_13, s_14})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_38 -> *Constructor_Control_Bind_Bind
Bind1_12_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_38
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_40 -> gopurs_runtime.Value
__local_var_14_40 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_40
// TAST (Let): functorStateT1_14_39 -> gopurs_runtime.Value
functorStateT1_14_39 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_40, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_16, s_17))
})
})
}))
_ = functorStateT1_14_39
// TAST (Let): __local_var_15_41 -> gopurs_runtime.Value
__local_var_15_41 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_16_42 -> gopurs_runtime.Value
pure_16_42 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_42
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_44 -> gopurs_runtime.Value
__local_var_18_44 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_44
// TAST (Let): functorStateT1_18_43 -> gopurs_runtime.Value
functorStateT1_18_43 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_44, "map"), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Tuple_Tuple)(v1_22.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_22.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_20, s_21))
})
})
}))
_ = functorStateT1_18_43
// TAST (Let): __local_var_19_45 -> gopurs_runtime.Value
__local_var_19_45 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_46 -> *Constructor_Control_Bind_Bind
Bind1_20_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_46
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_46.V1), gopurs_runtime.Apply(v_21, s_23), gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_22, (*Constructor_Data_Tuple_Tuple)(v1_24.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_24.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_19_45
// TAST (Let): Bind1_20_47 -> *Constructor_Control_Bind_Bind
Bind1_20_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_45, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_47
// TAST (Let): Applicative0_21_48 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_45, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_48
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_18_43
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_47.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_47.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_48.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_16_42, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_17, s_18})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_49 -> *Constructor_Control_Bind_Bind
Bind1_16_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_49
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_49.V1), gopurs_runtime.Apply(v_17, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_18, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_15_41
// TAST (Let): Bind1_16_50 -> *Constructor_Control_Bind_Bind
Bind1_16_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_41, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_50
// TAST (Let): Applicative0_17_51 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_41, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_51
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_14_39
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_50.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_50.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_51.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_38.V1), gopurs_runtime.Apply(v_13, s_15), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_14, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_11_24
// TAST (Let): Bind1_12_52 -> *Constructor_Control_Bind_Bind
Bind1_12_52 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_24, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_52
// TAST (Let): Applicative0_13_53 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_24, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_53
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_10_22
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_52.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_52.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_53.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_21.V1), gopurs_runtime.Apply(v_9, s_11), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_10, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_7_7
// TAST (Let): Bind1_8_54 -> *Constructor_Control_Bind_Bind
Bind1_8_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_54
// TAST (Let): Applicative0_9_55 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_55
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_6_5
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_54.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_54.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_55.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, s_6})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_56 -> *Constructor_Control_Bind_Bind
Bind1_4_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_56
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_58 -> gopurs_runtime.Value
__local_var_6_58 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_58
// TAST (Let): functorStateT1_6_57 -> gopurs_runtime.Value
functorStateT1_6_57 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_58, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_8, s_9))
})
})
}))
_ = functorStateT1_6_57
// TAST (Let): __local_var_7_59 -> gopurs_runtime.Value
__local_var_7_59 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_8_60 -> gopurs_runtime.Value
pure_8_60 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_60
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_62 -> gopurs_runtime.Value
__local_var_10_62 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_62
// TAST (Let): functorStateT1_10_61 -> gopurs_runtime.Value
functorStateT1_10_61 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_62, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_12, s_13))
})
})
}))
_ = functorStateT1_10_61
// TAST (Let): __local_var_11_63 -> gopurs_runtime.Value
__local_var_11_63 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_12_64 -> gopurs_runtime.Value
pure_12_64 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_64
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_66 -> gopurs_runtime.Value
__local_var_14_66 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_66
// TAST (Let): functorStateT1_14_65 -> gopurs_runtime.Value
functorStateT1_14_65 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_66, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_16, s_17))
})
})
}))
_ = functorStateT1_14_65
// TAST (Let): __local_var_15_67 -> gopurs_runtime.Value
__local_var_15_67 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_68 -> *Constructor_Control_Bind_Bind
Bind1_16_68 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_68
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_70 -> gopurs_runtime.Value
__local_var_18_70 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_70
// TAST (Let): functorStateT1_18_69 -> gopurs_runtime.Value
functorStateT1_18_69 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_70, "map"), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Tuple_Tuple)(v1_22.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_22.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_20, s_21))
})
})
}))
_ = functorStateT1_18_69
// TAST (Let): __local_var_19_71 -> gopurs_runtime.Value
__local_var_19_71 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_72 -> *Constructor_Control_Bind_Bind
Bind1_20_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_72
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_72.V1), gopurs_runtime.Apply(v_21, s_23), gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_22, (*Constructor_Data_Tuple_Tuple)(v1_24.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_24.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_19_71
// TAST (Let): Bind1_20_73 -> *Constructor_Control_Bind_Bind
Bind1_20_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_71, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_73
// TAST (Let): Applicative0_21_74 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_71, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_74
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_18_69
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_73.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_73.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_74.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_68.V1), gopurs_runtime.Apply(v_17, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_18, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_15_67
// TAST (Let): Bind1_16_75 -> *Constructor_Control_Bind_Bind
Bind1_16_75 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_67, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_75
// TAST (Let): Applicative0_17_76 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_76 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_67, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_76
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_14_65
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_75.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_75.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_76.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_12_64, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_13, s_14})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_77 -> *Constructor_Control_Bind_Bind
Bind1_12_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_77
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_79 -> gopurs_runtime.Value
__local_var_14_79 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_79
// TAST (Let): functorStateT1_14_78 -> gopurs_runtime.Value
functorStateT1_14_78 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_79, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_16, s_17))
})
})
}))
_ = functorStateT1_14_78
// TAST (Let): __local_var_15_80 -> gopurs_runtime.Value
__local_var_15_80 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_16_81 -> gopurs_runtime.Value
pure_16_81 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_81
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_83 -> gopurs_runtime.Value
__local_var_18_83 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_83
// TAST (Let): functorStateT1_18_82 -> gopurs_runtime.Value
functorStateT1_18_82 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_83, "map"), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Tuple_Tuple)(v1_22.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_22.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_20, s_21))
})
})
}))
_ = functorStateT1_18_82
// TAST (Let): __local_var_19_84 -> gopurs_runtime.Value
__local_var_19_84 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_85 -> *Constructor_Control_Bind_Bind
Bind1_20_85 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_85
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_21 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_85.V1), gopurs_runtime.Apply(v_21, s_23), gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_22, (*Constructor_Data_Tuple_Tuple)(v1_24.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_24.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_19_84
// TAST (Let): Bind1_20_86 -> *Constructor_Control_Bind_Bind
Bind1_20_86 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_84, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_86
// TAST (Let): Applicative0_21_87 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_87 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_84, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_87
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_18_82
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_86.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_86.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_87.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_16_81, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_17, s_18})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_88 -> *Constructor_Control_Bind_Bind
Bind1_16_88 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_88
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_88.V1), gopurs_runtime.Apply(v_17, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_18, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_15_80
// TAST (Let): Bind1_16_89 -> *Constructor_Control_Bind_Bind
Bind1_16_89 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_80, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_89
// TAST (Let): Applicative0_17_90 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_90 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_80, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_90
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_14_78
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_89.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_89.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_90.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_77.V1), gopurs_runtime.Apply(v_13, s_15), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_14, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_11_63
// TAST (Let): Bind1_12_91 -> *Constructor_Control_Bind_Bind
Bind1_12_91 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_63, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_91
// TAST (Let): Applicative0_13_92 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_92 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_63, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_92
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_10_61
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_91.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_91.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_92.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_8_60, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, s_10})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_93 -> *Constructor_Control_Bind_Bind
Bind1_8_93 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_93
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_95 -> gopurs_runtime.Value
__local_var_10_95 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_95
// TAST (Let): functorStateT1_10_94 -> gopurs_runtime.Value
functorStateT1_10_94 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_95, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_12, s_13))
})
})
}))
_ = functorStateT1_10_94
// TAST (Let): __local_var_11_96 -> gopurs_runtime.Value
__local_var_11_96 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_12_97 -> gopurs_runtime.Value
pure_12_97 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_97
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_99 -> gopurs_runtime.Value
__local_var_14_99 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_99
// TAST (Let): functorStateT1_14_98 -> gopurs_runtime.Value
functorStateT1_14_98 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_99, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_16, s_17))
})
})
}))
_ = functorStateT1_14_98
// TAST (Let): __local_var_15_100 -> gopurs_runtime.Value
__local_var_15_100 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_101 -> *Constructor_Control_Bind_Bind
Bind1_16_101 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_101
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_101.V1), gopurs_runtime.Apply(v_17, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_18, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_15_100
// TAST (Let): Bind1_16_102 -> *Constructor_Control_Bind_Bind
Bind1_16_102 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_100, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_102
// TAST (Let): Applicative0_17_103 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_103 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_100, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_103
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_14_98
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_102.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_102.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_103.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_12_97, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_13, s_14})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_104 -> *Constructor_Control_Bind_Bind
Bind1_12_104 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_104
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_104.V1), gopurs_runtime.Apply(v_13, s_15), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_14, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_11_96
// TAST (Let): Bind1_12_105 -> *Constructor_Control_Bind_Bind
Bind1_12_105 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_96, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_105
// TAST (Let): Applicative0_13_106 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_106 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_96, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_106
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_10_94
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_105.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_105.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_106.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_93.V1), gopurs_runtime.Apply(v_9, s_11), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_10, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_7_59
// TAST (Let): Bind1_8_107 -> *Constructor_Control_Bind_Bind
Bind1_8_107 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_59, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_107
// TAST (Let): Applicative0_9_108 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_108 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_59, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_108
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_6_57
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_107.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_107.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_108.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_56.V1), gopurs_runtime.Apply(v_5, s_7), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_6, (*Constructor_Data_Tuple_Tuple)(v1_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_8.UnsafePtr).V1)
}))
})
})
}))
}))
_ = monadStateT1_2_2
// TAST (Let): __local_var_3_109 -> *Constructor_Control_Monad_Monad
__local_var_3_109 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_109
// TAST (Let): Bind1_4_110 -> *Constructor_Control_Bind_Bind
Bind1_4_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_109.V1), gopurs_runtime.Value{}))
_ = Bind1_4_110
// TAST (Let): pure_5_111 -> gopurs_runtime.Value
pure_5_111 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_109.V0), gopurs_runtime.Value{}), "pure")
_ = pure_5_111
// TAST (Let): __local_var_6_112 -> gopurs_runtime.Value
__local_var_6_112 := gopurs_runtime.RecordGet(__local_var_1_1, "ask")
_ = __local_var_6_112
// TAST (Let): monadAskStateT1_1_0 -> gopurs_runtime.Value
monadAskStateT1_1_0 := gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_2
}), gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_110.V1), __local_var_6_112, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_111, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_8, s_7})})
}))
}))
_ = monadAskStateT1_1_0
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskStateT1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_113 -> gopurs_runtime.Value
__local_var_3_113 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), x_2)
_ = __local_var_3_113
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_113, gopurs_runtime.Apply(v_4, x_5))
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
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorStateT1_5_3 -> gopurs_runtime.Value
functorStateT1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_3
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_6 -> *Constructor_Control_Bind_Bind
Bind1_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_6
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_8
// TAST (Let): functorStateT1_9_7 -> gopurs_runtime.Value
functorStateT1_9_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_8, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_7
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_10 -> *Constructor_Control_Bind_Bind
Bind1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_10
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_10.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_9
// TAST (Let): Bind1_11_11 -> *Constructor_Control_Bind_Bind
Bind1_11_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_11
// TAST (Let): Applicative0_12_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_7
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_11.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_11.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_12.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_6.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_5
// TAST (Let): Bind1_7_13 -> *Constructor_Control_Bind_Bind
Bind1_7_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_13
// TAST (Let): Applicative0_8_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_3
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_14.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_4, s_5})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_15 -> *Constructor_Control_Bind_Bind
Bind1_3_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_15
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_17 -> gopurs_runtime.Value
__local_var_5_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_17
// TAST (Let): functorStateT1_5_16 -> gopurs_runtime.Value
functorStateT1_5_16 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_17, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_16
// TAST (Let): __local_var_6_18 -> gopurs_runtime.Value
__local_var_6_18 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_7_19 -> gopurs_runtime.Value
pure_7_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_19
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_21 -> gopurs_runtime.Value
__local_var_9_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_21
// TAST (Let): functorStateT1_9_20 -> gopurs_runtime.Value
functorStateT1_9_20 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_21, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_20
// TAST (Let): __local_var_10_22 -> gopurs_runtime.Value
__local_var_10_22 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_23 -> *Constructor_Control_Bind_Bind
Bind1_11_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_23
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_23.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_22
// TAST (Let): Bind1_11_24 -> *Constructor_Control_Bind_Bind
Bind1_11_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_22, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_24
// TAST (Let): Applicative0_12_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_22, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_20
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_24.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_24.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_25.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_19, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_8, s_9})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_26 -> *Constructor_Control_Bind_Bind
Bind1_7_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_26
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_26.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_18
// TAST (Let): Bind1_7_27 -> *Constructor_Control_Bind_Bind
Bind1_7_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_27
// TAST (Let): Applicative0_8_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_28
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_16
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_27.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_27.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_28.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_15.V1), gopurs_runtime.Apply(v_4, s_6), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1)
}))
})
})
}))
}))
_ = monadStateT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, s_prime_6})})
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
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorStateT1_5_3 -> gopurs_runtime.Value
functorStateT1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_3
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_6 -> *Constructor_Control_Bind_Bind
Bind1_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_6
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_8
// TAST (Let): functorStateT1_9_7 -> gopurs_runtime.Value
functorStateT1_9_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_8, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_7
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_10 -> *Constructor_Control_Bind_Bind
Bind1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_10
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_10.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_9
// TAST (Let): Bind1_11_11 -> *Constructor_Control_Bind_Bind
Bind1_11_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_11
// TAST (Let): Applicative0_12_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_7
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_11.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_11.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_12.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_6.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_5
// TAST (Let): Bind1_7_13 -> *Constructor_Control_Bind_Bind
Bind1_7_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_13
// TAST (Let): Applicative0_8_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_3
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_14.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_4, s_5})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_15 -> *Constructor_Control_Bind_Bind
Bind1_3_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_15
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_17 -> gopurs_runtime.Value
__local_var_5_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_17
// TAST (Let): functorStateT1_5_16 -> gopurs_runtime.Value
functorStateT1_5_16 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_17, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_16
// TAST (Let): __local_var_6_18 -> gopurs_runtime.Value
__local_var_6_18 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_7_19 -> gopurs_runtime.Value
pure_7_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_19
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_21 -> gopurs_runtime.Value
__local_var_9_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_21
// TAST (Let): functorStateT1_9_20 -> gopurs_runtime.Value
functorStateT1_9_20 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_21, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_20
// TAST (Let): __local_var_10_22 -> gopurs_runtime.Value
__local_var_10_22 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_23 -> *Constructor_Control_Bind_Bind
Bind1_11_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_23
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_23.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_22
// TAST (Let): Bind1_11_24 -> *Constructor_Control_Bind_Bind
Bind1_11_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_22, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_24
// TAST (Let): Applicative0_12_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_22, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_20
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_24.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_24.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_25.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_19, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_8, s_9})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_26 -> *Constructor_Control_Bind_Bind
Bind1_7_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_26
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_26.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_18
// TAST (Let): Bind1_7_27 -> *Constructor_Control_Bind_Bind
Bind1_7_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_27
// TAST (Let): Applicative0_8_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_28
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_16
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_27.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_27.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_28.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_15.V1), gopurs_runtime.Apply(v_4, s_6), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1)
}))
})
})
}))
}))
_ = monadStateT1_2_1
// TAST (Let): Bind1_3_30 -> *Constructor_Control_Bind_Bind
Bind1_3_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_30
// TAST (Let): pure_4_31 -> gopurs_runtime.Value
pure_4_31 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_31
// TAST (Let): __local_var_3_29 -> gopurs_runtime.Value
__local_var_3_29 := gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_30.V1), m_5, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_31, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_7, s_6})})
}))
})
})
_ = __local_var_3_29
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_29, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}))
}

func Call_Control_Monad_State_Trans_monadRecStateT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Applicative0_3_2 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_2
// TAST (Let): monadStateT1_4_3 -> gopurs_runtime.Value
monadStateT1_4_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_5_4 -> gopurs_runtime.Value
pure_5_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_4
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_6
// TAST (Let): functorStateT1_7_5 -> gopurs_runtime.Value
functorStateT1_7_5 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_6, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_9, s_10))
})
})
}))
_ = functorStateT1_7_5
// TAST (Let): __local_var_8_7 -> gopurs_runtime.Value
__local_var_8_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_8 -> *Constructor_Control_Bind_Bind
Bind1_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_8
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_10 -> gopurs_runtime.Value
__local_var_11_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_10
// TAST (Let): functorStateT1_11_9 -> gopurs_runtime.Value
functorStateT1_11_9 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_10, "map"), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_13, s_14))
})
})
}))
_ = functorStateT1_11_9
// TAST (Let): __local_var_12_11 -> gopurs_runtime.Value
__local_var_12_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_12 -> *Constructor_Control_Bind_Bind
Bind1_13_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_12
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_12.V1), gopurs_runtime.Apply(v_14, s_16), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_15, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_12_11
// TAST (Let): Bind1_13_13 -> *Constructor_Control_Bind_Bind
Bind1_13_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_13
// TAST (Let): Applicative0_14_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_11_9
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_13.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_13.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_14.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_8.V1), gopurs_runtime.Apply(v_10, s_12), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_11, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_8_7
// TAST (Let): Bind1_9_15 -> *Constructor_Control_Bind_Bind
Bind1_9_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_15
// TAST (Let): Applicative0_10_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_16
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_7_5
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_15.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_15.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_16.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_6, s_7})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_17 -> *Constructor_Control_Bind_Bind
Bind1_5_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_17
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_19 -> gopurs_runtime.Value
__local_var_7_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_19
// TAST (Let): functorStateT1_7_18 -> gopurs_runtime.Value
functorStateT1_7_18 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_19, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_9, s_10))
})
})
}))
_ = functorStateT1_7_18
// TAST (Let): __local_var_8_20 -> gopurs_runtime.Value
__local_var_8_20 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_9_21 -> gopurs_runtime.Value
pure_9_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_21
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_23 -> gopurs_runtime.Value
__local_var_11_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_23
// TAST (Let): functorStateT1_11_22 -> gopurs_runtime.Value
functorStateT1_11_22 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_23, "map"), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_13, s_14))
})
})
}))
_ = functorStateT1_11_22
// TAST (Let): __local_var_12_24 -> gopurs_runtime.Value
__local_var_12_24 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_25 -> *Constructor_Control_Bind_Bind
Bind1_13_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_25
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_25.V1), gopurs_runtime.Apply(v_14, s_16), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_15, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_12_24
// TAST (Let): Bind1_13_26 -> *Constructor_Control_Bind_Bind
Bind1_13_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_24, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_26
// TAST (Let): Applicative0_14_27 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_24, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_27
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_11_22
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_26.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_26.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_27.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_21, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_10, s_11})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_28 -> *Constructor_Control_Bind_Bind
Bind1_9_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_28
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_28.V1), gopurs_runtime.Apply(v_10, s_12), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_11, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_8_20
// TAST (Let): Bind1_9_29 -> *Constructor_Control_Bind_Bind
Bind1_9_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_20, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_29
// TAST (Let): Applicative0_10_30 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_20, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_30
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_7_18
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_29.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_29.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_30.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_17.V1), gopurs_runtime.Apply(v_6, s_8), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_7, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1)
}))
})
})
}))
}))
_ = monadStateT1_4_3
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_4_3
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(f_5, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1), gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 gopurs_runtime.Value
{
var __t_tag_31 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v2_9.UnsafePtr).V0
if (__t_tag_31.Type == 9 && __t_tag_31.IntVal == 525585346) {
__t33 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Tuple_Tuple)(v2_9.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v2_9.UnsafePtr).V1})}})}
goto end_branch_33
} else {

}
}
{
var __t_tag_32 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v2_9.UnsafePtr).V0
if (__t_tag_32.Type == 9 && __t_tag_32.IntVal == 60402430) {
__t33 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Data_Tuple_Tuple)(v2_9.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v2_9.UnsafePtr).V1})}})}
goto end_branch_33
} else {

}
}
{
__t33 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_33:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), __t33)
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_6, s_7})})
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
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorStateT1_5_3 -> gopurs_runtime.Value
functorStateT1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_3
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_6 -> *Constructor_Control_Bind_Bind
Bind1_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_6
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_8
// TAST (Let): functorStateT1_9_7 -> gopurs_runtime.Value
functorStateT1_9_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_8, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_7
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_10 -> *Constructor_Control_Bind_Bind
Bind1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_10
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_10.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_9
// TAST (Let): Bind1_11_11 -> *Constructor_Control_Bind_Bind
Bind1_11_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_11
// TAST (Let): Applicative0_12_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_7
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_11.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_11.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_12.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_6.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_5
// TAST (Let): Bind1_7_13 -> *Constructor_Control_Bind_Bind
Bind1_7_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_13
// TAST (Let): Applicative0_8_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_3
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_14.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_4, s_5})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_15 -> *Constructor_Control_Bind_Bind
Bind1_3_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_15
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_17 -> gopurs_runtime.Value
__local_var_5_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_17
// TAST (Let): functorStateT1_5_16 -> gopurs_runtime.Value
functorStateT1_5_16 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_17, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_16
// TAST (Let): __local_var_6_18 -> gopurs_runtime.Value
__local_var_6_18 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_7_19 -> gopurs_runtime.Value
pure_7_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_19
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_21 -> gopurs_runtime.Value
__local_var_9_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_21
// TAST (Let): functorStateT1_9_20 -> gopurs_runtime.Value
functorStateT1_9_20 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_21, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_20
// TAST (Let): __local_var_10_22 -> gopurs_runtime.Value
__local_var_10_22 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_23 -> *Constructor_Control_Bind_Bind
Bind1_11_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_23
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_23.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_22
// TAST (Let): Bind1_11_24 -> *Constructor_Control_Bind_Bind
Bind1_11_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_22, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_24
// TAST (Let): Applicative0_12_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_22, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_20
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_24.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_24.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_25.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_19, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_8, s_9})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_26 -> *Constructor_Control_Bind_Bind
Bind1_7_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_26
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_26.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_18
// TAST (Let): Bind1_7_27 -> *Constructor_Control_Bind_Bind
Bind1_7_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_27
// TAST (Let): Applicative0_8_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_28
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_16
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_27.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_27.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_28.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_15.V1), gopurs_runtime.Apply(v_4, s_6), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1)
}))
})
})
}))
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
// TAST (Let): pure_4_3 -> gopurs_runtime.Value
pure_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_3
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_5
// TAST (Let): functorStateT1_6_4 -> gopurs_runtime.Value
functorStateT1_6_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_5, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_8, s_9))
})
})
}))
_ = functorStateT1_6_4
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_7 -> *Constructor_Control_Bind_Bind
Bind1_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_7
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_9
// TAST (Let): functorStateT1_10_8 -> gopurs_runtime.Value
functorStateT1_10_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_9, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_12, s_13))
})
})
}))
_ = functorStateT1_10_8
// TAST (Let): __local_var_11_10 -> gopurs_runtime.Value
__local_var_11_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_11 -> *Constructor_Control_Bind_Bind
Bind1_12_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_11
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad1_1_0)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_11.V1), gopurs_runtime.Apply(v_13, s_15), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_14, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_11_10
// TAST (Let): Bind1_12_12 -> *Constructor_Control_Bind_Bind
Bind1_12_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_12
// TAST (Let): Applicative0_13_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_13
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_10_8
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_12.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_12.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_13.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_7.V1), gopurs_runtime.Apply(v_9, s_11), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_10, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_7_6
// TAST (Let): Bind1_8_14 -> *Constructor_Control_Bind_Bind
Bind1_8_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_14
// TAST (Let): Applicative0_9_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_15
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_6_4
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_14.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_14.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_15.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, s_6})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_16 -> *Constructor_Control_Bind_Bind
Bind1_4_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_16
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_18 -> gopurs_runtime.Value
__local_var_6_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_18
// TAST (Let): functorStateT1_6_17 -> gopurs_runtime.Value
functorStateT1_6_17 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_18, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_8, s_9))
})
})
}))
_ = functorStateT1_6_17
// TAST (Let): __local_var_7_19 -> gopurs_runtime.Value
__local_var_7_19 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_8_20 -> gopurs_runtime.Value
pure_8_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_20
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_22 -> gopurs_runtime.Value
__local_var_10_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_22
// TAST (Let): functorStateT1_10_21 -> gopurs_runtime.Value
functorStateT1_10_21 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_22, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_12, s_13))
})
})
}))
_ = functorStateT1_10_21
// TAST (Let): __local_var_11_23 -> gopurs_runtime.Value
__local_var_11_23 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_24 -> *Constructor_Control_Bind_Bind
Bind1_12_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_24
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad1_1_0)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_24.V1), gopurs_runtime.Apply(v_13, s_15), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_14, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_11_23
// TAST (Let): Bind1_12_25 -> *Constructor_Control_Bind_Bind
Bind1_12_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_23, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_25
// TAST (Let): Applicative0_13_26 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_23, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_26
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_10_21
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_25.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_25.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_26.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_8_20, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, s_10})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_27 -> *Constructor_Control_Bind_Bind
Bind1_8_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_27
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad1_1_0)
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_27.V1), gopurs_runtime.Apply(v_9, s_11), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_10, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_7_19
// TAST (Let): Bind1_8_28 -> *Constructor_Control_Bind_Bind
Bind1_8_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_19, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_28
// TAST (Let): Applicative0_9_29 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_19, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_29
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_6_17
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_28.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_28.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_29.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_16.V1), gopurs_runtime.Apply(v_5, s_7), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_6, (*Constructor_Data_Tuple_Tuple)(v1_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_8.UnsafePtr).V1)
}))
})
})
}))
}))
_ = monadStateT1_3_2
// TAST (Let): Bind1_4_31 -> *Constructor_Control_Bind_Bind
Bind1_4_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_31
// TAST (Let): pure_5_32 -> gopurs_runtime.Value
pure_5_32 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_32
// TAST (Let): __local_var_4_30 -> gopurs_runtime.Value
__local_var_4_30 := gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_31.V1), m_6, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_32, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_8, s_7})})
}))
})
})
_ = __local_var_4_30
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_30, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_5))
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
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): Applicative0_4_3 -> gopurs_runtime.Value
Applicative0_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_4_3
// TAST (Let): Monoid0_5_4 -> gopurs_runtime.Value
Monoid0_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0"), gopurs_runtime.Value{})
_ = Monoid0_5_4
// TAST (Let): Monad1_6_6 -> gopurs_runtime.Value
Monad1_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_6_6
// TAST (Let): Semigroup0_7_7 -> gopurs_runtime.Value
Semigroup0_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_7_7
// TAST (Let): monadStateT1_8_8 -> gopurs_runtime.Value
monadStateT1_8_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_9_9 -> gopurs_runtime.Value
pure_9_9 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_9
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_11
// TAST (Let): functorStateT1_11_10 -> gopurs_runtime.Value
functorStateT1_11_10 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_11, "map"), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_13, s_14))
})
})
}))
_ = functorStateT1_11_10
// TAST (Let): __local_var_12_12 -> gopurs_runtime.Value
__local_var_12_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_13_13 -> gopurs_runtime.Value
pure_13_13 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_13
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_15 -> gopurs_runtime.Value
__local_var_15_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_15
// TAST (Let): functorStateT1_15_14 -> gopurs_runtime.Value
functorStateT1_15_14 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_15, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_17, s_18))
})
})
}))
_ = functorStateT1_15_14
// TAST (Let): __local_var_16_16 -> gopurs_runtime.Value
__local_var_16_16 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_17 -> *Constructor_Control_Bind_Bind
Bind1_17_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_17
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_19 -> gopurs_runtime.Value
__local_var_19_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_19
// TAST (Let): functorStateT1_19_18 -> gopurs_runtime.Value
functorStateT1_19_18 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_19, "map"), gopurs_runtime.Func(func(v1_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_21, s_22))
})
})
}))
_ = functorStateT1_19_18
// TAST (Let): __local_var_20_20 -> gopurs_runtime.Value
__local_var_20_20 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_21 -> *Constructor_Control_Bind_Bind
Bind1_21_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_21
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_21.V1), gopurs_runtime.Apply(v_22, s_24), gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_23, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_20_20
// TAST (Let): Bind1_21_22 -> *Constructor_Control_Bind_Bind
Bind1_21_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_20, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_22
// TAST (Let): Applicative0_22_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_20, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_23
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_19_18
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_22.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_22.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_23.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_17.V1), gopurs_runtime.Apply(v_18, s_20), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_19, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_16_16
// TAST (Let): Bind1_17_24 -> *Constructor_Control_Bind_Bind
Bind1_17_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_16, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_24
// TAST (Let): Applicative0_18_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_16, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_15_14
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_24.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_24.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_25.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_13_13, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_14, s_15})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_26 -> *Constructor_Control_Bind_Bind
Bind1_13_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_26
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_28 -> gopurs_runtime.Value
__local_var_15_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_28
// TAST (Let): functorStateT1_15_27 -> gopurs_runtime.Value
functorStateT1_15_27 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_28, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_17, s_18))
})
})
}))
_ = functorStateT1_15_27
// TAST (Let): __local_var_16_29 -> gopurs_runtime.Value
__local_var_16_29 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_17_30 -> gopurs_runtime.Value
pure_17_30 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_17_30
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_32 -> gopurs_runtime.Value
__local_var_19_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_32
// TAST (Let): functorStateT1_19_31 -> gopurs_runtime.Value
functorStateT1_19_31 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_32, "map"), gopurs_runtime.Func(func(v1_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_21, s_22))
})
})
}))
_ = functorStateT1_19_31
// TAST (Let): __local_var_20_33 -> gopurs_runtime.Value
__local_var_20_33 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_34 -> *Constructor_Control_Bind_Bind
Bind1_21_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_34
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_36 -> gopurs_runtime.Value
__local_var_23_36 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_23_36
// TAST (Let): functorStateT1_23_35 -> gopurs_runtime.Value
functorStateT1_23_35 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_23_36, "map"), gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_24, (*Constructor_Data_Tuple_Tuple)(v1_27.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_27.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_25, s_26))
})
})
}))
_ = functorStateT1_23_35
// TAST (Let): __local_var_24_37 -> gopurs_runtime.Value
__local_var_24_37 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(_dollar__unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_25_38 -> *Constructor_Control_Bind_Bind
Bind1_25_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_25_38
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_25_38.V1), gopurs_runtime.Apply(v_26, s_28), gopurs_runtime.Func(func(v1_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_27, (*Constructor_Data_Tuple_Tuple)(v1_29.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_29.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_24_37
// TAST (Let): Bind1_25_39 -> *Constructor_Control_Bind_Bind
Bind1_25_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_37, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_25_39
// TAST (Let): Applicative0_26_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_26_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_37, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_26_40
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_23_35
}), gopurs_runtime.Func(func(f_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_25_39.V1), f_27, gopurs_runtime.Func(func(f_prime_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_25_39.V1), a_28, gopurs_runtime.Func(func(a_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_26_40.V1), gopurs_runtime.Apply(f_prime_29, a_prime_30))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_34.V1), gopurs_runtime.Apply(v_22, s_24), gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_23, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_20_33
// TAST (Let): Bind1_21_41 -> *Constructor_Control_Bind_Bind
Bind1_21_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_33, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_41
// TAST (Let): Applicative0_22_42 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_33, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_42
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_19_31
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_41.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_41.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_42.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_17_30, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_18, s_19})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_43 -> *Constructor_Control_Bind_Bind
Bind1_17_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_43
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_45 -> gopurs_runtime.Value
__local_var_19_45 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_45
// TAST (Let): functorStateT1_19_44 -> gopurs_runtime.Value
functorStateT1_19_44 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_45, "map"), gopurs_runtime.Func(func(v1_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_21, s_22))
})
})
}))
_ = functorStateT1_19_44
// TAST (Let): __local_var_20_46 -> gopurs_runtime.Value
__local_var_20_46 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_21_47 -> gopurs_runtime.Value
pure_21_47 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_47
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_49 -> gopurs_runtime.Value
__local_var_23_49 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_23_49
// TAST (Let): functorStateT1_23_48 -> gopurs_runtime.Value
functorStateT1_23_48 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_23_49, "map"), gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_24, (*Constructor_Data_Tuple_Tuple)(v1_27.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_27.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_25, s_26))
})
})
}))
_ = functorStateT1_23_48
// TAST (Let): __local_var_24_50 -> gopurs_runtime.Value
__local_var_24_50 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(_dollar__unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_25_51 -> *Constructor_Control_Bind_Bind
Bind1_25_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_25_51
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_25_51.V1), gopurs_runtime.Apply(v_26, s_28), gopurs_runtime.Func(func(v1_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_27, (*Constructor_Data_Tuple_Tuple)(v1_29.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_29.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_24_50
// TAST (Let): Bind1_25_52 -> *Constructor_Control_Bind_Bind
Bind1_25_52 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_50, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_25_52
// TAST (Let): Applicative0_26_53 -> *Constructor_Control_Applicative_Applicative
Applicative0_26_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_50, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_26_53
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_23_48
}), gopurs_runtime.Func(func(f_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_25_52.V1), f_27, gopurs_runtime.Func(func(f_prime_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_25_52.V1), a_28, gopurs_runtime.Func(func(a_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_26_53.V1), gopurs_runtime.Apply(f_prime_29, a_prime_30))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_21_47, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_22, s_23})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_54 -> *Constructor_Control_Bind_Bind
Bind1_21_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_54
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_54.V1), gopurs_runtime.Apply(v_22, s_24), gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_23, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_20_46
// TAST (Let): Bind1_21_55 -> *Constructor_Control_Bind_Bind
Bind1_21_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_46, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_55
// TAST (Let): Applicative0_22_56 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_46, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_56
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_19_44
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_55.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_55.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_56.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_43.V1), gopurs_runtime.Apply(v_18, s_20), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_19, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_16_29
// TAST (Let): Bind1_17_57 -> *Constructor_Control_Bind_Bind
Bind1_17_57 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_29, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_57
// TAST (Let): Applicative0_18_58 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_58 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_29, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_58
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_15_27
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_57.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_57.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_58.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_26.V1), gopurs_runtime.Apply(v_14, s_16), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_15, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_12_12
// TAST (Let): Bind1_13_59 -> *Constructor_Control_Bind_Bind
Bind1_13_59 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_59
// TAST (Let): Applicative0_14_60 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_60
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_11_10
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_59.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_59.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_60.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_9, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_10, s_11})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_61 -> *Constructor_Control_Bind_Bind
Bind1_9_61 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_61
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_63 -> gopurs_runtime.Value
__local_var_11_63 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_63
// TAST (Let): functorStateT1_11_62 -> gopurs_runtime.Value
functorStateT1_11_62 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_63, "map"), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_13, s_14))
})
})
}))
_ = functorStateT1_11_62
// TAST (Let): __local_var_12_64 -> gopurs_runtime.Value
__local_var_12_64 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_13_65 -> gopurs_runtime.Value
pure_13_65 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_65
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_67 -> gopurs_runtime.Value
__local_var_15_67 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_67
// TAST (Let): functorStateT1_15_66 -> gopurs_runtime.Value
functorStateT1_15_66 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_67, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_17, s_18))
})
})
}))
_ = functorStateT1_15_66
// TAST (Let): __local_var_16_68 -> gopurs_runtime.Value
__local_var_16_68 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_17_69 -> gopurs_runtime.Value
pure_17_69 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_17_69
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_71 -> gopurs_runtime.Value
__local_var_19_71 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_71
// TAST (Let): functorStateT1_19_70 -> gopurs_runtime.Value
functorStateT1_19_70 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_71, "map"), gopurs_runtime.Func(func(v1_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_21, s_22))
})
})
}))
_ = functorStateT1_19_70
// TAST (Let): __local_var_20_72 -> gopurs_runtime.Value
__local_var_20_72 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_73 -> *Constructor_Control_Bind_Bind
Bind1_21_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_73
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_75 -> gopurs_runtime.Value
__local_var_23_75 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_23_75
// TAST (Let): functorStateT1_23_74 -> gopurs_runtime.Value
functorStateT1_23_74 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_23_75, "map"), gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_24, (*Constructor_Data_Tuple_Tuple)(v1_27.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_27.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_25, s_26))
})
})
}))
_ = functorStateT1_23_74
// TAST (Let): __local_var_24_76 -> gopurs_runtime.Value
__local_var_24_76 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(_dollar__unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_25_77 -> *Constructor_Control_Bind_Bind
Bind1_25_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_25_77
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_25_77.V1), gopurs_runtime.Apply(v_26, s_28), gopurs_runtime.Func(func(v1_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_27, (*Constructor_Data_Tuple_Tuple)(v1_29.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_29.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_24_76
// TAST (Let): Bind1_25_78 -> *Constructor_Control_Bind_Bind
Bind1_25_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_76, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_25_78
// TAST (Let): Applicative0_26_79 -> *Constructor_Control_Applicative_Applicative
Applicative0_26_79 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_76, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_26_79
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_23_74
}), gopurs_runtime.Func(func(f_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_25_78.V1), f_27, gopurs_runtime.Func(func(f_prime_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_25_78.V1), a_28, gopurs_runtime.Func(func(a_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_26_79.V1), gopurs_runtime.Apply(f_prime_29, a_prime_30))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_73.V1), gopurs_runtime.Apply(v_22, s_24), gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_23, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_20_72
// TAST (Let): Bind1_21_80 -> *Constructor_Control_Bind_Bind
Bind1_21_80 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_72, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_80
// TAST (Let): Applicative0_22_81 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_72, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_81
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_19_70
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_80.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_80.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_81.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_17_69, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_18, s_19})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_82 -> *Constructor_Control_Bind_Bind
Bind1_17_82 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_82
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_84 -> gopurs_runtime.Value
__local_var_19_84 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_84
// TAST (Let): functorStateT1_19_83 -> gopurs_runtime.Value
functorStateT1_19_83 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_84, "map"), gopurs_runtime.Func(func(v1_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_21, s_22))
})
})
}))
_ = functorStateT1_19_83
// TAST (Let): __local_var_20_85 -> gopurs_runtime.Value
__local_var_20_85 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_21_86 -> gopurs_runtime.Value
pure_21_86 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_86
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_88 -> gopurs_runtime.Value
__local_var_23_88 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_23_88
// TAST (Let): functorStateT1_23_87 -> gopurs_runtime.Value
functorStateT1_23_87 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_23_88, "map"), gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_24, (*Constructor_Data_Tuple_Tuple)(v1_27.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_27.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_25, s_26))
})
})
}))
_ = functorStateT1_23_87
// TAST (Let): __local_var_24_89 -> gopurs_runtime.Value
__local_var_24_89 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(_dollar__unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_25_90 -> *Constructor_Control_Bind_Bind
Bind1_25_90 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_25_90
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_25_90.V1), gopurs_runtime.Apply(v_26, s_28), gopurs_runtime.Func(func(v1_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_27, (*Constructor_Data_Tuple_Tuple)(v1_29.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_29.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_24_89
// TAST (Let): Bind1_25_91 -> *Constructor_Control_Bind_Bind
Bind1_25_91 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_89, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_25_91
// TAST (Let): Applicative0_26_92 -> *Constructor_Control_Applicative_Applicative
Applicative0_26_92 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_89, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_26_92
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_24 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_23_87
}), gopurs_runtime.Func(func(f_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_25_91.V1), f_27, gopurs_runtime.Func(func(f_prime_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_25_91.V1), a_28, gopurs_runtime.Func(func(a_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_26_92.V1), gopurs_runtime.Apply(f_prime_29, a_prime_30))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_21_86, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_22, s_23})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_93 -> *Constructor_Control_Bind_Bind
Bind1_21_93 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_93
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_93.V1), gopurs_runtime.Apply(v_22, s_24), gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_23, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_20_85
// TAST (Let): Bind1_21_94 -> *Constructor_Control_Bind_Bind
Bind1_21_94 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_85, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_94
// TAST (Let): Applicative0_22_95 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_95 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_85, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_95
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_19_83
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_94.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_94.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_95.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_82.V1), gopurs_runtime.Apply(v_18, s_20), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_19, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_16_68
// TAST (Let): Bind1_17_96 -> *Constructor_Control_Bind_Bind
Bind1_17_96 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_68, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_96
// TAST (Let): Applicative0_18_97 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_97 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_68, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_97
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_15_66
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_96.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_96.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_97.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_13_65, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_14, s_15})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_98 -> *Constructor_Control_Bind_Bind
Bind1_13_98 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_98
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_100 -> gopurs_runtime.Value
__local_var_15_100 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_100
// TAST (Let): functorStateT1_15_99 -> gopurs_runtime.Value
functorStateT1_15_99 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_100, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_17, s_18))
})
})
}))
_ = functorStateT1_15_99
// TAST (Let): __local_var_16_101 -> gopurs_runtime.Value
__local_var_16_101 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_17_102 -> gopurs_runtime.Value
pure_17_102 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_17_102
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_104 -> gopurs_runtime.Value
__local_var_19_104 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_104
// TAST (Let): functorStateT1_19_103 -> gopurs_runtime.Value
functorStateT1_19_103 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_104, "map"), gopurs_runtime.Func(func(v1_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_21, s_22))
})
})
}))
_ = functorStateT1_19_103
// TAST (Let): __local_var_20_105 -> gopurs_runtime.Value
__local_var_20_105 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_106 -> *Constructor_Control_Bind_Bind
Bind1_21_106 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_106
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_106.V1), gopurs_runtime.Apply(v_22, s_24), gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_23, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_20_105
// TAST (Let): Bind1_21_107 -> *Constructor_Control_Bind_Bind
Bind1_21_107 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_105, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_107
// TAST (Let): Applicative0_22_108 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_108 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_105, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_108
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_19_103
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_107.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_107.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_108.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_17_102, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_18, s_19})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_109 -> *Constructor_Control_Bind_Bind
Bind1_17_109 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_109
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad1_6_6)
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_109.V1), gopurs_runtime.Apply(v_18, s_20), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_19, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_16_101
// TAST (Let): Bind1_17_110 -> *Constructor_Control_Bind_Bind
Bind1_17_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_101, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_110
// TAST (Let): Applicative0_18_111 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_111 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_101, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_111
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_15_99
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_110.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_110.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_111.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_98.V1), gopurs_runtime.Apply(v_14, s_16), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_15, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_12_64
// TAST (Let): Bind1_13_112 -> *Constructor_Control_Bind_Bind
Bind1_13_112 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_64, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_112
// TAST (Let): Applicative0_14_113 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_113 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_64, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_113
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_11_62
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_112.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_112.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_113.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_61.V1), gopurs_runtime.Apply(v_10, s_12), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_11, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1)
}))
})
})
}))
}))
_ = monadStateT1_8_8
// TAST (Let): Bind1_9_115 -> *Constructor_Control_Bind_Bind
Bind1_9_115 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_115
// TAST (Let): pure_10_116 -> gopurs_runtime.Value
pure_10_116 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_6_6, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_116
// TAST (Let): __local_var_9_114 -> gopurs_runtime.Value
__local_var_9_114 := gopurs_runtime.Func(func(m_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_115.V1), m_11, gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_10_116, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_13, s_12})})
}))
})
})
_ = __local_var_9_114
// TAST (Let): monadTellStateT1_6_5 -> gopurs_runtime.Value
monadTellStateT1_6_5 := gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_8_8
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_7_7
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_114, gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "tell"), x_10))
}))
_ = monadTellStateT1_6_5
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellStateT1_6_5
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_5_4
}), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), gopurs_runtime.Apply(m_7, s_8)), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_4_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V1})}, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0.UnsafePtr).V1})})
}))
})
}), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(m_7, s_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_4_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V1})}, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0.UnsafePtr).V1})})
})))
})
}))
}

func Call_Control_Monad_State_Trans_monadThrowStateT(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
// TAST (Let): Monad0_1_0 -> *Constructor_Control_Monad_Monad
Monad0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): monadStateT1_2_1 -> gopurs_runtime.Value
monadStateT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_4_3 -> gopurs_runtime.Value
pure_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_3
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_5
// TAST (Let): functorStateT1_6_4 -> gopurs_runtime.Value
functorStateT1_6_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_5, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_8, s_9))
})
})
}))
_ = functorStateT1_6_4
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_7 -> *Constructor_Control_Bind_Bind
Bind1_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_7
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_9
// TAST (Let): functorStateT1_10_8 -> gopurs_runtime.Value
functorStateT1_10_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_9, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_12, s_13))
})
})
}))
_ = functorStateT1_10_8
// TAST (Let): __local_var_11_10 -> gopurs_runtime.Value
__local_var_11_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_11 -> *Constructor_Control_Bind_Bind
Bind1_12_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_11
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_2_2)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_11.V1), gopurs_runtime.Apply(v_13, s_15), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_14, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_11_10
// TAST (Let): Bind1_12_12 -> *Constructor_Control_Bind_Bind
Bind1_12_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_12
// TAST (Let): Applicative0_13_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_13
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_10_8
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_12.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_12.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_13.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_7.V1), gopurs_runtime.Apply(v_9, s_11), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_10, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_7_6
// TAST (Let): Bind1_8_14 -> *Constructor_Control_Bind_Bind
Bind1_8_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_14
// TAST (Let): Applicative0_9_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_15
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_6_4
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_14.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_14.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_15.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, s_6})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_16 -> *Constructor_Control_Bind_Bind
Bind1_4_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_16
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_18 -> gopurs_runtime.Value
__local_var_6_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_18
// TAST (Let): functorStateT1_6_17 -> gopurs_runtime.Value
functorStateT1_6_17 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_18, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_8, s_9))
})
})
}))
_ = functorStateT1_6_17
// TAST (Let): __local_var_7_19 -> gopurs_runtime.Value
__local_var_7_19 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_8_20 -> gopurs_runtime.Value
pure_8_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_20
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_22 -> gopurs_runtime.Value
__local_var_10_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_22
// TAST (Let): functorStateT1_10_21 -> gopurs_runtime.Value
functorStateT1_10_21 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_22, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_12, s_13))
})
})
}))
_ = functorStateT1_10_21
// TAST (Let): __local_var_11_23 -> gopurs_runtime.Value
__local_var_11_23 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_24 -> *Constructor_Control_Bind_Bind
Bind1_12_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_24
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_2_2)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_24.V1), gopurs_runtime.Apply(v_13, s_15), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_14, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_11_23
// TAST (Let): Bind1_12_25 -> *Constructor_Control_Bind_Bind
Bind1_12_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_23, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_25
// TAST (Let): Applicative0_13_26 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_23, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_26
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_10_21
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_25.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_25.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_26.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_8_20, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, s_10})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_27 -> *Constructor_Control_Bind_Bind
Bind1_8_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_27
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_2_2)
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_27.V1), gopurs_runtime.Apply(v_9, s_11), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_10, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_7_19
// TAST (Let): Bind1_8_28 -> *Constructor_Control_Bind_Bind
Bind1_8_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_19, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_28
// TAST (Let): Applicative0_9_29 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_19, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_29
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_6_17
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_28.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_28.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_29.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_16.V1), gopurs_runtime.Apply(v_5, s_7), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_6, (*Constructor_Data_Tuple_Tuple)(v1_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_8.UnsafePtr).V1)
}))
})
})
}))
}))
_ = monadStateT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}), gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_30 -> *Constructor_Control_Bind_Bind
Bind1_4_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V1), gopurs_runtime.Value{}))
_ = Bind1_4_30
// TAST (Let): pure_5_31 -> gopurs_runtime.Value
pure_5_31 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_5_31
// TAST (Let): __local_var_6_32 -> gopurs_runtime.Value
__local_var_6_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_3)
_ = __local_var_6_32
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_30.V1), __local_var_6_32, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_31, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_8, s_7})})
}))
})
}))
}

func Call_Control_Monad_State_Trans_monadErrorStateT(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Monad0_2_2 -> *Constructor_Control_Monad_Monad
Monad0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_2_2
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): monadStateT1_3_3 -> gopurs_runtime.Value
monadStateT1_3_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_5_5 -> gopurs_runtime.Value
pure_5_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_5
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_7
// TAST (Let): functorStateT1_7_6 -> gopurs_runtime.Value
functorStateT1_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_7, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_9, s_10))
})
})
}))
_ = functorStateT1_7_6
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_9_9 -> gopurs_runtime.Value
pure_9_9 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_9
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_11
// TAST (Let): functorStateT1_11_10 -> gopurs_runtime.Value
functorStateT1_11_10 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_11, "map"), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_13, s_14))
})
})
}))
_ = functorStateT1_11_10
// TAST (Let): __local_var_12_12 -> gopurs_runtime.Value
__local_var_12_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_13 -> *Constructor_Control_Bind_Bind
Bind1_13_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_13
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_15 -> gopurs_runtime.Value
__local_var_15_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_15
// TAST (Let): functorStateT1_15_14 -> gopurs_runtime.Value
functorStateT1_15_14 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_15, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_17, s_18))
})
})
}))
_ = functorStateT1_15_14
// TAST (Let): __local_var_16_16 -> gopurs_runtime.Value
__local_var_16_16 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_17 -> *Constructor_Control_Bind_Bind
Bind1_17_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_17
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_17.V1), gopurs_runtime.Apply(v_18, s_20), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_19, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_16_16
// TAST (Let): Bind1_17_18 -> *Constructor_Control_Bind_Bind
Bind1_17_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_16, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_18
// TAST (Let): Applicative0_18_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_16, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_19
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_15_14
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_18.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_18.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_19.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_13.V1), gopurs_runtime.Apply(v_14, s_16), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_15, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_12_12
// TAST (Let): Bind1_13_20 -> *Constructor_Control_Bind_Bind
Bind1_13_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_20
// TAST (Let): Applicative0_14_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_21
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_11_10
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_20.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_20.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_21.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_9, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_10, s_11})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_22 -> *Constructor_Control_Bind_Bind
Bind1_9_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_22
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_24 -> gopurs_runtime.Value
__local_var_11_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_24
// TAST (Let): functorStateT1_11_23 -> gopurs_runtime.Value
functorStateT1_11_23 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_24, "map"), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_13, s_14))
})
})
}))
_ = functorStateT1_11_23
// TAST (Let): __local_var_12_25 -> gopurs_runtime.Value
__local_var_12_25 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_13_26 -> gopurs_runtime.Value
pure_13_26 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_26
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_28 -> gopurs_runtime.Value
__local_var_15_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_28
// TAST (Let): functorStateT1_15_27 -> gopurs_runtime.Value
functorStateT1_15_27 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_28, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_17, s_18))
})
})
}))
_ = functorStateT1_15_27
// TAST (Let): __local_var_16_29 -> gopurs_runtime.Value
__local_var_16_29 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_30 -> *Constructor_Control_Bind_Bind
Bind1_17_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_30
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_32 -> gopurs_runtime.Value
__local_var_19_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_32
// TAST (Let): functorStateT1_19_31 -> gopurs_runtime.Value
functorStateT1_19_31 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_32, "map"), gopurs_runtime.Func(func(v1_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_21, s_22))
})
})
}))
_ = functorStateT1_19_31
// TAST (Let): __local_var_20_33 -> gopurs_runtime.Value
__local_var_20_33 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_34 -> *Constructor_Control_Bind_Bind
Bind1_21_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_34
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_34.V1), gopurs_runtime.Apply(v_22, s_24), gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_23, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_20_33
// TAST (Let): Bind1_21_35 -> *Constructor_Control_Bind_Bind
Bind1_21_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_33, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_35
// TAST (Let): Applicative0_22_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_33, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_36
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_19_31
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_35.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_35.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_36.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_30.V1), gopurs_runtime.Apply(v_18, s_20), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_19, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_16_29
// TAST (Let): Bind1_17_37 -> *Constructor_Control_Bind_Bind
Bind1_17_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_29, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_37
// TAST (Let): Applicative0_18_38 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_29, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_38
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_15_27
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_37.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_37.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_38.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_13_26, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_14, s_15})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_39 -> *Constructor_Control_Bind_Bind
Bind1_13_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_39
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_41 -> gopurs_runtime.Value
__local_var_15_41 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_41
// TAST (Let): functorStateT1_15_40 -> gopurs_runtime.Value
functorStateT1_15_40 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_41, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_17, s_18))
})
})
}))
_ = functorStateT1_15_40
// TAST (Let): __local_var_16_42 -> gopurs_runtime.Value
__local_var_16_42 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_17_43 -> gopurs_runtime.Value
pure_17_43 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_17_43
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_45 -> gopurs_runtime.Value
__local_var_19_45 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_45
// TAST (Let): functorStateT1_19_44 -> gopurs_runtime.Value
functorStateT1_19_44 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_45, "map"), gopurs_runtime.Func(func(v1_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_21, s_22))
})
})
}))
_ = functorStateT1_19_44
// TAST (Let): __local_var_20_46 -> gopurs_runtime.Value
__local_var_20_46 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_47 -> *Constructor_Control_Bind_Bind
Bind1_21_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_47
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_47.V1), gopurs_runtime.Apply(v_22, s_24), gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_23, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_20_46
// TAST (Let): Bind1_21_48 -> *Constructor_Control_Bind_Bind
Bind1_21_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_46, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_48
// TAST (Let): Applicative0_22_49 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_46, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_49
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_19_44
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_48.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_48.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_49.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_17_43, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_18, s_19})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_50 -> *Constructor_Control_Bind_Bind
Bind1_17_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_50
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_50.V1), gopurs_runtime.Apply(v_18, s_20), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_19, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_16_42
// TAST (Let): Bind1_17_51 -> *Constructor_Control_Bind_Bind
Bind1_17_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_42, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_51
// TAST (Let): Applicative0_18_52 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_52 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_42, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_52
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_15_40
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_51.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_51.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_52.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_39.V1), gopurs_runtime.Apply(v_14, s_16), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_15, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_12_25
// TAST (Let): Bind1_13_53 -> *Constructor_Control_Bind_Bind
Bind1_13_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_25, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_53
// TAST (Let): Applicative0_14_54 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_25, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_54
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_11_23
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_53.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_53.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_54.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_22.V1), gopurs_runtime.Apply(v_10, s_12), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_11, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_8_8
// TAST (Let): Bind1_9_55 -> *Constructor_Control_Bind_Bind
Bind1_9_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_55
// TAST (Let): Applicative0_10_56 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_56
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_7_6
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_55.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_55.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_56.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_6, s_7})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_57 -> *Constructor_Control_Bind_Bind
Bind1_5_57 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_57
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_59 -> gopurs_runtime.Value
__local_var_7_59 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_59
// TAST (Let): functorStateT1_7_58 -> gopurs_runtime.Value
functorStateT1_7_58 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_59, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_9, s_10))
})
})
}))
_ = functorStateT1_7_58
// TAST (Let): __local_var_8_60 -> gopurs_runtime.Value
__local_var_8_60 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_9_61 -> gopurs_runtime.Value
pure_9_61 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_61
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_63 -> gopurs_runtime.Value
__local_var_11_63 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_63
// TAST (Let): functorStateT1_11_62 -> gopurs_runtime.Value
functorStateT1_11_62 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_63, "map"), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_13, s_14))
})
})
}))
_ = functorStateT1_11_62
// TAST (Let): __local_var_12_64 -> gopurs_runtime.Value
__local_var_12_64 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_13_65 -> gopurs_runtime.Value
pure_13_65 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_65
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_67 -> gopurs_runtime.Value
__local_var_15_67 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_67
// TAST (Let): functorStateT1_15_66 -> gopurs_runtime.Value
functorStateT1_15_66 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_67, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_17, s_18))
})
})
}))
_ = functorStateT1_15_66
// TAST (Let): __local_var_16_68 -> gopurs_runtime.Value
__local_var_16_68 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_69 -> *Constructor_Control_Bind_Bind
Bind1_17_69 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_69
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_71 -> gopurs_runtime.Value
__local_var_19_71 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_71
// TAST (Let): functorStateT1_19_70 -> gopurs_runtime.Value
functorStateT1_19_70 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_71, "map"), gopurs_runtime.Func(func(v1_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_21, s_22))
})
})
}))
_ = functorStateT1_19_70
// TAST (Let): __local_var_20_72 -> gopurs_runtime.Value
__local_var_20_72 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_73 -> *Constructor_Control_Bind_Bind
Bind1_21_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_73
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_73.V1), gopurs_runtime.Apply(v_22, s_24), gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_23, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_20_72
// TAST (Let): Bind1_21_74 -> *Constructor_Control_Bind_Bind
Bind1_21_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_72, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_74
// TAST (Let): Applicative0_22_75 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_75 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_72, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_75
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_19_70
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_74.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_74.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_75.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_69.V1), gopurs_runtime.Apply(v_18, s_20), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_19, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_16_68
// TAST (Let): Bind1_17_76 -> *Constructor_Control_Bind_Bind
Bind1_17_76 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_68, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_76
// TAST (Let): Applicative0_18_77 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_68, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_77
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_15_66
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_76.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_76.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_77.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_13_65, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_14, s_15})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_78 -> *Constructor_Control_Bind_Bind
Bind1_13_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_78
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_80 -> gopurs_runtime.Value
__local_var_15_80 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_80
// TAST (Let): functorStateT1_15_79 -> gopurs_runtime.Value
functorStateT1_15_79 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_80, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_17, s_18))
})
})
}))
_ = functorStateT1_15_79
// TAST (Let): __local_var_16_81 -> gopurs_runtime.Value
__local_var_16_81 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_17_82 -> gopurs_runtime.Value
pure_17_82 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_17_82
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_84 -> gopurs_runtime.Value
__local_var_19_84 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_84
// TAST (Let): functorStateT1_19_83 -> gopurs_runtime.Value
functorStateT1_19_83 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_84, "map"), gopurs_runtime.Func(func(v1_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_21, s_22))
})
})
}))
_ = functorStateT1_19_83
// TAST (Let): __local_var_20_85 -> gopurs_runtime.Value
__local_var_20_85 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_86 -> *Constructor_Control_Bind_Bind
Bind1_21_86 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_86
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_86.V1), gopurs_runtime.Apply(v_22, s_24), gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_23, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_25.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_20_85
// TAST (Let): Bind1_21_87 -> *Constructor_Control_Bind_Bind
Bind1_21_87 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_85, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_87
// TAST (Let): Applicative0_22_88 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_88 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_85, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_88
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_19_83
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_87.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_87.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_88.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_17_82, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_18, s_19})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_89 -> *Constructor_Control_Bind_Bind
Bind1_17_89 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_89
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_89.V1), gopurs_runtime.Apply(v_18, s_20), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_19, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_16_81
// TAST (Let): Bind1_17_90 -> *Constructor_Control_Bind_Bind
Bind1_17_90 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_81, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_90
// TAST (Let): Applicative0_18_91 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_91 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_81, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_91
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_15_79
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_90.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_90.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_91.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_78.V1), gopurs_runtime.Apply(v_14, s_16), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_15, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_12_64
// TAST (Let): Bind1_13_92 -> *Constructor_Control_Bind_Bind
Bind1_13_92 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_64, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_92
// TAST (Let): Applicative0_14_93 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_93 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_64, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_93
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_11_62
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_92.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_92.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_93.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_61, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_10, s_11})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_94 -> *Constructor_Control_Bind_Bind
Bind1_9_94 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_94
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_96 -> gopurs_runtime.Value
__local_var_11_96 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_96
// TAST (Let): functorStateT1_11_95 -> gopurs_runtime.Value
functorStateT1_11_95 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_96, "map"), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_13, s_14))
})
})
}))
_ = functorStateT1_11_95
// TAST (Let): __local_var_12_97 -> gopurs_runtime.Value
__local_var_12_97 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_13_98 -> gopurs_runtime.Value
pure_13_98 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_98
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_100 -> gopurs_runtime.Value
__local_var_15_100 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_100
// TAST (Let): functorStateT1_15_99 -> gopurs_runtime.Value
functorStateT1_15_99 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_100, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_17, s_18))
})
})
}))
_ = functorStateT1_15_99
// TAST (Let): __local_var_16_101 -> gopurs_runtime.Value
__local_var_16_101 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_102 -> *Constructor_Control_Bind_Bind
Bind1_17_102 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_102
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_102.V1), gopurs_runtime.Apply(v_18, s_20), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_19, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_16_101
// TAST (Let): Bind1_17_103 -> *Constructor_Control_Bind_Bind
Bind1_17_103 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_101, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_103
// TAST (Let): Applicative0_18_104 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_104 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_101, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_104
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_15_99
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_103.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_103.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_104.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_13_98, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_14, s_15})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_105 -> *Constructor_Control_Bind_Bind
Bind1_13_105 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_105
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_105.V1), gopurs_runtime.Apply(v_14, s_16), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_15, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_12_97
// TAST (Let): Bind1_13_106 -> *Constructor_Control_Bind_Bind
Bind1_13_106 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_97, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_106
// TAST (Let): Applicative0_14_107 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_107 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_97, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_107
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_11_95
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_106.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_106.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_107.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_94.V1), gopurs_runtime.Apply(v_10, s_12), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_11, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_8_60
// TAST (Let): Bind1_9_108 -> *Constructor_Control_Bind_Bind
Bind1_9_108 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_60, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_108
// TAST (Let): Applicative0_10_109 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_109 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_60, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_109
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_7_58
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_108.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_108.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_109.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_57.V1), gopurs_runtime.Apply(v_6, s_8), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_7, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1)
}))
})
})
}))
}))
_ = monadStateT1_3_3
// TAST (Let): monadThrowStateT1_1_0 -> gopurs_runtime.Value
monadThrowStateT1_1_0 := gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_3_3
}), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_110 -> *Constructor_Control_Bind_Bind
Bind1_5_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_2_2.V1), gopurs_runtime.Value{}))
_ = Bind1_5_110
// TAST (Let): pure_6_111 -> gopurs_runtime.Value
pure_6_111 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_2_2.V0), gopurs_runtime.Value{}), "pure")
_ = pure_6_111
// TAST (Let): __local_var_7_112 -> gopurs_runtime.Value
__local_var_7_112 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "throwError"), e_4)
_ = __local_var_7_112
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_110.V1), __local_var_7_112, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_111, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_9, s_8})})
}))
})
}))
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
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorStateT1_5_3 -> gopurs_runtime.Value
functorStateT1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_3
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_6 -> *Constructor_Control_Bind_Bind
Bind1_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_6
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_8
// TAST (Let): functorStateT1_9_7 -> gopurs_runtime.Value
functorStateT1_9_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_8, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_7
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_10 -> *Constructor_Control_Bind_Bind
Bind1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_10
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_10.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_9
// TAST (Let): Bind1_11_11 -> *Constructor_Control_Bind_Bind
Bind1_11_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_11
// TAST (Let): Applicative0_12_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_7
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_11.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_11.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_12.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_6.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_5
// TAST (Let): Bind1_7_13 -> *Constructor_Control_Bind_Bind
Bind1_7_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_13
// TAST (Let): Applicative0_8_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_3
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_14.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_4, s_5})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_15 -> *Constructor_Control_Bind_Bind
Bind1_3_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_15
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_17 -> gopurs_runtime.Value
__local_var_5_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_17
// TAST (Let): functorStateT1_5_16 -> gopurs_runtime.Value
functorStateT1_5_16 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_17, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_16
// TAST (Let): __local_var_6_18 -> gopurs_runtime.Value
__local_var_6_18 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_7_19 -> gopurs_runtime.Value
pure_7_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_19
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_21 -> gopurs_runtime.Value
__local_var_9_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_21
// TAST (Let): functorStateT1_9_20 -> gopurs_runtime.Value
functorStateT1_9_20 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_21, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_20
// TAST (Let): __local_var_10_22 -> gopurs_runtime.Value
__local_var_10_22 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_23 -> *Constructor_Control_Bind_Bind
Bind1_11_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_23
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_23.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_22
// TAST (Let): Bind1_11_24 -> *Constructor_Control_Bind_Bind
Bind1_11_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_22, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_24
// TAST (Let): Applicative0_12_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_22, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_20
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_24.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_24.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_25.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_19, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_8, s_9})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_26 -> *Constructor_Control_Bind_Bind
Bind1_7_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_26
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_26.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_18
// TAST (Let): Bind1_7_27 -> *Constructor_Control_Bind_Bind
Bind1_7_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_27
// TAST (Let): Applicative0_8_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_28
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_16
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_27.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_27.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_28.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_15.V1), gopurs_runtime.Apply(v_4, s_6), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1)
}))
})
})
}))
}))
_ = monadStateT1_2_1
// TAST (Let): Bind1_3_30 -> *Constructor_Control_Bind_Bind
Bind1_3_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_30
// TAST (Let): pure_4_31 -> gopurs_runtime.Value
pure_4_31 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_31
// TAST (Let): __local_var_3_29 -> gopurs_runtime.Value
__local_var_3_29 := gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_30.V1), m_5, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_31, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_7, s_6})})
}))
})
})
_ = __local_var_3_29
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_29, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_Control_Monad_State_Trans_monoidStateT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_1 -> gopurs_runtime.Value
pure_1_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_1
// TAST (Let): applicativeStateT1_1_0 -> *Constructor_Control_Applicative_Applicative
applicativeStateT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): functorStateT1_3_2 -> gopurs_runtime.Value
functorStateT1_3_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_4, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_5, s_6))
})
})
}))
_ = functorStateT1_3_2
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_5 -> *Constructor_Control_Bind_Bind
Bind1_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_5
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_7
// TAST (Let): functorStateT1_7_6 -> gopurs_runtime.Value
functorStateT1_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_7, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_9, s_10))
})
})
}))
_ = functorStateT1_7_6
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_9 -> *Constructor_Control_Bind_Bind
Bind1_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_9
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_9.V1), gopurs_runtime.Apply(v_10, s_12), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_11, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_8_8
// TAST (Let): Bind1_9_10 -> *Constructor_Control_Bind_Bind
Bind1_9_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_10
// TAST (Let): Applicative0_10_11 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_11
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_7_6
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_10.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_10.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_11.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_5.V1), gopurs_runtime.Apply(v_6, s_8), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_7, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_4_4
// TAST (Let): Bind1_5_12 -> *Constructor_Control_Bind_Bind
Bind1_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_12
// TAST (Let): Applicative0_6_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_13
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_3_2
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_12.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_12.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_13.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_2, s_3})})
})
})))
_ = applicativeStateT1_1_0
// TAST (Let): __local_var_2_16 -> gopurs_runtime.Value
__local_var_2_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_16
// TAST (Let): functorStateT1_2_15 -> gopurs_runtime.Value
functorStateT1_2_15 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_16, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Tuple_Tuple)(v1_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_6.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_4, s_5))
})
})
}))
_ = functorStateT1_2_15
// TAST (Let): __local_var_3_17 -> gopurs_runtime.Value
__local_var_3_17 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_4_18 -> gopurs_runtime.Value
pure_4_18 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_18
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_20 -> gopurs_runtime.Value
__local_var_6_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_20
// TAST (Let): functorStateT1_6_19 -> gopurs_runtime.Value
functorStateT1_6_19 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_20, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_8, s_9))
})
})
}))
_ = functorStateT1_6_19
// TAST (Let): __local_var_7_21 -> gopurs_runtime.Value
__local_var_7_21 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_8_22 -> gopurs_runtime.Value
pure_8_22 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_22
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_24 -> gopurs_runtime.Value
__local_var_10_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_24
// TAST (Let): functorStateT1_10_23 -> gopurs_runtime.Value
functorStateT1_10_23 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_24, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_12, s_13))
})
})
}))
_ = functorStateT1_10_23
// TAST (Let): __local_var_11_25 -> gopurs_runtime.Value
__local_var_11_25 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_26 -> *Constructor_Control_Bind_Bind
Bind1_12_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_26
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_28 -> gopurs_runtime.Value
__local_var_14_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_28
// TAST (Let): functorStateT1_14_27 -> gopurs_runtime.Value
functorStateT1_14_27 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_28, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_16, s_17))
})
})
}))
_ = functorStateT1_14_27
// TAST (Let): __local_var_15_29 -> gopurs_runtime.Value
__local_var_15_29 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_30 -> *Constructor_Control_Bind_Bind
Bind1_16_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_30
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_30.V1), gopurs_runtime.Apply(v_17, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_18, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_15_29
// TAST (Let): Bind1_16_31 -> *Constructor_Control_Bind_Bind
Bind1_16_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_29, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_31
// TAST (Let): Applicative0_17_32 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_29, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_32
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_14_27
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_31.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_31.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_32.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_26.V1), gopurs_runtime.Apply(v_13, s_15), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_14, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_11_25
// TAST (Let): Bind1_12_33 -> *Constructor_Control_Bind_Bind
Bind1_12_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_25, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_33
// TAST (Let): Applicative0_13_34 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_25, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_34
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_10_23
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_33.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_33.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_34.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_8_22, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, s_10})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_35 -> *Constructor_Control_Bind_Bind
Bind1_8_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_35
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_37 -> gopurs_runtime.Value
__local_var_10_37 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_37
// TAST (Let): functorStateT1_10_36 -> gopurs_runtime.Value
functorStateT1_10_36 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_37, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_12, s_13))
})
})
}))
_ = functorStateT1_10_36
// TAST (Let): __local_var_11_38 -> gopurs_runtime.Value
__local_var_11_38 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_12_39 -> gopurs_runtime.Value
pure_12_39 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_39
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_41 -> gopurs_runtime.Value
__local_var_14_41 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_41
// TAST (Let): functorStateT1_14_40 -> gopurs_runtime.Value
functorStateT1_14_40 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_41, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_18.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_16, s_17))
})
})
}))
_ = functorStateT1_14_40
// TAST (Let): __local_var_15_42 -> gopurs_runtime.Value
__local_var_15_42 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_43 -> *Constructor_Control_Bind_Bind
Bind1_16_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_43
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_43.V1), gopurs_runtime.Apply(v_17, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_18, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_20.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_15_42
// TAST (Let): Bind1_16_44 -> *Constructor_Control_Bind_Bind
Bind1_16_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_42, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_44
// TAST (Let): Applicative0_17_45 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_42, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_45
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_14_40
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_44.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_44.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_45.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_12_39, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_13, s_14})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_46 -> *Constructor_Control_Bind_Bind
Bind1_12_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_46
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_46.V1), gopurs_runtime.Apply(v_13, s_15), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_14, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_11_38
// TAST (Let): Bind1_12_47 -> *Constructor_Control_Bind_Bind
Bind1_12_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_38, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_47
// TAST (Let): Applicative0_13_48 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_38, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_48
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_10_36
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_47.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_47.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_48.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_35.V1), gopurs_runtime.Apply(v_9, s_11), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_10, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_7_21
// TAST (Let): Bind1_8_49 -> *Constructor_Control_Bind_Bind
Bind1_8_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_21, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_49
// TAST (Let): Applicative0_9_50 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_21, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_50
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_6_19
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_49.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_49.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_50.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_18, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_5, s_6})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_51 -> *Constructor_Control_Bind_Bind
Bind1_4_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_51
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_53 -> gopurs_runtime.Value
__local_var_6_53 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_53
// TAST (Let): functorStateT1_6_52 -> gopurs_runtime.Value
functorStateT1_6_52 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_53, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_8, s_9))
})
})
}))
_ = functorStateT1_6_52
// TAST (Let): __local_var_7_54 -> gopurs_runtime.Value
__local_var_7_54 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_8_55 -> gopurs_runtime.Value
pure_8_55 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_55
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_57 -> gopurs_runtime.Value
__local_var_10_57 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_57
// TAST (Let): functorStateT1_10_56 -> gopurs_runtime.Value
functorStateT1_10_56 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_57, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_14.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_12, s_13))
})
})
}))
_ = functorStateT1_10_56
// TAST (Let): __local_var_11_58 -> gopurs_runtime.Value
__local_var_11_58 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_59 -> *Constructor_Control_Bind_Bind
Bind1_12_59 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_59
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_59.V1), gopurs_runtime.Apply(v_13, s_15), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_14, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_16.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_11_58
// TAST (Let): Bind1_12_60 -> *Constructor_Control_Bind_Bind
Bind1_12_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_58, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_60
// TAST (Let): Applicative0_13_61 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_61 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_58, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_61
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_10_56
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_60.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_60.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_61.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_8_55, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_9, s_10})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_62 -> *Constructor_Control_Bind_Bind
Bind1_8_62 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_62
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_62.V1), gopurs_runtime.Apply(v_9, s_11), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_10, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_12.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_7_54
// TAST (Let): Bind1_8_63 -> *Constructor_Control_Bind_Bind
Bind1_8_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_54, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_63
// TAST (Let): Applicative0_9_64 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_54, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_64
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_6_52
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_63.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_63.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_64.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_51.V1), gopurs_runtime.Apply(v_5, s_7), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_6, (*Constructor_Data_Tuple_Tuple)(v1_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_8.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_3_17
// TAST (Let): Bind1_4_65 -> *Constructor_Control_Bind_Bind
Bind1_4_65 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_65
// TAST (Let): Applicative0_5_66 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_66 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_66
// TAST (Let): applyStateT1_2_14 -> *Constructor_Control_Apply_Apply
applyStateT1_2_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_2_15
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_65.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_65.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_66.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
})))
_ = applyStateT1_2_14
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_68 -> *Constructor_Data_Functor_Functor
Functor0_4_68 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyStateT1_2_14.V0), gopurs_runtime.Value{}))
_ = Functor0_4_68
// TAST (Let): __local_var_5_69 -> gopurs_runtime.Value
__local_var_5_69 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_5_69
// TAST (Let): semigroupStateT2_4_67 -> gopurs_runtime.Value
semigroupStateT2_4_67 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyStateT1_2_14.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_68.V0), __local_var_5_69, a_6), b_7)
})
}))
_ = semigroupStateT2_4_67
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupStateT2_4_67
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Tuple_Tuple)(v1_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_6.UnsafePtr).V1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Tuple_Tuple)(v1_8.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_8.UnsafePtr).V1})}
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
// TAST (Let): pure_1_1 -> gopurs_runtime.Value
pure_1_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_1
// TAST (Let): applicativeStateT1_1_0 -> gopurs_runtime.Value
applicativeStateT1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): functorStateT1_3_2 -> gopurs_runtime.Value
functorStateT1_3_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_4, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_5, s_6))
})
})
}))
_ = functorStateT1_3_2
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_5 -> *Constructor_Control_Bind_Bind
Bind1_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_5
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_7
// TAST (Let): functorStateT1_7_6 -> gopurs_runtime.Value
functorStateT1_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_7, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_9, s_10))
})
})
}))
_ = functorStateT1_7_6
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_9 -> *Constructor_Control_Bind_Bind
Bind1_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_9
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(dictMonad_0)
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_9.V1), gopurs_runtime.Apply(v_10, s_12), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_11, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_8_8
// TAST (Let): Bind1_9_10 -> *Constructor_Control_Bind_Bind
Bind1_9_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_10
// TAST (Let): Applicative0_10_11 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_11
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_7_6
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_10.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_10.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_11.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_5.V1), gopurs_runtime.Apply(v_6, s_8), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_7, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_4_4
// TAST (Let): Bind1_5_12 -> *Constructor_Control_Bind_Bind
Bind1_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_12
// TAST (Let): Applicative0_6_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_13
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_3_2
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_12.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_12.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_13.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_2, s_3})})
})
}))
_ = applicativeStateT1_1_0
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_15 -> gopurs_runtime.Value
__local_var_3_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_3_15
// TAST (Let): empty_4_16 -> gopurs_runtime.Value
empty_4_16 := gopurs_runtime.RecordGet(__local_var_3_15, "empty")
_ = empty_4_16
// TAST (Let): __local_var_5_18 -> gopurs_runtime.Value
__local_var_5_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_15, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_5_18
// TAST (Let): __local_var_6_20 -> gopurs_runtime.Value
__local_var_6_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_18, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_20
// TAST (Let): functorStateT1_6_19 -> gopurs_runtime.Value
functorStateT1_6_19 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_20, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_10.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_8, s_9))
})
})
}))
_ = functorStateT1_6_19
// TAST (Let): altStateT2_5_17 -> gopurs_runtime.Value
altStateT2_5_17 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_6_19
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_18, "alt"), gopurs_runtime.Apply(v_7, s_9), gopurs_runtime.Apply(v1_8, s_9))
})
})
}))
_ = altStateT2_5_17
// TAST (Let): plusStateT2_3_14 -> gopurs_runtime.Value
plusStateT2_3_14 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return altStateT2_5_17
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_4_16
}))
_ = plusStateT2_3_14
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeStateT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return plusStateT2_3_14
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
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorStateT1_5_3 -> gopurs_runtime.Value
functorStateT1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_3
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_6 -> *Constructor_Control_Bind_Bind
Bind1_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_6
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_8
// TAST (Let): functorStateT1_9_7 -> gopurs_runtime.Value
functorStateT1_9_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_8, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_7
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_10 -> *Constructor_Control_Bind_Bind
Bind1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_10
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_10.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_9
// TAST (Let): Bind1_11_11 -> *Constructor_Control_Bind_Bind
Bind1_11_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_11
// TAST (Let): Applicative0_12_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_12
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_7
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_11.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_11.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_12.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_6.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_5
// TAST (Let): Bind1_7_13 -> *Constructor_Control_Bind_Bind
Bind1_7_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_13
// TAST (Let): Applicative0_8_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_3
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_13.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_14.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_4, s_5})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_15 -> *Constructor_Control_Bind_Bind
Bind1_3_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_15
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_17 -> gopurs_runtime.Value
__local_var_5_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_17
// TAST (Let): functorStateT1_5_16 -> gopurs_runtime.Value
functorStateT1_5_16 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_17, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_16
// TAST (Let): __local_var_6_18 -> gopurs_runtime.Value
__local_var_6_18 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_7_19 -> gopurs_runtime.Value
pure_7_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_19
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_21 -> gopurs_runtime.Value
__local_var_9_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_21
// TAST (Let): functorStateT1_9_20 -> gopurs_runtime.Value
functorStateT1_9_20 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_21, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_20
// TAST (Let): __local_var_10_22 -> gopurs_runtime.Value
__local_var_10_22 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_23 -> *Constructor_Control_Bind_Bind
Bind1_11_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_23
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_23.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_22
// TAST (Let): Bind1_11_24 -> *Constructor_Control_Bind_Bind
Bind1_11_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_22, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_24
// TAST (Let): Applicative0_12_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_22, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_20
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_24.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_24.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_25.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_19, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_8, s_9})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_26 -> *Constructor_Control_Bind_Bind
Bind1_7_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_26
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_26.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_18
// TAST (Let): Bind1_7_27 -> *Constructor_Control_Bind_Bind
Bind1_7_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_27
// TAST (Let): Applicative0_8_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_18, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_28
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_16
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_27.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_27.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_28.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_15.V1), gopurs_runtime.Apply(v_4, s_6), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_7.UnsafePtr).V1)
}))
})
})
}))
}))
_ = monadStateT1_2_1
// TAST (Let): pure_3_31 -> gopurs_runtime.Value
pure_3_31 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_31
// TAST (Let): applicativeStateT1_3_30 -> gopurs_runtime.Value
applicativeStateT1_3_30 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_33 -> gopurs_runtime.Value
__local_var_5_33 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_33
// TAST (Let): functorStateT1_5_32 -> gopurs_runtime.Value
functorStateT1_5_32 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_33, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_9.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_7, s_8))
})
})
}))
_ = functorStateT1_5_32
// TAST (Let): __local_var_6_34 -> gopurs_runtime.Value
__local_var_6_34 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_7_35 -> gopurs_runtime.Value
pure_7_35 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_35
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_37 -> gopurs_runtime.Value
__local_var_9_37 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_37
// TAST (Let): functorStateT1_9_36 -> gopurs_runtime.Value
functorStateT1_9_36 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_37, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_36
// TAST (Let): __local_var_10_38 -> gopurs_runtime.Value
__local_var_10_38 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_39 -> *Constructor_Control_Bind_Bind
Bind1_11_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_39
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_41 -> gopurs_runtime.Value
__local_var_13_41 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_41
// TAST (Let): functorStateT1_13_40 -> gopurs_runtime.Value
functorStateT1_13_40 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_41, "map"), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_15, s_16))
})
})
}))
_ = functorStateT1_13_40
// TAST (Let): __local_var_14_42 -> gopurs_runtime.Value
__local_var_14_42 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_43 -> *Constructor_Control_Bind_Bind
Bind1_15_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_43
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_43.V1), gopurs_runtime.Apply(v_16, s_18), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_17, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_14_42
// TAST (Let): Bind1_15_44 -> *Constructor_Control_Bind_Bind
Bind1_15_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_42, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_44
// TAST (Let): Applicative0_16_45 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_42, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_45
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_13_40
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_44.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_44.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_45.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_39.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_38
// TAST (Let): Bind1_11_46 -> *Constructor_Control_Bind_Bind
Bind1_11_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_38, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_46
// TAST (Let): Applicative0_12_47 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_38, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_47
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_36
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_46.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_46.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_47.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_35, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_8, s_9})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_48 -> *Constructor_Control_Bind_Bind
Bind1_7_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_48
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_50 -> gopurs_runtime.Value
__local_var_9_50 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_50
// TAST (Let): functorStateT1_9_49 -> gopurs_runtime.Value
functorStateT1_9_49 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_50, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_13.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_11, s_12))
})
})
}))
_ = functorStateT1_9_49
// TAST (Let): __local_var_10_51 -> gopurs_runtime.Value
__local_var_10_51 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_11_52 -> gopurs_runtime.Value
pure_11_52 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_52
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_54 -> gopurs_runtime.Value
__local_var_13_54 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_54
// TAST (Let): functorStateT1_13_53 -> gopurs_runtime.Value
functorStateT1_13_53 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_54, "map"), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_15, s_16))
})
})
}))
_ = functorStateT1_13_53
// TAST (Let): __local_var_14_55 -> gopurs_runtime.Value
__local_var_14_55 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_56 -> *Constructor_Control_Bind_Bind
Bind1_15_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_56
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_58 -> gopurs_runtime.Value
__local_var_17_58 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_17_58
// TAST (Let): functorStateT1_17_57 -> gopurs_runtime.Value
functorStateT1_17_57 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_17_58, "map"), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_18, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_19, s_20))
})
})
}))
_ = functorStateT1_17_57
// TAST (Let): __local_var_18_59 -> gopurs_runtime.Value
__local_var_18_59 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_19_60 -> *Constructor_Control_Bind_Bind
Bind1_19_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_19_60
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_19_60.V1), gopurs_runtime.Apply(v_20, s_22), gopurs_runtime.Func(func(v1_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_21, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_18_59
// TAST (Let): Bind1_19_61 -> *Constructor_Control_Bind_Bind
Bind1_19_61 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_59, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_19_61
// TAST (Let): Applicative0_20_62 -> *Constructor_Control_Applicative_Applicative
Applicative0_20_62 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_59, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_20_62
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_17_57
}), gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_19_61.V1), f_21, gopurs_runtime.Func(func(f_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_19_61.V1), a_22, gopurs_runtime.Func(func(a_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_20_62.V1), gopurs_runtime.Apply(f_prime_23, a_prime_24))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_56.V1), gopurs_runtime.Apply(v_16, s_18), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_17, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_14_55
// TAST (Let): Bind1_15_63 -> *Constructor_Control_Bind_Bind
Bind1_15_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_55, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_63
// TAST (Let): Applicative0_16_64 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_55, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_64
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_13_53
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_63.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_63.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_64.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_11_52, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_12, s_13})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_65 -> *Constructor_Control_Bind_Bind
Bind1_11_65 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_65
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_67 -> gopurs_runtime.Value
__local_var_13_67 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_67
// TAST (Let): functorStateT1_13_66 -> gopurs_runtime.Value
functorStateT1_13_66 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_67, "map"), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_17.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_15, s_16))
})
})
}))
_ = functorStateT1_13_66
// TAST (Let): __local_var_14_68 -> gopurs_runtime.Value
__local_var_14_68 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_15_69 -> gopurs_runtime.Value
pure_15_69 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_15_69
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_71 -> gopurs_runtime.Value
__local_var_17_71 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_17_71
// TAST (Let): functorStateT1_17_70 -> gopurs_runtime.Value
functorStateT1_17_70 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_17_71, "map"), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_18, (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_21.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_19, s_20))
})
})
}))
_ = functorStateT1_17_70
// TAST (Let): __local_var_18_72 -> gopurs_runtime.Value
__local_var_18_72 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applicativeStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_19_73 -> *Constructor_Control_Bind_Bind
Bind1_19_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_19_73
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_19_73.V1), gopurs_runtime.Apply(v_20, s_22), gopurs_runtime.Func(func(v1_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_21, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_23.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_18_72
// TAST (Let): Bind1_19_74 -> *Constructor_Control_Bind_Bind
Bind1_19_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_72, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_19_74
// TAST (Let): Applicative0_20_75 -> *Constructor_Control_Applicative_Applicative
Applicative0_20_75 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_72, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_20_75
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_17_70
}), gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_19_74.V1), f_21, gopurs_runtime.Func(func(f_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_19_74.V1), a_22, gopurs_runtime.Func(func(a_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_20_75.V1), gopurs_runtime.Apply(f_prime_23, a_prime_24))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_15_69, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_16, s_17})})
})
}))
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_76 -> *Constructor_Control_Bind_Bind
Bind1_15_76 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_76
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_Trans_applyStateT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_76.V1), gopurs_runtime.Apply(v_16, s_18), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_17, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_19.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_14_68
// TAST (Let): Bind1_15_77 -> *Constructor_Control_Bind_Bind
Bind1_15_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_68, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_77
// TAST (Let): Applicative0_16_78 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_68, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_78
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_13_66
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_77.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_77.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_78.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_65.V1), gopurs_runtime.Apply(v_12, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_13, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_15.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_10_51
// TAST (Let): Bind1_11_79 -> *Constructor_Control_Bind_Bind
Bind1_11_79 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_51, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_79
// TAST (Let): Applicative0_12_80 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_80 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_51, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_80
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_9_49
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_79.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_79.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_80.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_48.V1), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_9, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1)
}))
})
})
}))
}))
_ = __local_var_6_34
// TAST (Let): Bind1_7_81 -> *Constructor_Control_Bind_Bind
Bind1_7_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_34, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_81
// TAST (Let): Applicative0_8_82 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_82 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_34, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_82
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_32
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_81.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_81.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_82.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_31, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_4, s_5})})
})
}))
_ = applicativeStateT1_3_30
// TAST (Let): __local_var_4_84 -> gopurs_runtime.Value
__local_var_4_84 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadPlus_0, "Alternative1"), gopurs_runtime.Value{}), "Plus1"), gopurs_runtime.Value{})
_ = __local_var_4_84
// TAST (Let): empty_5_85 -> gopurs_runtime.Value
empty_5_85 := gopurs_runtime.RecordGet(__local_var_4_84, "empty")
_ = empty_5_85
// TAST (Let): __local_var_6_87 -> gopurs_runtime.Value
__local_var_6_87 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_84, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_6_87
// TAST (Let): __local_var_7_89 -> gopurs_runtime.Value
__local_var_7_89 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_87, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_89
// TAST (Let): functorStateT1_7_88 -> gopurs_runtime.Value
functorStateT1_7_88 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_89, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v1_11.UnsafePtr).V1})}
}), gopurs_runtime.Apply(v_9, s_10))
})
})
}))
_ = functorStateT1_7_88
// TAST (Let): altStateT2_6_86 -> gopurs_runtime.Value
altStateT2_6_86 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_7_88
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_87, "alt"), gopurs_runtime.Apply(v_8, s_10), gopurs_runtime.Apply(v1_9, s_10))
})
})
}))
_ = altStateT2_6_86
// TAST (Let): plusStateT2_4_83 -> gopurs_runtime.Value
plusStateT2_4_83 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return altStateT2_6_86
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_5_85
}))
_ = plusStateT2_4_83
// TAST (Let): alternativeStateT1_3_29 -> gopurs_runtime.Value
alternativeStateT1_3_29 := gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeStateT1_3_30
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return plusStateT2_4_83
}))
_ = alternativeStateT1_3_29
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeStateT1_3_29
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

func Call_Control_Monad_State_Trans_mapStateT__1987836370(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2)))
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

func Call_Control_Monad_State_Trans_withStateT__2250856667(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2)))
}



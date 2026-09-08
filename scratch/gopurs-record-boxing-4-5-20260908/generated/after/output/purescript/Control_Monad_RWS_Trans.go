package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_RWS_Trans_RWSResult gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_RWSResult sync.Once
func Get_Control_Monad_RWS_Trans_RWSResult() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_RWSResult.Do(func() {
		cache_Control_Monad_RWS_Trans_RWSResult = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, value2}))}
})
})
})
	})
	return cache_Control_Monad_RWS_Trans_RWSResult
}

var cache_Control_Monad_RWS_Trans_RWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_RWST sync.Once
func Get_Control_Monad_RWS_Trans_RWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_RWST.Do(func() {
		cache_Control_Monad_RWS_Trans_RWST = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_RWST(x_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_RWST
}

var cache_Control_Monad_RWS_Trans_withRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_withRWST sync.Once
func Get_Control_Monad_RWS_Trans_withRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_withRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_withRWST = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_withRWST(f_0_box, m_1_box, r_2_box, s_3_box)
})
	})
	return cache_Control_Monad_RWS_Trans_withRWST
}

var cache_Control_Monad_RWS_Trans_runRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_runRWST sync.Once
func Get_Control_Monad_RWS_Trans_runRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_runRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_runRWST = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_runRWST(v_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_runRWST
}

var cache_Control_Monad_RWS_Trans_newtypeRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_newtypeRWST sync.Once
func Get_Control_Monad_RWS_Trans_newtypeRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_newtypeRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_newtypeRWST = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
	})
	return cache_Control_Monad_RWS_Trans_newtypeRWST
}

var cache_Control_Monad_RWS_Trans_monadTransRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadTransRWST sync.Once
func Get_Control_Monad_RWS_Trans_monadTransRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadTransRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadTransRWST = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_monadTransRWST(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_monadTransRWST
}

var cache_Control_Monad_RWS_Trans_mapRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_mapRWST sync.Once
func Get_Control_Monad_RWS_Trans_mapRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_mapRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_mapRWST = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_mapRWST(f_0_box, v_1_box, r_2_box, s_3_box)
})
	})
	return cache_Control_Monad_RWS_Trans_mapRWST
}

var cache_Control_Monad_RWS_Trans_lazyRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_lazyRWST sync.Once
func Get_Control_Monad_RWS_Trans_lazyRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_lazyRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_lazyRWST = gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer((&Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, r_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), r_1, s_2)
})}))}
	})
	return cache_Control_Monad_RWS_Trans_lazyRWST
}

var cache_Control_Monad_RWS_Trans_functorRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_functorRWST sync.Once
func Get_Control_Monad_RWS_Trans_functorRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_functorRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_functorRWST = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_functorRWST(dictFunctor_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_functorRWST
}

var cache_Control_Monad_RWS_Trans_execRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_execRWST sync.Once
func Get_Control_Monad_RWS_Trans_execRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_execRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_execRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_execRWST(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_RWS_Trans_execRWST
}

var cache_Control_Monad_RWS_Trans_evalRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_evalRWST sync.Once
func Get_Control_Monad_RWS_Trans_evalRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_evalRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_evalRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_evalRWST(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Control_Monad_RWS_Trans_evalRWST
}

var cache_Control_Monad_RWS_Trans_applyRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_applyRWST sync.Once
func Get_Control_Monad_RWS_Trans_applyRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_applyRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_applyRWST = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_applyRWST(dictBind_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_applyRWST
}

var cache_Control_Monad_RWS_Trans_bindRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_bindRWST sync.Once
func Get_Control_Monad_RWS_Trans_bindRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_bindRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_bindRWST = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_bindRWST(dictBind_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_bindRWST
}

var cache_Control_Monad_RWS_Trans_semigroupRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_semigroupRWST sync.Once
func Get_Control_Monad_RWS_Trans_semigroupRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_semigroupRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_semigroupRWST = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_semigroupRWST(dictBind_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_semigroupRWST
}

var cache_Control_Monad_RWS_Trans_applicativeRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_applicativeRWST sync.Once
func Get_Control_Monad_RWS_Trans_applicativeRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_applicativeRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_applicativeRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_applicativeRWST(dictMonad_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_applicativeRWST
}

var cache_Control_Monad_RWS_Trans_monadRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadRWST sync.Once
func Get_Control_Monad_RWS_Trans_monadRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_monadRWST(dictMonad_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_monadRWST
}

var cache_Control_Monad_RWS_Trans_monadAskRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadAskRWST sync.Once
func Get_Control_Monad_RWS_Trans_monadAskRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadAskRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadAskRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_monadAskRWST(dictMonad_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_monadAskRWST
}

var cache_Control_Monad_RWS_Trans_monadReaderRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadReaderRWST sync.Once
func Get_Control_Monad_RWS_Trans_monadReaderRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadReaderRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadReaderRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_monadReaderRWST(dictMonad_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_monadReaderRWST
}

var cache_Control_Monad_RWS_Trans_monadEffectRWS gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadEffectRWS sync.Once
func Get_Control_Monad_RWS_Trans_monadEffectRWS() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadEffectRWS.Do(func() {
		cache_Control_Monad_RWS_Trans_monadEffectRWS = gopurs_runtime.Func2(func(dictMonoid_0_box gopurs_runtime.Value, dictMonadEffect_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_monadEffectRWS(dictMonoid_0_box, dictMonadEffect_1_box)
})
	})
	return cache_Control_Monad_RWS_Trans_monadEffectRWS
}

var cache_Control_Monad_RWS_Trans_monadRecRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadRecRWST sync.Once
func Get_Control_Monad_RWS_Trans_monadRecRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadRecRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadRecRWST = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_monadRecRWST(dictMonadRec_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_monadRecRWST
}

var cache_Control_Monad_RWS_Trans_monadStateRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadStateRWST sync.Once
func Get_Control_Monad_RWS_Trans_monadStateRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadStateRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadStateRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_monadStateRWST(dictMonad_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_monadStateRWST
}

var cache_Control_Monad_RWS_Trans_monadTellRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadTellRWST sync.Once
func Get_Control_Monad_RWS_Trans_monadTellRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadTellRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadTellRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_monadTellRWST(dictMonad_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_monadTellRWST
}

var cache_Control_Monad_RWS_Trans_monadWriterRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadWriterRWST sync.Once
func Get_Control_Monad_RWS_Trans_monadWriterRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadWriterRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadWriterRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_monadWriterRWST(dictMonad_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_monadWriterRWST
}

var cache_Control_Monad_RWS_Trans_monadThrowRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadThrowRWST sync.Once
func Get_Control_Monad_RWS_Trans_monadThrowRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadThrowRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadThrowRWST = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_monadThrowRWST(dictMonadThrow_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_monadThrowRWST
}

var cache_Control_Monad_RWS_Trans_monadErrorRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadErrorRWST sync.Once
func Get_Control_Monad_RWS_Trans_monadErrorRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadErrorRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadErrorRWST = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_monadErrorRWST(dictMonadError_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_monadErrorRWST
}

var cache_Control_Monad_RWS_Trans_monadSTRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monadSTRWST sync.Once
func Get_Control_Monad_RWS_Trans_monadSTRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monadSTRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monadSTRWST = gopurs_runtime.Func2(func(dictMonoid_0_box gopurs_runtime.Value, dictMonadST_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_monadSTRWST(dictMonoid_0_box, dictMonadST_1_box)
})
	})
	return cache_Control_Monad_RWS_Trans_monadSTRWST
}

var cache_Control_Monad_RWS_Trans_monoidRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_monoidRWST sync.Once
func Get_Control_Monad_RWS_Trans_monoidRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_monoidRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_monoidRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_monoidRWST(dictMonad_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_monoidRWST
}

var cache_Control_Monad_RWS_Trans_altRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_altRWST sync.Once
func Get_Control_Monad_RWS_Trans_altRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_altRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_altRWST = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_altRWST(dictAlt_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_altRWST
}

var cache_Control_Monad_RWS_Trans_plusRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_plusRWST sync.Once
func Get_Control_Monad_RWS_Trans_plusRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_plusRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_plusRWST = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_plusRWST(dictPlus_0_box)
})
	})
	return cache_Control_Monad_RWS_Trans_plusRWST
}

var cache_Control_Monad_RWS_Trans_alternativeRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_alternativeRWST sync.Once
func Get_Control_Monad_RWS_Trans_alternativeRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_alternativeRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_alternativeRWST = gopurs_runtime.Func2(func(dictMonoid_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_alternativeRWST(dictMonoid_0_box, dictAlternative_1_box)
})
	})
	return cache_Control_Monad_RWS_Trans_alternativeRWST
}

type Constructor_Control_Monad_RWS_Trans_RWSResult[T_state any, T_result any, T_writer any] struct {
	Rc uint32
	V0 T_state
	V1 T_result
	V2 T_writer
}


func Call_Control_Monad_RWS_Trans_RWST(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_RWS_Trans_withRWST(f_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
// TAST (Let): __local_var_4_0 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar r1), (TypeVar s)])
__local_var_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, r_2, s_3))
_ = __local_var_4_0
return gopurs_runtime.Apply2(m_1, (__local_var_4_0).V0, (__local_var_4_0).V1)
}

func Call_Control_Monad_RWS_Trans_runRWST(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Control_Monad_RWS_Trans_monadTransRWST(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_0
// TAST (Let): pure_3_1 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
return gopurs_runtime.Func3(func(m_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), m_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_6, a_7, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}))})
}))
})
})}))}
}

func Call_Control_Monad_RWS_Trans_mapRWST(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply2(v_1, r_2, s_3))
}

func Call_Control_Monad_RWS_Trans_functorRWST(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value, s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0, gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_2, r_3, s_4))
})}))}
}

func Call_Control_Monad_RWS_Trans_execRWST(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(v_3, r_4, s_5), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V2}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
})
}

func Call_Control_Monad_RWS_Trans_evalRWST(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(v_3, r_4, s_5), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V2}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
})
}

func Call_Control_Monad_RWS_Trans_applyRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=Any
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_3 shape=App(Other) bindingType=Any
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): functorRWST1_3_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_3_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_5, r_6, s_7))
})})
_ = functorRWST1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_4 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_4
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_3_2)}
}), gopurs_runtime.Func4(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_6, r_8, s_9), gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_5 shape=Other bindingType=Any
__local_var_11_5 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V2
_ = __local_var_11_5
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_4.V0), __local_var_11_5, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_7, r_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V0))
}))
})}))}
})
}

func Call_Control_Monad_RWS_Trans_bindRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_2 shape=App(Other) bindingType=Any
Apply0_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_2_2
// TAST (Let): Functor0_3_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_2, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_3
// TAST (Let): __local_var_4_5 shape=App(Other) bindingType=Any
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_5
// TAST (Let): functorRWST1_4_4 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_4_4 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_5, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V0, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_9.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_6, r_7, s_8))
})})
_ = functorRWST1_4_4
// TAST (Let): applyRWST1__193435443_2_1 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
applyRWST1__193435443_2_1 := gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_6 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_6
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_4_4)}
}), gopurs_runtime.Func4(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_7, r_9, s_10), gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_7 shape=Other bindingType=Any
__local_var_12_7 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V2
_ = __local_var_12_7
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_3.V0), gopurs_runtime.Func(func(v3_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_13.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_13.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_6.V0), __local_var_12_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_13.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_8, r_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_8 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_4_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_8
// TAST (Let): applyRWST2_5_9 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_5_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_2_1, dictMonoid_3))
_ = applyRWST2_5_9
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_5_9)}
}), gopurs_runtime.Func4(func(v_6 gopurs_runtime.Value, f_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_6, r_8, s_9), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_10 shape=Other bindingType=Any
__local_var_11_10 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V2
_ = __local_var_11_10
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_8.V0), __local_var_11_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply3(f_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1, r_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0))
}))
})}))}
})
}

func Call_Control_Monad_RWS_Trans_semigroupRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): Apply0_1_1 shape=App(Other) bindingType=Any
Apply0_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_1
// TAST (Let): Functor0_2_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): __local_var_3_4 shape=App(Other) bindingType=Any
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorRWST1_3_3 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_3_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_5, r_6, s_7))
})})
_ = functorRWST1_3_3
// TAST (Let): applyRWST1__193435443_1_0 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
applyRWST1__193435443_1_0 := gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_5 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_5
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_3_3)}
}), gopurs_runtime.Func4(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_6, r_8, s_9), gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_6 shape=Other bindingType=Any
__local_var_11_6 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V2
_ = __local_var_11_6
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_5.V0), __local_var_11_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_12.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_7, r_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_3_7 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_1_0, dictMonoid_2))
_ = applyRWST2_3_7
return gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_5_8 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyRWST2_3_7.V0), gopurs_runtime.Value{}))
_ = Functor0_5_8
// TAST (Let): __local_var_6_9 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (TypeVar a))
__local_var_6_9 := gopurs_runtime.RecordGet(dictSemigroup_4, "append")
_ = __local_var_6_9
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_7 gopurs_runtime.Value, b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyRWST2_3_7.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_8.V0), __local_var_6_9, a_7), b_8)
})}))}
})
})
}

func Call_Control_Monad_RWS_Trans_applicativeRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): Apply0_3_3 shape=App(Other) bindingType=Any
Apply0_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_3_3
// TAST (Let): Functor0_4_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
// TAST (Let): __local_var_5_6 shape=App(Other) bindingType=Any
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_6
// TAST (Let): functorRWST1_5_5 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_5_5 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_6, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0, gopurs_runtime.Apply(f_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_7, r_8, s_9))
})})
_ = functorRWST1_5_5
// TAST (Let): applyRWST1__193435443_2_1 shape=Let(Let(Let(Let(Abs(Let(LitRecord)))))) bindingType=Any
applyRWST1__193435443_2_1 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_7 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_7
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_5_5)}
}), gopurs_runtime.Func4(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "bind"), gopurs_runtime.Apply2(v_8, r_10, s_11), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_8 shape=Other bindingType=Any
__local_var_13_8 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V2
_ = __local_var_13_8
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_7.V0), __local_var_13_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_9, r_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_4_9 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_4_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_2_1, dictMonoid_3))
_ = applyRWST2_4_9
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_4_9)}
}), gopurs_runtime.Func3(func(a_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_7, a_5, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")}))})
})}))}
})
}

func Call_Control_Monad_RWS_Trans_monadRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_1 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_1_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): Apply0_3_4 shape=App(Other) bindingType=Any
Apply0_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_3_4
// TAST (Let): Functor0_4_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_4, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
// TAST (Let): __local_var_5_7 shape=App(Other) bindingType=Any
__local_var_5_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_7
// TAST (Let): functorRWST1_5_6 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_5_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_7, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0, gopurs_runtime.Apply(f_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_7, r_8, s_9))
})})
_ = functorRWST1_5_6
// TAST (Let): applyRWST1__193435443_2_2 shape=Let(Let(Let(Let(Abs(Let(LitRecord)))))) bindingType=Any
applyRWST1__193435443_2_2 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_8 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_8
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_5_6)}
}), gopurs_runtime.Func4(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "bind"), gopurs_runtime.Apply2(v_8, r_10, s_11), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_9 shape=Other bindingType=Any
__local_var_13_9 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V2
_ = __local_var_13_9
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_8.V0), __local_var_13_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_9, r_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_2_2
// TAST (Let): applicativeRWST1__193435443_1_0 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
applicativeRWST1__193435443_1_0 := gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_4_10 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_4_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_2_2, dictMonoid_3))
_ = applyRWST2_4_10
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_4_10)}
}), gopurs_runtime.Func3(func(a_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_7, a_5, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")}))})
})}))}
})
_ = applicativeRWST1__193435443_1_0
// TAST (Let): __local_var_2_12 shape=App(Other) bindingType=Any
__local_var_2_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_12
// TAST (Let): Functor0_3_13 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_12, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_13
// TAST (Let): Apply0_4_15 shape=App(Other) bindingType=Any
Apply0_4_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_12, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_15
// TAST (Let): Functor0_5_16 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_15, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_16
// TAST (Let): __local_var_6_18 shape=App(Other) bindingType=Any
__local_var_6_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_15, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_18
// TAST (Let): functorRWST1_6_17 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_6_17 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_18, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V0, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_8, r_9, s_10))
})})
_ = functorRWST1_6_17
// TAST (Let): applyRWST1__193435443_4_14 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
applyRWST1__193435443_4_14 := gopurs_runtime.Func(func(dictMonoid_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_8_19 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_8_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_7, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_8_19
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_6_17)}
}), gopurs_runtime.Func4(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_12, "bind"), gopurs_runtime.Apply2(v_9, r_11, s_12), gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_20 shape=Other bindingType=Any
__local_var_14_20 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V2
_ = __local_var_14_20
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_16.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_19.V0), __local_var_14_20, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_10, r_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_4_14
// TAST (Let): bindRWST1__193435443_2_11 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
bindRWST1__193435443_2_11 := gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_21 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_6_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_21
// TAST (Let): applyRWST2_7_22 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_7_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_4_14, dictMonoid_5))
_ = applyRWST2_7_22
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_7_22)}
}), gopurs_runtime.Func4(func(v_8 gopurs_runtime.Value, f_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_12, "bind"), gopurs_runtime.Apply2(v_8, r_10, s_11), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_23 shape=Other bindingType=Any
__local_var_13_23 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V2
_ = __local_var_13_23
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_13.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_21.V0), __local_var_13_23, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply3(f_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V1, r_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V0))
}))
})}))}
})
_ = bindRWST1__193435443_2_11
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST2_4_24 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applicativeRWST2_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeRWST1__193435443_1_0, dictMonoid_3))
_ = applicativeRWST2_4_24
// TAST (Let): bindRWST2_5_25 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
bindRWST2_5_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindRWST1__193435443_2_11, dictMonoid_3))
_ = bindRWST2_5_25
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_4_24)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_5_25)}
})}))}
})
}

func Call_Control_Monad_RWS_Trans_monadAskRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar r), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar r), (TypeVar w)])]))
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): pure_2_3 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_2_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): Apply0_4_6 shape=App(Other) bindingType=Any
Apply0_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_6
// TAST (Let): Functor0_5_7 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_6, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_7
// TAST (Let): __local_var_6_9 shape=App(Other) bindingType=Any
__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_9
// TAST (Let): functorRWST1_6_8 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_6_8 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_9, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V0, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_8, r_9, s_10))
})})
_ = functorRWST1_6_8
// TAST (Let): applyRWST1__193435443_3_4 shape=Let(Let(Let(Let(Abs(Let(LitRecord)))))) bindingType=Any
applyRWST1__193435443_3_4 := gopurs_runtime.Func(func(dictMonoid_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_8_10 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_8_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_7, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_8_10
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_6_8)}
}), gopurs_runtime.Func4(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "bind"), gopurs_runtime.Apply2(v_9, r_11, s_12), gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_11 shape=Other bindingType=Any
__local_var_14_11 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V2
_ = __local_var_14_11
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_7.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_10.V0), __local_var_14_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_10, r_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_3_4
// TAST (Let): applicativeRWST1__193435443_2_2 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
applicativeRWST1__193435443_2_2 := gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_5_12 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_3_4, dictMonoid_4))
_ = applyRWST2_5_12
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_5_12)}
}), gopurs_runtime.Func3(func(a_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_3, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_8, a_6, gopurs_runtime.RecordGet(dictMonoid_4, "mempty")}))})
})}))}
})
_ = applicativeRWST1__193435443_2_2
// TAST (Let): __local_var_3_14 shape=App(Other) bindingType=Any
__local_var_3_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_14
// TAST (Let): Functor0_4_15 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_14, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_15
// TAST (Let): Apply0_5_17 shape=App(Other) bindingType=Any
Apply0_5_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_14, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_17
// TAST (Let): Functor0_6_18 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_17, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_18
// TAST (Let): __local_var_7_20 shape=App(Other) bindingType=Any
__local_var_7_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_17, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_20
// TAST (Let): functorRWST1_7_19 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_7_19 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_20, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V0, gopurs_runtime.Apply(f_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_9, r_10, s_11))
})})
_ = functorRWST1_7_19
// TAST (Let): applyRWST1__193435443_5_16 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
applyRWST1__193435443_5_16 := gopurs_runtime.Func(func(dictMonoid_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_21 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_9_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_8, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_21
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_7_19)}
}), gopurs_runtime.Func4(func(v_10 gopurs_runtime.Value, v1_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_14, "bind"), gopurs_runtime.Apply2(v_10, r_12, s_13), gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_22 shape=Other bindingType=Any
__local_var_15_22 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V2
_ = __local_var_15_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_18.V0), gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_21.V0), __local_var_15_22, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_11, r_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_5_16
// TAST (Let): bindRWST1__193435443_3_13 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
bindRWST1__193435443_3_13 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_23 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_7_23 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_23
// TAST (Let): applyRWST2_8_24 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_8_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_5_16, dictMonoid_6))
_ = applyRWST2_8_24
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_8_24)}
}), gopurs_runtime.Func4(func(v_9 gopurs_runtime.Value, f_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_14, "bind"), gopurs_runtime.Apply2(v_9, r_11, s_12), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_25 shape=Other bindingType=Any
__local_var_14_25 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V2
_ = __local_var_14_25
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_15.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_23.V0), __local_var_14_25, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply3(f_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V1, r_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V0))
}))
})}))}
})
_ = bindRWST1__193435443_3_13
// TAST (Let): monadRWST1__193435443_2_1 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadRWST1__193435443_2_1 := gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST2_5_26 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applicativeRWST2_5_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeRWST1__193435443_2_2, dictMonoid_4))
_ = applicativeRWST2_5_26
// TAST (Let): bindRWST2_6_27 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
bindRWST2_6_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindRWST1__193435443_3_13, dictMonoid_4))
_ = bindRWST2_6_27
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_5_26)}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_6_27)}
})}))}
})
_ = monadRWST1__193435443_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadRWST2_4_28 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadRWST2_4_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadRWST1__193435443_2_1, dictMonoid_3))
_ = monadRWST2_4_28
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_4_28)}
}), gopurs_runtime.Func2(func(r_5 gopurs_runtime.Value, s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_6, r_5, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")}))})
})}))}
})
}

func Call_Control_Monad_RWS_Trans_monadReaderRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_1 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar r), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar r), (TypeVar w)])]))
pure_1_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_1
// TAST (Let): pure_2_4 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_2_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_4
// TAST (Let): __local_var_3_6 shape=App(Other) bindingType=Any
__local_var_3_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_6
// TAST (Let): Apply0_4_7 shape=App(Other) bindingType=Any
Apply0_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_6, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_7
// TAST (Let): Functor0_5_8 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_7, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_8
// TAST (Let): __local_var_6_10 shape=App(Other) bindingType=Any
__local_var_6_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_7, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_10
// TAST (Let): functorRWST1_6_9 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_6_9 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_10, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V0, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_8, r_9, s_10))
})})
_ = functorRWST1_6_9
// TAST (Let): applyRWST1__193435443_3_5 shape=Let(Let(Let(Let(Abs(Let(LitRecord)))))) bindingType=Any
applyRWST1__193435443_3_5 := gopurs_runtime.Func(func(dictMonoid_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_8_11 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_8_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_7, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_8_11
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_6_9)}
}), gopurs_runtime.Func4(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_6, "bind"), gopurs_runtime.Apply2(v_9, r_11, s_12), gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_12 shape=Other bindingType=Any
__local_var_14_12 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V2
_ = __local_var_14_12
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_8.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_11.V0), __local_var_14_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_10, r_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_3_5
// TAST (Let): applicativeRWST1__193435443_2_3 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
applicativeRWST1__193435443_2_3 := gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_5_13 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_5_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_3_5, dictMonoid_4))
_ = applyRWST2_5_13
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_5_13)}
}), gopurs_runtime.Func3(func(a_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_4, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_8, a_6, gopurs_runtime.RecordGet(dictMonoid_4, "mempty")}))})
})}))}
})
_ = applicativeRWST1__193435443_2_3
// TAST (Let): __local_var_3_15 shape=App(Other) bindingType=Any
__local_var_3_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_15
// TAST (Let): Functor0_4_16 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_15, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_16
// TAST (Let): Apply0_5_18 shape=App(Other) bindingType=Any
Apply0_5_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_15, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_18
// TAST (Let): Functor0_6_19 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_18, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_19
// TAST (Let): __local_var_7_21 shape=App(Other) bindingType=Any
__local_var_7_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_18, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_21
// TAST (Let): functorRWST1_7_20 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_7_20 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_21, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V0, gopurs_runtime.Apply(f_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_9, r_10, s_11))
})})
_ = functorRWST1_7_20
// TAST (Let): applyRWST1__193435443_5_17 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
applyRWST1__193435443_5_17 := gopurs_runtime.Func(func(dictMonoid_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_22 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_9_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_8, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_22
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_7_20)}
}), gopurs_runtime.Func4(func(v_10 gopurs_runtime.Value, v1_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_15, "bind"), gopurs_runtime.Apply2(v_10, r_12, s_13), gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_23 shape=Other bindingType=Any
__local_var_15_23 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V2
_ = __local_var_15_23
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_19.V0), gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_22.V0), __local_var_15_23, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_11, r_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_5_17
// TAST (Let): bindRWST1__193435443_3_14 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
bindRWST1__193435443_3_14 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_24 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_7_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_24
// TAST (Let): applyRWST2_8_25 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_8_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_5_17, dictMonoid_6))
_ = applyRWST2_8_25
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_8_25)}
}), gopurs_runtime.Func4(func(v_9 gopurs_runtime.Value, f_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_15, "bind"), gopurs_runtime.Apply2(v_9, r_11, s_12), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_26 shape=Other bindingType=Any
__local_var_14_26 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V2
_ = __local_var_14_26
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_16.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_24.V0), __local_var_14_26, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply3(f_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V1, r_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V0))
}))
})}))}
})
_ = bindRWST1__193435443_3_14
// TAST (Let): monadRWST1__193435443_2_2 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadRWST1__193435443_2_2 := gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST2_5_27 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applicativeRWST2_5_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeRWST1__193435443_2_3, dictMonoid_4))
_ = applicativeRWST2_5_27
// TAST (Let): bindRWST2_6_28 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
bindRWST2_6_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindRWST1__193435443_3_14, dictMonoid_4))
_ = bindRWST2_6_28
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_5_27)}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_6_28)}
})}))}
})
_ = monadRWST1__193435443_2_2
// TAST (Let): monadAskRWST1__193435443_1_0 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
monadAskRWST1__193435443_1_0 := gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadRWST2_4_29 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadRWST2_4_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadRWST1__193435443_2_2, dictMonoid_3))
_ = monadRWST2_4_29
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_4_29)}
}), gopurs_runtime.Func2(func(r_5 gopurs_runtime.Value, s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_6, r_5, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")}))})
})}))}
})
_ = monadAskRWST1__193435443_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadAskRWST2_3_30 shape=App(Other) bindingType=(ADT ["Control","Monad","Reader","Class","MonadAsk"] [(TypeVar r), (Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadAskRWST2_3_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(monadAskRWST1__193435443_1_0, dictMonoid_2))
_ = monadAskRWST2_3_30
return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(monadAskRWST2_3_30)}
}), gopurs_runtime.Func4(func(f_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(m_5, gopurs_runtime.Apply(f_4, r_6), s_7)
})}))}
})
}

func Call_Control_Monad_RWS_Trans_monadEffectRWS(dictMonoid_0_loop gopurs_runtime.Value, dictMonadEffect_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
var dictMonadEffect_1 gopurs_runtime.Value = dictMonadEffect_1_loop
_ = dictMonadEffect_1
// TAST (Let): Monad0_2_0 shape=App(Other) bindingType=Any
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): pure_3_3 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_3_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_3
// TAST (Let): __local_var_4_5 shape=App(Other) bindingType=Any
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_4_5
// TAST (Let): Apply0_5_6 shape=App(Other) bindingType=Any
Apply0_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_6
// TAST (Let): Functor0_6_7 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_6, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_7
// TAST (Let): __local_var_7_9 shape=App(Other) bindingType=Any
__local_var_7_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_9
// TAST (Let): functorRWST1_7_8 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_7_8 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_9, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V0, gopurs_runtime.Apply(f_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_9, r_10, s_11))
})})
_ = functorRWST1_7_8
// TAST (Let): Semigroup0_8_10 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_8_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_8_10
// TAST (Let): applyRWST2_4_4 shape=Let(Let(Let(Let(Let(LitRecord))))) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_4_4 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_7_8)}
}), gopurs_runtime.Func4(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_5, "bind"), gopurs_runtime.Apply2(v_9, r_11, s_12), gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_11 shape=Other bindingType=Any
__local_var_14_11 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V2
_ = __local_var_14_11
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_7.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_10.V0), __local_var_14_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_10, r_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0))
}))
})})
_ = applyRWST2_4_4
// TAST (Let): applicativeRWST2_3_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applicativeRWST2_3_2 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_4_4)}
}), gopurs_runtime.Func3(func(a_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_3, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_7, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}))})
})})
_ = applicativeRWST2_3_2
// TAST (Let): __local_var_4_13 shape=App(Other) bindingType=Any
__local_var_4_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_4_13
// TAST (Let): Functor0_5_14 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_14
// TAST (Let): Semigroup0_6_15 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_6_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_15
// TAST (Let): Apply0_7_17 shape=App(Other) bindingType=Any
Apply0_7_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_17
// TAST (Let): Functor0_8_18 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_8_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_17, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_18
// TAST (Let): __local_var_9_20 shape=App(Other) bindingType=Any
__local_var_9_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_17, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_20
// TAST (Let): functorRWST1_9_19 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_9_19 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_10 gopurs_runtime.Value, v_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_20, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V0, gopurs_runtime.Apply(f_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_11, r_12, s_13))
})})
_ = functorRWST1_9_19
// TAST (Let): Semigroup0_10_21 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_10_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_10_21
// TAST (Let): applyRWST2_7_16 shape=Let(Let(Let(Let(LitRecord)))) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_7_16 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_9_19)}
}), gopurs_runtime.Func4(func(v_11 gopurs_runtime.Value, v1_12 gopurs_runtime.Value, r_13 gopurs_runtime.Value, s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_13, "bind"), gopurs_runtime.Apply2(v_11, r_13, s_14), gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_22 shape=Other bindingType=Any
__local_var_16_22 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V2
_ = __local_var_16_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_18.V0), gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_21.V0), __local_var_16_22, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_12, r_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V0))
}))
})})
_ = applyRWST2_7_16
// TAST (Let): bindRWST2_4_12 shape=Let(Let(Let(Let(LitRecord)))) bindingType=(ADT ["Control","Bind","Bind"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
bindRWST2_4_12 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_7_16)}
}), gopurs_runtime.Func4(func(v_8 gopurs_runtime.Value, f_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_13, "bind"), gopurs_runtime.Apply2(v_8, r_10, s_11), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_23 shape=Other bindingType=Any
__local_var_13_23 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V2
_ = __local_var_13_23
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_14.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_15.V0), __local_var_13_23, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply3(f_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V1, r_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V0))
}))
})})
_ = bindRWST2_4_12
// TAST (Let): monadRWST1_3_1 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadRWST1_3_1 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_3_2)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_4_12)}
})})
_ = monadRWST1_3_1
// TAST (Let): Bind1_4_24 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_24
// TAST (Let): pure_5_25 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_5_25 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_25
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer((&Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST1_3_1)}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_26 shape=App(Other) bindingType=(TypeVar c)
__local_var_7_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_1, "liftEffect"), x_6)
_ = __local_var_7_26
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_24.V1), __local_var_7_26, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_25, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_9, a_10, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}))})
}))
})
})}))}
}

func Call_Control_Monad_RWS_Trans_monadRecRWST(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Applicative0_3_2 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_2
// TAST (Let): pure_4_5 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_4_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_5
// TAST (Let): __local_var_5_7 shape=App(Other) bindingType=Any
__local_var_5_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_5_7
// TAST (Let): Apply0_6_8 shape=App(Other) bindingType=Any
Apply0_6_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_7, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_8
// TAST (Let): Functor0_7_9 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_8, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_9
// TAST (Let): __local_var_8_11 shape=App(Other) bindingType=Any
__local_var_8_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_8, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_11
// TAST (Let): functorRWST1_8_10 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_8_10 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_9 gopurs_runtime.Value, v_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_11, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V0, gopurs_runtime.Apply(f_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_10, r_11, s_12))
})})
_ = functorRWST1_8_10
// TAST (Let): applyRWST1__193435443_5_6 shape=Let(Let(Let(Let(Abs(Let(LitRecord)))))) bindingType=Any
applyRWST1__193435443_5_6 := gopurs_runtime.Func(func(dictMonoid_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_10_12 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_10_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_9, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_10_12
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_8_10)}
}), gopurs_runtime.Func4(func(v_11 gopurs_runtime.Value, v1_12 gopurs_runtime.Value, r_13 gopurs_runtime.Value, s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_7, "bind"), gopurs_runtime.Apply2(v_11, r_13, s_14), gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_13 shape=Other bindingType=Any
__local_var_16_13 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V2
_ = __local_var_16_13
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_9.V0), gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_12.V0), __local_var_16_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_12, r_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_5_6
// TAST (Let): applicativeRWST1__193435443_4_4 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
applicativeRWST1__193435443_4_4 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_7_14 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_7_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_5_6, dictMonoid_6))
_ = applyRWST2_7_14
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_7_14)}
}), gopurs_runtime.Func3(func(a_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_5, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_10, a_8, gopurs_runtime.RecordGet(dictMonoid_6, "mempty")}))})
})}))}
})
_ = applicativeRWST1__193435443_4_4
// TAST (Let): __local_var_5_16 shape=App(Other) bindingType=Any
__local_var_5_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_5_16
// TAST (Let): Functor0_6_17 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_16, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_17
// TAST (Let): Apply0_7_19 shape=App(Other) bindingType=Any
Apply0_7_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_16, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_19
// TAST (Let): Functor0_8_20 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_8_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_19, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_20
// TAST (Let): __local_var_9_22 shape=App(Other) bindingType=Any
__local_var_9_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_19, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_22
// TAST (Let): functorRWST1_9_21 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_9_21 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_10 gopurs_runtime.Value, v_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_22, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V0, gopurs_runtime.Apply(f_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_11, r_12, s_13))
})})
_ = functorRWST1_9_21
// TAST (Let): applyRWST1__193435443_7_18 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
applyRWST1__193435443_7_18 := gopurs_runtime.Func(func(dictMonoid_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_11_23 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_11_23 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_10, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_11_23
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_9_21)}
}), gopurs_runtime.Func4(func(v_12 gopurs_runtime.Value, v1_13 gopurs_runtime.Value, r_14 gopurs_runtime.Value, s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_16, "bind"), gopurs_runtime.Apply2(v_12, r_14, s_15), gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_24 shape=Other bindingType=Any
__local_var_17_24 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_16.UnsafePtr).V2
_ = __local_var_17_24
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_20.V0), gopurs_runtime.Func(func(v3_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_18.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_16.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_18.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_11_23.V0), __local_var_17_24, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_18.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_13, r_14, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_16.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_7_18
// TAST (Let): bindRWST1__193435443_5_15 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
bindRWST1__193435443_5_15 := gopurs_runtime.Func(func(dictMonoid_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_25 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_9_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_8, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_25
// TAST (Let): applyRWST2_10_26 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_10_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_7_18, dictMonoid_8))
_ = applyRWST2_10_26
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_10_26)}
}), gopurs_runtime.Func4(func(v_11 gopurs_runtime.Value, f_12 gopurs_runtime.Value, r_13 gopurs_runtime.Value, s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_16, "bind"), gopurs_runtime.Apply2(v_11, r_13, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_27 shape=Other bindingType=Any
__local_var_16_27 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_15.UnsafePtr).V2
_ = __local_var_16_27
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_17.V0), gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_25.V0), __local_var_16_27, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply3(f_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_15.UnsafePtr).V1, r_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_15.UnsafePtr).V0))
}))
})}))}
})
_ = bindRWST1__193435443_5_15
// TAST (Let): monadRWST1__193435443_4_3 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadRWST1__193435443_4_3 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST2_7_28 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applicativeRWST2_7_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeRWST1__193435443_4_4, dictMonoid_6))
_ = applicativeRWST2_7_28
// TAST (Let): bindRWST2_8_29 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
bindRWST2_8_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindRWST1__193435443_5_15, dictMonoid_6))
_ = bindRWST2_8_29
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_7_28)}
}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_8_29)}
})}))}
})
_ = monadRWST1__193435443_4_3
return gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_30 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_6_30 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_30
// TAST (Let): monadRWST2_7_31 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadRWST2_7_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadRWST1__193435443_4_3, dictMonoid_5))
_ = monadRWST2_7_31
return gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_7_31)}
}), gopurs_runtime.Func4(func(k_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_32 shape=Other bindingType=Any
__local_var_13_32 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V2
_ = __local_var_13_32
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply3(k_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V1, r_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0), gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
var __t_tag_33 gopurs_runtime.Value = (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1
if (__t_tag_33.Type == 9 && __t_tag_33.IntVal == 525585346) {
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_RWS_Trans_2567023824_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, (&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_30.V0), __local_var_13_32, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V2)})})))}
goto end_branch_35
} else {

}
}
{
var __t_tag_34 gopurs_runtime.Value = (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1
if (__t_tag_34.Type == 9 && __t_tag_34.IntVal == 60402430) {
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_RWS_Trans_4258103980_3603546092((&Constructor_Control_Monad_Rec_Class_Done[*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, (&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_30.V0), __local_var_13_32, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V2)})})))}
goto end_branch_35
} else {

}
}
{
__t35 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_35:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), __t35)
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_11, a_9, gopurs_runtime.RecordGet(dictMonoid_5, "mempty")}))})
})}))}
})
}

func Call_Control_Monad_RWS_Trans_monadStateRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): pure_2_3 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_2_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): Apply0_4_6 shape=App(Other) bindingType=Any
Apply0_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_6
// TAST (Let): Functor0_5_7 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_6, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_7
// TAST (Let): __local_var_6_9 shape=App(Other) bindingType=Any
__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_9
// TAST (Let): functorRWST1_6_8 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_6_8 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_9, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V0, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_8, r_9, s_10))
})})
_ = functorRWST1_6_8
// TAST (Let): applyRWST1__193435443_3_4 shape=Let(Let(Let(Let(Abs(Let(LitRecord)))))) bindingType=Any
applyRWST1__193435443_3_4 := gopurs_runtime.Func(func(dictMonoid_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_8_10 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_8_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_7, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_8_10
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_6_8)}
}), gopurs_runtime.Func4(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "bind"), gopurs_runtime.Apply2(v_9, r_11, s_12), gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_11 shape=Other bindingType=Any
__local_var_14_11 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V2
_ = __local_var_14_11
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_7.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_10.V0), __local_var_14_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_10, r_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_3_4
// TAST (Let): applicativeRWST1__193435443_2_2 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
applicativeRWST1__193435443_2_2 := gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_5_12 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_3_4, dictMonoid_4))
_ = applyRWST2_5_12
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_5_12)}
}), gopurs_runtime.Func3(func(a_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_3, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_8, a_6, gopurs_runtime.RecordGet(dictMonoid_4, "mempty")}))})
})}))}
})
_ = applicativeRWST1__193435443_2_2
// TAST (Let): __local_var_3_14 shape=App(Other) bindingType=Any
__local_var_3_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_14
// TAST (Let): Functor0_4_15 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_14, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_15
// TAST (Let): Apply0_5_17 shape=App(Other) bindingType=Any
Apply0_5_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_14, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_17
// TAST (Let): Functor0_6_18 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_17, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_18
// TAST (Let): __local_var_7_20 shape=App(Other) bindingType=Any
__local_var_7_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_17, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_20
// TAST (Let): functorRWST1_7_19 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_7_19 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_20, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V0, gopurs_runtime.Apply(f_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_9, r_10, s_11))
})})
_ = functorRWST1_7_19
// TAST (Let): applyRWST1__193435443_5_16 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
applyRWST1__193435443_5_16 := gopurs_runtime.Func(func(dictMonoid_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_21 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_9_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_8, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_21
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_7_19)}
}), gopurs_runtime.Func4(func(v_10 gopurs_runtime.Value, v1_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_14, "bind"), gopurs_runtime.Apply2(v_10, r_12, s_13), gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_22 shape=Other bindingType=Any
__local_var_15_22 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V2
_ = __local_var_15_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_18.V0), gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_21.V0), __local_var_15_22, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_11, r_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_5_16
// TAST (Let): bindRWST1__193435443_3_13 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
bindRWST1__193435443_3_13 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_23 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_7_23 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_23
// TAST (Let): applyRWST2_8_24 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_8_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_5_16, dictMonoid_6))
_ = applyRWST2_8_24
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_8_24)}
}), gopurs_runtime.Func4(func(v_9 gopurs_runtime.Value, f_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_14, "bind"), gopurs_runtime.Apply2(v_9, r_11, s_12), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_25 shape=Other bindingType=Any
__local_var_14_25 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V2
_ = __local_var_14_25
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_15.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_23.V0), __local_var_14_25, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply3(f_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V1, r_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V0))
}))
})}))}
})
_ = bindRWST1__193435443_3_13
// TAST (Let): monadRWST1__193435443_2_1 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadRWST1__193435443_2_1 := gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST2_5_26 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applicativeRWST2_5_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeRWST1__193435443_2_2, dictMonoid_4))
_ = applicativeRWST2_5_26
// TAST (Let): bindRWST2_6_27 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
bindRWST2_6_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindRWST1__193435443_3_13, dictMonoid_4))
_ = bindRWST2_6_27
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_5_26)}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_6_27)}
})}))}
})
_ = monadRWST1__193435443_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadRWST2_4_28 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadRWST2_4_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadRWST1__193435443_2_1, dictMonoid_3))
_ = monadRWST2_4_28
return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_State_Class_MonadState[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_4_28)}
}), gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_8_29 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar s)])
v1_8_29 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_5, s_7))
_ = v1_8_29
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (v1_8_29).V1, (v1_8_29).V0, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")}))})
})}))}
})
}

func Call_Control_Monad_RWS_Trans_monadTellRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), Unit, (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), Unit, (TypeVar w)])]))
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): pure_2_3 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_2_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_3
// TAST (Let): __local_var_3_5 shape=App(Other) bindingType=Any
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): Apply0_4_6 shape=App(Other) bindingType=Any
Apply0_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_6
// TAST (Let): Functor0_5_7 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_6, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_7
// TAST (Let): __local_var_6_9 shape=App(Other) bindingType=Any
__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_9
// TAST (Let): functorRWST1_6_8 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_6_8 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_9, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V0, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_11.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_8, r_9, s_10))
})})
_ = functorRWST1_6_8
// TAST (Let): applyRWST1__193435443_3_4 shape=Let(Let(Let(Let(Abs(Let(LitRecord)))))) bindingType=Any
applyRWST1__193435443_3_4 := gopurs_runtime.Func(func(dictMonoid_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_8_10 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_8_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_7, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_8_10
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_6_8)}
}), gopurs_runtime.Func4(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "bind"), gopurs_runtime.Apply2(v_9, r_11, s_12), gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_11 shape=Other bindingType=Any
__local_var_14_11 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V2
_ = __local_var_14_11
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_7.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_10.V0), __local_var_14_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_10, r_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_3_4
// TAST (Let): applicativeRWST1__193435443_2_2 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
applicativeRWST1__193435443_2_2 := gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_5_12 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_3_4, dictMonoid_4))
_ = applyRWST2_5_12
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_5_12)}
}), gopurs_runtime.Func3(func(a_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_3, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_8, a_6, gopurs_runtime.RecordGet(dictMonoid_4, "mempty")}))})
})}))}
})
_ = applicativeRWST1__193435443_2_2
// TAST (Let): __local_var_3_14 shape=App(Other) bindingType=Any
__local_var_3_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_14
// TAST (Let): Functor0_4_15 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_14, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_15
// TAST (Let): Apply0_5_17 shape=App(Other) bindingType=Any
Apply0_5_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_14, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_17
// TAST (Let): Functor0_6_18 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_17, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_18
// TAST (Let): __local_var_7_20 shape=App(Other) bindingType=Any
__local_var_7_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_17, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_20
// TAST (Let): functorRWST1_7_19 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_7_19 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_20, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V0, gopurs_runtime.Apply(f_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_9, r_10, s_11))
})})
_ = functorRWST1_7_19
// TAST (Let): applyRWST1__193435443_5_16 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
applyRWST1__193435443_5_16 := gopurs_runtime.Func(func(dictMonoid_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_21 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_9_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_8, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_21
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_7_19)}
}), gopurs_runtime.Func4(func(v_10 gopurs_runtime.Value, v1_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_14, "bind"), gopurs_runtime.Apply2(v_10, r_12, s_13), gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_22 shape=Other bindingType=Any
__local_var_15_22 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V2
_ = __local_var_15_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_18.V0), gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_21.V0), __local_var_15_22, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_11, r_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_5_16
// TAST (Let): bindRWST1__193435443_3_13 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
bindRWST1__193435443_3_13 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_23 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_7_23 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_23
// TAST (Let): applyRWST2_8_24 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_8_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_5_16, dictMonoid_6))
_ = applyRWST2_8_24
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_8_24)}
}), gopurs_runtime.Func4(func(v_9 gopurs_runtime.Value, f_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_14, "bind"), gopurs_runtime.Apply2(v_9, r_11, s_12), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_25 shape=Other bindingType=Any
__local_var_14_25 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V2
_ = __local_var_14_25
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_15.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_23.V0), __local_var_14_25, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply3(f_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V1, r_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V0))
}))
})}))}
})
_ = bindRWST1__193435443_3_13
// TAST (Let): monadRWST1__193435443_2_1 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadRWST1__193435443_2_1 := gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST2_5_26 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applicativeRWST2_5_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeRWST1__193435443_2_2, dictMonoid_4))
_ = applicativeRWST2_5_26
// TAST (Let): bindRWST2_6_27 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
bindRWST2_6_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindRWST1__193435443_3_13, dictMonoid_4))
_ = bindRWST2_6_27
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_5_26)}
}), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_6_27)}
})}))}
})
_ = monadRWST1__193435443_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_28 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_4_28 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_28
// TAST (Let): monadRWST2_5_29 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadRWST2_5_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadRWST1__193435443_2_1, dictMonoid_3))
_ = monadRWST2_5_29
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_5_29)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_4_28)}
}), gopurs_runtime.Func3(func(w_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_8, Get_Data_Unit_unit(), w_6}))})
})}))}
})
}

func Call_Control_Monad_RWS_Trans_monadWriterRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=Any
Applicative0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_2_1
// TAST (Let): pure_3_3 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), Unit, (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), Unit, (TypeVar w)])]))
pure_3_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_3
// TAST (Let): pure_4_6 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_4_6 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_6
// TAST (Let): __local_var_5_8 shape=App(Other) bindingType=Any
__local_var_5_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_5_8
// TAST (Let): Apply0_6_9 shape=App(Other) bindingType=Any
Apply0_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_8, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_9
// TAST (Let): Functor0_7_10 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_9, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_10
// TAST (Let): __local_var_8_12 shape=App(Other) bindingType=Any
__local_var_8_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_9, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_12
// TAST (Let): functorRWST1_8_11 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_8_11 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_9 gopurs_runtime.Value, v_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_12, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V0, gopurs_runtime.Apply(f_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_10, r_11, s_12))
})})
_ = functorRWST1_8_11
// TAST (Let): applyRWST1__193435443_5_7 shape=Let(Let(Let(Let(Abs(Let(LitRecord)))))) bindingType=Any
applyRWST1__193435443_5_7 := gopurs_runtime.Func(func(dictMonoid_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_10_13 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_10_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_9, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_10_13
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_8_11)}
}), gopurs_runtime.Func4(func(v_11 gopurs_runtime.Value, v1_12 gopurs_runtime.Value, r_13 gopurs_runtime.Value, s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_8, "bind"), gopurs_runtime.Apply2(v_11, r_13, s_14), gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_14 shape=Other bindingType=Any
__local_var_16_14 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V2
_ = __local_var_16_14
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_10.V0), gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_13.V0), __local_var_16_14, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_12, r_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_5_7
// TAST (Let): applicativeRWST1__193435443_4_5 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
applicativeRWST1__193435443_4_5 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_7_15 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_7_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_5_7, dictMonoid_6))
_ = applyRWST2_7_15
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_7_15)}
}), gopurs_runtime.Func3(func(a_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_6, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_10, a_8, gopurs_runtime.RecordGet(dictMonoid_6, "mempty")}))})
})}))}
})
_ = applicativeRWST1__193435443_4_5
// TAST (Let): __local_var_5_17 shape=App(Other) bindingType=Any
__local_var_5_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_5_17
// TAST (Let): Functor0_6_18 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_17, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_18
// TAST (Let): Apply0_7_20 shape=App(Other) bindingType=Any
Apply0_7_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_17, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_20
// TAST (Let): Functor0_8_21 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_8_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_20, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_21
// TAST (Let): __local_var_9_23 shape=App(Other) bindingType=Any
__local_var_9_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_20, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_23
// TAST (Let): functorRWST1_9_22 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_9_22 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_10 gopurs_runtime.Value, v_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_23, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V0, gopurs_runtime.Apply(f_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_11, r_12, s_13))
})})
_ = functorRWST1_9_22
// TAST (Let): applyRWST1__193435443_7_19 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
applyRWST1__193435443_7_19 := gopurs_runtime.Func(func(dictMonoid_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_11_24 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_11_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_10, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_11_24
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_9_22)}
}), gopurs_runtime.Func4(func(v_12 gopurs_runtime.Value, v1_13 gopurs_runtime.Value, r_14 gopurs_runtime.Value, s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_17, "bind"), gopurs_runtime.Apply2(v_12, r_14, s_15), gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_25 shape=Other bindingType=Any
__local_var_17_25 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_16.UnsafePtr).V2
_ = __local_var_17_25
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_21.V0), gopurs_runtime.Func(func(v3_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_18.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_16.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_18.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_11_24.V0), __local_var_17_25, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_18.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_13, r_14, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_16.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_7_19
// TAST (Let): bindRWST1__193435443_5_16 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
bindRWST1__193435443_5_16 := gopurs_runtime.Func(func(dictMonoid_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_26 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_9_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_8, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_26
// TAST (Let): applyRWST2_10_27 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_10_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_7_19, dictMonoid_8))
_ = applyRWST2_10_27
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_10_27)}
}), gopurs_runtime.Func4(func(v_11 gopurs_runtime.Value, f_12 gopurs_runtime.Value, r_13 gopurs_runtime.Value, s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_17, "bind"), gopurs_runtime.Apply2(v_11, r_13, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_28 shape=Other bindingType=Any
__local_var_16_28 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_15.UnsafePtr).V2
_ = __local_var_16_28
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_18.V0), gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_26.V0), __local_var_16_28, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply3(f_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_15.UnsafePtr).V1, r_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_15.UnsafePtr).V0))
}))
})}))}
})
_ = bindRWST1__193435443_5_16
// TAST (Let): monadRWST1__193435443_4_4 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadRWST1__193435443_4_4 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST2_7_29 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applicativeRWST2_7_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeRWST1__193435443_4_5, dictMonoid_6))
_ = applicativeRWST2_7_29
// TAST (Let): bindRWST2_8_30 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
bindRWST2_8_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindRWST1__193435443_5_16, dictMonoid_6))
_ = bindRWST2_8_30
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_7_29)}
}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_8_30)}
})}))}
})
_ = monadRWST1__193435443_4_4
// TAST (Let): monadTellRWST1__193435443_3_2 shape=Let(Let(Abs(Let(Let(LitRecord))))) bindingType=Any
monadTellRWST1__193435443_3_2 := gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_31 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_6_31 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_31
// TAST (Let): monadRWST2_7_32 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadRWST2_7_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadRWST1__193435443_4_4, dictMonoid_5))
_ = monadRWST2_7_32
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_7_32)}
}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_6_31)}
}), gopurs_runtime.Func3(func(w_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_3, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_10, Get_Data_Unit_unit(), w_8}))})
})}))}
})
_ = monadTellRWST1__193435443_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadTellRWST2_5_33 shape=App(Other) bindingType=(ADT ["Control","Monad","Writer","Class","MonadTell"] [(TypeVar w), (Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadTellRWST2_5_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(monadTellRWST1__193435443_3_2, dictMonoid_4))
_ = monadTellRWST2_5_33
return gopurs_runtime.Value{Type: 9, IntVal: 784743459, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(monadTellRWST2_5_33)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}
}), gopurs_runtime.Func3(func(m_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(m_6, r_7, s_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Rebox_Control_Monad_RWS_Trans_1947641166_2326785955((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0, (&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V2}), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V2})))})
}))
}), gopurs_runtime.Func3(func(m_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(m_6, r_7, s_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V2)}))})
}))
})}))}
})
}

func Call_Control_Monad_RWS_Trans_monadThrowRWST(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeVar m)])
Monad0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_1_0
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): pure_3_4 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_3_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_4
// TAST (Let): __local_var_4_6 shape=App(Other) bindingType=Any
__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_4_6
// TAST (Let): Apply0_5_7 shape=App(Other) bindingType=Any
Apply0_5_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_6, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_7
// TAST (Let): Functor0_6_8 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_7, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_8
// TAST (Let): __local_var_7_10 shape=App(Other) bindingType=Any
__local_var_7_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_7, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_10
// TAST (Let): functorRWST1_7_9 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_7_9 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_10, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V0, gopurs_runtime.Apply(f_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_9, r_10, s_11))
})})
_ = functorRWST1_7_9
// TAST (Let): applyRWST1__193435443_4_5 shape=Let(Let(Let(Let(Abs(Let(LitRecord)))))) bindingType=Any
applyRWST1__193435443_4_5 := gopurs_runtime.Func(func(dictMonoid_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_11 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_9_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_8, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_11
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_7_9)}
}), gopurs_runtime.Func4(func(v_10 gopurs_runtime.Value, v1_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_6, "bind"), gopurs_runtime.Apply2(v_10, r_12, s_13), gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_12 shape=Other bindingType=Any
__local_var_15_12 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V2
_ = __local_var_15_12
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_8.V0), gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_11.V0), __local_var_15_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_11, r_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_4_5
// TAST (Let): applicativeRWST1__193435443_3_3 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
applicativeRWST1__193435443_3_3 := gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_6_13 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_6_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_4_5, dictMonoid_5))
_ = applyRWST2_6_13
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_6_13)}
}), gopurs_runtime.Func3(func(a_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_4, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_9, a_7, gopurs_runtime.RecordGet(dictMonoid_5, "mempty")}))})
})}))}
})
_ = applicativeRWST1__193435443_3_3
// TAST (Let): __local_var_4_15 shape=App(Other) bindingType=Any
__local_var_4_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_4_15
// TAST (Let): Functor0_5_16 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_15, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_16
// TAST (Let): Apply0_6_18 shape=App(Other) bindingType=Any
Apply0_6_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_15, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_18
// TAST (Let): Functor0_7_19 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_18, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_19
// TAST (Let): __local_var_8_21 shape=App(Other) bindingType=Any
__local_var_8_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_18, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_21
// TAST (Let): functorRWST1_8_20 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_8_20 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_9 gopurs_runtime.Value, v_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_21, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V0, gopurs_runtime.Apply(f_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_10, r_11, s_12))
})})
_ = functorRWST1_8_20
// TAST (Let): applyRWST1__193435443_6_17 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
applyRWST1__193435443_6_17 := gopurs_runtime.Func(func(dictMonoid_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_10_22 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_10_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_9, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_10_22
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_8_20)}
}), gopurs_runtime.Func4(func(v_11 gopurs_runtime.Value, v1_12 gopurs_runtime.Value, r_13 gopurs_runtime.Value, s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_15, "bind"), gopurs_runtime.Apply2(v_11, r_13, s_14), gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_23 shape=Other bindingType=Any
__local_var_16_23 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V2
_ = __local_var_16_23
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_19.V0), gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_22.V0), __local_var_16_23, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_12, r_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_6_17
// TAST (Let): bindRWST1__193435443_4_14 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
bindRWST1__193435443_4_14 := gopurs_runtime.Func(func(dictMonoid_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_8_24 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_8_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_7, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_8_24
// TAST (Let): applyRWST2_9_25 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_9_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_6_17, dictMonoid_7))
_ = applyRWST2_9_25
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_9_25)}
}), gopurs_runtime.Func4(func(v_10 gopurs_runtime.Value, f_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_15, "bind"), gopurs_runtime.Apply2(v_10, r_12, s_13), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_26 shape=Other bindingType=Any
__local_var_15_26 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V2
_ = __local_var_15_26
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_16.V0), gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_24.V0), __local_var_15_26, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply3(f_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V1, r_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V0))
}))
})}))}
})
_ = bindRWST1__193435443_4_14
// TAST (Let): monadRWST1__193435443_2_1 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
monadRWST1__193435443_2_1 := gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST2_6_27 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applicativeRWST2_6_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeRWST1__193435443_3_3, dictMonoid_5))
_ = applicativeRWST2_6_27
// TAST (Let): bindRWST2_7_28 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
bindRWST2_7_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindRWST1__193435443_4_14, dictMonoid_5))
_ = bindRWST2_7_28
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_6_27)}
}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_7_28)}
})}))}
})
_ = monadRWST1__193435443_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadTransRWST1_4_29 shape=LitRecord bindingType=(ADT ["Control","Monad","Trans","Class","MonadTrans"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadTransRWST1_4_29 := (&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_30 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_5_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_30
// TAST (Let): pure_6_31 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_6_31 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_31
return gopurs_runtime.Func3(func(m_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_30.V1), m_7, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_31, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_9, a_10, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")}))})
}))
})
})})
_ = monadTransRWST1_4_29
// TAST (Let): monadRWST2_5_32 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadRWST2_5_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadRWST1__193435443_2_1, dictMonoid_3))
_ = monadRWST2_5_32
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_5_32)}
}), gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransRWST1_4_29.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_1_0)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_6))
})}))}
})
}

func Call_Control_Monad_RWS_Trans_monadErrorRWST(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Monad0_2_2 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(TypeVar m)])
Monad0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_2_2
// TAST (Let): __local_var_3_4 shape=App(Other) bindingType=Any
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): pure_4_6 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_4_6 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_6
// TAST (Let): __local_var_5_8 shape=App(Other) bindingType=Any
__local_var_5_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_5_8
// TAST (Let): Apply0_6_9 shape=App(Other) bindingType=Any
Apply0_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_8, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_9
// TAST (Let): Functor0_7_10 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_9, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_10
// TAST (Let): __local_var_8_12 shape=App(Other) bindingType=Any
__local_var_8_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_9, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_12
// TAST (Let): functorRWST1_8_11 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_8_11 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_9 gopurs_runtime.Value, v_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_12, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V0, gopurs_runtime.Apply(f_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_10, r_11, s_12))
})})
_ = functorRWST1_8_11
// TAST (Let): applyRWST1__193435443_5_7 shape=Let(Let(Let(Let(Abs(Let(LitRecord)))))) bindingType=Any
applyRWST1__193435443_5_7 := gopurs_runtime.Func(func(dictMonoid_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_10_13 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_10_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_9, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_10_13
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_8_11)}
}), gopurs_runtime.Func4(func(v_11 gopurs_runtime.Value, v1_12 gopurs_runtime.Value, r_13 gopurs_runtime.Value, s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_8, "bind"), gopurs_runtime.Apply2(v_11, r_13, s_14), gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_14 shape=Other bindingType=Any
__local_var_16_14 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V2
_ = __local_var_16_14
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_10.V0), gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_13.V0), __local_var_16_14, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_12, r_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_5_7
// TAST (Let): applicativeRWST1__193435443_4_5 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
applicativeRWST1__193435443_4_5 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_7_15 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_7_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_5_7, dictMonoid_6))
_ = applyRWST2_7_15
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_7_15)}
}), gopurs_runtime.Func3(func(a_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_6, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_10, a_8, gopurs_runtime.RecordGet(dictMonoid_6, "mempty")}))})
})}))}
})
_ = applicativeRWST1__193435443_4_5
// TAST (Let): __local_var_5_17 shape=App(Other) bindingType=Any
__local_var_5_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_5_17
// TAST (Let): Functor0_6_18 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_17, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_18
// TAST (Let): Apply0_7_20 shape=App(Other) bindingType=Any
Apply0_7_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_17, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_20
// TAST (Let): Functor0_8_21 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_8_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_20, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_21
// TAST (Let): __local_var_9_23 shape=App(Other) bindingType=Any
__local_var_9_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_20, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_23
// TAST (Let): functorRWST1_9_22 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_9_22 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_10 gopurs_runtime.Value, v_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_23, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V0, gopurs_runtime.Apply(f_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_11, r_12, s_13))
})})
_ = functorRWST1_9_22
// TAST (Let): applyRWST1__193435443_7_19 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
applyRWST1__193435443_7_19 := gopurs_runtime.Func(func(dictMonoid_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_11_24 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_11_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_10, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_11_24
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_9_22)}
}), gopurs_runtime.Func4(func(v_12 gopurs_runtime.Value, v1_13 gopurs_runtime.Value, r_14 gopurs_runtime.Value, s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_17, "bind"), gopurs_runtime.Apply2(v_12, r_14, s_15), gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_25 shape=Other bindingType=Any
__local_var_17_25 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_16.UnsafePtr).V2
_ = __local_var_17_25
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_21.V0), gopurs_runtime.Func(func(v3_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_18.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_16.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_18.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_11_24.V0), __local_var_17_25, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_18.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_13, r_14, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_16.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_7_19
// TAST (Let): bindRWST1__193435443_5_16 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
bindRWST1__193435443_5_16 := gopurs_runtime.Func(func(dictMonoid_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_26 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_9_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_8, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_26
// TAST (Let): applyRWST2_10_27 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_10_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_7_19, dictMonoid_8))
_ = applyRWST2_10_27
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_10_27)}
}), gopurs_runtime.Func4(func(v_11 gopurs_runtime.Value, f_12 gopurs_runtime.Value, r_13 gopurs_runtime.Value, s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_17, "bind"), gopurs_runtime.Apply2(v_11, r_13, s_14), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_28 shape=Other bindingType=Any
__local_var_16_28 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_15.UnsafePtr).V2
_ = __local_var_16_28
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_18.V0), gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_26.V0), __local_var_16_28, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply3(f_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_15.UnsafePtr).V1, r_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_15.UnsafePtr).V0))
}))
})}))}
})
_ = bindRWST1__193435443_5_16
// TAST (Let): monadRWST1__193435443_3_3 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
monadRWST1__193435443_3_3 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST2_7_29 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applicativeRWST2_7_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeRWST1__193435443_4_5, dictMonoid_6))
_ = applicativeRWST2_7_29
// TAST (Let): bindRWST2_8_30 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
bindRWST2_8_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(bindRWST1__193435443_5_16, dictMonoid_6))
_ = bindRWST2_8_30
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_7_29)}
}), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_8_30)}
})}))}
})
_ = monadRWST1__193435443_3_3
// TAST (Let): monadThrowRWST1__193435443_1_0 shape=Let(Let(Let(Abs(Let(Let(LitRecord)))))) bindingType=Any
monadThrowRWST1__193435443_1_0 := gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadTransRWST1_5_31 shape=LitRecord bindingType=(ADT ["Control","Monad","Trans","Class","MonadTrans"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadTransRWST1_5_31 := (&Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonad_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_32 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_6_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_32
// TAST (Let): pure_7_33 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_7_33 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_5, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_33
return gopurs_runtime.Func3(func(m_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_32.V1), m_8, gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_33, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_10, a_11, gopurs_runtime.RecordGet(dictMonoid_4, "mempty")}))})
}))
})
})})
_ = monadTransRWST1_5_31
// TAST (Let): monadRWST2_6_34 shape=App(Other) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadRWST2_6_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(monadRWST1__193435443_3_3, dictMonoid_4))
_ = monadRWST2_6_34
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_6_34)}
}), gopurs_runtime.Func(func(e_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransRWST1_5_31.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_2_2)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "throwError"), e_7))
})}))}
})
_ = monadThrowRWST1__193435443_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadThrowRWST2_3_35 shape=App(Other) bindingType=(ADT ["Control","Monad","Error","Class","MonadThrow"] [(TypeVar e), (Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
monadThrowRWST2_3_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(monadThrowRWST1__193435443_1_0, dictMonoid_2))
_ = monadThrowRWST2_3_35
return gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Error_Class_MonadError[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(monadThrowRWST2_3_35)}
}), gopurs_runtime.Func4(func(m_4 gopurs_runtime.Value, h_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), gopurs_runtime.Apply2(m_4, r_6, s_7), gopurs_runtime.Func(func(e_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(h_5, e_8, r_6, s_7)
}))
})}))}
})
}

func Call_Control_Monad_RWS_Trans_monadSTRWST(dictMonoid_0_loop gopurs_runtime.Value, dictMonadST_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
var dictMonadST_1 gopurs_runtime.Value = dictMonadST_1_loop
_ = dictMonadST_1
// TAST (Let): Monad0_2_0 shape=App(Other) bindingType=Any
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): pure_3_3 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_3_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_3
// TAST (Let): __local_var_4_5 shape=App(Other) bindingType=Any
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_4_5
// TAST (Let): Apply0_5_6 shape=App(Other) bindingType=Any
Apply0_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_6
// TAST (Let): Functor0_6_7 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_6, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_7
// TAST (Let): __local_var_7_9 shape=App(Other) bindingType=Any
__local_var_7_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_9
// TAST (Let): functorRWST1_7_8 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_7_8 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_9, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V0, gopurs_runtime.Apply(f_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_9, r_10, s_11))
})})
_ = functorRWST1_7_8
// TAST (Let): Semigroup0_8_10 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_8_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_8_10
// TAST (Let): applyRWST2_4_4 shape=Let(Let(Let(Let(Let(LitRecord))))) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_4_4 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_7_8)}
}), gopurs_runtime.Func4(func(v_9 gopurs_runtime.Value, v1_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_5, "bind"), gopurs_runtime.Apply2(v_9, r_11, s_12), gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_11 shape=Other bindingType=Any
__local_var_14_11 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V2
_ = __local_var_14_11
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_7.V0), gopurs_runtime.Func(func(v3_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_8_10.V0), __local_var_14_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_15.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_10, r_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0))
}))
})})
_ = applyRWST2_4_4
// TAST (Let): applicativeRWST2_3_2 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applicativeRWST2_3_2 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_4_4)}
}), gopurs_runtime.Func3(func(a_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_3, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_7, a_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}))})
})})
_ = applicativeRWST2_3_2
// TAST (Let): __local_var_4_13 shape=App(Other) bindingType=Any
__local_var_4_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_4_13
// TAST (Let): Functor0_5_14 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_14
// TAST (Let): Semigroup0_6_15 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_6_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_15
// TAST (Let): Apply0_7_17 shape=App(Other) bindingType=Any
Apply0_7_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_13, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_17
// TAST (Let): Functor0_8_18 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_8_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_17, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_18
// TAST (Let): __local_var_9_20 shape=App(Other) bindingType=Any
__local_var_9_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_17, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_20
// TAST (Let): functorRWST1_9_19 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_9_19 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_10 gopurs_runtime.Value, v_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_20, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V0, gopurs_runtime.Apply(f_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_14.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_11, r_12, s_13))
})})
_ = functorRWST1_9_19
// TAST (Let): Semigroup0_10_21 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_10_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_10_21
// TAST (Let): applyRWST2_7_16 shape=Let(Let(Let(Let(LitRecord)))) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_7_16 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_9_19)}
}), gopurs_runtime.Func4(func(v_11 gopurs_runtime.Value, v1_12 gopurs_runtime.Value, r_13 gopurs_runtime.Value, s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_13, "bind"), gopurs_runtime.Apply2(v_11, r_13, s_14), gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_22 shape=Other bindingType=Any
__local_var_16_22 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V2
_ = __local_var_16_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_18.V0), gopurs_runtime.Func(func(v3_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_21.V0), __local_var_16_22, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_17.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_12, r_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V0))
}))
})})
_ = applyRWST2_7_16
// TAST (Let): bindRWST2_4_12 shape=Let(Let(Let(Let(LitRecord)))) bindingType=(ADT ["Control","Bind","Bind"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
bindRWST2_4_12 := (&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_7_16)}
}), gopurs_runtime.Func4(func(v_8 gopurs_runtime.Value, f_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_13, "bind"), gopurs_runtime.Apply2(v_8, r_10, s_11), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_23 shape=Other bindingType=Any
__local_var_13_23 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V2
_ = __local_var_13_23
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_14.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_15.V0), __local_var_13_23, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply3(f_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V1, r_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_12.UnsafePtr).V0))
}))
})})
_ = bindRWST2_4_12
// TAST (Let): monadRWST1_3_1 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Monad","Monad"] [(Func [(TypeVar r), (TypeVar s')] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s'), (TypeVar a), (TypeVar w)])]))])
monadRWST1_3_1 := (&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_3_2)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_4_12)}
})})
_ = monadRWST1_3_1
// TAST (Let): Bind1_4_24 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_24
// TAST (Let): pure_5_25 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_5_25 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_25
return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST1_3_1)}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_26 shape=App(Other) bindingType=(TypeVar c)
__local_var_7_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_1, "liftST"), x_6)
_ = __local_var_7_26
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_24.V1), __local_var_7_26, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_25, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_9, a_10, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}))})
}))
})
})}))}
}

func Call_Control_Monad_RWS_Trans_monoidRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_1 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_1_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): Apply0_3_4 shape=App(Other) bindingType=Any
Apply0_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_3_4
// TAST (Let): Functor0_4_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_4, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
// TAST (Let): __local_var_5_7 shape=App(Other) bindingType=Any
__local_var_5_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_7
// TAST (Let): functorRWST1_5_6 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_5_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_7, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0, gopurs_runtime.Apply(f_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_7, r_8, s_9))
})})
_ = functorRWST1_5_6
// TAST (Let): applyRWST1__193435443_2_2 shape=Let(Let(Let(Let(Abs(Let(LitRecord)))))) bindingType=Any
applyRWST1__193435443_2_2 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_8 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_8
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_5_6)}
}), gopurs_runtime.Func4(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "bind"), gopurs_runtime.Apply2(v_8, r_10, s_11), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_9 shape=Other bindingType=Any
__local_var_13_9 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V2
_ = __local_var_13_9
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_8.V0), __local_var_13_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_9, r_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_2_2
// TAST (Let): applicativeRWST1__193435443_1_0 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
applicativeRWST1__193435443_1_0 := gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_4_10 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_4_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_2_2, dictMonoid_3))
_ = applyRWST2_4_10
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_4_10)}
}), gopurs_runtime.Func3(func(a_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_7, a_5, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")}))})
})}))}
})
_ = applicativeRWST1__193435443_1_0
// TAST (Let): __local_var_2_12 shape=App(Other) bindingType=Any
__local_var_2_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_12
// TAST (Let): Apply0_3_14 shape=App(Other) bindingType=Any
Apply0_3_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_12, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_3_14
// TAST (Let): Functor0_4_15 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_14, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_15
// TAST (Let): __local_var_5_17 shape=App(Other) bindingType=Any
__local_var_5_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_14, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_17
// TAST (Let): functorRWST1_5_16 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_5_16 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_17, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0, gopurs_runtime.Apply(f_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_7, r_8, s_9))
})})
_ = functorRWST1_5_16
// TAST (Let): applyRWST1__193435443_3_13 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
applyRWST1__193435443_3_13 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_18 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_7_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_18
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_5_16)}
}), gopurs_runtime.Func4(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_12, "bind"), gopurs_runtime.Apply2(v_8, r_10, s_11), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_19 shape=Other bindingType=Any
__local_var_13_19 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V2
_ = __local_var_13_19
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_15.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_18.V0), __local_var_13_19, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_14.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_9, r_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0))
}))
})}))}
})
_ = applyRWST1__193435443_3_13
// TAST (Let): semigroupRWST1__193435443_2_11 shape=Let(Let(Abs(Let(Abs(LitRecord))))) bindingType=Any
semigroupRWST1__193435443_2_11 := gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_5_20 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_5_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyRWST1__193435443_3_13, dictMonoid_4))
_ = applyRWST2_5_20
return gopurs_runtime.Func(func(dictSemigroup_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_7_21 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyRWST2_5_20.V0), gopurs_runtime.Value{}))
_ = Functor0_7_21
// TAST (Let): __local_var_8_22 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (TypeVar a))
__local_var_8_22 := gopurs_runtime.RecordGet(dictSemigroup_6, "append")
_ = __local_var_8_22
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_9 gopurs_runtime.Value, b_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyRWST2_5_20.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_21.V0), __local_var_8_22, a_9), b_10)
})}))}
})
})
_ = semigroupRWST1__193435443_2_11
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST2_4_23 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applicativeRWST2_4_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(applicativeRWST1__193435443_1_0, dictMonoid_3))
_ = applicativeRWST2_4_23
// TAST (Let): semigroupRWST2__193435443_5_24 shape=App(Other) bindingType=Any
semigroupRWST2__193435443_5_24 := gopurs_runtime.Apply(semigroupRWST1__193435443_2_11, dictMonoid_3)
_ = semigroupRWST2__193435443_5_24
return gopurs_runtime.Func(func(dictMonoid1_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupRWST3_7_25 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
semigroupRWST3_7_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(semigroupRWST2__193435443_5_24, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_6, "Semigroup0"), gopurs_runtime.Value{})))
_ = semigroupRWST3_7_25
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupRWST3_7_25)}
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeRWST2_4_23.V1), gopurs_runtime.RecordGet(dictMonoid1_6, "mempty"))}))}
})
})
}

func Call_Control_Monad_RWS_Trans_altRWST(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorRWST1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_1_0 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0, gopurs_runtime.Apply(f_2, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_3, r_4, s_5))
})})
_ = functorRWST1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_1_0)}
}), gopurs_runtime.Func4(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), gopurs_runtime.Apply2(v_2, r_4, s_5), gopurs_runtime.Apply2(v1_3, r_4, s_5))
})}))}
}

func Call_Control_Monad_RWS_Trans_plusRWST(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): empty_1_0 shape=Other bindingType=(TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])])
empty_1_0 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = empty_1_0
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): __local_var_3_4 shape=App(Other) bindingType=Any
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorRWST1_3_3 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_3_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_5, r_6, s_7))
})})
_ = functorRWST1_3_3
// TAST (Let): altRWST1_2_1 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Alt","Alt"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
altRWST1_2_1 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_3_3)}
}), gopurs_runtime.Func4(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "alt"), gopurs_runtime.Apply2(v_4, r_6, s_7), gopurs_runtime.Apply2(v1_5, r_6, s_7))
})})
_ = altRWST1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altRWST1_2_1)}
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_1_0
})}))}
}

func Call_Control_Monad_RWS_Trans_alternativeRWST(dictMonoid_0_loop gopurs_runtime.Value, dictAlternative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
var dictAlternative_1 gopurs_runtime.Value = dictAlternative_1_loop
_ = dictAlternative_1
// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=Any
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): empty_3_2 shape=Other bindingType=(TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])])
empty_3_2 := gopurs_runtime.RecordGet(__local_var_2_1, "empty")
_ = empty_3_2
// TAST (Let): __local_var_4_4 shape=App(Other) bindingType=Any
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_4_4
// TAST (Let): __local_var_5_6 shape=App(Other) bindingType=Any
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_6
// TAST (Let): functorRWST1_5_5 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_5_5 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_6, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V0, gopurs_runtime.Apply(f_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_10.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_7, r_8, s_9))
})})
_ = functorRWST1_5_5
// TAST (Let): altRWST1_4_3 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Alt","Alt"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
altRWST1_4_3 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_5_5)}
}), gopurs_runtime.Func4(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "alt"), gopurs_runtime.Apply2(v_6, r_8, s_9), gopurs_runtime.Apply2(v1_7, r_8, s_9))
})})
_ = altRWST1_4_3
// TAST (Let): plusRWST1_2_0 shape=Let(Let(Let(LitRecord))) bindingType=(ADT ["Control","Plus","Plus"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
plusRWST1_2_0 := (&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altRWST1_4_3)}
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_3_2
})})
_ = plusRWST1_2_0
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_4_8 shape=Other bindingType=(Func [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))
pure_4_8 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_8
// TAST (Let): __local_var_5_10 shape=App(Other) bindingType=Any
__local_var_5_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_5_10
// TAST (Let): Apply0_6_11 shape=App(Other) bindingType=Any
Apply0_6_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_10, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_11
// TAST (Let): Functor0_7_12 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_12
// TAST (Let): __local_var_8_14 shape=App(Other) bindingType=Any
__local_var_8_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_14
// TAST (Let): functorRWST1_8_13 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
functorRWST1_8_13 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(f_9 gopurs_runtime.Value, v_10 gopurs_runtime.Value, r_11 gopurs_runtime.Value, s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_14, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V0, gopurs_runtime.Apply(f_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_13.UnsafePtr).V2}))}
}), gopurs_runtime.Apply2(v_10, r_11, s_12))
})})
_ = functorRWST1_8_13
// TAST (Let): Semigroup0_9_15 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar w)])
Semigroup0_9_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_15
// TAST (Let): applyRWST2_5_9 shape=Let(Let(Let(Let(Let(LitRecord))))) bindingType=(ADT ["Control","Apply","Apply"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applyRWST2_5_9 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_8_13)}
}), gopurs_runtime.Func4(func(v_10 gopurs_runtime.Value, v1_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_10, "bind"), gopurs_runtime.Apply2(v_10, r_12, s_13), gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_16 shape=Other bindingType=Any
__local_var_15_16 := (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V2
_ = __local_var_15_16
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_12.V0), gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_15.V0), __local_var_15_16, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v3_16.UnsafePtr).V2)}))}
}), gopurs_runtime.Apply2(v1_11, r_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0))
}))
})})
_ = applyRWST2_5_9
// TAST (Let): applicativeRWST1_4_7 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Applicative","Applicative"] [(Func [(TypeVar r), (TypeVar s)] (TypeApp (TypeVar m) [(ADT ["Control","Monad","RWS","Trans","RWSResult"] [(TypeVar s), (TypeVar a), (TypeVar w)])]))])
applicativeRWST1_4_7 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_5_9)}
}), gopurs_runtime.Func3(func(a_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_8, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, s_8, a_6, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")}))})
})})
_ = applicativeRWST1_4_7
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST1_4_7)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusRWST1_2_0)}
})}))}
})
}

func Rebox_Control_Monad_RWS_Trans_1947641166_2326785955(in *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V1)}
		out.V2 = in.V2
	return out
}

func Rebox_Control_Monad_RWS_Trans_2567023824_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Control_Monad_RWS_Trans_4258103980_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[*Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value], *Constructor_Control_Monad_RWS_Trans_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}



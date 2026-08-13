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
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, value0, value1, value2})}
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
		cache_Control_Monad_RWS_Trans_newtypeRWST = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
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
		cache_Control_Monad_RWS_Trans_lazyRWST = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), r_1, s_2)
})
})
}))
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
return Call_Control_Monad_RWS_Trans_execRWST(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_RWS_Trans_execRWST
}

var cache_Control_Monad_RWS_Trans_evalRWST gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_evalRWST sync.Once
func Get_Control_Monad_RWS_Trans_evalRWST() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_evalRWST.Do(func() {
		cache_Control_Monad_RWS_Trans_evalRWST = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_evalRWST(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
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

var cache_Control_Monad_RWS_Trans_evalRWST__2982438712 gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_evalRWST__2982438712 sync.Once
func Get_Control_Monad_RWS_Trans_evalRWST__2982438712() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_evalRWST__2982438712.Do(func() {
		cache_Control_Monad_RWS_Trans_evalRWST__2982438712 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_evalRWST__2982438712(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_RWS_Trans_evalRWST__2982438712
}

var cache_Control_Monad_RWS_Trans_evalRWST__2802039376 gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_evalRWST__2802039376 sync.Once
func Get_Control_Monad_RWS_Trans_evalRWST__2802039376() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_evalRWST__2802039376.Do(func() {
		cache_Control_Monad_RWS_Trans_evalRWST__2802039376 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_evalRWST__2802039376(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Control_Monad_RWS_Trans_evalRWST__2802039376
}

var cache_Control_Monad_RWS_Trans_execRWST__2982438712 gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_execRWST__2982438712 sync.Once
func Get_Control_Monad_RWS_Trans_execRWST__2982438712() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_execRWST__2982438712.Do(func() {
		cache_Control_Monad_RWS_Trans_execRWST__2982438712 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_execRWST__2982438712(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_RWS_Trans_execRWST__2982438712
}

var cache_Control_Monad_RWS_Trans_execRWST__2802039376 gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_execRWST__2802039376 sync.Once
func Get_Control_Monad_RWS_Trans_execRWST__2802039376() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_execRWST__2802039376.Do(func() {
		cache_Control_Monad_RWS_Trans_execRWST__2802039376 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_execRWST__2802039376(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Control_Monad_RWS_Trans_execRWST__2802039376
}

var cache_Control_Monad_RWS_Trans_mapRWST__3506688348 gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_mapRWST__3506688348 sync.Once
func Get_Control_Monad_RWS_Trans_mapRWST__3506688348() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_mapRWST__3506688348.Do(func() {
		cache_Control_Monad_RWS_Trans_mapRWST__3506688348 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Call_Control_Monad_RWS_Trans_mapRWST__3506688348(f_0_box, v_1_box, r_2_box, s_3_box))}
})
	})
	return cache_Control_Monad_RWS_Trans_mapRWST__3506688348
}

var cache_Control_Monad_RWS_Trans_mapRWST__1363965404 gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_mapRWST__1363965404 sync.Once
func Get_Control_Monad_RWS_Trans_mapRWST__1363965404() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_mapRWST__1363965404.Do(func() {
		cache_Control_Monad_RWS_Trans_mapRWST__1363965404 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_RWS_Trans_mapRWST__1363965404(f_0_box, v_1_box, r_2_box, s_3_box)
})
	})
	return cache_Control_Monad_RWS_Trans_mapRWST__1363965404
}

var cache_Control_Monad_RWS_Trans_withRWST__673207610 gopurs_runtime.Value
var once_Control_Monad_RWS_Trans_withRWST__673207610 sync.Once
func Get_Control_Monad_RWS_Trans_withRWST__673207610() gopurs_runtime.Value {
	once_Control_Monad_RWS_Trans_withRWST__673207610.Do(func() {
		cache_Control_Monad_RWS_Trans_withRWST__673207610 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Call_Control_Monad_RWS_Trans_withRWST__673207610(f_0_box, m_1_box, r_2_box, s_3_box))}
})
	})
	return cache_Control_Monad_RWS_Trans_withRWST__673207610
}

type Constructor_Control_Monad_RWS_Trans_RWSResult struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
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
// TAST (Let): __local_var_4_0 -> *Constructor_Data_Tuple_Tuple
__local_var_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(f_0, r_2, s_3))
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
return gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_0 -> *Constructor_Control_Bind_Bind
Bind1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_0
// TAST (Let): pure_3_1 -> gopurs_runtime.Value
pure_3_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_1
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), m_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_6, a_7, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})
})
}))
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
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_5.UnsafePtr).V0, gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_5.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_5.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_2, r_3, s_4))
})
})
})
}))
}

func Call_Control_Monad_RWS_Trans_execRWST(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(v_3, r_4, s_5), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V2})})
}))
})
})
})
}

func Call_Control_Monad_RWS_Trans_evalRWST(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(v_3, r_4, s_5), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V2})})
}))
})
})
})
}

func Call_Control_Monad_RWS_Trans_applyRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): Apply0_1_0 -> gopurs_runtime.Value
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): functorRWST1_3_2 -> gopurs_runtime.Value
functorRWST1_3_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_5, r_6, s_7))
})
})
})
}))
_ = functorRWST1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_4 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_3_2
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_6, r_8, s_9), gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_5 -> gopurs_runtime.Value
__local_var_11_5 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_10.UnsafePtr).V2
_ = __local_var_11_5
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_12.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_10.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_12.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_4.V0), __local_var_11_5, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_12.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_7, r_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_10.UnsafePtr).V0))
}))
})
})
})
}))
})
}

func Call_Control_Monad_RWS_Trans_bindRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): applyRWST1_2_1 -> gopurs_runtime.Value
applyRWST1_2_1 := Call_Control_Monad_RWS_Trans_applyRWST(dictBind_0)
_ = applyRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_2
// TAST (Let): applyRWST2_5_3 -> gopurs_runtime.Value
applyRWST2_5_3 := gopurs_runtime.Apply(applyRWST1_2_1, dictMonoid_3)
_ = applyRWST2_5_3
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyRWST2_5_3
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_6, r_8, s_9), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_4 -> gopurs_runtime.Value
__local_var_11_4 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V2
_ = __local_var_11_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_12.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_12.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_2.V0), __local_var_11_4, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_12.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V1, r_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V0))
}))
})
})
})
}))
})
}

func Call_Control_Monad_RWS_Trans_semigroupRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): applyRWST1_1_0 -> gopurs_runtime.Value
applyRWST1_1_0 := Call_Control_Monad_RWS_Trans_applyRWST(dictBind_0)
_ = applyRWST1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_3_1 -> *Constructor_Control_Apply_Apply
applyRWST2_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(applyRWST1_1_0, dictMonoid_2))
_ = applyRWST2_3_1
return gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_5_2 -> *Constructor_Data_Functor_Functor
Functor0_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyRWST2_3_1.V0), gopurs_runtime.Value{}))
_ = Functor0_5_2
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := gopurs_runtime.RecordGet(dictSemigroup_4, "append")
_ = __local_var_6_3
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyRWST2_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_2.V0), __local_var_6_3, a_7), b_8)
})
}))
})
})
}

func Call_Control_Monad_RWS_Trans_applicativeRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): applyRWST1_2_1 -> gopurs_runtime.Value
applyRWST1_2_1 := Call_Control_Monad_RWS_Trans_applyRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = applyRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_4_2 -> gopurs_runtime.Value
applyRWST2_4_2 := gopurs_runtime.Apply(applyRWST1_2_1, dictMonoid_3)
_ = applyRWST2_4_2
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyRWST2_4_2
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_7, a_5, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")})})
})
})
}))
})
}

func Call_Control_Monad_RWS_Trans_monadRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): applicativeRWST1_1_0 -> gopurs_runtime.Value
applicativeRWST1_1_0 := Call_Control_Monad_RWS_Trans_applicativeRWST(dictMonad_0)
_ = applicativeRWST1_1_0
// TAST (Let): bindRWST1_2_1 -> gopurs_runtime.Value
bindRWST1_2_1 := Call_Control_Monad_RWS_Trans_bindRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = bindRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST2_4_2 -> gopurs_runtime.Value
applicativeRWST2_4_2 := gopurs_runtime.Apply(applicativeRWST1_1_0, dictMonoid_3)
_ = applicativeRWST2_4_2
// TAST (Let): bindRWST2_5_3 -> gopurs_runtime.Value
bindRWST2_5_3 := gopurs_runtime.Apply(bindRWST1_2_1, dictMonoid_3)
_ = bindRWST2_5_3
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeRWST2_4_2
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bindRWST2_5_3
}))
})
}

func Call_Control_Monad_RWS_Trans_monadAskRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): monadRWST1_2_1 -> gopurs_runtime.Value
monadRWST1_2_1 := Call_Control_Monad_RWS_Trans_monadRWST(dictMonad_0)
_ = monadRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadRWST2_4_2 -> gopurs_runtime.Value
monadRWST2_4_2 := gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3)
_ = monadRWST2_4_2
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_4_2
}), gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_6, r_5, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")})})
})
}))
})
}

func Call_Control_Monad_RWS_Trans_monadReaderRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): monadAskRWST1_1_0 -> gopurs_runtime.Value
monadAskRWST1_1_0 := Call_Control_Monad_RWS_Trans_monadAskRWST(dictMonad_0)
_ = monadAskRWST1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadAskRWST2_3_1 -> gopurs_runtime.Value
monadAskRWST2_3_1 := gopurs_runtime.Apply(monadAskRWST1_1_0, dictMonoid_2)
_ = monadAskRWST2_3_1
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskRWST2_3_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(m_5, gopurs_runtime.Apply(f_4, r_6), s_7)
})
})
})
}))
})
}

func Call_Control_Monad_RWS_Trans_monadEffectRWS(dictMonoid_0_loop gopurs_runtime.Value, dictMonadEffect_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
var dictMonadEffect_1 gopurs_runtime.Value = dictMonadEffect_1_loop
_ = dictMonadEffect_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): monadRWST1_3_1 -> gopurs_runtime.Value
monadRWST1_3_1 := gopurs_runtime.Apply(Call_Control_Monad_RWS_Trans_monadRWST(Monad0_2_0), dictMonoid_0)
_ = monadRWST1_3_1
// TAST (Let): Bind1_4_3 -> *Constructor_Control_Bind_Bind
Bind1_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_3
// TAST (Let): pure_5_4 -> gopurs_runtime.Value
pure_5_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_4
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_3.V1), m_6, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_4, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_8, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})
})
_ = __local_var_4_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST1_3_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_1, "liftEffect"), x_5))
}))
}

func Call_Control_Monad_RWS_Trans_monadRecRWST(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): monadRWST1_4_3 -> gopurs_runtime.Value
monadRWST1_4_3 := Call_Control_Monad_RWS_Trans_monadRWST(Monad0_1_0)
_ = monadRWST1_4_3
return gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_4 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_4
// TAST (Let): monadRWST2_7_5 -> gopurs_runtime.Value
monadRWST2_7_5 := gopurs_runtime.Apply(monadRWST1_4_3, dictMonoid_5)
_ = monadRWST2_7_5
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_7_5
}), gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_6 -> gopurs_runtime.Value
__local_var_13_6 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_12.UnsafePtr).V2
_ = __local_var_13_6
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply3(k_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_12.UnsafePtr).V1, r_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_12.UnsafePtr).V0), gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t_tag_7 gopurs_runtime.Value = (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_14.UnsafePtr).V1
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 525585346) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_14.UnsafePtr).V0, (*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_14.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_4.V0), __local_var_13_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_14.UnsafePtr).V2)})}})}
goto end_branch_9
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_14.UnsafePtr).V1
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 60402430) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_14.UnsafePtr).V0, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_14.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_4.V0), __local_var_13_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_14.UnsafePtr).V2)})}})}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), __t9)
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_11, a_9, gopurs_runtime.RecordGet(dictMonoid_5, "mempty")})})
})
})
})
}))
})
}

func Call_Control_Monad_RWS_Trans_monadStateRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): monadRWST1_2_1 -> gopurs_runtime.Value
monadRWST1_2_1 := Call_Control_Monad_RWS_Trans_monadRWST(dictMonad_0)
_ = monadRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadRWST2_4_2 -> gopurs_runtime.Value
monadRWST2_4_2 := gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3)
_ = monadRWST2_4_2
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_4_2
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_8_3 -> *Constructor_Data_Tuple_Tuple
v1_8_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_5, s_7))
_ = v1_8_3
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (v1_8_3).V1, (v1_8_3).V0, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")})})
})
})
}))
})
}

func Call_Control_Monad_RWS_Trans_monadTellRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): monadRWST1_2_1 -> gopurs_runtime.Value
monadRWST1_2_1 := Call_Control_Monad_RWS_Trans_monadRWST(dictMonad_0)
_ = monadRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_2 -> gopurs_runtime.Value
Semigroup0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_4_2
// TAST (Let): monadRWST2_5_3 -> gopurs_runtime.Value
monadRWST2_5_3 := gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3)
_ = monadRWST2_5_3
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_5_3
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_4_2
}), gopurs_runtime.Func(func(w_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_8, Get_Data_Unit_unit(), w_6})})
})
})
}))
})
}

func Call_Control_Monad_RWS_Trans_monadWriterRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 -> gopurs_runtime.Value
Applicative0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_2_1
// TAST (Let): monadTellRWST1_3_2 -> gopurs_runtime.Value
monadTellRWST1_3_2 := Call_Control_Monad_RWS_Trans_monadTellRWST(dictMonad_0)
_ = monadTellRWST1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadTellRWST2_5_3 -> gopurs_runtime.Value
monadTellRWST2_5_3 := gopurs_runtime.Apply(monadTellRWST1_3_2, dictMonoid_4)
_ = monadTellRWST2_5_3
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellRWST2_5_3
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonoid_4
}), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(m_6, r_7, s_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V2})}, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V2})})
}))
})
})
}), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(m_6, r_7, s_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V1.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V2)})})
}))
})
})
}))
})
}

func Call_Control_Monad_RWS_Trans_monadThrowRWST(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
// TAST (Let): Monad0_1_0 -> *Constructor_Control_Monad_Monad
Monad0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_1_0
// TAST (Let): monadRWST1_2_1 -> gopurs_runtime.Value
monadRWST1_2_1 := Call_Control_Monad_RWS_Trans_monadRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{}))
_ = monadRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadTransRWST1_4_2 -> *Constructor_Control_Monad_Trans_Class_MonadTrans
monadTransRWST1_4_2 := &Constructor_Control_Monad_Trans_Class_MonadTrans{1, gopurs_runtime.Func(func(dictMonad_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_3 -> *Constructor_Control_Bind_Bind
Bind1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_3
// TAST (Let): pure_6_4 -> gopurs_runtime.Value
pure_6_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_4
return gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_3.V1), m_7, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_6_4, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_9, a_10, gopurs_runtime.RecordGet(dictMonoid_3, "mempty")})})
}))
})
})
})
})}
_ = monadTransRWST1_4_2
// TAST (Let): monadRWST2_5_5 -> gopurs_runtime.Value
monadRWST2_5_5 := gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3)
_ = monadRWST2_5_5
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_5_5
}), gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransRWST1_4_2.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_1_0)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_6))
}))
})
}

func Call_Control_Monad_RWS_Trans_monadErrorRWST(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): monadThrowRWST1_1_0 -> gopurs_runtime.Value
monadThrowRWST1_1_0 := Call_Control_Monad_RWS_Trans_monadThrowRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowRWST1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadThrowRWST2_3_1 -> gopurs_runtime.Value
monadThrowRWST2_3_1 := gopurs_runtime.Apply(monadThrowRWST1_1_0, dictMonoid_2)
_ = monadThrowRWST2_3_1
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowRWST2_3_1
}), gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), gopurs_runtime.Apply2(m_4, r_6, s_7), gopurs_runtime.Func(func(e_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(h_5, e_8, r_6, s_7)
}))
})
})
})
}))
})
}

func Call_Control_Monad_RWS_Trans_monadSTRWST(dictMonoid_0_loop gopurs_runtime.Value, dictMonadST_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
var dictMonadST_1 gopurs_runtime.Value = dictMonadST_1_loop
_ = dictMonadST_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): monadRWST1_3_1 -> gopurs_runtime.Value
monadRWST1_3_1 := gopurs_runtime.Apply(Call_Control_Monad_RWS_Trans_monadRWST(Monad0_2_0), dictMonoid_0)
_ = monadRWST1_3_1
// TAST (Let): Bind1_4_3 -> *Constructor_Control_Bind_Bind
Bind1_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_3
// TAST (Let): pure_5_4 -> gopurs_runtime.Value
pure_5_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_4
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_3.V1), m_6, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_4, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_8, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})
})
_ = __local_var_4_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST1_3_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_1, "liftST"), x_5))
}))
}

func Call_Control_Monad_RWS_Trans_monoidRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): applicativeRWST1_1_0 -> gopurs_runtime.Value
applicativeRWST1_1_0 := Call_Control_Monad_RWS_Trans_applicativeRWST(dictMonad_0)
_ = applicativeRWST1_1_0
// TAST (Let): semigroupRWST1_2_1 -> gopurs_runtime.Value
semigroupRWST1_2_1 := Call_Control_Monad_RWS_Trans_semigroupRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = semigroupRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST2_4_2 -> *Constructor_Control_Applicative_Applicative
applicativeRWST2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(applicativeRWST1_1_0, dictMonoid_3))
_ = applicativeRWST2_4_2
// TAST (Let): semigroupRWST2_5_3 -> gopurs_runtime.Value
semigroupRWST2_5_3 := gopurs_runtime.Apply(semigroupRWST1_2_1, dictMonoid_3)
_ = semigroupRWST2_5_3
return gopurs_runtime.Func(func(dictMonoid1_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupRWST3_7_4 -> gopurs_runtime.Value
semigroupRWST3_7_4 := gopurs_runtime.Apply(semigroupRWST2_5_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupRWST3_7_4
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupRWST3_7_4
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeRWST2_4_2.V1), gopurs_runtime.RecordGet(dictMonoid1_6, "mempty")))
})
})
}

func Call_Control_Monad_RWS_Trans_altRWST(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorRWST1_1_0 -> gopurs_runtime.Value
functorRWST1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V0, gopurs_runtime.Apply(f_2, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_3, r_4, s_5))
})
})
})
}))
_ = functorRWST1_1_0
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), gopurs_runtime.Apply2(v_2, r_4, s_5), gopurs_runtime.Apply2(v1_3, r_4, s_5))
})
})
})
}))
}

func Call_Control_Monad_RWS_Trans_plusRWST(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
// TAST (Let): empty_1_0 -> gopurs_runtime.Value
empty_1_0 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = empty_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorRWST1_3_3 -> gopurs_runtime.Value
functorRWST1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_5, r_6, s_7))
})
})
})
}))
_ = functorRWST1_3_3
// TAST (Let): altRWST1_2_1 -> gopurs_runtime.Value
altRWST1_2_1 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_3_3
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "alt"), gopurs_runtime.Apply2(v_4, r_6, s_7), gopurs_runtime.Apply2(v1_5, r_6, s_7))
})
})
})
}))
_ = altRWST1_2_1
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return altRWST1_2_1
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_1_0
})
}))
}

func Call_Control_Monad_RWS_Trans_alternativeRWST(dictMonoid_0_loop gopurs_runtime.Value, dictAlternative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
var dictAlternative_1 gopurs_runtime.Value = dictAlternative_1_loop
_ = dictAlternative_1
// TAST (Let): plusRWST1_2_0 -> gopurs_runtime.Value
plusRWST1_2_0 := Call_Control_Monad_RWS_Trans_plusRWST(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Plus1"), gopurs_runtime.Value{}))
_ = plusRWST1_2_0
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST1_4_1 -> gopurs_runtime.Value
applicativeRWST1_4_1 := gopurs_runtime.Apply(Call_Control_Monad_RWS_Trans_applicativeRWST(dictMonad_3), dictMonoid_0)
_ = applicativeRWST1_4_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeRWST1_4_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return plusRWST1_2_0
}))
})
}

func Call_Control_Monad_RWS_Trans_evalRWST__2982438712(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(v_3, r_4, s_5), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V2})})
}))
})
})
})
}

func Call_Control_Monad_RWS_Trans_evalRWST__2802039376(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Identity_bindIdentity(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(__eta0_0, __eta1_1, __eta2_2)))}, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Identity_applicativeIdentity(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_3.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_3.UnsafePtr).V2})})))}
}))))}
}

func Call_Control_Monad_RWS_Trans_execRWST__2982438712(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(v_3, r_4, s_5), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V2})})
}))
})
})
})
}

func Call_Control_Monad_RWS_Trans_execRWST__2802039376(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Identity_bindIdentity(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(__eta0_0, __eta1_1, __eta2_2)))}, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Identity_applicativeIdentity(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_3.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_3.UnsafePtr).V2})})))}
}))))}
}

func Call_Control_Monad_RWS_Trans_mapRWST__3506688348(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) *Constructor_Control_Monad_RWS_Trans_RWSResult {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_RWS_Trans_RWSResult](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_RWS_Trans_RWSResult](gopurs_runtime.Apply2(v_1, r_2, s_3)))}))
}

func Call_Control_Monad_RWS_Trans_mapRWST__1363965404(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Control_Monad_RWS_Trans_withRWST__673207610(f_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) *Constructor_Control_Monad_RWS_Trans_RWSResult {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
// TAST (Let): __local_var_4_0 -> *Constructor_Data_Tuple_Tuple
__local_var_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(f_0, r_2, s_3))
_ = __local_var_4_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_RWS_Trans_RWSResult](gopurs_runtime.Apply2(m_1, (__local_var_4_0).V0, (__local_var_4_0).V1))
}



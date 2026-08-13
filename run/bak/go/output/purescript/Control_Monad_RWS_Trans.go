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
		cache_Control_Monad_RWS_Trans_newtypeRWST = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
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
		cache_Control_Monad_RWS_Trans_lazyRWST = gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(&Constructor_Control_Lazy_Lazy{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), r_1, s_2)
})
})
})})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Trans_Class_MonadTrans{1, gopurs_runtime.Func(func(dictMonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
})})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_5.UnsafePtr).V0, gopurs_runtime.Apply(f_1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_5.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_5.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_2, r_3, s_4))
})
})
})
})})}
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
// TAST (Let): functorRWST1_3_2 -> *Constructor_Data_Functor_Functor
functorRWST1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_5, r_6, s_7))
})
})
})
})))
_ = functorRWST1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_4 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_4
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_3_2)}
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
})})}
})
}

func Call_Control_Monad_RWS_Trans_bindRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_1 -> gopurs_runtime.Value
Apply0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_2_1
// TAST (Let): Functor0_3_2 -> *Constructor_Data_Functor_Functor
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_4
// TAST (Let): functorRWST1_4_3 -> *Constructor_Data_Functor_Functor
functorRWST1_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_9.UnsafePtr).V0, gopurs_runtime.Apply(f_5, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_9.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_9.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_6, r_7, s_8))
})
})
})
})))
_ = functorRWST1_4_3
return gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_5 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_5
// TAST (Let): Semigroup0_7_7 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_7
// TAST (Let): applyRWST2_7_6 -> *Constructor_Control_Apply_Apply
applyRWST2_7_6 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_4_3)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_8, r_10, s_11), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_8 -> gopurs_runtime.Value
__local_var_13_8 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_12.UnsafePtr).V2
_ = __local_var_13_8
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_14.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_12.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_14.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_7.V0), __local_var_13_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_14.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_9, r_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_12.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_7_6
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_7_6)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_8, r_10, s_11), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_9 -> gopurs_runtime.Value
__local_var_13_9 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_12.UnsafePtr).V2
_ = __local_var_13_9
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_14.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_14.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_5.V0), __local_var_13_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_14.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_12.UnsafePtr).V1, r_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_12.UnsafePtr).V0))
}))
})
})
})
})})}
})
}

func Call_Control_Monad_RWS_Trans_semigroupRWST(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): Apply0_1_0 -> gopurs_runtime.Value
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorRWST1_3_3 -> *Constructor_Data_Functor_Functor
functorRWST1_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_5, r_6, s_7))
})
})
})
})))
_ = functorRWST1_3_3
// TAST (Let): applyRWST1_3_2 -> gopurs_runtime.Value
applyRWST1_3_2 := gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_5 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_5
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_3_3)}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_6, r_8, s_9), gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_6 -> gopurs_runtime.Value
__local_var_11_6 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_10.UnsafePtr).V2
_ = __local_var_11_6
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v3_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_12.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_10.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_12.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_5.V0), __local_var_11_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_12.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_7, r_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_10.UnsafePtr).V0))
}))
})
})
})
})})}
})
_ = applyRWST1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_5_7 -> *Constructor_Control_Apply_Apply
applyRWST2_5_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(applyRWST1_3_2, dictMonoid_4))
_ = applyRWST2_5_7
return gopurs_runtime.Func(func(dictSemigroup_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_7_8 -> *Constructor_Data_Functor_Functor
Functor0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyRWST2_5_7.V0), gopurs_runtime.Value{}))
_ = Functor0_7_8
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.RecordGet(dictSemigroup_6, "append")
_ = __local_var_8_9
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyRWST2_5_7.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_8.V0), __local_var_8_9, a_9), b_10)
})
})})}
})
})
}

func Call_Control_Monad_RWS_Trans_applicativeRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): Apply0_3_2 -> gopurs_runtime.Value
Apply0_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_3_2
// TAST (Let): Functor0_4_4 -> *Constructor_Data_Functor_Functor
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_2, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_6
// TAST (Let): functorRWST1_5_5 -> *Constructor_Data_Functor_Functor
functorRWST1_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_6, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V0, gopurs_runtime.Apply(f_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_7, r_8, s_9))
})
})
})
})))
_ = functorRWST1_5_5
// TAST (Let): applyRWST1_4_3 -> gopurs_runtime.Value
applyRWST1_4_3 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_7 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_7
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_5_5)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply2(v_8, r_10, s_11), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_8 -> gopurs_runtime.Value
__local_var_13_8 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_12.UnsafePtr).V2
_ = __local_var_13_8
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_14.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_12.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_14.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_7.V0), __local_var_13_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_14.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_9, r_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_12.UnsafePtr).V0))
}))
})
})
})
})})}
})
_ = applyRWST1_4_3
return gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyRWST2_6_9 -> *Constructor_Control_Apply_Apply
applyRWST2_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(applyRWST1_4_3, dictMonoid_5))
_ = applyRWST2_6_9
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_6_9)}
}), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_9, a_7, gopurs_runtime.RecordGet(dictMonoid_5, "mempty")})})
})
})
})})}
})
}

func Call_Control_Monad_RWS_Trans_monadRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): Apply0_3_2 -> gopurs_runtime.Value
Apply0_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_3_2
// TAST (Let): Functor0_4_3 -> *Constructor_Data_Functor_Functor
Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_2, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_3
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_6
// TAST (Let): functorRWST1_5_5 -> *Constructor_Data_Functor_Functor
functorRWST1_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_6, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V0, gopurs_runtime.Apply(f_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_7, r_8, s_9))
})
})
})
})))
_ = functorRWST1_5_5
// TAST (Let): applicativeRWST1_5_4 -> gopurs_runtime.Value
applicativeRWST1_5_4 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_8 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_8
// TAST (Let): applyRWST2_7_7 -> *Constructor_Control_Apply_Apply
applyRWST2_7_7 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_5_5)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply2(v_8, r_10, s_11), gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_9 -> gopurs_runtime.Value
__local_var_13_9 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_12.UnsafePtr).V2
_ = __local_var_13_9
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), gopurs_runtime.Func(func(v3_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_14.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_12.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_14.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_8.V0), __local_var_13_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_14.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_9, r_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_12.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_7_7
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_7_7)}
}), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_10, a_8, gopurs_runtime.RecordGet(dictMonoid_6, "mempty")})})
})
})
})})}
})
_ = applicativeRWST1_5_4
// TAST (Let): __local_var_6_10 -> gopurs_runtime.Value
__local_var_6_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_6_10
// TAST (Let): Functor0_7_11 -> *Constructor_Data_Functor_Functor
Functor0_7_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_10, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_11
// TAST (Let): Apply0_8_12 -> gopurs_runtime.Value
Apply0_8_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_10, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_8_12
// TAST (Let): Functor0_9_13 -> *Constructor_Data_Functor_Functor
Functor0_9_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_12, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_13
// TAST (Let): __local_var_10_16 -> gopurs_runtime.Value
__local_var_10_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_12, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_16
// TAST (Let): functorRWST1_10_15 -> *Constructor_Data_Functor_Functor
functorRWST1_10_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_16, "map"), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_15.UnsafePtr).V0, gopurs_runtime.Apply(f_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_15.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_15.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_12, r_13, s_14))
})
})
})
})))
_ = functorRWST1_10_15
// TAST (Let): bindRWST1_10_14 -> gopurs_runtime.Value
bindRWST1_10_14 := gopurs_runtime.Func(func(dictMonoid_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_12_17 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_12_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_11, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_12_17
// TAST (Let): Semigroup0_13_19 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_13_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_11, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_13_19
// TAST (Let): applyRWST2_13_18 -> *Constructor_Control_Apply_Apply
applyRWST2_13_18 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_10_15)}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_10, "bind"), gopurs_runtime.Apply2(v_14, r_16, s_17), gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_20 -> gopurs_runtime.Value
__local_var_19_20 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V2
_ = __local_var_19_20
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_13.V0), gopurs_runtime.Func(func(v3_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_13_19.V0), __local_var_19_20, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_15, r_16, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_13_18
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_13_18)}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_10, "bind"), gopurs_runtime.Apply2(v_14, r_16, s_17), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_21 -> gopurs_runtime.Value
__local_var_19_21 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_18.UnsafePtr).V2
_ = __local_var_19_21
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_11.V0), gopurs_runtime.Func(func(v3_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_12_17.V0), __local_var_19_21, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_15, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_18.UnsafePtr).V1, r_16, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_18.UnsafePtr).V0))
}))
})
})
})
})})}
})
_ = bindRWST1_10_14
return gopurs_runtime.Func(func(dictMonoid_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeRWST2_12_22 -> *Constructor_Control_Applicative_Applicative
applicativeRWST2_12_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(applicativeRWST1_5_4, dictMonoid_11))
_ = applicativeRWST2_12_22
// TAST (Let): bindRWST2_13_23 -> *Constructor_Control_Bind_Bind
bindRWST2_13_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(bindRWST1_10_14, dictMonoid_11))
_ = bindRWST2_13_23
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_12_22)}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_13_23)}
})})}
})
}

func Call_Control_Monad_RWS_Trans_monadAskRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): Apply0_4_3 -> gopurs_runtime.Value
Apply0_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_3
// TAST (Let): Functor0_5_5 -> *Constructor_Data_Functor_Functor
Functor0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): functorRWST1_6_6 -> *Constructor_Data_Functor_Functor
functorRWST1_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_7, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_11.UnsafePtr).V0, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_11.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_11.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_8, r_9, s_10))
})
})
})
})))
_ = functorRWST1_6_6
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_7_8
// TAST (Let): Functor0_8_9 -> *Constructor_Data_Functor_Functor
Functor0_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_9
// TAST (Let): Apply0_9_10 -> gopurs_runtime.Value
Apply0_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_9_10
// TAST (Let): Functor0_10_11 -> *Constructor_Data_Functor_Functor
Functor0_10_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_10, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_11
// TAST (Let): __local_var_11_13 -> gopurs_runtime.Value
__local_var_11_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_10, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_13
// TAST (Let): functorRWST1_11_12 -> *Constructor_Data_Functor_Functor
functorRWST1_11_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_13, "map"), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_16.UnsafePtr).V0, gopurs_runtime.Apply(f_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_16.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_16.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_13, r_14, s_15))
})
})
})
})))
_ = functorRWST1_11_12
// TAST (Let): monadRWST1_5_4 -> gopurs_runtime.Value
monadRWST1_5_4 := gopurs_runtime.Func(func(dictMonoid_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_13_16 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_13_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_12, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_13_16
// TAST (Let): applyRWST2_13_15 -> *Constructor_Control_Apply_Apply
applyRWST2_13_15 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_6_6)}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply2(v_14, r_16, s_17), gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_17 -> gopurs_runtime.Value
__local_var_19_17 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V2
_ = __local_var_19_17
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_5.V0), gopurs_runtime.Func(func(v3_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_13_16.V0), __local_var_19_17, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_15, r_16, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_13_15
// TAST (Let): applicativeRWST2_13_14 -> *Constructor_Control_Applicative_Applicative
applicativeRWST2_13_14 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_13_15)}
}), gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_16, a_14, gopurs_runtime.RecordGet(dictMonoid_12, "mempty")})})
})
})
})}
_ = applicativeRWST2_13_14
// TAST (Let): Semigroup0_14_19 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_14_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_12, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_14_19
// TAST (Let): Semigroup0_15_21 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_15_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_12, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_15_21
// TAST (Let): applyRWST2_15_20 -> *Constructor_Control_Apply_Apply
applyRWST2_15_20 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_11_12)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_8, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_22 -> gopurs_runtime.Value
__local_var_21_22 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V2
_ = __local_var_21_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_11.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_15_21.V0), __local_var_21_22, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_17, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_15_20
// TAST (Let): bindRWST2_14_18 -> *Constructor_Control_Bind_Bind
bindRWST2_14_18 := &Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_15_20)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_8, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_23 -> gopurs_runtime.Value
__local_var_21_23 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V2
_ = __local_var_21_23
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_9.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_14_19.V0), __local_var_21_23, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_17, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V1, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = bindRWST2_14_18
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_13_14)}
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_14_18)}
})})}
})
_ = monadRWST1_5_4
return gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadRWST2_7_24 -> *Constructor_Control_Monad_Monad
monadRWST2_7_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(monadRWST1_5_4, dictMonoid_6))
_ = monadRWST2_7_24
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Reader_Class_MonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_7_24)}
}), gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_9, r_8, gopurs_runtime.RecordGet(dictMonoid_6, "mempty")})})
})
})})}
})
}

func Call_Control_Monad_RWS_Trans_monadReaderRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): Apply0_4_4 -> gopurs_runtime.Value
Apply0_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_4
// TAST (Let): Functor0_5_5 -> *Constructor_Data_Functor_Functor
Functor0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_4, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): functorRWST1_6_6 -> *Constructor_Data_Functor_Functor
functorRWST1_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_7, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_11.UnsafePtr).V0, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_11.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_11.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_8, r_9, s_10))
})
})
})
})))
_ = functorRWST1_6_6
// TAST (Let): __local_var_7_9 -> gopurs_runtime.Value
__local_var_7_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_7_9
// TAST (Let): Functor0_8_10 -> *Constructor_Data_Functor_Functor
Functor0_8_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_9, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_10
// TAST (Let): Apply0_9_11 -> gopurs_runtime.Value
Apply0_9_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_9, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_9_11
// TAST (Let): Functor0_10_12 -> *Constructor_Data_Functor_Functor
Functor0_10_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_12
// TAST (Let): __local_var_11_14 -> gopurs_runtime.Value
__local_var_11_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_14
// TAST (Let): functorRWST1_11_13 -> *Constructor_Data_Functor_Functor
functorRWST1_11_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_14, "map"), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_16.UnsafePtr).V0, gopurs_runtime.Apply(f_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_16.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_16.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_13, r_14, s_15))
})
})
})
})))
_ = functorRWST1_11_13
// TAST (Let): monadRWST1_7_8 -> gopurs_runtime.Value
monadRWST1_7_8 := gopurs_runtime.Func(func(dictMonoid_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_13_17 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_13_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_12, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_13_17
// TAST (Let): applyRWST2_13_16 -> *Constructor_Control_Apply_Apply
applyRWST2_13_16 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_6_6)}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "bind"), gopurs_runtime.Apply2(v_14, r_16, s_17), gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_18 -> gopurs_runtime.Value
__local_var_19_18 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V2
_ = __local_var_19_18
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_5.V0), gopurs_runtime.Func(func(v3_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_13_17.V0), __local_var_19_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_15, r_16, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_13_16
// TAST (Let): applicativeRWST2_13_15 -> *Constructor_Control_Applicative_Applicative
applicativeRWST2_13_15 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_13_16)}
}), gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_16, a_14, gopurs_runtime.RecordGet(dictMonoid_12, "mempty")})})
})
})
})}
_ = applicativeRWST2_13_15
// TAST (Let): Semigroup0_14_20 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_14_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_12, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_14_20
// TAST (Let): Semigroup0_15_22 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_15_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_12, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_15_22
// TAST (Let): applyRWST2_15_21 -> *Constructor_Control_Apply_Apply
applyRWST2_15_21 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_11_13)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_9, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_23 -> gopurs_runtime.Value
__local_var_21_23 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V2
_ = __local_var_21_23
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_12.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_15_22.V0), __local_var_21_23, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_17, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_15_21
// TAST (Let): bindRWST2_14_19 -> *Constructor_Control_Bind_Bind
bindRWST2_14_19 := &Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_15_21)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_9, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_24 -> gopurs_runtime.Value
__local_var_21_24 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V2
_ = __local_var_21_24
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_10.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_14_20.V0), __local_var_21_24, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_17, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V1, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = bindRWST2_14_19
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_13_15)}
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_14_19)}
})})}
})
_ = monadRWST1_7_8
// TAST (Let): monadAskRWST1_3_2 -> gopurs_runtime.Value
monadAskRWST1_3_2 := gopurs_runtime.Func(func(dictMonoid_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadRWST2_9_25 -> *Constructor_Control_Monad_Monad
monadRWST2_9_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(monadRWST1_7_8, dictMonoid_8))
_ = monadRWST2_9_25
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Reader_Class_MonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_9_25)}
}), gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_11, r_10, gopurs_runtime.RecordGet(dictMonoid_8, "mempty")})})
})
})})}
})
_ = monadAskRWST1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadAskRWST2_5_26 -> *Constructor_Control_Monad_Reader_Class_MonadAsk
monadAskRWST2_5_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk](gopurs_runtime.Apply(monadAskRWST1_3_2, dictMonoid_4))
_ = monadAskRWST2_5_26
return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Reader_Class_MonadReader{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(monadAskRWST2_5_26)}
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(m_7, gopurs_runtime.Apply(f_6, r_8), s_9)
})
})
})
})})}
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
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): Apply0_5_4 -> gopurs_runtime.Value
Apply0_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_4
// TAST (Let): Functor0_6_5 -> *Constructor_Data_Functor_Functor
Functor0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_4, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_5
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_7
// TAST (Let): functorRWST1_7_6 -> *Constructor_Data_Functor_Functor
functorRWST1_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_7, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_12.UnsafePtr).V0, gopurs_runtime.Apply(f_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_12.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_12.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_9, r_10, s_11))
})
})
})
})))
_ = functorRWST1_7_6
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_8_8
// TAST (Let): Functor0_9_9 -> *Constructor_Data_Functor_Functor
Functor0_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_9
// TAST (Let): Apply0_10_10 -> gopurs_runtime.Value
Apply0_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_10_10
// TAST (Let): Functor0_11_11 -> *Constructor_Data_Functor_Functor
Functor0_11_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_10_10, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_11_11
// TAST (Let): __local_var_12_13 -> gopurs_runtime.Value
__local_var_12_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_10_10, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_13
// TAST (Let): functorRWST1_12_12 -> *Constructor_Data_Functor_Functor
functorRWST1_12_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_12_13, "map"), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_17.UnsafePtr).V0, gopurs_runtime.Apply(f_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_17.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_17.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_14, r_15, s_16))
})
})
})
})))
_ = functorRWST1_12_12
// TAST (Let): Semigroup0_13_16 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_13_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_13_16
// TAST (Let): applyRWST2_13_15 -> *Constructor_Control_Apply_Apply
applyRWST2_13_15 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_7_6)}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "bind"), gopurs_runtime.Apply2(v_14, r_16, s_17), gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_17 -> gopurs_runtime.Value
__local_var_19_17 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V2
_ = __local_var_19_17
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_5.V0), gopurs_runtime.Func(func(v3_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_13_16.V0), __local_var_19_17, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_15, r_16, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_13_15
// TAST (Let): applicativeRWST2_13_14 -> *Constructor_Control_Applicative_Applicative
applicativeRWST2_13_14 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_13_15)}
}), gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_16, a_14, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
})
})
})}
_ = applicativeRWST2_13_14
// TAST (Let): Semigroup0_14_19 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_14_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_14_19
// TAST (Let): Semigroup0_15_21 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_15_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_15_21
// TAST (Let): applyRWST2_15_20 -> *Constructor_Control_Apply_Apply
applyRWST2_15_20 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_12_12)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_8, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_22 -> gopurs_runtime.Value
__local_var_21_22 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V2
_ = __local_var_21_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_11_11.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_15_21.V0), __local_var_21_22, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_17, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_15_20
// TAST (Let): bindRWST2_14_18 -> *Constructor_Control_Bind_Bind
bindRWST2_14_18 := &Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_15_20)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_8, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_23 -> gopurs_runtime.Value
__local_var_21_23 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V2
_ = __local_var_21_23
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_9.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_14_19.V0), __local_var_21_23, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_17, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V1, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = bindRWST2_14_18
// TAST (Let): monadRWST1_3_1 -> *Constructor_Control_Monad_Monad
monadRWST1_3_1 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_13_14)}
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_14_18)}
})}
_ = monadRWST1_3_1
// TAST (Let): Bind1_4_25 -> *Constructor_Control_Bind_Bind
Bind1_4_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_25
// TAST (Let): pure_5_26 -> gopurs_runtime.Value
pure_5_26 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_26
// TAST (Let): __local_var_4_24 -> gopurs_runtime.Value
__local_var_4_24 := gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_25.V1), m_6, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_26, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_8, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})
})
_ = __local_var_4_24
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(&Constructor_Effect_Class_MonadEffect{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST1_3_1)}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_24, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_1, "liftEffect"), x_5))
})})}
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
// TAST (Let): pure_4_3 -> gopurs_runtime.Value
pure_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_3
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): Apply0_6_5 -> gopurs_runtime.Value
Apply0_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_5
// TAST (Let): Functor0_7_7 -> *Constructor_Data_Functor_Functor
Functor0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_5, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_7
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_9
// TAST (Let): functorRWST1_8_8 -> *Constructor_Data_Functor_Functor
functorRWST1_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_9, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_13.UnsafePtr).V0, gopurs_runtime.Apply(f_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_13.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_13.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_10, r_11, s_12))
})
})
})
})))
_ = functorRWST1_8_8
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_9_10
// TAST (Let): Functor0_10_11 -> *Constructor_Data_Functor_Functor
Functor0_10_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_10, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_11
// TAST (Let): Apply0_11_12 -> gopurs_runtime.Value
Apply0_11_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_10, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_11_12
// TAST (Let): Functor0_12_13 -> *Constructor_Data_Functor_Functor
Functor0_12_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_11_12, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_12_13
// TAST (Let): __local_var_13_15 -> gopurs_runtime.Value
__local_var_13_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_11_12, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_15
// TAST (Let): functorRWST1_13_14 -> *Constructor_Data_Functor_Functor
functorRWST1_13_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_15, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_18.UnsafePtr).V0, gopurs_runtime.Apply(f_14, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_18.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_18.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_15, r_16, s_17))
})
})
})
})))
_ = functorRWST1_13_14
// TAST (Let): monadRWST1_7_6 -> gopurs_runtime.Value
monadRWST1_7_6 := gopurs_runtime.Func(func(dictMonoid_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_15_18 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_15_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_14, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_15_18
// TAST (Let): applyRWST2_15_17 -> *Constructor_Control_Apply_Apply
applyRWST2_15_17 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_8_8)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_19 -> gopurs_runtime.Value
__local_var_21_19 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V2
_ = __local_var_21_19
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_7.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_15_18.V0), __local_var_21_19, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_17, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_15_17
// TAST (Let): applicativeRWST2_15_16 -> *Constructor_Control_Applicative_Applicative
applicativeRWST2_15_16 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_15_17)}
}), gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_18, a_16, gopurs_runtime.RecordGet(dictMonoid_14, "mempty")})})
})
})
})}
_ = applicativeRWST2_15_16
// TAST (Let): Semigroup0_16_21 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_16_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_14, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_16_21
// TAST (Let): Semigroup0_17_23 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_17_23 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_14, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_17_23
// TAST (Let): applyRWST2_17_22 -> *Constructor_Control_Apply_Apply
applyRWST2_17_22 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_13_14)}
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_10, "bind"), gopurs_runtime.Apply2(v_18, r_20, s_21), gopurs_runtime.Func(func(v2_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_24 -> gopurs_runtime.Value
__local_var_23_24 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_22.UnsafePtr).V2
_ = __local_var_23_24
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_12_13.V0), gopurs_runtime.Func(func(v3_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_22.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_17_23.V0), __local_var_23_24, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_19, r_20, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_22.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_17_22
// TAST (Let): bindRWST2_16_20 -> *Constructor_Control_Bind_Bind
bindRWST2_16_20 := &Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_17_22)}
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_10, "bind"), gopurs_runtime.Apply2(v_18, r_20, s_21), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_25 -> gopurs_runtime.Value
__local_var_23_25 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_22.UnsafePtr).V2
_ = __local_var_23_25
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_11.V0), gopurs_runtime.Func(func(v3_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_16_21.V0), __local_var_23_25, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_19, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_22.UnsafePtr).V1, r_20, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_22.UnsafePtr).V0))
}))
})
})
})
})}
_ = bindRWST2_16_20
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_15_16)}
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_16_20)}
})})}
})
_ = monadRWST1_7_6
return gopurs_runtime.Func(func(dictMonoid_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_26 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_9_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_8, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_26
// TAST (Let): monadRWST2_10_27 -> *Constructor_Control_Monad_Monad
monadRWST2_10_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(monadRWST1_7_6, dictMonoid_8))
_ = monadRWST2_10_27
return gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_MonadRec{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_10_27)}
}), gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_28 -> gopurs_runtime.Value
__local_var_16_28 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_15.UnsafePtr).V2
_ = __local_var_16_28
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply3(k_11, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_15.UnsafePtr).V1, r_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_15.UnsafePtr).V0), gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 gopurs_runtime.Value
{
var __t_tag_29 gopurs_runtime.Value = (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_17.UnsafePtr).V1
if (__t_tag_29.Type == 9 && __t_tag_29.IntVal == 525585346) {
__t31 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_17.UnsafePtr).V0, (*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_17.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_26.V0), __local_var_16_28, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_17.UnsafePtr).V2)})}})}
goto end_branch_31
} else {

}
}
{
var __t_tag_30 gopurs_runtime.Value = (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_17.UnsafePtr).V1
if (__t_tag_30.Type == 9 && __t_tag_30.IntVal == 60402430) {
__t31 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_17.UnsafePtr).V0, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_17.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_26.V0), __local_var_16_28, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_17.UnsafePtr).V2)})}})}
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), __t31)
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_14, a_12, gopurs_runtime.RecordGet(dictMonoid_8, "mempty")})})
})
})
})
})})}
})
}

func Call_Control_Monad_RWS_Trans_monadStateRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): Apply0_4_3 -> gopurs_runtime.Value
Apply0_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_3
// TAST (Let): Functor0_5_5 -> *Constructor_Data_Functor_Functor
Functor0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): functorRWST1_6_6 -> *Constructor_Data_Functor_Functor
functorRWST1_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_7, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_11.UnsafePtr).V0, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_11.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_11.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_8, r_9, s_10))
})
})
})
})))
_ = functorRWST1_6_6
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_7_8
// TAST (Let): Functor0_8_9 -> *Constructor_Data_Functor_Functor
Functor0_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_9
// TAST (Let): Apply0_9_10 -> gopurs_runtime.Value
Apply0_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_9_10
// TAST (Let): Functor0_10_11 -> *Constructor_Data_Functor_Functor
Functor0_10_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_10, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_11
// TAST (Let): __local_var_11_13 -> gopurs_runtime.Value
__local_var_11_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_10, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_13
// TAST (Let): functorRWST1_11_12 -> *Constructor_Data_Functor_Functor
functorRWST1_11_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_13, "map"), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_16.UnsafePtr).V0, gopurs_runtime.Apply(f_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_16.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_16.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_13, r_14, s_15))
})
})
})
})))
_ = functorRWST1_11_12
// TAST (Let): monadRWST1_5_4 -> gopurs_runtime.Value
monadRWST1_5_4 := gopurs_runtime.Func(func(dictMonoid_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_13_16 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_13_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_12, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_13_16
// TAST (Let): applyRWST2_13_15 -> *Constructor_Control_Apply_Apply
applyRWST2_13_15 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_6_6)}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply2(v_14, r_16, s_17), gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_17 -> gopurs_runtime.Value
__local_var_19_17 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V2
_ = __local_var_19_17
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_5.V0), gopurs_runtime.Func(func(v3_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_13_16.V0), __local_var_19_17, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_15, r_16, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_13_15
// TAST (Let): applicativeRWST2_13_14 -> *Constructor_Control_Applicative_Applicative
applicativeRWST2_13_14 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_13_15)}
}), gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_16, a_14, gopurs_runtime.RecordGet(dictMonoid_12, "mempty")})})
})
})
})}
_ = applicativeRWST2_13_14
// TAST (Let): Semigroup0_14_19 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_14_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_12, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_14_19
// TAST (Let): Semigroup0_15_21 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_15_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_12, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_15_21
// TAST (Let): applyRWST2_15_20 -> *Constructor_Control_Apply_Apply
applyRWST2_15_20 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_11_12)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_8, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_22 -> gopurs_runtime.Value
__local_var_21_22 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V2
_ = __local_var_21_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_11.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_15_21.V0), __local_var_21_22, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_17, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_15_20
// TAST (Let): bindRWST2_14_18 -> *Constructor_Control_Bind_Bind
bindRWST2_14_18 := &Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_15_20)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_8, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_23 -> gopurs_runtime.Value
__local_var_21_23 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V2
_ = __local_var_21_23
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_9.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_14_19.V0), __local_var_21_23, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_17, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V1, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = bindRWST2_14_18
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_13_14)}
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_14_18)}
})})}
})
_ = monadRWST1_5_4
return gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadRWST2_7_24 -> *Constructor_Control_Monad_Monad
monadRWST2_7_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(monadRWST1_5_4, dictMonoid_6))
_ = monadRWST2_7_24
return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_State_Class_MonadState{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_7_24)}
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_11_25 -> *Constructor_Data_Tuple_Tuple
v1_11_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_8, s_10))
_ = v1_11_25
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (v1_11_25).V1, (v1_11_25).V0, gopurs_runtime.RecordGet(dictMonoid_6, "mempty")})})
})
})
})})}
})
}

func Call_Control_Monad_RWS_Trans_monadTellRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): Apply0_4_3 -> gopurs_runtime.Value
Apply0_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_4_3
// TAST (Let): Functor0_5_5 -> *Constructor_Data_Functor_Functor
Functor0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_5
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_4_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): functorRWST1_6_6 -> *Constructor_Data_Functor_Functor
functorRWST1_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_7, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_11.UnsafePtr).V0, gopurs_runtime.Apply(f_7, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_11.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_11.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_8, r_9, s_10))
})
})
})
})))
_ = functorRWST1_6_6
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_7_8
// TAST (Let): Functor0_8_9 -> *Constructor_Data_Functor_Functor
Functor0_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_9
// TAST (Let): Apply0_9_10 -> gopurs_runtime.Value
Apply0_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_9_10
// TAST (Let): Functor0_10_11 -> *Constructor_Data_Functor_Functor
Functor0_10_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_10, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_11
// TAST (Let): __local_var_11_13 -> gopurs_runtime.Value
__local_var_11_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_9_10, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_13
// TAST (Let): functorRWST1_11_12 -> *Constructor_Data_Functor_Functor
functorRWST1_11_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_13, "map"), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_16.UnsafePtr).V0, gopurs_runtime.Apply(f_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_16.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_16.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_13, r_14, s_15))
})
})
})
})))
_ = functorRWST1_11_12
// TAST (Let): monadRWST1_5_4 -> gopurs_runtime.Value
monadRWST1_5_4 := gopurs_runtime.Func(func(dictMonoid_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_13_16 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_13_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_12, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_13_16
// TAST (Let): applyRWST2_13_15 -> *Constructor_Control_Apply_Apply
applyRWST2_13_15 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_6_6)}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply2(v_14, r_16, s_17), gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_17 -> gopurs_runtime.Value
__local_var_19_17 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V2
_ = __local_var_19_17
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_5.V0), gopurs_runtime.Func(func(v3_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_13_16.V0), __local_var_19_17, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_15, r_16, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_13_15
// TAST (Let): applicativeRWST2_13_14 -> *Constructor_Control_Applicative_Applicative
applicativeRWST2_13_14 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_13_15)}
}), gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_16, a_14, gopurs_runtime.RecordGet(dictMonoid_12, "mempty")})})
})
})
})}
_ = applicativeRWST2_13_14
// TAST (Let): Semigroup0_14_19 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_14_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_12, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_14_19
// TAST (Let): Semigroup0_15_21 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_15_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_12, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_15_21
// TAST (Let): applyRWST2_15_20 -> *Constructor_Control_Apply_Apply
applyRWST2_15_20 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_11_12)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_8, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_22 -> gopurs_runtime.Value
__local_var_21_22 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V2
_ = __local_var_21_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_11.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_15_21.V0), __local_var_21_22, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_17, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_15_20
// TAST (Let): bindRWST2_14_18 -> *Constructor_Control_Bind_Bind
bindRWST2_14_18 := &Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_15_20)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_8, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_23 -> gopurs_runtime.Value
__local_var_21_23 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V2
_ = __local_var_21_23
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_9.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_14_19.V0), __local_var_21_23, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_17, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V1, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = bindRWST2_14_18
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_13_14)}
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_14_18)}
})})}
})
_ = monadRWST1_5_4
return gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_24 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_24
// TAST (Let): monadRWST2_8_25 -> *Constructor_Control_Monad_Monad
monadRWST2_8_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(monadRWST1_5_4, dictMonoid_6))
_ = monadRWST2_8_25
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Writer_Class_MonadTell{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_8_25)}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_7_24)}
}), gopurs_runtime.Func(func(w_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_11, Get_Data_Unit_unit(), w_9})})
})
})
})})}
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
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
// TAST (Let): pure_4_3 -> gopurs_runtime.Value
pure_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_3
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): Apply0_6_6 -> gopurs_runtime.Value
Apply0_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_6
// TAST (Let): Functor0_7_7 -> *Constructor_Data_Functor_Functor
Functor0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_6, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_7
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_9
// TAST (Let): functorRWST1_8_8 -> *Constructor_Data_Functor_Functor
functorRWST1_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_9, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_13.UnsafePtr).V0, gopurs_runtime.Apply(f_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_13.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_13.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_10, r_11, s_12))
})
})
})
})))
_ = functorRWST1_8_8
// TAST (Let): __local_var_9_11 -> gopurs_runtime.Value
__local_var_9_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_9_11
// TAST (Let): Functor0_10_12 -> *Constructor_Data_Functor_Functor
Functor0_10_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_11, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_12
// TAST (Let): Apply0_11_13 -> gopurs_runtime.Value
Apply0_11_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_11, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_11_13
// TAST (Let): Functor0_12_14 -> *Constructor_Data_Functor_Functor
Functor0_12_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_11_13, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_12_14
// TAST (Let): __local_var_13_16 -> gopurs_runtime.Value
__local_var_13_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_11_13, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_16
// TAST (Let): functorRWST1_13_15 -> *Constructor_Data_Functor_Functor
functorRWST1_13_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_16, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_18.UnsafePtr).V0, gopurs_runtime.Apply(f_14, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_18.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_18.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_15, r_16, s_17))
})
})
})
})))
_ = functorRWST1_13_15
// TAST (Let): monadRWST1_9_10 -> gopurs_runtime.Value
monadRWST1_9_10 := gopurs_runtime.Func(func(dictMonoid_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_15_19 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_15_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_14, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_15_19
// TAST (Let): applyRWST2_15_18 -> *Constructor_Control_Apply_Apply
applyRWST2_15_18 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_8_8)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_20 -> gopurs_runtime.Value
__local_var_21_20 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V2
_ = __local_var_21_20
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_7.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_15_19.V0), __local_var_21_20, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_17, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_15_18
// TAST (Let): applicativeRWST2_15_17 -> *Constructor_Control_Applicative_Applicative
applicativeRWST2_15_17 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_15_18)}
}), gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_18, a_16, gopurs_runtime.RecordGet(dictMonoid_14, "mempty")})})
})
})
})}
_ = applicativeRWST2_15_17
// TAST (Let): Semigroup0_16_22 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_16_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_14, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_16_22
// TAST (Let): Semigroup0_17_24 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_17_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_14, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_17_24
// TAST (Let): applyRWST2_17_23 -> *Constructor_Control_Apply_Apply
applyRWST2_17_23 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_13_15)}
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_11, "bind"), gopurs_runtime.Apply2(v_18, r_20, s_21), gopurs_runtime.Func(func(v2_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_25 -> gopurs_runtime.Value
__local_var_23_25 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_22.UnsafePtr).V2
_ = __local_var_23_25
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_12_14.V0), gopurs_runtime.Func(func(v3_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_22.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_17_24.V0), __local_var_23_25, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_19, r_20, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_22.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_17_23
// TAST (Let): bindRWST2_16_21 -> *Constructor_Control_Bind_Bind
bindRWST2_16_21 := &Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_17_23)}
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_11, "bind"), gopurs_runtime.Apply2(v_18, r_20, s_21), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_26 -> gopurs_runtime.Value
__local_var_23_26 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_22.UnsafePtr).V2
_ = __local_var_23_26
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_12.V0), gopurs_runtime.Func(func(v3_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_16_22.V0), __local_var_23_26, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_19, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_22.UnsafePtr).V1, r_20, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_22.UnsafePtr).V0))
}))
})
})
})
})}
_ = bindRWST2_16_21
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_15_17)}
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_16_21)}
})})}
})
_ = monadRWST1_9_10
// TAST (Let): monadTellRWST1_5_4 -> gopurs_runtime.Value
monadTellRWST1_5_4 := gopurs_runtime.Func(func(dictMonoid_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_11_27 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_11_27 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_10, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_11_27
// TAST (Let): monadRWST2_12_28 -> *Constructor_Control_Monad_Monad
monadRWST2_12_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(monadRWST1_9_10, dictMonoid_10))
_ = monadRWST2_12_28
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Writer_Class_MonadTell{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_12_28)}
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_11_27)}
}), gopurs_runtime.Func(func(w_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_15, Get_Data_Unit_unit(), w_13})})
})
})
})})}
})
_ = monadTellRWST1_5_4
return gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadTellRWST2_7_29 -> *Constructor_Control_Monad_Writer_Class_MonadTell
monadTellRWST2_7_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadTell](gopurs_runtime.Apply(monadTellRWST1_5_4, dictMonoid_6))
_ = monadTellRWST2_7_29
return gopurs_runtime.Value{Type: 9, IntVal: 784743459, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Writer_Class_MonadWriter{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(monadTellRWST2_7_29)}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_6))}
}), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(m_8, r_9, s_10), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_11.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_11.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_11.UnsafePtr).V2})}, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_11.UnsafePtr).V2})})
}))
})
})
}), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(m_8, r_9, s_10), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_11.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_11.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Data_Tuple_Tuple)((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_11.UnsafePtr).V1.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v_11.UnsafePtr).V2)})})
}))
})
})
})})}
})
}

func Call_Control_Monad_RWS_Trans_monadThrowRWST(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
// TAST (Let): Monad0_1_0 -> *Constructor_Control_Monad_Monad
Monad0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): Apply0_5_5 -> gopurs_runtime.Value
Apply0_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_5
// TAST (Let): Functor0_6_6 -> *Constructor_Data_Functor_Functor
Functor0_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_5, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_6
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_8
// TAST (Let): functorRWST1_7_7 -> *Constructor_Data_Functor_Functor
functorRWST1_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_8, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_12.UnsafePtr).V0, gopurs_runtime.Apply(f_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_12.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_12.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_9, r_10, s_11))
})
})
})
})))
_ = functorRWST1_7_7
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_8_9
// TAST (Let): Functor0_9_10 -> *Constructor_Data_Functor_Functor
Functor0_9_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_9, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_10
// TAST (Let): Apply0_10_11 -> gopurs_runtime.Value
Apply0_10_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_9, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_10_11
// TAST (Let): Functor0_11_12 -> *Constructor_Data_Functor_Functor
Functor0_11_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_10_11, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_11_12
// TAST (Let): __local_var_12_14 -> gopurs_runtime.Value
__local_var_12_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_10_11, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_14
// TAST (Let): functorRWST1_12_13 -> *Constructor_Data_Functor_Functor
functorRWST1_12_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_12_14, "map"), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_17.UnsafePtr).V0, gopurs_runtime.Apply(f_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_17.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_17.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_14, r_15, s_16))
})
})
})
})))
_ = functorRWST1_12_13
// TAST (Let): monadRWST1_5_4 -> gopurs_runtime.Value
monadRWST1_5_4 := gopurs_runtime.Func(func(dictMonoid_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_14_17 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_14_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_13, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_14_17
// TAST (Let): applyRWST2_14_16 -> *Constructor_Control_Apply_Apply
applyRWST2_14_16 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_7_7)}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "bind"), gopurs_runtime.Apply2(v_15, r_17, s_18), gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_18 -> gopurs_runtime.Value
__local_var_20_18 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_19.UnsafePtr).V2
_ = __local_var_20_18
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_6.V0), gopurs_runtime.Func(func(v3_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_21.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_19.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_21.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_14_17.V0), __local_var_20_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_21.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_16, r_17, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_19.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_14_16
// TAST (Let): applicativeRWST2_14_15 -> *Constructor_Control_Applicative_Applicative
applicativeRWST2_14_15 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_14_16)}
}), gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_17, a_15, gopurs_runtime.RecordGet(dictMonoid_13, "mempty")})})
})
})
})}
_ = applicativeRWST2_14_15
// TAST (Let): Semigroup0_15_20 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_15_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_13, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_15_20
// TAST (Let): Semigroup0_16_22 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_16_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_13, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_16_22
// TAST (Let): applyRWST2_16_21 -> *Constructor_Control_Apply_Apply
applyRWST2_16_21 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_12_13)}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_9, "bind"), gopurs_runtime.Apply2(v_17, r_19, s_20), gopurs_runtime.Func(func(v2_21 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_22_23 -> gopurs_runtime.Value
__local_var_22_23 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_21.UnsafePtr).V2
_ = __local_var_22_23
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_11_12.V0), gopurs_runtime.Func(func(v3_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_23.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_21.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_23.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_16_22.V0), __local_var_22_23, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_23.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_18, r_19, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_21.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_16_21
// TAST (Let): bindRWST2_15_19 -> *Constructor_Control_Bind_Bind
bindRWST2_15_19 := &Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_16_21)}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_9, "bind"), gopurs_runtime.Apply2(v_17, r_19, s_20), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_22_24 -> gopurs_runtime.Value
__local_var_22_24 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_21.UnsafePtr).V2
_ = __local_var_22_24
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_10.V0), gopurs_runtime.Func(func(v3_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_23.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_23.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_15_20.V0), __local_var_22_24, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_23.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_21.UnsafePtr).V1, r_19, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_21.UnsafePtr).V0))
}))
})
})
})
})}
_ = bindRWST2_15_19
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_14_15)}
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_15_19)}
})})}
})
_ = monadRWST1_5_4
return gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadTransRWST1_7_25 -> *Constructor_Control_Monad_Trans_Class_MonadTrans
monadTransRWST1_7_25 := &Constructor_Control_Monad_Trans_Class_MonadTrans{1, gopurs_runtime.Func(func(dictMonad_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_26 -> *Constructor_Control_Bind_Bind
Bind1_8_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_26
// TAST (Let): pure_9_27 -> gopurs_runtime.Value
pure_9_27 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_27
return gopurs_runtime.Func(func(m_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_26.V1), m_10, gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_27, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_12, a_13, gopurs_runtime.RecordGet(dictMonoid_6, "mempty")})})
}))
})
})
})
})}
_ = monadTransRWST1_7_25
// TAST (Let): monadRWST2_8_28 -> *Constructor_Control_Monad_Monad
monadRWST2_8_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(monadRWST1_5_4, dictMonoid_6))
_ = monadRWST2_8_28
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Error_Class_MonadThrow{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_8_28)}
}), gopurs_runtime.Func(func(e_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransRWST1_7_25.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_1_0)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_9))
})})}
})
}

func Call_Control_Monad_RWS_Trans_monadErrorRWST(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): Monad0_2_1 -> *Constructor_Control_Monad_Monad
Monad0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): pure_4_4 -> gopurs_runtime.Value
pure_4_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_4
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): Apply0_6_6 -> gopurs_runtime.Value
Apply0_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_6
// TAST (Let): Functor0_7_7 -> *Constructor_Data_Functor_Functor
Functor0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_6, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_7
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_10
// TAST (Let): functorRWST1_8_9 -> *Constructor_Data_Functor_Functor
functorRWST1_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_10, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_13.UnsafePtr).V0, gopurs_runtime.Apply(f_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_13.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_13.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_10, r_11, s_12))
})
})
})
})))
_ = functorRWST1_8_9
// TAST (Let): __local_var_9_11 -> gopurs_runtime.Value
__local_var_9_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_9_11
// TAST (Let): Functor0_10_12 -> *Constructor_Data_Functor_Functor
Functor0_10_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_11, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_12
// TAST (Let): Apply0_11_13 -> gopurs_runtime.Value
Apply0_11_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_11, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_11_13
// TAST (Let): Functor0_12_14 -> *Constructor_Data_Functor_Functor
Functor0_12_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_11_13, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_12_14
// TAST (Let): __local_var_13_16 -> gopurs_runtime.Value
__local_var_13_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_11_13, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_16
// TAST (Let): functorRWST1_13_15 -> *Constructor_Data_Functor_Functor
functorRWST1_13_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_16, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_18.UnsafePtr).V0, gopurs_runtime.Apply(f_14, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_18.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_18.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_15, r_16, s_17))
})
})
})
})))
_ = functorRWST1_13_15
// TAST (Let): monadRWST1_8_8 -> gopurs_runtime.Value
monadRWST1_8_8 := gopurs_runtime.Func(func(dictMonoid_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_15_19 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_15_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_14, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_15_19
// TAST (Let): applyRWST2_15_18 -> *Constructor_Control_Apply_Apply
applyRWST2_15_18 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_8_9)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_20 -> gopurs_runtime.Value
__local_var_21_20 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V2
_ = __local_var_21_20
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_7.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_15_19.V0), __local_var_21_20, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_17, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_15_18
// TAST (Let): applicativeRWST2_15_17 -> *Constructor_Control_Applicative_Applicative
applicativeRWST2_15_17 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_15_18)}
}), gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_4, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_18, a_16, gopurs_runtime.RecordGet(dictMonoid_14, "mempty")})})
})
})
})}
_ = applicativeRWST2_15_17
// TAST (Let): Semigroup0_16_22 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_16_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_14, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_16_22
// TAST (Let): Semigroup0_17_24 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_17_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_14, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_17_24
// TAST (Let): applyRWST2_17_23 -> *Constructor_Control_Apply_Apply
applyRWST2_17_23 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_13_15)}
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_11, "bind"), gopurs_runtime.Apply2(v_18, r_20, s_21), gopurs_runtime.Func(func(v2_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_25 -> gopurs_runtime.Value
__local_var_23_25 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_22.UnsafePtr).V2
_ = __local_var_23_25
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_12_14.V0), gopurs_runtime.Func(func(v3_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_22.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_17_24.V0), __local_var_23_25, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_19, r_20, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_22.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_17_23
// TAST (Let): bindRWST2_16_21 -> *Constructor_Control_Bind_Bind
bindRWST2_16_21 := &Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_17_23)}
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_11, "bind"), gopurs_runtime.Apply2(v_18, r_20, s_21), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_26 -> gopurs_runtime.Value
__local_var_23_26 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_22.UnsafePtr).V2
_ = __local_var_23_26
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_12.V0), gopurs_runtime.Func(func(v3_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_16_22.V0), __local_var_23_26, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_24.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_19, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_22.UnsafePtr).V1, r_20, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_22.UnsafePtr).V0))
}))
})
})
})
})}
_ = bindRWST2_16_21
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_15_17)}
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_16_21)}
})})}
})
_ = monadRWST1_8_8
// TAST (Let): monadThrowRWST1_3_2 -> gopurs_runtime.Value
monadThrowRWST1_3_2 := gopurs_runtime.Func(func(dictMonoid_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadTransRWST1_10_27 -> *Constructor_Control_Monad_Trans_Class_MonadTrans
monadTransRWST1_10_27 := &Constructor_Control_Monad_Trans_Class_MonadTrans{1, gopurs_runtime.Func(func(dictMonad_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_28 -> *Constructor_Control_Bind_Bind
Bind1_11_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_28
// TAST (Let): pure_12_29 -> gopurs_runtime.Value
pure_12_29 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_10, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_29
return gopurs_runtime.Func(func(m_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_28.V1), m_13, gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_12_29, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_15, a_16, gopurs_runtime.RecordGet(dictMonoid_9, "mempty")})})
}))
})
})
})
})}
_ = monadTransRWST1_10_27
// TAST (Let): monadRWST2_11_30 -> *Constructor_Control_Monad_Monad
monadRWST2_11_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(monadRWST1_8_8, dictMonoid_9))
_ = monadRWST2_11_30
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Error_Class_MonadThrow{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST2_11_30)}
}), gopurs_runtime.Func(func(e_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(monadTransRWST1_10_27.V0), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_2_1)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "throwError"), e_12))
})})}
})
_ = monadThrowRWST1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadThrowRWST2_5_31 -> *Constructor_Control_Monad_Error_Class_MonadThrow
monadThrowRWST2_5_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Error_Class_MonadThrow](gopurs_runtime.Apply(monadThrowRWST1_3_2, dictMonoid_4))
_ = monadThrowRWST2_5_31
return gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Error_Class_MonadError{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(monadThrowRWST2_5_31)}
}), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), gopurs_runtime.Apply2(m_6, r_8, s_9), gopurs_runtime.Func(func(e_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(h_7, e_10, r_8, s_9)
}))
})
})
})
})})}
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
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): Apply0_5_4 -> gopurs_runtime.Value
Apply0_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_5_4
// TAST (Let): Functor0_6_5 -> *Constructor_Data_Functor_Functor
Functor0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_4, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_5
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_5_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_7
// TAST (Let): functorRWST1_7_6 -> *Constructor_Data_Functor_Functor
functorRWST1_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_7, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_12.UnsafePtr).V0, gopurs_runtime.Apply(f_8, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_12.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_12.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_9, r_10, s_11))
})
})
})
})))
_ = functorRWST1_7_6
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_8_8
// TAST (Let): Functor0_9_9 -> *Constructor_Data_Functor_Functor
Functor0_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_9
// TAST (Let): Apply0_10_10 -> gopurs_runtime.Value
Apply0_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_10_10
// TAST (Let): Functor0_11_11 -> *Constructor_Data_Functor_Functor
Functor0_11_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_10_10, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_11_11
// TAST (Let): __local_var_12_13 -> gopurs_runtime.Value
__local_var_12_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_10_10, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_13
// TAST (Let): functorRWST1_12_12 -> *Constructor_Data_Functor_Functor
functorRWST1_12_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_12_13, "map"), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_17.UnsafePtr).V0, gopurs_runtime.Apply(f_13, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_17.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_17.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_14, r_15, s_16))
})
})
})
})))
_ = functorRWST1_12_12
// TAST (Let): Semigroup0_13_16 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_13_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_13_16
// TAST (Let): applyRWST2_13_15 -> *Constructor_Control_Apply_Apply
applyRWST2_13_15 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_7_6)}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "bind"), gopurs_runtime.Apply2(v_14, r_16, s_17), gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_17 -> gopurs_runtime.Value
__local_var_19_17 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V2
_ = __local_var_19_17
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_5.V0), gopurs_runtime.Func(func(v3_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_13_16.V0), __local_var_19_17, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_20.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_15, r_16, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_18.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_13_15
// TAST (Let): applicativeRWST2_13_14 -> *Constructor_Control_Applicative_Applicative
applicativeRWST2_13_14 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_13_15)}
}), gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_16, a_14, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
})
})
})}
_ = applicativeRWST2_13_14
// TAST (Let): Semigroup0_14_19 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_14_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_14_19
// TAST (Let): Semigroup0_15_21 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_15_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_15_21
// TAST (Let): applyRWST2_15_20 -> *Constructor_Control_Apply_Apply
applyRWST2_15_20 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_12_12)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_8, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_22 -> gopurs_runtime.Value
__local_var_21_22 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V2
_ = __local_var_21_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_11_11.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_15_21.V0), __local_var_21_22, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_17, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_15_20
// TAST (Let): bindRWST2_14_18 -> *Constructor_Control_Bind_Bind
bindRWST2_14_18 := &Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_15_20)}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_8, "bind"), gopurs_runtime.Apply2(v_16, r_18, s_19), gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_23 -> gopurs_runtime.Value
__local_var_21_23 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V2
_ = __local_var_21_23
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_9.V0), gopurs_runtime.Func(func(v3_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_14_19.V0), __local_var_21_23, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_22.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_17, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V1, r_18, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_20.UnsafePtr).V0))
}))
})
})
})
})}
_ = bindRWST2_14_18
// TAST (Let): monadRWST1_3_1 -> *Constructor_Control_Monad_Monad
monadRWST1_3_1 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST2_13_14)}
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(bindRWST2_14_18)}
})}
_ = monadRWST1_3_1
// TAST (Let): Bind1_4_25 -> *Constructor_Control_Bind_Bind
Bind1_4_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_25
// TAST (Let): pure_5_26 -> gopurs_runtime.Value
pure_5_26 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_26
// TAST (Let): __local_var_4_24 -> gopurs_runtime.Value
__local_var_4_24 := gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_25.V1), m_6, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_26, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_8, a_9, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
}))
})
})
})
_ = __local_var_4_24
return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_ST_Class_MonadST{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadRWST1_3_1)}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_24, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_1, "liftST"), x_5))
})})}
}

func Call_Control_Monad_RWS_Trans_monoidRWST(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): pure_1_0 -> gopurs_runtime.Value
pure_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): Apply0_3_2 -> gopurs_runtime.Value
Apply0_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_3_2
// TAST (Let): Functor0_4_3 -> *Constructor_Data_Functor_Functor
Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_2, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_3
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorRWST1_5_4 -> *Constructor_Data_Functor_Functor
functorRWST1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V0, gopurs_runtime.Apply(f_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_7, r_8, s_9))
})
})
})
})))
_ = functorRWST1_5_4
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_6_6
// TAST (Let): Apply0_7_7 -> gopurs_runtime.Value
Apply0_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_7
// TAST (Let): Functor0_8_8 -> *Constructor_Data_Functor_Functor
Functor0_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_7, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_8
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_7, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_10
// TAST (Let): functorRWST1_9_9 -> *Constructor_Data_Functor_Functor
functorRWST1_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_10, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_14.UnsafePtr).V0, gopurs_runtime.Apply(f_10, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_14.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_14.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_11, r_12, s_13))
})
})
})
})))
_ = functorRWST1_9_9
return gopurs_runtime.Func(func(dictMonoid_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_11_13 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_11_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_10, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_11_13
// TAST (Let): applyRWST2_11_12 -> *Constructor_Control_Apply_Apply
applyRWST2_11_12 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_5_4)}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply2(v_12, r_14, s_15), gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_14 -> gopurs_runtime.Value
__local_var_17_14 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_16.UnsafePtr).V2
_ = __local_var_17_14
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), gopurs_runtime.Func(func(v3_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_18.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_16.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_18.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_11_13.V0), __local_var_17_14, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_18.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_13, r_14, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_16.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_11_12
// TAST (Let): applicativeRWST2_11_11 -> *Constructor_Control_Applicative_Applicative
applicativeRWST2_11_11 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_11_12)}
}), gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_14, a_12, gopurs_runtime.RecordGet(dictMonoid_10, "mempty")})})
})
})
})}
_ = applicativeRWST2_11_11
// TAST (Let): Semigroup0_12_16 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_10, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_12_16
// TAST (Let): applyRWST2_12_15 -> *Constructor_Control_Apply_Apply
applyRWST2_12_15 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_9_9)}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_6, "bind"), gopurs_runtime.Apply2(v_13, r_15, s_16), gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_17 -> gopurs_runtime.Value
__local_var_18_17 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_17.UnsafePtr).V2
_ = __local_var_18_17
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_8.V0), gopurs_runtime.Func(func(v3_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_19.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_17.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_19.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_12_16.V0), __local_var_18_17, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_19.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_14, r_15, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_17.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_12_15
return gopurs_runtime.Func(func(dictMonoid1_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_14_19 -> *Constructor_Data_Functor_Functor
Functor0_14_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyRWST2_12_15.V0), gopurs_runtime.Value{}))
_ = Functor0_14_19
// TAST (Let): __local_var_15_20 -> gopurs_runtime.Value
__local_var_15_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_13, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_15_20
// TAST (Let): semigroupRWST3_14_18 -> *Constructor_Data_Semigroup_Semigroup
semigroupRWST3_14_18 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyRWST2_12_15.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_14_19.V0), __local_var_15_20, a_16), b_17)
})
})}
_ = semigroupRWST3_14_18
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupRWST3_14_18)}
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeRWST2_11_11.V1), gopurs_runtime.RecordGet(dictMonoid1_13, "mempty"))})}
})
})
}

func Call_Control_Monad_RWS_Trans_altRWST(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorRWST1_1_0 -> *Constructor_Data_Functor_Functor
functorRWST1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V0, gopurs_runtime.Apply(f_2, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_3, r_4, s_5))
})
})
})
})))
_ = functorRWST1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_1_0)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), gopurs_runtime.Apply2(v_2, r_4, s_5), gopurs_runtime.Apply2(v1_3, r_4, s_5))
})
})
})
})})}
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
// TAST (Let): functorRWST1_3_3 -> *Constructor_Data_Functor_Functor
functorRWST1_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_5, r_6, s_7))
})
})
})
})))
_ = functorRWST1_3_3
// TAST (Let): altRWST1_2_1 -> *Constructor_Control_Alt_Alt
altRWST1_2_1 := &Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_3_3)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "alt"), gopurs_runtime.Apply2(v_4, r_6, s_7), gopurs_runtime.Apply2(v1_5, r_6, s_7))
})
})
})
})}
_ = altRWST1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(&Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altRWST1_2_1)}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_1_0
})
})})}
}

func Call_Control_Monad_RWS_Trans_alternativeRWST(dictMonoid_0_loop gopurs_runtime.Value, dictAlternative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
var dictAlternative_1 gopurs_runtime.Value = dictAlternative_1_loop
_ = dictAlternative_1
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): empty_3_2 -> gopurs_runtime.Value
empty_3_2 := gopurs_runtime.RecordGet(__local_var_2_1, "empty")
_ = empty_3_2
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_4_4
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_6
// TAST (Let): functorRWST1_5_5 -> *Constructor_Data_Functor_Functor
functorRWST1_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_6, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V0, gopurs_runtime.Apply(f_6, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_7, r_8, s_9))
})
})
})
})))
_ = functorRWST1_5_5
// TAST (Let): altRWST1_4_3 -> *Constructor_Control_Alt_Alt
altRWST1_4_3 := &Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_5_5)}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "alt"), gopurs_runtime.Apply2(v_6, r_8, s_9), gopurs_runtime.Apply2(v1_7, r_8, s_9))
})
})
})
})}
_ = altRWST1_4_3
// TAST (Let): plusRWST1_2_0 -> *Constructor_Control_Plus_Plus
plusRWST1_2_0 := &Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altRWST1_4_3)}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_3_2
})
})}
_ = plusRWST1_2_0
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_4_8 -> gopurs_runtime.Value
pure_4_8 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_8
// TAST (Let): __local_var_5_9 -> gopurs_runtime.Value
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_5_9
// TAST (Let): Apply0_6_10 -> gopurs_runtime.Value
Apply0_6_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_9, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_10
// TAST (Let): Functor0_7_11 -> *Constructor_Data_Functor_Functor
Functor0_7_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_10, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_11
// TAST (Let): __local_var_8_13 -> gopurs_runtime.Value
__local_var_8_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_10, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_13
// TAST (Let): functorRWST1_8_12 -> *Constructor_Data_Functor_Functor
functorRWST1_8_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_13, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_13.UnsafePtr).V0, gopurs_runtime.Apply(f_9, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_13.UnsafePtr).V1), (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_13.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_10, r_11, s_12))
})
})
})
})))
_ = functorRWST1_8_12
// TAST (Let): Semigroup0_9_15 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_9_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_15
// TAST (Let): applyRWST2_9_14 -> *Constructor_Control_Apply_Apply
applyRWST2_9_14 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorRWST1_8_12)}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_9, "bind"), gopurs_runtime.Apply2(v_10, r_12, s_13), gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_16 -> gopurs_runtime.Value
__local_var_15_16 := (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_14.UnsafePtr).V2
_ = __local_var_15_16
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_11.V0), gopurs_runtime.Func(func(v3_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_16.UnsafePtr).V0, gopurs_runtime.Apply((*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_14.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_16.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_15.V0), __local_var_15_16, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v3_16.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_11, r_12, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v2_14.UnsafePtr).V0))
}))
})
})
})
})}
_ = applyRWST2_9_14
// TAST (Let): applicativeRWST1_4_7 -> *Constructor_Control_Applicative_Applicative
applicativeRWST1_4_7 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyRWST2_9_14)}
}), gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_8, gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_RWS_Trans_RWSResult{1, s_12, a_10, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})})
})
})
})}
_ = applicativeRWST1_4_7
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alternative_Alternative{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeRWST1_4_7)}
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusRWST1_2_0)}
})})}
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
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
var __local_var_3_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(__eta0_0, __eta1_1, __eta2_2)))}
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(__local_var_3_0.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(__local_var_3_0.UnsafePtr).V2})}
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
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
var __local_var_3_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(__eta0_0, __eta1_1, __eta2_2)))}
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(__local_var_3_0.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(__local_var_3_0.UnsafePtr).V2})}
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



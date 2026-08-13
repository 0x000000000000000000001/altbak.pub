package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Except_Trans_ExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_ExceptT sync.Once
func Get_Control_Monad_Except_Trans_ExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_ExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_ExceptT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_ExceptT(x_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_ExceptT
}

var cache_Control_Monad_Except_Trans_withExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_withExceptT sync.Once
func Get_Control_Monad_Except_Trans_withExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_withExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_withExceptT = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_withExceptT(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Monad_Except_Trans_withExceptT
}

var cache_Control_Monad_Except_Trans_runExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_runExceptT sync.Once
func Get_Control_Monad_Except_Trans_runExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_runExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_runExceptT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_runExceptT(v_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_runExceptT
}

var cache_Control_Monad_Except_Trans_newtypeExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_newtypeExceptT sync.Once
func Get_Control_Monad_Except_Trans_newtypeExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_newtypeExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_newtypeExceptT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Control_Monad_Except_Trans_newtypeExceptT
}

var cache_Control_Monad_Except_Trans_monadTransExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadTransExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadTransExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadTransExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadTransExceptT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), m_3, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_4})})
}))
})
}))
	})
	return cache_Control_Monad_Except_Trans_monadTransExceptT
}

var cache_Control_Monad_Except_Trans_lift gopurs_runtime.Value
var once_Control_Monad_Except_Trans_lift sync.Once
func Get_Control_Monad_Except_Trans_lift() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_lift.Do(func() {
		cache_Control_Monad_Except_Trans_lift = gopurs_runtime.RecordGet(Get_Control_Monad_Except_Trans_monadTransExceptT(), "lift")
	})
	return cache_Control_Monad_Except_Trans_lift
}

var cache_Control_Monad_Except_Trans_mapExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_mapExceptT sync.Once
func Get_Control_Monad_Except_Trans_mapExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_mapExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_mapExceptT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_mapExceptT(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Except_Trans_mapExceptT
}

var cache_Control_Monad_Except_Trans_functorExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_functorExceptT sync.Once
func Get_Control_Monad_Except_Trans_functorExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_functorExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_functorExceptT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_functorExceptT(dictFunctor_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_functorExceptT
}

var cache_Control_Monad_Except_Trans_except gopurs_runtime.Value
var once_Control_Monad_Except_Trans_except sync.Once
func Get_Control_Monad_Except_Trans_except() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_except.Do(func() {
		cache_Control_Monad_Except_Trans_except = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_except(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), x_1_box)
})
	})
	return cache_Control_Monad_Except_Trans_except
}

var cache_Control_Monad_Except_Trans_monadExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadExceptT
}

var cache_Control_Monad_Except_Trans_bindExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_bindExceptT sync.Once
func Get_Control_Monad_Except_Trans_bindExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_bindExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_bindExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_bindExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_bindExceptT
}

var cache_Control_Monad_Except_Trans_applyExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_applyExceptT sync.Once
func Get_Control_Monad_Except_Trans_applyExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_applyExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_applyExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_applyExceptT
}

var cache_Control_Monad_Except_Trans_applicativeExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_applicativeExceptT sync.Once
func Get_Control_Monad_Except_Trans_applicativeExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_applicativeExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_applicativeExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_applicativeExceptT
}

var cache_Control_Monad_Except_Trans_semigroupExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_semigroupExceptT sync.Once
func Get_Control_Monad_Except_Trans_semigroupExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_semigroupExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_semigroupExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_semigroupExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_semigroupExceptT
}

var cache_Control_Monad_Except_Trans_monadAskExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadAskExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadAskExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadAskExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadAskExceptT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadAskExceptT(dictMonadAsk_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadAskExceptT
}

var cache_Control_Monad_Except_Trans_monadReaderExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadReaderExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadReaderExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadReaderExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadReaderExceptT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadReaderExceptT(dictMonadReader_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadReaderExceptT
}

var cache_Control_Monad_Except_Trans_monadContExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadContExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadContExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadContExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadContExceptT = gopurs_runtime.Func(func(dictMonadCont_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadContExceptT(dictMonadCont_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadContExceptT
}

var cache_Control_Monad_Except_Trans_monadEffectExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadEffectExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadEffectExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadEffectExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadEffectExceptT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadEffectExceptT(dictMonadEffect_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadEffectExceptT
}

var cache_Control_Monad_Except_Trans_monadRecExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadRecExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadRecExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadRecExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadRecExceptT = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadRecExceptT(dictMonadRec_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadRecExceptT
}

var cache_Control_Monad_Except_Trans_monadStateExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadStateExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadStateExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadStateExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadStateExceptT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadStateExceptT(dictMonadState_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadStateExceptT
}

var cache_Control_Monad_Except_Trans_monadTellExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadTellExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadTellExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadTellExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadTellExceptT = gopurs_runtime.Func(func(dictMonadTell_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadTellExceptT(dictMonadTell_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadTellExceptT
}

var cache_Control_Monad_Except_Trans_monadWriterExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadWriterExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadWriterExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadWriterExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadWriterExceptT = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadWriterExceptT(dictMonadWriter_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadWriterExceptT
}

var cache_Control_Monad_Except_Trans_monadThrowExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadThrowExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadThrowExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadThrowExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadThrowExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadThrowExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadThrowExceptT
}

var cache_Control_Monad_Except_Trans_monadErrorExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadErrorExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadErrorExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadErrorExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadErrorExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadErrorExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadErrorExceptT
}

var cache_Control_Monad_Except_Trans_monadSTExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadSTExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadSTExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadSTExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadSTExceptT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadSTExceptT(dictMonadST_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadSTExceptT
}

var cache_Control_Monad_Except_Trans_monoidExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monoidExceptT sync.Once
func Get_Control_Monad_Except_Trans_monoidExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monoidExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monoidExceptT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monoidExceptT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monoidExceptT
}

var cache_Control_Monad_Except_Trans_altExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_altExceptT sync.Once
func Get_Control_Monad_Except_Trans_altExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_altExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_altExceptT = gopurs_runtime.Func2(func(dictSemigroup_0_box gopurs_runtime.Value, dictMonad_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_altExceptT(dictSemigroup_0_box, dictMonad_1_box)
})
	})
	return cache_Control_Monad_Except_Trans_altExceptT
}

var cache_Control_Monad_Except_Trans_plusExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_plusExceptT sync.Once
func Get_Control_Monad_Except_Trans_plusExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_plusExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_plusExceptT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_plusExceptT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_plusExceptT
}

var cache_Control_Monad_Except_Trans_alternativeExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_alternativeExceptT sync.Once
func Get_Control_Monad_Except_Trans_alternativeExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_alternativeExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_alternativeExceptT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_alternativeExceptT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_alternativeExceptT
}

var cache_Control_Monad_Except_Trans_monadPlusExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadPlusExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadPlusExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadPlusExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadPlusExceptT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_monadPlusExceptT(dictMonoid_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_monadPlusExceptT
}

var cache_Control_Monad_Except_Trans_mapExceptT__3966186360 gopurs_runtime.Value
var once_Control_Monad_Except_Trans_mapExceptT__3966186360 sync.Once
func Get_Control_Monad_Except_Trans_mapExceptT__3966186360() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_mapExceptT__3966186360.Do(func() {
		cache_Control_Monad_Except_Trans_mapExceptT__3966186360 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_mapExceptT__3966186360(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Except_Trans_mapExceptT__3966186360
}

var cache_Control_Monad_Except_Trans_mapExceptT__4285021944 gopurs_runtime.Value
var once_Control_Monad_Except_Trans_mapExceptT__4285021944 sync.Once
func Get_Control_Monad_Except_Trans_mapExceptT__4285021944() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_mapExceptT__4285021944.Do(func() {
		cache_Control_Monad_Except_Trans_mapExceptT__4285021944 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_mapExceptT__4285021944(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Except_Trans_mapExceptT__4285021944
}

var cache_Control_Monad_Except_Trans_mapExceptT__853646328 gopurs_runtime.Value
var once_Control_Monad_Except_Trans_mapExceptT__853646328 sync.Once
func Get_Control_Monad_Except_Trans_mapExceptT__853646328() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_mapExceptT__853646328.Do(func() {
		cache_Control_Monad_Except_Trans_mapExceptT__853646328 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_mapExceptT__853646328(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Except_Trans_mapExceptT__853646328
}

var cache_Control_Monad_Except_Trans_mapExceptT__1163275960 gopurs_runtime.Value
var once_Control_Monad_Except_Trans_mapExceptT__1163275960 sync.Once
func Get_Control_Monad_Except_Trans_mapExceptT__1163275960() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_mapExceptT__1163275960.Do(func() {
		cache_Control_Monad_Except_Trans_mapExceptT__1163275960 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_mapExceptT__1163275960(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Except_Trans_mapExceptT__1163275960
}

var cache_Control_Monad_Except_Trans_monadTransExceptT__4007330348 gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadTransExceptT__4007330348 sync.Once
func Get_Control_Monad_Except_Trans_monadTransExceptT__4007330348() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadTransExceptT__4007330348.Do(func() {
		cache_Control_Monad_Except_Trans_monadTransExceptT__4007330348 = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), m_3, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_4})})
}))
})
}))
	})
	return cache_Control_Monad_Except_Trans_monadTransExceptT__4007330348
}

var cache_Control_Monad_Except_Trans_runExceptT__509566043 gopurs_runtime.Value
var once_Control_Monad_Except_Trans_runExceptT__509566043 sync.Once
func Get_Control_Monad_Except_Trans_runExceptT__509566043() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_runExceptT__509566043.Do(func() {
		cache_Control_Monad_Except_Trans_runExceptT__509566043 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_runExceptT__509566043(v_0_box)
})
	})
	return cache_Control_Monad_Except_Trans_runExceptT__509566043
}

func Call_Control_Monad_Except_Trans_ExceptT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Except_Trans_withExceptT(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_3.Type == 9 && v2_3.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(v2_3.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_Either_Left)(v2_3.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), v_2)
}

func Call_Control_Monad_Except_Trans_runExceptT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Control_Monad_Except_Trans_mapExceptT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Except_Trans_functorExceptT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_1))
_ = __local_var_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, v_3)
})
}))
}

func Call_Control_Monad_Except_Trans_except(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), x_1)
}

func Call_Control_Monad_Except_Trans_monadExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_bindExceptT(dictMonad_0)
}))
}

func Call_Control_Monad_Except_Trans_bindExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), v_3, gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 3711209382) {
__t2 = gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_5.UnsafePtr).V0})})
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply(k_4, (*Constructor_Data_Either_Right)(v2_5.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
})
}))
}

func Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): functorExceptT1_1_0 -> gopurs_runtime.Value
functorExceptT1_1_0 := Call_Control_Monad_Except_Trans_functorExceptT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = functorExceptT1_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_bindExceptT(dictMonad_0)
}))
_ = __local_var_2_1
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): Applicative0_4_3 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_1_0
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_3.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}

func Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_2})})
})
_ = __local_var_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
}))
}

func Call_Control_Monad_Except_Trans_semigroupExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): applyExceptT1_1_0 -> *Constructor_Control_Apply_Apply
applyExceptT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0))
_ = applyExceptT1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_1 -> *Constructor_Data_Functor_Functor
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyExceptT1_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_2
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyExceptT1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), __local_var_4_2, a_5), b_6)
})
}))
})
}

func Call_Control_Monad_Except_Trans_monadAskExceptT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): monadExceptT1_1_0 -> gopurs_runtime.Value
monadExceptT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_bindExceptT(__local_var_1_1)
}))
_ = monadExceptT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_1_0
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_Except_Trans_monadTransExceptT(), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})))}, gopurs_runtime.RecordGet(dictMonadAsk_0, "ask")))
}

func Call_Control_Monad_Except_Trans_monadReaderExceptT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
// TAST (Let): monadAskExceptT1_1_0 -> gopurs_runtime.Value
monadAskExceptT1_1_0 := Call_Control_Monad_Except_Trans_monadAskExceptT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{}))
_ = monadAskExceptT1_1_0
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskExceptT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, v_4)
})
}))
}

func Call_Control_Monad_Except_Trans_monadContExceptT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): monadExceptT1_1_0 -> gopurs_runtime.Value
monadExceptT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_bindExceptT(__local_var_1_1)
}))
_ = monadExceptT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_3, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_4})})
}))
}))
}))
}

func Call_Control_Monad_Except_Trans_monadEffectExceptT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadExceptT1_2_1 -> gopurs_runtime.Value
monadExceptT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_bindExceptT(Monad0_1_0)
}))
_ = monadExceptT1_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Except_Trans_monadTransExceptT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}))
}

func Call_Control_Monad_Except_Trans_monadRecExceptT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): monadExceptT1_4_3 -> gopurs_runtime.Value
monadExceptT1_4_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_bindExceptT(Monad0_1_0)
}))
_ = monadExceptT1_4_3
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_4_3
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(f_5, a_6), gopurs_runtime.Func(func(m_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (m_prime_7.Type == 9 && m_prime_7.IntVal == 3711209382) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_prime_7.UnsafePtr).V0})}})}
goto end_branch_8
} else {

}
}
{
if (m_prime_7.Type == 9 && m_prime_7.IntVal == 2465973597) {
var __t7 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Data_Either_Right)(m_prime_7.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 525585346) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, (*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Either_Right)(m_prime_7.UnsafePtr).V0.UnsafePtr).V0})}
goto end_branch_7
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Data_Either_Right)(m_prime_7.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 60402430) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Data_Either_Right)(m_prime_7.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), __t8)
}))
}))
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, x_7)
})
}))
}

func Call_Control_Monad_Except_Trans_monadStateExceptT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
// TAST (Let): Monad0_1_0 -> *Constructor_Control_Monad_Monad
Monad0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): monadExceptT1_2_1 -> gopurs_runtime.Value
monadExceptT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_bindExceptT(__local_var_2_2)
}))
_ = monadExceptT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_Except_Trans_monadTransExceptT(), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_1_0)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), f_3))
}))
}

func Call_Control_Monad_Except_Trans_monadTellExceptT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
// TAST (Let): Monad1_1_0 -> gopurs_runtime.Value
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
// TAST (Let): Semigroup0_2_1 -> gopurs_runtime.Value
Semigroup0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_2_1
// TAST (Let): monadExceptT1_3_2 -> gopurs_runtime.Value
monadExceptT1_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_bindExceptT(Monad1_1_0)
}))
_ = monadExceptT1_3_2
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Except_Trans_monadTransExceptT(), "lift"), Monad1_1_0)
_ = __local_var_4_3
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_5))
}))
}

func Call_Control_Monad_Except_Trans_monadWriterExceptT(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): pure_4_3 -> gopurs_runtime.Value
pure_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_3
// TAST (Let): Applicative0_5_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_4
// TAST (Let): Monoid0_6_5 -> gopurs_runtime.Value
Monoid0_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0"), gopurs_runtime.Value{})
_ = Monoid0_6_5
// TAST (Let): monadTellExceptT1_7_6 -> gopurs_runtime.Value
monadTellExceptT1_7_6 := Call_Control_Monad_Except_Trans_monadTellExceptT(MonadTell1_1_0)
_ = monadTellExceptT1_7_6
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellExceptT1_7_6
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_6_5
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), v_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_7 -> gopurs_runtime.Value
__local_var_10_7 := (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V1
_ = __local_var_10_7
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), gopurs_runtime.Func(func(r_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, r_11, __local_var_10_7})}
}), (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0))
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), v_8, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (a_9.Type == 9 && a_9.IntVal == 3711209382) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(a_9.UnsafePtr).V0})}, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return x_10
})})}
goto end_branch_8
} else {

}
}
{
if (a_9.Type == 9 && a_9.IntVal == 2465973597) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Either_Right)(a_9.UnsafePtr).V0.UnsafePtr).V0})}, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Either_Right)(a_9.UnsafePtr).V0.UnsafePtr).V1})}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](__t8))})
})))
}))
}

func Call_Control_Monad_Except_Trans_monadThrowExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): monadExceptT1_1_0 -> gopurs_runtime.Value
monadExceptT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_bindExceptT(dictMonad_0)
}))
_ = monadExceptT1_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_2_2
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_3})})
})
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_1_0
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, x_3)
}))
}

func Call_Control_Monad_Except_Trans_monadErrorExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
// TAST (Let): monadThrowExceptT1_3_2 -> gopurs_runtime.Value
monadThrowExceptT1_3_2 := Call_Control_Monad_Except_Trans_monadThrowExceptT(dictMonad_0)
_ = monadThrowExceptT1_3_2
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowExceptT1_3_2
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), v_4, gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply(k_5, (*Constructor_Data_Either_Left)(v2_6.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(v2_6.UnsafePtr).V0})})
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))
})
}))
}

func Call_Control_Monad_Except_Trans_monadSTExceptT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadExceptT1_2_1 -> gopurs_runtime.Value
monadExceptT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_bindExceptT(Monad0_1_0)
}))
_ = monadExceptT1_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Except_Trans_monadTransExceptT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_Control_Monad_Except_Trans_monoidExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): applicativeExceptT1_1_0 -> *Constructor_Control_Applicative_Applicative
applicativeExceptT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0))
_ = applicativeExceptT1_1_0
// TAST (Let): semigroupExceptT1_2_1 -> gopurs_runtime.Value
semigroupExceptT1_2_1 := Call_Control_Monad_Except_Trans_semigroupExceptT(dictMonad_0)
_ = semigroupExceptT1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupExceptT2_4_2 -> gopurs_runtime.Value
semigroupExceptT2_4_2 := gopurs_runtime.Apply(semigroupExceptT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupExceptT2_4_2
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupExceptT2_4_2
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeExceptT1_1_0.V1), gopurs_runtime.RecordGet(dictMonoid_3, "mempty")))
})
}

func Call_Control_Monad_Except_Trans_altExceptT(dictSemigroup_0_loop gopurs_runtime.Value, dictMonad_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
var dictMonad_1 gopurs_runtime.Value = dictMonad_1_loop
_ = dictMonad_1
// TAST (Let): Bind1_2_0 -> *Constructor_Control_Bind_Bind
Bind1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_0
// TAST (Let): Applicative0_3_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_1
// TAST (Let): functorExceptT1_4_2 -> gopurs_runtime.Value
functorExceptT1_4_2 := Call_Control_Monad_Except_Trans_functorExceptT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = functorExceptT1_4_2
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_4_2
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), v_5, gopurs_runtime.Func(func(rm_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (rm_7.Type == 9 && rm_7.IntVal == 2465973597) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rm_7.UnsafePtr).V0})})
goto end_branch_5
} else {

}
}
{
if (rm_7.Type == 9 && rm_7.IntVal == 3711209382) {
// TAST (Let): __local_var_8_3 -> gopurs_runtime.Value
__local_var_8_3 := (*Constructor_Data_Either_Left)(rm_7.UnsafePtr).V0
_ = __local_var_8_3
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), v1_6, gopurs_runtime.Func(func(rn_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (rn_9.Type == 9 && rn_9.IntVal == 2465973597) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rn_9.UnsafePtr).V0})})
goto end_branch_4
} else {

}
}
{
if (rn_9.Type == 9 && rn_9.IntVal == 3711209382) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), __local_var_8_3, (*Constructor_Data_Either_Left)(rn_9.UnsafePtr).V0)})})
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
})
}))
}

func Call_Control_Monad_Except_Trans_plusExceptT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): mempty_1_0 -> gopurs_runtime.Value
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
// TAST (Let): altExceptT1_2_1 -> gopurs_runtime.Value
altExceptT1_2_1 := gopurs_runtime.Apply(Get_Control_Monad_Except_Trans_altExceptT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = altExceptT1_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): altExceptT2_4_2 -> gopurs_runtime.Value
altExceptT2_4_2 := gopurs_runtime.Apply(altExceptT1_2_1, dictMonad_3)
_ = altExceptT2_4_2
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return altExceptT2_4_2
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Monad_Except_Trans_monadThrowExceptT(dictMonad_3), "throwError"), mempty_1_0))
})
}

func Call_Control_Monad_Except_Trans_alternativeExceptT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): plusExceptT1_1_0 -> gopurs_runtime.Value
plusExceptT1_1_0 := Call_Control_Monad_Except_Trans_plusExceptT(dictMonoid_0)
_ = plusExceptT1_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applicativeExceptT1_3_1 -> gopurs_runtime.Value
applicativeExceptT1_3_1 := Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_2)
_ = applicativeExceptT1_3_1
// TAST (Let): plusExceptT2_4_2 -> gopurs_runtime.Value
plusExceptT2_4_2 := gopurs_runtime.Apply(plusExceptT1_1_0, dictMonad_2)
_ = plusExceptT2_4_2
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeExceptT1_3_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return plusExceptT2_4_2
}))
})
}

func Call_Control_Monad_Except_Trans_monadPlusExceptT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): alternativeExceptT1_1_0 -> gopurs_runtime.Value
alternativeExceptT1_1_0 := Call_Control_Monad_Except_Trans_alternativeExceptT(dictMonoid_0)
_ = alternativeExceptT1_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadExceptT1_3_1 -> gopurs_runtime.Value
monadExceptT1_3_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_2)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_bindExceptT(dictMonad_2)
}))
_ = monadExceptT1_3_1
// TAST (Let): alternativeExceptT2_4_2 -> gopurs_runtime.Value
alternativeExceptT2_4_2 := gopurs_runtime.Apply(alternativeExceptT1_1_0, dictMonad_2)
_ = alternativeExceptT2_4_2
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeExceptT2_4_2
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_3_1
}))
})
}

func Call_Control_Monad_Except_Trans_mapExceptT__3966186360(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Except_Trans_mapExceptT__4285021944(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Except_Trans_mapExceptT__853646328(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Except_Trans_mapExceptT__1163275960(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Except_Trans_runExceptT__509566043(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}


